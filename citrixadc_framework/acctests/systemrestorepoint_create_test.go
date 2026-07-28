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

// systemrestorepoint is an action-create + plain-delete + get resource.
// filename is Required and RequiresReplace, so there is no update step.
//
// PREREQUISITE:
//   - The appliance rejects "create system restorepoint" with errorcode 1067
//     ("Feature(s) not enabled") unless the auto-restore feature is enabled first
//     (CLI: "enable system autorestorefeature"). The test config therefore
//     declares a citrixadc_systemautorestorefeature_enable resource and makes the
//     restore point depend_on it so the feature is enabled before create.
//
// CAVEATS:
//   - The appliance enforces a MAXIMUM of 3 restore points. If create fails with
//     a cap error, remove stale tf-created restore points (e.g. via NITRO:
//     curl -X DELETE -u nsroot:<pw> http://<ns>/nitro/v1/config/systemrestorepoint/tf_restorepoint_test).
//   - Creating a restore point snapshots the config + a tech-support bundle; it is
//     resource-intensive and may take time, hence the generous timeout in the run command.
//   - Restore point creation stages a large tar (config + tech-support bundle) using
//     scratch space on the appliance root/tmp filesystem. On a testbed whose root
//     filesystem (/dev/md0) is (near-)full, create fails with errorcode 3428
//     ("Unable to create backup tar file"). This is an appliance/platform limitation,
//     not a provider or test defect.

const testAccSystemrestorepointCreate_basic = `

	resource "citrixadc_systemautorestorefeature_enable" "tf_systemautorestorefeature" {
	}

	resource "citrixadc_systemrestorepoint_create" "tf_systemrestorepoint" {
		filename   = "tf_restorepoint_test"
		depends_on = [citrixadc_systemautorestorefeature_enable.tf_systemautorestorefeature]
	}
`

func TestAccSystemrestorepointCreate_basic(t *testing.T) {
	t.Skip("Requires enabling System Auto Restore feature.")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemrestorepointCreateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemrestorepointCreate_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemrestorepointCreateExist("citrixadc_systemrestorepoint_create.tf_systemrestorepoint", nil),
					resource.TestCheckResourceAttr("citrixadc_systemrestorepoint_create.tf_systemrestorepoint", "filename", "tf_restorepoint_test"),
				),
			},
		},
	})
}

func TestAccSystemrestorepointCreate_import(t *testing.T) {
	t.Skip("TODO: Requires review")
	const resAddr = "citrixadc_systemrestorepoint_create.tf_systemrestorepoint"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemrestorepointCreateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemrestorepointCreate_basic,
			},
			{
				Config:                  testAccSystemrestorepointCreate_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckSystemrestorepointCreateExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No systemrestorepoint filename is set")
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
		data, err := client.FindResource(service.Systemrestorepoint.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("systemrestorepoint %s not found", n)
		}

		return nil
	}
}

func testAccCheckSystemrestorepointCreateDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_systemrestorepoint_create" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No filename is set")
		}

		_, err := client.FindResource(service.Systemrestorepoint.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("systemrestorepoint %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
