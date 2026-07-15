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

// ============================================================================
// IMPORTANT CAVEAT - direct cipher binding requires the DEFAULT SSL PROFILE to
// be DISABLED on the appliance.
//
// When the default SSL profile is ENABLED on the appliance, binding a cipher
// directly to an SSL vserver is NOT permitted - the NITRO API rejects it with a
// "Use profile command, vserver/service has profile feature enabled" style error.
// In that mode ciphers must be managed through the SSL profile instead.
//
// To run TestAccSslvserver_sslcipher_binding_basic / ...DataSource_basic the
// testbed must have the default SSL profile DISABLED (CLI: `unset ssl parameter
// -defaultProfile` / set defaultProfile to DISABLED). On a testbed with the
// non-default-SSL-profile setting (ADC_TESTBED == "STANDALONE_NON_DEFAULT_SSL_PROFILE")
// these tests will fail and should be skipped. Generate-only: no run performed.
// ============================================================================

import (
	"fmt"
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// sslvserver_sslcipher_binding joins an SSL vserver (the SSL view of an SSL-type
// lbvserver) with a built-in cipher / cipher alias (e.g. "HIGH").
//
// Composite ID = vservername,ciphername. By-name GET works:
// FindResourceArrayWithParams keyed on vservername, then match ciphername in the
// returned array.

const testAccSslvserver_sslcipher_binding_basic_step1 = `
	resource "citrixadc_lbvserver" "tf_sslvserver" {
		name        = "tf_sslvserver"
		servicetype = "SSL"
		ipv46       = "5.5.5.5"
		port        = 443
	}

	resource "citrixadc_sslvserver_sslcipher_binding" "tf_sslvserver_sslcipher_binding" {
		vservername = citrixadc_lbvserver.tf_sslvserver.name
		ciphername  = "HIGH"
	}
`

// step2 drops the binding but keeps the SSL vserver, so the binding's removal can
// be verified.
const testAccSslvserver_sslcipher_binding_basic_step2 = `
	resource "citrixadc_lbvserver" "tf_sslvserver" {
		name        = "tf_sslvserver"
		servicetype = "SSL"
		ipv46       = "5.5.5.5"
		port        = 443
	}
`

func TestAccSslvserver_sslcipher_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslvserver_sslcipher_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslvserver_sslcipher_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_sslcipher_bindingExist("citrixadc_sslvserver_sslcipher_binding.tf_sslvserver_sslcipher_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_sslvserver_sslcipher_binding.tf_sslvserver_sslcipher_binding", "vservername", "tf_sslvserver"),
					resource.TestCheckResourceAttr("citrixadc_sslvserver_sslcipher_binding.tf_sslvserver_sslcipher_binding", "ciphername", "HIGH"),
				),
			},
			{
				Config: testAccSslvserver_sslcipher_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_sslcipher_bindingNotExist("citrixadc_sslvserver_sslcipher_binding.tf_sslvserver_sslcipher_binding", "tf_sslvserver,HIGH"),
				),
			},
		},
	})
}

func TestAccSslvserver_sslcipher_binding_import(t *testing.T) {
	const resAddr = "citrixadc_sslvserver_sslcipher_binding.tf_sslvserver_sslcipher_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslvserver_sslcipher_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslvserver_sslcipher_binding_basic_step1,
			},
			{
				Config:                  testAccSslvserver_sslcipher_binding_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckSslvserver_sslcipher_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslvserver_sslcipher_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}
		vservername := idMap["vservername"]
		ciphername := idMap["ciphername"]

		// The typed sslvserver_sslcipher_binding GET is not reflected over REST on this
		// firmware; the binding is only visible via the umbrella sslvserver_binding
		// endpoint (under sslvserver_sslciphersuite_binding[]). See the resource Read.
		found, err := sslvserverSslcipherBindingExistsViaUmbrella(client, vservername, ciphername)
		if err != nil {
			return err
		}
		if !found {
			return fmt.Errorf("sslvserver_sslcipher_binding %s not found on ADC", rs.Primary.ID)
		}

		return nil
	}
}

// sslvserverSslcipherBindingExistsViaUmbrella reports whether the cipher is bound to
// the SSL vserver, reading the umbrella sslvserver_binding endpoint (the typed binding
// GET returns an empty body on this firmware).
func sslvserverSslcipherBindingExistsViaUmbrella(client *service.NitroClient, vservername, ciphername string) (bool, error) {
	umbrella, err := client.FindResourceArrayWithParams(service.FindParams{
		ResourceType:             service.Sslvserver_binding.Type(),
		ResourceName:             vservername,
		ResourceMissingErrorCode: 258,
	})
	if err != nil {
		return false, err
	}
	for _, vs := range umbrella {
		raw, ok := vs["sslvserver_sslciphersuite_binding"]
		if !ok {
			continue
		}
		rows, ok := raw.([]interface{})
		if !ok {
			continue
		}
		for _, cb := range rows {
			m, ok := cb.(map[string]interface{})
			if !ok {
				continue
			}
			if v, ok := m["ciphername"].(string); ok && v == ciphername {
				return true, nil
			}
			if v, ok := m["cipheraliasname"].(string); ok && v == ciphername {
				return true, nil
			}
		}
	}
	return false, nil
}

// testAccCheckSslvserver_sslcipher_bindingNotExist verifies the binding is gone while
// the parent vserver still exists (step2 keeps the vserver). id is "vservername,ciphername".
func testAccCheckSslvserver_sslcipher_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idSlice := strings.SplitN(id, ",", 2)
		if len(idSlice) != 2 {
			return fmt.Errorf("Invalid id string %v. Expected \"vservername,ciphername\".", id)
		}
		vservername := idSlice[0]
		ciphername := idSlice[1]

		found, err := sslvserverSslcipherBindingExistsViaUmbrella(client, vservername, ciphername)
		if err != nil {
			return nil
		}
		if found {
			return fmt.Errorf("sslvserver_sslcipher_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckSslvserver_sslcipher_bindingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslvserver_sslcipher_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}
		vservername := idMap["vservername"]
		ciphername := idMap["ciphername"]

		found, err := sslvserverSslcipherBindingExistsViaUmbrella(client, vservername, ciphername)
		if err != nil {
			continue
		}
		if found {
			return fmt.Errorf("sslvserver_sslcipher_binding %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

const testAccSslvserver_sslcipher_bindingDataSource_basic = `
	resource "citrixadc_lbvserver" "tf_sslvserver" {
		name        = "tf_sslvserver"
		servicetype = "SSL"
		ipv46       = "5.5.5.5"
		port        = 443
	}

	resource "citrixadc_sslvserver_sslcipher_binding" "tf_sslvserver_sslcipher_binding" {
		vservername = citrixadc_lbvserver.tf_sslvserver.name
		ciphername  = "HIGH"
	}

	data "citrixadc_sslvserver_sslcipher_binding" "tf_sslvserver_sslcipher_binding" {
		vservername = citrixadc_lbvserver.tf_sslvserver.name
		ciphername  = "HIGH"
		depends_on  = [citrixadc_sslvserver_sslcipher_binding.tf_sslvserver_sslcipher_binding]
	}
`

func TestAccSslvserver_sslcipher_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslvserver_sslcipher_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslvserver_sslcipher_bindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslvserver_sslcipher_binding.tf_sslvserver_sslcipher_binding", "vservername", "tf_sslvserver"),
					resource.TestCheckResourceAttr("data.citrixadc_sslvserver_sslcipher_binding.tf_sslvserver_sslcipher_binding", "ciphername", "HIGH"),
				),
			},
		},
	})
}
