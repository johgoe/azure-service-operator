// +build all appinsights

/*
Copyright 2019 Microsoft.

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

package controllers

import (
	"context"
	"strings"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestAppInsightsController(t *testing.T) {
	t.Parallel()
	defer PanicRecover()
	ctx := context.Background()
	assert := assert.New(t)

	rgName := tc.resourceGroupName
	rgLocation := tc.resourceGroupLocation
	appInsightsName := GenerateGroupName("appinsights")

	// Create an instance of Azure AppInsights
	appInsightsInstance := &azurev1alpha1.AppInsights{
		ObjectMeta: metav1.ObjectMeta{
			Name:      appInsightsName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AppInsightsSpec{
			Kind:            "web",
			Location:        rgLocation,
			ResourceGroup:   rgName,
			ApplicationType: "other",
		},
	}

	err := tc.k8sClient.Create(ctx, appInsightsInstance)
	assert.Equal(nil, err, "create appinsights record in k8s")

	appInsightsNamespacedName := types.NamespacedName{Name: appInsightsName, Namespace: "default"}

	// Wait for the AppInsights instance to be provisioned
	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, appInsightsNamespacedName, appInsightsInstance)
		return strings.Contains(appInsightsInstance.Status.Message, successMsg)
	}, tc.timeout, tc.retry, "awaiting appinsights instance creation")

	// Delete the service
	err = tc.k8sClient.Delete(ctx, appInsightsInstance)
	assert.Equal(nil, err, "deleting appinsights in k8s")

	// Wait for the AppInsights instance to be deleted
	assert.Eventually(func() bool {
		err := tc.k8sClient.Get(ctx, appInsightsNamespacedName, appInsightsInstance)
		return err != nil
	}, tc.timeout, tc.retry, "awaiting appInsightsInstance deletion")
}
