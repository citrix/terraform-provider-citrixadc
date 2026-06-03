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
)

// vpnglobal is a singleton on the ADC, so there is no parent resource to create.
// secureprivateaccessurl is the binding key itself (a literal URL string, max
// length 255) and not a reference to another resource, so the config is just the
// binding resource with a literal URL value.

const testAccVpnglobalSecureprivateaccessurlBinding_basic_step1 = `
	resource "citrixadc_vpnglobal_secureprivateaccessurl_binding" "tf_binding" {
		secureprivateaccessurl = "https://app.example.com/"
	}
`

const testAccVpnglobalSecureprivateaccessurlBinding_basic_step2 = `
	# Drop the binding to confirm proper deletion. vpnglobal is a singleton with no
	# participating entity to retain, so the config is empty.
`

func TestAccVpnglobalSecureprivateaccessurlBinding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnglobalSecureprivateaccessurlBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnglobalSecureprivateaccessurlBinding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnglobalSecureprivateaccessurlBindingExist("citrixadc_vpnglobal_secureprivateaccessurl_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnglobal_secureprivateaccessurl_binding.tf_binding", "secureprivateaccessurl", "https://app.example.com/"),
				),
			},
			{
				Config: testAccVpnglobalSecureprivateaccessurlBinding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnglobalSecureprivateaccessurlBindingNotExist("citrixadc_vpnglobal_secureprivateaccessurl_binding.tf_binding", "https://app.example.com/"),
				),
			},
		},
	})
}

func testAccCheckVpnglobalSecureprivateaccessurlBindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnglobal_secureprivateaccessurl_binding id is set")
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

		// ID is a plain value (single unique attr: secureprivateaccessurl)
		secureprivateaccessurl := rs.Primary.ID

		findParams := service.FindParams{
			ResourceType:             service.Vpnglobal_secureprivateaccessurl_binding.Type(),
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching secureprivateaccessurl
		found := false
		for _, v := range dataArr {
			if val, ok := v["secureprivateaccessurl"].(string); ok && val == secureprivateaccessurl {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("vpnglobal_secureprivateaccessurl_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnglobalSecureprivateaccessurlBindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		secureprivateaccessurl := id

		findParams := service.FindParams{
			ResourceType:             service.Vpnglobal_secureprivateaccessurl_binding.Type(),
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// A missing-resource error means the binding is gone, which is what we want.
		if err != nil {
			return nil
		}

		// Iterate through results to hopefully not find the one with the matching secureprivateaccessurl
		found := false
		for _, v := range dataArr {
			if val, ok := v["secureprivateaccessurl"].(string); ok && val == secureprivateaccessurl {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("vpnglobal_secureprivateaccessurl_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckVpnglobalSecureprivateaccessurlBindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnglobal_secureprivateaccessurl_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		secureprivateaccessurl := rs.Primary.ID

		findParams := service.FindParams{
			ResourceType:             service.Vpnglobal_secureprivateaccessurl_binding.Type(),
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// A missing-resource error means the binding is gone, which is what we want.
			continue
		}

		for _, v := range dataArr {
			if val, ok := v["secureprivateaccessurl"].(string); ok && val == secureprivateaccessurl {
				return fmt.Errorf("vpnglobal_secureprivateaccessurl_binding %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

const testAccVpnglobalSecureprivateaccessurlBindingDataSource_basic = `
	resource "citrixadc_vpnglobal_secureprivateaccessurl_binding" "tf_binding" {
		secureprivateaccessurl = "https://app.example.com/"
	}

	data "citrixadc_vpnglobal_secureprivateaccessurl_binding" "tf_binding" {
		secureprivateaccessurl = citrixadc_vpnglobal_secureprivateaccessurl_binding.tf_binding.secureprivateaccessurl
		depends_on             = [citrixadc_vpnglobal_secureprivateaccessurl_binding.tf_binding]
	}
`

func TestAccVpnglobalSecureprivateaccessurlBindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnglobalSecureprivateaccessurlBindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnglobal_secureprivateaccessurl_binding.tf_binding", "secureprivateaccessurl", "https://app.example.com/"),
				),
			},
		},
	})
}
