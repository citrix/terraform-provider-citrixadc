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

const testAccLbvserver_service_binding_basic_step1 = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_server" "tf_test_svr" {
	name = "192.168.43.33"
	ipaddress = "192.168.43.33"
}

resource "citrixadc_service" "tf_service" {
    name = "tf_service"
    ip = citrixadc_server.tf_test_svr.ipaddress
    servicetype  = "HTTP"
    port = 80
}

resource "citrixadc_lbvserver_service_binding" "tf_binding" {
  name = citrixadc_lbvserver.tf_lbvserver.name
  servicename = citrixadc_service.tf_service.name
  weight = 10
}
`

const testAccLbvserver_service_binding_basic_step2 = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_server" "tf_test_svr" {
	name = "192.168.43.33"
	ipaddress = "192.168.43.33"
}

resource "citrixadc_service" "tf_service" {
    name = "tf_service"
    ip = citrixadc_server.tf_test_svr.ipaddress
    servicetype  = "HTTP"
    port = 80
}

resource "citrixadc_lbvserver_service_binding" "tf_binding" {
  name = citrixadc_lbvserver.tf_lbvserver.name
  servicename = citrixadc_service.tf_service.name
  weight = 20
}
`

func TestAccLbvserver_service_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbvserver_service_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbvserver_service_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_service_bindingExist("citrixadc_lbvserver_service_binding.tf_binding", nil),
				),
			},
			{
				Config: testAccLbvserver_service_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_service_bindingExist("citrixadc_lbvserver_service_binding.tf_binding", nil),
				),
			},
		},
	})
}

const testAccLbvserver_service_bindingDataSource_basic = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_server" "tf_test_svr" {
	name = "192.168.43.33"
	ipaddress = "192.168.43.33"
}

resource "citrixadc_service" "tf_service" {
    name = "tf_service"
    ip = citrixadc_server.tf_test_svr.ipaddress
    servicetype  = "HTTP"
    port = 80
}

resource "citrixadc_lbvserver_service_binding" "tf_binding" {
  name = citrixadc_lbvserver.tf_lbvserver.name
  servicename = citrixadc_service.tf_service.name
  weight = 10
}

data "citrixadc_lbvserver_service_binding" "tf_binding" {
    name = citrixadc_lbvserver_service_binding.tf_binding.name
    servicename = citrixadc_lbvserver_service_binding.tf_binding.servicename
    depends_on = [citrixadc_lbvserver_service_binding.tf_binding]
}
`

func TestAccLbvserver_service_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLbvserver_service_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_service_binding.tf_binding", "name", "tf_lbvserver"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_service_binding.tf_binding", "servicename", "tf_service"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_service_binding.tf_binding", "weight", "10"),
				),
			},
		},
	})
}

const testAccLbvserver_service_binding_upgrade_basic = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_server" "tf_test_svr" {
	name = "192.168.43.33"
	ipaddress = "192.168.43.33"
}

resource "citrixadc_service" "tf_service" {
    name = "tf_service"
    ip = citrixadc_server.tf_test_svr.ipaddress
    servicetype  = "HTTP"
    port = 80
}

resource "citrixadc_lbvserver_service_binding" "tf_binding" {
  name = citrixadc_lbvserver.tf_lbvserver.name
  servicename = citrixadc_service.tf_service.name
  weight = 10
}
`

func TestAccLbvserver_service_binding_import(t *testing.T) {
	const resAddr = "citrixadc_lbvserver_service_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbvserver_service_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLbvserver_service_binding_basic_step1},
			{Config: testAccLbvserver_service_binding_basic_step1, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

// TestAccLbvserver_service_binding_sdkv2StateUpgrade verifies that state written by the
// last SDK v2 release (citrix/citrixadc 2.2.0) — which stores the legacy
// "name,servicename" composite ID — is upgraded cleanly by the current Framework
// provider.
//
// Step 1 creates the binding with the 2.2.0 provider from the registry, producing state
// with the legacy ID "tf_lbvserver,tf_service". Step 2 refreshes and applies the SAME
// config through the current Framework provider, exercising ParseIdString on the legacy
// ID. The Framework Read does NOT recompute the ID (see resource_schema.go), so the
// legacy ID is retained after the upgrade.
func TestAccLbvserver_service_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLbvserver_service_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create with the last SDK v2 release from the registry.
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccLbvserver_service_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_service_bindingExist("citrixadc_lbvserver_service_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_lbvserver_service_binding.tf_binding", "id", "tf_lbvserver,tf_service"),
				),
			},
			{
				// Step 2: refresh/apply the SAME config through the current Framework provider.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccLbvserver_service_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_service_bindingExist("citrixadc_lbvserver_service_binding.tf_binding", nil),
					// Framework Read recomputes the legacy comma-joined id into the new key:value form.
					resource.TestCheckResourceAttr("citrixadc_lbvserver_service_binding.tf_binding", "id", "name:tf_lbvserver,servicename:tf_service"),
				),
			},
		},
	})
}

func testAccCheckLbvserver_service_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbvserver_service_binding id is set")
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

		// Parse the ID with ParseIdString so both the new "key:value" format and the
		// legacy SDK v2 "name,servicename" comma format resolve correctly.
		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "servicename"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}
		name := idMap["name"]
		servicename := idMap["servicename"]

		findParams := service.FindParams{
			ResourceType:             "lbvserver_service_binding",
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
			if v["servicename"].(string) == servicename {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("lbvserver_service_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckLbvserver_service_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbvserver_service_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Lbvserver_service_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lbvserver_service_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
