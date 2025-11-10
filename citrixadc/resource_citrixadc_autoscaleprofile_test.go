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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

const testAccAutoscaleprofile_basic = `


resource "citrixadc_autoscaleprofile" "tf_autoscaleprofile" {
	name         = "my_autoscaleprofile"
	type         = "CLOUDSTACK"
	apikey       = "7c177611-4a18-42b0-a7c5-bfd811fd590f"
	url          = "www.service.example.com"
	sharedsecret = "YZEH6jkTqZWQ8r0o6kWj0mWruN3vXbtT"
	}
`
const testAccAutoscaleprofile_update = `


resource "citrixadc_autoscaleprofile" "tf_autoscaleprofile" {
	name         = "my_autoscaleprofile"
	type         = "CLOUDSTACK"
	apikey       = "88e0ae91-4cd0-4dd5-8fa1-dcd38165f4a2"
	url          = "www.service2.example.com"
	sharedsecret = "vruE8whIW8qnAvUGtT3EpmeIFp690nGo"
	}
`

func TestAccAutoscaleprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckAutoscaleprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAutoscaleprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAutoscaleprofileExist("citrixadc_autoscaleprofile.tf_autoscaleprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.tf_autoscaleprofile", "name", "my_autoscaleprofile"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.tf_autoscaleprofile", "type", "CLOUDSTACK"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.tf_autoscaleprofile", "apikey", "7c177611-4a18-42b0-a7c5-bfd811fd590f"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.tf_autoscaleprofile", "url", "www.service.example.com"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.tf_autoscaleprofile", "sharedsecret", "YZEH6jkTqZWQ8r0o6kWj0mWruN3vXbtT"),
				),
			},
			{
				Config: testAccAutoscaleprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAutoscaleprofileExist("citrixadc_autoscaleprofile.tf_autoscaleprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.tf_autoscaleprofile", "name", "my_autoscaleprofile"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.tf_autoscaleprofile", "type", "CLOUDSTACK"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.tf_autoscaleprofile", "apikey", "88e0ae91-4cd0-4dd5-8fa1-dcd38165f4a2"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.tf_autoscaleprofile", "url", "www.service2.example.com"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.tf_autoscaleprofile", "sharedsecret", "vruE8whIW8qnAvUGtT3EpmeIFp690nGo"),
				),
			},
		},
	})
}

func testAccCheckAutoscaleprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No autoscaleprofile name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Autoscaleprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("autoscaleprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckAutoscaleprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_autoscaleprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Autoscaleprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("autoscaleprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
