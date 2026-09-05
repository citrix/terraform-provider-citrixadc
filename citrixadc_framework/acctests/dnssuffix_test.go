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

const testAccDnssuffix_basic = `
	resource "citrixadc_dnssuffix" "tf_dnssuffix" {
		dnssuffix = "example.com"
	}
`

func TestAccDnssuffix_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnssuffixDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnssuffix_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnssuffixExist("citrixadc_dnssuffix.tf_dnssuffix", nil),
					resource.TestCheckResourceAttr("citrixadc_dnssuffix.tf_dnssuffix", "dnssuffix", "example.com"),
				),
			},
		},
	})
}

func testAccCheckDnssuffixExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnssuffix name is set")
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
		data, err := client.FindResource(service.Dnssuffix.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("dnssuffix %s not found", n)
		}

		return nil
	}
}

func testAccCheckDnssuffixDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnssuffix" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Dnssuffix.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("dnssuffix %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccDnssuffix_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_dnssuffix.tf_dnssuffix"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnssuffixDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnssuffix_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnssuffixExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Dnssuffix.Type(), "example.com"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccDnssuffix_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnssuffixExist(resAddr, nil)),
			},
		},
	})
}

func TestAccDnssuffix_import(t *testing.T) {
	const resAddr = "citrixadc_dnssuffix.tf_dnssuffix"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnssuffixDestroy,
		Steps: []resource.TestStep{
			{Config: testAccDnssuffix_basic},
			{
				Config:                  testAccDnssuffix_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccDnssuffix_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDnssuffixDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccDnssuffix_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnssuffixExist("citrixadc_dnssuffix.tf_dnssuffix", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccDnssuffix_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnssuffixExist("citrixadc_dnssuffix.tf_dnssuffix", nil)),
			},
		},
	})
}

const testAccDnssuffixDataSource_basic = `

	resource "citrixadc_dnssuffix" "tf_dnssuffix" {
		dnssuffix = "example.com"
	}
	
	data "citrixadc_dnssuffix" "tf_dnssuffix" {
		dnssuffix = citrixadc_dnssuffix.tf_dnssuffix.dnssuffix
	}
`

func TestAccDnssuffixDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnssuffixDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnssuffixDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_dnssuffix.tf_dnssuffix", "dnssuffix", "example.com"),
				),
			},
		},
	})
}
