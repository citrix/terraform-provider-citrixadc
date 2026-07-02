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

// callhome is a SINGLETON set-get parameter resource:
//   - Create/Update = UpdateUnnamedResource (set), Read = find-unnamed, Delete = no-op.
//   - Fixed ID "callhome"; the object always exists on the ADC and is never deleted.
//   - ipaddress and proxyauthservice are mutually exclusive, so they are never
//     configured together in these tests.

const testAccCallhome_basic_step1 = `
resource "citrixadc_callhome" "tf_callhome" {
  mode             = "Default"
  hbcustominterval = 10
  emailaddress     = "test@example.com"
  proxymode        = "NO"
}

`

func TestAccCallhome_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// Singleton: object always exists on the ADC, so the destroy check only
		// verifies the object is still present (it is never actually deleted).
		CheckDestroy: testAccCheckCallhomeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCallhome_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCallhomeExist("citrixadc_callhome.tf_callhome", nil),
					resource.TestCheckResourceAttr("citrixadc_callhome.tf_callhome", "mode", "Default"),
					resource.TestCheckResourceAttr("citrixadc_callhome.tf_callhome", "hbcustominterval", "10"),
					resource.TestCheckResourceAttr("citrixadc_callhome.tf_callhome", "emailaddress", "test@example.com"),
					resource.TestCheckResourceAttr("citrixadc_callhome.tf_callhome", "proxymode", "NO"),
				),
			},
		},
	})
}

func testAccCheckCallhomeExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No callhome ID is set")
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
		// Singleton set-get resource: read via find-unnamed (empty name).
		data, err := client.FindResource(service.Callhome.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("callhome %s not found", n)
		}

		return nil
	}
}

// Singleton destroy check: callhome is never deleted on the ADC, so this only
// confirms the object still exists after the config is destroyed (removed from
// Terraform state). It does NOT assert absence.
func testAccCheckCallhomeDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_callhome" {
			continue
		}
		// Singleton resource always exists on the ADC; a successful read is expected.
		_, err := client.FindResource(service.Callhome.Type(), "")
		if err != nil {
			return fmt.Errorf("callhome singleton unexpectedly missing after destroy: %v", err)
		}
	}
	return nil
}

const testAccCallhomeDataSource_basic = `

resource "citrixadc_callhome" "tf_callhome" {
  mode             = "Default"
  hbcustominterval = 10
  emailaddress     = "test@example.com"
  proxymode        = "NO"
}

data "citrixadc_callhome" "tf_callhome" {
  depends_on = [citrixadc_callhome.tf_callhome]
}
`

func TestAccCallhomeDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCallhomeDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_callhome.tf_callhome", "mode", "Default"),
					resource.TestCheckResourceAttr("data.citrixadc_callhome.tf_callhome", "hbcustominterval", "10"),
					resource.TestCheckResourceAttr("data.citrixadc_callhome.tf_callhome", "emailaddress", "test@example.com"),
					resource.TestCheckResourceAttr("data.citrixadc_callhome.tf_callhome", "proxymode", "NO"),
				),
			},
		},
	})
}
