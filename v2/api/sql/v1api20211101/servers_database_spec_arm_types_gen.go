// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Servers_Database_Spec_ARM struct {
	// Identity: The Azure Active Directory identity of the database.
	Identity *DatabaseIdentity_ARM `json:"identity,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Resource properties.
	Properties *DatabaseProperties_ARM `json:"properties,omitempty"`

	// Sku: The database SKU.
	// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition,
	// family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation`
	// REST API or one of the following commands:
	// ```azurecli
	// az sql db list-editions -l <location> -o table
	// ````
	// ```powershell
	// Get-AzSqlServerServiceObjective -Location <location>
	// ````
	Sku *Sku_ARM `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Servers_Database_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (database Servers_Database_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (database *Servers_Database_Spec_ARM) GetName() string {
	return database.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/databases"
func (database *Servers_Database_Spec_ARM) GetType() string {
	return "Microsoft.Sql/servers/databases"
}

// Azure Active Directory identity configuration for a resource.
type DatabaseIdentity_ARM struct {
	// Type: The identity type
	Type *DatabaseIdentity_Type `json:"type,omitempty"`
}

// The database's properties.
type DatabaseProperties_ARM struct {
	// AutoPauseDelay: Time in minutes after which database is automatically paused. A value of -1 means that automatic pause
	// is disabled
	AutoPauseDelay *int `json:"autoPauseDelay,omitempty"`

	// CatalogCollation: Collation of the metadata catalog.
	CatalogCollation *DatabaseProperties_CatalogCollation `json:"catalogCollation,omitempty"`

	// Collation: The collation of the database.
	Collation *string `json:"collation,omitempty"`

	// CreateMode: Specifies the mode of database creation.
	// Default: regular database creation.
	// Copy: creates a database as a copy of an existing database. sourceDatabaseId must be specified as the resource ID of the
	// source database.
	// Secondary: creates a database as a secondary replica of an existing database. sourceDatabaseId must be specified as the
	// resource ID of the existing primary database.
	// PointInTimeRestore: Creates a database by restoring a point in time backup of an existing database. sourceDatabaseId
	// must be specified as the resource ID of the existing database, and restorePointInTime must be specified.
	// Recovery: Creates a database by restoring a geo-replicated backup. sourceDatabaseId must be specified as the recoverable
	// database resource ID to restore.
	// Restore: Creates a database by restoring a backup of a deleted database. sourceDatabaseId must be specified. If
	// sourceDatabaseId is the database's original resource ID, then sourceDatabaseDeletionDate must be specified. Otherwise
	// sourceDatabaseId must be the restorable dropped database resource ID and sourceDatabaseDeletionDate is ignored.
	// restorePointInTime may also be specified to restore from an earlier point in time.
	// RestoreLongTermRetentionBackup: Creates a database by restoring from a long term retention vault.
	// recoveryServicesRecoveryPointResourceId must be specified as the recovery point resource ID.
	// Copy, Secondary, and RestoreLongTermRetentionBackup are not supported for DataWarehouse edition.
	CreateMode    *DatabaseProperties_CreateMode `json:"createMode,omitempty"`
	ElasticPoolId *string                        `json:"elasticPoolId,omitempty"`

	// FederatedClientId: The Client id used for cross tenant per database CMK scenario
	FederatedClientId *string `json:"federatedClientId,omitempty"`

	// HighAvailabilityReplicaCount: The number of secondary replicas associated with the database that are used to provide
	// high availability. Not applicable to a Hyperscale database within an elastic pool.
	HighAvailabilityReplicaCount *int `json:"highAvailabilityReplicaCount,omitempty"`

	// IsLedgerOn: Whether or not this database is a ledger database, which means all tables in the database are ledger tables.
	// Note: the value of this property cannot be changed after the database has been created.
	IsLedgerOn *bool `json:"isLedgerOn,omitempty"`

	// LicenseType: The license type to apply for this database. `LicenseIncluded` if you need a license, or `BasePrice` if you
	// have a license and are eligible for the Azure Hybrid Benefit.
	LicenseType                       *DatabaseProperties_LicenseType `json:"licenseType,omitempty"`
	LongTermRetentionBackupResourceId *string                         `json:"longTermRetentionBackupResourceId,omitempty"`

	// MaintenanceConfigurationId: Maintenance configuration id assigned to the database. This configuration defines the period
	// when the maintenance updates will occur.
	MaintenanceConfigurationId *string `json:"maintenanceConfigurationId,omitempty"`

	// MaxSizeBytes: The max size of the database expressed in bytes.
	MaxSizeBytes *int `json:"maxSizeBytes,omitempty"`

	// MinCapacity: Minimal capacity that database will always have allocated, if not paused
	MinCapacity *float64 `json:"minCapacity,omitempty"`

	// ReadScale: The state of read-only routing. If enabled, connections that have application intent set to readonly in their
	// connection string may be routed to a readonly secondary replica in the same region. Not applicable to a Hyperscale
	// database within an elastic pool.
	ReadScale                       *DatabaseProperties_ReadScale `json:"readScale,omitempty"`
	RecoverableDatabaseId           *string                       `json:"recoverableDatabaseId,omitempty"`
	RecoveryServicesRecoveryPointId *string                       `json:"recoveryServicesRecoveryPointId,omitempty"`

	// RequestedBackupStorageRedundancy: The storage account type to be used to store backups for this database.
	RequestedBackupStorageRedundancy *DatabaseProperties_RequestedBackupStorageRedundancy `json:"requestedBackupStorageRedundancy,omitempty"`
	RestorableDroppedDatabaseId      *string                                              `json:"restorableDroppedDatabaseId,omitempty"`

	// RestorePointInTime: Specifies the point in time (ISO8601 format) of the source database that will be restored to create
	// the new database.
	RestorePointInTime *string `json:"restorePointInTime,omitempty"`

	// SampleName: The name of the sample schema to apply when creating this database.
	SampleName *DatabaseProperties_SampleName `json:"sampleName,omitempty"`

	// SecondaryType: The secondary type of the database if it is a secondary.  Valid values are Geo and Named.
	SecondaryType *DatabaseProperties_SecondaryType `json:"secondaryType,omitempty"`

	// SourceDatabaseDeletionDate: Specifies the time that the database was deleted.
	SourceDatabaseDeletionDate *string `json:"sourceDatabaseDeletionDate,omitempty"`
	SourceDatabaseId           *string `json:"sourceDatabaseId,omitempty"`
	SourceResourceId           *string `json:"sourceResourceId,omitempty"`

	// ZoneRedundant: Whether or not this database is zone redundant, which means the replicas of this database will be spread
	// across multiple availability zones.
	ZoneRedundant *bool `json:"zoneRedundant,omitempty"`
}

// An ARM Resource SKU.
type Sku_ARM struct {
	// Capacity: Capacity of the particular SKU.
	Capacity *int `json:"capacity,omitempty"`

	// Family: If the service has different generations of hardware, for the same SKU, then that can be captured here.
	Family *string `json:"family,omitempty"`

	// Name: The name of the SKU, typically, a letter + Number code, e.g. P3.
	Name *string `json:"name,omitempty"`

	// Size: Size of the particular SKU
	Size *string `json:"size,omitempty"`

	// Tier: The tier or edition of the particular SKU, e.g. Basic, Premium.
	Tier *string `json:"tier,omitempty"`
}

// +kubebuilder:validation:Enum={"None","UserAssigned"}
type DatabaseIdentity_Type string

const (
	DatabaseIdentity_Type_None         = DatabaseIdentity_Type("None")
	DatabaseIdentity_Type_UserAssigned = DatabaseIdentity_Type("UserAssigned")
)
