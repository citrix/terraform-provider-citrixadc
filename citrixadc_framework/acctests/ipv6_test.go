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

const testAccIpv6_basic = `

resource "citrixadc_ipv6" "tf_ipv6" {
	ralearning        = "DISABLED"
	ndbasereachtime   = 4000
	routerredirection = "DISABLED"
	}
  
`
const testAccIpv6_update = `

resource "citrixadc_ipv6" "tf_ipv6" {
	ralearning        = "ENABLED"
	ndbasereachtime   = 4000
	routerredirection = "ENABLED"
	}
`

const testAccIpv6DataSource_basic = `

data "citrixadc_ipv6" "tf_ipv6_ds" {
	td = "0"
}
`

func TestAccIpv6_basic(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6Exist("citrixadc_ipv6.tf_ipv6", nil),
					resource.TestCheckResourceAttr("citrixadc_ipv6.tf_ipv6", "ralearning", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_ipv6.tf_ipv6", "ndbasereachtime", "4000"),
					resource.TestCheckResourceAttr("citrixadc_ipv6.tf_ipv6", "routerredirection", "DISABLED"),
				),
			},
			{
				Config: testAccIpv6_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6Exist("citrixadc_ipv6.tf_ipv6", nil),
					resource.TestCheckResourceAttr("citrixadc_ipv6.tf_ipv6", "ralearning", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_ipv6.tf_ipv6", "ndbasereachtime", "4000"),
					resource.TestCheckResourceAttr("citrixadc_ipv6.tf_ipv6", "routerredirection", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckIpv6Exist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ipv6 name is set")
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
		data, err := client.FindResource(service.Ipv6.Type(), rs.Primary.Attributes["td"])

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("ipv6 %s not found", n)
		}

		return nil
	}
}

func testAccCheckIpv6Destroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_ipv6" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Ipv6.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("ipv6 %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccIpv6_import(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	const resAddr = "citrixadc_ipv6.tf_ipv6"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6Destroy,
		Steps: []resource.TestStep{
			{Config: testAccIpv6_basic},
			{
				Config:                  testAccIpv6_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccIpv6_sdkv2StateUpgrade(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckIpv6Destroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccIpv6_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIpv6Exist("citrixadc_ipv6.tf_ipv6", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccIpv6_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIpv6Exist("citrixadc_ipv6.tf_ipv6", nil)),
			},
		},
	})
}

// testAccIpv6_unset_step1 sets the unset-eligible attributes to valid
// non-default values; step2 removes them so the provider must unset them
// (revert to the documented NITRO defaults).
const testAccIpv6_unset_step1 = `
resource "citrixadc_ipv6" "tf_unset" {
	dodad                = "ENABLED"
	ndbasereachtime      = 4000
	ndretransmissiontime = 2000
	ralearning           = "ENABLED"
	routerredirection    = "ENABLED"
}
`

const testAccIpv6_unset_step2 = `
resource "citrixadc_ipv6" "tf_unset" {
	# All unset-eligible attributes removed from config -> the provider must
	# unset them (revert to NITRO defaults).
}
`

func TestAccIpv6_unset(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpv6Destroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccIpv6_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6Exist("citrixadc_ipv6.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_ipv6.tf_unset", "dodad", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_ipv6.tf_unset", "ndbasereachtime", "4000"),
					resource.TestCheckResourceAttr("citrixadc_ipv6.tf_unset", "ndretransmissiontime", "2000"),
					resource.TestCheckResourceAttr("citrixadc_ipv6.tf_unset", "ralearning", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_ipv6.tf_unset", "routerredirection", "ENABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccIpv6_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpv6Exist("citrixadc_ipv6.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_ipv6.tf_unset", "dodad", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_ipv6.tf_unset", "ndbasereachtime", "30000"),
					resource.TestCheckResourceAttr("citrixadc_ipv6.tf_unset", "ndretransmissiontime", "1000"),
					resource.TestCheckResourceAttr("citrixadc_ipv6.tf_unset", "ralearning", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_ipv6.tf_unset", "routerredirection", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckIpv6ADCValue("0", "dodad", "DISABLED"),
					testAccCheckIpv6ADCValue("0", "ralearning", "DISABLED"),
					testAccCheckIpv6ADCValue("0", "routerredirection", "DISABLED"),
				),
			},
		},
	})
}

// testAccCheckIpv6ADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it. td identifies the ipv6 config entry.
func testAccCheckIpv6ADCValue(td, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Ipv6.Type(), td)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("ipv6 (td %s) not found on appliance", td)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("ipv6 (td %s): appliance attr %q = %q, want %q (unset did not revert it)", td, attr, got, want)
		}
		return nil
	}
}

func TestAccIpv6DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6DataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.citrixadc_ipv6.tf_ipv6_ds", "id"),
					resource.TestCheckResourceAttrSet("data.citrixadc_ipv6.tf_ipv6_ds", "ralearning"),
					resource.TestCheckResourceAttrSet("data.citrixadc_ipv6.tf_ipv6_ds", "ndbasereachtime"),
					resource.TestCheckResourceAttrSet("data.citrixadc_ipv6.tf_ipv6_ds", "routerredirection"),
					resource.TestCheckResourceAttrSet("data.citrixadc_ipv6.tf_ipv6_ds", "td"),
				),
			},
		},
	})
}
