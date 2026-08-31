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
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testAccAuthorizationpolicy_add = `
	resource "citrixadc_authorizationpolicy" "foo" {
		name = "tp-authorize-1"
		rule = "true"
		action = "ALLOW"
	}
`
const testAccAuthorizationpolicy_update = `
	resource "citrixadc_authorizationpolicy" "foo" {
		name = "tp-authorize-1"
		rule = "false"
		action = "ALLOW"
	}
`

const testAccAuthorizationpolicyDataSource_basic = `
	resource "citrixadc_authorizationpolicy" "foo" {
		name = "tp-authorize-datasource-1"
		rule = "true"
		action = "ALLOW"
	}

	data "citrixadc_authorizationpolicy" "foo" {
		name = citrixadc_authorizationpolicy.foo.name
	}
`

func TestAccAuthorizationpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthorizationpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthorizationpolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthorizationpolicyExist("citrixadc_authorizationpolicy.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_authorizationpolicy.foo", "name", "tp-authorize-1"),
					resource.TestCheckResourceAttr("citrixadc_authorizationpolicy.foo", "rule", "true"),
				),
			},
			{
				Config: testAccAuthorizationpolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthorizationpolicyExist("citrixadc_authorizationpolicy.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_authorizationpolicy.foo", "name", "tp-authorize-1"),
					resource.TestCheckResourceAttr("citrixadc_authorizationpolicy.foo", "rule", "false"),
				),
			},
		},
	})
}

func TestAccAuthorizationpolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthorizationpolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authorizationpolicy.foo", "name", "tp-authorize-datasource-1"),
					resource.TestCheckResourceAttr("data.citrixadc_authorizationpolicy.foo", "rule", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_authorizationpolicy.foo", "action", "ALLOW"),
				),
			},
		},
	})
}

func testAccCheckAuthorizationpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authorizationpolicy name is set")
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
		data, err := client.FindResource(service.Authorizationpolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authorizationpolicy %s not found", n)
		}

		return nil
	}
}

func TestAccAuthorizationpolicy_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_authorizationpolicy.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthorizationpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthorizationpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthorizationpolicyExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Authorizationpolicy.Type(), "tp-authorize-1"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAuthorizationpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthorizationpolicyExist(resAddr, nil)),
			},
		},
	})
}

func TestAccAuthorizationpolicy_import(t *testing.T) {
	const resAddr = "citrixadc_authorizationpolicy.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthorizationpolicyDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAuthorizationpolicy_add},
			{
				Config:                  testAccAuthorizationpolicy_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccAuthorizationpolicy_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAuthorizationpolicyDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccAuthorizationpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthorizationpolicyExist("citrixadc_authorizationpolicy.foo", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccAuthorizationpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthorizationpolicyExist("citrixadc_authorizationpolicy.foo", nil)),
			},
		},
	})
}

func testAccCheckAuthorizationpolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authorizationpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Authorizationpolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authorizationpolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
