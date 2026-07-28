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
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccMetricsprofile_basic_step1 = `
resource "citrixadc_metricsprofile" "tf_metricsprofile" {
  name                   = "tf_metricsprofile"
  outputmode             = "avro"
  metrics                = "ENABLED"
  servemode              = "Push"
  metricsexportfrequency = 30
}

`

const testAccMetricsprofile_basic_step2 = `
resource "citrixadc_metricsprofile" "tf_metricsprofile" {
  name                   = "tf_metricsprofile"
  outputmode             = "avro"
  metrics                = "DISABLED"
  servemode              = "Push"
  metricsexportfrequency = 120
}

`

func TestAccMetricsprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMetricsprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMetricsprofile_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMetricsprofileExist("citrixadc_metricsprofile.tf_metricsprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile.tf_metricsprofile", "name", "tf_metricsprofile"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile.tf_metricsprofile", "outputmode", "avro"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile.tf_metricsprofile", "metrics", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile.tf_metricsprofile", "servemode", "Push"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile.tf_metricsprofile", "metricsexportfrequency", "30"),
				),
			},
			{
				Config: testAccMetricsprofile_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMetricsprofileExist("citrixadc_metricsprofile.tf_metricsprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile.tf_metricsprofile", "name", "tf_metricsprofile"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile.tf_metricsprofile", "outputmode", "avro"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile.tf_metricsprofile", "metrics", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile.tf_metricsprofile", "servemode", "Push"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile.tf_metricsprofile", "metricsexportfrequency", "120"),
				),
			},
		},
	})
}

func TestAccMetricsprofile_import(t *testing.T) {
	const resAddr = "citrixadc_metricsprofile.tf_metricsprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMetricsprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMetricsprofile_basic_step1,
			},
			{
				Config:                  testAccMetricsprofile_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"metricsauthtoken_wo_version"},
			},
		},
	})
}

func testAccCheckMetricsprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No metricsprofile name is set")
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
		data, err := client.FindResource(service.Metricsprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("metricsprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckMetricsprofileDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_metricsprofile" {
			continue
		}

		_, err := client.FindResource(service.Metricsprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("metricsprofile %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

const testAccMetricsprofileDataSource_basic = `

resource "citrixadc_metricsprofile" "tf_metricsprofile" {
  name                   = "tf_metricsprofile"
  outputmode             = "avro"
  metrics                = "ENABLED"
  servemode              = "Push"
  metricsexportfrequency = 30
}

data "citrixadc_metricsprofile" "tf_metricsprofile" {
  name       = citrixadc_metricsprofile.tf_metricsprofile.name
  depends_on = [citrixadc_metricsprofile.tf_metricsprofile]
}
`

func TestAccMetricsprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccMetricsprofileDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_metricsprofile.tf_metricsprofile", "name", "tf_metricsprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_metricsprofile.tf_metricsprofile", "outputmode", "avro"),
					resource.TestCheckResourceAttr("data.citrixadc_metricsprofile.tf_metricsprofile", "metrics", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_metricsprofile.tf_metricsprofile", "servemode", "Push"),
					resource.TestCheckResourceAttr("data.citrixadc_metricsprofile.tf_metricsprofile", "metricsexportfrequency", "30"),
				),
			},
		},
	})
}

const testAccMetricsprofile_metricsauthtoken_step1 = `

	variable "metricsprofile_metricsauthtoken" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_metricsprofile" "tf_metricsprofile" {
		name               = "tf_metricsprofile_secret"
		metricsauthtoken   = var.metricsprofile_metricsauthtoken
		outputmode             = "avro"
		metrics                = "ENABLED"
		servemode              = "Push"
		metricsexportfrequency = 30
	}
`

const testAccMetricsprofile_metricsauthtoken_step2 = `

	 variable "metricsprofile_metricsauthtoken_2" {
	  type      = string
	  sensitive = true
	}

	 resource "citrixadc_metricsprofile" "tf_metricsprofile" {
		name               = "tf_metricsprofile_secret"
		metricsauthtoken   = var.metricsprofile_metricsauthtoken_2
		outputmode             = "avro"
		metrics                = "DISABLED"
		servemode              = "Push"
		metricsexportfrequency = 30
	}
`

func TestAccMetricsprofile_metricsauthtoken_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_metricsprofile_metricsauthtoken", "SplunkOldtoken123")
	t.Setenv("TF_VAR_metricsprofile_metricsauthtoken_2", "SplunkNewtoken456")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMetricsprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMetricsprofile_metricsauthtoken_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMetricsprofileExist("citrixadc_metricsprofile.tf_metricsprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile.tf_metricsprofile", "metrics", "ENABLED"),
				),
			},
			{
				Config: testAccMetricsprofile_metricsauthtoken_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMetricsprofileExist("citrixadc_metricsprofile.tf_metricsprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile.tf_metricsprofile", "metrics", "DISABLED"),
				),
			},
		},
	})
}

const testAccMetricsprofile_metricsauthtoken_wo_step1 = `

	variable "metricsprofile_metricsauthtoken_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_metricsprofile" "tf_metricsprofile" {
		name                        = "tf_metricsprofile_secret_wo"
		metricsauthtoken_wo         = var.metricsprofile_metricsauthtoken_wo
		metricsauthtoken_wo_version = 1
		outputmode             = "avro"
		metrics                = "ENABLED"
		servemode              = "Push"
		metricsexportfrequency = 30
	}
`

const testAccMetricsprofile_metricsauthtoken_wo_step2 = `

	 variable "metricsprofile_metricsauthtoken_wo_2" {
	  type      = string
	  sensitive = true
	}

	 resource "citrixadc_metricsprofile" "tf_metricsprofile" {
		name                        = "tf_metricsprofile_secret_wo"
		metricsauthtoken_wo         = var.metricsprofile_metricsauthtoken_wo_2
		metricsauthtoken_wo_version = 2
		outputmode             = "avro"
		metrics                = "ENABLED"
		servemode              = "Push"
		metricsexportfrequency = 30
	}
`

const testAccMetricsprofile_unset_step1 = `
resource "citrixadc_metricsprofile" "tf_unset" {
  name                   = "tf_test_metricsprofile_unset"
  metrics                = "ENABLED"
  metricsexportfrequency = 120
  outputmode             = "prometheus"
  servemode              = "Pull"
}
`

const testAccMetricsprofile_unset_step2 = `
resource "citrixadc_metricsprofile" "tf_unset" {
  name = "tf_test_metricsprofile_unset"
  # eligible attributes removed from config -> provider must unset them
}
`

func TestAccMetricsprofile_unset(t *testing.T) {
	// The resource's other tests (basic/import/datasource) run on the default
	// standalone testbed with no skip guard, so this test has none either.
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMetricsprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values apply and persist.
				Config: testAccMetricsprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMetricsprofileExist("citrixadc_metricsprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile.tf_unset", "metrics", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile.tf_unset", "metricsexportfrequency", "120"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile.tf_unset", "outputmode", "prometheus"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile.tf_unset", "servemode", "Pull"),
				),
			},
			{
				// Removing them must unset -> state reverts to NITRO defaults,
				// and the implicit post-apply plan must be empty.
				Config: testAccMetricsprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMetricsprofileExist("citrixadc_metricsprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile.tf_unset", "metrics", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile.tf_unset", "metricsexportfrequency", "30"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile.tf_unset", "outputmode", "avro"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile.tf_unset", "servemode", "Push"),
					// Independent appliance-level confirmation the unset took effect:
					testAccCheckMetricsprofileADCValue("tf_test_metricsprofile_unset", "metrics", "DISABLED"),
					testAccCheckMetricsprofileADCValue("tf_test_metricsprofile_unset", "metricsexportfrequency", "30"),
					testAccCheckMetricsprofileADCValue("tf_test_metricsprofile_unset", "outputmode", "avro"),
					testAccCheckMetricsprofileADCValue("tf_test_metricsprofile_unset", "servemode", "Push"),
				),
			},
		},
	})
}

// testAccCheckMetricsprofileADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckMetricsprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Metricsprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("metricsprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("metricsprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccMetricsprofile_metricsauthtoken_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_metricsprofile_metricsauthtoken_wo", "SplunkEphemeral_tok1")
	t.Setenv("TF_VAR_metricsprofile_metricsauthtoken_wo_2", "SplunkEphemeral_tok2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckMetricsprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMetricsprofile_metricsauthtoken_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMetricsprofileExist("citrixadc_metricsprofile.tf_metricsprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile.tf_metricsprofile", "metricsauthtoken_wo_version", "1"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile.tf_metricsprofile", "outputmode", "avro"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile.tf_metricsprofile", "metrics", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile.tf_metricsprofile", "servemode", "Push"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile.tf_metricsprofile", "metricsexportfrequency", "30"),
				),
			},
			{
				Config: testAccMetricsprofile_metricsauthtoken_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMetricsprofileExist("citrixadc_metricsprofile.tf_metricsprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile.tf_metricsprofile", "metricsauthtoken_wo_version", "2"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile.tf_metricsprofile", "outputmode", "avro"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile.tf_metricsprofile", "metrics", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile.tf_metricsprofile", "servemode", "Push"),
					resource.TestCheckResourceAttr("citrixadc_metricsprofile.tf_metricsprofile", "metricsexportfrequency", "30"),
				),
			},
		},
	})
}
