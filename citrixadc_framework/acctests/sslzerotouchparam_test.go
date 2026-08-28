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

const testAccSslzerotouchparam_basic = `

	resource "citrixadc_sslzerotouchparam" "tf_sslzerotouchparam" {
		ocspcachetimeout      = 60
		ocspbatchingdepth     = 4
		ocspbatchingdelay     = 100
		ocsptrustresponder    = "YES"
		ocspusenonce          = "DISABLED"
		ocsphttpmethod        = "GET"
		ocspproducedattimeskew = 600
	}

`
const testAccSslzerotouchparam_update = `

	resource "citrixadc_sslzerotouchparam" "tf_sslzerotouchparam" {
		ocspcachetimeout      = 120
		ocspbatchingdepth     = 8
		ocspbatchingdelay     = 200
		ocsptrustresponder    = "NO"
		ocspusenonce          = "ENABLED"
		ocsphttpmethod        = "POST"
		ocspproducedattimeskew = 900
	}

`

func TestAccSslzerotouchparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSslzerotouchparam_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslzerotouchparamExist("citrixadc_sslzerotouchparam.tf_sslzerotouchparam", nil),
					resource.TestCheckResourceAttr("citrixadc_sslzerotouchparam.tf_sslzerotouchparam", "ocspcachetimeout", "60"),
					resource.TestCheckResourceAttr("citrixadc_sslzerotouchparam.tf_sslzerotouchparam", "ocspbatchingdepth", "4"),
					resource.TestCheckResourceAttr("citrixadc_sslzerotouchparam.tf_sslzerotouchparam", "ocspbatchingdelay", "100"),
					resource.TestCheckResourceAttr("citrixadc_sslzerotouchparam.tf_sslzerotouchparam", "ocsptrustresponder", "YES"),
					resource.TestCheckResourceAttr("citrixadc_sslzerotouchparam.tf_sslzerotouchparam", "ocspusenonce", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslzerotouchparam.tf_sslzerotouchparam", "ocsphttpmethod", "GET"),
					resource.TestCheckResourceAttr("citrixadc_sslzerotouchparam.tf_sslzerotouchparam", "ocspproducedattimeskew", "600"),
				),
			},
			{
				Config: testAccSslzerotouchparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslzerotouchparamExist("citrixadc_sslzerotouchparam.tf_sslzerotouchparam", nil),
					resource.TestCheckResourceAttr("citrixadc_sslzerotouchparam.tf_sslzerotouchparam", "ocspcachetimeout", "120"),
					resource.TestCheckResourceAttr("citrixadc_sslzerotouchparam.tf_sslzerotouchparam", "ocspbatchingdepth", "8"),
					resource.TestCheckResourceAttr("citrixadc_sslzerotouchparam.tf_sslzerotouchparam", "ocspbatchingdelay", "200"),
					resource.TestCheckResourceAttr("citrixadc_sslzerotouchparam.tf_sslzerotouchparam", "ocsptrustresponder", "NO"),
					resource.TestCheckResourceAttr("citrixadc_sslzerotouchparam.tf_sslzerotouchparam", "ocspusenonce", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslzerotouchparam.tf_sslzerotouchparam", "ocsphttpmethod", "POST"),
					resource.TestCheckResourceAttr("citrixadc_sslzerotouchparam.tf_sslzerotouchparam", "ocspproducedattimeskew", "900"),
				),
			},
		},
	})
}

func testAccCheckSslzerotouchparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslzerotouchparam name is set")
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
		data, err := client.FindResource(service.Sslzerotouchparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("sslzerotouchparam %s not found", n)
		}

		return nil
	}
}

const testAccSslzerotouchparamDataSource_basic = `

	resource "citrixadc_sslzerotouchparam" "tf_sslzerotouchparam" {
		ocspcachetimeout   = 60
		ocsptrustresponder = "YES"
		ocspusenonce       = "DISABLED"
		ocsphttpmethod     = "GET"
	}

	data "citrixadc_sslzerotouchparam" "tf_sslzerotouchparam" {
		depends_on = [citrixadc_sslzerotouchparam.tf_sslzerotouchparam]
	}
`

func TestAccSslzerotouchparamDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSslzerotouchparamDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslzerotouchparam.tf_sslzerotouchparam", "ocspcachetimeout", "60"),
					resource.TestCheckResourceAttr("data.citrixadc_sslzerotouchparam.tf_sslzerotouchparam", "ocsptrustresponder", "YES"),
					resource.TestCheckResourceAttr("data.citrixadc_sslzerotouchparam.tf_sslzerotouchparam", "ocspusenonce", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_sslzerotouchparam.tf_sslzerotouchparam", "ocsphttpmethod", "GET"),
				),
			},
		},
	})
}

func TestAccSslzerotouchparam_import(t *testing.T) {
	const resAddr = "citrixadc_sslzerotouchparam.tf_sslzerotouchparam"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccSslzerotouchparam_basic},
			{
				Config:                  testAccSslzerotouchparam_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

// sslzerotouchparam is a singleton. Step 1 sets a set of unset-eligible string
// toggle attributes to valid non-default values; step 2 removes them from config
// so the provider must unset them (revert to the documented NITRO defaults).
const testAccSslzerotouchparam_unset_step1 = `

	resource "citrixadc_sslzerotouchparam" "tf_sslzerotouchparam" {
		ocsptrustresponder = "YES"
		ocspusenonce       = "DISABLED"
		ocsphttpmethod     = "GET"
	}
`

const testAccSslzerotouchparam_unset_step2 = `

	resource "citrixadc_sslzerotouchparam" "tf_sslzerotouchparam" {
		# All unset-eligible string attributes removed from config -> the provider
		# must unset them (revert to NITRO defaults).
	}
`

func TestAccSslzerotouchparam_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccSslzerotouchparam_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslzerotouchparamExist("citrixadc_sslzerotouchparam.tf_sslzerotouchparam", nil),
					resource.TestCheckResourceAttr("citrixadc_sslzerotouchparam.tf_sslzerotouchparam", "ocsptrustresponder", "YES"),
					resource.TestCheckResourceAttr("citrixadc_sslzerotouchparam.tf_sslzerotouchparam", "ocspusenonce", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslzerotouchparam.tf_sslzerotouchparam", "ocsphttpmethod", "GET"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccSslzerotouchparam_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslzerotouchparamExist("citrixadc_sslzerotouchparam.tf_sslzerotouchparam", nil),
					resource.TestCheckResourceAttr("citrixadc_sslzerotouchparam.tf_sslzerotouchparam", "ocspusenonce", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslzerotouchparam.tf_sslzerotouchparam", "ocsphttpmethod", "POST"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckSslzerotouchparamADCValue("ocspusenonce", "ENABLED"),
					testAccCheckSslzerotouchparamADCValue("ocsphttpmethod", "POST"),
				),
			},
		},
	})
}

// testAccCheckSslzerotouchparamADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckSslzerotouchparamADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Sslzerotouchparam.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("sslzerotouchparam not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("sslzerotouchparam: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

// TestAccSslzerotouchparam_selfHealing verifies that if the singleton drifts on
// the appliance (an attribute changed out-of-band), the next apply reconciles it
// back to the configured value.
func TestAccSslzerotouchparam_selfHealing(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSslzerotouchparam_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslzerotouchparamExist("citrixadc_sslzerotouchparam.tf_sslzerotouchparam", nil),
				),
			},
			{
				// Mutate the resource out-of-band, then confirm the plan is
				// non-empty (drift detected) and re-apply restores config.
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("Failed to get test client: %v", err)
					}
					drift := map[string]interface{}{
						"ocsphttpmethod": "POST",
					}
					if err := client.UpdateUnnamedResource(service.Sslzerotouchparam.Type(), &drift); err != nil {
						t.Fatalf("Failed to drift sslzerotouchparam: %v", err)
					}
				},
				Config: testAccSslzerotouchparam_basic,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{plancheck.ExpectNonEmptyPlan()},
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslzerotouchparamExist("citrixadc_sslzerotouchparam.tf_sslzerotouchparam", nil),
					resource.TestCheckResourceAttr("citrixadc_sslzerotouchparam.tf_sslzerotouchparam", "ocsphttpmethod", "GET"),
				),
			},
		},
	})
}
