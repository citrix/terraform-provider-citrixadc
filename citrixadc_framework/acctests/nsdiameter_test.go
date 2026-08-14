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

const testAccNsdiameter_add = `
	resource "citrixadc_nsdiameter" "tf_nsdiameter" {
		identity               = "citrixadc.com"
		realm                  = "com"
		serverclosepropagation = "OFF"
	}
`
const testAccNsdiameter_update = `
	resource "citrixadc_nsdiameter" "tf_nsdiameter" {
		identity               = "netscaler.com"
		realm                  = "com"
		serverclosepropagation = "ON"
	}
`

func TestAccNsdiameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNsdiameter_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsdiameterExist("citrixadc_nsdiameter.tf_nsdiameter", nil),
					resource.TestCheckResourceAttr("citrixadc_nsdiameter.tf_nsdiameter", "identity", "citrixadc.com"),
					resource.TestCheckResourceAttr("citrixadc_nsdiameter.tf_nsdiameter", "realm", "com"),
					resource.TestCheckResourceAttr("citrixadc_nsdiameter.tf_nsdiameter", "serverclosepropagation", "OFF"),
				),
			},
			{
				Config: testAccNsdiameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsdiameterExist("citrixadc_nsdiameter.tf_nsdiameter", nil),
					resource.TestCheckResourceAttr("citrixadc_nsdiameter.tf_nsdiameter", "identity", "netscaler.com"),
					resource.TestCheckResourceAttr("citrixadc_nsdiameter.tf_nsdiameter", "realm", "com"),
					resource.TestCheckResourceAttr("citrixadc_nsdiameter.tf_nsdiameter", "serverclosepropagation", "ON"),
				),
			},
		},
	})
}

func testAccCheckNsdiameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsdiameter name is set")
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
		data, err := client.FindResource(service.Nsdiameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nsdiameter %s not found", n)
		}

		return nil
	}
}

func testAccCheckNsdiameterDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nsdiameter" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Nsdiameter.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nsdiameter %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccNsdiameter_import(t *testing.T) {
	const resAddr = "citrixadc_nsdiameter.tf_nsdiameter"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccNsdiameter_add},
			{
				Config:                  testAccNsdiameter_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"serverclosepropagation"},
			},
		},
	})
}

func TestAccNsdiameter_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNsdiameterDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccNsdiameter_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsdiameterExist("citrixadc_nsdiameter.tf_nsdiameter", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccNsdiameter_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsdiameterExist("citrixadc_nsdiameter.tf_nsdiameter", nil)),
			},
		},
	})
}

// nsdiameter unset coverage. serverclosepropagation is the only spec-unsettable,
// mutable attribute (the unset payload in the NITRO spec lists only
// serverclosepropagation and ownernode; ownernode is CLIP-only / the GET key and
// cannot be set on a standalone node, so it is excluded). Step1 sets a non-default
// value ("YES"); step2 removes it, and the provider must unset it (revert to the
// NITRO default "NO").
const testAccNsdiameter_unset_step1 = `
	resource "citrixadc_nsdiameter" "tf_unset" {
		identity               = "citrixadc.com"
		realm                  = "com"
		serverclosepropagation = "YES"
	}
`

const testAccNsdiameter_unset_step2 = `
	resource "citrixadc_nsdiameter" "tf_unset" {
		identity = "citrixadc.com"
		realm    = "com"
		# serverclosepropagation removed from config -> provider must unset it
		# (revert to NITRO default "NO").
	}
`

func TestAccNsdiameter_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccNsdiameter_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsdiameterExist("citrixadc_nsdiameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nsdiameter.tf_unset", "serverclosepropagation", "YES"),
				),
			},
			{
				// Removing the attribute must unset it: state (read back from the
				// appliance) reverts to the NITRO default, and the implicit
				// post-apply plan must be empty.
				Config: testAccNsdiameter_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsdiameterExist("citrixadc_nsdiameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nsdiameter.tf_unset", "serverclosepropagation", "NO"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckNsdiameterADCValue("serverclosepropagation", "NO"),
				),
			},
		},
	})
}

// testAccCheckNsdiameterADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckNsdiameterADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Nsdiameter.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("nsdiameter not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("nsdiameter: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

const testAccNsdiameterDataSource_basic = `
	resource "citrixadc_nsdiameter" "tf_nsdiameter_ds" {
		identity               = "citrixadc.com"
		realm                  = "com"
		serverclosepropagation = "NO"
	}

	data "citrixadc_nsdiameter" "tf_nsdiameter_ds" {
		ownernode = -1
		depends_on = [citrixadc_nsdiameter.tf_nsdiameter_ds]
	}
`

func TestAccNsdiameterDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNsdiameterDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nsdiameter.tf_nsdiameter_ds", "identity", "citrixadc.com"),
					resource.TestCheckResourceAttr("data.citrixadc_nsdiameter.tf_nsdiameter_ds", "realm", "com"),
					resource.TestCheckResourceAttr("data.citrixadc_nsdiameter.tf_nsdiameter_ds", "serverclosepropagation", "NO"),
				),
			},
		},
	})
}
