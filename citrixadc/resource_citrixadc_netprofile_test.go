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
	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

func TestAccNetprofile_basic(t *testing.T) {
	if isCpxRun {
		t.Skip("CPX 12.0 is outdated for this resource")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetprofileDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetprofile_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetprofileExist("citrixadc_netprofile.tf_netprofile", nil),
				),
			},
			resource.TestStep{
				Config: testAccNetprofile_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetprofileExist("citrixadc_netprofile.tf_netprofile", nil),
				),
			},
			resource.TestStep{
				Config: testAccNetprofile_basic_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetprofileExist("citrixadc_netprofile.tf_netprofile", nil),
				),
			},
		},
	})
}

const testAccNetprofile_proxyprotocolaftertlshandshake = `
	resource "citrixadc_netprofile" "tf_netprofile_proxyprotocolaftertlshandshake" {
		name = "tf_netprofile2"
		proxyprotocol = "ENABLED"
		proxyprotocoltxversion = "V2"
		proxyprotocolaftertlshandshake = "ENABLED"
	}
`

func TestAccNetprofile_proxyprotocolaftertlshandshake(t *testing.T) {
	if isCpxRun {
		t.Skip("CPX 12.0 is outdated for this resource")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetprofileDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNetprofile_proxyprotocolaftertlshandshake,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetprofileExist("citrixadc_netprofile.tf_netprofile_proxyprotocolaftertlshandshake", nil),
					resource.TestCheckResourceAttr("citrixadc_netprofile.tf_netprofile_proxyprotocolaftertlshandshake", "proxyprotocolaftertlshandshake", "ENABLED"),
				),
			},
		},
	})
}
func testAccCheckNetprofileExist(n string, id *string) resource.TestCheckFunc {
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

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(netscaler.Netprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckNetprofileDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_netprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(netscaler.Netprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccNetprofile_basic_step1 = `

resource "citrixadc_netprofile" "tf_netprofile" {
    name = "tf_netprofile"
    proxyprotocol = "ENABLED"
    proxyprotocoltxversion = "V1"
}

`

const testAccNetprofile_basic_step2 = `

resource "citrixadc_netprofile" "tf_netprofile" {
    name = "tf_netprofile"
    proxyprotocol = "ENABLED"
    proxyprotocoltxversion = "V2"
}

`

const testAccNetprofile_basic_step3 = `

resource "citrixadc_netprofile" "tf_netprofile" {
    name = "tf_netprofile2"
    proxyprotocol = "ENABLED"
    proxyprotocoltxversion = "V2"
}

`
