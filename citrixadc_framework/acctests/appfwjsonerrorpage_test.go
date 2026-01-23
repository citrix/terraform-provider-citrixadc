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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

const testAccAppfwjsonerrorpage_basic = `
	resource "citrixadc_systemfile" "tf_jsonerrorpage" {
		filename     = "appfwjsonerrorpage.json"
		filelocation = "/var/tmp"
		filecontent  = file("testdata/appfwjsonerrorpage.json")
	}
	resource "citrixadc_appfwjsonerrorpage" "tf_appfwjsonerrorpage" {
		name       = "tf_appfwjsonerrorpage"
		src        = "local://appfwjsonerrorpage.json"
		depends_on = [citrixadc_systemfile.tf_jsonerrorpage]
		comment    = "TestingExample"
	}
`

func TestAccAppfwjsonerrorpage_basic(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwjsonerrorpageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwjsonerrorpage_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwjsonerrorpageExist("citrixadc_appfwjsonerrorpage.tf_appfwjsonerrorpage", nil),
				),
			},
		},
	})
}

func testAccCheckAppfwjsonerrorpageExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwjsonerrorpage name is set")
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
		data, err := client.FindResource("appfwjsonerrorpage", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appfwjsonerrorpage %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwjsonerrorpageDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwjsonerrorpage" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("appfwjsonerrorpage", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwjsonerrorpage %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
