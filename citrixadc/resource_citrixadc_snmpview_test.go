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
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

const testAccSnmpview_basic = `


resource "citrixadc_snmpview" "tf_snmpview" {
	name    = "test_name"
	subtree = "1.2.4.7"
	type    = "excluded"
  }
  
`
const testAccSnmpview_update = `


resource "citrixadc_snmpview" "tf_snmpview" {
	name    = "test_name"
	subtree = "1.2.4.8"
	type    = "included"
  }
  
`

func TestAccSnmpview_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSnmpviewDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpview_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpviewExist("citrixadc_snmpview.tf_snmpview", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpview.tf_snmpview", "subtree", "1.2.4.7"),
					resource.TestCheckResourceAttr("citrixadc_snmpview.tf_snmpview", "type", "excluded"),
				),
			},
			{
				Config: testAccSnmpview_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpviewExist("citrixadc_snmpview.tf_snmpview", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpview.tf_snmpview", "subtree", "1.2.4.8"),
					resource.TestCheckResourceAttr("citrixadc_snmpview.tf_snmpview", "type", "included"),
				),
			},
		},
	})
}

func testAccCheckSnmpviewExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No snmpview name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		dataArr, err := nsClient.FindAllResources(service.Snmpview.Type())

		snmpviewName := rs.Primary.ID
		// Unexpected error
		if err != nil {
			return err
		}
		if err != nil {
			return err
		}
		found := false
		for _, v := range dataArr {
			if v["name"].(string) == snmpviewName {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("gslbservicegroup %s not found", n)
		}

		return nil
	}
}

func testAccCheckSnmpviewDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_snmpview" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Snmpview.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("snmpview %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
