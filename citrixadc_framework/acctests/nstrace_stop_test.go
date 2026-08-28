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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// nstrace_stop wraps POST /nstrace?action=stop (stops a running packet trace). The
// stop action carries an empty payload and is idempotent (stopping when nothing is
// running succeeds). Action-only: no per-action GET, Delete is a state-only
// removal. See nstrace_start_test.go for the multi-action notes.
const testAccNstraceStop_basic = `
resource "citrixadc_nstrace_stop" "tf_stop" {
}
`

func TestAccNstraceStop_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNstraceStop_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("citrixadc_nstrace_stop.tf_stop", "id"),
					// stop is idempotent -> the appliance ends in the STOPPED state.
					testAccCheckNstraceState("STOPPED"),
				),
			},
		},
	})
}
