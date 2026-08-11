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

const testAccLsnpool_basic = `


	resource "citrixadc_lsnpool" "tf_lsnpool" {
		poolname            = "my_lsn_pool"
		nattype             = "DYNAMIC"
		portblockallocation = "DISABLED"
		maxportrealloctmq   = 50
		portrealloctimeout  = 50
	}
  
`

const testAccLsnpool_update = `

	resource "citrixadc_lsnpool" "tf_lsnpool" {
		poolname            = "my_lsn_pool"
		nattype             = "DYNAMIC"
		portblockallocation = "DISABLED"
		maxportrealloctmq   = 100
		portrealloctimeout  = 100
	}
  
`

func TestAccLsnpool_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnpoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnpool_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnpoolExist("citrixadc_lsnpool.tf_lsnpool", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnpool.tf_lsnpool", "poolname", "my_lsn_pool"),
					resource.TestCheckResourceAttr("citrixadc_lsnpool.tf_lsnpool", "nattype", "DYNAMIC"),
					resource.TestCheckResourceAttr("citrixadc_lsnpool.tf_lsnpool", "portblockallocation", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnpool.tf_lsnpool", "maxportrealloctmq", "50"),
					resource.TestCheckResourceAttr("citrixadc_lsnpool.tf_lsnpool", "portrealloctimeout", "50"),
				),
			},
			{
				Config: testAccLsnpool_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnpoolExist("citrixadc_lsnpool.tf_lsnpool", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnpool.tf_lsnpool", "poolname", "my_lsn_pool"),
					resource.TestCheckResourceAttr("citrixadc_lsnpool.tf_lsnpool", "nattype", "DYNAMIC"),
					resource.TestCheckResourceAttr("citrixadc_lsnpool.tf_lsnpool", "portblockallocation", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_lsnpool.tf_lsnpool", "maxportrealloctmq", "100"),
					resource.TestCheckResourceAttr("citrixadc_lsnpool.tf_lsnpool", "portrealloctimeout", "100"),
				),
			},
		},
	})
}

func TestAccLsnpool_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_lsnpool.tf_lsnpool"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnpoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnpool_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsnpoolExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Lsnpool.Type(), "my_lsn_pool"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccLsnpool_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsnpoolExist(resAddr, nil)),
			},
		},
	})
}

func TestAccLsnpool_import(t *testing.T) {
	const resAddr = "citrixadc_lsnpool.tf_lsnpool"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnpoolDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLsnpool_basic},
			{
				Config:                  testAccLsnpool_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccLsnpool_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLsnpoolDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccLsnpool_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLsnpoolExist("citrixadc_lsnpool.tf_lsnpool", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccLsnpool_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckLsnpoolExist("citrixadc_lsnpool.tf_lsnpool", nil)),
			},
		},
	})
}

// lsnpool unset test: maxportrealloctmq and portrealloctimeout are the only
// updatable (non-ForceNew) attributes that NITRO supports unsetting. Step 1
// sets them to non-default values; step 2 removes them from config so the
// provider unsets them, reverting to NITRO defaults (maxportrealloctmq=65536,
// portrealloctimeout=0).
const testAccLsnpool_unset_step1 = `
	resource "citrixadc_lsnpool" "tf_unset" {
		poolname            = "tf_test_lsnpool_unset"
		nattype             = "DYNAMIC"
		portblockallocation = "DISABLED"
		maxportrealloctmq   = 50
		portrealloctimeout  = 50
	}
`

const testAccLsnpool_unset_step2 = `
	resource "citrixadc_lsnpool" "tf_unset" {
		poolname            = "tf_test_lsnpool_unset"
		nattype             = "DYNAMIC"
		portblockallocation = "DISABLED"
		# maxportrealloctmq and portrealloctimeout removed -> provider unsets them.
	}
`

func TestAccLsnpool_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLsnpoolDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnpool_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnpoolExist("citrixadc_lsnpool.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnpool.tf_unset", "maxportrealloctmq", "50"),
					resource.TestCheckResourceAttr("citrixadc_lsnpool.tf_unset", "portrealloctimeout", "50"),
				),
			},
			{
				Config: testAccLsnpool_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLsnpoolExist("citrixadc_lsnpool.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_lsnpool.tf_unset", "maxportrealloctmq", "65536"),
					resource.TestCheckResourceAttr("citrixadc_lsnpool.tf_unset", "portrealloctimeout", "0"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckLsnpoolADCValue("tf_test_lsnpool_unset", "maxportrealloctmq", "65536"),
					testAccCheckLsnpoolADCValue("tf_test_lsnpool_unset", "portrealloctimeout", "0"),
				),
			},
		},
	})
}

// testAccCheckLsnpoolADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckLsnpoolADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Lsnpool.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("lsnpool %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("lsnpool %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func testAccCheckLsnpoolExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lsnpool name is set")
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
		data, err := client.FindResource("lsnpool", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("lsnpool %s not found", n)
		}

		return nil
	}
}

func testAccCheckLsnpoolDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lsnpool" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lsnpool", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("lsnpool %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccLsnpoolDataSource_basic = `

resource "citrixadc_lsnpool" "tf_lsnpool_ds" {
	poolname            = "my_lsn_pool_ds"
	nattype             = "DYNAMIC"
	portblockallocation = "DISABLED"
	maxportrealloctmq   = 50
	portrealloctimeout  = 50
}

data "citrixadc_lsnpool" "tf_lsnpool_ds" {
	poolname = citrixadc_lsnpool.tf_lsnpool_ds.poolname
}
`

func TestAccLsnpoolDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLsnpoolDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_lsnpool.tf_lsnpool_ds", "poolname", "my_lsn_pool_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnpool.tf_lsnpool_ds", "nattype", "DYNAMIC"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnpool.tf_lsnpool_ds", "portblockallocation", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnpool.tf_lsnpool_ds", "maxportrealloctmq", "50"),
					resource.TestCheckResourceAttr("data.citrixadc_lsnpool.tf_lsnpool_ds", "portrealloctimeout", "50"),
				),
			},
		},
	})
}
