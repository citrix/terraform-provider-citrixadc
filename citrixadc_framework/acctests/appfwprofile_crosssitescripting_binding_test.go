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

const testAccAppfwprofile_crosssitescripting_binding_basic = `

resource "citrixadc_appfwprofile_crosssitescripting_binding" "demo_binding1" {
  name                 = citrixadc_appfwprofile.demo_appfw.name
  crosssitescripting   = "file"
  isregex_xss          = "NOTREGEX"
  formactionurl_xss    = "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$"
  as_scan_location_xss = "FORMFIELD"
  as_value_type_xss    = "Tag"
  isvalueregex_xss     = "REGEX"
  as_value_expr_xss    = ".*"
  state                = "ENABLED"
}

resource "citrixadc_appfwprofile_crosssitescripting_binding" "demo_binding2" {
  name                 = citrixadc_appfwprofile.demo_appfw.name
  crosssitescripting   = "file"
  isregex_xss          = "NOTREGEX"
  formactionurl_xss    = "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/v1/resource/temp$"
  as_scan_location_xss = "FORMFIELD"
  as_value_type_xss    = "Tag"
  isvalueregex_xss     = "REGEX"
  as_value_expr_xss    = ".*"
  state                = "ENABLED"
}

resource "citrixadc_appfwprofile" "demo_appfw" {
	name                     = "demo_appfwprofile"
	type                     = ["HTML"]
  }
`

func TestAccAppfwprofile_crosssitescripting_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofile_crosssitescripting_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_crosssitescripting_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_crosssitescripting_bindingExist("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding1", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding1", "name", "demo_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding1", "crosssitescripting", "file"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding1", "formactionurl_xss", "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding1", "as_scan_location_xss", "FORMFIELD"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding1", "isregex_xss", "NOTREGEX"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding1", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding1", "as_value_type_xss", "Tag"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding1", "isvalueregex_xss", "REGEX"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding1", "as_value_expr_xss", ".*"),

					// Check second binding
					testAccCheckAppfwprofile_crosssitescripting_bindingExist("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding2", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding2", "name", "demo_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding2", "crosssitescripting", "file"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding2", "formactionurl_xss", "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/v1/resource/temp$"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding2", "as_scan_location_xss", "FORMFIELD"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding2", "isregex_xss", "NOTREGEX"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding2", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding2", "as_value_type_xss", "Tag"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding2", "isvalueregex_xss", "REGEX"),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding2", "as_value_expr_xss", ".*"),
				),
			},
		},
	})
}

func testAccCheckAppfwprofile_crosssitescripting_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_crosssitescripting_binding name is set")
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
		idMap, _, err := utils.ParseIdString(bindingId, []string{"name", "crosssitescripting", "formactionurl_xss", "as_scan_location_xss", "as_value_type_xss", "as_value_expr_xss"}, []string{"as_value_type_xss", "as_value_expr_xss"})
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", bindingId, err)
		}
		appFwName := idMap["name"]
		crosssitescripting := idMap["crosssitescripting"]
		formactionurl_xss := idMap["formactionurl_xss"]
		as_scan_location_xss := idMap["as_scan_location_xss"]
		as_value_type_xss := idMap["as_value_type_xss"]
		as_value_expr_xss := idMap["as_value_expr_xss"]

		findParams := service.FindParams{
			ResourceType:             service.Appfwprofile_crosssitescripting_binding.Type(),
			ResourceName:             appFwName,
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
			if v["crosssitescripting"].(string) == crosssitescripting {
				if v["formactionurl_xss"] != nil && v["as_scan_location_xss"] != nil && v["as_value_type_xss"] != nil && v["as_value_expr_xss"] != nil && v["as_value_type_xss"].(string) == as_value_type_xss && v["as_value_expr_xss"].(string) == as_value_expr_xss && v["as_scan_location_xss"].(string) == as_scan_location_xss && v["formactionurl_xss"].(string) == formactionurl_xss {
					foundIndex = i
					break
				}
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			return fmt.Errorf("Cannot find appfwprofile_crosssitescripting_binding ID %v", bindingId)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_crosssitescripting_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_crosssitescripting_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Appfwprofile_crosssitescripting_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwprofile_crosssitescripting_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAppfwprofile_crosssitescripting_bindingDataSource_basic = `
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
	resource "citrixadc_appfwprofile_crosssitescripting_binding" "tf_binding1" {
		name                 = citrixadc_appfwprofile.tf_appfwprofile.name
		crosssitescripting   = "file"
		isregex_xss          = "NOTREGEX"
		formactionurl_xss    = "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$"
		as_scan_location_xss = "FORMFIELD"
		as_value_type_xss    = "Tag"
		isvalueregex_xss     = "REGEX"
		as_value_expr_xss    = ".*"
		state                = "ENABLED"
	}

	data "citrixadc_appfwprofile_crosssitescripting_binding" "tf_binding1" {
		name                 = citrixadc_appfwprofile_crosssitescripting_binding.tf_binding1.name
		crosssitescripting   = citrixadc_appfwprofile_crosssitescripting_binding.tf_binding1.crosssitescripting
		formactionurl_xss    = citrixadc_appfwprofile_crosssitescripting_binding.tf_binding1.formactionurl_xss
		as_scan_location_xss = citrixadc_appfwprofile_crosssitescripting_binding.tf_binding1.as_scan_location_xss
		as_value_type_xss    = citrixadc_appfwprofile_crosssitescripting_binding.tf_binding1.as_value_type_xss
		as_value_expr_xss    = citrixadc_appfwprofile_crosssitescripting_binding.tf_binding1.as_value_expr_xss
	}
`

func TestAccAppfwprofile_crosssitescripting_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_crosssitescripting_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_crosssitescripting_binding.tf_binding1", "name", "tf_appfwprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_crosssitescripting_binding.tf_binding1", "crosssitescripting", "file"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_crosssitescripting_binding.tf_binding1", "formactionurl_xss", "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_crosssitescripting_binding.tf_binding1", "as_scan_location_xss", "FORMFIELD"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_crosssitescripting_binding.tf_binding1", "as_value_type_xss", "Tag"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_crosssitescripting_binding.tf_binding1", "as_value_expr_xss", ".*"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_crosssitescripting_binding.tf_binding1", "isregex_xss", "NOTREGEX"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_crosssitescripting_binding.tf_binding1", "isvalueregex_xss", "REGEX"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwprofile_crosssitescripting_binding.tf_binding1", "state", "ENABLED"),
				),
			},
		},
	})
}

const testAccAppfwprofile_crosssitescripting_binding_upgrade_basic = `

resource "citrixadc_appfwprofile_crosssitescripting_binding" "demo_binding1" {
  name                 = citrixadc_appfwprofile.demo_appfw.name
  crosssitescripting   = "file"
  isregex_xss          = "NOTREGEX"
  formactionurl_xss    = "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$"
  as_scan_location_xss = "FORMFIELD"
  as_value_type_xss    = "Tag"
  isvalueregex_xss     = "REGEX"
  as_value_expr_xss    = ".*"
  state                = "ENABLED"
}

resource "citrixadc_appfwprofile" "demo_appfw" {
	name                     = "demo_appfwprofile"
	type                     = ["HTML"]
  }
`

func TestAccAppfwprofile_crosssitescripting_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAppfwprofile_crosssitescripting_bindingDestroy,
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
				Config: testAccAppfwprofile_crosssitescripting_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_crosssitescripting_bindingExist("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding1", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding1", "id", "demo_appfwprofile,file,^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$,FORMFIELD,Tag,.*"),
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
				Config: testAccAppfwprofile_crosssitescripting_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_crosssitescripting_bindingExist("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding1", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprofile_crosssitescripting_binding.demo_binding1", "id", "as_scan_location_xss:FORMFIELD,as_value_expr_xss:.%2A,as_value_type_xss:Tag,crosssitescripting:file,formactionurl_xss:%5Ehttps%3A%2F%2Fsd2%5C-zgw%5C.test%5C.ctxns%5C.com%2Fapi%2Fdocument%2Fcontent%24,name:demo_appfwprofile"),
				),
			},
		},
	})
}

func TestAccAppfwprofile_crosssitescripting_binding_import(t *testing.T) {
	const resAddr = "citrixadc_appfwprofile_crosssitescripting_binding.demo_binding1"

	// Backward-compat: import via the LEGACY SDK v2 id. Rebuild the legacy positional id from
	// the current canonical key:value id (raw values, only the keys actually set, in legacy
	// order: name,crosssitescripting,formactionurl_xss,as_scan_location_xss,as_value_type_xss,as_value_expr_xss) so it matches exactly what SDK v2 wrote.
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
		ordr := []string{"name", "crosssitescripting", "formactionurl_xss", "as_scan_location_xss", "as_value_type_xss", "as_value_expr_xss"}
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
		CheckDestroy:             testAccCheckAppfwprofile_crosssitescripting_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAppfwprofile_crosssitescripting_binding_basic},
			{Config: testAccAppfwprofile_crosssitescripting_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
			{Config: testAccAppfwprofile_crosssitescripting_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func TestAccAppfwprofile_crosssitescripting_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_appfwprofile_crosssitescripting_binding.demo_binding1"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofile_crosssitescripting_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_crosssitescripting_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppfwprofile_crosssitescripting_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgsMap(service.Appfwprofile_crosssitescripting_binding.Type(), "demo_appfwprofile", map[string]string{
						"as_scan_location_xss": utils.UrlEncode("FORMFIELD"),
						"as_value_expr_xss":    utils.UrlEncode(".*"),
						"as_value_type_xss":    utils.UrlEncode("Tag"),
						"crosssitescripting":   utils.UrlEncode("file"),
						"formactionurl_xss":    utils.UrlEncode("^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$"),
					}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAppfwprofile_crosssitescripting_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppfwprofile_crosssitescripting_bindingExist(resAddr, nil)),
			},
		},
	})
}
