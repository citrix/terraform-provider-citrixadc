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

const testAccSubscriberradiusinterface_basic = `


resource "citrixadc_subscriberradiusinterface" "tf_subscriberradiusinterface" {
	listeningservice     = citrixadc_service.tf_service.name
	radiusinterimasstart = "ENABLED"
	}
  
  resource "citrixadc_service" "tf_service" {
	name        = "srad1"
	port        = 1813
	ip          = "192.0.0.206"
	servicetype = "RADIUSListener"
	}
  
`
const testAccSubscriberradiusinterface_update = `


resource "citrixadc_subscriberradiusinterface" "tf_subscriberradiusinterface" {
	listeningservice     = citrixadc_service.tf_service.name
	radiusinterimasstart = "DISABLED"
	}
  
  resource "citrixadc_service" "tf_service" {
	name        = "srad1"
	port        = 1813
	ip          = "192.0.0.206"
	servicetype = "RADIUSListener"
	}
  
  `

const testAccSubscriberradiusinterfaceDataSource_basic = `


resource "citrixadc_subscriberradiusinterface" "tf_subscriberradiusinterface" {
	listeningservice     = citrixadc_service.tf_service.name
	radiusinterimasstart = "ENABLED"
	}
  
  resource "citrixadc_service" "tf_service" {
	name        = "srad1"
	port        = 1813
	ip          = "192.0.0.206"
	servicetype = "RADIUSListener"
	}

data "citrixadc_subscriberradiusinterface" "tf_subscriberradiusinterface" {
	depends_on = [citrixadc_subscriberradiusinterface.tf_subscriberradiusinterface]
}
`

func TestAccSubscriberradiusinterface_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSubscriberradiusinterface_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubscriberradiusinterfaceExist("citrixadc_subscriberradiusinterface.tf_subscriberradiusinterface", nil),
					resource.TestCheckResourceAttr("citrixadc_subscriberradiusinterface.tf_subscriberradiusinterface", "listeningservice", "srad1"),
					resource.TestCheckResourceAttr("citrixadc_subscriberradiusinterface.tf_subscriberradiusinterface", "radiusinterimasstart", "ENABLED"),
				),
			},
			{
				Config: testAccSubscriberradiusinterface_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubscriberradiusinterfaceExist("citrixadc_subscriberradiusinterface.tf_subscriberradiusinterface", nil),
					resource.TestCheckResourceAttr("citrixadc_subscriberradiusinterface.tf_subscriberradiusinterface", "listeningservice", "srad1"),
					resource.TestCheckResourceAttr("citrixadc_subscriberradiusinterface.tf_subscriberradiusinterface", "radiusinterimasstart", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckSubscriberradiusinterfaceExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No subscriberradiusinterface name is set")
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
		data, err := client.FindResource("subscriberradiusinterface", "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("subscriberradiusinterface %s not found", n)
		}

		return nil
	}
}

func TestAccSubscriberradiusinterface_import(t *testing.T) {
	const resAddr = "citrixadc_subscriberradiusinterface.tf_subscriberradiusinterface"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccSubscriberradiusinterface_basic},
			{
				Config:                  testAccSubscriberradiusinterface_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccSubscriberradiusinterface_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccSubscriberradiusinterface_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSubscriberradiusinterfaceExist("citrixadc_subscriberradiusinterface.tf_subscriberradiusinterface", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccSubscriberradiusinterface_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSubscriberradiusinterfaceExist("citrixadc_subscriberradiusinterface.tf_subscriberradiusinterface", nil)),
			},
		},
	})
}

// radiusinterimasstart is the sole unset-eligible attribute (NITRO default
// "DISABLED"). Step1 sets it to the non-default "ENABLED"; step2 removes it
// from config so the provider must issue a NITRO unset and revert it to
// "DISABLED".
const testAccSubscriberradiusinterface_unset_step1 = `
resource "citrixadc_subscriberradiusinterface" "tf_unset" {
	listeningservice     = citrixadc_service.tf_service.name
	radiusinterimasstart = "ENABLED"
}

resource "citrixadc_service" "tf_service" {
	name        = "srad1"
	port        = 1813
	ip          = "192.0.0.206"
	servicetype = "RADIUSListener"
}
`

const testAccSubscriberradiusinterface_unset_step2 = `
resource "citrixadc_subscriberradiusinterface" "tf_unset" {
	listeningservice = citrixadc_service.tf_service.name
	# radiusinterimasstart removed from config -> provider must unset it
	# (revert to NITRO default "DISABLED").
}

resource "citrixadc_service" "tf_service" {
	name        = "srad1"
	port        = 1813
	ip          = "192.0.0.206"
	servicetype = "RADIUSListener"
}
`

func TestAccSubscriberradiusinterface_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccSubscriberradiusinterface_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubscriberradiusinterfaceExist("citrixadc_subscriberradiusinterface.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_subscriberradiusinterface.tf_unset", "radiusinterimasstart", "ENABLED"),
				),
			},
			{
				// Removing the attribute must unset it: state (read back from the
				// appliance) reverts to the NITRO default, and the implicit
				// post-apply plan must be empty.
				Config: testAccSubscriberradiusinterface_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubscriberradiusinterfaceExist("citrixadc_subscriberradiusinterface.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_subscriberradiusinterface.tf_unset", "radiusinterimasstart", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckSubscriberradiusinterfaceADCValue("radiusinterimasstart", "DISABLED"),
				),
			},
		},
	})
}

// testAccCheckSubscriberradiusinterfaceADCValue asserts an attribute's value
// directly on the appliance (not just in Terraform state), proving the unset
// actually reverted it.
func testAccCheckSubscriberradiusinterfaceADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Subscriberradiusinterface.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("subscriberradiusinterface not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("subscriberradiusinterface: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccSubscriberradiusinterfaceDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSubscriberradiusinterfaceDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_subscriberradiusinterface.tf_subscriberradiusinterface", "listeningservice", "srad1"),
					resource.TestCheckResourceAttr("data.citrixadc_subscriberradiusinterface.tf_subscriberradiusinterface", "radiusinterimasstart", "ENABLED"),
				),
			},
		},
	})
}
