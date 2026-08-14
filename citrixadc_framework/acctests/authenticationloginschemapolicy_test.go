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

const testAccAuthenticationloginschemapolicy_add = `
	resource "citrixadc_authenticationloginschema" "tf_loginschema" {
		name                    = "tf_loginschema"
		authenticationschema    = "LoginSchema/SingleAuth.xml"
		ssocredentials          = "YES"
		authenticationstrength  = "30"
		passwordcredentialindex = "10"
	}
	resource "citrixadc_authenticationloginschemapolicy" "tf_loginschemapolicy" {
		name      = "tf_loginschemapolicy"
		rule      = "true"
		action    = citrixadc_authenticationloginschema.tf_loginschema.name
		comment   = "sample_testing"
	}
`
const testAccAuthenticationloginschemapolicy_update = `
	resource "citrixadc_authenticationloginschema" "tf_loginschema" {
		name                    = "tf_loginschema"
		authenticationschema    = "LoginSchema/SingleAuth.xml"
		ssocredentials          = "YES"
		authenticationstrength  = "30"
		passwordcredentialindex = "10"
	}
	resource "citrixadc_authenticationloginschemapolicy" "tf_loginschemapolicy" {
		name      = "tf_loginschemapolicy"
		rule      = "false"
		action    = citrixadc_authenticationloginschema.tf_loginschema.name
		comment   = "samplenew_testing"
	}
`

func TestAccAuthenticationloginschemapolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationloginschemapolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationloginschemapolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationloginschemapolicyExist("citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy", "name", "tf_loginschemapolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy", "rule", "true"),
					resource.TestCheckResourceAttr("citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy", "comment", "sample_testing"),
				),
			},
			{
				Config: testAccAuthenticationloginschemapolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationloginschemapolicyExist("citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy", "name", "tf_loginschemapolicy"),
					resource.TestCheckResourceAttr("citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy", "rule", "false"),
					resource.TestCheckResourceAttr("citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy", "comment", "samplenew_testing"),
				),
			},
		},
	})
}

func TestAccAuthenticationloginschemapolicy_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAuthenticationloginschemapolicyDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccAuthenticationloginschemapolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationloginschemapolicyExist("citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccAuthenticationloginschemapolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationloginschemapolicyExist("citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy", nil)),
			},
		},
	})
}

func testAccCheckAuthenticationloginschemapolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationloginschemapolicy name is set")
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
		data, err := client.FindResource(service.Authenticationloginschemapolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationloginschemapolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationloginschemapolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationloginschemapolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Authenticationloginschemapolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationloginschemapolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccAuthenticationloginschemapolicy_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationloginschemapolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationloginschemapolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationloginschemapolicyExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Authenticationloginschemapolicy.Type(), "tf_loginschemapolicy"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAuthenticationloginschemapolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationloginschemapolicyExist(resAddr, nil)),
			},
		},
	})
}

func TestAccAuthenticationloginschemapolicy_import(t *testing.T) {
	const resAddr = "citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationloginschemapolicyDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAuthenticationloginschemapolicy_add},
			{
				Config:                  testAccAuthenticationloginschemapolicy_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

// comment is the sole cleanly-unsettable attribute: after a NITRO ?action=unset
// it reverts to its default (GET omits it). undefaction/logaction are excluded
// because their "unset" value reads back as the non-empty string "Use Global",
// which is rejected as a create/update input and would cause a perpetual diff.
const testAccAuthenticationloginschemapolicy_unset_step1 = `
	resource "citrixadc_authenticationloginschemapolicy" "tf_unset" {
		name    = "tf_test_lsp_unset"
		rule    = "true"
		action  = "NOOP"
		comment = "unset_probe_comment"
	}
`

const testAccAuthenticationloginschemapolicy_unset_step2 = `
	resource "citrixadc_authenticationloginschemapolicy" "tf_unset" {
		name   = "tf_test_lsp_unset"
		rule   = "true"
		action = "NOOP"
		# comment removed from config -> provider must unset it (revert to default).
	}
`

func TestAccAuthenticationloginschemapolicy_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationloginschemapolicyDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccAuthenticationloginschemapolicy_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationloginschemapolicyExist("citrixadc_authenticationloginschemapolicy.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationloginschemapolicy.tf_unset", "comment", "unset_probe_comment"),
					testAccCheckAuthenticationloginschemapolicyADCValue("tf_test_lsp_unset", "comment", "unset_probe_comment"),
				),
			},
			{
				// Removing comment must unset it: state reverts to the NITRO default
				// (GET omits comment -> null) and the implicit post-apply plan is empty.
				Config: testAccAuthenticationloginschemapolicy_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationloginschemapolicyExist("citrixadc_authenticationloginschemapolicy.tf_unset", nil),
					resource.TestCheckNoResourceAttr("citrixadc_authenticationloginschemapolicy.tf_unset", "comment"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAuthenticationloginschemapolicyADCValue("tf_test_lsp_unset", "comment", ""),
				),
			},
		},
	})
}

// testAccCheckAuthenticationloginschemapolicyADCValue asserts an attribute's value
// directly on the appliance (not just in Terraform state), proving the unset
// actually reverted it. A want of "" asserts the attribute is absent (default).
func testAccCheckAuthenticationloginschemapolicyADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Authenticationloginschemapolicy.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("authenticationloginschemapolicy %s not found on appliance", name)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("authenticationloginschemapolicy %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

const testAccAuthenticationloginschemapolicyDataSource_basic = `

resource "citrixadc_authenticationloginschema" "tf_loginschema_ds" {
	name                    = "tf_loginschema_ds"
	authenticationschema    = "LoginSchema/SingleAuth.xml"
	ssocredentials          = "YES"
	authenticationstrength  = "30"
	passwordcredentialindex = "10"
}

resource "citrixadc_authenticationloginschemapolicy" "tf_loginschemapolicy_ds" {
	name      = "tf_loginschemapolicy_ds"
	rule      = "true"
	action    = citrixadc_authenticationloginschema.tf_loginschema_ds.name
	comment   = "datasource_test"
}

data "citrixadc_authenticationloginschemapolicy" "tf_loginschemapolicy_ds" {
	name = citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy_ds.name
}
`

func TestAccAuthenticationloginschemapolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationloginschemapolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationloginschemapolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy_ds", "name", "tf_loginschemapolicy_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy_ds", "rule", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy_ds", "comment", "datasource_test"),
				),
			},
		},
	})
}
