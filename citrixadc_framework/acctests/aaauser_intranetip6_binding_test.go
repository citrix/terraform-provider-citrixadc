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
	"strings"
	"testing"
)

const testAccAaauser_intranetip6_binding_basic = `
	resource "citrixadc_aaauser" "tf_aaauser" {
		username = "user1"
		password = "my_pass"
	}
	resource "citrixadc_aaauser_intranetip6_binding" "tf_aaauser_intranetip6_binding" {
		username    = citrixadc_aaauser.tf_aaauser.username
		intranetip6 = "2003:db8:100::fb/128"
	}
`

const testAccAaauser_intranetip6_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_aaauser" "tf_aaauser" {
		username = "user1"
		password = "my_pass"
	}
`

func TestAccAaauser_intranetip6_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAaauser_intranetip6_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAaauser_intranetip6_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaauser_intranetip6_bindingExist("citrixadc_aaauser_intranetip6_binding.tf_aaauser_intranetip6_binding", nil),
				),
			},
			{
				Config: testAccAaauser_intranetip6_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaauser_intranetip6_bindingNotExist("citrixadc_aaauser_intranetip6_binding.tf_aaauser_intranetip6_binding", "user1,2003:db8:100::fb/128"),
				),
			},
		},
	})
}

func testAccCheckAaauser_intranetip6_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No aaauser_intranetip6_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 2)

		username := idSlice[0]
		intranetip6 := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "aaauser_intranetip6_binding",
			ResourceName:             username,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching intranetip6
		found := false
		for _, v := range dataArr {
			if v["intranetip6"].(string) == intranetip6 {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("aaauser_intranetip6_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAaauser_intranetip6_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		username := idSlice[0]
		intranetip6 := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "aaauser_intranetip6_binding",
			ResourceName:             username,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching intranetip6
		found := false
		for _, v := range dataArr {
			if v["intranetip6"].(string) == intranetip6 {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("aaauser_intranetip6_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAaauser_intranetip6_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_aaauser_intranetip6_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Aaauser_intranetip6_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("aaauser_intranetip6_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
