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

const testAccLsngroup_basic = `

	resource "citrixadc_lsnclient" "tf_lsnclient" {
		clientname = "my_lsnclient"
	}

	resource "citrixadc_lsngroup" "tf_lsngroup" {
		groupname     = "my_lsngroup"
		clientname    = resource.citrixadc_lsnclient.tf_lsnclient.clientname
		logging       = "DISABLED"
		nattype       = "DYNAMIC"
		snmptraplimit = 50
	}
`
const testAccLsngroup_update = `

	resource "citrixadc_lsnclient" "tf_lsnclient" {
		clientname = "my_lsnclient"
	}

	resource "citrixadc_lsngroup" "tf_lsngroup" {
		groupname     = "my_lsngroup"
		clientname    = resource.citrixadc_lsnclient.tf_lsnclient.clientname
		nattype       = "DYNAMIC"
		snmptraplimit = 100
	}
`

func TestAccLsngroup_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsngroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsngroup_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsngroupExist("citrixadc_lsngroup.tf_lsngroup", nil),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_lsngroup", "groupname", "my_lsngroup"),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_lsngroup", "clientname", "my_lsnclient"),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_lsngroup", "logging", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_lsngroup", "nattype", "DYNAMIC"),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_lsngroup", "snmptraplimit", "50"),
				),
			},
			{
				Config: testAccLsngroup_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsngroupExist("citrixadc_lsngroup.tf_lsngroup", nil),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_lsngroup", "groupname", "my_lsngroup"),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_lsngroup", "clientname", "my_lsnclient"),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_lsngroup", "nattype", "DYNAMIC"),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_lsngroup", "snmptraplimit", "100"),
				),
			},
		},
	})
}

func TestAccLsngroup_import(t *testing.T) {
	const resAddr = "citrixadc_lsngroup.tf_lsngroup"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsngroupDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLsngroup_basic},
			{
				Config:                  testAccLsngroup_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckLsngroupExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsngroup name is set")
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
		data, err := client.FindResource("lsngroup", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lsngroup %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsngroupDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsngroup" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lsngroup", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lsngroup %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

// testAccLsngroup_unset_step1 sets the unset-eligible (non-ForceNew) attributes
// to valid NON-default values.
const testAccLsngroup_unset_step1 = `

	resource "citrixadc_lsnclient" "tf_lsnclient_unset" {
		clientname = "tf_lsnclient_unset"
	}

	resource "citrixadc_lsngroup" "tf_unset" {
		groupname      = "tf_lsngroup_unset"
		clientname     = citrixadc_lsnclient.tf_lsnclient_unset.clientname
		nattype        = "DYNAMIC"
		ftp           = "DISABLED"
		ftpcm         = "ENABLED"
		pptp          = "ENABLED"
		rtspalg       = "ENABLED"
		sessionsync   = "DISABLED"
		sipalg        = "ENABLED"
		snmptraplimit = 50
	}
`

// testAccLsngroup_unset_step2 removes those attributes (key + required only) so
// the provider must unset them (revert to NITRO defaults).
const testAccLsngroup_unset_step2 = `

	resource "citrixadc_lsnclient" "tf_lsnclient_unset" {
		clientname = "tf_lsnclient_unset"
	}

	resource "citrixadc_lsngroup" "tf_unset" {
		groupname  = "tf_lsngroup_unset"
		clientname = citrixadc_lsnclient.tf_lsnclient_unset.clientname
		nattype    = "DYNAMIC"
	}
`

func TestAccLsngroup_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsngroupDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccLsngroup_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsngroupExist("citrixadc_lsngroup.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_unset", "ftp", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_unset", "ftpcm", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_unset", "pptp", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_unset", "rtspalg", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_unset", "sessionsync", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_unset", "sipalg", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_unset", "snmptraplimit", "50"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccLsngroup_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsngroupExist("citrixadc_lsngroup.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_unset", "ftp", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_unset", "ftpcm", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_unset", "pptp", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_unset", "rtspalg", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_unset", "sessionsync", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_unset", "sipalg", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsngroup.tf_unset", "snmptraplimit", "100"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckLsngroupADCValue("tf_lsngroup_unset", "ftp", "ENABLED"),
					testAccCheckLsngroupADCValue("tf_lsngroup_unset", "sessionsync", "ENABLED"),
					testAccCheckLsngroupADCValue("tf_lsngroup_unset", "snmptraplimit", "100"),
				),
			},
		},
	})
}

// testAccCheckLsngroupADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckLsngroupADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Lsngroup.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("lsngroup %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("lsngroup %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccLsngroup_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_lsngroup.tf_lsngroup"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsngroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsngroup_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsngroupExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Lsngroup.Type(), "my_lsngroup"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccLsngroup_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsngroupExist(resAddr, nil)),
			},
		},
	})
}

func TestAccLsngroup_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLsngroupDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccLsngroup_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsngroupExist("citrixadc_lsngroup.tf_lsngroup", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccLsngroup_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsngroupExist("citrixadc_lsngroup.tf_lsngroup", nil),
				),
			},
		},
	})
}

const testAccLsngroupDataSource_basic = `

	resource "citrixadc_lsnclient" "tf_lsnclient" {
		clientname = "my_lsnclient_ds"
	}

resource "citrixadc_lsngroup" "tf_lsngroup_ds" {
	groupname     = "my_lsngroup_ds"
	clientname    = citrixadc_lsnclient.tf_lsnclient.clientname
	logging       = "DISABLED"
	nattype       = "DYNAMIC"
	snmptraplimit = 50
}

data "citrixadc_lsngroup" "tf_lsngroup_ds" {
	groupname = citrixadc_lsngroup.tf_lsngroup_ds.groupname
}
`

func TestAccLsngroupDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLsngroupDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lsngroup.tf_lsngroup_ds", "groupname", "my_lsngroup_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_lsngroup.tf_lsngroup_ds", "clientname", "my_lsnclient_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_lsngroup.tf_lsngroup_ds", "logging", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_lsngroup.tf_lsngroup_ds", "nattype", "DYNAMIC"),
					resource.TestCheckResourceAttr("data.citrixadc_lsngroup.tf_lsngroup_ds", "snmptraplimit", "50"),
				),
			},
		},
	})
}
