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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccClusternodegroup_clusternode_binding_basic = `

resource "citrixadc_clusternodegroup" "tf_clusternodegroup" {
	name   = "my_tf_group"
	strict = "NO"
}

resource "citrixadc_clusternodegroup_clusternode_binding" "tf_clusternodegroup_clusternode_binding" {
	name = citrixadc_clusternodegroup.tf_clusternodegroup.name
	node = 0
}

`

// Step 2 keeps the nodegroup but drops the clusternode binding itself so proper
// deletion of the binding can be verified while the nodegroup still exists.
const testAccClusternodegroup_clusternode_binding_basic_step2 = `

resource "citrixadc_clusternodegroup" "tf_clusternodegroup" {
	name   = "my_tf_group"
	strict = "NO"
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
					testAccCheckClusternodegroup_clusternode_bindingNotExist("citrixadc_clusternodegroup_clusternode_binding.tf_clusternodegroup_clusternode_binding", "my_group,2"),
				),
			},
		},
	})
}

func TestAccClusternodegroup_clusternode_binding_import(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	const resAddr = "citrixadc_clusternodegroup_clusternode_binding.tf_clusternodegroup_clusternode_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckClusternodegroup_clusternode_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccClusternodegroup_clusternode_binding_basic},
			{Config: testAccClusternodegroup_clusternode_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
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

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		name := idSlice[0]
		node := idSlice[1]

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
	resource "citrixadc_clusternodegroup_clusternode_binding" "tf_clusternodegroup_clusternode_binding" {
		name = "my_tf_group"
		node = 0
	}

	data "citrixadc_clusternodegroup_clusternode_binding" "tf_clusternodegroup_clusternode_binding" {
		name = "my_tf_group"
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
					resource.TestCheckResourceAttr("data.citrixadc_clusternodegroup_clusternode_binding.tf_clusternodegroup_clusternode_binding", "name", "my_tf_group"),
					resource.TestCheckResourceAttr("data.citrixadc_clusternodegroup_clusternode_binding.tf_clusternodegroup_clusternode_binding", "node", "0"),
				),
			},
		},
	})
}

// testAccClusternodegroup_clusternode_binding_upgrade_basic mirrors the _basic
// config (a clusternodegroup nodegroup + a node bound by index). It is valid under
// BOTH the SDK v2 2.2.0 schema and the current framework schema, so it can be
// applied with the old provider in step 1 and re-planned with the new provider in
// step 2 of the state-upgrade test below. The terraform resource label is kept
// identical to the _basic config so the Exist/Destroy helpers and addresses match.
const testAccClusternodegroup_clusternode_binding_upgrade_basic = `

resource "citrixadc_clusternodegroup" "tf_clusternodegroup" {
	name   = "my_tf_group"
	strict = "NO"
}

resource "citrixadc_clusternodegroup_clusternode_binding" "tf_clusternodegroup_clusternode_binding" {
	name = citrixadc_clusternodegroup.tf_clusternodegroup.name
	node = 0
}

`

// TestAccClusternodegroup_clusternode_binding_sdkv2StateUpgrade verifies that a
// resource created by the LAST SDK v2 release (2.2.0) — which writes the legacy
// comma-joined id "name,node" — is refreshed and re-applied correctly by the
// CURRENT framework provider. Step 2 exercises the framework Read (which recomputes
// the id to the new "name:UrlEncode(value),node:UrlEncode(value)" format via
// SetAttrFromGet) on the legacy-id state.
func TestAccClusternodegroup_clusternode_binding_sdkv2StateUpgrade(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckClusternodegroup_clusternode_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release from the registry. This
			// writes state carrying the LEGACY comma-joined id "name,node".
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccClusternodegroup_clusternode_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroup_clusternode_bindingExist("citrixadc_clusternodegroup_clusternode_binding.tf_clusternodegroup_clusternode_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_clusternodegroup_clusternode_binding.tf_clusternodegroup_clusternode_binding", "id", "my_tf_group,0"),
				),
			},
			// Step 2: same config through the CURRENT framework provider. Terraform
			// refreshes the legacy-id state through the framework Read (recomputing
			// the id to the new key:value format) then plans/applies.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccClusternodegroup_clusternode_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroup_clusternode_bindingExist("citrixadc_clusternodegroup_clusternode_binding.tf_clusternodegroup_clusternode_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_clusternodegroup_clusternode_binding.tf_clusternodegroup_clusternode_binding", "id", "name:my_tf_group,node:0"),
				),
			},
		},
	})
}
