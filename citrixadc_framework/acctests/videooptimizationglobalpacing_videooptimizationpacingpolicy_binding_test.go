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

// global_binding (NO parent resource - binds a videooptimization pacing policy to the
// global pacing bind point, a singleton). Composite ID = policyname,priority,type.
// All binding attributes are RequiresReplace (no update path). globalbindtype is
// Computed/read-only (never set or asserted).
//
// Participating-entity config reused from existing acceptance tests:
//   - videooptimizationpacingpolicy_test.go (the bound POLICY: name + rule + action,
//     which itself depends on a videooptimizationpacingaction with rate)
//   - videooptimizationpacingaction_test.go (the action the policy references)
//
// No parent global resource is needed (the global pacing bind point is a singleton).
//
// NOTE: the video pacing feature is marked deprecated by NetScaler. The participating
// videooptimizationpacingpolicy / videooptimizationpacingaction acceptance tests are
// NOT skipped, so this binding test is generated normally (no t.Skip). If pacing
// creation begins to fail on newer ADC builds due to deprecation, add a t.Skip here
// noting pacing deprecation.

const testAccVideooptimizationglobalpacing_videooptimizationpacingpolicy_binding_basic_step1 = `
resource "citrixadc_videooptimizationpacingaction" "tf_pacingaction" {
  name = "tf_videooptimizationpacingaction"
  rate = 10
}

resource "citrixadc_videooptimizationpacingpolicy" "tf_pacingpolicy" {
  name   = "tf_videooptimizationpacingpolicy"
  rule   = "true"
  action = citrixadc_videooptimizationpacingaction.tf_pacingaction.name
}

resource "citrixadc_videooptimizationglobalpacing_videooptimizationpacingpolicy_binding" "tf_binding" {
  policyname = citrixadc_videooptimizationpacingpolicy.tf_pacingpolicy.name
  priority   = 100
  type       = "REQ_DEFAULT"

  depends_on = [
    citrixadc_videooptimizationpacingpolicy.tf_pacingpolicy,
  ]
}
`

// step2 drops the binding (but keeps the participating entities) to confirm the
// binding is deleted. All binding attributes are RequiresReplace (no update).
const testAccVideooptimizationglobalpacing_videooptimizationpacingpolicy_binding_basic_step2 = `
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

func TestAccVideooptimizationglobalpacing_videooptimizationpacingpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVideooptimizationglobalpacing_videooptimizationpacingpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVideooptimizationglobalpacing_videooptimizationpacingpolicy_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationglobalpacing_videooptimizationpacingpolicy_bindingExist("citrixadc_videooptimizationglobalpacing_videooptimizationpacingpolicy_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationglobalpacing_videooptimizationpacingpolicy_binding.tf_binding", "policyname", "tf_videooptimizationpacingpolicy"),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationglobalpacing_videooptimizationpacingpolicy_binding.tf_binding", "priority", "100"),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationglobalpacing_videooptimizationpacingpolicy_binding.tf_binding", "type", "REQ_DEFAULT"),
				),
			},
			{
				// Binding removed - verify it no longer exists on the ADC.
				Config: testAccVideooptimizationglobalpacing_videooptimizationpacingpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationglobalpacing_videooptimizationpacingpolicy_bindingNotExist("tf_videooptimizationpacingpolicy", 100, "REQ_DEFAULT"),
				),
			},
		},
	})
}

func testAccCheckVideooptimizationglobalpacing_videooptimizationpacingpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No videooptimizationglobalpacing_videooptimizationpacingpolicy_binding ID is set")
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
			ResourceType:             service.Videooptimizationglobalpacing_videooptimizationpacingpolicy_binding.Type(),
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
			return fmt.Errorf("videooptimizationglobalpacing_videooptimizationpacingpolicy_binding %s not found", n)
		}

		return nil
	}
}

// testAccCheckVideooptimizationglobalpacing_videooptimizationpacingpolicy_bindingNotExist
// verifies a specific binding (policyname/priority/type) is no longer present on the
// ADC, used after step2 drops it.
func testAccCheckVideooptimizationglobalpacing_videooptimizationpacingpolicy_bindingNotExist(policyname string, priority int64, bindType string) resource.TestCheckFunc {
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
			ResourceType:             service.Videooptimizationglobalpacing_videooptimizationpacingpolicy_binding.Type(),
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
						return fmt.Errorf("videooptimizationglobalpacing_videooptimizationpacingpolicy_binding %s:%d still exists", policyname, priority)
					}
				}
			}
		}

		return nil
	}
}

func testAccCheckVideooptimizationglobalpacing_videooptimizationpacingpolicy_bindingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_videooptimizationglobalpacing_videooptimizationpacingpolicy_binding" {
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
			ResourceType:             service.Videooptimizationglobalpacing_videooptimizationpacingpolicy_binding.Type(),
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
			return fmt.Errorf("videooptimizationglobalpacing_videooptimizationpacingpolicy_binding %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

const testAccVideooptimizationglobalpacing_videooptimizationpacingpolicy_bindingDataSource_basic = `
resource "citrixadc_videooptimizationpacingaction" "tf_pacingaction" {
  name = "tf_videooptimizationpacingaction_ds"
  rate = 10
}

resource "citrixadc_videooptimizationpacingpolicy" "tf_pacingpolicy" {
  name   = "tf_videooptimizationpacingpolicy_ds"
  rule   = "true"
  action = citrixadc_videooptimizationpacingaction.tf_pacingaction.name
}

resource "citrixadc_videooptimizationglobalpacing_videooptimizationpacingpolicy_binding" "tf_binding" {
  policyname = citrixadc_videooptimizationpacingpolicy.tf_pacingpolicy.name
  priority   = 100
  type       = "REQ_DEFAULT"

  depends_on = [
    citrixadc_videooptimizationpacingpolicy.tf_pacingpolicy,
  ]
}

data "citrixadc_videooptimizationglobalpacing_videooptimizationpacingpolicy_binding" "tf_binding" {
  policyname = citrixadc_videooptimizationglobalpacing_videooptimizationpacingpolicy_binding.tf_binding.policyname
  priority   = citrixadc_videooptimizationglobalpacing_videooptimizationpacingpolicy_binding.tf_binding.priority
  type       = citrixadc_videooptimizationglobalpacing_videooptimizationpacingpolicy_binding.tf_binding.type
  depends_on = [citrixadc_videooptimizationglobalpacing_videooptimizationpacingpolicy_binding.tf_binding]
}
`

func TestAccVideooptimizationglobalpacing_videooptimizationpacingpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccVideooptimizationglobalpacing_videooptimizationpacingpolicy_bindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_videooptimizationglobalpacing_videooptimizationpacingpolicy_binding.tf_binding", "policyname", "tf_videooptimizationpacingpolicy_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_videooptimizationglobalpacing_videooptimizationpacingpolicy_binding.tf_binding", "priority", "100"),
				),
			},
		},
	})
}
