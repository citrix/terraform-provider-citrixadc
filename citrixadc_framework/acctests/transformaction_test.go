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

const testAccTransformaction_basic_step1 = `
resource "citrixadc_transformprofile" "tf_trans_profile1" {
  name = "tf_trans_profile1"
}

resource "citrixadc_transformprofile" "tf_trans_profile2" {
  name = "tf_trans_profile2"
}

resource "citrixadc_transformaction" "tf_trans_action" {
  name = "tf_trans_action"
  profilename = citrixadc_transformprofile.tf_trans_profile1.name
  priority = 100
  requrlfrom = "http://m3.mydomain.com/(.*)"
  requrlinto = "https://exp-proxy-v1.api.mydomain.com/$1"
  resurlfrom = "https://exp-proxy-v1.api.mydomain.com/(.*)"
  resurlinto = "https://m3.mydomain.com/$1"
}
`

const testAccTransformaction_basic_step2 = `
resource "citrixadc_transformprofile" "tf_trans_profile1" {
  name = "tf_trans_profile1"
}

resource "citrixadc_transformprofile" "tf_trans_profile2" {
  name = "tf_trans_profile2"
}

resource "citrixadc_transformaction" "tf_trans_action" {
  name = "tf_trans_action"
  profilename = citrixadc_transformprofile.tf_trans_profile2.name
  priority = 100
  requrlfrom = "http://m4.mydomain.com/(.*)"
  requrlinto = "https://exp-proxy-v1.api.mydomain.com/$1"
  resurlfrom = "https://exp-proxy-v1.api.mydomain.com/(.*)"
  resurlinto = "https://m4.mydomain.com/$1"
}
`

const testAccTransformaction_basic_step3 = `
resource "citrixadc_transformprofile" "tf_trans_profile1" {
  name = "tf_trans_profile1"
}

resource "citrixadc_transformprofile" "tf_trans_profile2" {
  name = "tf_trans_profile2"
}

resource "citrixadc_transformaction" "tf_trans_action" {
  name = "tf_trans_action"
  profilename = citrixadc_transformprofile.tf_trans_profile2.name
  priority = 110
  requrlfrom = "http://m5.mydomain.com/(.*)"
  requrlinto = "https://exp-proxy-v1.api.mydomain.com/$1"
  resurlfrom = "https://exp-proxy-v1.api.mydomain.com/(.*)"
  resurlinto = "https://m5.mydomain.com/$1"
}
`

func TestAccTransformaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTransformactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTransformaction_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTransformactionExist("citrixadc_transformaction.tf_trans_action", nil),
				),
			},
			{
				Config: testAccTransformaction_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTransformactionExist("citrixadc_transformaction.tf_trans_action", nil),
				),
			},
			{
				Config: testAccTransformaction_basic_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTransformactionExist("citrixadc_transformaction.tf_trans_action", nil),
				),
			},
		},
	})
}

func TestAccTransformaction_import(t *testing.T) {
	const resAddr = "citrixadc_transformaction.tf_trans_action"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTransformactionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccTransformaction_basic_step1},
			{
				Config:                  testAccTransformaction_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckTransformactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No transformaction name is set")
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
		data, err := client.FindResource(service.Transformaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("transformaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckTransformactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_transformaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Transformaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("transformaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccTransformaction_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_transformaction.tf_trans_action"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTransformactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTransformaction_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTransformactionExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Transformaction.Type(), "tf_trans_action"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccTransformaction_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTransformactionExist(resAddr, nil)),
			},
		},
	})
}

func TestAccTransformaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckTransformactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccTransformaction_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTransformactionExist("citrixadc_transformaction.tf_trans_action", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccTransformaction_basic_step1,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckTransformactionExist("citrixadc_transformaction.tf_trans_action", nil)),
			},
		},
	})
}

// testAccCheckTransformactionADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it. An attribute that NITRO omits after unset is treated as "".
func testAccCheckTransformactionADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Transformaction.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("transformaction %s not found on appliance", name)
		}
		var got string
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("transformaction %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

const testAccTransformaction_unset_step1 = `
resource "citrixadc_transformprofile" "tf_trans_profile1" {
  name = "tf_trans_profile1"
}

resource "citrixadc_transformaction" "tf_unset" {
  name             = "tf_test_transformaction_unset"
  profilename      = citrixadc_transformprofile.tf_trans_profile1.name
  priority         = 100
  requrlfrom       = "http://m3.mydomain.com/(.*)"
  requrlinto       = "https://exp-proxy-v1.api.mydomain.com/$1"
  resurlfrom       = "https://exp-proxy-v1.api.mydomain.com/(.*)"
  resurlinto       = "https://m3.mydomain.com/$1"
  cookiedomainfrom = "old.mydomain.com"
  cookiedomaininto = "new.mydomain.com"
  comment          = "unset test comment"
  state            = "DISABLED"
}
`

const testAccTransformaction_unset_step2 = `
resource "citrixadc_transformprofile" "tf_trans_profile1" {
  name = "tf_trans_profile1"
}

resource "citrixadc_transformaction" "tf_unset" {
  name        = "tf_test_transformaction_unset"
  profilename = citrixadc_transformprofile.tf_trans_profile1.name
  priority    = 100
  # All unset-eligible attributes removed from config -> the provider must
  # unset them (revert to NITRO defaults).
}
`

func TestAccTransformaction_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTransformactionDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccTransformaction_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTransformactionExist("citrixadc_transformaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_transformaction.tf_unset", "requrlfrom", "http://m3.mydomain.com/(.*)"),
					resource.TestCheckResourceAttr("citrixadc_transformaction.tf_unset", "requrlinto", "https://exp-proxy-v1.api.mydomain.com/$1"),
					resource.TestCheckResourceAttr("citrixadc_transformaction.tf_unset", "resurlfrom", "https://exp-proxy-v1.api.mydomain.com/(.*)"),
					resource.TestCheckResourceAttr("citrixadc_transformaction.tf_unset", "resurlinto", "https://m3.mydomain.com/$1"),
					resource.TestCheckResourceAttr("citrixadc_transformaction.tf_unset", "cookiedomainfrom", "old.mydomain.com"),
					resource.TestCheckResourceAttr("citrixadc_transformaction.tf_unset", "cookiedomaininto", "new.mydomain.com"),
					resource.TestCheckResourceAttr("citrixadc_transformaction.tf_unset", "comment", "unset test comment"),
					resource.TestCheckResourceAttr("citrixadc_transformaction.tf_unset", "state", "DISABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccTransformaction_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTransformactionExist("citrixadc_transformaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_transformaction.tf_unset", "state", "ENABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckTransformactionADCValue("tf_test_transformaction_unset", "state", "ENABLED"),
					testAccCheckTransformactionADCValue("tf_test_transformaction_unset", "requrlfrom", ""),
					testAccCheckTransformactionADCValue("tf_test_transformaction_unset", "requrlinto", ""),
					testAccCheckTransformactionADCValue("tf_test_transformaction_unset", "resurlfrom", ""),
					testAccCheckTransformactionADCValue("tf_test_transformaction_unset", "resurlinto", ""),
					testAccCheckTransformactionADCValue("tf_test_transformaction_unset", "cookiedomainfrom", ""),
					testAccCheckTransformactionADCValue("tf_test_transformaction_unset", "cookiedomaininto", ""),
					testAccCheckTransformactionADCValue("tf_test_transformaction_unset", "comment", ""),
				),
			},
		},
	})
}

const testAccTransformactionDataSource_basic = `
resource "citrixadc_transformprofile" "tf_trans_profile1" {
  name = "tf_trans_profile1"
}

resource "citrixadc_transformaction" "tf_trans_action" {
  name = "tf_trans_action"
  profilename = citrixadc_transformprofile.tf_trans_profile1.name
  priority = 100
  requrlfrom = "http://m3.mydomain.com/(.*)"
  requrlinto = "https://exp-proxy-v1.api.mydomain.com/$1"
  resurlfrom = "https://exp-proxy-v1.api.mydomain.com/(.*)"
  resurlinto = "https://m3.mydomain.com/$1"
}

data "citrixadc_transformaction" "tf_trans_action" {
    name = citrixadc_transformaction.tf_trans_action.name
}
`

func TestAccTransformactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTransformactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_transformaction.tf_trans_action", "name", "tf_trans_action"),
					resource.TestCheckResourceAttr("data.citrixadc_transformaction.tf_trans_action", "profilename", "tf_trans_profile1"),
					resource.TestCheckResourceAttr("data.citrixadc_transformaction.tf_trans_action", "priority", "100"),
					resource.TestCheckResourceAttr("data.citrixadc_transformaction.tf_trans_action", "requrlfrom", "http://m3.mydomain.com/(.*)"),
					resource.TestCheckResourceAttr("data.citrixadc_transformaction.tf_trans_action", "requrlinto", "https://exp-proxy-v1.api.mydomain.com/$1"),
					resource.TestCheckResourceAttr("data.citrixadc_transformaction.tf_trans_action", "resurlfrom", "https://exp-proxy-v1.api.mydomain.com/(.*)"),
					resource.TestCheckResourceAttr("data.citrixadc_transformaction.tf_trans_action", "resurlinto", "https://m3.mydomain.com/$1"),
				),
			},
		},
	})
}
