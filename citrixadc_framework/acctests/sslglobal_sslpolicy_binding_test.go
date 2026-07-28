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

// Participating entities (lifted from sslpolicy_test.go):
//   - citrixadc_sslaction (referenced by the sslpolicy action)
//   - citrixadc_sslpolicy (the policy being bound globally)
// The binding is a global keyless binding; composite ID = policyname,type,priority.

const testAccSslglobalSslpolicyBinding_basic_step1 = `
	resource "citrixadc_sslaction" "tf_sslaction" {
		name                   = "tf_sslaction_glb"
		clientauth             = "DOCLIENTAUTH"
		clientcertverification = "Mandatory"
	}

	resource "citrixadc_sslpolicy" "tf_sslpolicy" {
		name   = "tf_sslpolicy_glb"
		rule   = "true"
		action = citrixadc_sslaction.tf_sslaction.name
	}

	resource "citrixadc_sslglobal_sslpolicy_binding" "tf_sslglobal_sslpolicy_binding" {
		policyname = citrixadc_sslpolicy.tf_sslpolicy.name
		priority   = 100
		type       = "CONTROL_OVERRIDE"

		depends_on = [citrixadc_sslpolicy.tf_sslpolicy]
	}
`

// step2 drops the binding (keeps the participating entities), verifying clean unbind.
const testAccSslglobalSslpolicyBinding_basic_step2 = `
	resource "citrixadc_sslaction" "tf_sslaction" {
		name                   = "tf_sslaction_glb"
		clientauth             = "DOCLIENTAUTH"
		clientcertverification = "Mandatory"
	}

	resource "citrixadc_sslpolicy" "tf_sslpolicy" {
		name   = "tf_sslpolicy_glb"
		rule   = "true"
		action = citrixadc_sslaction.tf_sslaction.name
	}
`

func TestAccSslglobalSslpolicyBinding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslglobalSslpolicyBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslglobalSslpolicyBinding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslglobalSslpolicyBindingExist("citrixadc_sslglobal_sslpolicy_binding.tf_sslglobal_sslpolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_sslglobal_sslpolicy_binding.tf_sslglobal_sslpolicy_binding", "policyname", "tf_sslpolicy_glb"),
					resource.TestCheckResourceAttr("citrixadc_sslglobal_sslpolicy_binding.tf_sslglobal_sslpolicy_binding", "priority", "100"),
					resource.TestCheckResourceAttr("citrixadc_sslglobal_sslpolicy_binding.tf_sslglobal_sslpolicy_binding", "type", "CONTROL_OVERRIDE"),
				),
			},
			{
				Config: testAccSslglobalSslpolicyBinding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslpolicyExist("citrixadc_sslpolicy.tf_sslpolicy", nil),
				),
			},
		},
	})
}

func TestAccSslglobalSslpolicyBinding_import(t *testing.T) {
	const resAddr = "citrixadc_sslglobal_sslpolicy_binding.tf_sslglobal_sslpolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslglobalSslpolicyBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslglobalSslpolicyBinding_basic_step1,
			},
			{
				Config:                  testAccSslglobalSslpolicyBinding_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckSslglobalSslpolicyBindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslglobal_sslpolicy_binding id is set")
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

		// Composite ID = policyname,type,priority (key:UrlEncode(value) form)
		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"policyname", "type", "priority"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		argsMap := make(map[string]string)
		if val, ok := idMap["type"]; ok && val != "" {
			argsMap["type"] = val
		}

		findParams := service.FindParams{
			ResourceType:             service.Sslglobal_sslpolicy_binding.Type(),
			ArgsMap:                  argsMap,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if pn, ok := v["policyname"].(string); !ok || pn != idMap["policyname"] {
				continue
			}
			if pVal, ok := v["priority"]; ok {
				gotPriority, _ := utils.ConvertToInt64(pVal)
				wantPriority, _ := strconv.ParseInt(idMap["priority"], 10, 64)
				if gotPriority != wantPriority {
					continue
				}
			}
			found = true
			break
		}

		if !found {
			return fmt.Errorf("sslglobal_sslpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslglobalSslpolicyBindingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslglobal_sslpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No id is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"policyname", "type", "priority"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		argsMap := make(map[string]string)
		if val, ok := idMap["type"]; ok && val != "" {
			argsMap["type"] = val
		}

		findParams := service.FindParams{
			ResourceType:             service.Sslglobal_sslpolicy_binding.Type(),
			ArgsMap:                  argsMap,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Error (e.g. nothing bound) means the binding is gone.
			continue
		}

		for _, v := range dataArr {
			if pn, ok := v["policyname"].(string); ok && pn == idMap["policyname"] {
				if pVal, ok := v["priority"]; ok {
					gotPriority, _ := utils.ConvertToInt64(pVal)
					wantPriority, _ := strconv.ParseInt(idMap["priority"], 10, 64)
					if gotPriority == wantPriority {
						return fmt.Errorf("sslglobal_sslpolicy_binding %s still exists", rs.Primary.ID)
					}
				}
			}
		}
	}

	return nil
}

const testAccSslglobalSslpolicyBindingDataSource_basic = `
	resource "citrixadc_sslaction" "tf_sslaction" {
		name                   = "tf_sslaction_glb"
		clientauth             = "DOCLIENTAUTH"
		clientcertverification = "Mandatory"
	}

	resource "citrixadc_sslpolicy" "tf_sslpolicy" {
		name   = "tf_sslpolicy_glb"
		rule   = "true"
		action = citrixadc_sslaction.tf_sslaction.name
	}

	resource "citrixadc_sslglobal_sslpolicy_binding" "tf_sslglobal_sslpolicy_binding" {
		policyname = citrixadc_sslpolicy.tf_sslpolicy.name
		priority   = 100
		type       = "CONTROL_OVERRIDE"

		depends_on = [citrixadc_sslpolicy.tf_sslpolicy]
	}

	data "citrixadc_sslglobal_sslpolicy_binding" "tf_sslglobal_sslpolicy_binding" {
		policyname = citrixadc_sslglobal_sslpolicy_binding.tf_sslglobal_sslpolicy_binding.policyname
		priority   = 100
		type       = "CONTROL_OVERRIDE"

		depends_on = [citrixadc_sslglobal_sslpolicy_binding.tf_sslglobal_sslpolicy_binding]
	}
`

func TestAccSslglobalSslpolicyBindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslglobalSslpolicyBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslglobalSslpolicyBindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslglobal_sslpolicy_binding.tf_sslglobal_sslpolicy_binding", "policyname", "tf_sslpolicy_glb"),
					resource.TestCheckResourceAttr("data.citrixadc_sslglobal_sslpolicy_binding.tf_sslglobal_sslpolicy_binding", "priority", "100"),
					resource.TestCheckResourceAttr("data.citrixadc_sslglobal_sslpolicy_binding.tf_sslglobal_sslpolicy_binding", "type", "CONTROL_OVERRIDE"),
				),
			},
		},
	})
}
