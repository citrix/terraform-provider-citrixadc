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

// nslimitsessions is an ACTION-ONLY resource: Create performs the NITRO
// `clear` action (POST ?action=clear) against an existing rate-limit
// identifier. There is no NITRO add/delete endpoint for the resource itself,
// so the test does a single apply with a state-only Exist check (no NITRO
// verify, no CheckDestroy).
//
// NOTE: To be MEANINGFUL, limitidentifier must reference a REAL
// nslimitidentifier configured on the appliance. Clearing a non-existent
// identifier may error on the ADC. Replace TODO_PLACEHOLDER with a real
// rate-limit identifier name before running.
const testAccNslimitsessions_basic_step1 = `
	resource "citrixadc_nslimitidentifier" "tf_nslimitidentifier" {
		limitidentifier  = "tf_nslimitidentifier"
		threshold        = 1
		timeslice        = 1000
		limittype        = "BURSTY"
		mode             = "REQUEST_RATE"
		maxbandwidth     = 0
		trapsintimeslice = 1
	}
	resource "citrixadc_nslimitsessions" "tf_nslimitsessions" {
	limitidentifier = citrixadc_nslimitidentifier.tf_nslimitidentifier.limitidentifier
	}

`

func TestAccNslimitsessions_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// Action-only resource: no CheckDestroy (NITRO exposes no delete endpoint).
		Steps: []resource.TestStep{
			{
				Config: testAccNslimitsessions_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNslimitsessionsExist("citrixadc_nslimitsessions.tf_nslimitsessions", nil),
					resource.TestCheckResourceAttrSet("citrixadc_nslimitsessions.tf_nslimitsessions", "id"),
					resource.TestCheckResourceAttr("citrixadc_nslimitsessions.tf_nslimitsessions", "limitidentifier", "tf_nslimitidentifier"),
				),
			},
		},
	})
}

// testAccCheckNslimitsessionsExist is a STATE-ONLY check. nslimitsessions is
// action-only (clear) and sessions are transient, so there is no NITRO GET to
// verify against; we only confirm the resource has a non-empty synthetic ID.
func testAccCheckNslimitsessionsExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nslimitsessions ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// No NITRO FindResource call: clear is an action and sessions are
		// transient, so there is nothing stable to read back from the ADC.
		return nil
	}
}

// Datasource test for nslimitsessions. Gated with t.Skip because it requires a
// REAL nslimitidentifier to exist on the appliance with active sessions to
// return meaningful data (mirrors the runtime/session skip precedent for
// transient/session resources).
const testAccNslimitsessionsDataSource_basic = `
	resource "citrixadc_nslimitidentifier" "tf_nslimitidentifier" {
		limitidentifier  = "tf_nslimitidentifier"
		threshold        = 1
		timeslice        = 1000
		limittype        = "BURSTY"
		mode             = "REQUEST_RATE"
		maxbandwidth     = 0
		trapsintimeslice = 1
	}
	resource "citrixadc_nslimitsessions" "tf_nslimitsessions" {
	limitidentifier = citrixadc_nslimitidentifier.tf_nslimitidentifier.limitidentifier
	}
	
	data "citrixadc_nslimitsessions" "tf_nslimitsessions" {
	limitidentifier = citrixadc_nslimitidentifier.tf_nslimitidentifier.limitidentifier
	}
`

func TestAccNslimitsessionsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNslimitsessionsDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nslimitsessions.tf_nslimitsessions", "limitidentifier", "tf_nslimitidentifier"),
				),
			},
		},
	})
}
