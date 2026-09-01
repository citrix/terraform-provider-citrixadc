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

const testAccLbmetrictable_add = `
resource "citrixadc_lbmetrictable" "tfAcc_lbmetrictable" {
        metrictable = "tf_lbmetrictable"
}
`

func TestAccLbmetrictable_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbmetrictableDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbmetrictable_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLbmetrictableExist("citrixadc_lbmetrictable.tfAcc_lbmetrictable", nil),
					resource.TestCheckResourceAttr("citrixadc_lbmetrictable.tfAcc_lbmetrictable", "metrictable", "tf_lbmetrictable"),
				),
			},
		},
	})
}

func testAccCheckLbmetrictableExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lb metrictable name is set")
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
		data, err := client.FindResource("lbmetrictable", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB metrictable %s not found", n)
		}

		return nil
	}
}

func testAccCheckLbmetrictableDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_lbmetrictable" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("lbmetrictable", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB metrictable %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccLbmetrictable_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_lbmetrictable.tfAcc_lbmetrictable"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbmetrictableDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLbmetrictable_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLbmetrictableExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Lbmetrictable.Type(), "tf_lbmetrictable"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccLbmetrictable_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLbmetrictableExist(resAddr, nil)),
			},
		},
	})
}

func TestAccLbmetrictable_import(t *testing.T) {
	const resAddr = "citrixadc_lbmetrictable.tfAcc_lbmetrictable"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLbmetrictableDestroy,
		Steps: []resource.TestStep{
			{Config: testAccLbmetrictable_add},
			{
				Config:                  testAccLbmetrictable_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccLbmetrictableDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccLbmetrictableDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					// id is the universal runtime-binding proof of a resolved data source.
					resource.TestCheckResourceAttrSet("data.citrixadc_lbmetrictable.tf_lbmetrictable_ds", "id"),
					resource.TestCheckResourceAttr("data.citrixadc_lbmetrictable.tf_lbmetrictable_ds", "metrictable", "tf_lbmetrictable_ds"),
				),
			},
		},
	})
}

func TestAccLbmetrictable_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckLbmetrictableDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccLbmetrictable_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLbmetrictableExist("citrixadc_lbmetrictable.tfAcc_lbmetrictable", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccLbmetrictable_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckLbmetrictableExist("citrixadc_lbmetrictable.tfAcc_lbmetrictable", nil)),
			},
		},
	})
}

const testAccLbmetrictableDataSource_basic = `

resource "citrixadc_lbmetrictable" "tf_lbmetrictable_ds" {
  metrictable = "tf_lbmetrictable_ds"
}

data "citrixadc_lbmetrictable" "tf_lbmetrictable_ds" {
  metrictable = citrixadc_lbmetrictable.tf_lbmetrictable_ds.metrictable
  depends_on  = [citrixadc_lbmetrictable.tf_lbmetrictable_ds]
}

`
