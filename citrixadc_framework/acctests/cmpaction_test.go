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

const testAccCmpaction_basic = `


	resource "citrixadc_cmpaction" "tf_cmpaction" {
		name    = "my_cmpaction"
		cmptype = "nocompress"
	}
  
`
const testAccCmpaction_update = `


	resource "citrixadc_cmpaction" "tf_cmpaction" {
		name    = "my_cmpaction"
		cmptype = "compress"
	}
 `

func TestAccCmpaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCmpactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCmpaction_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCmpactionExist("citrixadc_cmpaction.tf_cmpaction", nil),
					resource.TestCheckResourceAttr("citrixadc_cmpaction.tf_cmpaction", "name", "my_cmpaction"),
					resource.TestCheckResourceAttr("citrixadc_cmpaction.tf_cmpaction", "cmptype", "nocompress"),
				),
			},
			{
				Config: testAccCmpaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCmpactionExist("citrixadc_cmpaction.tf_cmpaction", nil),
					resource.TestCheckResourceAttr("citrixadc_cmpaction.tf_cmpaction", "name", "my_cmpaction"),
					resource.TestCheckResourceAttr("citrixadc_cmpaction.tf_cmpaction", "cmptype", "compress"),
				),
			},
		},
	})
}

func TestAccCmpaction_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_cmpaction.tf_cmpaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCmpactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCmpaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCmpactionExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Cmpaction.Type(), "my_cmpaction"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccCmpaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCmpactionExist(resAddr, nil)),
			},
		},
	})
}

func testAccCheckCmpactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cmpaction name is set")
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
		data, err := client.FindResource(service.Cmpaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("cmpaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckCmpactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_cmpaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Cmpaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("cmpaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccCmpaction_import(t *testing.T) {
	const resAddr = "citrixadc_cmpaction.tf_cmpaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCmpactionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccCmpaction_basic},
			{
				Config:                  testAccCmpaction_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccCmpaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckCmpactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccCmpaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCmpactionExist("citrixadc_cmpaction.tf_cmpaction", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccCmpaction_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckCmpactionExist("citrixadc_cmpaction.tf_cmpaction", nil)),
			},
		},
	})
}

// The cmpaction unset test covers addvaryheader, whose NITRO default is GLOBAL.
// Step 1 sets a non-default value; step 2 removes it so the provider unsets it
// (reverts to GLOBAL). cmptype is required and carried in both steps.
const testAccCmpaction_unset_step1 = `
resource "citrixadc_cmpaction" "tf_unset" {
  name          = "tf_test_cmpaction_unset"
  cmptype       = "compress"
  addvaryheader = "DISABLED"
}
`

const testAccCmpaction_unset_step2 = `
resource "citrixadc_cmpaction" "tf_unset" {
  name    = "tf_test_cmpaction_unset"
  cmptype = "compress"
  # addvaryheader removed from config -> the provider must unset it (revert to
  # the NITRO default, "GLOBAL").
}
`

func TestAccCmpaction_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCmpactionDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccCmpaction_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCmpactionExist("citrixadc_cmpaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_cmpaction.tf_unset", "addvaryheader", "DISABLED"),
				),
			},
			{
				// Removing the attribute must unset it: state (read back from the
				// appliance) reverts to the NITRO default, and the implicit
				// post-apply plan must be empty.
				Config: testAccCmpaction_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCmpactionExist("citrixadc_cmpaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_cmpaction.tf_unset", "addvaryheader", "GLOBAL"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckCmpactionADCValue("tf_test_cmpaction_unset", "addvaryheader", "GLOBAL"),
				),
			},
		},
	})
}

// testAccCheckCmpactionADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckCmpactionADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Cmpaction.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("cmpaction %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("cmpaction %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccCmpactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccCmpactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_cmpaction.tf_cmpaction_ds", "name", "tf_cmpaction_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_cmpaction.tf_cmpaction_ds", "cmptype", "nocompress"),
				),
			},
		},
	})
}

const testAccCmpactionDataSource_basic = `

resource "citrixadc_cmpaction" "tf_cmpaction_ds" {
    name    = "tf_cmpaction_ds"
    cmptype = "nocompress"
}

data "citrixadc_cmpaction" "tf_cmpaction_ds" {
    name = citrixadc_cmpaction.tf_cmpaction_ds.name
    depends_on = [citrixadc_cmpaction.tf_cmpaction_ds]
}

`
