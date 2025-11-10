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
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccSslprofile_sslcipher_binding_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckSslprofile_sslcipher_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslprofile_sslcipher_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofile_sslcipher_bindingExist("citrixadc_sslprofile_sslcipher_binding.tf_binding", nil),
					testAccCheckSslprofile_sslcipher_bindingExist("citrixadc_sslprofile_sslcipher_binding.tf_binding2", nil),
				),
			},
			{
				Config: testAccSslprofile_sslcipher_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofile_sslcipher_bindingExist("citrixadc_sslprofile_sslcipher_binding.tf_binding", nil),
					testAccCheckSslprofile_sslcipher_bindingExist("citrixadc_sslprofile_sslcipher_binding.tf_binding2", nil),
				),
			},
			{
				Config: testAccSslprofile_sslcipher_binding_basic_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofile_sslcipher_bindingExist("citrixadc_sslprofile_sslcipher_binding.tf_binding", nil),
				),
			},
		},
	})
}

func testAccCheckSslprofile_sslcipher_bindingExist(n string, id *string) resource.TestCheckFunc {
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

		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		bindingId := rs.Primary.ID
		idSlice := strings.Split(bindingId, ",")

		profileName := idSlice[0]
		cipherName := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "sslprofile_sslcipher_binding",
			ResourceName:             profileName,
			ResourceMissingErrorCode: 3248,
		}

		dataArr, err := client.FindResourceArrayWithParams(findParams)

		if err != nil {
			return err
		}

		if len(dataArr) == 0 {
			return fmt.Errorf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		}

		foundIndex := -1
		for i, v := range dataArr {
			if v["cipheraliasname"].(string) == cipherName {
				foundIndex = i
				break
			}
		}

		if foundIndex == -1 {
			return fmt.Errorf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams cipher name not found in array")
		}

		return nil
	}
}

func testAccCheckSslprofile_sslcipher_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslprofile_sslcipher_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		idSlice := strings.Split(rs.Primary.ID, ",")

		profileName := idSlice[0]

		findParams := service.FindParams{
			ResourceType:             "sslprofile_sslcipher_binding",
			ResourceName:             profileName,
			ResourceMissingErrorCode: 3248,
		}

		dataArr, err := client.FindResourceArrayWithParams(findParams)

		if err != nil {
			if strings.Contains(err.Error(), "\"errorcode\": 3248") {
				// Expected since ssl profile is already deleted
				return nil
			} else {
				// Unexpected error
				return err
			}
		}

		if len(dataArr) > 0 {
			return fmt.Errorf("sslcipher binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccSslprofile_sslcipher_binding_basic_step1 = `

resource "citrixadc_sslprofile" "tf_sslprofile" {
  name = "tf_sslprofile"

  ecccurvebindings = []

}

resource "citrixadc_sslprofile_sslcipher_binding" "tf_binding" {
    name = citrixadc_sslprofile.tf_sslprofile.name
    ciphername = "HIGH"
    cipherpriority = 10
}

resource "citrixadc_sslprofile_sslcipher_binding" "tf_binding2" {
    name = citrixadc_sslprofile.tf_sslprofile.name
    ciphername = "LOW"
    cipherpriority = 20
}

`

const testAccSslprofile_sslcipher_binding_basic_step2 = `

resource "citrixadc_sslprofile" "tf_sslprofile" {
  name = "tf_sslprofile"

  ecccurvebindings = []

}

resource "citrixadc_sslprofile_sslcipher_binding" "tf_binding" {
    name = citrixadc_sslprofile.tf_sslprofile.name
    ciphername = "HIGH"
    cipherpriority = 10
}

resource "citrixadc_sslprofile_sslcipher_binding" "tf_binding2" {
    name = citrixadc_sslprofile.tf_sslprofile.name
    ciphername = "LOW"
    cipherpriority = 30
}

`

const testAccSslprofile_sslcipher_binding_basic_step3 = `

resource "citrixadc_sslprofile" "tf_sslprofile" {
  name = "tf_sslprofile"

  ecccurvebindings = []

}

resource "citrixadc_sslprofile_sslcipher_binding" "tf_binding" {
    name = citrixadc_sslprofile.tf_sslprofile.name
    ciphername = "HIGH"
    cipherpriority = 10
}
`
