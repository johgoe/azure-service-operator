// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=sql.azure.com,resources=serversauditingsettings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=sql.azure.com,resources={serversauditingsettings/status,serversauditingsettings/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20211101.ServersAuditingSetting
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/BlobAuditing.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/auditingSettings/default
type ServersAuditingSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Servers_AuditingSetting_Spec   `json:"spec,omitempty"`
	Status            Servers_AuditingSetting_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ServersAuditingSetting{}

// GetConditions returns the conditions of the resource
func (setting *ServersAuditingSetting) GetConditions() conditions.Conditions {
	return setting.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (setting *ServersAuditingSetting) SetConditions(conditions conditions.Conditions) {
	setting.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &ServersAuditingSetting{}

// AzureName returns the Azure name of the resource (always "default")
func (setting *ServersAuditingSetting) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (setting ServersAuditingSetting) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (setting *ServersAuditingSetting) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (setting *ServersAuditingSetting) GetSpec() genruntime.ConvertibleSpec {
	return &setting.Spec
}

// GetStatus returns the status of this resource
func (setting *ServersAuditingSetting) GetStatus() genruntime.ConvertibleStatus {
	return &setting.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/auditingSettings"
func (setting *ServersAuditingSetting) GetType() string {
	return "Microsoft.Sql/servers/auditingSettings"
}

// NewEmptyStatus returns a new empty (blank) status
func (setting *ServersAuditingSetting) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Servers_AuditingSetting_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (setting *ServersAuditingSetting) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(setting.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  setting.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (setting *ServersAuditingSetting) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Servers_AuditingSetting_STATUS); ok {
		setting.Status = *st
		return nil
	}

	// Convert status to required version
	var st Servers_AuditingSetting_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	setting.Status = st
	return nil
}

// Hub marks that this ServersAuditingSetting is the hub type for conversion
func (setting *ServersAuditingSetting) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (setting *ServersAuditingSetting) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: setting.Spec.OriginalVersion,
		Kind:    "ServersAuditingSetting",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20211101.ServersAuditingSetting
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/BlobAuditing.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/auditingSettings/default
type ServersAuditingSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServersAuditingSetting `json:"items"`
}

// Storage version of v1api20211101.Servers_AuditingSetting_Spec
type Servers_AuditingSetting_Spec struct {
	AuditActionsAndGroups       []string `json:"auditActionsAndGroups,omitempty"`
	IsAzureMonitorTargetEnabled *bool    `json:"isAzureMonitorTargetEnabled,omitempty"`
	IsDevopsAuditEnabled        *bool    `json:"isDevopsAuditEnabled,omitempty"`
	IsManagedIdentityInUse      *bool    `json:"isManagedIdentityInUse,omitempty"`
	IsStorageSecondaryKeyInUse  *bool    `json:"isStorageSecondaryKeyInUse,omitempty"`
	OriginalVersion             string   `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a sql.azure.com/Server resource
	Owner                        *genruntime.KnownResourceReference `group:"sql.azure.com" json:"owner,omitempty" kind:"Server"`
	PropertyBag                  genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	QueueDelayMs                 *int                               `json:"queueDelayMs,omitempty"`
	RetentionDays                *int                               `json:"retentionDays,omitempty"`
	State                        *string                            `json:"state,omitempty"`
	StorageAccountAccessKey      *genruntime.SecretReference        `json:"storageAccountAccessKey,omitempty"`
	StorageAccountSubscriptionId *string                            `json:"storageAccountSubscriptionId,omitempty"`
	StorageEndpoint              *string                            `json:"storageEndpoint,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Servers_AuditingSetting_Spec{}

// ConvertSpecFrom populates our Servers_AuditingSetting_Spec from the provided source
func (setting *Servers_AuditingSetting_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == setting {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(setting)
}

// ConvertSpecTo populates the provided destination from our Servers_AuditingSetting_Spec
func (setting *Servers_AuditingSetting_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == setting {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(setting)
}

// Storage version of v1api20211101.Servers_AuditingSetting_STATUS
type Servers_AuditingSetting_STATUS struct {
	AuditActionsAndGroups        []string               `json:"auditActionsAndGroups,omitempty"`
	Conditions                   []conditions.Condition `json:"conditions,omitempty"`
	Id                           *string                `json:"id,omitempty"`
	IsAzureMonitorTargetEnabled  *bool                  `json:"isAzureMonitorTargetEnabled,omitempty"`
	IsDevopsAuditEnabled         *bool                  `json:"isDevopsAuditEnabled,omitempty"`
	IsManagedIdentityInUse       *bool                  `json:"isManagedIdentityInUse,omitempty"`
	IsStorageSecondaryKeyInUse   *bool                  `json:"isStorageSecondaryKeyInUse,omitempty"`
	Name                         *string                `json:"name,omitempty"`
	PropertyBag                  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	QueueDelayMs                 *int                   `json:"queueDelayMs,omitempty"`
	RetentionDays                *int                   `json:"retentionDays,omitempty"`
	State                        *string                `json:"state,omitempty"`
	StorageAccountSubscriptionId *string                `json:"storageAccountSubscriptionId,omitempty"`
	StorageEndpoint              *string                `json:"storageEndpoint,omitempty"`
	Type                         *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Servers_AuditingSetting_STATUS{}

// ConvertStatusFrom populates our Servers_AuditingSetting_STATUS from the provided source
func (setting *Servers_AuditingSetting_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == setting {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(setting)
}

// ConvertStatusTo populates the provided destination from our Servers_AuditingSetting_STATUS
func (setting *Servers_AuditingSetting_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == setting {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(setting)
}

func init() {
	SchemeBuilder.Register(&ServersAuditingSetting{}, &ServersAuditingSettingList{})
}
