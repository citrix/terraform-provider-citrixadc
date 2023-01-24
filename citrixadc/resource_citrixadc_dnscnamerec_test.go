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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
	"net/url"
)

const testAccDnscnamerec_basic = `


resource "citrixadc_dnscnamerec" "dnscnamerec" {
	aliasname = "citrixadc.cloud.com"
    canonicalname = "ctxwsp-citrixadc-fdproxy-global.trafficmanager.net"
    ttl = 3600
}
`

func TestAccDnscnamerec_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDnscnamerecDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDnscnamerec_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnscnamerecExist("citrixadc_dnscnamerec.dnscnamerec", nil),
					resource.TestCheckResourceAttr("citrixadc_dnscnamerec.dnscnamerec", "aliasname", "citrixadc.cloud.com"),
					resource.TestCheckResourceAttr("citrixadc_dnscnamerec.dnscnamerec", "canonicalname", "ctxwsp-citrixadc-fdproxy-global.trafficmanager.net"),
					resource.TestCheckResourceAttr("citrixadc_dnscnamerec.dnscnamerec", "ttl", "3600"),

				),
			},
		},
	})
}

func testAccCheckDnscnamerecExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnscnamerec name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Dnscnamerec.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("dnscnamerec %s not found", n)
		}

		return nil
	}
}

func testAccCheckDnscnamerecDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnscnamerec" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		argsMap := make(map[string]string)
		argsMap["ecssubnet"] = url.QueryEscape(rs.Primary.Attributes["ecssubnet"])
		findParams := service.FindParams{
			ResourceType: service.Dnscnamerec.Type(),
			ArgsMap:      argsMap,
		}
		_, err := nsClient.FindResourceArrayWithParams(findParams)

		if err == nil {
			return fmt.Errorf("dnscnamerec %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
