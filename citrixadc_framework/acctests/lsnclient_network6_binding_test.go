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

const testAccLsnclient_network6_binding_basic = `

resource "citrixadc_lsnclient_network6_binding" "tf_lsnclient_network6_binding" {
	clientname = "my_lsn_client"
	network6    = "2001:db8:5001::/96"
	}
`

const testAccLsnclient_network6_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
`

func TestAccLsnclient_network6_binding_basic(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this LSN resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnclient_network6_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnclient_network6_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnclient_network6_bindingExist("citrixadc_lsnclient_network6_binding.tf_lsnclient_network6_binding", nil),
				),
			},
			{
				Config: testAccLsnclient_network6_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnclient_network6_bindingNotExist("citrixadc_lsnclient_network6_binding.tf_lsnclient_network6_binding", "my_lsn_client,2001:db8:5001::/96"),
				),
			},
		},
	})
}

func testAccCheckLsnclient_network6_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsnclient_network6_binding id is set")
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

		clientname := idSlice[0]
		network6 := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "lsnclient_network6_binding",
			ResourceName:             clientname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching network6
		found := false
		for _, v := range dataArr {
			if strings.ToLower(v["network6"].(string)) == strings.ToLower(network6) {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("lsnclient_network6_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsnclient_network6_bindingNotExist(n string, id string) resource.TestCheckFunc {
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

		clientname := idSlice[0]
		network6 := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "lsnclient_network6_binding",
			ResourceName:             clientname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching network6
		found := false
		for _, v := range dataArr {
			if strings.ToLower(v["network6"].(string)) == strings.ToLower(network6) {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("lsnclient_network6_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckLsnclient_network6_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsnclient_network6_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lsnclient_network6_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lsnclient_network6_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
