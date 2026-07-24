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

const testAccAppflowparam_basic = `

resource "citrixadc_appflowparam" "tf_appflowparam" {
	templaterefresh     = 200
	flowrecordinterval  = 200
	httpcookie          = "ENABLED"
	httplocation        = "ENABLED"
	}
  
`
const testAccAppflowparam_update = `

resource "citrixadc_appflowparam" "tf_appflowparam" {
	templaterefresh     = 600
	flowrecordinterval  = 100
	httpcookie          = "DISABLED"
	httplocation        = "DISABLED"
	}
`

func TestAccAppflowparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAppflowparam_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowparamExist("citrixadc_appflowparam.tf_appflowparam", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "templaterefresh", "200"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "flowrecordinterval", "200"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "httpcookie", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "httplocation", "ENABLED"),
				),
			},
			{
				Config: testAccAppflowparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowparamExist("citrixadc_appflowparam.tf_appflowparam", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "templaterefresh", "600"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "flowrecordinterval", "100"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "httpcookie", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "httplocation", "DISABLED"),
				),
			},
		},
	})
}

func TestAccAppflowparam_import(t *testing.T) {
	const resAddr = "citrixadc_appflowparam.tf_appflowparam"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// appflowparam is a singleton (PARTIAL) resource; it has no delete on the
		// ADC, so there is no CheckDestroy.
		Steps: []resource.TestStep{
			{Config: testAccAppflowparam_basic},
			{
				// Import id is the synthetic constant "appflowparam-config" set in
				// Create; ImportStatePassthroughID uses the stored id.
				Config:            testAccAppflowparam_basic,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
				// analyticsauthtoken_wo_version is a write-only version tracker
				// (default 1) that is not returned by NITRO and cannot round-trip
				// through import.
				ImportStateVerifyIgnore: []string{"analyticsauthtoken_wo_version"},
			},
		},
	})
}

func testAccCheckAppflowparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appflowparam name is set")
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
		data, err := client.FindResource(service.Appflowparam.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appflowparam %s not found", n)
		}

		return nil
	}
}

const testAccAppflowparamDataSource_basic = `
	resource "citrixadc_appflowparam" "tf_appflowparam" {
		templaterefresh     = 200
		flowrecordinterval  = 200
		httpcookie          = "ENABLED"
		httplocation        = "ENABLED"
	}
	
	data "citrixadc_appflowparam" "tf_appflowparam" {
		depends_on = [citrixadc_appflowparam.tf_appflowparam]
	}
`

func TestAccAppflowparamDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAppflowparamDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appflowparam.tf_appflowparam", "templaterefresh", "200"),
					resource.TestCheckResourceAttr("data.citrixadc_appflowparam.tf_appflowparam", "flowrecordinterval", "200"),
					resource.TestCheckResourceAttr("data.citrixadc_appflowparam.tf_appflowparam", "httpcookie", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_appflowparam.tf_appflowparam", "httplocation", "ENABLED"),
				),
			},
		},
	})
}

// Test backward-compatible path: using analyticsauthtoken (Sensitive attribute)
const testAccAppflowparam_analyticsauthtoken_step1 = `

	variable "appflowparam_analyticsauthtoken" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_appflowparam" "tf_appflowparam" {
		templaterefresh    = 200
		flowrecordinterval = 200
		httpcookie         = "ENABLED"
		httplocation       = "ENABLED"
		analyticsauthtoken = var.appflowparam_analyticsauthtoken
	}
`

// Update backward-compatible path: change analyticsauthtoken value
const testAccAppflowparam_analyticsauthtoken_step2 = `

	variable "appflowparam_analyticsauthtoken_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_appflowparam" "tf_appflowparam" {
		templaterefresh    = 200
		flowrecordinterval = 200
		httpcookie         = "ENABLED"
		httplocation       = "ENABLED"
		analyticsauthtoken = var.appflowparam_analyticsauthtoken_2
	}
`

func TestAccAppflowparam_analyticsauthtoken_backward_compat(t *testing.T) {
	t.Setenv("TF_VAR_appflowparam_analyticsauthtoken", "authtoken_value1")
	t.Setenv("TF_VAR_appflowparam_analyticsauthtoken_2", "authtoken_value2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAppflowparam_analyticsauthtoken_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowparamExist("citrixadc_appflowparam.tf_appflowparam", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "templaterefresh", "200"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "flowrecordinterval", "200"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "httpcookie", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "httplocation", "ENABLED"),
				),
			},
			{
				Config: testAccAppflowparam_analyticsauthtoken_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowparamExist("citrixadc_appflowparam.tf_appflowparam", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "templaterefresh", "200"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "flowrecordinterval", "200"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "httpcookie", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "httplocation", "ENABLED"),
				),
			},
		},
	})
}

// Test ephemeral path: using analyticsauthtoken_wo (WriteOnly attribute) with version tracker
const testAccAppflowparam_analyticsauthtoken_wo_step1 = `

	variable "appflowparam_analyticsauthtoken_wo" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_appflowparam" "tf_appflowparam" {
		templaterefresh              = 200
		flowrecordinterval           = 200
		httpcookie                   = "ENABLED"
		httplocation                 = "ENABLED"
		analyticsauthtoken_wo        = var.appflowparam_analyticsauthtoken_wo
		analyticsauthtoken_wo_version = 1
	}
`

// Update ephemeral path: bump version to trigger update with new token
const testAccAppflowparam_analyticsauthtoken_wo_step2 = `

	variable "appflowparam_analyticsauthtoken_wo_2" {
	  type      = string
	  sensitive = true
	}

	resource "citrixadc_appflowparam" "tf_appflowparam" {
		templaterefresh              = 200
		flowrecordinterval           = 200
		httpcookie                   = "ENABLED"
		httplocation                 = "ENABLED"
		analyticsauthtoken_wo        = var.appflowparam_analyticsauthtoken_wo_2
		analyticsauthtoken_wo_version = 2
	}
`

func TestAccAppflowparam_analyticsauthtoken_wo_ephemeral(t *testing.T) {
	t.Setenv("TF_VAR_appflowparam_analyticsauthtoken_wo", "ephemeral_token1")
	t.Setenv("TF_VAR_appflowparam_analyticsauthtoken_wo_2", "ephemeral_token2")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAppflowparam_analyticsauthtoken_wo_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowparamExist("citrixadc_appflowparam.tf_appflowparam", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "templaterefresh", "200"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "flowrecordinterval", "200"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "httpcookie", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "httplocation", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "analyticsauthtoken_wo_version", "1"),
				),
			},
			{
				Config: testAccAppflowparam_analyticsauthtoken_wo_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowparamExist("citrixadc_appflowparam.tf_appflowparam", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "templaterefresh", "200"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "flowrecordinterval", "200"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "httpcookie", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "httplocation", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_appflowparam", "analyticsauthtoken_wo_version", "2"),
				),
			},
		},
	})
}

func TestAccAppflowparam_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccAppflowparam_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowparamExist("citrixadc_appflowparam.tf_appflowparam", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccAppflowparam_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowparamExist("citrixadc_appflowparam.tf_appflowparam", nil),
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// Unset acceptance test: proves ?action=unset reverts every unset-eligible
// attribute of appflowparam back to its NITRO default.
// ---------------------------------------------------------------------------

const testAccAppflowparam_unset_step1 = `
resource "citrixadc_appflowparam" "tf_unset" {
  aaausername                         = "ENABLED"
  appnamerefresh                      = 300
  cacheinsight                        = "ENABLED"
  clienttrafficonly                   = "YES"
  connectionchaining                  = "ENABLED"
  cqareporting                        = "ENABLED"
  distributedtracing                  = "ENABLED"
  disttracingsamplingrate             = 50
  emailaddress                        = "ENABLED"
  flowrecordinterval                  = 120
  gxsessionreporting                  = "ENABLED"
  httpauthorization                   = "ENABLED"
  httpcontenttype                     = "ENABLED"
  httpcookie                          = "ENABLED"
  httpdomain                          = "ENABLED"
  httphost                            = "ENABLED"
  httplocation                        = "ENABLED"
  httpmethod                          = "ENABLED"
  httpquerywithurl                    = "ENABLED"
  httpreferer                         = "ENABLED"
  httpsetcookie                       = "ENABLED"
  httpsetcookie2                      = "ENABLED"
  httpurl                             = "ENABLED"
  httpuseragent                       = "ENABLED"
  httpvia                             = "ENABLED"
  httpxforwardedfor                   = "ENABLED"
  identifiername                      = "ENABLED"
  identifiersessionname               = "ENABLED"
  logstreamovernsip                   = "ENABLED"
  lsnlogging                          = "ENABLED"
  securityinsightrecordinterval       = 300
  securityinsighttraffic              = "ENABLED"
  skipcacheredirectionhttptransaction = "ENABLED"
  subscriberawareness                 = "ENABLED"
  subscriberidobfuscation             = "ENABLED"
  tcpattackcounterinterval            = 60
  templaterefresh                     = 300
  timeseriesovernsip                  = "ENABLED"
  udppmtu                             = 1400
  urlcategory                         = "ENABLED"
  usagerecordinterval                 = 60
  videoinsight                        = "ENABLED"
  websaasappusagereporting            = "ENABLED"
}
`

const testAccAppflowparam_unset_step2 = `
resource "citrixadc_appflowparam" "tf_unset" {
  # all unset-eligible attributes removed from config -> provider must unset them
}
`

func TestAccAppflowparam_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		// appflowparam is a singleton (PARTIAL) resource; it has no delete on the
		// ADC, so there is no CheckDestroy.
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				// Non-default values apply and persist.
				Config: testAccAppflowparam_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowparamExist("citrixadc_appflowparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "aaausername", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "appnamerefresh", "300"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "cacheinsight", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "clienttrafficonly", "YES"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "connectionchaining", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "cqareporting", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "distributedtracing", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "disttracingsamplingrate", "50"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "emailaddress", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "flowrecordinterval", "120"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "gxsessionreporting", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "httpauthorization", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "httpcontenttype", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "httpcookie", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "httpdomain", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "httphost", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "httplocation", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "httpmethod", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "httpquerywithurl", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "httpreferer", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "httpsetcookie", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "httpsetcookie2", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "httpurl", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "httpuseragent", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "httpvia", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "httpxforwardedfor", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "identifiername", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "identifiersessionname", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "logstreamovernsip", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "lsnlogging", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "securityinsightrecordinterval", "300"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "securityinsighttraffic", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "skipcacheredirectionhttptransaction", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "subscriberawareness", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "subscriberidobfuscation", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "tcpattackcounterinterval", "60"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "templaterefresh", "300"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "timeseriesovernsip", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "udppmtu", "1400"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "urlcategory", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "usagerecordinterval", "60"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "videoinsight", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "websaasappusagereporting", "ENABLED"),
				),
			},
			{
				// Removing them must unset -> state reverts to NITRO defaults,
				// and the implicit post-apply plan must be empty.
				Config: testAccAppflowparam_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppflowparamExist("citrixadc_appflowparam.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "aaausername", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "appnamerefresh", "600"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "cacheinsight", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "clienttrafficonly", "NO"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "connectionchaining", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "cqareporting", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "distributedtracing", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "disttracingsamplingrate", "0"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "emailaddress", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "flowrecordinterval", "60"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "gxsessionreporting", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "httpauthorization", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "httpcontenttype", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "httpcookie", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "httpdomain", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "httphost", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "httplocation", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "httpmethod", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "httpquerywithurl", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "httpreferer", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "httpsetcookie", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "httpsetcookie2", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "httpurl", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "httpuseragent", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "httpvia", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "httpxforwardedfor", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "identifiername", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "identifiersessionname", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "logstreamovernsip", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "lsnlogging", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "securityinsightrecordinterval", "600"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "securityinsighttraffic", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "skipcacheredirectionhttptransaction", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "subscriberawareness", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "subscriberidobfuscation", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "tcpattackcounterinterval", "0"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "templaterefresh", "600"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "timeseriesovernsip", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "udppmtu", "1472"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "urlcategory", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "usagerecordinterval", "0"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "videoinsight", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_appflowparam.tf_unset", "websaasappusagereporting", "DISABLED"),
					// Independent appliance-level confirmation the unset actually reverted each attr.
					testAccCheckAppflowparamADCValue("aaausername", "DISABLED"),
					testAccCheckAppflowparamADCValue("appnamerefresh", "600"),
					testAccCheckAppflowparamADCValue("cacheinsight", "DISABLED"),
					testAccCheckAppflowparamADCValue("clienttrafficonly", "NO"),
					testAccCheckAppflowparamADCValue("connectionchaining", "DISABLED"),
					testAccCheckAppflowparamADCValue("cqareporting", "DISABLED"),
					testAccCheckAppflowparamADCValue("distributedtracing", "DISABLED"),
					testAccCheckAppflowparamADCValue("disttracingsamplingrate", "0"),
					testAccCheckAppflowparamADCValue("emailaddress", "DISABLED"),
					testAccCheckAppflowparamADCValue("flowrecordinterval", "60"),
					testAccCheckAppflowparamADCValue("gxsessionreporting", "DISABLED"),
					testAccCheckAppflowparamADCValue("httpauthorization", "DISABLED"),
					testAccCheckAppflowparamADCValue("httpcontenttype", "DISABLED"),
					testAccCheckAppflowparamADCValue("httpcookie", "DISABLED"),
					testAccCheckAppflowparamADCValue("httpdomain", "DISABLED"),
					testAccCheckAppflowparamADCValue("httphost", "DISABLED"),
					testAccCheckAppflowparamADCValue("httplocation", "DISABLED"),
					testAccCheckAppflowparamADCValue("httpmethod", "DISABLED"),
					testAccCheckAppflowparamADCValue("httpquerywithurl", "DISABLED"),
					testAccCheckAppflowparamADCValue("httpreferer", "DISABLED"),
					testAccCheckAppflowparamADCValue("httpsetcookie", "DISABLED"),
					testAccCheckAppflowparamADCValue("httpsetcookie2", "DISABLED"),
					testAccCheckAppflowparamADCValue("httpurl", "DISABLED"),
					testAccCheckAppflowparamADCValue("httpuseragent", "DISABLED"),
					testAccCheckAppflowparamADCValue("httpvia", "DISABLED"),
					testAccCheckAppflowparamADCValue("httpxforwardedfor", "DISABLED"),
					testAccCheckAppflowparamADCValue("identifiername", "DISABLED"),
					testAccCheckAppflowparamADCValue("identifiersessionname", "DISABLED"),
					testAccCheckAppflowparamADCValue("logstreamovernsip", "DISABLED"),
					testAccCheckAppflowparamADCValue("lsnlogging", "DISABLED"),
					testAccCheckAppflowparamADCValue("securityinsightrecordinterval", "600"),
					testAccCheckAppflowparamADCValue("securityinsighttraffic", "DISABLED"),
					testAccCheckAppflowparamADCValue("skipcacheredirectionhttptransaction", "DISABLED"),
					testAccCheckAppflowparamADCValue("subscriberawareness", "DISABLED"),
					testAccCheckAppflowparamADCValue("subscriberidobfuscation", "DISABLED"),
					testAccCheckAppflowparamADCValue("tcpattackcounterinterval", "0"),
					testAccCheckAppflowparamADCValue("templaterefresh", "600"),
					testAccCheckAppflowparamADCValue("timeseriesovernsip", "DISABLED"),
					testAccCheckAppflowparamADCValue("udppmtu", "1472"),
					testAccCheckAppflowparamADCValue("urlcategory", "DISABLED"),
					testAccCheckAppflowparamADCValue("usagerecordinterval", "0"),
					testAccCheckAppflowparamADCValue("videoinsight", "DISABLED"),
					testAccCheckAppflowparamADCValue("websaasappusagereporting", "DISABLED"),
				),
			},
		},
	})
}

// testAccCheckAppflowparamADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckAppflowparamADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Appflowparam.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("appflowparam not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("appflowparam: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}
