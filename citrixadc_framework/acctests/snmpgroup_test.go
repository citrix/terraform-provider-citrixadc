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
	"log"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testAccSnmpgroup_basic = `
	resource "citrixadc_snmpgroup" "tf_snmpgroup" {
	name          = "test_group"
	securitylevel = "noAuthNoPriv"
	readviewname  = "test_name"
	}
`
const testAccSnmpgroup_update = `
	resource "citrixadc_snmpgroup" "tf_snmpgroup" {
	name          = "test_group"
	securitylevel = "noAuthNoPriv"
	readviewname  = "test2_name"
	}
`

func TestAccSnmpgroup_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmpgroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpgroup_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpgroupExist("citrixadc_snmpgroup.tf_snmpgroup", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpgroup.tf_snmpgroup", "securitylevel", "noAuthNoPriv"),
					resource.TestCheckResourceAttr("citrixadc_snmpgroup.tf_snmpgroup", "readviewname", "test_name"),
				),
			},
			{
				Config: testAccSnmpgroup_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSnmpgroupExist("citrixadc_snmpgroup.tf_snmpgroup", nil),
					resource.TestCheckResourceAttr("citrixadc_snmpgroup.tf_snmpgroup", "securitylevel", "noAuthNoPriv"),
					resource.TestCheckResourceAttr("citrixadc_snmpgroup.tf_snmpgroup", "readviewname", "test2_name"),
				),
			},
		},
	})
}

func TestAccSnmpgroup_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_snmpgroup.tf_snmpgroup"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmpgroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpgroup_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSnmpgroupExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgs(service.Snmpgroup.Type(), "test_group", []string{"securitylevel:noAuthNoPriv"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccSnmpgroup_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSnmpgroupExist(resAddr, nil)),
			},
		},
	})
}

func testAccCheckSnmpgroupExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No snmpgroup name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}
		snmpgroupName := rs.Primary.ID

		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		dataArr, err := client.FindAllResources(service.Snmpgroup.Type())

		if err != nil {
			return err
		}

		if len(dataArr) == 0 {
			log.Printf("[WARN] citrix-provider: acceptance test: snmpgroup does not exist. Clearing state.")
			return nil
		}

		found := false
		for _, v := range dataArr {
			if v["name"] == snmpgroupName {
				found = true
				break
			}

		}
		if !found {
			return fmt.Errorf("snmpgroup %s not found", n)
		}

		return nil
	}
}

func testAccCheckSnmpgroupDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_snmpgroup" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		snmpgroupName := rs.Primary.ID

		dataArr, err := client.FindAllResources(service.Snmpgroup.Type())

		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if v["name"] == snmpgroupName {
				found = true
				break
			}

			if found {
				return fmt.Errorf("snmpgroup %s still exists", rs.Primary.ID)
			}

		}

	}
	return nil
}

func TestAccSnmpgroup_import(t *testing.T) {
	const resAddr = "citrixadc_snmpgroup.tf_snmpgroup"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmpgroupDestroy,
		Steps: []resource.TestStep{
			{Config: testAccSnmpgroup_basic},
			{
				Config:                  testAccSnmpgroup_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccSnmpgroup_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckSnmpgroupDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccSnmpgroup_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSnmpgroupExist("citrixadc_snmpgroup.tf_snmpgroup", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccSnmpgroup_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSnmpgroupExist("citrixadc_snmpgroup.tf_snmpgroup", nil)),
			},
		},
	})
}

func TestAccSnmpgroupDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSnmpgroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSnmpgroupDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_snmpgroup.tf_snmpgroup_ds", "name", "tf_group_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_snmpgroup.tf_snmpgroup_ds", "securitylevel", "noAuthNoPriv"),
					resource.TestCheckResourceAttr("data.citrixadc_snmpgroup.tf_snmpgroup_ds", "readviewname", "tf_view_ds"),
					// Universal runtime-binding proof for the data source.
					resource.TestCheckResourceAttrSet("data.citrixadc_snmpgroup.tf_snmpgroup_ds", "id"),
				),
			},
		},
	})
}

const testAccSnmpgroupDataSource_basic = `

resource "citrixadc_snmpgroup" "tf_snmpgroup_ds" {
	name          = "tf_group_ds"
	securitylevel = "noAuthNoPriv"
	readviewname  = "tf_view_ds"
}

data "citrixadc_snmpgroup" "tf_snmpgroup_ds" {
	name = citrixadc_snmpgroup.tf_snmpgroup_ds.name
	securitylevel = citrixadc_snmpgroup.tf_snmpgroup_ds.securitylevel
	depends_on = [citrixadc_snmpgroup.tf_snmpgroup_ds]
}
`
