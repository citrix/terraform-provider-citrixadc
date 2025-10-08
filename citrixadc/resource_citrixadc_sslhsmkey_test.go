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

const testAccSslhsmkey_basic = `
resource "citrixadc_sslhsmkey" "tf_hsmkey1" {
	hsmkeyname = "hsmkey1"
	hsmtype = "Fillme"
	password = "Fillme"
	serialnum = "Fillme"
}
`

func TestAccSslhsmkey_basic(t *testing.T) {
	if adcTestbed != "STANDALONE_HSM" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_HSM.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckSslhsmkeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslhsmkey_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslhsmkeyExist("citrixadc_sslhsmkey.tf_hsmkey1", nil),
				),
			},
		},
	})
}

func testAccCheckSslhsmkeyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslhsmkey name is set")
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
		data, err := client.FindResource(service.Sslhsmkey.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("sslhsmkey %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslhsmkeyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslhsmkey" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Sslhsmkey.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslhsmkey %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
