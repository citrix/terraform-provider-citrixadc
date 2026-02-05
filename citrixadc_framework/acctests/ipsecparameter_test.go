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

const testAccIpsecparameter_basic = `

resource "citrixadc_ipsecparameter" "tf_ipsecparameter" {
	ikeversion            = "V2"
	encalgo               = ["AES", "AES256"]
	hashalgo              = ["HMAC_SHA1", "HMAC_SHA256"]
	livenesscheckinterval = 50
	}
  
`
const testAccIpsecparameter_update = `

resource "citrixadc_ipsecparameter" "tf_ipsecparameter" {
	ikeversion            = "V1"
	encalgo               = ["AES", "AES256"]
	hashalgo              = ["HMAC_SHA1", "HMAC_SHA256"]
	livenesscheckinterval = 60
	}
  
`

const testAccIpsecparameterDataSource_basic = `

resource "citrixadc_ipsecparameter" "tf_ipsecparameter" {
	ikeversion            = "V2"
	encalgo               = ["AES", "AES256"]
	hashalgo              = ["HMAC_SHA1", "HMAC_SHA256"]
	livenesscheckinterval = 50
}

data "citrixadc_ipsecparameter" "tf_ipsecparameter_datasource" {
	depends_on = [citrixadc_ipsecparameter.tf_ipsecparameter]
}
`

func TestAccIpsecparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccIpsecparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpsecparameterExist("citrixadc_ipsecparameter.tf_ipsecparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_ipsecparameter.tf_ipsecparameter", "ikeversion", "V2"),
					resource.TestCheckResourceAttr("citrixadc_ipsecparameter.tf_ipsecparameter", "livenesscheckinterval", "50"),
				),
			},
			{
				Config: testAccIpsecparameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpsecparameterExist("citrixadc_ipsecparameter.tf_ipsecparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_ipsecparameter.tf_ipsecparameter", "ikeversion", "V1"),
					resource.TestCheckResourceAttr("citrixadc_ipsecparameter.tf_ipsecparameter", "livenesscheckinterval", "60"),
				),
			},
		},
	})
}

func testAccCheckIpsecparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ipsecparameter name is set")
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
		data, err := client.FindResource(service.Ipsecparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("ipsecparameter %s not found", n)
		}

		return nil
	}
}

func TestAccIpsecparameterDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIpsecparameterDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_ipsecparameter.tf_ipsecparameter_datasource", "ikeversion", "V2"),
					resource.TestCheckResourceAttr("data.citrixadc_ipsecparameter.tf_ipsecparameter_datasource", "livenesscheckinterval", "50"),
				),
			},
		},
	})
}
