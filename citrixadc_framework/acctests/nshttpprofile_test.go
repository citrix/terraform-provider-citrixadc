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
