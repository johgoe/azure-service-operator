// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20201101storage

import (
	v1api20220701s "github.com/Azure/azure-service-operator/v2/api/network/v1api20220701storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=network.azure.com,resources=loadbalancers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={loadbalancers/status,loadbalancers/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20201101.LoadBalancer
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2020-11-01/loadBalancer.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}
type LoadBalancer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadBalancer_Spec   `json:"spec,omitempty"`
	Status            LoadBalancer_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &LoadBalancer{}

// GetConditions returns the conditions of the resource
func (balancer *LoadBalancer) GetConditions() conditions.Conditions {
	return balancer.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (balancer *LoadBalancer) SetConditions(conditions conditions.Conditions) {
	balancer.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &LoadBalancer{}

// AzureName returns the Azure name of the resource
func (balancer *LoadBalancer) AzureName() string {
	return balancer.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (balancer LoadBalancer) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (balancer *LoadBalancer) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (balancer *LoadBalancer) GetSpec() genruntime.ConvertibleSpec {
	return &balancer.Spec
}

// GetStatus returns the status of this resource
func (balancer *LoadBalancer) GetStatus() genruntime.ConvertibleStatus {
	return &balancer.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/loadBalancers"
func (balancer *LoadBalancer) GetType() string {
	return "Microsoft.Network/loadBalancers"
}

// NewEmptyStatus returns a new empty (blank) status
func (balancer *LoadBalancer) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &LoadBalancer_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (balancer *LoadBalancer) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(balancer.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  balancer.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (balancer *LoadBalancer) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*LoadBalancer_STATUS); ok {
		balancer.Status = *st
		return nil
	}

	// Convert status to required version
	var st LoadBalancer_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	balancer.Status = st
	return nil
}

// Hub marks that this LoadBalancer is the hub type for conversion
func (balancer *LoadBalancer) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (balancer *LoadBalancer) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: balancer.Spec.OriginalVersion,
		Kind:    "LoadBalancer",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20201101.LoadBalancer
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2020-11-01/loadBalancer.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}
type LoadBalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancer `json:"items"`
}

// Storage version of v1api20201101.APIVersion
// +kubebuilder:validation:Enum={"2020-11-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2020-11-01")

// Storage version of v1api20201101.LoadBalancer_Spec
type LoadBalancer_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                string                                                     `json:"azureName,omitempty"`
	BackendAddressPools      []BackendAddressPool_LoadBalancer_SubResourceEmbedded      `json:"backendAddressPools,omitempty"`
	ExtendedLocation         *ExtendedLocation                                          `json:"extendedLocation,omitempty"`
	FrontendIPConfigurations []FrontendIPConfiguration_LoadBalancer_SubResourceEmbedded `json:"frontendIPConfigurations,omitempty"`
	InboundNatPools          []InboundNatPool                                           `json:"inboundNatPools,omitempty"`
	InboundNatRules          []InboundNatRule_LoadBalancer_SubResourceEmbedded          `json:"inboundNatRules,omitempty"`
	LoadBalancingRules       []LoadBalancingRule                                        `json:"loadBalancingRules,omitempty"`
	Location                 *string                                                    `json:"location,omitempty"`
	OriginalVersion          string                                                     `json:"originalVersion,omitempty"`
	OutboundRules            []OutboundRule                                             `json:"outboundRules,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner       *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	Probes      []Probe                            `json:"probes,omitempty"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Sku         *LoadBalancerSku                   `json:"sku,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &LoadBalancer_Spec{}

// ConvertSpecFrom populates our LoadBalancer_Spec from the provided source
func (balancer *LoadBalancer_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == balancer {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(balancer)
}

// ConvertSpecTo populates the provided destination from our LoadBalancer_Spec
func (balancer *LoadBalancer_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == balancer {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(balancer)
}

// Storage version of v1api20201101.LoadBalancer_STATUS
// LoadBalancer resource.
type LoadBalancer_STATUS struct {
	BackendAddressPools      []BackendAddressPool_STATUS_LoadBalancer_SubResourceEmbedded      `json:"backendAddressPools,omitempty"`
	Conditions               []conditions.Condition                                            `json:"conditions,omitempty"`
	Etag                     *string                                                           `json:"etag,omitempty"`
	ExtendedLocation         *ExtendedLocation_STATUS                                          `json:"extendedLocation,omitempty"`
	FrontendIPConfigurations []FrontendIPConfiguration_STATUS_LoadBalancer_SubResourceEmbedded `json:"frontendIPConfigurations,omitempty"`
	Id                       *string                                                           `json:"id,omitempty"`
	InboundNatPools          []InboundNatPool_STATUS                                           `json:"inboundNatPools,omitempty"`
	InboundNatRules          []InboundNatRule_STATUS_LoadBalancer_SubResourceEmbedded          `json:"inboundNatRules,omitempty"`
	LoadBalancingRules       []LoadBalancingRule_STATUS                                        `json:"loadBalancingRules,omitempty"`
	Location                 *string                                                           `json:"location,omitempty"`
	Name                     *string                                                           `json:"name,omitempty"`
	OutboundRules            []OutboundRule_STATUS                                             `json:"outboundRules,omitempty"`
	Probes                   []Probe_STATUS                                                    `json:"probes,omitempty"`
	PropertyBag              genruntime.PropertyBag                                            `json:"$propertyBag,omitempty"`
	ProvisioningState        *string                                                           `json:"provisioningState,omitempty"`
	ResourceGuid             *string                                                           `json:"resourceGuid,omitempty"`
	Sku                      *LoadBalancerSku_STATUS                                           `json:"sku,omitempty"`
	Tags                     map[string]string                                                 `json:"tags,omitempty"`
	Type                     *string                                                           `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &LoadBalancer_STATUS{}

// ConvertStatusFrom populates our LoadBalancer_STATUS from the provided source
func (balancer *LoadBalancer_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == balancer {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(balancer)
}

// ConvertStatusTo populates the provided destination from our LoadBalancer_STATUS
func (balancer *LoadBalancer_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == balancer {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(balancer)
}

// Storage version of v1api20201101.BackendAddressPool_LoadBalancer_SubResourceEmbedded
// Pool of backend IP addresses.
type BackendAddressPool_LoadBalancer_SubResourceEmbedded struct {
	LoadBalancerBackendAddresses []LoadBalancerBackendAddress `json:"loadBalancerBackendAddresses,omitempty"`
	Name                         *string                      `json:"name,omitempty"`
	PropertyBag                  genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20201101.BackendAddressPool_STATUS_LoadBalancer_SubResourceEmbedded
// Pool of backend IP addresses.
type BackendAddressPool_STATUS_LoadBalancer_SubResourceEmbedded struct {
	BackendIPConfigurations      []NetworkInterfaceIPConfiguration_STATUS_LoadBalancer_SubResourceEmbedded `json:"backendIPConfigurations,omitempty"`
	Etag                         *string                                                                   `json:"etag,omitempty"`
	Id                           *string                                                                   `json:"id,omitempty"`
	LoadBalancerBackendAddresses []LoadBalancerBackendAddress_STATUS                                       `json:"loadBalancerBackendAddresses,omitempty"`
	LoadBalancingRules           []SubResource_STATUS                                                      `json:"loadBalancingRules,omitempty"`
	Name                         *string                                                                   `json:"name,omitempty"`
	OutboundRule                 *SubResource_STATUS                                                       `json:"outboundRule,omitempty"`
	OutboundRules                []SubResource_STATUS                                                      `json:"outboundRules,omitempty"`
	PropertyBag                  genruntime.PropertyBag                                                    `json:"$propertyBag,omitempty"`
	ProvisioningState            *string                                                                   `json:"provisioningState,omitempty"`
	Type                         *string                                                                   `json:"type,omitempty"`
}

// Storage version of v1api20201101.ExtendedLocation
// ExtendedLocation complex type.
type ExtendedLocation struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// AssignProperties_From_ExtendedLocation populates our ExtendedLocation from the provided source ExtendedLocation
func (location *ExtendedLocation) AssignProperties_From_ExtendedLocation(source *v1api20220701s.ExtendedLocation) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Name
	location.Name = genruntime.ClonePointerToString(source.Name)

	// Type
	location.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		location.PropertyBag = propertyBag
	} else {
		location.PropertyBag = nil
	}

	// Invoke the augmentConversionForExtendedLocation interface (if implemented) to customize the conversion
	var locationAsAny any = location
	if augmentedLocation, ok := locationAsAny.(augmentConversionForExtendedLocation); ok {
		err := augmentedLocation.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_ExtendedLocation populates the provided destination ExtendedLocation from our ExtendedLocation
func (location *ExtendedLocation) AssignProperties_To_ExtendedLocation(destination *v1api20220701s.ExtendedLocation) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(location.PropertyBag)

	// Name
	destination.Name = genruntime.ClonePointerToString(location.Name)

	// Type
	destination.Type = genruntime.ClonePointerToString(location.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForExtendedLocation interface (if implemented) to customize the conversion
	var locationAsAny any = location
	if augmentedLocation, ok := locationAsAny.(augmentConversionForExtendedLocation); ok {
		err := augmentedLocation.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20201101.ExtendedLocation_STATUS
// ExtendedLocation complex type.
type ExtendedLocation_STATUS struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

// AssignProperties_From_ExtendedLocation_STATUS populates our ExtendedLocation_STATUS from the provided source ExtendedLocation_STATUS
func (location *ExtendedLocation_STATUS) AssignProperties_From_ExtendedLocation_STATUS(source *v1api20220701s.ExtendedLocation_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Name
	location.Name = genruntime.ClonePointerToString(source.Name)

	// Type
	location.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		location.PropertyBag = propertyBag
	} else {
		location.PropertyBag = nil
	}

	// Invoke the augmentConversionForExtendedLocation_STATUS interface (if implemented) to customize the conversion
	var locationAsAny any = location
	if augmentedLocation, ok := locationAsAny.(augmentConversionForExtendedLocation_STATUS); ok {
		err := augmentedLocation.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_ExtendedLocation_STATUS populates the provided destination ExtendedLocation_STATUS from our ExtendedLocation_STATUS
func (location *ExtendedLocation_STATUS) AssignProperties_To_ExtendedLocation_STATUS(destination *v1api20220701s.ExtendedLocation_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(location.PropertyBag)

	// Name
	destination.Name = genruntime.ClonePointerToString(location.Name)

	// Type
	destination.Type = genruntime.ClonePointerToString(location.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForExtendedLocation_STATUS interface (if implemented) to customize the conversion
	var locationAsAny any = location
	if augmentedLocation, ok := locationAsAny.(augmentConversionForExtendedLocation_STATUS); ok {
		err := augmentedLocation.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20201101.FrontendIPConfiguration_LoadBalancer_SubResourceEmbedded
// Frontend IP address of the load balancer.
type FrontendIPConfiguration_LoadBalancer_SubResourceEmbedded struct {
	Name                      *string                                               `json:"name,omitempty"`
	PrivateIPAddress          *string                                               `json:"privateIPAddress,omitempty"`
	PrivateIPAddressVersion   *string                                               `json:"privateIPAddressVersion,omitempty"`
	PrivateIPAllocationMethod *string                                               `json:"privateIPAllocationMethod,omitempty"`
	PropertyBag               genruntime.PropertyBag                                `json:"$propertyBag,omitempty"`
	PublicIPAddress           *PublicIPAddressSpec_LoadBalancer_SubResourceEmbedded `json:"publicIPAddress,omitempty"`
	PublicIPPrefix            *SubResource                                          `json:"publicIPPrefix,omitempty"`
	Subnet                    *Subnet_LoadBalancer_SubResourceEmbedded              `json:"subnet,omitempty"`
	Zones                     []string                                              `json:"zones,omitempty"`
}

// Storage version of v1api20201101.FrontendIPConfiguration_STATUS_LoadBalancer_SubResourceEmbedded
// Frontend IP address of the load balancer.
type FrontendIPConfiguration_STATUS_LoadBalancer_SubResourceEmbedded struct {
	Etag                      *string                                                  `json:"etag,omitempty"`
	Id                        *string                                                  `json:"id,omitempty"`
	InboundNatPools           []SubResource_STATUS                                     `json:"inboundNatPools,omitempty"`
	InboundNatRules           []SubResource_STATUS                                     `json:"inboundNatRules,omitempty"`
	LoadBalancingRules        []SubResource_STATUS                                     `json:"loadBalancingRules,omitempty"`
	Name                      *string                                                  `json:"name,omitempty"`
	OutboundRules             []SubResource_STATUS                                     `json:"outboundRules,omitempty"`
	PrivateIPAddress          *string                                                  `json:"privateIPAddress,omitempty"`
	PrivateIPAddressVersion   *string                                                  `json:"privateIPAddressVersion,omitempty"`
	PrivateIPAllocationMethod *string                                                  `json:"privateIPAllocationMethod,omitempty"`
	PropertyBag               genruntime.PropertyBag                                   `json:"$propertyBag,omitempty"`
	ProvisioningState         *string                                                  `json:"provisioningState,omitempty"`
	PublicIPAddress           *PublicIPAddress_STATUS_LoadBalancer_SubResourceEmbedded `json:"publicIPAddress,omitempty"`
	PublicIPPrefix            *SubResource_STATUS                                      `json:"publicIPPrefix,omitempty"`
	Subnet                    *Subnet_STATUS_LoadBalancer_SubResourceEmbedded          `json:"subnet,omitempty"`
	Type                      *string                                                  `json:"type,omitempty"`
	Zones                     []string                                                 `json:"zones,omitempty"`
}

// Storage version of v1api20201101.InboundNatPool
// Inbound NAT pool of the load balancer.
type InboundNatPool struct {
	BackendPort             *int                   `json:"backendPort,omitempty"`
	EnableFloatingIP        *bool                  `json:"enableFloatingIP,omitempty"`
	EnableTcpReset          *bool                  `json:"enableTcpReset,omitempty"`
	FrontendIPConfiguration *SubResource           `json:"frontendIPConfiguration,omitempty"`
	FrontendPortRangeEnd    *int                   `json:"frontendPortRangeEnd,omitempty"`
	FrontendPortRangeStart  *int                   `json:"frontendPortRangeStart,omitempty"`
	IdleTimeoutInMinutes    *int                   `json:"idleTimeoutInMinutes,omitempty"`
	Name                    *string                `json:"name,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol                *string                `json:"protocol,omitempty"`
}

// Storage version of v1api20201101.InboundNatPool_STATUS
// Inbound NAT pool of the load balancer.
type InboundNatPool_STATUS struct {
	BackendPort             *int                   `json:"backendPort,omitempty"`
	EnableFloatingIP        *bool                  `json:"enableFloatingIP,omitempty"`
	EnableTcpReset          *bool                  `json:"enableTcpReset,omitempty"`
	Etag                    *string                `json:"etag,omitempty"`
	FrontendIPConfiguration *SubResource_STATUS    `json:"frontendIPConfiguration,omitempty"`
	FrontendPortRangeEnd    *int                   `json:"frontendPortRangeEnd,omitempty"`
	FrontendPortRangeStart  *int                   `json:"frontendPortRangeStart,omitempty"`
	Id                      *string                `json:"id,omitempty"`
	IdleTimeoutInMinutes    *int                   `json:"idleTimeoutInMinutes,omitempty"`
	Name                    *string                `json:"name,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol                *string                `json:"protocol,omitempty"`
	ProvisioningState       *string                `json:"provisioningState,omitempty"`
	Type                    *string                `json:"type,omitempty"`
}

// Storage version of v1api20201101.InboundNatRule_LoadBalancer_SubResourceEmbedded
// Inbound NAT rule of the load balancer.
type InboundNatRule_LoadBalancer_SubResourceEmbedded struct {
	BackendPort             *int                   `json:"backendPort,omitempty"`
	EnableFloatingIP        *bool                  `json:"enableFloatingIP,omitempty"`
	EnableTcpReset          *bool                  `json:"enableTcpReset,omitempty"`
	FrontendIPConfiguration *SubResource           `json:"frontendIPConfiguration,omitempty"`
	FrontendPort            *int                   `json:"frontendPort,omitempty"`
	IdleTimeoutInMinutes    *int                   `json:"idleTimeoutInMinutes,omitempty"`
	Name                    *string                `json:"name,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol                *string                `json:"protocol,omitempty"`
}

// Storage version of v1api20201101.InboundNatRule_STATUS_LoadBalancer_SubResourceEmbedded
// Inbound NAT rule of the load balancer.
type InboundNatRule_STATUS_LoadBalancer_SubResourceEmbedded struct {
	BackendIPConfiguration  *NetworkInterfaceIPConfiguration_STATUS_LoadBalancer_SubResourceEmbedded `json:"backendIPConfiguration,omitempty"`
	BackendPort             *int                                                                     `json:"backendPort,omitempty"`
	EnableFloatingIP        *bool                                                                    `json:"enableFloatingIP,omitempty"`
	EnableTcpReset          *bool                                                                    `json:"enableTcpReset,omitempty"`
	Etag                    *string                                                                  `json:"etag,omitempty"`
	FrontendIPConfiguration *SubResource_STATUS                                                      `json:"frontendIPConfiguration,omitempty"`
	FrontendPort            *int                                                                     `json:"frontendPort,omitempty"`
	Id                      *string                                                                  `json:"id,omitempty"`
	IdleTimeoutInMinutes    *int                                                                     `json:"idleTimeoutInMinutes,omitempty"`
	Name                    *string                                                                  `json:"name,omitempty"`
	PropertyBag             genruntime.PropertyBag                                                   `json:"$propertyBag,omitempty"`
	Protocol                *string                                                                  `json:"protocol,omitempty"`
	ProvisioningState       *string                                                                  `json:"provisioningState,omitempty"`
	Type                    *string                                                                  `json:"type,omitempty"`
}

// Storage version of v1api20201101.LoadBalancerSku
// SKU of a load balancer.
type LoadBalancerSku struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

// Storage version of v1api20201101.LoadBalancerSku_STATUS
// SKU of a load balancer.
type LoadBalancerSku_STATUS struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

// Storage version of v1api20201101.LoadBalancingRule
// A load balancing rule for a load balancer.
type LoadBalancingRule struct {
	BackendAddressPool      *SubResource           `json:"backendAddressPool,omitempty"`
	BackendPort             *int                   `json:"backendPort,omitempty"`
	DisableOutboundSnat     *bool                  `json:"disableOutboundSnat,omitempty"`
	EnableFloatingIP        *bool                  `json:"enableFloatingIP,omitempty"`
	EnableTcpReset          *bool                  `json:"enableTcpReset,omitempty"`
	FrontendIPConfiguration *SubResource           `json:"frontendIPConfiguration,omitempty"`
	FrontendPort            *int                   `json:"frontendPort,omitempty"`
	IdleTimeoutInMinutes    *int                   `json:"idleTimeoutInMinutes,omitempty"`
	LoadDistribution        *string                `json:"loadDistribution,omitempty"`
	Name                    *string                `json:"name,omitempty"`
	Probe                   *SubResource           `json:"probe,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol                *string                `json:"protocol,omitempty"`
}

// Storage version of v1api20201101.LoadBalancingRule_STATUS
// A load balancing rule for a load balancer.
type LoadBalancingRule_STATUS struct {
	BackendAddressPool      *SubResource_STATUS    `json:"backendAddressPool,omitempty"`
	BackendPort             *int                   `json:"backendPort,omitempty"`
	DisableOutboundSnat     *bool                  `json:"disableOutboundSnat,omitempty"`
	EnableFloatingIP        *bool                  `json:"enableFloatingIP,omitempty"`
	EnableTcpReset          *bool                  `json:"enableTcpReset,omitempty"`
	Etag                    *string                `json:"etag,omitempty"`
	FrontendIPConfiguration *SubResource_STATUS    `json:"frontendIPConfiguration,omitempty"`
	FrontendPort            *int                   `json:"frontendPort,omitempty"`
	Id                      *string                `json:"id,omitempty"`
	IdleTimeoutInMinutes    *int                   `json:"idleTimeoutInMinutes,omitempty"`
	LoadDistribution        *string                `json:"loadDistribution,omitempty"`
	Name                    *string                `json:"name,omitempty"`
	Probe                   *SubResource_STATUS    `json:"probe,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol                *string                `json:"protocol,omitempty"`
	ProvisioningState       *string                `json:"provisioningState,omitempty"`
	Type                    *string                `json:"type,omitempty"`
}

// Storage version of v1api20201101.OutboundRule
// Outbound rule of the load balancer.
type OutboundRule struct {
	AllocatedOutboundPorts   *int                   `json:"allocatedOutboundPorts,omitempty"`
	BackendAddressPool       *SubResource           `json:"backendAddressPool,omitempty"`
	EnableTcpReset           *bool                  `json:"enableTcpReset,omitempty"`
	FrontendIPConfigurations []SubResource          `json:"frontendIPConfigurations,omitempty"`
	IdleTimeoutInMinutes     *int                   `json:"idleTimeoutInMinutes,omitempty"`
	Name                     *string                `json:"name,omitempty"`
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol                 *string                `json:"protocol,omitempty"`
}

// Storage version of v1api20201101.OutboundRule_STATUS
// Outbound rule of the load balancer.
type OutboundRule_STATUS struct {
	AllocatedOutboundPorts   *int                   `json:"allocatedOutboundPorts,omitempty"`
	BackendAddressPool       *SubResource_STATUS    `json:"backendAddressPool,omitempty"`
	EnableTcpReset           *bool                  `json:"enableTcpReset,omitempty"`
	Etag                     *string                `json:"etag,omitempty"`
	FrontendIPConfigurations []SubResource_STATUS   `json:"frontendIPConfigurations,omitempty"`
	Id                       *string                `json:"id,omitempty"`
	IdleTimeoutInMinutes     *int                   `json:"idleTimeoutInMinutes,omitempty"`
	Name                     *string                `json:"name,omitempty"`
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol                 *string                `json:"protocol,omitempty"`
	ProvisioningState        *string                `json:"provisioningState,omitempty"`
	Type                     *string                `json:"type,omitempty"`
}

// Storage version of v1api20201101.Probe
// A load balancer probe.
type Probe struct {
	IntervalInSeconds *int                   `json:"intervalInSeconds,omitempty"`
	Name              *string                `json:"name,omitempty"`
	NumberOfProbes    *int                   `json:"numberOfProbes,omitempty"`
	Port              *int                   `json:"port,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol          *string                `json:"protocol,omitempty"`
	RequestPath       *string                `json:"requestPath,omitempty"`
}

// Storage version of v1api20201101.Probe_STATUS
// A load balancer probe.
type Probe_STATUS struct {
	Etag               *string                `json:"etag,omitempty"`
	Id                 *string                `json:"id,omitempty"`
	IntervalInSeconds  *int                   `json:"intervalInSeconds,omitempty"`
	LoadBalancingRules []SubResource_STATUS   `json:"loadBalancingRules,omitempty"`
	Name               *string                `json:"name,omitempty"`
	NumberOfProbes     *int                   `json:"numberOfProbes,omitempty"`
	Port               *int                   `json:"port,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Protocol           *string                `json:"protocol,omitempty"`
	ProvisioningState  *string                `json:"provisioningState,omitempty"`
	RequestPath        *string                `json:"requestPath,omitempty"`
	Type               *string                `json:"type,omitempty"`
}

type augmentConversionForExtendedLocation interface {
	AssignPropertiesFrom(src *v1api20220701s.ExtendedLocation) error
	AssignPropertiesTo(dst *v1api20220701s.ExtendedLocation) error
}

type augmentConversionForExtendedLocation_STATUS interface {
	AssignPropertiesFrom(src *v1api20220701s.ExtendedLocation_STATUS) error
	AssignPropertiesTo(dst *v1api20220701s.ExtendedLocation_STATUS) error
}

// Storage version of v1api20201101.LoadBalancerBackendAddress
// Load balancer backend addresses.
type LoadBalancerBackendAddress struct {
	IpAddress                           *string                `json:"ipAddress,omitempty"`
	LoadBalancerFrontendIPConfiguration *SubResource           `json:"loadBalancerFrontendIPConfiguration,omitempty"`
	Name                                *string                `json:"name,omitempty"`
	PropertyBag                         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Subnet                              *SubResource           `json:"subnet,omitempty"`
	VirtualNetwork                      *SubResource           `json:"virtualNetwork,omitempty"`
}

// Storage version of v1api20201101.LoadBalancerBackendAddress_STATUS
// Load balancer backend addresses.
type LoadBalancerBackendAddress_STATUS struct {
	IpAddress                           *string                `json:"ipAddress,omitempty"`
	LoadBalancerFrontendIPConfiguration *SubResource_STATUS    `json:"loadBalancerFrontendIPConfiguration,omitempty"`
	Name                                *string                `json:"name,omitempty"`
	NetworkInterfaceIPConfiguration     *SubResource_STATUS    `json:"networkInterfaceIPConfiguration,omitempty"`
	PropertyBag                         genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Subnet                              *SubResource_STATUS    `json:"subnet,omitempty"`
	VirtualNetwork                      *SubResource_STATUS    `json:"virtualNetwork,omitempty"`
}

// Storage version of v1api20201101.NetworkInterfaceIPConfiguration_STATUS_LoadBalancer_SubResourceEmbedded
// IPConfiguration in a network interface.
type NetworkInterfaceIPConfiguration_STATUS_LoadBalancer_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20201101.PublicIPAddress_STATUS_LoadBalancer_SubResourceEmbedded
// Public IP address resource.
type PublicIPAddress_STATUS_LoadBalancer_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20201101.PublicIPAddressSpec_LoadBalancer_SubResourceEmbedded
// Public IP address resource.
type PublicIPAddressSpec_LoadBalancer_SubResourceEmbedded struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// Reference: Resource ID.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// Storage version of v1api20201101.Subnet_LoadBalancer_SubResourceEmbedded
// Subnet in a virtual network resource.
type Subnet_LoadBalancer_SubResourceEmbedded struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// Reference: Resource ID.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// Storage version of v1api20201101.Subnet_STATUS_LoadBalancer_SubResourceEmbedded
// Subnet in a virtual network resource.
type Subnet_STATUS_LoadBalancer_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&LoadBalancer{}, &LoadBalancerList{})
}
