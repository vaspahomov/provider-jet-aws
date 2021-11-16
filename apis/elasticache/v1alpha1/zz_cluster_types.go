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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CacheNodesObservation struct {
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`
}

type CacheNodesParameters struct {
}

type ClusterObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	CacheNodes []CacheNodesObservation `json:"cacheNodes,omitempty" tf:"cache_nodes,omitempty"`

	ClusterAddress *string `json:"clusterAddress,omitempty" tf:"cluster_address,omitempty"`

	ConfigurationEndpoint *string `json:"configurationEndpoint,omitempty" tf:"configuration_endpoint,omitempty"`

	EngineVersionActual *string `json:"engineVersionActual,omitempty" tf:"engine_version_actual,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ClusterParameters struct {

	// +kubebuilder:validation:Optional
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately,omitempty"`

	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// +kubebuilder:validation:Optional
	AzMode *string `json:"azMode,omitempty" tf:"az_mode,omitempty"`

	// +kubebuilder:validation:Optional
	Engine *string `json:"engine,omitempty" tf:"engine,omitempty"`

	// +kubebuilder:validation:Optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// +kubebuilder:validation:Optional
	FinalSnapshotIdentifier *string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier,omitempty"`

	// +kubebuilder:validation:Optional
	MaintenanceWindow *string `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// +kubebuilder:validation:Optional
	NodeType *string `json:"nodeType,omitempty" tf:"node_type,omitempty"`

	// +kubebuilder:validation:Optional
	NotificationTopicArn *string `json:"notificationTopicArn,omitempty" tf:"notification_topic_arn,omitempty"`

	// +kubebuilder:validation:Optional
	NumCacheNodes *int64 `json:"numCacheNodes,omitempty" tf:"num_cache_nodes,omitempty"`

	// +crossplane:generate:reference:type=ParameterGroup
	// +kubebuilder:validation:Optional
	ParameterGroupName *string `json:"parameterGroupName,omitempty" tf:"parameter_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ParameterGroupNameRef *v1.Reference `json:"parameterGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ParameterGroupNameSelector *v1.Selector `json:"parameterGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Port *int64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	PreferredAvailabilityZones []*string `json:"preferredAvailabilityZones,omitempty" tf:"preferred_availability_zones,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	ReplicationGroupID *string `json:"replicationGroupId,omitempty" tf:"replication_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupNames []*string `json:"securityGroupNames,omitempty" tf:"security_group_names,omitempty"`

	// +kubebuilder:validation:Optional
	SnapshotArns []*string `json:"snapshotArns,omitempty" tf:"snapshot_arns,omitempty"`

	// +kubebuilder:validation:Optional
	SnapshotName *string `json:"snapshotName,omitempty" tf:"snapshot_name,omitempty"`

	// +kubebuilder:validation:Optional
	SnapshotRetentionLimit *int64 `json:"snapshotRetentionLimit,omitempty" tf:"snapshot_retention_limit,omitempty"`

	// +kubebuilder:validation:Optional
	SnapshotWindow *string `json:"snapshotWindow,omitempty" tf:"snapshot_window,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetGroupName *string `json:"subnetGroupName,omitempty" tf:"subnet_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
