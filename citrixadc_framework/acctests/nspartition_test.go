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

const testAccNspartition_add = `
	resource "citrixadc_nspartition" "tf_nspartition" {
		partitionname = "tf_nspartition"
		maxbandwidth  = 1024
		minbandwidth  = 512
		maxconn       = 512
		maxmemlimit   = 11
	}
`
const testAccNspartition_update = `
	resource "citrixadc_nspartition" "tf_nspartition" {
		partitionname = "tf_nspartition"
		maxbandwidth  = 10240
		minbandwidth  = 512
		maxconn       = 512
		maxmemlimit   = 10
	}
`

func TestAccNspartition_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNspartitionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNspartition_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNspartitionExist("citrixadc_nspartition.tf_nspartition", nil),
					resource.TestCheckResourceAttr("citrixadc_nspartition.tf_nspartition", "partitionname", "tf_nspartition"),
					resource.TestCheckResourceAttr("citrixadc_nspartition.tf_nspartition", "maxbandwidth", "1024"),
					resource.TestCheckResourceAttr("citrixadc_nspartition.tf_nspartition", "maxmemlimit", "11"),
				),
			},
			{
				Config: testAccNspartition_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNspartitionExist("citrixadc_nspartition.tf_nspartition", nil),
					resource.TestCheckResourceAttr("citrixadc_nspartition.tf_nspartition", "partitionname", "tf_nspartition"),
					resource.TestCheckResourceAttr("citrixadc_nspartition.tf_nspartition", "maxbandwidth", "10240"),
					resource.TestCheckResourceAttr("citrixadc_nspartition.tf_nspartition", "maxmemlimit", "10"),
				),
			},
		},
	})
}

func testAccCheckNspartitionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nspartition name is set")
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
		data, err := client.FindResource(service.Nspartition.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nspartition %s not found", n)
		}

		return nil
	}
}

func testAccCheckNspartitionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nspartition" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Nspartition.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nspartition %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccNspartition_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_nspartition.tf_nspartition"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNspartitionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNspartition_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNspartitionExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Nspartition.Type(), "tf_nspartition"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccNspartition_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNspartitionExist(resAddr, nil)),
			},
		},
	})
}

func TestAccNspartition_import(t *testing.T) {
	const resAddr = "citrixadc_nspartition.tf_nspartition"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNspartitionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccNspartition_add},
			{
				Config:                  testAccNspartition_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"minbandwidth"},
			},
		},
	})
}

func TestAccNspartition_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNspartitionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccNspartition_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNspartitionExist("citrixadc_nspartition.tf_nspartition", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccNspartition_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNspartitionExist("citrixadc_nspartition.tf_nspartition", nil)),
			},
		},
	})
}

const testAccNspartitionDataSource_basic = `
	resource "citrixadc_nspartition" "tf_nspartition_ds" {
		partitionname = "tf_test_partition"
		maxbandwidth  = 10240
		minbandwidth  = 10240
		maxconn       = 1024
		maxmemlimit   = 10
	}

	data "citrixadc_nspartition" "tf_nspartition_ds" {
		partitionname = citrixadc_nspartition.tf_nspartition_ds.partitionname
	}
`

// nspartition unset test: step1 sets the unset-eligible numeric attributes to
// non-default values; step2 removes them from config so the provider must unset
// them back to the documented NITRO defaults (maxbandwidth=10240,
// minbandwidth=10240, maxconn=1024, maxmemlimit=10).
const testAccNspartition_unset_step1 = `
	resource "citrixadc_nspartition" "tf_unset" {
		partitionname = "tf_nspartition_unset"
		maxbandwidth  = 20480
		minbandwidth  = 2048
		maxconn       = 512
		maxmemlimit   = 20
	}
`

const testAccNspartition_unset_step2 = `
	resource "citrixadc_nspartition" "tf_unset" {
		partitionname = "tf_nspartition_unset"
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccNspartition_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNspartitionDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccNspartition_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNspartitionExist("citrixadc_nspartition.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nspartition.tf_unset", "maxbandwidth", "20480"),
					resource.TestCheckResourceAttr("citrixadc_nspartition.tf_unset", "minbandwidth", "2048"),
					resource.TestCheckResourceAttr("citrixadc_nspartition.tf_unset", "maxconn", "512"),
					resource.TestCheckResourceAttr("citrixadc_nspartition.tf_unset", "maxmemlimit", "20"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccNspartition_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNspartitionExist("citrixadc_nspartition.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nspartition.tf_unset", "maxbandwidth", "10240"),
					resource.TestCheckResourceAttr("citrixadc_nspartition.tf_unset", "minbandwidth", "10240"),
					resource.TestCheckResourceAttr("citrixadc_nspartition.tf_unset", "maxconn", "1024"),
					resource.TestCheckResourceAttr("citrixadc_nspartition.tf_unset", "maxmemlimit", "10"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckNspartitionADCValue("tf_nspartition_unset", "maxbandwidth", "10240"),
					testAccCheckNspartitionADCValue("tf_nspartition_unset", "maxconn", "1024"),
					testAccCheckNspartitionADCValue("tf_nspartition_unset", "maxmemlimit", "10"),
				),
			},
		},
	})
}

// testAccCheckNspartitionADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckNspartitionADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Nspartition.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("nspartition %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("nspartition %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccNspartitionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNspartitionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nspartition.tf_nspartition_ds", "partitionname", "tf_test_partition"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nspartition.tf_nspartition_ds", "id"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nspartition.tf_nspartition_ds", "maxbandwidth"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nspartition.tf_nspartition_ds", "maxconn"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nspartition.tf_nspartition_ds", "maxmemlimit"),
				),
			},
		},
	})
}
