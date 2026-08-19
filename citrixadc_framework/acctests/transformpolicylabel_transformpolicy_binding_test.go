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

const testAccTransformpolicylabel_transformpolicy_binding_basic = `
resource "citrixadc_transformprofile" "tf_trans_profile1" {
	name = "pro_1"
	}
  resource "citrixadc_transformpolicy" "tf_trans_policy" {
	  name = "tf_trans_policy"
	  profilename = citrixadc_transformprofile.tf_trans_profile1.name
	  rule = "http.REQ.URL.CONTAINS(\"test_url\")"
	}
  resource "citrixadc_transformpolicylabel" "transformpolicylabel" {
	labelname = "label_1"
	policylabeltype = "httpquic_req"
	}
  resource "citrixadc_transformpolicylabel_transformpolicy_binding" "transformpolicylabel_transformpolicy_binding"{
	 policyname = citrixadc_transformpolicy.tf_trans_policy.name
	  labelname = citrixadc_transformpolicylabel.transformpolicylabel.labelname
	  priority = 2
	}
`

const testAccTransformpolicylabel_transformpolicy_binding_basic_step2 = `
resource "citrixadc_transformprofile" "tf_trans_profile1" {
	name = "pro_1"
	}
  
  resource "citrixadc_transformpolicy" "tf_trans_policy" {
	  name = "tf_trans_policy"
	  profilename = citrixadc_transformprofile.tf_trans_profile1.name
	  rule = "http.REQ.URL.CONTAINS(\"test_url\")"
	}
  resource "citrixadc_transformpolicylabel" "transformpolicylabel" {
	labelname = "label_1"
	policylabeltype = "httpquic_req"
	}
`

func TestAccTransformpolicylabel_transformpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTransformpolicylabel_transformpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTransformpolicylabel_transformpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTransformpolicylabel_transformpolicy_bindingExist("citrixadc_transformpolicylabel_transformpolicy_binding.transformpolicylabel_transformpolicy_binding", nil),
				),
			},
			{
				Config: testAccTransformpolicylabel_transformpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTransformpolicylabel_transformpolicy_bindingNotExist("citrixadc_transformpolicylabel_transformpolicy_binding.transformpolicylabel_transformpolicy_binding", "label_1,tf_trans_policy"),
				),
			},
		},
	})
}

func TestAccTransformpolicylabel_transformpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_transformpolicylabel_transformpolicy_binding.transformpolicylabel_transformpolicy_binding"

	// Backward-compat: import via the LEGACY SDK v2 id. Rebuild the legacy positional id from
	// the current canonical key:value id (raw values, only the keys actually set, in legacy
	// order: labelname,policyname) so it matches exactly what SDK v2 wrote.
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
		ordr := []string{"labelname", "policyname"}
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
		CheckDestroy:             testAccCheckTransformpolicylabel_transformpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccTransformpolicylabel_transformpolicy_binding_basic},
			{Config: testAccTransformpolicylabel_transformpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
			{Config: testAccTransformpolicylabel_transformpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateIdFunc: legacyID, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func testAccCheckTransformpolicylabel_transformpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No transformpolicylabel_transformpolicy_binding id is set")
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

		idMap, _, err := utils.ParseIdString(bindingId, []string{"labelname", "policyname"}, nil)
		if err != nil {
			return err
		}
		labelname := idMap["labelname"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "transformpolicylabel_transformpolicy_binding",
			ResourceName:             labelname,
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
			return fmt.Errorf("transformpolicylabel_transformpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckTransformpolicylabel_transformpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
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

		labelname := idSlice[0]
		policyname := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "transformpolicylabel_transformpolicy_binding",
			ResourceName:             labelname,
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
			return fmt.Errorf("transformpolicylabel_transformpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckTransformpolicylabel_transformpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_transformpolicylabel_transformpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Transformpolicylabel_transformpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("transformpolicylabel_transformpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccTransformpolicylabel_transformpolicy_bindingDataSource_basic = `
resource "citrixadc_transformprofile" "tf_trans_profile1" {
	name = "pro_1"
	}
  resource "citrixadc_transformpolicy" "tf_trans_policy" {
	  name = "tf_trans_policy"
	  profilename = citrixadc_transformprofile.tf_trans_profile1.name
	  rule = "http.REQ.URL.CONTAINS(\"test_url\")"
	}
  resource "citrixadc_transformpolicylabel" "transformpolicylabel" {
	labelname = "label_1"
	policylabeltype = "httpquic_req"
	}
  resource "citrixadc_transformpolicylabel_transformpolicy_binding" "transformpolicylabel_transformpolicy_binding"{
	 policyname = citrixadc_transformpolicy.tf_trans_policy.name
	  labelname = citrixadc_transformpolicylabel.transformpolicylabel.labelname
	  priority = 2
	}

data "citrixadc_transformpolicylabel_transformpolicy_binding" "transformpolicylabel_transformpolicy_binding" {
	labelname  = citrixadc_transformpolicylabel_transformpolicy_binding.transformpolicylabel_transformpolicy_binding.labelname
	policyname = citrixadc_transformpolicylabel_transformpolicy_binding.transformpolicylabel_transformpolicy_binding.policyname
	depends_on = [citrixadc_transformpolicylabel_transformpolicy_binding.transformpolicylabel_transformpolicy_binding]
}
`

func TestAccTransformpolicylabel_transformpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTransformpolicylabel_transformpolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_transformpolicylabel_transformpolicy_binding.transformpolicylabel_transformpolicy_binding", "labelname", "label_1"),
					resource.TestCheckResourceAttr("data.citrixadc_transformpolicylabel_transformpolicy_binding.transformpolicylabel_transformpolicy_binding", "policyname", "tf_trans_policy"),
					resource.TestCheckResourceAttr("data.citrixadc_transformpolicylabel_transformpolicy_binding.transformpolicylabel_transformpolicy_binding", "priority", "2"),
				),
			},
		},
	})
}

const testAccTransformpolicylabel_transformpolicy_binding_upgrade_basic = `
resource "citrixadc_transformprofile" "tf_trans_profile1" {
	name = "pro_1"
	}
  resource "citrixadc_transformpolicy" "tf_trans_policy" {
	  name = "tf_trans_policy"
	  profilename = citrixadc_transformprofile.tf_trans_profile1.name
	  rule = "http.REQ.URL.CONTAINS(\"test_url\")"
	}
  resource "citrixadc_transformpolicylabel" "transformpolicylabel" {
	labelname = "label_1"
	policylabeltype = "httpquic_req"
	}
  resource "citrixadc_transformpolicylabel_transformpolicy_binding" "transformpolicylabel_transformpolicy_binding"{
	 policyname = citrixadc_transformpolicy.tf_trans_policy.name
	  labelname = citrixadc_transformpolicylabel.transformpolicylabel.labelname
	  priority = 2
	}
`

func TestAccTransformpolicylabel_transformpolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckTransformpolicylabel_transformpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create the resource with the last SDK v2 release (writes legacy id).
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccTransformpolicylabel_transformpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTransformpolicylabel_transformpolicy_bindingExist("citrixadc_transformpolicylabel_transformpolicy_binding.transformpolicylabel_transformpolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_transformpolicylabel_transformpolicy_binding.transformpolicylabel_transformpolicy_binding", "id", "label_1,tf_trans_policy"),
				),
			},
			// Step 2: refresh/plan/apply the legacy-id state through the current framework provider.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccTransformpolicylabel_transformpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTransformpolicylabel_transformpolicy_bindingExist("citrixadc_transformpolicylabel_transformpolicy_binding.transformpolicylabel_transformpolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_transformpolicylabel_transformpolicy_binding.transformpolicylabel_transformpolicy_binding", "id", "labelname:label_1,policyname:tf_trans_policy"),
				),
			},
		},
	})
}

func TestAccTransformpolicylabel_transformpolicy_binding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_transformpolicylabel_transformpolicy_binding.transformpolicylabel_transformpolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTransformpolicylabel_transformpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTransformpolicylabel_transformpolicy_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTransformpolicylabel_transformpolicy_bindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgsMap(service.Transformpolicylabel_transformpolicy_binding.Type(), "label_1", map[string]string{"policyname": "tf_trans_policy", "priority": "2"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccTransformpolicylabel_transformpolicy_binding_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTransformpolicylabel_transformpolicy_bindingExist(resAddr, nil)),
			},
		},
	})
}
