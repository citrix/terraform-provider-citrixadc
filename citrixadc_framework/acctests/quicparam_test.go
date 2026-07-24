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

const testAccQuicparam_add = `
	resource "citrixadc_quicparam" "tf_quicparam" {
		quicsecrettimeout = 3600
	}
`
const testAccQuicparam_update = `
	resource "citrixadc_quicparam" "tf_quicparam" {
		quicsecrettimeout = 7200
	}
`

func TestAccQuicparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccQuicparam_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckQuicparamExist("citrixadc_quicparam.tf_quicparam", nil),
					resource.TestCheckResourceAttr("citrixadc_quicparam.tf_quicparam", "quicsecrettimeout", "3600"),
				),
			},
			{
				Config: testAccQuicparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckQuicparamExist("citrixadc_quicparam.tf_quicparam", nil),
					resource.TestCheckResourceAttr("citrixadc_quicparam.tf_quicparam", "quicsecrettimeout", "7200"),
				),
			},
		},
	})
}

func TestAccQuicparam_import(t *testing.T) {
	const resAddr = "citrixadc_quicparam.tf_quicparam"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccQuicparam_add},
			{
				Config:                  testAccQuicparam_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckQuicparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No quicparam name is set")
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
		data, err := client.FindResource(service.Quicparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("quicparam %s not found", n)
		}

		return nil
	}
}

const testAccQuicparam_unset_step1 = `
	resource "citrixadc_quicparam" "tf_unset" {
		quicsecrettimeout = 7200
	}
`

const testAccQuicparam_unset_step2 = `
	resource "citrixadc_quicparam" "tf_unset" {
		# quicsecrettimeout removed from config -> provider must unset it
	}
`

func TestAccQuicparam_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default value applies and persists.
				Config: testAccQuicparam_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckQuicparamExist("citrixadc_quicparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_quicparam.tf_unset", "quicsecrettimeout", "7200"),
				),
			},
			{
				// Removing it must unset -> state reverts to the NITRO default,
				// and the implicit post-apply plan must be empty.
				Config: testAccQuicparam_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckQuicparamExist("citrixadc_quicparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_quicparam.tf_unset", "quicsecrettimeout", "3600"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckQuicparamADCValue("quicsecrettimeout", "3600"),
				),
			},
		},
	})
}

// testAccCheckQuicparamADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
// quicparam is a singleton, so it is looked up with an empty resource name.
func testAccCheckQuicparamADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Quicparam.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("quicparam not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("quicparam: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

const testAccQuicparamDataSource_basic = `

	resource "citrixadc_quicparam" "tf_quicparam" {
		quicsecrettimeout = 3600
	}

	data "citrixadc_quicparam" "tf_quicparam" {
		depends_on = [citrixadc_quicparam.tf_quicparam]
	}
`

func TestAccQuicparamDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccQuicparamDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_quicparam.tf_quicparam", "quicsecrettimeout", "3600"),
				),
			},
		},
	})
}
