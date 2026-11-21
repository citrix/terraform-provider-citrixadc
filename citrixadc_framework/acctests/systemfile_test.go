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
	"encoding/base64"
	"fmt"
	"net/url"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccSystemfile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemfile_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemfileExist("citrixadc_systemfile.tf_file", nil, []string{"/var/tmp", "hello.txt"}),
				),
			},
			{
				Config: testAccSystemfile_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemfileExist("citrixadc_systemfile.tf_file", nil, []string{"/tmp", "helloandbye.txt"}),
				),
			},
		},
	})
}

func TestAccSystemfile_base64(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemfile_base64_plaintext,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemfileExist("citrixadc_systemfile.tf_file_plain", nil, []string{"/var/tmp", "plain.txt"}),
				),
			},
			{
				Config: testAccSystemfile_base64_encoded,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemfileExist("citrixadc_systemfile.tf_file_encoded", nil, []string{"/var/tmp", "encoded.txt"}),
				),
			},
		},
	})
}

func TestAccSystemfile_import(t *testing.T) {
	const resAddr = "citrixadc_systemfile.tf_file"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemfileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccSystemfile_basic_step1},
			{
				Config:            testAccSystemfile_basic_step1,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// filecontent_wo is write-only (never in state); filecontent_wo_version is a
				// Computed version tracker not populated by the import Read (its Default is
				// applied at plan time, so it reads back null right after import).
				ImportStateVerifyIgnore: []string{"filecontent_wo", "filecontent_wo_version"},
			},
		},
	})
}

// testAccCheckSystemfileContent verifies the file content on the appliance matches
// the expected plaintext. Used by the write-only (filecontent_wo) test where the
// content is intentionally absent from Terraform state, so the only way to assert
// the content flowed through is to read it back from the box and decode it.
func testAccCheckSystemfileContent(location, name, expectedPlain string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		argsMap := make(map[string]string)
		argsMap["filelocation"] = url.QueryEscape(location)
		argsMap["filename"] = url.QueryEscape(name)
		findParams := service.FindParams{ResourceType: "systemfile", ArgsMap: argsMap}
		arr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}
		if len(arr) == 0 {
			return fmt.Errorf("systemfile %s/%s not found", location, name)
		}
		raw, _ := arr[0]["filecontent"].(string)
		decoded, err := base64.StdEncoding.DecodeString(raw)
		if err != nil {
			return fmt.Errorf("systemfile %s/%s content is not valid base64: %v", location, name, err)
		}
		if string(decoded) != expectedPlain {
			return fmt.Errorf("systemfile %s/%s content mismatch: got %q, want %q", location, name, string(decoded), expectedPlain)
		}
		return nil
	}
}

func testAccCheckSystemfileExist(n string, id *string, pathData []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lb vserver name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		argsMap := make(map[string]string)
		argsMap["filelocation"] = url.QueryEscape(pathData[0])
		argsMap["filename"] = url.QueryEscape(pathData[1])
		findParams := service.FindParams{
			ResourceType: "systemfile",
			ArgsMap:      argsMap,
		}
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		dataArray, err := client.FindResourceArrayWithParams(findParams)

		if err != nil {
			return err
		}

		if len(dataArray) == 0 {
			return fmt.Errorf("systemfile %s not found", n)
		}

		return nil
	}
}

func testAccCheckSystemfileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_systemfile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Systemfile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccSystemfile_basic_step1 = `

resource "citrixadc_systemfile" "tf_file" {
    filename = "hello.txt"
    filelocation = "/var/tmp"
    filecontent = "hello"
}
`

const testAccSystemfile_basic_step2 = `

resource "citrixadc_systemfile" "tf_file" {
    filename = "helloandbye.txt"
    filelocation = "/tmp"
    filecontent = "hello and goodbye"
}
`

const testAccSystemfile_base64_plaintext = `

resource "citrixadc_systemfile" "tf_file_plain" {
    filename = "plain.txt"
    filelocation = "/var/tmp"
    filecontent = "This is plain text content"
    fileencoding = "BASE64"
}
`

const testAccSystemfile_base64_encoded = `

resource "citrixadc_systemfile" "tf_file_encoded" {
    filename = "encoded.txt"
    filelocation = "/var/tmp"
    filecontent = "SGVsbG8gV29ybGQhIFRoaXMgaXMgYSBiYXNlNjQgZW5jb2RlZCBleGFtcGxlIGZpbGU="
    fileencoding = "BASE64"
}
`

const testAccSystemfile_filecontent_wo_step1 = `

variable "systemfile_filecontent_wo" {
  type      = string
  sensitive = true
}

resource "citrixadc_systemfile" "tf_file_wo" {
    filename               = "wo_secret.txt"
    filelocation           = "/var/tmp"
    filecontent_wo         = var.systemfile_filecontent_wo
    filecontent_wo_version = 1
}
`

const testAccSystemfile_filecontent_wo_step2 = `

variable "systemfile_filecontent_wo_2" {
  type      = string
  sensitive = true
}

resource "citrixadc_systemfile" "tf_file_wo" {
    filename               = "wo_secret.txt"
    filelocation           = "/var/tmp"
    filecontent_wo         = var.systemfile_filecontent_wo_2
    filecontent_wo_version = 2
}
`

// TestAccSystemfile_filecontent_wo_ephemeral exercises the write-only content path:
// the file content is supplied via filecontent_wo (never stored in state) and a
// filecontent_wo_version bump rotates it (destroy+recreate, since systemfile has no
// NITRO 'set'). It asserts the content actually reached the appliance while neither
// filecontent_wo nor the plaintext filecontent leaks into Terraform state.
func TestAccSystemfile_filecontent_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_systemfile_filecontent_wo", "write-only secret content v1")
	t.Setenv("TF_VAR_systemfile_filecontent_wo_2", "write-only secret content v2 rotated")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemfile_filecontent_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemfileExist("citrixadc_systemfile.tf_file_wo", nil, []string{"/var/tmp", "wo_secret.txt"}),
					// The content reached the box even though it is absent from state.
					testAccCheckSystemfileContent("/var/tmp", "wo_secret.txt", "write-only secret content v1"),
					resource.TestCheckResourceAttr("citrixadc_systemfile.tf_file_wo", "filecontent_wo_version", "1"),
					// The secret and its write-only sibling are never persisted to state.
					resource.TestCheckNoResourceAttr("citrixadc_systemfile.tf_file_wo", "filecontent"),
					resource.TestCheckNoResourceAttr("citrixadc_systemfile.tf_file_wo", "filecontent_wo"),
				),
			},
			{
				Config: testAccSystemfile_filecontent_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemfileExist("citrixadc_systemfile.tf_file_wo", nil, []string{"/var/tmp", "wo_secret.txt"}),
					testAccCheckSystemfileContent("/var/tmp", "wo_secret.txt", "write-only secret content v2 rotated"),
					resource.TestCheckResourceAttr("citrixadc_systemfile.tf_file_wo", "filecontent_wo_version", "2"),
					resource.TestCheckNoResourceAttr("citrixadc_systemfile.tf_file_wo", "filecontent"),
				),
			},
		},
	})
}

const testAccSystemfileDataSource_basic = `
resource "citrixadc_systemfile" "tf_file_ds" {
    filename = "datasource_test.txt"
    filelocation = "/var/tmp"
    filecontent = "datasource test content"
    fileencoding = "BASE64"
}

data "citrixadc_systemfile" "tf_file_datasource" {
    filelocation = citrixadc_systemfile.tf_file_ds.filelocation
    filename = citrixadc_systemfile.tf_file_ds.filename
}
`

func TestAccSystemfile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckSystemfileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccSystemfile_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSystemfileExist("citrixadc_systemfile.tf_file", nil, []string{"/var/tmp", "hello.txt"})),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccSystemfile_basic_step1,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSystemfileExist("citrixadc_systemfile.tf_file", nil, []string{"/var/tmp", "hello.txt"})),
			},
		},
	})
}

func TestAccSystemfileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemfileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_systemfile.tf_file_datasource", "filename", "datasource_test.txt"),
					resource.TestCheckResourceAttr("data.citrixadc_systemfile.tf_file_datasource", "filelocation", "/var/tmp"),
					resource.TestCheckResourceAttr("data.citrixadc_systemfile.tf_file_datasource", "fileencoding", "BASE64"),
					resource.TestCheckResourceAttrSet("data.citrixadc_systemfile.tf_file_datasource", "filecontent"),
					resource.TestCheckResourceAttrSet("data.citrixadc_systemfile.tf_file_datasource", "id"),
					// Read-only file metadata that is always populated for an existing file.
					resource.TestCheckResourceAttrSet("data.citrixadc_systemfile.tf_file_datasource", "filesize"),
				),
			},
		},
	})
}

// TestAccSystemfile_gh1141_partitionFilelocation is the regression test for GH #1141.
// In an admin partition the ADC reports systemfile filelocation in partition-relative
// normalized form (/var/tmp/) regardless of the path used to create or query it, so the
// provider must NOT overwrite the user's configured filelocation with that GET echo.
//
// The file is created in partition context (provider do_login=true + partition=...) using
// the ABSOLUTE path /var/partitions/<p>/tmp, exactly as the reporter did. Pre-fix the
// apply fails with "Provider produced inconsistent result after apply (.filelocation: was
// .../tmp, but now /var/tmp/)". Post-fix the configured filelocation is preserved and the
// post-apply plan is empty (idempotent).
func TestAccSystemfile_gh1141_partitionFilelocation(t *testing.T) {
	if isCpxRun {
		t.Skip("admin partitions are not supported on CPX")
	}
	const partName = "tf1141part"
	const resAddr = "citrixadc_systemfile.tf1141"
	absLoc := "/var/partitions/" + partName + "/tmp"

	// Create the admin partition out-of-band (default context) BEFORE the
	// partition-scoped provider configures and switches into it. Skip (not fail) if the
	// appliance does not permit partition creation (e.g. a protected cluster/HA box), so
	// the test is self-gating across the testbed.
	client, err := testAccGetFrameworkClient()
	if err != nil {
		t.Fatalf("gh1141: get client: %v", err)
	}
	if _, err := client.AddResource(service.Nspartition.Type(), partName,
		map[string]interface{}{"partitionname": partName}); err != nil {
		t.Skipf("gh1141: cannot create admin partition %q on this appliance: %v", partName, err)
	}
	t.Cleanup(func() {
		// Best-effort teardown: removing the partition also removes any file left in it.
		_ = client.DeleteResource(service.Nspartition.Type(), partName)
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemfile_gh1141(partName, absLoc),
				Check: resource.ComposeTestCheckFunc(
					// The configured absolute path must survive the read-back, NOT be
					// replaced by the partition-normalized "/var/tmp/" echo.
					resource.TestCheckResourceAttr(resAddr, "filelocation", absLoc),
					resource.TestCheckResourceAttr(resAddr, "id", absLoc+"/tf1141.txt"),
				),
				// Idempotency: no diff after apply+refresh (guards the absolute-path and
				// trailing-slash round-trip once the clobber is removed).
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PostApplyPostRefresh: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
				},
			},
		},
	})
}

func testAccSystemfile_gh1141(partName, absLoc string) string {
	return fmt.Sprintf(`
provider "citrixadc" {
  do_login  = true
  partition = "%s"
}

resource "citrixadc_systemfile" "tf1141" {
  filename     = "tf1141.txt"
  filelocation = "%s"
  filecontent  = "gh1141 partition filelocation"
}
`, partName, absLoc)
}
