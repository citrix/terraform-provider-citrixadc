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

const testAccSubscribergxinterface_basic = `

	resource "citrixadc_service" "tf_service" {
		name 		= "pcrf-svc1"
		port 		= 80
		ip 			= "10.202.22.12"
		servicetype = "HTTP"
	}
	resource "citrixadc_subscribergxinterface" "tf_subscribergxinterface" {
		service        = citrixadc_service.tf_service.name
		pcrfrealm      = "myrealm.com"
		healthcheck    = "YES"
		servicepathavp = [26009]
		healthcheckttl = 30
	}
`

const testAccSubscribergxinterfaceDataSource_basic = `

	resource "citrixadc_service" "tf_service" {
		name 		= "pcrf-svc1"
		port 		= 3868
		ip 			= "10.202.22.12"
		servicetype = "DIAMETER"
	}
	resource "citrixadc_subscribergxinterface" "tf_subscribergxinterface" {
		service        = citrixadc_service.tf_service.name
		pcrfrealm      = "myrealm.com"
		healthcheck    = "YES"
		servicepathavp = [26009]
		healthcheckttl = 30
	}

	data "citrixadc_subscribergxinterface" "tf_subscribergxinterface" {
		depends_on = [citrixadc_subscribergxinterface.tf_subscribergxinterface]
	}
`
const testAccSubscribergxinterface_update = `

	resource "citrixadc_service" "tf_service" {
		name 		= "pcrf-svc1"
		port 		= 80
		ip 			= "10.202.22.12"
		servicetype = "HTTP"
	}

	resource "citrixadc_subscribergxinterface" "tf_subscribergxinterface" {
		service        = citrixadc_service.tf_service.name
		pcrfrealm      = "myrealm2.com"
		healthcheck    = "NO"
		servicepathavp = [26010]
		healthcheckttl = 40
	}
`

func TestAccSubscribergxinterface_basic(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSubscribergxinterface_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubscribergxinterfaceExist("citrixadc_subscribergxinterface.tf_subscribergxinterface", nil),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_subscribergxinterface", "service", "pcrf-svc1"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_subscribergxinterface", "pcrfrealm", "myrealm.com"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_subscribergxinterface", "healthcheck", "YES"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_subscribergxinterface", "healthcheckttl", "30"),
				),
			},
			{
				Config: testAccSubscribergxinterface_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubscribergxinterfaceExist("citrixadc_subscribergxinterface.tf_subscribergxinterface", nil),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_subscribergxinterface", "service", "pcrf-svc1"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_subscribergxinterface", "pcrfrealm", "myrealm2.com"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_subscribergxinterface", "healthcheck", "NO"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_subscribergxinterface", "healthcheckttl", "40"),
				),
			},
		},
	})
}

func TestAccSubscribergxinterface_import(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	const resAddr = "citrixadc_subscribergxinterface.tf_subscribergxinterface"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccSubscribergxinterface_basic},
			{
				Config:                  testAccSubscribergxinterface_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckSubscribergxinterfaceExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No subscribergxinterface name is set")
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
		data, err := client.FindResource("subscribergxinterface", "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("subscribergxinterface %s not found", n)
		}

		return nil
	}
}

func TestAccSubscribergxinterface_sdkv2StateUpgrade(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccSubscribergxinterface_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSubscribergxinterfaceExist("citrixadc_subscribergxinterface.tf_subscribergxinterface", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccSubscribergxinterface_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSubscribergxinterfaceExist("citrixadc_subscribergxinterface.tf_subscribergxinterface", nil)),
			},
		},
	})
}

// The subscribergxinterface unset test covers the unset-eligible attributes:
// every attribute wired into attributesToUnset (those the NITRO spec lists in the
// unset payload AND that have a documented server default). service/vserver/
// servicepathavp/servicepathvendorid/pcrfrealm/nodeid are excluded (no documented
// unset default or not in the unset payload).
const testAccSubscribergxinterface_unset_step1 = `
	resource "citrixadc_subscribergxinterface" "tf_unset" {
		cerrequesttimeout         = 10
		healthcheck               = "YES"
		healthcheckttl            = 60
		holdonsubscriberabsence   = "NO"
		idlettl                   = 1000
		negativettl               = 700
		negativettllimitedsuccess = "YES"
		purgesdbongxfailure       = "YES"
		requestretryattempts      = 5
		requesttimeout            = 20
		revalidationtimeout       = 100
	}
`

const testAccSubscribergxinterface_unset_step2 = `
	resource "citrixadc_subscribergxinterface" "tf_unset" {
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to their documented NITRO defaults).
	}
`

func TestAccSubscribergxinterface_unset(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccSubscribergxinterface_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubscribergxinterfaceExist("citrixadc_subscribergxinterface.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_unset", "cerrequesttimeout", "10"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_unset", "healthcheck", "YES"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_unset", "healthcheckttl", "60"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_unset", "holdonsubscriberabsence", "NO"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_unset", "idlettl", "1000"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_unset", "negativettl", "700"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_unset", "negativettllimitedsuccess", "YES"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_unset", "purgesdbongxfailure", "YES"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_unset", "requestretryattempts", "5"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_unset", "requesttimeout", "20"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_unset", "revalidationtimeout", "100"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to the documented NITRO defaults, and the implicit
				// post-apply plan must be empty.
				Config: testAccSubscribergxinterface_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubscribergxinterfaceExist("citrixadc_subscribergxinterface.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_unset", "cerrequesttimeout", "0"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_unset", "healthcheck", "NO"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_unset", "healthcheckttl", "30"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_unset", "holdonsubscriberabsence", "YES"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_unset", "idlettl", "900"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_unset", "negativettl", "600"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_unset", "negativettllimitedsuccess", "NO"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_unset", "purgesdbongxfailure", "NO"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_unset", "requestretryattempts", "3"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_unset", "requesttimeout", "10"),
					resource.TestCheckResourceAttr("citrixadc_subscribergxinterface.tf_unset", "revalidationtimeout", "0"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckSubscribergxinterfaceADCValue("healthcheck", "NO"),
					testAccCheckSubscribergxinterfaceADCValue("holdonsubscriberabsence", "YES"),
					testAccCheckSubscribergxinterfaceADCValue("requesttimeout", "10"),
				),
			},
		},
	})
}

// testAccCheckSubscribergxinterfaceADCValue asserts an attribute's value directly
// on the appliance (not just in Terraform state), proving the unset actually
// reverted it. subscribergxinterface is a singleton, so it is read with an empty
// resource name.
func testAccCheckSubscribergxinterfaceADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Subscribergxinterface.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("subscribergxinterface not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("subscribergxinterface: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccSubscribergxinterfaceDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSubscribergxinterfaceDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_subscribergxinterface.tf_subscribergxinterface", "service", "pcrf-svc1"),
					resource.TestCheckResourceAttr("data.citrixadc_subscribergxinterface.tf_subscribergxinterface", "pcrfrealm", "myrealm.com"),
					resource.TestCheckResourceAttr("data.citrixadc_subscribergxinterface.tf_subscribergxinterface", "healthcheck", "YES"),
					resource.TestCheckResourceAttr("data.citrixadc_subscribergxinterface.tf_subscribergxinterface", "healthcheckttl", "30"),
				),
			},
		},
	})
}
