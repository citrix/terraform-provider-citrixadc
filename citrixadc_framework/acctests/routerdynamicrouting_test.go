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

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccRouterdynamicrouting_basic(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	if isCpxRun {
		t.Skip("Feature not supported on CPX")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRouterdynamicrouting_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouterdynamicroutingExist("citrixadc_routerdynamicrouting.tf_dynamicrouting", nil),
				),
			},
		},
	})
}

func TestAccRouterdynamicrouting_import(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	if isCpxRun {
		t.Skip("Feature not supported on CPX")
	}
	const resAddr = "citrixadc_routerdynamicrouting.tf_dynamicrouting"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccRouterdynamicrouting_basic},
			{
				Config:                  testAccRouterdynamicrouting_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccRouterdynamicrouting_sdkv2StateUpgrade(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	if isCpxRun {
		t.Skip("Feature not supported on CPX")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccRouterdynamicrouting_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckRouterdynamicroutingExist("citrixadc_routerdynamicrouting.tf_dynamicrouting", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccRouterdynamicrouting_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckRouterdynamicroutingExist("citrixadc_routerdynamicrouting.tf_dynamicrouting", nil)),
			},
		},
	})
}

func testAccCheckRouterdynamicroutingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lb vserver name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		return nil
	}
}

const testAccRouterdynamicrouting_basic = `

resource "citrixadc_routerdynamicrouting" "tf_dynamicrouting" {
    commandlines = [
        "router bgp 101",
        "neighbor 192.168.5.1 remote-as 100",
        "redistribute kernel",
    ]
}

`

// const testAccRouterdynamicroutingDataSource_basic = `
// 	resource "citrixadc_routerdynamicrouting" "tf_dynamicrouting_ds" {
// 		commandlines = [
// 			"router bgp 101",
// 			"neighbor 192.168.5.1 remote-as 100",
// 			"redistribute kernel",
// 		]
// 	}

// 	data "citrixadc_routerdynamicrouting" "tf_dynamicrouting_ds" {
// 		commandstring = "show ip bgp summary"
// 		depends_on = [citrixadc_routerdynamicrouting.tf_dynamicrouting_ds]
// 	}
// `

// func TestAccRouterdynamicroutingDataSource_basic(t *testing.T) {
// 	t.Skip("TODO: Need to find a way to test this resource!")
// 	if isCpxRun {
// 		t.Skip("Feature not supported on CPX")
// 	}
// 	resource.Test(t, resource.TestCase{
// 		PreCheck:                 func() { testAccPreCheck(t) },
// 		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccRouterdynamicroutingDataSource_basic,
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttr("data.citrixadc_routerdynamicrouting.tf_dynamicrouting_ds", "commandstring", "show ip bgp summary"),
// 					resource.TestCheckResourceAttrSet("data.citrixadc_routerdynamicrouting.tf_dynamicrouting_ds", "id"),
// 				),
// 			},
// 		},
// 	})
// }
