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

const testAccLsnparameter_basic = `

	resource "citrixadc_lsnparameter" "tf_lsnparameter" {
		sessionsync          = "ENABLED"
		subscrsessionremoval = "ENABLED"
	}
`
const testAccLsnparameter_update = `

	resource "citrixadc_lsnparameter" "tf_lsnparameter" {
		sessionsync          = "DISABLED"
		subscrsessionremoval = "DISABLED"
	}
`

func TestAccLsnparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnparameterExist("citrixadc_lsnparameter.tf_lsnparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnparameter.tf_lsnparameter", "sessionsync", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnparameter.tf_lsnparameter", "subscrsessionremoval", "ENABLED"),
				),
			},
			{
				Config: testAccLsnparameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnparameterExist("citrixadc_lsnparameter.tf_lsnparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnparameter.tf_lsnparameter", "sessionsync", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnparameter.tf_lsnparameter", "subscrsessionremoval", "DISABLED"),
				),
			},
		},
	})
}

func TestAccLsnparameter_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccLsnparameter_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsnparameterExist("citrixadc_lsnparameter.tf_lsnparameter", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccLsnparameter_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckLsnparameterExist("citrixadc_lsnparameter.tf_lsnparameter", nil)),
			},
		},
	})
}

// Unset test: step1 sets the unset-eligible attributes to non-default values;
// step2 removes them from config so the provider must unset them, reverting to
// the documented NITRO defaults (sessionsync=ENABLED, subscrsessionremoval=DISABLED).
const testAccLsnparameter_unset_step1 = `
	resource "citrixadc_lsnparameter" "tf_unset" {
		sessionsync          = "DISABLED"
		subscrsessionremoval = "ENABLED"
	}
`

const testAccLsnparameter_unset_step2 = `
	resource "citrixadc_lsnparameter" "tf_unset" {
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccLsnparameter_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccLsnparameter_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnparameterExist("citrixadc_lsnparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnparameter.tf_unset", "sessionsync", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnparameter.tf_unset", "subscrsessionremoval", "ENABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccLsnparameter_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnparameterExist("citrixadc_lsnparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnparameter.tf_unset", "sessionsync", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnparameter.tf_unset", "subscrsessionremoval", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckLsnparameterADCValue("sessionsync", "ENABLED"),
					testAccCheckLsnparameterADCValue("subscrsessionremoval", "DISABLED"),
				),
			},
		},
	})
}

// testAccCheckLsnparameterADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckLsnparameterADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Lsnparameter.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("lsnparameter not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("lsnparameter: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func testAccCheckLsnparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsnparameter name is set")
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
		data, err := client.FindResource("lsnparameter", "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lsnparameter %s not found", n)
		}

		return nil
	}
}

func TestAccLsnparameter_import(t *testing.T) {
	const resAddr = "citrixadc_lsnparameter.tf_lsnparameter"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccLsnparameter_basic},
			{
				Config:                  testAccLsnparameter_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccLsnparameterDataSource_basic = `

resource "citrixadc_lsnparameter" "tf_lsnparameter_ds" {
	sessionsync          = "ENABLED"
	subscrsessionremoval = "ENABLED"
}

data "citrixadc_lsnparameter" "tf_lsnparameter_ds" {
	depends_on = [citrixadc_lsnparameter.tf_lsnparameter_ds]
}
`

func TestAccLsnparameterDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnparameterDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lsnparameter.tf_lsnparameter_ds", "sessionsync", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnparameter.tf_lsnparameter_ds", "subscrsessionremoval", "ENABLED"),
				),
			},
		},
	})
}
