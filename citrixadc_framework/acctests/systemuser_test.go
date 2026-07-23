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

func TestAccSystemuser_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemuserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemuser_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemuserExist("citrixadc_systemuser.tf_user", nil),
				),
			},
			{
				Config: testAccSystemuser_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemuserExist("citrixadc_systemuser.tf_user", nil),
				),
			},
		},
	})
}

func TestAccSystemuser_import(t *testing.T) {
	const resAddr = "citrixadc_systemuser.tf_user"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemuserDestroy,
		Steps: []resource.TestStep{
			{Config: testAccSystemuser_basic_step1},
			{
				Config:            testAccSystemuser_basic_step1,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// password: NITRO returns only the hashed password (tracked in
				// hashedpassword), so the plaintext cannot round-trip.
				// password_wo_version: a Computed version tracker not returned by
				// the API and not read back on import.
				// cmdpolicybinding: import starts with a null set, so the inline
				// bindings are not read back and cannot round-trip.
				ImportStateVerifyIgnore: []string{"password", "password_wo_version", "cmdpolicybinding"},
			},
		},
	})
}

func testAccCheckSystemuserExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lb vserver name is set")
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
		data, err := client.FindResource(service.Systemuser.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckSystemuserDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_systemuser" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Systemuser.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccSystemuser_basic_step1 = `
resource "citrixadc_systemuser" "tf_user" {
    username = "tf_user"
    password = "tf_password"
    timeout = 900

    cmdpolicybinding {
        policyname = "superuser"
        priority = 100
	}

    cmdpolicybinding {
        policyname = "network"
        priority = 200
	}
}
`

const testAccSystemuser_basic_step2 = `
resource "citrixadc_systemuser" "tf_user" {
    username = "tf_user"
    password = "tf_password"
    timeout = 200

}
`

const testAccSystemuserDataSource_basic = `
resource "citrixadc_systemuser" "tf_user" {
    username = "tf_user"
    password = "tf_password"
    timeout = 900

    cmdpolicybinding {
        policyname = "superuser"
        priority = 100
	}

    cmdpolicybinding {
        policyname = "network"
        priority = 200
	}
}

data "citrixadc_systemuser" "tf_user" {
    username = citrixadc_systemuser.tf_user.username
}
`

func TestAccSystemuserDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemuserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemuserDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_systemuser.tf_user", "username", "tf_user"),
					resource.TestCheckResourceAttr("data.citrixadc_systemuser.tf_user", "timeout", "900"),
				),
			},
		},
	})
}

// Test backward-compatible path: using password (Sensitive attribute)
const testAccSystemuser_password_step1 = `

	variable "systemuser_password" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_systemuser" "tf_user_secret" {
		username = "tf_user_secret"
		password = var.systemuser_password
		timeout  = 900
	}
`

const testAccSystemuser_password_step2 = `

	variable "systemuser_password_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_systemuser" "tf_user_secret" {
		username = "tf_user_secret"
		password = var.systemuser_password_2
		timeout  = 900
	}
`

func TestAccSystemuser_password_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_systemuser_password", "oldpass123")
	t.Setenv("TF_VAR_systemuser_password_2", "newpass456")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemuserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemuser_password_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemuserExist("citrixadc_systemuser.tf_user_secret", nil),
					resource.TestCheckResourceAttr("citrixadc_systemuser.tf_user_secret", "username", "tf_user_secret"),
					resource.TestCheckResourceAttr("citrixadc_systemuser.tf_user_secret", "timeout", "900"),
				),
			},
			{
				Config: testAccSystemuser_password_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemuserExist("citrixadc_systemuser.tf_user_secret", nil),
					resource.TestCheckResourceAttr("citrixadc_systemuser.tf_user_secret", "username", "tf_user_secret"),
					resource.TestCheckResourceAttr("citrixadc_systemuser.tf_user_secret", "timeout", "900"),
				),
			},
		},
	})
}

// Test ephemeral path: using password_wo (WriteOnly attribute) with version tracker
const testAccSystemuser_password_wo_step1 = `

	variable "systemuser_password_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_systemuser" "tf_user_ephemeral" {
		username            = "tf_user_ephemeral"
		password_wo         = var.systemuser_password_wo
		password_wo_version = 1
		timeout             = 900
	}
`

const testAccSystemuser_password_wo_step2 = `

	variable "systemuser_password_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_systemuser" "tf_user_ephemeral" {
		username            = "tf_user_ephemeral"
		password_wo         = var.systemuser_password_wo_2
		password_wo_version = 2
		timeout             = 900
	}
`

func TestAccSystemuser_password_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_systemuser_password_wo", "ephemeral_pass1")
	t.Setenv("TF_VAR_systemuser_password_wo_2", "ephemeral_pass2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemuserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemuser_password_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemuserExist("citrixadc_systemuser.tf_user_ephemeral", nil),
					resource.TestCheckResourceAttr("citrixadc_systemuser.tf_user_ephemeral", "password_wo_version", "1"),
				),
			},
			{
				Config: testAccSystemuser_password_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemuserExist("citrixadc_systemuser.tf_user_ephemeral", nil),
					resource.TestCheckResourceAttr("citrixadc_systemuser.tf_user_ephemeral", "password_wo_version", "2"),
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// Unset support test
// ---------------------------------------------------------------------------

// Step 1: unset-eligible attributes (externalauth, logging, timeout) set to
// non-default values so we can prove they take effect.
const testAccSystemuser_unset_step1 = `
resource "citrixadc_systemuser" "tf_unset" {
    username     = "tf_test_systemuser_unset"
    password     = "tf_unset_password"
    externalauth = "DISABLED"
    logging      = "ENABLED"
    timeout      = 300
}
`

// Step 2: the unset-eligible attributes are removed from configuration, so the
// provider must issue ?action=unset and the appliance reverts each to its
// default (externalauth=ENABLED, logging=DISABLED, timeout=900).
const testAccSystemuser_unset_step2 = `
resource "citrixadc_systemuser" "tf_unset" {
    username = "tf_test_systemuser_unset"
    password = "tf_unset_password"
}
`

func TestAccSystemuser_unset(t *testing.T) {
	// No skip guards on the other systemuser tests: they run on the default
	// standalone testbed, so this test mirrors that (no guard).
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemuserDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values apply and persist.
				Config: testAccSystemuser_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemuserExist("citrixadc_systemuser.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_systemuser.tf_unset", "externalauth", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemuser.tf_unset", "logging", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemuser.tf_unset", "timeout", "300"),
				),
			},
			{
				// Removing them must unset -> state reverts to NITRO defaults,
				// and the implicit post-apply plan must be empty.
				Config: testAccSystemuser_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemuserExist("citrixadc_systemuser.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_systemuser.tf_unset", "externalauth", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemuser.tf_unset", "logging", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemuser.tf_unset", "timeout", "900"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckSystemuserADCValue("tf_test_systemuser_unset", "externalauth", "ENABLED"),
					testAccCheckSystemuserADCValue("tf_test_systemuser_unset", "logging", "DISABLED"),
					testAccCheckSystemuserADCValue("tf_test_systemuser_unset", "timeout", "900"),
				),
			},
		},
	})
}

// testAccCheckSystemuserADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckSystemuserADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Systemuser.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("systemuser %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("systemuser %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccSystemuser_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckSystemuserDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccSystemuser_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemuserExist("citrixadc_systemuser.tf_user", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccSystemuser_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemuserExist("citrixadc_systemuser.tf_user", nil),
				),
			},
		},
	})
}
