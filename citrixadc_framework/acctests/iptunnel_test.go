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

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccIptunnel_basic_step1 = `
resource "citrixadc_iptunnel" "tf_iptunnel" {
    name = "tf_iptunnel"
	protocol = "GENEVE"
    remote = "66.0.0.11"
    remotesubnetmask = "255.255.255.255"
    local = "*"
	vnid = 100
	tosinherit = "DISABLED"
	destport = 1088
	vlantagging = "DISABLED"
}
`

const testAccIptunnel_basic_step2 = `
resource "citrixadc_iptunnel" "tf_iptunnel" {
    name = "tf_iptunnel"
	protocol = "GENEVE"
    remote = "66.0.0.10"
    remotesubnetmask = "255.255.255.255"
    local = "*"
	vnid = 100
	tosinherit = "ENABLED"
	destport = 2088
	vlantagging = "ENABLED"
}
`

func TestAccIptunnel_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIptunnelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIptunnel_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIptunnelExist("citrixadc_iptunnel.tf_iptunnel", nil),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "name", "tf_iptunnel"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "protocol", "GENEVE"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "remote", "66.0.0.11"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "remotesubnetmask", "255.255.255.255"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "local", "*"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "vnid", "100"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "tosinherit", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "destport", "1088"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "vlantagging", "DISABLED"),
				),
			},
			{
				Config: testAccIptunnel_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIptunnelExist("citrixadc_iptunnel.tf_iptunnel", nil),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "name", "tf_iptunnel"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "protocol", "GENEVE"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "remote", "66.0.0.10"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "remotesubnetmask", "255.255.255.255"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "local", "*"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "vnid", "100"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "tosinherit", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "destport", "2088"),
					resource.TestCheckResourceAttr("citrixadc_iptunnel.tf_iptunnel", "vlantagging", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckIptunnelExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No iptunnel name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Iptunnel.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("iptunnel %s not found", n)
		}

		return nil
	}
}

func testAccCheckIptunnelDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_iptunnel" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Iptunnel.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("iptunnel %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
