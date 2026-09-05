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
	//"strconv"
)

const testAccRoute6_basic = `

resource "citrixadc_vlan" "tf_vlan" {
    vlanid = 2
}
resource "citrixadc_route6" "tf_route6" {
	network  = "2001:db8:85a3::/64"
	vlan     = citrixadc_vlan.tf_vlan.vlanid
	weight   = 5
	distance = 3
	}
  
`
const testAccRoute6_update = `

resource "citrixadc_vlan" "tf_vlan" {
    vlanid = 2
}
resource "citrixadc_route6" "tf_route6" {
	network  = "2001:db8:85a3::/64"
	vlan     = citrixadc_vlan.tf_vlan.vlanid
	weight   = 6
	distance = 4
	}
  
`

func TestAccRoute6_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRoute6Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRoute6_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoute6Exist("citrixadc_route6.tf_route6", nil),
					resource.TestCheckResourceAttr("citrixadc_route6.tf_route6", "vlan", "2"),
					resource.TestCheckResourceAttr("citrixadc_route6.tf_route6", "weight", "5"),
					resource.TestCheckResourceAttr("citrixadc_route6.tf_route6", "distance", "3"),
				),
			},
			{
				Config: testAccRoute6_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoute6Exist("citrixadc_route6.tf_route6", nil),
					resource.TestCheckResourceAttr("citrixadc_route6.tf_route6", "vlan", "2"),
					resource.TestCheckResourceAttr("citrixadc_route6.tf_route6", "weight", "6"),
					resource.TestCheckResourceAttr("citrixadc_route6.tf_route6", "distance", "4"),
				),
			},
		},
	})
}

func TestAccRoute6_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_route6.tf_route6"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRoute6Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRoute6_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckRoute6Exist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgs(service.Route6.Type(), "", []string{"network:2001%3Adb8%3A85a3%3A%3A%2F64", "vlan:2"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccRoute6_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckRoute6Exist(resAddr, nil)),
			},
		},
	})
}

func TestAccRoute6_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckRoute6Destroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccRoute6_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckRoute6Exist("citrixadc_route6.tf_route6", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccRoute6_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckRoute6Exist("citrixadc_route6.tf_route6", nil)),
			},
		},
	})
}

// testAccRoute6_unset_step1 sets the unset-eligible mutable attributes to valid
// NON-default values.
const testAccRoute6_unset_step1 = `
resource "citrixadc_vlan" "tf_vlan" {
    vlanid = 3
}
resource "citrixadc_route6" "tf_unset" {
	network  = "2001:db8:aaaa::/64"
	vlan     = citrixadc_vlan.tf_vlan.vlanid
	weight   = 5
	distance = 3
	cost     = 7
	msr      = "ENABLED"
}
`

// testAccRoute6_unset_step2 removes the unset-eligible attributes (keeping only
// the key + route-identity attrs); the provider must NITRO-unset them so they
// revert to the documented defaults.
const testAccRoute6_unset_step2 = `
resource "citrixadc_vlan" "tf_vlan" {
    vlanid = 3
}
resource "citrixadc_route6" "tf_unset" {
	network = "2001:db8:aaaa::/64"
	vlan    = citrixadc_vlan.tf_vlan.vlanid
}
`

func TestAccRoute6_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRoute6Destroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccRoute6_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoute6Exist("citrixadc_route6.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_route6.tf_unset", "weight", "5"),
					resource.TestCheckResourceAttr("citrixadc_route6.tf_unset", "distance", "3"),
					resource.TestCheckResourceAttr("citrixadc_route6.tf_unset", "cost", "7"),
					resource.TestCheckResourceAttr("citrixadc_route6.tf_unset", "msr", "ENABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state reverts to the
				// documented NITRO defaults and the implicit post-apply plan is empty.
				Config: testAccRoute6_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRoute6Exist("citrixadc_route6.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_route6.tf_unset", "weight", "1"),
					resource.TestCheckResourceAttr("citrixadc_route6.tf_unset", "distance", "1"),
					resource.TestCheckResourceAttr("citrixadc_route6.tf_unset", "cost", "1"),
					resource.TestCheckResourceAttr("citrixadc_route6.tf_unset", "msr", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckRoute6ADCValue("2001:db8:aaaa::/64", "3", "weight", "1"),
					testAccCheckRoute6ADCValue("2001:db8:aaaa::/64", "3", "msr", "DISABLED"),
				),
			},
		},
	})
}

// testAccCheckRoute6ADCValue asserts an attribute's value directly on the
// appliance (route6 is unnamed, so match on network + vlan), proving the unset
// actually reverted it.
func testAccCheckRoute6ADCValue(network, vlan, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		dataArr, err := client.FindAllResources(service.Route6.Type())
		if err != nil {
			return err
		}
		for _, v := range dataArr {
			if fmt.Sprintf("%v", v["network"]) == network && fmt.Sprintf("%v", v["vlan"]) == vlan {
				got := strings.TrimSpace(fmt.Sprintf("%v", v[attr]))
				if got != want {
					return fmt.Errorf("route6 %s (vlan %s): appliance attr %q = %q, want %q (unset did not revert it)", network, vlan, attr, got, want)
				}
				return nil
			}
		}
		return fmt.Errorf("route6 %s (vlan %s) not found on appliance", network, vlan)
	}
}

func testAccCheckRoute6Exist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No route6 name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}
		route6Network := rs.Primary.ID
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		dataArr, err := client.FindAllResources(service.Route6.Type())
		found := false
		for _, v := range dataArr {
			if v["network"] == route6Network &&
				v["vlan"] == rs.Primary.Attributes["vlan"] {
				found = true
				break
			}
		}

		if err != nil {
			return err
		}

		if !found {
			return fmt.Errorf("route6 %s not found", n)
		}

		return nil
	}
}

func testAccCheckRoute6Destroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_route6" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		route6Network := rs.Primary.ID
		dataArr, err := client.FindAllResources(service.Route6.Type())
		if err != nil {
			return err
		}
		found := false
		for _, v := range dataArr {
			if v["network"] == route6Network &&
				v["vlan"] == rs.Primary.Attributes["vlan"] {
				found = true
				break
			}
		}
		if found {
			return fmt.Errorf("route6 %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccRoute6DataSource_basic = `

resource "citrixadc_vlan" "tf_vlan" {
    vlanid = 2
}

resource "citrixadc_route6" "tf_route6" {
	network  = "2001:db8:85a3::/64"
	vlan     = citrixadc_vlan.tf_vlan.vlanid
	weight   = 5
	distance = 3
	td       = 0
}

data "citrixadc_route6" "tf_route6" {
	network = citrixadc_route6.tf_route6.network
	td      = citrixadc_route6.tf_route6.td
	depends_on = [citrixadc_route6.tf_route6]
}
`

func TestAccRoute6DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRoute6Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRoute6DataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_route6.tf_route6", "network", "2001:db8:85a3::/64"),
					resource.TestCheckResourceAttr("data.citrixadc_route6.tf_route6", "vlan", "2"),
					resource.TestCheckResourceAttr("data.citrixadc_route6.tf_route6", "weight", "5"),
					resource.TestCheckResourceAttr("data.citrixadc_route6.tf_route6", "distance", "3"),
					// Universal runtime-binding proof that the data source resolved.
					resource.TestCheckResourceAttrSet("data.citrixadc_route6.tf_route6", "id"),
				),
			},
		},
	})
}
