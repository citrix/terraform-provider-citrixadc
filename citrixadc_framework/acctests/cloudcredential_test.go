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

// cloudcredential is a SINGLETON set-get resource.
// - Create/Update  = PUT (UpdateUnnamedResource)
// - Read           = FindResource("") used only to confirm existence
// - Delete         = no-op (state removal only)
// The NITRO GET returns ONLY "isset"; it does NOT echo back tenantidentifier,
// applicationid, or applicationsecret. State therefore retains the config
// values (asserted below), but the appliance cannot be used to detect drift on
// these attributes. No datasource exists (removed under Pattern 13), so no
// datasource test is generated. No CheckDestroy for singletons.

const testAccCloudcredential_basic_step1 = `
resource "citrixadc_cloudcredential" "tf_cloudcredential" {
  tenantidentifier  = "11111111-1111-1111-1111-111111111111"
  applicationid     = "22222222-2222-2222-2222-222222222222"
  applicationsecret = "cloudsecret_step1"
}

`

const testAccCloudcredential_basic_step2 = `
resource "citrixadc_cloudcredential" "tf_cloudcredential" {
  tenantidentifier  = "11111111-1111-1111-1111-111111111111"
  applicationid     = "33333333-3333-3333-3333-333333333333"
  applicationsecret = "cloudsecret_step1"
}

`

func TestAccCloudcredential_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudcredential_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudcredentialExist("citrixadc_cloudcredential.tf_cloudcredential", nil),
					resource.TestCheckResourceAttr("citrixadc_cloudcredential.tf_cloudcredential", "tenantidentifier", "11111111-1111-1111-1111-111111111111"),
					resource.TestCheckResourceAttr("citrixadc_cloudcredential.tf_cloudcredential", "applicationid", "22222222-2222-2222-2222-222222222222"),
					resource.TestCheckResourceAttr("citrixadc_cloudcredential.tf_cloudcredential", "applicationsecret", "cloudsecret_step1"),
				),
			},
			{
				Config: testAccCloudcredential_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudcredentialExist("citrixadc_cloudcredential.tf_cloudcredential", nil),
					resource.TestCheckResourceAttr("citrixadc_cloudcredential.tf_cloudcredential", "tenantidentifier", "11111111-1111-1111-1111-111111111111"),
					resource.TestCheckResourceAttr("citrixadc_cloudcredential.tf_cloudcredential", "applicationid", "33333333-3333-3333-3333-333333333333"),
					resource.TestCheckResourceAttr("citrixadc_cloudcredential.tf_cloudcredential", "applicationsecret", "cloudsecret_step1"),
				),
			},
		},
	})
}

func testAccCheckCloudcredentialExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cloudcredential name is set")
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
		// Singleton: GET takes no name; the response only confirms existence.
		data, err := client.FindResource(service.Cloudcredential.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("cloudcredential %s not found", n)
		}

		return nil
	}
}

// Test ephemeral path: using applicationsecret_wo (WriteOnly attribute) with version tracker.
// Bump applicationsecret_wo_version between steps to trigger an update.
const testAccCloudcredential_applicationsecret_wo_step1 = `

	variable "cloudcredential_applicationsecret_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_cloudcredential" "tf_cloudcredential" {
		tenantidentifier             = "11111111-1111-1111-1111-111111111111"
		applicationid                = "22222222-2222-2222-2222-222222222222"
		applicationsecret_wo         = var.cloudcredential_applicationsecret_wo
		applicationsecret_wo_version = 1
	}
`

// Update ephemeral path: bump version to trigger update with new secret.
const testAccCloudcredential_applicationsecret_wo_step2 = `

	 variable "cloudcredential_applicationsecret_wo_2" {
	  type      = string
	  sensitive = true
	}

	 resource "citrixadc_cloudcredential" "tf_cloudcredential" {
		tenantidentifier             = "11111111-1111-1111-1111-111111111111"
		applicationid                = "22222222-2222-2222-2222-222222222222"
		applicationsecret_wo         = var.cloudcredential_applicationsecret_wo_2
		applicationsecret_wo_version = 2
	}
`

func TestAccCloudcredential_applicationsecret_wo_ephemeral(t *testing.T) {
	t.Skip("TODO: Requires review")
	t.Setenv("TF_VAR_cloudcredential_applicationsecret_wo", "ephemeral_secret1")
	t.Setenv("TF_VAR_cloudcredential_applicationsecret_wo_2", "ephemeral_secret2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudcredential_applicationsecret_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudcredentialExist("citrixadc_cloudcredential.tf_cloudcredential", nil),
					// write-only value is not stored in state; version tracker is.
					resource.TestCheckResourceAttr("citrixadc_cloudcredential.tf_cloudcredential", "applicationsecret_wo_version", "1"),
					resource.TestCheckResourceAttr("citrixadc_cloudcredential.tf_cloudcredential", "applicationid", "22222222-2222-2222-2222-222222222222"),
				),
			},
			{
				Config: testAccCloudcredential_applicationsecret_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudcredentialExist("citrixadc_cloudcredential.tf_cloudcredential", nil),
					// version bumped confirms the update path was triggered.
					resource.TestCheckResourceAttr("citrixadc_cloudcredential.tf_cloudcredential", "applicationsecret_wo_version", "2"),
					resource.TestCheckResourceAttr("citrixadc_cloudcredential.tf_cloudcredential", "applicationid", "22222222-2222-2222-2222-222222222222"),
				),
			},
		},
	})
}
