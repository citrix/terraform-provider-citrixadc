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

const testAccAppfwprofile_xmlxss_binding_basic = `
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
	resource "citrixadc_appfwprofile_xmlxss_binding" "tf_binding1" {
		name                    = citrixadc_appfwprofile.tf_appfwprofile.name
		xmlxss                  = "tf_xmlxss"
		state                   = "ENABLED"
		alertonly               = "ON"
		isregex_xmlxss          = "NOTREGEX"
		isautodeployed          = "AUTODEPLOYED"
	}
	resource "citrixadc_appfwprofile_xmlxss_binding" "tf_binding2" {
		name                    = citrixadc_appfwprofile.tf_appfwprofile.name
		xmlxss                  = "new_tf_xmlxss"
		state                   = "ENABLED"
		alertonly               = "ON"
		isregex_xmlxss          = "NOTREGEX"
		isautodeployed          = "AUTODEPLOYED"
	}
`

const testAccAppfwprofile_xmlxss_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
`

func TestAccAppfwprofile_xmlxss_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofile_xmlxss_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_xmlxss_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_xmlxss_bindingExist("citrixadc_appfwprofile_xmlxss_binding.tf_binding1", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlxss_binding.tf_binding1", "name", "tf_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlxss_binding.tf_binding1", "xmlxss", "tf_xmlxss"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlxss_binding.tf_binding1", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlxss_binding.tf_binding1", "alertonly", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlxss_binding.tf_binding1", "isautodeployed", "AUTODEPLOYED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlxss_binding.tf_binding1", "isregex_xmlxss", "NOTREGEX"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlxss_binding.tf_binding1", "as_scan_location_xmlxss", "ELEMENT"),
					testAccCheckAppfwprofile_xmlxss_bindingExist("citrixadc_appfwprofile_xmlxss_binding.tf_binding2", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlxss_binding.tf_binding2", "name", "tf_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlxss_binding.tf_binding2", "xmlxss", "new_tf_xmlxss"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlxss_binding.tf_binding2", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlxss_binding.tf_binding2", "alertonly", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlxss_binding.tf_binding2", "isautodeployed", "AUTODEPLOYED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlxss_binding.tf_binding2", "isregex_xmlxss", "NOTREGEX"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlxss_binding.tf_binding2", "as_scan_location_xmlxss", "ELEMENT"),
				),
			},
			{
				Config: testAccAppfwprofile_xmlxss_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_xmlxss_bindingNotExist("citrixadc_appfwprofile_xmlxss_binding.tf_binding", "tf_appfwprofile,tf_xmlxss,ELEMENT"),
				),
			},
		},
	})
}

func testAccCheckAppfwprofile_xmlxss_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_xmlxss_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "xmlxss", "as_scan_location_xmlxss"}, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}
		name := idMap["name"]
		xmlxss := idMap["xmlxss"]
		as_scan_location_xmlxss := idMap["as_scan_location_xmlxss"]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_xmlxss_binding",
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
			if v["xmlxss"].(string) == xmlxss {
				if v["as_scan_location_xmlxss"].(string) == as_scan_location_xmlxss {
					found = true
					break
				}
			}
		}

		if !found {
			return fmt.Errorf("appfwprofile_xmlxss_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_xmlxss_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 3)

		name := idSlice[0]
		xmlxss := idSlice[1]
		as_scan_location_xmlxss := idSlice[2]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_xmlxss_binding",
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
			if v["xmlxss"].(string) == xmlxss {
				if v["as_scan_location_xmlxss"].(string) == as_scan_location_xmlxss {
					found = true
					break
				}
			}
		}

		if found {
			return fmt.Errorf("appfwprofile_xmlxss_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_xmlxss_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_xmlxss_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Appfwprofile_xmlxss_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwprofile_xmlxss_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAppfwprofileXmlxssBindingDataSource_basic = `
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
	resource "citrixadc_appfwprofile_xmlxss_binding" "tf_binding1" {
		name                    = citrixadc_appfwprofile.tf_appfwprofile.name
		xmlxss                  = "tf_xmlxss"
		state                   = "ENABLED"
		alertonly               = "ON"
		isregex_xmlxss          = "NOTREGEX"
		isautodeployed          = "AUTODEPLOYED"
	}

	data "citrixadc_appfwprofile_xmlxss_binding" "tf_binding_data" {
		name                    = citrixadc_appfwprofile_xmlxss_binding.tf_binding1.name
		xmlxss                  = citrixadc_appfwprofile_xmlxss_binding.tf_binding1.xmlxss
		as_scan_location_xmlxss = citrixadc_appfwprofile_xmlxss_binding.tf_binding1.as_scan_location_xmlxss
		depends_on              = [citrixadc_appfwprofile_xmlxss_binding.tf_binding1]
	}
`

func TestAccAppfwprofileXmlxssBindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofileXmlxssBindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_xmlxss_binding.tf_binding_data", "name", "tf_appfwprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_xmlxss_binding.tf_binding_data", "xmlxss", "tf_xmlxss"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_xmlxss_binding.tf_binding_data", "state", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_xmlxss_binding.tf_binding_data", "alertonly", "OFF"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_xmlxss_binding.tf_binding_data", "isautodeployed", "NOTAUTODEPLOYED"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_xmlxss_binding.tf_binding_data", "isregex_xmlxss", "NOTREGEX"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_xmlxss_binding.tf_binding_data", "as_scan_location_xmlxss", "ELEMENT"),
				),
			},
		},
	})
}

const testAccAppfwprofile_xmlxss_binding_upgrade_basic = `
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
	resource "citrixadc_appfwprofile_xmlxss_binding" "tf_binding1" {
		name                    = citrixadc_appfwprofile.tf_appfwprofile.name
		xmlxss                  = "tf_xmlxss"
		state                   = "ENABLED"
		alertonly               = "ON"
		isregex_xmlxss          = "NOTREGEX"
		isautodeployed          = "AUTODEPLOYED"
	}
`

func TestAccAppfwprofile_xmlxss_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAppfwprofile_xmlxss_bindingDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create the binding with the last SDK v2 release (2.2.0),
				// which writes state using the legacy comma-joined id.
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.0.0",
					},
				},
				Config: testAccAppfwprofile_xmlxss_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_xmlxss_bindingExist("citrixadc_appfwprofile_xmlxss_binding.tf_binding1", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlxss_binding.tf_binding1", "id", "tf_appfwprofile,tf_xmlxss,ELEMENT"),
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
				Config: testAccAppfwprofile_xmlxss_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_xmlxss_bindingExist("citrixadc_appfwprofile_xmlxss_binding.tf_binding1", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_xmlxss_binding.tf_binding1", "id", "as_scan_location_xmlxss:ELEMENT,name:tf_appfwprofile,xmlxss:tf_xmlxss"),
				),
			},
		},
	})
}

func TestAccAppfwprofile_xmlxss_binding_import(t *testing.T) {
	const resAddr = "citrixadc_appfwprofile_xmlxss_binding.tf_binding1"

	// Backward-compat: import via the LEGACY SDK v2 id. Rebuild the legacy positional id from
	// the current canonical key:value id (raw values, only the keys actually set, in legacy
	// order: name,xmlxss,as_scan_location_xmlxss) so it matches exactly what SDK v2 wrote.
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
		ordr := []string{"name", "xmlxss", "as_scan_location_xmlxss"}
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
		CheckDestroy:             testAccCheckAppfwprofile_xmlxss_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAppfwprofile_xmlxss_binding_basic},
			{Config: testAccAppfwprofile_xmlxss_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"alertonly", "isautodeployed"}},
			{Config: testAccAppfwprofile_xmlxss_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"alertonly", "isautodeployed"}},
		},
	})
}

func TestAccAppfwprofile_xmlxss_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_appfwprofile_xmlxss_binding.tf_binding1"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofile_xmlxss_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_xmlxss_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppfwprofile_xmlxss_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgs(service.Appfwprofile_xmlxss_binding.Type(), "tf_appfwprofile", []string{"xmlxss:tf_xmlxss", "as_scan_location_xmlxss:ELEMENT"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAppfwprofile_xmlxss_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppfwprofile_xmlxss_bindingExist(resAddr, nil)),
			},
		},
	})
}
