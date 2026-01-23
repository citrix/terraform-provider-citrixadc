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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strings"
	"testing"
)

const testAccCmppolicylabel_cmppolicy_binding_basic = `

resource "citrixadc_cmppolicylabel_cmppolicy_binding" "tf_cmppolicylabel_cmppolicy_binding" {
	policyname = citrixadc_cmppolicy.tf_cmppolicy.name
	labelname  = citrixadc_cmppolicylabel.tf_cmppolicylabel.labelname
	priority   = 100
	}

  resource "citrixadc_cmppolicylabel" "tf_cmppolicylabel" {
	labelname = "my_cmppolicy_label"
	type      = "RES"
	}
  resource "citrixadc_cmppolicy" "tf_cmppolicy" {
	name      = "tf_cmppolicy"
	rule      = "HTTP.RES.HEADER(\"Content-Type\").CONTAINS(\"text\")"
	resaction = "COMPRESS"
	}
`

const testAccCmppolicylabel_cmppolicy_binding_basic_step2 = `

resource "citrixadc_cmppolicylabel" "tf_cmppolicylabel" {
	labelname = "my_cmppolicy_label"
	type      = "RES"
	}
resource "citrixadc_cmppolicy" "tf_cmppolicy" {
	name      = "tf_cmppolicy"
	rule      = "HTTP.RES.HEADER(\"Content-Type\").CONTAINS(\"text\")"
	resaction = "COMPRESS"
	}
`

func TestAccCmppolicylabel_cmppolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCmppolicylabel_cmppolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCmppolicylabel_cmppolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCmppolicylabel_cmppolicy_bindingExist("citrixadc_cmppolicylabel_cmppolicy_binding.tf_cmppolicylabel_cmppolicy_binding", nil),
				),
			},
			{
				Config: testAccCmppolicylabel_cmppolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCmppolicylabel_cmppolicy_bindingNotExist("citrixadc_cmppolicylabel_cmppolicy_binding.tf_cmppolicylabel_cmppolicy_binding", "my_cmppolicy_label,tf_cmppolicy"),
				),
			},
		},
	})
}

func testAccCheckCmppolicylabel_cmppolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cmppolicylabel_cmppolicy_binding id is set")
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

		idSlice := strings.SplitN(bindingId, ",", 2)

		labelname := idSlice[0]
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "cmppolicylabel_cmppolicy_binding",
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
			return fmt.Errorf("cmppolicylabel_cmppolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckCmppolicylabel_cmppolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		labelname := idSlice[0]
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "cmppolicylabel_cmppolicy_binding",
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
			return fmt.Errorf("cmppolicylabel_cmppolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckCmppolicylabel_cmppolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_cmppolicylabel_cmppolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Cmppolicylabel_cmppolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("cmppolicylabel_cmppolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
