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

const testAccSystemautosaveparam_basic = `

	resource "citrixadc_systemautosaveparam" "tf_systemautosaveparam" {
		status                = "ENABLED"
		periodicsave          = "ENABLED"
		periodicsavefrequency = 1440
	}

`
const testAccSystemautosaveparam_update = `

	resource "citrixadc_systemautosaveparam" "tf_systemautosaveparam" {
		status                = "DISABLED"
		periodicsave          = "ENABLED"
		periodicsavefrequency = 720
	}

`

func TestAccSystemautosaveparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemautosaveparamDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemautosaveparam_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemautosaveparamExist("citrixadc_systemautosaveparam.tf_systemautosaveparam", nil),
					resource.TestCheckResourceAttr("citrixadc_systemautosaveparam.tf_systemautosaveparam", "status", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemautosaveparam.tf_systemautosaveparam", "periodicsave", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemautosaveparam.tf_systemautosaveparam", "periodicsavefrequency", "1440"),
					// Independent appliance-level confirmation.
					testAccCheckSystemautosaveparamADCValue("status", "ENABLED"),
					testAccCheckSystemautosaveparamADCValue("periodicsave", "ENABLED"),
					testAccCheckSystemautosaveparamADCValue("periodicsavefrequency", "1440"),
				),
			},
			{
				Config: testAccSystemautosaveparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemautosaveparamExist("citrixadc_systemautosaveparam.tf_systemautosaveparam", nil),
					resource.TestCheckResourceAttr("citrixadc_systemautosaveparam.tf_systemautosaveparam", "status", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemautosaveparam.tf_systemautosaveparam", "periodicsave", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemautosaveparam.tf_systemautosaveparam", "periodicsavefrequency", "720"),
					testAccCheckSystemautosaveparamADCValue("status", "DISABLED"),
					testAccCheckSystemautosaveparamADCValue("periodicsavefrequency", "720"),
				),
			},
		},
	})
}

func testAccCheckSystemautosaveparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No systemautosaveparam id is set")
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
		data, err := client.FindResource(service.Systemautosaveparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("systemautosaveparam %s not found", n)
		}

		return nil
	}
}

// systemautosaveparam is a global configuration singleton with no NITRO delete
// operation; there is nothing to assert on destroy.
func testAccCheckSystemautosaveparamDestroy(s *terraform.State) error {
	return nil
}

// testAccCheckSystemautosaveparamADCValue asserts an attribute's value directly
// on the appliance (not just in Terraform state). A missing key is treated as an
// empty value, which is how the appliance reports an unset attribute.
func testAccCheckSystemautosaveparamADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Systemautosaveparam.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("systemautosaveparam not found on appliance")
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("systemautosaveparam: appliance attr %q = %q, want %q", attr, got, want)
		}
		return nil
	}
}

const testAccSystemautosaveparamDataSource_basic = `

	resource "citrixadc_systemautosaveparam" "tf_systemautosaveparam" {
		status                = "ENABLED"
		periodicsave          = "ENABLED"
		periodicsavefrequency = 1440
	}

	data "citrixadc_systemautosaveparam" "tf_systemautosaveparam" {
		depends_on = [citrixadc_systemautosaveparam.tf_systemautosaveparam]
	}
`

func TestAccSystemautosaveparamDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemautosaveparamDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemautosaveparamDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_systemautosaveparam.tf_systemautosaveparam", "status", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_systemautosaveparam.tf_systemautosaveparam", "periodicsave", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_systemautosaveparam.tf_systemautosaveparam", "periodicsavefrequency", "1440"),
				),
			},
		},
	})
}

func TestAccSystemautosaveparam_import(t *testing.T) {
	const resAddr = "citrixadc_systemautosaveparam.tf_systemautosaveparam"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemautosaveparamDestroy,
		Steps: []resource.TestStep{
			{Config: testAccSystemautosaveparam_basic},
			{
				Config:            testAccSystemautosaveparam_basic,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// systemautosaveparam is a singleton. Step 1 sets the unset-eligible attributes
// (status, periodicsave, periodicsavefrequency) to non-default values; step 2
// removes them from config so the provider must unset them (revert to the NITRO
// defaults: status=DEFAULT, periodicsave=DISABLED, periodicsavefrequency=720).
const testAccSystemautosaveparam_unset_step1 = `

	resource "citrixadc_systemautosaveparam" "tf_systemautosaveparam" {
		status                = "ENABLED"
		periodicsave          = "ENABLED"
		periodicsavefrequency = 1440
	}
`

const testAccSystemautosaveparam_unset_step2 = `

	resource "citrixadc_systemautosaveparam" "tf_systemautosaveparam" {
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccSystemautosaveparam_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemautosaveparamDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccSystemautosaveparam_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemautosaveparamExist("citrixadc_systemautosaveparam.tf_systemautosaveparam", nil),
					resource.TestCheckResourceAttr("citrixadc_systemautosaveparam.tf_systemautosaveparam", "status", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemautosaveparam.tf_systemautosaveparam", "periodicsave", "ENABLED"),
					testAccCheckSystemautosaveparamADCValue("status", "ENABLED"),
					testAccCheckSystemautosaveparamADCValue("periodicsave", "ENABLED"),
				),
			},
			{
				// Removing the attributes must unset them: the appliance reverts
				// them to their defaults, and the implicit post-apply plan must be
				// empty.
				Config: testAccSystemautosaveparam_unset_step2,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemautosaveparamExist("citrixadc_systemautosaveparam.tf_systemautosaveparam", nil),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckSystemautosaveparamADCValue("status", "DEFAULT"),
					testAccCheckSystemautosaveparamADCValue("periodicsave", "DISABLED"),
				),
			},
		},
	})
}
