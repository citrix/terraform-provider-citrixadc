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
	"reflect"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccNsparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNsparam_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsparamExist("citrixadc_nsparam.tf_nsparam", nil, map[string]interface{}{"maxconn": "10", "useproxyport": "DISABLED"}),
				),
			},
			{
				Config: testAccNsparam_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsparamExist("citrixadc_nsparam.tf_nsparam", nil, map[string]interface{}{"maxconn": "0", "useproxyport": "ENABLED"}),
				),
			},
			{
				Config: testAccNsparam_basic_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsparamExist("citrixadc_nsparam.tf_nsparam", nil, map[string]interface{}{"icaports": []int{84, 85}, "secureicaports": []int{8443, 9443}, "ipttl": 150}),
				),
			},
		},
	})
}

func testAccCheckNsparamExist(n string, id *string, expectedValues map[string]interface{}) resource.TestCheckFunc {
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

		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Nsparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("NS parameters %s not found", n)
		}

		// Iterate through all expected values and validate them
		for key, expectedValue := range expectedValues {
			if actualValue, exists := data[key]; !exists {
				return fmt.Errorf("Expected key %q not found in retrieved data", key)
			} else if !compareValues(expectedValue, actualValue) {
				return fmt.Errorf("Expected value for %q differs. Expected: %v, Retrieved: %v",
					key, expectedValue, actualValue)
			}
		}

		return nil
	}
}

// compareValues compares two values, handling slices and different types properly
func compareValues(expected, actual interface{}) bool {
	// Handle nil cases
	if expected == nil && actual == nil {
		return true
	}
	if expected == nil || actual == nil {
		return false
	}

	// Use reflect to get detailed type information
	expectedVal := reflect.ValueOf(expected)
	actualVal := reflect.ValueOf(actual)

	// If both are slices, compare them element by element
	if expectedVal.Kind() == reflect.Slice && actualVal.Kind() == reflect.Slice {
		if expectedVal.Len() != actualVal.Len() {
			return false
		}

		// Compare each element
		for i := 0; i < expectedVal.Len(); i++ {
			expectedElem := expectedVal.Index(i).Interface()
			actualElem := actualVal.Index(i).Interface()

			// Convert both to strings for comparison (since API might return strings)
			expectedStr := fmt.Sprintf("%v", expectedElem)
			actualStr := fmt.Sprintf("%v", actualElem)

			if expectedStr != actualStr {
				return false
			}
		}
		return true
	}

	// For non-slice types, convert both to strings and compare
	// This handles the case where API returns strings but we expect ints
	expectedStr := fmt.Sprintf("%v", expected)
	actualStr := fmt.Sprintf("%v", actual)

	return expectedStr == actualStr
}

const testAccNsparam_basic_step1 = `

resource "citrixadc_nsparam" "tf_nsparam" {
  maxconn = 10
  useproxyport = "DISABLED"
}
`

const testAccNsparam_basic_step2 = `

resource "citrixadc_nsparam" "tf_nsparam" {
  maxconn = 0
  useproxyport = "ENABLED"
}
`
const testAccNsparam_basic_step3 = `

resource "citrixadc_nsparam" "tf_nsparam" {
  icaports = [ 84, 85]
  secureicaports = [ 8443, 9443]
  ipttl = 150
}
`

func TestAccNsparamDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNsparamDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.citrixadc_nsparam.test", "id"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nsparam.test", "timezone"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nsparam.test", "advancedanalyticsstats"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nsparam.test", "aftpallowrandomsourceport"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nsparam.test", "internaluserlogin"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nsparam.test", "ipttl"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nsparam.test", "useproxyport"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nsparam.test", "securecookie"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nsparam.test", "proxyprotocol"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nsparam.test", "tcpcip"),
				),
			},
		},
	})
}

const testAccNsparamDataSource_basic = `
data "citrixadc_nsparam" "test" {
}
`
