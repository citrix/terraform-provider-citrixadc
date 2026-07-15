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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// Participating-entity config reused from cspolicylabel_test.go (cspolicylabel,
// cspolicylabeltype = HTTP) and cspolicy_test.go / csvserver_cspolicy_binding_test.go
// (cspolicy with a rule, plus a target lbvserver used as targetvserver).

const testAccCspolicylabel_cspolicy_binding_basic_step1 = `
resource "citrixadc_cspolicylabel" "tf_cspolicylabel" {
  cspolicylabeltype = "HTTP"
  labelname         = "tf_cspolicylabel"
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_cspollabel_lb"
  servicetype = "HTTP"
}

resource "citrixadc_cspolicy" "tf_cspolicy" {
  policyname = "tf_cspollabel_policy"
  rule       = "CLIENT.IP.SRC.SUBNET(24).EQ(10.217.84.0)"
}

resource "citrixadc_cspolicylabel_cspolicy_binding" "tf_binding" {
  labelname     = citrixadc_cspolicylabel.tf_cspolicylabel.labelname
  policyname    = citrixadc_cspolicy.tf_cspolicy.policyname
  priority      = 100
  targetvserver = citrixadc_lbvserver.tf_lbvserver.name

  depends_on = [
    citrixadc_cspolicylabel.tf_cspolicylabel,
    citrixadc_cspolicy.tf_cspolicy,
    citrixadc_lbvserver.tf_lbvserver,
  ]
}
`

// step2 drops the binding (but keeps the participating entities) to confirm the
// binding is deleted. All binding attributes are RequiresReplace (no update).
const testAccCspolicylabel_cspolicy_binding_basic_step2 = `
resource "citrixadc_cspolicylabel" "tf_cspolicylabel" {
  cspolicylabeltype = "HTTP"
  labelname         = "tf_cspolicylabel"
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_cspollabel_lb"
  servicetype = "HTTP"
}

resource "citrixadc_cspolicy" "tf_cspolicy" {
  policyname = "tf_cspollabel_policy"
  rule       = "CLIENT.IP.SRC.SUBNET(24).EQ(10.217.84.0)"
}
`

func TestAccCspolicylabel_cspolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCspolicylabel_cspolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCspolicylabel_cspolicy_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCspolicylabel_cspolicy_bindingExist("citrixadc_cspolicylabel_cspolicy_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_cspolicylabel_cspolicy_binding.tf_binding", "labelname", "tf_cspolicylabel"),
					resource.TestCheckResourceAttr("citrixadc_cspolicylabel_cspolicy_binding.tf_binding", "policyname", "tf_cspollabel_policy"),
					resource.TestCheckResourceAttr("citrixadc_cspolicylabel_cspolicy_binding.tf_binding", "priority", "100"),
					resource.TestCheckResourceAttr("citrixadc_cspolicylabel_cspolicy_binding.tf_binding", "targetvserver", "tf_cspollabel_lb"),
				),
			},
			{
				// Binding removed - CheckDestroy-style verification that it no longer exists.
				Config: testAccCspolicylabel_cspolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCspolicylabel_cspolicy_bindingNotExist("tf_cspolicylabel", "tf_cspollabel_policy"),
				),
			},
		},
	})
}

func TestAccCspolicylabel_cspolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_cspolicylabel_cspolicy_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCspolicylabel_cspolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCspolicylabel_cspolicy_binding_basic_step1,
			},
			{
				Config:                  testAccCspolicylabel_cspolicy_binding_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckCspolicylabel_cspolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cspolicylabel_cspolicy_binding ID is set")
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
			ResourceType:             service.Cspolicylabel_cspolicy_binding.Type(),
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
			return fmt.Errorf("cspolicylabel_cspolicy_binding %s not found", n)
		}

		return nil
	}
}

// testAccCheckCspolicylabel_cspolicy_bindingNotExist verifies a specific binding
// (labelname/policyname) is no longer present on the ADC, used after step2 drops it.
func testAccCheckCspolicylabel_cspolicy_bindingNotExist(labelname, policyname string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		findParams := service.FindParams{
			ResourceType:             service.Cspolicylabel_cspolicy_binding.Type(),
			ResourceName:             labelname,
			ResourceMissingErrorCode: 258,
		}

		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		for _, v := range dataArr {
			if val, ok := v["policyname"].(string); ok && val == policyname {
				return fmt.Errorf("cspolicylabel_cspolicy_binding %s:%s still exists", labelname, policyname)
			}
		}

		return nil
	}
}

func testAccCheckCspolicylabel_cspolicy_bindingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_cspolicylabel_cspolicy_binding" {
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
			ResourceType:             service.Cspolicylabel_cspolicy_binding.Type(),
			ResourceName:             labelname,
			ResourceMissingErrorCode: 258,
		}

		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Parent cspolicylabel already deleted (errorcode 3087) => binding is gone too.
			if strings.Contains(err.Error(), "errorcode 3087") {
				continue
			}
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
			return fmt.Errorf("cspolicylabel_cspolicy_binding %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

const testAccCspolicylabel_cspolicy_bindingDataSource_basic = `
resource "citrixadc_cspolicylabel" "tf_cspolicylabel" {
  cspolicylabeltype = "HTTP"
  labelname         = "tf_cspolicylabel_ds"
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_cspollabel_lb_ds"
  servicetype = "HTTP"
}

resource "citrixadc_cspolicy" "tf_cspolicy" {
  policyname = "tf_cspollabel_policy_ds"
  rule       = "CLIENT.IP.SRC.SUBNET(24).EQ(10.217.84.0)"
}

resource "citrixadc_cspolicylabel_cspolicy_binding" "tf_binding" {
  labelname     = citrixadc_cspolicylabel.tf_cspolicylabel.labelname
  policyname    = citrixadc_cspolicy.tf_cspolicy.policyname
  priority      = 100
  targetvserver = citrixadc_lbvserver.tf_lbvserver.name

  depends_on = [
    citrixadc_cspolicylabel.tf_cspolicylabel,
    citrixadc_cspolicy.tf_cspolicy,
    citrixadc_lbvserver.tf_lbvserver,
  ]
}

data "citrixadc_cspolicylabel_cspolicy_binding" "tf_binding" {
  labelname  = citrixadc_cspolicylabel_cspolicy_binding.tf_binding.labelname
  policyname = citrixadc_cspolicylabel_cspolicy_binding.tf_binding.policyname
  depends_on = [citrixadc_cspolicylabel_cspolicy_binding.tf_binding]
}
`

func TestAccCspolicylabel_cspolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccCspolicylabel_cspolicy_bindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_cspolicylabel_cspolicy_binding.tf_binding", "labelname", "tf_cspolicylabel_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_cspolicylabel_cspolicy_binding.tf_binding", "policyname", "tf_cspollabel_policy_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_cspolicylabel_cspolicy_binding.tf_binding", "priority", "100"),
					resource.TestCheckResourceAttr("data.citrixadc_cspolicylabel_cspolicy_binding.tf_binding", "targetvserver", "tf_cspollabel_lb_ds"),
				),
			},
		},
	})
}
