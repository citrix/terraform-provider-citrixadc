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

// sslcrlfile is an Import-as-create resource (POST ?action=Import), with no
// update (every attribute is RequiresReplace) and a keyless delete
// (DELETE ?args=name:<name>). The `src` attribute must point at a reachable
// CRL file (an http(s) URL or a local: path of a file already on the
// appliance). It is a write-only Import input that NITRO does not echo back,
// so it is never asserted.
//
// doSslcrlfilePreChecks (see helpers_test.go) uploads testdata/sample.crl to
// /var/tmp as a prerequisite, so src = "local:sample.crl" resolves on the appliance.

const testAccSslcrlfile_basic_step1 = `
resource "citrixadc_sslcrlfile" "tf_sslcrlfile" {
  name = "tf_sslcrlfile"
  src  = "local:sample.crl"
}

`

func TestAccSslcrlfile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcrlfilePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslcrlfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcrlfile_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcrlfileExist("citrixadc_sslcrlfile.tf_sslcrlfile", nil),
					resource.TestCheckResourceAttr("citrixadc_sslcrlfile.tf_sslcrlfile", "name", "tf_sslcrlfile"),
				),
			},
		},
	})
}

func TestAccSslcrlfile_import(t *testing.T) {
	const resAddr = "citrixadc_sslcrlfile.tf_sslcrlfile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcrlfilePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslcrlfileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccSslcrlfile_basic_step1},
			{
				Config:            testAccSslcrlfile_basic_step1,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// `src` is a write-only Import input that NITRO does not echo back,
				// so it cannot round-trip through import.
				ImportStateVerifyIgnore: []string{"src"},
			},
		},
	})
}

func testAccCheckSslcrlfileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslcrlfile name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		// sslcrlfile has NO get-by-name endpoint (GET /sslcrlfile/<name> => 400,
		// errorcode 1090). Get all records and filter by name.
		allResources, err := client.FindAllResources(service.Sslcrlfile.Type())
		if err != nil {
			return err
		}

		found := false
		for _, v := range allResources {
			if name, ok := v["name"].(string); ok && name == rs.Primary.ID {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("sslcrlfile %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslcrlfileDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslcrlfile" {
			continue
		}

		// sslcrlfile has NO get-by-name endpoint. Get all records and filter by
		// name; a list error is treated as destroyed.
		allResources, err := client.FindAllResources(service.Sslcrlfile.Type())
		if err != nil {
			continue
		}

		for _, v := range allResources {
			if name, ok := v["name"].(string); ok && name == rs.Primary.ID {
				return fmt.Errorf("sslcrlfile %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

const testAccSslcrlfileDataSource_basic = `

resource "citrixadc_sslcrlfile" "tf_sslcrlfile" {
  name = "tf_sslcrlfile"
  src  = "local:sample.crl"
}

data "citrixadc_sslcrlfile" "tf_sslcrlfile" {
  name       = citrixadc_sslcrlfile.tf_sslcrlfile.name
  depends_on = [citrixadc_sslcrlfile.tf_sslcrlfile]
}
`

func TestAccSslcrlfileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcrlfilePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslcrlfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcrlfileDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslcrlfile.tf_sslcrlfile", "name", "tf_sslcrlfile"),
				),
			},
		},
	})
}
