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

const testAccAppqoecustomresp_basic = `

	resource "citrixadc_appqoecustomresp" "tf_appqoecustomresp" {
		name   = "my_appqoecustomresp"
		src   = "local://index.html"
	}
`

func TestAccAppqoecustomresp_basic(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckAppqoecustomrespDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppqoecustomresp_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppqoecustomrespExist("citrixadc_appqoecustomresp.tf_appqoecustomresp", nil),
				),
			},
		},
	})
}

func testAccCheckAppqoecustomrespExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appqoecustomresp name is set")
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
		dataArr, err := client.FindAllResources(service.Appqoecustomresp.Type())
		if err != nil {
			return err
		}
		if len(dataArr) == 0 {
			return fmt.Errorf("appqoecustomresp %s not found", n)
		}

		found := false
		for _, v := range dataArr {
			if v["name"].(string) == rs.Primary.ID {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("appqoecustomresp %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppqoecustomrespDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appqoecustomresp" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Appqoecustomresp.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appqoecustomresp %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
