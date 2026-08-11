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

const testAccNsassignment_add = `
	resource "citrixadc_nsvariable" "tf_nsvariable" {
		name          = "tf_nsvariable"
		type          = "text(20)"
		scope         = "global"
		iffull        = "undef"
		ifvaluetoobig = "undef"
		ifnovalue     = "init"
		comment       = "Testing"
	}
	resource "citrixadc_nsassignment" "tf_nsassignment" {
		name     = "tf_nsassignment"
		variable = join("", ["$", citrixadc_nsvariable.tf_nsvariable.name])
		set      = 1
		comment  = "Testing"
	}
`
const testAccNsassignment_update = `
	resource "citrixadc_nsvariable" "tf_nsvariable" {
		name          = "tf_nsvariable"
		type          = "text(20)"
		scope         = "global"
		iffull        = "undef"
		ifvaluetoobig = "undef"
		ifnovalue     = "init"
		comment       = "Testing"
	}
	resource "citrixadc_nsassignment" "tf_nsassignment" {
		name     = "tf_nsassignment"
		variable = join("", ["$", citrixadc_nsvariable.tf_nsvariable.name])
		set      = 1
		comment  = "Testing_updated"
	}
`

func TestAccNsassignment_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsassignmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsassignment_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsassignmentExist("citrixadc_nsassignment.tf_nsassignment", nil),
					resource.TestCheckResourceAttr("citrixadc_nsassignment.tf_nsassignment", "name", "tf_nsassignment"),
					resource.TestCheckResourceAttr("citrixadc_nsassignment.tf_nsassignment", "set", "1"),
					resource.TestCheckResourceAttr("citrixadc_nsassignment.tf_nsassignment", "comment", "Testing"),
				),
			},
			{
				Config: testAccNsassignment_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsassignmentExist("citrixadc_nsassignment.tf_nsassignment", nil),
					resource.TestCheckResourceAttr("citrixadc_nsassignment.tf_nsassignment", "name", "tf_nsassignment"),
					resource.TestCheckResourceAttr("citrixadc_nsassignment.tf_nsassignment", "set", "1"),
					resource.TestCheckResourceAttr("citrixadc_nsassignment.tf_nsassignment", "comment", "Testing_updated"),
				),
			},
		},
	})
}

func testAccCheckNsassignmentExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsassignment name is set")
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
		data, err := client.FindResource(service.Nsassignment.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nsassignment %s not found", n)
		}

		return nil
	}
}

func testAccCheckNsassignmentDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nsassignment" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Nsassignment.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nsassignment %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccNsassignment_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_nsassignment.tf_nsassignment"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsassignmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsassignment_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsassignmentExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Nsassignment.Type(), "tf_nsassignment"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccNsassignment_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsassignmentExist(resAddr, nil)),
			},
		},
	})
}

func TestAccNsassignment_import(t *testing.T) {
	const resAddr = "citrixadc_nsassignment.tf_nsassignment"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsassignmentDestroy,
		Steps: []resource.TestStep{
			{Config: testAccNsassignment_add},
			{
				Config:                  testAccNsassignment_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccNsassignmentDataSource_basic = `
	resource "citrixadc_nsvariable" "tf_nsvariable" {
		name          = "tf_nsvariable"
		type          = "text(20)"
		scope         = "global"
		iffull        = "undef"
		ifvaluetoobig = "undef"
		ifnovalue     = "init"
		comment       = "Testing"
	}
	resource "citrixadc_nsassignment" "tf_nsassignment" {
		name     = "tf_nsassignment"
		variable = join("", ["$", citrixadc_nsvariable.tf_nsvariable.name])
		set      = "1"
		comment  = "Testing"
	}

	data "citrixadc_nsassignment" "tf_nsassignment_data" {
		name = citrixadc_nsassignment.tf_nsassignment.name
	}
`

func TestAccNsassignment_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNsassignmentDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccNsassignment_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsassignmentExist("citrixadc_nsassignment.tf_nsassignment", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccNsassignment_add,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckNsassignmentExist("citrixadc_nsassignment.tf_nsassignment", nil)),
			},
		},
	})
}

// The nsassignment unset test covers the only spec-unsettable mutable
// attribute: comment. Step 1 sets it to a non-default value; step 2 removes it
// from config so the provider must issue a NITRO unset (reverting to the empty
// default).
const testAccNsassignment_unset_step1 = `
	resource "citrixadc_nsvariable" "tf_nsvariable" {
		name          = "tf_nsvariable"
		type          = "text(20)"
		scope         = "global"
		iffull        = "undef"
		ifvaluetoobig = "undef"
		ifnovalue     = "init"
		comment       = "Testing"
	}
	resource "citrixadc_nsassignment" "tf_unset" {
		name     = "tf_test_nsassignment_unset"
		variable = join("", ["$", citrixadc_nsvariable.tf_nsvariable.name])
		set      = "1"
		comment  = "unset_me"
	}
`

const testAccNsassignment_unset_step2 = `
	resource "citrixadc_nsvariable" "tf_nsvariable" {
		name          = "tf_nsvariable"
		type          = "text(20)"
		scope         = "global"
		iffull        = "undef"
		ifvaluetoobig = "undef"
		ifnovalue     = "init"
		comment       = "Testing"
	}
	resource "citrixadc_nsassignment" "tf_unset" {
		name     = "tf_test_nsassignment_unset"
		variable = join("", ["$", citrixadc_nsvariable.tf_nsvariable.name])
		set      = "1"
		# comment removed from config -> the provider must unset it.
	}
`

func TestAccNsassignment_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsassignmentDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccNsassignment_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsassignmentExist("citrixadc_nsassignment.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nsassignment.tf_unset", "comment", "unset_me"),
				),
			},
			{
				// Removing comment must unset it: state reverts to the NITRO
				// default (empty), the implicit post-apply plan must be empty,
				// and the appliance confirms the revert.
				Config: testAccNsassignment_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsassignmentExist("citrixadc_nsassignment.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nsassignment.tf_unset", "comment", ""),
					testAccCheckNsassignmentADCValue("tf_test_nsassignment_unset", "comment", ""),
				),
			},
		},
	})
}

// testAccCheckNsassignmentADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it.
func testAccCheckNsassignmentADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Nsassignment.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("nsassignment %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		// NITRO omits an empty/default string attribute from GET, so a reverted
		// value comes back as a missing key (nil) -> treat that as empty.
		if _, present := data[attr]; !present || got == "<nil>" {
			got = ""
		}
		if got != want {
			return fmt.Errorf("nsassignment %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccNsassignmentDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsassignmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsassignmentDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nsassignment.tf_nsassignment_data", "name", "tf_nsassignment"),
					resource.TestCheckResourceAttr("data.citrixadc_nsassignment.tf_nsassignment_data", "set", "1"),
					resource.TestCheckResourceAttr("data.citrixadc_nsassignment.tf_nsassignment_data", "comment", "Testing"),
				),
			},
		},
	})
}
