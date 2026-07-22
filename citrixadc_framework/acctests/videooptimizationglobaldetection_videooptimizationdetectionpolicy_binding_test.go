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
	"strconv"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// global_binding (NO parent resource - binds a videooptimization detection policy to
// the global detection bind point, a singleton). Composite ID = policyname,priority,type.
// All binding attributes are RequiresReplace (no update path). globalbindtype is
// Computed/read-only (never set or asserted).
//
// Participating-entity config reused from existing acceptance tests:
//   - videooptimizationdetectionpolicy_test.go (the bound POLICY: name + rule + action,
//     which itself depends on a videooptimizationdetectionaction with
//     type = clear_text_abr)
//
// No parent global resource is needed (the global detection bind point is a singleton).

const testAccVideooptimizationglobaldetection_videooptimizationdetectionpolicy_binding_basic_step1 = `
resource "citrixadc_videooptimizationdetectionaction" "tf_detectionaction" {
  name = "tf_videooptimizationdetectionaction"
  type = "clear_text_abr"
}

resource "citrixadc_videooptimizationdetectionpolicy" "tf_detectionpolicy" {
  name   = "tf_videooptimizationdetectionpolicy"
  rule   = "true"
  action = citrixadc_videooptimizationdetectionaction.tf_detectionaction.name
}

resource "citrixadc_videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding" "tf_binding" {
  policyname = citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy.name
  priority   = 100
  type       = "REQ_DEFAULT"

  depends_on = [
    citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy,
  ]
}
`

// step2 drops the binding (but keeps the participating entities) to confirm the
// binding is deleted. All binding attributes are RequiresReplace (no update).
const testAccVideooptimizationglobaldetection_videooptimizationdetectionpolicy_binding_basic_step2 = `
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

func TestAccVideooptimizationglobaldetection_videooptimizationdetectionpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVideooptimizationglobaldetection_videooptimizationdetectionpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVideooptimizationglobaldetection_videooptimizationdetectionpolicy_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationglobaldetection_videooptimizationdetectionpolicy_bindingExist("citrixadc_videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding.tf_binding", "policyname", "tf_videooptimizationdetectionpolicy"),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding.tf_binding", "priority", "100"),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding.tf_binding", "type", "REQ_DEFAULT"),
				),
			},
			{
				// Binding removed - verify it no longer exists on the ADC.
				Config: testAccVideooptimizationglobaldetection_videooptimizationdetectionpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationglobaldetection_videooptimizationdetectionpolicy_bindingNotExist("tf_videooptimizationdetectionpolicy", 100, "REQ_DEFAULT"),
				),
			},
		},
	})
}

func TestAccVideooptimizationglobaldetection_videooptimizationdetectionpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVideooptimizationglobaldetection_videooptimizationdetectionpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVideooptimizationglobaldetection_videooptimizationdetectionpolicy_binding_basic_step1,
			},
			{
				Config:                  testAccVideooptimizationglobaldetection_videooptimizationdetectionpolicy_binding_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckVideooptimizationglobaldetection_videooptimizationdetectionpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding ID is set")
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

		policyname := idMap["policyname"]
		priority := idMap["priority"]

		// Global binding - no parent resource name; filter the returned array by the
		// key attributes (policyname + priority).
		argsMap := make(map[string]string)
		if val, ok := idMap["type"]; ok && val != "" {
			argsMap["type"] = val
		}

		findParams := service.FindParams{
			ResourceType:             service.Videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding.Type(),
			ArgsMap:                  argsMap,
			ResourceMissingErrorCode: 258,
		}

		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if val, ok := v["policyname"].(string); ok && val == policyname {
				if pVal, ok := v["priority"]; ok {
					pInt, _ := utils.ConvertToInt64(pVal)
					idInt, _ := strconv.ParseInt(priority, 10, 64)
					if pInt == idInt {
						found = true
						break
					}
				}
			}
		}

		if !found {
			return fmt.Errorf("videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding %s not found", n)
		}

		return nil
	}
}

// testAccCheckVideooptimizationglobaldetection_videooptimizationdetectionpolicy_bindingNotExist
// verifies a specific binding (policyname/priority/type) is no longer present on the
// ADC, used after step2 drops it.
func testAccCheckVideooptimizationglobaldetection_videooptimizationdetectionpolicy_bindingNotExist(policyname string, priority int64, bindType string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		argsMap := make(map[string]string)
		if bindType != "" {
			argsMap["type"] = bindType
		}

		findParams := service.FindParams{
			ResourceType:             service.Videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding.Type(),
			ArgsMap:                  argsMap,
			ResourceMissingErrorCode: 258,
		}

		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		for _, v := range dataArr {
			if val, ok := v["policyname"].(string); ok && val == policyname {
				if pVal, ok := v["priority"]; ok {
					pInt, _ := utils.ConvertToInt64(pVal)
					if pInt == priority {
						return fmt.Errorf("videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding %s:%d still exists", policyname, priority)
					}
				}
			}
		}

		return nil
	}
}

func testAccCheckVideooptimizationglobaldetection_videooptimizationdetectionpolicy_bindingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %q: %v", rs.Primary.ID, err)
		}

		policyname := idMap["policyname"]
		priority := idMap["priority"]

		argsMap := make(map[string]string)
		if val, ok := idMap["type"]; ok && val != "" {
			argsMap["type"] = val
		}

		findParams := service.FindParams{
			ResourceType:             service.Videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding.Type(),
			ArgsMap:                  argsMap,
			ResourceMissingErrorCode: 258,
		}

		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if val, ok := v["policyname"].(string); ok && val == policyname {
				if pVal, ok := v["priority"]; ok {
					pInt, _ := utils.ConvertToInt64(pVal)
					idInt, _ := strconv.ParseInt(priority, 10, 64)
					if pInt == idInt {
						found = true
						break
					}
				}
			}
		}

		if found {
			return fmt.Errorf("videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

const testAccVideooptimizationglobaldetection_videooptimizationdetectionpolicy_bindingDataSource_basic = `
resource "citrixadc_videooptimizationdetectionaction" "tf_detectionaction" {
  name = "tf_videooptimizationdetectionaction_ds"
  type = "clear_text_abr"
}

resource "citrixadc_videooptimizationdetectionpolicy" "tf_detectionpolicy" {
  name   = "tf_videooptimizationdetectionpolicy_ds"
  rule   = "true"
  action = citrixadc_videooptimizationdetectionaction.tf_detectionaction.name
}

resource "citrixadc_videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding" "tf_binding" {
  policyname = citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy.name
  priority   = 100
  type       = "REQ_DEFAULT"

  depends_on = [
    citrixadc_videooptimizationdetectionpolicy.tf_detectionpolicy,
  ]
}

data "citrixadc_videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding" "tf_binding" {
  policyname = citrixadc_videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding.tf_binding.policyname
  priority   = citrixadc_videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding.tf_binding.priority
  type       = citrixadc_videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding.tf_binding.type
  depends_on = [citrixadc_videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding.tf_binding]
}
`

func TestAccVideooptimizationglobaldetection_videooptimizationdetectionpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccVideooptimizationglobaldetection_videooptimizationdetectionpolicy_bindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding.tf_binding", "policyname", "tf_videooptimizationdetectionpolicy_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding.tf_binding", "priority", "100"),
				),
			},
		},
	})
}
