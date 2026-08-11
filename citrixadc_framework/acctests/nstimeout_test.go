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

const testAccNstimeout_basic = `


resource "citrixadc_nstimeout" "tf_nstimeout" {
	zombie     = 60
	anyclient     = 2000
	server     = 2000
	httpclient = 2000
	reducedrsttimeout = 10
	}
  
`
const testAccNstimeout_update = `


resource "citrixadc_nstimeout" "tf_nstimeout" {
	zombie     = 70
	anyclient     = 2300
	server     = 2400
	httpclient = 2500
	reducedrsttimeout = 15
	}
  
`

func TestAccNstimeout_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNstimeout_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstimeoutExist("citrixadc_nstimeout.tf_nstimeout", nil),
					resource.TestCheckResourceAttr("citrixadc_nstimeout.tf_nstimeout", "zombie", "60"),
					resource.TestCheckResourceAttr("citrixadc_nstimeout.tf_nstimeout", "anyclient", "2000"),
					resource.TestCheckResourceAttr("citrixadc_nstimeout.tf_nstimeout", "server", "2000"),
					resource.TestCheckResourceAttr("citrixadc_nstimeout.tf_nstimeout", "httpclient", "2000"),
					resource.TestCheckResourceAttr("citrixadc_nstimeout.tf_nstimeout", "reducedrsttimeout", "10"),
				),
			},
			{
				Config: testAccNstimeout_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstimeoutExist("citrixadc_nstimeout.tf_nstimeout", nil),
					resource.TestCheckResourceAttr("citrixadc_nstimeout.tf_nstimeout", "zombie", "70"),
					resource.TestCheckResourceAttr("citrixadc_nstimeout.tf_nstimeout", "anyclient", "2300"),
					resource.TestCheckResourceAttr("citrixadc_nstimeout.tf_nstimeout", "server", "2400"),
					resource.TestCheckResourceAttr("citrixadc_nstimeout.tf_nstimeout", "httpclient", "2500"),
					resource.TestCheckResourceAttr("citrixadc_nstimeout.tf_nstimeout", "reducedrsttimeout", "15"),
				),
			},
		},
	})
}

func testAccCheckNstimeoutExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nstimeout name is set")
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
		data, err := client.FindResource(service.Nstimeout.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nstimeout %s not found", n)
		}

		return nil
	}
}

func TestAccNstimeout_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccNstimeout_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNstimeoutExist("citrixadc_nstimeout.tf_nstimeout", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccNstimeout_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckNstimeoutExist("citrixadc_nstimeout.tf_nstimeout", nil)),
			},
		},
	})
}

// nstimeout is a singleton. step1 sets the unset-eligible attributes to
// non-default values; step2 removes them from config so the provider must unset
// them, reverting each to its documented NITRO default.
const testAccNstimeout_unset_step1 = `
resource "citrixadc_nstimeout" "tf_unset" {
	zombie            = 90
	anyclient         = 1500
	server            = 1500
	httpclient        = 1500
	reducedrsttimeout = 20
}
`

const testAccNstimeout_unset_step2 = `
resource "citrixadc_nstimeout" "tf_unset" {
	# All unset-eligible attributes removed from config -> the provider must
	# unset them (revert to NITRO defaults).
}
`

func TestAccNstimeout_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccNstimeout_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstimeoutExist("citrixadc_nstimeout.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nstimeout.tf_unset", "zombie", "90"),
					resource.TestCheckResourceAttr("citrixadc_nstimeout.tf_unset", "anyclient", "1500"),
					resource.TestCheckResourceAttr("citrixadc_nstimeout.tf_unset", "server", "1500"),
					resource.TestCheckResourceAttr("citrixadc_nstimeout.tf_unset", "httpclient", "1500"),
					resource.TestCheckResourceAttr("citrixadc_nstimeout.tf_unset", "reducedrsttimeout", "20"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccNstimeout_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstimeoutExist("citrixadc_nstimeout.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nstimeout.tf_unset", "zombie", "120"),
					resource.TestCheckResourceAttr("citrixadc_nstimeout.tf_unset", "anyclient", "0"),
					resource.TestCheckResourceAttr("citrixadc_nstimeout.tf_unset", "server", "0"),
					resource.TestCheckResourceAttr("citrixadc_nstimeout.tf_unset", "httpclient", "0"),
					resource.TestCheckResourceAttr("citrixadc_nstimeout.tf_unset", "reducedrsttimeout", "0"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckNstimeoutADCValue("zombie", "120"),
					testAccCheckNstimeoutADCValue("anyclient", "0"),
					testAccCheckNstimeoutADCValue("reducedrsttimeout", "0"),
				),
			},
		},
	})
}

// testAccCheckNstimeoutADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckNstimeoutADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Nstimeout.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("nstimeout not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("nstimeout: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccNstimeout_import(t *testing.T) {
	const resAddr = "citrixadc_nstimeout.tf_nstimeout"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccNstimeout_basic},
			{
				Config:                  testAccNstimeout_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccNstimeoutDataSource_basic = `

	resource "citrixadc_nstimeout" "tf_nstimeout" {
		zombie     = 60
		client     = 2000
		server     = 2000
		httpclient = 2000
		reducedrsttimeout = 10
	}

	data "citrixadc_nstimeout" "tf_nstimeout_data" {
		depends_on = [citrixadc_nstimeout.tf_nstimeout]
	}
`

func TestAccNstimeoutDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNstimeoutDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nstimeout.tf_nstimeout_data", "zombie", "60"),
					resource.TestCheckResourceAttr("data.citrixadc_nstimeout.tf_nstimeout_data", "client", "2000"),
					resource.TestCheckResourceAttr("data.citrixadc_nstimeout.tf_nstimeout_data", "server", "2000"),
					resource.TestCheckResourceAttr("data.citrixadc_nstimeout.tf_nstimeout_data", "httpclient", "2000"),
					resource.TestCheckResourceAttr("data.citrixadc_nstimeout.tf_nstimeout_data", "reducedrsttimeout", "10"),
				),
			},
		},
	})
}
