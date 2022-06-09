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
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

const testAccNspbr6_add = `

	resource "citrixadc_iptunnel" "tf_iptunnel" {
		name             = "tf_iptunnel"
		remote           = "66.0.0.11"
		remotesubnetmask = "255.255.255.255"
		local            = "*"
	}
	resource "citrixadc_nspbr6" "tf_nspbr6" {
		name     = "tf_nspbr6"
		action   = "ALLOW"
		protocol = "ICMPV6"
		priority = 20
		state    = "ENABLED"
		iptunnel = citrixadc_iptunnel.tf_iptunnel.name
	}
`
const testAccNspbr6_update = `

	resource "citrixadc_iptunnel" "tf_iptunnel" {
		name             = "tf_iptunnel"
		remote           = "66.0.0.11"
		remotesubnetmask = "255.255.255.255"
		local            = "*"
	}
	resource "citrixadc_nspbr6" "tf_nspbr6" {
		name     = "tf_nspbr6"
		action   = "ALLOW"
		protocol = "TCP"
		priority = 30
		state    = "DISABLED"
		iptunnel = citrixadc_iptunnel.tf_iptunnel.name
	}
`

func TestAccNspbr6_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNspbr6Destroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNspbr6_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNspbr6Exist("citrixadc_nspbr6.tf_nspbr6", nil),
					resource.TestCheckResourceAttr("citrixadc_nspbr6.tf_nspbr6", "name", "tf_nspbr6"),
					resource.TestCheckResourceAttr("citrixadc_nspbr6.tf_nspbr6", "protocol", "ICMPV6"),
					resource.TestCheckResourceAttr("citrixadc_nspbr6.tf_nspbr6", "priority", "20"),
					resource.TestCheckResourceAttr("citrixadc_nspbr6.tf_nspbr6", "state", "ENABLED"),
				),
			},
			resource.TestStep{
				Config: testAccNspbr6_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNspbr6Exist("citrixadc_nspbr6.tf_nspbr6", nil),
					resource.TestCheckResourceAttr("citrixadc_nspbr6.tf_nspbr6", "name", "tf_nspbr6"),
					resource.TestCheckResourceAttr("citrixadc_nspbr6.tf_nspbr6", "protocol", "TCP"),
					resource.TestCheckResourceAttr("citrixadc_nspbr6.tf_nspbr6", "priority", "30"),
					resource.TestCheckResourceAttr("citrixadc_nspbr6.tf_nspbr6", "state", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckNspbr6Exist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nspbr6 name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Nspbr6.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nspbr6 %s not found", n)
		}

		return nil
	}
}

func testAccCheckNspbr6Destroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nspbr6" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Nspbr6.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nspbr6 %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
