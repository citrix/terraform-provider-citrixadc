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

const testAccNstcpbufparam_add = `
	resource "citrixadc_nstcpbufparam" "tf_nstcpbufparam" {
		size     = 32
		memlimit = 8
	}
`
const testAccNstcpbufparam_update = `
	resource "citrixadc_nstcpbufparam" "tf_nstcpbufparam" {
		size     = 64
		memlimit = 16
	}
`

func TestAccNstcpbufparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNstcpbufparam_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstcpbufparamExist("citrixadc_nstcpbufparam.tf_nstcpbufparam", nil),
					resource.TestCheckResourceAttr("citrixadc_nstcpbufparam.tf_nstcpbufparam", "size", "32"),
					resource.TestCheckResourceAttr("citrixadc_nstcpbufparam.tf_nstcpbufparam", "memlimit", "8"),
				),
			},
			{
				Config: testAccNstcpbufparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstcpbufparamExist("citrixadc_nstcpbufparam.tf_nstcpbufparam", nil),
					resource.TestCheckResourceAttr("citrixadc_nstcpbufparam.tf_nstcpbufparam", "size", "64"),
					resource.TestCheckResourceAttr("citrixadc_nstcpbufparam.tf_nstcpbufparam", "memlimit", "16"),
				),
			},
		},
	})
}

func testAccCheckNstcpbufparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nstcpbufparam name is set")
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
		data, err := client.FindResource(service.Nstcpbufparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nstcpbufparam %s not found", n)
		}

		return nil
	}
}

func TestAccNstcpbufparam_import(t *testing.T) {
	const resAddr = "citrixadc_nstcpbufparam.tf_nstcpbufparam"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccNstcpbufparam_add},
			{
				Config:                  testAccNstcpbufparam_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccNstcpbufparam_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccNstcpbufparam_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstcpbufparamExist("citrixadc_nstcpbufparam.tf_nstcpbufparam", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccNstcpbufparam_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstcpbufparamExist("citrixadc_nstcpbufparam.tf_nstcpbufparam", nil),
				),
			},
		},
	})
}

const testAccNstcpbufparam_unset_step1 = `
	resource "citrixadc_nstcpbufparam" "tf_unset" {
		size     = 32
		memlimit = 8
	}
`

const testAccNstcpbufparam_unset_step2 = `
	resource "citrixadc_nstcpbufparam" "tf_unset" {
		# size and memlimit removed from config -> the provider must unset them
		# (revert to NITRO defaults of 64).
	}
`

func TestAccNstcpbufparam_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccNstcpbufparam_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstcpbufparamExist("citrixadc_nstcpbufparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nstcpbufparam.tf_unset", "size", "32"),
					resource.TestCheckResourceAttr("citrixadc_nstcpbufparam.tf_unset", "memlimit", "8"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccNstcpbufparam_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstcpbufparamExist("citrixadc_nstcpbufparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nstcpbufparam.tf_unset", "size", "64"),
					resource.TestCheckResourceAttr("citrixadc_nstcpbufparam.tf_unset", "memlimit", "64"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckNstcpbufparamADCValue("size", "64"),
					testAccCheckNstcpbufparamADCValue("memlimit", "64"),
				),
			},
		},
	})
}

// testAccCheckNstcpbufparamADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckNstcpbufparamADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Nstcpbufparam.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("nstcpbufparam not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("nstcpbufparam: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccNstcpbufparamDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNstcpbufparamDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nstcpbufparam.tf_nstcpbufparam_ds", "size", "32"),
					resource.TestCheckResourceAttr("data.citrixadc_nstcpbufparam.tf_nstcpbufparam_ds", "memlimit", "8"),
				),
			},
		},
	})
}

const testAccNstcpbufparamDataSource_basic = `

	resource "citrixadc_nstcpbufparam" "tf_nstcpbufparam_ds" {
		size     = 32
		memlimit = 8
	}

	data "citrixadc_nstcpbufparam" "tf_nstcpbufparam_ds" {
		depends_on = [citrixadc_nstcpbufparam.tf_nstcpbufparam_ds]
	}
`
