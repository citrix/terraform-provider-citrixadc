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

const testAccStreamselector_basic = `

resource "citrixadc_streamselector" "tf_streamselector" {
	name = "my_streamselector"
	rule = ["HTTP.REQ.URL", "CLIENT.IP.SRC"]
	}
  
`
const testAccStreamselector_update = `

resource "citrixadc_streamselector" "tf_streamselector" {
	name = "my_streamselector"
	rule = ["HTTP.REQ.URL", "CLIENT.IP.SRC", "true"]
	}
  
`

func TestAccStreamselector_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckStreamselectorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccStreamselector_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStreamselectorExist("citrixadc_streamselector.tf_streamselector", nil),
				),
			},
		},
	})
}

func testAccCheckStreamselectorExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No streamselector name is set")
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
		data, err := client.FindResource(service.Streamselector.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("streamselector %s not found", n)
		}

		return nil
	}
}

func testAccCheckStreamselectorDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_streamselector" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Streamselector.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("streamselector %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccStreamselector_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_streamselector.tf_streamselector"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckStreamselectorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccStreamselector_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckStreamselectorExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Streamselector.Type(), "my_streamselector"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccStreamselector_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckStreamselectorExist(resAddr, nil)),
			},
		},
	})
}

func TestAccStreamselector_import(t *testing.T) {
	const resAddr = "citrixadc_streamselector.tf_streamselector"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckStreamselectorDestroy,
		Steps: []resource.TestStep{
			{Config: testAccStreamselector_basic},
			{
				Config:                  testAccStreamselector_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccStreamselectorDataSource_basic = `
	resource "citrixadc_streamselector" "tf_streamselector" {
		name = "my_streamselector"
		rule = ["HTTP.REQ.URL", "CLIENT.IP.SRC"]
	}

	data "citrixadc_streamselector" "tf_streamselector_datasource" {
		name = citrixadc_streamselector.tf_streamselector.name
	}
`

func TestAccStreamselector_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckStreamselectorDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccStreamselector_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckStreamselectorExist("citrixadc_streamselector.tf_streamselector", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccStreamselector_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckStreamselectorExist("citrixadc_streamselector.tf_streamselector", nil)),
			},
		},
	})
}

func TestAccStreamselectorDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccStreamselectorDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_streamselector.tf_streamselector_datasource", "name", "my_streamselector"),
					resource.TestCheckResourceAttrSet("data.citrixadc_streamselector.tf_streamselector_datasource", "id"),
				),
			},
		},
	})
}
