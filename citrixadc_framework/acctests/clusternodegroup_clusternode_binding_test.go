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

const testAccClusternodegroup_clusternode_binding_basic = `

resource "citrixadc_clusternodegroup" "tf_ng_node" {
	name   = "tf_ng_node"
	strict = "NO"
	sticky = "YES"
}

resource "citrixadc_clusternodegroup_clusternode_binding" "tf_clusternodegroup_clusternode_binding" {
	name = citrixadc_clusternodegroup.tf_ng_node.name
	node = 0
	depends_on = [citrixadc_clusternodegroup.tf_ng_node]
	}

`

const testAccClusternodegroup_clusternode_binding_basic_step2 = `

resource "citrixadc_clusternodegroup" "tf_ng_node" {
	name   = "tf_ng_node"
	strict = "NO"
	sticky = "YES"
}

`

func TestAccClusternodegroup_clusternode_binding_basic(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckClusternodegroup_clusternode_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClusternodegroup_clusternode_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroup_clusternode_bindingExist("citrixadc_clusternodegroup_clusternode_binding.tf_clusternodegroup_clusternode_binding", nil),
				),
			},
			{
				Config: testAccClusternodegroup_clusternode_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroup_clusternode_bindingNotExist("citrixadc_clusternodegroup_clusternode_binding.tf_clusternodegroup_clusternode_binding", "tf_ng_node,0"),
				),
			},
		},
	})
}

func testAccCheckClusternodegroup_clusternode_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No clusternodegroup_clusternode_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "node"}, nil)
		if err != nil {
			return err
		}

		name := idMap["name"]
		node := idMap["node"]

		findParams := service.FindParams{
			ResourceType:             "clusternodegroup_clusternode_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching node
		found := false
		for _, v := range dataArr {
			if v["node"].(string) == node {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("clusternodegroup_clusternode_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckClusternodegroup_clusternode_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(id, []string{"name", "node"}, nil)
		if err != nil {
			return err
		}

		name := idMap["name"]
		node := idMap["node"]

		findParams := service.FindParams{
			ResourceType:             "clusternodegroup_clusternode_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching node
		found := false
		for _, v := range dataArr {
			if v["node"].(string) == node {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("clusternodegroup_clusternode_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckClusternodegroup_clusternode_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_clusternodegroup_clusternode_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Clusternodegroup_clusternode_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("clusternodegroup_clusternode_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccClusternodegroup_clusternode_bindingDataSource_basic = `
	resource "citrixadc_clusternodegroup" "tf_ng_node_ds" {
		name   = "tf_ng_node_ds"
		strict = "NO"
		sticky = "YES"
	}

	resource "citrixadc_clusternodegroup_clusternode_binding" "tf_clusternodegroup_clusternode_binding" {
		name = citrixadc_clusternodegroup.tf_ng_node_ds.name
		node = 0
		depends_on = [citrixadc_clusternodegroup.tf_ng_node_ds]
	}

	data "citrixadc_clusternodegroup_clusternode_binding" "tf_clusternodegroup_clusternode_binding" {
		name = citrixadc_clusternodegroup_clusternode_binding.tf_clusternodegroup_clusternode_binding.name
		node = 0
		depends_on = [citrixadc_clusternodegroup_clusternode_binding.tf_clusternodegroup_clusternode_binding]
	}
`

func TestAccclusternodegroup_clusternode_bindingDataSource_basic(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccClusternodegroup_clusternode_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_clusternodegroup_clusternode_binding.tf_clusternodegroup_clusternode_binding", "name", "tf_ng_node_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_clusternodegroup_clusternode_binding.tf_clusternodegroup_clusternode_binding", "node", "0"),
				),
			},
		},
	})
}
