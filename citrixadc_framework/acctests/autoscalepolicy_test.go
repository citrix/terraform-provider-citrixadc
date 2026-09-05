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

const testAccAutoscalepolicy_basic = `


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
resource "citrixadc_autoscalepolicy" "tf_autoscalepolicy" {
	name         = "my_autoscaleprofile"
	rule         = "true"
	action       = citrixadc_autoscaleaction.tf_autoscaleaction.name
	}
`
const testAccAutoscalepolicy_update = `

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
resource "citrixadc_autoscalepolicy" "tf_autoscalepolicy" {
	name         = "my_autoscaleprofile"
	rule         = "false"
	action       = citrixadc_autoscaleaction.tf_autoscaleaction.name
	}
`

const testAccAutoscalepolicyDataSource_basic = `

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
resource "citrixadc_autoscalepolicy" "tf_autoscalepolicy" {
	name         = "my_autoscalepolicy"
	rule         = "true"
	action       = citrixadc_autoscaleaction.tf_autoscaleaction.name
}

data "citrixadc_autoscalepolicy" "tf_autoscalepolicy" {
	name = citrixadc_autoscalepolicy.tf_autoscalepolicy.name
}
`

func TestAccAutoscalepolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAutoscalepolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAutoscalepolicy_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAutoscalepolicyExist("citrixadc_autoscalepolicy.tf_autoscalepolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_autoscalepolicy.tf_autoscalepolicy", "name", "my_autoscaleprofile"),
					resource.TestCheckResourceAttr("citrixadc_autoscalepolicy.tf_autoscalepolicy", "rule", "true"),
					resource.TestCheckResourceAttr("citrixadc_autoscalepolicy.tf_autoscalepolicy", "action", "my_autoscaleaction"),
				),
			},
			{
				Config: testAccAutoscalepolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAutoscalepolicyExist("citrixadc_autoscalepolicy.tf_autoscalepolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_autoscalepolicy.tf_autoscalepolicy", "name", "my_autoscaleprofile"),
					resource.TestCheckResourceAttr("citrixadc_autoscalepolicy.tf_autoscalepolicy", "rule", "false"),
					resource.TestCheckResourceAttr("citrixadc_autoscalepolicy.tf_autoscalepolicy", "action", "my_autoscaleaction"),
				),
			},
		},
	})
}

func TestAccAutoscalepolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAutoscalepolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.citrixadc_autoscalepolicy.tf_autoscalepolicy", "name", "citrixadc_autoscalepolicy.tf_autoscalepolicy", "name"),
					resource.TestCheckResourceAttr("data.citrixadc_autoscalepolicy.tf_autoscalepolicy", "rule", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_autoscalepolicy.tf_autoscalepolicy", "action", "my_autoscaleaction"),
					// Universal runtime-binding proof.
					resource.TestCheckResourceAttrSet("data.citrixadc_autoscalepolicy.tf_autoscalepolicy", "id"),
					// Read-only counter-style attribute always populated by the appliance.
					resource.TestCheckResourceAttrSet("data.citrixadc_autoscalepolicy.tf_autoscalepolicy", "hits"),
				),
			},
		},
	})
}

func TestAccAutoscalepolicy_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_autoscalepolicy.tf_autoscalepolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAutoscalepolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAutoscalepolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAutoscalepolicyExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Autoscalepolicy.Type(), "my_autoscaleprofile"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAutoscalepolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAutoscalepolicyExist(resAddr, nil)),
			},
		},
	})
}

func TestAccAutoscalepolicy_import(t *testing.T) {
	const resAddr = "citrixadc_autoscalepolicy.tf_autoscalepolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAutoscalepolicyDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAutoscalepolicy_basic},
			{
				Config:                  testAccAutoscalepolicy_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccAutoscalepolicy_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAutoscalepolicyDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccAutoscalepolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAutoscalepolicyExist("citrixadc_autoscalepolicy.tf_autoscalepolicy", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccAutoscalepolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAutoscalepolicyExist("citrixadc_autoscalepolicy.tf_autoscalepolicy", nil)),
			},
		},
	})
}

// The autoscalepolicy unset test exercises the two mutable, non-mandatory
// attributes the NITRO spec lists as unsettable: comment and logaction. (rule
// and action are mandatory create arguments with no server default, so they are
// not unset-eligible.) Step 1 sets both to non-default values; step 2 removes
// them from config, so the provider must issue ?action=unset and the appliance
// must revert them to their (empty) defaults.
const testAccAutoscalepolicy_unset_step1 = `
resource "citrixadc_autoscaleprofile" "tf_autoscaleprofile" {
	name         = "tf_unset_profile"
	type         = "CLOUDSTACK"
	apikey       = "7c177611-4a18-42b0-a7c5-bfd811fd590f"
	url          = "www.service.example.com"
	sharedsecret = "YZEH6jkTqZWQ8r0o6kWj0mWruN3vXbtT"
}
resource "citrixadc_autoscaleaction" "tf_autoscaleaction" {
	name        = "tf_unset_action"
	type        = "SCALE_UP"
	profilename = citrixadc_autoscaleprofile.tf_autoscaleprofile.name
	vserver     = "my_vserver"
	parameters  = "my_parameters"
}
resource "citrixadc_auditmessageaction" "tf_msgaction" {
	name              = "tf_unset_msgaction"
	loglevel          = "INFORMATIONAL"
	stringbuilderexpr = "\"autoscale unset test\""
}
resource "citrixadc_autoscalepolicy" "tf_unset" {
	name      = "tf_test_autoscalepolicy_unset"
	rule      = "true"
	action    = citrixadc_autoscaleaction.tf_autoscaleaction.name
	comment   = "unset_comment"
	logaction = citrixadc_auditmessageaction.tf_msgaction.name
}
`

const testAccAutoscalepolicy_unset_step2 = `
resource "citrixadc_autoscaleprofile" "tf_autoscaleprofile" {
	name         = "tf_unset_profile"
	type         = "CLOUDSTACK"
	apikey       = "7c177611-4a18-42b0-a7c5-bfd811fd590f"
	url          = "www.service.example.com"
	sharedsecret = "YZEH6jkTqZWQ8r0o6kWj0mWruN3vXbtT"
}
resource "citrixadc_autoscaleaction" "tf_autoscaleaction" {
	name        = "tf_unset_action"
	type        = "SCALE_UP"
	profilename = citrixadc_autoscaleprofile.tf_autoscaleprofile.name
	vserver     = "my_vserver"
	parameters  = "my_parameters"
}
resource "citrixadc_auditmessageaction" "tf_msgaction" {
	name              = "tf_unset_msgaction"
	loglevel          = "INFORMATIONAL"
	stringbuilderexpr = "\"autoscale unset test\""
}
resource "citrixadc_autoscalepolicy" "tf_unset" {
	name   = "tf_test_autoscalepolicy_unset"
	rule   = "true"
	action = citrixadc_autoscaleaction.tf_autoscaleaction.name
	# comment and logaction removed from config -> provider must unset them.
}
`

func TestAccAutoscalepolicy_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAutoscalepolicyDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccAutoscalepolicy_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAutoscalepolicyExist("citrixadc_autoscalepolicy.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_autoscalepolicy.tf_unset", "comment", "unset_comment"),
					resource.TestCheckResourceAttr("citrixadc_autoscalepolicy.tf_unset", "logaction", "tf_unset_msgaction"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to the empty NITRO defaults, the implicit
				// post-apply plan must be empty, and the appliance confirms it.
				Config: testAccAutoscalepolicy_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAutoscalepolicyExist("citrixadc_autoscalepolicy.tf_unset", nil),
					// After unset, NITRO omits comment/logaction from GET, so they read
					// back as null (absent from state) -- assert the revert directly on
					// the appliance instead, and rely on the implicit empty-plan check.
					testAccCheckAutoscalepolicyADCValue("tf_test_autoscalepolicy_unset", "comment", ""),
					testAccCheckAutoscalepolicyADCValue("tf_test_autoscalepolicy_unset", "logaction", ""),
				),
			},
		},
	})
}

// testAccCheckAutoscalepolicyADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it. An unset attribute is omitted from the GET response, which reads
// back as an empty string here.
func testAccCheckAutoscalepolicyADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Autoscalepolicy.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("autoscalepolicy %s not found on appliance", name)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("autoscalepolicy %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func testAccCheckAutoscalepolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No autoscalepolicy name is set")
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
		data, err := client.FindResource(service.Autoscalepolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("autoscalepolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckAutoscalepolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_autoscalepolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Autoscalepolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("autoscalepolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
