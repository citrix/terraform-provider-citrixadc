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
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// NOTE on the nstrace_start resource:
//   - nstrace is a NITRO object supporting multiple actions (start / stop) plus a
//     get. Mirroring the systemscalablemgmtthreads package, each action is its own
//     action-only resource. This one wraps POST /nstrace?action=start, which begins
//     a packet trace with the configured options.
//   - There is no `set nstrace` op (NITRO errorcode 1088), so it is a fire-once
//     action: Create starts the trace, Read/Update are no-ops, and Delete is a
//     state-only removal that does NOT stop the trace (use citrixadc_nstrace_stop).
//   - The Check verifies the trace reached the RUNNING state on the appliance; the
//     CheckDestroy stops the trace so no live capture is left running on the box.

const testAccNstraceStart_basic = `
resource "citrixadc_nstrace_start" "tf_start" {
	nf          = 1
	time        = 60
	size        = 0
	filename    = "tf_nstrace_acc"
	traceformat = "NSCAP"
	filesize    = 1024
}
`

func TestAccNstraceStart_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// nstrace has no delete; the resource Delete is a state-only no-op, so stop
		// the trace here to leave the appliance clean (no running capture).
		CheckDestroy: testAccNstraceStopCleanup,
		Steps: []resource.TestStep{
			{
				Config: testAccNstraceStart_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("citrixadc_nstrace_start.tf_start", "id"),
					testAccCheckNstraceState("RUNNING"),
				),
			},
		},
	})
}

// nstraceCurrentState returns the appliance's current nstrace state (RUNNING /
// STOPPED), read from the get(all) response.
func nstraceCurrentState(client *service.NitroClient) (string, error) {
	data, err := client.FindResource(service.Nstrace.Type(), "")
	if err != nil {
		return "", err
	}
	if v, ok := data["state"]; ok && v != nil {
		return strings.TrimSpace(fmt.Sprintf("%v", v)), nil
	}
	return "", nil
}

// testAccCheckNstraceState asserts the live nstrace state on the appliance.
func testAccCheckNstraceState(want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		got, err := nstraceCurrentState(client)
		if err != nil {
			return err
		}
		if got != want {
			return fmt.Errorf("nstrace state = %q, want %q", got, want)
		}
		return nil
	}
}

// testAccNstraceStopCleanup stops any running trace (the resource Delete does not)
// and verifies the appliance is left in the STOPPED state. stop is idempotent.
func testAccNstraceStopCleanup(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}
	// stop is idempotent (succeeds even if nothing is running).
	_ = client.ActOnResource(service.Nstrace.Type(), map[string]interface{}{}, "stop")
	got, err := nstraceCurrentState(client)
	if err != nil {
		return err
	}
	if got == "RUNNING" {
		return fmt.Errorf("nstrace still RUNNING after stop cleanup")
	}
	return nil
}
