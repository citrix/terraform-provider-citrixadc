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

const testAccAppflowcollector_basic = `
resource "citrixadc_appflowcollector" "tf_appflowcollector" {
	name      = "tf_collector3"
	ipaddress = "192.168.2.3"
	transport = "logstream"
	port      =  80
	}
`
const testAccAppflowcollector_update = `
resource "citrixadc_appflowcollector" "tf_appflowcollector" {
	name      = "tf_collector3"
	ipaddress = "192.168.2.4"
	transport = "rest"
	port      = 90
	}
`

func TestAccAppflowcollector_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppflowcollectorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppflowcollector_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowcollectorExist("citrixadc_appflowcollector.tf_appflowcollector", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowcollector.tf_appflowcollector", "name", "tf_collector3"),
					resource.TestCheckResourceAttr("citrixadc_appflowcollector.tf_appflowcollector", "transport", "logstream"),
					resource.TestCheckResourceAttr("citrixadc_appflowcollector.tf_appflowcollector", "ipaddress", "192.168.2.3"),
					resource.TestCheckResourceAttr("citrixadc_appflowcollector.tf_appflowcollector", "port", "80"),
				),
			},
			{
				Config: testAccAppflowcollector_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowcollectorExist("citrixadc_appflowcollector.tf_appflowcollector", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowcollector.tf_appflowcollector", "name", "tf_collector3"),
					resource.TestCheckResourceAttr("citrixadc_appflowcollector.tf_appflowcollector", "transport", "rest"),
					resource.TestCheckResourceAttr("citrixadc_appflowcollector.tf_appflowcollector", "ipaddress", "192.168.2.4"),
					resource.TestCheckResourceAttr("citrixadc_appflowcollector.tf_appflowcollector", "port", "90"),
				),
			},
		},
	})
}

func testAccCheckAppflowcollectorExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appflowcollector name is set")
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
		data, err := client.FindResource(service.Appflowcollector.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appflowcollector %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppflowcollectorDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appflowcollector" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Appflowcollector.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appflowcollector %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAppflowcollectorDataSource_basic = `

resource "citrixadc_appflowcollector" "tf_appflowcollector" {
	name      = "tf_collector3"
	ipaddress = "192.168.2.3"
	transport = "logstream"
	port      =  80
}

data "citrixadc_appflowcollector" "tf_appflowcollector" {
	name = citrixadc_appflowcollector.tf_appflowcollector.name
}
`

func TestAccAppflowcollectorDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAppflowcollectorDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appflowcollector.tf_appflowcollector", "name", "tf_collector3"),
					resource.TestCheckResourceAttr("data.citrixadc_appflowcollector.tf_appflowcollector", "ipaddress", "192.168.2.3"),
					resource.TestCheckResourceAttr("data.citrixadc_appflowcollector.tf_appflowcollector", "transport", "logstream"),
					resource.TestCheckResourceAttr("data.citrixadc_appflowcollector.tf_appflowcollector", "port", "80"),
					// Universal runtime-binding proof for the data source.
					resource.TestCheckResourceAttrSet("data.citrixadc_appflowcollector.tf_appflowcollector", "id"),
				),
			},
		},
	})
}

func TestAccAppflowcollector_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_appflowcollector.tf_appflowcollector"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppflowcollectorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppflowcollector_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppflowcollectorExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Appflowcollector.Type(), "tf_collector3"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccAppflowcollector_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppflowcollectorExist(resAddr, nil)),
			},
		},
	})
}

// The appflowcollector unset test covers the sole reliably unsettable mutable
// attribute: port (reverts to the NITRO default 4739). ipaddress is mandatory
// at add time and cannot be unset ("Invalid IP address"); netprofile's unset is
// silently ignored by NITRO (value persists), so neither is wired/tested.
const testAccAppflowcollector_unset_step1 = `
resource "citrixadc_appflowcollector" "tf_unset" {
	name      = "tf_afc_unset"
	ipaddress = "192.168.9.9"
	port      = 5000
}
`

const testAccAppflowcollector_unset_step2 = `
resource "citrixadc_appflowcollector" "tf_unset" {
	name      = "tf_afc_unset"
	ipaddress = "192.168.9.9"
	# port removed from config -> the provider must unset it (revert to 4739).
}
`

func TestAccAppflowcollector_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppflowcollectorDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccAppflowcollector_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowcollectorExist("citrixadc_appflowcollector.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowcollector.tf_unset", "port", "5000"),
				),
			},
			{
				// Removing port must unset it: state (read back from the appliance)
				// reverts to the NITRO default, and the implicit post-apply plan
				// must be empty.
				Config: testAccAppflowcollector_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowcollectorExist("citrixadc_appflowcollector.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowcollector.tf_unset", "port", "4739"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAppflowcollectorADCValue("tf_afc_unset", "port", "4739"),
				),
			},
		},
	})
}

// testAccCheckAppflowcollectorADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckAppflowcollectorADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Appflowcollector.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("appflowcollector %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("appflowcollector %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccAppflowcollector_import(t *testing.T) {
	const resAddr = "citrixadc_appflowcollector.tf_appflowcollector"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppflowcollectorDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAppflowcollector_basic},
			{
				Config:                  testAccAppflowcollector_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccAppflowcollector_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAppflowcollectorDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccAppflowcollector_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppflowcollectorExist("citrixadc_appflowcollector.tf_appflowcollector", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccAppflowcollector_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppflowcollectorExist("citrixadc_appflowcollector.tf_appflowcollector", nil)),
			},
		},
	})
}
