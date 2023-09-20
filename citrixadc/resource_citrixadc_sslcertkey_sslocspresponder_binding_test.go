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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"strings"
	"testing"
)

const testAccSslcertkey_sslocspresponder_binding_basic = `
	resource "citrixadc_sslcertkey" "tf_sslcertkey" {
		certkey            = "tf_sslcertkey"
		cert               = "/nsconfig/ssl/certificate1.crt"
		key                = "/nsconfig/ssl/key1.pem"
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
		cert               = "/nsconfig/ssl/certificate1.crt"
		key                = "/nsconfig/ssl/key1.pem"
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
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSslcertkey_sslocspresponder_bindingDestroy,
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

		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 2)

		certkey := idSlice[0]
		ocspresponder := idSlice[1]

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
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		certkey := idSlice[0]
		ocspresponder := idSlice[1]

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

func testAccCheckSslcertkey_sslocspresponder_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslcertkey_sslocspresponder_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Sslcertkey_sslocspresponder_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslcertkey_sslocspresponder_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
