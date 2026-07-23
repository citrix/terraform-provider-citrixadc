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

// Step 1: all unset-eligible attributes set to non-default (but valid) values.
const testAccQuicprofile_unset_step1 = `
resource "citrixadc_quicprofile" "tf_unset" {
  name                           = "tf_test_quicprofile_unset"
  ackdelayexponent               = 5
  activeconnectionidlimit        = 4
  activeconnectionmigration      = "DISABLED"
  congestionctrlalgorithm        = "CUBIC"
  initialmaxdata                 = 2097152
  initialmaxstreamdatabidilocal  = 524288
  initialmaxstreamdatabidiremote = 524288
  initialmaxstreamdatauni        = 524288
  initialmaxstreamsbidi          = 200
  initialmaxstreamsuni           = 20
  maxackdelay                    = 25
  maxidletimeout                 = 200
  maxudpdatagramsperburst        = 10
  maxudppayloadsize              = 1500
  newtokenvalidityperiod         = 600
  retrytokenvalidityperiod       = 20
  statelessaddressvalidation     = "DISABLED"
}
`

// Step 2: eligible attributes removed from config -> provider must unset them,
// reverting each to its ADC default.
const testAccQuicprofile_unset_step2 = `
resource "citrixadc_quicprofile" "tf_unset" {
  name = "tf_test_quicprofile_unset"
}
`

func TestAccQuicprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckQuicprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values apply and persist.
				Config: testAccQuicprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckQuicprofileExist("citrixadc_quicprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "ackdelayexponent", "5"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "activeconnectionidlimit", "4"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "activeconnectionmigration", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "congestionctrlalgorithm", "CUBIC"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "initialmaxdata", "2097152"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "initialmaxstreamdatabidilocal", "524288"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "initialmaxstreamdatabidiremote", "524288"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "initialmaxstreamdatauni", "524288"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "initialmaxstreamsbidi", "200"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "initialmaxstreamsuni", "20"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "maxackdelay", "25"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "maxidletimeout", "200"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "maxudpdatagramsperburst", "10"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "maxudppayloadsize", "1500"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "newtokenvalidityperiod", "600"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "retrytokenvalidityperiod", "20"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "statelessaddressvalidation", "DISABLED"),
				),
			},
			{
				// Removing them must unset -> state reverts to NITRO defaults,
				// and the implicit post-apply plan must be empty.
				Config: testAccQuicprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckQuicprofileExist("citrixadc_quicprofile.tf_unset", nil),
					// State reverted to ADC defaults.
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "ackdelayexponent", "3"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "activeconnectionidlimit", "3"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "activeconnectionmigration", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "congestionctrlalgorithm", "Default"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "initialmaxdata", "1048576"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "initialmaxstreamdatabidilocal", "262144"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "initialmaxstreamdatabidiremote", "262144"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "initialmaxstreamdatauni", "262144"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "initialmaxstreamsbidi", "100"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "initialmaxstreamsuni", "10"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "maxackdelay", "20"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "maxidletimeout", "180"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "maxudpdatagramsperburst", "8"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "maxudppayloadsize", "1472"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "newtokenvalidityperiod", "300"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "retrytokenvalidityperiod", "10"),
					resource.TestCheckResourceAttr("citrixadc_quicprofile.tf_unset", "statelessaddressvalidation", "ENABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckQuicprofileADCValue("tf_test_quicprofile_unset", "ackdelayexponent", "3"),
					testAccCheckQuicprofileADCValue("tf_test_quicprofile_unset", "activeconnectionidlimit", "3"),
					testAccCheckQuicprofileADCValue("tf_test_quicprofile_unset", "activeconnectionmigration", "ENABLED"),
					testAccCheckQuicprofileADCValue("tf_test_quicprofile_unset", "congestionctrlalgorithm", "Default"),
					testAccCheckQuicprofileADCValue("tf_test_quicprofile_unset", "initialmaxdata", "1048576"),
					testAccCheckQuicprofileADCValue("tf_test_quicprofile_unset", "initialmaxstreamdatabidilocal", "262144"),
					testAccCheckQuicprofileADCValue("tf_test_quicprofile_unset", "initialmaxstreamdatabidiremote", "262144"),
					testAccCheckQuicprofileADCValue("tf_test_quicprofile_unset", "initialmaxstreamdatauni", "262144"),
					testAccCheckQuicprofileADCValue("tf_test_quicprofile_unset", "initialmaxstreamsbidi", "100"),
					testAccCheckQuicprofileADCValue("tf_test_quicprofile_unset", "initialmaxstreamsuni", "10"),
					testAccCheckQuicprofileADCValue("tf_test_quicprofile_unset", "maxackdelay", "20"),
					testAccCheckQuicprofileADCValue("tf_test_quicprofile_unset", "maxidletimeout", "180"),
					testAccCheckQuicprofileADCValue("tf_test_quicprofile_unset", "maxudpdatagramsperburst", "8"),
					testAccCheckQuicprofileADCValue("tf_test_quicprofile_unset", "maxudppayloadsize", "1472"),
					testAccCheckQuicprofileADCValue("tf_test_quicprofile_unset", "newtokenvalidityperiod", "300"),
					testAccCheckQuicprofileADCValue("tf_test_quicprofile_unset", "retrytokenvalidityperiod", "10"),
					testAccCheckQuicprofileADCValue("tf_test_quicprofile_unset", "statelessaddressvalidation", "ENABLED"),
				),
			},
		},
	})
}

// testAccCheckQuicprofileADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckQuicprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Quicprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("quicprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("quicprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}
