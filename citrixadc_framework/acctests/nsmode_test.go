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

const testAccNsmode_basic_step1 = `
resource "citrixadc_nsmode" "tf_nsmode" {
    usip = true
	cka = true
}
`

const testAccNsmode_basic_step2 = `
resource "citrixadc_nsmode" "tf_nsmode" {
    usip = false
	cka = false
}
`

func TestAccNsmode_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNsmode_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testEnsureNsmodes([]string{"usip", "cka"}, true),
				),
			},
			{
				Config: testAccNsmode_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testEnsureNsmodes([]string{"usip", "cka"}, false),
				),
			},
		},
	})
}

func testEnsureNsmodes(modes []string, expectedState bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		findParams := service.FindParams{
			ResourceType: "nsmode",
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}
		if len(dataArr) != 1 {
			return fmt.Errorf("Unexpected fetched nsmode result %v", dataArr)
		}
		data := dataArr[0]
		for _, mode := range modes {
			if val, ok := data[mode]; ok {
				if val.(bool) != expectedState {
					return fmt.Errorf("Wrong mode value for %s. Expected %v, found %v", mode, expectedState, val.(bool))
				}
			} else {
				return fmt.Errorf("Cannot find mode %s in retrieved modes list", mode)
			}
		}
		return nil
	}
}

func TestAccNsmodeDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNsmodeDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.citrixadc_nsmode.test", "id"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nsmode.test", "fr"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nsmode.test", "l2"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nsmode.test", "usip"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nsmode.test", "cka"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nsmode.test", "tcpb"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nsmode.test", "mbf"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nsmode.test", "edge"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nsmode.test", "usnip"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nsmode.test", "l3"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nsmode.test", "pmtud"),
				),
			},
		},
	})
}

const testAccNsmodeDataSource_basic = `
data "citrixadc_nsmode" "test" {
}
`

func TestAccNsmode_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccNsmode_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testEnsureNsmodes([]string{"usip", "cka"}, true),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccNsmode_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testEnsureNsmodes([]string{"usip", "cka"}, true),
				),
			},
		},
	})
}

func TestAccNsmode_import(t *testing.T) {
	const resAddr = "citrixadc_nsmode.tf_nsmode"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccNsmode_basic_step1},
			{
				Config:                  testAccNsmode_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}
