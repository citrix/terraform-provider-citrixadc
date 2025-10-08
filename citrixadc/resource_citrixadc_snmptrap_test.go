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
	"log"
	"strings"
	"testing"
)

const testAccSnmptrap_basic = `
	resource "citrixadc_snmptrap" "tf_snmptrap" {
		severity        = "Major"
		trapclass       = "specific"
		trapdestination = "192.168.2.2"
	}
`
const testAccSnmptrap_update = `
	resource "citrixadc_snmptrap" "tf_snmptrap" {
		severity        = "Minor"
		trapclass       = "specific"
		trapdestination = "192.168.2.2"
	}
`

func TestAccSnmptrap_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckSnmptrapDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmptrap_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmptrapExist("citrixadc_snmptrap.tf_snmptrap", nil),
					resource.TestCheckResourceAttr("citrixadc_snmptrap.tf_snmptrap", "trapclass", "specific"),
					resource.TestCheckResourceAttr("citrixadc_snmptrap.tf_snmptrap", "trapdestination", "192.168.2.2"),
					resource.TestCheckResourceAttr("citrixadc_snmptrap.tf_snmptrap", "severity", "Major"),
				),
			},
			{
				Config: testAccSnmptrap_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmptrapExist("citrixadc_snmptrap.tf_snmptrap", nil),
					resource.TestCheckResourceAttr("citrixadc_snmptrap.tf_snmptrap", "trapclass", "specific"),
					resource.TestCheckResourceAttr("citrixadc_snmptrap.tf_snmptrap", "trapdestination", "192.168.2.2"),
					resource.TestCheckResourceAttr("citrixadc_snmptrap.tf_snmptrap", "severity", "Minor"),
				),
			},
		},
	})
}

func testAccCheckSnmptrapExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No snmptrap name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		snmptrapId := rs.Primary.ID
		idSlice := strings.SplitN(snmptrapId, ",", 3)

		trapclass := idSlice[0]
		trapdestination := idSlice[1]
		version := idSlice[2]

		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		dataArr, err := client.FindAllResources(service.Snmptrap.Type())

		if err != nil {
			return err
		}

		if len(dataArr) == 0 {
			log.Printf("[WARN] citrix-provider: acceptance test: snmptrap does not exist. Clearing state.")
			return nil
		}

		// if len(dataArray) > 1 {
		// 	return fmt.Errorf("[ERROR] citrix-provider: acceptance test: multiple entries found for snmptrap")
		// }

		found := false
		for _, v := range dataArr {
			if v["trapclass"].(string) == trapclass && v["trapdestination"] == trapdestination && v["version"] == version {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("snmptrap %s not found", n)
		}

		return nil
	}
}

func testAccCheckSnmptrapDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_snmptrap" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}
		snmptrapId := rs.Primary.ID
		idSlice := strings.SplitN(snmptrapId, ",", 3)

		trapclass := idSlice[0]
		trapdestination := idSlice[1]
		version := idSlice[2]

		dataArr, err := client.FindAllResources(service.Snmptrap.Type())

		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if v["trapclass"].(string) == trapclass && v["trapdestination"] == trapdestination && v["version"] == version {
				found = true
				break
			}
		}
		if found {
			return fmt.Errorf("snmptrap %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
