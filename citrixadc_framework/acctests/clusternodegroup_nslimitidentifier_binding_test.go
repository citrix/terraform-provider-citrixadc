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

const testAccClusternodegroup_nslimitidentifier_binding_basic = `

	resource "citrixadc_clusternodegroup" "tf_clusternodegroup" {
		name   = "my_tf_group"
		strict = "NO"
	}

	resource "citrixadc_clusternodegroup_clusternode_binding" "tf_clusternodegroup_clusternode_binding" {
		name = citrixadc_clusternodegroup.tf_clusternodegroup.name
		node = 0
	}

	resource "citrixadc_nslimitidentifier" "tf_nslimitidentifier" {
		limitidentifier = "my_ns_limit_identifier_ds"
		threshold       = 100
		timeslice       = 1000
	}

	resource "citrixadc_clusternodegroup_nslimitidentifier_binding" "tf_clusternodegroup_nslimitidentifier_binding" {
		name           = citrixadc_clusternodegroup.tf_clusternodegroup.name
		identifiername = citrixadc_nslimitidentifier.tf_nslimitidentifier.limitidentifier
		depends_on     = [citrixadc_clusternodegroup_clusternode_binding.tf_clusternodegroup_clusternode_binding]
	}

`

// Step 2 keeps the participating entities (nodegroup, clusternode binding and
// nslimitidentifier) but drops the binding itself so proper deletion of the
// binding can be verified while the endpoints still exist.
const testAccClusternodegroup_nslimitidentifier_binding_basic_step2 = `

	resource "citrixadc_clusternodegroup" "tf_clusternodegroup" {
		name   = "my_tf_group"
		strict = "NO"
	}

	resource "citrixadc_clusternodegroup_clusternode_binding" "tf_clusternodegroup_clusternode_binding" {
		name = citrixadc_clusternodegroup.tf_clusternodegroup.name
		node = 0
	}

	resource "citrixadc_nslimitidentifier" "tf_nslimitidentifier" {
		limitidentifier = "my_ns_limit_identifier_ds"
		threshold       = 100
		timeslice       = 1000
	}
`

func TestAccClusternodegroup_nslimitidentifier_binding_basic(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckClusternodegroup_nslimitidentifier_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClusternodegroup_nslimitidentifier_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroup_nslimitidentifier_bindingExist("citrixadc_clusternodegroup_nslimitidentifier_binding.tf_clusternodegroup_nslimitidentifier_binding", nil),
				),
			},
			{
				Config: testAccClusternodegroup_nslimitidentifier_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroup_nslimitidentifier_bindingNotExist("citrixadc_clusternodegroup_nslimitidentifier_binding.tf_clusternodegroup_nslimitidentifier_binding", "my_group,my_nslimitidentifier"),
				),
			},
		},
	})
}

func testAccCheckClusternodegroup_nslimitidentifier_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No clusternodegroup_nslimitidentifier_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "identifiername"}, nil)
		if err != nil {
			return err
		}
		name := idMap["name"]
		identifiername := idMap["identifiername"]

		findParams := service.FindParams{
			ResourceType:             "clusternodegroup_nslimitidentifier_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching identifiername
		found := false
		for _, v := range dataArr {
			if v["identifiername"].(string) == identifiername {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("clusternodegroup_nslimitidentifier_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckClusternodegroup_nslimitidentifier_bindingNotExist(n string, id string) resource.TestCheckFunc {
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
		identifiername := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "clusternodegroup_nslimitidentifier_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching identifiername
		found := false
		for _, v := range dataArr {
			if v["identifiername"].(string) == identifiername {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("clusternodegroup_nslimitidentifier_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckClusternodegroup_nslimitidentifier_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_clusternodegroup_nslimitidentifier_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Clusternodegroup_nslimitidentifier_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("clusternodegroup_nslimitidentifier_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccClusternodegroup_nslimitidentifier_bindingDataSource_basic = `

	resource "citrixadc_nslimitidentifier" "tf_nslimitidentifier" {
		limitidentifier = "my_ns_limit_identifier_ds"
		threshold       = 100
		timeslice       = 1000
	}

	resource "citrixadc_clusternodegroup_nslimitidentifier_binding" "tf_clusternodegroup_nslimitidentifier_binding" {
		name           = "my_tf_group"
		identifiername = citrixadc_nslimitidentifier.tf_nslimitidentifier.limitidentifier
	}

	data "citrixadc_clusternodegroup_nslimitidentifier_binding" "tf_clusternodegroup_nslimitidentifier_binding" {
		name           = citrixadc_clusternodegroup_nslimitidentifier_binding.tf_clusternodegroup_nslimitidentifier_binding.name
		identifiername = citrixadc_clusternodegroup_nslimitidentifier_binding.tf_clusternodegroup_nslimitidentifier_binding.identifiername
		depends_on = [citrixadc_clusternodegroup_nslimitidentifier_binding.tf_clusternodegroup_nslimitidentifier_binding]
	}
`

func TestAccclusternodegroup_nslimitidentifier_bindingDataSource_basic(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccClusternodegroup_nslimitidentifier_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_clusternodegroup_nslimitidentifier_binding.tf_clusternodegroup_nslimitidentifier_binding", "name", "my_tf_group"),
					resource.TestCheckResourceAttr("data.citrixadc_clusternodegroup_nslimitidentifier_binding.tf_clusternodegroup_nslimitidentifier_binding", "identifiername", "my_ns_limit_identifier_ds"),
				),
			},
		},
	})
}

// testAccClusternodegroup_nslimitidentifier_binding_upgrade_basic reuses the _basic
// config (binding + prerequisite nslimitidentifier). It is valid under BOTH the SDK v2
// 2.2.0 schema and the current Framework schema because the migration restored the SDK
// v2 attribute names.
const testAccClusternodegroup_nslimitidentifier_binding_upgrade_basic = `

	resource "citrixadc_clusternodegroup" "tf_clusternodegroup" {
		name   = "my_tf_group"
		strict = "NO"
	}

	resource "citrixadc_clusternodegroup_clusternode_binding" "tf_clusternodegroup_clusternode_binding" {
		name = citrixadc_clusternodegroup.tf_clusternodegroup.name
		node = 0
	}

	resource "citrixadc_nslimitidentifier" "tf_nslimitidentifier" {
		limitidentifier = "my_ns_limit_identifier_ds"
		threshold       = 100
		timeslice       = 1000
	}

	resource "citrixadc_clusternodegroup_nslimitidentifier_binding" "tf_clusternodegroup_nslimitidentifier_binding" {
		name           = citrixadc_clusternodegroup.tf_clusternodegroup.name
		identifiername = citrixadc_nslimitidentifier.tf_nslimitidentifier.limitidentifier
		depends_on     = [citrixadc_clusternodegroup_clusternode_binding.tf_clusternodegroup_clusternode_binding]
	}
`

// TestAccClusternodegroup_nslimitidentifier_binding_sdkv2StateUpgrade verifies that
// state written by the last SDK v2 release (legacy comma-separated ID) is correctly
// upgraded when the same config is subsequently managed by the current Framework
// provider. Step 1 creates the binding with citrix/citrixadc 2.2.0 (writes the legacy
// id "my_tf_group,my_ns_limit_identifier_ds"). Step 2 refreshes/plans/applies the same
// config through the Framework provider, exercising ParseIdString on the legacy id;
// because the Framework recomputes the id on Read (SetAttrFromGet), the id upgrades to
// the new "key:value" form "identifiername:my_ns_limit_identifier_ds,name:my_tf_group".
func TestAccClusternodegroup_nslimitidentifier_binding_sdkv2StateUpgrade(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resourceAddr := "citrixadc_clusternodegroup_nslimitidentifier_binding.tf_clusternodegroup_nslimitidentifier_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckClusternodegroup_nslimitidentifier_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> state carries the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccClusternodegroup_nslimitidentifier_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroup_nslimitidentifier_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "my_tf_group,my_ns_limit_identifier_ds"),
				),
			},
			// Step 2: refresh/plan/apply the SAME config through the current Framework
			// provider. The legacy-id state is read via ParseIdString and the id is
			// recomputed to the new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccClusternodegroup_nslimitidentifier_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroup_nslimitidentifier_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "identifiername:my_ns_limit_identifier_ds,name:my_tf_group"),
				),
			},
		},
	})
}

func TestAccClusternodegroup_nslimitidentifier_binding_import(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	const resAddr = "citrixadc_clusternodegroup_nslimitidentifier_binding.tf_clusternodegroup_nslimitidentifier_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckClusternodegroup_nslimitidentifier_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccClusternodegroup_nslimitidentifier_binding_basic},
			{Config: testAccClusternodegroup_nslimitidentifier_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
