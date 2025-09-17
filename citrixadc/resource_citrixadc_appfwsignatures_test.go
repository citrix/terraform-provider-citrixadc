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
)

const testAccAppfwsignatures_basic = `
	resource "citrixadc_systemfile" "tf_signature" {
		filename     = "appfw_signatures.xml"
		filelocation = "/var/tmp"
		filecontent  = file("testdata/appfw_signatures.xml")
	}
	resource "citrixadc_appfwsignatures" "tf_appfwsignatures" {
		name       = "tf_appfwsignatures"
		src        = "local://appfw_signatures.xml"
		depends_on = [citrixadc_systemfile.tf_signature]
		comment    = "TestingExample"
	}
`

func TestAccAppfwsignatures_basic(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAppfwsignaturesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwsignatures_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwsignaturesExist("citrixadc_appfwsignatures.tf_appfwsignatures", nil),
				),
			},
		},
	})
}

func testAccCheckAppfwsignaturesExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwsignatures name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Appfwsignatures.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appfwsignatures %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwsignaturesDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwsignatures" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Appfwsignatures.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwsignatures %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
