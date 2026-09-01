/*
Copyright 2021 Citrix Systems, Inc

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

const testAccLbgroup_basic = `
# The cookiedomain, rule and usevserverpersistency properties variabled cannot
# be updated and so they were deliberately left out of the test suite.

resource "citrixadc_lbgroup" "tf_lbgroup" {
	name = "tf_lbgroup"
	persistencetype = "RULE"
	persistencebackup = "SOURCEIP"
	backuppersistencetimeout = 10.0
	persistmask = "255.255.255.0"
	v6persistmasklen = 64
	timeout = 10.0
}
`

const testAccLbgroup_update_properties = `
resource "citrixadc_lbgroup" "tf_lbgroup" {
	name = "tf_lbgroup"
	persistencetype = "COOKIEINSERT"
	persistencebackup = "SOURCEIP"
	backuppersistencetimeout = 15.0
	persistmask = "255.255.254.0"
	cookiename = "tf_cookie_1"
	v6persistmasklen = 96
	timeout = 15.0
}
`

const testAccLbgroupDataSource_basic = `
resource "citrixadc_lbgroup" "tf_lbgroup" {
	name = "tf_lbgroup_ds"
	persistencetype = "COOKIEINSERT"
	persistencebackup = "SOURCEIP"
	backuppersistencetimeout = 10
	persistmask = "255.255.255.0"
	cookiename = "test_cookie"
	v6persistmasklen = 64
	timeout = 10
}

data "citrixadc_lbgroup" "tf_lbgroup" {
	name = citrixadc_lbgroup.tf_lbgroup.name
}
`

func TestAccLbgroup_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbgroupDestroy,
		Steps: []resource.TestStep{
			// create Lbgroup
			{
				Config: testAccLbgroup_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbgroupExist("citrixadc_lbgroup.tf_lbgroup", nil),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_lbgroup", "persistencetype", "RULE"),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_lbgroup", "persistencebackup", "SOURCEIP"),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_lbgroup", "backuppersistencetimeout", "10"),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_lbgroup", "persistmask", "255.255.255.0"),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_lbgroup", "v6persistmasklen", "64"),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_lbgroup", "timeout", "10"),
					testAccCheckUserAgent(),
				),
			},
			// update Lbgroup properties
			{
				Config: testAccLbgroup_update_properties,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbgroupExist("citrixadc_lbgroup.tf_lbgroup", nil),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_lbgroup", "persistencetype", "COOKIEINSERT"),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_lbgroup", "persistencebackup", "SOURCEIP"),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_lbgroup", "backuppersistencetimeout", "15"),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_lbgroup", "persistmask", "255.255.254.0"),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_lbgroup", "cookiename", "tf_cookie_1"),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_lbgroup", "v6persistmasklen", "96"),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_lbgroup", "timeout", "15"),
					testAccCheckUserAgent(),
				),
			},
		},
	})
}

func TestAccLbgroupDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLbgroupDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					// id is the universal runtime-binding proof of a resolved data source.
					resource.TestCheckResourceAttrSet("data.citrixadc_lbgroup.tf_lbgroup", "id"),
					resource.TestCheckResourceAttr("data.citrixadc_lbgroup.tf_lbgroup", "name", "tf_lbgroup_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_lbgroup.tf_lbgroup", "persistencetype", "COOKIEINSERT"),
					resource.TestCheckResourceAttr("data.citrixadc_lbgroup.tf_lbgroup", "persistencebackup", "SOURCEIP"),
					resource.TestCheckResourceAttr("data.citrixadc_lbgroup.tf_lbgroup", "backuppersistencetimeout", "10"),
					resource.TestCheckResourceAttr("data.citrixadc_lbgroup.tf_lbgroup", "persistmask", "255.255.255.0"),
					resource.TestCheckResourceAttr("data.citrixadc_lbgroup.tf_lbgroup", "cookiename", "test_cookie"),
					resource.TestCheckResourceAttr("data.citrixadc_lbgroup.tf_lbgroup", "v6persistmasklen", "64"),
					resource.TestCheckResourceAttr("data.citrixadc_lbgroup.tf_lbgroup", "timeout", "10"),
				),
			},
		},
	})
}

func TestAccLbgroup_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_lbgroup.tf_lbgroup"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbgroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbgroup_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLbgroupExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Lbgroup.Type(), "tf_lbgroup"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccLbgroup_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLbgroupExist(resAddr, nil)),
			},
		},
	})
}

func TestAccLbgroup_import(t *testing.T) {
	const resAddr = "citrixadc_lbgroup.tf_lbgroup"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbgroupDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLbgroup_basic},
			{
				Config:                  testAccLbgroup_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccLbgroup_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLbgroupDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccLbgroup_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLbgroupExist("citrixadc_lbgroup.tf_lbgroup", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccLbgroup_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLbgroupExist("citrixadc_lbgroup.tf_lbgroup", nil)),
			},
		},
	})
}

// testAccLbgroup_unset covers the unset-eligible attributes with documented
// NITRO server defaults: backuppersistencetimeout (2), timeout (2),
// v6persistmasklen (128) and usevserverpersistency (DISABLED). Step 1 sets them
// to non-default values; step 2 removes them so the provider must unset them
// (revert to the NITRO defaults). persistencetype/persistencebackup are kept as
// prerequisites in both steps.
const testAccLbgroup_unset_step1 = `
resource "citrixadc_lbgroup" "tf_unset" {
	name                     = "tf_lbgroup_unset"
	persistencetype          = "RULE"
	persistencebackup        = "SOURCEIP"
	backuppersistencetimeout = 20
	timeout                  = 20
	v6persistmasklen         = 64
}
`

const testAccLbgroup_unset_step2 = `
resource "citrixadc_lbgroup" "tf_unset" {
	name              = "tf_lbgroup_unset"
	persistencetype   = "RULE"
	persistencebackup = "SOURCEIP"
	# All unset-eligible attributes removed from config -> the provider must
	# unset them (revert to NITRO defaults).
}
`

func TestAccLbgroup_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbgroupDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccLbgroup_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbgroupExist("citrixadc_lbgroup.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_unset", "backuppersistencetimeout", "20"),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_unset", "timeout", "20"),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_unset", "v6persistmasklen", "64"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccLbgroup_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbgroupExist("citrixadc_lbgroup.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_unset", "backuppersistencetimeout", "2"),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_unset", "timeout", "2"),
					resource.TestCheckResourceAttr("citrixadc_lbgroup.tf_unset", "v6persistmasklen", "128"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckLbgroupADCValue("tf_lbgroup_unset", "timeout", "2"),
					testAccCheckLbgroupADCValue("tf_lbgroup_unset", "backuppersistencetimeout", "2"),
				),
			},
		},
	})
}

// testAccCheckLbgroupADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckLbgroupADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Lbgroup.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("lbgroup %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("lbgroup %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func testAccCheckLbgroupExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Lbgroup name is set")
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
		data, err := client.FindResource("lbgroup", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lbgroup %s not found", n)
		}

		return nil
	}
}

func testAccCheckLbgroupDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbgroup" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lbgroup", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lbgroup %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
