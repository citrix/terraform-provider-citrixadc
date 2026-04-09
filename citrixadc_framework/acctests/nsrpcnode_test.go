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

func TestAccNsrpcnode_basic(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	if isCpxRun {
		t.Skip("Operation not permitted under CPX")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNsrpcnode_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsrpcnodeExist("citrixadc_nsrpcnode.tf_nsrpcnode", nil),
				),
			},
			{
				Config: testAccNsrpcnode_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsrpcnodeExist("citrixadc_nsrpcnode.tf_nsrpcnode", nil),
				),
			},
		},
	})
}

func testAccCheckNsrpcnodeExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No id is set")
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
		data, err := client.FindResource(service.Nsrpcnode.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("RPC node %s not found", n)
		}

		return nil
	}
}

const testAccNsrpcnode_basic_step1 = `

resource "citrixadc_nsrpcnode" "tf_nsrpcnode" {
    ipaddress = "10.101.132.123"
    password = "CADS123$%^"
    secure = "ON"
    srcip = "10.101.132.123"
}
`

const testAccNsrpcnode_basic_step2 = `

resource "citrixadc_nsrpcnode" "tf_nsrpcnode" {
    ipaddress = "10.101.132.123"
    password = "CADS123$%^"
    secure = "OFF"
    srcip = "10.101.132.123"
}
`

const testAccNsrpcnodeDataSource_basic = `

resource "citrixadc_nsrpcnode" "tf_nsrpcnode" {
    ipaddress = "10.101.132.123"
    password = "CADS123$%^"
    secure = "ON"
    srcip = "10.101.132.123"
}

data "citrixadc_nsrpcnode" "tf_nsrpcnode" {
    ipaddress = citrixadc_nsrpcnode.tf_nsrpcnode.ipaddress
}
`

// Test backward-compatible path: using password (Sensitive attribute)
const testAccNsrpcnode_password_step1 = `

	variable "nsrpcnode_password" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_nsrpcnode" "tf_nsrpcnode" {
		ipaddress = "10.101.132.123"
		password  = var.nsrpcnode_password
		secure    = "ON"
		srcip     = "10.101.132.123"
	}
`

// Update backward-compatible path: change password value
const testAccNsrpcnode_password_step2 = `

	variable "nsrpcnode_password_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_nsrpcnode" "tf_nsrpcnode" {
		ipaddress = "10.101.132.123"
		password  = var.nsrpcnode_password_2
		secure    = "ON"
		srcip     = "10.101.132.123"
	}
`

func TestAccNsrpcnode_password_backward_compat(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	if isCpxRun {
		t.Skip("Operation not permitted under CPX")
	}
	t.Setenv("TF_VAR_nsrpcnode_password", "oldpassword123")
	t.Setenv("TF_VAR_nsrpcnode_password_2", "newpassword456")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNsrpcnode_password_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsrpcnodeExist("citrixadc_nsrpcnode.tf_nsrpcnode", nil),
					resource.TestCheckResourceAttr("citrixadc_nsrpcnode.tf_nsrpcnode", "secure", "ON"),
				),
			},
			{
				Config: testAccNsrpcnode_password_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsrpcnodeExist("citrixadc_nsrpcnode.tf_nsrpcnode", nil),
					resource.TestCheckResourceAttr("citrixadc_nsrpcnode.tf_nsrpcnode", "secure", "ON"),
				),
			},
		},
	})
}

// Test ephemeral path: using password_wo (WriteOnly attribute) with version tracker
const testAccNsrpcnode_password_wo_step1 = `

	variable "nsrpcnode_password_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_nsrpcnode" "tf_nsrpcnode" {
		ipaddress          = "10.101.132.123"
		password_wo        = var.nsrpcnode_password_wo
		password_wo_version = 1
		secure             = "ON"
		srcip              = "10.101.132.123"
	}
`

// Update ephemeral path: bump version to trigger update with new password
const testAccNsrpcnode_password_wo_step2 = `

	variable "nsrpcnode_password_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_nsrpcnode" "tf_nsrpcnode" {
		ipaddress          = "10.101.132.123"
		password_wo        = var.nsrpcnode_password_wo_2
		password_wo_version = 2
		secure             = "ON"
		srcip              = "10.101.132.123"
	}
`

func TestAccNsrpcnode_password_wo_ephemeral(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	if isCpxRun {
		t.Skip("Operation not permitted under CPX")
	}
	t.Setenv("TF_VAR_nsrpcnode_password_wo", "ephemeral_pass1")
	t.Setenv("TF_VAR_nsrpcnode_password_wo_2", "ephemeral_pass2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNsrpcnode_password_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsrpcnodeExist("citrixadc_nsrpcnode.tf_nsrpcnode", nil),
					resource.TestCheckResourceAttr("citrixadc_nsrpcnode.tf_nsrpcnode", "password_wo_version", "1"),
					resource.TestCheckResourceAttr("citrixadc_nsrpcnode.tf_nsrpcnode", "secure", "ON"),
				),
			},
			{
				Config: testAccNsrpcnode_password_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsrpcnodeExist("citrixadc_nsrpcnode.tf_nsrpcnode", nil),
					resource.TestCheckResourceAttr("citrixadc_nsrpcnode.tf_nsrpcnode", "password_wo_version", "2"),
					resource.TestCheckResourceAttr("citrixadc_nsrpcnode.tf_nsrpcnode", "secure", "ON"),
				),
			},
		},
	})
}

func TestAccNsrpcnodeDataSource_basic(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	if isCpxRun {
		t.Skip("Operation not permitted under CPX")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNsrpcnodeDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nsrpcnode.tf_nsrpcnode", "ipaddress", "10.101.132.123"),
					resource.TestCheckResourceAttr("data.citrixadc_nsrpcnode.tf_nsrpcnode", "secure", "ON"),
				),
			},
		},
	})
}
