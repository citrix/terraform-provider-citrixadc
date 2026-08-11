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

const testAccTunneltrafficpolicy_basic = `


resource "citrixadc_tunneltrafficpolicy" "tf_tunneltrafficpolicy" {
	name   = "my_tunneltrafficpolicy"
	rule   = "true"
	action = "COMPRESS"
	}
  
`
const testAccTunneltrafficpolicy_update = `


resource "citrixadc_tunneltrafficpolicy" "tf_tunneltrafficpolicy" {
	name   = "my_tunneltrafficpolicy"
	rule   = "false"
	action = "NOCOMPRESS"
	}
  
`

func TestAccTunneltrafficpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTunneltrafficpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTunneltrafficpolicy_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTunneltrafficpolicyExist("citrixadc_tunneltrafficpolicy.tf_tunneltrafficpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_tunneltrafficpolicy.tf_tunneltrafficpolicy", "name", "my_tunneltrafficpolicy"),
					resource.TestCheckResourceAttr("citrixadc_tunneltrafficpolicy.tf_tunneltrafficpolicy", "rule", "true"),
					resource.TestCheckResourceAttr("citrixadc_tunneltrafficpolicy.tf_tunneltrafficpolicy", "action", "COMPRESS"),
				),
			},
			{
				Config: testAccTunneltrafficpolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTunneltrafficpolicyExist("citrixadc_tunneltrafficpolicy.tf_tunneltrafficpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_tunneltrafficpolicy.tf_tunneltrafficpolicy", "name", "my_tunneltrafficpolicy"),
					resource.TestCheckResourceAttr("citrixadc_tunneltrafficpolicy.tf_tunneltrafficpolicy", "rule", "false"),
					resource.TestCheckResourceAttr("citrixadc_tunneltrafficpolicy.tf_tunneltrafficpolicy", "action", "NOCOMPRESS"),
				),
			},
		},
	})
}

func TestAccTunneltrafficpolicy_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_tunneltrafficpolicy.tf_tunneltrafficpolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTunneltrafficpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTunneltrafficpolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTunneltrafficpolicyExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Tunneltrafficpolicy.Type(), "my_tunneltrafficpolicy"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccTunneltrafficpolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTunneltrafficpolicyExist(resAddr, nil)),
			},
		},
	})
}

func TestAccTunneltrafficpolicy_import(t *testing.T) {
	const resAddr = "citrixadc_tunneltrafficpolicy.tf_tunneltrafficpolicy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTunneltrafficpolicyDestroy,
		Steps: []resource.TestStep{
			{Config: testAccTunneltrafficpolicy_basic},
			{
				Config:                  testAccTunneltrafficpolicy_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckTunneltrafficpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No tunneltrafficpolicy name is set")
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
		data, err := client.FindResource(service.Tunneltrafficpolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("tunneltrafficpolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckTunneltrafficpolicyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_tunneltrafficpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Tunneltrafficpolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("tunneltrafficpolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccTunneltrafficpolicy_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckTunneltrafficpolicyDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccTunneltrafficpolicy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTunneltrafficpolicyExist("citrixadc_tunneltrafficpolicy.tf_tunneltrafficpolicy", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccTunneltrafficpolicy_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckTunneltrafficpolicyExist("citrixadc_tunneltrafficpolicy.tf_tunneltrafficpolicy", nil)),
			},
		},
	})
}

// tunneltrafficpolicy's unset-eligible attributes are comment and logaction,
// the two attributes NITRO's ?action=unset payload documents. Both are
// Optional+Computed and unset cleanly (revert to the empty NITRO default).
// action is not unsettable and rule/name are policy-core/key, so they are
// excluded. Setting logaction requires a prerequisite audit message action.
const testAccTunneltrafficpolicy_unset_step1 = `
resource "citrixadc_auditmessageaction" "tf_msgaction" {
	name              = "tf_unset_ttp_msgaction"
	loglevel          = "INFORMATIONAL"
	stringbuilderexpr = "\"tunneltrafficpolicy unset test\""
}
resource "citrixadc_tunneltrafficpolicy" "tf_unset" {
	name      = "tf_test_ttp_unset"
	rule      = "true"
	action    = "COMPRESS"
	comment   = "unset test comment"
	logaction = citrixadc_auditmessageaction.tf_msgaction.name
}
`

const testAccTunneltrafficpolicy_unset_step2 = `
resource "citrixadc_auditmessageaction" "tf_msgaction" {
	name              = "tf_unset_ttp_msgaction"
	loglevel          = "INFORMATIONAL"
	stringbuilderexpr = "\"tunneltrafficpolicy unset test\""
}
resource "citrixadc_tunneltrafficpolicy" "tf_unset" {
	name   = "tf_test_ttp_unset"
	rule   = "true"
	action = "COMPRESS"
	# comment and logaction removed from config -> provider must unset them.
}
`

func TestAccTunneltrafficpolicy_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTunneltrafficpolicyDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccTunneltrafficpolicy_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTunneltrafficpolicyExist("citrixadc_tunneltrafficpolicy.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_tunneltrafficpolicy.tf_unset", "comment", "unset test comment"),
					resource.TestCheckResourceAttr("citrixadc_tunneltrafficpolicy.tf_unset", "logaction", "tf_unset_ttp_msgaction"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from the
				// appliance) reverts to the empty NITRO default, the implicit
				// post-apply plan must be empty, and the appliance confirms it.
				Config: testAccTunneltrafficpolicy_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTunneltrafficpolicyExist("citrixadc_tunneltrafficpolicy.tf_unset", nil),
					// After unset, NITRO omits these from GET, so they read back as null
					// (absent from state) -- assert the revert directly on the appliance
					// instead, and rely on the implicit empty-plan check.
					testAccCheckTunneltrafficpolicyADCValue("tf_test_ttp_unset", "comment", ""),
					testAccCheckTunneltrafficpolicyADCValue("tf_test_ttp_unset", "logaction", ""),
				),
			},
		},
	})
}

// testAccCheckTunneltrafficpolicyADCValue asserts an attribute's value directly
// on the appliance (not just in Terraform state), proving the unset actually
// reverted it. An unset attribute is omitted from the GET response, which reads
// back as an empty string here.
func testAccCheckTunneltrafficpolicyADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Tunneltrafficpolicy.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("tunneltrafficpolicy %s not found on appliance", name)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("tunneltrafficpolicy %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccTunneltrafficpolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccTunneltrafficpolicyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_tunneltrafficpolicy.tf_tunneltrafficpolicy", "name", "my_tunneltrafficpolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_tunneltrafficpolicy.tf_tunneltrafficpolicy", "rule", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_tunneltrafficpolicy.tf_tunneltrafficpolicy", "action", "COMPRESS"),
				),
			},
		},
	})
}

const testAccTunneltrafficpolicyDataSource_basic = `
resource "citrixadc_tunneltrafficpolicy" "tf_tunneltrafficpolicy" {
	name   = "my_tunneltrafficpolicy"
	rule   = "true"
	action = "COMPRESS"
}

data "citrixadc_tunneltrafficpolicy" "tf_tunneltrafficpolicy" {
    name = citrixadc_tunneltrafficpolicy.tf_tunneltrafficpolicy.name
    depends_on = [citrixadc_tunneltrafficpolicy.tf_tunneltrafficpolicy]
}
`
