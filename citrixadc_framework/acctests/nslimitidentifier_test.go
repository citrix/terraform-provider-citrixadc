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

const testAccNslimitidentifier_add = `
	resource "citrixadc_nslimitidentifier" "tf_nslimitidentifier" {
		limitidentifier  = "tf_nslimitidentifier"
		threshold        = 1
		timeslice        = 1000
		limittype        = "BURSTY"
		mode             = "REQUEST_RATE"
		maxbandwidth     = 0
		trapsintimeslice = 1
	}
`
const testAccNslimitidentifier_update = `
	resource "citrixadc_nslimitidentifier" "tf_nslimitidentifier" {
		limitidentifier  = "tf_nslimitidentifier"
		threshold        = 2
		timeslice        = 2000
		limittype        = "BURSTY"
		mode             = "REQUEST_RATE"
		maxbandwidth     = 0
		trapsintimeslice = 1
	}
`

func TestAccNslimitidentifier_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNslimitidentifierDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNslimitidentifier_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNslimitidentifierExist("citrixadc_nslimitidentifier.tf_nslimitidentifier", nil),
					resource.TestCheckResourceAttr("citrixadc_nslimitidentifier.tf_nslimitidentifier", "limitidentifier", "tf_nslimitidentifier"),
					resource.TestCheckResourceAttr("citrixadc_nslimitidentifier.tf_nslimitidentifier", "threshold", "1"),
					resource.TestCheckResourceAttr("citrixadc_nslimitidentifier.tf_nslimitidentifier", "timeslice", "1000"),
				),
			},
			{
				Config: testAccNslimitidentifier_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNslimitidentifierExist("citrixadc_nslimitidentifier.tf_nslimitidentifier", nil),
					resource.TestCheckResourceAttr("citrixadc_nslimitidentifier.tf_nslimitidentifier", "limitidentifier", "tf_nslimitidentifier"),
					resource.TestCheckResourceAttr("citrixadc_nslimitidentifier.tf_nslimitidentifier", "threshold", "2"),
					resource.TestCheckResourceAttr("citrixadc_nslimitidentifier.tf_nslimitidentifier", "timeslice", "2000"),
				),
			},
		},
	})
}

func testAccCheckNslimitidentifierExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nslimitidentifier name is set")
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
		data, err := client.FindResource(service.Nslimitidentifier.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nslimitidentifier %s not found", n)
		}

		return nil
	}
}

func testAccCheckNslimitidentifierDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nslimitidentifier" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Nslimitidentifier.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nslimitidentifier %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
