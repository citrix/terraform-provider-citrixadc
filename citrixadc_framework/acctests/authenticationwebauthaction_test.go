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

const testAccAuthenticationwebauthaction_add = `
	resource "citrixadc_authenticationwebauthaction" "tf_webauthaction" {
		name                       = "tf_webauthaction"
		serverip                   = "1.2.3.4"
		serverport                 = 8080
		fullreqexpr                = "TRUE"
		scheme                     = "https"
		successrule                = "http.RES.STATUS.EQ(200)"
		defaultauthenticationgroup = "old_group"
	}
`
const testAccAuthenticationwebauthaction_update = `
	resource "citrixadc_authenticationwebauthaction" "tf_webauthaction" {
		name                       = "tf_webauthaction"
		serverip                   = "1.2.3.4"
		serverport                 = 8080
		fullreqexpr                = "FALSE"
		scheme                     = "http"
		successrule                = "http.RES.STATUS.EQ(200)"
		defaultauthenticationgroup = "new_group"
	}
`

func TestAccAuthenticationwebauthaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationwebauthactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationwebauthaction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationwebauthactionExist("citrixadc_authenticationwebauthaction.tf_webauthaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationwebauthaction.tf_webauthaction", "name", "tf_webauthaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationwebauthaction.tf_webauthaction", "fullreqexpr", "TRUE"),
					resource.TestCheckResourceAttr("citrixadc_authenticationwebauthaction.tf_webauthaction", "scheme", "https"),
				),
			},
			{
				Config: testAccAuthenticationwebauthaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationwebauthactionExist("citrixadc_authenticationwebauthaction.tf_webauthaction", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationwebauthaction.tf_webauthaction", "name", "tf_webauthaction"),
					resource.TestCheckResourceAttr("citrixadc_authenticationwebauthaction.tf_webauthaction", "fullreqexpr", "FALSE"),
					resource.TestCheckResourceAttr("citrixadc_authenticationwebauthaction.tf_webauthaction", "scheme", "http"),
				),
			},
		},
	})
}

func TestAccAuthenticationwebauthaction_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_authenticationwebauthaction.tf_webauthaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationwebauthactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationwebauthaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationwebauthactionExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Authenticationwebauthaction.Type(), "tf_webauthaction"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAuthenticationwebauthaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationwebauthactionExist(resAddr, nil)),
			},
		},
	})
}

func TestAccAuthenticationwebauthaction_import(t *testing.T) {
	const resAddr = "citrixadc_authenticationwebauthaction.tf_webauthaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationwebauthactionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAuthenticationwebauthaction_add},
			{
				Config:                  testAccAuthenticationwebauthaction_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckAuthenticationwebauthactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationwebauthaction name is set")
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
		data, err := client.FindResource(service.Authenticationwebauthaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationwebauthaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationwebauthactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationwebauthaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Authenticationwebauthaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationwebauthaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccAuthenticationwebauthaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAuthenticationwebauthactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccAuthenticationwebauthaction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationwebauthactionExist("citrixadc_authenticationwebauthaction.tf_webauthaction", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAuthenticationwebauthaction_add,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckAuthenticationwebauthactionExist("citrixadc_authenticationwebauthaction.tf_webauthaction", nil)),
			},
		},
	})
}

// Unset test: step1 sets the unset-eligible attributes to non-default values;
// step2 removes them from config so the provider must unset them (revert to the
// NITRO default, which for these expression attributes is empty/absent).
// fullreqexpr is NOT unset-eligible: NITRO rejects unsetting it ("Please specify
// full request expression") and rolls back atomically, so it is kept in config.
const testAccAuthenticationwebauthaction_unset_step1 = `
	resource "citrixadc_authenticationwebauthaction" "tf_unset" {
		name                       = "tf_test_webauthaction_unset"
		serverip                   = "1.2.3.4"
		serverport                 = 8080
		scheme                     = "http"
		successrule                = "true"
		fullreqexpr                = "TRUE"
		defaultauthenticationgroup = "grp1"
		attribute1                 = "HTTP.RES.BODY(100)"
		attribute2                 = "HTTP.RES.BODY(100)"
		attribute3                 = "HTTP.RES.BODY(100)"
		attribute4                 = "HTTP.RES.BODY(100)"
		attribute5                 = "HTTP.RES.BODY(100)"
		attribute6                 = "HTTP.RES.BODY(100)"
		attribute7                 = "HTTP.RES.BODY(100)"
		attribute8                 = "HTTP.RES.BODY(100)"
		attribute9                 = "HTTP.RES.BODY(100)"
		attribute10                = "HTTP.RES.BODY(100)"
		attribute11                = "HTTP.RES.BODY(100)"
		attribute12                = "HTTP.RES.BODY(100)"
		attribute13                = "HTTP.RES.BODY(100)"
		attribute14                = "HTTP.RES.BODY(100)"
		attribute15                = "HTTP.RES.BODY(100)"
		attribute16                = "HTTP.RES.BODY(100)"
	}
`

const testAccAuthenticationwebauthaction_unset_step2 = `
	resource "citrixadc_authenticationwebauthaction" "tf_unset" {
		name        = "tf_test_webauthaction_unset"
		serverip    = "1.2.3.4"
		serverport  = 8080
		scheme      = "http"
		successrule = "true"
		fullreqexpr = "TRUE"
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO default, empty/absent).
	}
`

func TestAccAuthenticationwebauthaction_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationwebauthactionDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccAuthenticationwebauthaction_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationwebauthactionExist("citrixadc_authenticationwebauthaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationwebauthaction.tf_unset", "defaultauthenticationgroup", "grp1"),
					resource.TestCheckResourceAttr("citrixadc_authenticationwebauthaction.tf_unset", "attribute1", "HTTP.RES.BODY(100)"),
					resource.TestCheckResourceAttr("citrixadc_authenticationwebauthaction.tf_unset", "attribute16", "HTTP.RES.BODY(100)"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to the NITRO default (empty), and the implicit
				// post-apply plan must be empty.
				Config: testAccAuthenticationwebauthaction_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationwebauthactionExist("citrixadc_authenticationwebauthaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationwebauthaction.tf_unset", "defaultauthenticationgroup", ""),
					resource.TestCheckResourceAttr("citrixadc_authenticationwebauthaction.tf_unset", "attribute1", ""),
					resource.TestCheckResourceAttr("citrixadc_authenticationwebauthaction.tf_unset", "attribute16", ""),
					// Independent appliance-level confirmation the unset took effect
					// (the attributes are absent from the GET response).
					testAccCheckAuthenticationwebauthactionADCValue("tf_test_webauthaction_unset", "defaultauthenticationgroup", ""),
					testAccCheckAuthenticationwebauthactionADCValue("tf_test_webauthaction_unset", "attribute1", ""),
					testAccCheckAuthenticationwebauthactionADCValue("tf_test_webauthaction_unset", "attribute16", ""),
				),
			},
		},
	})
}

// testAccCheckAuthenticationwebauthactionADCValue asserts an attribute's value
// directly on the appliance. Unset-eligible attributes revert to absent, which
// is reported as the empty string.
func testAccCheckAuthenticationwebauthactionADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Authenticationwebauthaction.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("authenticationwebauthaction %s not found on appliance", name)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("authenticationwebauthaction %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

const testAccAuthenticationwebauthactionDataSource_basic = `
	resource "citrixadc_authenticationwebauthaction" "tf_webauthaction" {
		name                       = "tf_webauthaction_ds"
		serverip                   = "1.2.3.4"
		serverport                 = 8080
		fullreqexpr                = "TRUE"
		scheme                     = "https"
		successrule                = "http.RES.STATUS.EQ(200)"
		defaultauthenticationgroup = "test_group"
	}

	data "citrixadc_authenticationwebauthaction" "tf_webauthaction_ds" {
		name = citrixadc_authenticationwebauthaction.tf_webauthaction.name
	}
`

func TestAccAuthenticationwebauthactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationwebauthactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationwebauthaction.tf_webauthaction_ds", "name", "tf_webauthaction_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationwebauthaction.tf_webauthaction_ds", "serverip", "1.2.3.4"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationwebauthaction.tf_webauthaction_ds", "serverport", "8080"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationwebauthaction.tf_webauthaction_ds", "fullreqexpr", "TRUE"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationwebauthaction.tf_webauthaction_ds", "scheme", "https"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationwebauthaction.tf_webauthaction_ds", "successrule", "http.RES.STATUS.EQ(200)"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationwebauthaction.tf_webauthaction_ds", "defaultauthenticationgroup", "test_group"),
				),
			},
		},
	})
}
