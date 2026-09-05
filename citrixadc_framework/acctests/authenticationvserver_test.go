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

const testAccAuthenticationvserver_add = `

	resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
		name  		   = "tf_authenticationvserver"
		servicetype    = "SSL"
		comment 	   = "Hello"
		authentication = "ON"
		state          = "ENABLED"
	}
`
const testAccAuthenticationvserver_update = `

	resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
		name  		   = "tf_authenticationvserver"
		servicetype    = "SSL"
		comment 	   = "New"
		authentication = "ON"
		state          = "DISABLED"
	}
`

func TestAccAuthenticationvserver_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationvserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationvserver_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationvserverExist("citrixadc_authenticationvserver.tf_authenticationvserver", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationvserver.tf_authenticationvserver", "name", "tf_authenticationvserver"),
					resource.TestCheckResourceAttr("citrixadc_authenticationvserver.tf_authenticationvserver", "comment", "Hello"),
					resource.TestCheckResourceAttr("citrixadc_authenticationvserver.tf_authenticationvserver", "state", "ENABLED"),
				),
			},
			{
				Config: testAccAuthenticationvserver_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationvserverExist("citrixadc_authenticationvserver.tf_authenticationvserver", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationvserver.tf_authenticationvserver", "name", "tf_authenticationvserver"),
					resource.TestCheckResourceAttr("citrixadc_authenticationvserver.tf_authenticationvserver", "comment", "New"),
					resource.TestCheckResourceAttr("citrixadc_authenticationvserver.tf_authenticationvserver", "state", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationvserverExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationvserver name is set")
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
		data, err := client.FindResource(service.Authenticationvserver.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationvserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationvserverDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationvserver" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Authenticationvserver.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationvserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccAuthenticationvserver_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_authenticationvserver.tf_authenticationvserver"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationvserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationvserver_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationvserverExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Authenticationvserver.Type(), "tf_authenticationvserver"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAuthenticationvserver_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationvserverExist(resAddr, nil)),
			},
		},
	})
}

func TestAccAuthenticationvserver_import(t *testing.T) {
	const resAddr = "citrixadc_authenticationvserver.tf_authenticationvserver"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationvserverDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAuthenticationvserver_add},
			{
				Config:                  testAccAuthenticationvserver_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"state"},
			},
		},
	})
}

func TestAccAuthenticationvserver_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAuthenticationvserverDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccAuthenticationvserver_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationvserverExist("citrixadc_authenticationvserver.tf_authenticationvserver", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccAuthenticationvserver_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationvserverExist("citrixadc_authenticationvserver.tf_authenticationvserver", nil)),
			},
		},
	})
}

// The authenticationvserver unset test covers the two unset-eligible attributes
// that have a documented NITRO server default (authentication -> "ON",
// appflowlog -> "ENABLED"). The other spec-unsettable attributes lack a
// documented default and are excluded.
const testAccAuthenticationvserver_unset_step1 = `
	resource "citrixadc_authenticationvserver" "tf_unset" {
		name           = "tf_authvserver_unset"
		servicetype    = "SSL"
		authentication = "OFF"
		appflowlog     = "DISABLED"
	}
`

const testAccAuthenticationvserver_unset_step2 = `
	resource "citrixadc_authenticationvserver" "tf_unset" {
		name        = "tf_authvserver_unset"
		servicetype = "SSL"
		# unset-eligible attributes removed from config -> the provider must unset
		# them (revert to NITRO defaults).
	}
`

func TestAccAuthenticationvserver_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationvserverDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccAuthenticationvserver_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationvserverExist("citrixadc_authenticationvserver.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationvserver.tf_unset", "authentication", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_authenticationvserver.tf_unset", "appflowlog", "DISABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state reverts to the
				// documented NITRO defaults and the implicit post-apply plan is empty.
				Config: testAccAuthenticationvserver_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationvserverExist("citrixadc_authenticationvserver.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationvserver.tf_unset", "authentication", "ON"),
					resource.TestCheckResourceAttr("citrixadc_authenticationvserver.tf_unset", "appflowlog", "ENABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAuthenticationvserverADCValue("tf_authvserver_unset", "authentication", "ON"),
					testAccCheckAuthenticationvserverADCValue("tf_authvserver_unset", "appflowlog", "ENABLED"),
				),
			},
		},
	})
}

// testAccCheckAuthenticationvserverADCValue asserts an attribute's value directly
// on the appliance (not just in Terraform state), proving the unset reverted it.
func testAccCheckAuthenticationvserverADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Authenticationvserver.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("authenticationvserver %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("authenticationvserver %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccAuthenticationvserverDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationvserverDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationvserver.tf_authenticationvserver_ds", "name", "tf_authenticationvserver_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationvserver.tf_authenticationvserver_ds", "servicetype", "SSL"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationvserver.tf_authenticationvserver_ds", "comment", "DataSource Test"),
					// Universal runtime-binding proof.
					resource.TestCheckResourceAttrSet("data.citrixadc_authenticationvserver.tf_authenticationvserver_ds", "id"),
					// Read-only runtime/state attribute always populated for a live vserver.
					resource.TestCheckResourceAttrSet("data.citrixadc_authenticationvserver.tf_authenticationvserver_ds", "curstate"),
				),
			},
		},
	})
}

const testAccAuthenticationvserverDataSource_basic = `

resource "citrixadc_authenticationvserver" "tf_authenticationvserver_ds" {
	name           = "tf_authenticationvserver_ds"
	servicetype    = "SSL"
	comment        = "DataSource Test"
	authentication = "ON"
	state          = "ENABLED"
}

data "citrixadc_authenticationvserver" "tf_authenticationvserver_ds" {
	name = citrixadc_authenticationvserver.tf_authenticationvserver_ds.name
	depends_on = [citrixadc_authenticationvserver.tf_authenticationvserver_ds]
}

`
