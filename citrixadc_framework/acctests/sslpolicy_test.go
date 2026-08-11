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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccSslpolicy_add = `
	resource "citrixadc_sslaction" "foo" {
	name                   = "tf_sslaction"
	clientauth             = "DOCLIENTAUTH"
	clientcertverification = "Mandatory"
	}

	resource "citrixadc_sslpolicy" "foo" {
	name   = "tf_sslpolicy"
	rule   = "false"
	action = citrixadc_sslaction.foo.name
	}
`

const testAccSslpolicy_update = `
	resource "citrixadc_sslaction" "foo" {
	name                   = "tf_sslaction"
	clientauth             = "DOCLIENTAUTH"
	clientcertverification = "Mandatory"
	}

	resource "citrixadc_sslpolicy" "foo" {
	name   = "tf_sslpolicy"
	rule   = "true"
	action = citrixadc_sslaction.foo.name
	}
`

const testAccSslpolicyDataSource_basic = `
	resource "citrixadc_sslaction" "foo" {
	name                   = "tf_sslaction"
	clientauth             = "DOCLIENTAUTH"
	clientcertverification = "Mandatory"
	}

	resource "citrixadc_sslpolicy" "foo" {
	name   = "tf_sslpolicy"
	rule   = "false"
	action = citrixadc_sslaction.foo.name
	}

	data "citrixadc_sslpolicy" "foo" {
		name = citrixadc_sslpolicy.foo.name
	}
`

func TestAccSslpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslpolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslpolicyExist("citrixadc_sslpolicy.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslpolicy.foo", "name", "tf_sslpolicy"),
					resource.TestCheckResourceAttr("citrixadc_sslpolicy.foo", "rule", "false"),
					resource.TestCheckResourceAttr("citrixadc_sslpolicy.foo", "action", "tf_sslaction"),
				),
			},
			{
				Config: testAccSslpolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslpolicyExist("citrixadc_sslpolicy.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslpolicy.foo", "name", "tf_sslpolicy"),
					resource.TestCheckResourceAttr("citrixadc_sslpolicy.foo", "rule", "true"),
					resource.TestCheckResourceAttr("citrixadc_sslpolicy.foo", "action", "tf_sslaction"),
				),
			},
		},
	})
}

func testAccCheckSslpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SSL Policy name is set")
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
		data, err := client.FindResource(service.Sslpolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("SSL Policy %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslpolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Sslpolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("SSL Policy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccSslpolicy_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_sslpolicy.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSslpolicyExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Sslpolicy.Type(), "tf_sslpolicy"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccSslpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSslpolicyExist(resAddr, nil)),
			},
		},
	})
}

func TestAccSslpolicy_import(t *testing.T) {
	const resAddr = "citrixadc_sslpolicy.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslpolicyDestroy,
		Steps: []resource.TestStep{
			{Config: testAccSslpolicy_add},
			{
				Config:                  testAccSslpolicy_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccSslpolicy_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckSslpolicyDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccSslpolicy_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSslpolicyExist("citrixadc_sslpolicy.foo", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccSslpolicy_add,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckSslpolicyExist("citrixadc_sslpolicy.foo", nil)),
			},
		},
	})
}

// testAccSslpolicy_unset_step1 sets the unset-eligible attributes (comment,
// undefaction) to valid non-default values.
const testAccSslpolicy_unset_step1 = `
	resource "citrixadc_sslaction" "foo" {
	name                   = "tf_sslaction"
	clientauth             = "DOCLIENTAUTH"
	clientcertverification = "Mandatory"
	}

	resource "citrixadc_sslpolicy" "foo" {
	name        = "tf_sslpolicy"
	rule        = "true"
	action      = citrixadc_sslaction.foo.name
	undefaction = "RESET"
	comment     = "managed by tf"
	}
`

// testAccSslpolicy_unset_step2 removes comment and undefaction from config so
// the provider must unset them (revert to NITRO defaults, empty).
const testAccSslpolicy_unset_step2 = `
	resource "citrixadc_sslaction" "foo" {
	name                   = "tf_sslaction"
	clientauth             = "DOCLIENTAUTH"
	clientcertverification = "Mandatory"
	}

	resource "citrixadc_sslpolicy" "foo" {
	name   = "tf_sslpolicy"
	rule   = "true"
	action = citrixadc_sslaction.foo.name
	}
`

func TestAccSslpolicy_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslpolicyDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccSslpolicy_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslpolicyExist("citrixadc_sslpolicy.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_sslpolicy.foo", "undefaction", "RESET"),
					resource.TestCheckResourceAttr("citrixadc_sslpolicy.foo", "comment", "managed by tf"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the NITRO defaults (empty), and the
				// implicit post-apply plan must be empty.
				Config: testAccSslpolicy_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslpolicyExist("citrixadc_sslpolicy.foo", nil),
					resource.TestCheckNoResourceAttr("citrixadc_sslpolicy.foo", "undefaction"),
					resource.TestCheckNoResourceAttr("citrixadc_sslpolicy.foo", "comment"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckSslpolicyADCValue("tf_sslpolicy", "undefaction", ""),
					testAccCheckSslpolicyADCValue("tf_sslpolicy", "comment", ""),
				),
			},
		},
	})
}

// testAccCheckSslpolicyADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it. NITRO omits unset attributes from GET, so an absent value is treated as
// empty.
func testAccCheckSslpolicyADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Sslpolicy.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("sslpolicy %s not found on appliance", name)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("sslpolicy %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccSslpolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslpolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslpolicy.foo", "name", "tf_sslpolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_sslpolicy.foo", "rule", "false"),
					resource.TestCheckResourceAttr("data.citrixadc_sslpolicy.foo", "action", "tf_sslaction"),
				),
			},
		},
	})
}
