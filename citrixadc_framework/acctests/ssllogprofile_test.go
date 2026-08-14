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

const testAccSsllogprofile_basic = `
resource "citrixadc_ssllogprofile" "demo_ssllogprofile" {
    name = "demo_ssllogprofile"
    ssllogclauth = "ENABLED"
    ssllogclauthfailures = "ENABLED"
    sslloghs = "ENABLED"
    sslloghsfailures = "ENABLED"	
}
`

const testAccSsllogprofile_update = `
resource "citrixadc_ssllogprofile" "demo_ssllogprofile" {
    name = "demo_ssllogprofile"
    ssllogclauth = "DISABLED"
    ssllogclauthfailures = "ENABLED"
    sslloghs = "DISABLED"
    sslloghsfailures = "ENABLED"	
}
`

const testAccSsllogprofileDataSource_basic = `
resource "citrixadc_ssllogprofile" "demo_ssllogprofile" {
    name = "demo_ssllogprofile"
    ssllogclauth = "ENABLED"
    ssllogclauthfailures = "ENABLED"
    sslloghs = "ENABLED"
    sslloghsfailures = "ENABLED"	
}

data "citrixadc_ssllogprofile" "demo_ssllogprofile" {
    name = citrixadc_ssllogprofile.demo_ssllogprofile.name
}
`

func TestAccSsllogprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSsllogprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSsllogprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSsllogprofileExist("citrixadc_ssllogprofile.demo_ssllogprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_ssllogprofile.demo_ssllogprofile", "name", "demo_ssllogprofile"),
					resource.TestCheckResourceAttr("citrixadc_ssllogprofile.demo_ssllogprofile", "ssllogclauth", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_ssllogprofile.demo_ssllogprofile", "ssllogclauthfailures", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_ssllogprofile.demo_ssllogprofile", "sslloghs", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_ssllogprofile.demo_ssllogprofile", "sslloghsfailures", "ENABLED"),
				),
			},
			{
				Config: testAccSsllogprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSsllogprofileExist("citrixadc_ssllogprofile.demo_ssllogprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_ssllogprofile.demo_ssllogprofile", "name", "demo_ssllogprofile"),
					resource.TestCheckResourceAttr("citrixadc_ssllogprofile.demo_ssllogprofile", "ssllogclauth", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_ssllogprofile.demo_ssllogprofile", "ssllogclauthfailures", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_ssllogprofile.demo_ssllogprofile", "sslloghs", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_ssllogprofile.demo_ssllogprofile", "sslloghsfailures", "ENABLED"),
				),
			},
		},
	})
}

func TestAccSsllogprofile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_ssllogprofile.demo_ssllogprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSsllogprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSsllogprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSsllogprofileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Ssllogprofile.Type(), "demo_ssllogprofile"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccSsllogprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSsllogprofileExist(resAddr, nil)),
			},
		},
	})
}

func testAccCheckSsllogprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ssllogprofile name is set")
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
		data, err := client.FindResource("ssllogprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("ssllogprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckSsllogprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_ssllogprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("ssllogprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("ssllogprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccSsllogprofile_import(t *testing.T) {
	const resAddr = "citrixadc_ssllogprofile.demo_ssllogprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSsllogprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccSsllogprofile_basic},
			{
				Config:                  testAccSsllogprofile_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccSsllogprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckSsllogprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccSsllogprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSsllogprofileExist("citrixadc_ssllogprofile.demo_ssllogprofile", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccSsllogprofile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSsllogprofileExist("citrixadc_ssllogprofile.demo_ssllogprofile", nil)),
			},
		},
	})
}

const testAccSsllogprofile_unset_step1 = `
resource "citrixadc_ssllogprofile" "tf_unset" {
    name                 = "tf_test_ssllogprofile_unset"
    ssllogclauth         = "ENABLED"
    ssllogclauthfailures = "ENABLED"
    sslloghs             = "ENABLED"
    sslloghsfailures     = "ENABLED"
}
`

const testAccSsllogprofile_unset_step2 = `
resource "citrixadc_ssllogprofile" "tf_unset" {
    name = "tf_test_ssllogprofile_unset"
    # All unset-eligible attributes removed from config -> the provider must
    # unset them (revert to NITRO defaults, "DISABLED").
}
`

func TestAccSsllogprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSsllogprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccSsllogprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSsllogprofileExist("citrixadc_ssllogprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_ssllogprofile.tf_unset", "ssllogclauth", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_ssllogprofile.tf_unset", "ssllogclauthfailures", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_ssllogprofile.tf_unset", "sslloghs", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_ssllogprofile.tf_unset", "sslloghsfailures", "ENABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccSsllogprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSsllogprofileExist("citrixadc_ssllogprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_ssllogprofile.tf_unset", "ssllogclauth", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_ssllogprofile.tf_unset", "ssllogclauthfailures", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_ssllogprofile.tf_unset", "sslloghs", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_ssllogprofile.tf_unset", "sslloghsfailures", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckSsllogprofileADCValue("tf_test_ssllogprofile_unset", "ssllogclauth", "DISABLED"),
					testAccCheckSsllogprofileADCValue("tf_test_ssllogprofile_unset", "sslloghs", "DISABLED"),
				),
			},
		},
	})
}

// testAccCheckSsllogprofileADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckSsllogprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Ssllogprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("ssllogprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("ssllogprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccSsllogprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSsllogprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_ssllogprofile.demo_ssllogprofile", "name", "demo_ssllogprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_ssllogprofile.demo_ssllogprofile", "ssllogclauth", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_ssllogprofile.demo_ssllogprofile", "ssllogclauthfailures", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_ssllogprofile.demo_ssllogprofile", "sslloghs", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_ssllogprofile.demo_ssllogprofile", "sslloghsfailures", "ENABLED"),
				),
			},
		},
	})
}
