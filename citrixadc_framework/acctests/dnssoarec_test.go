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

const testAccDnssoarec_basic_step1 = `
resource "citrixadc_dnssoarec" "tf_dnssoarec" {
	domain =  "test.com"
	originserver  = "10.2.3.5"
	contact =  "other"
	expire = 1800
	refresh = 4000
}
`

const testAccDnssoarec_basic_step2 = `
resource "citrixadc_dnssoarec" "tf_dnssoarec" {
	domain =  "test.com"
	originserver  = "10.2.3.6"
	contact =  "some_other"
	expire = 1600
	refresh = 3600
}
`

const testAccDnssoarecDataSource_basic = `

resource "citrixadc_dnssoarec" "tf_dnssoarec" {
	domain =  "test.com"
	originserver  = "10.2.3.5"
	contact =  "other"
	expire = 1800
	refresh = 4000
}

data "citrixadc_dnssoarec" "tf_dnssoarec_ds" {
	domain = citrixadc_dnssoarec.tf_dnssoarec.domain
	depends_on = [citrixadc_dnssoarec.tf_dnssoarec]
}

`

func TestAccDnssoarec_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnssoarecDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnssoarec_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnssoarecExist("citrixadc_dnssoarec.tf_dnssoarec", nil),
				),
			},
			{
				Config: testAccDnssoarec_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnssoarecExist("citrixadc_dnssoarec.tf_dnssoarec", nil),
				),
			},
		},
	})
}

func TestAccDnssoarecDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDnssoarecDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_dnssoarec.tf_dnssoarec_ds", "domain", "test.com"),
					resource.TestCheckResourceAttr("data.citrixadc_dnssoarec.tf_dnssoarec_ds", "originserver", "10.2.3.5"),
					resource.TestCheckResourceAttr("data.citrixadc_dnssoarec.tf_dnssoarec_ds", "contact", "other"),
					resource.TestCheckResourceAttr("data.citrixadc_dnssoarec.tf_dnssoarec_ds", "expire", "1800"),
					resource.TestCheckResourceAttr("data.citrixadc_dnssoarec.tf_dnssoarec_ds", "refresh", "4000"),
				),
			},
		},
	})
}

func TestAccDnssoarec_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_dnssoarec.tf_dnssoarec"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnssoarecDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnssoarec_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnssoarecExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Dnssoarec.Type(), "test.com"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccDnssoarec_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnssoarecExist(resAddr, nil)),
			},
		},
	})
}

func TestAccDnssoarec_import(t *testing.T) {
	const resAddr = "citrixadc_dnssoarec.tf_dnssoarec"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnssoarecDestroy,
		Steps: []resource.TestStep{
			{Config: testAccDnssoarec_basic_step1},
			{
				Config:                  testAccDnssoarec_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccDnssoarec_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDnssoarecDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccDnssoarec_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnssoarecExist("citrixadc_dnssoarec.tf_dnssoarec", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccDnssoarec_basic_step1,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckDnssoarecExist("citrixadc_dnssoarec.tf_dnssoarec", nil)),
			},
		},
	})
}

// dnssoarec unset test: the NITRO-unsettable attributes (per the dnssoarec
// spec's unset payload) are serial, refresh, retry, expire, minimum and ttl,
// each with a documented server default. Step 1 sets them to non-default values;
// step 2 removes them so the provider unsets them and the appliance reverts to
// the documented NITRO defaults.
const testAccDnssoarec_unset_step1 = `
resource "citrixadc_dnssoarec" "tf_unset" {
	domain       = "unset.test.com"
	originserver = "10.2.3.5"
	contact      = "admin.unset.test.com"
	serial       = 200
	refresh      = 4000
	retry        = 5
	expire       = 1800
	minimum      = 10
	ttl          = 7200
}
`

const testAccDnssoarec_unset_step2 = `
resource "citrixadc_dnssoarec" "tf_unset" {
	domain       = "unset.test.com"
	originserver = "10.2.3.5"
	contact      = "admin.unset.test.com"
	# unset-eligible attributes removed -> provider unsets them (NITRO defaults).
}
`

func TestAccDnssoarec_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnssoarecDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccDnssoarec_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnssoarecExist("citrixadc_dnssoarec.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_dnssoarec.tf_unset", "serial", "200"),
					resource.TestCheckResourceAttr("citrixadc_dnssoarec.tf_unset", "refresh", "4000"),
					resource.TestCheckResourceAttr("citrixadc_dnssoarec.tf_unset", "retry", "5"),
					resource.TestCheckResourceAttr("citrixadc_dnssoarec.tf_unset", "expire", "1800"),
					resource.TestCheckResourceAttr("citrixadc_dnssoarec.tf_unset", "minimum", "10"),
					resource.TestCheckResourceAttr("citrixadc_dnssoarec.tf_unset", "ttl", "7200"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults and the
				// implicit post-apply plan must be empty.
				Config: testAccDnssoarec_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnssoarecExist("citrixadc_dnssoarec.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_dnssoarec.tf_unset", "serial", "100"),
					resource.TestCheckResourceAttr("citrixadc_dnssoarec.tf_unset", "refresh", "3600"),
					resource.TestCheckResourceAttr("citrixadc_dnssoarec.tf_unset", "retry", "3"),
					resource.TestCheckResourceAttr("citrixadc_dnssoarec.tf_unset", "expire", "3600"),
					resource.TestCheckResourceAttr("citrixadc_dnssoarec.tf_unset", "minimum", "5"),
					resource.TestCheckResourceAttr("citrixadc_dnssoarec.tf_unset", "ttl", "3600"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckDnssoarecADCValue("unset.test.com", "serial", "100"),
					testAccCheckDnssoarecADCValue("unset.test.com", "expire", "3600"),
					testAccCheckDnssoarecADCValue("unset.test.com", "ttl", "3600"),
				),
			},
		},
	})
}

// testAccCheckDnssoarecADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckDnssoarecADCValue(domain, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Dnssoarec.Type(), domain)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("dnssoarec %s not found on appliance", domain)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("dnssoarec %s: appliance attr %q = %q, want %q (unset did not revert it)", domain, attr, got, want)
		}
		return nil
	}
}

func testAccCheckDnssoarecExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnssoarec name is set")
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
		data, err := client.FindResource(service.Dnssoarec.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("dnssoarec %s not found", n)
		}

		return nil
	}
}

func testAccCheckDnssoarecDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnssoarec" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Dnssoarec.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("dnssoarec %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
