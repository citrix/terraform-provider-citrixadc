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

const testAccSmppparam_basic = `

resource "citrixadc_smppparam" "tf_smppparam" {
	clientmode = "TRANSCEIVER"
	msgqueue   = "OFF"
	addrnpi    = 40
	addrton    = 40
	}
  
`
const testAccSmppparam_update = `

resource "citrixadc_smppparam" "tf_smppparam" {
	clientmode = "TRANSMITTERONLY"
	msgqueue   = "ON"
	addrnpi    = 50
	addrton    = 50
	}
  
`

func TestAccSmppparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSmppparam_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmppparamExist("citrixadc_smppparam.tf_smppparam", nil),
					resource.TestCheckResourceAttr("citrixadc_smppparam.tf_smppparam", "clientmode", "TRANSCEIVER"),
					resource.TestCheckResourceAttr("citrixadc_smppparam.tf_smppparam", "msgqueue", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_smppparam.tf_smppparam", "addrnpi", "40"),
					resource.TestCheckResourceAttr("citrixadc_smppparam.tf_smppparam", "addrton", "40"),
				),
			},
			{
				Config: testAccSmppparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmppparamExist("citrixadc_smppparam.tf_smppparam", nil),
					resource.TestCheckResourceAttr("citrixadc_smppparam.tf_smppparam", "clientmode", "TRANSMITTERONLY"),
					resource.TestCheckResourceAttr("citrixadc_smppparam.tf_smppparam", "msgqueue", "ON"),
					resource.TestCheckResourceAttr("citrixadc_smppparam.tf_smppparam", "addrnpi", "50"),
					resource.TestCheckResourceAttr("citrixadc_smppparam.tf_smppparam", "addrton", "50"),
				),
			},
		},
	})
}

func TestAccSmppparam_import(t *testing.T) {
	const resAddr = "citrixadc_smppparam.tf_smppparam"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccSmppparam_basic},
			{
				Config:                  testAccSmppparam_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckSmppparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No smppparam name is set")
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
		data, err := client.FindResource("smppparam", "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("smppparam %s not found", n)
		}

		return nil
	}
}

const testAccSmppparamDataSource_basic = `

	resource "citrixadc_smppparam" "tf_smppparam" {
		clientmode = "TRANSCEIVER"
		msgqueue   = "OFF"
		addrnpi    = 40
		addrton    = 40
	}

	data "citrixadc_smppparam" "tf_smppparam" {
		depends_on = [citrixadc_smppparam.tf_smppparam]
	}
`

func TestAccSmppparam_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccSmppparam_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSmppparamExist("citrixadc_smppparam.tf_smppparam", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccSmppparam_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSmppparamExist("citrixadc_smppparam.tf_smppparam", nil)),
			},
		},
	})
}

// smppparam is a singleton config resource. All read/write attributes have
// documented NITRO defaults and support the unset operation.
const testAccSmppparam_unset_step1 = `
resource "citrixadc_smppparam" "tf_unset" {
	clientmode   = "TRANSMITTERONLY"
	msgqueue     = "ON"
	msgqueuesize = 5000
	addrnpi      = 50
	addrton      = 50
	addrrange    = "abc*"
}
`

const testAccSmppparam_unset_step2 = `
resource "citrixadc_smppparam" "tf_unset" {
	# All unset-eligible attributes removed from config -> the provider must
	# unset them (revert to NITRO defaults).
}
`

func TestAccSmppparam_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccSmppparam_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmppparamExist("citrixadc_smppparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_smppparam.tf_unset", "clientmode", "TRANSMITTERONLY"),
					resource.TestCheckResourceAttr("citrixadc_smppparam.tf_unset", "msgqueue", "ON"),
					resource.TestCheckResourceAttr("citrixadc_smppparam.tf_unset", "msgqueuesize", "5000"),
					resource.TestCheckResourceAttr("citrixadc_smppparam.tf_unset", "addrnpi", "50"),
					resource.TestCheckResourceAttr("citrixadc_smppparam.tf_unset", "addrton", "50"),
					resource.TestCheckResourceAttr("citrixadc_smppparam.tf_unset", "addrrange", "abc*"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccSmppparam_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSmppparamExist("citrixadc_smppparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_smppparam.tf_unset", "clientmode", "TRANSCEIVER"),
					resource.TestCheckResourceAttr("citrixadc_smppparam.tf_unset", "msgqueue", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_smppparam.tf_unset", "msgqueuesize", "10000"),
					resource.TestCheckResourceAttr("citrixadc_smppparam.tf_unset", "addrnpi", "0"),
					resource.TestCheckResourceAttr("citrixadc_smppparam.tf_unset", "addrton", "0"),
					resource.TestCheckResourceAttr("citrixadc_smppparam.tf_unset", "addrrange", `\d*`),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckSmppparamADCValue("clientmode", "TRANSCEIVER"),
					testAccCheckSmppparamADCValue("msgqueue", "OFF"),
					testAccCheckSmppparamADCValue("msgqueuesize", "10000"),
				),
			},
		},
	})
}

// testAccCheckSmppparamADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it. smppparam is a singleton, so it is fetched with an empty name.
func testAccCheckSmppparamADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Smppparam.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("smppparam not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("smppparam: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccSmppparamDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSmppparamDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_smppparam.tf_smppparam", "clientmode", "TRANSCEIVER"),
					resource.TestCheckResourceAttr("data.citrixadc_smppparam.tf_smppparam", "msgqueue", "OFF"),
				),
			},
		},
	})
}
