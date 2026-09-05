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

const testAccLsnappsattributes_basic = `


resource "citrixadc_lsnappsattributes" "tf_lsnappsattributes" {
	name              = "my_lsn_appattributes"
	transportprotocol = "TCP"
	port              = 90
	sessiontimeout    = 40
	}
  
`
const testAccLsnappsattributes_update = `


resource "citrixadc_lsnappsattributes" "tf_lsnappsattributes" {
	name              = "my_lsn_appattributes"
	transportprotocol = "TCP"
	port              = 90
	sessiontimeout    = 60
	}
  
`

func TestAccLsnappsattributes_basic(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this LSN resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnappsattributesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnappsattributes_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnappsattributesExist("citrixadc_lsnappsattributes.tf_lsnappsattributes", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnappsattributes.tf_lsnappsattributes", "name", "my_lsn_appattributes"),
					resource.TestCheckResourceAttr("citrixadc_lsnappsattributes.tf_lsnappsattributes", "transportprotocol", "TCP"),
					resource.TestCheckResourceAttr("citrixadc_lsnappsattributes.tf_lsnappsattributes", "port", "90"),
					resource.TestCheckResourceAttr("citrixadc_lsnappsattributes.tf_lsnappsattributes", "sessiontimeout", "40"),
				),
			},
			{
				Config: testAccLsnappsattributes_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnappsattributesExist("citrixadc_lsnappsattributes.tf_lsnappsattributes", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnappsattributes.tf_lsnappsattributes", "name", "my_lsn_appattributes"),
					resource.TestCheckResourceAttr("citrixadc_lsnappsattributes.tf_lsnappsattributes", "transportprotocol", "TCP"),
					resource.TestCheckResourceAttr("citrixadc_lsnappsattributes.tf_lsnappsattributes", "port", "90"),
					resource.TestCheckResourceAttr("citrixadc_lsnappsattributes.tf_lsnappsattributes", "sessiontimeout", "60"),
				),
			},
		},
	})
}

func TestAccLsnappsattributes_selfHealing(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this LSN resource!")
	const resAddr = "citrixadc_lsnappsattributes.tf_lsnappsattributes"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnappsattributesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnappsattributes_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsnappsattributesExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Lsnappsattributes.Type(), "my_lsn_appattributes"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccLsnappsattributes_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsnappsattributesExist(resAddr, nil)),
			},
		},
	})
}

func TestAccLsnappsattributes_import(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this LSN resource!")
	const resAddr = "citrixadc_lsnappsattributes.tf_lsnappsattributes"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnappsattributesDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLsnappsattributes_basic},
			{
				Config:                  testAccLsnappsattributes_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccLsnappsattributes_sdkv2StateUpgrade(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this LSN resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLsnappsattributesDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccLsnappsattributes_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsnappsattributesExist("citrixadc_lsnappsattributes.tf_lsnappsattributes", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccLsnappsattributes_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsnappsattributesExist("citrixadc_lsnappsattributes.tf_lsnappsattributes", nil)),
			},
		},
	})
}

// sessiontimeout is the only unsettable attribute (spec unset operation lists
// only sessiontimeout; name/transportprotocol are keys and port is ForceNew).
// Its NITRO default is 30.
const testAccLsnappsattributes_unset_step1 = `
resource "citrixadc_lsnappsattributes" "tf_unset" {
	name              = "tf_test_lsnappsattributes_unset"
	transportprotocol = "TCP"
	port              = "90"
	sessiontimeout    = 40
}
`

const testAccLsnappsattributes_unset_step2 = `
resource "citrixadc_lsnappsattributes" "tf_unset" {
	name              = "tf_test_lsnappsattributes_unset"
	transportprotocol = "TCP"
	port              = "90"
	# sessiontimeout removed from config -> provider must unset it (revert to NITRO default 30).
}
`

func TestAccLsnappsattributes_unset(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this LSN resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnappsattributesDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccLsnappsattributes_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnappsattributesExist("citrixadc_lsnappsattributes.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnappsattributes.tf_unset", "sessiontimeout", "40"),
				),
			},
			{
				// Removing sessiontimeout must unset it: state (read back from the
				// appliance) reverts to the documented NITRO default, and the
				// implicit post-apply plan must be empty.
				Config: testAccLsnappsattributes_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnappsattributesExist("citrixadc_lsnappsattributes.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnappsattributes.tf_unset", "sessiontimeout", "30"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckLsnappsattributesADCValue("tf_test_lsnappsattributes_unset", "sessiontimeout", "30"),
				),
			},
		},
	})
}

// testAccCheckLsnappsattributesADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckLsnappsattributesADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Lsnappsattributes.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("lsnappsattributes %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("lsnappsattributes %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func testAccCheckLsnappsattributesExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsnappsattributes name is set")
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
		data, err := client.FindResource("lsnappsattributes", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lsnappsattributes %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsnappsattributesDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsnappsattributes" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lsnappsattributes", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lsnappsattributes %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLsnappsattributesDataSource_basic = `

resource "citrixadc_lsnappsattributes" "tf_lsnappsattributes" {
	name              = "my_lsn_appattributes_ds"
	transportprotocol = "TCP"
	port              = 90
	sessiontimeout    = 40
}

data "citrixadc_lsnappsattributes" "tf_lsnappsattributes_ds" {
	name = citrixadc_lsnappsattributes.tf_lsnappsattributes.name
	depends_on = [citrixadc_lsnappsattributes.tf_lsnappsattributes]
}
`

func TestAccLsnappsattributesDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnappsattributesDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lsnappsattributes.tf_lsnappsattributes_ds", "name", "my_lsn_appattributes_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnappsattributes.tf_lsnappsattributes_ds", "transportprotocol", "TCP"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnappsattributes.tf_lsnappsattributes_ds", "port", "90"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnappsattributes.tf_lsnappsattributes_ds", "sessiontimeout", "40"),
				),
			},
		},
	})
}
