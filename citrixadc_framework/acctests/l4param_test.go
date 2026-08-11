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

const testAccL4param_add = `
	resource "citrixadc_l4param" "tf_l4param" {
		l2connmethod = "Channel"
		l4switch     = "ENABLED"
	}
`
const testAccL4param_update = `
	resource "citrixadc_l4param" "tf_l4param" {
		l2connmethod = "MacVlanChannel"
		l4switch     = "DISABLED"
	}
`

func TestAccL4param_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccL4param_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckL4paramExist("citrixadc_l4param.tf_l4param", nil),
					resource.TestCheckResourceAttr("citrixadc_l4param.tf_l4param", "l2connmethod", "Channel"),
					resource.TestCheckResourceAttr("citrixadc_l4param.tf_l4param", "l4switch", "ENABLED"),
				),
			},
			{
				Config: testAccL4param_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckL4paramExist("citrixadc_l4param.tf_l4param", nil),
					resource.TestCheckResourceAttr("citrixadc_l4param.tf_l4param", "l2connmethod", "MacVlanChannel"),
					resource.TestCheckResourceAttr("citrixadc_l4param.tf_l4param", "l4switch", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckL4paramExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No l4param name is set")
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
		data, err := client.FindResource(service.L4param.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("l4param %s not found", n)
		}

		return nil
	}
}

func TestAccL4param_import(t *testing.T) {
	const resAddr = "citrixadc_l4param.tf_l4param"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccL4param_add},
			{
				Config:                  testAccL4param_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccL4paramDataSource_basic = `

	resource "citrixadc_l4param" "tf_l4param" {
		l2connmethod = "Channel"
		l4switch     = "ENABLED"
	}

	data "citrixadc_l4param" "tf_l4param" {
		depends_on = [citrixadc_l4param.tf_l4param]
	}
`

func TestAccL4param_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccL4param_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckL4paramExist("citrixadc_l4param.tf_l4param", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccL4param_add,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckL4paramExist("citrixadc_l4param.tf_l4param", nil)),
			},
		},
	})
}

// l4param unset test: step1 sets both mutable attrs to non-default values,
// step2 removes them so the provider must unset them (revert to NITRO
// defaults: l2connmethod=MacVlanChannel, l4switch=DISABLED).
const testAccL4param_unset_step1 = `
	resource "citrixadc_l4param" "tf_unset" {
		l2connmethod = "Channel"
		l4switch     = "ENABLED"
	}
`

const testAccL4param_unset_step2 = `
	resource "citrixadc_l4param" "tf_unset" {
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccL4param_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccL4param_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckL4paramExist("citrixadc_l4param.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_l4param.tf_unset", "l2connmethod", "Channel"),
					resource.TestCheckResourceAttr("citrixadc_l4param.tf_unset", "l4switch", "ENABLED"),
				),
			},
			{
				Config: testAccL4param_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckL4paramExist("citrixadc_l4param.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_l4param.tf_unset", "l2connmethod", "MacVlanChannel"),
					resource.TestCheckResourceAttr("citrixadc_l4param.tf_unset", "l4switch", "DISABLED"),
					testAccCheckL4paramADCValue("l2connmethod", "MacVlanChannel"),
					testAccCheckL4paramADCValue("l4switch", "DISABLED"),
				),
			},
		},
	})
}

// testAccCheckL4paramADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckL4paramADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.L4param.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("l4param not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("l4param: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccL4paramDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccL4paramDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_l4param.tf_l4param", "l2connmethod", "Channel"),
					resource.TestCheckResourceAttr("data.citrixadc_l4param.tf_l4param", "l4switch", "ENABLED"),
				),
			},
		},
	})
}
