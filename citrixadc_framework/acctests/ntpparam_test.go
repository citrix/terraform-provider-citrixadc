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

const testAccNtpparam_basic = `
	resource "citrixadc_ntpparam" "tf_ntpparam" {
		authentication = "YES"
		trustedkey     = [123, 456]
		autokeylogsec  = 15
		revokelogsec   = 20
	}
`
const testAccNtpparam_update = `
	resource "citrixadc_ntpparam" "tf_ntpparam" {
		authentication = "NO"
		trustedkey     = [1234, 4567]
		autokeylogsec  = 10
		revokelogsec   = 12
	}
`

func TestAccNtpparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNtpparam_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNtpparamExist("citrixadc_ntpparam.tf_ntpparam", nil),
					resource.TestCheckResourceAttr("citrixadc_ntpparam.tf_ntpparam", "authentication", "YES"),
					//resource.TestCheckResourceAttr("citrixadc_ntpparam.tf_ntpparam", "trustedkey", "[123, 456]"),
					resource.TestCheckResourceAttr("citrixadc_ntpparam.tf_ntpparam", "autokeylogsec", "15"),
					resource.TestCheckResourceAttr("citrixadc_ntpparam.tf_ntpparam", "revokelogsec", "20"),
				),
			},
			{
				Config: testAccNtpparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNtpparamExist("citrixadc_ntpparam.tf_ntpparam", nil),
					resource.TestCheckResourceAttr("citrixadc_ntpparam.tf_ntpparam", "authentication", "NO"),
					//resource.TestCheckResourceAttr("citrixadc_ntpparam.tf_ntpparam", "trustedkey", "[1234, 4567]"),
					resource.TestCheckResourceAttr("citrixadc_ntpparam.tf_ntpparam", "autokeylogsec", "10"),
					resource.TestCheckResourceAttr("citrixadc_ntpparam.tf_ntpparam", "revokelogsec", "12"),
				),
			},
		},
	})
}

func testAccCheckNtpparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ntpparam name is set")
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
		data, err := client.FindResource(service.Ntpparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("ntpparam %s not found", n)
		}

		return nil
	}
}
func TestAccNtpparamDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNtpparamDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.citrixadc_ntpparam.test", "id"),
					resource.TestCheckResourceAttrSet("data.citrixadc_ntpparam.test", "authentication"),
				),
			},
		},
	})
}

const testAccNtpparamDataSource_basic = `
data "citrixadc_ntpparam" "test" {
}
`

func TestAccNtpparam_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccNtpparam_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNtpparamExist("citrixadc_ntpparam.tf_ntpparam", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccNtpparam_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckNtpparamExist("citrixadc_ntpparam.tf_ntpparam", nil)),
			},
		},
	})
}

// testAccNtpparam_unset_step1 sets the unset-eligible scalar attributes to
// valid NON-default values.
const testAccNtpparam_unset_step1 = `
	resource "citrixadc_ntpparam" "tf_unset" {
		authentication = "NO"
		autokeylogsec  = 10
		revokelogsec   = 20
	}
`

// testAccNtpparam_unset_step2 removes all unset-eligible attributes -> the
// provider must unset them (revert to NITRO defaults).
const testAccNtpparam_unset_step2 = `
	resource "citrixadc_ntpparam" "tf_unset" {
	}
`

func TestAccNtpparam_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccNtpparam_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNtpparamExist("citrixadc_ntpparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_ntpparam.tf_unset", "authentication", "NO"),
					resource.TestCheckResourceAttr("citrixadc_ntpparam.tf_unset", "autokeylogsec", "10"),
					resource.TestCheckResourceAttr("citrixadc_ntpparam.tf_unset", "revokelogsec", "20"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccNtpparam_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNtpparamExist("citrixadc_ntpparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_ntpparam.tf_unset", "authentication", "YES"),
					resource.TestCheckResourceAttr("citrixadc_ntpparam.tf_unset", "autokeylogsec", "12"),
					resource.TestCheckResourceAttr("citrixadc_ntpparam.tf_unset", "revokelogsec", "16"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckNtpparamADCValue("authentication", "YES"),
					testAccCheckNtpparamADCValue("autokeylogsec", "12"),
					testAccCheckNtpparamADCValue("revokelogsec", "16"),
				),
			},
		},
	})
}

// testAccCheckNtpparamADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckNtpparamADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Ntpparam.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("ntpparam not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("ntpparam: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccNtpparam_import(t *testing.T) {
	const resAddr = "citrixadc_ntpparam.tf_ntpparam"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccNtpparam_basic},
			{
				Config:                  testAccNtpparam_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}
