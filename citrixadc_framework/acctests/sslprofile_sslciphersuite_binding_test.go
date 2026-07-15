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

// Participating entity (lifted from sslprofile_test.go):
//   - citrixadc_sslprofile (the SSL profile the ciphersuite is bound to)
// ciphername uses the built-in cipher alias "HIGH" (the same alias the inline
// sslprofile cipherbindings test in sslprofile_test.go uses successfully).
// Composite ID = name,ciphername.

const testAccSslprofileSslciphersuiteBinding_basic_step1 = `
	resource "citrixadc_sslprofile" "tf_sslprofile" {
		name             = "tf_sslprofile_cs"
		ecccurvebindings = []
	}

	resource "citrixadc_sslprofile_sslciphersuite_binding" "tf_sslprofile_sslciphersuite_binding" {
		name           = citrixadc_sslprofile.tf_sslprofile.name
		ciphername     = "HIGH"
		cipherpriority = 10

		depends_on = [citrixadc_sslprofile.tf_sslprofile]
	}
`

// step2 drops the binding (keeps the profile), verifying clean unbind.
const testAccSslprofileSslciphersuiteBinding_basic_step2 = `
	resource "citrixadc_sslprofile" "tf_sslprofile" {
		name             = "tf_sslprofile_cs"
		ecccurvebindings = []
	}
`

func TestAccSslprofileSslciphersuiteBinding_basic(t *testing.T) {
	if adcTestbed != "STANDALONE_DEFAULT_SSL_PROFILE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_DEFAULT_SSL_PROFILE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslprofileSslciphersuiteBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslprofileSslciphersuiteBinding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileSslciphersuiteBindingExist("citrixadc_sslprofile_sslciphersuite_binding.tf_sslprofile_sslciphersuite_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile_sslciphersuite_binding.tf_sslprofile_sslciphersuite_binding", "name", "tf_sslprofile_cs"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile_sslciphersuite_binding.tf_sslprofile_sslciphersuite_binding", "ciphername", "HIGH"),
					resource.TestCheckResourceAttr("citrixadc_sslprofile_sslciphersuite_binding.tf_sslprofile_sslciphersuite_binding", "cipherpriority", "10"),
				),
			},
			{
				Config: testAccSslprofileSslciphersuiteBinding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofileExist("citrixadc_sslprofile.tf_sslprofile", nil),
				),
			},
		},
	})
}

func TestAccSslprofileSslciphersuiteBinding_import(t *testing.T) {
	if adcTestbed != "STANDALONE_DEFAULT_SSL_PROFILE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_DEFAULT_SSL_PROFILE.", adcTestbed)
	}
	const resAddr = "citrixadc_sslprofile_sslciphersuite_binding.tf_sslprofile_sslciphersuite_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslprofileSslciphersuiteBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslprofileSslciphersuiteBinding_basic_step1,
			},
			{
				Config:                  testAccSslprofileSslciphersuiteBinding_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckSslprofileSslciphersuiteBindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslprofile_sslciphersuite_binding id is set")
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

		// Composite ID = name,ciphername (key:UrlEncode(value) form)
		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"name", "ciphername"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		// The typed sslprofile_sslciphersuite_binding GET is not reflected over REST on
		// this firmware - the by-name endpoint (even with ?filter=ciphername:<v>) returns
		// an empty "Done" body even when the cipher is bound. Try the typed GET first
		// (for firmwares that do reflect it); on empty, confirm via the umbrella
		// sslprofile_binding endpoint, which DOES reflect bound ciphers under
		// sslprofile_sslcipher_binding[] (keyed by cipheraliasname / ciphername).
		findParams := service.FindParams{
			ResourceType:             service.Sslprofile_sslciphersuite_binding.Type(),
			ResourceName:             idMap["name"],
			FilterMap:                map[string]string{"ciphername": idMap["ciphername"]},
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}
		for _, v := range dataArr {
			if cn, ok := v["ciphername"].(string); ok && cn == idMap["ciphername"] {
				return nil
			}
		}

		if sslprofileCipherBoundViaUmbrella(client, idMap["name"], idMap["ciphername"]) {
			return nil
		}

		return fmt.Errorf("sslprofile_sslciphersuite_binding %s not found", n)
	}
}

// sslprofileCipherBoundViaUmbrella reports whether the given cipher is bound to the
// profile, read via the umbrella sslprofile_binding endpoint (the typed binding GET is
// not reflected over REST on this firmware). Matched on cipheraliasname (aliases) or
// ciphername (individual ciphers).
func sslprofileCipherBoundViaUmbrella(client *service.NitroClient, name, ciphername string) bool {
	umbrella, err := client.FindResourceArrayWithParams(service.FindParams{
		ResourceType:             service.Sslprofile_binding.Type(),
		ResourceName:             name,
		ResourceMissingErrorCode: 258,
	})
	if err != nil {
		return false
	}
	for _, prof := range umbrella {
		rows, ok := prof["sslprofile_sslcipher_binding"].([]interface{})
		if !ok {
			continue
		}
		for _, cb := range rows {
			m, ok := cb.(map[string]interface{})
			if !ok {
				continue
			}
			if v, ok := m["cipheraliasname"].(string); ok && v == ciphername {
				return true
			}
			if v, ok := m["ciphername"].(string); ok && v == ciphername {
				return true
			}
		}
	}
	return false
}

func testAccCheckSslprofileSslciphersuiteBindingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslprofile_sslciphersuite_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No id is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"name", "ciphername"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		// The typed binding GET is not reflected over REST on this firmware, so confirm
		// removal via the umbrella sslprofile_binding endpoint (same path the Exist check
		// uses). If the cipher is still bound there, the destroy did not take effect.
		if sslprofileCipherBoundViaUmbrella(client, idMap["name"], idMap["ciphername"]) {
			return fmt.Errorf("sslprofile_sslciphersuite_binding %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

const testAccSslprofileSslciphersuiteBindingDataSource_basic = `
	resource "citrixadc_sslprofile" "tf_sslprofile" {
		name             = "tf_sslprofile_cs"
		ecccurvebindings = []
	}

	resource "citrixadc_sslprofile_sslciphersuite_binding" "tf_sslprofile_sslciphersuite_binding" {
		name           = citrixadc_sslprofile.tf_sslprofile.name
		ciphername     = "HIGH"
		cipherpriority = 10

		depends_on = [citrixadc_sslprofile.tf_sslprofile]
	}

	data "citrixadc_sslprofile_sslciphersuite_binding" "tf_sslprofile_sslciphersuite_binding" {
		name       = citrixadc_sslprofile_sslciphersuite_binding.tf_sslprofile_sslciphersuite_binding.name
		ciphername = "HIGH"

		depends_on = [citrixadc_sslprofile_sslciphersuite_binding.tf_sslprofile_sslciphersuite_binding]
	}
`

func TestAccSslprofileSslciphersuiteBindingDataSource_basic(t *testing.T) {
	if adcTestbed != "STANDALONE_DEFAULT_SSL_PROFILE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_DEFAULT_SSL_PROFILE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslprofileSslciphersuiteBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslprofileSslciphersuiteBindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslprofile_sslciphersuite_binding.tf_sslprofile_sslciphersuite_binding", "name", "tf_sslprofile_cs"),
					resource.TestCheckResourceAttr("data.citrixadc_sslprofile_sslciphersuite_binding.tf_sslprofile_sslciphersuite_binding", "ciphername", "HIGH"),
				),
			},
		},
	})
}
