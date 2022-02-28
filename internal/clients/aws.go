package clients

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	xpawsclient "github.com/crossplane/provider-aws/pkg/clients"
	"github.com/crossplane/terrajet/pkg/terraform"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane-contrib/provider-jet-aws/apis/v1alpha1"
)

const (
	// AWS credentials environment variable names
	envSessionToken    = "AWS_SESSION_TOKEN"
	envAccessKeyID     = "AWS_ACCESS_KEY_ID"
	envSecretAccessKey = "AWS_SECRET_ACCESS_KEY"

	fmtEnvVar = "%s=%s"
)

// TerraformSetupBuilder returns Terraform setup with provider specific
// configuration like provider credentials used to connect to cloud APIs in the
// expected form of a Terraform provider.
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn { //nolint:gocyclo
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		if mg.GetProviderConfigReference() == nil {
			return ps, errors.New("no providerConfigRef provided")
		}
		pc := &v1alpha1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: mg.GetProviderConfigReference().Name, Namespace: mg.GetNamespace()}, pc); err != nil {
			return ps, errors.Wrap(err, "cannot get referenced Provider")
		}

		region, err := getRegion(mg)
		if err != nil {
			return ps, errors.Wrap(err, "cannot get region")
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1alpha1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, "cannot track ProviderConfig usage")
		}

		var cfg *aws.Config
		switch s := pc.Spec.Credentials.Source; s { //nolint:exhaustive
		case xpv1.CredentialsSourceInjectedIdentity:
			if cfg, err = xpawsclient.UsePodServiceAccount(ctx, []byte{}, xpawsclient.DefaultSection, region); err != nil {
				return ps, errors.Wrap(err, "failed to use pod service account")
			}
		default:
			data, err := resource.CommonCredentialExtractor(ctx, s, client, pc.Spec.Credentials.CommonCredentialSelectors)
			if err != nil {
				return ps, errors.Wrap(err, "cannot get credentials")
			}
			if cfg, err = xpawsclient.UseProviderSecret(ctx, data, xpawsclient.DefaultSection, region); err != nil {
				return ps, errors.Wrap(err, "failed to use provider secret")
			}
		}
		awsConf := xpawsclient.SetResolver(ctx, mg, cfg)
		creds, err := awsConf.Credentials.Retrieve(ctx)
		if err != nil {
			return ps, errors.Wrap(err, "failed to retrieve aws credentials from aws config")
		}

		// TODO(hasan): figure out what other values could be possible set here.
		//   e.g. what about setting an assume_role section: https://registry.terraform.io/providers/hashicorp/aws/latest/docs#argument-reference
		tfCfg := map[string]interface{}{}
		tfCfg["region"] = awsConf.Region
		if awsConf.Region == "" {
			// Some resources, like iam group, do not have a notion of region
			// hence we have no region in their schema. However, terraform still
			// attempts validating region in provider config and does not like
			// both empty string or not setting it at all. We need to skip
			// region validation in this case.
			tfCfg["skip_region_validation"] = true
		}

		ps.Configuration = tfCfg
		// set credentials environment
		ps.Env = []string{
			fmt.Sprintf(fmtEnvVar, envAccessKeyID, creds.AccessKeyID),
			fmt.Sprintf(fmtEnvVar, envSecretAccessKey, creds.SecretAccessKey),
			fmt.Sprintf(fmtEnvVar, envSessionToken, creds.SessionToken),
		}

		return ps, err
	}
}

func getRegion(obj runtime.Object) (string, error) {
	fromMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return "", errors.Wrap(err, "cannot convert to unstructured")
	}
	r, err := fieldpath.Pave(fromMap).GetString("spec.forProvider.region")
	if fieldpath.IsNotFound(err) {
		// Region is not required for all resources, e.g. resource in "iam"
		// group.
		return "", nil
	}
	return r, err
}
