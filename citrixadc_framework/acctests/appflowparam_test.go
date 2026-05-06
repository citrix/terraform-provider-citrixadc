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

const testAccAppflowparam_basic = `

resource "citrixadc_appflowparam" "tf_appflowparam" {
	templaterefresh     = 200
	flowrecordinterval  = 200
	httpcookie          = "ENABLED"
	httplocation        = "ENABLED"
	}
  
`
const testAccAppflowparam_update = `

resource "citrixadc_appflowparam" "tf_appflowparam" {
	templaterefresh     = 600
	flowrecordinterval  = 100
	httpcookie          = "DISABLED"
	httplocation        = "DISABLED"
	}
`

func TestAccAppflowparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppflowparam_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowparamExist("citrixadc_appflowparam.tf_appflowparam", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "templaterefresh", "200"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "flowrecordinterval", "200"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "httpcookie", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "httplocation", "ENABLED"),
				),
			},
			{
				Config: testAccAppflowparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowparamExist("citrixadc_appflowparam.tf_appflowparam", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "templaterefresh", "600"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "flowrecordinterval", "100"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "httpcookie", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "httplocation", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckAppflowparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appflowparam name is set")
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
		data, err := client.FindResource(service.Appflowparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appflowparam %s not found", n)
		}

		return nil
	}
}

const testAccAppflowparamDataSource_basic = `
	resource "citrixadc_appflowparam" "tf_appflowparam" {
		templaterefresh     = 200
		flowrecordinterval  = 200
		httpcookie          = "ENABLED"
		httplocation        = "ENABLED"
	}
	
	data "citrixadc_appflowparam" "tf_appflowparam" {
		depends_on = [citrixadc_appflowparam.tf_appflowparam]
	}
`

func TestAccAppflowparamDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAppflowparamDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appflowparam.tf_appflowparam", "templaterefresh", "200"),
					resource.TestCheckResourceAttr("data.citrixadc_appflowparam.tf_appflowparam", "flowrecordinterval", "200"),
					resource.TestCheckResourceAttr("data.citrixadc_appflowparam.tf_appflowparam", "httpcookie", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_appflowparam.tf_appflowparam", "httplocation", "ENABLED"),
				),
			},
		},
	})
}

// Test backward-compatible path: using analyticsauthtoken (Sensitive attribute)
const testAccAppflowparam_analyticsauthtoken_step1 = `

	variable "appflowparam_analyticsauthtoken" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_appflowparam" "tf_appflowparam" {
		templaterefresh    = 200
		flowrecordinterval = 200
		httpcookie         = "ENABLED"
		httplocation       = "ENABLED"
		analyticsauthtoken = var.appflowparam_analyticsauthtoken
	}
`

// Update backward-compatible path: change analyticsauthtoken value
const testAccAppflowparam_analyticsauthtoken_step2 = `

	variable "appflowparam_analyticsauthtoken_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_appflowparam" "tf_appflowparam" {
		templaterefresh    = 200
		flowrecordinterval = 200
		httpcookie         = "ENABLED"
		httplocation       = "ENABLED"
		analyticsauthtoken = var.appflowparam_analyticsauthtoken_2
	}
`

func TestAccAppflowparam_analyticsauthtoken_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_appflowparam_analyticsauthtoken", "authtoken_value1")
	t.Setenv("TF_VAR_appflowparam_analyticsauthtoken_2", "authtoken_value2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAppflowparam_analyticsauthtoken_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowparamExist("citrixadc_appflowparam.tf_appflowparam", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "templaterefresh", "200"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "flowrecordinterval", "200"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "httpcookie", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "httplocation", "ENABLED"),
				),
			},
			{
				Config: testAccAppflowparam_analyticsauthtoken_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowparamExist("citrixadc_appflowparam.tf_appflowparam", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "templaterefresh", "200"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "flowrecordinterval", "200"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "httpcookie", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "httplocation", "ENABLED"),
				),
			},
		},
	})
}

// Test ephemeral path: using analyticsauthtoken_wo (WriteOnly attribute) with version tracker
const testAccAppflowparam_analyticsauthtoken_wo_step1 = `

	variable "appflowparam_analyticsauthtoken_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_appflowparam" "tf_appflowparam" {
		templaterefresh              = 200
		flowrecordinterval           = 200
		httpcookie                   = "ENABLED"
		httplocation                 = "ENABLED"
		analyticsauthtoken_wo        = var.appflowparam_analyticsauthtoken_wo
		analyticsauthtoken_wo_version = 1
	}
`

// Update ephemeral path: bump version to trigger update with new token
const testAccAppflowparam_analyticsauthtoken_wo_step2 = `

	variable "appflowparam_analyticsauthtoken_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_appflowparam" "tf_appflowparam" {
		templaterefresh              = 200
		flowrecordinterval           = 200
		httpcookie                   = "ENABLED"
		httplocation                 = "ENABLED"
		analyticsauthtoken_wo        = var.appflowparam_analyticsauthtoken_wo_2
		analyticsauthtoken_wo_version = 2
	}
`

func TestAccAppflowparam_analyticsauthtoken_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_appflowparam_analyticsauthtoken_wo", "ephemeral_token1")
	t.Setenv("TF_VAR_appflowparam_analyticsauthtoken_wo_2", "ephemeral_token2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAppflowparam_analyticsauthtoken_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowparamExist("citrixadc_appflowparam.tf_appflowparam", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "templaterefresh", "200"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "flowrecordinterval", "200"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "httpcookie", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "httplocation", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "analyticsauthtoken_wo_version", "1"),
				),
			},
			{
				Config: testAccAppflowparam_analyticsauthtoken_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowparamExist("citrixadc_appflowparam.tf_appflowparam", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "templaterefresh", "200"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "flowrecordinterval", "200"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "httpcookie", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "httplocation", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "analyticsauthtoken_wo_version", "2"),
				),
			},
		},
	})
}
