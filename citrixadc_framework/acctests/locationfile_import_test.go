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
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

// testAccLocationfileImport_basic exercises a genuine locationfile import.
//
// The import action (POST locationfile?action=import) needs a real source file
// reachable via `src`. We provide one without any external download by first
// uploading a small netscaler-format location DB into /var/tmp using the
// (SDK v2, muxed) citrixadc_systemfile resource, then importing it via the
// local:// scheme (which resolves to /var/tmp on the appliance). depends_on
// guarantees the source file exists before the import runs.
const testAccLocationfileImport_basic = `

resource "citrixadc_systemfile" "tf_locationfile_src" {
	filename     = "tf_locationfile_src"
	filelocation = "/var/tmp"
	filecontent  = "\"1.0.0.0-1.0.0.255\",\"North America.United States.California.San Jose\"\n\"2.0.0.0-2.0.0.255\",\"Europe.United Kingdom.England.London\"\n"
}

resource "citrixadc_locationfile_import" "tf_locationfile_import" {
	locationfile = "tf_locationfile_import"
	src          = "local://tf_locationfile_src"
	format       = "netscaler"
	depends_on   = [citrixadc_systemfile.tf_locationfile_src]
}
`

func TestAccLocationfileImport_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccLocationfileImport_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLocationfileImportExist("citrixadc_locationfile_import.tf_locationfile_import", nil),
					resource.TestCheckResourceAttr("citrixadc_locationfile_import.tf_locationfile_import", "locationfile", "tf_locationfile_import"),
					resource.TestCheckResourceAttr("citrixadc_locationfile_import.tf_locationfile_import", "src", "local://tf_locationfile_src"),
					resource.TestCheckResourceAttr("citrixadc_locationfile_import.tf_locationfile_import", "format", "netscaler"),
				),
			},
		},
	})
}

func testAccCheckLocationfileImportExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No locationfile name is set")
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
		data, err := client.FindResource(service.Locationfile.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("locationfile %s not found", n)
		}

		return nil
	}
}
