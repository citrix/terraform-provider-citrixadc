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
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// NOTE: vpnglobal is a singleton on the ADC, so there is no parent resource to create.
// secureprivateaccessprofile is the binding key and is a reference to a Secure Private Access
// Profile. The profile is created inline by the citrixadc_vpnsecureprivateaccessprofile
// resource (name tf_spa_profile), so the binding no longer requires a pre-staged SPA profile
// on the ADC.
const testAccVpnglobalVpnsecureprivateaccessprofileBinding_prereq = `
	resource "citrixadc_vpnsecureprivateaccessprofile" "tf_spa_profile" {
		name                        = "tf_spa_profile"
		url                         = "https://spa.example.com"
		forceclienttype             = "ON"
		chromeenterprisepremiummode = "OFF"
	}
`

const testAccVpnglobalVpnsecureprivateaccessprofileBinding_basic_step1 = testAccVpnglobalVpnsecureprivateaccessprofileBinding_prereq + `
	resource "citrixadc_vpnglobal_vpnsecureprivateaccessprofile_binding" "tf_binding" {
		secureprivateaccessprofile = citrixadc_vpnsecureprivateaccessprofile.tf_spa_profile.name
	}
`

// Step 2 drops the binding but keeps the profile so the binding delete is exercised.
const testAccVpnglobalVpnsecureprivateaccessprofileBinding_basic_step2 = testAccVpnglobalVpnsecureprivateaccessprofileBinding_prereq

func TestAccVpnglobalVpnsecureprivateaccessprofileBinding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnglobalVpnsecureprivateaccessprofileBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnglobalVpnsecureprivateaccessprofileBinding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnglobalVpnsecureprivateaccessprofileBindingExist("citrixadc_vpnglobal_vpnsecureprivateaccessprofile_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnglobal_vpnsecureprivateaccessprofile_binding.tf_binding", "secureprivateaccessprofile", "tf_spa_profile"),
				),
			},
			{
				Config: testAccVpnglobalVpnsecureprivateaccessprofileBinding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnglobalVpnsecureprivateaccessprofileBindingNotExist("citrixadc_vpnglobal_vpnsecureprivateaccessprofile_binding.tf_binding", "tf_spa_profile"),
				),
			},
		},
	})
}

func TestAccVpnglobalVpnsecureprivateaccessprofileBinding_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_vpnglobal_vpnsecureprivateaccessprofile_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnglobalVpnsecureprivateaccessprofileBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnglobalVpnsecureprivateaccessprofileBinding_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnglobalVpnsecureprivateaccessprofileBindingExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgs(service.Vpnglobal_vpnsecureprivateaccessprofile_binding.Type(), "", []string{"secureprivateaccessprofile:tf_spa_profile"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccVpnglobalVpnsecureprivateaccessprofileBinding_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckVpnglobalVpnsecureprivateaccessprofileBindingExist(resAddr, nil)),
			},
		},
	})
}

func TestAccVpnglobalVpnsecureprivateaccessprofileBinding_import(t *testing.T) {
	const resAddr = "citrixadc_vpnglobal_vpnsecureprivateaccessprofile_binding.tf_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnglobalVpnsecureprivateaccessprofileBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnglobalVpnsecureprivateaccessprofileBinding_basic_step1,
			},
			{
				Config:                  testAccVpnglobalVpnsecureprivateaccessprofileBinding_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckVpnglobalVpnsecureprivateaccessprofileBindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnglobal_vpnsecureprivateaccessprofile_binding id is set")
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

		// ID is a plain value (single unique attr: secureprivateaccessprofile)
		secureprivateaccessprofile := rs.Primary.ID

		findParams := service.FindParams{
			ResourceType:             service.Vpnglobal_vpnsecureprivateaccessprofile_binding.Type(),
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching secureprivateaccessprofile
		found := false
		for _, v := range dataArr {
			if val, ok := v["secureprivateaccessprofile"].(string); ok && val == secureprivateaccessprofile {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("vpnglobal_vpnsecureprivateaccessprofile_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnglobalVpnsecureprivateaccessprofileBindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		secureprivateaccessprofile := id

		findParams := service.FindParams{
			ResourceType:             service.Vpnglobal_vpnsecureprivateaccessprofile_binding.Type(),
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// A missing-resource error means the binding is gone, which is what we want.
		if err != nil {
			return nil
		}

		// Iterate through results to hopefully not find the one with the matching profile
		found := false
		for _, v := range dataArr {
			if val, ok := v["secureprivateaccessprofile"].(string); ok && val == secureprivateaccessprofile {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("vpnglobal_vpnsecureprivateaccessprofile_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckVpnglobalVpnsecureprivateaccessprofileBindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnglobal_vpnsecureprivateaccessprofile_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		secureprivateaccessprofile := rs.Primary.ID

		findParams := service.FindParams{
			ResourceType:             service.Vpnglobal_vpnsecureprivateaccessprofile_binding.Type(),
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// A missing-resource error means the binding is gone, which is what we want.
			continue
		}

		for _, v := range dataArr {
			if val, ok := v["secureprivateaccessprofile"].(string); ok && val == secureprivateaccessprofile {
				return fmt.Errorf("vpnglobal_vpnsecureprivateaccessprofile_binding %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

const testAccVpnglobalVpnsecureprivateaccessprofileBindingDataSource_basic = testAccVpnglobalVpnsecureprivateaccessprofileBinding_prereq + `
	resource "citrixadc_vpnglobal_vpnsecureprivateaccessprofile_binding" "tf_binding" {
		secureprivateaccessprofile = citrixadc_vpnsecureprivateaccessprofile.tf_spa_profile.name
	}

	data "citrixadc_vpnglobal_vpnsecureprivateaccessprofile_binding" "tf_binding" {
		secureprivateaccessprofile = citrixadc_vpnglobal_vpnsecureprivateaccessprofile_binding.tf_binding.secureprivateaccessprofile
		depends_on                 = [citrixadc_vpnglobal_vpnsecureprivateaccessprofile_binding.tf_binding]
	}
`

func TestAccVpnglobalVpnsecureprivateaccessprofileBindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnglobalVpnsecureprivateaccessprofileBindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_vpnglobal_vpnsecureprivateaccessprofile_binding.tf_binding", "secureprivateaccessprofile", "tf_spa_profile"),
				),
			},
		},
	})
}
