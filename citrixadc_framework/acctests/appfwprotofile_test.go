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

// NOTE: appfwprotofile has every attribute marked RequiresReplace, and NITRO
// exposes no in-place update endpoint. Step 2 therefore mutates the `name`
// (and other inputs) to force a destroy + recreate, which is the only
// "update" path available for this resource.
const testAccAppfwprotofile_basic_step1 = `
resource "citrixadc_appfwprotofile" "tf_appfwprotofile" {
  name    = "tf_appfwprotofile"
  src     = "local:tftest.proto"
  comment = "test_comment"
}

`

const testAccAppfwprotofile_basic_step2 = `
resource "citrixadc_appfwprotofile" "tf_appfwprotofile" {
  name    = "tf_appfwprotofile_2"
  src     = "local:tftest.proto"
  comment = "test_comment_updated"
}

`

func TestAccAppfwprotofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprotofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprotofile_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprotofileExist("citrixadc_appfwprotofile.tf_appfwprotofile", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprotofile.tf_appfwprotofile", "name", "tf_appfwprotofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwprotofile.tf_appfwprotofile", "src", "local:tftest.proto"),
					resource.TestCheckResourceAttr("citrixadc_appfwprotofile.tf_appfwprotofile", "comment", "test_comment"),
				),
			},
			{
				Config: testAccAppfwprotofile_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprotofileExist("citrixadc_appfwprotofile.tf_appfwprotofile", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwprotofile.tf_appfwprotofile", "name", "tf_appfwprotofile_2"),
					resource.TestCheckResourceAttr("citrixadc_appfwprotofile.tf_appfwprotofile", "src", "local:tftest.proto"),
					resource.TestCheckResourceAttr("citrixadc_appfwprotofile.tf_appfwprotofile", "comment", "test_comment_updated"),
				),
			},
		},
	})
}

// doAppfwprotofilePreChecks uploads the gRPC schema file to /var/tmp so the
// resource's `src = "local:tftest.proto"` Import resolves on the appliance.
func doAppfwprotofilePreChecks(t *testing.T) {
	testAccPreCheck(t)

	c, err := testHelperInstantiateClient("", "", "", false)
	if err != nil {
		t.Fatalf("Failed to instantiate client. %v\n", err)
	}
	if err := uploadTestdataFile(c, t, "tftest.proto", "/var/tmp"); err != nil {
		t.Errorf("%v", err)
	}
}

func TestAccAppfwprotofile_import(t *testing.T) {
	const resAddr = "citrixadc_appfwprotofile.tf_appfwprotofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doAppfwprotofilePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprotofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAppfwprotofile_basic_step1},
			{
				Config:            testAccAppfwprotofile_basic_step1,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// `comment` is a write-only Import input that NITRO never echoes
				// back, and `src` is a write-only Import input that NITRO
				// normalizes (strips the "local:" prefix); neither can round-trip
				// through import.
				ImportStateVerifyIgnore: []string{"comment", "src"},
			},
		},
	})
}

func testAccCheckAppfwprotofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprotofile name is set")
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
		data, err := client.FindResource(service.Appfwprotofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appfwprotofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwprotofileDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprotofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Appfwprotofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwprotofile %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

// Datasource test - the data source has a precondition that the underlying
// resource must exist first. `depends_on` ensures the resource is created
// before the data source attempts to read it.
const testAccAppfwprotofileDataSource_basic = `

resource "citrixadc_appfwprotofile" "tf_appfwprotofile" {
  name    = "tf_appfwprotofile"
  src     = "local:tftest.proto"
  comment = "test_comment"
}

data "citrixadc_appfwprotofile" "tf_appfwprotofile" {
  name       = citrixadc_appfwprotofile.tf_appfwprotofile.name
  depends_on = [citrixadc_appfwprotofile.tf_appfwprotofile]
}
`

func TestAccAppfwprotofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprotofileDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appfwprotofile.tf_appfwprotofile", "name", "tf_appfwprotofile"),
				),
			},
		},
	})
}
