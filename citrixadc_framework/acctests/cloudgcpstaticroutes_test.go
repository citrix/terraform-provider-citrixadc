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
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// needsGCPFixtureSkip documents why cloudgcpstaticroutes acc tests are guarded.
// The resource pushes static routes to GCP and enabling it (status = ENABLED,
// non-empty project) only succeeds on an ADC actually running in Google Cloud.
// The tests below use only enum/length-valid values from the NITRO doc, but they
// are skipped until run against a GCP-hosted appliance.
const needsGCPFixtureSkip = "needs GCP-hosted ADC fixture: cloudgcpstaticroutes pushes static routes to GCP; enabling requires a GCP environment"

const testAccCloudgcpstaticroutes_basic_step1 = `

	resource "citrixadc_cloudgcpstaticroutes" "tf_cloudgcpstaticroutes" {
		status  = "ENABLED"
		project = "tf-acc-project-1"
	}
`

const testAccCloudgcpstaticroutes_basic_step2 = `

	resource "citrixadc_cloudgcpstaticroutes" "tf_cloudgcpstaticroutes" {
		status  = "DISABLED"
		project = "tf-acc-project-2"
	}
`

func TestAccCloudgcpstaticroutes_basic(t *testing.T) {
	t.Skip(needsGCPFixtureSkip)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// Singleton resource: it always exists on the ADC and is never deleted,
		// so no CheckDestroy is used.
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudgcpstaticroutes_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudgcpstaticroutesExist("citrixadc_cloudgcpstaticroutes.tf_cloudgcpstaticroutes", nil),
					resource.TestCheckResourceAttr("citrixadc_cloudgcpstaticroutes.tf_cloudgcpstaticroutes", "status", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_cloudgcpstaticroutes.tf_cloudgcpstaticroutes", "project", "tf-acc-project-1"),
				),
			},
			{
				Config: testAccCloudgcpstaticroutes_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudgcpstaticroutesExist("citrixadc_cloudgcpstaticroutes.tf_cloudgcpstaticroutes", nil),
					resource.TestCheckResourceAttr("citrixadc_cloudgcpstaticroutes.tf_cloudgcpstaticroutes", "status", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_cloudgcpstaticroutes.tf_cloudgcpstaticroutes", "project", "tf-acc-project-2"),
				),
			},
		},
	})
}

func TestAccCloudgcpstaticroutes_import(t *testing.T) {
	t.Skip(needsGCPFixtureSkip)
	const resAddr = "citrixadc_cloudgcpstaticroutes.tf_cloudgcpstaticroutes"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// Singleton resource: it always exists on the ADC and is never deleted,
		// so no CheckDestroy is used.
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{Config: testAccCloudgcpstaticroutes_basic_step1},
			{
				Config:                  testAccCloudgcpstaticroutes_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccCloudgcpstaticroutes_selfHealing(t *testing.T) {
	t.Skip(needsGCPFixtureSkip)
	const resAddr = "citrixadc_cloudgcpstaticroutes.tf_cloudgcpstaticroutes"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudgcpstaticroutes_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudgcpstaticroutesExist(resAddr, nil),
				),
			},
			{
				// Out-of-band unset of the tracked attributes should be
				// re-reconciled by the next apply of the same config.
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("Failed to get test client: %v", err)
					}
					if err := client.ActOnResource(service.Cloudgcpstaticroutes.Type(), map[string]interface{}{"status": true, "project": true}, "unset"); err != nil {
						t.Fatalf("Failed to unset cloudgcpstaticroutes out-of-band: %v", err)
					}
				},
				Config: testAccCloudgcpstaticroutes_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudgcpstaticroutesExist(resAddr, nil),
					resource.TestCheckResourceAttr(resAddr, "status", "ENABLED"),
					resource.TestCheckResourceAttr(resAddr, "project", "tf-acc-project-1"),
				),
			},
		},
	})
}

func testAccCheckCloudgcpstaticroutesExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cloudgcpstaticroutes name is set")
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
		// Singleton set-get resource: read with an empty name.
		data, err := client.FindResource(service.Cloudgcpstaticroutes.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("cloudgcpstaticroutes %s not found", n)
		}

		return nil
	}
}

const testAccCloudgcpstaticroutesDataSource_basic = `
	resource "citrixadc_cloudgcpstaticroutes" "tf_cloudgcpstaticroutes" {
		status  = "ENABLED"
		project = "tf-acc-project-1"
	}

	data "citrixadc_cloudgcpstaticroutes" "tf_cloudgcpstaticroutes" {
		depends_on = [citrixadc_cloudgcpstaticroutes.tf_cloudgcpstaticroutes]
	}
`

func TestAccCloudgcpstaticroutesDataSource_basic(t *testing.T) {
	t.Skip(needsGCPFixtureSkip)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudgcpstaticroutesDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_cloudgcpstaticroutes.tf_cloudgcpstaticroutes", "status", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_cloudgcpstaticroutes.tf_cloudgcpstaticroutes", "project", "tf-acc-project-1"),
				),
			},
		},
	})
}

// Step 1: both unset-eligible attributes set to non-default values.
const testAccCloudgcpstaticroutes_unset_step1 = `
	resource "citrixadc_cloudgcpstaticroutes" "tf_unset" {
		status  = "ENABLED"
		project = "tf-acc-project-1"
	}
`

// Step 2: eligible attributes removed from config -> provider must unset them,
// reverting each to its NITRO default.
const testAccCloudgcpstaticroutes_unset_step2 = `
	resource "citrixadc_cloudgcpstaticroutes" "tf_unset" {
	}
`

func TestAccCloudgcpstaticroutes_unset(t *testing.T) {
	t.Skip(needsGCPFixtureSkip)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// Singleton resource: it always exists on the ADC and is never deleted,
		// so no CheckDestroy is used.
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				// Non-default values apply and persist.
				Config: testAccCloudgcpstaticroutes_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudgcpstaticroutesExist("citrixadc_cloudgcpstaticroutes.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_cloudgcpstaticroutes.tf_unset", "status", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_cloudgcpstaticroutes.tf_unset", "project", "tf-acc-project-1"),
				),
			},
			{
				// Removing them must unset -> state reverts to NITRO defaults,
				// and the implicit post-apply plan must be empty.
				Config: testAccCloudgcpstaticroutes_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudgcpstaticroutesExist("citrixadc_cloudgcpstaticroutes.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_cloudgcpstaticroutes.tf_unset", "status", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckCloudgcpstaticroutesADCValue("status", "DISABLED"),
				),
			},
		},
	})
}

// testAccCheckCloudgcpstaticroutesADCValue asserts an attribute's value directly
// on the appliance (not just in Terraform state), proving the unset actually
// reverted it. cloudgcpstaticroutes is a singleton, so it is read with an empty
// name.
func testAccCheckCloudgcpstaticroutesADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Cloudgcpstaticroutes.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("cloudgcpstaticroutes not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("cloudgcpstaticroutes: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}
