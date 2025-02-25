// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210701

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Workspace_Spec_ARM struct {
	// Identity: The identity of the resource.
	Identity *Identity_ARM `json:"identity,omitempty"`

	// Location: Specifies the location of the resource.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: The properties of the machine learning workspace.
	Properties *WorkspaceProperties_ARM `json:"properties,omitempty"`

	// Sku: The sku of the workspace.
	Sku *Sku_ARM `json:"sku,omitempty"`

	// SystemData: System data
	SystemData *SystemData_ARM `json:"systemData,omitempty"`

	// Tags: Contains resource tags defined as key/value pairs.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Workspace_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-07-01"
func (workspace Workspace_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (workspace *Workspace_Spec_ARM) GetName() string {
	return workspace.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.MachineLearningServices/workspaces"
func (workspace *Workspace_Spec_ARM) GetType() string {
	return "Microsoft.MachineLearningServices/workspaces"
}

// Identity for the resource.
type Identity_ARM struct {
	// Type: The identity type.
	Type *Identity_Type `json:"type,omitempty"`
}

// Sku of the resource
type Sku_ARM struct {
	// Name: Name of the sku
	Name *string `json:"name,omitempty"`

	// Tier: Tier of the sku like Basic or Enterprise
	Tier *string `json:"tier,omitempty"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemData_ARM struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemData_CreatedByType `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource last modification (UTC)
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemData_LastModifiedByType `json:"lastModifiedByType,omitempty"`
}

// The properties of a machine learning workspace.
type WorkspaceProperties_ARM struct {
	// AllowPublicAccessWhenBehindVnet: The flag to indicate whether to allow public access when behind VNet.
	AllowPublicAccessWhenBehindVnet *bool   `json:"allowPublicAccessWhenBehindVnet,omitempty"`
	ApplicationInsights             *string `json:"applicationInsights,omitempty"`
	ContainerRegistry               *string `json:"containerRegistry,omitempty"`

	// Description: The description of this workspace.
	Description *string `json:"description,omitempty"`

	// DiscoveryUrl: Url for the discovery service to identify regional endpoints for machine learning experimentation services
	DiscoveryUrl *string `json:"discoveryUrl,omitempty"`

	// Encryption: The encryption settings of Azure ML workspace.
	Encryption *EncryptionProperty_ARM `json:"encryption,omitempty"`

	// FriendlyName: The friendly name for this workspace. This name in mutable
	FriendlyName *string `json:"friendlyName,omitempty"`

	// HbiWorkspace: The flag to signal HBI data in the workspace and reduce diagnostic data collected by the service
	HbiWorkspace *bool `json:"hbiWorkspace,omitempty"`

	// ImageBuildCompute: The compute name for image build
	ImageBuildCompute           *string `json:"imageBuildCompute,omitempty"`
	KeyVault                    *string `json:"keyVault,omitempty"`
	PrimaryUserAssignedIdentity *string `json:"primaryUserAssignedIdentity,omitempty"`

	// PublicNetworkAccess: Whether requests from Public Network are allowed.
	PublicNetworkAccess *WorkspaceProperties_PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// ServiceManagedResourcesSettings: The service managed resource settings.
	ServiceManagedResourcesSettings *ServiceManagedResourcesSettings_ARM `json:"serviceManagedResourcesSettings,omitempty"`

	// SharedPrivateLinkResources: The list of shared private link resources in this workspace.
	SharedPrivateLinkResources []SharedPrivateLinkResource_ARM `json:"sharedPrivateLinkResources,omitempty"`
	StorageAccount             *string                         `json:"storageAccount,omitempty"`
}

type EncryptionProperty_ARM struct {
	// Identity: The identity that will be used to access the key vault for encryption at rest.
	Identity *IdentityForCmk_ARM `json:"identity,omitempty"`

	// KeyVaultProperties: Customer Key vault properties.
	KeyVaultProperties *KeyVaultProperties_ARM `json:"keyVaultProperties,omitempty"`

	// Status: Indicates whether or not the encryption is enabled for the workspace.
	Status *EncryptionProperty_Status `json:"status,omitempty"`
}

// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned,UserAssigned","UserAssigned"}
type Identity_Type string

const (
	Identity_Type_None                       = Identity_Type("None")
	Identity_Type_SystemAssigned             = Identity_Type("SystemAssigned")
	Identity_Type_SystemAssignedUserAssigned = Identity_Type("SystemAssigned,UserAssigned")
	Identity_Type_UserAssigned               = Identity_Type("UserAssigned")
)

type ServiceManagedResourcesSettings_ARM struct {
	// CosmosDb: The settings for the service managed cosmosdb account.
	CosmosDb *CosmosDbSettings_ARM `json:"cosmosDb,omitempty"`
}

type SharedPrivateLinkResource_ARM struct {
	// Name: Unique name of the private link.
	Name *string `json:"name,omitempty"`

	// Properties: Resource properties.
	Properties *SharedPrivateLinkResourceProperty_ARM `json:"properties,omitempty"`
}

// +kubebuilder:validation:Enum={"Application","Key","ManagedIdentity","User"}
type SystemData_CreatedByType string

const (
	SystemData_CreatedByType_Application     = SystemData_CreatedByType("Application")
	SystemData_CreatedByType_Key             = SystemData_CreatedByType("Key")
	SystemData_CreatedByType_ManagedIdentity = SystemData_CreatedByType("ManagedIdentity")
	SystemData_CreatedByType_User            = SystemData_CreatedByType("User")
)

// +kubebuilder:validation:Enum={"Application","Key","ManagedIdentity","User"}
type SystemData_LastModifiedByType string

const (
	SystemData_LastModifiedByType_Application     = SystemData_LastModifiedByType("Application")
	SystemData_LastModifiedByType_Key             = SystemData_LastModifiedByType("Key")
	SystemData_LastModifiedByType_ManagedIdentity = SystemData_LastModifiedByType("ManagedIdentity")
	SystemData_LastModifiedByType_User            = SystemData_LastModifiedByType("User")
)

type CosmosDbSettings_ARM struct {
	// CollectionsThroughput: The throughput of the collections in cosmosdb database
	CollectionsThroughput *int `json:"collectionsThroughput,omitempty"`
}

// Identity that will be used to access key vault for encryption at rest
type IdentityForCmk_ARM struct {
	// UserAssignedIdentity: The ArmId of the user assigned identity that will be used to access the customer managed key vault
	UserAssignedIdentity *string `json:"userAssignedIdentity,omitempty"`
}

type KeyVaultProperties_ARM struct {
	// IdentityClientId: For future use - The client id of the identity which will be used to access key vault.
	IdentityClientId *string `json:"identityClientId,omitempty"`

	// KeyIdentifier: Key vault uri to access the encryption key.
	KeyIdentifier *string `json:"keyIdentifier,omitempty"`

	// KeyVaultArmId: The ArmId of the keyVault where the customer owned encryption key is present.
	KeyVaultArmId *string `json:"keyVaultArmId,omitempty"`
}

// Properties of a shared private link resource.
type SharedPrivateLinkResourceProperty_ARM struct {
	// GroupId: The private link resource group id.
	GroupId               *string `json:"groupId,omitempty"`
	PrivateLinkResourceId *string `json:"privateLinkResourceId,omitempty"`

	// RequestMessage: Request message.
	RequestMessage *string `json:"requestMessage,omitempty"`

	// Status: Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
	Status *PrivateEndpointServiceConnectionStatus `json:"status,omitempty"`
}
