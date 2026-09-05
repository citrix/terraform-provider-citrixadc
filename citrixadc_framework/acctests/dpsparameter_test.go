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

const testAccDpsparameter_basic = `

	resource "citrixadc_dpsparameter" "tf_dpsparameter" {
		customerid = "customer123"
		deployment = "COMMERCIAL"
		serviceurl = "https://example.citrixcloud.net"
	}

`
const testAccDpsparameter_update = `

	resource "citrixadc_dpsparameter" "tf_dpsparameter" {
		customerid = "customer456"
		deployment = "GOV"
		serviceurl = "https://example-gov.citrixcloud.net"
	}

`

func TestAccDpsparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDpsparameterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDpsparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDpsparameterExist("citrixadc_dpsparameter.tf_dpsparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_dpsparameter.tf_dpsparameter", "customerid", "customer123"),
					resource.TestCheckResourceAttr("citrixadc_dpsparameter.tf_dpsparameter", "deployment", "COMMERCIAL"),
					resource.TestCheckResourceAttr("citrixadc_dpsparameter.tf_dpsparameter", "serviceurl", "https://example.citrixcloud.net"),
					// Independent appliance-level confirmation.
					testAccCheckDpsparameterADCValue("customerid", "customer123"),
					testAccCheckDpsparameterADCValue("deployment", "COMMERCIAL"),
					testAccCheckDpsparameterADCValue("serviceurl", "https://example.citrixcloud.net"),
				),
			},
			{
				Config: testAccDpsparameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDpsparameterExist("citrixadc_dpsparameter.tf_dpsparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_dpsparameter.tf_dpsparameter", "customerid", "customer456"),
					resource.TestCheckResourceAttr("citrixadc_dpsparameter.tf_dpsparameter", "deployment", "GOV"),
					resource.TestCheckResourceAttr("citrixadc_dpsparameter.tf_dpsparameter", "serviceurl", "https://example-gov.citrixcloud.net"),
					testAccCheckDpsparameterADCValue("customerid", "customer456"),
					testAccCheckDpsparameterADCValue("deployment", "GOV"),
					testAccCheckDpsparameterADCValue("serviceurl", "https://example-gov.citrixcloud.net"),
				),
			},
		},
	})
}

func testAccCheckDpsparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dpsparameter id is set")
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
		data, err := client.FindResource(service.Dpsparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("dpsparameter %s not found", n)
		}

		return nil
	}
}

// dpsparameter is a global configuration singleton with no NITRO delete
// operation; there is nothing to assert on destroy.
func testAccCheckDpsparameterDestroy(s *terraform.State) error {
	return nil
}

// testAccCheckDpsparameterADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state). A missing key is treated as an empty
// value, which is how the appliance reports an unset attribute.
func testAccCheckDpsparameterADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Dpsparameter.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("dpsparameter not found on appliance")
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("dpsparameter: appliance attr %q = %q, want %q", attr, got, want)
		}
		return nil
	}
}

const testAccDpsparameterDataSource_basic = `

	resource "citrixadc_dpsparameter" "tf_dpsparameter" {
		customerid = "customer123"
		deployment = "COMMERCIAL"
		serviceurl = "https://example.citrixcloud.net"
	}

	data "citrixadc_dpsparameter" "tf_dpsparameter" {
		depends_on = [citrixadc_dpsparameter.tf_dpsparameter]
	}
`

func TestAccDpsparameterDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDpsparameterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDpsparameterDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_dpsparameter.tf_dpsparameter", "customerid", "customer123"),
					resource.TestCheckResourceAttr("data.citrixadc_dpsparameter.tf_dpsparameter", "deployment", "COMMERCIAL"),
					resource.TestCheckResourceAttr("data.citrixadc_dpsparameter.tf_dpsparameter", "serviceurl", "https://example.citrixcloud.net"),
				),
			},
		},
	})
}

func TestAccDpsparameter_import(t *testing.T) {
	const resAddr = "citrixadc_dpsparameter.tf_dpsparameter"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDpsparameterDestroy,
		Steps: []resource.TestStep{
			{Config: testAccDpsparameter_basic},
			{
				Config:            testAccDpsparameter_basic,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// dpsparameter is a singleton. Step 1 sets all three unset-eligible attributes
// (customerid, deployment, serviceurl) to non-default values; step 2 removes
// them from config so the provider must unset them (revert to NITRO defaults).
const testAccDpsparameter_unset_step1 = `

	resource "citrixadc_dpsparameter" "tf_dpsparameter" {
		customerid = "customer123"
		deployment = "COMMERCIAL"
		serviceurl = "https://example.citrixcloud.net"
	}
`

const testAccDpsparameter_unset_step2 = `

	resource "citrixadc_dpsparameter" "tf_dpsparameter" {
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccDpsparameter_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDpsparameterDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccDpsparameter_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDpsparameterExist("citrixadc_dpsparameter.tf_dpsparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_dpsparameter.tf_dpsparameter", "customerid", "customer123"),
					resource.TestCheckResourceAttr("citrixadc_dpsparameter.tf_dpsparameter", "serviceurl", "https://example.citrixcloud.net"),
					testAccCheckDpsparameterADCValue("customerid", "customer123"),
					testAccCheckDpsparameterADCValue("serviceurl", "https://example.citrixcloud.net"),
				),
			},
			{
				// Removing the attributes must unset them: the appliance reverts
				// them to their defaults, and the implicit post-apply plan must be
				// empty.
				Config: testAccDpsparameter_unset_step2,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDpsparameterExist("citrixadc_dpsparameter.tf_dpsparameter", nil),
					// Independent appliance-level confirmation the revert-to-default
					// took effect. Each attribute reverts to its KNOWN non-null NITRO
					// default (verified on the appliance): customerid="None",
					// deployment="COMMERCIAL",
					// serviceurl="https://device-posture-controller.cloud.com".
					testAccCheckDpsparameterADCValue("customerid", "None"),
					testAccCheckDpsparameterADCValue("deployment", "COMMERCIAL"),
					testAccCheckDpsparameterADCValue("serviceurl", "https://device-posture-controller.cloud.com"),
				),
			},
		},
	})
}
