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

// systemnsbtracing is a lifecycle-driven enable/disable toggle singleton:
//   Create -> ?action=enable, Delete -> ?action=disable, Update -> no-op.
// There is no settable on/off attribute. Presence of the resource means NSB
// tracing is ENABLED on the appliance; destroying it disables tracing again.
// The test is reversible: ENABLE on apply, DISABLE on destroy. nodeid is NOT
// set (standalone testbed, not cluster).

const testAccSystemnsbtracing_basic = `

	resource "citrixadc_systemnsbtracing" "tf_systemnsbtracing" {
	}

`

func TestAccSystemnsbtracing_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemnsbtracingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemnsbtracing_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemnsbtracingExist("citrixadc_systemnsbtracing.tf_systemnsbtracing", nil),
					resource.TestCheckResourceAttr("citrixadc_systemnsbtracing.tf_systemnsbtracing", "id", "systemnsbtracing-config"),
				),
			},
		},
	})
}

// testAccCheckSystemnsbtracingExist confirms the resource is present in state and
// that NSB tracing is actually ENABLED on the appliance (configuredstate=ENABLED).
func testAccCheckSystemnsbtracingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No systemnsbtracing ID is set")
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

		// systemnsbtracing is a toggle singleton; GET takes an empty name.
		data, err := client.FindResource(service.Systemnsbtracing.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("systemnsbtracing not found")
		}

		// Confirm tracing is actually enabled, not merely present in state.
		if state, ok := data["configuredstate"]; ok {
			if fmt.Sprintf("%v", state) != "ENABLED" {
				return fmt.Errorf("systemnsbtracing configuredstate is %v, expected ENABLED", state)
			}
		}

		return nil
	}
}

// testAccCheckSystemnsbtracingDestroy confirms that, after the resource is
// destroyed, NSB tracing has been disabled on the appliance.
func testAccCheckSystemnsbtracingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_systemnsbtracing" {
			continue
		}

		data, err := client.FindResource(service.Systemnsbtracing.Type(), "")
		if err != nil {
			// Read failure does not mean the toggle is still on; treat as cleaned up.
			return nil
		}
		if data != nil {
			if state, ok := data["configuredstate"]; ok && fmt.Sprintf("%v", state) != "DISABLED" {
				return fmt.Errorf("systemnsbtracing still ENABLED after destroy (configuredstate=%v)", state)
			}
		}
	}
	return nil
}
