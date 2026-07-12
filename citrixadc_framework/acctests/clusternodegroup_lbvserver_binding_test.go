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

const testAccClusternodegroup_lbvserver_binding_basic = `

	resource "citrixadc_clusternodegroup" "tf_clusternodegroup" {
		name   = "my_tf_group"
		strict = "NO"
	}

	resource "citrixadc_clusternodegroup_clusternode_binding" "tf_clusternodegroup_clusternode_binding" {
		name = citrixadc_clusternodegroup.tf_clusternodegroup.name
		node = 0
	}

	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name        = "my_lb_vserver_ds"
		servicetype = "HTTP"
	}

	resource "citrixadc_clusternodegroup_lbvserver_binding" "tf_clusternodegroup_lbvserver_binding" {
		name       = citrixadc_clusternodegroup.tf_clusternodegroup.name
		vserver    = citrixadc_lbvserver.tf_lbvserver.name
		depends_on = [citrixadc_clusternodegroup_clusternode_binding.tf_clusternodegroup_clusternode_binding]
	}
`

const testAccClusternodegroup_lbvserver_binding_basic_step2 = `
	# Keep the participating entities (nodegroup, clusternode binding and lbvserver)
	# but drop the binding itself to verify proper deletion while endpoints exist.

	resource "citrixadc_clusternodegroup" "tf_clusternodegroup" {
		name   = "my_tf_group"
		strict = "NO"
	}

	resource "citrixadc_clusternodegroup_clusternode_binding" "tf_clusternodegroup_clusternode_binding" {
		name = citrixadc_clusternodegroup.tf_clusternodegroup.name
		node = 0
	}

	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name        = "my_lb_vserver_ds"
		servicetype = "HTTP"
	}
`

func TestAccClusternodegroup_lbvserver_binding_basic(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckClusternodegroup_lbvserver_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClusternodegroup_lbvserver_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroup_lbvserver_bindingExist("citrixadc_clusternodegroup_lbvserver_binding.tf_clusternodegroup_lbvserver_binding", nil),
				),
			},
			{
				Config: testAccClusternodegroup_lbvserver_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroup_lbvserver_bindingNotExist("citrixadc_clusternodegroup_lbvserver_binding.tf_clusternodegroup_lbvserver_binding", "my_tf_group,my_lb_vserver_ds"),
				),
			},
		},
	})
}

// testAccClusternodegroup_lbvserver_binding_upgrade_basic mirrors the _basic config and must be
// valid under BOTH the last SDK v2 release (2.2.0) schema AND the current provider schema.
const testAccClusternodegroup_lbvserver_binding_upgrade_basic = `

	resource "citrixadc_clusternodegroup" "tf_clusternodegroup" {
		name   = "my_tf_group"
		strict = "NO"
	}

	resource "citrixadc_clusternodegroup_clusternode_binding" "tf_clusternodegroup_clusternode_binding" {
		name = citrixadc_clusternodegroup.tf_clusternodegroup.name
		node = 0
	}

	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name        = "my_lb_vserver_ds"
		servicetype = "HTTP"
	}

	resource "citrixadc_clusternodegroup_lbvserver_binding" "tf_clusternodegroup_lbvserver_binding" {
		name       = citrixadc_clusternodegroup.tf_clusternodegroup.name
		vserver    = citrixadc_lbvserver.tf_lbvserver.name
		depends_on = [citrixadc_clusternodegroup_clusternode_binding.tf_clusternodegroup_clusternode_binding]
	}
`

// TestAccClusternodegroup_lbvserver_binding_sdkv2StateUpgrade verifies that state written by the
// last SDK v2 release (2.2.0) upgrades cleanly to the current provider.
//
// Note: the clusternodegroup_lbvserver_binding resource is now served by the Plugin Framework, so its
// Read recomputes the id to the new key:value form on refresh. Step 1 (SDK v2 2.2.0) writes the legacy
// comma-separated id; step 2 (current framework provider) upgrades it to the new name:vserver form.
func TestAccClusternodegroup_lbvserver_binding_sdkv2StateUpgrade(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckClusternodegroup_lbvserver_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release (2.2.0). State is written with the legacy comma-separated id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccClusternodegroup_lbvserver_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroup_lbvserver_bindingExist("citrixadc_clusternodegroup_lbvserver_binding.tf_clusternodegroup_lbvserver_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_clusternodegroup_lbvserver_binding.tf_clusternodegroup_lbvserver_binding", "id", "my_tf_group,my_lb_vserver_ds"),
				),
			},
			// Step 2: refresh/apply the same config through the current provider (exercising the state upgrade).
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccClusternodegroup_lbvserver_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroup_lbvserver_bindingExist("citrixadc_clusternodegroup_lbvserver_binding.tf_clusternodegroup_lbvserver_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_clusternodegroup_lbvserver_binding.tf_clusternodegroup_lbvserver_binding", "id", "name:my_tf_group,vserver:my_lb_vserver_ds"),
				),
			},
		},
	})
}

func TestAccClusternodegroup_lbvserver_binding_import(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	const resAddr = "citrixadc_clusternodegroup_lbvserver_binding.tf_clusternodegroup_lbvserver_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckClusternodegroup_lbvserver_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccClusternodegroup_lbvserver_binding_basic},
			{Config: testAccClusternodegroup_lbvserver_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func testAccCheckClusternodegroup_lbvserver_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No clusternodegroup_lbvserver_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "vserver"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}
		name := idMap["name"]
		vserver := idMap["vserver"]

		findParams := service.FindParams{
			ResourceType:             "clusternodegroup_lbvserver_binding",
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
			return fmt.Errorf("clusternodegroup_lbvserver_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckClusternodegroup_lbvserver_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(id, []string{"name", "vserver"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", id, err)
		}
		name := idMap["name"]
		vserver := idMap["vserver"]

		findParams := service.FindParams{
			ResourceType:             "clusternodegroup_lbvserver_binding",
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
			return fmt.Errorf("clusternodegroup_lbvserver_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckClusternodegroup_lbvserver_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_clusternodegroup_lbvserver_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Clusternodegroup_lbvserver_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("clusternodegroup_lbvserver_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccClusternodegroup_lbvserver_bindingDataSource_basic = `

	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name        = "my_lb_vserver_ds"
		servicetype = "HTTP"
	}

	resource "citrixadc_clusternodegroup_lbvserver_binding" "tf_clusternodegroup_lbvserver_binding" {
		name    = "my_tf_group"
		vserver = citrixadc_lbvserver.tf_lbvserver.name
	}

	data "citrixadc_clusternodegroup_lbvserver_binding" "tf_clusternodegroup_lbvserver_binding" {
		name    = citrixadc_clusternodegroup_lbvserver_binding.tf_clusternodegroup_lbvserver_binding.name
		vserver = citrixadc_clusternodegroup_lbvserver_binding.tf_clusternodegroup_lbvserver_binding.vserver
		depends_on = [citrixadc_clusternodegroup_lbvserver_binding.tf_clusternodegroup_lbvserver_binding]
	}
`

func TestAccclusternodegroup_lbvserver_bindingDataSource_basic(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccClusternodegroup_lbvserver_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_clusternodegroup_lbvserver_binding.tf_clusternodegroup_lbvserver_binding", "name", "my_tf_group"),
					resource.TestCheckResourceAttr("data.citrixadc_clusternodegroup_lbvserver_binding.tf_clusternodegroup_lbvserver_binding", "vserver", "my_lb_vserver_ds"),
				),
			},
		},
	})
}
