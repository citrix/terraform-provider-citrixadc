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

const testAccContentinspectionaction_basic = `
	resource "citrixadc_nsicapprofile" "tf_nsicapprofile" {
		name             = "new-profile"
		uri              = "/example"
		mode             = "REQMOD"
		reqtimeout       = 4
		reqtimeoutaction = "RESET"
		preview          = "ENABLED"
		previewlength    = 4096
	}
	resource "citrixadc_contentinspectionaction" "tf_contentinspectionaction" {
		name            = "my_ci_action"
		type            = "ICAP"
		icapprofilename = citrixadc_nsicapprofile.tf_nsicapprofile.name
		serverip      = "2.2.2.2"
		ifserverdown    = "DROP"
	}
`
const testAccContentinspectionaction_update = `

	resource "citrixadc_nsicapprofile" "tf_nsicapprofile" {
		name             = "new-profile"
		uri              = "/example"
		mode             = "REQMOD"
		reqtimeout       = 4
		reqtimeoutaction = "RESET"
		preview          = "ENABLED"
		previewlength    = 4096
	}
	resource "citrixadc_contentinspectionaction" "tf_contentinspectionaction" {
		name            = "my_ci_action"
		type            = "NOINSPECTION"
	}
`

func TestAccContentinspectionaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckContentinspectionactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccContentinspectionaction_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionactionExist("citrixadc_contentinspectionaction.tf_contentinspectionaction", nil),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionaction.tf_contentinspectionaction", "name", "my_ci_action"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionaction.tf_contentinspectionaction", "type", "ICAP"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionaction.tf_contentinspectionaction", "icapprofilename", "new-profile"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionaction.tf_contentinspectionaction", "serverip", "2.2.2.2"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionaction.tf_contentinspectionaction", "ifserverdown", "DROP"),
				),
			},
			{
				Config: testAccContentinspectionaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionactionExist("citrixadc_contentinspectionaction.tf_contentinspectionaction", nil),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionaction.tf_contentinspectionaction", "name", "my_ci_action"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionaction.tf_contentinspectionaction", "type", "NOINSPECTION"),
				),
			},
		},
	})
}

func TestAccContentinspectionaction_import(t *testing.T) {
	const resAddr = "citrixadc_contentinspectionaction.tf_contentinspectionaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckContentinspectionactionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccContentinspectionaction_basic},
			{
				Config:                  testAccContentinspectionaction_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckContentinspectionactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No contentinspectionaction name is set")
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
		data, err := client.FindResource("contentinspectionaction", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("contentinspectionaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckContentinspectionactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_contentinspectionaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("contentinspectionaction", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("contentinspectionaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccContentinspectionactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccContentinspectionactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_contentinspectionaction.tf_contentinspectionaction_ds", "name", "my_ci_action_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_contentinspectionaction.tf_contentinspectionaction_ds", "type", "ICAP"),
					resource.TestCheckResourceAttr("data.citrixadc_contentinspectionaction.tf_contentinspectionaction_ds", "icapprofilename", "new-profile-ds"),
					resource.TestCheckResourceAttr("data.citrixadc_contentinspectionaction.tf_contentinspectionaction_ds", "serverip", "3.3.3.3"),
					resource.TestCheckResourceAttr("data.citrixadc_contentinspectionaction.tf_contentinspectionaction_ds", "ifserverdown", "CONTINUE"),
					// Universal runtime-binding proof for the data source.
					resource.TestCheckResourceAttrSet("data.citrixadc_contentinspectionaction.tf_contentinspectionaction_ds", "id"),
				),
			},
		},
	})
}

func TestAccContentinspectionaction_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_contentinspectionaction.tf_contentinspectionaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckContentinspectionactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccContentinspectionaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckContentinspectionactionExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Contentinspectionaction.Type(), "my_ci_action"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccContentinspectionaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckContentinspectionactionExist(resAddr, nil)),
			},
		},
	})
}

func TestAccContentinspectionaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckContentinspectionactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccContentinspectionaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckContentinspectionactionExist("citrixadc_contentinspectionaction.tf_contentinspectionaction", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccContentinspectionaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckContentinspectionactionExist("citrixadc_contentinspectionaction.tf_contentinspectionaction", nil)),
			},
		},
	})
}

// The contentinspectionaction unset test covers the two spec-unsettable,
// mutable attributes (serverport, ifserverdown) on an ICAP action. Step 1 sets
// them to non-default values; step 2 removes them so the provider must issue a
// NITRO ?action=unset, reverting them to the documented defaults
// (serverport=1344, ifserverdown=RESET).
const testAccContentinspectionaction_unset_step1 = `
	resource "citrixadc_nsicapprofile" "tf_nsicapprofile_unset" {
		name             = "unset-profile"
		uri              = "/example"
		mode             = "REQMOD"
		reqtimeout       = 4
		reqtimeoutaction = "RESET"
		preview          = "ENABLED"
		previewlength    = 4096
	}
	resource "citrixadc_contentinspectionaction" "tf_unset" {
		name            = "tf_ci_action_unset"
		type            = "ICAP"
		icapprofilename = citrixadc_nsicapprofile.tf_nsicapprofile_unset.name
		serverip        = "2.2.2.2"
		serverport      = 2048
		ifserverdown    = "DROP"
	}
`

const testAccContentinspectionaction_unset_step2 = `
	resource "citrixadc_nsicapprofile" "tf_nsicapprofile_unset" {
		name             = "unset-profile"
		uri              = "/example"
		mode             = "REQMOD"
		reqtimeout       = 4
		reqtimeoutaction = "RESET"
		preview          = "ENABLED"
		previewlength    = 4096
	}
	resource "citrixadc_contentinspectionaction" "tf_unset" {
		name            = "tf_ci_action_unset"
		type            = "ICAP"
		icapprofilename = citrixadc_nsicapprofile.tf_nsicapprofile_unset.name
		serverip        = "2.2.2.2"
		# serverport and ifserverdown removed from config -> provider must unset
		# them (revert to NITRO defaults 1344 / RESET).
	}
`

func TestAccContentinspectionaction_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckContentinspectionactionDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccContentinspectionaction_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionactionExist("citrixadc_contentinspectionaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionaction.tf_unset", "serverport", "2048"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionaction.tf_unset", "ifserverdown", "DROP"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccContentinspectionaction_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckContentinspectionactionExist("citrixadc_contentinspectionaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionaction.tf_unset", "serverport", "1344"),
					resource.TestCheckResourceAttr("citrixadc_contentinspectionaction.tf_unset", "ifserverdown", "RESET"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckContentinspectionactionADCValue("tf_ci_action_unset", "serverport", "1344"),
					testAccCheckContentinspectionactionADCValue("tf_ci_action_unset", "ifserverdown", "RESET"),
				),
			},
		},
	})
}

// testAccCheckContentinspectionactionADCValue asserts an attribute's value
// directly on the appliance (not just in Terraform state), proving the unset
// actually reverted it.
func testAccCheckContentinspectionactionADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Contentinspectionaction.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("contentinspectionaction %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("contentinspectionaction %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

const testAccContentinspectionactionDataSource_basic = `

resource "citrixadc_nsicapprofile" "tf_nsicapprofile_ds" {
	name             = "new-profile-ds"
	uri              = "/example"
	mode             = "REQMOD"
	reqtimeout       = 4
	reqtimeoutaction = "RESET"
	preview          = "ENABLED"
	previewlength    = 4096
}

resource "citrixadc_contentinspectionaction" "tf_contentinspectionaction_ds" {
	name            = "my_ci_action_ds"
	type            = "ICAP"
	icapprofilename = citrixadc_nsicapprofile.tf_nsicapprofile_ds.name
	serverip        = "3.3.3.3"
	ifserverdown    = "CONTINUE"
}

data "citrixadc_contentinspectionaction" "tf_contentinspectionaction_ds" {
	name = citrixadc_contentinspectionaction.tf_contentinspectionaction_ds.name
	depends_on = [citrixadc_contentinspectionaction.tf_contentinspectionaction_ds]
}

`
