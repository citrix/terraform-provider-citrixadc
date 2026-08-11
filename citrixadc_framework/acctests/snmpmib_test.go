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

const testAccSnmpmib_basic = `

resource "citrixadc_snmpmib" "tf_snmpmib" {
	contact  = "phone_number"
	name     = "my_name"
	location = "LOCATION"
	customid = "CUSTOMER_ID"
	}
  
`
const testAccSnmpmib_update = `

resource "citrixadc_snmpmib" "tf_snmpmib" {
	contact  = "phone_number2"
	name     = "my_name2"
	location = "LOCATION2"
	customid = "CUSTOMER_ID2"
	}
  
`

func TestAccSnmpmib_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpmib_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpmibExist("citrixadc_snmpmib.tf_snmpmib", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpmib.tf_snmpmib", "contact", "phone_number"),
					resource.TestCheckResourceAttr("citrixadc_snmpmib.tf_snmpmib", "name", "my_name"),
					resource.TestCheckResourceAttr("citrixadc_snmpmib.tf_snmpmib", "location", "LOCATION"),
					resource.TestCheckResourceAttr("citrixadc_snmpmib.tf_snmpmib", "customid", "CUSTOMER_ID"),
				),
			},
			{
				Config: testAccSnmpmib_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpmibExist("citrixadc_snmpmib.tf_snmpmib", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpmib.tf_snmpmib", "contact", "phone_number2"),
					resource.TestCheckResourceAttr("citrixadc_snmpmib.tf_snmpmib", "name", "my_name2"),
					resource.TestCheckResourceAttr("citrixadc_snmpmib.tf_snmpmib", "location", "LOCATION2"),
					resource.TestCheckResourceAttr("citrixadc_snmpmib.tf_snmpmib", "customid", "CUSTOMER_ID2"),
				),
			},
		},
	})
}

func TestAccSnmpmib_import(t *testing.T) {
	const resAddr = "citrixadc_snmpmib.tf_snmpmib"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccSnmpmib_basic},
			{
				Config:                  testAccSnmpmib_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckSnmpmibExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No snmpmib name is set")
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
		data, err := client.FindResource(service.Snmpmib.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("snmpmib %s not found", n)
		}

		return nil
	}
}

func TestAccSnmpmibDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpmibDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_snmpmib.tf_snmpmib_ds", "contact", "phone_number_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_snmpmib.tf_snmpmib_ds", "name", "my_name_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_snmpmib.tf_snmpmib_ds", "location", "LOCATION_DS"),
				),
			},
		},
	})
}

func TestAccSnmpmib_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccSnmpmib_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSnmpmibExist("citrixadc_snmpmib.tf_snmpmib", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccSnmpmib_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckSnmpmibExist("citrixadc_snmpmib.tf_snmpmib", nil)),
			},
		},
	})
}

// snmpmib is a singleton global config. step1 sets the unset-eligible
// read/write attributes to non-default values; step2 removes them so the
// provider must unset them, reverting to the documented NITRO defaults.
const testAccSnmpmib_unset_step1 = `
resource "citrixadc_snmpmib" "tf_unset" {
	contact  = "tf_contact"
	name     = "tf_name"
	location = "tf_location"
	customid = "tf_customid"
}
`

const testAccSnmpmib_unset_step2 = `
resource "citrixadc_snmpmib" "tf_unset" {
	# All unset-eligible attributes removed from config -> the provider must
	# unset them (revert to NITRO defaults).
}
`

func TestAccSnmpmib_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccSnmpmib_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpmibExist("citrixadc_snmpmib.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpmib.tf_unset", "contact", "tf_contact"),
					resource.TestCheckResourceAttr("citrixadc_snmpmib.tf_unset", "name", "tf_name"),
					resource.TestCheckResourceAttr("citrixadc_snmpmib.tf_unset", "location", "tf_location"),
					resource.TestCheckResourceAttr("citrixadc_snmpmib.tf_unset", "customid", "tf_customid"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccSnmpmib_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpmibExist("citrixadc_snmpmib.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpmib.tf_unset", "contact", "WebMaster (default)"),
					resource.TestCheckResourceAttr("citrixadc_snmpmib.tf_unset", "name", "NetScaler"),
					resource.TestCheckResourceAttr("citrixadc_snmpmib.tf_unset", "location", "POP (default)"),
					resource.TestCheckResourceAttr("citrixadc_snmpmib.tf_unset", "customid", "Default"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckSnmpmibADCValue("contact", "WebMaster (default)"),
					testAccCheckSnmpmibADCValue("name", "NetScaler"),
					testAccCheckSnmpmibADCValue("location", "POP (default)"),
					testAccCheckSnmpmibADCValue("customid", "Default"),
				),
			},
		},
	})
}

// testAccCheckSnmpmibADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it. snmpmib is a singleton, so it is fetched with an empty id.
func testAccCheckSnmpmibADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Snmpmib.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("snmpmib not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("snmpmib: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

const testAccSnmpmibDataSource_basic = `

resource "citrixadc_snmpmib" "tf_snmpmib_ds" {
	contact  = "phone_number_ds"
	name     = "my_name_ds"
	location = "LOCATION_DS"
	customid = "CUSTOMER_ID_DS"
}

data "citrixadc_snmpmib" "tf_snmpmib_ds" {
	ownernode = -1
	depends_on = [citrixadc_snmpmib.tf_snmpmib_ds]
}
`
