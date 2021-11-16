/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/iam/v1alpha1"
	common "github.com/crossplane-contrib/provider-tf-aws/config/common"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Cluster.
func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.CapacityProviders),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.CapacityProvidersRefs,
		Selector:      mg.Spec.ForProvider.CapacityProvidersSelector,
		To: reference.To{
			List:    &CapacityProviderList{},
			Managed: &CapacityProvider{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CapacityProviders")
	}
	mg.Spec.ForProvider.CapacityProviders = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.CapacityProvidersRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this Service.
func (mg *Service) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Cluster),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.ClusterRef,
		Selector:     mg.Spec.ForProvider.ClusterSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Cluster")
	}
	mg.Spec.ForProvider.Cluster = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IamRole),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.IamRoleRef,
		Selector:     mg.Spec.ForProvider.IamRoleSelector,
		To: reference.To{
			List:    &v1alpha1.RoleList{},
			Managed: &v1alpha1.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.IamRole")
	}
	mg.Spec.ForProvider.IamRole = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.IamRoleRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TaskDefinition.
func (mg *TaskDefinition) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ExecutionRoleArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.ExecutionRoleArnRef,
		Selector:     mg.Spec.ForProvider.ExecutionRoleArnSelector,
		To: reference.To{
			List:    &v1alpha1.RoleList{},
			Managed: &v1alpha1.Role{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ExecutionRoleArn")
	}
	mg.Spec.ForProvider.ExecutionRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ExecutionRoleArnRef = rsp.ResolvedReference

	return nil
}
