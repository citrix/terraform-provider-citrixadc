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

const testAccLocationparameter_add = `
	resource "citrixadc_locationparameter" "tf_locationpara" {
		context            = "geographic"
		q1label            = "asia"
		matchwildcardtoany = "YES"
	}
`
const testAccLocationparameter_update = `
	resource "citrixadc_locationparameter" "tf_locationpara" {
		context            = "geographic"
		q1label            = "europe"
		matchwildcardtoany = "NO"
	}
`

const testAccLocationparameterDataSource_basic = `
	resource "citrixadc_locationparameter" "tf_locationpara" {
		context            = "geographic"
		q1label            = "asia"
		matchwildcardtoany = "YES"
	}

	data "citrixadc_locationparameter" "tf_locationpara" {
		depends_on = [citrixadc_locationparameter.tf_locationpara]
	}
`

func TestAccLocationparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccLocationparameter_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLocationparameterExist("citrixadc_locationparameter.tf_locationpara", nil),
					resource.TestCheckResourceAttr("citrixadc_locationparameter.tf_locationpara", "context", "geographic"),
					resource.TestCheckResourceAttr("citrixadc_locationparameter.tf_locationpara", "q1label", "asia"),
					resource.TestCheckResourceAttr("citrixadc_locationparameter.tf_locationpara", "matchwildcardtoany", "YES"),
				),
			},
			{
				Config: testAccLocationparameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLocationparameterExist("citrixadc_locationparameter.tf_locationpara", nil),
					resource.TestCheckResourceAttr("citrixadc_locationparameter.tf_locationpara", "context", "geographic"),
					resource.TestCheckResourceAttr("citrixadc_locationparameter.tf_locationpara", "matchwildcardtoany", "NO"),
				),
			},
		},
	})
}

func testAccCheckLocationparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No locationparameter name is set")
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
		data, err := client.FindResource(service.Locationparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("locationparameter %s not found", n)
		}

		return nil
	}
}

func TestAccLocationparameter_import(t *testing.T) {
	const resAddr = "citrixadc_locationparameter.tf_locationpara"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccLocationparameter_add},
			{
				Config:                  testAccLocationparameter_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccLocationparameter_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccLocationparameter_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLocationparameterExist("citrixadc_locationparameter.tf_locationpara", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccLocationparameter_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLocationparameterExist("citrixadc_locationparameter.tf_locationpara", nil)),
			},
		},
	})
}

// The locationparameter unset test covers matchwildcardtoany, the only
// mutable attribute with a documented NITRO server default (NO). Step 1 sets
// it to a non-default value; step 2 removes it from config, and the provider
// must unset it so the appliance reverts it to the default.
const testAccLocationparameter_unset_step1 = `
	resource "citrixadc_locationparameter" "tf_unset" {
		context            = "geographic"
		matchwildcardtoany = "YES"
	}
`

const testAccLocationparameter_unset_step2 = `
	resource "citrixadc_locationparameter" "tf_unset" {
		context = "geographic"
		# matchwildcardtoany removed from config -> provider must unset it
		# (revert to NITRO default, "NO").
	}
`

func TestAccLocationparameter_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccLocationparameter_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLocationparameterExist("citrixadc_locationparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_locationparameter.tf_unset", "matchwildcardtoany", "YES"),
				),
			},
			{
				// Removing the attribute must unset it: state (read back from the
				// appliance) reverts to the documented NITRO default, and the
				// implicit post-apply plan must be empty.
				Config: testAccLocationparameter_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLocationparameterExist("citrixadc_locationparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_locationparameter.tf_unset", "matchwildcardtoany", "NO"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckLocationparameterADCValue("matchwildcardtoany", "NO"),
				),
			},
		},
	})
}

// testAccCheckLocationparameterADCValue asserts an attribute's value directly
// on the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckLocationparameterADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Locationparameter.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("locationparameter not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("locationparameter: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccLocationparameterDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccLocationparameterDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_locationparameter.tf_locationpara", "context", "geographic"),
					resource.TestCheckResourceAttr("data.citrixadc_locationparameter.tf_locationpara", "q1label", "asia"),
					resource.TestCheckResourceAttr("data.citrixadc_locationparameter.tf_locationpara", "matchwildcardtoany", "YES"),
				),
			},
		},
	})
}
