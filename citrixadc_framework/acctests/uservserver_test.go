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

const testAccUservserver_basic = `

	resource "citrixadc_userprotocol" "tf_userprotocol" {
		name      = "MQTT"
		transport = "TCP"
		extension = "mqtt_code"
		comment   = "my_comment"
	}

	resource "citrixadc_lbvserver" "tf_defaultlb" {
		name        = "tf_defaultlb"
		servicetype = "USER_TCP"
	}

	resource "citrixadc_uservserver" "tf_uservserver" {
		name         = "my_user_vserver"
		userprotocol = "MQTT"
		ipaddress    = "10.222.74.180"
		port         = 3200
		defaultlb    = citrixadc_lbvserver.tf_defaultlb.name
		depends_on   = [citrixadc_userprotocol.tf_userprotocol, citrixadc_lbvserver.tf_defaultlb]
	}

`
const testAccUservserver_update = `

	resource "citrixadc_userprotocol" "tf_userprotocol" {
		name      = "MQTT"
		transport = "TCP"
		extension = "mqtt_code"
		comment   = "my_comment"
	}

	resource "citrixadc_lbvserver" "tf_defaultlb" {
		name        = "tf_defaultlb"
		servicetype = "USER_TCP"
	}

	resource "citrixadc_uservserver" "tf_uservserver" {
		name         = "my_user_vserver"
		userprotocol = "MQTT"
		ipaddress    = "10.222.74.200"
		port         = 3500
		defaultlb    = citrixadc_lbvserver.tf_defaultlb.name
	}
`

const testAccUservserverDataSource_basic = `

	resource "citrixadc_userprotocol" "tf_userprotocol" {
		name      = "MQTT"
		transport = "TCP"
		extension = "mqtt_code"
		comment   = "my_comment"
	}

	resource "citrixadc_lbvserver" "tf_defaultlb" {
		name        = "tf_defaultlb"
		servicetype = "USER_TCP"
	}

	resource "citrixadc_uservserver" "tf_uservserver" {
		name         = "my_user_vserver"
		userprotocol = "MQTT"
		ipaddress    = "10.222.74.180"
		port         = 3200
		defaultlb    = citrixadc_lbvserver.tf_defaultlb.name
		depends_on   = [citrixadc_userprotocol.tf_userprotocol, citrixadc_lbvserver.tf_defaultlb]
	}

	data "citrixadc_uservserver" "tf_uservserver" {
		name = citrixadc_uservserver.tf_uservserver.name
		depends_on = [citrixadc_uservserver.tf_uservserver]
	}
`

func TestAccUservserver_basic(t *testing.T) {
	t.Skip("TODO: Requires adding new ns extension. Refer https://docs.netscaler.com/en-us/citrix-adc/current-release/citrix-adc-extensions/citrix-adc-protocol-extensions/tutorial-examples!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckUservserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccUservserver_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUservserverExist("citrixadc_uservserver.tf_uservserver", nil),
					resource.TestCheckResourceAttr("citrixadc_uservserver.tf_uservserver", "name", "my_user_vserver"),
					resource.TestCheckResourceAttr("citrixadc_uservserver.tf_uservserver", "userprotocol", "MQTT"),
					resource.TestCheckResourceAttr("citrixadc_uservserver.tf_uservserver", "ipaddress", "10.222.74.180"),
					resource.TestCheckResourceAttr("citrixadc_uservserver.tf_uservserver", "port", "3200"),
					resource.TestCheckResourceAttr("citrixadc_uservserver.tf_uservserver", "defaultlb", "tf_defaultlb"),
				),
			},
			{
				Config: testAccUservserver_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUservserverExist("citrixadc_uservserver.tf_uservserver", nil),
					resource.TestCheckResourceAttr("citrixadc_uservserver.tf_uservserver", "name", "my_user_vserver"),
					resource.TestCheckResourceAttr("citrixadc_uservserver.tf_uservserver", "userprotocol", "MQTT"),
					resource.TestCheckResourceAttr("citrixadc_uservserver.tf_uservserver", "ipaddress", "10.222.74.200"),
					resource.TestCheckResourceAttr("citrixadc_uservserver.tf_uservserver", "port", "3500"),
					resource.TestCheckResourceAttr("citrixadc_uservserver.tf_uservserver", "defaultlb", "tf_defaultlb"),
				),
			},
		},
	})
}

func TestAccUservserver_selfHealing(t *testing.T) {
	t.Skip("TODO: Requires adding new ns extension. Refer https://docs.netscaler.com/en-us/citrix-adc/current-release/citrix-adc-extensions/citrix-adc-protocol-extensions/tutorial-examples!")
	const resAddr = "citrixadc_uservserver.tf_uservserver"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckUservserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccUservserver_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckUservserverExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Uservserver.Type(), "my_user_vserver"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccUservserver_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckUservserverExist(resAddr, nil)),
			},
		},
	})
}

func testAccCheckUservserverExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No uservserver name is set")
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
		data, err := client.FindResource("uservserver", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("uservserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckUservserverDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_uservserver" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("uservserver", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("uservserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccUservserver_import(t *testing.T) {
	t.Skip("TODO: Requires adding new ns extension. Refer https://docs.netscaler.com/en-us/citrix-adc/current-release/citrix-adc-extensions/citrix-adc-protocol-extensions/tutorial-examples!")
	const resAddr = "citrixadc_uservserver.tf_uservserver"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckUservserverDestroy,
		Steps: []resource.TestStep{
			{Config: testAccUservserver_basic},
			{
				Config:                  testAccUservserver_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccUservserver_sdkv2StateUpgrade(t *testing.T) {
	t.Skip("TODO: Requires adding new ns extension. Refer https://docs.netscaler.com/en-us/citrix-adc/current-release/citrix-adc-extensions/citrix-adc-protocol-extensions/tutorial-examples!")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckUservserverDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccUservserver_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckUservserverExist("citrixadc_uservserver.tf_uservserver", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccUservserver_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckUservserverExist("citrixadc_uservserver.tf_uservserver", nil)),
			},
		},
	})
}

// The uservserver unset test exercises the two spec-unsettable attributes
// (comment and Params). Step 1 sets them to non-default values; step 2 removes
// them from config, so the provider must issue a NITRO unset that reverts them
// to their (empty) defaults.
const testAccUservserver_unset_step1 = `

	resource "citrixadc_userprotocol" "tf_userprotocol" {
		name      = "MQTT"
		transport = "TCP"
		extension = "mqtt_code"
		comment   = "my_comment"
	}

	resource "citrixadc_lbvserver" "tf_defaultlb" {
		name        = "tf_defaultlb"
		servicetype = "USER_TCP"
	}

	resource "citrixadc_uservserver" "tf_unset" {
		name         = "my_user_vserver"
		userprotocol = "MQTT"
		ipaddress    = "10.222.74.180"
		port         = 3200
		defaultlb    = citrixadc_lbvserver.tf_defaultlb.name
		comment      = "unset_test_comment"
		params       = "unset_test_params"
		depends_on   = [citrixadc_userprotocol.tf_userprotocol, citrixadc_lbvserver.tf_defaultlb]
	}
`

const testAccUservserver_unset_step2 = `

	resource "citrixadc_userprotocol" "tf_userprotocol" {
		name      = "MQTT"
		transport = "TCP"
		extension = "mqtt_code"
		comment   = "my_comment"
	}

	resource "citrixadc_lbvserver" "tf_defaultlb" {
		name        = "tf_defaultlb"
		servicetype = "USER_TCP"
	}

	resource "citrixadc_uservserver" "tf_unset" {
		name         = "my_user_vserver"
		userprotocol = "MQTT"
		ipaddress    = "10.222.74.180"
		port         = 3200
		defaultlb    = citrixadc_lbvserver.tf_defaultlb.name
		# comment and params removed from config -> provider must unset them.
		depends_on   = [citrixadc_userprotocol.tf_userprotocol, citrixadc_lbvserver.tf_defaultlb]
	}
`

func TestAccUservserver_unset(t *testing.T) {
	t.Skip("TODO: Requires adding new ns extension. Refer https://docs.netscaler.com/en-us/citrix-adc/current-release/citrix-adc-extensions/citrix-adc-protocol-extensions/tutorial-examples!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckUservserverDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccUservserver_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUservserverExist("citrixadc_uservserver.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_uservserver.tf_unset", "comment", "unset_test_comment"),
					resource.TestCheckResourceAttr("citrixadc_uservserver.tf_unset", "params", "unset_test_params"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the empty NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccUservserver_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckUservserverExist("citrixadc_uservserver.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_uservserver.tf_unset", "comment", ""),
					resource.TestCheckResourceAttr("citrixadc_uservserver.tf_unset", "params", ""),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckUservserverADCValue("my_user_vserver", "comment", ""),
					testAccCheckUservserverADCValue("my_user_vserver", "Params", ""),
				),
			},
		},
	})
}

// testAccCheckUservserverADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it.
func testAccCheckUservserverADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Uservserver.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("uservserver %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if want == "" && (data[attr] == nil || got == "") {
			return nil
		}
		if got != want {
			return fmt.Errorf("uservserver %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccUservserverDataSource_basic(t *testing.T) {
	t.Skip("TODO: Requires adding new ns extension. Refer https://docs.netscaler.com/en-us/citrix-adc/current-release/citrix-adc-extensions/citrix-adc-protocol-extensions/tutorial-examples!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccUservserverDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_uservserver.tf_uservserver", "name", "my_user_vserver"),
					resource.TestCheckResourceAttr("data.citrixadc_uservserver.tf_uservserver", "userprotocol", "MQTT"),
					resource.TestCheckResourceAttr("data.citrixadc_uservserver.tf_uservserver", "ipaddress", "10.222.74.180"),
					resource.TestCheckResourceAttr("data.citrixadc_uservserver.tf_uservserver", "port", "3200"),
					// Universal runtime-binding proof for the data source read.
					resource.TestCheckResourceAttrSet("data.citrixadc_uservserver.tf_uservserver", "id"),
				),
			},
		},
	})
}
