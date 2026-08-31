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

const testAccNssimpleacl6_basic = `

	resource "citrixadc_nssimpleacl6" "tf_nssimpleacl6" {
		aclname   = "tf_nssimpleacl6"
		aclaction = "DENY"
		srcipv6   = "3ffe:192:168:215::82"
		destport  = 123
		protocol  = "TCP"
		ttl       = 600
	}
`

func TestAccNssimpleacl6_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNssimpleacl6Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNssimpleacl6_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNssimpleacl6Exist("citrixadc_nssimpleacl6.tf_nssimpleacl6", nil),
					resource.TestCheckResourceAttr("citrixadc_nssimpleacl6.tf_nssimpleacl6", "aclname", "tf_nssimpleacl6"),
					resource.TestCheckResourceAttr("citrixadc_nssimpleacl6.tf_nssimpleacl6", "aclaction", "DENY"),
					resource.TestCheckResourceAttr("citrixadc_nssimpleacl6.tf_nssimpleacl6", "srcipv6", "3ffe:192:168:215::82"),
				),
			},
		},
	})
}

func TestAccNssimpleacl6_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_nssimpleacl6.tf_nssimpleacl6"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNssimpleacl6Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNssimpleacl6_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNssimpleacl6Exist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Nssimpleacl6.Type(), "tf_nssimpleacl6"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccNssimpleacl6_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNssimpleacl6Exist(resAddr, nil)),
			},
		},
	})
}

func testAccCheckNssimpleacl6Exist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nssimpleacl6 name is set")
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
		data, err := client.FindResource(service.Nssimpleacl6.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nssimpleacl6 %s not found", n)
		}

		return nil
	}
}

func testAccCheckNssimpleacl6Destroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nssimpleacl6" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Nssimpleacl6.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nssimpleacl6 %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccNssimpleacl6_import(t *testing.T) {
	const resAddr = "citrixadc_nssimpleacl6.tf_nssimpleacl6"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNssimpleacl6Destroy,
		Steps: []resource.TestStep{
			{Config: testAccNssimpleacl6_basic},
			{
				Config:                  testAccNssimpleacl6_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ttl"},
			},
		},
	})
}

func TestAccNssimpleacl6_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNssimpleacl6Destroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccNssimpleacl6_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNssimpleacl6Exist("citrixadc_nssimpleacl6.tf_nssimpleacl6", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccNssimpleacl6_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNssimpleacl6Exist("citrixadc_nssimpleacl6.tf_nssimpleacl6", nil)),
			},
		},
	})
}

func TestAccNssimpleacl6DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNssimpleacl6DataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nssimpleacl6.tf_simpleacl6_ds", "aclname", "tf_simpleacl6_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_nssimpleacl6.tf_simpleacl6_ds", "aclaction", "DENY"),
				),
			},
		},
	})
}

const testAccNssimpleacl6DataSource_basic = `

	resource "citrixadc_nssimpleacl6" "tf_simpleacl6_ds" {
		aclname   = "tf_simpleacl6_ds"
		aclaction = "DENY"
		srcipv6   = "2001:db8:1::1"
		protocol  = "TCP"
		destport  = 443
		ttl       = 7200
	}

	data "citrixadc_nssimpleacl6" "tf_simpleacl6_ds" {
		aclname    = citrixadc_nssimpleacl6.tf_simpleacl6_ds.aclname
		depends_on = [citrixadc_nssimpleacl6.tf_simpleacl6_ds]
	}
`
