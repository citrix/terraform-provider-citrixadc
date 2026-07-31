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
	"net/url"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// testAccLocationfileImport6_basic exercises a genuine IPv6 locationfile import.
//
// The import action (POST locationfile6?action=import) needs a real source file
// reachable via `src`. We provide one without any external download by first
// uploading a small netscaler6-format IPv6 location DB into /var/tmp using the
// (SDK v2, muxed) citrixadc_systemfile resource, then importing it via the
// local:// scheme (which resolves to /var/tmp on the appliance). depends_on
// guarantees the source file exists before the import runs.
//
// NOTE: the import action only accepts http/https/ftp AND local:// sources
// (file:// is rejected with NITRO errorcode 3234); local://<name> resolves to
// /var/tmp, which is why the uploaded source can be referenced by bare filename.
const testAccLocationfileImport6_basic = `

resource "citrixadc_systemfile" "tf_locationfile6_src" {
	filename     = "tf_locationfile6_src"
	filelocation = "/var/tmp"
	filecontent  = "\"2001:db8::-2001:db8::ffff\",\"North America.United States.California.San Jose\"\n\"2001:db9::-2001:db9::ffff\",\"Europe.United Kingdom.England.London\"\n"
}

resource "citrixadc_locationfile6_import" "tf_locationfile6_import" {
	locationfile = "tf_locationfile6_import"
	src          = "local://tf_locationfile6_src"
	format       = "netscaler6"
	depends_on   = [citrixadc_systemfile.tf_locationfile6_src]
}
`

func TestAccLocationfile6Import_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckLocationfile6ImportDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLocationfileImport6_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLocationfile6ImportExist("citrixadc_locationfile6_import.tf_locationfile6_import", nil),
					resource.TestCheckResourceAttr("citrixadc_locationfile6_import.tf_locationfile6_import", "locationfile", "tf_locationfile6_import"),
					resource.TestCheckResourceAttr("citrixadc_locationfile6_import.tf_locationfile6_import", "src", "local://tf_locationfile6_src"),
					resource.TestCheckResourceAttr("citrixadc_locationfile6_import.tf_locationfile6_import", "format", "netscaler6"),
				),
			},
		},
	})
}

func testAccCheckLocationfile6ImportExist(n string, id *string) resource.TestCheckFunc {
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

		// The import action copies the source file into /var/netscaler/locdb. For
		// IPv6, GET locationfile6 returns no object until a file is actually loaded
		// (unlike GET locationfile for IPv4, which always returns a singleton
		// status object), so we verify the import genuinely landed by confirming
		// the imported file is present in /var/netscaler/locdb.
		// The nitro-go client does not URL-encode ArgsMap values, and NITRO
		// rejects the raw slashes in a filelocation path, so pre-encode it.
		locationfileName := rs.Primary.Attributes["locationfile"]
		findParams := service.FindParams{
			ResourceType: service.Systemfile.Type(),
			ArgsMap:      map[string]string{"filelocation": url.QueryEscape("/var/netscaler/locdb")},
		}
		files, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return fmt.Errorf("Failed to list /var/netscaler/locdb: %v", err)
		}
		for _, f := range files {
			if fn, ok := f["filename"].(string); ok && fn == locationfileName {
				return nil
			}
		}
		return fmt.Errorf("imported locationfile6 %q not found in /var/netscaler/locdb", locationfileName)
	}
}

// testAccCheckLocationfile6ImportDestroy is a best-effort cleanup hook. The
// import action has no inverse (the resource's Delete is a no-op), so terraform
// destroy legitimately leaves the imported file in /var/netscaler/locdb and its
// entry in the locdb import index (mapping-spdbfile). We remove both here so the
// appliance is left in its pre-test state and the test stays rerunnable: a repeat
// import of the same Locationfile name otherwise fails with NITRO errorcode 3198
// ("Object already exists") because the name is still registered in the index.
// Cleanup failures are not treated as test failures (there is nothing to verify
// as "destroyed" for a no-op-delete action resource).
func testAccCheckLocationfile6ImportDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return nil
	}
	// The nitro-go client does not URL-encode args, and NITRO rejects raw slashes
	// in a filelocation path, so pre-encode it.
	locdbArgs := []string{"filelocation:" + url.QueryEscape("/var/netscaler/locdb")}
	indexReset := false
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_locationfile6_import" {
			continue
		}
		if name := rs.Primary.Attributes["locationfile"]; name != "" {
			// Remove the imported copy from /var/netscaler/locdb.
			_ = client.DeleteResourceWithArgs(service.Systemfile.Type(), name, locdbArgs)
		}
		if !indexReset {
			// Reset the import index so the imported name(s) are freed. The index
			// regenerates automatically on the next import.
			_ = client.DeleteResourceWithArgs(service.Systemfile.Type(), "mapping-spdbfile", locdbArgs)
			indexReset = true
		}
	}
	return nil
}
