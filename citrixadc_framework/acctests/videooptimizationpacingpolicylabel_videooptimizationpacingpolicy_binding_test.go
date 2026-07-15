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

// binding_with_parent. Composite ID = labelname,policyname,priority.
// All binding attributes are RequiresReplace (no update path).
//
// NOTE: video optimization pacing is deprecated in recent NetScaler releases,
// but it remains configurable and the sibling videooptimizationpacingpolicy /
// pacingpolicylabel tests are not skipped, so this test is generated normally
// (no t.Skip).
//
// Participating-entity config reused from existing acceptance tests:
//   - videooptimizationpacingpolicylabel_test.go (the parent LABEL:
//     labelname + policylabeltype = videoopt_req)
//   - videooptimizationpacingpolicy_test.go (the bound POLICY:
//     name + rule + action, which itself depends on a
//     videooptimizationpacingaction with rate = 10)

const testAccVideooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding_basic_step1 = `
resource "citrixadc_videooptimizationpacingpolicylabel" "tf_videooptimizationpacingpolicylabel" {
  labelname       = "tf_videoopt_pacing_pl"
  policylabeltype = "videoopt_req"
  comment         = "test_comment"
}

resource "citrixadc_videooptimizationpacingaction" "tf_pacingaction" {
  name = "tf_videooptimizationpacingaction"
  rate = 10
}

resource "citrixadc_videooptimizationpacingpolicy" "tf_pacingpolicy" {
  name   = "tf_videooptimizationpacingpolicy"
  rule   = "true"
  action = citrixadc_videooptimizationpacingaction.tf_pacingaction.name
}

resource "citrixadc_videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding" "tf_binding" {
  labelname  = citrixadc_videooptimizationpacingpolicylabel.tf_videooptimizationpacingpolicylabel.labelname
  policyname = citrixadc_videooptimizationpacingpolicy.tf_pacingpolicy.name
  priority   = 100

  depends_on = [
    citrixadc_videooptimizationpacingpolicylabel.tf_videooptimizationpacingpolicylabel,
    citrixadc_videooptimizationpacingpolicy.tf_pacingpolicy,
  ]
}
`

// step2 drops the binding (but keeps the participating entities) to confirm the
// binding is deleted. All binding attributes are RequiresReplace (no update).
const testAccVideooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding_basic_step2 = `
resource "citrixadc_videooptimizationpacingpolicylabel" "tf_videooptimizationpacingpolicylabel" {
  labelname       = "tf_videoopt_pacing_pl"
  policylabeltype = "videoopt_req"
  comment         = "test_comment"
}

resource "citrixadc_videooptimizationpacingaction" "tf_pacingaction" {
  name = "tf_videooptimizationpacingaction"
  rate = 10
}

resource "citrixadc_videooptimizationpacingpolicy" "tf_pacingpolicy" {
  name   = "tf_videooptimizationpacingpolicy"
  rule   = "true"
  action = citrixadc_videooptimizationpacingaction.tf_pacingaction.name
}
`

func TestAccVideooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVideooptimizationpacingpolicylabel_videooptimizationpacingpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVideooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationpacingpolicylabel_videooptimizationpacingpolicy_bindingExist("citrixadc_videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding.tf_binding", "labelname", "tf_videoopt_pacing_pl"),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding.tf_binding", "policyname", "tf_videooptimizationpacingpolicy"),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding.tf_binding", "priority", "100"),
				),
			},
			{
				// Binding removed - verify it no longer exists on the ADC.
				Config: testAccVideooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationpacingpolicylabel_videooptimizationpacingpolicy_bindingNotExist("tf_videoopt_pacing_pl", "tf_videooptimizationpacingpolicy"),
				),
			},
		},
	})
}

func TestAccVideooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVideooptimizationpacingpolicylabel_videooptimizationpacingpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVideooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding_basic_step1,
			},
			{
				Config:                  testAccVideooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
			{
				// Drop the binding (but keep the participating entities) so the
				// binding is removed from state before the framework's final
				// destroy. This mirrors the basic test and avoids CheckDestroy
				// querying bindings under an already-deleted parent label
				// (which returns NITRO errorcode 3087 instead of 258).
				Config: testAccVideooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationpacingpolicylabel_videooptimizationpacingpolicy_bindingNotExist("tf_videoopt_pacing_pl", "tf_videooptimizationpacingpolicy"),
				),
			},
		},
	})
}

func testAccCheckVideooptimizationpacingpolicylabel_videooptimizationpacingpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %q: %v", rs.Primary.ID, err)
		}

		labelname := idMap["labelname"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             service.Videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding.Type(),
			ResourceName:             labelname,
			ResourceMissingErrorCode: 258,
		}

		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if val, ok := v["policyname"].(string); ok && val == policyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding %s not found", n)
		}

		return nil
	}
}

// testAccCheckVideooptimizationpacingpolicylabel_videooptimizationpacingpolicy_bindingNotExist
// verifies a specific binding (labelname/policyname) is no longer present on the
// ADC, used after step2 drops it.
func testAccCheckVideooptimizationpacingpolicylabel_videooptimizationpacingpolicy_bindingNotExist(labelname, policyname string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		findParams := service.FindParams{
			ResourceType:             service.Videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding.Type(),
			ResourceName:             labelname,
			ResourceMissingErrorCode: 258,
		}

		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		for _, v := range dataArr {
			if val, ok := v["policyname"].(string); ok && val == policyname {
				return fmt.Errorf("videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding %s:%s still exists", labelname, policyname)
			}
		}

		return nil
	}
}

func testAccCheckVideooptimizationpacingpolicylabel_videooptimizationpacingpolicy_bindingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %q: %v", rs.Primary.ID, err)
		}

		labelname := idMap["labelname"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             service.Videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding.Type(),
			ResourceName:             labelname,
			ResourceMissingErrorCode: 258,
		}

		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if val, ok := v["policyname"].(string); ok && val == policyname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

const testAccVideooptimizationpacingpolicylabel_videooptimizationpacingpolicy_bindingDataSource_basic = `
resource "citrixadc_videooptimizationpacingpolicylabel" "tf_videooptimizationpacingpolicylabel" {
  labelname       = "tf_videoopt_pacing_pl_ds"
  policylabeltype = "videoopt_req"
  comment         = "test_comment"
}

resource "citrixadc_videooptimizationpacingaction" "tf_pacingaction" {
  name = "tf_videooptimizationpacingaction_ds"
  rate = 10
}

resource "citrixadc_videooptimizationpacingpolicy" "tf_pacingpolicy" {
  name   = "tf_videooptimizationpacingpolicy_ds"
  rule   = "true"
  action = citrixadc_videooptimizationpacingaction.tf_pacingaction.name
}

resource "citrixadc_videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding" "tf_binding" {
  labelname  = citrixadc_videooptimizationpacingpolicylabel.tf_videooptimizationpacingpolicylabel.labelname
  policyname = citrixadc_videooptimizationpacingpolicy.tf_pacingpolicy.name
  priority   = 100

  depends_on = [
    citrixadc_videooptimizationpacingpolicylabel.tf_videooptimizationpacingpolicylabel,
    citrixadc_videooptimizationpacingpolicy.tf_pacingpolicy,
  ]
}

data "citrixadc_videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding" "tf_binding" {
  labelname  = citrixadc_videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding.tf_binding.labelname
  policyname = citrixadc_videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding.tf_binding.policyname
  priority   = citrixadc_videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding.tf_binding.priority
  depends_on = [citrixadc_videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding.tf_binding]
}
`

func TestAccVideooptimizationpacingpolicylabel_videooptimizationpacingpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccVideooptimizationpacingpolicylabel_videooptimizationpacingpolicy_bindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding.tf_binding", "labelname", "tf_videoopt_pacing_pl_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding.tf_binding", "policyname", "tf_videooptimizationpacingpolicy_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding.tf_binding", "priority", "100"),
				),
			},
		},
	})
}
