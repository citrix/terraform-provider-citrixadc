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
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testAccUserprotocol_basic = `

	resource "citrixadc_userprotocol" "tf_userprotocol" {
		name      = "my_userprotocol"
		transport = "TCP"
		extension = "mqtt_code"
		comment   = "my_comment"
	}
`
const testAccUserprotocol_update = `

	resource "citrixadc_userprotocol" "tf_userprotocol" {
		name      = "my_userprotocol"
		transport = "SSL"
		extension = "mqtt_code"
		comment   = "my_new_comment"
	}
`

const testAccUserprotocolDataSource_basic = `

	resource "citrixadc_userprotocol" "tf_userprotocol" {
		name      = "MQTT"
		transport = "TCP"
		extension = "mqtt_code"
		comment   = "my_comment"
	}

	data "citrixadc_userprotocol" "tf_userprotocol" {
		name = citrixadc_userprotocol.tf_userprotocol.name
	}
`

func TestAccUserprotocol_basic(t *testing.T) {
	t.Skip("TODO: Requires adding new ns extension. Refer https://docs.netscaler.com/en-us/citrix-adc/current-release/citrix-adc-extensions/citrix-adc-protocol-extensions/tutorial-examples!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckUserprotocolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccUserprotocol_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserprotocolExist("citrixadc_userprotocol.tf_userprotocol", nil),
					resource.TestCheckResourceAttr("citrixadc_userprotocol.tf_userprotocol", "name", "my_userprotocol"),
					resource.TestCheckResourceAttr("citrixadc_userprotocol.tf_userprotocol", "transport", "TCP"),
					resource.TestCheckResourceAttr("citrixadc_userprotocol.tf_userprotocol", "extension", "mqtt_code"),
					resource.TestCheckResourceAttr("citrixadc_userprotocol.tf_userprotocol", "comment", "my_comment"),
				),
			},
			{
				Config: testAccUserprotocol_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserprotocolExist("citrixadc_userprotocol.tf_userprotocol", nil),
					resource.TestCheckResourceAttr("citrixadc_userprotocol.tf_userprotocol", "name", "my_userprotocol"),
					resource.TestCheckResourceAttr("citrixadc_userprotocol.tf_userprotocol", "transport", "SSL"),
					resource.TestCheckResourceAttr("citrixadc_userprotocol.tf_userprotocol", "extension", "mqtt_code"),
					resource.TestCheckResourceAttr("citrixadc_userprotocol.tf_userprotocol", "comment", "my_new_comment"),
				),
			},
		},
	})
}

func TestAccUserprotocol_import(t *testing.T) {
	t.Skip("TODO: Requires adding new ns extension. Refer https://docs.netscaler.com/en-us/citrix-adc/current-release/citrix-adc-extensions/citrix-adc-protocol-extensions/tutorial-examples!")
	const resAddr = "citrixadc_userprotocol.tf_userprotocol"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckUserprotocolDestroy,
		Steps: []resource.TestStep{
			{Config: testAccUserprotocol_basic},
			{
				Config:                  testAccUserprotocol_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckUserprotocolExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No userprotocol name is set")
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
		data, err := client.FindResource("userprotocol", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("userprotocol %s not found", n)
		}

		return nil
	}
}

func testAccCheckUserprotocolDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_userprotocol" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("userprotocol", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("userprotocol %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccUserprotocol_selfHealing(t *testing.T) {
	t.Skip("TODO: Requires adding new ns extension. Refer https://docs.netscaler.com/en-us/citrix-adc/current-release/citrix-adc-extensions/citrix-adc-protocol-extensions/tutorial-examples!")
	const resAddr = "citrixadc_userprotocol.tf_userprotocol"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckUserprotocolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccUserprotocol_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckUserprotocolExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Userprotocol.Type(), "my_userprotocol"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccUserprotocol_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckUserprotocolExist(resAddr, nil)),
			},
		},
	})
}

func TestAccUserprotocol_sdkv2StateUpgrade(t *testing.T) {
	t.Skip("TODO: Requires adding new ns extension. Refer https://docs.netscaler.com/en-us/citrix-adc/current-release/citrix-adc-extensions/citrix-adc-protocol-extensions/tutorial-examples!")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckUserprotocolDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccUserprotocol_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckUserprotocolExist("citrixadc_userprotocol.tf_userprotocol", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccUserprotocol_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckUserprotocolExist("citrixadc_userprotocol.tf_userprotocol", nil)),
			},
		},
	})
}

// The userprotocol unset test covers the single unset-eligible attribute,
// `comment` (name/transport/extension are key/ForceNew). Step 1 sets it to a
// non-default value; step 2 removes it from config so the provider unsets it
// and the appliance reverts it to the NITRO default (empty comment).
const testAccUserprotocol_unset_step1 = `

	resource "citrixadc_userprotocol" "tf_unset" {
		name      = "tf_test_userprotocol_unset"
		transport = "TCP"
		extension = "mqtt_code"
		comment   = "my_comment"
	}
`

const testAccUserprotocol_unset_step2 = `

	resource "citrixadc_userprotocol" "tf_unset" {
		name      = "tf_test_userprotocol_unset"
		transport = "TCP"
		extension = "mqtt_code"
		# comment removed from config -> the provider must unset it (revert to the
		# NITRO default, an empty comment).
	}
`

func TestAccUserprotocol_unset(t *testing.T) {
	t.Skip("TODO: Requires adding new ns extension. Refer https://docs.netscaler.com/en-us/citrix-adc/current-release/citrix-adc-extensions/citrix-adc-protocol-extensions/tutorial-examples!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckUserprotocolDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccUserprotocol_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserprotocolExist("citrixadc_userprotocol.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_userprotocol.tf_unset", "comment", "my_comment"),
				),
			},
			{
				// Removing comment must unset it: state (read back from the
				// appliance) reverts to the NITRO default, and the implicit
				// post-apply plan must be empty.
				Config: testAccUserprotocol_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUserprotocolExist("citrixadc_userprotocol.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_userprotocol.tf_unset", "comment", ""),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckUserprotocolADCValue("tf_test_userprotocol_unset", "comment", ""),
				),
			},
		},
	})
}

// testAccCheckUserprotocolADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it.
func testAccCheckUserprotocolADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Userprotocol.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("userprotocol %s not found on appliance", name)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = fmt.Sprintf("%v", v)
		}
		if got != want {
			return fmt.Errorf("userprotocol %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccUserprotocolDataSource_basic(t *testing.T) {
	t.Skip("TODO: Requires adding new ns extension. Refer https://docs.netscaler.com/en-us/citrix-adc/current-release/citrix-adc-extensions/citrix-adc-protocol-extensions/tutorial-examples!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUserprotocolDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_userprotocol.tf_userprotocol", "name", "MQTT"),
					resource.TestCheckResourceAttr("data.citrixadc_userprotocol.tf_userprotocol", "transport", "TCP"),
					resource.TestCheckResourceAttr("data.citrixadc_userprotocol.tf_userprotocol", "extension", "mqtt_code"),
					resource.TestCheckResourceAttr("data.citrixadc_userprotocol.tf_userprotocol", "comment", "my_comment"),
				),
			},
		},
	})
}
