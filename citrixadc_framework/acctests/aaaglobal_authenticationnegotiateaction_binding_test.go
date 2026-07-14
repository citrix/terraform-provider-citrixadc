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

// aaaglobal_authenticationnegotiateaction_binding is a KEYLESS global binding.
// aaaglobal is a singleton (no parent name); the composite ID is the plain
// windowsprofile value. The bound entity (a negotiate profile) is created via the
// existing authenticationnegotiateaction acceptance test config, then bound globally.

// step1: create the negotiate action + global binding
const testAccAaaglobalAuthenticationnegotiateactionBinding_basic_step1 = `
resource "citrixadc_authenticationnegotiateaction" "tf_negotiateaction" {
	name                       = "tf_negotiateaction_binding"
	domain                     = "DomainName"
	domainuser                 = "username"
	domainuserpasswd           = "password"
	ntlmpath                   = "http://www.example.com/"
	defaultauthenticationgroup = "grpname"
}

resource "citrixadc_aaaglobal_authenticationnegotiateaction_binding" "tf_aaaglobal_authenticationnegotiateaction_binding" {
	windowsprofile = citrixadc_authenticationnegotiateaction.tf_negotiateaction.name
	depends_on     = [citrixadc_authenticationnegotiateaction.tf_negotiateaction]
}
`

// step2: drop the binding (keep the negotiate action) to verify the binding is deleted
const testAccAaaglobalAuthenticationnegotiateactionBinding_basic_step2 = `
resource "citrixadc_authenticationnegotiateaction" "tf_negotiateaction" {
	name                       = "tf_negotiateaction_binding"
	domain                     = "DomainName"
	domainuser                 = "username"
	domainuserpasswd           = "password"
	ntlmpath                   = "http://www.example.com/"
	defaultauthenticationgroup = "grpname"
}
`

func TestAccAaaglobalAuthenticationnegotiateactionBinding_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAaaglobalAuthenticationnegotiateactionBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAaaglobalAuthenticationnegotiateactionBinding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaglobalAuthenticationnegotiateactionBindingExist("citrixadc_aaaglobal_authenticationnegotiateaction_binding.tf_aaaglobal_authenticationnegotiateaction_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaglobal_authenticationnegotiateaction_binding.tf_aaaglobal_authenticationnegotiateaction_binding", "windowsprofile", "tf_negotiateaction_binding"),
				),
			},
			{
				// Binding dropped from config: verify it no longer exists on the ADC.
				Config: testAccAaaglobalAuthenticationnegotiateactionBinding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaglobalAuthenticationnegotiateactionBindingNotExist(t, "tf_negotiateaction_binding"),
				),
			},
		},
	})
}

func testAccCheckAaaglobalAuthenticationnegotiateactionBindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No aaaglobal_authenticationnegotiateaction_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Keyless global binding: ID is the plain windowsprofile value.
		windowsprofile := rs.Primary.ID

		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		// Mirror the resource Read: list all bindings (no parent name) and filter on windowsprofile.
		findParams := service.FindParams{
			ResourceType:             service.Aaaglobal_authenticationnegotiateaction_binding.Type(),
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if val, ok := v["windowsprofile"].(string); ok && val == windowsprofile {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("aaaglobal_authenticationnegotiateaction_binding %s not found", windowsprofile)
		}

		return nil
	}
}

// testAccCheckAaaglobalAuthenticationnegotiateactionBindingNotExist verifies that the
// binding for the given windowsprofile is absent from the ADC.
func testAccCheckAaaglobalAuthenticationnegotiateactionBindingNotExist(t *testing.T, windowsprofile string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		findParams := service.FindParams{
			ResourceType:             service.Aaaglobal_authenticationnegotiateaction_binding.Type(),
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		for _, v := range dataArr {
			if val, ok := v["windowsprofile"].(string); ok && val == windowsprofile {
				return fmt.Errorf("aaaglobal_authenticationnegotiateaction_binding %s still exists after being dropped from config", windowsprofile)
			}
		}

		return nil
	}
}

func testAccCheckAaaglobalAuthenticationnegotiateactionBindingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_aaaglobal_authenticationnegotiateaction_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No id is set")
		}

		// Keyless global binding: ID is the plain windowsprofile value.
		windowsprofile := rs.Primary.ID

		findParams := service.FindParams{
			ResourceType:             service.Aaaglobal_authenticationnegotiateaction_binding.Type(),
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Treat a missing/empty resource as successfully destroyed.
			continue
		}

		for _, v := range dataArr {
			if val, ok := v["windowsprofile"].(string); ok && val == windowsprofile {
				return fmt.Errorf("aaaglobal_authenticationnegotiateaction_binding %s still exists", windowsprofile)
			}
		}
	}

	return nil
}

const testAccAaaglobalAuthenticationnegotiateactionBindingDataSource_basic = `
resource "citrixadc_authenticationnegotiateaction" "tf_negotiateaction_ds" {
	name                       = "tf_negotiateaction_binding_ds"
	domain                     = "DomainName"
	domainuser                 = "username"
	domainuserpasswd           = "password"
	ntlmpath                   = "http://www.example.com/"
	defaultauthenticationgroup = "grpname"
}

resource "citrixadc_aaaglobal_authenticationnegotiateaction_binding" "tf_aaaglobal_authenticationnegotiateaction_binding_ds" {
	windowsprofile = citrixadc_authenticationnegotiateaction.tf_negotiateaction_ds.name
	depends_on     = [citrixadc_authenticationnegotiateaction.tf_negotiateaction_ds]
}

data "citrixadc_aaaglobal_authenticationnegotiateaction_binding" "tf_aaaglobal_authenticationnegotiateaction_binding_ds" {
	windowsprofile = citrixadc_aaaglobal_authenticationnegotiateaction_binding.tf_aaaglobal_authenticationnegotiateaction_binding_ds.windowsprofile
	depends_on     = [citrixadc_aaaglobal_authenticationnegotiateaction_binding.tf_aaaglobal_authenticationnegotiateaction_binding_ds]
}
`

func TestAccAaaglobalAuthenticationnegotiateactionBindingDataSource_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAaaglobalAuthenticationnegotiateactionBindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_aaaglobal_authenticationnegotiateaction_binding.tf_aaaglobal_authenticationnegotiateaction_binding_ds", "windowsprofile", "tf_negotiateaction_binding_ds"),
				),
			},
		},
	})
}
