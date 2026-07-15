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

// cloudtunnelparameter is a SINGLETON set-get resource (fixed id "cloudtunnelparameter-config").
// Create/Update = PUT (UpdateUnnamedResource), Read = FindResource(""), Delete = no-op state removal.
// No CheckDestroy is used because a singleton always exists on the ADC.
//
// TODO_PLACEHOLDER (TESTBED PREREQUISITE): This resource is gated behind a cloud-tunnel
// license/feature. On testbeds where the feature is NOT enabled, the ADC returns
// "ERROR: Feature not supported in this release". The resource and datasource Read paths
// tolerate this gracefully, but the acceptance test as a whole requires a cloud-tunnel
// license/feature to be provisioned on the appliance before it can pass. If the testbed
// does not have this feature, the whole test should be skipped/gated.
//
// TODO_PLACEHOLDER (EXIST CHECK): testAccCheckCloudtunnelparameterExist calls
// FindResource(service.Cloudtunnelparameter.Type(), ""). On a testbed where the cloud-tunnel
// feature is unlicensed/unsupported this FindResource call may error out ("Feature not
// supported in this release"). If so, adjust or skip the exist-check (e.g. tolerate the
// feature-not-supported error, or remove the exist-check from the Check funcs) once you know
// whether the target testbed has the feature enabled.
//
// The 4 attributes (controllerfqdn, fqdn, resourcelocation, subnetresourcelocationmappings)
// are free-form strings; the sample FQDN/values below are plausible but may need to be
// replaced with valid testbed-specific values (marked TODO_PLACEHOLDER inline).

const testAccCloudtunnelparameter_basic_step1 = `
resource "citrixadc_cloudtunnelparameter" "tf_cloudtunnelparameter" {
  controllerfqdn   = "controller.cloud.example.com" // TODO_PLACEHOLDER: use a valid controller FQDN for the testbed
  fqdn             = "tunnel.cloud.example.com"      // TODO_PLACEHOLDER: use a valid tunnel FQDN for the testbed
  resourcelocation = "rl-primary"                    // TODO_PLACEHOLDER: use a valid resource-location id for the testbed
}

`

const testAccCloudtunnelparameter_basic_step2 = `
resource "citrixadc_cloudtunnelparameter" "tf_cloudtunnelparameter" {
  controllerfqdn   = "controller.cloud.example.com" // TODO_PLACEHOLDER: use a valid controller FQDN for the testbed
  fqdn             = "tunnel.cloud.example.com"      // TODO_PLACEHOLDER: use a valid tunnel FQDN for the testbed
  resourcelocation = "rl-secondary"                  // TODO_PLACEHOLDER: use a valid resource-location id for the testbed
}

`

func TestAccCloudtunnelparameter_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	// TODO_PLACEHOLDER (PREREQUISITE): requires a cloud-tunnel license/feature on the testbed.
	// If the target testbed does not have the feature enabled, skip this test.
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudtunnelparameter_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudtunnelparameterExist("citrixadc_cloudtunnelparameter.tf_cloudtunnelparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_cloudtunnelparameter.tf_cloudtunnelparameter", "controllerfqdn", "controller.cloud.example.com"),
					resource.TestCheckResourceAttr("citrixadc_cloudtunnelparameter.tf_cloudtunnelparameter", "fqdn", "tunnel.cloud.example.com"),
					resource.TestCheckResourceAttr("citrixadc_cloudtunnelparameter.tf_cloudtunnelparameter", "resourcelocation", "rl-primary"),
				),
			},
			{
				Config: testAccCloudtunnelparameter_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudtunnelparameterExist("citrixadc_cloudtunnelparameter.tf_cloudtunnelparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_cloudtunnelparameter.tf_cloudtunnelparameter", "controllerfqdn", "controller.cloud.example.com"),
					resource.TestCheckResourceAttr("citrixadc_cloudtunnelparameter.tf_cloudtunnelparameter", "fqdn", "tunnel.cloud.example.com"),
					resource.TestCheckResourceAttr("citrixadc_cloudtunnelparameter.tf_cloudtunnelparameter", "resourcelocation", "rl-secondary"),
				),
			},
		},
	})
}

func TestAccCloudtunnelparameter_import(t *testing.T) {
	t.Skip("TODO: Requires review")
	// cloudtunnelparameter is a singleton whose Terraform id is a synthetic constant
	// ("cloudtunnelparameter-config"). ImportStatePassthroughID uses that stored id, so
	// no ImportStateIdFunc is needed.
	const resAddr = "citrixadc_cloudtunnelparameter.tf_cloudtunnelparameter"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{Config: testAccCloudtunnelparameter_basic_step1},
			{
				Config:                  testAccCloudtunnelparameter_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckCloudtunnelparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cloudtunnelparameter name is set")
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
		//
		// TODO_PLACEHOLDER (TESTBED): On a testbed where the cloud-tunnel feature is
		// unlicensed/unsupported, this FindResource may return an error such as
		// "ERROR: Feature not supported in this release". If the target testbed does not
		// have the feature enabled, either skip this exist-check or tolerate that specific
		// feature-not-supported error here.
		data, err := client.FindResource(service.Cloudtunnelparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("cloudtunnelparameter %s not found", n)
		}

		return nil
	}
}

const testAccCloudtunnelparameterDataSource_basic = `

resource "citrixadc_cloudtunnelparameter" "tf_cloudtunnelparameter" {
  controllerfqdn   = "controller.cloud.example.com" // TODO_PLACEHOLDER: use a valid controller FQDN for the testbed
  fqdn             = "tunnel.cloud.example.com"      // TODO_PLACEHOLDER: use a valid tunnel FQDN for the testbed
  resourcelocation = "rl-primary"                    // TODO_PLACEHOLDER: use a valid resource-location id for the testbed
}

data "citrixadc_cloudtunnelparameter" "tf_cloudtunnelparameter" {
  depends_on = [citrixadc_cloudtunnelparameter.tf_cloudtunnelparameter]
}
`

func TestAccCloudtunnelparameterDataSource_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	// TODO_PLACEHOLDER (TESTBED / FEATURE-GATED): This datasource reads a singleton that is
	// gated behind a cloud-tunnel license/feature. On a testbed where the feature is NOT
	// enabled the GET tolerates the failure gracefully and the datasource attributes may come
	// back empty/null, so the TestCheckResourceAttr assertions below will not match. Enable the
	// cloud-tunnel feature on the testbed before expecting these assertions to pass, or skip
	// this test on testbeds without the feature.
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudtunnelparameterDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_cloudtunnelparameter.tf_cloudtunnelparameter", "controllerfqdn", "controller.cloud.example.com"),
					resource.TestCheckResourceAttr("data.citrixadc_cloudtunnelparameter.tf_cloudtunnelparameter", "fqdn", "tunnel.cloud.example.com"),
					resource.TestCheckResourceAttr("data.citrixadc_cloudtunnelparameter.tf_cloudtunnelparameter", "resourcelocation", "rl-primary"),
				),
			},
		},
	})
}
