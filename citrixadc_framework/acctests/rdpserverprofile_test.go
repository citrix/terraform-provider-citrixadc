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

const testAccRdpserverprofile_basic = `

resource "citrixadc_rdpserverprofile" "tf_rdpserverprofile" {
	name           = "my_rdpserverprofile"
	psk            = "key"
	rdpredirection = "ENABLE"
	rdpport        = 4000
	}
  
`

const testAccRdpserverprofile_update = `


resource "citrixadc_rdpserverprofile" "tf_rdpserverprofile" {
	name           = "my_rdpserverprofile"
	psk            = "key"
	rdpredirection = "DISABLE"
	rdpport        = 4100
	}
  
`

func TestAccRdpserverprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRdpserverprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRdpserverprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRdpserverprofileExist("citrixadc_rdpserverprofile.tf_rdpserverprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_rdpserverprofile.tf_rdpserverprofile", "name", "my_rdpserverprofile"),
					resource.TestCheckResourceAttr("citrixadc_rdpserverprofile.tf_rdpserverprofile", "psk", "key"),
					resource.TestCheckResourceAttr("citrixadc_rdpserverprofile.tf_rdpserverprofile", "rdpredirection", "ENABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpserverprofile.tf_rdpserverprofile", "rdpport", "4000"),
				),
			},
			{
				Config: testAccRdpserverprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRdpserverprofileExist("citrixadc_rdpserverprofile.tf_rdpserverprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_rdpserverprofile.tf_rdpserverprofile", "name", "my_rdpserverprofile"),
					resource.TestCheckResourceAttr("citrixadc_rdpserverprofile.tf_rdpserverprofile", "psk", "key"),
					resource.TestCheckResourceAttr("citrixadc_rdpserverprofile.tf_rdpserverprofile", "rdpredirection", "DISABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpserverprofile.tf_rdpserverprofile", "rdpport", "4100"),
				),
			},
		},
	})
}

func testAccCheckRdpserverprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No rdpserverprofile name is set")
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
		data, err := client.FindResource("rdpserverprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("rdpserverprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckRdpserverprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_rdpserverprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("rdpserverprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("rdpserverprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccRdpserverprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRdpserverprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_rdpserverprofile.test", "name", "tf_rdpserverprofile"),
					resource.TestCheckResourceAttrSet("data.citrixadc_rdpserverprofile.test", "id"),
					resource.TestCheckResourceAttrSet("data.citrixadc_rdpserverprofile.test", "rdpip"),
					resource.TestCheckResourceAttrSet("data.citrixadc_rdpserverprofile.test", "rdpport"),
				),
			},
		},
	})
}

const testAccRdpserverprofileDataSource_basic = `
resource "citrixadc_rdpserverprofile" "tf_rdpserverprofile" {
	name    = "tf_rdpserverprofile"
	rdpip   = "192.168.2.10"
	rdpport = 3389
	psk     = "secret789"
}

data "citrixadc_rdpserverprofile" "test" {
	name = citrixadc_rdpserverprofile.tf_rdpserverprofile.name
}
`

// Test backward-compatible path: using psk (Sensitive attribute)
const testAccRdpserverprofile_psk_step1 = `
	variable "rdpserverprofile_psk" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_rdpserverprofile" "tf_rdpserverprofile_ephem" {
		name           = "tf_rdpserverprofile_ephem"
		psk            = var.rdpserverprofile_psk
		rdpredirection = "ENABLE"
		rdpport        = 3389
	}
`

const testAccRdpserverprofile_psk_step2 = `
	variable "rdpserverprofile_psk_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_rdpserverprofile" "tf_rdpserverprofile_ephem" {
		name           = "tf_rdpserverprofile_ephem"
		psk            = var.rdpserverprofile_psk_2
		rdpredirection = "ENABLE"
		rdpport        = 3389
	}
`

func TestAccRdpserverprofile_psk_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_rdpserverprofile_psk", "presharedkey1")
	t.Setenv("TF_VAR_rdpserverprofile_psk_2", "presharedkey2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRdpserverprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRdpserverprofile_psk_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRdpserverprofileExist("citrixadc_rdpserverprofile.tf_rdpserverprofile_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_rdpserverprofile.tf_rdpserverprofile_ephem", "name", "tf_rdpserverprofile_ephem"),
					resource.TestCheckResourceAttr("citrixadc_rdpserverprofile.tf_rdpserverprofile_ephem", "rdpredirection", "ENABLE"),
				),
			},
			{
				Config: testAccRdpserverprofile_psk_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRdpserverprofileExist("citrixadc_rdpserverprofile.tf_rdpserverprofile_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_rdpserverprofile.tf_rdpserverprofile_ephem", "name", "tf_rdpserverprofile_ephem"),
					resource.TestCheckResourceAttr("citrixadc_rdpserverprofile.tf_rdpserverprofile_ephem", "rdpredirection", "ENABLE"),
				),
			},
		},
	})
}

func TestAccRdpserverprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckRdpserverprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccRdpserverprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRdpserverprofileExist("citrixadc_rdpserverprofile.tf_rdpserverprofile", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccRdpserverprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRdpserverprofileExist("citrixadc_rdpserverprofile.tf_rdpserverprofile", nil),
				),
			},
		},
	})
}

// Test ephemeral path: using psk_wo (WriteOnly attribute) with version tracker
const testAccRdpserverprofile_psk_wo_step1 = `
	variable "rdpserverprofile_psk_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_rdpserverprofile" "tf_rdpserverprofile_ephem" {
		name           = "tf_rdpserverprofile_ephem"
		psk_wo         = var.rdpserverprofile_psk_wo
		psk_wo_version = 1
		rdpredirection = "ENABLE"
		rdpport        = 3389
	}
`

const testAccRdpserverprofile_psk_wo_step2 = `
	variable "rdpserverprofile_psk_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_rdpserverprofile" "tf_rdpserverprofile_ephem" {
		name           = "tf_rdpserverprofile_ephem"
		psk_wo         = var.rdpserverprofile_psk_wo_2
		psk_wo_version = 2
		rdpredirection = "ENABLE"
		rdpport        = 3389
	}
`

func TestAccRdpserverprofile_psk_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_rdpserverprofile_psk_wo", "ephem_psk1")
	t.Setenv("TF_VAR_rdpserverprofile_psk_wo_2", "ephem_psk2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRdpserverprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRdpserverprofile_psk_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRdpserverprofileExist("citrixadc_rdpserverprofile.tf_rdpserverprofile_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_rdpserverprofile.tf_rdpserverprofile_ephem", "name", "tf_rdpserverprofile_ephem"),
					resource.TestCheckResourceAttr("citrixadc_rdpserverprofile.tf_rdpserverprofile_ephem", "psk_wo_version", "1"),
				),
			},
			{
				Config: testAccRdpserverprofile_psk_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRdpserverprofileExist("citrixadc_rdpserverprofile.tf_rdpserverprofile_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_rdpserverprofile.tf_rdpserverprofile_ephem", "name", "tf_rdpserverprofile_ephem"),
					resource.TestCheckResourceAttr("citrixadc_rdpserverprofile.tf_rdpserverprofile_ephem", "psk_wo_version", "2"),
				),
			},
		},
	})
}
