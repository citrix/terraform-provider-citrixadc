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

// nsextension is created via the NITRO ?action=Import endpoint, which uploads
// the extension file referenced by `src`. For this test to actually run, `src`
// must point to a reachable import source containing a valid .lua extension —
// either a URL (e.g. "http://<host>/tftest_extension.lua") or a file already
// uploaded to the appliance under local:// (e.g. via the systemfile resource,
// mirroring how policyurlset / policypatsetfile stage their import source).
//
// TODO_PLACEHOLDER: replace `src` below with a real reachable .lua extension
// (an uploaded local file or an accessible URL). The value "local:tftest_extension.lua"
// is a placeholder and will fail Import until a real source is supplied.
//
// step1 imports the extension and verifies the name. `src` is NOT asserted: it
// is a write-only import input that the NITRO GET does not return.
// step2 updates a settable attribute (comment) to exercise the set/update path.
const testAccNsextension_basic_step1 = `
resource "citrixadc_nsextension" "tf_nsextension" {
  name    = "tf_nsextension"
  src     = "local:tftest_extension.lua"
  comment = "created by acceptance test"
}

`

const testAccNsextension_basic_step2 = `
resource "citrixadc_nsextension" "tf_nsextension" {
  name    = "tf_nsextension"
  src     = "local:tftest_extension.lua"
  comment = "updated by acceptance test"
}

`

func TestAccNsextension_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsextensionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsextension_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsextensionExist("citrixadc_nsextension.tf_nsextension", nil),
					resource.TestCheckResourceAttr("citrixadc_nsextension.tf_nsextension", "name", "tf_nsextension"),
					resource.TestCheckResourceAttr("citrixadc_nsextension.tf_nsextension", "comment", "created by acceptance test"),
					// src is a write-only import input and is not returned by GET - do not assert it.
				),
			},
			{
				Config: testAccNsextension_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsextensionExist("citrixadc_nsextension.tf_nsextension", nil),
					resource.TestCheckResourceAttr("citrixadc_nsextension.tf_nsextension", "name", "tf_nsextension"),
					resource.TestCheckResourceAttr("citrixadc_nsextension.tf_nsextension", "comment", "updated by acceptance test"),
				),
			},
		},
	})
}

func TestAccNsextension_import(t *testing.T) {
	t.Skip("TODO: Requires review")
	const resAddr = "citrixadc_nsextension.tf_nsextension"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsextensionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsextension_basic_step1,
			},
			{
				Config:                  testAccNsextension_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckNsextensionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsextension name is set")
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
		data, err := client.FindResource(service.Nsextension.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nsextension %s not found", n)
		}

		return nil
	}
}

func testAccCheckNsextensionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nsextension" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Nsextension.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nsextension %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

// TODO_PLACEHOLDER: `src` must point to a real reachable .lua extension for this
// datasource test to run (see the note above the basic test).
const testAccNsextensionDataSource_basic = `

resource "citrixadc_nsextension" "tf_nsextension" {
  name    = "tf_nsextension_ds"
  src     = "local:tftest_extension.lua"
  comment = "datasource test"
}

data "citrixadc_nsextension" "tf_nsextension" {
  name       = citrixadc_nsextension.tf_nsextension.name
  depends_on = [citrixadc_nsextension.tf_nsextension]
}
`

func TestAccNsextensionDataSource_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNsextensionDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nsextension.tf_nsextension", "name", "tf_nsextension_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_nsextension.tf_nsextension", "comment", "datasource test"),
				),
			},
		},
	})
}
