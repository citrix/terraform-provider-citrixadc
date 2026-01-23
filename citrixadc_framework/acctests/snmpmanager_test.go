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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

const testAccSnmpmanager_basic = `

resource "citrixadc_snmpmanager" "tf_snmpmanager" {
	ipaddress          = "192.168.2.4"
	}
  	
`

const testAccSnmpmanager_update = `

resource "citrixadc_snmpmanager" "tf_snmpmanager" {
	ipaddress          = "192.168.2.4"
	netmask            = "255.255.255.252"
	}
	
`

func TestAccSnmpmanager_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmpmanagerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpmanager_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpmanagerExist("citrixadc_snmpmanager.tf_snmpmanager", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpmanager.tf_snmpmanager", "ipaddress", "192.168.2.4"),
					resource.TestCheckResourceAttr("citrixadc_snmpmanager.tf_snmpmanager", "netmask", "255.255.255.255"),
				),
			},
			{
				Config: testAccSnmpmanager_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpmanagerExist("citrixadc_snmpmanager.tf_snmpmanager", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpmanager.tf_snmpmanager", "ipaddress", "192.168.2.4"),
					resource.TestCheckResourceAttr("citrixadc_snmpmanager.tf_snmpmanager", "netmask", "255.255.255.252"),
				),
			},
		},
	})
}

func testAccCheckSnmpmanagerExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No snmpmanager name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		snmpmanagerName := rs.Primary.ID
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		dataArr, err := client.FindAllResources(service.Snmpmanager.Type())

		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if v["ipaddress"] == snmpmanagerName {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("snmpmanager %s not found", n)
		}

		return nil
	}
}

func testAccCheckSnmpmanagerDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_snmpmanager" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		snmpmanagerName := rs.Primary.ID

		dataArr, err := client.FindAllResources(service.Snmpmanager.Type())

		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if v["ipaddress"] == snmpmanagerName {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("snmpmanager %s still exists", snmpmanagerName)
		}

	}

	return nil
}
