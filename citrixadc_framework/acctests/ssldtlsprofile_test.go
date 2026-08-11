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
	initialretrytimeout = 2
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
		initialretrytimeout = 3
	}
`

const testAccSsldtlsprofileDataSource_basic = `
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
	initialretrytimeout = 2
}

data "citrixadc_ssldtlsprofile" "tf_ssldtlsprofile" {
	name = citrixadc_ssldtlsprofile.tf_ssldtlsprofile.name
}
`

func TestAccSsldtlsprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSsldtlsprofileDestroy,
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
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_ssldtlsprofile", "initialretrytimeout", "2"),
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
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_ssldtlsprofile", "initialretrytimeout", "3"),
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

		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Ssldtlsprofile.Type(), rs.Primary.ID)

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
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_ssldtlsprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Ssldtlsprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("ssldtlsprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccSsldtlsprofile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_ssldtlsprofile.tf_ssldtlsprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSsldtlsprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSsldtlsprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSsldtlsprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Ssldtlsprofile.Type(), "tf_ssldtlsprofile"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccSsldtlsprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSsldtlsprofileExist(resAddr, nil)),
			},
		},
	})
}

func TestAccSsldtlsprofile_import(t *testing.T) {
	const resAddr = "citrixadc_ssldtlsprofile.tf_ssldtlsprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSsldtlsprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccSsldtlsprofile_basic},
			{
				Config:                  testAccSsldtlsprofile_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccSsldtlsprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckSsldtlsprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccSsldtlsprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSsldtlsprofileExist("citrixadc_ssldtlsprofile.tf_ssldtlsprofile", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccSsldtlsprofile_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckSsldtlsprofileExist("citrixadc_ssldtlsprofile.tf_ssldtlsprofile", nil)),
			},
		},
	})
}

// The ssldtlsprofile unset test sets every unset-eligible read/write attribute
// to a valid non-default value, then removes them all from config. The provider
// must issue a NITRO unset so each attribute reverts to its documented default.
const testAccSsldtlsprofile_unset_step1 = `
resource "citrixadc_ssldtlsprofile" "tf_unset" {
	name                 = "tf_ssldtlsprofile_unset"
	helloverifyrequest   = "DISABLED"
	initialretrytimeout  = 2
	maxbadmacignorecount = 128
	maxholdqlen          = 64
	maxpacketsize        = 125
	maxrecordsize        = 250
	maxretrytime         = 6
	pmtudiscovery        = "ENABLED"
	terminatesession     = "ENABLED"
}
`

const testAccSsldtlsprofile_unset_step2 = `
resource "citrixadc_ssldtlsprofile" "tf_unset" {
	name = "tf_ssldtlsprofile_unset"
	# All unset-eligible attributes removed from config -> the provider must
	# unset them (revert to the documented NITRO defaults).
}
`

func TestAccSsldtlsprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSsldtlsprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccSsldtlsprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSsldtlsprofileExist("citrixadc_ssldtlsprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_unset", "helloverifyrequest", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_unset", "initialretrytimeout", "2"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_unset", "maxbadmacignorecount", "128"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_unset", "maxholdqlen", "64"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_unset", "maxpacketsize", "125"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_unset", "maxrecordsize", "250"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_unset", "maxretrytime", "6"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_unset", "pmtudiscovery", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_unset", "terminatesession", "ENABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccSsldtlsprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSsldtlsprofileExist("citrixadc_ssldtlsprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_unset", "helloverifyrequest", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_unset", "initialretrytimeout", "3"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_unset", "maxbadmacignorecount", "100"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_unset", "maxholdqlen", "32"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_unset", "maxpacketsize", "120"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_unset", "maxrecordsize", "1459"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_unset", "maxretrytime", "3"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_unset", "pmtudiscovery", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_ssldtlsprofile.tf_unset", "terminatesession", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckSsldtlsprofileADCValue("tf_ssldtlsprofile_unset", "helloverifyrequest", "ENABLED"),
					testAccCheckSsldtlsprofileADCValue("tf_ssldtlsprofile_unset", "pmtudiscovery", "DISABLED"),
					testAccCheckSsldtlsprofileADCValue("tf_ssldtlsprofile_unset", "maxrecordsize", "1459"),
				),
			},
		},
	})
}

// testAccCheckSsldtlsprofileADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckSsldtlsprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Ssldtlsprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("ssldtlsprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("ssldtlsprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccSsldtlsprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSsldtlsprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_ssldtlsprofile.tf_ssldtlsprofile", "name", "tf_ssldtlsprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_ssldtlsprofile.tf_ssldtlsprofile", "helloverifyrequest", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_ssldtlsprofile.tf_ssldtlsprofile", "maxbadmacignorecount", "128"),
					resource.TestCheckResourceAttr("data.citrixadc_ssldtlsprofile.tf_ssldtlsprofile", "maxholdqlen", "64"),
					resource.TestCheckResourceAttr("data.citrixadc_ssldtlsprofile.tf_ssldtlsprofile", "maxpacketsize", "125"),
					resource.TestCheckResourceAttr("data.citrixadc_ssldtlsprofile.tf_ssldtlsprofile", "maxrecordsize", "250"),
					resource.TestCheckResourceAttr("data.citrixadc_ssldtlsprofile.tf_ssldtlsprofile", "maxretrytime", "5"),
					resource.TestCheckResourceAttr("data.citrixadc_ssldtlsprofile.tf_ssldtlsprofile", "pmtudiscovery", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_ssldtlsprofile.tf_ssldtlsprofile", "terminatesession", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_ssldtlsprofile.tf_ssldtlsprofile", "initialretrytimeout", "2"),
				),
			},
		},
	})
}
