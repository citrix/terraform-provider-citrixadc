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
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccPolicydataset_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckPolicydatasetExist("citrixadc_policydataset.tf_dataset", nil)),
			},
		},
	})
}

const testAccPolicydataset_unset_step1 = `
resource "citrixadc_policydataset" "tf_unset" {
  name    = "tf_test_policydataset_unset"
  type    = "ipv4"
  dynamic = "YES"
}
`

const testAccPolicydataset_unset_step2 = `
resource "citrixadc_policydataset" "tf_unset" {
  name = "tf_test_policydataset_unset"
  type = "ipv4"
  # dynamic removed from config -> the provider must unset it (revert to NITRO default "NO").
}
`

func TestAccPolicydataset_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPolicydatasetDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccPolicydataset_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicydatasetExist("citrixadc_policydataset.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_policydataset.tf_unset", "dynamic", "YES"),
				),
			},
			{
				// Removing the attribute must unset it: state (read back from the
				// appliance) reverts to the documented NITRO default, and the
				// implicit post-apply plan must be empty.
				Config: testAccPolicydataset_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicydatasetExist("citrixadc_policydataset.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_policydataset.tf_unset", "dynamic", "NO"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckPolicydatasetADCValue("tf_test_policydataset_unset", "dynamic", "NO"),
				),
			},
		},
	})
}

// testAccCheckPolicydatasetADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckPolicydatasetADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Policydataset.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("policydataset %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("policydataset %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
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
