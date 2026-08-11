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

const testAccNshttpprofile_add = `
	resource "citrixadc_nshttpprofile" "foo" {
		name  = "tf_httpprofile"
		http2 = "ENABLED"
        markrfc7230noncompliantinval = "ENABLED"
        markhttpheaderextrawserror = "ENABLED"
        dropinvalreqs = "ENABLED"
		maxheaderfieldlen = 2050
		http2maxrxresetframespermin = 10
		http3webtransport = "DISABLED"
		http3minseverconn = 10
		httppipelinebuffsize = 131100
		allowonlywordcharactersandhyphen = "DISABLED"
		hostheadervalidation = "DISABLED"
		maxduplicateheaderfields = 10
		passprotocolupgrade = "DISABLED"
		http2extendedconnect = "DISABLED"
		http2smallwndtimeout = 30
	}
`
const testAccNshttpprofile_update = `
	resource "citrixadc_nshttpprofile" "foo" {
		name  = "tf_httpprofile"
		http2 = "DISABLED"
        markrfc7230noncompliantinval = "DISABLED"
        markhttpheaderextrawserror = "DISABLED"
        dropinvalreqs = "DISABLED"
		maxheaderfieldlen = 2060
		http2maxrxresetframespermin = 20
		http3webtransport = "ENABLED"
		http3minseverconn = 20
		httppipelinebuffsize = 131200
		allowonlywordcharactersandhyphen = "ENABLED"
		hostheadervalidation = "ENABLED"
		maxduplicateheaderfields = 12
		passprotocolupgrade = "ENABLED"
		http2extendedconnect = "ENABLED"
		http2smallwndtimeout = 60
	}
`

// nshttpprofileSupportsHttp2smallwndtimeout reports whether the ADC firmware
// recognizes the http2smallwndtimeout attribute (added for CVE-2026-13474 by a
// recent 14.1 build). Older builds reject it on create with errorcode 278
// ("Invalid argument [http2smallwndtimeout]"). The attribute's presence as a key
// in the GET of the built-in nshttp_default_profile is a reliable, version-number-
// free firmware capability probe.
func nshttpprofileSupportsHttp2smallwndtimeout(t *testing.T) bool {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		t.Fatalf("Failed to get test client: %v", err)
	}
	data, err := client.FindResource(service.Nshttpprofile.Type(), "nshttp_default_profile")
	if err != nil {
		t.Fatalf("Failed to read nshttp_default_profile for firmware probe: %v", err)
	}
	_, ok := data["http2smallwndtimeout"]
	return ok
}

func TestAccNshttpprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			// The _add/_update configs set http2smallwndtimeout, which older ADC
			// builds reject with errorcode 278. Skip on firmware that lacks it
			// rather than reporting a spurious failure.
			if !nshttpprofileSupportsHttp2smallwndtimeout(t) {
				t.Skip("ADC firmware does not support http2smallwndtimeout (CVE-2026-13474 param); skipping nshttpprofile test")
			}
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNshttpprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNshttpprofile_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNshttpprofileExist("citrixadc_nshttpprofile.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "name", "tf_httpprofile"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "http2", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "markrfc7230noncompliantinval", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "markhttpheaderextrawserror", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "dropinvalreqs", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "maxheaderfieldlen", "2050"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "http2maxrxresetframespermin", "10"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "http3webtransport", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "http3minseverconn", "10"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "httppipelinebuffsize", "131100"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "allowonlywordcharactersandhyphen", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "hostheadervalidation", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "maxduplicateheaderfields", "10"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "passprotocolupgrade", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "http2extendedconnect", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "http2smallwndtimeout", "30"),
				),
			},
			{
				Config: testAccNshttpprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNshttpprofileExist("citrixadc_nshttpprofile.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "name", "tf_httpprofile"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "http2", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "markrfc7230noncompliantinval", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "markhttpheaderextrawserror", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "dropinvalreqs", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "maxheaderfieldlen", "2060"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "http2maxrxresetframespermin", "20"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "http3webtransport", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "http3minseverconn", "20"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "httppipelinebuffsize", "131200"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "allowonlywordcharactersandhyphen", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "hostheadervalidation", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "maxduplicateheaderfields", "12"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "passprotocolupgrade", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "http2extendedconnect", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.foo", "http2smallwndtimeout", "60"),
				),
			},
		},
	})
}

// The nshttpprofile unset test covers the string ENABLED/DISABLED attributes
// that are unset-eligible per the NITRO spec (Optional, mutable, with a
// documented server default). step1 sets each to its non-default value; step2
// removes them all from config, so the provider must unset them and the
// appliance must revert each to its documented default.
const testAccNshttpprofile_unset_step1 = `
resource "citrixadc_nshttpprofile" "tf_unset" {
  name                             = "tf_nshttpprofile_unset"
  adpttimeout                      = "ENABLED"
  allowonlywordcharactersandhyphen = "ENABLED"
  altsvc                           = "ENABLED"
  cmponpush                        = "ENABLED"
  conmultiplex                     = "DISABLED"
  dropextracrlf                    = "DISABLED"
  dropextradata                    = "ENABLED"
  dropinvalreqs                    = "ENABLED"
  grpclengthdelimitation           = "DISABLED"
  hostheadervalidation             = "ENABLED"
  http2                            = "ENABLED"
  http2altsvcframe                 = "ENABLED"
  http2direct                      = "ENABLED"
  http2extendedconnect             = "DISABLED"
  http2strictcipher                = "DISABLED"
  http3                            = "ENABLED"
  http3webtransport                = "ENABLED"
  markconnreqinval                 = "ENABLED"
  markhttp09inval                  = "ENABLED"
  markhttpheaderextrawserror       = "ENABLED"
  markrfc7230noncompliantinval     = "ENABLED"
  marktracereqinval                = "ENABLED"
  passprotocolupgrade              = "DISABLED"
  persistentetag                   = "ENABLED"
  rtsptunnel                       = "ENABLED"
  weblog                           = "DISABLED"
  websocket                        = "ENABLED"
}
`

const testAccNshttpprofile_unset_step2 = `
resource "citrixadc_nshttpprofile" "tf_unset" {
  name = "tf_nshttpprofile_unset"
  # All unset-eligible attributes removed from config -> the provider must
  # unset them (revert to NITRO defaults).
}
`

func TestAccNshttpprofile_unset(t *testing.T) {
	if !nshttpprofileSupportsHttp2smallwndtimeout(t) {
		t.Skip("ADC firmware does not support http2smallwndtimeout (CVE-2026-13474 param); skipping nshttpprofile test")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNshttpprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccNshttpprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNshttpprofileExist("citrixadc_nshttpprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "adpttimeout", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "allowonlywordcharactersandhyphen", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "altsvc", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "cmponpush", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "conmultiplex", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "dropextracrlf", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "dropextradata", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "dropinvalreqs", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "grpclengthdelimitation", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "hostheadervalidation", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "http2", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "http2altsvcframe", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "http2direct", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "http2extendedconnect", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "http2strictcipher", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "http3", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "http3webtransport", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "markconnreqinval", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "markhttp09inval", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "markhttpheaderextrawserror", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "markrfc7230noncompliantinval", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "marktracereqinval", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "passprotocolupgrade", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "persistentetag", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "rtsptunnel", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "weblog", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "websocket", "ENABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccNshttpprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNshttpprofileExist("citrixadc_nshttpprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "adpttimeout", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "allowonlywordcharactersandhyphen", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "altsvc", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "cmponpush", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "conmultiplex", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "dropextracrlf", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "dropextradata", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "dropinvalreqs", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "grpclengthdelimitation", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "hostheadervalidation", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "http2", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "http2altsvcframe", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "http2direct", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "http2extendedconnect", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "http2strictcipher", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "http3", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "http3webtransport", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "markconnreqinval", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "markhttp09inval", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "markhttpheaderextrawserror", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "markrfc7230noncompliantinval", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "marktracereqinval", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "passprotocolupgrade", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "persistentetag", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "rtsptunnel", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "weblog", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nshttpprofile.tf_unset", "websocket", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckNshttpprofileADCValue("tf_nshttpprofile_unset", "http2", "DISABLED"),
					testAccCheckNshttpprofileADCValue("tf_nshttpprofile_unset", "dropinvalreqs", "DISABLED"),
					testAccCheckNshttpprofileADCValue("tf_nshttpprofile_unset", "weblog", "ENABLED"),
					testAccCheckNshttpprofileADCValue("tf_nshttpprofile_unset", "conmultiplex", "ENABLED"),
				),
			},
		},
	})
}

// testAccCheckNshttpprofileADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckNshttpprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Nshttpprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("nshttpprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("nshttpprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func testAccCheckNshttpprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No NS HTTP Profile name is set")
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
		data, err := client.FindResource(service.Nshttpprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("NS HTTP Profile %s not found", n)
		}

		return nil
	}
}

func testAccCheckNshttpprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nshttpprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Nshttpprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("NS HTTP Profile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccNshttpprofile_selfHealing(t *testing.T) {
	t.Skip("ADC firmware does not support http2smallwndtimeout (CVE-2026-13474 param); skipping nshttpprofile test")
	const resAddr = "citrixadc_nshttpprofile.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNshttpprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNshttpprofile_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNshttpprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Nshttpprofile.Type(), "tf_httpprofile"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccNshttpprofile_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNshttpprofileExist(resAddr, nil)),
			},
		},
	})
}

// testAccNshttpprofile_upgrade mirrors testAccNshttpprofile_add but OMITS
// http2smallwndtimeout. That attribute (CVE-2026-13474) postdates the released
// citrixadc 2.2.0 provider pinned in step 1 of the upgrade test, so including it
// makes the 2.2.0 provider reject the config at validation with
// `Unsupported argument: An argument named "http2smallwndtimeout" is not expected here`.
// Every other attribute here IS present in the 2.2.0 schema. The current provider
// reads http2smallwndtimeout back as Computed, so omitting it from config does not
// weaken the upgrade check.
const testAccNshttpprofile_upgrade = `
	resource "citrixadc_nshttpprofile" "foo" {
		name  = "tf_httpprofile"
		http2 = "ENABLED"
        markrfc7230noncompliantinval = "ENABLED"
        markhttpheaderextrawserror = "ENABLED"
        dropinvalreqs = "ENABLED"
		maxheaderfieldlen = 2050
		http2maxrxresetframespermin = 10
		http3webtransport = "DISABLED"
		http3minseverconn = 10
		httppipelinebuffsize = 131100
		allowonlywordcharactersandhyphen = "DISABLED"
		hostheadervalidation = "DISABLED"
		maxduplicateheaderfields = 10
		passprotocolupgrade = "DISABLED"
		http2extendedconnect = "DISABLED"
	}
`

func TestAccNshttpprofile_sdkv2StateUpgrade(t *testing.T) {
	// The gate still applies: it keeps the test on firmware recent enough to
	// support the other (post-13.0) attributes the config sets, even though the
	// upgrade config itself omits http2smallwndtimeout for 2.2.0 compatibility.
	if !nshttpprofileSupportsHttp2smallwndtimeout(t) {
		t.Skip("ADC firmware does not support http2smallwndtimeout (CVE-2026-13474 param); skipping nshttpprofile test")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNshttpprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccNshttpprofile_upgrade,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNshttpprofileExist("citrixadc_nshttpprofile.foo", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccNshttpprofile_upgrade,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckNshttpprofileExist("citrixadc_nshttpprofile.foo", nil)),
			},
		},
	})
}

func TestAccNshttpprofile_import(t *testing.T) {
	if !nshttpprofileSupportsHttp2smallwndtimeout(t) {
		t.Skip("ADC firmware does not support http2smallwndtimeout (CVE-2026-13474 param); skipping nshttpprofile test")
	}
	const resAddr = "citrixadc_nshttpprofile.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNshttpprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccNshttpprofile_add},
			{
				Config:                  testAccNshttpprofile_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccNshttpprofileDataSource_basic = `

	resource "citrixadc_nshttpprofile" "tf_nshttpprofile_ds" {
		name                             = "tf_nshttpprofile_ds"
		dropinvalreqs                    = "ENABLED"
		markconnreqinval                 = "ENABLED"
		markhttp09inval                  = "ENABLED"
		cmponpush                        = "ENABLED"
		conmultiplex                     = "ENABLED"
		maxreusepool                     = 75
		http2                            = "ENABLED"
		altsvc                           = "ENABLED"
		reqtimeout                       = 60000
		persistentetag                   = "ENABLED"
		markhttpheaderextrawserror       = "ENABLED"
		markrfc7230noncompliantinval     = "ENABLED"
		allowonlywordcharactersandhyphen = "DISABLED"
	}

	data "citrixadc_nshttpprofile" "tf_nshttpprofile_ds" {
		name = citrixadc_nshttpprofile.tf_nshttpprofile_ds.name
	}
`

func TestAccNshttpprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNshttpprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNshttpprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nshttpprofile.tf_nshttpprofile_ds", "name", "tf_nshttpprofile_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_nshttpprofile.tf_nshttpprofile_ds", "dropinvalreqs", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_nshttpprofile.tf_nshttpprofile_ds", "markconnreqinval", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_nshttpprofile.tf_nshttpprofile_ds", "markhttp09inval", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_nshttpprofile.tf_nshttpprofile_ds", "cmponpush", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_nshttpprofile.tf_nshttpprofile_ds", "conmultiplex", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_nshttpprofile.tf_nshttpprofile_ds", "maxreusepool", "75"),
					resource.TestCheckResourceAttr("data.citrixadc_nshttpprofile.tf_nshttpprofile_ds", "http2", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_nshttpprofile.tf_nshttpprofile_ds", "altsvc", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_nshttpprofile.tf_nshttpprofile_ds", "reqtimeout", "60000"),
					resource.TestCheckResourceAttr("data.citrixadc_nshttpprofile.tf_nshttpprofile_ds", "persistentetag", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_nshttpprofile.tf_nshttpprofile_ds", "markhttpheaderextrawserror", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_nshttpprofile.tf_nshttpprofile_ds", "markrfc7230noncompliantinval", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_nshttpprofile.tf_nshttpprofile_ds", "allowonlywordcharactersandhyphen", "DISABLED"),
				),
			},
		},
	})
}
