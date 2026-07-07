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
