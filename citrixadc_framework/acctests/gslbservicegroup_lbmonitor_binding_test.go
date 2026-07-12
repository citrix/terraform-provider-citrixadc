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

const testAccGslbservicegroup_lbmonitor_binding_basic = `

resource "citrixadc_gslbservicegroup_lbmonitor_binding" "tf_gslbservicegroup_lbmonitor_binding" {
	weight           = 20
	servicegroupname = citrixadc_gslbservicegroup.tf_gslbservicegroup.servicegroupname
	monitor_name      = citrixadc_lbmonitor.tfmonitor1.monitorname
  
	}
  
  resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
	servicegroupname = "test_gslbvservicegroup"
	servicetype      = "HTTP"
	cip              = "DISABLED"
	healthmonitor    = "NO"
	sitename         = citrixadc_gslbsite.site_local.sitename
	}
  
  resource "citrixadc_gslbsite" "site_local" {
	sitename        = "Site-Local"
	siteipaddress   = "172.31.96.234"
	sessionexchange = "DISABLED"
	sitepassword    = "password123"
	}
  
  resource "citrixadc_lbmonitor" "tfmonitor1" {
	monitorname = "tf_monitor"
	type        = "HTTP"
	}
`

const testAccGslbservicegroup_lbmonitor_binding_basic_step2 = `

  resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
	servicegroupname = "test_gslbvservicegroup"
	servicetype      = "HTTP"
	cip              = "DISABLED"
	healthmonitor    = "NO"
	sitename         = citrixadc_gslbsite.site_local.sitename
	}
  
  resource "citrixadc_gslbsite" "site_local" {
	sitename        = "Site-Local"
	siteipaddress   = "172.31.96.234"
	sessionexchange = "DISABLED"
	sitepassword    = "password123"
	}
  
  resource "citrixadc_lbmonitor" "tfmonitor1" {
	monitorname = "tf_monitor"
	type        = "HTTP"
	}
`

func TestAccGslbservicegroup_lbmonitor_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckGslbservicegroup_lbmonitor_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccGslbservicegroup_lbmonitor_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbservicegroup_lbmonitor_bindingExist("citrixadc_gslbservicegroup_lbmonitor_binding.tf_gslbservicegroup_lbmonitor_binding", nil),
				),
			},
			{
				Config: testAccGslbservicegroup_lbmonitor_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbservicegroup_lbmonitor_bindingNotExist("citrixadc_gslbservicegroup_lbmonitor_binding.tf_gslbservicegroup_lbmonitor_binding", "test_gslbvservicegroup,tf_monitor"),
				),
			},
		},
	})
}

func testAccCheckGslbservicegroup_lbmonitor_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No gslbservicegroup_lbmonitor_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"servicegroupname", "monitor_name"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %v: %v", bindingId, err)
		}
		servicegroupname := idMap["servicegroupname"]
		monitor_name := idMap["monitor_name"]

		findParams := service.FindParams{
			ResourceType:             "gslbservicegroup_lbmonitor_binding",
			ResourceName:             servicegroupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching monitor_name
		found := false
		for _, v := range dataArr {
			if v["monitor_name"].(string) == monitor_name {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("gslbservicegroup_lbmonitor_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckGslbservicegroup_lbmonitor_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(id, []string{"servicegroupname", "monitor_name"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %v: %v", id, err)
		}
		servicegroupname := idMap["servicegroupname"]
		monitor_name := idMap["monitor_name"]

		findParams := service.FindParams{
			ResourceType:             "gslbservicegroup_lbmonitor_binding",
			ResourceName:             servicegroupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching monitor_name
		found := false
		for _, v := range dataArr {
			if v["monitor_name"].(string) == monitor_name {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("gslbservicegroup_lbmonitor_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckGslbservicegroup_lbmonitor_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_gslbservicegroup_lbmonitor_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("gslbservicegroup_lbmonitor_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("gslbservicegroup_lbmonitor_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccGslbservicegroup_lbmonitor_bindingDataSource_basic = `

resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
	servicegroupname = "test_gslbvservicegroup"
	servicetype      = "HTTP"
	cip              = "DISABLED"
	healthmonitor    = "NO"
	sitename         = citrixadc_gslbsite.site_local.sitename
}

resource "citrixadc_gslbsite" "site_local" {
	sitename        = "Site-Local"
	siteipaddress   = "172.31.96.234"
	sessionexchange = "DISABLED"
	sitepassword    = "password123"
}

resource "citrixadc_lbmonitor" "tfmonitor1" {
	monitorname = "tf_monitor"
	type        = "HTTP"
}

resource "citrixadc_gslbservicegroup_lbmonitor_binding" "tf_gslbservicegroup_lbmonitor_binding" {
	weight           = 20
	servicegroupname = citrixadc_gslbservicegroup.tf_gslbservicegroup.servicegroupname
	monitor_name     = citrixadc_lbmonitor.tfmonitor1.monitorname
}

data "citrixadc_gslbservicegroup_lbmonitor_binding" "tf_gslbservicegroup_lbmonitor_binding" {
	servicegroupname = citrixadc_gslbservicegroup_lbmonitor_binding.tf_gslbservicegroup_lbmonitor_binding.servicegroupname
	monitor_name     = citrixadc_gslbservicegroup_lbmonitor_binding.tf_gslbservicegroup_lbmonitor_binding.monitor_name
	depends_on       = [citrixadc_gslbservicegroup_lbmonitor_binding.tf_gslbservicegroup_lbmonitor_binding]
}
`

func TestAccGslbservicegroup_lbmonitor_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGslbservicegroup_lbmonitor_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_gslbservicegroup_lbmonitor_binding.tf_gslbservicegroup_lbmonitor_binding", "servicegroupname", "test_gslbvservicegroup"),
					resource.TestCheckResourceAttr("data.citrixadc_gslbservicegroup_lbmonitor_binding.tf_gslbservicegroup_lbmonitor_binding", "monitor_name", "tf_monitor"),
					resource.TestCheckResourceAttr("data.citrixadc_gslbservicegroup_lbmonitor_binding.tf_gslbservicegroup_lbmonitor_binding", "weight", "20"),
				),
			},
		},
	})
}

// testAccGslbservicegroup_lbmonitor_binding_upgrade_basic reuses the _basic config
// (the binding plus all prerequisite resources). It is valid under BOTH the SDK v2
// 2.2.0 schema and the current provider schema because it uses the SDK v2 attribute
// names that the migration restored.
const testAccGslbservicegroup_lbmonitor_binding_upgrade_basic = `

resource "citrixadc_gslbservicegroup_lbmonitor_binding" "tf_gslbservicegroup_lbmonitor_binding" {
	weight           = 20
	servicegroupname = citrixadc_gslbservicegroup.tf_gslbservicegroup.servicegroupname
	monitor_name     = citrixadc_lbmonitor.tfmonitor1.monitorname
}

resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
	servicegroupname = "test_gslbvservicegroup"
	servicetype      = "HTTP"
	cip              = "DISABLED"
	healthmonitor    = "NO"
	sitename         = citrixadc_gslbsite.site_local.sitename
}

resource "citrixadc_gslbsite" "site_local" {
	sitename        = "Site-Local"
	siteipaddress   = "172.31.96.234"
	sessionexchange = "DISABLED"
	sitepassword    = "password123"
}

resource "citrixadc_lbmonitor" "tfmonitor1" {
	monitorname = "tf_monitor"
	type        = "HTTP"
}
`

// TestAccGslbservicegroup_lbmonitor_binding_sdkv2StateUpgrade verifies that state
// written by the last SDK v2 release (legacy comma-separated id) is correctly read
// and reconciled when the same config is subsequently managed by the current
// provider. Step 1 creates the binding with citrix/citrixadc 2.2.0, which writes the
// legacy id "test_gslbvservicegroup,tf_monitor". Step 2 refreshes/plans/applies the
// SAME config through the current provider.
//
// NOTE: on this branch the citrixadc_gslbservicegroup_lbmonitor_binding RESOURCE is
// still served by the SDK v2 provider (only the datasource is registered in the
// Framework provider; the Framework resource stub is not wired into Resources()).
// The SDK v2 Read does not recompute the id, so the id remains the legacy value after
// the step-2 refresh; step 2 therefore asserts the Exist helper only (the "id
// recompute on Read" behavior does not apply to this resource).
func TestAccGslbservicegroup_lbmonitor_binding_sdkv2StateUpgrade(t *testing.T) {
	resourceAddr := "citrixadc_gslbservicegroup_lbmonitor_binding.tf_gslbservicegroup_lbmonitor_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckGslbservicegroup_lbmonitor_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> state carries the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccGslbservicegroup_lbmonitor_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbservicegroup_lbmonitor_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "test_gslbvservicegroup,tf_monitor"),
				),
			},
			// Step 2: refresh/plan/apply the SAME config through the current provider,
			// exercising the read of the legacy-id state.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccGslbservicegroup_lbmonitor_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGslbservicegroup_lbmonitor_bindingExist(resourceAddr, nil),
				),
			},
		},
	})
}

func TestAccGslbservicegroup_lbmonitor_binding_import(t *testing.T) {
	const resAddr = "citrixadc_gslbservicegroup_lbmonitor_binding.tf_gslbservicegroup_lbmonitor_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckGslbservicegroup_lbmonitor_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccGslbservicegroup_lbmonitor_binding_basic},
			{Config: testAccGslbservicegroup_lbmonitor_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
