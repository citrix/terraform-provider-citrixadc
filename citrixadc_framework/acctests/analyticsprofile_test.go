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

const testAccAnalyticsprofile_basic = `

	resource "citrixadc_analyticsprofile" "tf_analyticsprofile" {
		name             = "my_analyticsprofile"
		type             = "webinsight"
		httppagetracking = "DISABLED"
		httpurl          = "DISABLED"
	}
`
const testAccAnalyticsprofile_update = `

	resource "citrixadc_analyticsprofile" "tf_analyticsprofile" {
		name             = "my_analyticsprofile"
		type             = "webinsight"
		httppagetracking = "ENABLED"
		httpurl          = "ENABLED"
	}
`

func TestAccAnalyticsprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAnalyticsprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAnalyticsprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAnalyticsprofileExist("citrixadc_analyticsprofile.tf_analyticsprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_analyticsprofile", "name", "my_analyticsprofile"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_analyticsprofile", "type", "webinsight"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_analyticsprofile", "httppagetracking", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_analyticsprofile", "httpurl", "DISABLED"),
				),
			},
			{
				Config: testAccAnalyticsprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAnalyticsprofileExist("citrixadc_analyticsprofile.tf_analyticsprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_analyticsprofile", "name", "my_analyticsprofile"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_analyticsprofile", "type", "webinsight"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_analyticsprofile", "httppagetracking", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_analyticsprofile", "httpurl", "ENABLED"),
				),
			},
		},
	})
}

func TestAccAnalyticsprofile_import(t *testing.T) {
	const resAddr = "citrixadc_analyticsprofile.tf_analyticsprofile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAnalyticsprofileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccAnalyticsprofile_basic},
			{
				Config:            testAccAnalyticsprofile_basic,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// analyticsauthtoken_wo_version is a write-only version tracker that
				// NITRO does not echo back, so it cannot round-trip through import.
				ImportStateVerifyIgnore: []string{"analyticsauthtoken_wo_version"},
			},
		},
	})
}

func testAccCheckAnalyticsprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No analyticsprofile name is set")
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
		data, err := client.FindResource("analyticsprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("analyticsprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckAnalyticsprofileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_analyticsprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("analyticsprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("analyticsprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAnalyticsprofileDataSource_basic = `

	resource "citrixadc_analyticsprofile" "tf_analyticsprofile" {
		name             = "my_analyticsprofile"
		type             = "webinsight"
		httppagetracking = "DISABLED"
		httpurl          = "DISABLED"
	}
	
	data "citrixadc_analyticsprofile" "tf_analyticsprofile" {
		name = citrixadc_analyticsprofile.tf_analyticsprofile.name
	}
`

func TestAccAnalyticsprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAnalyticsprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAnalyticsprofileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_analyticsprofile.tf_analyticsprofile", "name", "my_analyticsprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_analyticsprofile.tf_analyticsprofile", "type", "webinsight"),
					resource.TestCheckResourceAttr("data.citrixadc_analyticsprofile.tf_analyticsprofile", "httppagetracking", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_analyticsprofile.tf_analyticsprofile", "httpurl", "DISABLED"),
				),
			},
		},
	})
}

const testAccAnalyticsprofile_analyticsauthtoken_step1 = `
	variable "analyticsprofile_analyticsauthtoken" {
		type      = string
		sensitive = true
	}

	resource "citrixadc_analyticsprofile" "tf_analyticsprofile" {
		name                = "my_analyticsprofile"
		type                = "webinsight"
		httppagetracking    = "DISABLED"
		httpurl             = "DISABLED"
		analyticsauthtoken  = var.analyticsprofile_analyticsauthtoken
	}
`

const testAccAnalyticsprofile_analyticsauthtoken_step2 = `
	variable "analyticsprofile_analyticsauthtoken_2" {
		type      = string
		sensitive = true
	}

	resource "citrixadc_analyticsprofile" "tf_analyticsprofile" {
		name                = "my_analyticsprofile"
		type                = "webinsight"
		httppagetracking    = "DISABLED"
		httpurl             = "DISABLED"
		analyticsauthtoken  = var.analyticsprofile_analyticsauthtoken_2
	}
`

func TestAccAnalyticsprofile_analyticsauthtoken_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_analyticsprofile_analyticsauthtoken", "value1")
	t.Setenv("TF_VAR_analyticsprofile_analyticsauthtoken_2", "value2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAnalyticsprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAnalyticsprofile_analyticsauthtoken_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAnalyticsprofileExist("citrixadc_analyticsprofile.tf_analyticsprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_analyticsprofile", "name", "my_analyticsprofile"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_analyticsprofile", "type", "webinsight"),
				),
			},
			{
				Config: testAccAnalyticsprofile_analyticsauthtoken_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAnalyticsprofileExist("citrixadc_analyticsprofile.tf_analyticsprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_analyticsprofile", "name", "my_analyticsprofile"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_analyticsprofile", "type", "webinsight"),
				),
			},
		},
	})
}

const testAccAnalyticsprofile_analyticsauthtoken_wo_step1 = `
	variable "analyticsprofile_analyticsauthtoken_wo" {
		type      = string
		sensitive = true
	}

	resource "citrixadc_analyticsprofile" "tf_analyticsprofile" {
		name                          = "my_analyticsprofile"
		type                          = "webinsight"
		httppagetracking              = "DISABLED"
		httpurl                       = "DISABLED"
		analyticsauthtoken_wo         = var.analyticsprofile_analyticsauthtoken_wo
		analyticsauthtoken_wo_version = 1
	}
`

const testAccAnalyticsprofile_analyticsauthtoken_wo_step2 = `
	variable "analyticsprofile_analyticsauthtoken_wo_2" {
		type      = string
		sensitive = true
	}

	resource "citrixadc_analyticsprofile" "tf_analyticsprofile" {
		name                          = "my_analyticsprofile"
		type                          = "webinsight"
		httppagetracking              = "DISABLED"
		httpurl                       = "DISABLED"
		analyticsauthtoken_wo         = var.analyticsprofile_analyticsauthtoken_wo_2
		analyticsauthtoken_wo_version = 2
	}
`

func TestAccAnalyticsprofile_analyticsauthtoken_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_analyticsprofile_analyticsauthtoken_wo", "ephemeral_value1")
	t.Setenv("TF_VAR_analyticsprofile_analyticsauthtoken_wo_2", "ephemeral_value2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAnalyticsprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAnalyticsprofile_analyticsauthtoken_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAnalyticsprofileExist("citrixadc_analyticsprofile.tf_analyticsprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_analyticsprofile", "name", "my_analyticsprofile"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_analyticsprofile", "type", "webinsight"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_analyticsprofile", "analyticsauthtoken_wo_version", "1"),
				),
			},
			{
				Config: testAccAnalyticsprofile_analyticsauthtoken_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAnalyticsprofileExist("citrixadc_analyticsprofile.tf_analyticsprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_analyticsprofile", "name", "my_analyticsprofile"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_analyticsprofile", "type", "webinsight"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_analyticsprofile", "analyticsauthtoken_wo_version", "2"),
				),
			},
		},
	})
}

func TestAccAnalyticsprofile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckAnalyticsprofileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccAnalyticsprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAnalyticsprofileExist("citrixadc_analyticsprofile.tf_analyticsprofile", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAnalyticsprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAnalyticsprofileExist("citrixadc_analyticsprofile.tf_analyticsprofile", nil),
				),
			},
		},
	})
}

// The analyticsprofile unset test covers the type-independent unset-eligible
// attributes for type "webinsight": every attribute wired into
// attributesToUnset that applies cleanly for a webinsight profile. Six other
// wired attributes (auditlogs, cqareporting, events, metrics, outputmode,
// topn) are NOT type-independent -- each requires a different, non-webinsight
// type (timeseries/tcpinsight/streaminsight) and NITRO rejects them on a
// webinsight profile with "Argument pre-requisite missing" -- so they cannot
// be exercised on this single-resource unset test and are excluded here.
const testAccAnalyticsprofile_unset_step1 = `
resource "citrixadc_analyticsprofile" "tf_unset" {
  name                       = "tf_test_analyticsprofile_unset"
  type                       = "webinsight"
  allhttpheaders             = "ENABLED"
  grpcstatus                 = "ENABLED"
  httpauthentication         = "ENABLED"
  httpclientsidemeasurements = "ENABLED"
  httpcontenttype            = "ENABLED"
  httpcookie                 = "ENABLED"
  httpdomainname             = "ENABLED"
  httphost                   = "ENABLED"
  httplocation               = "ENABLED"
  httpmethod                 = "ENABLED"
  httppagetracking           = "ENABLED"
  httpreferer                = "ENABLED"
  httpsetcookie              = "ENABLED"
  httpsetcookie2             = "ENABLED"
  httpurl                    = "ENABLED"
  httpurlquery               = "ENABLED"
  httpuseragent              = "ENABLED"
  httpvia                    = "ENABLED"
  httpxforwardedforheader    = "ENABLED"
  integratedcache            = "ENABLED"
  urlcategory                = "ENABLED"
}
`

const testAccAnalyticsprofile_unset_step2 = `
resource "citrixadc_analyticsprofile" "tf_unset" {
  name = "tf_test_analyticsprofile_unset"
  type = "webinsight"
  # All unset-eligible attributes removed from config -> the provider must
  # unset them (revert to NITRO defaults, "DISABLED").
}
`

func TestAccAnalyticsprofile_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAnalyticsprofileDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccAnalyticsprofile_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAnalyticsprofileExist("citrixadc_analyticsprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "allhttpheaders", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "grpcstatus", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httpauthentication", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httpclientsidemeasurements", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httpcontenttype", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httpcookie", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httpdomainname", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httphost", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httplocation", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httpmethod", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httppagetracking", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httpreferer", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httpsetcookie", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httpsetcookie2", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httpurl", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httpurlquery", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httpuseragent", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httpvia", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httpxforwardedforheader", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "integratedcache", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "urlcategory", "ENABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccAnalyticsprofile_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAnalyticsprofileExist("citrixadc_analyticsprofile.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "allhttpheaders", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "grpcstatus", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httpauthentication", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httpclientsidemeasurements", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httpcontenttype", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httpcookie", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httpdomainname", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httphost", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httplocation", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httpmethod", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httppagetracking", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httpreferer", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httpsetcookie", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httpsetcookie2", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httpurl", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httpurlquery", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httpuseragent", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httpvia", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "httpxforwardedforheader", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "integratedcache", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_analyticsprofile.tf_unset", "urlcategory", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAnalyticsprofileADCValue("tf_test_analyticsprofile_unset", "allhttpheaders", "DISABLED"),
					testAccCheckAnalyticsprofileADCValue("tf_test_analyticsprofile_unset", "httpurl", "DISABLED"),
					testAccCheckAnalyticsprofileADCValue("tf_test_analyticsprofile_unset", "integratedcache", "DISABLED"),
					testAccCheckAnalyticsprofileADCValue("tf_test_analyticsprofile_unset", "urlcategory", "DISABLED"),
				),
			},
		},
	})
}

// testAccCheckAnalyticsprofileADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckAnalyticsprofileADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Analyticsprofile.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("analyticsprofile %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("analyticsprofile %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}
