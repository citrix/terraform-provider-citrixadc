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

const testAccIptunnelparam_add = `
	resource "citrixadc_iptunnelparam" "tf_iptunnelparam" {
		dropfrag             = "YES"
		dropfragcputhreshold = 1
		srciproundrobin      = "NO"
		enablestrictrx       = "YES"
		enablestricttx       = "YES"
		useclientsourceip    = "YES"
	}
`
const testAccIptunnelparam_update = `
	resource "citrixadc_iptunnelparam" "tf_iptunnelparam" {
		dropfrag             = "NO"
		dropfragcputhreshold = 1
		srciproundrobin      = "NO"
		enablestrictrx       = "NO"
		enablestricttx       = "NO"
		useclientsourceip    = "NO"
	}
`

func TestAccIptunnelparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccIptunnelparam_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIptunnelparamExist("citrixadc_iptunnelparam.tf_iptunnelparam", nil),
					resource.TestCheckResourceAttr("citrixadc_iptunnelparam.tf_iptunnelparam", "dropfrag", "YES"),
					resource.TestCheckResourceAttr("citrixadc_iptunnelparam.tf_iptunnelparam", "enablestrictrx", "YES"),
				),
			},
			{
				Config: testAccIptunnelparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIptunnelparamExist("citrixadc_iptunnelparam.tf_iptunnelparam", nil),
					resource.TestCheckResourceAttr("citrixadc_iptunnelparam.tf_iptunnelparam", "dropfrag", "NO"),
					resource.TestCheckResourceAttr("citrixadc_iptunnelparam.tf_iptunnelparam", "enablestrictrx", "NO"),
				),
			},
		},
	})
}

func TestAccIptunnelparam_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccIptunnelparam_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIptunnelparamExist("citrixadc_iptunnelparam.tf_iptunnelparam", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccIptunnelparam_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIptunnelparamExist("citrixadc_iptunnelparam.tf_iptunnelparam", nil),
				),
			},
		},
	})
}

func testAccCheckIptunnelparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No iptunnelparam name is set")
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
		data, err := client.FindResource(service.Iptunnelparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("iptunnelparam %s not found", n)
		}

		return nil
	}
}

// iptunnelparam is a singleton. Step 1 sets the unset-eligible attributes to
// non-default values; step 2 removes them so the provider must unset them,
// reverting the appliance to the documented NITRO defaults ("NO"/0).
const testAccIptunnelparam_unset_step1 = `
	resource "citrixadc_iptunnelparam" "tf_unset" {
		dropfrag             = "YES"
		dropfragcputhreshold = 50
		srciproundrobin      = "YES"
		enablestrictrx       = "YES"
		enablestricttx       = "YES"
		useclientsourceip    = "YES"
	}
`

const testAccIptunnelparam_unset_step2 = `
	resource "citrixadc_iptunnelparam" "tf_unset" {
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccIptunnelparam_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccIptunnelparam_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIptunnelparamExist("citrixadc_iptunnelparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_iptunnelparam.tf_unset", "dropfrag", "YES"),
					resource.TestCheckResourceAttr("citrixadc_iptunnelparam.tf_unset", "dropfragcputhreshold", "50"),
					resource.TestCheckResourceAttr("citrixadc_iptunnelparam.tf_unset", "srciproundrobin", "YES"),
					resource.TestCheckResourceAttr("citrixadc_iptunnelparam.tf_unset", "enablestrictrx", "YES"),
					resource.TestCheckResourceAttr("citrixadc_iptunnelparam.tf_unset", "enablestricttx", "YES"),
					resource.TestCheckResourceAttr("citrixadc_iptunnelparam.tf_unset", "useclientsourceip", "YES"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccIptunnelparam_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIptunnelparamExist("citrixadc_iptunnelparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_iptunnelparam.tf_unset", "dropfrag", "NO"),
					resource.TestCheckResourceAttr("citrixadc_iptunnelparam.tf_unset", "dropfragcputhreshold", "0"),
					resource.TestCheckResourceAttr("citrixadc_iptunnelparam.tf_unset", "srciproundrobin", "NO"),
					resource.TestCheckResourceAttr("citrixadc_iptunnelparam.tf_unset", "enablestrictrx", "NO"),
					resource.TestCheckResourceAttr("citrixadc_iptunnelparam.tf_unset", "enablestricttx", "NO"),
					resource.TestCheckResourceAttr("citrixadc_iptunnelparam.tf_unset", "useclientsourceip", "NO"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckIptunnelparamADCValue("dropfrag", "NO"),
					testAccCheckIptunnelparamADCValue("useclientsourceip", "NO"),
				),
			},
		},
	})
}

// testAccCheckIptunnelparamADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckIptunnelparamADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Iptunnelparam.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("iptunnelparam not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("iptunnelparam: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccIptunnelparam_import(t *testing.T) {
	const resAddr = "citrixadc_iptunnelparam.tf_iptunnelparam"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccIptunnelparam_add},
			{
				Config:                  testAccIptunnelparam_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccIptunnelparamDataSource_basic = `

	resource "citrixadc_iptunnelparam" "tf_iptunnelparam" {
		dropfrag             = "YES"
		dropfragcputhreshold = 1
		srciproundrobin      = "NO"
		enablestrictrx       = "YES"
		enablestricttx       = "YES"
		useclientsourceip    = "YES"
	}

	data "citrixadc_iptunnelparam" "tf_iptunnelparam" {
		depends_on = [citrixadc_iptunnelparam.tf_iptunnelparam]
	}
`

func TestAccIptunnelparamDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccIptunnelparamDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_iptunnelparam.tf_iptunnelparam", "dropfrag", "YES"),
					resource.TestCheckResourceAttr("data.citrixadc_iptunnelparam.tf_iptunnelparam", "dropfragcputhreshold", "1"),
					resource.TestCheckResourceAttr("data.citrixadc_iptunnelparam.tf_iptunnelparam", "srciproundrobin", "NO"),
					resource.TestCheckResourceAttr("data.citrixadc_iptunnelparam.tf_iptunnelparam", "enablestrictrx", "YES"),
					resource.TestCheckResourceAttr("data.citrixadc_iptunnelparam.tf_iptunnelparam", "enablestricttx", "YES"),
					resource.TestCheckResourceAttr("data.citrixadc_iptunnelparam.tf_iptunnelparam", "useclientsourceip", "YES"),
				),
			},
		},
	})
}
