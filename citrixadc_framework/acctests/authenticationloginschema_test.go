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

const testAccAuthenticationloginschema_add = `
	resource "citrixadc_authenticationloginschema" "tf_loginschema" {
		name                    = "tf_loginschema"
		authenticationschema    = "LoginSchema/SingleAuth.xml"
		ssocredentials          = "NO"
		authenticationstrength  = "30"
		passwordcredentialindex = "10"
	}
`
const testAccAuthenticationloginschema_update = `
	resource "citrixadc_authenticationloginschema" "tf_loginschema" {
		name                    = "tf_loginschema"
		authenticationschema    = "LoginSchema/SingleAuth.xml"
		ssocredentials          = "YES"
		authenticationstrength  = "20"
		passwordcredentialindex = "10"
	}
`

func TestAccAuthenticationloginschema_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationloginschemaDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationloginschema_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationloginschemaExist("citrixadc_authenticationloginschema.tf_loginschema", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationloginschema.tf_loginschema", "name", "tf_loginschema"),
					resource.TestCheckResourceAttr("citrixadc_authenticationloginschema.tf_loginschema", "ssocredentials", "NO"),
					resource.TestCheckResourceAttr("citrixadc_authenticationloginschema.tf_loginschema", "authenticationstrength", "30"),
				),
			},
			{
				Config: testAccAuthenticationloginschema_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationloginschemaExist("citrixadc_authenticationloginschema.tf_loginschema", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationloginschema.tf_loginschema", "name", "tf_loginschema"),
					resource.TestCheckResourceAttr("citrixadc_authenticationloginschema.tf_loginschema", "ssocredentials", "YES"),
					resource.TestCheckResourceAttr("citrixadc_authenticationloginschema.tf_loginschema", "authenticationstrength", "20"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationloginschemaExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationloginschema name is set")
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
		data, err := client.FindResource(service.Authenticationloginschema.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationloginschema %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationloginschemaDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationloginschema" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Authenticationloginschema.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationloginschema %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccAuthenticationloginschema_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_authenticationloginschema.tf_loginschema"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationloginschemaDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationloginschema_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationloginschemaExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Authenticationloginschema.Type(), "tf_loginschema"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAuthenticationloginschema_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationloginschemaExist(resAddr, nil)),
			},
		},
	})
}

// The authenticationloginschema unset test covers ssocredentials, the only
// mutable attribute with a documented NITRO default ("NO"). Step 1 sets it to a
// non-default value ("YES"); step 2 removes it from config, so the provider must
// unset it and the appliance reverts it to the default.
const testAccAuthenticationloginschema_unset_step1 = `
	resource "citrixadc_authenticationloginschema" "tf_unset" {
		name                 = "tf_loginschema_unset"
		authenticationschema = "LoginSchema/SingleAuth.xml"
		ssocredentials       = "YES"
	}
`

const testAccAuthenticationloginschema_unset_step2 = `
	resource "citrixadc_authenticationloginschema" "tf_unset" {
		name                 = "tf_loginschema_unset"
		authenticationschema = "LoginSchema/SingleAuth.xml"
		# ssocredentials removed from config -> provider must unset it (revert to "NO").
	}
`

func TestAccAuthenticationloginschema_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationloginschemaDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationloginschema_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationloginschemaExist("citrixadc_authenticationloginschema.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationloginschema.tf_unset", "ssocredentials", "YES"),
				),
			},
			{
				Config: testAccAuthenticationloginschema_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationloginschemaExist("citrixadc_authenticationloginschema.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationloginschema.tf_unset", "ssocredentials", "NO"),
					testAccCheckAuthenticationloginschemaADCValue("tf_loginschema_unset", "ssocredentials", "NO"),
				),
			},
		},
	})
}

// testAccCheckAuthenticationloginschemaADCValue asserts an attribute's value
// directly on the appliance, proving the unset actually reverted it.
func testAccCheckAuthenticationloginschemaADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Authenticationloginschema.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("authenticationloginschema %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("authenticationloginschema %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccAuthenticationloginschema_import(t *testing.T) {
	const resAddr = "citrixadc_authenticationloginschema.tf_loginschema"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationloginschemaDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAuthenticationloginschema_add},
			{
				Config:                  testAccAuthenticationloginschema_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccAuthenticationloginschema_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAuthenticationloginschemaDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccAuthenticationloginschema_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationloginschemaExist("citrixadc_authenticationloginschema.tf_loginschema", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccAuthenticationloginschema_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationloginschemaExist("citrixadc_authenticationloginschema.tf_loginschema", nil)),
			},
		},
	})
}

const testAccAuthenticationloginschemaDataSource_basic = `
	resource "citrixadc_authenticationloginschema" "tf_loginschema_ds" {
		name                    = "tf_loginschema_ds"
		authenticationschema    = "LoginSchema/SingleAuth.xml"
		ssocredentials          = "NO"
		authenticationstrength  = "30"
		passwordcredentialindex = "10"
	}
	data "citrixadc_authenticationloginschema" "tf_loginschema_ds" {
		name = citrixadc_authenticationloginschema.tf_loginschema_ds.name
	}
`

func TestAccAuthenticationloginschemaDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationloginschemaDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationloginschema.tf_loginschema_ds", "name", "tf_loginschema_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationloginschema.tf_loginschema_ds", "authenticationschema", "LoginSchema/SingleAuth.xml"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationloginschema.tf_loginschema_ds", "ssocredentials", "NO"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationloginschema.tf_loginschema_ds", "authenticationstrength", "30"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationloginschema.tf_loginschema_ds", "passwordcredentialindex", "10"),
					resource.TestCheckResourceAttrSet("data.citrixadc_authenticationloginschema.tf_loginschema_ds", "id"),
				),
			},
		},
	})
}
