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

const testAccLsnclient_nsacl_binding_basic = `

resource "citrixadc_lsnclient" "tf_lsnclient" {
	clientname = "my_lsn_client"
}

resource "citrixadc_nsacl" "foo" {

  aclaction = "ALLOW"
  aclname = "test_acl"
  destipval = "192.168.1.33"
  protocol = "TCP"
  srcportval = "45-1024"
  priority = "100"

}

resource "citrixadc_lsnclient_nsacl_binding" "tf_lsnclient_nsacl_binding" {
	clientname = citrixadc_lsnclient.tf_lsnclient.clientname
	aclname    = citrixadc_nsacl.foo.aclname
}
`

const testAccLsnclient_nsacl_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
`

const testAccLsnclient_nsacl_bindingDataSource_basic = `

resource "citrixadc_lsnclient" "tf_lsnclient" {
	clientname = "my_lsn_client"
}

resource "citrixadc_nsacl" "foo" {

  aclaction = "ALLOW"
  aclname = "test_acl"
  destipval = "192.168.1.33"
  protocol = "TCP"
  srcportval = "45-1024"
  priority = "100"

}

resource "citrixadc_lsnclient_nsacl_binding" "tf_lsnclient_nsacl_binding" {
	clientname = citrixadc_lsnclient.tf_lsnclient.clientname
	aclname    = citrixadc_nsacl.foo.aclname
}

data "citrixadc_lsnclient_nsacl_binding" "tf_lsnclient_nsacl_binding" {
	clientname = citrixadc_lsnclient_nsacl_binding.tf_lsnclient_nsacl_binding.clientname
	aclname    = citrixadc_lsnclient_nsacl_binding.tf_lsnclient_nsacl_binding.aclname
	depends_on = [citrixadc_lsnclient_nsacl_binding.tf_lsnclient_nsacl_binding]
}
`

func TestAccLsnclient_nsacl_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnclient_nsacl_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnclient_nsacl_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnclient_nsacl_bindingExist("citrixadc_lsnclient_nsacl_binding.tf_lsnclient_nsacl_binding", nil),
				),
			},
			{
				Config: testAccLsnclient_nsacl_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnclient_nsacl_bindingNotExist("citrixadc_lsnclient_nsacl_binding.tf_lsnclient_nsacl_binding", "my_lsnclient,my_acl"),
				),
			},
		},
	})
}

func testAccCheckLsnclient_nsacl_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsnclient_nsacl_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"clientname", "aclname"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID: %v", err)
		}
		clientname := idMap["clientname"]
		aclname := idMap["aclname"]

		findParams := service.FindParams{
			ResourceType:             "lsnclient_nsacl_binding",
			ResourceName:             clientname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching aclname
		found := false
		for _, v := range dataArr {
			if v["aclname"].(string) == aclname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("lsnclient_nsacl_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsnclient_nsacl_bindingNotExist(n string, id string) resource.TestCheckFunc {
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

		clientname := idSlice[0]
		aclname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "lsnclient_nsacl_binding",
			ResourceName:             clientname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching aclname
		found := false
		for _, v := range dataArr {
			if v["aclname"].(string) == aclname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("lsnclient_nsacl_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckLsnclient_nsacl_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsnclient_nsacl_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lsnclient_nsacl_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lsnclient_nsacl_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccLsnclient_nsacl_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnclient_nsacl_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lsnclient_nsacl_binding.tf_lsnclient_nsacl_binding", "clientname", "my_lsn_client"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnclient_nsacl_binding.tf_lsnclient_nsacl_binding", "aclname", "test_acl"),
				),
			},
		},
	})
}

// testAccLsnclient_nsacl_binding_upgrade_basic reuses the _basic config (binding +
// all prerequisite resources). It is valid under BOTH the SDK v2 2.2.0 schema and the
// current Framework schema because the migration restored the SDK v2 attribute names.
const testAccLsnclient_nsacl_binding_upgrade_basic = `

resource "citrixadc_lsnclient" "tf_lsnclient" {
	clientname = "my_lsn_client"
}

resource "citrixadc_nsacl" "foo" {

  aclaction = "ALLOW"
  aclname = "test_acl"
  destipval = "192.168.1.33"
  protocol = "TCP"
  srcportval = "45-1024"
  priority = "100"

}

resource "citrixadc_lsnclient_nsacl_binding" "tf_lsnclient_nsacl_binding" {
	clientname = citrixadc_lsnclient.tf_lsnclient.clientname
	aclname    = citrixadc_nsacl.foo.aclname
}
`

// TestAccLsnclient_nsacl_binding_sdkv2StateUpgrade verifies that state written by the
// last SDK v2 release (legacy comma-separated ID) is correctly upgraded when the same
// config is subsequently managed by the current Framework provider. Step 1 creates the
// binding with citrix/citrixadc 2.2.0 (writes the legacy id "my_lsn_client,test_acl").
// Step 2 refreshes/plans/applies the same config through the Framework provider,
// exercising ParseIdString on the legacy id; because the Framework recomputes the id on
// Read (SetAttrFromGet), the id upgrades to the new "key:value" form.
func TestAccLsnclient_nsacl_binding_sdkv2StateUpgrade(t *testing.T) {
	resourceAddr := "citrixadc_lsnclient_nsacl_binding.tf_lsnclient_nsacl_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLsnclient_nsacl_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> state carries the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccLsnclient_nsacl_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnclient_nsacl_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "my_lsn_client,test_acl"),
				),
			},
			// Step 2: refresh/plan/apply the SAME config through the current Framework
			// provider. The legacy-id state is read via ParseIdString and the id is
			// recomputed to the new key:value format.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccLsnclient_nsacl_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnclient_nsacl_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "clientname:my_lsn_client,aclname:test_acl"),
				),
			},
		},
	})
}

func TestAccLsnclient_nsacl_binding_import(t *testing.T) {
	const resAddr = "citrixadc_lsnclient_nsacl_binding.tf_lsnclient_nsacl_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnclient_nsacl_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLsnclient_nsacl_binding_basic},
			{Config: testAccLsnclient_nsacl_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}
