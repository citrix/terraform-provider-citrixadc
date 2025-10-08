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

const testAccBotsignature_basic = `
	resource "citrixadc_systemfile" "tf_signature" {
		filename     = "bot_signature.json"
		filelocation = "/var/tmp"
		filecontent  = file("testdata/bot_signatures.json")
	}
	resource "citrixadc_botsignature" "tf_botsignature" {
		name       = "tf_botsignature"
		src        = "local://bot_signature.json"
		depends_on = [citrixadc_systemfile.tf_signature]
		comment    = "TestingExample"
	}
`

func TestAccBotsignature_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckBotsignatureDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBotsignature_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBotsignatureExist("citrixadc_botsignature.tf_botsignature", nil),
				),
			},
		},
	})
}

func testAccCheckBotsignatureExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No botsignature name is set")
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
		data, err := client.FindResource("botsignature", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("botsignature %s not found", n)
		}

		return nil
	}
}

func testAccCheckBotsignatureDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_botsignature" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("botsignature", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("botsignature %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
