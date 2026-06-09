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

// lbmonitor_service_binding is a "no-GET" binding (NITRO Pattern 13):
//   - Create  = UpdateUnnamedResource (HTTP PUT) — the NITRO "add" verb.
//   - Read    = no-op; NITRO exposes no get/get(all)/count endpoint, so prior
//     state is preserved unchanged. The binding CANNOT be read back.
//   - Update  = no-op; every attribute is RequiresReplace.
//   - Delete  = DeleteResourceWithArgs (monitorname URL key + servicename arg).
//   - The datasource was removed by FeatureDeveloper (no GET endpoint), so there
//     is intentionally NO datasource acceptance test here.
//
// Because there is no GET-by-id, the Exist check below is STATE-ONLY: it only
// asserts that the resource recorded an ID in Terraform state. It deliberately
// does NOT call client.FindResource on the binding, since the appliance offers
// no endpoint to look one up. For the same reason CheckDestroy is omitted — we
// cannot query the appliance to confirm the binding was removed.

const testAccLbmonitorServiceBinding_basic = `

resource "citrixadc_lbmonitor" "tf_lbmonitor" {
  monitorname = "tf_test_lbmon_svc_binding"
  type        = "PING"
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_test_lbmon_svc_binding_lb"
  ipv46       = "10.202.11.11"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_service" "tf_service" {
  name        = "tf_test_lbmon_svc_binding_svc"
  lbvserver   = citrixadc_lbvserver.tf_lbvserver.name
  ip          = "10.202.22.12"
  port        = 80
  servicetype = "HTTP"

  depends_on = [citrixadc_lbvserver.tf_lbvserver]
}

resource "citrixadc_lbmonitor_service_binding" "tf_binding" {
  monitorname = citrixadc_lbmonitor.tf_lbmonitor.monitorname
  servicename = citrixadc_service.tf_service.name

  depends_on = [
    citrixadc_lbmonitor.tf_lbmonitor,
    citrixadc_service.tf_service,
  ]
}
`

func TestAccLbmonitorServiceBinding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// No CheckDestroy: lbmonitor_service_binding has no NITRO GET endpoint,
		// so the binding's removal cannot be confirmed against the appliance.
		Steps: []resource.TestStep{
			{
				Config: testAccLbmonitorServiceBinding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmonitorServiceBindingExist("citrixadc_lbmonitor_service_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor_service_binding.tf_binding", "monitorname", "tf_test_lbmon_svc_binding"),
					resource.TestCheckResourceAttr("citrixadc_lbmonitor_service_binding.tf_binding", "servicename", "tf_test_lbmon_svc_binding_svc"),
					// resource.TestCheckResourceAttr("citrixadc_lbmonitor_service_binding.tf_binding", "weight", "10"),
				),
			},
		},
	})
}

// testAccCheckLbmonitorServiceBindingExist is STATE-ONLY by design.
//
// lbmonitor_service_binding has no NITRO GET endpoint (add/delete only), so we
// cannot call client.FindResource to confirm the binding on the appliance. We
// verify only that the resource exists in Terraform state with a non-empty ID.
func testAccCheckLbmonitorServiceBindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbmonitor_service_binding ID is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// No NITRO GET endpoint for this binding (Pattern 13): the appliance
		// exposes no get/get(all)/count, so we intentionally do NOT call
		// client.FindResource. Presence of a non-empty state ID is the only
		// verification possible here.
		return nil
	}
}
