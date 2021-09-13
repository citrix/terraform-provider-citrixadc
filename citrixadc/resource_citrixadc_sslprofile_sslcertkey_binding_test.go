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
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const testAccSslprofile_sslcertkey_binding_basic_step1 = `
resource "citrixadc_sslprofile" "tfUnit_sslprofile-hello" {
	name = "tfUnit_sslprofile-hello"

	// ecccurvebindings is REQUIRED attribute.
	// The default ecccurvebindings will be DELETED and only the explicitly given ecccurvebindings will be retained
	// To unbind all the ecccurvebindings, an empty list [] is to be assinged to ecccurvebindings attribute

	ecccurvebindings = ["P_256"]
	sslinterception = "ENABLED"
  
}
  
resource "citrixadc_sslcertkey" "tf_sslcertkey" {
	  certkey = "tf_sslcertkey"
	  cert = "/nsconfig/ssl/ns-root.cert"
	  key = "/nsconfig/ssl/ns-root.key"
}
	
resource "citrixadc_sslprofile_sslcertkey_binding" "demo_sslprofile_sslcertkey_binding" {
	  name = citrixadc_sslprofile.tfUnit_sslprofile-hello.name
	  sslicacertkey = citrixadc_sslcertkey.tf_sslcertkey.certkey 
}
  `

const testAccSslprofile_sslcertkey_binding_basic_step2 = `
resource "citrixadc_sslprofile" "tfUnit_sslprofile-hello" {
	name = "tfUnit_sslprofile-hello"

	// ecccurvebindings is REQUIRED attribute.
	// The default ecccurvebindings will be DELETED and only the explicitly given ecccurvebindings will be retained
	// To unbind all the ecccurvebindings, an empty list [] is to be assinged to ecccurvebindings attribute
	
	ecccurvebindings = ["P_256"]
	sslinterception = "ENABLED"
  
}
  
resource "citrixadc_sslcertkey" "tf_sslcertkey" {
	  certkey = "tf_sslcertkey"
	  cert = "/nsconfig/ssl/ns-root.cert"
	  key = "/nsconfig/ssl/ns-root.key"
}
`

func TestAccSslprofile_sslcertkey_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSslprofile_sslcertkey_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSslprofile_sslcertkey_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofile_sslcertkey_bindingExist("citrixadc_sslprofile_sslcertkey_binding.demo_sslprofile_sslcertkey_binding", nil),
				),
			},
			resource.TestStep{
				Config: testAccSslprofile_sslcertkey_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofile_sslcertkey_bindingNotExist("citrixadc_sslprofile_sslcertkey_binding.demo_sslprofile_sslcertkey_binding", "tfUnit_sslprofile-hello,tf_sslcertkey"),
				),
			},
		},
	})
}

func testAccCheckSslprofile_sslcertkey_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslprofile_sslcertkey_binding id is set")
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

		name := idSlice[0]
		sslicacertkey := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "sslprofile_sslcertkey_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 3248,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right monitor name
		found := false
		for _, v := range dataArr {
			if v["sslicacertkey"].(string) == sslicacertkey {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("sslprofile_sslcertkey_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslprofile_sslcertkey_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}

		idSlice := strings.SplitN(id, ",", 2)

		name := idSlice[0]
		sslicacertkey := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "sslprofile_sslcertkey_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 3248,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right monitor name
		found := false
		for _, v := range dataArr {
			if v["sslicacertkey"].(string) == sslicacertkey {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("sslprofile_sslcertkey_binding %s not deleted", n)
		}

		return nil
	}
}

func testAccCheckSslprofile_sslcertkey_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslprofile_sslcertkey_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("sslprofile_sslcertkey_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslprofile_sslcertkey_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
