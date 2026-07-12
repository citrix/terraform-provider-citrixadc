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

func TestAccServicegroup_lbmonitor_binding_import(t *testing.T) {
	const resAddr = "citrixadc_servicegroup_lbmonitor_binding.bind1"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckServicegroup_lbmonitor_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccServicegroup_lbmonitor_binding_basic_step1},
			{Config: testAccServicegroup_lbmonitor_binding_basic_step1, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func TestAccServicegroup_lbmonitor_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckServicegroup_lbmonitor_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServicegroup_lbmonitor_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroup_lbmonitor_bindingExist("citrixadc_servicegroup_lbmonitor_binding.bind1", nil),
					testAccCheckServicegroup_lbmonitor_bindingExist("citrixadc_servicegroup_lbmonitor_binding.bind2", nil),
				),
			},
			{
				Config: testAccServicegroup_lbmonitor_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroup_lbmonitor_bindingExist("citrixadc_servicegroup_lbmonitor_binding.bind1", nil),
					testAccCheckServicegroup_lbmonitor_bindingExist("citrixadc_servicegroup_lbmonitor_binding.bind2", nil),
				),
			},
			{
				Config: testAccServicegroup_lbmonitor_binding_basic_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroup_lbmonitor_bindingExist("citrixadc_servicegroup_lbmonitor_binding.bind1", nil),
				),
			},
		},
	})
}

func testAccCheckServicegroup_lbmonitor_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lb vserver name is set")
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

		servicegroupLbmonitorBindingId := rs.Primary.ID
		idMap, _, err := utils.ParseIdString(servicegroupLbmonitorBindingId, []string{"servicegroupname", "monitorname"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", servicegroupLbmonitorBindingId, err)
		}
		servicegroupName := idMap["servicegroupname"]
		monitorName := idMap["monitorname"]

		findParams := service.FindParams{
			ResourceType:             "servicegroup_lbmonitor_binding",
			ResourceName:             servicegroupName,
			ResourceMissingErrorCode: 258,
		}

		dataArr, err := client.FindResourceArrayWithParams(findParams)

		if err != nil {
			return err
		}

		found := false

		for _, v := range dataArr {
			if v["monitor_name"].(string) == monitorName {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckServicegroup_lbmonitor_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_servicegroup_lbmonitor_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Servicegroup_lbmonitor_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccServicegroup_lbmonitor_binding_basic_step1 = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_lbvserver"
  ipv46       = "192.168.13.46"
  port        = "80"
  servicetype = "HTTP"
}

resource "citrixadc_lbmonitor" "tfmonitor1" {
  monitorname = "tf-monitor1"
  type        = "HTTP"
}

resource "citrixadc_lbmonitor" "tfmonitor2" {
  monitorname = "tfmonitor2"
  type        = "PING"
}

resource "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "tf_servicegroup"
  lbvservers       = [citrixadc_lbvserver.tf_lbvserver.name]
  servicetype      = "HTTP"
  servicegroupmembers = [
    "192.168.33.33:80:1",
  ]

}

resource "citrixadc_servicegroup_lbmonitor_binding" "bind1" {
    servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
    monitorname = citrixadc_lbmonitor.tfmonitor1.monitorname
    weight = 80
}

resource "citrixadc_servicegroup_lbmonitor_binding" "bind2" {
    servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
    monitorname = citrixadc_lbmonitor.tfmonitor2.monitorname
    weight = 20
}

`

const testAccServicegroup_lbmonitor_binding_basic_step2 = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_lbvserver"
  ipv46       = "192.168.13.46"
  port        = "80"
  servicetype = "HTTP"
}

resource "citrixadc_lbmonitor" "tfmonitor1" {
  monitorname = "tf-monitor1"
  type        = "HTTP"
}

resource "citrixadc_lbmonitor" "tfmonitor2" {
  monitorname = "tfmonitor2"
  type        = "PING"
}

resource "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "tf_servicegroup"
  lbvservers       = [citrixadc_lbvserver.tf_lbvserver.name]
  servicetype      = "HTTP"
  servicegroupmembers = [
    "192.168.33.33:80:1",
  ]

}

resource "citrixadc_servicegroup_lbmonitor_binding" "bind1" {
    servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
    monitorname = citrixadc_lbmonitor.tfmonitor1.monitorname
    weight = 50
}

resource "citrixadc_servicegroup_lbmonitor_binding" "bind2" {
    servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
    monitorname = citrixadc_lbmonitor.tfmonitor2.monitorname
    weight = 50
}

`

const testAccServicegroup_lbmonitor_binding_basic_step3 = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_lbvserver"
  ipv46       = "192.168.13.46"
  port        = "80"
  servicetype = "HTTP"
}

resource "citrixadc_lbmonitor" "tfmonitor1" {
  monitorname = "tf-monitor1"
  type        = "HTTP"
}

resource "citrixadc_lbmonitor" "tfmonitor2" {
  monitorname = "tfmonitor2"
  type        = "PING"
}

resource "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "tf_servicegroup"
  lbvservers       = [citrixadc_lbvserver.tf_lbvserver.name]
  servicetype      = "HTTP"
  servicegroupmembers = [
    "192.168.33.33:80:1",
  ]

}

resource "citrixadc_servicegroup_lbmonitor_binding" "bind1" {
    servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
    monitorname = citrixadc_lbmonitor.tfmonitor1.monitorname
    weight = 50
}

`

const testAccServicegroup_lbmonitor_bindingDataSource_basic = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_lbvserver"
  ipv46       = "192.168.13.46"
  port        = "80"
  servicetype = "HTTP"
}

resource "citrixadc_lbmonitor" "tfmonitor1" {
  monitorname = "tf-monitor1"
  type        = "HTTP"
}

resource "citrixadc_lbmonitor" "tfmonitor2" {
  monitorname = "tfmonitor2"
  type        = "PING"
}

resource "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "tf_servicegroup"
  lbvservers       = [citrixadc_lbvserver.tf_lbvserver.name]
  servicetype      = "HTTP"
  servicegroupmembers = [
    "192.168.33.33:80:1",
  ]

}

resource "citrixadc_servicegroup_lbmonitor_binding" "bind1" {
    servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
    monitorname = citrixadc_lbmonitor.tfmonitor1.monitorname
    weight = 80
    port = 0
}

resource "citrixadc_servicegroup_lbmonitor_binding" "bind2" {
    servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
    monitorname = citrixadc_lbmonitor.tfmonitor2.monitorname
    weight = 20
}

data "citrixadc_servicegroup_lbmonitor_binding" "bind1" {
    servicegroupname = citrixadc_servicegroup_lbmonitor_binding.bind1.servicegroupname
    monitor_name = citrixadc_servicegroup_lbmonitor_binding.bind1.monitorname
    depends_on = [citrixadc_servicegroup_lbmonitor_binding.bind1]
}
`

func TestAccServicegroup_lbmonitor_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccServicegroup_lbmonitor_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_servicegroup_lbmonitor_binding.bind1", "servicegroupname", "tf_servicegroup"),
					resource.TestCheckResourceAttr("data.citrixadc_servicegroup_lbmonitor_binding.bind1", "monitor_name", "tf-monitor1"),
					resource.TestCheckResourceAttr("data.citrixadc_servicegroup_lbmonitor_binding.bind1", "weight", "80"),
				),
			},
		},
	})
}

// testAccServicegroup_lbmonitor_binding_upgrade_basic mirrors the _basic_step1
// config (lbvserver + two lbmonitors + servicegroup + two lbmonitor bindings). It
// is valid under BOTH the SDK v2 2.2.0 schema and the current framework schema, so
// it can be applied with the old provider in step 1 and re-planned with the new
// provider in step 2 of the state-upgrade test below.
const testAccServicegroup_lbmonitor_binding_upgrade_basic = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_lbvserver"
  ipv46       = "192.168.13.46"
  port        = "80"
  servicetype = "HTTP"
}

resource "citrixadc_lbmonitor" "tfmonitor1" {
  monitorname = "tf-monitor1"
  type        = "HTTP"
}

resource "citrixadc_lbmonitor" "tfmonitor2" {
  monitorname = "tfmonitor2"
  type        = "PING"
}

resource "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "tf_servicegroup"
  lbvservers       = [citrixadc_lbvserver.tf_lbvserver.name]
  servicetype      = "HTTP"
  servicegroupmembers = [
    "192.168.33.33:80:1",
  ]

}

resource "citrixadc_servicegroup_lbmonitor_binding" "bind1" {
    servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
    monitorname = citrixadc_lbmonitor.tfmonitor1.monitorname
    weight = 80
}

resource "citrixadc_servicegroup_lbmonitor_binding" "bind2" {
    servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
    monitorname = citrixadc_lbmonitor.tfmonitor2.monitorname
    weight = 20
}
`

// TestAccServicegroup_lbmonitor_binding_sdkv2StateUpgrade verifies that bindings
// created by the LAST SDK v2 release (2.2.0) — which write the legacy comma-joined
// id "servicegroupname,monitorname" — are refreshed and re-applied correctly by the
// CURRENT framework provider. Step 2 exercises ParseIdString on the legacy id during
// the framework Read.
//
// This resource's SetAttrFromGet RECOMPUTES data.Id into the new
// "servicegroupname:<v>,monitorname:<v>" format on every Read, so after the step-2
// refresh the id upgrades to the canonical new format (asserted below).
func TestAccServicegroup_lbmonitor_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckServicegroup_lbmonitor_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release from the registry. This
			// writes state carrying the LEGACY comma-joined id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccServicegroup_lbmonitor_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroup_lbmonitor_bindingExist("citrixadc_servicegroup_lbmonitor_binding.bind1", nil),
					testAccCheckServicegroup_lbmonitor_bindingExist("citrixadc_servicegroup_lbmonitor_binding.bind2", nil),
					resource.TestCheckResourceAttr("citrixadc_servicegroup_lbmonitor_binding.bind1", "id", "tf_servicegroup,tf-monitor1"),
					resource.TestCheckResourceAttr("citrixadc_servicegroup_lbmonitor_binding.bind2", "id", "tf_servicegroup,tfmonitor2"),
				),
			},
			// Step 2: same config through the CURRENT framework provider. Terraform
			// refreshes the legacy-id state through the framework Read (exercising
			// ParseIdString on the legacy id) then plans/applies. The framework Read
			// recomputes the id into the new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccServicegroup_lbmonitor_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroup_lbmonitor_bindingExist("citrixadc_servicegroup_lbmonitor_binding.bind1", nil),
					testAccCheckServicegroup_lbmonitor_bindingExist("citrixadc_servicegroup_lbmonitor_binding.bind2", nil),
					resource.TestCheckResourceAttr("citrixadc_servicegroup_lbmonitor_binding.bind1", "id", "servicegroupname:tf_servicegroup,monitorname:tf-monitor1"),
					resource.TestCheckResourceAttr("citrixadc_servicegroup_lbmonitor_binding.bind2", "id", "servicegroupname:tf_servicegroup,monitorname:tfmonitor2"),
				),
			},
		},
	})
}
