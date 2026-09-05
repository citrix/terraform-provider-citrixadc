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

const testAccTmtrafficaction_basic = `


	resource "citrixadc_tmtrafficaction" "tf_tmtrafficaction" {
		name             = "my_traffic_action"
		apptimeout       = 5
		sso              = "OFF"
		persistentcookie = "ON"
	}
`

const testAccTmtrafficaction_update = `


	resource "citrixadc_tmtrafficaction" "tf_tmtrafficaction" {
		name             = "my_traffic_action"
		apptimeout       = 10
		sso              = "ON"
		persistentcookie = "OFF"
	}
`

func TestAccTmtrafficaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTmtrafficactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTmtrafficaction_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmtrafficactionExist("citrixadc_tmtrafficaction.tf_tmtrafficaction", nil),
					resource.TestCheckResourceAttr("citrixadc_tmtrafficaction.tf_tmtrafficaction", "name", "my_traffic_action"),
					resource.TestCheckResourceAttr("citrixadc_tmtrafficaction.tf_tmtrafficaction", "apptimeout", "5"),
					resource.TestCheckResourceAttr("citrixadc_tmtrafficaction.tf_tmtrafficaction", "sso", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_tmtrafficaction.tf_tmtrafficaction", "persistentcookie", "ON"),
				),
			},
			{
				Config: testAccTmtrafficaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmtrafficactionExist("citrixadc_tmtrafficaction.tf_tmtrafficaction", nil),
					resource.TestCheckResourceAttr("citrixadc_tmtrafficaction.tf_tmtrafficaction", "name", "my_traffic_action"),
					resource.TestCheckResourceAttr("citrixadc_tmtrafficaction.tf_tmtrafficaction", "apptimeout", "10"),
					resource.TestCheckResourceAttr("citrixadc_tmtrafficaction.tf_tmtrafficaction", "sso", "ON"),
					resource.TestCheckResourceAttr("citrixadc_tmtrafficaction.tf_tmtrafficaction", "persistentcookie", "OFF"),
				),
			},
		},
	})
}

func testAccCheckTmtrafficactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No tmtrafficaction name is set")
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
		data, err := client.FindResource(service.Tmtrafficaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("tmtrafficaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckTmtrafficactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_tmtrafficaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Tmtrafficaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("tmtrafficaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccTmtrafficaction_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_tmtrafficaction.tf_tmtrafficaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTmtrafficactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTmtrafficaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTmtrafficactionExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Tmtrafficaction.Type(), "my_traffic_action"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccTmtrafficaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTmtrafficactionExist(resAddr, nil)),
			},
		},
	})
}

func TestAccTmtrafficaction_import(t *testing.T) {
	const resAddr = "citrixadc_tmtrafficaction.tf_tmtrafficaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTmtrafficactionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccTmtrafficaction_basic},
			{
				Config:                  testAccTmtrafficaction_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccTmtrafficaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckTmtrafficactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccTmtrafficaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTmtrafficactionExist("citrixadc_tmtrafficaction.tf_tmtrafficaction", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccTmtrafficaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTmtrafficactionExist("citrixadc_tmtrafficaction.tf_tmtrafficaction", nil)),
			},
		},
	})
}

// testAccTmtrafficaction_unset_step1 sets the unset-eligible attributes to
// valid non-default values; step2 removes them so the provider must unset them
// (revert to the documented NITRO defaults).
const testAccTmtrafficaction_unset_step1 = `
	resource "citrixadc_tmtrafficaction" "tf_unset" {
		name             = "tf_test_tmtrafficaction_unset"
		apptimeout       = 5
		persistentcookie = "ON"
		userexpression   = "http.req.user.name"
		passwdexpression = "http.req.user.passwd"
	}
`

const testAccTmtrafficaction_unset_step2 = `
	resource "citrixadc_tmtrafficaction" "tf_unset" {
		name       = "tf_test_tmtrafficaction_unset"
		apptimeout = 5
		# persistentcookie, userexpression and passwdexpression removed from
		# config -> the provider must unset them (revert to NITRO defaults).
	}
`

func TestAccTmtrafficaction_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTmtrafficactionDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccTmtrafficaction_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmtrafficactionExist("citrixadc_tmtrafficaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_tmtrafficaction.tf_unset", "persistentcookie", "ON"),
					resource.TestCheckResourceAttr("citrixadc_tmtrafficaction.tf_unset", "userexpression", "http.req.user.name"),
					resource.TestCheckResourceAttr("citrixadc_tmtrafficaction.tf_unset", "passwdexpression", "http.req.user.passwd"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults and the
				// implicit post-apply plan must be empty.
				Config: testAccTmtrafficaction_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmtrafficactionExist("citrixadc_tmtrafficaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_tmtrafficaction.tf_unset", "persistentcookie", "OFF"),
					// After Option B (unsetOnRemove) these converted attrs read back
					// NULL: NITRO omits them from GET and there is no Default to inject.
					resource.TestCheckNoResourceAttr("citrixadc_tmtrafficaction.tf_unset", "userexpression"),
					resource.TestCheckNoResourceAttr("citrixadc_tmtrafficaction.tf_unset", "passwdexpression"),
					// Independent appliance-level confirmation the unset took effect
					// (unset attributes are omitted from the NITRO GET response).
					testAccCheckTmtrafficactionADCUnset("tf_test_tmtrafficaction_unset", "persistentcookie"),
					testAccCheckTmtrafficactionADCUnset("tf_test_tmtrafficaction_unset", "userexpression"),
				),
			},
		},
	})
}

// testAccCheckTmtrafficactionADCUnset asserts that an attribute is absent from
// the appliance's GET response (unset attributes revert to defaults and are
// omitted), proving the unset actually took effect.
func testAccCheckTmtrafficactionADCUnset(name, attr string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Tmtrafficaction.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("tmtrafficaction %s not found on appliance", name)
		}
		if val, ok := data[attr]; ok && val != nil && strings.TrimSpace(fmt.Sprintf("%v", val)) != "" {
			return fmt.Errorf("tmtrafficaction %s: appliance attr %q = %q, want unset/absent", name, attr, val)
		}
		return nil
	}
}

const testAccTmtrafficactionDataSource_basic = `


	resource "citrixadc_tmtrafficaction" "tf_tmtrafficaction" {
		name             = "my_traffic_action"
		apptimeout       = 5
		sso              = "OFF"
		persistentcookie = "ON"
	}

data "citrixadc_tmtrafficaction" "tf_tmtrafficaction" {
    name = citrixadc_tmtrafficaction.tf_tmtrafficaction.name
}
`

func TestAccTmtrafficactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTmtrafficactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_tmtrafficaction.tf_tmtrafficaction", "name", "my_traffic_action"),
					resource.TestCheckResourceAttr("data.citrixadc_tmtrafficaction.tf_tmtrafficaction", "apptimeout", "5"),
					resource.TestCheckResourceAttr("data.citrixadc_tmtrafficaction.tf_tmtrafficaction", "sso", "OFF"),
					resource.TestCheckResourceAttr("data.citrixadc_tmtrafficaction.tf_tmtrafficaction", "persistentcookie", "ON"),
				),
			},
		},
	})
}
