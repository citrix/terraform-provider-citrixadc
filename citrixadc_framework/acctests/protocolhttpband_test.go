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

// --- Unset test -------------------------------------------------------------
//
// protocolhttpband has unset support wired for reqbandsize (default 100) and
// respbandsize (default 1024). Step 1 sets both to non-default values; step 2
// removes them from config so the provider issues ?action=unset, reverting each
// to its ADC default. The revert is asserted both in Terraform state and, for
// independent appliance-level confirmation, directly on the ADC via the
// stats show's `bandrange` field (keyed by the mandatory `type` filter:
// type=REQUEST reports reqbandsize, type=RESPONSE reports respbandsize).

const testAccProtocolhttpband_unset_step1 = `
resource "citrixadc_protocolhttpband" "tf_unset" {
  reqbandsize  = 200
  respbandsize = 2048
}
`

const testAccProtocolhttpband_unset_step2 = `
resource "citrixadc_protocolhttpband" "tf_unset" {
  # reqbandsize / respbandsize removed from config -> provider must unset them,
  # reverting to the ADC defaults (100 / 1024).
}
`

func TestAccProtocolhttpband_unset(t *testing.T) {
	// The resource's other tests (TestAccProtocolhttpband_basic) run on the
	// default standalone testbed with no skip guard, so none is added here.
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// Singleton settings resource with a no-op Delete: nothing to destroy.
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				// Non-default values apply and persist.
				Config: testAccProtocolhttpband_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProtocolhttpbandExist("citrixadc_protocolhttpband.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_protocolhttpband.tf_unset", "reqbandsize", "200"),
					resource.TestCheckResourceAttr("citrixadc_protocolhttpband.tf_unset", "respbandsize", "2048"),
					testAccCheckProtocolhttpbandADCBandrange("REQUEST", "200"),
					testAccCheckProtocolhttpbandADCBandrange("RESPONSE", "2048"),
				),
			},
			{
				// Removing them must unset -> state reverts to NITRO defaults,
				// and the implicit post-apply plan must be empty.
				Config: testAccProtocolhttpband_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProtocolhttpbandExist("citrixadc_protocolhttpband.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_protocolhttpband.tf_unset", "reqbandsize", "100"),
					resource.TestCheckResourceAttr("citrixadc_protocolhttpband.tf_unset", "respbandsize", "1024"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckProtocolhttpbandADCBandrange("REQUEST", "100"),
					testAccCheckProtocolhttpbandADCBandrange("RESPONSE", "1024"),
				),
			},
		},
	})
}

// testAccCheckProtocolhttpbandADCBandrange asserts, directly on the appliance,
// the configured band size for the given band type (REQUEST -> reqbandsize,
// RESPONSE -> respbandsize). NITRO exposes only a stats-only `show` keyed by a
// mandatory `type` filter whose `bandrange` field echoes the configured band
// size for that type, so this proves the unset actually reverted the value on
// the ADC rather than just in Terraform state.
func testAccCheckProtocolhttpbandADCBandrange(bandType, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		findParams := service.FindParams{
			ResourceType: service.Protocolhttpband.Type(),
			ArgsMap:      map[string]string{"type": bandType},
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}
		if len(dataArr) == 0 {
			return fmt.Errorf("protocolhttpband (type=%s) not found on appliance", bandType)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", dataArr[0]["bandrange"]))
		if got != want {
			return fmt.Errorf("protocolhttpband type=%s: appliance bandrange = %q, want %q (unset did not revert it)", bandType, got, want)
		}
		return nil
	}
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
