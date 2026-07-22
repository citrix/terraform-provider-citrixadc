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
// Participating-entity config reused from existing acceptance tests:
//   - videooptimizationdetectionpolicylabel_test.go (the parent LABEL:
//     labelname + policylabeltype = videoopt_req)
//   - videooptimizationdetectionpolicy_test.go (the bound POLICY:
//     name + rule + action, which itself depends on a
//     videooptimizationdetectionaction with type = clear_text_abr)

const testAccVideooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding_basic_step1 = `
resource "citrixadc_videooptimizationdetectionpolicylabel" "tf_videooptimizationdetectionpolicylabel" {
  labelname       = "tf_videoopt_detection_pl"
  policylabeltype = "videoopt_req"
  comment         = "test_comment"
}

resource "citrixadc_videooptimizationdetectionaction" "tf_detectionaction" {
  name = "tf_videooptimizationdetectionaction"
  type = "clear_text_abr"
}

resource "citrixadc_videooptimizationdetectionpolicy" "tf_detectionpolicy" {
  name   = "tf_videooptimizationdetectionpolicy"
  rule   = "true"
  action = citrixadc_videooptimizationdetectionaction.tf_detectionaction.name
}

resource "citrixadc_videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding" "tf_binding" {
  labelname  = citrixadc_videooptimizationdetectionpolicylabel.tf_videooptimizationdetectionpolicylabel.labelname
  policyname = citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy.name
  priority   = 100

  depends_on = [
    citrixadc_videooptimizationdetectionpolicylabel.tf_videooptimizationdetectionpolicylabel,
    citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy,
  ]
}
`

// step2 drops the binding (but keeps the participating entities) to confirm the
// binding is deleted. All binding attributes are RequiresReplace (no update).
const testAccVideooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding_basic_step2 = `
resource "citrixadc_videooptimizationdetectionpolicylabel" "tf_videooptimizationdetectionpolicylabel" {
  labelname       = "tf_videoopt_detection_pl"
  policylabeltype = "videoopt_req"
  comment         = "test_comment"
}

resource "citrixadc_videooptimizationdetectionaction" "tf_detectionaction" {
  name = "tf_videooptimizationdetectionaction"
  type = "clear_text_abr"
}

resource "citrixadc_videooptimizationdetectionpolicy" "tf_detectionpolicy" {
  name   = "tf_videooptimizationdetectionpolicy"
  rule   = "true"
  action = citrixadc_videooptimizationdetectionaction.tf_detectionaction.name
}
`

func TestAccVideooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVideooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVideooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_bindingExist("citrixadc_videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding.tf_binding", "labelname", "tf_videoopt_detection_pl"),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding.tf_binding", "policyname", "tf_videooptimizationdetectionpolicy"),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding.tf_binding", "priority", "100"),
				),
			},
			{
				// Binding removed - verify it no longer exists on the ADC.
				Config: testAccVideooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_bindingNotExist("tf_videoopt_detection_pl", "tf_videooptimizationdetectionpolicy"),
				),
			},
		},
	})
}

func TestAccVideooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVideooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVideooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding_basic_step1,
			},
			{
				Config:                  testAccVideooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckVideooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding ID is set")
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
			ResourceType:             service.Videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding.Type(),
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
			return fmt.Errorf("videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding %s not found", n)
		}

		return nil
	}
}

// testAccCheckVideooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_bindingNotExist
// verifies a specific binding (labelname/policyname) is no longer present on the
// ADC, used after step2 drops it.
func testAccCheckVideooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_bindingNotExist(labelname, policyname string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		findParams := service.FindParams{
			ResourceType:             service.Videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding.Type(),
			ResourceName:             labelname,
			ResourceMissingErrorCode: 258,
		}

		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		for _, v := range dataArr {
			if val, ok := v["policyname"].(string); ok && val == policyname {
				return fmt.Errorf("videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding %s:%s still exists", labelname, policyname)
			}
		}

		return nil
	}
}

func testAccCheckVideooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_bindingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding" {
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
			ResourceType:             service.Videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding.Type(),
			ResourceName:             labelname,
			ResourceMissingErrorCode: 258,
		}

		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// The parent label is destroyed before this check runs, so the
			// NITRO GET on the (now absent) label returns an error. That means
			// the binding is gone too - treat it as destroyed.
			continue
		}

		found := false
		for _, v := range dataArr {
			if val, ok := v["policyname"].(string); ok && val == policyname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

const testAccVideooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_bindingDataSource_basic = `
resource "citrixadc_videooptimizationdetectionpolicylabel" "tf_videooptimizationdetectionpolicylabel" {
  labelname       = "tf_videoopt_detection_pl_ds"
  policylabeltype = "videoopt_req"
  comment         = "test_comment"
}

resource "citrixadc_videooptimizationdetectionaction" "tf_detectionaction" {
  name = "tf_videooptimizationdetectionaction_ds"
  type = "clear_text_abr"
}

resource "citrixadc_videooptimizationdetectionpolicy" "tf_detectionpolicy" {
  name   = "tf_videooptimizationdetectionpolicy_ds"
  rule   = "true"
  action = citrixadc_videooptimizationdetectionaction.tf_detectionaction.name
}

resource "citrixadc_videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding" "tf_binding" {
  labelname  = citrixadc_videooptimizationdetectionpolicylabel.tf_videooptimizationdetectionpolicylabel.labelname
  policyname = citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy.name
  priority   = 100

  depends_on = [
    citrixadc_videooptimizationdetectionpolicylabel.tf_videooptimizationdetectionpolicylabel,
    citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy,
  ]
}

data "citrixadc_videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding" "tf_binding" {
  labelname  = citrixadc_videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding.tf_binding.labelname
  policyname = citrixadc_videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding.tf_binding.policyname
  priority   = citrixadc_videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding.tf_binding.priority
  depends_on = [citrixadc_videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding.tf_binding]
}
`

func TestAccVideooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccVideooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_bindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding.tf_binding", "labelname", "tf_videoopt_detection_pl_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding.tf_binding", "policyname", "tf_videooptimizationdetectionpolicy_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_videooptimizationdetectionpolicylabel_videooptimizationdetectionpolicy_binding.tf_binding", "priority", "100"),
				),
			},
		},
	})
}
