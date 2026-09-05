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
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testAccAuthenticationauthnprofile_add = `

	resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
		name           = "tf_authenticationvserver"
		servicetype    = "SSL"
		comment        = "new_vserver"
		authentication = "ON"
		state          = "DISABLED"
	}
	resource "citrixadc_authenticationauthnprofile" "tf_authenticationauthnprofile" {
		name                = "tf_name"
		authnvsname         = citrixadc_authenticationvserver.tf_authenticationvserver.name
		authenticationhost  = "hostname"
		authenticationlevel = "20"
	}
`
const testAccAuthenticationauthnprofile_update = `

	resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
		name           = "tf_authenticationvserver"
		servicetype    = "SSL"
		comment        = "new_vserver"
		authentication = "ON"
		state          = "DISABLED"
	}
	resource "citrixadc_authenticationauthnprofile" "tf_authenticationauthnprofile" {
		name                = "tf_name"
		authnvsname         = citrixadc_authenticationvserver.tf_authenticationvserver.name
		authenticationhost  = "newhostname"
		authenticationlevel = "30"
	}
`

func TestAccAuthenticationauthnprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationauthnprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationauthnprofile_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationauthnprofileExist("citrixadc_authenticationauthnprofile.tf_authenticationauthnprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationauthnprofile.tf_authenticationauthnprofile", "name", "tf_name"),
					resource.TestCheckResourceAttr("citrixadc_authenticationauthnprofile.tf_authenticationauthnprofile", "authenticationhost", "hostname"),
					resource.TestCheckResourceAttr("citrixadc_authenticationauthnprofile.tf_authenticationauthnprofile", "authenticationlevel", "20"),
				),
			},
			{
				Config: testAccAuthenticationauthnprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationauthnprofileExist("citrixadc_authenticationauthnprofile.tf_authenticationauthnprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationauthnprofile.tf_authenticationauthnprofile", "name", "tf_name"),
					resource.TestCheckResourceAttr("citrixadc_authenticationauthnprofile.tf_authenticationauthnprofile", "authenticationhost", "newhostname"),
					resource.TestCheckResourceAttr("citrixadc_authenticationauthnprofile.tf_authenticationauthnprofile", "authenticationlevel", "30"),
				),
			},
		},
	})
}

func TestAccAuthenticationauthnprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAuthenticationauthnprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccAuthenticationauthnprofile_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationauthnprofileExist("citrixadc_authenticationauthnprofile.tf_authenticationauthnprofile", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccAuthenticationauthnprofile_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationauthnprofileExist("citrixadc_authenticationauthnprofile.tf_authenticationauthnprofile", nil)),
			},
		},
	})
}

func testAccCheckAuthenticationauthnprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No authenticationauthnprofile name is set")
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
		data, err := client.FindResource(service.Authenticationauthnprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("authenticationauthnprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuthenticationauthnprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_authenticationauthnprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Authenticationauthnprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("authenticationauthnprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccAuthenticationauthnprofile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_authenticationauthnprofile.tf_authenticationauthnprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationauthnprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationauthnprofile_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationauthnprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Authenticationauthnprofile.Type(), "tf_name"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAuthenticationauthnprofile_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuthenticationauthnprofileExist(resAddr, nil)),
			},
		},
	})
}

func TestAccAuthenticationauthnprofile_import(t *testing.T) {
	const resAddr = "citrixadc_authenticationauthnprofile.tf_authenticationauthnprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationauthnprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAuthenticationauthnprofile_add},
			{
				Config:                  testAccAuthenticationauthnprofile_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

// authenticationlevel is the only NITRO-unsettable mutable attribute for this
// resource (the unset spec payload lists authenticationdomain and
// authenticationlevel, but authenticationdomain reads back absent/null after
// unset which cannot round-trip through an Optional+Computed default, so only
// authenticationlevel is wired). NITRO default for authenticationlevel is 0.
const testAccAuthenticationauthnprofile_unset_step1 = `

	resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
		name           = "tf_authenticationvserver"
		servicetype    = "SSL"
		comment        = "new_vserver"
		authentication = "ON"
		state          = "DISABLED"
	}
	resource "citrixadc_authenticationauthnprofile" "tf_unset" {
		name                = "tf_test_authnprofile_unset"
		authnvsname         = citrixadc_authenticationvserver.tf_authenticationvserver.name
		authenticationlevel = 25
	}
`

const testAccAuthenticationauthnprofile_unset_step2 = `

	resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
		name           = "tf_authenticationvserver"
		servicetype    = "SSL"
		comment        = "new_vserver"
		authentication = "ON"
		state          = "DISABLED"
	}
	resource "citrixadc_authenticationauthnprofile" "tf_unset" {
		name        = "tf_test_authnprofile_unset"
		authnvsname = citrixadc_authenticationvserver.tf_authenticationvserver.name
		# authenticationlevel removed from config -> provider must unset it
		# (revert to NITRO default 0).
	}
`

func TestAccAuthenticationauthnprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthenticationauthnprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value applied and persisted.
				Config: testAccAuthenticationauthnprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationauthnprofileExist("citrixadc_authenticationauthnprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationauthnprofile.tf_unset", "authenticationlevel", "25"),
				),
			},
			{
				// Removing the attribute must unset it: state (read back from the
				// appliance) reverts to the NITRO default, and the implicit
				// post-apply plan must be empty.
				Config: testAccAuthenticationauthnprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthenticationauthnprofileExist("citrixadc_authenticationauthnprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_authenticationauthnprofile.tf_unset", "authenticationlevel", "0"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAuthenticationauthnprofileADCValue("tf_test_authnprofile_unset", "authenticationlevel", "0"),
				),
			},
		},
	})
}

// testAccCheckAuthenticationauthnprofileADCValue asserts an attribute's value
// directly on the appliance (not just in Terraform state), proving the unset
// actually reverted it.
func testAccCheckAuthenticationauthnprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Authenticationauthnprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("authenticationauthnprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("authenticationauthnprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

const testAccAuthenticationauthnprofileDataSource_basic = `

	resource "citrixadc_authenticationvserver" "tf_authenticationvserver" {
		name           = "tf_authenticationvserver"
		servicetype    = "SSL"
		comment        = "new_vserver"
		authentication = "ON"
		state          = "DISABLED"
	}
	resource "citrixadc_authenticationauthnprofile" "tf_authenticationauthnprofile" {
		name                = "tf_name"
		authnvsname         = citrixadc_authenticationvserver.tf_authenticationvserver.name
		authenticationhost  = "hostname"
		authenticationlevel = "20"
	}

	data "citrixadc_authenticationauthnprofile" "tf_authenticationauthnprofile_datasource" {
		name = citrixadc_authenticationauthnprofile.tf_authenticationauthnprofile.name
	}
`

func TestAccAuthenticationauthnprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAuthenticationauthnprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_authenticationauthnprofile.tf_authenticationauthnprofile_datasource", "name", "tf_name"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationauthnprofile.tf_authenticationauthnprofile_datasource", "authnvsname", "tf_authenticationvserver"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationauthnprofile.tf_authenticationauthnprofile_datasource", "authenticationhost", "hostname"),
					resource.TestCheckResourceAttr("data.citrixadc_authenticationauthnprofile.tf_authenticationauthnprofile_datasource", "authenticationlevel", "20"),
				),
			},
		},
	})
}
