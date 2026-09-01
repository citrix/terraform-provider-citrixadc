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

const testAccAppflowpolicy_basic = `
	resource "citrixadc_appflowpolicy" "tf_appflowpolicy" {
		name   = "test_policy"
		action = citrixadc_appflowaction.tf_appflowaction.name
		rule   = "client.TCP.DSTPORT.EQ(22)"
	}
	
	resource "citrixadc_appflowaction" "tf_appflowaction" {
		name            = "test_action"
		collectors      = [citrixadc_appflowcollector.tf_appflowcollector.name, citrixadc_appflowcollector.tf_appflowcollector2.name, ]
		securityinsight = "ENABLED"
		botinsight      = "ENABLED"
		videoanalytics  = "ENABLED"
	}
	resource "citrixadc_appflowcollector" "tf_appflowcollector" {
		name      = "tf_collector"
		ipaddress = "192.168.2.2"
		port      = 80
	}
	resource "citrixadc_appflowcollector" "tf_appflowcollector2" {
		name      = "tf2_collector"
		ipaddress = "192.168.2.3"
		port      = 80
	}
  
`

const testAccAppflowpolicy_update = `
	resource "citrixadc_appflowpolicy" "tf_appflowpolicy" {
		name   = "test_policy"
		action = citrixadc_appflowaction.tf_appflowaction.name
		rule   = "client.TCP.DSTPORT.EQ(25)"
	}
	
	resource "citrixadc_appflowaction" "tf_appflowaction" {
		name            = "test_action"
		collectors      = [citrixadc_appflowcollector.tf_appflowcollector.name, citrixadc_appflowcollector.tf_appflowcollector2.name, ]
		securityinsight = "ENABLED"
		botinsight      = "ENABLED"
		videoanalytics  = "ENABLED"
	}
	resource "citrixadc_appflowcollector" "tf_appflowcollector" {
		name      = "tf_collector"
		ipaddress = "192.168.2.2"
		port      = 80
	}
	resource "citrixadc_appflowcollector" "tf_appflowcollector2" {
		name      = "tf2_collector"
		ipaddress = "192.168.2.3"
		port      = 80
	}
  
`

func TestAccAppflowpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppflowpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppflowpolicy_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowpolicyExist("citrixadc_appflowpolicy.tf_appflowpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowpolicy.tf_appflowpolicy", "name", "test_policy"),
					resource.TestCheckResourceAttr("citrixadc_appflowpolicy.tf_appflowpolicy", "action", "test_action"),
					resource.TestCheckResourceAttr("citrixadc_appflowpolicy.tf_appflowpolicy", "rule", "client.TCP.DSTPORT.EQ(22)"),
				),
			},
			{
				Config: testAccAppflowpolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowpolicyExist("citrixadc_appflowpolicy.tf_appflowpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowpolicy.tf_appflowpolicy", "name", "test_policy"),
					resource.TestCheckResourceAttr("citrixadc_appflowpolicy.tf_appflowpolicy", "action", "test_action"),
					resource.TestCheckResourceAttr("citrixadc_appflowpolicy.tf_appflowpolicy", "rule", "client.TCP.DSTPORT.EQ(25)"),
				),
			},
		},
	})
}

func testAccCheckAppflowpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appflowpolicy name is set")
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
		data, err := client.FindResource(service.Appflowpolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appflowpolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppflowpolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appflowpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Appflowpolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appflowpolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAppflowpolicyDataSource_basic = `
	resource "citrixadc_appflowpolicy" "tf_appflowpolicy" {
		name   = "test_policy"
		action = citrixadc_appflowaction.tf_appflowaction.name
		rule   = "client.TCP.DSTPORT.EQ(22)"
	}
	
	resource "citrixadc_appflowaction" "tf_appflowaction" {
		name            = "test_action"
		collectors      = [citrixadc_appflowcollector.tf_appflowcollector.name, citrixadc_appflowcollector.tf_appflowcollector2.name, ]
		securityinsight = "ENABLED"
		botinsight      = "ENABLED"
		videoanalytics  = "ENABLED"
	}
	resource "citrixadc_appflowcollector" "tf_appflowcollector" {
		name      = "tf_collector"
		ipaddress = "192.168.2.2"
		port      = 80
	}
	resource "citrixadc_appflowcollector" "tf_appflowcollector2" {
		name      = "tf2_collector"
		ipaddress = "192.168.2.3"
		port      = 80
	}

	data "citrixadc_appflowpolicy" "tf_appflowpolicy" {
		name = citrixadc_appflowpolicy.tf_appflowpolicy.name
	}
`

func TestAccAppflowpolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAppflowpolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appflowpolicy.tf_appflowpolicy", "name", "test_policy"),
					resource.TestCheckResourceAttr("data.citrixadc_appflowpolicy.tf_appflowpolicy", "action", "test_action"),
					resource.TestCheckResourceAttr("data.citrixadc_appflowpolicy.tf_appflowpolicy", "rule", "client.TCP.DSTPORT.EQ(22)"),
					// Universal runtime-binding proof for the data source.
					resource.TestCheckResourceAttrSet("data.citrixadc_appflowpolicy.tf_appflowpolicy", "id"),
					// Read-only counter metadata exposed only by the data source.
					resource.TestCheckResourceAttrSet("data.citrixadc_appflowpolicy.tf_appflowpolicy", "hits"),
					resource.TestCheckResourceAttrSet("data.citrixadc_appflowpolicy.tf_appflowpolicy", "undefhits"),
				),
			},
		},
	})
}

func TestAccAppflowpolicy_import(t *testing.T) {
	const resAddr = "citrixadc_appflowpolicy.tf_appflowpolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppflowpolicyDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAppflowpolicy_basic},
			{
				Config:                  testAccAppflowpolicy_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccAppflowpolicy_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAppflowpolicyDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccAppflowpolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppflowpolicyExist("citrixadc_appflowpolicy.tf_appflowpolicy", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccAppflowpolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppflowpolicyExist("citrixadc_appflowpolicy.tf_appflowpolicy", nil)),
			},
		},
	})
}

// The appflowpolicy unset test covers the two spec-unsettable, mutable
// attributes (comment, undefaction). Step 1 sets them to non-default values;
// step 2 removes them from config, so the provider must unset them and the
// appliance reverts them to their defaults (absent -> empty string).
const testAccAppflowpolicy_unset_step1 = `
	resource "citrixadc_appflowcollector" "tf_unset_col" {
		name      = "tf_unset_col"
		ipaddress = "192.168.2.4"
		port      = 80
	}
	resource "citrixadc_appflowaction" "tf_unset_action" {
		name       = "tf_unset_action"
		collectors = [citrixadc_appflowcollector.tf_unset_col.name]
	}
	resource "citrixadc_appflowaction" "tf_unset_undefaction" {
		name       = "tf_unset_undefaction"
		collectors = [citrixadc_appflowcollector.tf_unset_col.name]
	}
	resource "citrixadc_appflowpolicy" "tf_unset" {
		name        = "tf_unset_policy"
		action      = citrixadc_appflowaction.tf_unset_action.name
		rule        = "client.TCP.DSTPORT.EQ(22)"
		comment     = "unset test comment"
		undefaction = citrixadc_appflowaction.tf_unset_undefaction.name
	}
`

const testAccAppflowpolicy_unset_step2 = `
	resource "citrixadc_appflowcollector" "tf_unset_col" {
		name      = "tf_unset_col"
		ipaddress = "192.168.2.4"
		port      = 80
	}
	resource "citrixadc_appflowaction" "tf_unset_action" {
		name       = "tf_unset_action"
		collectors = [citrixadc_appflowcollector.tf_unset_col.name]
	}
	resource "citrixadc_appflowaction" "tf_unset_undefaction" {
		name       = "tf_unset_undefaction"
		collectors = [citrixadc_appflowcollector.tf_unset_col.name]
	}
	resource "citrixadc_appflowpolicy" "tf_unset" {
		name   = "tf_unset_policy"
		action = citrixadc_appflowaction.tf_unset_action.name
		rule   = "client.TCP.DSTPORT.EQ(22)"
		# comment and undefaction removed from config -> provider must unset them.
	}
`

func TestAccAppflowpolicy_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppflowpolicyDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccAppflowpolicy_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowpolicyExist("citrixadc_appflowpolicy.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowpolicy.tf_unset", "comment", "unset test comment"),
					resource.TestCheckResourceAttr("citrixadc_appflowpolicy.tf_unset", "undefaction", "tf_unset_undefaction"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the defaults, and the implicit
				// post-apply plan must be empty.
				Config: testAccAppflowpolicy_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowpolicyExist("citrixadc_appflowpolicy.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowpolicy.tf_unset", "comment", ""),
					resource.TestCheckResourceAttr("citrixadc_appflowpolicy.tf_unset", "undefaction", ""),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAppflowpolicyADCValue("tf_unset_policy", "comment", ""),
					testAccCheckAppflowpolicyADCValue("tf_unset_policy", "undefaction", ""),
				),
			},
		},
	})
}

// testAccCheckAppflowpolicyADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it. An attribute absent from the GET response is treated as "".
func testAccCheckAppflowpolicyADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Appflowpolicy.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("appflowpolicy %s not found on appliance", name)
		}
		got := ""
		if raw, ok := data[attr]; ok && raw != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", raw))
		}
		if got != want {
			return fmt.Errorf("appflowpolicy %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccAppflowpolicy_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_appflowpolicy.tf_appflowpolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppflowpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppflowpolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppflowpolicyExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Appflowpolicy.Type(), "test_policy"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAppflowpolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppflowpolicyExist(resAddr, nil)),
			},
		},
	})
}
