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

// sslservice_sslcipher_binding binds a cipher group/alias/individual cipher
// (ciphername) to an SSL service (servicename). Participating entity:
//   - citrixadc_service of servicetype SSL (the SSL service is the SSL view of that service)
//
// ciphername here is a built-in cipher alias ("HIGH"). cipheraliasname,
// cipherdefaulton and description are read-only / GET-only attributes (Computed).
//
// CAVEAT: binding a cipher directly to an SSL service may be REJECTED by the ADC if
// the default SSL profile is ENABLED - NITRO returns errorcode 3740 ("Use profile
// command to ..."). On a STANDALONE_NON_DEFAULT_SSL_PROFILE testbed (or with the
// default profile disabled) this test should succeed. This file is generate-only.
//
// Composite ID = ciphername:<v>,servicename:<v>. The exist/destroy checks read the
// binding array for the servicename and match on ciphername.

// step1: create the binding
const testAccSslserviceSslcipherBinding_basic_step1 = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_sslsvc_cipher_lb"
  ipv46       = "10.33.55.35"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_service" "tf_service" {
  servicetype = "SSL"
  name        = "tf_sslsvc_cipher"
  ipaddress   = "10.77.33.24"
  ip          = "10.77.33.24"
  port        = "443"
  lbvserver   = citrixadc_lbvserver.tf_lbvserver.name
}

resource "citrixadc_sslservice_sslcipher_binding" "tf_binding" {
  servicename = citrixadc_service.tf_service.name
  ciphername  = "HIGH"

  depends_on = [citrixadc_service.tf_service]
}
`

// step2: drop the binding (service remains)
const testAccSslserviceSslcipherBinding_basic_step2 = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_sslsvc_cipher_lb"
  ipv46       = "10.33.55.35"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_service" "tf_service" {
  servicetype = "SSL"
  name        = "tf_sslsvc_cipher"
  ipaddress   = "10.77.33.24"
  ip          = "10.77.33.24"
  port        = "443"
  lbvserver   = citrixadc_lbvserver.tf_lbvserver.name
}
`

func TestAccSslserviceSslcipherBinding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslserviceSslcipherBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslserviceSslcipherBinding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslserviceSslcipherBindingExist("citrixadc_sslservice_sslcipher_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_sslservice_sslcipher_binding.tf_binding", "servicename", "tf_sslsvc_cipher"),
					resource.TestCheckResourceAttr("citrixadc_sslservice_sslcipher_binding.tf_binding", "ciphername", "HIGH"),
				),
			},
			{
				// Binding dropped; verify it no longer exists on the ADC.
				Config: testAccSslserviceSslcipherBinding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslserviceSslcipherBindingNotExist("tf_sslsvc_cipher", "HIGH"),
				),
			},
		},
	})
}

func testAccCheckSslserviceSslcipherBindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslservice_sslcipher_binding ID is set")
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
		servicename := idMap["servicename"]
		ciphername := idMap["ciphername"]

		// The typed sslservice_sslcipher_binding GET is not reflected over REST on this
		// firmware; the binding is only visible via the umbrella sslservice_binding
		// endpoint (under sslservice_sslciphersuite_binding[]). See the resource Read.
		found, err := sslserviceSslcipherBindingExistsViaUmbrella(client, servicename, ciphername)
		if err != nil {
			return err
		}
		if !found {
			return fmt.Errorf("sslservice_sslcipher_binding %s not found on ADC", rs.Primary.ID)
		}

		return nil
	}
}

// sslserviceSslcipherBindingExistsViaUmbrella reports whether the cipher is bound to
// the SSL service, reading the umbrella sslservice_binding endpoint (the typed binding
// GET returns an empty body on this firmware).
func sslserviceSslcipherBindingExistsViaUmbrella(client *service.NitroClient, servicename, ciphername string) (bool, error) {
	umbrella, err := client.FindResourceArrayWithParams(service.FindParams{
		ResourceType:             service.Sslservice_binding.Type(),
		ResourceName:             servicename,
		ResourceMissingErrorCode: 258,
	})
	if err != nil {
		return false, err
	}
	for _, svc := range umbrella {
		raw, ok := svc["sslservice_sslciphersuite_binding"]
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

// testAccCheckSslserviceSslcipherBindingNotExist verifies the binding is gone
// while the parent service still exists (step2 keeps the service).
func testAccCheckSslserviceSslcipherBindingNotExist(servicename, ciphername string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		found, err := sslserviceSslcipherBindingExistsViaUmbrella(client, servicename, ciphername)
		if err != nil {
			return nil
		}
		if found {
			return fmt.Errorf("sslservice_sslcipher_binding for %s/%s still exists", servicename, ciphername)
		}
		return nil
	}
}

func testAccCheckSslserviceSslcipherBindingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslservice_sslcipher_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}
		servicename := idMap["servicename"]
		ciphername := idMap["ciphername"]

		found, err := sslserviceSslcipherBindingExistsViaUmbrella(client, servicename, ciphername)
		if err != nil {
			continue
		}
		if found {
			return fmt.Errorf("sslservice_sslcipher_binding %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

const testAccSslserviceSslcipherBindingDataSource_basic = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_sslsvc_cipher_lb"
  ipv46       = "10.33.55.35"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_service" "tf_service" {
  servicetype = "SSL"
  name        = "tf_sslsvc_cipher"
  ipaddress   = "10.77.33.24"
  ip          = "10.77.33.24"
  port        = "443"
  lbvserver   = citrixadc_lbvserver.tf_lbvserver.name
}

resource "citrixadc_sslservice_sslcipher_binding" "tf_binding" {
  servicename = citrixadc_service.tf_service.name
  ciphername  = "HIGH"

  depends_on = [citrixadc_service.tf_service]
}

data "citrixadc_sslservice_sslcipher_binding" "tf_binding" {
  servicename = citrixadc_sslservice_sslcipher_binding.tf_binding.servicename
  ciphername  = citrixadc_sslservice_sslcipher_binding.tf_binding.ciphername
  depends_on  = [citrixadc_sslservice_sslcipher_binding.tf_binding]
}
`

func TestAccSslserviceSslcipherBindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslserviceSslcipherBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslserviceSslcipherBindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslservice_sslcipher_binding.tf_binding", "servicename", "tf_sslsvc_cipher"),
					resource.TestCheckResourceAttr("data.citrixadc_sslservice_sslcipher_binding.tf_binding", "ciphername", "HIGH"),
				),
			},
		},
	})
}
