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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccNslimitidentifier_add = `
	resource "citrixadc_nslimitidentifier" "tf_nslimitidentifier" {
		limitidentifier  = "tf_nslimitidentifier"
		threshold        = 1
		timeslice        = 1000
		limittype        = "BURSTY"
		mode             = "REQUEST_RATE"
		maxbandwidth     = 0
		trapsintimeslice = 1
	}
`
const testAccNslimitidentifier_update = `
	resource "citrixadc_nslimitidentifier" "tf_nslimitidentifier" {
		limitidentifier  = "tf_nslimitidentifier"
		threshold        = 2
		timeslice        = 2000
		limittype        = "BURSTY"
		mode             = "REQUEST_RATE"
		maxbandwidth     = 0
		trapsintimeslice = 1
	}
`

func TestAccNslimitidentifier_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNslimitidentifierDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNslimitidentifier_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNslimitidentifierExist("citrixadc_nslimitidentifier.tf_nslimitidentifier", nil),
					resource.TestCheckResourceAttr("citrixadc_nslimitidentifier.tf_nslimitidentifier", "limitidentifier", "tf_nslimitidentifier"),
					resource.TestCheckResourceAttr("citrixadc_nslimitidentifier.tf_nslimitidentifier", "threshold", "1"),
					resource.TestCheckResourceAttr("citrixadc_nslimitidentifier.tf_nslimitidentifier", "timeslice", "1000"),
				),
			},
			{
				Config: testAccNslimitidentifier_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNslimitidentifierExist("citrixadc_nslimitidentifier.tf_nslimitidentifier", nil),
					resource.TestCheckResourceAttr("citrixadc_nslimitidentifier.tf_nslimitidentifier", "limitidentifier", "tf_nslimitidentifier"),
					resource.TestCheckResourceAttr("citrixadc_nslimitidentifier.tf_nslimitidentifier", "threshold", "2"),
					resource.TestCheckResourceAttr("citrixadc_nslimitidentifier.tf_nslimitidentifier", "timeslice", "2000"),
				),
			},
		},
	})
}

func testAccCheckNslimitidentifierExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nslimitidentifier name is set")
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
		data, err := client.FindResource(service.Nslimitidentifier.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nslimitidentifier %s not found", n)
		}

		return nil
	}
}

func testAccCheckNslimitidentifierDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nslimitidentifier" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Nslimitidentifier.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nslimitidentifier %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccNslimitidentifier_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_nslimitidentifier.tf_nslimitidentifier"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNslimitidentifierDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNslimitidentifier_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNslimitidentifierExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Nslimitidentifier.Type(), "tf_nslimitidentifier"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccNslimitidentifier_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNslimitidentifierExist(resAddr, nil)),
			},
		},
	})
}

func TestAccNslimitidentifier_import(t *testing.T) {
	const resAddr = "citrixadc_nslimitidentifier.tf_nslimitidentifier"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNslimitidentifierDestroy,
		Steps: []resource.TestStep{
			{Config: testAccNslimitidentifier_add},
			{
				Config:                  testAccNslimitidentifier_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccNslimitidentifierDataSource_basic = `

	resource "citrixadc_nslimitidentifier" "tf_nslimitidentifier_ds" {
		limitidentifier  = "tf_nslimitidentifier_ds"
		threshold        = 5
		timeslice        = 3000
		limittype        = "BURSTY"
		mode             = "REQUEST_RATE"
		maxbandwidth     = 100
		trapsintimeslice = 2
	}

	data "citrixadc_nslimitidentifier" "tf_nslimitidentifier_ds_data" {
		limitidentifier = citrixadc_nslimitidentifier.tf_nslimitidentifier_ds.limitidentifier
	}
`

func TestAccNslimitidentifier_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNslimitidentifierDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccNslimitidentifier_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNslimitidentifierExist("citrixadc_nslimitidentifier.tf_nslimitidentifier", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccNslimitidentifier_add,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckNslimitidentifierExist("citrixadc_nslimitidentifier.tf_nslimitidentifier", nil)),
			},
		},
	})
}

// step1 sets the unset-eligible attributes to valid non-default values.
// mode is kept as REQUEST_RATE in both steps because threshold, timeslice and
// limittype are only meaningful in that mode (prerequisite for the appliance).
const testAccNslimitidentifier_unset_step1 = `
	resource "citrixadc_nslimitidentifier" "tf_unset" {
		limitidentifier  = "tf_nslimitidentifier_unset"
		mode             = "REQUEST_RATE"
		limittype        = "SMOOTH"
		threshold        = 50
		timeslice        = 2000
		maxbandwidth     = 100
		trapsintimeslice = 2
	}
`

// step2 removes all unset-eligible attributes -> the provider must unset them
// (revert to NITRO defaults).
const testAccNslimitidentifier_unset_step2 = `
	resource "citrixadc_nslimitidentifier" "tf_unset" {
		limitidentifier = "tf_nslimitidentifier_unset"
		mode            = "REQUEST_RATE"
	}
`

func TestAccNslimitidentifier_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNslimitidentifierDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccNslimitidentifier_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNslimitidentifierExist("citrixadc_nslimitidentifier.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nslimitidentifier.tf_unset", "limittype", "SMOOTH"),
					resource.TestCheckResourceAttr("citrixadc_nslimitidentifier.tf_unset", "threshold", "50"),
					resource.TestCheckResourceAttr("citrixadc_nslimitidentifier.tf_unset", "timeslice", "2000"),
					resource.TestCheckResourceAttr("citrixadc_nslimitidentifier.tf_unset", "maxbandwidth", "100"),
					resource.TestCheckResourceAttr("citrixadc_nslimitidentifier.tf_unset", "trapsintimeslice", "2"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccNslimitidentifier_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNslimitidentifierExist("citrixadc_nslimitidentifier.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nslimitidentifier.tf_unset", "limittype", "BURSTY"),
					resource.TestCheckResourceAttr("citrixadc_nslimitidentifier.tf_unset", "threshold", "1"),
					resource.TestCheckResourceAttr("citrixadc_nslimitidentifier.tf_unset", "timeslice", "1000"),
					resource.TestCheckResourceAttr("citrixadc_nslimitidentifier.tf_unset", "maxbandwidth", "0"),
					resource.TestCheckResourceAttr("citrixadc_nslimitidentifier.tf_unset", "trapsintimeslice", "0"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckNslimitidentifierADCValue("tf_nslimitidentifier_unset", "limittype", "BURSTY"),
					testAccCheckNslimitidentifierADCValue("tf_nslimitidentifier_unset", "threshold", "1"),
				),
			},
		},
	})
}

// testAccCheckNslimitidentifierADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckNslimitidentifierADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Nslimitidentifier.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("nslimitidentifier %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("nslimitidentifier %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccNslimitidentifierDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNslimitidentifierDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nslimitidentifier.tf_nslimitidentifier_ds_data", "limitidentifier", "tf_nslimitidentifier_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_nslimitidentifier.tf_nslimitidentifier_ds_data", "threshold", "5"),
					resource.TestCheckResourceAttr("data.citrixadc_nslimitidentifier.tf_nslimitidentifier_ds_data", "timeslice", "3000"),
					resource.TestCheckResourceAttr("data.citrixadc_nslimitidentifier.tf_nslimitidentifier_ds_data", "limittype", "BURSTY"),
					resource.TestCheckResourceAttr("data.citrixadc_nslimitidentifier.tf_nslimitidentifier_ds_data", "mode", "REQUEST_RATE"),
					resource.TestCheckResourceAttr("data.citrixadc_nslimitidentifier.tf_nslimitidentifier_ds_data", "maxbandwidth", "100"),
					resource.TestCheckResourceAttr("data.citrixadc_nslimitidentifier.tf_nslimitidentifier_ds_data", "trapsintimeslice", "2"),
				),
			},
		},
	})
}
