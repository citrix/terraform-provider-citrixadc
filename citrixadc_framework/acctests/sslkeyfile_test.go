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

// sslkeyfile is an Import-as-create resource (POST ?action=Import), with no
// update (every attribute is RequiresReplace) and a keyless delete
// (DELETE ?args=name:<name>). The `src` attribute must point at a reachable
// key file (an http(s) URL or a local: path of a file already on the
// appliance). It is a write-only Import input that NITRO does not echo back,
// so it is never asserted. `password` is an optional secret (passphrase for an
// encrypted key file) expanded into the write-only triple
// password / password_wo / password_wo_version; secrets are never asserted.
//
// doSslkeyfilePreChecks (see helpers_test.go) uploads testdata/sample.key and
// testdata/sample_enc.key to /var/tmp as prerequisites. The basic/datasource
// tests use the UNENCRYPTED src = "local:sample.key"; the password_wo ephemeral
// test uses the ENCRYPTED src = "local:sample_enc.key" (the supplied passphrase
// decrypts it on Import). Secrets/src are never asserted.

const testAccSslkeyfile_basic_step1 = `
resource "citrixadc_sslkeyfile" "tf_sslkeyfile" {
  name = "tf_sslkeyfile"
  src  = "local:sample.key"
}

`

func TestAccSslkeyfile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslkeyfilePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslkeyfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslkeyfile_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslkeyfileExist("citrixadc_sslkeyfile.tf_sslkeyfile", nil),
					resource.TestCheckResourceAttr("citrixadc_sslkeyfile.tf_sslkeyfile", "name", "tf_sslkeyfile"),
				),
			},
		},
	})
}

func testAccCheckSslkeyfileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslkeyfile name is set")
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
		// sslkeyfile has NO get-by-name endpoint (GET /sslkeyfile/<name> => 400,
		// errorcode 1090). Get all records and filter by name.
		allResources, err := client.FindAllResources(service.Sslkeyfile.Type())
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
			return fmt.Errorf("sslkeyfile %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslkeyfileDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslkeyfile" {
			continue
		}

		// sslkeyfile has NO get-by-name endpoint. Get all records and filter by
		// name; a list error is treated as destroyed.
		allResources, err := client.FindAllResources(service.Sslkeyfile.Type())
		if err != nil {
			continue
		}

		for _, v := range allResources {
			if name, ok := v["name"].(string); ok && name == rs.Primary.ID {
				return fmt.Errorf("sslkeyfile %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

// Write-only ephemeral test for the password secret. The key file referenced by
// `src` must be encrypted with the supplied passphrase for the Import to
// succeed; bump password_wo_version between steps to trigger replacement. The
// secret value itself is never stored in state and is never asserted.
const testAccSslkeyfile_password_wo_step1 = `

	variable "sslkeyfile_password_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslkeyfile" "tf_sslkeyfile" {
		name                = "tf_sslkeyfile"
		src                 = "local:sample_enc.key"
		password_wo         = var.sslkeyfile_password_wo
		password_wo_version = 1
	}
`

func TestAccSslkeyfile_password_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_sslkeyfile_password_wo", "1234567")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslkeyfilePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslkeyfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslkeyfile_password_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslkeyfileExist("citrixadc_sslkeyfile.tf_sslkeyfile", nil),
					resource.TestCheckResourceAttr("citrixadc_sslkeyfile.tf_sslkeyfile", "password_wo_version", "1"),
					resource.TestCheckResourceAttr("citrixadc_sslkeyfile.tf_sslkeyfile", "name", "tf_sslkeyfile"),
				),
			},
		},
	})
}

const testAccSslkeyfileDataSource_basic = `

resource "citrixadc_sslkeyfile" "tf_sslkeyfile" {
  name = "tf_sslkeyfile"
  src  = "local:sample.key"
}

data "citrixadc_sslkeyfile" "tf_sslkeyfile" {
  name       = citrixadc_sslkeyfile.tf_sslkeyfile.name
  depends_on = [citrixadc_sslkeyfile.tf_sslkeyfile]
}
`

func TestAccSslkeyfileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslkeyfilePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslkeyfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslkeyfileDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslkeyfile.tf_sslkeyfile", "name", "tf_sslkeyfile"),
				),
			},
		},
	})
}
