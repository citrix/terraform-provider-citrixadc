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

const testAccLsngroup_lsntransportprofile_binding_basic = `

resource "citrixadc_lsngroup_lsntransportprofile_binding" "tf_lsngroup_lsntransportprofile_binding" {
	groupname            = "my_lsn_group"
	transportprofilename = "my_lsntransportfile"
	}
  
`

const testAccLsngroup_lsntransportprofile_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
`

func TestAccLsngroup_lsntransportprofile_binding_basic(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this LSN resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckLsngroup_lsntransportprofile_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsngroup_lsntransportprofile_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsngroup_lsntransportprofile_bindingExist("citrixadc_lsngroup_lsntransportprofile_binding.tf_lsngroup_lsntransportprofile_binding", nil),
				),
			},
			{
				Config: testAccLsngroup_lsntransportprofile_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsngroup_lsntransportprofile_bindingNotExist("citrixadc_lsngroup_lsntransportprofile_binding.tf_lsngroup_lsntransportprofile_binding", "my_lsn_group,my_lsntransportfile"),
				),
			},
		},
	})
}

func testAccCheckLsngroup_lsntransportprofile_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsngroup_lsntransportprofile_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 2)

		groupname := idSlice[0]
		transportprofilename := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "lsngroup_lsntransportprofile_binding",
			ResourceName:             groupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching transportprofilename
		found := false
		for _, v := range dataArr {
			if v["transportprofilename"].(string) == transportprofilename {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("lsngroup_lsntransportprofile_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsngroup_lsntransportprofile_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		groupname := idSlice[0]
		transportprofilename := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "lsngroup_lsntransportprofile_binding",
			ResourceName:             groupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching transportprofilename
		found := false
		for _, v := range dataArr {
			if v["transportprofilename"].(string) == transportprofilename {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("lsngroup_lsntransportprofile_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckLsngroup_lsntransportprofile_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsngroup_lsntransportprofile_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lsngroup_lsntransportprofile_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lsngroup_lsntransportprofile_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
