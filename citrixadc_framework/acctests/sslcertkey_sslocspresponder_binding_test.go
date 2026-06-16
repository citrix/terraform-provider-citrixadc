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

const testAccSslcertkey_sslocspresponder_binding_basic = `
	resource "citrixadc_sslcertkey" "tf_sslcertkey" {
		certkey            = "tf_sslcertkey"
		cert               = "/nsconfig/ssl/rootcert2.cert"
		key                = "/nsconfig/ssl/rootcert2.key"
		notificationperiod = 40
		expirymonitor      = "ENABLED"
	}
	resource "citrixadc_sslocspresponder" "tf_sslocspresponder" {
		name = "tf_sslocspresponder"
		url  = "http://www.google.com"
	}
	resource "citrixadc_sslcertkey_sslocspresponder_binding" "tf_binding" {
		certkey 	  = citrixadc_sslcertkey.tf_sslcertkey.certkey
		ocspresponder = citrixadc_sslocspresponder.tf_sslocspresponder.name
		priority      = 90
	}
`

const testAccSslcertkey_sslocspresponder_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_sslcertkey" "tf_sslcertkey" {
		certkey            = "tf_sslcertkey"
		cert               = "/nsconfig/ssl/rootcert2.cert"
		key                = "/nsconfig/ssl/rootcert2.key"
		notificationperiod = 40
		expirymonitor      = "ENABLED"
	}
	resource "citrixadc_sslocspresponder" "tf_sslocspresponder" {
		name = "tf_sslocspresponder"
		url  = "http://www.google.com"
	}
`

func TestAccSslcertkey_sslocspresponder_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslcertkey_sslocspresponder_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcertkey_sslocspresponder_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertkey_sslocspresponder_bindingExist("citrixadc_sslcertkey_sslocspresponder_binding.tf_binding", nil),
				),
			},
			{
				Config: testAccSslcertkey_sslocspresponder_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertkey_sslocspresponder_bindingNotExist("citrixadc_sslcertkey_sslocspresponder_binding.tf_binding", "tf_sslcertkey,tf_sslocspresponder"),
				),
			},
		},
	})
}

func testAccCheckSslcertkey_sslocspresponder_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslcertkey_sslocspresponder_binding id is set")
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

		bindingId := rs.Primary.ID

		idMap, _, err := utils.ParseIdString(bindingId, []string{"certkey", "ocspresponder"}, nil)
		if err != nil {
			return err
		}

		certkey := idMap["certkey"]
		ocspresponder := idMap["ocspresponder"]

		findParams := service.FindParams{
			ResourceType:             "sslcertkey_sslocspresponder_binding",
			ResourceName:             certkey,
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
			if v["ocspresponder"].(string) == ocspresponder {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("sslcertkey_sslocspresponder_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslcertkey_sslocspresponder_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(id, []string{"certkey", "ocspresponder"}, nil)
		if err != nil {
			return err
		}

		certkey := idMap["certkey"]
		ocspresponder := idMap["ocspresponder"]

		findParams := service.FindParams{
			ResourceType:             "sslcertkey_sslocspresponder_binding",
			ResourceName:             certkey,
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
			if v["ocspresponder"].(string) == ocspresponder {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("sslcertkey_sslocspresponder_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

const testAccSslcertkey_sslocspresponder_bindingDataSource_basic = `
	resource "citrixadc_sslcertkey" "tf_sslcertkey" {
		certkey = "tf_sslcertkey"
		cert    = "/nsconfig/ssl/rootcert2.cert"
		key     = "/nsconfig/ssl/rootcert2.key"
	}
	resource "citrixadc_sslocspresponder" "tf_sslocspresponder" {
		name = "tf_sslocspresponder"
		url  = "http://www.google.com"
	}
	resource "citrixadc_sslcertkey_sslocspresponder_binding" "tf_binding" {
		certkey 	  = citrixadc_sslcertkey.tf_sslcertkey.certkey
		ocspresponder = citrixadc_sslocspresponder.tf_sslocspresponder.name
		priority      = 90
	}

	data "citrixadc_sslcertkey_sslocspresponder_binding" "tf_binding" {
		certkey       = citrixadc_sslcertkey_sslocspresponder_binding.tf_binding.certkey
		ocspresponder = citrixadc_sslcertkey_sslocspresponder_binding.tf_binding.ocspresponder
		depends_on    = [citrixadc_sslcertkey_sslocspresponder_binding.tf_binding]
	}
`

func testAccCheckSslcertkey_sslocspresponder_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslcertkey_sslocspresponder_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"certkey", "ocspresponder"}, nil)
		if err != nil {
			return err
		}
		certkey := idMap["certkey"]

		_, err = client.FindResource(service.Sslcertkey_sslocspresponder_binding.Type(), certkey)
		if err == nil {
			return fmt.Errorf("sslcertkey_sslocspresponder_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccSslcertkey_sslocspresponder_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeyPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcertkey_sslocspresponder_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslcertkey_sslocspresponder_binding.tf_binding", "certkey", "tf_sslcertkey"),
					resource.TestCheckResourceAttr("data.citrixadc_sslcertkey_sslocspresponder_binding.tf_binding", "ocspresponder", "tf_sslocspresponder"),
					resource.TestCheckResourceAttr("data.citrixadc_sslcertkey_sslocspresponder_binding.tf_binding", "priority", "90"),
				),
			},
		},
	})
}
