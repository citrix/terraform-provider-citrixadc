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

const testAccTmformssoaction_basic = `

	resource "citrixadc_tmformssoaction" "tf_tmformssoaction" {
		name           = "my_formsso_action"
		actionurl      = "/logon.php"
		userfield      = "loginID"
		passwdfield    = "passwd"
		ssosuccessrule = "HTTP.RES.HEADER(\"Set-Cookie\").CONTAINS(\"LogonID\")"
	}
`
const testAccTmformssoaction_update = `

	resource "citrixadc_tmformssoaction" "tf_tmformssoaction" {
		name           = "my_formsso_action"
		actionurl      = "/main/logon.php"
		userfield      = "loginID2"
		passwdfield    = "passwd2"
		ssosuccessrule = "HTTP.RES.HEADER(\"Set-Cookie\").CONTAINS(\"LogonID\")"
	}
`

const testAccTmformssoactionDataSource_basic = `

	resource "citrixadc_tmformssoaction" "tf_tmformssoaction" {
		name           = "my_formsso_action"
		actionurl      = "/logon.php"
		userfield      = "loginID"
		passwdfield    = "passwd"
		ssosuccessrule = "HTTP.RES.HEADER(\"Set-Cookie\").CONTAINS(\"LogonID\")"
	}

	data "citrixadc_tmformssoaction" "tf_tmformssoaction" {
		name = citrixadc_tmformssoaction.tf_tmformssoaction.name
	}
`

func TestAccTmformssoaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTmformssoactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTmformssoaction_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmformssoactionExist("citrixadc_tmformssoaction.tf_tmformssoaction", nil),
					resource.TestCheckResourceAttr("citrixadc_tmformssoaction.tf_tmformssoaction", "name", "my_formsso_action"),
					resource.TestCheckResourceAttr("citrixadc_tmformssoaction.tf_tmformssoaction", "actionurl", "/logon.php"),
					resource.TestCheckResourceAttr("citrixadc_tmformssoaction.tf_tmformssoaction", "userfield", "loginID"),
					resource.TestCheckResourceAttr("citrixadc_tmformssoaction.tf_tmformssoaction", "passwdfield", "passwd"),
					resource.TestCheckResourceAttr("citrixadc_tmformssoaction.tf_tmformssoaction", "ssosuccessrule", "HTTP.RES.HEADER(\"Set-Cookie\").CONTAINS(\"LogonID\")"),
				),
			},
			{
				Config: testAccTmformssoaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmformssoactionExist("citrixadc_tmformssoaction.tf_tmformssoaction", nil),
					resource.TestCheckResourceAttr("citrixadc_tmformssoaction.tf_tmformssoaction", "name", "my_formsso_action"),
					resource.TestCheckResourceAttr("citrixadc_tmformssoaction.tf_tmformssoaction", "actionurl", "/main/logon.php"),
					resource.TestCheckResourceAttr("citrixadc_tmformssoaction.tf_tmformssoaction", "userfield", "loginID2"),
					resource.TestCheckResourceAttr("citrixadc_tmformssoaction.tf_tmformssoaction", "passwdfield", "passwd2"),
					resource.TestCheckResourceAttr("citrixadc_tmformssoaction.tf_tmformssoaction", "ssosuccessrule", "HTTP.RES.HEADER(\"Set-Cookie\").CONTAINS(\"LogonID\")"),
				),
			},
		},
	})
}

func testAccCheckTmformssoactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No tmformssoaction name is set")
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
		data, err := client.FindResource(service.Tmformssoaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("tmformssoaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckTmformssoactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_tmformssoaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Tmformssoaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("tmformssoaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccTmformssoaction_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_tmformssoaction.tf_tmformssoaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTmformssoactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTmformssoaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTmformssoactionExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Tmformssoaction.Type(), "my_formsso_action"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccTmformssoaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTmformssoactionExist(resAddr, nil)),
			},
		},
	})
}

func TestAccTmformssoaction_import(t *testing.T) {
	const resAddr = "citrixadc_tmformssoaction.tf_tmformssoaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTmformssoactionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccTmformssoaction_basic},
			{
				Config:                  testAccTmformssoaction_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccTmformssoaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckTmformssoactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccTmformssoaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTmformssoactionExist("citrixadc_tmformssoaction.tf_tmformssoaction", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccTmformssoaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTmformssoactionExist("citrixadc_tmformssoaction.tf_tmformssoaction", nil)),
			},
		},
	})
}

// Unset test: nvtype, responsesize and submitmethod have documented NITRO
// defaults (DYNAMIC / 8096 / GET). submitmethod applies only to STATIC
// name-value type, so step1 sets nvtype = STATIC to make submitmethod valid.
const testAccTmformssoaction_unset_step1 = `
	resource "citrixadc_tmformssoaction" "tf_unset" {
		name           = "tf_test_tmformssoaction_unset"
		actionurl      = "/logon.php"
		userfield      = "loginID"
		passwdfield    = "passwd"
		ssosuccessrule = "HTTP.RES.HEADER(\"Set-Cookie\").CONTAINS(\"LogonID\")"
		nvtype         = "STATIC"
		submitmethod   = "POST"
		responsesize   = 4096
	}
`

const testAccTmformssoaction_unset_step2 = `
	resource "citrixadc_tmformssoaction" "tf_unset" {
		name           = "tf_test_tmformssoaction_unset"
		actionurl      = "/logon.php"
		userfield      = "loginID"
		passwdfield    = "passwd"
		ssosuccessrule = "HTTP.RES.HEADER(\"Set-Cookie\").CONTAINS(\"LogonID\")"
		# nvtype, submitmethod, responsesize removed -> provider must unset them
		# (revert to NITRO defaults DYNAMIC / GET / 8096).
	}
`

func TestAccTmformssoaction_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTmformssoactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTmformssoaction_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmformssoactionExist("citrixadc_tmformssoaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_tmformssoaction.tf_unset", "nvtype", "STATIC"),
					resource.TestCheckResourceAttr("citrixadc_tmformssoaction.tf_unset", "submitmethod", "POST"),
					resource.TestCheckResourceAttr("citrixadc_tmformssoaction.tf_unset", "responsesize", "4096"),
				),
			},
			{
				Config: testAccTmformssoaction_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmformssoactionExist("citrixadc_tmformssoaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_tmformssoaction.tf_unset", "nvtype", "DYNAMIC"),
					resource.TestCheckResourceAttr("citrixadc_tmformssoaction.tf_unset", "submitmethod", "GET"),
					resource.TestCheckResourceAttr("citrixadc_tmformssoaction.tf_unset", "responsesize", "8096"),
					testAccCheckTmformssoactionADCValue("tf_test_tmformssoaction_unset", "nvtype", "DYNAMIC"),
					testAccCheckTmformssoactionADCValue("tf_test_tmformssoaction_unset", "submitmethod", "GET"),
					testAccCheckTmformssoactionADCValue("tf_test_tmformssoaction_unset", "responsesize", "8096"),
				),
			},
		},
	})
}

// testAccCheckTmformssoactionADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset took effect.
func testAccCheckTmformssoactionADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Tmformssoaction.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("tmformssoaction %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("tmformssoaction %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccTmformssoactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTmformssoactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTmformssoactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_tmformssoaction.tf_tmformssoaction", "name", "my_formsso_action"),
					resource.TestCheckResourceAttr("data.citrixadc_tmformssoaction.tf_tmformssoaction", "actionurl", "/logon.php"),
					resource.TestCheckResourceAttr("data.citrixadc_tmformssoaction.tf_tmformssoaction", "userfield", "loginID"),
					resource.TestCheckResourceAttr("data.citrixadc_tmformssoaction.tf_tmformssoaction", "passwdfield", "passwd"),
					resource.TestCheckResourceAttr("data.citrixadc_tmformssoaction.tf_tmformssoaction", "ssosuccessrule", "HTTP.RES.HEADER(\"Set-Cookie\").CONTAINS(\"LogonID\")"),
				),
			},
		},
	})
}
