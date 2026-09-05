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

const testAccAaaproxyparam_basic = `

	resource "citrixadc_aaaproxyparam" "tf_aaaproxyparam" {
		proxy              = "10.1.1.1:8080"
		proxyauthorization = "basic"
		proxyusername      = "proxyuser"
		proxypassword      = "proxypass123"
	}

`
const testAccAaaproxyparam_update = `

	resource "citrixadc_aaaproxyparam" "tf_aaaproxyparam" {
		proxy              = "http://10.2.2.2:3128"
		proxyauthorization = "basic"
		proxyusername      = "proxyuser2"
		proxypassword      = "proxypass456"
	}

`

func TestAccAaaproxyparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAaaproxyparamDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAaaproxyparam_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaproxyparamExist("citrixadc_aaaproxyparam.tf_aaaproxyparam", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaproxyparam.tf_aaaproxyparam", "proxy", "10.1.1.1:8080"),
					resource.TestCheckResourceAttr("citrixadc_aaaproxyparam.tf_aaaproxyparam", "proxyauthorization", "basic"),
					resource.TestCheckResourceAttr("citrixadc_aaaproxyparam.tf_aaaproxyparam", "proxyusername", "proxyuser"),
					resource.TestCheckResourceAttr("citrixadc_aaaproxyparam.tf_aaaproxyparam", "proxypassword", "proxypass123"),
					// Independent appliance-level confirmation.
					testAccCheckAaaproxyparamADCValue("proxy", "10.1.1.1:8080"),
					testAccCheckAaaproxyparamADCValue("proxyauthorization", "basic"),
					testAccCheckAaaproxyparamADCValue("proxyusername", "proxyuser"),
				),
			},
			{
				Config: testAccAaaproxyparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaproxyparamExist("citrixadc_aaaproxyparam.tf_aaaproxyparam", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaproxyparam.tf_aaaproxyparam", "proxy", "http://10.2.2.2:3128"),
					resource.TestCheckResourceAttr("citrixadc_aaaproxyparam.tf_aaaproxyparam", "proxyauthorization", "basic"),
					resource.TestCheckResourceAttr("citrixadc_aaaproxyparam.tf_aaaproxyparam", "proxyusername", "proxyuser2"),
					resource.TestCheckResourceAttr("citrixadc_aaaproxyparam.tf_aaaproxyparam", "proxypassword", "proxypass456"),
					testAccCheckAaaproxyparamADCValue("proxy", "http://10.2.2.2:3128"),
					testAccCheckAaaproxyparamADCValue("proxyusername", "proxyuser2"),
				),
			},
		},
	})
}

func testAccCheckAaaproxyparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No aaaproxyparam id is set")
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
		data, err := client.FindResource(service.Aaaproxyparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("aaaproxyparam %s not found", n)
		}

		return nil
	}
}

// aaaproxyparam is a global configuration singleton with no NITRO delete
// operation; there is nothing to assert on destroy.
func testAccCheckAaaproxyparamDestroy(s *terraform.State) error {
	return nil
}

// testAccCheckAaaproxyparamADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state). A missing key is treated as an empty
// value, which is how the appliance reports an unset attribute.
func testAccCheckAaaproxyparamADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Aaaproxyparam.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("aaaproxyparam not found on appliance")
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("aaaproxyparam: appliance attr %q = %q, want %q", attr, got, want)
		}
		return nil
	}
}

const testAccAaaproxyparamDataSource_basic = `

	resource "citrixadc_aaaproxyparam" "tf_aaaproxyparam" {
		proxy              = "10.1.1.1:8080"
		proxyauthorization = "basic"
		proxyusername      = "proxyuser"
		proxypassword      = "proxypass123"
	}

	data "citrixadc_aaaproxyparam" "tf_aaaproxyparam" {
		depends_on = [citrixadc_aaaproxyparam.tf_aaaproxyparam]
	}
`

func TestAccAaaproxyparamDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAaaproxyparamDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAaaproxyparamDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_aaaproxyparam.tf_aaaproxyparam", "proxy", "10.1.1.1:8080"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaproxyparam.tf_aaaproxyparam", "proxyauthorization", "basic"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaproxyparam.tf_aaaproxyparam", "proxyusername", "proxyuser"),
				),
			},
		},
	})
}

func TestAccAaaproxyparam_import(t *testing.T) {
	const resAddr = "citrixadc_aaaproxyparam.tf_aaaproxyparam"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAaaproxyparamDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAaaproxyparam_basic},
			{
				Config:            testAccAaaproxyparam_basic,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// proxypassword is a secret retained from config (NITRO returns only
				// an encrypted form); proxypassword_wo_version is a config-only tracker
				// not returned by GET. Neither round-trips on import.
				ImportStateVerifyIgnore: []string{"proxypassword", "proxypassword_wo_version"},
			},
		},
	})
}

// aaaproxyparam is a singleton. Step 1 sets the two unset-eligible attributes
// (proxy and proxyauthorization) to non-default values; step 2 removes them from
// config so the provider must unset them (revert to the NITRO defaults). Note:
// proxyusername/proxypassword do NOT support the NITRO unset operation, so they
// are not exercised here.
const testAccAaaproxyparam_unset_step1 = `

	resource "citrixadc_aaaproxyparam" "tf_aaaproxyparam" {
		proxy              = "10.1.1.1:8080"
		proxyauthorization = "basic"
		proxyusername      = "proxyuser"
		proxypassword      = "proxypass123"
	}
`

const testAccAaaproxyparam_unset_step2 = `

	resource "citrixadc_aaaproxyparam" "tf_aaaproxyparam" {
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccAaaproxyparam_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAaaproxyparamDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccAaaproxyparam_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaproxyparamExist("citrixadc_aaaproxyparam.tf_aaaproxyparam", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaproxyparam.tf_aaaproxyparam", "proxy", "10.1.1.1:8080"),
					resource.TestCheckResourceAttr("citrixadc_aaaproxyparam.tf_aaaproxyparam", "proxyauthorization", "basic"),
					testAccCheckAaaproxyparamADCValue("proxy", "10.1.1.1:8080"),
					testAccCheckAaaproxyparamADCValue("proxyauthorization", "basic"),
				),
			},
			{
				// Removing the attributes must unset them: the appliance reverts
				// them to their defaults (empty), and the implicit post-apply plan
				// must be empty.
				Config: testAccAaaproxyparam_unset_step2,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{
						plancheck.ExpectEmptyPlan(),
					},
				},
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaproxyparamExist("citrixadc_aaaproxyparam.tf_aaaproxyparam", nil),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAaaproxyparamADCValue("proxy", ""),
					testAccCheckAaaproxyparamADCValue("proxyauthorization", ""),
				),
			},
		},
	})
}

// TestAccAaaproxyparam_proxypassword_backward_compat exercises the legacy plain
// (non write-only) proxypassword attribute: it is supplied via a sensitive
// variable and changed between steps, mirroring aaakcdaccount's kcdpassword
// backward-compat test.
const testAccAaaproxyparam_proxypassword_step1 = `
	variable "aaaproxyparam_proxypassword" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_aaaproxyparam" "tf_aaaproxyparam" {
		proxy              = "10.1.1.1:8080"
		proxyauthorization = "basic"
		proxyusername      = "proxyuser"
		proxypassword      = var.aaaproxyparam_proxypassword
	}
`
const testAccAaaproxyparam_proxypassword_step2 = `
	variable "aaaproxyparam_proxypassword_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_aaaproxyparam" "tf_aaaproxyparam" {
		proxy              = "10.1.1.1:8080"
		proxyauthorization = "basic"
		proxyusername      = "proxyuser"
		proxypassword      = var.aaaproxyparam_proxypassword_2
	}
`

func TestAccAaaproxyparam_proxypassword_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_aaaproxyparam_proxypassword", "value1")
	t.Setenv("TF_VAR_aaaproxyparam_proxypassword_2", "value2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAaaproxyparamDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAaaproxyparam_proxypassword_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaproxyparamExist("citrixadc_aaaproxyparam.tf_aaaproxyparam", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaproxyparam.tf_aaaproxyparam", "proxyusername", "proxyuser"),
					resource.TestCheckResourceAttr("citrixadc_aaaproxyparam.tf_aaaproxyparam", "proxyauthorization", "basic"),
				),
			},
			{
				Config: testAccAaaproxyparam_proxypassword_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaproxyparamExist("citrixadc_aaaproxyparam.tf_aaaproxyparam", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaproxyparam.tf_aaaproxyparam", "proxyusername", "proxyuser"),
					resource.TestCheckResourceAttr("citrixadc_aaaproxyparam.tf_aaaproxyparam", "proxyauthorization", "basic"),
				),
			},
		},
	})
}

// TestAccAaaproxyparam_proxypassword_wo_ephemeral exercises the write-only
// (ephemeral) proxypassword_wo secret + proxypassword_wo_version tracker,
// mirroring aaakcdaccount's kcdpassword_wo ephemeral test. The secret is supplied
// via a sensitive variable and rotated by bumping the version; it must never be
// persisted to state.
const testAccAaaproxyparam_proxypassword_wo_step1 = `
	variable "aaaproxyparam_proxypassword_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_aaaproxyparam" "tf_aaaproxyparam" {
		proxy                    = "10.1.1.1:8080"
		proxyauthorization       = "basic"
		proxyusername            = "proxyuser"
		proxypassword_wo         = var.aaaproxyparam_proxypassword_wo
		proxypassword_wo_version = 1
	}
`
const testAccAaaproxyparam_proxypassword_wo_step2 = `
	variable "aaaproxyparam_proxypassword_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_aaaproxyparam" "tf_aaaproxyparam" {
		proxy                    = "10.1.1.1:8080"
		proxyauthorization       = "basic"
		proxyusername            = "proxyuser"
		proxypassword_wo         = var.aaaproxyparam_proxypassword_wo_2
		proxypassword_wo_version = 2
	}
`

func TestAccAaaproxyparam_proxypassword_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_aaaproxyparam_proxypassword_wo", "ephemeral_value1")
	t.Setenv("TF_VAR_aaaproxyparam_proxypassword_wo_2", "ephemeral_value2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAaaproxyparamDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAaaproxyparam_proxypassword_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaproxyparamExist("citrixadc_aaaproxyparam.tf_aaaproxyparam", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaproxyparam.tf_aaaproxyparam", "proxypassword_wo_version", "1"),
					resource.TestCheckResourceAttr("citrixadc_aaaproxyparam.tf_aaaproxyparam", "proxyusername", "proxyuser"),
					// The write-only secret and the plain secret must never be
					// persisted to Terraform state.
					resource.TestCheckNoResourceAttr("citrixadc_aaaproxyparam.tf_aaaproxyparam", "proxypassword_wo"),
					resource.TestCheckNoResourceAttr("citrixadc_aaaproxyparam.tf_aaaproxyparam", "proxypassword"),
				),
			},
			{
				Config: testAccAaaproxyparam_proxypassword_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaproxyparamExist("citrixadc_aaaproxyparam.tf_aaaproxyparam", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaproxyparam.tf_aaaproxyparam", "proxypassword_wo_version", "2"),
					resource.TestCheckResourceAttr("citrixadc_aaaproxyparam.tf_aaaproxyparam", "proxyusername", "proxyuser"),
					resource.TestCheckNoResourceAttr("citrixadc_aaaproxyparam.tf_aaaproxyparam", "proxypassword_wo"),
					resource.TestCheckNoResourceAttr("citrixadc_aaaproxyparam.tf_aaaproxyparam", "proxypassword"),
				),
			},
		},
	})
}
