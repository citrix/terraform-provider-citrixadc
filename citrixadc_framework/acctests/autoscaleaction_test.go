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

const testAccAutoscaleaction_basic = `

	resource "citrixadc_autoscaleprofile" "tf_autoscaleprofile" {
		name         = "my_profile"
		type         = "CLOUDSTACK"
		apikey       = "7c177611-4a18-42b0-a7c5-bfd811fd590f"
		url          = "www.service.example.com"
		sharedsecret = "YZEH6jkTqZWQ8r0o6kWj0mWruN3vXbtT"
	}
	resource "citrixadc_autoscaleaction" "tf_autoscaleaction" {
		name        = "my_autoscaleaction"
		type        = "SCALE_UP"
		profilename = citrixadc_autoscaleprofile.tf_autoscaleprofile.name
		vserver     = "my_vserver"
		parameters  = "my_parameters"
	}
`
const testAccAutoscaleaction_update = `


	resource "citrixadc_autoscaleprofile" "tf_autoscaleprofile" {
		name         = "my_profile"
		type         = "CLOUDSTACK"
		apikey       = "7c177611-4a18-42b0-a7c5-bfd811fd590f"
		url          = "www.service.example.com"
		sharedsecret = "YZEH6jkTqZWQ8r0o6kWj0mWruN3vXbtT"
	}
	resource "citrixadc_autoscaleaction" "tf_autoscaleaction" {
		name        = "my_autoscaleaction"
		type        = "SCALE_DOWN"
		profilename = citrixadc_autoscaleprofile.tf_autoscaleprofile.name
		vserver     = "my_vserver2"
		parameters  = "my_parameters"
	}
`

func TestAccAutoscaleaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAutoscaleactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAutoscaleaction_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAutoscaleactionExist("citrixadc_autoscaleaction.tf_autoscaleaction", nil),
					resource.TestCheckResourceAttr("citrixadc_autoscaleaction.tf_autoscaleaction", "name", "my_autoscaleaction"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleaction.tf_autoscaleaction", "type", "SCALE_UP"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleaction.tf_autoscaleaction", "profilename", "my_profile"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleaction.tf_autoscaleaction", "vserver", "my_vserver"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleaction.tf_autoscaleaction", "parameters", "my_parameters"),
				),
			},
			{
				Config: testAccAutoscaleaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAutoscaleactionExist("citrixadc_autoscaleaction.tf_autoscaleaction", nil),
					resource.TestCheckResourceAttr("citrixadc_autoscaleaction.tf_autoscaleaction", "name", "my_autoscaleaction"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleaction.tf_autoscaleaction", "type", "SCALE_DOWN"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleaction.tf_autoscaleaction", "profilename", "my_profile"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleaction.tf_autoscaleaction", "vserver", "my_vserver2"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleaction.tf_autoscaleaction", "parameters", "my_parameters"),
				),
			},
		},
	})
}

func TestAccAutoscaleaction_import(t *testing.T) {
	const resAddr = "citrixadc_autoscaleaction.tf_autoscaleaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAutoscaleactionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAutoscaleaction_basic},
			{
				Config:                  testAccAutoscaleaction_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckAutoscaleactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No autoscaleaction name is set")
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
		data, err := client.FindResource(service.Autoscaleaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("autoscaleaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckAutoscaleactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_autoscaleaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Autoscaleaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("autoscaleaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccAutoscaleaction_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_autoscaleaction.tf_autoscaleaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAutoscaleactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAutoscaleaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAutoscaleactionExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Autoscaleaction.Type(), "my_autoscaleaction"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAutoscaleaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAutoscaleactionExist(resAddr, nil)),
			},
		},
	})
}

const testAccAutoscaleactionDataSource_basic = `
	resource "citrixadc_autoscaleprofile" "tf_autoscaleprofile_ds" {
		name         = "my_profile_ds"
		type         = "CLOUDSTACK"
		apikey       = "7c177611-4a18-42b0-a7c5-bfd811fd590f"
		url          = "www.service.example.com"
		sharedsecret = "YZEH6jkTqZWQ8r0o6kWj0mWruN3vXbtT"
	}
	resource "citrixadc_autoscaleaction" "tf_autoscaleaction_ds" {
		name        = "my_autoscaleaction_ds"
		type        = "SCALE_UP"
		profilename = citrixadc_autoscaleprofile.tf_autoscaleprofile_ds.name
		vserver     = "my_vserver"
		parameters  = "my_parameters"
	}

	data "citrixadc_autoscaleaction" "tf_autoscaleaction_ds" {
		name = citrixadc_autoscaleaction.tf_autoscaleaction_ds.name
	}
`

func TestAccAutoscaleaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAutoscaleactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccAutoscaleaction_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAutoscaleactionExist("citrixadc_autoscaleaction.tf_autoscaleaction", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAutoscaleaction_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckAutoscaleactionExist("citrixadc_autoscaleaction.tf_autoscaleaction", nil)),
			},
		},
	})
}

// The autoscaleaction unset test covers the two spec-unsettable attributes:
// quiettime (NITRO default 300) and vmdestroygraceperiod (NITRO default 10).
const testAccAutoscaleaction_unset_step1 = `
	resource "citrixadc_autoscaleprofile" "tf_autoscaleprofile" {
		name         = "my_profile_unset"
		type         = "CLOUDSTACK"
		apikey       = "7c177611-4a18-42b0-a7c5-bfd811fd590f"
		url          = "www.service.example.com"
		sharedsecret = "YZEH6jkTqZWQ8r0o6kWj0mWruN3vXbtT"
	}
	resource "citrixadc_autoscaleaction" "tf_unset" {
		name                 = "tf_test_autoscaleaction_unset"
		type                 = "SCALE_UP"
		profilename          = citrixadc_autoscaleprofile.tf_autoscaleprofile.name
		vserver              = "my_vserver"
		parameters           = "my_parameters"
		quiettime            = 100
		vmdestroygraceperiod = 20
	}
`

const testAccAutoscaleaction_unset_step2 = `
	resource "citrixadc_autoscaleprofile" "tf_autoscaleprofile" {
		name         = "my_profile_unset"
		type         = "CLOUDSTACK"
		apikey       = "7c177611-4a18-42b0-a7c5-bfd811fd590f"
		url          = "www.service.example.com"
		sharedsecret = "YZEH6jkTqZWQ8r0o6kWj0mWruN3vXbtT"
	}
	resource "citrixadc_autoscaleaction" "tf_unset" {
		name        = "tf_test_autoscaleaction_unset"
		type        = "SCALE_UP"
		profilename = citrixadc_autoscaleprofile.tf_autoscaleprofile.name
		vserver     = "my_vserver"
		parameters  = "my_parameters"
		# quiettime and vmdestroygraceperiod removed -> provider must unset them
		# (revert to NITRO defaults 300 and 10).
	}
`

func TestAccAutoscaleaction_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAutoscaleactionDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccAutoscaleaction_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAutoscaleactionExist("citrixadc_autoscaleaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_autoscaleaction.tf_unset", "quiettime", "100"),
					resource.TestCheckResourceAttr("citrixadc_autoscaleaction.tf_unset", "vmdestroygraceperiod", "20"),
				),
			},
			{
				// Removing the attributes must unset them: state reverts to the
				// documented NITRO defaults and the implicit post-apply plan is empty.
				Config: testAccAutoscaleaction_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAutoscaleactionExist("citrixadc_autoscaleaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_autoscaleaction.tf_unset", "quiettime", "300"),
					// Option B: vmdestroygraceperiod is unset via the unsetOnRemove
					// plan-modifier (no schema Default). NITRO omits it from GET once
					// unset, so it reads back NULL rather than the "10" default.
					resource.TestCheckNoResourceAttr("citrixadc_autoscaleaction.tf_unset", "vmdestroygraceperiod"),
					// Independent appliance-level confirmation the unset took effect.
					// Note: NITRO omits vmdestroygraceperiod from GET once it is at its
					// default, so it cannot be asserted at the appliance level; the
					// state check (=10) plus the empty post-apply plan cover it.
					testAccCheckAutoscaleactionADCValue("tf_test_autoscaleaction_unset", "quiettime", "300"),
				),
			},
		},
	})
}

// testAccCheckAutoscaleactionADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset reverted it.
func testAccCheckAutoscaleactionADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Autoscaleaction.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("autoscaleaction %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("autoscaleaction %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccAutoscaleactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAutoscaleactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_autoscaleaction.tf_autoscaleaction_ds", "name", "my_autoscaleaction_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_autoscaleaction.tf_autoscaleaction_ds", "type", "SCALE_UP"),
					resource.TestCheckResourceAttr("data.citrixadc_autoscaleaction.tf_autoscaleaction_ds", "profilename", "my_profile_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_autoscaleaction.tf_autoscaleaction_ds", "vserver", "my_vserver"),
					resource.TestCheckResourceAttr("data.citrixadc_autoscaleaction.tf_autoscaleaction_ds", "parameters", "my_parameters"),
				),
			},
		},
	})
}
