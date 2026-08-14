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

const testAccPcpprofile_basic = `


resource "citrixadc_pcpprofile" "tf_pcpprofile" {
	name               = "my_pcpprofile"
	mapping            = "ENABLED"
	peer               = "ENABLED"
	}
  
`
const testAccPcpprofile_update = `


resource "citrixadc_pcpprofile" "tf_pcpprofile" {
	name               = "my_pcpprofile"
	mapping            = "DISABLED"
	peer               = "DISABLED"
	}
  
`

func TestAccPcpprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPcpprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPcpprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPcpprofileExist("citrixadc_pcpprofile.tf_pcpprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_pcpprofile.tf_pcpprofile", "name", "my_pcpprofile"),
					resource.TestCheckResourceAttr("citrixadc_pcpprofile.tf_pcpprofile", "mapping", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_pcpprofile.tf_pcpprofile", "peer", "ENABLED"),
				),
			},
			{
				Config: testAccPcpprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPcpprofileExist("citrixadc_pcpprofile.tf_pcpprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_pcpprofile.tf_pcpprofile", "name", "my_pcpprofile"),
					resource.TestCheckResourceAttr("citrixadc_pcpprofile.tf_pcpprofile", "mapping", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_pcpprofile.tf_pcpprofile", "peer", "DISABLED"),
				),
			},
		},
	})
}

func testAccCheckPcpprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No pcpprofile name is set")
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
		data, err := client.FindResource("pcpprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("pcpprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckPcpprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_pcpprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("pcpprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("pcpprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccPcpprofile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_pcpprofile.tf_pcpprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPcpprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPcpprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckPcpprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Pcpprofile.Type(), "my_pcpprofile"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccPcpprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckPcpprofileExist(resAddr, nil)),
			},
		},
	})
}

func TestAccPcpprofile_import(t *testing.T) {
	const resAddr = "citrixadc_pcpprofile.tf_pcpprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPcpprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccPcpprofile_basic},
			{
				Config:                  testAccPcpprofile_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccPcpprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckPcpprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccPcpprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckPcpprofileExist("citrixadc_pcpprofile.tf_pcpprofile", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccPcpprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckPcpprofileExist("citrixadc_pcpprofile.tf_pcpprofile", nil)),
			},
		},
	})
}

// The pcpprofile unset test covers every unset-eligible attribute. Step 1 sets
// all six mutable attributes to valid non-default values; step 2 removes them so
// the provider unsets them (reverting to the documented NITRO defaults).
const testAccPcpprofile_unset_step1 = `
resource "citrixadc_pcpprofile" "tf_unset" {
  name               = "tf_test_pcpprofile_unset"
  mapping            = "DISABLED"
  peer               = "DISABLED"
  thirdparty         = "ENABLED"
  minmaplife         = 240
  maxmaplife         = 43200
  announcemulticount = "20"
}
`

const testAccPcpprofile_unset_step2 = `
resource "citrixadc_pcpprofile" "tf_unset" {
  name = "tf_test_pcpprofile_unset"
  # All unset-eligible attributes removed from config -> the provider must
  # unset them (revert to NITRO defaults).
}
`

func TestAccPcpprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPcpprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccPcpprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPcpprofileExist("citrixadc_pcpprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_pcpprofile.tf_unset", "mapping", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_pcpprofile.tf_unset", "peer", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_pcpprofile.tf_unset", "thirdparty", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_pcpprofile.tf_unset", "minmaplife", "240"),
					resource.TestCheckResourceAttr("citrixadc_pcpprofile.tf_unset", "maxmaplife", "43200"),
					resource.TestCheckResourceAttr("citrixadc_pcpprofile.tf_unset", "announcemulticount", "20"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccPcpprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPcpprofileExist("citrixadc_pcpprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_pcpprofile.tf_unset", "mapping", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_pcpprofile.tf_unset", "peer", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_pcpprofile.tf_unset", "thirdparty", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_pcpprofile.tf_unset", "minmaplife", "120"),
					resource.TestCheckResourceAttr("citrixadc_pcpprofile.tf_unset", "maxmaplife", "86400"),
					resource.TestCheckResourceAttr("citrixadc_pcpprofile.tf_unset", "announcemulticount", "10"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckPcpprofileADCValue("tf_test_pcpprofile_unset", "mapping", "ENABLED"),
					testAccCheckPcpprofileADCValue("tf_test_pcpprofile_unset", "thirdparty", "DISABLED"),
				),
			},
		},
	})
}

// testAccCheckPcpprofileADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckPcpprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Pcpprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("pcpprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("pcpprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccPcpprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPcpprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPcpprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_pcpprofile.tf_pcpprofile_ds", "name", "my_pcpprofile_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_pcpprofile.tf_pcpprofile_ds", "mapping", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_pcpprofile.tf_pcpprofile_ds", "peer", "ENABLED"),
				),
			},
		},
	})
}

const testAccPcpprofileDataSource_basic = `

resource "citrixadc_pcpprofile" "tf_pcpprofile_ds" {
	name    = "my_pcpprofile_ds"
	mapping = "ENABLED"
	peer    = "ENABLED"
}

data "citrixadc_pcpprofile" "tf_pcpprofile_ds" {
	name = citrixadc_pcpprofile.tf_pcpprofile_ds.name
	depends_on = [citrixadc_pcpprofile.tf_pcpprofile_ds]
}
`
