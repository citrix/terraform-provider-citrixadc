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
	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)

func TestAccServicegroup_lbmonitor_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckServicegroup_lbmonitor_bindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccServicegroup_lbmonitor_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroup_lbmonitor_bindingExist("citrixadc_servicegroup_lbmonitor_binding.bind1", nil),
					testAccCheckServicegroup_lbmonitor_bindingExist("citrixadc_servicegroup_lbmonitor_binding.bind2", nil),
				),
			},
			resource.TestStep{
				Config: testAccServicegroup_lbmonitor_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroup_lbmonitor_bindingExist("citrixadc_servicegroup_lbmonitor_binding.bind1", nil),
					testAccCheckServicegroup_lbmonitor_bindingExist("citrixadc_servicegroup_lbmonitor_binding.bind2", nil),
				),
			},
			resource.TestStep{
				Config: testAccServicegroup_lbmonitor_binding_basic_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroup_lbmonitor_bindingExist("citrixadc_servicegroup_lbmonitor_binding.bind1", nil),
				),
			},
		},
	})
}

func testAccCheckServicegroup_lbmonitor_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lb vserver name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

		servicegroupLbmonitorBindingId := rs.Primary.ID
		idSlice := strings.Split(servicegroupLbmonitorBindingId, ",")
		servicegroupName := idSlice[0]
		monitorName := idSlice[1]

		findParams := netscaler.FindParams{
			ResourceType:             "servicegroup_lbmonitor_binding",
			ResourceName:             servicegroupName,
			ResourceMissingErrorCode: 258,
		}

		dataArr, err := nsClient.FindResourceArrayWithParams(findParams)

		if err != nil {
			return err
		}

		found := false

		for _, v := range dataArr {
			if v["monitor_name"].(string) == monitorName {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckServicegroup_lbmonitor_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_servicegroup_lbmonitor_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(netscaler.Servicegroup_lbmonitor_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccServicegroup_lbmonitor_binding_basic_step1 = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_lbvserver"
  ipv46       = "192.168.13.46"
  port        = "80"
  servicetype = "HTTP"
}

resource "citrixadc_lbmonitor" "tfmonitor1" {
  monitorname = "tf-monitor1"
  type        = "HTTP"
}

resource "citrixadc_lbmonitor" "tfmonitor2" {
  monitorname = "tfmonitor2"
  type        = "PING"
}

resource "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "tf_servicegroup"
  lbvservers       = [citrixadc_lbvserver.tf_lbvserver.name]
  servicetype      = "HTTP"
  servicegroupmembers = [
    "192.168.33.33:80:1",
  ]

}

resource "citrixadc_servicegroup_lbmonitor_binding" "bind1" {
    servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
    monitorname = citrixadc_lbmonitor.tfmonitor1.monitorname
    weight = 80
}

resource "citrixadc_servicegroup_lbmonitor_binding" "bind2" {
    servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
    monitorname = citrixadc_lbmonitor.tfmonitor2.monitorname
    weight = 20
}

`

const testAccServicegroup_lbmonitor_binding_basic_step2 = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_lbvserver"
  ipv46       = "192.168.13.46"
  port        = "80"
  servicetype = "HTTP"
}

resource "citrixadc_lbmonitor" "tfmonitor1" {
  monitorname = "tf-monitor1"
  type        = "HTTP"
}

resource "citrixadc_lbmonitor" "tfmonitor2" {
  monitorname = "tfmonitor2"
  type        = "PING"
}

resource "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "tf_servicegroup"
  lbvservers       = [citrixadc_lbvserver.tf_lbvserver.name]
  servicetype      = "HTTP"
  servicegroupmembers = [
    "192.168.33.33:80:1",
  ]

}

resource "citrixadc_servicegroup_lbmonitor_binding" "bind1" {
    servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
    monitorname = citrixadc_lbmonitor.tfmonitor1.monitorname
    weight = 50
}

resource "citrixadc_servicegroup_lbmonitor_binding" "bind2" {
    servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
    monitorname = citrixadc_lbmonitor.tfmonitor2.monitorname
    weight = 50
}

`

const testAccServicegroup_lbmonitor_binding_basic_step3 = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_lbvserver"
  ipv46       = "192.168.13.46"
  port        = "80"
  servicetype = "HTTP"
}

resource "citrixadc_lbmonitor" "tfmonitor1" {
  monitorname = "tf-monitor1"
  type        = "HTTP"
}

resource "citrixadc_lbmonitor" "tfmonitor2" {
  monitorname = "tfmonitor2"
  type        = "PING"
}

resource "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "tf_servicegroup"
  lbvservers       = [citrixadc_lbvserver.tf_lbvserver.name]
  servicetype      = "HTTP"
  servicegroupmembers = [
    "192.168.33.33:80:1",
  ]

}

resource "citrixadc_servicegroup_lbmonitor_binding" "bind1" {
    servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
    monitorname = citrixadc_lbmonitor.tfmonitor1.monitorname
    weight = 50
}

`
