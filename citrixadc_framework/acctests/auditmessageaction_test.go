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

func TestAccAuditmessageaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuditmessageactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuditmessageaction_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditmessageactionExist("citrixadc_auditmessageaction.tf_msgaction", nil),
				),
			},
			{
				Config: testAccAuditmessageaction_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditmessageactionExist("citrixadc_auditmessageaction.tf_msgaction", nil),
				),
			},
			{
				Config: testAccAuditmessageaction_basic_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditmessageactionExist("citrixadc_auditmessageaction.tf_msgaction", nil),
				),
			},
		},
	})
}

func testAccCheckAuditmessageactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
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
		data, err := client.FindResource(service.Auditmessageaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Audit message action %s not found", n)
		}

		return nil
	}
}

func testAccCheckAuditmessageactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_auditmessageaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Auditmessageaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAuditmessageaction_basic_step1 = `

resource "citrixadc_auditmessageaction" "tf_msgaction" {
    name = "tf_msgaction"
    loglevel = "NOTICE"
    stringbuilderexpr = "\"hello\""
    logtonewnslog = "YES"
}

`

const testAccAuditmessageaction_basic_step2 = `

resource "citrixadc_auditmessageaction" "tf_msgaction" {
    name = "tf_msgaction"
    loglevel = "DEBUG"
    stringbuilderexpr = "\"hello and bye\""
    logtonewnslog = "NO"
}

`

const testAccAuditmessageaction_basic_step3 = `

resource "citrixadc_auditmessageaction" "tf_msgaction" {
    name = "tf_msgaction2"
    loglevel = "NOTICE"
    stringbuilderexpr = "\"hello\""
    logtonewnslog = "YES"
}

`

func TestAccAuditmessageaction_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_auditmessageaction.tf_msgaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuditmessageactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAuditmessageaction_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuditmessageactionExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Auditmessageaction.Type(), "tf_msgaction"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAuditmessageaction_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAuditmessageactionExist(resAddr, nil)),
			},
		},
	})
}

func TestAccAuditmessageaction_import(t *testing.T) {
	const resAddr = "citrixadc_auditmessageaction.tf_msgaction"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuditmessageactionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAuditmessageaction_basic_step1},
			{
				Config:                  testAccAuditmessageaction_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccAuditmessageaction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAuditmessageactionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccAuditmessageaction_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditmessageactionExist("citrixadc_auditmessageaction.tf_msgaction", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAuditmessageaction_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditmessageactionExist("citrixadc_auditmessageaction.tf_msgaction", nil),
				),
			},
		},
	})
}

// The auditmessageaction unset test covers logtonewnslog, the only
// spec-unsettable attribute that round-trips through GET. bypasssafetycheck is
// also listed in the NITRO unset payload but is never echoed back by GET (even
// when set), so it cannot be asserted or safely defaulted and is excluded.
const testAccAuditmessageaction_unset_step1 = `
resource "citrixadc_auditmessageaction" "tf_unset" {
    name              = "tf_test_amsgaction_unset"
    loglevel          = "NOTICE"
    stringbuilderexpr = "\"hello\""
    logtonewnslog     = "YES"
}
`

const testAccAuditmessageaction_unset_step2 = `
resource "citrixadc_auditmessageaction" "tf_unset" {
    name              = "tf_test_amsgaction_unset"
    loglevel          = "NOTICE"
    stringbuilderexpr = "\"hello\""
    # logtonewnslog removed from config -> provider must unset it (revert to "NO").
}
`

func TestAccAuditmessageaction_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuditmessageactionDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value applied and persisted.
				Config: testAccAuditmessageaction_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditmessageactionExist("citrixadc_auditmessageaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_auditmessageaction.tf_unset", "logtonewnslog", "YES"),
				),
			},
			{
				// Removing the attribute must unset it: state (read back from the
				// appliance) reverts to the NITRO default "NO", and the implicit
				// post-apply plan must be empty.
				Config: testAccAuditmessageaction_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuditmessageactionExist("citrixadc_auditmessageaction.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_auditmessageaction.tf_unset", "logtonewnslog", "NO"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAuditmessageactionADCValue("tf_test_amsgaction_unset", "logtonewnslog", "NO"),
				),
			},
		},
	})
}

// testAccCheckAuditmessageactionADCValue asserts an attribute's value directly
// on the appliance (not just in Terraform state), proving the unset reverted it.
func testAccCheckAuditmessageactionADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Auditmessageaction.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("auditmessageaction %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("auditmessageaction %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccAuditmessageactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAuditmessageactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_auditmessageaction.tf_msgaction", "name", "tf_msgaction_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_auditmessageaction.tf_msgaction", "loglevel", "NOTICE"),
					resource.TestCheckResourceAttr("data.citrixadc_auditmessageaction.tf_msgaction", "stringbuilderexpr", "\"hello from datasource\""),
					resource.TestCheckResourceAttr("data.citrixadc_auditmessageaction.tf_msgaction", "logtonewnslog", "YES"),
				),
			},
		},
	})
}

const testAccAuditmessageactionDataSource_basic = `

resource "citrixadc_auditmessageaction" "tf_msgaction" {
    name = "tf_msgaction_ds"
    loglevel = "NOTICE"
    stringbuilderexpr = "\"hello from datasource\""
    logtonewnslog = "YES"
}

data "citrixadc_auditmessageaction" "tf_msgaction" {
    name = citrixadc_auditmessageaction.tf_msgaction.name
    depends_on = [citrixadc_auditmessageaction.tf_msgaction]
}

`
