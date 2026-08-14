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

const testAccIcapolicy_basic = `

	resource "citrixadc_icaaction" "tf_icaaction" {
		name              = "my_ica_action"
		accessprofilename = "default_ica_accessprofile"
	}
	resource "citrixadc_icapolicy" "tf_icapolicy" {
		name   = "my_ica_policy"
		rule   = true
		action = citrixadc_icaaction.tf_icaaction.name
	}
`

const testAccIcapolicy_update = `

	resource "citrixadc_icaaction" "tf_icaaction" {
		name              = "my_ica_action"
		accessprofilename = "default_ica_accessprofile"
	}

	resource "citrixadc_icapolicy" "tf_icapolicy" {
		name   = "my_ica_policy"
		rule   = false
		action = citrixadc_icaaction.tf_icaaction.name
	}
`

func TestAccIcapolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIcapolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIcapolicy_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcapolicyExist("citrixadc_icapolicy.tf_icapolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_icapolicy.tf_icapolicy", "name", "my_ica_policy"),
					resource.TestCheckResourceAttr("citrixadc_icapolicy.tf_icapolicy", "rule", "true"),
					resource.TestCheckResourceAttr("citrixadc_icapolicy.tf_icapolicy", "action", "my_ica_action"),
				),
			},
			{
				Config: testAccIcapolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcapolicyExist("citrixadc_icapolicy.tf_icapolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_icapolicy.tf_icapolicy", "name", "my_ica_policy"),
					resource.TestCheckResourceAttr("citrixadc_icapolicy.tf_icapolicy", "rule", "false"),
					resource.TestCheckResourceAttr("citrixadc_icapolicy.tf_icapolicy", "action", "my_ica_action"),
				),
			},
		},
	})
}

func testAccCheckIcapolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No icapolicy name is set")
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
		data, err := client.FindResource("icapolicy", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("icapolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckIcapolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_icapolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("icapolicy", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("icapolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccIcapolicy_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_icapolicy.tf_icapolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIcapolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIcapolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIcapolicyExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Icapolicy.Type(), "my_ica_policy"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccIcapolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIcapolicyExist(resAddr, nil)),
			},
		},
	})
}

func TestAccIcapolicy_import(t *testing.T) {
	const resAddr = "citrixadc_icapolicy.tf_icapolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIcapolicyDestroy,
		Steps: []resource.TestStep{
			{Config: testAccIcapolicy_basic},
			{
				Config:                  testAccIcapolicy_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccIcapolicy_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckIcapolicyDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccIcapolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIcapolicyExist("citrixadc_icapolicy.tf_icapolicy", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccIcapolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIcapolicyExist("citrixadc_icapolicy.tf_icapolicy", nil)),
			},
		},
	})
}

// Unset test: step1 sets the unset-eligible mutable attributes (comment,
// logaction) to valid non-default values; step2 removes them from config so the
// provider issues a NITRO unset, reverting them to their defaults (absent -> "").
const testAccIcapolicy_unset_step1 = `

	resource "citrixadc_icaaction" "tf_icaaction_unset" {
		name              = "tf_ica_action_unset"
		accessprofilename = "default_ica_accessprofile"
	}

	resource "citrixadc_auditmessageaction" "tf_msgaction_unset" {
		name              = "tf_ica_msgaction_unset"
		loglevel          = "NOTICE"
		stringbuilderexpr = "\"ica unset test\""
	}

	resource "citrixadc_icapolicy" "tf_icapolicy_unset" {
		name      = "tf_ica_policy_unset"
		rule      = true
		action    = citrixadc_icaaction.tf_icaaction_unset.name
		comment   = "managed by terraform"
		logaction = citrixadc_auditmessageaction.tf_msgaction_unset.name
	}
`

const testAccIcapolicy_unset_step2 = `

	resource "citrixadc_icaaction" "tf_icaaction_unset" {
		name              = "tf_ica_action_unset"
		accessprofilename = "default_ica_accessprofile"
	}

	resource "citrixadc_auditmessageaction" "tf_msgaction_unset" {
		name              = "tf_ica_msgaction_unset"
		loglevel          = "NOTICE"
		stringbuilderexpr = "\"ica unset test\""
	}

	resource "citrixadc_icapolicy" "tf_icapolicy_unset" {
		name   = "tf_ica_policy_unset"
		rule   = true
		action = citrixadc_icaaction.tf_icaaction_unset.name
		# comment and logaction removed from config -> the provider must unset them
		# (revert to NITRO defaults, i.e. absent / empty string).
	}
`

func TestAccIcapolicy_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIcapolicyDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccIcapolicy_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcapolicyExist("citrixadc_icapolicy.tf_icapolicy_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_icapolicy.tf_icapolicy_unset", "comment", "managed by terraform"),
					resource.TestCheckResourceAttr("citrixadc_icapolicy.tf_icapolicy_unset", "logaction", "tf_ica_msgaction_unset"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to the defaults, and the implicit post-apply plan
				// must be empty.
				Config: testAccIcapolicy_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcapolicyExist("citrixadc_icapolicy.tf_icapolicy_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_icapolicy.tf_icapolicy_unset", "comment", ""),
					resource.TestCheckResourceAttr("citrixadc_icapolicy.tf_icapolicy_unset", "logaction", ""),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckIcapolicyADCValue("tf_ica_policy_unset", "comment", ""),
					testAccCheckIcapolicyADCValue("tf_ica_policy_unset", "logaction", ""),
				),
			},
		},
	})
}

// testAccCheckIcapolicyADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it. An absent attribute is treated as the empty string.
func testAccCheckIcapolicyADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Icapolicy.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("icapolicy %s not found on appliance", name)
		}
		got := ""
		if raw, ok := data[attr]; ok && raw != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", raw))
		}
		if got != want {
			return fmt.Errorf("icapolicy %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccIcapolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIcapolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_icapolicy.tf_icapolicy_ds", "name", "my_ica_policy_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_icapolicy.tf_icapolicy_ds", "rule", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_icapolicy.tf_icapolicy_ds", "action", "my_ica_action_ds"),
					resource.TestCheckResourceAttrSet("data.citrixadc_icapolicy.tf_icapolicy_ds", "id"),
				),
			},
		},
	})
}

const testAccIcapolicyDataSource_basic = `

	resource "citrixadc_icaaction" "tf_icaaction_ds" {
		name              = "my_ica_action_ds"
		accessprofilename = "default_ica_accessprofile"
	}

	resource "citrixadc_icapolicy" "tf_icapolicy_ds" {
		name   = "my_ica_policy_ds"
		rule   = true
		action = citrixadc_icaaction.tf_icaaction_ds.name
	}

	data "citrixadc_icapolicy" "tf_icapolicy_ds" {
		name = citrixadc_icapolicy.tf_icapolicy_ds.name
	}
`
