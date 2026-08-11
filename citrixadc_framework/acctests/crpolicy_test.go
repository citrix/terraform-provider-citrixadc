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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccCrpolicy_add = `

resource "citrixadc_crpolicy" "crpolicy" {
    policyname = "crpolicy1"
    rule = "true"
    action = "CACHE"
}

`
const testAccCrpolicy_update = `

resource "citrixadc_crpolicy" "crpolicy" {
    policyname = "crpolicy1"
    rule = "CLIENT.IP.SRC.IN_SUBNET(1.1.1.1/24)"
    action = "ORIGIN"
}

`

func TestAccCrpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCrpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCrpolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrpolicyExist("citrixadc_crpolicy.crpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_crpolicy.crpolicy", "rule", "true"),
					resource.TestCheckResourceAttr("citrixadc_crpolicy.crpolicy", "action", "CACHE"),
				),
			},
			{
				Config: testAccCrpolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrpolicyExist("citrixadc_crpolicy.crpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_crpolicy.crpolicy", "rule", "CLIENT.IP.SRC.IN_SUBNET(1.1.1.1/24)"),
					resource.TestCheckResourceAttr("citrixadc_crpolicy.crpolicy", "action", "ORIGIN"),
				),
			},
		},
	})
}

func TestAccCrpolicy_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_crpolicy.crpolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCrpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCrpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCrpolicyExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Crpolicy.Type(), "crpolicy1"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccCrpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCrpolicyExist(resAddr, nil)),
			},
		},
	})
}

func TestAccCrpolicy_import(t *testing.T) {
	const resAddr = "citrixadc_crpolicy.crpolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCrpolicyDestroy,
		Steps: []resource.TestStep{
			{Config: testAccCrpolicy_add},
			{
				Config:                  testAccCrpolicy_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckCrpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No crpolicy name is set")
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
		data, err := client.FindResource(service.Crpolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("crpolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckCrpolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_crpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Crpolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("crpolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccCrpolicy_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckCrpolicyDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccCrpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCrpolicyExist("citrixadc_crpolicy.crpolicy", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccCrpolicy_add,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckCrpolicyExist("citrixadc_crpolicy.crpolicy", nil)),
			},
		},
	})
}

const testAccCrpolicyDataSource_basic = `

resource "citrixadc_crpolicy" "crpolicy" {
    policyname = "crpolicy_datasource_test"
    rule = "true"
    action = "CACHE"
}

data "citrixadc_crpolicy" "crpolicy_data" {
    policyname = citrixadc_crpolicy.crpolicy.policyname
}
`

// crpolicy's only unset-eligible attribute is logaction: action is not
// unsettable (NITRO rejects ?action=unset with "Invalid argument [action]") and
// rule/policyname are required/key. logaction is Optional+Computed and unsets
// cleanly (revert to the empty NITRO default). Setting it requires a prerequisite
// audit message action.
const testAccCrpolicy_unset_step1 = `
resource "citrixadc_auditmessageaction" "tf_msgaction" {
	name              = "tf_unset_crpolicy_msgaction"
	loglevel          = "INFORMATIONAL"
	stringbuilderexpr = "\"crpolicy unset test\""
}
resource "citrixadc_crpolicy" "tf_unset" {
	policyname = "tf_test_crpolicy_unset"
	rule       = "true"
	action     = "CACHE"
	logaction  = citrixadc_auditmessageaction.tf_msgaction.name
}
`

const testAccCrpolicy_unset_step2 = `
resource "citrixadc_auditmessageaction" "tf_msgaction" {
	name              = "tf_unset_crpolicy_msgaction"
	loglevel          = "INFORMATIONAL"
	stringbuilderexpr = "\"crpolicy unset test\""
}
resource "citrixadc_crpolicy" "tf_unset" {
	policyname = "tf_test_crpolicy_unset"
	rule       = "true"
	action     = "CACHE"
	# logaction removed from config -> provider must unset it.
}
`

func TestAccCrpolicy_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCrpolicyDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccCrpolicy_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrpolicyExist("citrixadc_crpolicy.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_crpolicy.tf_unset", "logaction", "tf_unset_crpolicy_msgaction"),
				),
			},
			{
				// Removing the attribute must unset it: state (read back from the
				// appliance) reverts to the empty NITRO default, the implicit
				// post-apply plan must be empty, and the appliance confirms it.
				Config: testAccCrpolicy_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrpolicyExist("citrixadc_crpolicy.tf_unset", nil),
					// After unset, NITRO omits logaction from GET, so it reads back as
					// null (absent from state) -- assert the revert directly on the
					// appliance instead, and rely on the implicit empty-plan check.
					testAccCheckCrpolicyADCValue("tf_test_crpolicy_unset", "logaction", ""),
				),
			},
		},
	})
}

// testAccCheckCrpolicyADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it. An unset attribute is omitted from the GET response, which reads back as an
// empty string here.
func testAccCheckCrpolicyADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Crpolicy.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("crpolicy %s not found on appliance", name)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("crpolicy %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccCrpolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCrpolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_crpolicy.crpolicy_data", "policyname", "crpolicy_datasource_test"),
					resource.TestCheckResourceAttr("data.citrixadc_crpolicy.crpolicy_data", "rule", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_crpolicy.crpolicy_data", "action", "CACHE"),
				),
			},
		},
	})
}
