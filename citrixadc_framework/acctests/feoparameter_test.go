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

const testAccFeoparameter_basic = `
	
resource "citrixadc_feoparameter" "tf_feoparameter" {
		jpegqualitypercent = 10
		cssinlinethressize = 100
		jsinlinethressize  = 50
		imginlinethressize = 1
	}
  
`

const testAccFeoparameter_update = `
	resource "citrixadc_feoparameter" "tf_feoparameter" {
		jpegqualitypercent = 0
		cssinlinethressize = 50
		jsinlinethressize  = 100
		imginlinethressize = 20
	}
  
`

func TestAccFeoparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccFeoparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFeoparameterExist("citrixadc_feoparameter.tf_feoparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_feoparameter.tf_feoparameter", "jpegqualitypercent", "10"),
					resource.TestCheckResourceAttr("citrixadc_feoparameter.tf_feoparameter", "cssinlinethressize", "100"),
					resource.TestCheckResourceAttr("citrixadc_feoparameter.tf_feoparameter", "jsinlinethressize", "50"),
					resource.TestCheckResourceAttr("citrixadc_feoparameter.tf_feoparameter", "imginlinethressize", "1"),
				),
			},
			{
				Config: testAccFeoparameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFeoparameterExist("citrixadc_feoparameter.tf_feoparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_feoparameter.tf_feoparameter", "jpegqualitypercent", "0"),
					resource.TestCheckResourceAttr("citrixadc_feoparameter.tf_feoparameter", "cssinlinethressize", "50"),
					resource.TestCheckResourceAttr("citrixadc_feoparameter.tf_feoparameter", "jsinlinethressize", "100"),
					resource.TestCheckResourceAttr("citrixadc_feoparameter.tf_feoparameter", "imginlinethressize", "20"),
				),
			},
		},
	})
}

func testAccCheckFeoparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No feoparameter name is set")
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
		data, err := client.FindResource("feoparameter", "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("feoparameter %s not found", n)
		}

		return nil
	}
}

func TestAccFeoparameter_import(t *testing.T) {
	const resAddr = "citrixadc_feoparameter.tf_feoparameter"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccFeoparameter_basic},
			{
				Config:                  testAccFeoparameter_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccFeoparameterDataSource_basic = `
	resource "citrixadc_feoparameter" "tf_feoparameter" {
		jpegqualitypercent = 10
		cssinlinethressize = 100
		jsinlinethressize  = 50
		imginlinethressize = 1
	}

	data "citrixadc_feoparameter" "feoparameter" {
		depends_on = [citrixadc_feoparameter.tf_feoparameter]
	}
`

func TestAccFeoparameter_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccFeoparameter_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckFeoparameterExist("citrixadc_feoparameter.tf_feoparameter", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccFeoparameter_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckFeoparameterExist("citrixadc_feoparameter.tf_feoparameter", nil)),
			},
		},
	})
}

// step1 sets all four unsettable attributes to non-default values; step2 removes
// them from config, so the provider must unset them and the appliance reverts to
// the documented NITRO defaults (jpegqualitypercent=75, the *inlinethressize=1024).
const testAccFeoparameter_unset_step1 = `
	resource "citrixadc_feoparameter" "tf_unset" {
		jpegqualitypercent = 10
		cssinlinethressize = 100
		jsinlinethressize  = 50
		imginlinethressize = 20
	}
`

const testAccFeoparameter_unset_step2 = `
	resource "citrixadc_feoparameter" "tf_unset" {
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccFeoparameter_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccFeoparameter_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFeoparameterExist("citrixadc_feoparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_feoparameter.tf_unset", "jpegqualitypercent", "10"),
					resource.TestCheckResourceAttr("citrixadc_feoparameter.tf_unset", "cssinlinethressize", "100"),
					resource.TestCheckResourceAttr("citrixadc_feoparameter.tf_unset", "jsinlinethressize", "50"),
					resource.TestCheckResourceAttr("citrixadc_feoparameter.tf_unset", "imginlinethressize", "20"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to NITRO defaults and the post-apply plan is empty.
				Config: testAccFeoparameter_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFeoparameterExist("citrixadc_feoparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_feoparameter.tf_unset", "jpegqualitypercent", "75"),
					resource.TestCheckResourceAttr("citrixadc_feoparameter.tf_unset", "cssinlinethressize", "1024"),
					resource.TestCheckResourceAttr("citrixadc_feoparameter.tf_unset", "jsinlinethressize", "1024"),
					resource.TestCheckResourceAttr("citrixadc_feoparameter.tf_unset", "imginlinethressize", "1024"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckFeoparameterADCValue("jpegqualitypercent", "75"),
					testAccCheckFeoparameterADCValue("cssinlinethressize", "1024"),
				),
			},
		},
	})
}

// testAccCheckFeoparameterADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckFeoparameterADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Feoparameter.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("feoparameter not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("feoparameter: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccFeoparameterDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccFeoparameterDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.citrixadc_feoparameter.feoparameter", "id"),
					resource.TestCheckResourceAttr("data.citrixadc_feoparameter.feoparameter", "jpegqualitypercent", "10"),
					resource.TestCheckResourceAttr("data.citrixadc_feoparameter.feoparameter", "cssinlinethressize", "100"),
					resource.TestCheckResourceAttr("data.citrixadc_feoparameter.feoparameter", "jsinlinethressize", "50"),
					resource.TestCheckResourceAttr("data.citrixadc_feoparameter.feoparameter", "imginlinethressize", "1"),
				),
			},
		},
	})
}
