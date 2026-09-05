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

const testAccTmsessionpolicy_basic = `

	resource "citrixadc_tmsessionaction" "tf_tmsessionaction" {
		name                       = "tf_tmsessaction"
		sesstimeout                = 10
		defaultauthorizationaction = "ALLOW"
		sso                        = "OFF"
	}
	resource "citrixadc_tmsessionpolicy" "tf_tmsessionpolicy" {
		name   = "my_tmsession_policy"
		rule   = "true"
		action = citrixadc_tmsessionaction.tf_tmsessionaction.name
	}
`
const testAccTmsessionpolicy_update = `


	resource "citrixadc_tmsessionaction" "tf_tmsessionaction" {
		name                       = "tf_tmsessaction"
		sesstimeout                = 10
		defaultauthorizationaction = "ALLOW"
		sso                        = "OFF"
	}
	resource "citrixadc_tmsessionpolicy" "tf_tmsessionpolicy" {
		name   = "my_tmsession_policy"
		rule   = "false"
		action = citrixadc_tmsessionaction.tf_tmsessionaction.name
	}
`

func TestAccTmsessionpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTmsessionpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTmsessionpolicy_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmsessionpolicyExist("citrixadc_tmsessionpolicy.tf_tmsessionpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_tmsessionpolicy.tf_tmsessionpolicy", "name", "my_tmsession_policy"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionpolicy.tf_tmsessionpolicy", "rule", "true"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionpolicy.tf_tmsessionpolicy", "action", "tf_tmsessaction"),
				),
			},
			{
				Config: testAccTmsessionpolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmsessionpolicyExist("citrixadc_tmsessionpolicy.tf_tmsessionpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_tmsessionpolicy.tf_tmsessionpolicy", "name", "my_tmsession_policy"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionpolicy.tf_tmsessionpolicy", "rule", "false"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionpolicy.tf_tmsessionpolicy", "action", "tf_tmsessaction"),
				),
			},
		},
	})
}

func TestAccTmsessionpolicy_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_tmsessionpolicy.tf_tmsessionpolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTmsessionpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTmsessionpolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTmsessionpolicyExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Tmsessionpolicy.Type(), "my_tmsession_policy"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccTmsessionpolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTmsessionpolicyExist(resAddr, nil)),
			},
		},
	})
}

func TestAccTmsessionpolicy_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckTmsessionpolicyDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccTmsessionpolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTmsessionpolicyExist("citrixadc_tmsessionpolicy.tf_tmsessionpolicy", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccTmsessionpolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTmsessionpolicyExist("citrixadc_tmsessionpolicy.tf_tmsessionpolicy", nil)),
			},
		},
	})
}

func testAccCheckTmsessionpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No tmsessionpolicy name is set")
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
		data, err := client.FindResource(service.Tmsessionpolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("tmsessionpolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckTmsessionpolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_tmsessionpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Tmsessionpolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("tmsessionpolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccTmsessionpolicy_import(t *testing.T) {
	const resAddr = "citrixadc_tmsessionpolicy.tf_tmsessionpolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTmsessionpolicyDestroy,
		Steps: []resource.TestStep{
			{Config: testAccTmsessionpolicy_basic},
			{
				Config:                  testAccTmsessionpolicy_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccTmsessionpolicyDataSource_basic = `

	resource "citrixadc_tmsessionaction" "tf_tmsessionaction" {
		name                       = "tf_tmsessaction"
		sesstimeout                = 10
		defaultauthorizationaction = "ALLOW"
		sso                        = "OFF"
	}
	resource "citrixadc_tmsessionpolicy" "tf_tmsessionpolicy" {
		name   = "my_tmsession_policy"
		rule   = "true"
		action = citrixadc_tmsessionaction.tf_tmsessionaction.name
	}

data "citrixadc_tmsessionpolicy" "tf_tmsessionpolicy" {
    name = citrixadc_tmsessionpolicy.tf_tmsessionpolicy.name
}
`

func TestAccTmsessionpolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTmsessionpolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					// "id" is the universal runtime-binding proof (equals name).
					resource.TestCheckResourceAttrSet("data.citrixadc_tmsessionpolicy.tf_tmsessionpolicy", "id"),
					resource.TestCheckResourceAttr("data.citrixadc_tmsessionpolicy.tf_tmsessionpolicy", "name", "my_tmsession_policy"),
					resource.TestCheckResourceAttr("data.citrixadc_tmsessionpolicy.tf_tmsessionpolicy", "rule", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_tmsessionpolicy.tf_tmsessionpolicy", "action", "tf_tmsessaction"),
				),
			},
		},
	})
}
