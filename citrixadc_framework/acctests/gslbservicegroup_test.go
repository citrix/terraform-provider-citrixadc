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

const testAccGslbservicegroup_add = `


resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
	servicegroupname = "test_gslbvservicegroup"
	servicetype      = "HTTP"
	cip              = "DISABLED"
	healthmonitor    = "NO"
	sitename         = citrixadc_gslbsite.site_local.sitename
	}
  resource "citrixadc_gslbsite" "site_local" {
	sitename        = "Site-Local"
	siteipaddress   = "172.31.96.234"
	sessionexchange = "DISABLED"
	sitepassword    = "password123"
	}
  
`

const testAccGslbservicegroup_update = `


resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
	servicegroupname = "test_gslbvservicegroup"
	servicetype      = "HTTP"
	cip              = "ENABLED"
	healthmonitor    = "NO"
	sitename         = citrixadc_gslbsite.site_local.sitename
	}
  resource "citrixadc_gslbsite" "site_local" {
	sitename        = "Site-Local"
	siteipaddress   = "172.31.96.234"
	sessionexchange = "DISABLED"
	sitepassword    = "password123"
	}
  
`

const testAccGslbservicegroupDataSource_basic = `

resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
	servicegroupname = "test_gslbvservicegroup"
	servicetype      = "HTTP"
	cip              = "DISABLED"
	healthmonitor    = "NO"
	sitename         = citrixadc_gslbsite.site_local.sitename
}
resource "citrixadc_gslbsite" "site_local" {
	sitename        = "Site-Local"
	siteipaddress   = "172.31.96.234"
	sessionexchange = "DISABLED"
	sitepassword    = "password123"
}

data "citrixadc_gslbservicegroup" "gslbservicegroup_data" {
	servicegroupname = citrixadc_gslbservicegroup.tf_gslbservicegroup.servicegroupname
}
`

func TestAccGslbservicegroup_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckGslbservicegroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGslbservicegroup_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbservicegroupExist("citrixadc_gslbservicegroup.tf_gslbservicegroup", nil),
					resource.TestCheckResourceAttr("citrixadc_gslbservicegroup.tf_gslbservicegroup", "servicegroupname", "test_gslbvservicegroup"),
					resource.TestCheckResourceAttr("citrixadc_gslbservicegroup.tf_gslbservicegroup", "cip", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_gslbservicegroup.tf_gslbservicegroup", "sitename", "Site-Local"),
				),
			},
			{
				Config: testAccGslbservicegroup_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbservicegroupExist("citrixadc_gslbservicegroup.tf_gslbservicegroup", nil),
					resource.TestCheckResourceAttr("citrixadc_gslbservicegroup.tf_gslbservicegroup", "servicegroupname", "test_gslbvservicegroup"),
					resource.TestCheckResourceAttr("citrixadc_gslbservicegroup.tf_gslbservicegroup", "cip", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_gslbservicegroup.tf_gslbservicegroup", "sitename", "Site-Local"),
				),
			},
		},
	})
}

func TestAccGslbservicegroup_import(t *testing.T) {
	const resAddr = "citrixadc_gslbservicegroup.tf_gslbservicegroup"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckGslbservicegroupDestroy,
		Steps: []resource.TestStep{
			{Config: testAccGslbservicegroup_add},
			{
				Config:                  testAccGslbservicegroup_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccGslbservicegroupDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGslbservicegroupDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					// Universal runtime-binding proof that the data source resolved.
					resource.TestCheckResourceAttrSet("data.citrixadc_gslbservicegroup.gslbservicegroup_data", "id"),
					resource.TestCheckResourceAttr("data.citrixadc_gslbservicegroup.gslbservicegroup_data", "servicegroupname", "test_gslbvservicegroup"),
					resource.TestCheckResourceAttr("data.citrixadc_gslbservicegroup.gslbservicegroup_data", "servicetype", "HTTP"),
					resource.TestCheckResourceAttr("data.citrixadc_gslbservicegroup.gslbservicegroup_data", "cip", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_gslbservicegroup.gslbservicegroup_data", "healthmonitor", "NO"),
					resource.TestCheckResourceAttr("data.citrixadc_gslbservicegroup.gslbservicegroup_data", "sitename", "Site-Local"),
				),
			},
		},
	})
}

func TestAccGslbservicegroup_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_gslbservicegroup.tf_gslbservicegroup"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckGslbservicegroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGslbservicegroup_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckGslbservicegroupExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Gslbservicegroup.Type(), "test_gslbvservicegroup"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccGslbservicegroup_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckGslbservicegroupExist(resAddr, nil)),
			},
		},
	})
}

func TestAccGslbservicegroup_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckGslbservicegroupDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccGslbservicegroup_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckGslbservicegroupExist("citrixadc_gslbservicegroup.tf_gslbservicegroup", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccGslbservicegroup_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckGslbservicegroupExist("citrixadc_gslbservicegroup.tf_gslbservicegroup", nil)),
			},
		},
	})
}

// The gslbservicegroup unset test covers the mutable, documented-default,
// unset-eligible attributes: appflowlog (ENABLED), downstateflush (ENABLED),
// healthmonitor (YES). Step 1 sets them to non-default values; step 2 removes
// them from config so the provider issues the NITRO unset, reverting each to its
// appliance default.
const testAccGslbservicegroup_unset_step1 = `
resource "citrixadc_gslbservicegroup" "tf_unset" {
	servicegroupname = "tf_gslbsg_unset"
	servicetype      = "HTTP"
	sitename         = citrixadc_gslbsite.site_local.sitename
	appflowlog       = "DISABLED"
	downstateflush   = "DISABLED"
	healthmonitor    = "NO"
}
resource "citrixadc_gslbsite" "site_local" {
	sitename        = "Site-Local"
	siteipaddress   = "172.31.96.234"
	sessionexchange = "DISABLED"
	sitepassword    = "password123"
}
`

const testAccGslbservicegroup_unset_step2 = `
resource "citrixadc_gslbservicegroup" "tf_unset" {
	servicegroupname = "tf_gslbsg_unset"
	servicetype      = "HTTP"
	sitename         = citrixadc_gslbsite.site_local.sitename
	# unset-eligible attributes removed -> provider must revert to NITRO defaults.
}
resource "citrixadc_gslbsite" "site_local" {
	sitename        = "Site-Local"
	siteipaddress   = "172.31.96.234"
	sessionexchange = "DISABLED"
	sitepassword    = "password123"
}
`

func TestAccGslbservicegroup_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckGslbservicegroupDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccGslbservicegroup_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbservicegroupExist("citrixadc_gslbservicegroup.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_gslbservicegroup.tf_unset", "appflowlog", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_gslbservicegroup.tf_unset", "downstateflush", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_gslbservicegroup.tf_unset", "healthmonitor", "NO"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccGslbservicegroup_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbservicegroupExist("citrixadc_gslbservicegroup.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_gslbservicegroup.tf_unset", "appflowlog", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_gslbservicegroup.tf_unset", "downstateflush", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_gslbservicegroup.tf_unset", "healthmonitor", "YES"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckGslbservicegroupADCValue("tf_gslbsg_unset", "appflowlog", "ENABLED"),
					testAccCheckGslbservicegroupADCValue("tf_gslbsg_unset", "downstateflush", "ENABLED"),
					testAccCheckGslbservicegroupADCValue("tf_gslbsg_unset", "healthmonitor", "YES"),
				),
			},
		},
	})
}

// testAccCheckGslbservicegroupADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckGslbservicegroupADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Gslbservicegroup.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("gslbservicegroup %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("gslbservicegroup %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func testAccCheckGslbservicegroupExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No gslbservicegroup name is set")
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
		data, err := client.FindResource("gslbservicegroup", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("gslbservicegroup %s not found", n)
		}

		return nil
	}
}

func testAccCheckGslbservicegroupDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_gslbservicegroup" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("gslbservicegroup", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("gslbservicegroup %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
