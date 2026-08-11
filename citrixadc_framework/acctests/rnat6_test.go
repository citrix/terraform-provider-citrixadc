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

const testAccRnat6_basic = `


resource "citrixadc_rnat6" "tf_rnat6" {
	name             = "my_rnat6"
	network          = "2003::/64"
	srcippersistency = "DISABLED"
	}
  
`
const testAccRnat6_update = `


resource "citrixadc_rnat6" "tf_rnat6" {
	name             = "my_rnat6"
	network          = "2003::/64"
	srcippersistency = "ENABLED"
	}
  
`

func TestAccRnat6_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccRnat6_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRnat6Exist("citrixadc_rnat6.tf_rnat6", nil),
					resource.TestCheckResourceAttr("citrixadc_rnat6.tf_rnat6", "network", "2003::/64"),
					resource.TestCheckResourceAttr("citrixadc_rnat6.tf_rnat6", "srcippersistency", "DISABLED"),
				),
			},
			{
				Config: testAccRnat6_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRnat6Exist("citrixadc_rnat6.tf_rnat6", nil),
					resource.TestCheckResourceAttr("citrixadc_rnat6.tf_rnat6", "network", "2003::/64"),
					resource.TestCheckResourceAttr("citrixadc_rnat6.tf_rnat6", "srcippersistency", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckRnat6Exist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No rnat6 name is set")
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
		data, err := client.FindResource(service.Rnat6.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("rnat6 %s not found", n)
		}

		return nil
	}
}

func TestAccRnat6DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRnat6DataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_rnat6.test", "name", "tf_rnat6_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_rnat6.test", "network", "2004::/64"),
					resource.TestCheckResourceAttrSet("data.citrixadc_rnat6.test", "id"),
				),
			},
		},
	})
}

const testAccRnat6DataSource_basic = `
resource "citrixadc_rnat6" "test" {
	name    = "tf_rnat6_ds"
	network = "2004::/64"
}

data "citrixadc_rnat6" "test" {
	name = citrixadc_rnat6.test.name
}
`

func TestAccRnat6_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccRnat6_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckRnat6Exist("citrixadc_rnat6.tf_rnat6", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccRnat6_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckRnat6Exist("citrixadc_rnat6.tf_rnat6", nil)),
			},
		},
	})
}

// rnat6 unset coverage: srcippersistency is the only type-independent,
// mutable, spec-unsettable attribute with a documented NITRO default
// (DISABLED). ownergroup (default DEFAULT_NG) requires a cluster node group to
// set a non-default value and redirectport has no documented default, so both
// are excluded.
const testAccRnat6_unset_step1 = `
resource "citrixadc_rnat6" "tf_unset" {
	name             = "tf_rnat6_unset"
	network          = "2005::/64"
	srcippersistency = "ENABLED"
}
`

const testAccRnat6_unset_step2 = `
resource "citrixadc_rnat6" "tf_unset" {
	name    = "tf_rnat6_unset"
	network = "2005::/64"
	# srcippersistency removed -> provider must unset it (revert to DISABLED).
}
`

func TestAccRnat6_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccRnat6_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRnat6Exist("citrixadc_rnat6.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_rnat6.tf_unset", "srcippersistency", "ENABLED"),
				),
			},
			{
				// Removing the attribute must unset it: state (read back from the
				// appliance) reverts to the NITRO default and the implicit
				// post-apply plan must be empty.
				Config: testAccRnat6_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRnat6Exist("citrixadc_rnat6.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_rnat6.tf_unset", "srcippersistency", "DISABLED"),
					testAccCheckRnat6ADCValue("tf_rnat6_unset", "srcippersistency", "DISABLED"),
				),
			},
		},
	})
}

// testAccCheckRnat6ADCValue asserts an attribute's value directly on the
// appliance, proving the unset actually reverted it.
func testAccCheckRnat6ADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Rnat6.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("rnat6 %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("rnat6 %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccRnat6_import(t *testing.T) {
	const resAddr = "citrixadc_rnat6.tf_rnat6"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccRnat6_basic},
			{
				Config:                  testAccRnat6_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}
