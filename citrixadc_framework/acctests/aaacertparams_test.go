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

const testAccAaacertparams_basic = `


	resource "citrixadc_aaacertparams" "tf_aaacertparams" {
		usernamefield              = "Subject:CN"
		groupnamefield             = "Subject:OU"
		defaultauthenticationgroup = 50
	}
`
const testAccAaacertparams_update = `


	resource "citrixadc_aaacertparams" "tf_aaacertparams" {
		usernamefield              = "Subject:CW"
		groupnamefield             = "Subject:OW"
		defaultauthenticationgroup = 50
	}
`

func TestAccAaacertparams_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAaacertparams_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaacertparamsExist("citrixadc_aaacertparams.tf_aaacertparams", nil),
					resource.TestCheckResourceAttr("citrixadc_aaacertparams.tf_aaacertparams", "usernamefield", "Subject:CN"),
					resource.TestCheckResourceAttr("citrixadc_aaacertparams.tf_aaacertparams", "groupnamefield", "Subject:OU"),
				),
			},
			{
				Config: testAccAaacertparams_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaacertparamsExist("citrixadc_aaacertparams.tf_aaacertparams", nil),
					resource.TestCheckResourceAttr("citrixadc_aaacertparams.tf_aaacertparams", "usernamefield", "Subject:CW"),
					resource.TestCheckResourceAttr("citrixadc_aaacertparams.tf_aaacertparams", "groupnamefield", "Subject:OW"),
				),
			},
		},
	})
}

func testAccCheckAaacertparamsExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No aaacertparams name is set")
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
		data, err := client.FindResource(service.Aaacertparams.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("aaacertparams %s not found", n)
		}

		return nil
	}
}

func TestAccAaacertparams_import(t *testing.T) {
	const resAddr = "citrixadc_aaacertparams.tf_aaacertparams"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccAaacertparams_basic},
			{
				Config:                  testAccAaacertparams_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccAaacertparams_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccAaacertparams_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAaacertparamsExist("citrixadc_aaacertparams.tf_aaacertparams", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccAaacertparams_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAaacertparamsExist("citrixadc_aaacertparams.tf_aaacertparams", nil)),
			},
		},
	})
}

const testAccAaacertparamsDataSource_basic = `


	resource "citrixadc_aaacertparams" "tf_aaacertparams" {
		usernamefield              = "Subject:CN"
		groupnamefield             = "Subject:OU"
		defaultauthenticationgroup = 50
	}
	
	data "citrixadc_aaacertparams" "tf_aaacertparams" {
		depends_on = [citrixadc_aaacertparams.tf_aaacertparams]
	}
`

// aaacertparams is a singleton config resource. All three read/write attributes
// (usernamefield, groupnamefield, defaultauthenticationgroup) support the NITRO
// unset operation and revert to no value (absent from GET) when unset.
const testAccAaacertparams_unset_step1 = `
	resource "citrixadc_aaacertparams" "tf_unset" {
		usernamefield              = "Subject:CW"
		groupnamefield             = "Subject:OW"
		defaultauthenticationgroup = "tf_unset_grp"
	}
`

const testAccAaacertparams_unset_step2 = `
	resource "citrixadc_aaacertparams" "tf_unset" {
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults, i.e. no value).
	}
`

func TestAccAaacertparams_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccAaacertparams_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaacertparamsExist("citrixadc_aaacertparams.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_aaacertparams.tf_unset", "usernamefield", "Subject:CW"),
					resource.TestCheckResourceAttr("citrixadc_aaacertparams.tf_unset", "groupnamefield", "Subject:OW"),
					resource.TestCheckResourceAttr("citrixadc_aaacertparams.tf_unset", "defaultauthenticationgroup", "tf_unset_grp"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to the NITRO defaults, and the implicit
				// post-apply plan must be empty.
				Config: testAccAaacertparams_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaacertparamsExist("citrixadc_aaacertparams.tf_unset", nil),
					resource.TestCheckNoResourceAttr("citrixadc_aaacertparams.tf_unset", "usernamefield"),
					resource.TestCheckNoResourceAttr("citrixadc_aaacertparams.tf_unset", "groupnamefield"),
					resource.TestCheckNoResourceAttr("citrixadc_aaacertparams.tf_unset", "defaultauthenticationgroup"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAaacertparamsADCValue("usernamefield", ""),
					testAccCheckAaacertparamsADCValue("groupnamefield", ""),
					testAccCheckAaacertparamsADCValue("defaultauthenticationgroup", ""),
				),
			},
		},
	})
}

// testAccCheckAaacertparamsADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it. For unset string attributes NITRO omits the key entirely from GET.
func testAccCheckAaacertparamsADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Aaacertparams.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("aaacertparams not found on appliance")
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("aaacertparams: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccAaacertparamsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAaacertparamsDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_aaacertparams.tf_aaacertparams", "usernamefield", "Subject:CN"),
					resource.TestCheckResourceAttr("data.citrixadc_aaacertparams.tf_aaacertparams", "groupnamefield", "Subject:OU"),
					resource.TestCheckResourceAttr("data.citrixadc_aaacertparams.tf_aaacertparams", "defaultauthenticationgroup", "50"),
				),
			},
		},
	})
}
