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

// cloudawsparam is a singleton set/get parameter resource (fixed id "cloudawsparam-config").
// No CheckDestroy is used: singleton parameter resources always exist on the ADC and cannot be deleted.
//
// NOTE: rolearn is a standard AWS IAM Role ARN (arn:aws:iam::<account-id>:role/<role-name>).
// Verified against the live AWS-hosted ADC (3.239.99.1): the appliance accepts the ARN string
// as-is (errorcode 0) without validating that the IAM role actually exists/assumable, and the
// GET echoes rolearn back, so the value round-trips into state.

const testAccCloudawsparam_basic_step1 = `
resource "citrixadc_cloudawsparam" "tf_cloudawsparam" {
  rolearn = "arn:aws:iam::123456789012:role/example"
}

`

func TestAccCloudawsparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudawsparam_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudawsparamExist("citrixadc_cloudawsparam.tf_cloudawsparam", nil),
					resource.TestCheckResourceAttr("citrixadc_cloudawsparam.tf_cloudawsparam", "rolearn", "arn:aws:iam::123456789012:role/example"),
				),
			},
		},
	})
}

func testAccCheckCloudawsparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cloudawsparam name is set")
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
		// Singleton resource - find without an ID
		data, err := client.FindResource(service.Cloudawsparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("cloudawsparam %s not found", n)
		}

		return nil
	}
}

const testAccCloudawsparamDataSource_basic = `

resource "citrixadc_cloudawsparam" "tf_cloudawsparam" {
  rolearn = "arn:aws:iam::123456789012:role/example"
}

data "citrixadc_cloudawsparam" "tf_cloudawsparam" {
  depends_on = [citrixadc_cloudawsparam.tf_cloudawsparam]
}
`

func TestAccCloudawsparamDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudawsparamDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_cloudawsparam.tf_cloudawsparam", "rolearn", "arn:aws:iam::123456789012:role/example"),
				),
			},
		},
	})
}
