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
// The binding's participating entity is the appfwpolicy (and its appfwprofile
// prerequisite), reused from appfwpolicy_test.go.

const testAccVpnglobalAppfwpolicyBinding_basic_step1 = `
	resource citrixadc_appfwprofile tf_appfwprofile {
		name = "tf_appfwprofile"
	}

	resource citrixadc_appfwpolicy tf_appfwpolicy {
		name        = "tf_appfwpolicy"
		profilename = citrixadc_appfwprofile.tf_appfwprofile.name
		rule        = "true"
	}

	resource "citrixadc_vpnglobal_appfwpolicy_binding" "tf_binding" {
		policyname = citrixadc_appfwpolicy.tf_appfwpolicy.name
		priority   = 90
		secondary  = false

		depends_on = [citrixadc_appfwpolicy.tf_appfwpolicy]
	}
`

const testAccVpnglobalAppfwpolicyBinding_basic_step2 = `
	# Keep the participating entities without the actual binding to confirm proper deletion.
	resource citrixadc_appfwprofile tf_appfwprofile {
		name = "tf_appfwprofile"
		type = ["HTML"]
	}

	resource citrixadc_appfwpolicy tf_appfwpolicy {
		name        = "tf_appfwpolicy"
		profilename = citrixadc_appfwprofile.tf_appfwprofile.name
		rule        = "ns_true"
	}
`

func TestAccVpnglobalAppfwpolicyBinding_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnglobalAppfwpolicyBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnglobalAppfwpolicyBinding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnglobalAppfwpolicyBindingExist("citrixadc_vpnglobal_appfwpolicy_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnglobal_appfwpolicy_binding.tf_binding", "policyname", "tf_appfwpolicy"),
					resource.TestCheckResourceAttr("citrixadc_vpnglobal_appfwpolicy_binding.tf_binding", "priority", "90"),
				),
			},
			{
				Config: testAccVpnglobalAppfwpolicyBinding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnglobalAppfwpolicyBindingNotExist("citrixadc_vpnglobal_appfwpolicy_binding.tf_binding", "tf_appfwpolicy"),
				),
			},
		},
	})
}

func testAccCheckVpnglobalAppfwpolicyBindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnglobal_appfwpolicy_binding id is set")
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

		// ID is a plain value (single unique attr: policyname)
		policyname := rs.Primary.ID

		findParams := service.FindParams{
			ResourceType:             service.Vpnglobal_appfwpolicy_binding.Type(),
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching policyname
		found := false
		for _, v := range dataArr {
			if val, ok := v["policyname"].(string); ok && val == policyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("vpnglobal_appfwpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnglobalAppfwpolicyBindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		policyname := id

		findParams := service.FindParams{
			ResourceType:             service.Vpnglobal_appfwpolicy_binding.Type(),
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching policyname
		found := false
		for _, v := range dataArr {
			if val, ok := v["policyname"].(string); ok && val == policyname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("vpnglobal_appfwpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckVpnglobalAppfwpolicyBindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnglobal_appfwpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		policyname := rs.Primary.ID

		findParams := service.FindParams{
			ResourceType:             service.Vpnglobal_appfwpolicy_binding.Type(),
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// A missing-resource error means the binding is gone, which is what we want.
			continue
		}

		for _, v := range dataArr {
			if val, ok := v["policyname"].(string); ok && val == policyname {
				return fmt.Errorf("vpnglobal_appfwpolicy_binding %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

const testAccVpnglobalAppfwpolicyBindingDataSource_basic = `
	resource citrixadc_appfwprofile tf_appfwprofile {
		name = "tf_appfwprofile"
		type = ["HTML"]
	}

	resource citrixadc_appfwpolicy tf_appfwpolicy {
		name        = "tf_appfwpolicy"
		profilename = citrixadc_appfwprofile.tf_appfwprofile.name
		rule        = "ns_true"
	}

	resource "citrixadc_vpnglobal_appfwpolicy_binding" "tf_binding" {
		policyname = citrixadc_appfwpolicy.tf_appfwpolicy.name
		priority   = 90
		secondary  = false

		depends_on = [citrixadc_appfwpolicy.tf_appfwpolicy]
	}

	data "citrixadc_vpnglobal_appfwpolicy_binding" "tf_binding" {
		policyname = citrixadc_vpnglobal_appfwpolicy_binding.tf_binding.policyname
		depends_on = [citrixadc_vpnglobal_appfwpolicy_binding.tf_binding]
	}
`

func TestAccVpnglobalAppfwpolicyBindingDataSource_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnglobalAppfwpolicyBindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnglobal_appfwpolicy_binding.tf_binding", "policyname", "tf_appfwpolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_vpnglobal_appfwpolicy_binding.tf_binding", "priority", "90"),
				),
			},
		},
	})
}
