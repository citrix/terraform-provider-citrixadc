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

const testAccDnsview_basic = `


	resource "citrixadc_dnsview" "tf_dnsview" {
		viewname = "view3"
		
	}
`

func TestAccDnsview_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnsviewDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsview_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsviewExist("citrixadc_dnsview.tf_dnsview", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsview.tf_dnsview", "viewname", "view3"),
				),
			},
		},
	})
}

func testAccCheckDnsviewExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnsview name is set")
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
		data, err := client.FindResource(service.Dnsview.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("dnsview %s not found", n)
		}

		return nil
	}
}

func testAccCheckDnsviewDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnsview" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Dnsview.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("dnsview %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccDnsview_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_dnsview.tf_dnsview"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnsviewDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsview_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnsviewExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Dnsview.Type(), "view3"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccDnsview_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnsviewExist(resAddr, nil)),
			},
		},
	})
}

const testAccDnsviewDataSource_basic = `
	resource "citrixadc_dnsview" "tf_dnsview_ds" {
		viewname = "view_ds_test"
	}

	data "citrixadc_dnsview" "tf_dnsview_ds" {
		viewname = citrixadc_dnsview.tf_dnsview_ds.viewname
	}
`

func TestAccDnsviewDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsviewDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_dnsview.tf_dnsview_ds", "viewname", "view_ds_test"),
					resource.TestCheckResourceAttr("data.citrixadc_dnsview.tf_dnsview_ds", "id", "view_ds_test"),
				),
			},
		},
	})
}

func TestAccDnsview_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDnsviewDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccDnsview_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnsviewExist("citrixadc_dnsview.tf_dnsview", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccDnsview_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnsviewExist("citrixadc_dnsview.tf_dnsview", nil)),
			},
		},
	})
}

func TestAccDnsview_import(t *testing.T) {
	const resAddr = "citrixadc_dnsview.tf_dnsview"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnsviewDestroy,
		Steps: []resource.TestStep{
			{Config: testAccDnsview_basic},
			{
				Config:                  testAccDnsview_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}
