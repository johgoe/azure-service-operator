// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210401

// Deprecated version of StorageAccounts_ManagementPolicy_STATUS. Use v1api20210401.StorageAccounts_ManagementPolicy_STATUS instead
type StorageAccounts_ManagementPolicy_STATUS_ARM struct {
	Id         *string                                `json:"id,omitempty"`
	Name       *string                                `json:"name,omitempty"`
	Properties *ManagementPolicyProperties_STATUS_ARM `json:"properties,omitempty"`
	Type       *string                                `json:"type,omitempty"`
}

// Deprecated version of ManagementPolicyProperties_STATUS. Use v1api20210401.ManagementPolicyProperties_STATUS instead
type ManagementPolicyProperties_STATUS_ARM struct {
	LastModifiedTime *string                            `json:"lastModifiedTime,omitempty"`
	Policy           *ManagementPolicySchema_STATUS_ARM `json:"policy,omitempty"`
}

// Deprecated version of ManagementPolicySchema_STATUS. Use v1api20210401.ManagementPolicySchema_STATUS instead
type ManagementPolicySchema_STATUS_ARM struct {
	Rules []ManagementPolicyRule_STATUS_ARM `json:"rules,omitempty"`
}

// Deprecated version of ManagementPolicyRule_STATUS. Use v1api20210401.ManagementPolicyRule_STATUS instead
type ManagementPolicyRule_STATUS_ARM struct {
	Definition *ManagementPolicyDefinition_STATUS_ARM `json:"definition,omitempty"`
	Enabled    *bool                                  `json:"enabled,omitempty"`
	Name       *string                                `json:"name,omitempty"`
	Type       *ManagementPolicyRule_Type_STATUS      `json:"type,omitempty"`
}

// Deprecated version of ManagementPolicyDefinition_STATUS. Use v1api20210401.ManagementPolicyDefinition_STATUS instead
type ManagementPolicyDefinition_STATUS_ARM struct {
	Actions *ManagementPolicyAction_STATUS_ARM `json:"actions,omitempty"`
	Filters *ManagementPolicyFilter_STATUS_ARM `json:"filters,omitempty"`
}

// Deprecated version of ManagementPolicyAction_STATUS. Use v1api20210401.ManagementPolicyAction_STATUS instead
type ManagementPolicyAction_STATUS_ARM struct {
	BaseBlob *ManagementPolicyBaseBlob_STATUS_ARM `json:"baseBlob,omitempty"`
	Snapshot *ManagementPolicySnapShot_STATUS_ARM `json:"snapshot,omitempty"`
	Version  *ManagementPolicyVersion_STATUS_ARM  `json:"version,omitempty"`
}

// Deprecated version of ManagementPolicyFilter_STATUS. Use v1api20210401.ManagementPolicyFilter_STATUS instead
type ManagementPolicyFilter_STATUS_ARM struct {
	BlobIndexMatch []TagFilter_STATUS_ARM `json:"blobIndexMatch,omitempty"`
	BlobTypes      []string               `json:"blobTypes,omitempty"`
	PrefixMatch    []string               `json:"prefixMatch,omitempty"`
}

// Deprecated version of ManagementPolicyBaseBlob_STATUS. Use v1api20210401.ManagementPolicyBaseBlob_STATUS instead
type ManagementPolicyBaseBlob_STATUS_ARM struct {
	Delete                      *DateAfterModification_STATUS_ARM `json:"delete,omitempty"`
	EnableAutoTierToHotFromCool *bool                             `json:"enableAutoTierToHotFromCool,omitempty"`
	TierToArchive               *DateAfterModification_STATUS_ARM `json:"tierToArchive,omitempty"`
	TierToCool                  *DateAfterModification_STATUS_ARM `json:"tierToCool,omitempty"`
}

// Deprecated version of ManagementPolicySnapShot_STATUS. Use v1api20210401.ManagementPolicySnapShot_STATUS instead
type ManagementPolicySnapShot_STATUS_ARM struct {
	Delete        *DateAfterCreation_STATUS_ARM `json:"delete,omitempty"`
	TierToArchive *DateAfterCreation_STATUS_ARM `json:"tierToArchive,omitempty"`
	TierToCool    *DateAfterCreation_STATUS_ARM `json:"tierToCool,omitempty"`
}

// Deprecated version of ManagementPolicyVersion_STATUS. Use v1api20210401.ManagementPolicyVersion_STATUS instead
type ManagementPolicyVersion_STATUS_ARM struct {
	Delete        *DateAfterCreation_STATUS_ARM `json:"delete,omitempty"`
	TierToArchive *DateAfterCreation_STATUS_ARM `json:"tierToArchive,omitempty"`
	TierToCool    *DateAfterCreation_STATUS_ARM `json:"tierToCool,omitempty"`
}

// Deprecated version of TagFilter_STATUS. Use v1api20210401.TagFilter_STATUS instead
type TagFilter_STATUS_ARM struct {
	Name  *string `json:"name,omitempty"`
	Op    *string `json:"op,omitempty"`
	Value *string `json:"value,omitempty"`
}

// Deprecated version of DateAfterCreation_STATUS. Use v1api20210401.DateAfterCreation_STATUS instead
type DateAfterCreation_STATUS_ARM struct {
	DaysAfterCreationGreaterThan *float64 `json:"daysAfterCreationGreaterThan,omitempty"`
}

// Deprecated version of DateAfterModification_STATUS. Use v1api20210401.DateAfterModification_STATUS instead
type DateAfterModification_STATUS_ARM struct {
	DaysAfterLastAccessTimeGreaterThan *float64 `json:"daysAfterLastAccessTimeGreaterThan,omitempty"`
	DaysAfterModificationGreaterThan   *float64 `json:"daysAfterModificationGreaterThan,omitempty"`
}
