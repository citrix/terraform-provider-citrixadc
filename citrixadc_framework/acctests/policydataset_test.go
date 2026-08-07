/*
Copyright 2016 Citrix Systems, Inc

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
package citrixadc

import (
	"fmt"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccPolicydataset_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPolicydatasetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicydataset_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicydatasetExist("citrixadc_policydataset.tf_dataset", nil),
				),
			},
		},
	})
}

func testAccCheckPolicydatasetExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dataset name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Policydataset.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Dataset %s not found", n)
		}

		return nil
	}
}

func testAccCheckPolicydatasetDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_policydataset" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Policydataset.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("dataset %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccPolicydataset_basic = `

resource "citrixadc_policydataset" "tf_dataset" {
  name    = "tf_dataset"
  type    = "number"
}

`

const testAccPolicydatasetDataSource_basic = `
	resource "citrixadc_policydataset" "tf_dataset_ds" {
		name = "tf_dataset_ds"
		type = "ipv4"
	}

	data "citrixadc_policydataset" "tf_dataset_ds" {
		name = citrixadc_policydataset.tf_dataset_ds.name
	}
`

func TestAccPolicydataset_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_policydataset.tf_dataset"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPolicydatasetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicydataset_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckPolicydatasetExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Policydataset.Type(), "tf_dataset"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccPolicydataset_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckPolicydatasetExist(resAddr, nil)),
			},
		},
	})
}

func TestAccPolicydataset_import(t *testing.T) {
	const resAddr = "citrixadc_policydataset.tf_dataset"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPolicydatasetDestroy,
		Steps: []resource.TestStep{
			{Config: testAccPolicydataset_basic},
			{
				Config:                  testAccPolicydataset_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccPolicydataset_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckPolicydatasetDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccPolicydataset_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckPolicydatasetExist("citrixadc_policydataset.tf_dataset", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccPolicydataset_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckPolicydatasetExist("citrixadc_policydataset.tf_dataset", nil)),
			},
		},
	})
}

func TestAccPolicydatasetDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicydatasetDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_policydataset.tf_dataset_ds", "name", "tf_dataset_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_policydataset.tf_dataset_ds", "type", "ipv4"),
				),
			},
		},
	})
}
