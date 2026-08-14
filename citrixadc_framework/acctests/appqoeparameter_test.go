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

const testAccAppqoeparameter_basic = `

	resource "citrixadc_appqoeparameter" "tf_appqoeparameter" {
		sessionlife         = 300
		avgwaitingclient    = 400
		maxaltrespbandwidth = 50
		dosattackthresh     = 100
	}
`
const testAccAppqoeparameter_update = `

	resource "citrixadc_appqoeparameter" "tf_appqoeparameter" {
		sessionlife         = 400
		avgwaitingclient    = 300
		maxaltrespbandwidth = 100
		dosattackthresh     = 50
	}
`

func TestAccAppqoeparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAppqoeparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppqoeparameterExist("citrixadc_appqoeparameter.tf_appqoeparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_appqoeparameter.tf_appqoeparameter", "sessionlife", "300"),
					resource.TestCheckResourceAttr("citrixadc_appqoeparameter.tf_appqoeparameter", "avgwaitingclient", "400"),
					resource.TestCheckResourceAttr("citrixadc_appqoeparameter.tf_appqoeparameter", "maxaltrespbandwidth", "50"),
					resource.TestCheckResourceAttr("citrixadc_appqoeparameter.tf_appqoeparameter", "dosattackthresh", "100"),
				),
			},
			{
				Config: testAccAppqoeparameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppqoeparameterExist("citrixadc_appqoeparameter.tf_appqoeparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_appqoeparameter.tf_appqoeparameter", "sessionlife", "400"),
					resource.TestCheckResourceAttr("citrixadc_appqoeparameter.tf_appqoeparameter", "avgwaitingclient", "300"),
					resource.TestCheckResourceAttr("citrixadc_appqoeparameter.tf_appqoeparameter", "maxaltrespbandwidth", "100"),
					resource.TestCheckResourceAttr("citrixadc_appqoeparameter.tf_appqoeparameter", "dosattackthresh", "50"),
				),
			},
		},
	})
}

func testAccCheckAppqoeparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appqoeparameter name is set")
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
		data, err := client.FindResource(service.Appqoeparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appqoeparameter %s not found", n)
		}

		return nil
	}
}

// appqoeparameter is a singleton. The unset test sets all four mutable
// attributes to non-default values, then removes them from config; the provider
// must unset them so the appliance reverts to the documented NITRO defaults.
const testAccAppqoeparameter_unset_step1 = `
	resource "citrixadc_appqoeparameter" "tf_unset" {
		sessionlife         = 400
		avgwaitingclient    = 500
		maxaltrespbandwidth = 200
		dosattackthresh     = 1000
	}
`

const testAccAppqoeparameter_unset_step2 = `
	resource "citrixadc_appqoeparameter" "tf_unset" {
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccAppqoeparameter_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccAppqoeparameter_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppqoeparameterExist("citrixadc_appqoeparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_appqoeparameter.tf_unset", "sessionlife", "400"),
					resource.TestCheckResourceAttr("citrixadc_appqoeparameter.tf_unset", "avgwaitingclient", "500"),
					resource.TestCheckResourceAttr("citrixadc_appqoeparameter.tf_unset", "maxaltrespbandwidth", "200"),
					resource.TestCheckResourceAttr("citrixadc_appqoeparameter.tf_unset", "dosattackthresh", "1000"),
				),
			},
			{
				// Removing the attributes must unset them: state reverts to the
				// documented NITRO defaults, and the implicit post-apply plan is empty.
				Config: testAccAppqoeparameter_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppqoeparameterExist("citrixadc_appqoeparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_appqoeparameter.tf_unset", "sessionlife", "300"),
					resource.TestCheckResourceAttr("citrixadc_appqoeparameter.tf_unset", "avgwaitingclient", "1000000"),
					resource.TestCheckResourceAttr("citrixadc_appqoeparameter.tf_unset", "maxaltrespbandwidth", "100"),
					resource.TestCheckResourceAttr("citrixadc_appqoeparameter.tf_unset", "dosattackthresh", "2000"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAppqoeparameterADCValue("sessionlife", "300"),
					testAccCheckAppqoeparameterADCValue("dosattackthresh", "2000"),
				),
			},
		},
	})
}

// testAccCheckAppqoeparameterADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckAppqoeparameterADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Appqoeparameter.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("appqoeparameter not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("appqoeparameter: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccAppqoeparameter_import(t *testing.T) {
	const resAddr = "citrixadc_appqoeparameter.tf_appqoeparameter"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccAppqoeparameter_basic},
			{
				Config:                  testAccAppqoeparameter_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccAppqoeparameter_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccAppqoeparameter_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppqoeparameterExist("citrixadc_appqoeparameter.tf_appqoeparameter", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccAppqoeparameter_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppqoeparameterExist("citrixadc_appqoeparameter.tf_appqoeparameter", nil)),
			},
		},
	})
}

const testAccAppqoeparameterDataSource_basic = `

	resource "citrixadc_appqoeparameter" "tf_appqoeparameter" {
		sessionlife         = 300
		avgwaitingclient    = 400
		maxaltrespbandwidth = 50
		dosattackthresh     = 100
	}

	data "citrixadc_appqoeparameter" "tf_appqoeparameter" {
		depends_on = [citrixadc_appqoeparameter.tf_appqoeparameter]
	}
`

func TestAccAppqoeparameterDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAppqoeparameterDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appqoeparameter.tf_appqoeparameter", "sessionlife", "300"),
					resource.TestCheckResourceAttr("data.citrixadc_appqoeparameter.tf_appqoeparameter", "avgwaitingclient", "400"),
					resource.TestCheckResourceAttr("data.citrixadc_appqoeparameter.tf_appqoeparameter", "maxaltrespbandwidth", "50"),
					resource.TestCheckResourceAttr("data.citrixadc_appqoeparameter.tf_appqoeparameter", "dosattackthresh", "100"),
				),
			},
		},
	})
}
