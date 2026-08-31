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

const testAccNspbr6_add = `

	resource "citrixadc_iptunnel" "tf_iptunnel" {
		name             = "tf_iptunnel"
		remote           = "66.0.0.11"
		remotesubnetmask = "255.255.255.255"
		local            = "*"
	}
	resource "citrixadc_nspbr6" "tf_nspbr6" {
		name     = "tf_nspbr6"
		action   = "ALLOW"
		protocol = "ICMPV6"
		priority = 20
		state    = "ENABLED"
		iptunnel = citrixadc_iptunnel.tf_iptunnel.name
	}
`
const testAccNspbr6_update = `

	resource "citrixadc_iptunnel" "tf_iptunnel" {
		name             = "tf_iptunnel"
		remote           = "66.0.0.11"
		remotesubnetmask = "255.255.255.255"
		local            = "*"
	}
	resource "citrixadc_nspbr6" "tf_nspbr6" {
		name     = "tf_nspbr6"
		action   = "ALLOW"
		protocol = "TCP"
		priority = 30
		state    = "DISABLED"
		iptunnel = citrixadc_iptunnel.tf_iptunnel.name
	}
`

func TestAccNspbr6_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNspbr6Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNspbr6_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNspbr6Exist("citrixadc_nspbr6.tf_nspbr6", nil),
					resource.TestCheckResourceAttr("citrixadc_nspbr6.tf_nspbr6", "name", "tf_nspbr6"),
					resource.TestCheckResourceAttr("citrixadc_nspbr6.tf_nspbr6", "protocol", "ICMPV6"),
					resource.TestCheckResourceAttr("citrixadc_nspbr6.tf_nspbr6", "priority", "20"),
					resource.TestCheckResourceAttr("citrixadc_nspbr6.tf_nspbr6", "state", "ENABLED"),
				),
			},
			{
				Config: testAccNspbr6_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNspbr6Exist("citrixadc_nspbr6.tf_nspbr6", nil),
					resource.TestCheckResourceAttr("citrixadc_nspbr6.tf_nspbr6", "name", "tf_nspbr6"),
					resource.TestCheckResourceAttr("citrixadc_nspbr6.tf_nspbr6", "protocol", "TCP"),
					resource.TestCheckResourceAttr("citrixadc_nspbr6.tf_nspbr6", "priority", "30"),
					resource.TestCheckResourceAttr("citrixadc_nspbr6.tf_nspbr6", "state", "DISABLED"),
				),
			},
		},
	})
}

func TestAccNspbr6_import(t *testing.T) {
	const resAddr = "citrixadc_nspbr6.tf_nspbr6"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNspbr6Destroy,
		Steps: []resource.TestStep{
			{Config: testAccNspbr6_add},
			{
				Config:                  testAccNspbr6_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckNspbr6Exist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nspbr6 name is set")
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
		data, err := client.FindResource(service.Nspbr6.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nspbr6 %s not found", n)
		}

		return nil
	}
}

func testAccCheckNspbr6Destroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nspbr6" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Nspbr6.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nspbr6 %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccNspbr6_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_nspbr6.tf_nspbr6"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNspbr6Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNspbr6_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNspbr6Exist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Nspbr6.Type(), "tf_nspbr6"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccNspbr6_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNspbr6Exist(resAddr, nil)),
			},
		},
	})
}

func TestAccNspbr6_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNspbr6Destroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccNspbr6_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNspbr6Exist("citrixadc_nspbr6.tf_nspbr6", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccNspbr6_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNspbr6Exist("citrixadc_nspbr6.tf_nspbr6", nil)),
			},
		},
	})
}

const testAccNspbr6DataSource_basic = `

	resource "citrixadc_iptunnel" "tf_iptunnel" {
		name             = "tf_iptunnel"
		remote           = "66.0.0.11"
		remotesubnetmask = "255.255.255.255"
		local            = "*"
	}
	resource "citrixadc_nspbr6" "tf_nspbr6" {
		name     = "tf_nspbr6_ds"
		action   = "ALLOW"
		protocol = "ICMPV6"
		priority = 25
		state    = "ENABLED"
		iptunnel = citrixadc_iptunnel.tf_iptunnel.name
	}

	data "citrixadc_nspbr6" "tf_nspbr6" {
		name   = citrixadc_nspbr6.tf_nspbr6.name
		detail = false
	}
`

// The nspbr6 unset test covers the spec-unsettable mutable attribute with a
// documented NITRO server default that applies cleanly here: msr (default
// "DISABLED"). Step 1 sets a non-default value; step 2 removes it from config so
// the provider must unset it, reverting to the NITRO default.
const testAccNspbr6_unset_step1 = `
	resource "citrixadc_iptunnel" "tf_iptunnel" {
		name             = "tf_iptunnel_unset"
		remote           = "66.0.0.11"
		remotesubnetmask = "255.255.255.255"
		local            = "*"
	}
	resource "citrixadc_nspbr6" "tf_unset" {
		name     = "tf_nspbr6_unset"
		action   = "ALLOW"
		iptunnel = citrixadc_iptunnel.tf_iptunnel.name
		msr      = "ENABLED"
	}
`

const testAccNspbr6_unset_step2 = `
	resource "citrixadc_iptunnel" "tf_iptunnel" {
		name             = "tf_iptunnel_unset"
		remote           = "66.0.0.11"
		remotesubnetmask = "255.255.255.255"
		local            = "*"
	}
	resource "citrixadc_nspbr6" "tf_unset" {
		name     = "tf_nspbr6_unset"
		action   = "ALLOW"
		iptunnel = citrixadc_iptunnel.tf_iptunnel.name
		# msr removed from config -> provider must unset it (revert to NITRO
		# default "DISABLED").
	}
`

func TestAccNspbr6_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNspbr6Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNspbr6_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNspbr6Exist("citrixadc_nspbr6.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nspbr6.tf_unset", "msr", "ENABLED"),
				),
			},
			{
				Config: testAccNspbr6_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNspbr6Exist("citrixadc_nspbr6.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nspbr6.tf_unset", "msr", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckNspbr6ADCValue("tf_nspbr6_unset", "msr", "DISABLED"),
				),
			},
		},
	})
}

// testAccCheckNspbr6ADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckNspbr6ADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Nspbr6.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("nspbr6 %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("nspbr6 %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccNspbr6DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNspbr6DataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nspbr6.tf_nspbr6", "name", "tf_nspbr6_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_nspbr6.tf_nspbr6", "action", "ALLOW"),
					resource.TestCheckResourceAttr("data.citrixadc_nspbr6.tf_nspbr6", "protocol", "ICMPV6"),
					resource.TestCheckResourceAttr("data.citrixadc_nspbr6.tf_nspbr6", "priority", "25"),
				),
			},
		},
	})
}
