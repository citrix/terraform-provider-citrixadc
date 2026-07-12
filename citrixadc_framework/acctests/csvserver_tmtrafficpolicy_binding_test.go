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

const testAccCsvserver_tmtrafficpolicy_binding_basic = `

	resource "citrixadc_tmtrafficaction" "tf_tmtrafficaction" {
		name             = "my_trafficaction"
		apptimeout       = 5
		sso              = "OFF"
		persistentcookie = "ON"
	}
	resource "citrixadc_tmtrafficpolicy" "tf_tmtrafficpolicy" {
		name   = "tf_tmttrafficpolicy"
		rule   = "true"
		action = citrixadc_tmtrafficaction.tf_tmtrafficaction.name
	}
	resource "citrixadc_csvserver_tmtrafficpolicy_binding" "tf_csvserver_tmtrafficpolicy_binding" {
		name 		= citrixadc_csvserver.tf_csvserver.name
		policyname	= citrixadc_tmtrafficpolicy.tf_tmtrafficpolicy.name
		priority 	= 1
	}

	resource "citrixadc_csvserver" "tf_csvserver" {
		name 		= "tf_csvserver"
		ipv46 		= "10.202.11.11"
		port 		= 8080
		servicetype = "HTTP"
	}
`

const testAccCsvserver_tmtrafficpolicy_binding_basic_step2 = `

	resource "citrixadc_tmtrafficaction" "tf_tmtrafficaction" {
		name             = "my_trafficaction"
		apptimeout       = 5
		sso              = "OFF"
		persistentcookie = "ON"
	}
	resource "citrixadc_tmtrafficpolicy" "tf_tmtrafficpolicy" {
		name   = "tf_tmttrafficpolicy"
		rule   = "true"
		action = citrixadc_tmtrafficaction.tf_tmtrafficaction.name
	}
	resource "citrixadc_csvserver" "tf_csvserver" {
		name 		= "tf_csvserver"
		ipv46 		= "10.202.11.11"
		port 		= 8080
		servicetype = "HTTP"
	}
`

func TestAccCsvserver_tmtrafficpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsvserver_tmtrafficpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_tmtrafficpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_tmtrafficpolicy_bindingExist("citrixadc_csvserver_tmtrafficpolicy_binding.tf_csvserver_tmtrafficpolicy_binding", nil),
				),
			},
			{
				Config: testAccCsvserver_tmtrafficpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_tmtrafficpolicy_bindingNotExist("citrixadc_csvserver_tmtrafficpolicy_binding.tf_csvserver_tmtrafficpolicy_binding", "tf_csvserver,tf_tmttrafficpolicy"),
				),
			},
		},
	})
}

func TestAccCsvserver_tmtrafficpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_csvserver_tmtrafficpolicy_binding.tf_csvserver_tmtrafficpolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCsvserver_tmtrafficpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccCsvserver_tmtrafficpolicy_binding_basic},
			{Config: testAccCsvserver_tmtrafficpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func testAccCheckCsvserver_tmtrafficpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No csvserver_tmtrafficpolicy_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "policyname"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}
		name := idMap["name"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "csvserver_tmtrafficpolicy_binding",
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
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("csvserver_tmtrafficpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckCsvserver_tmtrafficpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
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

		csvserverName := idSlice[0]
		policyName := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "csvserver_tmtrafficpolicy_binding",
			ResourceName:             csvserverName,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right policy name
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyName {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("csvserver_tmtrafficpolicy_binding %s was found, but is should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckCsvserver_tmtrafficpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_csvserver_tmtrafficpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Csvserver_tmtrafficpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("csvserver_tmtrafficpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccCsvserver_tmtrafficpolicy_binding_upgrade_basic = `

	resource "citrixadc_tmtrafficaction" "tf_tmtrafficaction" {
		name             = "my_trafficaction"
		apptimeout       = 5
		sso              = "OFF"
		persistentcookie = "ON"
	}
	resource "citrixadc_tmtrafficpolicy" "tf_tmtrafficpolicy" {
		name   = "tf_tmttrafficpolicy"
		rule   = "true"
		action = citrixadc_tmtrafficaction.tf_tmtrafficaction.name
	}
	resource "citrixadc_csvserver_tmtrafficpolicy_binding" "tf_csvserver_tmtrafficpolicy_binding" {
		name 		= citrixadc_csvserver.tf_csvserver.name
		policyname	= citrixadc_tmtrafficpolicy.tf_tmtrafficpolicy.name
		priority 	= 1
	}

	resource "citrixadc_csvserver" "tf_csvserver" {
		name 		= "tf_csvserver"
		ipv46 		= "10.202.11.11"
		port 		= 8080
		servicetype = "HTTP"
	}
`

// TestAccCsvserver_tmtrafficpolicy_binding_sdkv2StateUpgrade verifies that state written by
// the last SDK v2 release (v2.2.0), which uses the legacy comma-joined id
// (name,policyname), is transparently upgraded by the current Framework provider.
// Step 1 creates the resource with the external SDK v2 provider (legacy id).
// Step 2 refreshes/plans/applies the SAME config through the Framework provider; its Read
// re-derives the canonical new-format id (name:<v>,policyname:<v>) via SetAttrFromGet.
func TestAccCsvserver_tmtrafficpolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckCsvserver_tmtrafficpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Create with the last SDK v2 release from the registry -> legacy id.
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccCsvserver_tmtrafficpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_tmtrafficpolicy_bindingExist("citrixadc_csvserver_tmtrafficpolicy_binding.tf_csvserver_tmtrafficpolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_csvserver_tmtrafficpolicy_binding.tf_csvserver_tmtrafficpolicy_binding", "id", "tf_csvserver,tf_tmttrafficpolicy"),
				),
			},
			{
				// Refresh legacy-id state through the current Framework provider -> id upgraded.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccCsvserver_tmtrafficpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCsvserver_tmtrafficpolicy_bindingExist("citrixadc_csvserver_tmtrafficpolicy_binding.tf_csvserver_tmtrafficpolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_csvserver_tmtrafficpolicy_binding.tf_csvserver_tmtrafficpolicy_binding", "id", "name:tf_csvserver,policyname:tf_tmttrafficpolicy"),
				),
			},
		},
	})
}

func TestAccCsvserver_tmtrafficpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCsvserver_tmtrafficpolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_tmtrafficpolicy_binding.test", "name", "tf_csvserver"),
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_tmtrafficpolicy_binding.test", "policyname", "tf_tmttrafficpolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_csvserver_tmtrafficpolicy_binding.test", "priority", "1"),
				),
			},
		},
	})
}

const testAccCsvserver_tmtrafficpolicy_bindingDataSource_basic = `

resource "citrixadc_tmtrafficaction" "tf_tmtrafficaction" {
	name             = "my_trafficaction"
	apptimeout       = 5
	sso              = "OFF"
	persistentcookie = "ON"
}
resource "citrixadc_tmtrafficpolicy" "tf_tmtrafficpolicy" {
	name   = "tf_tmttrafficpolicy"
	rule   = "true"
	action = citrixadc_tmtrafficaction.tf_tmtrafficaction.name
}
resource "citrixadc_csvserver" "tf_csvserver" {
	name 		= "tf_csvserver"
	ipv46 		= "10.202.11.11"
	port 		= 8080
	servicetype = "HTTP"
}
resource "citrixadc_csvserver_tmtrafficpolicy_binding" "tf_csvserver_tmtrafficpolicy_binding" {
	name 		= citrixadc_csvserver.tf_csvserver.name
	policyname	= citrixadc_tmtrafficpolicy.tf_tmtrafficpolicy.name
	priority 	= 1
	bindpoint   = "REQUEST"
}

data "citrixadc_csvserver_tmtrafficpolicy_binding" "test" {
	name 		= "tf_csvserver"
	policyname	= "tf_tmttrafficpolicy"
	depends_on = [citrixadc_csvserver_tmtrafficpolicy_binding.tf_csvserver_tmtrafficpolicy_binding]
}
`
