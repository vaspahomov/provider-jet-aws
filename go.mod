module github.com/crossplane-contrib/provider-jet-aws

go 1.16

require (
	github.com/aws/aws-sdk-go-v2 v0.23.0
	github.com/crossplane/crossplane-runtime v0.15.1-0.20211004150827-579c1833b513
	github.com/crossplane/crossplane-tools v0.0.0-20210916125540-071de511ae8e
	github.com/crossplane/provider-aws v0.19.0
	github.com/crossplane/terrajet v0.4.0-rc.0.0.20220128111246-e5aaa1790fe6
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.7.0
	github.com/pkg/errors v0.9.1
	github.com/terraform-providers/terraform-provider-aws v1.60.1-0.20210811232925-d6f99829ec3f
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	k8s.io/apimachinery v0.23.0
	k8s.io/client-go v0.23.0
	k8s.io/utils v0.0.0-20211208161948-7d6a63dca704
	sigs.k8s.io/controller-runtime v0.11.0
	sigs.k8s.io/controller-tools v0.8.0
)

replace github.com/crossplane/terrajet => github.com/vaspahomov/terrajet v0.2.1-0.20220216131301-55dd18cdd805
replace github.com/crossplane/crossplane-runtime => github.com/vaspahomov/crossplane-runtime v0.14.1-0.20220202070154-70fc88a4d127

replace github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/gdavison/terraform-plugin-sdk/v2 v2.0.2-0.20210714181518-b5a3dc95a675
