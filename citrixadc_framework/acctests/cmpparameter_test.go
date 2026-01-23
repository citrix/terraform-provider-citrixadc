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

const testAccCmpparameter_basic = `


resource "citrixadc_cmpparameter" "tf_cmpparameter" {
	cmplevel    = "optimal"
	quantumsize = 20
	servercmp   = "OFF"
	randomgzipfilenameminlength = "12"
	randomgzipfilenamemaxlength = "20"
	randomgzipfilename = "ENABLED"
	}
`
const testAccCmpparameter_update = `


	resource "citrixadc_cmpparameter" "tf_cmpparameter" {
		cmplevel    = "bestspeed"
		quantumsize = 30
		servercmp   = "ON"
		randomgzipfilenameminlength = "14"
		randomgzipfilenamemaxlength = "22"
		randomgzipfilename = "DISABLED"
	}
`

func TestAccCmpparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccCmpparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCmpparameterExist("citrixadc_cmpparameter.tf_cmpparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_cmpparameter", "cmplevel", "optimal"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_cmpparameter", "quantumsize", "20"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_cmpparameter", "servercmp", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_cmpparameter", "randomgzipfilenameminlength", "12"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_cmpparameter", "randomgzipfilenamemaxlength", "20"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_cmpparameter", "randomgzipfilename", "ENABLED"),
				),
			},
			{
				Config: testAccCmpparameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCmpparameterExist("citrixadc_cmpparameter.tf_cmpparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_cmpparameter", "cmplevel", "bestspeed"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_cmpparameter", "quantumsize", "30"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_cmpparameter", "servercmp", "ON"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_cmpparameter", "randomgzipfilenameminlength", "14"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_cmpparameter", "randomgzipfilenamemaxlength", "22"),
					resource.TestCheckResourceAttr("citrixadc_cmpparameter.tf_cmpparameter", "randomgzipfilename", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckCmpparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cmpparameter name is set")
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
		data, err := client.FindResource(service.Cmpparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("cmpparameter %s not found", n)
		}

		return nil
	}
}
