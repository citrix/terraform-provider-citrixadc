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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// nsconfig is a singleton that always exists on the appliance, so the datasource
// needs no resource dependency.
const testAccNsconfigDataSource_basic = `
	data "citrixadc_nsconfig" "tf_nsconfig" {
	}
`

func TestAccNsconfigDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNsconfigDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nsconfig.tf_nsconfig", "id", "nsconfig-config"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nsconfig.tf_nsconfig", "ipaddress"),
				),
			},
		},
	})
}
