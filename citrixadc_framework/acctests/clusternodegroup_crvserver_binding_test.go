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
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccClusternodegroup_crvserver_binding_basic = `

	resource "citrixadc_crvserver" "tf_crvserver" {
		name        = "my_cache_redirection_server_ds"
		servicetype = "HTTP"
	}

	resource "citrixadc_clusternodegroup_crvserver_binding" "tf_clusternodegroup_crvserver_binding" {
		name = "my_tf_group"
		vserver = citrixadc_crvserver.tf_crvserver.name
	}
`

const testAccClusternodegroup_crvserver_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
`

func TestAccClusternodegroup_crvserver_binding_basic(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckClusternodegroup_crvserver_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClusternodegroup_crvserver_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroup_crvserver_bindingExist("citrixadc_clusternodegroup_crvserver_binding.tf_clusternodegroup_crvserver_binding", nil),
				),
			},
			{
				Config: testAccClusternodegroup_crvserver_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroup_crvserver_bindingNotExist("citrixadc_clusternodegroup_crvserver_binding.tf_clusternodegroup_crvserver_binding", "my_cr_group,my_crvserver"),
				),
			},
		},
	})
}

func testAccCheckClusternodegroup_crvserver_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No clusternodegroup_crvserver_binding id is set")
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

		name := idSlice[0]
		vserver := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "clusternodegroup_crvserver_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching vserver
		found := false
		for _, v := range dataArr {
			if v["vserver"].(string) == vserver {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("clusternodegroup_crvserver_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckClusternodegroup_crvserver_bindingNotExist(n string, id string) resource.TestCheckFunc {
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

		name := idSlice[0]
		vserver := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "clusternodegroup_crvserver_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching vserver
		found := false
		for _, v := range dataArr {
			if v["vserver"].(string) == vserver {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("clusternodegroup_crvserver_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckClusternodegroup_crvserver_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_clusternodegroup_crvserver_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Clusternodegroup_crvserver_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("clusternodegroup_crvserver_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccClusternodegroup_crvserver_bindingDataSource_basic = `

	resource "citrixadc_crvserver" "tf_crvserver" {
		name        = "my_cache_redirection_server_ds"
		servicetype = "HTTP"
	}

	resource "citrixadc_clusternodegroup_crvserver_binding" "tf_clusternodegroup_crvserver_binding" {
		name    = "my_tf_group"
		vserver = citrixadc_crvserver.tf_crvserver.name
	}

	data "citrixadc_clusternodegroup_crvserver_binding" "tf_clusternodegroup_crvserver_binding" {
		name    = citrixadc_clusternodegroup_crvserver_binding.tf_clusternodegroup_crvserver_binding.name
		vserver = citrixadc_clusternodegroup_crvserver_binding.tf_clusternodegroup_crvserver_binding.vserver
		depends_on = [citrixadc_clusternodegroup_crvserver_binding.tf_clusternodegroup_crvserver_binding]
	}
`

func TestAccclusternodegroup_crvserver_bindingDataSource_basic(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccClusternodegroup_crvserver_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_clusternodegroup_crvserver_binding.tf_clusternodegroup_crvserver_binding", "name", "my_tf_group"),
					resource.TestCheckResourceAttr("data.citrixadc_clusternodegroup_crvserver_binding.tf_clusternodegroup_crvserver_binding", "vserver", "my_cache_redirection_server_ds"),
				),
			},
		},
	})
}
