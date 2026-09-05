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

const testAccLocationfile_basic = `

	resource "citrixadc_locationfile" "tf_locationfile" {
		locationfile = "/var/netscaler/inbuilt_db/Citrix_Netscaler_InBuilt_GeoIP_DB_IPv4"
		format       = "netscaler"
	}
`

func TestAccLocationfile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLocationfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLocationfile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLocationfileExist("citrixadc_locationfile.tf_locationfile", nil),
					resource.TestCheckResourceAttr("citrixadc_locationfile.tf_locationfile", "locationfile", "/var/netscaler/inbuilt_db/Citrix_Netscaler_InBuilt_GeoIP_DB_IPv4"),
					resource.TestCheckResourceAttr("citrixadc_locationfile.tf_locationfile", "format", "netscaler"),
				),
			},
		},
	})
}

func testAccCheckLocationfileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No locationfile name is set")
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
		data, err := client.FindResource(service.Locationfile.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("locationfile %s not found", n)
		}

		return nil
	}
}

func testAccCheckLocationfileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_locationfile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		data, _ := client.FindResource(service.Locationfile.Type(), "")
		// if err == nil {
		// 	return fmt.Errorf("locationfile %s still exists", rs.Primary.ID)
		// }
		if data["locationfile"] == rs.Primary.Attributes["locationfile"] {
			return fmt.Errorf("locationfile %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

func TestAccLocationfile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_locationfile.tf_locationfile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLocationfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLocationfile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLocationfileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Locationfile.Type(), ""); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccLocationfile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLocationfileExist(resAddr, nil)),
			},
		},
	})
}

func TestAccLocationfile_import(t *testing.T) {
	const resAddr = "citrixadc_locationfile.tf_locationfile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLocationfileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLocationfile_basic},
			{
				Config:                  testAccLocationfile_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccLocationfileDataSource_basic = `

	resource "citrixadc_locationfile" "tf_locationfile" {
		locationfile = "/var/netscaler/inbuilt_db/Citrix_Netscaler_InBuilt_GeoIP_DB_IPv4"
		format       = "netscaler"
	}

	data "citrixadc_locationfile" "tf_locationfile" {
		depends_on = [citrixadc_locationfile.tf_locationfile]
	}
`

func TestAccLocationfile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLocationfileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccLocationfile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLocationfileExist("citrixadc_locationfile.tf_locationfile", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccLocationfile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLocationfileExist("citrixadc_locationfile.tf_locationfile", nil)),
			},
		},
	})
}

func TestAccLocationfileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccLocationfileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_locationfile.tf_locationfile", "locationfile", "/var/netscaler/inbuilt_db/Citrix_Netscaler_InBuilt_GeoIP_DB_IPv4"),
					resource.TestCheckResourceAttr("data.citrixadc_locationfile.tf_locationfile", "format", "netscaler"),
					// id is the universal runtime-binding proof for the data source.
					resource.TestCheckResourceAttrSet("data.citrixadc_locationfile.tf_locationfile", "id"),
					// curlocfilestatus is a status field the appliance always reports
					// for a loaded location file.
					resource.TestCheckResourceAttrSet("data.citrixadc_locationfile.tf_locationfile", "curlocfilestatus"),
				),
			},
		},
	})
}
