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

const testAccQuicprofile_basic_step1 = `
resource "citrixadc_quicprofile" "tf_quicprofile" {
  name                       = "tf_quicprofile"
  maxidletimeout             = 180
  congestionctrlalgorithm    = "CUBIC"
  activeconnectionmigration  = "ENABLED"
  statelessaddressvalidation = "ENABLED"
  ackdelayexponent           = 3
  activeconnectionidlimit    = 3
  maxackdelay                = 20
  maxudppayloadsize          = 1472
}

`

const testAccQuicprofile_basic_step2 = `
resource "citrixadc_quicprofile" "tf_quicprofile" {
  name                       = "tf_quicprofile"
  maxidletimeout             = 200
  congestionctrlalgorithm    = "BBR"
  activeconnectionmigration  = "DISABLED"
  statelessaddressvalidation = "DISABLED"
  ackdelayexponent           = 5
  activeconnectionidlimit    = 4
  maxackdelay                = 25
  maxudppayloadsize          = 1500
}

`

func TestAccQuicprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckQuicprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccQuicprofile_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckQuicprofileExist("citrixadc_quicprofile.tf_quicprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_quicprofile", "name", "tf_quicprofile"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_quicprofile", "maxidletimeout", "180"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_quicprofile", "congestionctrlalgorithm", "CUBIC"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_quicprofile", "activeconnectionmigration", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_quicprofile", "statelessaddressvalidation", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_quicprofile", "ackdelayexponent", "3"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_quicprofile", "activeconnectionidlimit", "3"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_quicprofile", "maxackdelay", "20"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_quicprofile", "maxudppayloadsize", "1472"),
				),
			},
			{
				Config: testAccQuicprofile_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckQuicprofileExist("citrixadc_quicprofile.tf_quicprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_quicprofile", "name", "tf_quicprofile"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_quicprofile", "maxidletimeout", "200"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_quicprofile", "congestionctrlalgorithm", "BBR"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_quicprofile", "activeconnectionmigration", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_quicprofile", "statelessaddressvalidation", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_quicprofile", "ackdelayexponent", "5"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_quicprofile", "activeconnectionidlimit", "4"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_quicprofile", "maxackdelay", "25"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_quicprofile", "maxudppayloadsize", "1500"),
				),
			},
		},
	})
}

func TestAccQuicprofile_import(t *testing.T) {
	const resAddr = "citrixadc_quicprofile.tf_quicprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckQuicprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccQuicprofile_basic_step1,
			},
			{
				Config:                  testAccQuicprofile_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckQuicprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No quicprofile name is set")
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
		data, err := client.FindResource(service.Quicprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("quicprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckQuicprofileDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_quicprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Quicprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("quicprofile %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

const testAccQuicprofileDataSource_basic = `

resource "citrixadc_quicprofile" "tf_quicprofile" {
  name                       = "tf_quicprofile"
  maxidletimeout             = 180
  congestionctrlalgorithm    = "CUBIC"
  activeconnectionmigration  = "ENABLED"
  statelessaddressvalidation = "ENABLED"
  ackdelayexponent           = 3
  activeconnectionidlimit    = 3
  maxackdelay                = 20
  maxudppayloadsize          = 1472
}

data "citrixadc_quicprofile" "tf_quicprofile" {
  name       = citrixadc_quicprofile.tf_quicprofile.name
  depends_on = [citrixadc_quicprofile.tf_quicprofile]
}
`

func TestAccQuicprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckQuicprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccQuicprofileDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_quicprofile.tf_quicprofile", "name", "tf_quicprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_quicprofile.tf_quicprofile", "maxidletimeout", "180"),
					resource.TestCheckResourceAttr("data.citrixadc_quicprofile.tf_quicprofile", "congestionctrlalgorithm", "CUBIC"),
					resource.TestCheckResourceAttr("data.citrixadc_quicprofile.tf_quicprofile", "activeconnectionmigration", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_quicprofile.tf_quicprofile", "statelessaddressvalidation", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_quicprofile.tf_quicprofile", "ackdelayexponent", "3"),
					resource.TestCheckResourceAttr("data.citrixadc_quicprofile.tf_quicprofile", "activeconnectionidlimit", "3"),
					resource.TestCheckResourceAttr("data.citrixadc_quicprofile.tf_quicprofile", "maxackdelay", "20"),
					resource.TestCheckResourceAttr("data.citrixadc_quicprofile.tf_quicprofile", "maxudppayloadsize", "1472"),
				),
			},
		},
	})
}
