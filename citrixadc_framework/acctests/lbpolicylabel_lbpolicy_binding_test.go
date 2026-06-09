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

// lbpolicylabel_lbpolicy_binding is a binding_with_parent resource: parent
// lbpolicylabel (key labelname) <-> bound lbpolicy (policyname). Composite ID =
// labelname:<v>,policyname:<v>. Read/Exist parse the ID, GET by-name keyed on
// labelname, then filter the returned array on policyname. All attributes are
// RequiresReplace (no update); Create=PUT (UpdateUnnamedResource). Delete uses
// labelname (URL key) + policyname (arg).
//
// The participating entities are created first (HCL lifted from
// lbpolicylabel_test.go and lbpolicy_test.go, including the chained lbaction the
// lbpolicy requires). The lbpolicylabel uses policylabeltype HTTP and the lbpolicy
// is an HTTP-type lb policy so the two are compatible for binding.

const testAccLbpolicylabel_lbpolicy_binding_basic_step1 = `

resource "citrixadc_lbpolicylabel" "tf_lbpolicylabel" {
  labelname       = "tf_lbpolicylabel"
  policylabeltype = "HTTP"
  comment         = "test label"
}

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

resource "citrixadc_lbpolicylabel_lbpolicy_binding" "tf_lbpolicylabel_lbpolicy_binding" {
  labelname  = citrixadc_lbpolicylabel.tf_lbpolicylabel.labelname
  policyname = citrixadc_lbpolicy.tf_pol.name
  priority   = 100
  depends_on = [citrixadc_lbpolicylabel.tf_lbpolicylabel, citrixadc_lbpolicy.tf_pol]
}
`

const testAccLbpolicylabel_lbpolicy_binding_basic_step2 = `
	# Keep the participating entities but drop the actual binding to verify proper deletion

resource "citrixadc_lbpolicylabel" "tf_lbpolicylabel" {
  labelname       = "tf_lbpolicylabel"
  policylabeltype = "HTTP"
  comment         = "test label"
}

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

func TestAccLbpolicylabel_lbpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbpolicylabel_lbpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbpolicylabel_lbpolicy_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbpolicylabel_lbpolicy_bindingExist("citrixadc_lbpolicylabel_lbpolicy_binding.tf_lbpolicylabel_lbpolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_lbpolicylabel_lbpolicy_binding.tf_lbpolicylabel_lbpolicy_binding", "labelname", "tf_lbpolicylabel"),
					resource.TestCheckResourceAttr("citrixadc_lbpolicylabel_lbpolicy_binding.tf_lbpolicylabel_lbpolicy_binding", "policyname", "tf_pol"),
					resource.TestCheckResourceAttr("citrixadc_lbpolicylabel_lbpolicy_binding.tf_lbpolicylabel_lbpolicy_binding", "priority", "100"),
				),
			},
			{
				Config: testAccLbpolicylabel_lbpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbpolicylabel_lbpolicy_bindingNotExist("citrixadc_lbpolicylabel_lbpolicy_binding.tf_lbpolicylabel_lbpolicy_binding", "tf_lbpolicylabel", "tf_pol"),
				),
			},
		},
	})
}

func testAccCheckLbpolicylabel_lbpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbpolicylabel_lbpolicy_binding id is set")
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

		// binding_with_parent: parse composite ID, GET by-name keyed on labelname,
		// then filter the returned array on policyname.
		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"labelname", "policyname"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}
		labelname := idMap["labelname"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             service.Lbpolicylabel_lbpolicy_binding.Type(),
			ResourceName:             labelname,
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
			return fmt.Errorf("lbpolicylabel_lbpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckLbpolicylabel_lbpolicy_bindingNotExist(n string, labelname string, policyname string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		findParams := service.FindParams{
			ResourceType:             service.Lbpolicylabel_lbpolicy_binding.Type(),
			ResourceName:             labelname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Missing parent / empty array means the binding was destroyed - OK
		if err != nil {
			return nil
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
			return fmt.Errorf("lbpolicylabel_lbpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckLbpolicylabel_lbpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbpolicylabel_lbpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"labelname", "policyname"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}
		labelname := idMap["labelname"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             service.Lbpolicylabel_lbpolicy_binding.Type(),
			ResourceName:             labelname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Missing resource (empty array / 258) means it was destroyed - OK
			continue
		}

		for _, v := range dataArr {
			if val, ok := v["policyname"].(string); ok && val == policyname {
				return fmt.Errorf("lbpolicylabel_lbpolicy_binding %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

const testAccLbpolicylabel_lbpolicy_bindingDataSource_basic = `

resource "citrixadc_lbpolicylabel" "tf_lbpolicylabel" {
  labelname       = "tf_lbpolicylabel"
  policylabeltype = "HTTP"
  comment         = "test label"
}

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

resource "citrixadc_lbpolicylabel_lbpolicy_binding" "tf_lbpolicylabel_lbpolicy_binding" {
  labelname  = citrixadc_lbpolicylabel.tf_lbpolicylabel.labelname
  policyname = citrixadc_lbpolicy.tf_pol.name
  priority   = 100
  depends_on = [citrixadc_lbpolicylabel.tf_lbpolicylabel, citrixadc_lbpolicy.tf_pol]
}

data "citrixadc_lbpolicylabel_lbpolicy_binding" "tf_lbpolicylabel_lbpolicy_binding" {
  labelname  = citrixadc_lbpolicylabel_lbpolicy_binding.tf_lbpolicylabel_lbpolicy_binding.labelname
  policyname = citrixadc_lbpolicylabel_lbpolicy_binding.tf_lbpolicylabel_lbpolicy_binding.policyname
  depends_on = [citrixadc_lbpolicylabel_lbpolicy_binding.tf_lbpolicylabel_lbpolicy_binding]
}
`

func TestAccLbpolicylabel_lbpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLbpolicylabel_lbpolicy_bindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lbpolicylabel_lbpolicy_binding.tf_lbpolicylabel_lbpolicy_binding", "labelname", "tf_lbpolicylabel"),
					resource.TestCheckResourceAttr("data.citrixadc_lbpolicylabel_lbpolicy_binding.tf_lbpolicylabel_lbpolicy_binding", "policyname", "tf_pol"),
					resource.TestCheckResourceAttr("data.citrixadc_lbpolicylabel_lbpolicy_binding.tf_lbpolicylabel_lbpolicy_binding", "priority", "100"),
				),
			},
		},
	})
}
