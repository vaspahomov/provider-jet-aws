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

package ecr

import (
	"github.com/crossplane-contrib/terrajet/pkg/config"

	"github.com/crossplane-contrib/provider-tf-aws/config/common"
)

// Configure adds configurations for ecrs group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_ecr_repository", func(r *config.Resource) {
		r.References = map[string]config.Reference{
			"encryption_configuration.kms_key": {
				Type:      "github.com/crossplane-contrib/provider-tf-aws/apis/kms/v1alpha1.Key",
				Extractor: common.PathARNExtractor,
			},
		}
	})
}
