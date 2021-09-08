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
	"strings"
	"testing"
)

const testAccSslvserver_sslciphersuite_binding_basic = `
	resource "citrixadc_sslvserver_sslciphersuite_binding" "tf_sslvserver_sslciphersuite_binding" {
		ciphername = "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256"
		vservername = citrixadc_lbvserver.tf_sslvserver.name
	}

	resource "citrixadc_sslvserver_sslciphersuite_binding" "tf_sslvserver_sslciphersuite_binding2" {
		ciphername = "TLS1.3-CHACHA20-POLY1305-SHA256"
		vservername = citrixadc_lbvserver.tf_sslvserver.name
	}

	resource "citrixadc_lbvserver" "tf_sslvserver" {
		name = "tf_sslvserver"
		servicetype = "SSL"
		ipv46 = "5.5.5.5"
		port = 80
	}
`

const testAccSslvserver_sslciphersuite_binding_basic_step2 = `
	resource "citrixadc_lbvserver" "tf_sslvserver" {
		name = "tf_sslvserver"
		servicetype = "SSL"
		ipv46 = "5.5.5.5"
		port = 80
	}
`

func TestAccSslvserver_sslciphersuite_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSslvserver_sslciphersuite_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSslvserver_sslciphersuite_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_sslciphersuite_bindingExist("citrixadc_sslvserver_sslciphersuite_binding.tf_sslvserver_sslciphersuite_binding", nil),
				),
			},
			resource.TestStep{
				Config: testAccSslvserver_sslciphersuite_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_sslciphersuite_bindingExist("citrixadc_sslvserver_sslciphersuite_binding.tf_sslvserver_sslciphersuite_binding2", nil),
				),
			},
			resource.TestStep{
				Config: testAccSslvserver_sslciphersuite_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_sslciphersuite_bindingNotExist("citrixadc_sslvserver_sslciphersuite_binding.tf_sslvserver_sslciphersuite_binding", "tf_sslvserver,TLS1.2-ECDHE-RSA-AES128-GCM-SHA256"),
				),
			},
			resource.TestStep{
				Config: testAccSslvserver_sslciphersuite_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_sslciphersuite_bindingNotExist("citrixadc_sslvserver_sslciphersuite_binding.tf_sslvserver_sslciphersuite_binding2", "tf_sslvserver,TLS1.3-CHACHA20-POLY1305-SHA256"),
				),
			},
		},
	})
}

func testAccCheckSslvserver_sslciphersuite_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslvserver_sslciphersuite_binding id is set")
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

		vservername := idSlice[0]
		ciphername := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "sslvserver_sslciphersuite_binding",
			ResourceName:             vservername,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right monitor name
		found := false
		for _, v := range dataArr {
			if v["ciphername"].(string) == ciphername {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("sslvserver_sslciphersuite_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslvserver_sslciphersuite_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		name := idSlice[0]
		ciphername := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "sslvserver_sslciphersuite_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right certkey name
		found := false
		for _, v := range dataArr {
			if v["ciphername"].(string) == ciphername {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("sslvserver_sslciphersuite_binding %s was found, but is should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckSslvserver_sslciphersuite_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslvserver_sslciphersuite_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Sslvserver_sslciphersuite_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslvserver_sslciphersuite_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
