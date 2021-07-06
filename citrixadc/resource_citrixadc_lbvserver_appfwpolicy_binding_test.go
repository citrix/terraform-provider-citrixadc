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

const testAccLbvserver_appfwpolicy_binding_basic = `
	resource citrixadc_lbvserver_appfwpolicy_binding demo_binding {
		name = citrixadc_lbvserver.demo_lb.name
		priority = 100
		bindpoint = "REQUEST"
		policyname  = citrixadc_appfwpolicy.demo_appfwpolicy.name
		labelname = citrixadc_lbvserver.demo_lb.name
		gotopriorityexpression = "END"
		invoke = true
		labeltype = "reqvserver"
	}

	resource citrixadc_lbvserver demo_lb {
	name        = "demo_lb"
	ipv46       = "1.1.1.1"
	port        = "80"
	servicetype = "HTTP"
	}

	resource citrixadc_appfwprofile demo_appfwprofile {
		name = "demo_appfwprofile"
		bufferoverflowaction = ["none"]
		contenttypeaction = ["none"]
		cookieconsistencyaction = ["none"]
		creditcard = ["none"]
		creditcardaction = ["none"]
		crosssitescriptingaction = ["none"]
		csrftagaction = ["none"]
		denyurlaction = ["none"]
		dynamiclearning = ["none"]
		fieldconsistencyaction = ["none"]
		fieldformataction = ["none"]
		fileuploadtypesaction = ["none"]
		inspectcontenttypes = ["none"]
		jsondosaction = ["none"]
		jsonsqlinjectionaction = ["none"]
		jsonxssaction = ["none"]
		multipleheaderaction = ["none"]
		sqlinjectionaction = ["none"]
		starturlaction = ["none"]
		type = ["HTML"]
		xmlattachmentaction = ["none"]
		xmldosaction = ["none"]
		xmlformataction = ["none"]
		xmlsoapfaultaction = ["none"]
		xmlsqlinjectionaction = ["none"]
		xmlvalidationaction = ["none"]
		xmlwsiaction = ["none"]
		xmlxssaction = ["none"]
	}

	resource citrixadc_appfwpolicy demo_appfwpolicy {
		name = "demo_appfwpolicy"
		profilename = citrixadc_appfwprofile.demo_appfwprofile.name
		rule = "true"
	}

`

func TestAccLbvserver_appfwpolicy_binding_basic(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLbvserver_appfwpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccLbvserver_appfwpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_appfwpolicy_bindingExist("citrixadc_lbvserver_appfwpolicy_binding.demo_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_lbvserver_appfwpolicy_binding.demo_binding", "name", "demo_lb"),
					resource.TestCheckResourceAttr("citrixadc_lbvserver_appfwpolicy_binding.demo_binding", "priority", "100"),
					resource.TestCheckResourceAttr("citrixadc_lbvserver_appfwpolicy_binding.demo_binding", "bindpoint", "REQUEST"),
					resource.TestCheckResourceAttr("citrixadc_lbvserver_appfwpolicy_binding.demo_binding", "policyname", "demo_appfwpolicy"),
					resource.TestCheckResourceAttr("citrixadc_lbvserver_appfwpolicy_binding.demo_binding", "gotopriorityexpression", "END"),
					resource.TestCheckResourceAttr("citrixadc_lbvserver_appfwpolicy_binding.demo_binding", "labeltype", "reqvserver"),
					resource.TestCheckResourceAttr("citrixadc_lbvserver_appfwpolicy_binding.demo_binding", "invoke", "true"),
				),
			},
		},
	})
}

func testAccCheckLbvserver_appfwpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbvserver_appfwpolicy_binding name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		bindingId := rs.Primary.ID
		idSlice := strings.SplitN(bindingId, ",", 2)
		lbvserverName := idSlice[0]
		appfwPolicyName := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             service.Lbvserver_appfwpolicy_binding.Type(),
			ResourceName:             lbvserverName,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := nsClient.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right policy name
		foundIndex := -1
		for i, v := range dataArr {
			if v["policyname"].(string) == appfwPolicyName {
				foundIndex = i
				break
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			return fmt.Errorf("Cannot find lbvserver_appfwpolicy_binding ID %v", bindingId)
		}

		return nil
	}
}

func testAccCheckLbvserver_appfwpolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbvserver_appfwpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Lbvserver_appfwpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lbvserver_appfwpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
