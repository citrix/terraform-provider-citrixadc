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

const testAccResponderpolicylabel_responderpolicy_binding_basic = `

resource "citrixadc_responderpolicylabel_responderpolicy_binding" "tf_responderpolicylabel_responderpolicy_binding" {
	labelname = citrixadc_responderpolicylabel.tf_responderpolicylabel.labelname
	policyname = citrixadc_responderpolicy.tf_responderpolicy.name
	priority = 5  
	gotopriorityexpression = "END"
	invoke = "false"
}

resource "citrixadc_responderpolicylabel" "tf_responderpolicylabel" {
	labelname = "tf_responderpolicylabel"
	policylabeltype = "HTTP"
}

resource "citrixadc_responderpolicy" "tf_responderpolicy" {
	name    = "tf_responderpolicy"
	action = "NOOP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"nosuchthing\")"
}
`

const testAccResponderpolicylabel_responderpolicy_binding_basic_step2 = `

	resource "citrixadc_responderpolicylabel" "tf_responderpolicylabel" {
		labelname = "tf_responderpolicylabel"
		policylabeltype = "HTTP"
	}

	resource "citrixadc_responderpolicy" "tf_responderpolicy" {
		name    = "tf_responderpolicy"
		action = "NOOP"
		rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"nosuchthing\")"
	}
`

func TestAccResponderpolicylabel_responderpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckResponderpolicylabel_responderpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResponderpolicylabel_responderpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResponderpolicylabel_responderpolicy_bindingExist("citrixadc_responderpolicylabel_responderpolicy_binding.tf_responderpolicylabel_responderpolicy_binding", nil),
				),
			},
			{
				Config: testAccResponderpolicylabel_responderpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResponderpolicylabel_responderpolicy_bindingNotExist("citrixadc_responderpolicylabel_responderpolicy_binding.tf_responderpolicylabel_responderpolicy_binding", "tf_responderpolicylabel,tf_responderpolicy"),
				),
			},
		},
	})
}

func testAccCheckResponderpolicylabel_responderpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No responderpolicylabel_responderpolicy_binding id is set")
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
			ResourceType:             "responderpolicylabel_responderpolicy_binding",
			ResourceName:             labelname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching policyname
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("responderpolicylabel_responderpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckResponderpolicylabel_responderpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		labelname := idSlice[0]
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "responderpolicylabel_responderpolicy_binding",
			ResourceName:             labelname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching policyname
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("responderpolicylabel_responderpolicy_binding %s was found, but is should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckResponderpolicylabel_responderpolicy_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_responderpolicylabel_responderpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Responderpolicylabel_responderpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("responderpolicylabel_responderpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
