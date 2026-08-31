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

const testAccCrvserver_add = `

	resource "citrixadc_crvserver" "crvserver" {
		name = "my_vserver"
		servicetype = "HTTP"
		arp = "OFF"
	}
`

const testAccCrvserver_update = `

	resource "citrixadc_crvserver" "crvserver" {
		name = "my_vserver"
		servicetype = "HTTP"
		arp = "ON"
	}
`

func TestAccCrvserver_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCrvserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCrvserver_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrvserverExist("citrixadc_crvserver.crvserver", nil),
					resource.TestCheckResourceAttr("citrixadc_crvserver.crvserver", "name", "my_vserver"),
					resource.TestCheckResourceAttr("citrixadc_crvserver.crvserver", "servicetype", "HTTP"),
					resource.TestCheckResourceAttr("citrixadc_crvserver.crvserver", "arp", "OFF"),
				),
			},
			{
				Config: testAccCrvserver_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrvserverExist("citrixadc_crvserver.crvserver", nil),
					resource.TestCheckResourceAttr("citrixadc_crvserver.crvserver", "name", "my_vserver"),
					resource.TestCheckResourceAttr("citrixadc_crvserver.crvserver", "servicetype", "HTTP"),
					resource.TestCheckResourceAttr("citrixadc_crvserver.crvserver", "arp", "ON"),
				),
			},
		},
	})
}

func testAccCheckCrvserverExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No crvserver name is set")
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
		data, err := client.FindResource(service.Crvserver.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("crvserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckCrvserverDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_crvserver" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Crvserver.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("crvserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccCrvserver_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_crvserver.crvserver"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCrvserverDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCrvserver_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCrvserverExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Crvserver.Type(), "my_vserver"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccCrvserver_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCrvserverExist(resAddr, nil)),
			},
		},
	})
}

func TestAccCrvserver_import(t *testing.T) {
	const resAddr = "citrixadc_crvserver.crvserver"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCrvserverDestroy,
		Steps: []resource.TestStep{
			{Config: testAccCrvserver_add},
			{
				Config:                  testAccCrvserver_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccCrvserver_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckCrvserverDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccCrvserver_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCrvserverExist("citrixadc_crvserver.crvserver", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccCrvserver_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCrvserverExist("citrixadc_crvserver.crvserver", nil)),
			},
		},
	})
}

func TestAccCrvserverDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccCrvserverDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_crvserver.tf_crvserver_ds", "name", "my_vserver_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_crvserver.tf_crvserver_ds", "servicetype", "HTTP"),
					resource.TestCheckResourceAttr("data.citrixadc_crvserver.tf_crvserver_ds", "arp", "OFF"),
				),
			},
		},
	})
}

// The crvserver unset test covers the type-independent, mutable attributes that
// have a documented NITRO default and that NITRO's unset operation reverts
// cleanly on an HTTP cr vserver. Other defaulted attributes (redirect, via,
// reuse, downstateflush, disableprimaryondown, rhistate, useoriginipportforcache)
// are NOT included: NITRO silently ignores the unset for them (the configured
// value persists), so they cannot be unset cleanly.
const testAccCrvserver_unset_step1 = `
	resource "citrixadc_crvserver" "tf_unset" {
		name            = "tf_crvserver_unset"
		servicetype     = "HTTP"
		appflowlog      = "DISABLED"
		icmpvsrresponse = "ACTIVE"
		useportrange    = "ON"
	}
`

const testAccCrvserver_unset_step2 = `
	resource "citrixadc_crvserver" "tf_unset" {
		name        = "tf_crvserver_unset"
		servicetype = "HTTP"
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to the documented NITRO defaults).
	}
`

func TestAccCrvserver_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCrvserverDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccCrvserver_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrvserverExist("citrixadc_crvserver.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_crvserver.tf_unset", "appflowlog", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_crvserver.tf_unset", "icmpvsrresponse", "ACTIVE"),
					resource.TestCheckResourceAttr("citrixadc_crvserver.tf_unset", "useportrange", "ON"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccCrvserver_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrvserverExist("citrixadc_crvserver.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_crvserver.tf_unset", "appflowlog", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_crvserver.tf_unset", "icmpvsrresponse", "PASSIVE"),
					resource.TestCheckResourceAttr("citrixadc_crvserver.tf_unset", "useportrange", "OFF"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckCrvserverADCValue("tf_crvserver_unset", "appflowlog", "ENABLED"),
					testAccCheckCrvserverADCValue("tf_crvserver_unset", "icmpvsrresponse", "PASSIVE"),
					testAccCheckCrvserverADCValue("tf_crvserver_unset", "useportrange", "OFF"),
				),
			},
		},
	})
}

// testAccCheckCrvserverADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckCrvserverADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Crvserver.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("crvserver %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("crvserver %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

const testAccCrvserverDataSource_basic = `

resource "citrixadc_crvserver" "tf_crvserver_ds" {
	name        = "my_vserver_ds"
	servicetype = "HTTP"
	arp         = "OFF"
}

data "citrixadc_crvserver" "tf_crvserver_ds" {
	name       = citrixadc_crvserver.tf_crvserver_ds.name
	depends_on = [citrixadc_crvserver.tf_crvserver_ds]
}

`
