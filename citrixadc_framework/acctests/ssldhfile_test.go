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

// ssldhfile is an Import-as-create resource (POST ?action=Import), with no
// update (every attribute is RequiresReplace) and a keyless delete
// (DELETE ?args=name:<name>). The `src` attribute must point at a reachable
// DH parameters file (an http(s) URL or a local: path of a file already on the
// appliance). It is a write-only Import input that NITRO does not echo back,
// so it is never asserted.
//
// doSsldhfilePreChecks (see helpers_test.go) uploads testdata/dhparam.pem to
// /var/tmp as a prerequisite, so src = "local:dhparam.pem" resolves on the appliance.

const testAccSsldhfile_basic_step1 = `
resource "citrixadc_ssldhfile" "tf_ssldhfile" {
  name = "tf_ssldhfile"
  src  = "local:dhparam.pem"
}

`

func TestAccSsldhfile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSsldhfilePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSsldhfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSsldhfile_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSsldhfileExist("citrixadc_ssldhfile.tf_ssldhfile", nil),
					resource.TestCheckResourceAttr("citrixadc_ssldhfile.tf_ssldhfile", "name", "tf_ssldhfile"),
				),
			},
		},
	})
}

func testAccCheckSsldhfileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ssldhfile name is set")
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
		// ssldhfile has NO get-by-name endpoint (GET /ssldhfile/<name> => 400,
		// errorcode 1090). Get all records and filter by name.
		allResources, err := client.FindAllResources(service.Ssldhfile.Type())
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
			return fmt.Errorf("ssldhfile %s not found", n)
		}

		return nil
	}
}

func testAccCheckSsldhfileDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_ssldhfile" {
			continue
		}

		// ssldhfile has NO get-by-name endpoint. Get all records and filter by
		// name; a list error is treated as destroyed.
		allResources, err := client.FindAllResources(service.Ssldhfile.Type())
		if err != nil {
			continue
		}

		for _, v := range allResources {
			if name, ok := v["name"].(string); ok && name == rs.Primary.ID {
				return fmt.Errorf("ssldhfile %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

const testAccSsldhfileDataSource_basic = `

resource "citrixadc_ssldhfile" "tf_ssldhfile" {
  name = "tf_ssldhfile"
  src  = "local:dhparam.pem"
}

data "citrixadc_ssldhfile" "tf_ssldhfile" {
  name       = citrixadc_ssldhfile.tf_ssldhfile.name
  depends_on = [citrixadc_ssldhfile.tf_ssldhfile]
}
`

func TestAccSsldhfileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSsldhfilePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSsldhfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSsldhfileDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_ssldhfile.tf_ssldhfile", "name", "tf_ssldhfile"),
				),
			},
		},
	})
}
