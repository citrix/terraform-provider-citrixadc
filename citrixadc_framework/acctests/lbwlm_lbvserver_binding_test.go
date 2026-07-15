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

// lbwlm_lbvserver_binding joins a deprecated Work Load Manager (lbwlm, key wlmname)
// to an lbvserver (vservername). The participating-entity HCL is lifted from
// lbwlm_test.go and lbvserver_test.go. All binding attributes are RequiresReplace
// and there is no update path, so step1 creates the binding and step2 drops it to
// exercise unbind/destroy.

const testAccLbwlmLbvserverBinding_basic_step1 = `
resource "citrixadc_lbwlm" "tf_lbwlm" {
  wlmname   = "tf_lbwlm"
  lbuid     = "TF_LBWLM_UID"
  ipaddress = "10.222.74.128"
  port      = 3060
  katimeout = 2
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_lbvserver"
  servicetype = "HTTP"
  ipv46       = "10.202.11.11"
  port        = 80
}

resource "citrixadc_lbwlm_lbvserver_binding" "tf_binding" {
  wlmname     = citrixadc_lbwlm.tf_lbwlm.wlmname
  vservername = citrixadc_lbvserver.tf_lbvserver.name
  depends_on  = [citrixadc_lbwlm.tf_lbwlm, citrixadc_lbvserver.tf_lbvserver]
}

`

// step2 drops the binding (parents remain) to exercise the unbind/destroy path.
const testAccLbwlmLbvserverBinding_basic_step2 = `
resource "citrixadc_lbwlm" "tf_lbwlm" {
  wlmname   = "tf_lbwlm"
  lbuid     = "TF_LBWLM_UID"
  ipaddress = "10.222.74.128"
  port      = 3060
  katimeout = 2
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_lbvserver"
  servicetype = "HTTP"
  ipv46       = "10.202.11.11"
  port        = 80
}

`

func TestAccLbwlmLbvserverBinding_basic(t *testing.T) {
	t.Skip("TODO: Requires review.")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbwlmLbvserverBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbwlmLbvserverBinding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbwlmLbvserverBindingExist("citrixadc_lbwlm_lbvserver_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_lbwlm_lbvserver_binding.tf_binding", "wlmname", "tf_lbwlm"),
					resource.TestCheckResourceAttr("citrixadc_lbwlm_lbvserver_binding.tf_binding", "vservername", "tf_lbvserver"),
				),
			},
			{
				Config: testAccLbwlmLbvserverBinding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbwlmLbvserverBindingNotExist("citrixadc_lbwlm_lbvserver_binding.tf_binding", "tf_lbwlm", "tf_lbvserver"),
				),
			},
		},
	})
}

func TestAccLbwlmLbvserverBinding_import(t *testing.T) {
	t.Skip("TODO: Requires review.")
	const resAddr = "citrixadc_lbwlm_lbvserver_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbwlmLbvserverBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbwlmLbvserverBinding_basic_step1,
			},
			{
				Config:                  testAccLbwlmLbvserverBinding_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckLbwlmLbvserverBindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbwlm_lbvserver_binding id is set")
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
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		wlmname := idMap["wlmname"]
		vservername := idMap["vservername"]

		findParams := service.FindParams{
			ResourceType:             service.Lbwlm_lbvserver_binding.Type(),
			ResourceName:             wlmname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if val, ok := v["vservername"].(string); ok && val == vservername {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("lbwlm_lbvserver_binding %s not found", rs.Primary.ID)
		}

		return nil
	}
}

// testAccCheckLbwlmLbvserverBindingNotExist confirms the binding has been unbound
// while the parent lbwlm still exists (step2 drops only the binding).
func testAccCheckLbwlmLbvserverBindingNotExist(n string, wlmname string, vservername string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		findParams := service.FindParams{
			ResourceType:             service.Lbwlm_lbvserver_binding.Type(),
			ResourceName:             wlmname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// No bindings at all is an acceptable "not found" outcome.
			return nil
		}

		for _, v := range dataArr {
			if val, ok := v["vservername"].(string); ok && val == vservername {
				return fmt.Errorf("lbwlm_lbvserver_binding (%s,%s) still exists", wlmname, vservername)
			}
		}

		return nil
	}
}

func testAccCheckLbwlmLbvserverBindingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbwlm_lbvserver_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No id is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		wlmname := idMap["wlmname"]
		vservername := idMap["vservername"]

		findParams := service.FindParams{
			ResourceType:             service.Lbwlm_lbvserver_binding.Type(),
			ResourceName:             wlmname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Parent gone or no bindings - binding is destroyed.
			continue
		}

		for _, v := range dataArr {
			if val, ok := v["vservername"].(string); ok && val == vservername {
				return fmt.Errorf("lbwlm_lbvserver_binding %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

const testAccLbwlmLbvserverBindingDataSource_basic = `

resource "citrixadc_lbwlm" "tf_lbwlm" {
  wlmname   = "tf_lbwlm"
  lbuid     = "TF_LBWLM_UID"
  ipaddress = "10.222.74.128"
  port      = 3060
  katimeout = 2
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_lbvserver"
  servicetype = "HTTP"
  ipv46       = "10.202.11.11"
  port        = 80
}

resource "citrixadc_lbwlm_lbvserver_binding" "tf_binding" {
  wlmname     = citrixadc_lbwlm.tf_lbwlm.wlmname
  vservername = citrixadc_lbvserver.tf_lbvserver.name
  depends_on  = [citrixadc_lbwlm.tf_lbwlm, citrixadc_lbvserver.tf_lbvserver]
}

data "citrixadc_lbwlm_lbvserver_binding" "tf_binding_data" {
  wlmname     = citrixadc_lbwlm_lbvserver_binding.tf_binding.wlmname
  vservername = citrixadc_lbwlm_lbvserver_binding.tf_binding.vservername
  depends_on  = [citrixadc_lbwlm_lbvserver_binding.tf_binding]
}
`

func TestAccLbwlmLbvserverBindingDataSource_basic(t *testing.T) {
	t.Skip("TODO: Requires review.")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbwlmLbvserverBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbwlmLbvserverBindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lbwlm_lbvserver_binding.tf_binding_data", "wlmname", "tf_lbwlm"),
					resource.TestCheckResourceAttr("data.citrixadc_lbwlm_lbvserver_binding.tf_binding_data", "vservername", "tf_lbvserver"),
				),
			},
		},
	})
}
