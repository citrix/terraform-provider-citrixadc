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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccRewriteaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRewriteactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRewriteaction_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewriteactionExist("citrixadc_rewriteaction.tf_rewrite_action", nil),
				),
			},
			{
				Config: testAccRewriteaction_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewriteactionExist("citrixadc_rewriteaction.tf_rewrite_action", nil),
				),
			},
		},
	})
}

func TestAccRewriteaction_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_rewriteaction.tf_rewrite_action"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRewriteactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRewriteaction_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckRewriteactionExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Rewriteaction.Type(), "tf_rewrite_action"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccRewriteaction_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckRewriteactionExist(resAddr, nil)),
			},
		},
	})
}

func testAccCheckRewriteactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lb vserver name is set")
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
		data, err := client.FindResource(service.Rewriteaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckRewriteactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_rewriteaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Rewriteaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccRewriteaction_step1 = `

resource "citrixadc_rewriteaction" "tf_rewrite_action" {
    name = "tf_rewrite_action"
    target = "HTTP.REQ.HOSTNAME"
    type = "delete"
}
`

const testAccRewriteaction_step2 = `

resource "citrixadc_rewriteaction" "tf_rewrite_action" {
    name = "tf_rewrite_action"
    target = "HTTP.REQ.COOKIE"
    type = "delete"
}
`

func TestAccRewriteaction_import(t *testing.T) {
	const resAddr = "citrixadc_rewriteaction.tf_rewrite_action"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRewriteactionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccRewriteaction_step1},
			{
				Config:                  testAccRewriteaction_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccRewriteaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckRewriteactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccRewriteaction_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckRewriteactionExist("citrixadc_rewriteaction.tf_rewrite_action", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccRewriteaction_step1,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckRewriteactionExist("citrixadc_rewriteaction.tf_rewrite_action", nil)),
			},
		},
	})
}

const testAccRewriteactionDataSource_basic = `
resource "citrixadc_rewriteaction" "tf_rewrite_action_ds" {
    name = "tf_rewrite_action_ds"
    target = "HTTP.REQ.HOSTNAME"
    type = "delete"
    comment = "datasource test comment"
}

data "citrixadc_rewriteaction" "tf_rewrite_action_ds" {
  name = citrixadc_rewriteaction.tf_rewrite_action_ds.name
}
`

// The rewriteaction unset test covers the unset-eligible attributes that can be
// both set and cleanly unset: comment and refinesearch. It uses a REPLACE_ALL
// action so refinesearch (valid only on the *_ALL body-expression action types)
// is settable and optional; the mandatory REPLACE_ALL arguments (target, search,
// stringbuilderexpr) are kept across both steps. stringbuilderexpr itself is in
// the NITRO unset payload but is mandatory whenever it is settable, so it cannot
// be exercised on a single-resource unset test and is excluded here.
const testAccRewriteaction_unset_step1 = `
resource "citrixadc_rewriteaction" "tf_unset" {
  name              = "tf_test_rewriteaction_unset"
  type              = "replace_all"
  target            = "HTTP.RES.BODY(1000)"
  search            = "text(\"hello\")"
  stringbuilderexpr = "\"world\""
  refinesearch      = "extend(10, 20)"
  comment           = "tf unset test comment"
}
`

const testAccRewriteaction_unset_step2 = `
resource "citrixadc_rewriteaction" "tf_unset" {
  name              = "tf_test_rewriteaction_unset"
  type              = "replace_all"
  target            = "HTTP.RES.BODY(1000)"
  search            = "text(\"hello\")"
  stringbuilderexpr = "\"world\""
  # comment and refinesearch removed from config -> the provider must unset them
  # (revert to NITRO defaults, empty).
}
`

func TestAccRewriteaction_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRewriteactionDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccRewriteaction_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewriteactionExist("citrixadc_rewriteaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_rewriteaction.tf_unset", "comment", "tf unset test comment"),
					resource.TestCheckResourceAttr("citrixadc_rewriteaction.tf_unset", "refinesearch", "extend(10, 20)"),
				),
			},
			{
				// Removing the attributes must unset them: NITRO omits the unset
				// fields from GET so state reads back NULL (no Default to inject a
				// value), and the implicit post-apply plan must be empty.
				Config: testAccRewriteaction_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewriteactionExist("citrixadc_rewriteaction.tf_unset", nil),
					resource.TestCheckNoResourceAttr("citrixadc_rewriteaction.tf_unset", "comment"),
					resource.TestCheckNoResourceAttr("citrixadc_rewriteaction.tf_unset", "refinesearch"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckRewriteactionADCValue("tf_test_rewriteaction_unset", "comment", ""),
					testAccCheckRewriteactionADCValue("tf_test_rewriteaction_unset", "refinesearch", ""),
				),
			},
		},
	})
}

// testAccCheckRewriteactionADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it. A missing/nil value is treated as the empty-string default.
func testAccCheckRewriteactionADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Rewriteaction.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("rewriteaction %s not found on appliance", name)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("rewriteaction %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccRewriteactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRewriteactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_rewriteaction.tf_rewrite_action_ds", "name", "tf_rewrite_action_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_rewriteaction.tf_rewrite_action_ds", "type", "delete"),
					resource.TestCheckResourceAttr("data.citrixadc_rewriteaction.tf_rewrite_action_ds", "target", "HTTP.REQ.HOSTNAME"),
					resource.TestCheckResourceAttr("data.citrixadc_rewriteaction.tf_rewrite_action_ds", "comment", "datasource test comment"),
				),
			},
		},
	})
}
