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

const testAccRadiusnode_basic = `
	resource "citrixadc_radiusnode" "tf_radiusnode" {
		nodeprefix = "10.10.10.10/32"
		radkey     = "secret"
	}
`

func TestAccRadiusnode_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRadiusnodeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRadiusnode_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusnodeExist("citrixadc_radiusnode.tf_radiusnode", nil),
					resource.TestCheckResourceAttr("citrixadc_radiusnode.tf_radiusnode", "nodeprefix", "10.10.10.10/32"),
				),
			},
		},
	})
}

func TestAccRadiusnode_import(t *testing.T) {
	const resAddr = "citrixadc_radiusnode.tf_radiusnode"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRadiusnodeDestroy,
		Steps: []resource.TestStep{
			{Config: testAccRadiusnode_basic},
			{
				Config:            testAccRadiusnode_basic,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// radkey is a Sensitive secret that NITRO never echoes back, and
				// radkey_wo_version is a client-side version tracker that is not
				// returned by the API - neither can round-trip through import.
				ImportStateVerifyIgnore: []string{"radkey", "radkey_wo_version"},
			},
		},
	})
}

func testAccCheckRadiusnodeExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No radiusnode name is set")
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
		data, err := client.FindResource("radiusnode", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("radiusnode %s not found", n)
		}

		return nil
	}
}

func testAccCheckRadiusnodeDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_radiusnode" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("radiusnode", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("radiusnode %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccRadiusnodeDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRadiusnodeDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_radiusnode.test", "nodeprefix", "192.168.1.0/24"),
					resource.TestCheckResourceAttrSet("data.citrixadc_radiusnode.test", "id"),
				),
			},
		},
	})
}

const testAccRadiusnodeDataSource_basic = `
resource "citrixadc_radiusnode" "tf_radiusnode" {
	nodeprefix = "192.168.1.0/24"
	radkey     = "secret123"
}

data "citrixadc_radiusnode" "test" {
	nodeprefix = citrixadc_radiusnode.tf_radiusnode.nodeprefix
}
`

// Test backward-compatible path: using radkey (Sensitive attribute)
const testAccRadiusnode_radkey_step1 = `
	variable "radiusnode_radkey" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_radiusnode" "tf_radiusnode_ephem" {
		nodeprefix = "10.20.30.0/24"
		radkey     = var.radiusnode_radkey
	}
`

const testAccRadiusnode_radkey_step2 = `
	variable "radiusnode_radkey_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_radiusnode" "tf_radiusnode_ephem" {
		nodeprefix = "10.20.30.0/24"
		radkey     = var.radiusnode_radkey_2
	}
`

func TestAccRadiusnode_radkey_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_radiusnode_radkey", "secretkey1")
	t.Setenv("TF_VAR_radiusnode_radkey_2", "secretkey2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRadiusnodeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRadiusnode_radkey_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusnodeExist("citrixadc_radiusnode.tf_radiusnode_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_radiusnode.tf_radiusnode_ephem", "nodeprefix", "10.20.30.0/24"),
				),
			},
			{
				Config: testAccRadiusnode_radkey_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusnodeExist("citrixadc_radiusnode.tf_radiusnode_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_radiusnode.tf_radiusnode_ephem", "nodeprefix", "10.20.30.0/24"),
				),
			},
		},
	})
}

func TestAccRadiusnode_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckRadiusnodeDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccRadiusnode_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusnodeExist("citrixadc_radiusnode.tf_radiusnode", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccRadiusnode_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusnodeExist("citrixadc_radiusnode.tf_radiusnode", nil),
				),
			},
		},
	})
}

// Test ephemeral path: using radkey_wo (WriteOnly attribute) with version tracker
const testAccRadiusnode_radkey_wo_step1 = `
	variable "radiusnode_radkey_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_radiusnode" "tf_radiusnode_ephem" {
		nodeprefix        = "10.20.30.0/24"
		radkey_wo         = var.radiusnode_radkey_wo
		radkey_wo_version = 1
	}
`

const testAccRadiusnode_radkey_wo_step2 = `
	variable "radiusnode_radkey_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_radiusnode" "tf_radiusnode_ephem" {
		nodeprefix        = "10.20.30.0/24"
		radkey_wo         = var.radiusnode_radkey_wo_2
		radkey_wo_version = 2
	}
`

func TestAccRadiusnode_radkey_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_radiusnode_radkey_wo", "ephemeral_key1")
	t.Setenv("TF_VAR_radiusnode_radkey_wo_2", "ephemeral_key2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRadiusnodeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRadiusnode_radkey_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusnodeExist("citrixadc_radiusnode.tf_radiusnode_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_radiusnode.tf_radiusnode_ephem", "nodeprefix", "10.20.30.0/24"),
					resource.TestCheckResourceAttr("citrixadc_radiusnode.tf_radiusnode_ephem", "radkey_wo_version", "1"),
				),
			},
			{
				Config: testAccRadiusnode_radkey_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRadiusnodeExist("citrixadc_radiusnode.tf_radiusnode_ephem", nil),
					resource.TestCheckResourceAttr("citrixadc_radiusnode.tf_radiusnode_ephem", "nodeprefix", "10.20.30.0/24"),
					resource.TestCheckResourceAttr("citrixadc_radiusnode.tf_radiusnode_ephem", "radkey_wo_version", "2"),
				),
			},
		},
	})
}
