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

const testAccStreamidentifier_basic = `

resource "citrixadc_streamidentifier" "tf_streamidentifier" {
	name         = "my_streamidentifier"
	selectorname = "my_streamselector"
	samplecount  = 10
	sort         = "CONNECTIONS"
	snmptrap     = "ENABLED"
  }
  
`

const testAccStreamidentifier_update = `

resource "citrixadc_streamidentifier" "tf_streamidentifier" {
	name         = "my_streamidentifier"
	selectorname = "my_streamselector"
	samplecount  = 20
	sort         = "REQUESTS"
	snmptrap     = "DISABLED"
  }
  
`

func TestAccStreamidentifier_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckStreamidentifierDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccStreamidentifier_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStreamidentifierExist("citrixadc_streamidentifier.tf_streamidentifier", nil),
					resource.TestCheckResourceAttr("citrixadc_streamidentifier.tf_streamidentifier", "name", "my_streamidentifier"),
					resource.TestCheckResourceAttr("citrixadc_streamidentifier.tf_streamidentifier", "selectorname", "my_streamselector"),
					resource.TestCheckResourceAttr("citrixadc_streamidentifier.tf_streamidentifier", "samplecount", "10"),
					resource.TestCheckResourceAttr("citrixadc_streamidentifier.tf_streamidentifier", "sort", "CONNECTIONS"),
					resource.TestCheckResourceAttr("citrixadc_streamidentifier.tf_streamidentifier", "snmptrap", "ENABLED"),
				),
			},
			{
				Config: testAccStreamidentifier_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStreamidentifierExist("citrixadc_streamidentifier.tf_streamidentifier", nil),
					resource.TestCheckResourceAttr("citrixadc_streamidentifier.tf_streamidentifier", "name", "my_streamidentifier"),
					resource.TestCheckResourceAttr("citrixadc_streamidentifier.tf_streamidentifier", "selectorname", "my_streamselector"),
					resource.TestCheckResourceAttr("citrixadc_streamidentifier.tf_streamidentifier", "samplecount", "20"),
					resource.TestCheckResourceAttr("citrixadc_streamidentifier.tf_streamidentifier", "sort", "REQUESTS"),
					resource.TestCheckResourceAttr("citrixadc_streamidentifier.tf_streamidentifier", "snmptrap", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckStreamidentifierExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No streamidentifier name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Streamidentifier.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("streamidentifier %s not found", n)
		}

		return nil
	}
}

func testAccCheckStreamidentifierDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_streamidentifier" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Streamidentifier.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("streamidentifier %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
