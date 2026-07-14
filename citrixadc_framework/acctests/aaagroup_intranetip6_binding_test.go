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

// IPv6 note: intranetip6 contains literal ':' characters, so the composite ID is
// "groupname:<v>,intranetip6:<v>,numaddr:<v>" with each value URL-encoded ('%3A').
// The check functions below parse the ID with utils.ParseIdString (which returns
// the intranetip6 value URL-DECODED, raw ':' restored) and then filter the binding
// array on the raw intranetip6 value AND numaddr, exactly as the resource Read does.

const testAccAaagroup_intranetip6_binding_basic_step1 = `

	resource "citrixadc_aaagroup" "tf_aaagroup" {
		groupname = "my_group"
		weight    = 100
	}
	resource "citrixadc_aaagroup_intranetip6_binding" "tf_aaagroup_intranetip6_binding" {
		groupname   = citrixadc_aaagroup.tf_aaagroup.groupname
		intranetip6 = "2001:db8::1"
		numaddr     = 1
		depends_on  = [citrixadc_aaagroup.tf_aaagroup]
	}

`

const testAccAaagroup_intranetip6_binding_basic_step2 = `
	# Keep the parent aaagroup without the actual binding to verify proper deletion
	resource "citrixadc_aaagroup" "tf_aaagroup" {
		groupname = "my_group"
		weight    = 100
	}
`

const testAccAaagroupIntranetip6BindingDataSource_basic = `

	resource "citrixadc_aaagroup" "tf_aaagroup" {
		groupname = "my_group"
		weight    = 100
	}
	resource "citrixadc_aaagroup_intranetip6_binding" "tf_aaagroup_intranetip6_binding" {
		groupname   = citrixadc_aaagroup.tf_aaagroup.groupname
		intranetip6 = "2001:db8::1"
		numaddr     = 1
		depends_on  = [citrixadc_aaagroup.tf_aaagroup]
	}

	data "citrixadc_aaagroup_intranetip6_binding" "tf_aaagroup_intranetip6_binding" {
		groupname   = citrixadc_aaagroup_intranetip6_binding.tf_aaagroup_intranetip6_binding.groupname
		intranetip6 = citrixadc_aaagroup_intranetip6_binding.tf_aaagroup_intranetip6_binding.intranetip6
		depends_on  = [citrixadc_aaagroup_intranetip6_binding.tf_aaagroup_intranetip6_binding]
	}
`

func TestAccAaagroup_intranetip6_binding_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAaagroup_intranetip6_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAaagroup_intranetip6_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaagroup_intranetip6_bindingExist("citrixadc_aaagroup_intranetip6_binding.tf_aaagroup_intranetip6_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_aaagroup_intranetip6_binding.tf_aaagroup_intranetip6_binding", "groupname", "my_group"),
					resource.TestCheckResourceAttr("citrixadc_aaagroup_intranetip6_binding.tf_aaagroup_intranetip6_binding", "intranetip6", "2001:db8::1"),
					resource.TestCheckResourceAttr("citrixadc_aaagroup_intranetip6_binding.tf_aaagroup_intranetip6_binding", "numaddr", "1"),
				),
			},
			{
				// Binding dropped; the parent aaagroup remains. Verify the binding was deleted.
				Config: testAccAaagroup_intranetip6_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaagroup_intranetip6_bindingNotExist(
						"citrixadc_aaagroup_intranetip6_binding.tf_aaagroup_intranetip6_binding",
						"my_group", "2001:db8::1", 1),
				),
			},
		},
	})
}

func TestAccAaagroupIntranetip6BindingDataSource_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAaagroupIntranetip6BindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_aaagroup_intranetip6_binding.tf_aaagroup_intranetip6_binding", "groupname", "my_group"),
					resource.TestCheckResourceAttr("data.citrixadc_aaagroup_intranetip6_binding.tf_aaagroup_intranetip6_binding", "intranetip6", "2001:db8::1"),
				),
			},
		},
	})
}

func testAccCheckAaagroup_intranetip6_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No aaagroup_intranetip6_binding id is set")
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

		// ParseIdString returns intranetip6 URL-decoded (raw ':' restored) so it
		// matches the API value directly. numaddr is returned as its decoded string.
		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		groupname := idMap["groupname"]
		intranetip6 := idMap["intranetip6"]
		numaddr := idMap["numaddr"]

		findParams := service.FindParams{
			ResourceType:             service.Aaagroup_intranetip6_binding.Type(),
			ResourceName:             groupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		// Filter on both intranetip6 AND numaddr, matching the resource Read.
		found := false
		for _, v := range dataArr {
			val, ok := v["intranetip6"].(string)
			if !ok || val != intranetip6 {
				continue
			}
			if numaddr != "" {
				apiNum, apiErr := utils.ConvertToInt64(v["numaddr"])
				idNum, idErr := utils.ConvertToInt64(numaddr)
				if apiErr != nil || idErr != nil || apiNum != idNum {
					continue
				}
			}
			found = true
			break
		}

		if !found {
			return fmt.Errorf("aaagroup_intranetip6_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAaagroup_intranetip6_bindingNotExist(n string, groupname string, intranetip6 string, numaddr int64) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		findParams := service.FindParams{
			ResourceType:             service.Aaagroup_intranetip6_binding.Type(),
			ResourceName:             groupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		// Hopefully NOT find a binding matching intranetip6 AND numaddr.
		found := false
		for _, v := range dataArr {
			val, ok := v["intranetip6"].(string)
			if !ok || val != intranetip6 {
				continue
			}
			apiNum, apiErr := utils.ConvertToInt64(v["numaddr"])
			if apiErr != nil || apiNum != numaddr {
				continue
			}
			found = true
			break
		}

		if found {
			return fmt.Errorf("aaagroup_intranetip6_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAaagroup_intranetip6_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_aaagroup_intranetip6_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		// Parse the composite ID (intranetip6 is URL-decoded by ParseIdString).
		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}

		groupname := idMap["groupname"]
		intranetip6 := idMap["intranetip6"]
		numaddr := idMap["numaddr"]

		findParams := service.FindParams{
			ResourceType:             service.Aaagroup_intranetip6_binding.Type(),
			ResourceName:             groupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		// Parent group gone or no bindings - treated as destroyed.
		if err != nil {
			continue
		}

		for _, v := range dataArr {
			val, ok := v["intranetip6"].(string)
			if !ok || val != intranetip6 {
				continue
			}
			if numaddr != "" {
				apiNum, apiErr := utils.ConvertToInt64(v["numaddr"])
				idNum, idErr := utils.ConvertToInt64(numaddr)
				if apiErr != nil || idErr != nil || apiNum != idNum {
					continue
				}
			}
			return fmt.Errorf("aaagroup_intranetip6_binding %s still exists", rs.Primary.ID)
		}
	}

	return nil
}
