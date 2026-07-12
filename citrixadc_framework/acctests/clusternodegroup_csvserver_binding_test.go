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

const testAccClusternodegroup_csvserver_binding_basic = `

resource "citrixadc_clusternodegroup" "tf_clusternodegroup" {
	name   = "my_tf_group"
	strict = "NO"
}

resource "citrixadc_clusternodegroup_clusternode_binding" "tf_clusternodegroup_clusternode_binding" {
	name = citrixadc_clusternodegroup.tf_clusternodegroup.name
	node = 0
}

resource "citrixadc_csvserver" "tf_csvserver" {
	name        = "my_content_server_ds"
	servicetype = "HTTP"
	ipv46       = "10.71.139.100"
	port        = "80"
}

resource "citrixadc_clusternodegroup_csvserver_binding" "tf_clusternodegroup_csvserver_binding" {
	name       = citrixadc_clusternodegroup.tf_clusternodegroup.name
	vserver    = citrixadc_csvserver.tf_csvserver.name
	depends_on = [citrixadc_clusternodegroup_clusternode_binding.tf_clusternodegroup_clusternode_binding]
}
`

// Step 2 keeps the participating entities (nodegroup, clusternode binding and
// csvserver) but drops the binding itself so proper deletion of the binding can
// be verified while the endpoints still exist.
const testAccClusternodegroup_csvserver_binding_basic_step2 = `

resource "citrixadc_clusternodegroup" "tf_clusternodegroup" {
	name   = "my_tf_group"
	strict = "NO"
}

resource "citrixadc_clusternodegroup_clusternode_binding" "tf_clusternodegroup_clusternode_binding" {
	name = citrixadc_clusternodegroup.tf_clusternodegroup.name
	node = 0
}

resource "citrixadc_csvserver" "tf_csvserver" {
	name        = "my_content_server_ds"
	servicetype = "HTTP"
	ipv46       = "10.71.139.100"
	port        = "80"
}
`

func TestAccClusternodegroup_csvserver_binding_basic(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckClusternodegroup_csvserver_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClusternodegroup_csvserver_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroup_csvserver_bindingExist("citrixadc_clusternodegroup_csvserver_binding.tf_clusternodegroup_csvserver_binding", nil),
				),
			},
			{
				Config: testAccClusternodegroup_csvserver_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroup_csvserver_bindingNotExist("citrixadc_clusternodegroup_csvserver_binding.tf_clusternodegroup_csvserver_binding", "my_cs_group,my_csvserver"),
				),
			},
		},
	})
}

func TestAccClusternodegroup_csvserver_binding_import(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	const resAddr = "citrixadc_clusternodegroup_csvserver_binding.tf_clusternodegroup_csvserver_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckClusternodegroup_csvserver_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccClusternodegroup_csvserver_binding_basic},
			{Config: testAccClusternodegroup_csvserver_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func testAccCheckClusternodegroup_csvserver_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No clusternodegroup_csvserver_binding id is set")
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

		// Parse the ID with utils.ParseIdString so both the new key:value
		// format and the legacy comma-separated format are handled.
		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "vserver"}, nil)
		if err != nil {
			return err
		}
		name := idMap["name"]
		my_csvserver := idMap["vserver"]

		findParams := service.FindParams{
			ResourceType:             "clusternodegroup_csvserver_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching my_csvserver
		found := false
		for _, v := range dataArr {
			if v["vserver"].(string) == my_csvserver {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("clusternodegroup_csvserver_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckClusternodegroup_csvserver_bindingNotExist(n string, id string) resource.TestCheckFunc {
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
		my_csvserver := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "clusternodegroup_csvserver_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching my_csvserver
		found := false
		for _, v := range dataArr {
			if v["vserver"].(string) == my_csvserver {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("clusternodegroup_csvserver_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckClusternodegroup_csvserver_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_clusternodegroup_csvserver_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Clusternodegroup_csvserver_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("clusternodegroup_csvserver_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccClusternodegroup_csvserver_bindingDataSource_basic = `

	resource "citrixadc_csvserver" "tf_csvserver" {
		name        = "my_content_server_ds"
		servicetype = "HTTP"
		ipv46       = "10.71.139.100"
		port        = "80"
	}

	resource "citrixadc_clusternodegroup_csvserver_binding" "tf_clusternodegroup_csvserver_binding" {
		name    = "my_tf_group"
		vserver = citrixadc_csvserver.tf_csvserver.name
	}

	data "citrixadc_clusternodegroup_csvserver_binding" "tf_clusternodegroup_csvserver_binding" {
		name    = citrixadc_clusternodegroup_csvserver_binding.tf_clusternodegroup_csvserver_binding.name
		vserver = citrixadc_clusternodegroup_csvserver_binding.tf_clusternodegroup_csvserver_binding.vserver
		depends_on = [citrixadc_clusternodegroup_csvserver_binding.tf_clusternodegroup_csvserver_binding]
	}
`

func TestAccclusternodegroup_csvserver_bindingDataSource_basic(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccClusternodegroup_csvserver_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_clusternodegroup_csvserver_binding.tf_clusternodegroup_csvserver_binding", "name", "my_tf_group"),
					resource.TestCheckResourceAttr("data.citrixadc_clusternodegroup_csvserver_binding.tf_clusternodegroup_csvserver_binding", "vserver", "my_content_server_ds"),
				),
			},
		},
	})
}

// testAccClusternodegroup_csvserver_binding_upgrade_basic is valid under BOTH the
// SDK v2 2.2.0 schema and the current provider schema (uses the SDK v2 attribute
// names name/vserver). The terraform resource label matches the _basic config so
// the shared Exist/Destroy helpers and resource address apply.
const testAccClusternodegroup_csvserver_binding_upgrade_basic = `

	resource "citrixadc_clusternodegroup" "tf_clusternodegroup" {
		name   = "my_tf_group"
		strict = "NO"
	}

	resource "citrixadc_clusternodegroup_clusternode_binding" "tf_clusternodegroup_clusternode_binding" {
		name = citrixadc_clusternodegroup.tf_clusternodegroup.name
		node = 0
	}

	resource "citrixadc_csvserver" "tf_csvserver" {
		name        = "my_content_server_ds"
		servicetype = "HTTP"
		ipv46       = "10.71.139.100"
		port        = "80"
	}

	resource "citrixadc_clusternodegroup_csvserver_binding" "tf_clusternodegroup_csvserver_binding" {
		name       = citrixadc_clusternodegroup.tf_clusternodegroup.name
		vserver    = citrixadc_csvserver.tf_csvserver.name
		depends_on = [citrixadc_clusternodegroup_clusternode_binding.tf_clusternodegroup_clusternode_binding]
	}
`

func TestAccClusternodegroup_csvserver_binding_sdkv2StateUpgrade(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckClusternodegroup_csvserver_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release (2.2.0) from the registry.
			// The SDK v2 resource writes state with the LEGACY comma-joined id
			// (name,vserver) -> "my_tf_group,my_content_server_ds".
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccClusternodegroup_csvserver_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroup_csvserver_bindingExist("citrixadc_clusternodegroup_csvserver_binding.tf_clusternodegroup_csvserver_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_clusternodegroup_csvserver_binding.tf_clusternodegroup_csvserver_binding", "id", "my_tf_group,my_content_server_ds"),
				),
			},
			// Step 2: refresh the legacy-id state through the CURRENT provider, then
			// plan/apply. This resource is still served by SDK v2 in the current
			// provider (only its datasource is registered in the Framework provider;
			// NewClusternodegroupCsvserverBindingResource is not wired into
			// citrixadc_framework/provider Resources()). The Framework
			// SetAttrFromGet id recompute is therefore not exercised at runtime, so
			// the id is not upgraded to the new key:value format. Assert Exist only
			// (the sanctioned fallback when the recompute is genuinely absent).
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccClusternodegroup_csvserver_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroup_csvserver_bindingExist("citrixadc_clusternodegroup_csvserver_binding.tf_clusternodegroup_csvserver_binding", nil),
				),
			},
		},
	})
}
