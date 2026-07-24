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

const testAccAppfwsettings_basic_step1 = `
resource "citrixadc_appfwsettings" "tf_appfwsettings" {
  defaultprofile           = "APPFW_DROP"
  undefaction              = "APPFW_DROP"
  sessiontimeout           = 800
  learnratelimit           = 300
  sessionlifetime          = 1000
  sessioncookiename        = "test_ns_id"
  importsizelimit          = 134217700
  signatureautoupdate      = "ON"
  signatureurl             = "https://example.com"
  cookiepostencryptprefix  = "ENCRYPTED"
  geolocationlogging       = "ON"
  ceflogging               = "ON"
  entitydecoding           = "ON"
  useconfigurablesecretkey = "ON"
  sessionlimit             = 0
  malformedreqaction = [
	"none",
  ]
  centralizedlearning = "ON"
  proxyport           = 9090
}

`

const testAccAppfwsettings_basic_step2 = `
resource "citrixadc_appfwsettings" "tf_appfwsettings" {
  defaultprofile           = "APPFW_BYPASS"
  undefaction              = "APPFW_BLOCK"
  sessiontimeout           = 900
  learnratelimit           = 400
  sessionlifetime          = 0
  sessioncookiename        = "citrix_ns_id"
  importsizelimit          = 134217728
  signatureautoupdate      = "OFF"
  signatureurl             = "https://s3.amazonaws.com/NSAppFwSignatures/SignaturesMapping.xml"
  cookiepostencryptprefix  = "ENC"
  geolocationlogging       = "OFF"
  ceflogging               = "OFF"
  entitydecoding           = "OFF"
  useconfigurablesecretkey = "OFF"
  sessionlimit             = 100000
  malformedreqaction = [
    "block",
    "log",
    "stats"
  ]
  centralizedlearning = "OFF"
  proxyport           = 8080
}

`

func TestAccAppfwsettings_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwsettings_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwsettingsExist("citrixadc_appfwsettings.tf_appfwsettings", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "defaultprofile", "APPFW_DROP"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "undefaction", "APPFW_DROP"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "sessiontimeout", "800"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "learnratelimit", "300"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "sessionlifetime", "1000"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "sessioncookiename", "test_ns_id"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "importsizelimit", "134217700"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "signatureautoupdate", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "signatureurl", "https://example.com"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "cookiepostencryptprefix", "ENCRYPTED"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "geolocationlogging", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "ceflogging", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "entitydecoding", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "useconfigurablesecretkey", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "sessionlimit", "0"),
					// Attribute not present for check
					//resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "malformedreqaction", "[\"none\"]"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "centralizedlearning", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "proxyport", "9090"),
				),
			},
			{
				Config: testAccAppfwsettings_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwsettingsExist("citrixadc_appfwsettings.tf_appfwsettings", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "defaultprofile", "APPFW_BYPASS"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "undefaction", "APPFW_BLOCK"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "sessiontimeout", "900"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "learnratelimit", "400"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "sessionlifetime", "0"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "sessioncookiename", "citrix_ns_id"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "importsizelimit", "134217728"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "signatureautoupdate", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "signatureurl", "https://s3.amazonaws.com/NSAppFwSignatures/SignaturesMapping.xml"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "cookiepostencryptprefix", "ENC"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "geolocationlogging", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "ceflogging", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "entitydecoding", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "useconfigurablesecretkey", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "sessionlimit", "100000"),
					// Attribute not present for check
					//resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "malformedreqaction", "[\"block\", \"log\", \"stats\"]"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "centralizedlearning", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "proxyport", "8080"),
				),
			},
		},
	})
}

func TestAccAppfwsettings_import(t *testing.T) {
	const resAddr = "citrixadc_appfwsettings.tf_appfwsettings"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{Config: testAccAppfwsettings_basic_step1},
			{
				Config:            testAccAppfwsettings_basic_step1,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// proxypassword_wo_version is a write-only version tracker that is
				// not returned by NITRO, so it cannot round-trip through import.
				ImportStateVerifyIgnore: []string{"proxypassword_wo_version"},
			},
		},
	})
}

func testAccCheckAppfwsettingsExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwsettings name is set")
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
		data, err := client.FindResource(service.Appfwsettings.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appfwsettings %s not found", n)
		}

		return nil
	}
}

// Test backward-compatible path: using proxypassword (Sensitive attribute)
const testAccAppfwsettings_proxypassword_step1 = `

	variable "appfwsettings_proxypassword" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_appfwsettings" "tf_appfwsettings" {
		proxyusername = "testuser"
		proxypassword = var.appfwsettings_proxypassword
		proxyport     = 9090
	}
`

// Update backward-compatible path: change proxypassword value
const testAccAppfwsettings_proxypassword_step2 = `

	 variable "appfwsettings_proxypassword_2" {
	  type      = string
	  sensitive = true
	}

	 resource "citrixadc_appfwsettings" "tf_appfwsettings" {
		proxyusername = "testuser"
		proxypassword = var.appfwsettings_proxypassword_2
		proxyport     = 9090
	}
`

func TestAccAppfwsettings_proxypassword_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_appfwsettings_proxypassword", "oldpassword123")
	t.Setenv("TF_VAR_appfwsettings_proxypassword_2", "newpassword456")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwsettings_proxypassword_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwsettingsExist("citrixadc_appfwsettings.tf_appfwsettings", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "proxyport", "9090"),
				),
			},
			{
				Config: testAccAppfwsettings_proxypassword_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwsettingsExist("citrixadc_appfwsettings.tf_appfwsettings", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "proxyport", "9090"),
				),
			},
		},
	})
}

// Test ephemeral path: using proxypassword_wo (WriteOnly attribute) with version tracker
const testAccAppfwsettings_proxypassword_wo_step1 = `

	variable "appfwsettings_proxypassword_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_appfwsettings" "tf_appfwsettings" {
		proxyusername            = "testuser"
		proxypassword_wo         = var.appfwsettings_proxypassword_wo
		proxypassword_wo_version = 1
		proxyport                = 9090
	}
`

// Update ephemeral path: bump version to trigger update with new password
const testAccAppfwsettings_proxypassword_wo_step2 = `

	 variable "appfwsettings_proxypassword_wo_2" {
	  type      = string
	  sensitive = true
	}

	 resource "citrixadc_appfwsettings" "tf_appfwsettings" {
		proxyusername            = "testuser"
		proxypassword_wo         = var.appfwsettings_proxypassword_wo_2
		proxypassword_wo_version = 2
		proxyport                = 9090
	}
`

func TestAccAppfwsettings_proxypassword_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_appfwsettings_proxypassword_wo", "ephemeral_pass1")
	t.Setenv("TF_VAR_appfwsettings_proxypassword_wo_2", "ephemeral_pass2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwsettings_proxypassword_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwsettingsExist("citrixadc_appfwsettings.tf_appfwsettings", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "proxypassword_wo_version", "1"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "proxyport", "9090"),
				),
			},
			{
				Config: testAccAppfwsettings_proxypassword_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwsettingsExist("citrixadc_appfwsettings.tf_appfwsettings", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "proxypassword_wo_version", "2"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_appfwsettings", "proxyport", "9090"),
				),
			},
		},
	})
}

const testAccAppfwsettingsDataSource_basic = `

resource "citrixadc_appfwsettings" "tf_appfwsettings" {
  defaultprofile           = "APPFW_DROP"
  undefaction              = "APPFW_DROP"
  sessiontimeout           = 800
  learnratelimit           = 300
  sessionlifetime          = 1000
  sessioncookiename        = "test_ns_id"
  importsizelimit          = 134217700
  signatureautoupdate      = "ON"
  signatureurl             = "https://example.com"
  cookiepostencryptprefix  = "ENCRYPTED"
  geolocationlogging       = "ON"
  ceflogging               = "ON"
  entitydecoding           = "ON"
  useconfigurablesecretkey = "ON"
  sessionlimit             = 0
  malformedreqaction = [
    "none",
  ]
  centralizedlearning = "ON"
  proxyport           = 9090
}

data "citrixadc_appfwsettings" "tf_appfwsettings" {
  depends_on = [citrixadc_appfwsettings.tf_appfwsettings]
}
`

func TestAccAppfwsettingsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwsettingsDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appfwsettings.tf_appfwsettings", "defaultprofile", "APPFW_DROP"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwsettings.tf_appfwsettings", "undefaction", "APPFW_DROP"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwsettings.tf_appfwsettings", "sessiontimeout", "800"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwsettings.tf_appfwsettings", "learnratelimit", "300"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwsettings.tf_appfwsettings", "sessionlifetime", "1000"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwsettings.tf_appfwsettings", "sessioncookiename", "test_ns_id"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwsettings.tf_appfwsettings", "importsizelimit", "134217700"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwsettings.tf_appfwsettings", "signatureautoupdate", "ON"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwsettings.tf_appfwsettings", "signatureurl", "https://example.com"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwsettings.tf_appfwsettings", "cookiepostencryptprefix", "ENCRYPTED"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwsettings.tf_appfwsettings", "geolocationlogging", "ON"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwsettings.tf_appfwsettings", "ceflogging", "ON"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwsettings.tf_appfwsettings", "entitydecoding", "ON"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwsettings.tf_appfwsettings", "useconfigurablesecretkey", "ON"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwsettings.tf_appfwsettings", "sessionlimit", "0"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwsettings.tf_appfwsettings", "centralizedlearning", "ON"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwsettings.tf_appfwsettings", "proxyport", "9090"),
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// Unset test: proves ?action=unset support for the unset-eligible attributes.
// Step 1 sets every eligible attribute to a valid non-default value; step 2
// removes them from config so the provider must issue unset -> each reverts to
// its NITRO default (asserted in state AND directly on the appliance), with an
// empty post-apply plan.
// ---------------------------------------------------------------------------
const testAccAppfwsettings_unset_step1 = `
resource "citrixadc_appfwsettings" "tf_unset" {
  ceflogging               = "ON"
  centralizedlearning      = "ON"
  cookieflags              = "all"
  defaultprofile           = "APPFW_DROP"
  entitydecoding           = "ON"
  geolocationlogging       = "ON"
  importsizelimit          = 134217700
  learnratelimit           = 300
  proxyport                = 9090
  sessionlifetime          = 1000
  sessionlimit             = 50000
  sessiontimeout           = 800
  signatureautoupdate      = "ON"
  signatureurl             = "https://example.com"
  undefaction              = "APPFW_DROP"
  useconfigurablesecretkey = "ON"
}
`

const testAccAppfwsettings_unset_step2 = `
resource "citrixadc_appfwsettings" "tf_unset" {
  # All unset-eligible attributes removed from config -> provider must unset them.
}
`

func TestAccAppfwsettings_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				// Non-default values apply and persist.
				Config: testAccAppfwsettings_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwsettingsExist("citrixadc_appfwsettings.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "ceflogging", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "centralizedlearning", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "cookieflags", "all"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "defaultprofile", "APPFW_DROP"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "entitydecoding", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "geolocationlogging", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "importsizelimit", "134217700"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "learnratelimit", "300"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "proxyport", "9090"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "sessionlifetime", "1000"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "sessionlimit", "50000"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "sessiontimeout", "800"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "signatureautoupdate", "ON"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "signatureurl", "https://example.com"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "undefaction", "APPFW_DROP"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "useconfigurablesecretkey", "ON"),
				),
			},
			{
				// Removing them must unset -> state reverts to NITRO defaults,
				// and the implicit post-apply plan must be empty.
				Config: testAccAppfwsettings_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwsettingsExist("citrixadc_appfwsettings.tf_unset", nil),
					// State reverted to defaults.
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "ceflogging", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "centralizedlearning", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "cookieflags", "none"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "defaultprofile", "APPFW_BYPASS"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "entitydecoding", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "geolocationlogging", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "importsizelimit", "134217728"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "learnratelimit", "400"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "proxyport", "8080"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "sessionlifetime", "0"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "sessionlimit", "100000"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "sessiontimeout", "900"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "signatureautoupdate", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "signatureurl", "https://s3.amazonaws.com/NSAppFwSignatures/SignaturesMapping.xml"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "undefaction", "APPFW_BLOCK"),
					resource.TestCheckResourceAttr("citrixadc_appfwsettings.tf_unset", "useconfigurablesecretkey", "OFF"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAppfwsettingsADCValue("ceflogging", "OFF"),
					testAccCheckAppfwsettingsADCValue("centralizedlearning", "OFF"),
					testAccCheckAppfwsettingsADCValue("cookieflags", "none"),
					testAccCheckAppfwsettingsADCValue("defaultprofile", "APPFW_BYPASS"),
					testAccCheckAppfwsettingsADCValue("entitydecoding", "OFF"),
					testAccCheckAppfwsettingsADCValue("geolocationlogging", "OFF"),
					testAccCheckAppfwsettingsADCValue("importsizelimit", "134217728"),
					testAccCheckAppfwsettingsADCValue("learnratelimit", "400"),
					testAccCheckAppfwsettingsADCValue("proxyport", "8080"),
					testAccCheckAppfwsettingsADCValue("sessionlifetime", "0"),
					testAccCheckAppfwsettingsADCValue("sessionlimit", "100000"),
					testAccCheckAppfwsettingsADCValue("sessiontimeout", "900"),
					testAccCheckAppfwsettingsADCValue("signatureautoupdate", "OFF"),
					testAccCheckAppfwsettingsADCValue("signatureurl", "https://s3.amazonaws.com/NSAppFwSignatures/SignaturesMapping.xml"),
					testAccCheckAppfwsettingsADCValue("undefaction", "APPFW_BLOCK"),
					testAccCheckAppfwsettingsADCValue("useconfigurablesecretkey", "OFF"),
				),
			},
		},
	})
}

// testAccCheckAppfwsettingsADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
// appfwsettings is a singleton, so it is fetched with an empty name.
func testAccCheckAppfwsettingsADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Appfwsettings.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("appfwsettings not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("appfwsettings: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccAppfwsettings_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccAppfwsettings_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwsettingsExist("citrixadc_appfwsettings.tf_appfwsettings", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAppfwsettings_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwsettingsExist("citrixadc_appfwsettings.tf_appfwsettings", nil),
				),
			},
		},
	})
}
