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
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"fmt"
	"strings"
	"testing"
)

func TestAccCsvserver_cspolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCsvserver_cspolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_cspolicy_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_cspolicy_bindingExist("citrixadc_csvserver_cspolicy_binding.tf_csvscspolbind", nil),
					testAccCheckCsvserver_cspolicy_bindingExist("citrixadc_csvserver_cspolicy_binding.tf_csvscspolbind_extra", nil),
				),
			},
			{
				Config: testAccCsvserver_cspolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_cspolicy_bindingExist("citrixadc_csvserver_cspolicy_binding.tf_csvscspolbind", nil),
				),
			},
		},
	})
}

func testAccCheckCsvserver_cspolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lb vserver name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		bindingId := rs.Primary.ID
		idSlice := strings.Split(bindingId, ",")
		name := idSlice[0]
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "csvserver_cspolicy_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}

		dataArr, err := client.FindResourceArrayWithParams(findParams)

		if err != nil {
			return err
		}

		found := false

		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckCsvserver_cspolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_csvserver_cspolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		bindingId := rs.Primary.ID
		idSlice := strings.Split(bindingId, ",")
		name := idSlice[0]
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "csvserver_cspolicy_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}

		dataArr, err := client.FindResourceArrayWithParams(findParams)

		if err != nil {
			return err
		}

		found := false

		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("binding %s still exists", bindingId)
		}

	}

	return nil
}

const testAccCsvserver_cspolicy_binding_basic_step1 = `
resource "citrixadc_csvserver" "tf_csvserver" {
  ipv46       = "10.10.10.22"
  name        = "tf_csvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_cspolicy" "tf_cspolicy" {
  policyname      = "tf_cspolicy"
  rule            = "CLIENT.IP.SRC.SUBNET(24).EQ(10.217.85.0)"
}

resource "citrixadc_csvserver_cspolicy_binding" "tf_csvscspolbind" {
    name = citrixadc_csvserver.tf_csvserver.name
    policyname = citrixadc_cspolicy.tf_cspolicy.policyname
    priority = 100
    targetlbvserver = citrixadc_lbvserver.tf_lbvserver.name
}

resource "citrixadc_cspolicy" "tf_cspolicy_extra" {
  policyname      = "tf_cspolicy_extra"
  rule            = "CLIENT.IP.SRC.SUBNET(24).EQ(10.217.86.0)"
}

resource "citrixadc_csvserver_cspolicy_binding" "tf_csvscspolbind_extra" {
    name = citrixadc_csvserver.tf_csvserver.name
    policyname = citrixadc_cspolicy.tf_cspolicy_extra.policyname
    priority = 110
    targetlbvserver = citrixadc_lbvserver.tf_lbvserver.name
}
`

const testAccCsvserver_cspolicy_binding_basic_step2 = `
resource "citrixadc_csvserver" "tf_csvserver" {
  ipv46       = "10.10.10.22"
  name        = "tf_csvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_cspolicy" "tf_cspolicy" {
  policyname      = "tf_cspolicy"
  rule            = "CLIENT.IP.SRC.SUBNET(24).EQ(10.217.85.0)"
}

resource "citrixadc_csvserver_cspolicy_binding" "tf_csvscspolbind" {
    name = citrixadc_csvserver.tf_csvserver.name
    policyname = citrixadc_cspolicy.tf_cspolicy.policyname
    priority = 100
    targetlbvserver = citrixadc_lbvserver.tf_lbvserver.name
}

resource "citrixadc_cspolicy" "tf_cspolicy_extra" {
  policyname      = "tf_cspolicy_extra"
  rule            = "CLIENT.IP.SRC.SUBNET(24).EQ(10.217.86.0)"
}

`
