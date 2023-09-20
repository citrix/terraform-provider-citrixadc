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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const testAccSslpolicylabel_sslpolicy_binding_basic_step1 = `
resource "citrixadc_sslaction" "certinsertact" {
	name       = "certinsertact"
	clientcert = "ENABLED"
	certheader = "CERT"
	}
	
	resource "citrixadc_sslpolicy" "certinsert_pol" {
	name   = "certinsert_pol"
	rule   = "false"
	action = citrixadc_sslaction.certinsertact.name
	}
	
	resource "citrixadc_sslpolicylabel" "ssl_pol_label" {
		labelname = "ssl_pol_label"
		type = "DATA"	
	}
	
	resource "citrixadc_sslpolicylabel_sslpolicy_binding" "demo_sslpolicylabel_sslpolicy_binding" {
		gotopriorityexpression = "END"
		invoke = true
		labelname = citrixadc_sslpolicylabel.ssl_pol_label.labelname
		labeltype = "policylabel"
		policyname = citrixadc_sslpolicy.certinsert_pol.name
		priority = 56       
		invokelabelname = "ssl_pol_label"
	}
	`

const testAccSslpolicylabel_sslpolicy_binding_basic_step2 = `
resource "citrixadc_sslaction" "certinsertact" {
	name       = "certinsertact"
	clientcert = "ENABLED"
	certheader = "CERT"
	}
	
	resource "citrixadc_sslpolicy" "certinsert_pol" {
	name   = "certinsert_pol"
	rule   = "false"
	action = citrixadc_sslaction.certinsertact.name
	}
	
	resource "citrixadc_sslpolicylabel" "ssl_pol_label" {
		labelname = "ssl_pol_label"
		type = "DATA"	
	}
`

func TestAccSslpolicylabel_sslpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSslpolicylabel_sslpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslpolicylabel_sslpolicy_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslpolicylabel_sslpolicy_bindingExist("citrixadc_sslpolicylabel_sslpolicy_binding.demo_sslpolicylabel_sslpolicy_binding", nil),
				),
			},
			{
				Config: testAccSslpolicylabel_sslpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslpolicylabel_sslpolicy_bindingNotExist("citrixadc_sslpolicylabel_sslpolicy_binding.demo_sslpolicylabel_sslpolicy_binding", "ssl_pol_label,certinsert_pol"),
				),
			},
		},
	})
}

func testAccCheckSslpolicylabel_sslpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslpolicylabel_sslpolicy_binding id is set")
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

		labelname := idSlice[0]
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "sslpolicylabel_sslpolicy_binding",
			ResourceName:             labelname,
			ResourceMissingErrorCode: 3248,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right monitor labelname
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("sslpolicylabel_sslpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslpolicylabel_sslpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}

		idSlice := strings.SplitN(id, ",", 2)

		labelname := idSlice[0]
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "sslpolicylabel_sslpolicy_binding",
			ResourceName:             labelname,
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
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("sslpolicylabel_sslpolicy_binding %s not deleted", n)
		}

		return nil
	}
}

func testAccCheckSslpolicylabel_sslpolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslpolicylabel_sslpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No labelname is set")
		}

		_, err := nsClient.FindResource(service.Sslpolicylabel_sslpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslpolicylabel_sslpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
