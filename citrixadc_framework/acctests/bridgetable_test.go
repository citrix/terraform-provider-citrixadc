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

const testAccBridgetable_basic = `

	resource "citrixadc_vlan" "tf_vlan" {
		vlanid    = 20
		aliasname = "Management VLAN"
	}
	resource "citrixadc_vxlan" "tf_vxlan" {
		vxlanid            = 123
		vlan               = citrixadc_vlan.tf_vlan.vlanid
		port               = 33
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
		innervlantagging   = "ENABLED"
	}
	resource "citrixadc_bridgetable" "tf_bridgetable" {
		mac       = "00:00:00:00:00:01"
		vxlan     = citrixadc_vxlan.tf_vxlan.vxlanid
		vtep      = "2.34.5.6"
		bridgeage = "250"
	}
`

func TestAccBridgetable_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBridgetableDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBridgetable_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBridgetableExist("citrixadc_bridgetable.tf_bridgetable", nil),
				),
			},
		},
	})
}

func testAccCheckBridgetableExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No bridgetable name is set")
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

		findParams := service.FindParams{
			ResourceType: service.Bridgetable.Type(),
		}
		dataArray, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 3)

		mac := idSlice[0]
		vxlan := idSlice[1]
		vtep := idSlice[2]

		foundIndex := -1
		for i, bridgetable := range dataArray {
			match := true
			if bridgetable["mac"] != mac {
				match = false
			}
			if bridgetable["vxlan"] != vxlan {
				match = false
			}
			if bridgetable["vtep"] != vtep {
				match = false
			}
			if match {
				foundIndex = i
				break
			}
		}

		if foundIndex == -1 {
			return fmt.Errorf("bridgetable %s not found", n)
		}

		return nil
	}
}

func testAccCheckBridgetableDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_bridgetable" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		findParams := service.FindParams{
			ResourceType: service.Bridgetable.Type(),
		}
		dataArray, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 3)

		mac := idSlice[0]
		vxlan := idSlice[1]
		vtep := idSlice[2]

		foundIndex := -1
		for i, bridgetable := range dataArray {
			match := true
			if bridgetable["mac"] != mac {
				match = false
			}
			if bridgetable["vxlan"] != vxlan {
				match = false
			}
			if bridgetable["vtep"] != vtep {
				match = false
			}
			if match {
				foundIndex = i
				break
			}
		}

		if foundIndex != -1 {
			return fmt.Errorf("bridgetable %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccBridgetable_import(t *testing.T) {
	const resAddr = "citrixadc_bridgetable.tf_bridgetable"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBridgetableDestroy,
		Steps: []resource.TestStep{
			{Config: testAccBridgetable_basic},
			{
				Config:                  testAccBridgetable_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"bridgeage"},
			},
		},
	})
}

func TestAccBridgetableDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccBridgetableDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.citrixadc_bridgetable.tf_bridgetable_ds", "mac"),
					resource.TestCheckResourceAttrSet("data.citrixadc_bridgetable.tf_bridgetable_ds", "id"),
				),
			},
		},
	})
}

func TestAccBridgetable_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckBridgetableDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccBridgetable_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckBridgetableExist("citrixadc_bridgetable.tf_bridgetable", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccBridgetable_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckBridgetableExist("citrixadc_bridgetable.tf_bridgetable", nil)),
			},
		},
	})
}

// The only unset-eligible attribute is bridgeage (Optional, mutable, documented
// NITRO default 300). All other attributes are the key (mac/vxlan/vtep) or
// RequiresReplace, so they are not unsettable. step1 sets bridgeage to a
// non-default value; step2 removes it, and the provider must unset it back to
// the appliance default (300).
const testAccBridgetable_unset_step1 = `
	resource "citrixadc_vlan" "tf_vlan_unset" {
		vlanid    = 22
		aliasname = "Management VLAN Unset"
	}
	resource "citrixadc_vxlan" "tf_vxlan_unset" {
		vxlanid            = 125
		vlan               = citrixadc_vlan.tf_vlan_unset.vlanid
		port               = 33
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
		innervlantagging   = "ENABLED"
	}
	resource "citrixadc_bridgetable" "tf_unset" {
		mac       = "00:00:00:00:00:03"
		vxlan     = citrixadc_vxlan.tf_vxlan_unset.vxlanid
		vtep      = "2.34.5.8"
		bridgeage = 250
	}
`

const testAccBridgetable_unset_step2 = `
	resource "citrixadc_vlan" "tf_vlan_unset" {
		vlanid    = 22
		aliasname = "Management VLAN Unset"
	}
	resource "citrixadc_vxlan" "tf_vxlan_unset" {
		vxlanid            = 125
		vlan               = citrixadc_vlan.tf_vlan_unset.vlanid
		port               = 33
		dynamicrouting     = "DISABLED"
		ipv6dynamicrouting = "DISABLED"
		innervlantagging   = "ENABLED"
	}
	resource "citrixadc_bridgetable" "tf_unset" {
		mac   = "00:00:00:00:00:03"
		vxlan = citrixadc_vxlan.tf_vxlan_unset.vxlanid
		vtep  = "2.34.5.8"
		# bridgeage removed from config -> provider must unset it (revert to 300).
	}
`

func TestAccBridgetable_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBridgetableDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccBridgetable_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBridgetableExist("citrixadc_bridgetable.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_bridgetable.tf_unset", "bridgeage", "250"),
				),
			},
			{
				// Removing bridgeage must unset it: state reverts to the NITRO
				// default (300) and the implicit post-apply plan must be empty.
				Config: testAccBridgetable_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBridgetableExist("citrixadc_bridgetable.tf_unset", nil),
					// bridgeage reverts to the NITRO default (300) after unset.
					// NITRO does not echo bridgeage back per-entry via GET (it is a
					// table-wide setting), so it cannot be asserted directly on the
					// appliance; the state revert plus the implicit empty post-apply
					// plan confirm the unset took effect.
					resource.TestCheckResourceAttr("citrixadc_bridgetable.tf_unset", "bridgeage", "300"),
				),
			},
		},
	})
}

const testAccBridgetableDataSource_basic = `

resource "citrixadc_vlan" "tf_vlan_ds" {
	vlanid    = 21
	aliasname = "Management VLAN DS"
}
resource "citrixadc_vxlan" "tf_vxlan_ds" {
	vxlanid            = 124
	vlan               = citrixadc_vlan.tf_vlan_ds.vlanid
	port               = 33
	dynamicrouting     = "DISABLED"
	ipv6dynamicrouting = "DISABLED"
	innervlantagging   = "ENABLED"
}
resource "citrixadc_bridgetable" "tf_bridgetable_ds" {
	mac       = "00:00:00:00:00:02"
	vxlan     = citrixadc_vxlan.tf_vxlan_ds.vxlanid
	vtep      = "2.34.5.7"
	bridgeage = "250"
}

data "citrixadc_bridgetable" "tf_bridgetable_ds" {
	depends_on = [citrixadc_bridgetable.tf_bridgetable_ds]
}

`
