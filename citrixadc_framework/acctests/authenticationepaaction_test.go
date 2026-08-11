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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccAuthenticationepaaction_add = `

	resource "citrixadc_authenticationepaaction" "tf_epaaction" {
		name            = "tf_epaaction"
		csecexpr        = "sys.client_expr (\"app_0_MAC-BROWSER_1001_VERSION_<=_10.0.3\")"
		defaultepagroup = "new_group"
		deletefiles     = "old_files"
		killprocess     = "old_process"
	}

	resource "citrixadc_authenticationepaaction" "tf_epaaction2" {
		name            = "tf_epaaction2"
		deviceposture	 = "DISABLED"
		defaultepagroup = "new_group"
	}
`
const testAccAuthenticationepaaction_update = `

	resource "citrixadc_authenticationepaaction" "tf_epaaction" {
		name            = "tf_epaaction"
		csecexpr        = "sys.client_expr (\"app_0_MAC-BROWSER_1001_VERSION_<=_10.0.3\")"
		defaultepagroup = "new_group"
		deletefiles     = "new_files"
		killprocess     = "new_process"
	}

	resource "citrixadc_authenticationepaaction" "tf_epaaction2" {
		name            = "tf_epaaction2"
		deviceposture	 = "ENABLED"
		defaultepagroup = "new_group"
	}
`

func TestAccAuthenticationepaaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationepaactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationepaaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationepaactionExist("citrixadc_authenticationepaaction.tf_epaaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationepaaction.tf_epaaction", "name", "tf_epaaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationepaaction.tf_epaaction", "deletefiles", "old_files"),
					resource.TestCheckResourceAttr("citrixadc_authenticationepaaction.tf_epaaction", "killprocess", "old_process"),
					testAccCheckAuthenticationepaactionExist("citrixadc_authenticationepaaction.tf_epaaction2", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationepaaction.tf_epaaction2", "name", "tf_epaaction2"),
					resource.TestCheckResourceAttr("citrixadc_authenticationepaaction.tf_epaaction2", "deviceposture", "DISABLED"),
				),
			},
			{
				Config: testAccAuthenticationepaaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationepaactionExist("citrixadc_authenticationepaaction.tf_epaaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationepaaction.tf_epaaction", "name", "tf_epaaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationepaaction.tf_epaaction", "deletefiles", "new_files"),
					resource.TestCheckResourceAttr("citrixadc_authenticationepaaction.tf_epaaction", "killprocess", "new_process"),
					testAccCheckAuthenticationepaactionExist("citrixadc_authenticationepaaction.tf_epaaction2", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationepaaction.tf_epaaction2", "name", "tf_epaaction2"),
					resource.TestCheckResourceAttr("citrixadc_authenticationepaaction.tf_epaaction2", "deviceposture", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationepaactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationepaaction name is set")
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
		data, err := client.FindResource("authenticationepaaction", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationepaaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationepaactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationepaaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("aAuthenticationepaaction", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationepaaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAuthenticationepaactionDataSource_basic = `

	resource "citrixadc_authenticationepaaction" "tf_epaaction_ds" {
		name            = "tf_epaaction_ds"
		csecexpr        = "sys.client_expr (\"app_0_MAC-BROWSER_1001_VERSION_<=_10.0.3\")"
		defaultepagroup = "new_group"
		deletefiles     = "old_files"
		killprocess     = "old_process"
	}

	data "citrixadc_authenticationepaaction" "tf_epaaction_ds" {
		name = citrixadc_authenticationepaaction.tf_epaaction_ds.name
	}
`

func TestAccAuthenticationepaaction_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_authenticationepaaction.tf_epaaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationepaactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationepaaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationepaactionExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Authenticationepaaction.Type(), "tf_epaaction"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAuthenticationepaaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationepaactionExist(resAddr, nil)),
			},
		},
	})
}

func TestAccAuthenticationepaaction_import(t *testing.T) {
	const resAddr = "citrixadc_authenticationepaaction.tf_epaaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationepaactionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAuthenticationepaaction_add},
			{
				Config:                  testAccAuthenticationepaaction_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccAuthenticationepaaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAuthenticationepaactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccAuthenticationepaaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationepaactionExist("citrixadc_authenticationepaaction.tf_epaaction", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAuthenticationepaaction_add,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckAuthenticationepaactionExist("citrixadc_authenticationepaaction.tf_epaaction", nil)),
			},
		},
	})
}

// The unset test covers the four NITRO unset-eligible attributes for
// authenticationepaaction (killprocess, deletefiles, defaultepagroup,
// quarantinegroup). csecexpr and deviceposture are NOT listed in the NITRO
// unset operation, so they are not unset-eligible and are excluded.
const testAccAuthenticationepaaction_unset_step1 = `
	resource "citrixadc_authenticationepaaction" "tf_unset" {
		name            = "tf_epaaction_unset"
		csecexpr        = "sys.client_expr (\"app_0_MAC-BROWSER_1001_VERSION_<=_10.0.3\")"
		killprocess     = "some_process"
		deletefiles     = "/tmp/somefile"
		defaultepagroup = "def_group"
		quarantinegroup = "quar_group"
	}
`

const testAccAuthenticationepaaction_unset_step2 = `
	resource "citrixadc_authenticationepaaction" "tf_unset" {
		name     = "tf_epaaction_unset"
		# csecexpr is a prerequisite for killprocess and is not unset-eligible, so
		# it is retained. All unset-eligible attributes are removed from config ->
		# the provider must unset them (revert to NITRO defaults / absent).
		csecexpr = "sys.client_expr (\"app_0_MAC-BROWSER_1001_VERSION_<=_10.0.3\")"
	}
`

func TestAccAuthenticationepaaction_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationepaactionDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccAuthenticationepaaction_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationepaactionExist("citrixadc_authenticationepaaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationepaaction.tf_unset", "killprocess", "some_process"),
					resource.TestCheckResourceAttr("citrixadc_authenticationepaaction.tf_unset", "deletefiles", "/tmp/somefile"),
					resource.TestCheckResourceAttr("citrixadc_authenticationepaaction.tf_unset", "defaultepagroup", "def_group"),
					resource.TestCheckResourceAttr("citrixadc_authenticationepaaction.tf_unset", "quarantinegroup", "quar_group"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the NITRO defaults (absent/empty), and
				// the implicit post-apply plan must be empty.
				Config: testAccAuthenticationepaaction_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationepaactionExist("citrixadc_authenticationepaaction.tf_unset", nil),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAuthenticationepaactionADCValue("tf_epaaction_unset", "killprocess", ""),
					testAccCheckAuthenticationepaactionADCValue("tf_epaaction_unset", "deletefiles", ""),
					testAccCheckAuthenticationepaactionADCValue("tf_epaaction_unset", "defaultepagroup", ""),
					testAccCheckAuthenticationepaactionADCValue("tf_epaaction_unset", "quarantinegroup", ""),
				),
			},
		},
	})
}

// testAccCheckAuthenticationepaactionADCValue asserts an attribute's value
// directly on the appliance (not just in Terraform state), proving the unset
// actually reverted it. An absent/nil value is treated as the empty string.
func testAccCheckAuthenticationepaactionADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Authenticationepaaction.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("authenticationepaaction %s not found on appliance", name)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("authenticationepaaction %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccAuthenticationepaactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationepaactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationepaaction.tf_epaaction_ds", "name", "tf_epaaction_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationepaaction.tf_epaaction_ds", "csecexpr", "sys.client_expr (\"app_0_MAC-BROWSER_1001_VERSION_<=_10.0.3\")"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationepaaction.tf_epaaction_ds", "defaultepagroup", "new_group"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationepaaction.tf_epaaction_ds", "deletefiles", "old_files"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationepaaction.tf_epaaction_ds", "killprocess", "old_process"),
				),
			},
		},
	})
}
