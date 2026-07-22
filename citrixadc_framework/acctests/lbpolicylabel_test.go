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

// lbpolicylabel is a named resource with NO in-place update: NITRO exposes only
// add/delete/rename and all schema attributes are RequiresReplace. Therefore the
// basic test uses a SINGLE step (create + verify) rather than a create/update pair.
const testAccLbpolicylabel_basic_step1 = `
resource "citrixadc_lbpolicylabel" "tf_lbpolicylabel" {
  labelname       = "tf_lbpolicylabel"
  policylabeltype = "HTTP"
  comment         = "test label"
}

`

func TestAccLbpolicylabel_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbpolicylabelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbpolicylabel_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbpolicylabelExist("citrixadc_lbpolicylabel.tf_lbpolicylabel", nil),
					resource.TestCheckResourceAttr("citrixadc_lbpolicylabel.tf_lbpolicylabel", "labelname", "tf_lbpolicylabel"),
					resource.TestCheckResourceAttr("citrixadc_lbpolicylabel.tf_lbpolicylabel", "policylabeltype", "HTTP"),
					resource.TestCheckResourceAttr("citrixadc_lbpolicylabel.tf_lbpolicylabel", "comment", "test label"),
				),
			},
		},
	})
}

func TestAccLbpolicylabel_import(t *testing.T) {
	const resAddr = "citrixadc_lbpolicylabel.tf_lbpolicylabel"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbpolicylabelDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLbpolicylabel_basic_step1},
			{Config: testAccLbpolicylabel_basic_step1, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func testAccCheckLbpolicylabelExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbpolicylabel name is set")
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
		data, err := client.FindResource(service.Lbpolicylabel.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lbpolicylabel %s not found", n)
		}

		return nil
	}
}

func testAccCheckLbpolicylabelDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbpolicylabel" {
			continue
		}

		_, err := client.FindResource(service.Lbpolicylabel.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lbpolicylabel %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

// lbpolicylabel supports an in-place RENAME via the `newname` attribute (NITRO
// ?action=rename). Changing `newname` triggers an Update (NOT a replacement):
// the live object on the ADC is physically renamed and the resource ID is set to
// the new name. The other attributes (labelname/policylabeltype/comment) remain
// RequiresReplace.
//
// Step 1 creates the label under "tf_lbpolicylabel" (no newname).
// Step 2 keeps the SAME labelname/policylabeltype but adds newname, which renames
// the live object to "tf_lbpolicylabel_renamed".
const testAccLbpolicylabel_rename_step1 = `
resource "citrixadc_lbpolicylabel" "tf_lbpolicylabel" {
  labelname       = "tf_lbpolicylabel"
  policylabeltype = "HTTP"
}

`

const testAccLbpolicylabel_rename_step2 = `
resource "citrixadc_lbpolicylabel" "tf_lbpolicylabel" {
  labelname       = "tf_lbpolicylabel"
  policylabeltype = "HTTP"
  newname         = "tf_lbpolicylabel_renamed"
}

`

func TestAccLbpolicylabel_rename(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbpolicylabelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbpolicylabel_rename_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbpolicylabelExist("citrixadc_lbpolicylabel.tf_lbpolicylabel", nil),
					resource.TestCheckResourceAttr("citrixadc_lbpolicylabel.tf_lbpolicylabel", "labelname", "tf_lbpolicylabel"),
					resource.TestCheckResourceAttr("citrixadc_lbpolicylabel.tf_lbpolicylabel", "policylabeltype", "HTTP"),
				),
			},
			{
				Config: testAccLbpolicylabel_rename_step2,
				Check: resource.ComposeTestCheckFunc(
					// labelname stays the configured value; the rename only changes
					// the LIVE object name (and the resource ID), not the labelname attr.
					resource.TestCheckResourceAttr("citrixadc_lbpolicylabel.tf_lbpolicylabel", "labelname", "tf_lbpolicylabel"),
					resource.TestCheckResourceAttr("citrixadc_lbpolicylabel.tf_lbpolicylabel", "policylabeltype", "HTTP"),
					resource.TestCheckResourceAttr("citrixadc_lbpolicylabel.tf_lbpolicylabel", "newname", "tf_lbpolicylabel_renamed"),
					// After rename, the resource ID points at the new name.
					resource.TestCheckResourceAttr("citrixadc_lbpolicylabel.tf_lbpolicylabel", "id", "tf_lbpolicylabel_renamed"),
					// Distinguishes RENAME from RECREATE: the live object must now exist
					// under the NEW name and NOT under the old one.
					testAccCheckLbpolicylabelRenamed("citrixadc_lbpolicylabel.tf_lbpolicylabel", "tf_lbpolicylabel", "tf_lbpolicylabel_renamed"),
				),
			},
		},
	})
}

// testAccCheckLbpolicylabelRenamed verifies the in-place rename actually happened
// on the ADC rather than a silent recreate. A genuine rename leaves the live
// object addressable under newName and removes the oldName; a recreate (which
// would re-run the add path, where `newname` is excluded from the payload) would
// instead leave the object named oldName. It also asserts the resource ID tracks
// the new name.
func testAccCheckLbpolicylabelRenamed(n, oldName, newName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID != newName {
			return fmt.Errorf("Expected resource ID to be the renamed name %q, got %q", newName, rs.Primary.ID)
		}

		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		// The renamed object MUST exist under the new name.
		data, err := client.FindResource(service.Lbpolicylabel.Type(), newName)
		if err != nil {
			return fmt.Errorf("renamed lbpolicylabel %q not found on ADC: %v", newName, err)
		}
		if data == nil {
			return fmt.Errorf("renamed lbpolicylabel %q not found on ADC", newName)
		}

		// The old name MUST be gone (a recreate would have left it under oldName).
		oldData, err := client.FindResource(service.Lbpolicylabel.Type(), oldName)
		if err == nil && oldData != nil {
			return fmt.Errorf("old lbpolicylabel %q still exists on ADC - rename did not occur (likely a recreate)", oldName)
		}

		return nil
	}
}

const testAccLbpolicylabelDataSource_basic = `

resource "citrixadc_lbpolicylabel" "tf_lbpolicylabel" {
  labelname       = "tf_lbpolicylabel"
  policylabeltype = "HTTP"
  comment         = "test label"
}

data "citrixadc_lbpolicylabel" "tf_lbpolicylabel" {
  labelname  = citrixadc_lbpolicylabel.tf_lbpolicylabel.labelname
  depends_on = [citrixadc_lbpolicylabel.tf_lbpolicylabel]
}
`

func TestAccLbpolicylabelDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLbpolicylabelDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lbpolicylabel.tf_lbpolicylabel", "labelname", "tf_lbpolicylabel"),
					resource.TestCheckResourceAttr("data.citrixadc_lbpolicylabel.tf_lbpolicylabel", "policylabeltype", "HTTP"),
				),
			},
		},
	})
}
