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

const testAccCrvserver_appfwpolicy_binding_basic = `

resource "citrixadc_crvserver" "crvserver" {
    name = "my_vserver"
    servicetype = "HTTP"
    arp = "OFF"
}
resource "citrixadc_appfwprofile" "demo_appfwprofile" {
    name = "demo_appfwprofile"
    bufferoverflowaction = ["none"]
    contenttypeaction = ["none"]
    cookieconsistencyaction = ["none"]
    creditcard = ["none"]
    creditcardaction = ["none"]
    crosssitescriptingaction = ["none"]
    csrftagaction = ["none"]
    denyurlaction = ["none"]
    dynamiclearning = ["none"]
    fieldconsistencyaction = ["none"]
    fieldformataction = ["none"]
    fileuploadtypesaction = ["none"]
    inspectcontenttypes = ["none"]
    jsondosaction = ["none"]
    jsonsqlinjectionaction = ["none"]
    jsonxssaction = ["none"]
    multipleheaderaction = ["none"]
    sqlinjectionaction = ["none"]
    starturlaction = ["none"]
    type = ["HTML"]
    xmlattachmentaction = ["none"]
    xmldosaction = ["none"]
    xmlformataction = ["none"]
    xmlsoapfaultaction = ["none"]
    xmlsqlinjectionaction = ["none"]
    xmlvalidationaction = ["none"]
    xmlwsiaction = ["none"]
    xmlxssaction = ["none"]
}

resource "citrixadc_appfwpolicy" "demo_appfwpolicy1" {
    name = "demo_appfwpolicy1"
    profilename = citrixadc_appfwprofile.demo_appfwprofile.name
    rule = "true"
}

resource "citrixadc_crvserver_appfwpolicy_binding" "crvserver_appfwpolicy_binding" {
    name = citrixadc_crvserver.crvserver.name
    policyname = citrixadc_appfwpolicy.demo_appfwpolicy1.name
    priority = 20
}
`

const testAccCrvserver_appfwpolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_crvserver" "crvserver" {
		name = "my_vserver"
		servicetype = "HTTP"
		arp = "OFF"
	}
	resource "citrixadc_appfwprofile" "demo_appfwprofile" {
		name = "demo_appfwprofile"
		bufferoverflowaction = ["none"]
		contenttypeaction = ["none"]
		cookieconsistencyaction = ["none"]
		creditcard = ["none"]
		creditcardaction = ["none"]
		crosssitescriptingaction = ["none"]
		csrftagaction = ["none"]
		denyurlaction = ["none"]
		dynamiclearning = ["none"]
		fieldconsistencyaction = ["none"]
		fieldformataction = ["none"]
		fileuploadtypesaction = ["none"]
		inspectcontenttypes = ["none"]
		jsondosaction = ["none"]
		jsonsqlinjectionaction = ["none"]
		jsonxssaction = ["none"]
		multipleheaderaction = ["none"]
		sqlinjectionaction = ["none"]
		starturlaction = ["none"]
		type = ["HTML"]
		xmlattachmentaction = ["none"]
		xmldosaction = ["none"]
		xmlformataction = ["none"]
		xmlsoapfaultaction = ["none"]
		xmlsqlinjectionaction = ["none"]
		xmlvalidationaction = ["none"]
		xmlwsiaction = ["none"]
		xmlxssaction = ["none"]
	}
	
	resource "citrixadc_appfwpolicy" "demo_appfwpolicy1" {
		name = "demo_appfwpolicy1"
		profilename = citrixadc_appfwprofile.demo_appfwprofile.name
		rule = "true"
	}
`

func TestAccCrvserver_appfwpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCrvserver_appfwpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCrvserver_appfwpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrvserver_appfwpolicy_bindingExist("citrixadc_crvserver_appfwpolicy_binding.crvserver_appfwpolicy_binding", nil),
				),
			},
			{
				Config: testAccCrvserver_appfwpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrvserver_appfwpolicy_bindingNotExist("citrixadc_crvserver_appfwpolicy_binding.crvserver_appfwpolicy_binding", "my_vserver,demo_appfwpolicy1"),
				),
			},
		},
	})
}

func testAccCheckCrvserver_appfwpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No crvserver_appfwpolicy_binding id is set")
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
			ResourceType:             "crvserver_appfwpolicy_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching policyname
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("crvserver_appfwpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckCrvserver_appfwpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(id, []string{"name", "policyname"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", id, err)
		}
		name := idMap["name"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "crvserver_appfwpolicy_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching policyname
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("crvserver_appfwpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckCrvserver_appfwpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_crvserver_appfwpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Crvserver_appfwpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("crvserver_appfwpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccCrvserver_appfwpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_crvserver_appfwpolicy_binding.crvserver_appfwpolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCrvserver_appfwpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccCrvserver_appfwpolicy_binding_basic},
			{Config: testAccCrvserver_appfwpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

const testAccCrvserver_appfwpolicy_bindingDataSource_basic = `

	resource "citrixadc_crvserver" "crvserver" {
		name = "my_vserver_ds"
		servicetype = "HTTP"
		arp = "OFF"
	}
	resource "citrixadc_appfwpolicy" "tf_appfwpolicy" {
		name      = "tf_appfwpolicy_ds"
		rule      = "true"
		profilename = "APPFW_BYPASS"
	}
	
	resource "citrixadc_crvserver_appfwpolicy_binding" "crvserver_appfwpolicy_binding" {
		name = citrixadc_crvserver.crvserver.name
		policyname = citrixadc_appfwpolicy.tf_appfwpolicy.name
		priority = 1
	}

	data "citrixadc_crvserver_appfwpolicy_binding" "crvserver_appfwpolicy_binding" {
		name       = citrixadc_crvserver_appfwpolicy_binding.crvserver_appfwpolicy_binding.name
		policyname = citrixadc_crvserver_appfwpolicy_binding.crvserver_appfwpolicy_binding.policyname
		depends_on = [citrixadc_crvserver_appfwpolicy_binding.crvserver_appfwpolicy_binding]
	}
`

func TestAcccrvserver_appfwpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccCrvserver_appfwpolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_crvserver_appfwpolicy_binding.crvserver_appfwpolicy_binding", "name", "my_vserver_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_crvserver_appfwpolicy_binding.crvserver_appfwpolicy_binding", "policyname", "tf_appfwpolicy_ds"),
				),
			},
		},
	})
}

const testAccCrvserver_appfwpolicy_binding_upgrade_basic = `

resource "citrixadc_crvserver" "crvserver" {
    name = "my_vserver"
    servicetype = "HTTP"
    arp = "OFF"
}
resource "citrixadc_appfwprofile" "demo_appfwprofile" {
    name = "demo_appfwprofile"
    bufferoverflowaction = ["none"]
    contenttypeaction = ["none"]
    cookieconsistencyaction = ["none"]
    creditcard = ["none"]
    creditcardaction = ["none"]
    crosssitescriptingaction = ["none"]
    csrftagaction = ["none"]
    denyurlaction = ["none"]
    dynamiclearning = ["none"]
    fieldconsistencyaction = ["none"]
    fieldformataction = ["none"]
    fileuploadtypesaction = ["none"]
    inspectcontenttypes = ["none"]
    jsondosaction = ["none"]
    jsonsqlinjectionaction = ["none"]
    jsonxssaction = ["none"]
    multipleheaderaction = ["none"]
    sqlinjectionaction = ["none"]
    starturlaction = ["none"]
    type = ["HTML"]
    xmlattachmentaction = ["none"]
    xmldosaction = ["none"]
    xmlformataction = ["none"]
    xmlsoapfaultaction = ["none"]
    xmlsqlinjectionaction = ["none"]
    xmlvalidationaction = ["none"]
    xmlwsiaction = ["none"]
    xmlxssaction = ["none"]
}

resource "citrixadc_appfwpolicy" "demo_appfwpolicy1" {
    name = "demo_appfwpolicy1"
    profilename = citrixadc_appfwprofile.demo_appfwprofile.name
    rule = "true"
}

resource "citrixadc_crvserver_appfwpolicy_binding" "crvserver_appfwpolicy_binding" {
    name = citrixadc_crvserver.crvserver.name
    policyname = citrixadc_appfwpolicy.demo_appfwpolicy1.name
    priority = 20
}
`

func TestAccCrvserver_appfwpolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckCrvserver_appfwpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create the binding with the last SDK v2 release (2.2.0),
				// which writes state using the legacy comma-joined id.
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccCrvserver_appfwpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrvserver_appfwpolicy_bindingExist("citrixadc_crvserver_appfwpolicy_binding.crvserver_appfwpolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_crvserver_appfwpolicy_binding.crvserver_appfwpolicy_binding", "id", "my_vserver,demo_appfwpolicy1"),
				),
			},
			{
				// Step 2: refresh/plan the legacy-id state through the current
				// framework provider. Read exercises ParseIdString on the legacy id
				// and SetAttrFromGet recomputes the id into the new key:value form.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccCrvserver_appfwpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCrvserver_appfwpolicy_bindingExist("citrixadc_crvserver_appfwpolicy_binding.crvserver_appfwpolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_crvserver_appfwpolicy_binding.crvserver_appfwpolicy_binding", "id", "name:my_vserver,policyname:demo_appfwpolicy1"),
				),
			},
		},
	})
}
