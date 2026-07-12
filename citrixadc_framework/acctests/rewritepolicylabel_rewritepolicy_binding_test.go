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

const testAccRewritepolicylabel_rewritepolicy_binding_basic = `

resource "citrixadc_rewritepolicylabel_rewritepolicy_binding" "tf_rewritepolicylabel_rewritepolicy_binding" {
	labelname = citrixadc_rewritepolicylabel.tf_rewritepolicylabel.labelname
	policyname = citrixadc_rewritepolicy.tf_rewrite_policy.name
	gotopriorityexpression = "END"
	priority = 5   
}

resource "citrixadc_rewritepolicylabel" "tf_rewritepolicylabel" {
	labelname = "tf_rewritepolicylabel"
	transform = "http_req"
}

resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
	name = "tf_rewrite_policy"
	action = "DROP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"
}
`

const testAccRewritepolicylabel_rewritepolicy_binding_basic_step2 = `

	resource "citrixadc_rewritepolicylabel" "tf_rewritepolicylabel" {
		labelname = "tf_rewritepolicylabel"
		transform = "http_req"
	}

	resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
		name = "tf_rewrite_policy"
		action = "DROP"
		rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"
	}
`

func TestAccRewritepolicylabel_rewritepolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_rewritepolicylabel_rewritepolicy_binding.tf_rewritepolicylabel_rewritepolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRewritepolicylabel_rewritepolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccRewritepolicylabel_rewritepolicy_binding_basic},
			{Config: testAccRewritepolicylabel_rewritepolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func TestAccRewritepolicylabel_rewritepolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRewritepolicylabel_rewritepolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRewritepolicylabel_rewritepolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewritepolicylabel_rewritepolicy_bindingExist("citrixadc_rewritepolicylabel_rewritepolicy_binding.tf_rewritepolicylabel_rewritepolicy_binding", nil),
				),
			},
			{
				Config: testAccRewritepolicylabel_rewritepolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewritepolicylabel_rewritepolicy_bindingNotExist("citrixadc_rewritepolicylabel_rewritepolicy_binding.tf_rewritepolicylabel_rewritepolicy_binding", "tf_rewritepolicylabel,tf_rewrite_policy"),
				),
			},
		},
	})
}

func testAccCheckRewritepolicylabel_rewritepolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No rewritepolicylabel_rewritepolicy_binding id is set")
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

		// ID-parse: support both the new key:value ID format and the legacy comma format.
		idMap, _, err := utils.ParseIdString(bindingId, []string{"labelname", "policyname"}, nil)
		if err != nil {
			return err
		}
		name := idMap["labelname"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "rewritepolicylabel_rewritepolicy_binding",
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
			return fmt.Errorf("rewritepolicylabel_rewritepolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckRewritepolicylabel_rewritepolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		// ID-parse: support both the new key:value ID format and the legacy comma format.
		idMap, _, err := utils.ParseIdString(id, []string{"labelname", "policyname"}, nil)
		if err != nil {
			return err
		}
		name := idMap["labelname"]
		policyname := idMap["policyname"]

		findParams := service.FindParams{
			ResourceType:             "rewritepolicylabel_rewritepolicy_binding",
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
			return fmt.Errorf("rewritepolicylabel_rewritepolicy_binding %s was found, but is should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckRewritepolicylabel_rewritepolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_rewritepolicylabel_rewritepolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Rewritepolicylabel_rewritepolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("rewritepolicylabel_rewritepolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccRewritepolicylabel_rewritepolicy_bindingDataSource_basic = `

resource "citrixadc_rewritepolicylabel" "tf_rewritepolicylabel" {
	labelname = "tf_rewritepolicylabel"
	transform = "http_req"
}

resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
	name = "tf_rewrite_policy"
	action = "DROP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"
}

resource "citrixadc_rewritepolicylabel_rewritepolicy_binding" "tf_rewritepolicylabel_rewritepolicy_binding" {
	labelname = citrixadc_rewritepolicylabel.tf_rewritepolicylabel.labelname
	policyname = citrixadc_rewritepolicy.tf_rewrite_policy.name
	gotopriorityexpression = "END"
	priority = 5   
}

data "citrixadc_rewritepolicylabel_rewritepolicy_binding" "tf_rewritepolicylabel_rewritepolicy_binding" {
	labelname = citrixadc_rewritepolicylabel_rewritepolicy_binding.tf_rewritepolicylabel_rewritepolicy_binding.labelname
	policyname = citrixadc_rewritepolicylabel_rewritepolicy_binding.tf_rewritepolicylabel_rewritepolicy_binding.policyname
	depends_on = [citrixadc_rewritepolicylabel_rewritepolicy_binding.tf_rewritepolicylabel_rewritepolicy_binding]
}
`

func TestAccRewritepolicylabel_rewritepolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRewritepolicylabel_rewritepolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_rewritepolicylabel_rewritepolicy_binding.tf_rewritepolicylabel_rewritepolicy_binding", "labelname", "tf_rewritepolicylabel"),
					resource.TestCheckResourceAttr("data.citrixadc_rewritepolicylabel_rewritepolicy_binding.tf_rewritepolicylabel_rewritepolicy_binding", "policyname", "tf_rewrite_policy"),
					resource.TestCheckResourceAttr("data.citrixadc_rewritepolicylabel_rewritepolicy_binding.tf_rewritepolicylabel_rewritepolicy_binding", "priority", "5"),
					resource.TestCheckResourceAttr("data.citrixadc_rewritepolicylabel_rewritepolicy_binding.tf_rewritepolicylabel_rewritepolicy_binding", "gotopriorityexpression", "END"),
				),
			},
		},
	})
}

// testAccRewritepolicylabel_rewritepolicy_binding_upgrade_basic mirrors the _basic config
// but is written so it is valid under BOTH the last SDK v2 release (2.2.0) schema and the
// current Framework schema, so the same HCL can be applied by each provider across the two
// steps of the state-upgrade test.
const testAccRewritepolicylabel_rewritepolicy_binding_upgrade_basic = `

resource "citrixadc_rewritepolicylabel_rewritepolicy_binding" "tf_rewritepolicylabel_rewritepolicy_binding" {
	labelname = citrixadc_rewritepolicylabel.tf_rewritepolicylabel.labelname
	policyname = citrixadc_rewritepolicy.tf_rewrite_policy.name
	gotopriorityexpression = "END"
	priority = 5
}

resource "citrixadc_rewritepolicylabel" "tf_rewritepolicylabel" {
	labelname = "tf_rewritepolicylabel"
	transform = "http_req"
}

resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
	name = "tf_rewrite_policy"
	action = "DROP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"
}
`

// TestAccRewritepolicylabel_rewritepolicy_binding_sdkv2StateUpgrade verifies that state
// written by the last SDK v2 provider release (legacy comma-separated ID) is read and
// upgraded cleanly by the current Framework provider.
//
// Step 1 creates the binding with citrix/citrixadc 2.2.0, which writes the legacy ID
// "<labelname>,<policyname>".
// Step 2 re-applies the SAME config through the current (Framework) provider. Its Read
// parses the legacy ID (via ParseIdString), then SetAttrFromGet re-derives the canonical
// new-format ID "labelname:<v>,policyname:<v>,priority:<v>".
func TestAccRewritepolicylabel_rewritepolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		// Providers are specified per-step (ExternalProviders in step 1, framework factories
		// in step 2), so they must NOT also be set at the TestCase level.
		CheckDestroy: testAccCheckRewritepolicylabel_rewritepolicy_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release (writes the legacy comma ID).
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccRewritepolicylabel_rewritepolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewritepolicylabel_rewritepolicy_bindingExist("citrixadc_rewritepolicylabel_rewritepolicy_binding.tf_rewritepolicylabel_rewritepolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_rewritepolicylabel_rewritepolicy_binding.tf_rewritepolicylabel_rewritepolicy_binding", "id", "tf_rewritepolicylabel,tf_rewrite_policy"),
				),
			},
			// Step 2: same config through the current Framework provider; Read upgrades the ID.
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccRewritepolicylabel_rewritepolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewritepolicylabel_rewritepolicy_bindingExist("citrixadc_rewritepolicylabel_rewritepolicy_binding.tf_rewritepolicylabel_rewritepolicy_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_rewritepolicylabel_rewritepolicy_binding.tf_rewritepolicylabel_rewritepolicy_binding", "id", "labelname:tf_rewritepolicylabel,policyname:tf_rewrite_policy,priority:5"),
				),
			},
		},
	})
}
