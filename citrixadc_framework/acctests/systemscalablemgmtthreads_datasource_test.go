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

// The citrixadc_systemscalablemgmtthreads data source reads the live feature state
// (configuredstate / effectivestate) via the NITRO get. The enable/disable actions
// are separate action-only resources (citrixadc_systemscalablemgmtthreads_enable /
// _disable).
//
// SKIPPED by default: the Scalable Management Threads feature is platform-gated. On
// standard VPX/lab appliances even the get returns NITRO errorcode 1501 "Operation
// not supported on this platform". Remove the t.Skip only on a platform that
// supports the feature (then configuredstate reflects ENABLED/DISABLED).
const testAccSystemscalablemgmtthreadsDataSource_basic = `
data "citrixadc_systemscalablemgmtthreads" "tf_systemscalablemgmtthreads" {
}
`

func TestAccSystemscalablemgmtthreadsDataSource_basic(t *testing.T) {
	t.Skip("Skipping systemscalablemgmtthreads data source test: the Scalable Management Threads feature is platform-gated (NITRO errorcode 1501 \"Operation not supported on this platform\" on standard appliances)")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemscalablemgmtthreadsDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.citrixadc_systemscalablemgmtthreads.tf_systemscalablemgmtthreads", "configuredstate"),
				),
			},
		},
	})
}
