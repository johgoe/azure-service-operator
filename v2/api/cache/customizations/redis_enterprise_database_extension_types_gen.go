// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v1api20210301 "github.com/Azure/azure-service-operator/v2/api/cache/v1api20210301"
	v1api20210301s "github.com/Azure/azure-service-operator/v2/api/cache/v1api20210301storage"
	v20210301 "github.com/Azure/azure-service-operator/v2/api/cache/v1beta20210301"
	v20210301s "github.com/Azure/azure-service-operator/v2/api/cache/v1beta20210301storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type RedisEnterpriseDatabaseExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *RedisEnterpriseDatabaseExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v1api20210301.RedisEnterpriseDatabase{},
		&v1api20210301s.RedisEnterpriseDatabase{},
		&v20210301.RedisEnterpriseDatabase{},
		&v20210301s.RedisEnterpriseDatabase{}}
}
