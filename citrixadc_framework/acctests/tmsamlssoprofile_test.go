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

const testAccTmsamlssoprofile_basic = `

	resource "citrixadc_tmsamlssoprofile" "tf_tmsamlssoprofile" {
		name                        = "my_tmsamlssoprofile"
		assertionconsumerserviceurl = "https://service.example.com"
		sendpassword                = "OFF"
		relaystaterule              = "true"
	}
  
`
const testAccTmsamlssoprofile_update = `

	resource "citrixadc_tmsamlssoprofile" "tf_tmsamlssoprofile" {
		name                        = "my_tmsamlssoprofile"
		assertionconsumerserviceurl = "https://service.example2.com"
		sendpassword                = "ON"
		relaystaterule              = "false"
	}
  
`

func TestAccTmsamlssoprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTmsamlssoprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTmsamlssoprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmsamlssoprofileExist("citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile", "name", "my_tmsamlssoprofile"),
					resource.TestCheckResourceAttr("citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile", "assertionconsumerserviceurl", "https://service.example.com"),
					resource.TestCheckResourceAttr("citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile", "sendpassword", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile", "relaystaterule", "true"),
				),
			},
			{
				Config: testAccTmsamlssoprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmsamlssoprofileExist("citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile", "name", "my_tmsamlssoprofile"),
					resource.TestCheckResourceAttr("citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile", "assertionconsumerserviceurl", "https://service.example2.com"),
					resource.TestCheckResourceAttr("citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile", "sendpassword", "ON"),
					resource.TestCheckResourceAttr("citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile", "relaystaterule", "false"),
				),
			},
		},
	})
}

func testAccCheckTmsamlssoprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No tmsamlssoprofile name is set")
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
		data, err := client.FindResource(service.Tmsamlssoprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("tmsamlssoprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckTmsamlssoprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_tmsamlssoprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Tmsamlssoprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("tmsamlssoprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccTmsamlssoprofile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTmsamlssoprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTmsamlssoprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTmsamlssoprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Tmsamlssoprofile.Type(), "my_tmsamlssoprofile"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccTmsamlssoprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTmsamlssoprofileExist(resAddr, nil)),
			},
		},
	})
}

func TestAccTmsamlssoprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckTmsamlssoprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccTmsamlssoprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTmsamlssoprofileExist("citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccTmsamlssoprofile_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckTmsamlssoprofileExist("citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile", nil)),
			},
		},
	})
}

func TestAccTmsamlssoprofile_import(t *testing.T) {
	const resAddr = "citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTmsamlssoprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccTmsamlssoprofile_basic},
			{
				Config:                  testAccTmsamlssoprofile_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"sendpassword"},
			},
		},
	})
}

// Unset test: step1 sets the unset-eligible attributes to valid non-default
// values; step2 removes them from config so the provider must unset them
// (revert to the documented NITRO defaults).
const testAccTmsamlssoprofile_unset_step1 = `
	resource "citrixadc_tmsamlssoprofile" "tf_unset" {
		name                        = "tf_test_tmsamlssoprofile_unset"
		assertionconsumerserviceurl = "https://service.example.com"
		digestmethod                = "SHA1"
		nameidformat                = "persistent"
		signassertion               = "NONE"
		signaturealg                = "RSA-SHA1"
		skewtime                    = 10
	}
`

const testAccTmsamlssoprofile_unset_step2 = `
	resource "citrixadc_tmsamlssoprofile" "tf_unset" {
		name                        = "tf_test_tmsamlssoprofile_unset"
		assertionconsumerserviceurl = "https://service.example.com"
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccTmsamlssoprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTmsamlssoprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccTmsamlssoprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmsamlssoprofileExist("citrixadc_tmsamlssoprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_tmsamlssoprofile.tf_unset", "digestmethod", "SHA1"),
					resource.TestCheckResourceAttr("citrixadc_tmsamlssoprofile.tf_unset", "nameidformat", "persistent"),
					resource.TestCheckResourceAttr("citrixadc_tmsamlssoprofile.tf_unset", "signassertion", "NONE"),
					resource.TestCheckResourceAttr("citrixadc_tmsamlssoprofile.tf_unset", "signaturealg", "RSA-SHA1"),
					resource.TestCheckResourceAttr("citrixadc_tmsamlssoprofile.tf_unset", "skewtime", "10"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccTmsamlssoprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmsamlssoprofileExist("citrixadc_tmsamlssoprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_tmsamlssoprofile.tf_unset", "digestmethod", "SHA256"),
					resource.TestCheckResourceAttr("citrixadc_tmsamlssoprofile.tf_unset", "nameidformat", "transient"),
					resource.TestCheckResourceAttr("citrixadc_tmsamlssoprofile.tf_unset", "signassertion", "ASSERTION"),
					resource.TestCheckResourceAttr("citrixadc_tmsamlssoprofile.tf_unset", "signaturealg", "RSA-SHA256"),
					resource.TestCheckResourceAttr("citrixadc_tmsamlssoprofile.tf_unset", "skewtime", "5"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckTmsamlssoprofileADCValue("tf_test_tmsamlssoprofile_unset", "digestmethod", "SHA256"),
					testAccCheckTmsamlssoprofileADCValue("tf_test_tmsamlssoprofile_unset", "signassertion", "ASSERTION"),
				),
			},
		},
	})
}

// testAccCheckTmsamlssoprofileADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckTmsamlssoprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Tmsamlssoprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("tmsamlssoprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("tmsamlssoprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

const testAccTmsamlssoprofileDataSource_basic = `

	resource "citrixadc_tmsamlssoprofile" "tf_tmsamlssoprofile" {
		name                        = "my_tmsamlssoprofile"
		assertionconsumerserviceurl = "https://service.example.com"
		sendpassword                = "OFF"
		relaystaterule              = "true"
	}

data "citrixadc_tmsamlssoprofile" "tf_tmsamlssoprofile" {
    name = citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile.name
}
`

func TestAccTmsamlssoprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTmsamlssoprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile", "name", "my_tmsamlssoprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile", "assertionconsumerserviceurl", "https://service.example.com"),
					resource.TestCheckResourceAttr("data.citrixadc_tmsamlssoprofile.tf_tmsamlssoprofile", "relaystaterule", "true"),
				),
			},
		},
	})
}
