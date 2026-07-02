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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// protocolhttpband is a SINGLETON SETTINGS resource (like rnatparam/quicparam):
//   - Create = set, Update = set (real update), Delete = no-op (no NITRO rm verb).
//   - Read is a no-op: NITRO exposes only a stats-only `show` (keyed by a
//     mandatory `type` filter) and no config readback of reqbandsize/respbandsize.
//     Because of that, the Exist check below only verifies presence/state and does
//     NOT call FindResource against the ADC (a stats-only show is not a reliable
//     config readback). Assertions therefore validate the values held in state.
//   - No CheckDestroy (Delete is a no-op, the object cannot be deleted/read back).
//   - No datasource test (datasource removed - no NITRO GET config endpoint).
//   - No ephemeral test (no secret attributes).

const testAccProtocolhttpband_basic_step1 = `
resource "citrixadc_protocolhttpband" "tf_protocolhttpband" {
  reqbandsize  = 100
  respbandsize = 1024
}
`

const testAccProtocolhttpband_basic_step2 = `
resource "citrixadc_protocolhttpband" "tf_protocolhttpband" {
  reqbandsize  = 200
  respbandsize = 2048
}
`

func TestAccProtocolhttpband_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// Singleton settings resource with a no-op Delete: nothing to destroy/verify.
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: testAccProtocolhttpband_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProtocolhttpbandExist("citrixadc_protocolhttpband.tf_protocolhttpband", nil),
					resource.TestCheckResourceAttr("citrixadc_protocolhttpband.tf_protocolhttpband", "reqbandsize", "100"),
					resource.TestCheckResourceAttr("citrixadc_protocolhttpband.tf_protocolhttpband", "respbandsize", "1024"),
				),
			},
			{
				Config: testAccProtocolhttpband_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProtocolhttpbandExist("citrixadc_protocolhttpband.tf_protocolhttpband", nil),
					resource.TestCheckResourceAttr("citrixadc_protocolhttpband.tf_protocolhttpband", "reqbandsize", "200"),
					resource.TestCheckResourceAttr("citrixadc_protocolhttpband.tf_protocolhttpband", "respbandsize", "2048"),
				),
			},
		},
	})
}

// testAccCheckProtocolhttpbandExist verifies the singleton is present in Terraform
// state. It intentionally does NOT call client.FindResource: protocolhttpband has
// no config-readback NITRO endpoint (only a stats-only `show` keyed by a mandatory
// `type` filter), so a GET is not a reliable existence check. State presence plus
// the fixed synthetic ID is the correct signal for this singleton settings resource.
func testAccCheckProtocolhttpbandExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No protocolhttpband ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		return nil
	}
}
