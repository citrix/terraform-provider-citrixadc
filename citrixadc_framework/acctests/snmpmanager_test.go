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

const testAccSnmpmanager_basic = `

resource "citrixadc_snmpmanager" "tf_snmpmanager" {
	ipaddress          = "192.168.2.4"
	}
  	
`

const testAccSnmpmanager_update = `

resource "citrixadc_snmpmanager" "tf_snmpmanager" {
	ipaddress          = "192.168.2.4"
	netmask            = "255.255.255.252"
	}
	
`

func TestAccSnmpmanager_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmpmanagerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpmanager_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpmanagerExist("citrixadc_snmpmanager.tf_snmpmanager", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpmanager.tf_snmpmanager", "ipaddress", "192.168.2.4"),
					resource.TestCheckResourceAttr("citrixadc_snmpmanager.tf_snmpmanager", "netmask", "255.255.255.255"),
				),
			},
			{
				Config: testAccSnmpmanager_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpmanagerExist("citrixadc_snmpmanager.tf_snmpmanager", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpmanager.tf_snmpmanager", "ipaddress", "192.168.2.4"),
					resource.TestCheckResourceAttr("citrixadc_snmpmanager.tf_snmpmanager", "netmask", "255.255.255.252"),
				),
			},
		},
	})
}

func testAccCheckSnmpmanagerExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No snmpmanager name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		snmpmanagerName := rs.Primary.ID
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		dataArr, err := client.FindAllResources(service.Snmpmanager.Type())

		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if v["ipaddress"] == snmpmanagerName {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("snmpmanager %s not found", n)
		}

		return nil
	}
}

func testAccCheckSnmpmanagerDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_snmpmanager" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		snmpmanagerName := rs.Primary.ID

		dataArr, err := client.FindAllResources(service.Snmpmanager.Type())

		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if v["ipaddress"] == snmpmanagerName {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("snmpmanager %s still exists", snmpmanagerName)
		}

	}

	return nil
}

func TestAccSnmpmanager_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckSnmpmanagerDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccSnmpmanager_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSnmpmanagerExist("citrixadc_snmpmanager.tf_snmpmanager", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccSnmpmanager_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSnmpmanagerExist("citrixadc_snmpmanager.tf_snmpmanager", nil)),
			},
		},
	})
}

func TestAccSnmpmanagerDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmpmanagerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpmanagerDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_snmpmanager.tf_snmpmanager_ds", "ipaddress", "192.168.2.10"),
					resource.TestCheckResourceAttr("data.citrixadc_snmpmanager.tf_snmpmanager_ds", "netmask", "255.255.255.255"),
					// Universal runtime-binding proof for the data source.
					resource.TestCheckResourceAttrSet("data.citrixadc_snmpmanager.tf_snmpmanager_ds", "id"),
				),
			},
		},
	})
}

func TestAccSnmpmanager_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_snmpmanager.tf_snmpmanager"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmpmanagerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpmanager_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSnmpmanagerExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgs(service.Snmpmanager.Type(), "192.168.2.4", []string{"netmask:255.255.255.255"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccSnmpmanager_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSnmpmanagerExist(resAddr, nil)),
			},
		},
	})
}

// domainresolveretry is the only unset-eligible attribute: it is Optional and
// updateable (not the key, not ForceNew, not read-only) and has a documented
// NITRO default of 5. NITRO rejects it on IP-based managers ("Domain resolve
// retry cannot be given for IP based manager"), so the unset test uses a
// host-name based manager, which in turn requires a DNS name server.
const testAccSnmpmanager_unset_step1 = `
resource "citrixadc_dnsnameserver" "tf_unset_ns" {
  ip = "8.8.8.8"
}

resource "citrixadc_snmpmanager" "tf_unset" {
  ipaddress          = "tfunsethost.example.com"
  domainresolveretry = 10
  depends_on         = [citrixadc_dnsnameserver.tf_unset_ns]
}
`

const testAccSnmpmanager_unset_step2 = `
resource "citrixadc_dnsnameserver" "tf_unset_ns" {
  ip = "8.8.8.8"
}

resource "citrixadc_snmpmanager" "tf_unset" {
  ipaddress  = "tfunsethost.example.com"
  depends_on = [citrixadc_dnsnameserver.tf_unset_ns]
  # domainresolveretry removed from config -> the provider must unset it
  # (revert to the NITRO default of 5).
}
`

func TestAccSnmpmanager_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmpmanagerDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccSnmpmanager_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpmanagerExist("citrixadc_snmpmanager.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpmanager.tf_unset", "ipaddress", "tfunsethost.example.com"),
					resource.TestCheckResourceAttr("citrixadc_snmpmanager.tf_unset", "domainresolveretry", "10"),
					testAccCheckSnmpmanagerADCValue("tfunsethost.example.com", "domainresolveretry", "10"),
				),
			},
			{
				// Removing the attribute must unset it: state (read back from the
				// appliance) reverts to the documented NITRO default (5), and the
				// implicit post-apply plan must be empty.
				Config: testAccSnmpmanager_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpmanagerExist("citrixadc_snmpmanager.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpmanager.tf_unset", "domainresolveretry", "5"),
					testAccCheckSnmpmanagerADCValue("tfunsethost.example.com", "domainresolveretry", "5"),
				),
			},
		},
	})
}

// testAccCheckSnmpmanagerADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it. snmpmanager is an array resource keyed by ipaddress.
func testAccCheckSnmpmanagerADCValue(ipaddress, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		dataArr, err := client.FindAllResources(service.Snmpmanager.Type())
		if err != nil {
			return err
		}
		for _, v := range dataArr {
			if ip, ok := v["ipaddress"].(string); ok && ip == ipaddress {
				got := strings.TrimSpace(fmt.Sprintf("%v", v[attr]))
				if got != want {
					return fmt.Errorf("snmpmanager %s: appliance attr %q = %q, want %q (unset did not revert it)", ipaddress, attr, got, want)
				}
				return nil
			}
		}
		return fmt.Errorf("snmpmanager %s not found on appliance", ipaddress)
	}
}

func TestAccSnmpmanager_import(t *testing.T) {
	const resAddr = "citrixadc_snmpmanager.tf_snmpmanager"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmpmanagerDestroy,
		Steps: []resource.TestStep{
			{Config: testAccSnmpmanager_basic},
			{
				Config:                  testAccSnmpmanager_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccSnmpmanagerDataSource_basic = `

resource "citrixadc_snmpmanager" "tf_snmpmanager_ds" {
	ipaddress = "192.168.2.10"
	netmask   = "255.255.255.255"
}

data "citrixadc_snmpmanager" "tf_snmpmanager_ds" {
	ipaddress = citrixadc_snmpmanager.tf_snmpmanager_ds.ipaddress
	netmask   = citrixadc_snmpmanager.tf_snmpmanager_ds.netmask
	depends_on = [citrixadc_snmpmanager.tf_snmpmanager_ds]
}
`
