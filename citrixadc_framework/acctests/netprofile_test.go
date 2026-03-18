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

func TestAccNetprofile_basic(t *testing.T) {
	if isCpxRun {
		t.Skip("CPX 12.0 is outdated for this resource")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetprofile_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetprofileExist("citrixadc_netprofile.tf_netprofile", nil),
				),
			},
			{
				Config: testAccNetprofile_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetprofileExist("citrixadc_netprofile.tf_netprofile", nil),
				),
			},
			{
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
		name                   = "tf_netprofile2"
		proxyprotocol          = "ENABLED"
		proxyprotocoltxversion = "V2"
	}
`

func TestAccNetprofile_proxyprotocolaftertlshandshake(t *testing.T) {
	if isCpxRun {
		t.Skip("CPX 12.0 is outdated for this resource")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetprofile_proxyprotocolaftertlshandshake,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetprofileExist("citrixadc_netprofile.tf_netprofile_proxyprotocolaftertlshandshake", nil),
					// resource.TestCheckResourceAttr("citrixadc_netprofile.tf_netprofile_proxyprotocolaftertlshandshake", "proxyprotocolaftertlshandshake", "ENABLED"),
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

		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Netprofile.Type(), rs.Primary.ID)

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
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_netprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Netprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccNetprofile_basic_step1 = `

resource "citrixadc_netprofile" "tf_netprofile" {
    name 				   = "tf_netprofile"
    proxyprotocol 		   = "ENABLED"
    proxyprotocoltxversion = "V1"
}

`

const testAccNetprofile_basic_step2 = `

resource "citrixadc_netprofile" "tf_netprofile" {
    name 				   = "tf_netprofile"
    proxyprotocol          = "ENABLED"
    proxyprotocoltxversion = "V2"
}

`

const testAccNetprofile_basic_step3 = `

resource "citrixadc_netprofile" "tf_netprofile" {
    name 				   = "tf_netprofile2"
    proxyprotocol 		   = "ENABLED"
    proxyprotocoltxversion = "V2"
}

`

const testAccNetprofileDataSource_basic = `

	resource "citrixadc_netprofile" "tf_netprofile" {
		name                   = "tf_netprofile_ds"
		proxyprotocol          = "ENABLED"
		proxyprotocoltxversion = "V1"
		srcippersistency       = "ENABLED"
	}

	data "citrixadc_netprofile" "tf_netprofile_ds" {
		name = citrixadc_netprofile.tf_netprofile.name
	}
`

func TestAccNetprofileDataSource_basic(t *testing.T) {
	if isCpxRun {
		t.Skip("CPX 12.0 is outdated for this resource")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNetprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_netprofile.tf_netprofile_ds", "name", "tf_netprofile_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_netprofile.tf_netprofile_ds", "proxyprotocol", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_netprofile.tf_netprofile_ds", "proxyprotocoltxversion", "V1"),
					resource.TestCheckResourceAttr("data.citrixadc_netprofile.tf_netprofile_ds", "srcippersistency", "ENABLED"),
				),
			},
		},
	})
}
