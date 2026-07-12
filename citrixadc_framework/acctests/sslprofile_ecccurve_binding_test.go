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
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccSslprofile_ecccurve_binding_basic = `
	resource "citrixadc_sslprofile" "tf_sslprofile" {
		name = "tf_sslprofile"
	}

	resource "citrixadc_sslprofile_ecccurve_binding" "tf_sslprofile_ecccurve_binding" {
		name                             = citrixadc_sslprofile.tf_sslprofile.name
		ecccurvename                     = ["X_25519", "P_521", "P_384"]
		remove_existing_ecccurve_binding = true
	}

`

const testAccSslprofile_ecccurve_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
	resource "citrixadc_sslprofile" "tf_sslprofile" {
		name = "tf_sslprofile"
	}
`

func TestAccSslprofile_ecccurve_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslprofile_ecccurve_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslprofile_ecccurve_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofile_ecccurve_bindingExist("citrixadc_sslprofile_ecccurve_binding.tf_sslprofile_ecccurve_binding", nil),
				),
			},
			{
				Config: testAccSslprofile_ecccurve_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofile_ecccurve_bindingNotExist("citrixadc_sslprofile_ecccurve_binding.tf_sslprofile_ecccurve_binding", "tf_sslprofile,"),
				),
			},
		},
	})
}

func testAccCheckSslprofile_ecccurve_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslprofile_ecccurve_binding id is set")
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

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 2)

		name := idSlice[0]
		ecccurvename := []string{"X_25519", "P_521", "P_384"}

		findParams := service.FindParams{
			ResourceType:             "sslprofile_ecccurve_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		var arr_ecccurve []string

		for _, v := range dataArr {
			arr_ecccurve = append(arr_ecccurve, v["ecccurvename"].(string))
		}

		if !equalStringSlices(ecccurvename, arr_ecccurve) {
			return fmt.Errorf("Expected 3 ecccurves, but found %d", len(arr_ecccurve))
		}

		return nil
	}
}

func equalStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func testAccCheckSslprofile_ecccurve_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		name := idSlice[0]
		// ecccurvename := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "sslprofile_ecccurve_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		if len(dataArr) != 0 {
			return fmt.Errorf("sslprofile_ecccurve_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckSslprofile_ecccurve_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslprofile_ecccurve_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Sslprofile_ecccurve_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslprofile_ecccurve_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccSslprofile_ecccurve_bindingDataSource_basic = `

resource "citrixadc_sslparameter" "tf_sslparam" {
	defaultprofile = "ENABLED"
}

resource "citrixadc_sslprofile" "tf_sslprofile" {
	name             = "tf_sslprofile_ds"
	ecccurvebindings = ["P_256", "P_384"]
	depends_on       = [citrixadc_sslparameter.tf_sslparam]
}

data "citrixadc_sslprofile_ecccurve_binding" "tf_sslprofile_ecccurve_binding" {
	name         = citrixadc_sslprofile.tf_sslprofile.name
	ecccurvename = "P_256"
}
`

func TestAccSslprofile_ecccurve_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslprofile_ecccurve_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslprofile_ecccurve_binding.tf_sslprofile_ecccurve_binding", "name", "tf_sslprofile_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_sslprofile_ecccurve_binding.tf_sslprofile_ecccurve_binding", "ecccurvename", "P_256"),
				),
			},
		},
	})
}

// testAccSslprofile_ecccurve_binding_upgrade_basic reuses the _basic config values
// (same resource labels, same profile name and ecccurve list) so it is valid under
// BOTH the last SDK v2 release (2.2.0) schema and the current Framework schema.
const testAccSslprofile_ecccurve_binding_upgrade_basic = `
	resource "citrixadc_sslprofile" "tf_sslprofile" {
		name = "tf_sslprofile"
	}

	resource "citrixadc_sslprofile_ecccurve_binding" "tf_sslprofile_ecccurve_binding" {
		name                             = citrixadc_sslprofile.tf_sslprofile.name
		ecccurvename                     = ["X_25519", "P_521", "P_384"]
		remove_existing_ecccurve_binding = true
	}
`

// TestAccSslprofile_ecccurve_binding_sdkv2StateUpgrade verifies that state written by the
// last SDK v2 release (with the legacy comma-composite id) upgrades cleanly through the
// current Framework provider, which recomputes the id to the new canonical form on Read.
func TestAccSslprofile_ecccurve_binding_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckSslprofile_ecccurve_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release (2.2.0). State carries the
			// legacy id: fmt.Sprintf("%s,%s", name, ecccurvename) => "tf_sslprofile,[X_25519 P_521 P_384]".
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccSslprofile_ecccurve_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofile_ecccurve_bindingExist("citrixadc_sslprofile_ecccurve_binding.tf_sslprofile_ecccurve_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile_ecccurve_binding.tf_sslprofile_ecccurve_binding", "id", "tf_sslprofile,[X_25519 P_521 P_384]"),
				),
			},
			// Step 2: refresh/plan/apply the SAME config through the current Framework provider.
			// Read (readSslprofileEcccurveBindingFromApi) recomputes data.Id to the new canonical
			// form (the plain SSL profile name), so the id upgrades from legacy to "tf_sslprofile".
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccSslprofile_ecccurve_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslprofile_ecccurve_bindingExist("citrixadc_sslprofile_ecccurve_binding.tf_sslprofile_ecccurve_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_sslprofile_ecccurve_binding.tf_sslprofile_ecccurve_binding", "id", "tf_sslprofile"),
				),
			},
		},
	})
}

func TestAccSslprofile_ecccurve_binding_import(t *testing.T) {
	const resAddr = "citrixadc_sslprofile_ecccurve_binding.tf_sslprofile_ecccurve_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslprofile_ecccurve_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccSslprofile_ecccurve_binding_basic},
			{Config: testAccSslprofile_ecccurve_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{"remove_existing_ecccurve_binding"}},
		},
	})
}
