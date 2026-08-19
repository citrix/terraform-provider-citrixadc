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
	"net/url"
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testAccAppfwprofile_xmlattachmenturl_binding_basic = `
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
	resource "citrixadc_appfwprofile_xmlattachmenturl_binding" "tf_binding" {
		name                          = citrixadc_appfwprofile.tf_appfwprofile.name
		xmlattachmenturl              = ".*"
		xmlattachmentcontenttype      = "abc*"
		alertonly                     = "ON"
		state                         = "ENABLED"
		isautodeployed                = "AUTODEPLOYED"
		comment                       = "Testing"
		xmlattachmentcontenttypecheck = "ON"
		xmlmaxattachmentsize          = "1000"
		xmlmaxattachmentsizecheck     = "ON"
	}
`

const testAccAppfwprofile_xmlattachmenturl_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
`

func TestAccAppfwprofile_xmlattachmenturl_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofile_xmlattachmenturl_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_xmlattachmenturl_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_xmlattachmenturl_bindingExist("citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "name", "tf_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "xmlattachmenturl", ".*"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "xmlattachmentcontenttype", "abc*"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "alertonly", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "isautodeployed", "AUTODEPLOYED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "comment", "Testing"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "xmlattachmentcontenttypecheck", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "xmlmaxattachmentsize", "1000"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "xmlmaxattachmentsizecheck", "ON"),
				),
			},
			{
				Config: testAccAppfwprofile_xmlattachmenturl_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_xmlattachmenturl_bindingNotExist("citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "tf_appfwprofile,.*"),
				),
			},
		},
	})
}

func testAccCheckAppfwprofile_xmlattachmenturl_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_xmlattachmenturl_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "xmlattachmenturl"}, nil)
		if err != nil {
			return err
		}
		name := idMap["name"]
		xmlattachmenturl := idMap["xmlattachmenturl"]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_xmlattachmenturl_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["xmlattachmenturl"].(string) == xmlattachmenturl {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("appfwprofile_xmlattachmenturl_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_xmlattachmenturl_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		idMap, _, err := utils.ParseIdString(id, []string{"name", "xmlattachmenturl"}, nil)
		if err != nil {
			return err
		}
		name := idMap["name"]
		xmlattachmenturl := idMap["xmlattachmenturl"]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_xmlattachmenturl_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["xmlattachmenturl"].(string) == xmlattachmenturl {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("appfwprofile_xmlattachmenturl_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_xmlattachmenturl_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_xmlattachmenturl_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Appfwprofile_xmlattachmenturl_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwprofile_xmlattachmenturl_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAppfwprofileXmlattachmenturlBindingDataSource_basic = `
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
	resource "citrixadc_appfwprofile_xmlattachmenturl_binding" "tf_binding" {
		name                          = citrixadc_appfwprofile.tf_appfwprofile.name
		xmlattachmenturl              = ".*"
		xmlattachmentcontenttype      = "abc*"
		alertonly                     = "ON"
		state                         = "ENABLED"
		isautodeployed                = "AUTODEPLOYED"
		comment                       = "Testing"
		xmlattachmentcontenttypecheck = "ON"
		xmlmaxattachmentsize          = "1000"
		xmlmaxattachmentsizecheck     = "ON"
	}

	data "citrixadc_appfwprofile_xmlattachmenturl_binding" "tf_binding" {
		name             = citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding.name
		xmlattachmenturl = citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding.xmlattachmenturl
		depends_on       = [citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding]
	}
`

func TestAccAppfwprofileXmlattachmenturlBindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofileXmlattachmenturlBindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "name", "tf_appfwprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "xmlattachmenturl", ".*"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "xmlattachmentcontenttype", "abc*"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "alertonly", "OFF"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "state", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "isautodeployed", "NOTAUTODEPLOYED"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "comment", "Testing"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "xmlattachmentcontenttypecheck", "ON"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "xmlmaxattachmentsize", "1000"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "xmlmaxattachmentsizecheck", "ON"),
				),
			},
		},
	})
}

const testAccAppfwprofile_xmlattachmenturl_binding_upgrade_basic = `
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
	resource "citrixadc_appfwprofile_xmlattachmenturl_binding" "tf_binding" {
		name                          = citrixadc_appfwprofile.tf_appfwprofile.name
		xmlattachmenturl              = ".*"
		xmlattachmentcontenttype      = "abc*"
		alertonly                     = "ON"
		state                         = "ENABLED"
		isautodeployed                = "AUTODEPLOYED"
		comment                       = "Testing"
		xmlattachmentcontenttypecheck = "ON"
		xmlmaxattachmentsize          = "1000"
		xmlmaxattachmentsizecheck     = "ON"
	}
`

func TestAccAppfwprofile_xmlattachmenturl_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAppfwprofile_xmlattachmenturl_bindingDestroy,
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
				Config: testAccAppfwprofile_xmlattachmenturl_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_xmlattachmenturl_bindingExist("citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "id", "tf_appfwprofile,.*"),
				),
			},
			{
				// Step 2: refresh/plan the legacy-id state through the current
				// framework provider. Read exercises ParseIdString on the legacy id
				// and SetAttrFromGet recomputes the id into the new key:value form.
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccAppfwprofile_xmlattachmenturl_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_xmlattachmenturl_bindingExist("citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding", "id", "name:tf_appfwprofile,xmlattachmenturl:.%2A"),
				),
			},
		},
	})
}

func TestAccAppfwprofile_xmlattachmenturl_binding_import(t *testing.T) {
	const resAddr = "citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding"

	// Backward-compat: import via the LEGACY SDK v2 id. Rebuild the legacy positional id from
	// the current canonical key:value id (raw values, only the keys actually set, in legacy
	// order: name,xmlattachmenturl) so it matches exactly what SDK v2 wrote.
	legacyID := func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resAddr]
		if !ok {
			return "", fmt.Errorf("resource not found in state: %s", resAddr)
		}
		kv := map[string]string{}
		for _, p := range strings.Split(rs.Primary.ID, ",") {
			if i := strings.Index(p, ":"); i >= 0 {
				v, _ := url.QueryUnescape(p[i+1:])
				kv[p[:i]] = v
			}
		}
		ordr := []string{"name", "xmlattachmenturl"}
		parts := make([]string, 0, len(ordr))
		for _, k := range ordr {
			if v, ok := kv[k]; ok {
				parts = append(parts, v)
			}
		}
		// Fallback: a positional (non key:value) id has no key:value parts to reorder; import it as-is.
		if len(parts) == 0 {
			return rs.Primary.ID, nil
		}
		return strings.Join(parts, ","), nil
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofile_xmlattachmenturl_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAppfwprofile_xmlattachmenturl_binding_basic},
			{Config: testAccAppfwprofile_xmlattachmenturl_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"alertonly", "isautodeployed"}},
			{Config: testAccAppfwprofile_xmlattachmenturl_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"alertonly", "isautodeployed"}},
		},
	})
}

func TestAccAppfwprofile_xmlattachmenturl_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_appfwprofile_xmlattachmenturl_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofile_xmlattachmenturl_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_xmlattachmenturl_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppfwprofile_xmlattachmenturl_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgsMap(service.Appfwprofile_xmlattachmenturl_binding.Type(), "tf_appfwprofile", map[string]string{"xmlattachmenturl": ".%2A"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAppfwprofile_xmlattachmenturl_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppfwprofile_xmlattachmenturl_bindingExist(resAddr, nil)),
			},
		},
	})
}
