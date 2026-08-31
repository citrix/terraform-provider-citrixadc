/*
Copyright 2021 Citrix Systems, Inc

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

func TestAccBotprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBotprofileDestroy,
		Steps: []resource.TestStep{
			// create botprofile
			{
				Config: testAccBotprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotprofileExist("citrixadc_botprofile.tf_botprofile", nil),
					testAccCheckUserAgent(),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "addcookieflags", "secure"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "dfprequestlimit", "25"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "headlessbrowserdetection", "ON"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "sessioncookiename", "testCookie"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "sessiontimeout", "1200"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "spoofedreqaction.0", "LOG"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "verboseloglevel", "HTTP_FULL_HEADER"),
				),
			},
			// update botprofile actions
			{
				Config: testAccBotprofile_update_actions,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotprofileExist("citrixadc_botprofile.tf_botprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "devicefingerprintaction.0", "LOG"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "devicefingerprintaction.1", "DROP"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "trapaction.0", "LOG"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "trapaction.1", "DROP"),
					testAccCheckUserAgent(),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "addcookieflags", "secure"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "dfprequestlimit", "50"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "headlessbrowserdetection", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "sessioncookiename", "testCookie1"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "sessiontimeout", "1800"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "spoofedreqaction.0", "DROP"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "verboseloglevel", "NONE"),
				),
			},
			// update botprofile properties
			{
				Config: testAccBotprofile_update_properties,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotprofileExist("citrixadc_botprofile.tf_botprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "errorurl", "http://www.citrix.com/products/citrix-adc/"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "trapurl", "/http://www.citrix.com/products/citrix-adc/"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "comment", "tf_botprofile comment 1"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "bot_enable_white_list", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "bot_enable_black_list", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "bot_enable_rate_limit", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "devicefingerprint", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "bot_enable_ip_reputation", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "trap", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_botprofile", "bot_enable_tps", "OFF"),
					testAccCheckUserAgent(),
				),
			},
		},
	})
}

func testAccCheckBotprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Botprofile name is set")
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
		data, err := client.FindResource("botprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Botprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckBotprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_Botprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("botprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Botprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccBotprofile_basic = `
resource "citrixadc_botprofile" "tf_botprofile" {
	name = "tf_botprofile"
	errorurl = "http://www.citrix.com"
	trapurl = "/http://www.citrix.com"
	comment = "tf_botprofile comment"
	bot_enable_white_list = "ON"
	bot_enable_black_list = "ON"
	bot_enable_rate_limit = "ON"
	devicefingerprint = "ON"
	devicefingerprintaction = ["LOG", "RESET"]
	bot_enable_ip_reputation = "ON"
	trap = "ON"
	trapaction = ["LOG", "RESET"]
	bot_enable_tps = "ON"
	addcookieflags	= "secure"
	dfprequestlimit = "25"
	headlessbrowserdetection = "ON"
	sessioncookiename = "testCookie"
	sessiontimeout = "1200"
	spoofedreqaction = ["LOG"]
	verboseloglevel = "HTTP_FULL_HEADER"
}
`

const testAccBotprofile_update_actions = `
resource "citrixadc_botprofile" "tf_botprofile" {
	name = "tf_botprofile"
	devicefingerprintaction = ["LOG", "DROP"]
	trapaction = ["LOG", "DROP"]
	addcookieflags	= "secure"
	dfprequestlimit = "50"
	headlessbrowserdetection = "OFF"
	sessioncookiename = "testCookie1"
	sessiontimeout = "1800"
	spoofedreqaction = ["DROP"]
	verboseloglevel = "NONE"
}
`

const testAccBotprofile_update_properties = `
resource "citrixadc_botprofile" "tf_botprofile" {
	name = "tf_botprofile"
	errorurl = "http://www.citrix.com/products/citrix-adc/"
	trapurl = "/http://www.citrix.com/products/citrix-adc/"
	comment = "tf_botprofile comment 1"
	bot_enable_white_list = "OFF"
	bot_enable_black_list = "OFF"
	bot_enable_rate_limit = "OFF"
	devicefingerprint = "OFF"
	bot_enable_ip_reputation = "OFF"
	trap = "OFF"
	bot_enable_tps = "OFF"
}
`

// The botprofile unset test covers the type-independent, unset-eligible string
// flag attributes that carry a documented NITRO default. Setting them to
// non-default values in step1 and removing them from config in step2 must make
// the provider unset them (revert to the NITRO defaults).
const testAccBotprofile_unset_step1 = `
resource "citrixadc_botprofile" "tf_unset" {
	name                     = "tf_botprofile_unset"
	errorurl                 = "http://www.citrix.com"
	trapurl                  = "/http://www.citrix.com"
	devicefingerprintaction  = ["LOG", "RESET"]
	trapaction               = ["LOG", "RESET"]
	addcookieflags           = "secure"
	bot_enable_white_list    = "ON"
	bot_enable_black_list    = "ON"
	bot_enable_rate_limit    = "ON"
	bot_enable_ip_reputation = "ON"
	bot_enable_tps           = "ON"
	devicefingerprint        = "ON"
	trap                     = "ON"
	headlessbrowserdetection = "ON"
	verboseloglevel          = "HTTP_FULL_HEADER"
}
`

const testAccBotprofile_unset_step2 = `
resource "citrixadc_botprofile" "tf_unset" {
	name = "tf_botprofile_unset"
	# All unset-eligible attributes removed from config -> the provider must
	# unset them (revert to the documented NITRO defaults).
}
`

func TestAccBotprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBotprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccBotprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotprofileExist("citrixadc_botprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_unset", "addcookieflags", "secure"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_unset", "bot_enable_white_list", "ON"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_unset", "bot_enable_black_list", "ON"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_unset", "bot_enable_rate_limit", "ON"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_unset", "bot_enable_ip_reputation", "ON"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_unset", "bot_enable_tps", "ON"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_unset", "devicefingerprint", "ON"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_unset", "trap", "ON"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_unset", "headlessbrowserdetection", "ON"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_unset", "verboseloglevel", "HTTP_FULL_HEADER"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccBotprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotprofileExist("citrixadc_botprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_unset", "addcookieflags", "httpOnly"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_unset", "bot_enable_white_list", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_unset", "bot_enable_black_list", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_unset", "bot_enable_rate_limit", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_unset", "bot_enable_ip_reputation", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_unset", "bot_enable_tps", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_unset", "devicefingerprint", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_unset", "trap", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_unset", "headlessbrowserdetection", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_botprofile.tf_unset", "verboseloglevel", "NONE"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckBotprofileADCValue("tf_botprofile_unset", "bot_enable_white_list", "OFF"),
					testAccCheckBotprofileADCValue("tf_botprofile_unset", "addcookieflags", "httpOnly"),
					testAccCheckBotprofileADCValue("tf_botprofile_unset", "verboseloglevel", "NONE"),
				),
			},
		},
	})
}

// testAccCheckBotprofileADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it.
func testAccCheckBotprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Botprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("botprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("botprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccBotprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckBotprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccBotprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckBotprofileExist("citrixadc_botprofile.tf_botprofile", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccBotprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckBotprofileExist("citrixadc_botprofile.tf_botprofile", nil)),
			},
		},
	})
}

func TestAccBotprofile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_botprofile.tf_botprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBotprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBotprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckBotprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Botprofile.Type(), "tf_botprofile"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccBotprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckBotprofileExist(resAddr, nil)),
			},
		},
	})
}

func TestAccBotprofile_import(t *testing.T) {
	const resAddr = "citrixadc_botprofile.tf_botprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckBotprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccBotprofile_basic},
			{
				Config:                  testAccBotprofile_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccBotprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccBotprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_botprofile.tf_botprofile_ds", "name", "tf_botprofile_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_botprofile.tf_botprofile_ds", "addcookieflags", "secure"),
					resource.TestCheckResourceAttr("data.citrixadc_botprofile.tf_botprofile_ds", "dfprequestlimit", "25"),
					resource.TestCheckResourceAttr("data.citrixadc_botprofile.tf_botprofile_ds", "headlessbrowserdetection", "ON"),
					resource.TestCheckResourceAttr("data.citrixadc_botprofile.tf_botprofile_ds", "sessioncookiename", "dsCookie"),
					resource.TestCheckResourceAttr("data.citrixadc_botprofile.tf_botprofile_ds", "sessiontimeout", "1200"),
					resource.TestCheckResourceAttr("data.citrixadc_botprofile.tf_botprofile_ds", "verboseloglevel", "HTTP_FULL_HEADER"),
					resource.TestCheckResourceAttr("data.citrixadc_botprofile.tf_botprofile_ds", "comment", "DATASOURCE TEST COMMENT"),
				),
			},
		},
	})
}

const testAccBotprofileDataSource_basic = `

resource "citrixadc_botprofile" "tf_botprofile_ds" {
	name = "tf_botprofile_ds"
	errorurl = "http://www.citrix.com"
	trapurl = "/http://www.citrix.com"
	comment = "DATASOURCE TEST COMMENT"
	bot_enable_white_list = "ON"
	bot_enable_black_list = "ON"
	bot_enable_rate_limit = "ON"
	devicefingerprint = "ON"
	bot_enable_ip_reputation = "ON"
	trap = "ON"
	bot_enable_tps = "ON"
	addcookieflags = "secure"
	dfprequestlimit = "25"
	headlessbrowserdetection = "ON"
	sessioncookiename = "dsCookie"
	sessiontimeout = "1200"
	verboseloglevel = "HTTP_FULL_HEADER"
}

data "citrixadc_botprofile" "tf_botprofile_ds" {
	name = citrixadc_botprofile.tf_botprofile_ds.name
	depends_on = [citrixadc_botprofile.tf_botprofile_ds]
}

`
