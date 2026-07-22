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

func TestAccLbprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbprofileExist("citrixadc_lbprofile.tf_lbprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "lbprofilename", "tf_lbprofile"),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "dbslb", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "processlocal", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "httponlycookieflag", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "lbhashfingers", "258"),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "lbhashalgorithm", "PRAC"),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "storemqttclientidandusername", "YES"),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "proximityfromself", "NO"),
					testAccCheckUserAgent(),
				),
			},
			{
				Config: testAccLbprofile_basic_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbprofileExist("citrixadc_lbprofile.tf_lbprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "lbprofilename", "tf_lbprofile"),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "dbslb", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "processlocal", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "httponlycookieflag", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "lbhashfingers", "255"),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "lbhashalgorithm", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "storemqttclientidandusername", "NO"),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile", "proximityfromself", "YES"),
					testAccCheckUserAgent(),
				),
			},
		},
	})
}

func TestAccLbprofile_import(t *testing.T) {
	const resAddr = "citrixadc_lbprofile.tf_lbprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLbprofile_basic},
			{
				Config:            testAccLbprofile_basic,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// cookiepassphrase_wo_version is a write-only version tracker that
				// defaults to 1 in config but is never read back from NITRO, so it
				// cannot round-trip through import.
				ImportStateVerifyIgnore: []string{"cookiepassphrase_wo_version"},
			},
		},
	})
}

func testAccCheckLbprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lbprofile name is set")
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
		data, err := client.FindResource("lbprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Lbprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckLbprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_Lbprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lbprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Lbprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLbprofile_basic = `
resource "citrixadc_lbprofile" "tf_lbprofile" {
    lbprofilename = "tf_lbprofile"
    dbslb = "ENABLED"
	processlocal = "DISABLED"
	httponlycookieflag = "ENABLED"
	lbhashfingers = 258
	lbhashalgorithm = "PRAC"
	storemqttclientidandusername = "YES"
	proximityfromself = "NO"
}

`

const testAccLbprofile_basic_update = `

resource "citrixadc_lbprofile" "tf_lbprofile" {
    lbprofilename = "tf_lbprofile"
	dbslb = "DISABLED"
	processlocal = "ENABLED"
	httponlycookieflag = "DISABLED"
	lbhashfingers = 255
	lbhashalgorithm = "DEFAULT"
	storemqttclientidandusername = "NO"
	proximityfromself = "YES"
    
}

`

const testAccLbprofileDataSource_basic = `
resource "citrixadc_lbprofile" "tf_lbprofile_ds" {
    lbprofilename = "tf_lbprofile_ds"
    dbslb = "ENABLED"
	processlocal = "DISABLED"
	httponlycookieflag = "ENABLED"
	lbhashfingers = 258
	lbhashalgorithm = "PRAC"
	storemqttclientidandusername = "YES"
	proximityfromself = "NO"
}

data "citrixadc_lbprofile" "tf_lbprofile_ds" {
	lbprofilename = citrixadc_lbprofile.tf_lbprofile_ds.lbprofilename
	depends_on = [citrixadc_lbprofile.tf_lbprofile_ds]
}
`

func TestAccLbprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccLbprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lbprofile.tf_lbprofile_ds", "lbprofilename", "tf_lbprofile_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_lbprofile.tf_lbprofile_ds", "dbslb", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_lbprofile.tf_lbprofile_ds", "processlocal", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_lbprofile.tf_lbprofile_ds", "httponlycookieflag", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_lbprofile.tf_lbprofile_ds", "lbhashfingers", "258"),
					resource.TestCheckResourceAttr("data.citrixadc_lbprofile.tf_lbprofile_ds", "lbhashalgorithm", "PRAC"),
					resource.TestCheckResourceAttr("data.citrixadc_lbprofile.tf_lbprofile_ds", "storemqttclientidandusername", "YES"),
					resource.TestCheckResourceAttr("data.citrixadc_lbprofile.tf_lbprofile_ds", "proximityfromself", "NO"),
				),
			},
		},
	})
}

// Backward-compatible path: sensitive cookiepassphrase attribute
const testAccLbprofile_cookiepassphrase_step1 = `

	variable "lbprofile_cookiepassphrase" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_lbprofile" "tf_lbprofile_cp" {
		lbprofilename    = "tf_lbprofile_cp"
		cookiepassphrase = var.lbprofile_cookiepassphrase
	}
`

// Update backward-compatible path: change cookiepassphrase value
const testAccLbprofile_cookiepassphrase_step2 = `

	variable "lbprofile_cookiepassphrase_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_lbprofile" "tf_lbprofile_cp" {
		lbprofilename    = "tf_lbprofile_cp"
		cookiepassphrase = var.lbprofile_cookiepassphrase_2
	}
`

func TestAccLbprofile_cookiepassphrase_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_lbprofile_cookiepassphrase", "oldpassphrase123")
	t.Setenv("TF_VAR_lbprofile_cookiepassphrase_2", "newpassphrase456")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbprofile_cookiepassphrase_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbprofileExist("citrixadc_lbprofile.tf_lbprofile_cp", nil),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile_cp", "lbprofilename", "tf_lbprofile_cp"),
				),
			},
			{
				Config: testAccLbprofile_cookiepassphrase_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbprofileExist("citrixadc_lbprofile.tf_lbprofile_cp", nil),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile_cp", "lbprofilename", "tf_lbprofile_cp"),
				),
			},
		},
	})
}

// Ephemeral path: using cookiepassphrase_wo (WriteOnly attribute) with version tracker
const testAccLbprofile_cookiepassphrase_wo_step1 = `

	variable "lbprofile_cookiepassphrase_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_lbprofile" "tf_lbprofile_cp_wo" {
		lbprofilename               = "tf_lbprofile_cp_wo"
		cookiepassphrase_wo         = var.lbprofile_cookiepassphrase_wo
		cookiepassphrase_wo_version = 1
	}
`

// Update ephemeral path: bump version to trigger update with new passphrase
const testAccLbprofile_cookiepassphrase_wo_step2 = `

	variable "lbprofile_cookiepassphrase_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_lbprofile" "tf_lbprofile_cp_wo" {
		lbprofilename               = "tf_lbprofile_cp_wo"
		cookiepassphrase_wo         = var.lbprofile_cookiepassphrase_wo_2
		cookiepassphrase_wo_version = 2
	}
`

func TestAccLbprofile_cookiepassphrase_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_lbprofile_cookiepassphrase_wo", "ephemeral_pass1")
	t.Setenv("TF_VAR_lbprofile_cookiepassphrase_wo_2", "ephemeral_pass2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbprofile_cookiepassphrase_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbprofileExist("citrixadc_lbprofile.tf_lbprofile_cp_wo", nil),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile_cp_wo", "lbprofilename", "tf_lbprofile_cp_wo"),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile_cp_wo", "cookiepassphrase_wo_version", "1"),
				),
			},
			{
				Config: testAccLbprofile_cookiepassphrase_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbprofileExist("citrixadc_lbprofile.tf_lbprofile_cp_wo", nil),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile_cp_wo", "lbprofilename", "tf_lbprofile_cp_wo"),
					resource.TestCheckResourceAttr("citrixadc_lbprofile.tf_lbprofile_cp_wo", "cookiepassphrase_wo_version", "2"),
				),
			},
		},
	})
}

func TestAccLbprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLbprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccLbprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbprofileExist("citrixadc_lbprofile.tf_lbprofile", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccLbprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbprofileExist("citrixadc_lbprofile.tf_lbprofile", nil),
				),
			},
		},
	})
}
