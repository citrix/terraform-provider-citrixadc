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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

const testAccIpsecprofile_basic = `


	resource "citrixadc_ipsecprofile" "tf_ipsecprofile" {
		name                  = "my_ipsecprofile"
		ikeversion            = "V2"
		encalgo               = ["AES", "3DES"]
		hashalgo              = ["HMAC_SHA1", "HMAC_SHA256"]
		livenesscheckinterval = 50
		psk                   = "GCC5VcY0TQ+0TfjGwCrR+cQthm5UnBPB"
	}
  
`
const testAccIpsecprofile_update = `


	resource "citrixadc_ipsecprofile" "tf_ipsecprofile" {
		name                  = "my_ipsecprofile"
		ikeversion            = "V1"
		encalgo               = ["AES", "3DES"]
		hashalgo              = ["HMAC_SHA1", "HMAC_SHA256"]
		livenesscheckinterval = 60
		psk                   = "XzuIoPthvSrBl/pD9OS+eiGyTJ6y5wuf"

	}
  
`

func TestAccIpsecprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIpsecprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIpsecprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpsecprofileExist("citrixadc_ipsecprofile.tf_ipsecprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_ipsecprofile.tf_ipsecprofile", "name", "my_ipsecprofile"),
					resource.TestCheckResourceAttr("citrixadc_ipsecprofile.tf_ipsecprofile", "ikeversion", "V2"),
					resource.TestCheckResourceAttr("citrixadc_ipsecprofile.tf_ipsecprofile", "livenesscheckinterval", "50"),
					resource.TestCheckResourceAttr("citrixadc_ipsecprofile.tf_ipsecprofile", "psk", "GCC5VcY0TQ+0TfjGwCrR+cQthm5UnBPB"),
				),
			},
			{
				Config: testAccIpsecprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpsecprofileExist("citrixadc_ipsecprofile.tf_ipsecprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_ipsecprofile.tf_ipsecprofile", "name", "my_ipsecprofile"),
					resource.TestCheckResourceAttr("citrixadc_ipsecprofile.tf_ipsecprofile", "ikeversion", "V1"),
					resource.TestCheckResourceAttr("citrixadc_ipsecprofile.tf_ipsecprofile", "livenesscheckinterval", "60"),
					resource.TestCheckResourceAttr("citrixadc_ipsecprofile.tf_ipsecprofile", "psk", "XzuIoPthvSrBl/pD9OS+eiGyTJ6y5wuf"),
				),
			},
		},
	})
}

func testAccCheckIpsecprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ipsecprofile name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Ipsecprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("ipsecprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckIpsecprofileDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_ipsecprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Ipsecprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("ipsecprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
