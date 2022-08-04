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
	"log"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)
const testAccSnmpgroup_basic = `
	resource "citrixadc_snmpgroup" "tf_snmpgroup" {
	name          = "test_group"
	securitylevel = "noAuthNoPriv"
	readviewname  = "test_name"
	}
`
const testAccSnmpgroup_update = `
	resource "citrixadc_snmpgroup" "tf_snmpgroup" {
	name          = "test_group"
	securitylevel = "noAuthNoPriv"
	readviewname  = "test2_name"
	}
`

func TestAccSnmpgroup_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSnmpgroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSnmpgroup_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpgroupExist("citrixadc_snmpgroup.tf_snmpgroup", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpgroup.tf_snmpgroup", "securitylevel", "noAuthNoPriv" ),
					resource.TestCheckResourceAttr("citrixadc_snmpgroup.tf_snmpgroup", "readviewname", "test_name" ),

				),
			},
			resource.TestStep{
				Config: testAccSnmpgroup_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpgroupExist("citrixadc_snmpgroup.tf_snmpgroup", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpgroup.tf_snmpgroup", "securitylevel", "noAuthNoPriv" ),
					resource.TestCheckResourceAttr("citrixadc_snmpgroup.tf_snmpgroup", "readviewname", "test2_name" ),

				),
			},
		},
	})
}

func testAccCheckSnmpgroupExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No snmpgroup name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}
		snmpgroupName := rs.Primary.ID

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		dataArr, err := nsClient.FindAllResources(service.Snmpgroup.Type())

		if err != nil {
			return err
		}
		
		if len(dataArr) == 0 {
			log.Printf("[WARN] citrix-provider: acceptance test: snmpgroup does not exist. Clearing state.")
			return nil
		}
		
		found := false
		for _, v := range  dataArr {
			if v["name"] == snmpgroupName {
				found = true
				break
			}

		}
		if !found {
			return fmt.Errorf("snmpgroup %s not found", n)
		}

		return nil
	}
}

func testAccCheckSnmpgroupDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_snmpgroup" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		snmpgroupName := rs.Primary.ID

		dataArr, err := nsClient.FindAllResources(service.Snmpgroup.Type())
		
		if err != nil {
			return err
		}
		
		found := false
		for _, v := range  dataArr{
			if v["name"] == snmpgroupName {
				found = true
				break
			}
		
			if found {
				return fmt.Errorf("snmpgroup %s still exists", rs.Primary.ID)
			}

		}
	
	}
	return nil
}
