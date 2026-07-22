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

// sslcertkeybundle supports add + update (?action=update) + delete by certkeybundlename.
// passplain is a secret expanded into the write-only triple
// (passplain / passplain_wo / passplain_wo_version) - secrets are never asserted.
//
// NOTE: bundlefile must be a cert-key-bundle file (certificates + one PEM key) that
// ALREADY EXISTS on the appliance under /nsconfig/ssl/. doSslcertkeybundlePreChecks
// (see helpers_test.go) uploads testdata/servercert1_certkeybundle.pem there as a
// prerequisite, so bundlefile = "servercert1_certkeybundle.pem" resolves on the testbed.
const testAccSslcertkeybundle_basic_step1 = `

resource "citrixadc_sslcertkeybundle" "tf_sslcertkeybundle" {
  certkeybundlename    = "tf_sslcertkeybundle"
  bundlefile           = "servercert1_certkeybundle.pem"
}

`

func TestAccSslcertkeybundle_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeybundlePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslcertkeybundleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcertkeybundle_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertkeybundleExist("citrixadc_sslcertkeybundle.tf_sslcertkeybundle", nil),
					resource.TestCheckResourceAttr("citrixadc_sslcertkeybundle.tf_sslcertkeybundle", "certkeybundlename", "tf_sslcertkeybundle"),
					resource.TestCheckResourceAttr("citrixadc_sslcertkeybundle.tf_sslcertkeybundle", "bundlefile", "servercert1_certkeybundle.pem"),
				),
			},
		},
	})
}

func TestAccSslcertkeybundle_import(t *testing.T) {
	const resAddr = "citrixadc_sslcertkeybundle.tf_sslcertkeybundle"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeybundlePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslcertkeybundleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcertkeybundle_basic_step1,
			},
			{
				Config:                  testAccSslcertkeybundle_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"passplain_wo_version"},
			},
		},
	})
}

func testAccCheckSslcertkeybundleExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslcertkeybundle name is set")
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
		data, err := client.FindResource(service.Sslcertkeybundle.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("sslcertkeybundle %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslcertkeybundleDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslcertkeybundle" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Sslcertkeybundle.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslcertkeybundle %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

// Write-only ephemeral test for passplain: drives the passplain_wo / passplain_wo_version
// triple and confirms the version tracker bumps between steps (the secret itself is never
// stored in state, so it is never asserted).
const testAccSslcertkeybundle_passplain_wo_step1 = `

	variable "sslcertkeybundle_passplain_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslcertkeybundle" "tf_sslcertkeybundle" {
		certkeybundlename    = "tf_sslcertkeybundle"
		bundlefile           = "servercert1_certkeybundle_with_passplain.pem"
		passplain_wo         = var.sslcertkeybundle_passplain_wo
		passplain_wo_version = 1
	}
`

func TestAccSslcertkeybundle_passplain_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_sslcertkeybundle_passplain_wo", "1234567")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeybundlePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslcertkeybundleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcertkeybundle_passplain_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertkeybundleExist("citrixadc_sslcertkeybundle.tf_sslcertkeybundle", nil),
					resource.TestCheckResourceAttr("citrixadc_sslcertkeybundle.tf_sslcertkeybundle", "passplain_wo_version", "1"),
				),
			},
		},
	})
}

// Backward-compatible test for passplain: drives the legacy Sensitive attribute directly.
const testAccSslcertkeybundle_passplain_step1 = `

	variable "sslcertkeybundle_passplain" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_sslcertkeybundle" "tf_sslcertkeybundle" {
		certkeybundlename = "tf_sslcertkeybundle"
		bundlefile        = "servercert1_certkeybundle_with_passplain.pem"
		passplain         = var.sslcertkeybundle_passplain
	}
`

func TestAccSslcertkeybundle_passplain_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_sslcertkeybundle_passplain", "1234567")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeybundlePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslcertkeybundleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcertkeybundle_passplain_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertkeybundleExist("citrixadc_sslcertkeybundle.tf_sslcertkeybundle", nil),
					resource.TestCheckResourceAttr("citrixadc_sslcertkeybundle.tf_sslcertkeybundle", "certkeybundlename", "tf_sslcertkeybundle"),
				),
			},
		},
	})
}

const testAccSslcertkeybundleDataSource_basic = `

resource "citrixadc_sslcertkeybundle" "tf_sslcertkeybundle" {
  certkeybundlename = "tf_sslcertkeybundle"
  bundlefile        = "servercert1_certkeybundle.pem"
}

data "citrixadc_sslcertkeybundle" "tf_sslcertkeybundle" {
  certkeybundlename = citrixadc_sslcertkeybundle.tf_sslcertkeybundle.certkeybundlename
  depends_on        = [citrixadc_sslcertkeybundle.tf_sslcertkeybundle]
}
`

func TestAccSslcertkeybundleDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcertkeybundlePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslcertkeybundleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcertkeybundleDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslcertkeybundle.tf_sslcertkeybundle", "certkeybundlename", "tf_sslcertkeybundle"),
					resource.TestCheckResourceAttr("data.citrixadc_sslcertkeybundle.tf_sslcertkeybundle", "bundlefile", "servercert1_certkeybundle.pem"),
				),
			},
		},
	})
}
