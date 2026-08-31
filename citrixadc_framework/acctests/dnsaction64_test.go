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

const testAccDnsaction64_add = `


resource "citrixadc_dnsaction64" "dnsaction64" {
	actionname = "default_DNS64_action1"
    prefix = "64:ff9c::/96"
    mappedrule = "DNS.RR.RDATA.IP.IN_SUBNET(10.0.0.0/8)"
    excluderule = "DNS.RR.RDATA.IPV6.IN_SUBNET(::ffff:0:0/96)"
}

`

const testAccDnsaction64_update = `

resource "citrixadc_dnsaction64" "dnsaction64" {
	actionname = "default_DNS64_action1"
    prefix = "64:ff9c::/96"
    mappedrule = "DNS.RR.TYPE.EQ(A)"
    excluderule = "DNS.RR.TYPE.EQ(AAAA)"
}

`

const testAccDnsaction64DataSource_basic = `

resource "citrixadc_dnsaction64" "dnsaction64" {
	actionname = "default_DNS64_action1"
    prefix = "64:ff9c::/96"
    mappedrule = "DNS.RR.RDATA.IP.IN_SUBNET(10.0.0.0/8)"
    excluderule = "DNS.RR.RDATA.IPV6.IN_SUBNET(::ffff:0:0/96)"
}

data "citrixadc_dnsaction64" "dnsaction64_datasource" {
    actionname = citrixadc_dnsaction64.dnsaction64.actionname
}

`

func TestAccDnsaction64_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnsaction64Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsaction64_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsaction64Exist("citrixadc_dnsaction64.dnsaction64", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsaction64.dnsaction64", "actionname", "default_DNS64_action1"),
					resource.TestCheckResourceAttr("citrixadc_dnsaction64.dnsaction64", "prefix", "64:ff9c::/96"),
					resource.TestCheckResourceAttr("citrixadc_dnsaction64.dnsaction64", "mappedrule", "DNS.RR.RDATA.IP.IN_SUBNET(10.0.0.0/8)"),
					resource.TestCheckResourceAttr("citrixadc_dnsaction64.dnsaction64", "excluderule", "DNS.RR.RDATA.IPV6.IN_SUBNET(::ffff:0:0/96)"),
				),
			},
			{
				Config: testAccDnsaction64_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsaction64Exist("citrixadc_dnsaction64.dnsaction64", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsaction64.dnsaction64", "actionname", "default_DNS64_action1"),
					resource.TestCheckResourceAttr("citrixadc_dnsaction64.dnsaction64", "prefix", "64:ff9c::/96"),
					resource.TestCheckResourceAttr("citrixadc_dnsaction64.dnsaction64", "mappedrule", "DNS.RR.TYPE.EQ(A)"),
					resource.TestCheckResourceAttr("citrixadc_dnsaction64.dnsaction64", "excluderule", "DNS.RR.TYPE.EQ(AAAA)"),
				),
			},
		},
	})
}

func testAccCheckDnsaction64Exist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnsaction64 name is set")
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
		data, err := client.FindResource(service.Dnsaction64.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("dnsaction64 %s not found", n)
		}

		return nil
	}
}

func testAccCheckDnsaction64Destroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnsaction64" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Dnsaction64.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("dnsaction64 %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccDnsaction64_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_dnsaction64.dnsaction64"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnsaction64Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsaction64_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnsaction64Exist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Dnsaction64.Type(), "default_DNS64_action1"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccDnsaction64_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnsaction64Exist(resAddr, nil)),
			},
		},
	})
}

func TestAccDnsaction64_import(t *testing.T) {
	const resAddr = "citrixadc_dnsaction64.dnsaction64"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnsaction64Destroy,
		Steps: []resource.TestStep{
			{Config: testAccDnsaction64_add},
			{
				Config:                  testAccDnsaction64_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccDnsaction64_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDnsaction64Destroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccDnsaction64_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnsaction64Exist("citrixadc_dnsaction64.dnsaction64", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccDnsaction64_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnsaction64Exist("citrixadc_dnsaction64.dnsaction64", nil)),
			},
		},
	})
}

// dnsaction64 unset test: mappedrule and excluderule are Optional+Computed
// expression attributes. NITRO omits them from GET once unset (revert-to-empty),
// so removing them from config must trigger the provider's unset flow.
const testAccDnsaction64_unset_step1 = `
resource "citrixadc_dnsaction64" "tf_unset" {
	actionname  = "tf_test_dnsaction64_unset"
	prefix      = "64:ff9b::/96"
	mappedrule  = "DNS.RR.TYPE.EQ(A)"
	excluderule = "DNS.RR.TYPE.EQ(AAAA)"
}
`

const testAccDnsaction64_unset_step2 = `
resource "citrixadc_dnsaction64" "tf_unset" {
	actionname = "tf_test_dnsaction64_unset"
	prefix     = "64:ff9b::/96"
	# mappedrule and excluderule removed from config -> the provider must unset
	# them (NITRO reverts them to empty / omits them from GET).
}
`

func TestAccDnsaction64_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnsaction64Destroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccDnsaction64_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsaction64Exist("citrixadc_dnsaction64.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsaction64.tf_unset", "mappedrule", "DNS.RR.TYPE.EQ(A)"),
					resource.TestCheckResourceAttr("citrixadc_dnsaction64.tf_unset", "excluderule", "DNS.RR.TYPE.EQ(AAAA)"),
				),
			},
			{
				// Removing the attributes must unset them: state reverts to the
				// NITRO default (empty) and the implicit post-apply plan is empty.
				Config: testAccDnsaction64_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsaction64Exist("citrixadc_dnsaction64.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsaction64.tf_unset", "mappedrule", ""),
					resource.TestCheckResourceAttr("citrixadc_dnsaction64.tf_unset", "excluderule", ""),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckDnsaction64ADCValue("tf_test_dnsaction64_unset", "mappedrule", ""),
					testAccCheckDnsaction64ADCValue("tf_test_dnsaction64_unset", "excluderule", ""),
				),
			},
		},
	})
}

// testAccCheckDnsaction64ADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted
// it. NITRO omits unset expression attributes from GET, so an absent value is
// treated as the empty default.
func testAccCheckDnsaction64ADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Dnsaction64.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("dnsaction64 %s not found on appliance", name)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = strings.TrimSpace(fmt.Sprintf("%v", v))
		}
		if got != want {
			return fmt.Errorf("dnsaction64 %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccDnsaction64DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsaction64DataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_dnsaction64.dnsaction64_datasource", "actionname", "default_DNS64_action1"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsaction64.dnsaction64_datasource", "prefix", "64:ff9c::/96"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsaction64.dnsaction64_datasource", "mappedrule", "DNS.RR.RDATA.IP.IN_SUBNET(10.0.0.0/8)"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsaction64.dnsaction64_datasource", "excluderule", "DNS.RR.RDATA.IPV6.IN_SUBNET(::ffff:0:0/96)"),
				),
			},
		},
	})
}
