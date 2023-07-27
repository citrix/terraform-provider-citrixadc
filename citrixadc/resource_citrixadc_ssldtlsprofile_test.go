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

const testAccSsldtlsprofile_basic = `
resource "citrixadc_ssldtlsprofile" "tf_ssldtlsprofile" {
	name = "tf_ssldtlsprofile"
	helloverifyrequest = "ENABLED"
	maxbadmacignorecount = 128
	maxholdqlen = 64
	maxpacketsize = 125
	maxrecordsize = 250
	maxretrytime = 5
	pmtudiscovery = "DISABLED"
	terminatesession = "ENABLED"
}
`

const testAccSsldtlsprofile_basic_update = `
	resource "citrixadc_ssldtlsprofile" "tf_ssldtlsprofile" {
		name = "tf_ssldtlsprofile"
		helloverifyrequest = "DISABLED"
		maxbadmacignorecount = 129
		maxholdqlen = 65
		maxpacketsize = 126
		maxrecordsize = 251
		maxretrytime = 6
		pmtudiscovery = "ENABLED"
		terminatesession = "DISABLED"
	}
`

func TestAccSsldtlsprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSsldtlsprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSsldtlsprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSsldtlsprofileExist("citrixadc_ssldtlsprofile.tf_ssldtlsprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_ssldtlsprofile", "name", "tf_ssldtlsprofile"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_ssldtlsprofile", "helloverifyrequest", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_ssldtlsprofile", "maxbadmacignorecount", "128"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_ssldtlsprofile", "maxholdqlen", "64"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_ssldtlsprofile", "maxpacketsize", "125"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_ssldtlsprofile", "maxrecordsize", "250"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_ssldtlsprofile", "maxretrytime", "5"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_ssldtlsprofile", "pmtudiscovery", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_ssldtlsprofile", "terminatesession", "ENABLED"),
				),
			},
			{
				Config: testAccSsldtlsprofile_basic_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSsldtlsprofileExist("citrixadc_ssldtlsprofile.tf_ssldtlsprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_ssldtlsprofile", "name", "tf_ssldtlsprofile"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_ssldtlsprofile", "helloverifyrequest", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_ssldtlsprofile", "maxbadmacignorecount", "129"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_ssldtlsprofile", "maxholdqlen", "65"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_ssldtlsprofile", "maxpacketsize", "126"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_ssldtlsprofile", "maxrecordsize", "251"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_ssldtlsprofile", "maxretrytime", "6"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_ssldtlsprofile", "pmtudiscovery", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_ssldtlsprofile", "terminatesession", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckSsldtlsprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ssldtlsprofile name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Ssldtlsprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("ssldtlsprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckSsldtlsprofileDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_ssldtlsprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Ssldtlsprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("ssldtlsprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
