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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AttachmentObservation struct {
	AttachmentID *string `json:"attachmentId,omitempty" tf:"attachment_id,omitempty"`
}

type AttachmentParameters struct {

	// +kubebuilder:validation:Required
	DeviceIndex *int64 `json:"deviceIndex" tf:"device_index,omitempty"`

	// +crossplane:generate:reference:type=Instance
	// +kubebuilder:validation:Optional
	Instance *string `json:"instance,omitempty" tf:"instance,omitempty"`

	// +kubebuilder:validation:Optional
	InstanceRef *v1.Reference `json:"instanceRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	InstanceSelector *v1.Selector `json:"instanceSelector,omitempty" tf:"-"`
}

type NetworkInterfaceObservation_2 struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	MacAddress *string `json:"macAddress,omitempty" tf:"mac_address,omitempty"`

	OutpostArn *string `json:"outpostArn,omitempty" tf:"outpost_arn,omitempty"`

	PrivateDNSName *string `json:"privateDnsName,omitempty" tf:"private_dns_name,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type NetworkInterfaceParameters_2 struct {

	// +kubebuilder:validation:Optional
	Attachment []AttachmentParameters `json:"attachment,omitempty" tf:"attachment,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6AddressCount *int64 `json:"ipv6AddressCount,omitempty" tf:"ipv6_address_count,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6Addresses []*string `json:"ipv6Addresses,omitempty" tf:"ipv6_addresses,omitempty"`

	// +kubebuilder:validation:Optional
	InterfaceType *string `json:"interfaceType,omitempty" tf:"interface_type,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateIps []*string `json:"privateIps,omitempty" tf:"private_ips,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateIpsCount *int64 `json:"privateIpsCount,omitempty" tf:"private_ips_count,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SecurityGroupRefs []v1.Reference `json:"securityGroupRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SecurityGroupSelector *v1.Selector `json:"securityGroupSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupSelector
	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// +kubebuilder:validation:Optional
	SourceDestCheck *bool `json:"sourceDestCheck,omitempty" tf:"source_dest_check,omitempty"`

	// +crossplane:generate:reference:type=Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// NetworkInterfaceSpec defines the desired state of NetworkInterface
type NetworkInterfaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkInterfaceParameters_2 `json:"forProvider"`
}

// NetworkInterfaceStatus defines the observed state of NetworkInterface.
type NetworkInterfaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkInterfaceObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkInterface is the Schema for the NetworkInterfaces API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type NetworkInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkInterfaceSpec   `json:"spec"`
	Status            NetworkInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkInterfaceList contains a list of NetworkInterfaces
type NetworkInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkInterface `json:"items"`
}

// Repository type metadata.
var (
	NetworkInterface_Kind             = "NetworkInterface"
	NetworkInterface_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkInterface_Kind}.String()
	NetworkInterface_KindAPIVersion   = NetworkInterface_Kind + "." + CRDGroupVersion.String()
	NetworkInterface_GroupVersionKind = CRDGroupVersion.WithKind(NetworkInterface_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkInterface{}, &NetworkInterfaceList{})
}
