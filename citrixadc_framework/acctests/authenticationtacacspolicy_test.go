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

const testAccAuthenticationtacacspolicy_add = `
	resource "citrixadc_authenticationtacacsaction" "tf_tacacsaction" {
		name            = "tf_tacacsaction"
		serverip        = "1.2.3.4"
		serverport      = 8080
		authtimeout     = 5
		authorization   = "ON"
		accounting      = "ON"
		auditfailedcmds = "ON"
		groupattrname   = "group"
	}
	resource "citrixadc_authenticationtacacspolicy" "tf_tacacspolicy" {
		name= "tf_tacacspolicy"
		rule= "NS_TRUE"
		reqaction= citrixadc_authenticationtacacsaction.tf_tacacsaction.name
		
	}
`
const testAccAuthenticationtacacspolicy_update = `
	resource "citrixadc_authenticationtacacsaction" "tf_tacacsaction" {
		name            = "tf_tacacsaction"
		serverip        = "1.2.3.4"
		serverport      = 8080
		authtimeout     = 5
		authorization   = "ON"
		accounting      = "ON"
		auditfailedcmds = "ON"
		groupattrname   = "group"
	}
	resource "citrixadc_authenticationtacacspolicy" "tf_tacacspolicy" {
		name= "tf_tacacspolicy"
		rule= "NS_FALSE"
		reqaction= citrixadc_authenticationtacacsaction.tf_tacacsaction.name
		
	}
`

func TestAccAuthenticationtacacspolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationtacacspolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationtacacspolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationtacacspolicyExist("citrixadc_authenticationtacacspolicy.tf_tacacspolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationtacacspolicy.tf_tacacspolicy", "name", "tf_tacacspolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationtacacspolicy.tf_tacacspolicy", "rule", "NS_TRUE"),
				),
			},
			{
				Config: testAccAuthenticationtacacspolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationtacacspolicyExist("citrixadc_authenticationtacacspolicy.tf_tacacspolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationtacacspolicy.tf_tacacspolicy", "name", "tf_tacacspolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationtacacspolicy.tf_tacacspolicy", "rule", "NS_FALSE"),
				),
			},
		},
	})
}

func TestAccAuthenticationtacacspolicy_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAuthenticationtacacspolicyDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccAuthenticationtacacspolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationtacacspolicyExist("citrixadc_authenticationtacacspolicy.tf_tacacspolicy", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccAuthenticationtacacspolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationtacacspolicyExist("citrixadc_authenticationtacacspolicy.tf_tacacspolicy", nil),
				),
			},
		},
	})
}

func TestAccAuthenticationtacacspolicy_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_authenticationtacacspolicy.tf_tacacspolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationtacacspolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationtacacspolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationtacacspolicyExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Authenticationtacacspolicy.Type(), "tf_tacacspolicy"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAuthenticationtacacspolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationtacacspolicyExist(resAddr, nil)),
			},
		},
	})
}

func TestAccAuthenticationtacacspolicy_import(t *testing.T) {
	const resAddr = "citrixadc_authenticationtacacspolicy.tf_tacacspolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationtacacspolicyDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAuthenticationtacacspolicy_add},
			{
				Config:                  testAccAuthenticationtacacspolicy_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckAuthenticationtacacspolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationtacacspolicy name is set")
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
		data, err := client.FindResource(service.Authenticationtacacspolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationtacacspolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationtacacspolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationtacacspolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Authenticationtacacspolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationtacacspolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAuthenticationtacacspolicyDataSource_basic = `
	resource "citrixadc_authenticationtacacsaction" "tf_tacacsaction_ds" {
		name            = "tf_tacacsaction_ds"
		serverip        = "1.2.3.5"
		serverport      = 8081
		authtimeout     = 6
		authorization   = "ON"
		accounting      = "ON"
		auditfailedcmds = "ON"
		groupattrname   = "group"
	}
	resource "citrixadc_authenticationtacacspolicy" "tf_tacacspolicy_ds" {
		name      = "tf_tacacspolicy_ds"
		rule      = "NS_TRUE"
		reqaction = citrixadc_authenticationtacacsaction.tf_tacacsaction_ds.name
	}
	data "citrixadc_authenticationtacacspolicy" "tf_tacacspolicy_data" {
		name = citrixadc_authenticationtacacspolicy.tf_tacacspolicy_ds.name
	}
`

func TestAccAuthenticationtacacspolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationtacacspolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationtacacspolicy.tf_tacacspolicy_data", "name", "tf_tacacspolicy_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationtacacspolicy.tf_tacacspolicy_data", "rule", "NS_TRUE"),
					resource.TestCheckResourceAttrPair("data.citrixadc_authenticationtacacspolicy.tf_tacacspolicy_data", "reqaction", "citrixadc_authenticationtacacspolicy.tf_tacacspolicy_ds", "reqaction"),
				),
			},
		},
	})
}
