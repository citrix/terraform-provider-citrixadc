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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// apiprofile_apispec_binding is a binding_with_parent resource.
//   - Parent (URL key): name  -> the apiprofile name.
//   - Bound entity (array filter + delete arg): apispec -> the apispec name.
//   - Composite ID = "apispec:<UrlEncode>,name:<UrlEncode>".
//
// Both name and apispec are RequiresReplace; there is no update path. The test
// reuses the participating-entity config lifted from apiprofile_test.go and
// apispec_test.go (which in turn requires apispecfile + the uploaded sample
// spec files). The PreCheck therefore uses doApiSpecPreChecks to upload the
// sample_apispec*.yaml files that the apispec/apispecfile resources consume.
//
// Dependency chain:
//   apispecfile (uploads sample spec) -> apispec -> apiprofile (independent)
//   -> apiprofile_apispec_binding (references apiprofile.name + apispec.name).

const testAccApiprofileApispecBinding_basic_step1 = `
resource "citrixadc_apispecfile" "tf_apispecfile" {
  name      = "test_apispecfile"
  src       = "local://sample_apispec.yaml"
  overwrite = true
}

resource "citrixadc_apispec" "tf_apispec" {
  name = "tf_apispec"
  file = citrixadc_apispecfile.tf_apispecfile.name
  type = "OAS"
}

resource "citrixadc_apiprofile" "tf_apiprofile" {
  name          = "test_apiprofile"
  apivisibility = "ENABLED"
}

resource "citrixadc_apiprofile_apispec_binding" "tf_apiprofile_apispec_binding" {
  name       = citrixadc_apiprofile.tf_apiprofile.name
  apispec    = citrixadc_apispec.tf_apispec.name
  depends_on = [citrixadc_apiprofile.tf_apiprofile, citrixadc_apispec.tf_apispec]
}
`

// step2: drop the binding (keep the participating entities) to verify the
// binding is deleted from the ADC.
const testAccApiprofileApispecBinding_basic_step2 = `
resource "citrixadc_apispecfile" "tf_apispecfile" {
  name      = "test_apispecfile"
  src       = "local://sample_apispec.yaml"
  overwrite = true
}

resource "citrixadc_apispec" "tf_apispec" {
  name = "tf_apispec"
  file = citrixadc_apispecfile.tf_apispecfile.name
  type = "OAS"
}

resource "citrixadc_apiprofile" "tf_apiprofile" {
  name          = "test_apiprofile"
  apivisibility = "ENABLED"
}
`

func TestAccApiprofileApispecBinding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doApiSpecPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckApiprofileApispecBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApiprofileApispecBinding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApiprofileApispecBindingExist("citrixadc_apiprofile_apispec_binding.tf_apiprofile_apispec_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_apiprofile_apispec_binding.tf_apiprofile_apispec_binding", "name", "test_apiprofile"),
					resource.TestCheckResourceAttr("citrixadc_apiprofile_apispec_binding.tf_apiprofile_apispec_binding", "apispec", "tf_apispec"),
				),
			},
			{
				// Binding dropped from config: verify it no longer exists on the ADC.
				Config: testAccApiprofileApispecBinding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApiprofileApispecBindingNotExist(t, "test_apiprofile", "tf_apispec"),
				),
			},
		},
	})
}

func TestAccApiprofileApispecBinding_import(t *testing.T) {
	const resAddr = "citrixadc_apiprofile_apispec_binding.tf_apiprofile_apispec_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doApiSpecPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckApiprofileApispecBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccApiprofileApispecBinding_basic_step1,
			},
			{
				Config:                  testAccApiprofileApispecBinding_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckApiprofileApispecBindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No apiprofile_apispec_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Composite ID = "apispec:<value>,name:<value>".
		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}
		nameValue := idMap["name"]
		apispecValue := idMap["apispec"]

		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		// Mirror the resource Read: list bindings under the parent name, filter on apispec.
		findParams := service.FindParams{
			ResourceType:             service.Apiprofile_apispec_binding.Type(),
			ResourceName:             nameValue,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if val, ok := v["apispec"].(string); ok && val == apispecValue {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("apiprofile_apispec_binding %s not found", rs.Primary.ID)
		}

		return nil
	}
}

// testAccCheckApiprofileApispecBindingNotExist verifies that the binding for the
// given apiprofile name + apispec is absent from the ADC.
func testAccCheckApiprofileApispecBindingNotExist(t *testing.T, name string, apispec string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		findParams := service.FindParams{
			ResourceType:             service.Apiprofile_apispec_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Treat a missing/empty resource as successfully unbound.
			return nil
		}

		for _, v := range dataArr {
			if val, ok := v["apispec"].(string); ok && val == apispec {
				return fmt.Errorf("apiprofile_apispec_binding (name=%s, apispec=%s) still exists after being dropped from config", name, apispec)
			}
		}

		return nil
	}
}

func testAccCheckApiprofileApispecBindingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_apiprofile_apispec_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No id is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}
		nameValue := idMap["name"]
		apispecValue := idMap["apispec"]

		findParams := service.FindParams{
			ResourceType:             service.Apiprofile_apispec_binding.Type(),
			ResourceName:             nameValue,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Treat a missing/empty resource as successfully destroyed.
			continue
		}

		for _, v := range dataArr {
			if val, ok := v["apispec"].(string); ok && val == apispecValue {
				return fmt.Errorf("apiprofile_apispec_binding %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

const testAccApiprofileApispecBindingDataSource_basic = `
resource "citrixadc_apispecfile" "tf_apispecfile" {
  name      = "test_apispecfile"
  src       = "local://sample_apispec.yaml"
  overwrite = true
}

resource "citrixadc_apispec" "tf_apispec" {
  name = "tf_apispec"
  file = citrixadc_apispecfile.tf_apispecfile.name
  type = "OAS"
}

resource "citrixadc_apiprofile" "tf_apiprofile" {
  name          = "test_apiprofile"
  apivisibility = "ENABLED"
}

resource "citrixadc_apiprofile_apispec_binding" "tf_apiprofile_apispec_binding" {
  name       = citrixadc_apiprofile.tf_apiprofile.name
  apispec    = citrixadc_apispec.tf_apispec.name
  depends_on = [citrixadc_apiprofile.tf_apiprofile, citrixadc_apispec.tf_apispec]
}

data "citrixadc_apiprofile_apispec_binding" "tf_apiprofile_apispec_binding" {
  name       = citrixadc_apiprofile_apispec_binding.tf_apiprofile_apispec_binding.name
  apispec    = citrixadc_apiprofile_apispec_binding.tf_apiprofile_apispec_binding.apispec
  depends_on = [citrixadc_apiprofile_apispec_binding.tf_apiprofile_apispec_binding]
}
`

func TestAccApiprofileApispecBindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doApiSpecPreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccApiprofileApispecBindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_apiprofile_apispec_binding.tf_apiprofile_apispec_binding", "name", "test_apiprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_apiprofile_apispec_binding.tf_apiprofile_apispec_binding", "apispec", "tf_apispec"),
				),
			},
		},
	})
}
