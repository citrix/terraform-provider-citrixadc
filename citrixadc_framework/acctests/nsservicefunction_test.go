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
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testAccNsservicefunction_add = `
	resource "citrixadc_vlan" "tf_vlan" {
		vlanid    = 20
		aliasname = "Management VLAN"
	}
	resource "citrixadc_vlan" "tf_vlan_1" {
		vlanid    = 30
		aliasname = "Management VLAN"
	}
	resource "citrixadc_nsservicefunction" "tf_servicefunc" {
		servicefunctionname = "tf_servicefunc"
		ingressvlan         = citrixadc_vlan.tf_vlan.vlanid
	}
`
const testAccNsservicefunction_update = `
	resource "citrixadc_vlan" "tf_vlan" {
		vlanid    = 20
		aliasname = "Management VLAN"
	}
	resource "citrixadc_vlan" "tf_vlan_1" {
		vlanid    = 30
		aliasname = "Management VLAN"
	}
	resource "citrixadc_nsservicefunction" "tf_servicefunc" {
		servicefunctionname = "tf_servicefunc"
		ingressvlan         = citrixadc_vlan.tf_vlan_1.vlanid
	}
`

func TestAccNsservicefunction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsservicefunctionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsservicefunction_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsservicefunctionExist("citrixadc_nsservicefunction.tf_servicefunc", nil),
					resource.TestCheckResourceAttr("citrixadc_nsservicefunction.tf_servicefunc", "servicefunctionname", "tf_servicefunc"),
					resource.TestCheckResourceAttr("citrixadc_nsservicefunction.tf_servicefunc", "ingressvlan", "20"),
				),
			},
			{
				Config: testAccNsservicefunction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsservicefunctionExist("citrixadc_nsservicefunction.tf_servicefunc", nil),
					resource.TestCheckResourceAttr("citrixadc_nsservicefunction.tf_servicefunc", "servicefunctionname", "tf_servicefunc"),
					resource.TestCheckResourceAttr("citrixadc_nsservicefunction.tf_servicefunc", "ingressvlan", "30"),
				),
			},
		},
	})
}

func testAccCheckNsservicefunctionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsservicefunction name is set")
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
		data, err := client.FindResource(service.Nsservicefunction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nsservicefunction %s not found", n)
		}

		return nil
	}
}

func testAccCheckNsservicefunctionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nsservicefunction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Nsservicefunction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nsservicefunction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccNsservicefunction_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_nsservicefunction.tf_servicefunc"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsservicefunctionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsservicefunction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsservicefunctionExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Nsservicefunction.Type(), "tf_servicefunc"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccNsservicefunction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsservicefunctionExist(resAddr, nil)),
			},
		},
	})
}

func TestAccNsservicefunction_import(t *testing.T) {
	const resAddr = "citrixadc_nsservicefunction.tf_servicefunc"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsservicefunctionDestroy,
		Steps: []resource.TestStep{
			{Config: testAccNsservicefunction_add},
			{
				Config:                  testAccNsservicefunction_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccNsservicefunctionDataSource_basic = `
	resource "citrixadc_vlan" "tf_vlan" {
		vlanid    = 25
		aliasname = "Test VLAN"
	}
	resource "citrixadc_nsservicefunction" "tf_servicefunc" {
		servicefunctionname = "tf_servicefunc_ds"
		ingressvlan         = citrixadc_vlan.tf_vlan.vlanid
	}

	data "citrixadc_nsservicefunction" "tf_servicefunc" {
		servicefunctionname = citrixadc_nsservicefunction.tf_servicefunc.servicefunctionname
	}
`

func TestAccNsservicefunction_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNsservicefunctionDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccNsservicefunction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsservicefunctionExist("citrixadc_nsservicefunction.tf_servicefunc", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccNsservicefunction_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsservicefunctionExist("citrixadc_nsservicefunction.tf_servicefunc", nil)),
			},
		},
	})
}

func TestAccNsservicefunctionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNsservicefunctionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nsservicefunction.tf_servicefunc", "servicefunctionname", "tf_servicefunc_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_nsservicefunction.tf_servicefunc", "ingressvlan", "25"),
				),
			},
		},
	})
}
