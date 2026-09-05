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

// The citrixadc_nsmigration data source reads the live migration singleton via the
// NITRO get (dumpsession + the read-only migration status fields). Unlike the
// start/stop/complete actions, the get IS supported on standalone appliances
// (migrationstatus reports "Migration is not yet started"), so this test runs by
// default. The actions are separate action-only resources
// (citrixadc_nsmigration_start / _stop / _complete).
const testAccNsmigrationDataSource_basic = `
data "citrixadc_nsmigration" "tf_nsmigration" {
}
`

func TestAccNsmigrationDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNsmigrationDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.citrixadc_nsmigration.tf_nsmigration", "id"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nsmigration.tf_nsmigration", "dumpsession"),
				),
			},
		},
	})
}
