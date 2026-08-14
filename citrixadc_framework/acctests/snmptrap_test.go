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
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"fmt"
	"log"
	"strings"
	"testing"
)

const testAccSnmptrap_basic = `
	resource "citrixadc_snmptrap" "tf_snmptrap" {
		severity        = "Major"
		trapclass       = "specific"
		trapdestination = "192.168.2.2"
	}
`
const testAccSnmptrap_update = `
	resource "citrixadc_snmptrap" "tf_snmptrap" {
		severity        = "Minor"
		trapclass       = "specific"
		trapdestination = "192.168.2.2"
	}
`

func TestAccSnmptrap_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmptrapDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmptrap_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmptrapExist("citrixadc_snmptrap.tf_snmptrap", nil),
					resource.TestCheckResourceAttr("citrixadc_snmptrap.tf_snmptrap", "trapclass", "specific"),
					resource.TestCheckResourceAttr("citrixadc_snmptrap.tf_snmptrap", "trapdestination", "192.168.2.2"),
					resource.TestCheckResourceAttr("citrixadc_snmptrap.tf_snmptrap", "severity", "Major"),
				),
			},
			{
				Config: testAccSnmptrap_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmptrapExist("citrixadc_snmptrap.tf_snmptrap", nil),
					resource.TestCheckResourceAttr("citrixadc_snmptrap.tf_snmptrap", "trapclass", "specific"),
					resource.TestCheckResourceAttr("citrixadc_snmptrap.tf_snmptrap", "trapdestination", "192.168.2.2"),
					resource.TestCheckResourceAttr("citrixadc_snmptrap.tf_snmptrap", "severity", "Minor"),
				),
			},
		},
	})
}

func testAccCheckSnmptrapExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No snmptrap name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		snmptrapId := rs.Primary.ID
		idSlice := strings.SplitN(snmptrapId, ",", 3)

		trapclass := idSlice[0]
		trapdestination := idSlice[1]
		version := idSlice[2]

		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		dataArr, err := client.FindAllResources(service.Snmptrap.Type())

		if err != nil {
			return err
		}

		if len(dataArr) == 0 {
			log.Printf("[WARN] citrix-provider: acceptance test: snmptrap does not exist. Clearing state.")
			return nil
		}

		// if len(dataArray) > 1 {
		// 	return fmt.Errorf("[ERROR] citrix-provider: acceptance test: multiple entries found for snmptrap")
		// }

		found := false
		for _, v := range dataArr {
			if v["trapclass"].(string) == trapclass && v["trapdestination"] == trapdestination && v["version"] == version {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("snmptrap %s not found", n)
		}

		return nil
	}
}

func testAccCheckSnmptrapDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_snmptrap" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}
		snmptrapId := rs.Primary.ID
		idSlice := strings.SplitN(snmptrapId, ",", 3)

		trapclass := idSlice[0]
		trapdestination := idSlice[1]
		version := idSlice[2]

		dataArr, err := client.FindAllResources(service.Snmptrap.Type())

		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if v["trapclass"].(string) == trapclass && v["trapdestination"] == trapdestination && v["version"] == version {
				found = true
				break
			}
		}
		if found {
			return fmt.Errorf("snmptrap %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccSnmptrap_unset_step1 = `
	resource "citrixadc_snmptrap" "tf_unset" {
		trapclass       = "specific"
		trapdestination = "192.168.5.5"
		version         = "V2"
		severity        = "Critical"
		destport        = 1620
		allpartitions   = "ENABLED"
	}
`

// Unsettable attributes removed; only the identity keys remain.
const testAccSnmptrap_unset_step2 = `
	resource "citrixadc_snmptrap" "tf_unset" {
		trapclass       = "specific"
		trapdestination = "192.168.5.5"
		version         = "V2"
	}
`

func TestAccSnmptrap_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmptrapDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccSnmptrap_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmptrapExist("citrixadc_snmptrap.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_snmptrap.tf_unset", "severity", "Critical"),
					resource.TestCheckResourceAttr("citrixadc_snmptrap.tf_unset", "destport", "1620"),
					resource.TestCheckResourceAttr("citrixadc_snmptrap.tf_unset", "allpartitions", "ENABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccSnmptrap_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmptrapExist("citrixadc_snmptrap.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_snmptrap.tf_unset", "severity", "Unknown"),
					resource.TestCheckResourceAttr("citrixadc_snmptrap.tf_unset", "destport", "162"),
					resource.TestCheckResourceAttr("citrixadc_snmptrap.tf_unset", "allpartitions", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckSnmptrapADCValue("specific", "192.168.5.5", "V2", "destport", "162"),
					testAccCheckSnmptrapADCValue("specific", "192.168.5.5", "V2", "allpartitions", "DISABLED"),
				),
			},
		},
	})
}

// testAccCheckSnmptrapADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it. snmptrap is an array/filter resource so it is located by its composite key.
func testAccCheckSnmptrapADCValue(trapclass, trapdestination, version, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		dataArr, err := client.FindAllResources(service.Snmptrap.Type())
		if err != nil {
			return err
		}
		for _, v := range dataArr {
			if v["trapclass"] == trapclass && v["trapdestination"] == trapdestination && v["version"] == version {
				got := strings.TrimSpace(fmt.Sprintf("%v", v[attr]))
				if got != want {
					return fmt.Errorf("snmptrap: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
				}
				return nil
			}
		}
		return fmt.Errorf("snmptrap %s,%s,%s not found on appliance", trapclass, trapdestination, version)
	}
}

func TestAccSnmptrap_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_snmptrap.tf_snmptrap"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmptrapDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmptrap_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSnmptrapExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgs(service.Snmptrap.Type(), "specific", []string{"trapdestination:192.168.2.2", "version:V2"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccSnmptrap_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSnmptrapExist(resAddr, nil)),
			},
		},
	})
}

func TestAccSnmptrap_import(t *testing.T) {
	const resAddr = "citrixadc_snmptrap.tf_snmptrap"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmptrapDestroy,
		Steps: []resource.TestStep{
			{Config: testAccSnmptrap_basic},
			{
				Config:                  testAccSnmptrap_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccSnmptrap_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckSnmptrapDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccSnmptrap_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSnmptrapExist("citrixadc_snmptrap.tf_snmptrap", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccSnmptrap_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSnmptrapExist("citrixadc_snmptrap.tf_snmptrap", nil)),
			},
		},
	})
}

func TestAccSnmptrapDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmptrapDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmptrapDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_snmptrap.tf_snmptrap_ds", "trapclass", "specific"),
					resource.TestCheckResourceAttr("data.citrixadc_snmptrap.tf_snmptrap_ds", "trapdestination", "192.168.2.25"),
					resource.TestCheckResourceAttr("data.citrixadc_snmptrap.tf_snmptrap_ds", "version", "V2"),
					resource.TestCheckResourceAttr("data.citrixadc_snmptrap.tf_snmptrap_ds", "td", "0"),
				),
			},
		},
	})
}

const testAccSnmptrapDataSource_basic = `

resource "citrixadc_snmptrap" "tf_snmptrap_ds" {
	severity        = "Major"
	trapclass       = "specific"
	trapdestination = "192.168.2.25"
	version         = "V2"
	td              = 0
}

data "citrixadc_snmptrap" "tf_snmptrap_ds" {
	trapclass       = citrixadc_snmptrap.tf_snmptrap_ds.trapclass
	trapdestination = citrixadc_snmptrap.tf_snmptrap_ds.trapdestination
	version         = citrixadc_snmptrap.tf_snmptrap_ds.version
	td              = citrixadc_snmptrap.tf_snmptrap_ds.td
	depends_on      = [citrixadc_snmptrap.tf_snmptrap_ds]
}
`
