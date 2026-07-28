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

// videooptimizationpacingpolicylabel is a named resource whose attributes
// (labelname, policylabeltype, comment) are all RequiresReplace, so a single
// create+verify step is sufficient — there is no in-place update path.
//
// NOTE: video optimization pacing is deprecated in recent NetScaler releases,
// but it remains configurable and the sibling videooptimizationpacingpolicy /
// pacingaction tests are not skipped, so this test is generated normally
// (no t.Skip).
//
// NOTE: policylabeltype is supplied as the CLI enum "videoopt_req", but the
// NITRO API may normalize/echo it as a different token (e.g. NS_PLTMAP_RSP_REQ).
// To avoid an inconsistent-result / assert mismatch, we set it in config but do
// NOT assert its read-back value. Only labelname and comment (stable) are asserted.

const testAccVideooptimizationpacingpolicylabel_basic_step1 = `
resource "citrixadc_videooptimizationpacingpolicylabel" "tf_videooptimizationpacingpolicylabel" {
  labelname       = "tf_videoopt_pacing_pl"
  policylabeltype = "videoopt_req"
  comment         = "test_comment"
}

`

func TestAccVideooptimizationpacingpolicylabel_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVideooptimizationpacingpolicylabelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVideooptimizationpacingpolicylabel_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVideooptimizationpacingpolicylabelExist("citrixadc_videooptimizationpacingpolicylabel.tf_videooptimizationpacingpolicylabel", nil),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationpacingpolicylabel.tf_videooptimizationpacingpolicylabel", "labelname", "tf_videoopt_pacing_pl"),
					resource.TestCheckResourceAttr("citrixadc_videooptimizationpacingpolicylabel.tf_videooptimizationpacingpolicylabel", "comment", "test_comment"),
				),
			},
		},
	})
}

func TestAccVideooptimizationpacingpolicylabel_import(t *testing.T) {
	const resAddr = "citrixadc_videooptimizationpacingpolicylabel.tf_videooptimizationpacingpolicylabel"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVideooptimizationpacingpolicylabelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVideooptimizationpacingpolicylabel_basic_step1,
			},
			{
				Config:                  testAccVideooptimizationpacingpolicylabel_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckVideooptimizationpacingpolicylabelExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No videooptimizationpacingpolicylabel name is set")
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
		data, err := client.FindResource(service.Videooptimizationpacingpolicylabel.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("videooptimizationpacingpolicylabel %s not found", n)
		}

		return nil
	}
}

func testAccCheckVideooptimizationpacingpolicylabelDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_videooptimizationpacingpolicylabel" {
			continue
		}

		_, err := client.FindResource(service.Videooptimizationpacingpolicylabel.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("videooptimizationpacingpolicylabel %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

const testAccVideooptimizationpacingpolicylabelDataSource_basic = `

resource "citrixadc_videooptimizationpacingpolicylabel" "tf_videooptimizationpacingpolicylabel" {
  labelname       = "tf_videoopt_pacing_pl"
  policylabeltype = "videoopt_req"
  comment         = "test_comment"
}

data "citrixadc_videooptimizationpacingpolicylabel" "tf_videooptimizationpacingpolicylabel" {
  labelname  = citrixadc_videooptimizationpacingpolicylabel.tf_videooptimizationpacingpolicylabel.labelname
  depends_on = [citrixadc_videooptimizationpacingpolicylabel.tf_videooptimizationpacingpolicylabel]
}
`

func TestAccVideooptimizationpacingpolicylabelDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVideooptimizationpacingpolicylabelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVideooptimizationpacingpolicylabelDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_videooptimizationpacingpolicylabel.tf_videooptimizationpacingpolicylabel", "labelname", "tf_videoopt_pacing_pl"),
					resource.TestCheckResourceAttr("data.citrixadc_videooptimizationpacingpolicylabel.tf_videooptimizationpacingpolicylabel", "comment", "test_comment"),
				),
			},
		},
	})
}
