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

const testAccNsmode_basic_step1 = `
resource "citrixadc_nsmode" "tf_nsmode" {
    usip = true
	cka = true
}
`

const testAccNsmode_basic_step2 = `
resource "citrixadc_nsmode" "tf_nsmode" {
    usip = false
	cka = false
}
`

func TestAccNsmode_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNsmode_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testEnsureNsmodes([]string{"usip", "cka"}, true),
				),
			},
			{
				Config: testAccNsmode_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testEnsureNsmodes([]string{"usip", "cka"}, false),
				),
			},
		},
	})
}

func testEnsureNsmodes(modes []string, expectedState bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		findParams := service.FindParams{
			ResourceType: "nsmode",
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}
		if len(dataArr) != 1 {
			return fmt.Errorf("Unexpected fetched nsmode result %v", dataArr)
		}
		data := dataArr[0]
		for _, mode := range modes {
			if val, ok := data[mode]; ok {
				if val.(bool) != expectedState {
					return fmt.Errorf("Wrong mode value for %s. Expected %v, found %v", mode, expectedState, val.(bool))
				}
			} else {
				return fmt.Errorf("Cannot find mode %s in retrieved modes list", mode)
			}
		}
		return nil
	}
}
