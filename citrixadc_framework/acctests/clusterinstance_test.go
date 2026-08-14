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

const testAccClusterinstance_basic = `

resource "citrixadc_clusterinstance" "tf_clusterinstance" {
	clid          = 1
	deadinterval  = 5
	hellointerval = 600
	}
  
`
const testAccClusterinstance_update = `

resource "citrixadc_clusterinstance" "tf_clusterinstance" {
	clid          = 1
	deadinterval  = 8
	hellointerval = 400
	}
  
`

func TestAccClusterinstance_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckClusterinstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClusterinstance_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterinstanceExist("citrixadc_clusterinstance.tf_clusterinstance", nil),
					resource.TestCheckResourceAttr("citrixadc_clusterinstance.tf_clusterinstance", "deadinterval", "5"),
					resource.TestCheckResourceAttr("citrixadc_clusterinstance.tf_clusterinstance", "hellointerval", "600"),
				),
			},
			{
				Config: testAccClusterinstance_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterinstanceExist("citrixadc_clusterinstance.tf_clusterinstance", nil),
					resource.TestCheckResourceAttr("citrixadc_clusterinstance.tf_clusterinstance", "deadinterval", "8"),
					resource.TestCheckResourceAttr("citrixadc_clusterinstance.tf_clusterinstance", "hellointerval", "400"),
				),
			},
		},
	})
}

func TestAccClusterinstance_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_clusterinstance.tf_clusterinstance"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckClusterinstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClusterinstance_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckClusterinstanceExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Clusterinstance.Type(), "1"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccClusterinstance_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckClusterinstanceExist(resAddr, nil)),
			},
		},
	})
}

func TestAccClusterinstance_import(t *testing.T) {
	const resAddr = "citrixadc_clusterinstance.tf_clusterinstance"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckClusterinstanceDestroy,
		Steps: []resource.TestStep{
			{Config: testAccClusterinstance_basic},
			{
				Config:                  testAccClusterinstance_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckClusterinstanceExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No clusterinstance name is set")
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
		data, err := client.FindResource(service.Clusterinstance.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("clusterinstance %s not found", n)
		}

		return nil
	}
}

func testAccCheckClusterinstanceDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_clusterinstance" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Clusterinstance.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("clusterinstance %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccClusterinstance_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckClusterinstanceDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccClusterinstance_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckClusterinstanceExist("citrixadc_clusterinstance.tf_clusterinstance", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccClusterinstance_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckClusterinstanceExist("citrixadc_clusterinstance.tf_clusterinstance", nil)),
			},
		},
	})
}

const testAccClusterinstance_unset_step1 = `
resource "citrixadc_clusterinstance" "tf_unset" {
	clid                       = 1
	deadinterval               = 5
	hellointerval              = 400
	preemption                 = "ENABLED"
	quorumtype                 = "NONE"
	processlocal               = "ENABLED"
	retainconnectionsoncluster = "YES"
	backplanebasedview         = "ENABLED"
	syncstatusstrictmode       = "ENABLED"
	dfdretainl2params          = "ENABLED"
	clusterproxyarp            = "DISABLED"
	secureheartbeats           = "ENABLED"
}
`

const testAccClusterinstance_unset_step2 = `
resource "citrixadc_clusterinstance" "tf_unset" {
	clid = 1
	# All unset-eligible attributes removed from config -> the provider must
	# unset them (revert to their documented NITRO defaults).
}
`

func TestAccClusterinstance_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckClusterinstanceDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccClusterinstance_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterinstanceExist("citrixadc_clusterinstance.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_clusterinstance.tf_unset", "deadinterval", "5"),
					resource.TestCheckResourceAttr("citrixadc_clusterinstance.tf_unset", "hellointerval", "400"),
					resource.TestCheckResourceAttr("citrixadc_clusterinstance.tf_unset", "preemption", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_clusterinstance.tf_unset", "quorumtype", "NONE"),
					resource.TestCheckResourceAttr("citrixadc_clusterinstance.tf_unset", "processlocal", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_clusterinstance.tf_unset", "retainconnectionsoncluster", "YES"),
					resource.TestCheckResourceAttr("citrixadc_clusterinstance.tf_unset", "backplanebasedview", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_clusterinstance.tf_unset", "syncstatusstrictmode", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_clusterinstance.tf_unset", "dfdretainl2params", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_clusterinstance.tf_unset", "clusterproxyarp", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_clusterinstance.tf_unset", "secureheartbeats", "ENABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccClusterinstance_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusterinstanceExist("citrixadc_clusterinstance.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_clusterinstance.tf_unset", "deadinterval", "3"),
					resource.TestCheckResourceAttr("citrixadc_clusterinstance.tf_unset", "hellointerval", "200"),
					resource.TestCheckResourceAttr("citrixadc_clusterinstance.tf_unset", "preemption", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_clusterinstance.tf_unset", "quorumtype", "MAJORITY"),
					resource.TestCheckResourceAttr("citrixadc_clusterinstance.tf_unset", "processlocal", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_clusterinstance.tf_unset", "retainconnectionsoncluster", "NO"),
					resource.TestCheckResourceAttr("citrixadc_clusterinstance.tf_unset", "backplanebasedview", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_clusterinstance.tf_unset", "syncstatusstrictmode", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_clusterinstance.tf_unset", "dfdretainl2params", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_clusterinstance.tf_unset", "clusterproxyarp", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_clusterinstance.tf_unset", "secureheartbeats", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckClusterinstanceADCValue("1", "quorumtype", "MAJORITY"),
					testAccCheckClusterinstanceADCValue("1", "deadinterval", "3"),
					testAccCheckClusterinstanceADCValue("1", "clusterproxyarp", "ENABLED"),
				),
			},
		},
	})
}

// testAccCheckClusterinstanceADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckClusterinstanceADCValue(id, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Clusterinstance.Type(), id)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("clusterinstance %s not found on appliance", id)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("clusterinstance %s: appliance attr %q = %q, want %q (unset did not revert it)", id, attr, got, want)
		}
		return nil
	}
}

func TestAccClusterinstanceDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccClusterinstanceDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_clusterinstance.tf_clusterinstance_ds", "clid", "1"),
					resource.TestCheckResourceAttr("data.citrixadc_clusterinstance.tf_clusterinstance_ds", "deadinterval", "5"),
					resource.TestCheckResourceAttr("data.citrixadc_clusterinstance.tf_clusterinstance_ds", "hellointerval", "600"),
				),
			},
		},
	})
}

const testAccClusterinstanceDataSource_basic = `

resource "citrixadc_clusterinstance" "tf_clusterinstance_ds" {
	clid          = 1
	deadinterval  = 5
	hellointerval = 600
}

data "citrixadc_clusterinstance" "tf_clusterinstance_ds" {
	clid = citrixadc_clusterinstance.tf_clusterinstance_ds.clid
	depends_on = [citrixadc_clusterinstance.tf_clusterinstance_ds]
}

`
