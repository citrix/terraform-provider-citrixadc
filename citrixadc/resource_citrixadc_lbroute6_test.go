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
	"log"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const testAccLbroute6_basic = `
resource "citrixadc_nsip6" "tf1_nsip6" {
    ipv6address = "22::1/64"
	vserver = "DISABLED"
}

resource "citrixadc_nsip6" "tf2_nsip6" {
    ipv6address = "33::1/64"
	vserver = "DISABLED"
}

resource "citrixadc_nsip6" "tf3_nsip6" {
    ipv6address = "44::1/64"
	vserver = "DISABLED"
}

resource "citrixadc_lbvserver" "llb6" {
    name = "llb6"
    servicetype = "ANY"
    persistencetype = "NONE"
    lbmethod = "ROUNDROBIN"
}

resource "citrixadc_service" "r4" {
    name = "r4"
    ip = "22::2"
    servicetype  = "ANY"
    port = 65535

    depends_on = [citrixadc_nsip6.tf1_nsip6]
}

resource "citrixadc_service" "r5" {
    name = "r5"
    ip = "33::2"
    servicetype  = "ANY"
    port = 65535

    depends_on = [citrixadc_nsip6.tf2_nsip6]

}

resource "citrixadc_service" "r6" {
    name = "r6"
    ip = "44::2"
    servicetype  = "ANY"
    port = 65535

    depends_on = [citrixadc_nsip6.tf3_nsip6]

}

resource "citrixadc_lbvserver_service_binding" "tf_binding4" {
  name = citrixadc_lbvserver.llb6.name
  servicename = citrixadc_service.r4.name
  weight = 10
}

resource "citrixadc_lbvserver_service_binding" "tf_binding5" {
  name = citrixadc_lbvserver.llb6.name
  servicename = citrixadc_service.r5.name
  weight = 10
}

resource "citrixadc_lbvserver_service_binding" "tf_binding6" {
  name = citrixadc_lbvserver.llb6.name
  servicename = citrixadc_service.r6.name
  weight = 10
}

resource "citrixadc_lbroute6" "demo_route6" {
    network = "66::/64"
    gatewayname = citrixadc_lbvserver.llb6.name
}`

func TestAccLbroute6_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLbroute6Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbroute6_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbroute6Exist("citrixadc_lbroute6.demo_route6", nil),
				),
			},
		},
	})
}

func testAccCheckLbroute6Exist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbroute6 name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		network := rs.Primary.ID
		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

		findParams := service.FindParams{
			ResourceType: service.Lbroute6.Type(),
		}

		dataArray, err := nsClient.FindResourceArrayWithParams(findParams)

		foundIndex := -1
		for i, lbroute6 := range dataArray {
			if lbroute6["network"] == network {
				foundIndex = i
				break
			}
		}

		if err != nil {
			log.Printf("[WARN] citrix-provider: acceptance test: Clearing lbroute6 state %s", network)
			return err
		}

		if foundIndex == -1 {
			return fmt.Errorf("Could not find lbroute6 with network %v", network)
		}

		return nil
	}
}

func testAccCheckLbroute6Destroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbroute6" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		network := rs.Primary.ID

		findParams := service.FindParams{
			ResourceType: service.Lbroute6.Type(),
		}
		dataArray, err := nsClient.FindResourceArrayWithParams(findParams)

		if err != nil {
			return err
		}

		foundIndex := -1
		for i, lbroute6 := range dataArray {
			if lbroute6["network"] == network {
				foundIndex = i
				break
			}
		}

		if foundIndex != -1 {
			return fmt.Errorf("LB route6 still exists with network %v", network)
		}

	}

	return nil
}
