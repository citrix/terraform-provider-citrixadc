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
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
	"net/url"
)

const testAccVpnglobal_sslcertkey_binding_basic = `

resource "citrixadc_sslcertkey" "foo" {
	certkey = "sample_ssl_cert"
	cert    = "/var/tmp/certificate1.crt"
	key     = "/var/tmp/key1.pem"
  }
  resource "citrixadc_vpnglobal_sslcertkey_binding" "tf_vpnglobal_sslcertkey_binding" {
	certkeyname = citrixadc_sslcertkey.foo.certkey
  }
`

const testAccVpnglobal_sslcertkey_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_sslcertkey" "foo" {
		certkey = "sample_ssl_cert"
		cert    = "/var/tmp/certificate1.crt"
		key     = "/var/tmp/key1.pem"
	  }
	  
`

func TestAccVpnglobal_sslcertkey_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { PreCheckSslceriKey(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVpnglobal_sslcertkey_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVpnglobal_sslcertkey_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnglobal_sslcertkey_bindingExist("citrixadc_vpnglobal_sslcertkey_binding.tf_vpnglobal_sslcertkey_binding", nil),
				),
			},
			resource.TestStep{
				Config: testAccVpnglobal_sslcertkey_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnglobal_sslcertkey_bindingNotExist("citrixadc_vpnglobal_sslcertkey_binding.tf_vpnglobal_sslcertkey_binding", "certkeyname"),
				),
			},
		},
	})
}

func testAccCheckVpnglobal_sslcertkey_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnglobal_sslcertkey_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client := testAccProvider.Meta().(*NetScalerNitroClient).client
		certkeyname , _ := url.QueryUnescape(rs.Primary.ID)

		findParams := service.FindParams{
			ResourceType:             "vpnglobal_sslcertkey_binding",
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["certkeyname"].(string) == certkeyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("vpnglobal_sslcertkey_binding %s not found", n)
		}

		return nil
	}
}

func PreCheckSslceriKey(t *testing.T) {
	testAccPreCheck(t)

	uploads := []string{
		"certificate1.crt",
		"key1.pem",
	}

	c, err := testHelperInstantiateClient("", "", "", false)
	if err != nil {
		t.Fatalf("Failed to instantiate client. %v\n", err)
	}

	//c := testAccProvider.Meta().(*NetScalerNitroClient)
	for _, filename := range uploads {
		err := uploadTestdataFile(c, t, filename, "/var/tmp")
		if err != nil {
			t.Errorf(err.Error())
		}
	}
}


func testAccCheckVpnglobal_sslcertkey_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client
		certkeyname := id

		findParams := service.FindParams{
			ResourceType:             "vpnglobal_sslcertkey_binding",
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["certkeyname"].(string) == certkeyname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("vpnglobal_sslcertkey_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckVpnglobal_sslcertkey_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnglobal_sslcertkey_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("Vpnglobal_sslcertkey_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnglobal_sslcertkey_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
