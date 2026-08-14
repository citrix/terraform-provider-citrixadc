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

const testAccNsappflowcollector_basic = `
	resource "citrixadc_nsappflowcollector" "tf_appflowcollector" {
		name      = "tf_appflowcollector"
		ipaddress = "1.2.4.1"
		port      = 30
	}
`

const testAccNsappflowcollectorDataSource_basic = `
	resource "citrixadc_nsappflowcollector" "tf_appflowcollector" {
		name      = "tf_appflowcollector_ds"
		ipaddress = "1.2.4.2"
		port      = 31
	}

	data "citrixadc_nsappflowcollector" "tf_appflowcollector" {
		name = citrixadc_nsappflowcollector.tf_appflowcollector.name
		depends_on = [citrixadc_nsappflowcollector.tf_appflowcollector]
	}
`

func TestAccNsappflowcollector_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsappflowcollectorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsappflowcollector_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsappflowcollectorExist("citrixadc_nsappflowcollector.tf_appflowcollector", nil),
					resource.TestCheckResourceAttr("citrixadc_nsappflowcollector.tf_appflowcollector", "name", "tf_appflowcollector"),
					resource.TestCheckResourceAttr("citrixadc_nsappflowcollector.tf_appflowcollector", "ipaddress", "1.2.4.1"),
					resource.TestCheckResourceAttr("citrixadc_nsappflowcollector.tf_appflowcollector", "port", "30"),
				),
			},
		},
	})
}

func testAccCheckNsappflowcollectorExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsappflowcollector name is set")
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
		data, err := client.FindResource(service.Nsappflowcollector.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nsappflowcollector %s not found", n)
		}

		return nil
	}
}

func testAccCheckNsappflowcollectorDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nsappflowcollector" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Nsappflowcollector.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nsappflowcollector %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccNsappflowcollector_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_nsappflowcollector.tf_appflowcollector"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsappflowcollectorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsappflowcollector_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsappflowcollectorExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Nsappflowcollector.Type(), "tf_appflowcollector"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccNsappflowcollector_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsappflowcollectorExist(resAddr, nil)),
			},
		},
	})
}

func TestAccNsappflowcollector_import(t *testing.T) {
	const resAddr = "citrixadc_nsappflowcollector.tf_appflowcollector"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsappflowcollectorDestroy,
		Steps: []resource.TestStep{
			{Config: testAccNsappflowcollector_basic},
			{
				Config:                  testAccNsappflowcollector_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccNsappflowcollector_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNsappflowcollectorDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccNsappflowcollector_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsappflowcollectorExist("citrixadc_nsappflowcollector.tf_appflowcollector", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccNsappflowcollector_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsappflowcollectorExist("citrixadc_nsappflowcollector.tf_appflowcollector", nil)),
			},
		},
	})
}

func TestAccNsappflowcollectorDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNsappflowcollectorDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nsappflowcollector.tf_appflowcollector", "name", "tf_appflowcollector_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_nsappflowcollector.tf_appflowcollector", "ipaddress", "1.2.4.2"),
					resource.TestCheckResourceAttr("data.citrixadc_nsappflowcollector.tf_appflowcollector", "port", "31"),
				),
			},
		},
	})
}
