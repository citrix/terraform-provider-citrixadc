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

const testAccSslcertfile_basic = `
	resource "citrixadc_sslcertfile" "tf_sslcertfile" {
		name = "tf_sslcertfile"
		src = "local://certificate1.crt"
	}
`

func TestAccSslcertfile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSslcertfileDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSslcertfile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertfileExist("citrixadc_sslcertfile.tf_sslcertfile", nil),
				),
			},
		},
	})
}

func testAccCheckSslcertfileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslcertfile name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		dataArr, err := nsClient.FindAllResources(service.Sslcertfile.Type())

		if err != nil {
			return err
		}
		found := false
		for _, v := range dataArr {
			if v["name"].(string) == rs.Primary.Attributes["name"] {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("sslcertfile %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslcertfileDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslcertfile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		dataArr, err := nsClient.FindAllResources(service.Sslcertfile.Type())
		
		if err != nil {
			return err
		}
		
		found := false
		for _, v := range dataArr {
			if v["name"].(string) == rs.Primary.Attributes["name"] {
				found = true
				break
			}
		}
		if found {
			return fmt.Errorf("sslcertfile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
