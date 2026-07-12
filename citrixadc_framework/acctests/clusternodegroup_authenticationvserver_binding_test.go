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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strings"
	"testing"
)

const testAccClusternodegroup_authenticationvserver_binding_basic = `

resource "citrixadc_clusternodegroup" "tf_clusternodegroup" {
	name   = "my_tf_group"
	strict = "NO"
}

resource "citrixadc_clusternodegroup_clusternode_binding" "tf_clusternodegroup_clusternode_binding" {
	name = citrixadc_clusternodegroup.tf_clusternodegroup.name
	node = 0
}

resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
	name           = "my_authentication_server"
	servicetype    = "SSL"
	authentication = "ON"
	state          = "ENABLED"
}

resource "citrixadc_clusternodegroup_authenticationvserver_binding" "tf_clusternodegroup_authenticationvserver_binding" {
	name       = citrixadc_clusternodegroup.tf_clusternodegroup.name
	vserver    = citrixadc_authenticationvserver.tf_authenticationvserver.name
	depends_on = [citrixadc_clusternodegroup_clusternode_binding.tf_clusternodegroup_clusternode_binding]
}
`

// Step 2 keeps the participating entities (nodegroup, clusternode binding and
// authenticationvserver) but drops the binding itself so proper deletion of the
// binding can be verified while the endpoints still exist.
const testAccClusternodegroup_authenticationvserver_binding_basic_step2 = `

resource "citrixadc_clusternodegroup" "tf_clusternodegroup" {
	name   = "my_tf_group"
	strict = "NO"
}

resource "citrixadc_clusternodegroup_clusternode_binding" "tf_clusternodegroup_clusternode_binding" {
	name = citrixadc_clusternodegroup.tf_clusternodegroup.name
	node = 0
}

resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
	name           = "my_authentication_server"
	servicetype    = "SSL"
	authentication = "ON"
	state          = "ENABLED"
}
`

func TestAccClusternodegroup_authenticationvserver_binding_basic(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckClusternodegroup_authenticationvserver_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClusternodegroup_authenticationvserver_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroup_authenticationvserver_bindingExist("citrixadc_clusternodegroup_authenticationvserver_binding.tf_clusternodegroup_authenticationvserver_binding", nil),
				),
			},
			{
				Config: testAccClusternodegroup_authenticationvserver_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroup_authenticationvserver_bindingNotExist("citrixadc_clusternodegroup_authenticationvserver_binding.tf_clusternodegroup_authenticationvserver_binding", "my_tf_group,my_authentication_server"),
				),
			},
		},
	})
}

func testAccCheckClusternodegroup_authenticationvserver_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No clusternodegroup_authenticationvserver_binding id is set")
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
			return err
		}
		name := idMap["name"]
		vserver := idMap["vserver"]

		findParams := service.FindParams{
			ResourceType:             "clusternodegroup_authenticationvserver_binding",
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
			return fmt.Errorf("clusternodegroup_authenticationvserver_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckClusternodegroup_authenticationvserver_bindingNotExist(n string, id string) resource.TestCheckFunc {
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
			ResourceType:             "clusternodegroup_authenticationvserver_binding",
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
			return fmt.Errorf("clusternodegroup_authenticationvserver_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckClusternodegroup_authenticationvserver_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_clusternodegroup_authenticationvserver_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Clusternodegroup_authenticationvserver_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("clusternodegroup_authenticationvserver_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccClusternodegroup_authenticationvserver_bindingDataSource_basic = `

	resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
		name           = "my_authentication_server_ds"
		servicetype    = "SSL"
		authentication = "ON"
		state          = "ENABLED"
	}

	resource "citrixadc_clusternodegroup_authenticationvserver_binding" "tf_clusternodegroup_authenticationvserver_binding" {
		name    = "my_tf_group"
		vserver = citrixadc_authenticationvserver.tf_authenticationvserver.name
	}

	data "citrixadc_clusternodegroup_authenticationvserver_binding" "tf_clusternodegroup_authenticationvserver_binding" {
		name    = citrixadc_clusternodegroup_authenticationvserver_binding.tf_clusternodegroup_authenticationvserver_binding.name
		vserver = citrixadc_clusternodegroup_authenticationvserver_binding.tf_clusternodegroup_authenticationvserver_binding.vserver
		depends_on = [citrixadc_clusternodegroup_authenticationvserver_binding.tf_clusternodegroup_authenticationvserver_binding]
	}
`

func TestAccclusternodegroup_authenticationvserver_bindingDataSource_basic(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccClusternodegroup_authenticationvserver_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_clusternodegroup_authenticationvserver_binding.tf_clusternodegroup_authenticationvserver_binding", "name", "my_tf_group"),
					resource.TestCheckResourceAttr("data.citrixadc_clusternodegroup_authenticationvserver_binding.tf_clusternodegroup_authenticationvserver_binding", "vserver", "my_authentication_server_ds"),
				),
			},
		},
	})
}

// testAccClusternodegroup_authenticationvserver_binding_upgrade_basic reuses the _basic
// config values (binding name "my_tf_group", vserver "my_authentication_server") and adds
// the required nodegroup prerequisites (a clusternodegroup + a clusternode binding to
// activate it) so the binding can actually be created on a cluster testbed. It is valid
// under BOTH the SDK v2 2.2.0 schema and the current Framework schema because the SDK v2
// attribute names (name, vserver) are preserved.
const testAccClusternodegroup_authenticationvserver_binding_upgrade_basic = `

resource "citrixadc_clusternodegroup" "tf_clusternodegroup" {
	name   = "my_tf_group"
	strict = "NO"
}

resource "citrixadc_clusternodegroup_clusternode_binding" "tf_clusternodegroup_clusternode_binding" {
	name = citrixadc_clusternodegroup.tf_clusternodegroup.name
	node = 0
}

resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
	name           = "my_authentication_server"
	servicetype    = "SSL"
	authentication = "ON"
	state          = "ENABLED"
}

resource "citrixadc_clusternodegroup_authenticationvserver_binding" "tf_clusternodegroup_authenticationvserver_binding" {
	name       = citrixadc_clusternodegroup.tf_clusternodegroup.name
	vserver    = citrixadc_authenticationvserver.tf_authenticationvserver.name
	depends_on = [citrixadc_clusternodegroup_clusternode_binding.tf_clusternodegroup_clusternode_binding]
}
`

// TestAccClusternodegroup_authenticationvserver_binding_sdkv2StateUpgrade verifies that
// state written by the last SDK v2 release (legacy comma-separated ID) is correctly
// handled when the same config is subsequently managed by the current provider.
// Step 1 creates the binding with citrix/citrixadc 2.2.0 (writes the legacy id
// "my_tf_group,my_authentication_server"). Step 2 refreshes/plans/applies the same
// config through the current provider; because the Framework recomputes the id on Read
// (SetAttrFromGet re-derives data.Id), the id upgrades to the new "key:value" form.
func TestAccClusternodegroup_authenticationvserver_binding_sdkv2StateUpgrade(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resourceAddr := "citrixadc_clusternodegroup_authenticationvserver_binding.tf_clusternodegroup_authenticationvserver_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckClusternodegroup_authenticationvserver_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> state carries the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccClusternodegroup_authenticationvserver_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroup_authenticationvserver_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "my_tf_group,my_authentication_server"),
				),
			},
			// Step 2: refresh/plan/apply the SAME config through the current provider. The
			// legacy-id state is read via ParseIdString and the id is recomputed to the new
			// key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccClusternodegroup_authenticationvserver_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroup_authenticationvserver_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "name:my_tf_group,vserver:my_authentication_server"),
				),
			},
		},
	})
}

func TestAccClusternodegroup_authenticationvserver_binding_import(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	const resAddr = "citrixadc_clusternodegroup_authenticationvserver_binding.tf_clusternodegroup_authenticationvserver_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckClusternodegroup_authenticationvserver_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccClusternodegroup_authenticationvserver_binding_basic},
			{Config: testAccClusternodegroup_authenticationvserver_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
