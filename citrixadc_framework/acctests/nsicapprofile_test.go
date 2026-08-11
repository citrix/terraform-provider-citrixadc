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

const testAccNsicapprofile_add = `
	resource "citrixadc_nsicapprofile" "tf_nsicapprofile" {
		name             = "tf_nsicapprofile"
		uri              = "/example"
		mode             = "REQMOD"
		reqtimeout       = 4
		reqtimeoutaction = "RESET"
		preview          = "ENABLED"
		previewlength    = 4096
	}
`
const testAccNsicapprofile_update = `
	resource "citrixadc_nsicapprofile" "tf_nsicapprofile" {
		name             = "tf_nsicapprofile"
		uri              = "/hello"
		mode             = "REQMOD"
		reqtimeout       = 4
		reqtimeoutaction = "RESET"
		preview          = "DISABLED"
		previewlength    = 4096
	}
`

const testAccNsicapprofileDataSource_basic = `
	resource "citrixadc_nsicapprofile" "tf_nsicapprofile" {
		name             = "tf_nsicapprofile_ds"
		uri              = "/avscan"
		mode             = "REQMOD"
		reqtimeout       = 30
		reqtimeoutaction = "RESET"
		preview          = "ENABLED"
		previewlength    = 2048
		allow204         = "ENABLED"
		connectionkeepalive = "ENABLED"
	}

	data "citrixadc_nsicapprofile" "nsicapprofile_data" {
		name = citrixadc_nsicapprofile.tf_nsicapprofile.name
	}
`

func TestAccNsicapprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsicapprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsicapprofile_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsicapprofileExist("citrixadc_nsicapprofile.tf_nsicapprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_nsicapprofile.tf_nsicapprofile", "name", "tf_nsicapprofile"),
					resource.TestCheckResourceAttr("citrixadc_nsicapprofile.tf_nsicapprofile", "uri", "/example"),
					resource.TestCheckResourceAttr("citrixadc_nsicapprofile.tf_nsicapprofile", "preview", "ENABLED"),
				),
			},
			{
				Config: testAccNsicapprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsicapprofileExist("citrixadc_nsicapprofile.tf_nsicapprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_nsicapprofile.tf_nsicapprofile", "name", "tf_nsicapprofile"),
					resource.TestCheckResourceAttr("citrixadc_nsicapprofile.tf_nsicapprofile", "uri", "/hello"),
					resource.TestCheckResourceAttr("citrixadc_nsicapprofile.tf_nsicapprofile", "preview", "DISABLED"),
				),
			},
		},
	})
}

func TestAccNsicapprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNsicapprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nsicapprofile.nsicapprofile_data", "name", "tf_nsicapprofile_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_nsicapprofile.nsicapprofile_data", "uri", "/avscan"),
					resource.TestCheckResourceAttr("data.citrixadc_nsicapprofile.nsicapprofile_data", "mode", "REQMOD"),
					resource.TestCheckResourceAttr("data.citrixadc_nsicapprofile.nsicapprofile_data", "preview", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_nsicapprofile.nsicapprofile_data", "previewlength", "2048"),
					resource.TestCheckResourceAttr("data.citrixadc_nsicapprofile.nsicapprofile_data", "allow204", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_nsicapprofile.nsicapprofile_data", "connectionkeepalive", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_nsicapprofile.nsicapprofile_data", "reqtimeout", "30"),
					resource.TestCheckResourceAttr("data.citrixadc_nsicapprofile.nsicapprofile_data", "reqtimeoutaction", "RESET"),
				),
			},
		},
	})
}

func TestAccNsicapprofile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_nsicapprofile.tf_nsicapprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsicapprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsicapprofile_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsicapprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Nsicapprofile.Type(), "tf_nsicapprofile"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccNsicapprofile_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsicapprofileExist(resAddr, nil)),
			},
		},
	})
}

func TestAccNsicapprofile_import(t *testing.T) {
	const resAddr = "citrixadc_nsicapprofile.tf_nsicapprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsicapprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccNsicapprofile_add},
			{
				Config:                  testAccNsicapprofile_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccNsicapprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNsicapprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccNsicapprofile_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsicapprofileExist("citrixadc_nsicapprofile.tf_nsicapprofile", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccNsicapprofile_add,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckNsicapprofileExist("citrixadc_nsicapprofile.tf_nsicapprofile", nil)),
			},
		},
	})
}

// The nsicapprofile unset test covers the spec-unsettable attributes that have
// a documented NITRO server default: preview (DISABLED), previewlength (4096),
// connectionkeepalive (ENABLED), allow204 (ENABLED), reqtimeout (0) and
// reqtimeoutaction (RESET). uri and mode are Required/mandatory (not
// unsettable); the free-form string attributes (hostheader, useragent,
// queryparams, inserticapheaders, inserthttprequest, logaction) have no
// documented default and are excluded.
const testAccNsicapprofile_unset_step1 = `
resource "citrixadc_nsicapprofile" "tf_unset" {
  name                = "tf_nsicapprofile_unset"
  uri                 = "/avscan"
  mode                = "REQMOD"
  preview             = "ENABLED"
  previewlength       = 2048
  connectionkeepalive = "DISABLED"
  allow204            = "DISABLED"
  reqtimeout          = 30
  reqtimeoutaction    = "BYPASS"
}
`

const testAccNsicapprofile_unset_step2 = `
resource "citrixadc_nsicapprofile" "tf_unset" {
  name = "tf_nsicapprofile_unset"
  uri  = "/avscan"
  mode = "REQMOD"
  # All unset-eligible attributes removed from config -> the provider must
  # unset them (revert to NITRO defaults).
}
`

func TestAccNsicapprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsicapprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccNsicapprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsicapprofileExist("citrixadc_nsicapprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nsicapprofile.tf_unset", "preview", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nsicapprofile.tf_unset", "previewlength", "2048"),
					resource.TestCheckResourceAttr("citrixadc_nsicapprofile.tf_unset", "connectionkeepalive", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nsicapprofile.tf_unset", "allow204", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nsicapprofile.tf_unset", "reqtimeout", "30"),
					resource.TestCheckResourceAttr("citrixadc_nsicapprofile.tf_unset", "reqtimeoutaction", "BYPASS"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccNsicapprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsicapprofileExist("citrixadc_nsicapprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nsicapprofile.tf_unset", "preview", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nsicapprofile.tf_unset", "previewlength", "4096"),
					resource.TestCheckResourceAttr("citrixadc_nsicapprofile.tf_unset", "connectionkeepalive", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nsicapprofile.tf_unset", "allow204", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nsicapprofile.tf_unset", "reqtimeout", "0"),
					resource.TestCheckResourceAttr("citrixadc_nsicapprofile.tf_unset", "reqtimeoutaction", "RESET"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckNsicapprofileADCValue("tf_nsicapprofile_unset", "preview", "DISABLED"),
					testAccCheckNsicapprofileADCValue("tf_nsicapprofile_unset", "reqtimeoutaction", "RESET"),
				),
			},
		},
	})
}

// testAccCheckNsicapprofileADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckNsicapprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Nsicapprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("nsicapprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("nsicapprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func testAccCheckNsicapprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsicapprofile name is set")
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
		data, err := client.FindResource("nsicapprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nsicapprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckNsicapprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nsicapprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("nsicapprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nsicapprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
