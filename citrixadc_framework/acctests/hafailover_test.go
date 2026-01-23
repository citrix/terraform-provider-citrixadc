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

const testAccHafailover_basic = `

	resource "citrixadc_hafailover" "tf_failover" {
		
		ipaddress = "10.222.74.152"
		state = "Secondary"
		force = true
	}
`

const testAccHafailover_basic_update = `

	resource "citrixadc_hafailover" "tf_failover" {
		
		ipaddress = "10.222.74.152"
		state = "Primary"
		force = true
	}
`

func TestAccHafailover_basic(t *testing.T) {
	if adcTestbed != "HA_PAIR" {
		t.Skipf("ADC testbed is %s. Expected HA_PAIR.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccHafailover_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("citrixadc_hafailover.tf_failover", "state", "Secondary"),
				),
			},
			{
				Config: testAccHafailover_basic_update,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("citrixadc_hafailover.tf_failover", "state", "Primary"),
				),
			},
		},
	})
}
