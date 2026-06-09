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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// lbglobal_lbpolicy_binding is a KEYLESS global binding (lbglobal singleton has no
// parent name). Composite ID = policyname:<v>,type:<v>. Read/Exist use
// FindResourceArrayWithParams with the `type` arg filter, then match on policyname.
// All attributes are RequiresReplace (no update); Create=PUT (UpdateUnnamedResource).
// Delete uses policyname + type args.

const testAccLbglobal_lbpolicy_binding_basic_step1 = `

resource "citrixadc_lbaction" "tf_act" {
	name  = "tf_act"
	type  = "SELECTIONORDER"
	value = [1]
}

resource "citrixadc_lbpolicy" "tf_pol" {
	name   = "tf_pol"
	rule   = "true"
	action = citrixadc_lbaction.tf_act.name
}

resource "citrixadc_lbglobal_lbpolicy_binding" "tf_lbglobal_lbpolicy_binding" {
	policyname = citrixadc_lbpolicy.tf_pol.name
	priority   = 100
	type       = "REQ_DEFAULT"
	depends_on = [citrixadc_lbpolicy.tf_pol]
}
`

const testAccLbglobal_lbpolicy_binding_basic_step2 = `
	# Keep the participating entities but drop the actual binding to verify proper deletion

	resource "citrixadc_lbaction" "tf_act" {
		name  = "tf_act"
		type  = "SELECTIONORDER"
		value = [1]
	}

	resource "citrixadc_lbpolicy" "tf_pol" {
		name   = "tf_pol"
		rule   = "true"
		action = citrixadc_lbaction.tf_act.name
	}
`

func TestAccLbglobal_lbpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbglobal_lbpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbglobal_lbpolicy_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbglobal_lbpolicy_bindingExist("citrixadc_lbglobal_lbpolicy_binding.tf_lbglobal_lbpolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_lbglobal_lbpolicy_binding.tf_lbglobal_lbpolicy_binding", "policyname", "tf_pol"),
					resource.TestCheckResourceAttr("citrixadc_lbglobal_lbpolicy_binding.tf_lbglobal_lbpolicy_binding", "priority", "100"),
					resource.TestCheckResourceAttr("citrixadc_lbglobal_lbpolicy_binding.tf_lbglobal_lbpolicy_binding", "type", "REQ_DEFAULT"),
				),
			},
			{
				Config: testAccLbglobal_lbpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbglobal_lbpolicy_bindingNotExist("citrixadc_lbglobal_lbpolicy_binding.tf_lbglobal_lbpolicy_binding", "tf_pol", "REQ_DEFAULT"),
				),
			},
		},
	})
}

func testAccCheckLbglobal_lbpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbglobal_lbpolicy_binding id is set")
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

		// Keyless global binding: filter the GET by the `type` arg, then match on policyname.
		policyname := rs.Primary.Attributes["policyname"]
		typename := rs.Primary.Attributes["type"]
		findParams := service.FindParams{
			ResourceType:             service.Lbglobal_lbpolicy_binding.Type(),
			ArgsMap:                  map[string]string{"type": typename},
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one matching policyname
		found := false
		for _, v := range dataArr {
			if val, ok := v["policyname"].(string); ok && val == policyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("lbglobal_lbpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckLbglobal_lbpolicy_bindingNotExist(n string, policyname string, typename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		findParams := service.FindParams{
			ResourceType:             service.Lbglobal_lbpolicy_binding.Type(),
			ArgsMap:                  map[string]string{"type": typename},
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully NOT find the matching policyname
		found := false
		for _, v := range dataArr {
			if val, ok := v["policyname"].(string); ok && val == policyname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("lbglobal_lbpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckLbglobal_lbpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbglobal_lbpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		// Keyless global binding: filter by `type` arg and ensure no entry matches policyname.
		policyname := rs.Primary.Attributes["policyname"]
		typename := rs.Primary.Attributes["type"]
		findParams := service.FindParams{
			ResourceType:             service.Lbglobal_lbpolicy_binding.Type(),
			ArgsMap:                  map[string]string{"type": typename},
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Missing resource (empty array / 258) means it was destroyed - OK
			continue
		}

		for _, v := range dataArr {
			if val, ok := v["policyname"].(string); ok && val == policyname {
				return fmt.Errorf("lbglobal_lbpolicy_binding %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

const testAccLbglobal_lbpolicy_bindingDataSource_basic = `

resource "citrixadc_lbaction" "tf_act" {
	name  = "tf_act"
	type  = "SELECTIONORDER"
	value = [1]
}

resource "citrixadc_lbpolicy" "tf_pol" {
	name   = "tf_pol"
	rule   = "true"
	action = citrixadc_lbaction.tf_act.name
}

resource "citrixadc_lbglobal_lbpolicy_binding" "tf_lbglobal_lbpolicy_binding" {
	policyname = citrixadc_lbpolicy.tf_pol.name
	priority   = 100
	type       = "REQ_DEFAULT"
	depends_on = [citrixadc_lbpolicy.tf_pol]
}

data "citrixadc_lbglobal_lbpolicy_binding" "tf_lbglobal_lbpolicy_binding" {
	policyname = citrixadc_lbglobal_lbpolicy_binding.tf_lbglobal_lbpolicy_binding.policyname
	type       = citrixadc_lbglobal_lbpolicy_binding.tf_lbglobal_lbpolicy_binding.type
	depends_on = [citrixadc_lbglobal_lbpolicy_binding.tf_lbglobal_lbpolicy_binding]
}
`

func TestAccLbglobal_lbpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLbglobal_lbpolicy_bindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lbglobal_lbpolicy_binding.tf_lbglobal_lbpolicy_binding", "policyname", "tf_pol"),
					resource.TestCheckResourceAttr("data.citrixadc_lbglobal_lbpolicy_binding.tf_lbglobal_lbpolicy_binding", "priority", "100"),
					resource.TestCheckResourceAttr("data.citrixadc_lbglobal_lbpolicy_binding.tf_lbglobal_lbpolicy_binding", "type", "REQ_DEFAULT"),
				),
			},
		},
	})
}
