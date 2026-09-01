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

const testAccDnspolicylabel_add = `


resource "citrixadc_dnspolicylabel" "dnspolicylabel" {
	labelname = "label1"
	transform = "dns_req"
	
	}
`

const testAccDnspolicylabelDataSource_basic = `
	resource "citrixadc_dnspolicylabel" "dnspolicylabel" {
		labelname = "label1"
		transform = "dns_req"
	}

	data "citrixadc_dnspolicylabel" "dnspolicylabel" {
		labelname = citrixadc_dnspolicylabel.dnspolicylabel.labelname
		depends_on = [citrixadc_dnspolicylabel.dnspolicylabel]
	}
`

func TestAccDnspolicylabel_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnspolicylabelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnspolicylabel_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnspolicylabelExist("citrixadc_dnspolicylabel.dnspolicylabel", nil),
					resource.TestCheckResourceAttr("citrixadc_dnspolicylabel.dnspolicylabel", "labelname", "label1"),
					resource.TestCheckResourceAttr("citrixadc_dnspolicylabel.dnspolicylabel", "transform", "dns_req"),
				),
			},
		},
	})
}

func testAccCheckDnspolicylabelExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnspolicylabel name is set")
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
		data, err := client.FindResource(service.Dnspolicylabel.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("dnspolicylabel %s not found", n)
		}

		return nil
	}
}

func testAccCheckDnspolicylabelDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnspolicylabel" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Dnspolicylabel.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("dnspolicylabel %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccDnspolicylabel_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_dnspolicylabel.dnspolicylabel"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnspolicylabelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnspolicylabel_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnspolicylabelExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Dnspolicylabel.Type(), "label1"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccDnspolicylabel_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnspolicylabelExist(resAddr, nil)),
			},
		},
	})
}

func TestAccDnspolicylabel_import(t *testing.T) {
	const resAddr = "citrixadc_dnspolicylabel.dnspolicylabel"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnspolicylabelDestroy,
		Steps: []resource.TestStep{
			{Config: testAccDnspolicylabel_add},
			{
				Config:                  testAccDnspolicylabel_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccDnspolicylabel_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDnspolicylabelDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccDnspolicylabel_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnspolicylabelExist("citrixadc_dnspolicylabel.dnspolicylabel", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccDnspolicylabel_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnspolicylabelExist("citrixadc_dnspolicylabel.dnspolicylabel", nil)),
			},
		},
	})
}

func TestAccDnspolicylabelDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccDnspolicylabelDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_dnspolicylabel.dnspolicylabel", "labelname", "label1"),
					resource.TestCheckResourceAttr("data.citrixadc_dnspolicylabel.dnspolicylabel", "transform", "dns_req"),
					// id is the universal runtime-binding proof.
					resource.TestCheckResourceAttrSet("data.citrixadc_dnspolicylabel.dnspolicylabel", "id"),
				),
			},
		},
	})
}
