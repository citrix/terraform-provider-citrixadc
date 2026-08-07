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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccAuthenticationnoauthaction_add = `
	resource "citrixadc_authenticationnoauthaction" "tf_noauthaction" {
		name                       = "tf_noauthaction"
		defaultauthenticationgroup = "old_group"
	}
`
const testAccAuthenticationnoauthaction_update = `
	resource "citrixadc_authenticationnoauthaction" "tf_noauthaction" {
		name                       = "tf_noauthaction"
		defaultauthenticationgroup = "new_group"
	}
`

const testAccAuthenticationnoauthactionDataSource_basic = `
	resource "citrixadc_authenticationnoauthaction" "tf_noauthaction" {
		name                       = "tf_noauthaction_ds"
		defaultauthenticationgroup = "test_group"
	}

	data "citrixadc_authenticationnoauthaction" "tf_noauthaction_ds" {
		name = citrixadc_authenticationnoauthaction.tf_noauthaction.name
	}
`

func TestAccAuthenticationnoauthaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationnoauthactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationnoauthaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationnoauthactionExist("citrixadc_authenticationnoauthaction.tf_noauthaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationnoauthaction.tf_noauthaction", "name", "tf_noauthaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationnoauthaction.tf_noauthaction", "defaultauthenticationgroup", "old_group"),
				),
			},
			{
				Config: testAccAuthenticationnoauthaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationnoauthactionExist("citrixadc_authenticationnoauthaction.tf_noauthaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationnoauthaction.tf_noauthaction", "name", "tf_noauthaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationnoauthaction.tf_noauthaction", "defaultauthenticationgroup", "new_group"),
				),
			},
		},
	})
}

func testAccCheckAuthenticationnoauthactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationnoauthaction name is set")
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
		data, err := client.FindResource("authenticationnoauthaction", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationnoauthaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationnoauthactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationnoauthaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("authenticationnoauthaction", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationnoauthaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccAuthenticationnoauthaction_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_authenticationnoauthaction.tf_noauthaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationnoauthactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationnoauthaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationnoauthactionExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Authenticationnoauthaction.Type(), "tf_noauthaction"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAuthenticationnoauthaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationnoauthactionExist(resAddr, nil)),
			},
		},
	})
}

func TestAccAuthenticationnoauthaction_import(t *testing.T) {
	const resAddr = "citrixadc_authenticationnoauthaction.tf_noauthaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationnoauthactionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAuthenticationnoauthaction_add},
			{
				Config:                  testAccAuthenticationnoauthaction_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccAuthenticationnoauthaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAuthenticationnoauthactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccAuthenticationnoauthaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationnoauthactionExist("citrixadc_authenticationnoauthaction.tf_noauthaction", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAuthenticationnoauthaction_add,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckAuthenticationnoauthactionExist("citrixadc_authenticationnoauthaction.tf_noauthaction", nil)),
			},
		},
	})
}

func TestAccAuthenticationnoauthactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationnoauthactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.citrixadc_authenticationnoauthaction.tf_noauthaction_ds", "name", "citrixadc_authenticationnoauthaction.tf_noauthaction", "name"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationnoauthaction.tf_noauthaction_ds", "name", "tf_noauthaction_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationnoauthaction.tf_noauthaction_ds", "defaultauthenticationgroup", "test_group"),
				),
			},
		},
	})
}
