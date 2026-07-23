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

const testAccRdpclientprofile_basic = `


resource "citrixadc_rdpclientprofile" "tf_rdpclientprofile" {
	name              = "my_rdpclientprofile"
	rdpurloverride    = "ENABLE"
	redirectclipboard = "ENABLE"
	redirectdrives    = "ENABLE"
	rdpvalidateclientip = "ENABLE"
	}
  
`
const testAccRdpclientprofile_update = `


resource "citrixadc_rdpclientprofile" "tf_rdpclientprofile" {
	name              = "my_rdpclientprofile"
	rdpurloverride    = "DISABLE"
	redirectclipboard = "DISABLE"
	redirectdrives    = "DISABLE"
	rdpvalidateclientip = "DISABLE"
	}
  
`

func TestAccRdpclientprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRdpclientprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRdpclientprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRdpclientprofileExist("citrixadc_rdpclientprofile.tf_rdpclientprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_rdpclientprofile", "name", "my_rdpclientprofile"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_rdpclientprofile", "rdpurloverride", "ENABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_rdpclientprofile", "redirectclipboard", "ENABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_rdpclientprofile", "redirectdrives", "ENABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_rdpclientprofile", "rdpvalidateclientip", "ENABLE"),
				),
			},
			{
				Config: testAccRdpclientprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRdpclientprofileExist("citrixadc_rdpclientprofile.tf_rdpclientprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_rdpclientprofile", "name", "my_rdpclientprofile"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_rdpclientprofile", "rdpurloverride", "DISABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_rdpclientprofile", "redirectclipboard", "DISABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_rdpclientprofile", "redirectdrives", "DISABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_rdpclientprofile", "rdpvalidateclientip", "DISABLE"),
				),
			},
		},
	})
}

func TestAccRdpclientprofile_import(t *testing.T) {
	const resAddr = "citrixadc_rdpclientprofile.tf_rdpclientprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRdpclientprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccRdpclientprofile_basic},
			{
				Config:            testAccRdpclientprofile_basic,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// psk_wo_version is a write-only version tracker (defaulted to 1) that
				// is not returned by NITRO, so it cannot round-trip through import.
				ImportStateVerifyIgnore: []string{"psk_wo_version"},
			},
		},
	})
}

func testAccCheckRdpclientprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No rdpclientprofile name is set")
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
		data, err := client.FindResource("rdpclientprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("rdpclientprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckRdpclientprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_rdpclientprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("rdpclientprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("rdpclientprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccRdpclientprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRdpclientprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_rdpclientprofile.test", "name", "tf_rdpclientprofile"),
					resource.TestCheckResourceAttrSet("data.citrixadc_rdpclientprofile.test", "id"),
					resource.TestCheckResourceAttrSet("data.citrixadc_rdpclientprofile.test", "addusernameinrdpfile"),
					resource.TestCheckResourceAttrSet("data.citrixadc_rdpclientprofile.test", "keyboardhook"),
				),
			},
		},
	})
}

const testAccRdpclientprofileDataSource_basic = `
resource "citrixadc_rdpclientprofile" "tf_rdpclientprofile" {
	name                  = "tf_rdpclientprofile"
	addusernameinrdpfile  = "YES"
	keyboardhook          = "InFullScreenMode"
}

data "citrixadc_rdpclientprofile" "test" {
	name = citrixadc_rdpclientprofile.tf_rdpclientprofile.name
}
`

// Test backward-compatible path: using psk (Sensitive attribute)
const testAccRdpclientprofile_psk_step1 = `
	variable "rdpclientprofile_psk" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_rdpclientprofile" "tf_rdpclientprofile_ephem" {
		name           = "tf_rdpclientprofile_ephem"
		rdpurloverride = "ENABLE"
		psk            = var.rdpclientprofile_psk
	}
`

const testAccRdpclientprofile_psk_step2 = `
	variable "rdpclientprofile_psk_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_rdpclientprofile" "tf_rdpclientprofile_ephem" {
		name           = "tf_rdpclientprofile_ephem"
		rdpurloverride = "ENABLE"
		psk            = var.rdpclientprofile_psk_2
	}
`

func TestAccRdpclientprofile_psk_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_rdpclientprofile_psk", "presharedkey1")
	t.Setenv("TF_VAR_rdpclientprofile_psk_2", "presharedkey2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRdpclientprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRdpclientprofile_psk_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRdpclientprofileExist("citrixadc_rdpclientprofile.tf_rdpclientprofile_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_rdpclientprofile_ephem", "name", "tf_rdpclientprofile_ephem"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_rdpclientprofile_ephem", "rdpurloverride", "ENABLE"),
				),
			},
			{
				Config: testAccRdpclientprofile_psk_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRdpclientprofileExist("citrixadc_rdpclientprofile.tf_rdpclientprofile_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_rdpclientprofile_ephem", "name", "tf_rdpclientprofile_ephem"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_rdpclientprofile_ephem", "rdpurloverride", "ENABLE"),
				),
			},
		},
	})
}

// Test ephemeral path: using psk_wo (WriteOnly attribute) with version tracker
const testAccRdpclientprofile_psk_wo_step1 = `
	variable "rdpclientprofile_psk_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_rdpclientprofile" "tf_rdpclientprofile_ephem" {
		name           = "tf_rdpclientprofile_ephem"
		rdpurloverride = "ENABLE"
		psk_wo         = var.rdpclientprofile_psk_wo
		psk_wo_version = 1
	}
`

const testAccRdpclientprofile_psk_wo_step2 = `
	variable "rdpclientprofile_psk_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_rdpclientprofile" "tf_rdpclientprofile_ephem" {
		name           = "tf_rdpclientprofile_ephem"
		rdpurloverride = "ENABLE"
		psk_wo         = var.rdpclientprofile_psk_wo_2
		psk_wo_version = 2
	}
`

func TestAccRdpclientprofile_psk_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_rdpclientprofile_psk_wo", "ephem_psk1")
	t.Setenv("TF_VAR_rdpclientprofile_psk_wo_2", "ephem_psk2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRdpclientprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRdpclientprofile_psk_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRdpclientprofileExist("citrixadc_rdpclientprofile.tf_rdpclientprofile_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_rdpclientprofile_ephem", "name", "tf_rdpclientprofile_ephem"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_rdpclientprofile_ephem", "psk_wo_version", "1"),
				),
			},
			{
				Config: testAccRdpclientprofile_psk_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRdpclientprofileExist("citrixadc_rdpclientprofile.tf_rdpclientprofile_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_rdpclientprofile_ephem", "name", "tf_rdpclientprofile_ephem"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_rdpclientprofile_ephem", "psk_wo_version", "2"),
				),
			},
		},
	})
}

// Step 1: every unset-eligible attribute set to a valid non-default value.
const testAccRdpclientprofile_unset_step1 = `
resource "citrixadc_rdpclientprofile" "tf_unset" {
	name                 = "tf_test_rdpclientprofile_unset"
	addusernameinrdpfile = "YES"
	audiocapturemode     = "ENABLE"
	keyboardhook         = "OnRemote"
	multimonitorsupport  = "DISABLE"
	randomizerdpfilename = "YES"
	rdpcookievalidity    = 120
	rdpurloverride       = "DISABLE"
	rdpvalidateclientip  = "ENABLE"
	redirectclipboard    = "DISABLE"
	redirectcomports     = "ENABLE"
	redirectdrives       = "ENABLE"
	redirectpnpdevices   = "ENABLE"
	redirectprinters     = "DISABLE"
	videoplaybackmode    = "DISABLE"
}
`

// Step 2: all eligible attributes removed from config -> provider must unset them,
// reverting each to its NITRO default.
const testAccRdpclientprofile_unset_step2 = `
resource "citrixadc_rdpclientprofile" "tf_unset" {
	name = "tf_test_rdpclientprofile_unset"
}
`

func TestAccRdpclientprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRdpclientprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values apply and persist.
				Config: testAccRdpclientprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRdpclientprofileExist("citrixadc_rdpclientprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_unset", "addusernameinrdpfile", "YES"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_unset", "audiocapturemode", "ENABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_unset", "keyboardhook", "OnRemote"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_unset", "multimonitorsupport", "DISABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_unset", "randomizerdpfilename", "YES"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_unset", "rdpcookievalidity", "120"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_unset", "rdpurloverride", "DISABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_unset", "rdpvalidateclientip", "ENABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_unset", "redirectclipboard", "DISABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_unset", "redirectcomports", "ENABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_unset", "redirectdrives", "ENABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_unset", "redirectpnpdevices", "ENABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_unset", "redirectprinters", "DISABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_unset", "videoplaybackmode", "DISABLE"),
				),
			},
			{
				// Removing them must unset -> state reverts to NITRO defaults,
				// and the implicit post-apply plan must be empty.
				Config: testAccRdpclientprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRdpclientprofileExist("citrixadc_rdpclientprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_unset", "addusernameinrdpfile", "NO"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_unset", "audiocapturemode", "DISABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_unset", "keyboardhook", "InFullScreenMode"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_unset", "multimonitorsupport", "ENABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_unset", "randomizerdpfilename", "NO"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_unset", "rdpcookievalidity", "60"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_unset", "rdpurloverride", "ENABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_unset", "rdpvalidateclientip", "DISABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_unset", "redirectclipboard", "ENABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_unset", "redirectcomports", "DISABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_unset", "redirectdrives", "DISABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_unset", "redirectpnpdevices", "DISABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_unset", "redirectprinters", "ENABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_unset", "videoplaybackmode", "ENABLE"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckRdpclientprofileADCValue("tf_test_rdpclientprofile_unset", "addusernameinrdpfile", "NO"),
					testAccCheckRdpclientprofileADCValue("tf_test_rdpclientprofile_unset", "audiocapturemode", "DISABLE"),
					testAccCheckRdpclientprofileADCValue("tf_test_rdpclientprofile_unset", "keyboardhook", "InFullScreenMode"),
					testAccCheckRdpclientprofileADCValue("tf_test_rdpclientprofile_unset", "multimonitorsupport", "ENABLE"),
					testAccCheckRdpclientprofileADCValue("tf_test_rdpclientprofile_unset", "randomizerdpfilename", "NO"),
					testAccCheckRdpclientprofileADCValue("tf_test_rdpclientprofile_unset", "rdpcookievalidity", "60"),
					testAccCheckRdpclientprofileADCValue("tf_test_rdpclientprofile_unset", "rdpurloverride", "ENABLE"),
					testAccCheckRdpclientprofileADCValue("tf_test_rdpclientprofile_unset", "rdpvalidateclientip", "DISABLE"),
					testAccCheckRdpclientprofileADCValue("tf_test_rdpclientprofile_unset", "redirectclipboard", "ENABLE"),
					testAccCheckRdpclientprofileADCValue("tf_test_rdpclientprofile_unset", "redirectcomports", "DISABLE"),
					testAccCheckRdpclientprofileADCValue("tf_test_rdpclientprofile_unset", "redirectdrives", "DISABLE"),
					testAccCheckRdpclientprofileADCValue("tf_test_rdpclientprofile_unset", "redirectpnpdevices", "DISABLE"),
					testAccCheckRdpclientprofileADCValue("tf_test_rdpclientprofile_unset", "redirectprinters", "ENABLE"),
					testAccCheckRdpclientprofileADCValue("tf_test_rdpclientprofile_unset", "videoplaybackmode", "ENABLE"),
				),
			},
		},
	})
}

// testAccCheckRdpclientprofileADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckRdpclientprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Rdpclientprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("rdpclientprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("rdpclientprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccRdpclientprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckRdpclientprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccRdpclientprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRdpclientprofileExist("citrixadc_rdpclientprofile.tf_rdpclientprofile", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccRdpclientprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRdpclientprofileExist("citrixadc_rdpclientprofile.tf_rdpclientprofile", nil),
				),
			},
		},
	})
}
