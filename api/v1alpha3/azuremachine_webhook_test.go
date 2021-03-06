/*
Copyright 2020 The Kubernetes Authors.

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

package v1alpha3

import (
	"testing"
)

func TestAzureMachine_ValidateCreate(t *testing.T) {
	tests := []struct {
		name    string
		machine *AzureMachine
		wantErr bool
	}{
		{
			name:    "azuremachine with marketplace image - full",
			machine: createMachineWithtMarketPlaceImage(t, "PUB1234", "OFFER1234", "SKU1234", "1.0.0"),
			wantErr: false,
		},
		{
			name:    "azuremachine with marketplace image - missing publisher",
			machine: createMachineWithtMarketPlaceImage(t, "", "OFFER1234", "SKU1234", "1.0.0"),
			wantErr: true,
		},
		{
			name:    "azuremachine with shared gallery image - full",
			machine: createMachineWithSharedImage(t, "SUB123", "RG123", "NAME123", "GALLERY1", "1.0.0"),
			wantErr: false,
		},
		{
			name:    "azuremachine with marketplace image - missing subscription",
			machine: createMachineWithSharedImage(t, "", "RG123", "NAME123", "GALLERY1", "1.0.0"),
			wantErr: true,
		},
		{
			name:    "azuremachine with image by - with id",
			machine: createMachineWithImageByID(t, "ID123"),
			wantErr: false,
		},
		{
			name:    "azuremachine with image by - without id",
			machine: createMachineWithImageByID(t, ""),
			wantErr: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.machine.ValidateCreate(); (err != nil) != tc.wantErr {
				t.Errorf("ValidateCreate() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func createMachineWithSharedImage(t *testing.T, subscriptionID, resourceGroup, name, gallery, version string) *AzureMachine {
	image := &Image{
		SharedGallery: &AzureSharedGalleryImage{
			SubscriptionID: subscriptionID,
			ResourceGroup:  resourceGroup,
			Name:           name,
			Gallery:        gallery,
			Version:        version,
		},
	}

	return &AzureMachine{
		Spec: AzureMachineSpec{
			Image: image,
		},
	}

}

func createMachineWithtMarketPlaceImage(t *testing.T, publisher, offer, sku, version string) *AzureMachine {
	image := &Image{
		Marketplace: &AzureMarketplaceImage{
			Publisher: publisher,
			Offer:     offer,
			SKU:       sku,
			Version:   version,
		},
	}

	return &AzureMachine{
		Spec: AzureMachineSpec{
			Image: image,
		},
	}
}

func createMachineWithImageByID(t *testing.T, imageID string) *AzureMachine {
	image := &Image{
		ID: &imageID,
	}

	return &AzureMachine{
		Spec: AzureMachineSpec{
			Image: image,
		},
	}
}
