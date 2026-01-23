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
	}
`

func TestAccNshttpprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
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
