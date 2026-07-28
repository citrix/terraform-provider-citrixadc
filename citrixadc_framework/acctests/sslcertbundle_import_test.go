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

// sslcertbundle is Import-as-create (ActOnResource "Import"), no update (all attrs
// RequiresReplace), delete by name. GET is list-and-filter by name, and NITRO does NOT
// echo back the src - so the basic test asserts only "name", never "src".
//
// src is the Import source. doSslcertbundlePreChecks (see helpers_test.go) uploads
// testdata/servercert1_bundle.pem to /var/tmp as a prerequisite, so the
// `local:servercert1_bundle.pem` source resolves on the appliance.
const testAccSslcertbundleImport_basic_step1 = `
resource "citrixadc_sslcertbundle_import" "tf_sslcertbundle" {
  name = "tf_sslcertbundle"
  src  = "local:servercert1_bundle.pem"
}

`

func TestAccSslcertbundleImport_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertbundlePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslcertbundleImportDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcertbundleImport_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertbundleImportExist("citrixadc_sslcertbundle_import.tf_sslcertbundle", nil),
					resource.TestCheckResourceAttr("citrixadc_sslcertbundle_import.tf_sslcertbundle", "name", "tf_sslcertbundle"),
					// src is an Import-only input and is not echoed back by NITRO GET - do not assert it.
				),
			},
		},
	})
}

func TestAccSslcertbundleImport_import(t *testing.T) {
	const resAddr = "citrixadc_sslcertbundle_import.tf_sslcertbundle"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertbundlePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslcertbundleImportDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcertbundleImport_basic_step1,
			},
			{
				Config:            testAccSslcertbundleImport_basic_step1,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// src is an Import-only input and is not echoed back by NITRO GET,
				// so it cannot round-trip on import.
				ImportStateVerifyIgnore: []string{"src"},
			},
		},
	})
}

func testAccCheckSslcertbundleImportExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslcertbundle name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		// GET has no get-by-name endpoint; list all and filter by name.
		allResources, err := client.FindAllResources(service.Sslcertbundle.Type())
		if err != nil {
			return err
		}

		found := false
		for _, v := range allResources {
			if name, ok := v["name"].(string); ok && name == rs.Primary.ID {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("sslcertbundle %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslcertbundleImportDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslcertbundle_import" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		allResources, err := client.FindAllResources(service.Sslcertbundle.Type())
		if err != nil {
			// An error fetching the (now-empty) list is acceptable as destroyed.
			continue
		}
		for _, v := range allResources {
			if name, ok := v["name"].(string); ok && name == rs.Primary.ID {
				return fmt.Errorf("sslcertbundle %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}
