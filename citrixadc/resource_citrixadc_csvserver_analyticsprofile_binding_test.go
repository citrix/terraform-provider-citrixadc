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
	"strings"
	"testing"
)

const testAccCsvserver_analyticsprofile_binding_basic = `
	resource "citrixadc_csvserver_analyticsprofile_binding" "tf_csvserver_analyticsprofile_binding" {
        name = citrixadc_csvserver.tf_csvserver.name
		analyticsprofile = "ns_analytics_global_profile"
	}

	resource "citrixadc_csvserver" "tf_csvserver" {
		name = "tf_csvserver"
		ipv46 = "1.1.1.2"
		port = 80
		servicetype = "HTTP"
	}
`

const testAccCsvserver_analyticsprofile_binding_basic_step2 = `
	resource "citrixadc_csvserver" "tf_csvserver" {
		name = "tf_csvserver"
		ipv46 = "1.1.1.2"
		port = 80
		servicetype = "HTTP"
	}
`

func TestAccCsvserver_analyticsprofile_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCsvserver_analyticsprofile_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCsvserver_analyticsprofile_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_analyticsprofile_bindingExist("citrixadc_csvserver_analyticsprofile_binding.tf_csvserver_analyticsprofile_binding", nil),
				),
			},
			resource.TestStep{
				Config: testAccCsvserver_analyticsprofile_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_analyticsprofile_bindingNotExist("citrixadc_csvserver_analyticsprofile_binding.tf_csvserver_analyticsprofile_binding", "tf_csvserver,ns_analytics_global_profile"),
				),
			},
		},
	})
}

func testAccCheckCsvserver_analyticsprofile_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No csvserver_analyticsprofile_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 2)

		name := idSlice[0]
		analyticsprofile := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "csvserver_analyticsprofile_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right monitor name
		found := false
		for _, v := range dataArr {
			if v["analyticsprofile"].(string) == analyticsprofile {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("csvserver_analyticsprofile_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckCsvserver_analyticsprofile_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		name := idSlice[0]
		analyticsprofile := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "csvserver_analyticsprofile_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right profile name
		found := false
		for _, v := range dataArr {
			if v["analyticsprofile"].(string) == analyticsprofile {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("csvserver_analyticsprofile_binding %s was found, but is should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckCsvserver_analyticsprofile_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_csvserver_analyticsprofile_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("csvserver_analyticsprofile_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("csvserver_analyticsprofile_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
