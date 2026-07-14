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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// cloudparameter is a SINGLETON set-get resource (fixed id "cloudparameter-config").
// Create/Update = PUT (UpdateUnnamedResource), Read = FindResource(""), Delete = no-op state removal.
// No CheckDestroy is used because a singleton always exists on the ADC.
//
// NOTE: activationcode is NOT returned by GET. Its config value is preserved in state,
// so we do NOT assert it against the datasource (it is null there).
//
// TODO_PLACEHOLDER values below (controllerfqdn, instanceid, customerid, resourcelocation,
// activationcode) may be validated by the appliance in a real Citrix Cloud / NGS Connector
// onboarding flow. If the testbed rejects the sample values, replace them with valid
// testbed-specific values.

const testAccCloudparameter_basic_step1 = `
resource "citrixadc_cloudparameter" "tf_cloudparameter" {
  controllerfqdn     = "TODO_PLACEHOLDER" // e.g. "adm.cloud.com"
  controllerport     = 443
  deployment         = "Production"
  connectorresidence = "Aws"
}

`

const testAccCloudparameter_basic_step2 = `
resource "citrixadc_cloudparameter" "tf_cloudparameter" {
  controllerfqdn     = "TODO_PLACEHOLDER" // e.g. "adm.cloud.com"
  controllerport     = 8443
  deployment         = "Staging"
  connectorresidence = "Azure"
}

`

func TestAccCloudparameter_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudparameter_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudparameterExist("citrixadc_cloudparameter.tf_cloudparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_cloudparameter.tf_cloudparameter", "controllerfqdn", "TODO_PLACEHOLDER"),
					resource.TestCheckResourceAttr("citrixadc_cloudparameter.tf_cloudparameter", "controllerport", "443"),
					resource.TestCheckResourceAttr("citrixadc_cloudparameter.tf_cloudparameter", "deployment", "Production"),
					resource.TestCheckResourceAttr("citrixadc_cloudparameter.tf_cloudparameter", "connectorresidence", "Aws"),
				),
			},
			{
				Config: testAccCloudparameter_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudparameterExist("citrixadc_cloudparameter.tf_cloudparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_cloudparameter.tf_cloudparameter", "controllerfqdn", "TODO_PLACEHOLDER"),
					resource.TestCheckResourceAttr("citrixadc_cloudparameter.tf_cloudparameter", "controllerport", "8443"),
					resource.TestCheckResourceAttr("citrixadc_cloudparameter.tf_cloudparameter", "deployment", "Staging"),
					resource.TestCheckResourceAttr("citrixadc_cloudparameter.tf_cloudparameter", "connectorresidence", "Azure"),
				),
			},
		},
	})
}

func testAccCheckCloudparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cloudparameter name is set")
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
		// Singleton set-get resource: Read via FindResource with empty name.
		data, err := client.FindResource(service.Cloudparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("cloudparameter %s not found", n)
		}

		return nil
	}
}

const testAccCloudparameterDataSource_basic = `

resource "citrixadc_cloudparameter" "tf_cloudparameter" {
  controllerfqdn     = "TODO_PLACEHOLDER" // e.g. "adm.cloud.com"
  controllerport     = 443
  deployment         = "Production"
  connectorresidence = "Aws"
}

data "citrixadc_cloudparameter" "tf_cloudparameter" {
  depends_on = [citrixadc_cloudparameter.tf_cloudparameter]
}
`

func TestAccCloudparameterDataSource_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudparameterDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					// GET-readable attributes only. activationcode is NOT asserted here
					// because it is not returned by GET (null on the datasource).
					resource.TestCheckResourceAttr("data.citrixadc_cloudparameter.tf_cloudparameter", "controllerfqdn", "TODO_PLACEHOLDER"),
					resource.TestCheckResourceAttr("data.citrixadc_cloudparameter.tf_cloudparameter", "controllerport", "443"),
					resource.TestCheckResourceAttr("data.citrixadc_cloudparameter.tf_cloudparameter", "deployment", "Production"),
					resource.TestCheckResourceAttr("data.citrixadc_cloudparameter.tf_cloudparameter", "connectorresidence", "Aws"),
				),
			},
		},
	})
}
