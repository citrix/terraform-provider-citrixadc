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

const testAccLbvserver_cmppolicy_binding_basic_step1 = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}


resource "citrixadc_cmppolicy" "tf_cmppolicy" {
    name = "tf_cmppolicy"
    rule = "HTTP.RES.HEADER(\"Content-Type\").CONTAINS(\"text\")"
    resaction = "COMPRESS"
}

resource "citrixadc_lbvserver_cmppolicy_binding" "tf_bind" {
    name = citrixadc_lbvserver.tf_lbvserver.name
    policyname = citrixadc_cmppolicy.tf_cmppolicy.name
    priority = 100
    bindpoint = "RESPONSE"
    gotopriorityexpression = "END"
}

`

const testAccLbvserver_cmppolicy_binding_basic_step2 = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}


resource "citrixadc_cmppolicy" "tf_cmppolicy" {
    name = "tf_cmppolicy"
    rule = "HTTP.RES.HEADER(\"Content-Type\").CONTAINS(\"text\")"
    resaction = "COMPRESS"
}

resource "citrixadc_lbvserver_cmppolicy_binding" "tf_bind" {
    name = citrixadc_lbvserver.tf_lbvserver.name
    policyname = citrixadc_cmppolicy.tf_cmppolicy.name
    priority = 110
    bindpoint = "RESPONSE"
    gotopriorityexpression = "END"
}

`

func TestAccLbvserver_cmppolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbvserver_cmppolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbvserver_cmppolicy_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_cmppolicy_bindingExist("citrixadc_lbvserver_cmppolicy_binding.tf_bind", nil),
				),
			},
			{
				Config: testAccLbvserver_cmppolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_cmppolicy_bindingExist("citrixadc_lbvserver_cmppolicy_binding.tf_bind", nil),
				),
			},
		},
	})
}

func testAccCheckLbvserver_cmppolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbvserver_cmppolicy_binding name is set")
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
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}
		name := idMap["name"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "lbvserver_cmppolicy_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the right policy name
		foundIndex := -1
		for i, v := range dataArr {
			if v["policyname"].(string) == policyname {
				foundIndex = i
				break
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			return fmt.Errorf("Cannot find lbvserver_cmppolicy_binding %s", bindingId)
		}

		return nil
	}
}

func testAccCheckLbvserver_cmppolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbvserver_cmppolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Lbvserver_cmppolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lbvserver_cmppolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLbvserver_cmppolicy_bindingDataSource_basic = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}


resource "citrixadc_cmppolicy" "tf_cmppolicy" {
    name = "tf_cmppolicy"
    rule = "HTTP.RES.HEADER(\"Content-Type\").CONTAINS(\"text\")"
    resaction = "COMPRESS"
}

resource "citrixadc_lbvserver_cmppolicy_binding" "tf_bind" {
    name = citrixadc_lbvserver.tf_lbvserver.name
    policyname = citrixadc_cmppolicy.tf_cmppolicy.name
    priority = 100
    bindpoint = "RESPONSE"
    gotopriorityexpression = "END"
}

data "citrixadc_lbvserver_cmppolicy_binding" "tf_bind" {
    name = citrixadc_lbvserver_cmppolicy_binding.tf_bind.name
    policyname = citrixadc_lbvserver_cmppolicy_binding.tf_bind.policyname
    bindpoint = citrixadc_lbvserver_cmppolicy_binding.tf_bind.bindpoint
    depends_on = [citrixadc_lbvserver_cmppolicy_binding.tf_bind]
}
`

func TestAccLbvserver_cmppolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLbvserver_cmppolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_cmppolicy_binding.tf_bind", "name", "tf_lbvserver"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_cmppolicy_binding.tf_bind", "policyname", "tf_cmppolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_cmppolicy_binding.tf_bind", "bindpoint", "RESPONSE"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_cmppolicy_binding.tf_bind", "priority", "100"),
					resource.TestCheckResourceAttr("data.citrixadc_lbvserver_cmppolicy_binding.tf_bind", "gotopriorityexpression", "END"),
				),
			},
		},
	})
}

// Config for the SDK v2 -> Framework state-upgrade test. Reuses the _basic_step1
// values and is valid under BOTH the last SDK v2 release (2.2.0) schema and the
// current Framework schema (uses only SDK v2 attribute names).
const testAccLbvserver_cmppolicy_binding_upgrade_basic = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}


resource "citrixadc_cmppolicy" "tf_cmppolicy" {
    name = "tf_cmppolicy"
    rule = "HTTP.RES.HEADER(\"Content-Type\").CONTAINS(\"text\")"
    resaction = "COMPRESS"
}

resource "citrixadc_lbvserver_cmppolicy_binding" "tf_bind" {
    name = citrixadc_lbvserver.tf_lbvserver.name
    policyname = citrixadc_cmppolicy.tf_cmppolicy.name
    priority = 100
    bindpoint = "RESPONSE"
    gotopriorityexpression = "END"
}

`

func TestAccLbvserver_cmppolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_lbvserver_cmppolicy_binding.tf_bind"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbvserver_cmppolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLbvserver_cmppolicy_binding_basic_step1},
			{Config: testAccLbvserver_cmppolicy_binding_basic_step1, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

// TestAccLbvserver_cmppolicy_binding_sdkv2StateUpgrade verifies that state written
// by the last SDK v2 release (legacy comma-joined id) is transparently upgraded by
// the current Framework provider. Step 1 creates the binding with citrix/citrixadc
// 2.2.0 (legacy id "name,policyname"); step 2 refreshes/plans the same config
// through the current Framework provider, whose Read parses the legacy id and
// recomputes it to the new "name:<v>,policyname:<v>" format (SetAttrFromGet).
func TestAccLbvserver_cmppolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLbvserver_cmppolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create with the last SDK v2 release, writing the legacy id.
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccLbvserver_cmppolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_cmppolicy_bindingExist("citrixadc_lbvserver_cmppolicy_binding.tf_bind", nil),
					resource.TestCheckResourceAttr("citrixadc_lbvserver_cmppolicy_binding.tf_bind", "id", "tf_lbvserver,tf_cmppolicy"),
				),
			},
			{
				// Step 2: refresh/apply the same config through the current Framework
				// provider. Read exercises ParseIdString on the legacy id, then
				// recomputes the id to the new key:value canonical format.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccLbvserver_cmppolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbvserver_cmppolicy_bindingExist("citrixadc_lbvserver_cmppolicy_binding.tf_bind", nil),
					resource.TestCheckResourceAttr("citrixadc_lbvserver_cmppolicy_binding.tf_bind", "id", "name:tf_lbvserver,policyname:tf_cmppolicy"),
				),
			},
		},
	})
}
