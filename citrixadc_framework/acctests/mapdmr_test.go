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

const testAccMapdmr_basic = `

	resource "citrixadc_mapdmr" "tf_mapdmr" {
		name         = "tf_mapdmr"
		bripv6prefix = "2002:db8::/64"
	}
`

func TestAccMapdmr_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMapdmrDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMapdmr_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMapdmrExist("citrixadc_mapdmr.tf_mapdmr", nil),
					resource.TestCheckResourceAttr("citrixadc_mapdmr.tf_mapdmr", "name", "tf_mapdmr"),
					resource.TestCheckResourceAttr("citrixadc_mapdmr.tf_mapdmr", "bripv6prefix", "2002:db8::/64"),
				),
			},
		},
	})
}

func testAccCheckMapdmrExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No mapdmr name is set")
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
		data, err := client.FindResource("mapdmr", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("mapdmr %s not found", n)
		}

		return nil
	}
}

func testAccCheckMapdmrDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_mapdmr" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("mapdmr", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("mapdmr %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccMapdmr_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_mapdmr.tf_mapdmr"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMapdmrDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMapdmr_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckMapdmrExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Mapdmr.Type(), "tf_mapdmr"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccMapdmr_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckMapdmrExist(resAddr, nil)),
			},
		},
	})
}

func TestAccMapdmr_import(t *testing.T) {
	const resAddr = "citrixadc_mapdmr.tf_mapdmr"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMapdmrDestroy,
		Steps: []resource.TestStep{
			{Config: testAccMapdmr_basic},
			{
				Config:                  testAccMapdmr_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccMapdmr_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckMapdmrDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccMapdmr_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckMapdmrExist("citrixadc_mapdmr.tf_mapdmr", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccMapdmr_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckMapdmrExist("citrixadc_mapdmr.tf_mapdmr", nil)),
			},
		},
	})
}

const testAccMapdmrDataSource_basic = `

	resource "citrixadc_mapdmr" "tf_mapdmr_ds" {
		name         = "tf_mapdmr_ds"
		bripv6prefix = "2002:db8::/64"
	}

	data "citrixadc_mapdmr" "tf_mapdmr_ds_data" {
		name = citrixadc_mapdmr.tf_mapdmr_ds.name
	}
`

func TestAccMapdmrDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMapdmrDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_mapdmr.tf_mapdmr_ds_data", "name", "tf_mapdmr_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_mapdmr.tf_mapdmr_ds_data", "bripv6prefix", "2002:db8::/64"),
				),
			},
		},
	})
}
