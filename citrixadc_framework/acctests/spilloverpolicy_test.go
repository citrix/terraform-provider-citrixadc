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

const testAccSpilloverpolicy_add = `
	resource "citrixadc_spilloveraction" "tf_spilloveraction" {
		name   = "my_spilloveraction"
		action = "SPILLOVER"
	}
	resource "citrixadc_spilloverpolicy" "tf_spilloverpolicy" {
		name    = "tf_spilloverpolicy"
		rule    = "true"
		action  = citrixadc_spilloveraction.tf_spilloveraction.name
		comment = "This is example of spilloverpolicy"
	}
`
const testAccSpilloverpolicy_update = `
	resource "citrixadc_spilloveraction" "tf_spilloveraction" {
		name   = "my_spilloveraction"
		action = "SPILLOVER"
	}
	resource "citrixadc_spilloverpolicy" "tf_spilloverpolicy" {
		name    = "tf_spilloverpolicy"
		rule    = "false"
		action  = citrixadc_spilloveraction.tf_spilloveraction.name
		comment = "This is example of spilloverpolicy"
	}
`

const testAccSpilloverpolicyDataSource_basic = `
resource "citrixadc_spilloveraction" "tf_spilloveraction" {
	name   = "my_spilloveraction_ds"
	action = "SPILLOVER"
}
resource "citrixadc_spilloverpolicy" "tf_spilloverpolicy" {
	name    = "tf_spilloverpolicy_ds"
	rule    = "true"
	action  = citrixadc_spilloveraction.tf_spilloveraction.name
	comment = "This is example of spilloverpolicy"
}

data "citrixadc_spilloverpolicy" "tf_spilloverpolicy" {
	name = citrixadc_spilloverpolicy.tf_spilloverpolicy.name
}
`

func TestAccSpilloverpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSpilloverpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSpilloverpolicy_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpilloverpolicyExist("citrixadc_spilloverpolicy.tf_spilloverpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_spilloverpolicy.tf_spilloverpolicy", "name", "tf_spilloverpolicy"),
					resource.TestCheckResourceAttr("citrixadc_spilloverpolicy.tf_spilloverpolicy", "rule", "true"),
				),
			},
			{
				Config: testAccSpilloverpolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSpilloverpolicyExist("citrixadc_spilloverpolicy.tf_spilloverpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_spilloverpolicy.tf_spilloverpolicy", "name", "tf_spilloverpolicy"),
					resource.TestCheckResourceAttr("citrixadc_spilloverpolicy.tf_spilloverpolicy", "rule", "false"),
				),
			},
		},
	})
}

func TestAccSpilloverpolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSpilloverpolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_spilloverpolicy.tf_spilloverpolicy", "name", "tf_spilloverpolicy_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_spilloverpolicy.tf_spilloverpolicy", "rule", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_spilloverpolicy.tf_spilloverpolicy", "comment", "This is example of spilloverpolicy"),
				),
			},
		},
	})
}

func testAccCheckSpilloverpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No spilloverpolicy name is set")
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
		data, err := client.FindResource(service.Spilloverpolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("spilloverpolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckSpilloverpolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_spilloverpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Spilloverpolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("spilloverpolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
