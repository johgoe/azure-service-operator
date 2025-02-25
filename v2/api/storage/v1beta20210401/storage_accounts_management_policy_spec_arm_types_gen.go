// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210401

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of StorageAccounts_ManagementPolicy_Spec. Use v1api20210401.StorageAccounts_ManagementPolicy_Spec instead
type StorageAccounts_ManagementPolicy_Spec_ARM struct {
	Name       string                          `json:"name,omitempty"`
	Properties *ManagementPolicyProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &StorageAccounts_ManagementPolicy_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (policy StorageAccounts_ManagementPolicy_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (policy *StorageAccounts_ManagementPolicy_Spec_ARM) GetName() string {
	return policy.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/managementPolicies"
func (policy *StorageAccounts_ManagementPolicy_Spec_ARM) GetType() string {
	return "Microsoft.Storage/storageAccounts/managementPolicies"
}

// Deprecated version of ManagementPolicyProperties. Use v1api20210401.ManagementPolicyProperties instead
type ManagementPolicyProperties_ARM struct {
	Policy *ManagementPolicySchema_ARM `json:"policy,omitempty"`
}

// Deprecated version of ManagementPolicySchema. Use v1api20210401.ManagementPolicySchema instead
type ManagementPolicySchema_ARM struct {
	Rules []ManagementPolicyRule_ARM `json:"rules,omitempty"`
}

// Deprecated version of ManagementPolicyRule. Use v1api20210401.ManagementPolicyRule instead
type ManagementPolicyRule_ARM struct {
	Definition *ManagementPolicyDefinition_ARM `json:"definition,omitempty"`
	Enabled    *bool                           `json:"enabled,omitempty"`
	Name       *string                         `json:"name,omitempty"`
	Type       *ManagementPolicyRule_Type      `json:"type,omitempty"`
}

// Deprecated version of ManagementPolicyDefinition. Use v1api20210401.ManagementPolicyDefinition instead
type ManagementPolicyDefinition_ARM struct {
	Actions *ManagementPolicyAction_ARM `json:"actions,omitempty"`
	Filters *ManagementPolicyFilter_ARM `json:"filters,omitempty"`
}

// Deprecated version of ManagementPolicyAction. Use v1api20210401.ManagementPolicyAction instead
type ManagementPolicyAction_ARM struct {
	BaseBlob *ManagementPolicyBaseBlob_ARM `json:"baseBlob,omitempty"`
	Snapshot *ManagementPolicySnapShot_ARM `json:"snapshot,omitempty"`
	Version  *ManagementPolicyVersion_ARM  `json:"version,omitempty"`
}

// Deprecated version of ManagementPolicyFilter. Use v1api20210401.ManagementPolicyFilter instead
type ManagementPolicyFilter_ARM struct {
	BlobIndexMatch []TagFilter_ARM `json:"blobIndexMatch,omitempty"`
	BlobTypes      []string        `json:"blobTypes,omitempty"`
	PrefixMatch    []string        `json:"prefixMatch,omitempty"`
}

// Deprecated version of ManagementPolicyBaseBlob. Use v1api20210401.ManagementPolicyBaseBlob instead
type ManagementPolicyBaseBlob_ARM struct {
	Delete                      *DateAfterModification_ARM `json:"delete,omitempty"`
	EnableAutoTierToHotFromCool *bool                      `json:"enableAutoTierToHotFromCool,omitempty"`
	TierToArchive               *DateAfterModification_ARM `json:"tierToArchive,omitempty"`
	TierToCool                  *DateAfterModification_ARM `json:"tierToCool,omitempty"`
}

// Deprecated version of ManagementPolicySnapShot. Use v1api20210401.ManagementPolicySnapShot instead
type ManagementPolicySnapShot_ARM struct {
	Delete        *DateAfterCreation_ARM `json:"delete,omitempty"`
	TierToArchive *DateAfterCreation_ARM `json:"tierToArchive,omitempty"`
	TierToCool    *DateAfterCreation_ARM `json:"tierToCool,omitempty"`
}

// Deprecated version of ManagementPolicyVersion. Use v1api20210401.ManagementPolicyVersion instead
type ManagementPolicyVersion_ARM struct {
	Delete        *DateAfterCreation_ARM `json:"delete,omitempty"`
	TierToArchive *DateAfterCreation_ARM `json:"tierToArchive,omitempty"`
	TierToCool    *DateAfterCreation_ARM `json:"tierToCool,omitempty"`
}

// Deprecated version of TagFilter. Use v1api20210401.TagFilter instead
type TagFilter_ARM struct {
	Name  *string `json:"name,omitempty"`
	Op    *string `json:"op,omitempty"`
	Value *string `json:"value,omitempty"`
}

// Deprecated version of DateAfterCreation. Use v1api20210401.DateAfterCreation instead
type DateAfterCreation_ARM struct {
	DaysAfterCreationGreaterThan *int `json:"daysAfterCreationGreaterThan,omitempty"`
}

// Deprecated version of DateAfterModification. Use v1api20210401.DateAfterModification instead
type DateAfterModification_ARM struct {
	DaysAfterLastAccessTimeGreaterThan *int `json:"daysAfterLastAccessTimeGreaterThan,omitempty"`
	DaysAfterModificationGreaterThan   *int `json:"daysAfterModificationGreaterThan,omitempty"`
}
