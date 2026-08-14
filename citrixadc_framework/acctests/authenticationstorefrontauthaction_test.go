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

const testAccAuthenticationstorefrontauthaction_add = `
	resource "citrixadc_authenticationstorefrontauthaction" "tf_storefront" {
		name                       = "tf_storefront"
		serverurl                  = "http://www.example.com/"
		domain                     = "domainname"
		defaultauthenticationgroup = "group_name"
	}
`
const testAccAuthenticationstorefrontauthaction_update = `
	resource "citrixadc_authenticationstorefrontauthaction" "tf_storefront" {
		name                       = "tf_storefront"
		serverurl                  = "http://www.example.com/"
		domain                     = "new_domainname"
		defaultauthenticationgroup = "new_groupname"
	}
`

func TestAccAuthenticationstorefrontauthaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationstorefrontauthactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationstorefrontauthaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationstorefrontauthactionExist("citrixadc_authenticationstorefrontauthaction.tf_storefront", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationstorefrontauthaction.tf_storefront", "name", "tf_storefront"),
					resource.TestCheckResourceAttr("citrixadc_authenticationstorefrontauthaction.tf_storefront", "domain", "domainname"),
				),
			},
			{
				Config: testAccAuthenticationstorefrontauthaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationstorefrontauthactionExist("citrixadc_authenticationstorefrontauthaction.tf_storefront", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationstorefrontauthaction.tf_storefront", "name", "tf_storefront"),
					resource.TestCheckResourceAttr("citrixadc_authenticationstorefrontauthaction.tf_storefront", "domain", "new_domainname"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationstorefrontauthactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationstorefrontauthaction name is set")
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
		data, err := client.FindResource("authenticationstorefrontauthaction", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationstorefrontauthaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationstorefrontauthactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationstorefrontauthaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("authenticationstorefrontauthaction", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationstorefrontauthaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAuthenticationstorefrontauthactionDataSource_basic = `
	resource "citrixadc_authenticationstorefrontauthaction" "tf_storefront_ds" {
		name                       = "tf_storefront_ds"
		serverurl                  = "http://www.example.com/"
		domain                     = "domainname_ds"
		defaultauthenticationgroup = "group_name_ds"
	}

	data "citrixadc_authenticationstorefrontauthaction" "tf_storefront_ds_data" {
		name = citrixadc_authenticationstorefrontauthaction.tf_storefront_ds.name
	}
`

func TestAccAuthenticationstorefrontauthaction_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_authenticationstorefrontauthaction.tf_storefront"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationstorefrontauthactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationstorefrontauthaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationstorefrontauthactionExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Authenticationstorefrontauthaction.Type(), "tf_storefront"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAuthenticationstorefrontauthaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationstorefrontauthactionExist(resAddr, nil)),
			},
		},
	})
}

func TestAccAuthenticationstorefrontauthaction_import(t *testing.T) {
	const resAddr = "citrixadc_authenticationstorefrontauthaction.tf_storefront"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationstorefrontauthactionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAuthenticationstorefrontauthaction_add},
			{
				Config:                  testAccAuthenticationstorefrontauthaction_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccAuthenticationstorefrontauthaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAuthenticationstorefrontauthactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccAuthenticationstorefrontauthaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationstorefrontauthactionExist("citrixadc_authenticationstorefrontauthaction.tf_storefront", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccAuthenticationstorefrontauthaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationstorefrontauthactionExist("citrixadc_authenticationstorefrontauthaction.tf_storefront", nil)),
			},
		},
	})
}

const testAccAuthenticationstorefrontauthaction_unset_step1 = `
	resource "citrixadc_authenticationstorefrontauthaction" "tf_unset" {
		name                       = "tf_test_storefront_unset"
		serverurl                  = "http://www.example.com/"
		domain                     = "unset_domain"
		defaultauthenticationgroup = "unset_group"
	}
`

const testAccAuthenticationstorefrontauthaction_unset_step2 = `
	resource "citrixadc_authenticationstorefrontauthaction" "tf_unset" {
		name      = "tf_test_storefront_unset"
		serverurl = "http://www.example.com/"
		# domain and defaultauthenticationgroup removed from config -> the provider
		# must unset them (revert to NITRO defaults, i.e. absent/empty).
	}
`

func TestAccAuthenticationstorefrontauthaction_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationstorefrontauthactionDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccAuthenticationstorefrontauthaction_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationstorefrontauthactionExist("citrixadc_authenticationstorefrontauthaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationstorefrontauthaction.tf_unset", "domain", "unset_domain"),
					resource.TestCheckResourceAttr("citrixadc_authenticationstorefrontauthaction.tf_unset", "defaultauthenticationgroup", "unset_group"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to the NITRO defaults (empty), and the implicit
				// post-apply plan must be empty.
				Config: testAccAuthenticationstorefrontauthaction_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationstorefrontauthactionExist("citrixadc_authenticationstorefrontauthaction.tf_unset", nil),
					resource.TestCheckNoResourceAttr("citrixadc_authenticationstorefrontauthaction.tf_unset", "domain"),
					resource.TestCheckNoResourceAttr("citrixadc_authenticationstorefrontauthaction.tf_unset", "defaultauthenticationgroup"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAuthenticationstorefrontauthactionADCValue("tf_test_storefront_unset", "domain", ""),
					testAccCheckAuthenticationstorefrontauthactionADCValue("tf_test_storefront_unset", "defaultauthenticationgroup", ""),
				),
			},
		},
	})
}

// testAccCheckAuthenticationstorefrontauthactionADCValue asserts an attribute's
// value directly on the appliance (not just in Terraform state), proving the
// unset actually reverted it. After unset NITRO omits the attribute from GET, so
// the expected value is empty.
func testAccCheckAuthenticationstorefrontauthactionADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Authenticationstorefrontauthaction.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("authenticationstorefrontauthaction %s not found on appliance", name)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("authenticationstorefrontauthaction %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccAuthenticationstorefrontauthactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationstorefrontauthactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationstorefrontauthaction.tf_storefront_ds_data", "name", "tf_storefront_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationstorefrontauthaction.tf_storefront_ds_data", "serverurl", "http://www.example.com/"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationstorefrontauthaction.tf_storefront_ds_data", "domain", "domainname_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationstorefrontauthaction.tf_storefront_ds_data", "defaultauthenticationgroup", "group_name_ds"),
				),
			},
		},
	})
}
