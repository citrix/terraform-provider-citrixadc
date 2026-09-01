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
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testAccTransformpolicy_basic_step1 = `
resource "citrixadc_transformprofile" "tf_trans_profile1" {
  name = "tf_trans_profile1"
}

resource "citrixadc_transformprofile" "tf_trans_profile2" {
  name = "tf_trans_profile2"
}

resource "citrixadc_transformpolicy" "tf_trans_policy" {
    name = "tf_trans_policy"
    profilename = citrixadc_transformprofile.tf_trans_profile1.name
    rule = "http.REQ.URL.CONTAINS(\"test_url\")"
}
`

const testAccTransformpolicy_basic_step2 = `
resource "citrixadc_transformprofile" "tf_trans_profile1" {
  name = "tf_trans_profile1"
}

resource "citrixadc_transformprofile" "tf_trans_profile2" {
  name = "tf_trans_profile2"
}

resource "citrixadc_transformpolicy" "tf_trans_policy" {
    name = "tf_trans_policy"
    profilename = citrixadc_transformprofile.tf_trans_profile2.name
    rule = "http.REQ.URL.CONTAINS(\"test_url_other\")"
}
`

func TestAccTransformpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTransformpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTransformpolicy_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTransformpolicyExist("citrixadc_transformpolicy.tf_trans_policy", nil),
				),
			},
			{
				Config: testAccTransformpolicy_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTransformpolicyExist("citrixadc_transformpolicy.tf_trans_policy", nil),
				),
			},
		},
	})
}

func TestAccTransformpolicy_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_transformpolicy.tf_trans_policy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTransformpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTransformpolicy_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTransformpolicyExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Transformpolicy.Type(), "tf_trans_policy"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccTransformpolicy_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTransformpolicyExist(resAddr, nil)),
			},
		},
	})
}

func testAccCheckTransformpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No transformpolicy name is set")
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
		data, err := client.FindResource(service.Transformpolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("transformpolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckTransformpolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_transformpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Transformpolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("transformpolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccTransformpolicy_import(t *testing.T) {
	const resAddr = "citrixadc_transformpolicy.tf_trans_policy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTransformpolicyDestroy,
		Steps: []resource.TestStep{
			{Config: testAccTransformpolicy_basic_step1},
			{
				Config:                  testAccTransformpolicy_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccTransformpolicy_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckTransformpolicyDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccTransformpolicy_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTransformpolicyExist("citrixadc_transformpolicy.tf_trans_policy", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccTransformpolicy_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTransformpolicyExist("citrixadc_transformpolicy.tf_trans_policy", nil)),
			},
		},
	})
}

// The transformpolicy unset test covers the spec-unsettable attributes
// (comment, logaction). Step 1 sets them to non-default values; step 2 removes
// them from config so the provider issues the NITRO ?action=unset, reverting
// them to their appliance defaults (empty).
const testAccTransformpolicy_unset_step1 = `
resource "citrixadc_transformprofile" "tf_trans_profile1" {
  name = "tf_trans_profile1"
}

resource "citrixadc_auditmessageaction" "tf_unset_msgaction" {
  name              = "tf_unset_msgaction"
  loglevel          = "NOTICE"
  stringbuilderexpr = "\"hello\""
  logtonewnslog     = "YES"
}

resource "citrixadc_transformpolicy" "tf_trans_policy_unset" {
  name        = "tf_trans_policy_unset"
  profilename = citrixadc_transformprofile.tf_trans_profile1.name
  rule        = "http.REQ.URL.CONTAINS(\"test_url\")"
  comment     = "managed by terraform"
  logaction   = citrixadc_auditmessageaction.tf_unset_msgaction.name
}
`

const testAccTransformpolicy_unset_step2 = `
resource "citrixadc_transformprofile" "tf_trans_profile1" {
  name = "tf_trans_profile1"
}

resource "citrixadc_auditmessageaction" "tf_unset_msgaction" {
  name              = "tf_unset_msgaction"
  loglevel          = "NOTICE"
  stringbuilderexpr = "\"hello\""
  logtonewnslog     = "YES"
}

resource "citrixadc_transformpolicy" "tf_trans_policy_unset" {
  name        = "tf_trans_policy_unset"
  profilename = citrixadc_transformprofile.tf_trans_profile1.name
  rule        = "http.REQ.URL.CONTAINS(\"test_url\")"
  # comment and logaction removed -> provider must unset them.
}
`

func TestAccTransformpolicy_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTransformpolicyDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccTransformpolicy_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTransformpolicyExist("citrixadc_transformpolicy.tf_trans_policy_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_transformpolicy.tf_trans_policy_unset", "comment", "managed by terraform"),
					resource.TestCheckResourceAttr("citrixadc_transformpolicy.tf_trans_policy_unset", "logaction", "tf_unset_msgaction"),
				),
			},
			{
				// Removing the attributes must unset them: the appliance reverts
				// them to their defaults (empty) and the implicit post-apply plan
				// must be empty.
				Config: testAccTransformpolicy_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTransformpolicyExist("citrixadc_transformpolicy.tf_trans_policy_unset", nil),
					testAccCheckTransformpolicyADCValue("tf_trans_policy_unset", "comment", ""),
					testAccCheckTransformpolicyADCValue("tf_trans_policy_unset", "logaction", ""),
				),
			},
		},
	})
}

// testAccCheckTransformpolicyADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset reverted it.
func testAccCheckTransformpolicyADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Transformpolicy.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("transformpolicy %s not found on appliance", name)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("transformpolicy %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccTransformpolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccTransformpolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_transformpolicy.tf_trans_policy", "name", "tf_trans_policy"),
					resource.TestCheckResourceAttr("data.citrixadc_transformpolicy.tf_trans_policy", "profilename", "tf_trans_profile1"),
					resource.TestCheckResourceAttr("data.citrixadc_transformpolicy.tf_trans_policy", "rule", "http.REQ.URL.CONTAINS(\"test_url\")"),
					resource.TestCheckResourceAttrSet("data.citrixadc_transformpolicy.tf_trans_policy", "id"),
					// Read-only metadata exposed only by the data source.
					resource.TestCheckResourceAttrSet("data.citrixadc_transformpolicy.tf_trans_policy", "hits"),
				),
			},
		},
	})
}

const testAccTransformpolicyDataSource_basic = `
resource "citrixadc_transformprofile" "tf_trans_profile1" {
  name = "tf_trans_profile1"
}

resource "citrixadc_transformpolicy" "tf_trans_policy" {
    name = "tf_trans_policy"
    profilename = citrixadc_transformprofile.tf_trans_profile1.name
    rule = "http.REQ.URL.CONTAINS(\"test_url\")"
}

data "citrixadc_transformpolicy" "tf_trans_policy" {
    name = citrixadc_transformpolicy.tf_trans_policy.name
    depends_on = [citrixadc_transformpolicy.tf_trans_policy]
}
`
