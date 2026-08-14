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

const testAccIpsecparameter_basic = `

resource "citrixadc_ipsecparameter" "tf_ipsecparameter" {
	ikeversion            = "V2"
	encalgo               = ["AES", "AES256"]
	hashalgo              = ["HMAC_SHA1", "HMAC_SHA256"]
	livenesscheckinterval = 50
	}
  
`
const testAccIpsecparameter_update = `

resource "citrixadc_ipsecparameter" "tf_ipsecparameter" {
	ikeversion            = "V1"
	encalgo               = ["AES", "AES256"]
	hashalgo              = ["HMAC_SHA1", "HMAC_SHA256"]
	livenesscheckinterval = 60
	}
  
`

const testAccIpsecparameterDataSource_basic = `

resource "citrixadc_ipsecparameter" "tf_ipsecparameter" {
	ikeversion            = "V2"
	encalgo               = ["AES", "AES256"]
	hashalgo              = ["HMAC_SHA1", "HMAC_SHA256"]
	livenesscheckinterval = 50
}

data "citrixadc_ipsecparameter" "tf_ipsecparameter_datasource" {
	depends_on = [citrixadc_ipsecparameter.tf_ipsecparameter]
}
`

func TestAccIpsecparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccIpsecparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpsecparameterExist("citrixadc_ipsecparameter.tf_ipsecparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_ipsecparameter.tf_ipsecparameter", "ikeversion", "V2"),
					resource.TestCheckResourceAttr("citrixadc_ipsecparameter.tf_ipsecparameter", "livenesscheckinterval", "50"),
				),
			},
			{
				Config: testAccIpsecparameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpsecparameterExist("citrixadc_ipsecparameter.tf_ipsecparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_ipsecparameter.tf_ipsecparameter", "ikeversion", "V1"),
					resource.TestCheckResourceAttr("citrixadc_ipsecparameter.tf_ipsecparameter", "livenesscheckinterval", "60"),
				),
			},
		},
	})
}

func testAccCheckIpsecparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ipsecparameter name is set")
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
		data, err := client.FindResource(service.Ipsecparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("ipsecparameter %s not found", n)
		}

		return nil
	}
}

func TestAccIpsecparameter_import(t *testing.T) {
	const resAddr = "citrixadc_ipsecparameter.tf_ipsecparameter"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccIpsecparameter_basic},
			{
				Config:                  testAccIpsecparameter_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccIpsecparameter_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccIpsecparameter_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIpsecparameterExist("citrixadc_ipsecparameter.tf_ipsecparameter", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccIpsecparameter_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIpsecparameterExist("citrixadc_ipsecparameter.tf_ipsecparameter", nil)),
			},
		},
	})
}

// ipsecparameter is a singleton (global) resource. Step1 sets every
// unset-eligible attribute to a valid non-default value; step2 removes them all
// so the provider must unset them (revert to the documented NITRO defaults).
const testAccIpsecparameter_unset_step1 = `
resource "citrixadc_ipsecparameter" "tf_unset" {
	ikeversion            = "V1"
	encalgo               = ["AES256"]
	hashalgo              = ["HMAC_SHA1"]
	lifetime              = 480
	livenesscheckinterval = 50
	replaywindowsize      = 8192
	ikeretryinterval      = 120
	perfectforwardsecrecy = "ENABLE"
	retransmissiontime    = 5
}
`

const testAccIpsecparameter_unset_step2 = `
resource "citrixadc_ipsecparameter" "tf_unset" {
	# All unset-eligible scalar attributes removed from config -> the provider
	# must unset them (revert to the NITRO defaults). encalgo/hashalgo are list
	# attributes without a schema Default (sticky on removal) so they are kept in
	# config and are not part of the unset set.
	encalgo  = ["AES256"]
	hashalgo = ["HMAC_SHA1"]
}
`

func TestAccIpsecparameter_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccIpsecparameter_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpsecparameterExist("citrixadc_ipsecparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_ipsecparameter.tf_unset", "ikeversion", "V1"),
					resource.TestCheckResourceAttr("citrixadc_ipsecparameter.tf_unset", "encalgo.0", "AES256"),
					resource.TestCheckResourceAttr("citrixadc_ipsecparameter.tf_unset", "hashalgo.0", "HMAC_SHA1"),
					resource.TestCheckResourceAttr("citrixadc_ipsecparameter.tf_unset", "lifetime", "480"),
					resource.TestCheckResourceAttr("citrixadc_ipsecparameter.tf_unset", "livenesscheckinterval", "50"),
					resource.TestCheckResourceAttr("citrixadc_ipsecparameter.tf_unset", "replaywindowsize", "8192"),
					resource.TestCheckResourceAttr("citrixadc_ipsecparameter.tf_unset", "ikeretryinterval", "120"),
					resource.TestCheckResourceAttr("citrixadc_ipsecparameter.tf_unset", "perfectforwardsecrecy", "ENABLE"),
					resource.TestCheckResourceAttr("citrixadc_ipsecparameter.tf_unset", "retransmissiontime", "5"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccIpsecparameter_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpsecparameterExist("citrixadc_ipsecparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_ipsecparameter.tf_unset", "ikeversion", "V2"),
					resource.TestCheckResourceAttr("citrixadc_ipsecparameter.tf_unset", "lifetime", "28800"),
					resource.TestCheckResourceAttr("citrixadc_ipsecparameter.tf_unset", "livenesscheckinterval", "10"),
					resource.TestCheckResourceAttr("citrixadc_ipsecparameter.tf_unset", "replaywindowsize", "9216"),
					resource.TestCheckResourceAttr("citrixadc_ipsecparameter.tf_unset", "ikeretryinterval", "60"),
					resource.TestCheckResourceAttr("citrixadc_ipsecparameter.tf_unset", "perfectforwardsecrecy", "DISABLE"),
					resource.TestCheckResourceAttr("citrixadc_ipsecparameter.tf_unset", "retransmissiontime", "1"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckIpsecparameterADCValue("ikeversion", "V2"),
					testAccCheckIpsecparameterADCValue("perfectforwardsecrecy", "DISABLE"),
					testAccCheckIpsecparameterADCValue("lifetime", "28800"),
				),
			},
		},
	})
}

// testAccCheckIpsecparameterADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckIpsecparameterADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Ipsecparameter.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("ipsecparameter not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("ipsecparameter: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccIpsecparameterDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIpsecparameterDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_ipsecparameter.tf_ipsecparameter_datasource", "ikeversion", "V2"),
					resource.TestCheckResourceAttr("data.citrixadc_ipsecparameter.tf_ipsecparameter_datasource", "livenesscheckinterval", "50"),
				),
			},
		},
	})
}
