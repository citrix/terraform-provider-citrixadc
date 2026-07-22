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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// rnat_retainsourceportset_binding is an immutable bind/unbind resource (no update
// action). Both attributes are RequiresReplace, so the basic test has a single bind
// step plus a step2 that drops the binding to verify unbind/CheckDestroy.
//
// Prerequisites: a parent RNAT rule must exist (referenced by `name`). The parent
// rnat config below is lifted from rnat_test.go / rnat_nsip_binding_test.go.
//
// TODO_PLACEHOLDER note: `retainsourceportrange` uses "1024-2048" (a well-formed
// int-range). Adjust to a valid, non-conflicting port range for your testbed if
// the ADC rejects it (e.g. if the range overlaps an existing binding).

const testAccRnat_retainsourceportset_binding_basic = `

	resource "citrixadc_rnat" "tfrnat" {
		name             = "my_rnat_rsps"
		network          = "10.2.2.0"
		netmask          = "255.255.255.255"
		useproxyport     = "DISABLED"
		srcippersistency = "DISABLED"
		connfailover     = "DISABLED"
	}
	resource "citrixadc_rnat_retainsourceportset_binding" "tf_rnat_retainsourceportset_binding" {
		name                  = citrixadc_rnat.tfrnat.name
		retainsourceportrange = "1024-2048"
	}

`

const testAccRnat_retainsourceportset_binding_basic_step2 = `
	# Keep the parent rnat rule without the actual binding to check proper deletion

	resource "citrixadc_rnat" "tfrnat" {
		name             = "my_rnat_rsps"
		network          = "10.2.2.0"
		netmask          = "255.255.255.255"
		useproxyport     = "DISABLED"
		srcippersistency = "DISABLED"
		connfailover     = "DISABLED"
	}
`

func TestAccRnat_retainsourceportset_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRnat_retainsourceportset_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRnat_retainsourceportset_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRnat_retainsourceportset_bindingExist("citrixadc_rnat_retainsourceportset_binding.tf_rnat_retainsourceportset_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_rnat_retainsourceportset_binding.tf_rnat_retainsourceportset_binding", "name", "my_rnat_rsps"),
					resource.TestCheckResourceAttr("citrixadc_rnat_retainsourceportset_binding.tf_rnat_retainsourceportset_binding", "retainsourceportrange", "1024-2048"),
				),
			},
			{
				Config: testAccRnat_retainsourceportset_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRnat_retainsourceportset_bindingNotExist("citrixadc_rnat_retainsourceportset_binding.tf_rnat_retainsourceportset_binding", "my_rnat_rsps", "1024-2048"),
				),
			},
		},
	})
}

func TestAccRnat_retainsourceportset_binding_import(t *testing.T) {
	const resAddr = "citrixadc_rnat_retainsourceportset_binding.tf_rnat_retainsourceportset_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRnat_retainsourceportset_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRnat_retainsourceportset_binding_basic,
			},
			{
				Config:                  testAccRnat_retainsourceportset_binding_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckRnat_retainsourceportset_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No rnat_retainsourceportset_binding id is set")
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

		// ID format: name:<enc>,retainsourceportrange:<enc>
		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}
		name := idMap["name"]
		retainsourceportrange := idMap["retainsourceportrange"]

		findParams := service.FindParams{
			ResourceType:             service.Rnat_retainsourceportset_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching retainsourceportrange
		found := false
		for _, v := range dataArr {
			if val, ok := v["retainsourceportrange"].(string); ok && val == retainsourceportrange {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("rnat_retainsourceportset_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckRnat_retainsourceportset_bindingNotExist(n string, name string, retainsourceportrange string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		findParams := service.FindParams{
			ResourceType:             service.Rnat_retainsourceportset_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the matching retainsourceportrange
		found := false
		for _, v := range dataArr {
			if val, ok := v["retainsourceportrange"].(string); ok && val == retainsourceportrange {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("rnat_retainsourceportset_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckRnat_retainsourceportset_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_rnat_retainsourceportset_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}
		name := idMap["name"]
		retainsourceportrange := idMap["retainsourceportrange"]

		findParams := service.FindParams{
			ResourceType:             service.Rnat_retainsourceportset_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Parent gone / no bindings - treated as destroyed
			continue
		}

		for _, v := range dataArr {
			if val, ok := v["retainsourceportrange"].(string); ok && val == retainsourceportrange {
				return fmt.Errorf("rnat_retainsourceportset_binding %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

const testAccRnat_retainsourceportset_bindingDataSource_basic = `

	resource "citrixadc_rnat" "tfrnat" {
		name             = "my_rnat_rsps"
		network          = "10.2.2.0"
		netmask          = "255.255.255.255"
		useproxyport     = "DISABLED"
		srcippersistency = "DISABLED"
		connfailover     = "DISABLED"
	}
	resource "citrixadc_rnat_retainsourceportset_binding" "tf_rnat_retainsourceportset_binding" {
		name                  = citrixadc_rnat.tfrnat.name
		retainsourceportrange = "1024-2048"
	}

	data "citrixadc_rnat_retainsourceportset_binding" "tf_rnat_retainsourceportset_binding" {
		name                  = citrixadc_rnat_retainsourceportset_binding.tf_rnat_retainsourceportset_binding.name
		retainsourceportrange = citrixadc_rnat_retainsourceportset_binding.tf_rnat_retainsourceportset_binding.retainsourceportrange
		depends_on            = [citrixadc_rnat_retainsourceportset_binding.tf_rnat_retainsourceportset_binding]
	}
`

func TestAccRnat_retainsourceportset_bindingDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccRnat_retainsourceportset_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_rnat_retainsourceportset_binding.tf_rnat_retainsourceportset_binding", "name", "my_rnat_rsps"),
					resource.TestCheckResourceAttr("data.citrixadc_rnat_retainsourceportset_binding.tf_rnat_retainsourceportset_binding", "retainsourceportrange", "1024-2048"),
				),
			},
		},
	})
}
