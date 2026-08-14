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

const testAccNstrafficdomain_basic = `
	resource "citrixadc_nstrafficdomain" "tf_trafficdomain" {
		td        = 2
		aliasname = "tf_trafficdomain"
		vmac      = "ENABLED"
	}
`

func TestAccNstrafficdomain_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNstrafficdomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNstrafficdomain_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNstrafficdomainExist("citrixadc_nstrafficdomain.tf_trafficdomain", nil),
					resource.TestCheckResourceAttr("citrixadc_nstrafficdomain.tf_trafficdomain", "td", "2"),
					resource.TestCheckResourceAttr("citrixadc_nstrafficdomain.tf_trafficdomain", "aliasname", "tf_trafficdomain"),
					resource.TestCheckResourceAttr("citrixadc_nstrafficdomain.tf_trafficdomain", "vmac", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckNstrafficdomainExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nstrafficdomain name is set")
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
		data, err := client.FindResource(service.Nstrafficdomain.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nstrafficdomain %s not found", n)
		}

		return nil
	}
}

func testAccCheckNstrafficdomainDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nstrafficdomain" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Nstrafficdomain.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nstrafficdomain %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccNstrafficdomain_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_nstrafficdomain.tf_trafficdomain"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNstrafficdomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNstrafficdomain_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNstrafficdomainExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Nstrafficdomain.Type(), "2"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccNstrafficdomain_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNstrafficdomainExist(resAddr, nil)),
			},
		},
	})
}

func TestAccNstrafficdomain_import(t *testing.T) {
	const resAddr = "citrixadc_nstrafficdomain.tf_trafficdomain"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNstrafficdomainDestroy,
		Steps: []resource.TestStep{
			{Config: testAccNstrafficdomain_basic},
			{
				Config:                  testAccNstrafficdomain_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccNstrafficdomainDataSource_basic = `

	resource "citrixadc_nstrafficdomain" "tf_trafficdomain" {
		td        = 3
		aliasname = "tf_trafficdomain_ds"
		vmac      = "ENABLED"
	}

	data "citrixadc_nstrafficdomain" "tf_trafficdomain_data" {
		td = citrixadc_nstrafficdomain.tf_trafficdomain.td
	}
`

func TestAccNstrafficdomain_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNstrafficdomainDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccNstrafficdomain_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNstrafficdomainExist("citrixadc_nstrafficdomain.tf_trafficdomain", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccNstrafficdomain_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNstrafficdomainExist("citrixadc_nstrafficdomain.tf_trafficdomain", nil)),
			},
		},
	})
}

func TestAccNstrafficdomainDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNstrafficdomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNstrafficdomainDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nstrafficdomain.tf_trafficdomain_data", "td", "3"),
					resource.TestCheckResourceAttr("data.citrixadc_nstrafficdomain.tf_trafficdomain_data", "aliasname", "tf_trafficdomain_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_nstrafficdomain.tf_trafficdomain_data", "vmac", "ENABLED"),
				),
			},
		},
	})
}
