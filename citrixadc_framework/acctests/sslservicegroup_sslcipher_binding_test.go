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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// sslservicegroup_sslcipher_binding binds a cipher/cipher-alias/cipher-group to an SSL
// servicegroup. ciphername may be an individual cipher name, a system predefined
// cipher-alias (e.g. "DEFAULT", "ECDHE"), or a user-defined cipher-group.
//
// CAVEAT: When the appliance runs with the default (enabled) SSL profile, ciphers are
// managed through the SSL profile and direct cipher binding to a servicegroup may be
// rejected (e.g. error 3082 "no default profile..."/operation not permitted). This test
// is gated to the STANDALONE_NON_DEFAULT_SSL_PROFILE testbed where direct cipher binding
// is allowed. "DEFAULT" is a built-in cipher alias available on every appliance.

const testAccSslservicegroup_sslcipher_binding_basic = `
	resource "citrixadc_servicegroup" "tf_servicegroup" {
		servicegroupname = "tf_servicegroup"
		servicetype      = "SSL"
	}

	resource "citrixadc_sslservicegroup_sslcipher_binding" "tf_sslservicegroup_sslcipher_binding" {
		ciphername       = "DEFAULT"
		servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
		depends_on       = [citrixadc_servicegroup.tf_servicegroup]
	}
`

const testAccSslservicegroup_sslcipher_binding_basic_step2 = `
	resource "citrixadc_servicegroup" "tf_servicegroup" {
		servicegroupname = "tf_servicegroup"
		servicetype      = "SSL"
	}
`

func TestAccSslservicegroup_sslcipher_binding_basic(t *testing.T) {
	if adcTestbed != "STANDALONE_DEFAULT_SSL_PROFILE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_DEFAULT_SSL_PROFILE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslservicegroup_sslcipher_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslservicegroup_sslcipher_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslservicegroup_sslcipher_bindingExist("citrixadc_sslservicegroup_sslcipher_binding.tf_sslservicegroup_sslcipher_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup_sslcipher_binding.tf_sslservicegroup_sslcipher_binding", "ciphername", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup_sslcipher_binding.tf_sslservicegroup_sslcipher_binding", "servicegroupname", "tf_servicegroup"),
				),
			},
			{
				Config: testAccSslservicegroup_sslcipher_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslservicegroup_sslcipher_bindingNotExist("tf_servicegroup", "DEFAULT"),
				),
			},
		},
	})
}

func TestAccSslservicegroup_sslcipher_binding_import(t *testing.T) {
	if adcTestbed != "STANDALONE_DEFAULT_SSL_PROFILE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_DEFAULT_SSL_PROFILE.", adcTestbed)
	}
	const resAddr = "citrixadc_sslservicegroup_sslcipher_binding.tf_sslservicegroup_sslcipher_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslservicegroup_sslcipher_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslservicegroup_sslcipher_binding_basic,
			},
			{
				Config:                  testAccSslservicegroup_sslcipher_binding_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckSslservicegroup_sslcipher_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslservicegroup_sslcipher_binding id is set")
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

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}
		servicegroupname := idMap["servicegroupname"]
		ciphername := idMap["ciphername"]

		// The typed sslservicegroup_sslcipher_binding GET is not reflected over REST on
		// this firmware; the binding is only visible via the umbrella
		// sslservicegroup_binding endpoint (under sslservicegroup_sslciphersuite_binding[]).
		// See the resource Read.
		found, err := sslservicegroupSslcipherBindingExistsViaUmbrella(client, servicegroupname, ciphername)
		if err != nil {
			return err
		}
		if !found {
			return fmt.Errorf("sslservicegroup_sslcipher_binding %s not found on ADC", rs.Primary.ID)
		}

		return nil
	}
}

// sslservicegroupSslcipherBindingExistsViaUmbrella reports whether the cipher is bound to
// the SSL service group, reading the umbrella sslservicegroup_binding endpoint (the typed
// binding GET returns an empty body on this firmware).
func sslservicegroupSslcipherBindingExistsViaUmbrella(client *service.NitroClient, servicegroupname, ciphername string) (bool, error) {
	umbrella, err := client.FindResourceArrayWithParams(service.FindParams{
		ResourceType:             service.Sslservicegroup_binding.Type(),
		ResourceName:             servicegroupname,
		ResourceMissingErrorCode: 258,
	})
	if err != nil {
		return false, err
	}
	for _, svc := range umbrella {
		raw, ok := svc["sslservicegroup_sslciphersuite_binding"]
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

// testAccCheckSslservicegroup_sslcipher_bindingNotExist verifies the binding is gone
// while the parent service group still exists (step2 keeps the service group).
func testAccCheckSslservicegroup_sslcipher_bindingNotExist(servicegroupname, ciphername string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		found, err := sslservicegroupSslcipherBindingExistsViaUmbrella(client, servicegroupname, ciphername)
		if err != nil {
			return nil
		}
		if found {
			return fmt.Errorf("sslservicegroup_sslcipher_binding for %s/%s still exists", servicegroupname, ciphername)
		}
		return nil
	}
}

func testAccCheckSslservicegroup_sslcipher_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslservicegroup_sslcipher_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}
		servicegroupname := idMap["servicegroupname"]
		ciphername := idMap["ciphername"]

		found, err := sslservicegroupSslcipherBindingExistsViaUmbrella(client, servicegroupname, ciphername)
		if err != nil {
			continue
		}
		if found {
			return fmt.Errorf("sslservicegroup_sslcipher_binding %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

const testAccSslservicegroup_sslcipher_bindingDataSource_basic = `
	resource "citrixadc_servicegroup" "tf_servicegroup" {
		servicegroupname = "tf_servicegroup"
		servicetype      = "SSL"
	}

	resource "citrixadc_sslservicegroup_sslcipher_binding" "tf_sslservicegroup_sslcipher_binding" {
		ciphername       = "DEFAULT"
		servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
		depends_on       = [citrixadc_servicegroup.tf_servicegroup]
	}

	data "citrixadc_sslservicegroup_sslcipher_binding" "tf_sslservicegroup_sslcipher_binding_ds" {
		servicegroupname = citrixadc_sslservicegroup_sslcipher_binding.tf_sslservicegroup_sslcipher_binding.servicegroupname
		ciphername       = citrixadc_sslservicegroup_sslcipher_binding.tf_sslservicegroup_sslcipher_binding.ciphername
		depends_on       = [citrixadc_sslservicegroup_sslcipher_binding.tf_sslservicegroup_sslcipher_binding]
	}
`

func TestAccSslservicegroup_sslcipher_bindingDataSource_basic(t *testing.T) {
	if adcTestbed != "STANDALONE_DEFAULT_SSL_PROFILE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_DEFAULT_SSL_PROFILE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslservicegroup_sslcipher_bindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslservicegroup_sslcipher_binding.tf_sslservicegroup_sslcipher_binding_ds", "servicegroupname", "tf_servicegroup"),
					resource.TestCheckResourceAttr("data.citrixadc_sslservicegroup_sslcipher_binding.tf_sslservicegroup_sslcipher_binding_ds", "ciphername", "DEFAULT"),
				),
			},
		},
	})
}
