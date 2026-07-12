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

const testAccLbvserver_analyticsprofile_binding_basic = `
resource "citrixadc_lbvserver_analyticsprofile_binding" "foo" {
	name = citrixadc_lbvserver.test_server.name
    analyticsprofile = "ns_analytics_global_profile"

}

resource "citrixadc_lbvserver" "test_server" {
	name = "test_server"
	servicetype = "HTTP"
}

`

func TestAccLbvserver_analyticsprofile_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbvserver_analyticsprofile_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbvserver_analyticsprofile_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_analyticsprofile_bindingExist("citrixadc_lbvserver_analyticsprofile_binding.foo", nil),
				),
			},
		},
	})
}

func testAccCheckLbvserver_analyticsprofile_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbvserver_analyticsprofile_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "analyticsprofile"}, nil)
		if err != nil {
			return err
		}
		name := idMap["name"]
		analyticsprofile := idMap["analyticsprofile"]

		findParams := service.FindParams{
			ResourceType:             "lbvserver_analyticsprofile_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right monitor name
		found := false
		for _, v := range dataArr {
			if v["analyticsprofile"].(string) == analyticsprofile {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("lbvserver_analyticsprofile_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckLbvserver_analyticsprofile_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbvserver_analyticsprofile_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lbvserver_analyticsprofile_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lbvserver_analyticsprofile_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLbvserver_analyticsprofile_bindingDataSource_basic = `
resource "citrixadc_lbvserver" "test_server" {
	name = "test_server"
	servicetype = "HTTP"
}

resource "citrixadc_lbvserver_analyticsprofile_binding" "foo" {
	name = citrixadc_lbvserver.test_server.name
    analyticsprofile = "ns_analytics_global_profile"
}

data "citrixadc_lbvserver_analyticsprofile_binding" "foo" {
	name = citrixadc_lbvserver_analyticsprofile_binding.foo.name
	analyticsprofile = citrixadc_lbvserver_analyticsprofile_binding.foo.analyticsprofile
	depends_on = [citrixadc_lbvserver_analyticsprofile_binding.foo]
}
`

func TestAccLbvserver_analyticsprofile_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLbvserver_analyticsprofile_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_analyticsprofile_binding.foo", "name", "test_server"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_analyticsprofile_binding.foo", "analyticsprofile", "ns_analytics_global_profile"),
				),
			},
		},
	})
}

// testAccLbvserver_analyticsprofile_binding_upgrade_basic reuses the _basic config
// (binding + its prerequisite lbvserver). It is valid under BOTH the SDK v2 2.2.0
// schema and the current Framework schema because the migration restored the SDK v2
// attribute names (name, analyticsprofile).
const testAccLbvserver_analyticsprofile_binding_upgrade_basic = `
resource "citrixadc_lbvserver_analyticsprofile_binding" "foo" {
	name             = citrixadc_lbvserver.test_server.name
	analyticsprofile = "ns_analytics_global_profile"
}

resource "citrixadc_lbvserver" "test_server" {
	name        = "test_server"
	servicetype = "HTTP"
}
`

// TestAccLbvserver_analyticsprofile_binding_sdkv2StateUpgrade verifies that state written
// by the last SDK v2 release is correctly upgraded when the same config is subsequently
// managed by the current Framework provider. Step 1 creates the binding with
// citrix/citrixadc 2.2.0 (writes the legacy comma id "test_server,ns_analytics_global_profile" —
// the SDK v2 d.SetId(fmt.Sprintf("%s,%s", name, analyticsprofile))). Step 2 refreshes/plans/
// applies the same config through the Framework provider, exercising ParseIdString on the
// legacy id; the Framework recomputes the id on Read (SetAttrFromGet), so the canonical
// new-format id becomes "analyticsprofile:ns_analytics_global_profile,name:test_server".
func TestAccLbvserver_analyticsprofile_binding_sdkv2StateUpgrade(t *testing.T) {
	resourceAddr := "citrixadc_lbvserver_analyticsprofile_binding.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLbvserver_analyticsprofile_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> state carries the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccLbvserver_analyticsprofile_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_analyticsprofile_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "test_server,ns_analytics_global_profile"),
				),
			},
			// Step 2: refresh/plan/apply the SAME config through the current Framework
			// provider. The legacy-id state is read via ParseIdString and the id is
			// recomputed on Read into the new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccLbvserver_analyticsprofile_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_analyticsprofile_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "analyticsprofile:ns_analytics_global_profile,name:test_server"),
				),
			},
		},
	})
}

func TestAccLbvserver_analyticsprofile_binding_import(t *testing.T) {
	const resAddr = "citrixadc_lbvserver_analyticsprofile_binding.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbvserver_analyticsprofile_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLbvserver_analyticsprofile_binding_basic},
			{Config: testAccLbvserver_analyticsprofile_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
