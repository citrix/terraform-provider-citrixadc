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

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccNsfeature_basic(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNsfeature_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEnabledDisabledFeatures([]string{"cs", "lb"}, []string{"ssl", "appfw"}),
				),
			},
			resource.TestStep{
				Config: testAccNsfeature_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEnabledDisabledFeatures([]string{"cs"}, []string{"ssl", "appfw", "lb"}),
				),
			},
			resource.TestStep{
				Config: testAccNsfeature_basic_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEnabledDisabledFeatures([]string{"cs", "ssl"}, []string{"appfw", "lb"}),
				),
			},
			resource.TestStep{
				Config: testAccNsfeature_basic_step4,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEnabledDisabledFeatures([]string{"appfw", "lb"}, []string{"cs", "ssl"}),
				),
			},
		},
	})
}

func testAccCheckEnabledDisabledFeatures(enabledFeatures, disabledFeatures []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := client.ListEnabledFeatures()
		if err != nil {
			return err
		}
		featuresRead := make([]string, len(data))
		for i, val := range data {
			featuresRead[i] = strings.ToLower(val)
		}

		// Find enabled
		for _, enabledFeature := range enabledFeatures {
			found := false
			for _, featureRead := range featuresRead {
				if featureRead == enabledFeature {
					found = true
					break
				}
			}
			if !found {
				return fmt.Errorf("Feature should be enabled %v", enabledFeature)
			}
		}

		// Find disabled
		for _, disabledFeature := range disabledFeatures {
			found := false
			for _, featureRead := range featuresRead {
				if featureRead == disabledFeature {
					found = true
					break
				}
			}
			if found {
				return fmt.Errorf("Feature should be disabled %v", disabledFeature)
			}
		}

		return nil
	}

}

const testAccNsfeature_basic_step1 = `
resource "citrixadc_nsfeature" "tf_nsfeature" {
    cs = true
    lb = true
    ssl = false
    appfw = false
}

`

const testAccNsfeature_basic_step2 = `
resource "citrixadc_nsfeature" "tf_nsfeature" {
    cs = true
    lb = false
    ssl = false
    appfw = false
}

`

const testAccNsfeature_basic_step3 = `
resource "citrixadc_nsfeature" "tf_nsfeature" {
    cs = true
    lb = false
    ssl = true
    appfw = false
}

`

const testAccNsfeature_basic_step4 = `
resource "citrixadc_nsfeature" "tf_nsfeature" {
    cs = false
    lb = true
    ssl = false
    appfw = true
}

`
