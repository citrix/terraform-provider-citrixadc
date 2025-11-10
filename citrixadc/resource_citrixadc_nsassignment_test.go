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

const testAccNsassignment_add = `
	resource "citrixadc_nsvariable" "tf_nsvariable" {
		name          = "tf_nsvariable"
		type          = "text(20)"
		scope         = "global"
		iffull        = "undef"
		ifvaluetoobig = "undef"
		ifnovalue     = "init"
		comment       = "Testing"
	}
	resource "citrixadc_nsassignment" "tf_nsassignment" {
		name     = "tf_nsassignment"
		variable = join("", ["$", citrixadc_nsvariable.tf_nsvariable.name])
		set      = 1
		comment  = "Testing"
	}
`
const testAccNsassignment_update = `
	resource "citrixadc_nsvariable" "tf_nsvariable" {
		name          = "tf_nsvariable"
		type          = "text(20)"
		scope         = "global"
		iffull        = "undef"
		ifvaluetoobig = "undef"
		ifnovalue     = "init"
		comment       = "Testing"
	}
	resource "citrixadc_nsassignment" "tf_nsassignment" {
		name     = "tf_nsassignment"
		variable = join("", ["$", citrixadc_nsvariable.tf_nsvariable.name])
		set      = 1
		comment  = "Testing_updated"
	}
`

func TestAccNsassignment_basic(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckNsassignmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsassignment_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsassignmentExist("citrixadc_nsassignment.tf_nsassignment", nil),
					resource.TestCheckResourceAttr("citrixadc_nsassignment.tf_nsassignment", "name", "tf_nsassignment"),
					resource.TestCheckResourceAttr("citrixadc_nsassignment.tf_nsassignment", "set", "1"),
					resource.TestCheckResourceAttr("citrixadc_nsassignment.tf_nsassignment", "comment", "Testing"),
				),
			},
			{
				Config: testAccNsassignment_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsassignmentExist("citrixadc_nsassignment.tf_nsassignment", nil),
					resource.TestCheckResourceAttr("citrixadc_nsassignment.tf_nsassignment", "name", "tf_nsassignment"),
					resource.TestCheckResourceAttr("citrixadc_nsassignment.tf_nsassignment", "set", "1"),
					resource.TestCheckResourceAttr("citrixadc_nsassignment.tf_nsassignment", "comment", "Testing_updated"),
				),
			},
		},
	})
}

func testAccCheckNsassignmentExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsassignment name is set")
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
		data, err := client.FindResource(service.Nsassignment.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nsassignment %s not found", n)
		}

		return nil
	}
}

func testAccCheckNsassignmentDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nsassignment" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Nsassignment.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nsassignment %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
