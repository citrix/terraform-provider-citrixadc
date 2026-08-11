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

const testAccRewriteparam_basic = `
	resource "citrixadc_rewriteparam" "tf_rewriteparam" {
		timeout = 5
		undefaction = "RESET"
	}
`

const testAccRewriteparam_basic_update = `
	resource "citrixadc_rewriteparam" "tf_rewriteparam" {
		timeout = 6
		undefaction = "DROP"
	}
`

func TestAccRewriteparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRewriteparamDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRewriteparam_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewriteparamExist("citrixadc_rewriteparam.tf_rewriteparam", nil),
					resource.TestCheckResourceAttr("citrixadc_rewriteparam.tf_rewriteparam", "timeout", "5"),
					resource.TestCheckResourceAttr("citrixadc_rewriteparam.tf_rewriteparam", "undefaction", "RESET"),
				),
			},
			{
				Config: testAccRewriteparam_basic_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewriteparamExist("citrixadc_rewriteparam.tf_rewriteparam", nil),
					resource.TestCheckResourceAttr("citrixadc_rewriteparam.tf_rewriteparam", "timeout", "6"),
					resource.TestCheckResourceAttr("citrixadc_rewriteparam.tf_rewriteparam", "undefaction", "DROP"),
				),
			},
		},
	})
}

func testAccCheckRewriteparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No rewriteparam name is set")
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
		data, err := client.FindResource(service.Rewriteparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("rewriteparam %s not found", n)
		}

		return nil
	}
}

func testAccCheckRewriteparamDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_rewriteparam" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Rewriteparam.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("rewriteparam %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccRewriteparamDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRewriteparamDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.citrixadc_rewriteparam.test", "id"),
					resource.TestCheckResourceAttrSet("data.citrixadc_rewriteparam.test", "timeout"),
					resource.TestCheckResourceAttrSet("data.citrixadc_rewriteparam.test", "undefaction"),
				),
			},
		},
	})
}

func TestAccRewriteparam_import(t *testing.T) {
	const resAddr = "citrixadc_rewriteparam.tf_rewriteparam"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRewriteparamDestroy,
		Steps: []resource.TestStep{
			{Config: testAccRewriteparam_basic},
			{
				Config:                  testAccRewriteparam_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccRewriteparamDataSource_basic = `
data "citrixadc_rewriteparam" "test" {
}
`

// rewriteparam unset: step1 sets both mutable attrs to non-default values;
// step2 removes them so the provider must unset them, reverting to the NITRO
// documented defaults (timeout=3900, undefaction="NOREWRITE").
const testAccRewriteparam_unset_step1 = `
	resource "citrixadc_rewriteparam" "tf_unset" {
		timeout     = 2000
		undefaction = "RESET"
	}
`

const testAccRewriteparam_unset_step2 = `
	resource "citrixadc_rewriteparam" "tf_unset" {
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccRewriteparam_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRewriteparamDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccRewriteparam_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewriteparamExist("citrixadc_rewriteparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_rewriteparam.tf_unset", "timeout", "2000"),
					resource.TestCheckResourceAttr("citrixadc_rewriteparam.tf_unset", "undefaction", "RESET"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccRewriteparam_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewriteparamExist("citrixadc_rewriteparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_rewriteparam.tf_unset", "timeout", "3900"),
					resource.TestCheckResourceAttr("citrixadc_rewriteparam.tf_unset", "undefaction", "NOREWRITE"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckRewriteparamADCValue("timeout", "3900"),
					testAccCheckRewriteparamADCValue("undefaction", "NOREWRITE"),
				),
			},
		},
	})
}

// testAccCheckRewriteparamADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckRewriteparamADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Rewriteparam.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("rewriteparam not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("rewriteparam: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccRewriteparam_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckRewriteparamDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccRewriteparam_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckRewriteparamExist("citrixadc_rewriteparam.tf_rewriteparam", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccRewriteparam_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckRewriteparamExist("citrixadc_rewriteparam.tf_rewriteparam", nil)),
			},
		},
	})
}
