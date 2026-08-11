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

const testAccNspbr_basic = `


	resource "citrixadc_nspbr" "tf_nspbr" {
		name = "my_nspbr"
		action = "DENY"
	}
`
const testAccNspbr_update = `


	resource "citrixadc_nspbr" "tf_nspbr" {
		name = "my_nspbr"
		action = "ALLOW"
		nexthop = "true"
		nexthopval = "10.222.74.128"
	}
`

func TestAccNspbr_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNspbrDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNspbr_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNspbrExist("citrixadc_nspbr.tf_nspbr", nil),
					resource.TestCheckResourceAttr("citrixadc_nspbr.tf_nspbr", "name", "my_nspbr"),
					resource.TestCheckResourceAttr("citrixadc_nspbr.tf_nspbr", "action", "DENY"),
				),
			},
			// Commenting out update test, because this requires valid ip in the same subnet as nexthop
			// {
			// 	Config: testAccNspbr_update,
			// 	Check: resource.ComposeTestCheckFunc(
			// 		testAccCheckNspbrExist("citrixadc_nspbr.tf_nspbr", nil),
			// 		resource.TestCheckResourceAttr("citrixadc_nspbr.tf_nspbr", "name", "my_nspbr"),
			// 		resource.TestCheckResourceAttr("citrixadc_nspbr.tf_nspbr", "action", "ALLOW"),
			// 		resource.TestCheckResourceAttr("citrixadc_nspbr.tf_nspbr", "nexthopval", "10.222.74.128"),
			// 	),
			// },
		},
	})
}

func testAccCheckNspbrExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nspbr name is set")
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
		data, err := client.FindResource(service.Nspbr.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nspbr %s not found", n)
		}

		return nil
	}
}

func testAccCheckNspbrDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nspbr" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Nspbr.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nspbr %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccNspbr_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_nspbr.tf_nspbr"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNspbrDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNspbr_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNspbrExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Nspbr.Type(), "my_nspbr"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccNspbr_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNspbrExist(resAddr, nil)),
			},
		},
	})
}

const testAccNspbrDataSource_basic = `
	resource "citrixadc_nspbr" "tf_nspbr_ds" {
		name     = "tf_test_nspbr_ds"
		action   = "DENY"
		srcip    = true
		srcipop  = "="
		srcipval = "192.0.2.0-192.0.2.255"
		destip   = true
		destipop = "="
		destipval = "203.0.113.0-203.0.113.255"
		priority = 100
	}

	data "citrixadc_nspbr" "tf_nspbr_ds" {
		name = citrixadc_nspbr.tf_nspbr_ds.name
	}
`

func TestAccNspbr_import(t *testing.T) {
	const resAddr = "citrixadc_nspbr.tf_nspbr"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNspbrDestroy,
		Steps: []resource.TestStep{
			{Config: testAccNspbr_basic},
			{
				Config:                  testAccNspbr_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccNspbr_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNspbrDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccNspbr_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNspbrExist("citrixadc_nspbr.tf_nspbr", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccNspbr_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckNspbrExist("citrixadc_nspbr.tf_nspbr", nil)),
			},
		},
	})
}

// testAccNspbr_unset_step1 sets unset-eligible attributes to non-default values.
// srcmac is a prerequisite for srcmacmask; srcmacmask defaults to "000000000000".
const testAccNspbr_unset_step1 = `
	resource "citrixadc_nspbr" "tf_unset" {
		name       = "tf_test_nspbr_unset"
		action     = "DENY"
		srcmac     = "00:11:22:33:44:55"
		srcmacmask = "000000111111"
	}
`

// testAccNspbr_unset_step2 removes srcmacmask (key + required + its co-prerequisite
// srcmac retained), which must unset srcmacmask back to the NITRO default. NITRO
// rejects setting srcmacmask without srcmac, so srcmac stays configured.
const testAccNspbr_unset_step2 = `
	resource "citrixadc_nspbr" "tf_unset" {
		name   = "tf_test_nspbr_unset"
		action = "DENY"
		srcmac = "00:11:22:33:44:55"
	}
`

func TestAccNspbr_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNspbrDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccNspbr_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNspbrExist("citrixadc_nspbr.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nspbr.tf_unset", "srcmacmask", "000000111111"),
				),
			},
			{
				// Removing the attributes must unset them: state reverts to the
				// documented NITRO defaults, and the implicit post-apply plan is empty.
				Config: testAccNspbr_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNspbrExist("citrixadc_nspbr.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nspbr.tf_unset", "srcmac", "00:11:22:33:44:55"),
					// srcmacmask was unset: NITRO omits it from GET once back at the
					// default ("000000000000"), so state carries no value for it.
					resource.TestCheckNoResourceAttr("citrixadc_nspbr.tf_unset", "srcmacmask"),
					// Independent appliance-level confirmation the unset took effect:
					// the appliance must no longer report the non-default value.
					testAccCheckNspbrADCNotValue("tf_test_nspbr_unset", "srcmacmask", "000000111111"),
				),
			},
		},
	})
}

// testAccCheckNspbrADCNotValue asserts an attribute on the appliance (not just in
// Terraform state) no longer holds a given value, proving the unset reverted it.
// Used for omit-on-default attributes that NITRO drops from GET once at default.
func testAccCheckNspbrADCNotValue(name, attr, notWant string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Nspbr.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("nspbr %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got == notWant {
			return fmt.Errorf("nspbr %s: appliance attr %q still = %q (unset did not revert it)", name, attr, got)
		}
		return nil
	}
}

func TestAccNspbrDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNspbrDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nspbr.tf_nspbr_ds", "name", "tf_test_nspbr_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_nspbr.tf_nspbr_ds", "action", "DENY"),
					resource.TestCheckResourceAttr("data.citrixadc_nspbr.tf_nspbr_ds", "srcipval", "192.0.2.0-192.0.2.255"),
					resource.TestCheckResourceAttr("data.citrixadc_nspbr.tf_nspbr_ds", "destipval", "203.0.113.0-203.0.113.255"),
					resource.TestCheckResourceAttr("data.citrixadc_nspbr.tf_nspbr_ds", "priority", "100"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nspbr.tf_nspbr_ds", "id"),
				),
			},
		},
	})
}
