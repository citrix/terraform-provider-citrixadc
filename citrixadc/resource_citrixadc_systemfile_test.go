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
	"net/url"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccSystemfile_basic(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSystemfileDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSystemfile_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemfileExist("citrixadc_systemfile.tf_file", nil, []string{"/var/tmp", "hello.txt"}),
				),
			},
			resource.TestStep{
				Config: testAccSystemfile_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemfileExist("citrixadc_systemfile.tf_file", nil, []string{"/tmp", "helloandbye.txt"}),
				),
			},
		},
	})
}

func testAccCheckSystemfileExist(n string, id *string, pathData []string) resource.TestCheckFunc {
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

		argsMap := make(map[string]string)
		argsMap["filelocation"] = url.QueryEscape(pathData[0])
		argsMap["filename"] = url.QueryEscape(pathData[1])
		findParams := service.FindParams{
			ResourceType: "systemfile",
			ArgsMap:      argsMap,
		}
		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		dataArray, err := nsClient.FindResourceArrayWithParams(findParams)

		if err != nil {
			return err
		}

		if len(dataArray) == 0 {
			return fmt.Errorf("systemfile %s not found", n)
		}

		return nil
	}
}

func testAccCheckSystemfileDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_systemfile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Systemfile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccSystemfile_basic_step1 = `

resource "citrixadc_systemfile" "tf_file" {
    filename = "hello.txt"
    filelocation = "/var/tmp"
    filecontent = "hello"
}
`

const testAccSystemfile_basic_step2 = `

resource "citrixadc_systemfile" "tf_file" {
    filename = "helloandbye.txt"
    filelocation = "/tmp"
    filecontent = "hello and goodbye"
}
`
