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

const testAccIpsecalgprofile_basic = `
	resource "citrixadc_ipsecalgprofile" "tf_ipsecalgprofile" {
		name              = "my_ipsecalgprofile"
		ikesessiontimeout = 50
		espsessiontimeout = 20
		connfailover      = "DISABLED"
	}
  
`

const testAccIpsecalgprofile_update = `
	resource "citrixadc_ipsecalgprofile" "tf_ipsecalgprofile" {
		name              = "my_ipsecalgprofile"
		ikesessiontimeout = 40
		espsessiontimeout = 30
		connfailover      = "ENABLED"
	}
  
`

func TestAccIpsecalgprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpsecalgprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIpsecalgprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpsecalgprofileExist("citrixadc_ipsecalgprofile.tf_ipsecalgprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_ipsecalgprofile.tf_ipsecalgprofile", "name", "my_ipsecalgprofile"),
					resource.TestCheckResourceAttr("citrixadc_ipsecalgprofile.tf_ipsecalgprofile", "ikesessiontimeout", "50"),
					resource.TestCheckResourceAttr("citrixadc_ipsecalgprofile.tf_ipsecalgprofile", "espsessiontimeout", "20"),
					resource.TestCheckResourceAttr("citrixadc_ipsecalgprofile.tf_ipsecalgprofile", "connfailover", "DISABLED"),
				),
			},
			{
				Config: testAccIpsecalgprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpsecalgprofileExist("citrixadc_ipsecalgprofile.tf_ipsecalgprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_ipsecalgprofile.tf_ipsecalgprofile", "name", "my_ipsecalgprofile"),
					resource.TestCheckResourceAttr("citrixadc_ipsecalgprofile.tf_ipsecalgprofile", "ikesessiontimeout", "40"),
					resource.TestCheckResourceAttr("citrixadc_ipsecalgprofile.tf_ipsecalgprofile", "espsessiontimeout", "30"),
					resource.TestCheckResourceAttr("citrixadc_ipsecalgprofile.tf_ipsecalgprofile", "connfailover", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckIpsecalgprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ipsecalgprofile name is set")
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
		data, err := client.FindResource("ipsecalgprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("ipsecalgprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckIpsecalgprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_ipsecalgprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("ipsecalgprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("ipsecalgprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccIpsecalgprofile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_ipsecalgprofile.tf_ipsecalgprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpsecalgprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIpsecalgprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIpsecalgprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Ipsecalgprofile.Type(), "my_ipsecalgprofile"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccIpsecalgprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIpsecalgprofileExist(resAddr, nil)),
			},
		},
	})
}

func TestAccIpsecalgprofile_import(t *testing.T) {
	const resAddr = "citrixadc_ipsecalgprofile.tf_ipsecalgprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpsecalgprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccIpsecalgprofile_basic},
			{
				Config:                  testAccIpsecalgprofile_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccIpsecalgprofileDataSource_basic = `
	resource "citrixadc_ipsecalgprofile" "tf_ipsecalgprofile_ds" {
		name              = "my_ipsecalgprofile_ds"
		ikesessiontimeout = 50
		espsessiontimeout = 20
		connfailover      = "DISABLED"
	}

	data "citrixadc_ipsecalgprofile" "tf_ipsecalgprofile_ds" {
		name = citrixadc_ipsecalgprofile.tf_ipsecalgprofile_ds.name
	}
`

func TestAccIpsecalgprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckIpsecalgprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccIpsecalgprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIpsecalgprofileExist("citrixadc_ipsecalgprofile.tf_ipsecalgprofile", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccIpsecalgprofile_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckIpsecalgprofileExist("citrixadc_ipsecalgprofile.tf_ipsecalgprofile", nil)),
			},
		},
	})
}

func TestAccIpsecalgprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIpsecalgprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_ipsecalgprofile.tf_ipsecalgprofile_ds", "name", "my_ipsecalgprofile_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_ipsecalgprofile.tf_ipsecalgprofile_ds", "ikesessiontimeout", "50"),
					resource.TestCheckResourceAttr("data.citrixadc_ipsecalgprofile.tf_ipsecalgprofile_ds", "espsessiontimeout", "20"),
					resource.TestCheckResourceAttr("data.citrixadc_ipsecalgprofile.tf_ipsecalgprofile_ds", "connfailover", "DISABLED"),
				),
			},
		},
	})
}
