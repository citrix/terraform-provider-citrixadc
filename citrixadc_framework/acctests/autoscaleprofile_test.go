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

const testAccAutoscaleprofile_basic = `


resource "citrixadc_autoscaleprofile" "tf_autoscaleprofile" {
	name         = "my_autoscaleprofile"
	type         = "CLOUDSTACK"
	apikey       = "7c177611-4a18-42b0-a7c5-bfd811fd590f"
	url          = "www.service.example.com"
	sharedsecret = "YZEH6jkTqZWQ8r0o6kWj0mWruN3vXbtT"
	}
`

const testAccAutoscaleprofileDataSource_basic = `

resource "citrixadc_autoscaleprofile" "tf_autoscaleprofile" {
	name         = "my_autoscaleprofile"
	type         = "CLOUDSTACK"
	apikey       = "7c177611-4a18-42b0-a7c5-bfd811fd590f"
	url          = "www.service.example.com"
	sharedsecret = "YZEH6jkTqZWQ8r0o6kWj0mWruN3vXbtT"
}

data "citrixadc_autoscaleprofile" "tf_autoscaleprofile" {
	name = citrixadc_autoscaleprofile.tf_autoscaleprofile.name
}
`
const testAccAutoscaleprofile_update = `


resource "citrixadc_autoscaleprofile" "tf_autoscaleprofile" {
	name         = "my_autoscaleprofile"
	type         = "CLOUDSTACK"
	apikey       = "88e0ae91-4cd0-4dd5-8fa1-dcd38165f4a2"
	url          = "www.service2.example.com"
	sharedsecret = "vruE8whIW8qnAvUGtT3EpmeIFp690nGo"
	}
`

func TestAccAutoscaleprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAutoscaleprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAutoscaleprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAutoscaleprofileExist("citrixadc_autoscaleprofile.tf_autoscaleprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.tf_autoscaleprofile", "name", "my_autoscaleprofile"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.tf_autoscaleprofile", "type", "CLOUDSTACK"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.tf_autoscaleprofile", "apikey", "7c177611-4a18-42b0-a7c5-bfd811fd590f"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.tf_autoscaleprofile", "url", "www.service.example.com"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.tf_autoscaleprofile", "sharedsecret", "YZEH6jkTqZWQ8r0o6kWj0mWruN3vXbtT"),
				),
			},
			{
				Config: testAccAutoscaleprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAutoscaleprofileExist("citrixadc_autoscaleprofile.tf_autoscaleprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.tf_autoscaleprofile", "name", "my_autoscaleprofile"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.tf_autoscaleprofile", "type", "CLOUDSTACK"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.tf_autoscaleprofile", "apikey", "88e0ae91-4cd0-4dd5-8fa1-dcd38165f4a2"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.tf_autoscaleprofile", "url", "www.service2.example.com"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.tf_autoscaleprofile", "sharedsecret", "vruE8whIW8qnAvUGtT3EpmeIFp690nGo"),
				),
			},
		},
	})
}

func testAccCheckAutoscaleprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No autoscaleprofile name is set")
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
		data, err := client.FindResource(service.Autoscaleprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("autoscaleprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckAutoscaleprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_autoscaleprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Autoscaleprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("autoscaleprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

// --- apikey backward-compatible (sensitive attribute) tests ---

const testAccAutoscaleprofile_apikey_step1 = `
variable "autoscaleprofile_apikey" {
  type      = string
  sensitive = true
}

resource "citrixadc_autoscaleprofile" "test_apikey" {
  name         = "tf_autoscaleprofile_apikey"
  type         = "CLOUDSTACK"
  apikey       = var.autoscaleprofile_apikey
  url          = "www.service.example.com"
  sharedsecret = "YZEH6jkTqZWQ8r0o6kWj0mWruN3vXbtT"
}
`

const testAccAutoscaleprofile_apikey_step2 = `
variable "autoscaleprofile_apikey_2" {
  type      = string
  sensitive = true
}

resource "citrixadc_autoscaleprofile" "test_apikey" {
  name         = "tf_autoscaleprofile_apikey"
  type         = "CLOUDSTACK"
  apikey       = var.autoscaleprofile_apikey_2
  url          = "www.service.example.com"
  sharedsecret = "YZEH6jkTqZWQ8r0o6kWj0mWruN3vXbtT"
}
`

func TestAccAutoscaleprofile_apikey_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_autoscaleprofile_apikey", "7c177611-4a18-42b0-a7c5-bfd811fd590f")
	t.Setenv("TF_VAR_autoscaleprofile_apikey_2", "88e0ae91-4cd0-4dd5-8fa1-dcd38165f4a2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAutoscaleprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAutoscaleprofile_apikey_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAutoscaleprofileExist("citrixadc_autoscaleprofile.test_apikey", nil),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.test_apikey", "name", "tf_autoscaleprofile_apikey"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.test_apikey", "type", "CLOUDSTACK"),
				),
			},
			{
				Config: testAccAutoscaleprofile_apikey_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAutoscaleprofileExist("citrixadc_autoscaleprofile.test_apikey", nil),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.test_apikey", "name", "tf_autoscaleprofile_apikey"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.test_apikey", "type", "CLOUDSTACK"),
				),
			},
		},
	})
}

// --- apikey write-only (ephemeral) tests ---

const testAccAutoscaleprofile_apikey_wo_step1 = `
variable "autoscaleprofile_apikey_wo" {
  type      = string
  sensitive = true
}

resource "citrixadc_autoscaleprofile" "test_apikey_wo" {
  name             = "tf_autoscaleprofile_apikey_wo"
  type             = "CLOUDSTACK"
  apikey_wo        = var.autoscaleprofile_apikey_wo
  apikey_wo_version = 1
  url              = "www.service.example.com"
  sharedsecret     = "YZEH6jkTqZWQ8r0o6kWj0mWruN3vXbtT"
}
`

const testAccAutoscaleprofile_apikey_wo_step2 = `
variable "autoscaleprofile_apikey_wo_2" {
  type      = string
  sensitive = true
}

resource "citrixadc_autoscaleprofile" "test_apikey_wo" {
  name             = "tf_autoscaleprofile_apikey_wo"
  type             = "CLOUDSTACK"
  apikey_wo        = var.autoscaleprofile_apikey_wo_2
  apikey_wo_version = 2
  url              = "www.service.example.com"
  sharedsecret     = "YZEH6jkTqZWQ8r0o6kWj0mWruN3vXbtT"
}
`

func TestAccAutoscaleprofile_apikey_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_autoscaleprofile_apikey_wo", "7c177611-4a18-42b0-a7c5-bfd811fd590f")
	t.Setenv("TF_VAR_autoscaleprofile_apikey_wo_2", "88e0ae91-4cd0-4dd5-8fa1-dcd38165f4a2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAutoscaleprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAutoscaleprofile_apikey_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAutoscaleprofileExist("citrixadc_autoscaleprofile.test_apikey_wo", nil),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.test_apikey_wo", "apikey_wo_version", "1"),
				),
			},
			{
				Config: testAccAutoscaleprofile_apikey_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAutoscaleprofileExist("citrixadc_autoscaleprofile.test_apikey_wo", nil),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.test_apikey_wo", "apikey_wo_version", "2"),
				),
			},
		},
	})
}

// --- sharedsecret backward-compatible (sensitive attribute) tests ---

const testAccAutoscaleprofile_sharedsecret_step1 = `
variable "autoscaleprofile_sharedsecret" {
  type      = string
  sensitive = true
}

resource "citrixadc_autoscaleprofile" "test_sharedsecret" {
  name         = "tf_autoscaleprofile_ss"
  type         = "CLOUDSTACK"
  apikey       = "7c177611-4a18-42b0-a7c5-bfd811fd590f"
  url          = "www.service.example.com"
  sharedsecret = var.autoscaleprofile_sharedsecret
}
`

const testAccAutoscaleprofile_sharedsecret_step2 = `
variable "autoscaleprofile_sharedsecret_2" {
  type      = string
  sensitive = true
}

resource "citrixadc_autoscaleprofile" "test_sharedsecret" {
  name         = "tf_autoscaleprofile_ss"
  type         = "CLOUDSTACK"
  apikey       = "7c177611-4a18-42b0-a7c5-bfd811fd590f"
  url          = "www.service.example.com"
  sharedsecret = var.autoscaleprofile_sharedsecret_2
}
`

func TestAccAutoscaleprofile_sharedsecret_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_autoscaleprofile_sharedsecret", "YZEH6jkTqZWQ8r0o6kWj0mWruN3vXbtT")
	t.Setenv("TF_VAR_autoscaleprofile_sharedsecret_2", "vruE8whIW8qnAvUGtT3EpmeIFp690nGo")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAutoscaleprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAutoscaleprofile_sharedsecret_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAutoscaleprofileExist("citrixadc_autoscaleprofile.test_sharedsecret", nil),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.test_sharedsecret", "name", "tf_autoscaleprofile_ss"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.test_sharedsecret", "type", "CLOUDSTACK"),
				),
			},
			{
				Config: testAccAutoscaleprofile_sharedsecret_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAutoscaleprofileExist("citrixadc_autoscaleprofile.test_sharedsecret", nil),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.test_sharedsecret", "name", "tf_autoscaleprofile_ss"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.test_sharedsecret", "type", "CLOUDSTACK"),
				),
			},
		},
	})
}

// --- sharedsecret write-only (ephemeral) tests ---

const testAccAutoscaleprofile_sharedsecret_wo_step1 = `
variable "autoscaleprofile_sharedsecret_wo" {
  type      = string
  sensitive = true
}

resource "citrixadc_autoscaleprofile" "test_sharedsecret_wo" {
  name                    = "tf_autoscaleprofile_ss_wo"
  type                    = "CLOUDSTACK"
  apikey                  = "7c177611-4a18-42b0-a7c5-bfd811fd590f"
  url                     = "www.service.example.com"
  sharedsecret_wo         = var.autoscaleprofile_sharedsecret_wo
  sharedsecret_wo_version = 1
}
`

const testAccAutoscaleprofile_sharedsecret_wo_step2 = `
variable "autoscaleprofile_sharedsecret_wo_2" {
  type      = string
  sensitive = true
}

resource "citrixadc_autoscaleprofile" "test_sharedsecret_wo" {
  name                    = "tf_autoscaleprofile_ss_wo"
  type                    = "CLOUDSTACK"
  apikey                  = "7c177611-4a18-42b0-a7c5-bfd811fd590f"
  url                     = "www.service.example.com"
  sharedsecret_wo         = var.autoscaleprofile_sharedsecret_wo_2
  sharedsecret_wo_version = 2
}
`

func TestAccAutoscaleprofile_sharedsecret_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_autoscaleprofile_sharedsecret_wo", "YZEH6jkTqZWQ8r0o6kWj0mWruN3vXbtT")
	t.Setenv("TF_VAR_autoscaleprofile_sharedsecret_wo_2", "vruE8whIW8qnAvUGtT3EpmeIFp690nGo")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAutoscaleprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAutoscaleprofile_sharedsecret_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAutoscaleprofileExist("citrixadc_autoscaleprofile.test_sharedsecret_wo", nil),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.test_sharedsecret_wo", "sharedsecret_wo_version", "1"),
				),
			},
			{
				Config: testAccAutoscaleprofile_sharedsecret_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAutoscaleprofileExist("citrixadc_autoscaleprofile.test_sharedsecret_wo", nil),
					resource.TestCheckResourceAttr("citrixadc_autoscaleprofile.test_sharedsecret_wo", "sharedsecret_wo_version", "2"),
				),
			},
		},
	})
}

func TestAccAutoscaleprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAutoscaleprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_autoscaleprofile.tf_autoscaleprofile", "name", "my_autoscaleprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_autoscaleprofile.tf_autoscaleprofile", "type", "CLOUDSTACK"),
					resource.TestCheckResourceAttr("data.citrixadc_autoscaleprofile.tf_autoscaleprofile", "url", "www.service.example.com"),
					resource.TestCheckResourceAttrSet("data.citrixadc_autoscaleprofile.tf_autoscaleprofile", "apikey"),
					resource.TestCheckResourceAttrSet("data.citrixadc_autoscaleprofile.tf_autoscaleprofile", "sharedsecret"),
				),
			},
		},
	})
}
