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

const testAccDnspolicy_add = `
resource "citrixadc_dnspolicy" "dnspolicy" {
	name = "policy_A"
	rule = "CLIENT.IP.SRC.IN_SUBNET(1.1.1.1/24)"
	drop = "YES"
	}
`
const testAccDnspolicy_update = `
resource "citrixadc_dnspolicy" "dnspolicy" {
	name = "policy_A"
	rule = "dns.req.question.type.ne(aaaa)"
	drop = "NO"
	}
`

const testAccDnspolicyDataSource_basic = `
resource "citrixadc_dnspolicy" "dnspolicy" {
	name = "policy_A"
	rule = "CLIENT.IP.SRC.IN_SUBNET(1.1.1.1/24)"
	drop = "YES"
}

data "citrixadc_dnspolicy" "dnspolicy_data" {
	name = citrixadc_dnspolicy.dnspolicy.name
}
`

func TestAccDnspolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnspolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnspolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnspolicyExist("citrixadc_dnspolicy.dnspolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_dnspolicy.dnspolicy", "name", "policy_A"),
					resource.TestCheckResourceAttr("citrixadc_dnspolicy.dnspolicy", "rule", "CLIENT.IP.SRC.IN_SUBNET(1.1.1.1/24)"),
					resource.TestCheckResourceAttr("citrixadc_dnspolicy.dnspolicy", "drop", "YES"),
				),
			},
			{
				Config: testAccDnspolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnspolicyExist("citrixadc_dnspolicy.dnspolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_dnspolicy.dnspolicy", "name", "policy_A"),
					resource.TestCheckResourceAttr("citrixadc_dnspolicy.dnspolicy", "rule", "dns.req.question.type.ne(aaaa)"),
					resource.TestCheckResourceAttr("citrixadc_dnspolicy.dnspolicy", "drop", "NO"),
				),
			},
		},
	})
}

func TestAccDnspolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDnspolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_dnspolicy.dnspolicy_data", "name", "policy_A"),
					resource.TestCheckResourceAttr("data.citrixadc_dnspolicy.dnspolicy_data", "rule", "CLIENT.IP.SRC.IN_SUBNET(1.1.1.1/24)"),
				),
			},
		},
	})
}

func TestAccDnspolicy_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_dnspolicy.dnspolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnspolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnspolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnspolicyExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Dnspolicy.Type(), "policy_A"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccDnspolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnspolicyExist(resAddr, nil)),
			},
		},
	})
}

func TestAccDnspolicy_import(t *testing.T) {
	const resAddr = "citrixadc_dnspolicy.dnspolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnspolicyDestroy,
		Steps: []resource.TestStep{
			{Config: testAccDnspolicy_add},
			{
				Config:                  testAccDnspolicy_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"drop"},
			},
		},
	})
}

func TestAccDnspolicy_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDnspolicyDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccDnspolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnspolicyExist("citrixadc_dnspolicy.dnspolicy", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccDnspolicy_add,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckDnspolicyExist("citrixadc_dnspolicy.dnspolicy", nil)),
			},
		},
	})
}

// The only unset-eligible dnspolicy attribute is logaction. All other mutable
// attributes (drop, cachebypass, actionname, viewname, preferredlocation,
// preferredloclist) are rejected by NITRO with "Invalid argument" on unset, so
// they are not wired for unset. logaction requires a messagelog action, so an
// auditmessageaction prerequisite is created; drop=YES supplies the mandatory
// policy action and is kept unchanged across both steps.
const testAccDnspolicy_unset_step1 = `
resource "citrixadc_auditmessageaction" "tf_dnspolicy_ma" {
  name              = "tf_dnspolicy_unset_ma"
  loglevel          = "INFORMATIONAL"
  stringbuilderexpr = "\"dnspolicy unset test\""
}

resource "citrixadc_dnspolicy" "tf_unset" {
  name      = "tf_dnspolicy_unset"
  rule      = "CLIENT.IP.SRC.IN_SUBNET(1.1.1.1/24)"
  drop      = "YES"
  logaction = citrixadc_auditmessageaction.tf_dnspolicy_ma.name
}
`

const testAccDnspolicy_unset_step2 = `
resource "citrixadc_auditmessageaction" "tf_dnspolicy_ma" {
  name              = "tf_dnspolicy_unset_ma"
  loglevel          = "INFORMATIONAL"
  stringbuilderexpr = "\"dnspolicy unset test\""
}

resource "citrixadc_dnspolicy" "tf_unset" {
  name = "tf_dnspolicy_unset"
  rule = "CLIENT.IP.SRC.IN_SUBNET(1.1.1.1/24)"
  drop = "YES"
  # logaction removed from config -> the provider must unset it (revert to the
  # NITRO default, an empty value).
}
`

func TestAccDnspolicy_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnspolicyDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccDnspolicy_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnspolicyExist("citrixadc_dnspolicy.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_dnspolicy.tf_unset", "logaction", "tf_dnspolicy_unset_ma"),
					testAccCheckDnspolicyADCValue("tf_dnspolicy_unset", "logaction", "tf_dnspolicy_unset_ma"),
				),
			},
			{
				// Removing logaction must unset it: state (read back from the
				// appliance) reverts to the NITRO default and the implicit
				// post-apply plan must be empty.
				Config: testAccDnspolicy_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnspolicyExist("citrixadc_dnspolicy.tf_unset", nil),
					// Option B: no Default is injected, so after unset NITRO omits
					// logaction and it reads back as null/absent in state.
					resource.TestCheckNoResourceAttr("citrixadc_dnspolicy.tf_unset", "logaction"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckDnspolicyADCValue("tf_dnspolicy_unset", "logaction", ""),
				),
			},
		},
	})
}

// testAccCheckDnspolicyADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it. An absent attribute is treated as the empty string.
func testAccCheckDnspolicyADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Dnspolicy.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("dnspolicy %s not found on appliance", name)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("dnspolicy %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func testAccCheckDnspolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnspolicy name is set")
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
		data, err := client.FindResource(service.Dnspolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("dnspolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckDnspolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnspolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Dnspolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("dnspolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
