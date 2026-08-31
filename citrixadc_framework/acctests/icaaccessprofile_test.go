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

const testAccIcaaccessprofile_basic = `


	resource "citrixadc_icaaccessprofile" "tf_icaaccessprofile" {
		name                   = "my_ica_accessprofile"
		connectclientlptports  = "DEFAULT"
		localremotedatasharing = "DEFAULT"
		wiaredirection		 = "DISABLED"
		smartcardredirection	= "DISABLED"
		fido2redirection		 = "DISABLED"
		draganddrop			 = "DISABLED"
		clienttwaindeviceredirection = "DISABLED"
	}
	
`
const testAccIcaaccessprofile_update = `


	resource "citrixadc_icaaccessprofile" "tf_icaaccessprofile" {
		name                   = "my_ica_accessprofile"
		connectclientlptports  = "DISABLED"
		localremotedatasharing = "DISABLED"
		wiaredirection		 = "DEFAULT"
		smartcardredirection	= "DEFAULT"
		fido2redirection		 = "DEFAULT"
		draganddrop			 = "DEFAULT"
		clienttwaindeviceredirection = "DEFAULT"
	}
	
`

const testAccIcaaccessprofileDataSource_basic = `

	resource "citrixadc_icaaccessprofile" "tf_icaaccessprofile" {
		name                   = "my_ica_accessprofile"
		connectclientlptports  = "DEFAULT"
		localremotedatasharing = "DEFAULT"
		wiaredirection		 = "DISABLED"
		smartcardredirection	= "DISABLED"
		fido2redirection		 = "DISABLED"
		draganddrop			 = "DISABLED"
		clienttwaindeviceredirection = "DISABLED"
	}

	data "citrixadc_icaaccessprofile" "icaaccessprofile_data" {
		name = citrixadc_icaaccessprofile.tf_icaaccessprofile.name
	}
`

func TestAccIcaaccessprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIcaaccessprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIcaaccessprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcaaccessprofileExist("citrixadc_icaaccessprofile.tf_icaaccessprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "name", "my_ica_accessprofile"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "connectclientlptports", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "localremotedatasharing", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "wiaredirection", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "smartcardredirection", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "fido2redirection", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "draganddrop", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "clienttwaindeviceredirection", "DISABLED"),
				),
			},
			{
				Config: testAccIcaaccessprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcaaccessprofileExist("citrixadc_icaaccessprofile.tf_icaaccessprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "name", "my_ica_accessprofile"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "connectclientlptports", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "localremotedatasharing", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "wiaredirection", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "smartcardredirection", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "fido2redirection", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "draganddrop", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_icaaccessprofile", "clienttwaindeviceredirection", "DEFAULT"),
				),
			},
		},
	})
}

func testAccCheckIcaaccessprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No icaaccessprofile name is set")
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
		data, err := client.FindResource("icaaccessprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("icaaccessprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckIcaaccessprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_icaaccessprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("icaaccessprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("icaaccessprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccIcaaccessprofile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_icaaccessprofile.tf_icaaccessprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIcaaccessprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccIcaaccessprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIcaaccessprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Icaaccessprofile.Type(), "my_ica_accessprofile"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccIcaaccessprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIcaaccessprofileExist(resAddr, nil)),
			},
		},
	})
}

func TestAccIcaaccessprofile_import(t *testing.T) {
	const resAddr = "citrixadc_icaaccessprofile.tf_icaaccessprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIcaaccessprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccIcaaccessprofile_basic},
			{
				Config:                  testAccIcaaccessprofile_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccIcaaccessprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckIcaaccessprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccIcaaccessprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIcaaccessprofileExist("citrixadc_icaaccessprofile.tf_icaaccessprofile", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccIcaaccessprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckIcaaccessprofileExist("citrixadc_icaaccessprofile.tf_icaaccessprofile", nil)),
			},
		},
	})
}

// All 14 mutable attributes are Optional+Computed with a documented NITRO
// default of "DISABLED", so every one is unset-eligible. Step 1 sets them all
// to the non-default "DEFAULT"; step 2 removes them and the provider must unset
// them (revert to "DISABLED").
const testAccIcaaccessprofile_unset_step1 = `
resource "citrixadc_icaaccessprofile" "tf_unset" {
  name                         = "tf_test_icaaccessprofile_unset"
  clientaudioredirection       = "DEFAULT"
  clientclipboardredirection   = "DEFAULT"
  clientcomportredirection     = "DEFAULT"
  clientdriveredirection       = "DEFAULT"
  clientprinterredirection     = "DEFAULT"
  clienttwaindeviceredirection = "DEFAULT"
  clientusbdriveredirection    = "DEFAULT"
  connectclientlptports        = "DEFAULT"
  draganddrop                  = "DEFAULT"
  fido2redirection             = "DEFAULT"
  localremotedatasharing       = "DEFAULT"
  multistream                  = "DEFAULT"
  smartcardredirection         = "DEFAULT"
  wiaredirection               = "DEFAULT"
}
`

const testAccIcaaccessprofile_unset_step2 = `
resource "citrixadc_icaaccessprofile" "tf_unset" {
  name = "tf_test_icaaccessprofile_unset"
  # All unset-eligible attributes removed from config -> the provider must
  # unset them (revert to NITRO defaults, "DISABLED").
}
`

func TestAccIcaaccessprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIcaaccessprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccIcaaccessprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcaaccessprofileExist("citrixadc_icaaccessprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_unset", "clientaudioredirection", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_unset", "clientclipboardredirection", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_unset", "clientcomportredirection", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_unset", "clientdriveredirection", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_unset", "clientprinterredirection", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_unset", "clienttwaindeviceredirection", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_unset", "clientusbdriveredirection", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_unset", "connectclientlptports", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_unset", "draganddrop", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_unset", "fido2redirection", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_unset", "localremotedatasharing", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_unset", "multistream", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_unset", "smartcardredirection", "DEFAULT"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_unset", "wiaredirection", "DEFAULT"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccIcaaccessprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIcaaccessprofileExist("citrixadc_icaaccessprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_unset", "clientaudioredirection", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_unset", "clientclipboardredirection", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_unset", "clientcomportredirection", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_unset", "clientdriveredirection", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_unset", "clientprinterredirection", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_unset", "clienttwaindeviceredirection", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_unset", "clientusbdriveredirection", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_unset", "connectclientlptports", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_unset", "draganddrop", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_unset", "fido2redirection", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_unset", "localremotedatasharing", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_unset", "multistream", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_unset", "smartcardredirection", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_icaaccessprofile.tf_unset", "wiaredirection", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckIcaaccessprofileADCValue("tf_test_icaaccessprofile_unset", "clientaudioredirection", "DISABLED"),
					testAccCheckIcaaccessprofileADCValue("tf_test_icaaccessprofile_unset", "draganddrop", "DISABLED"),
					testAccCheckIcaaccessprofileADCValue("tf_test_icaaccessprofile_unset", "wiaredirection", "DISABLED"),
				),
			},
		},
	})
}

// testAccCheckIcaaccessprofileADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckIcaaccessprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Icaaccessprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("icaaccessprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("icaaccessprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccIcaaccessprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIcaaccessprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_icaaccessprofile.icaaccessprofile_data", "name", "my_ica_accessprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_icaaccessprofile.icaaccessprofile_data", "connectclientlptports", "DEFAULT"),
					resource.TestCheckResourceAttr("data.citrixadc_icaaccessprofile.icaaccessprofile_data", "localremotedatasharing", "DEFAULT"),
					resource.TestCheckResourceAttr("data.citrixadc_icaaccessprofile.icaaccessprofile_data", "wiaredirection", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_icaaccessprofile.icaaccessprofile_data", "smartcardredirection", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_icaaccessprofile.icaaccessprofile_data", "fido2redirection", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_icaaccessprofile.icaaccessprofile_data", "draganddrop", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_icaaccessprofile.icaaccessprofile_data", "clienttwaindeviceredirection", "DISABLED"),
				),
			},
		},
	})
}
