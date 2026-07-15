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

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Participating entities are created from the working HCL of their own acceptance
// tests:
//   - lbvserver        (reused from lbvserver_test.go: name + servicetype HTTP + ipv46 + port)
//   - auditnslogpolicy (reused from auditnslogpolicy_test.go: name + rule + built-in
//     action SETASLEARNNSLOG_ACT, so no chained auditnslogaction/auditsyslog is needed)

const testAccLbvserver_auditnslogpolicy_binding_basic_step1 = `

resource "citrixadc_lbvserver" "tf_lbvserver" {
	name        = "tf_lbvserver_audit"
	servicetype = "HTTP"
	ipv46       = "10.202.11.12"
	port        = 80
}

resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
	name   = "tf_auditnslogpolicy_bind"
	rule   = "true"
	action = "SETASLEARNNSLOG_ACT"
}

resource "citrixadc_lbvserver_auditnslogpolicy_binding" "tf_binding" {
	name       = citrixadc_lbvserver.tf_lbvserver.name
	policyname = citrixadc_auditnslogpolicy.tf_auditnslogpolicy.name
	priority   = 100

	depends_on = [
		citrixadc_lbvserver.tf_lbvserver,
		citrixadc_auditnslogpolicy.tf_auditnslogpolicy,
	]
}
`

// Step 2 drops the binding (keeps the participating entities) to verify deletion.
const testAccLbvserver_auditnslogpolicy_binding_basic_step2 = `

resource "citrixadc_lbvserver" "tf_lbvserver" {
	name        = "tf_lbvserver_audit"
	servicetype = "HTTP"
	ipv46       = "10.202.11.12"
	port        = 80
}

resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
	name   = "tf_auditnslogpolicy_bind"
	rule   = "true"
	action = "SETASLEARNNSLOG_ACT"
}
`

func TestAccLbvserver_auditnslogpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbvserver_auditnslogpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbvserver_auditnslogpolicy_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_auditnslogpolicy_bindingExist("citrixadc_lbvserver_auditnslogpolicy_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_lbvserver_auditnslogpolicy_binding.tf_binding", "name", "tf_lbvserver_audit"),
					resource.TestCheckResourceAttr("citrixadc_lbvserver_auditnslogpolicy_binding.tf_binding", "policyname", "tf_auditnslogpolicy_bind"),
					resource.TestCheckResourceAttr("citrixadc_lbvserver_auditnslogpolicy_binding.tf_binding", "priority", "100"),
				),
			},
			{
				Config: testAccLbvserver_auditnslogpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_auditnslogpolicy_bindingNotExist("citrixadc_lbvserver", "tf_lbvserver_audit", "tf_auditnslogpolicy_bind"),
				),
			},
		},
	})
}

func TestAccLbvserver_auditnslogpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_lbvserver_auditnslogpolicy_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbvserver_auditnslogpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbvserver_auditnslogpolicy_binding_basic_step1,
			},
			{
				Config:                  testAccLbvserver_auditnslogpolicy_binding_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckLbvserver_auditnslogpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbvserver_auditnslogpolicy_binding id is set")
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

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"name", "policyname"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}
		name := idMap["name"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             service.Lbvserver_auditnslogpolicy_binding.Type(),
			ResourceName:             name,
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
			return fmt.Errorf("lbvserver_auditnslogpolicy_binding %s not found", n)
		}

		return nil
	}
}

// testAccCheckLbvserver_auditnslogpolicy_bindingNotExist verifies the binding was
// removed while the parent lbvserver still exists (step 2 drops only the binding).
func testAccCheckLbvserver_auditnslogpolicy_bindingNotExist(_ string, name string, policyname string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		findParams := service.FindParams{
			ResourceType:             service.Lbvserver_auditnslogpolicy_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		for _, v := range dataArr {
			if val, ok := v["policyname"].(string); ok && val == policyname {
				return fmt.Errorf("lbvserver_auditnslogpolicy_binding still exists for %s/%s", name, policyname)
			}
		}

		return nil
	}
}

func testAccCheckLbvserver_auditnslogpolicy_bindingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbvserver_auditnslogpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No id is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, []string{"name", "policyname"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}
		name := idMap["name"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             service.Lbvserver_auditnslogpolicy_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Parent gone / no bindings - binding is destroyed
			continue
		}

		for _, v := range dataArr {
			if val, ok := v["policyname"].(string); ok && val == policyname {
				return fmt.Errorf("lbvserver_auditnslogpolicy_binding %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

const testAccLbvserver_auditnslogpolicy_bindingDataSource_basic = `

resource "citrixadc_lbvserver" "tf_lbvserver" {
	name        = "tf_lbvserver_audit"
	servicetype = "HTTP"
	ipv46       = "10.202.11.12"
	port        = 80
}

resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
	name   = "tf_auditnslogpolicy_bind"
	rule   = "true"
	action = "SETASLEARNNSLOG_ACT"
}

resource "citrixadc_lbvserver_auditnslogpolicy_binding" "tf_binding" {
	name       = citrixadc_lbvserver.tf_lbvserver.name
	policyname = citrixadc_auditnslogpolicy.tf_auditnslogpolicy.name
	priority   = 100

	depends_on = [
		citrixadc_lbvserver.tf_lbvserver,
		citrixadc_auditnslogpolicy.tf_auditnslogpolicy,
	]
}

data "citrixadc_lbvserver_auditnslogpolicy_binding" "tf_binding" {
	name       = citrixadc_lbvserver_auditnslogpolicy_binding.tf_binding.name
	policyname = citrixadc_lbvserver_auditnslogpolicy_binding.tf_binding.policyname
	depends_on = [citrixadc_lbvserver_auditnslogpolicy_binding.tf_binding]
}
`

func TestAccLbvserver_auditnslogpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbvserver_auditnslogpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbvserver_auditnslogpolicy_bindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_auditnslogpolicy_binding.tf_binding", "name", "tf_lbvserver_audit"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_auditnslogpolicy_binding.tf_binding", "policyname", "tf_auditnslogpolicy_bind"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_auditnslogpolicy_binding.tf_binding", "priority", "100"),
				),
			},
		},
	})
}
