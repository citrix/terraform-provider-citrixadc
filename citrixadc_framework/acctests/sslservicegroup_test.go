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
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testAccSslservicegroup_basic = `
	resource "citrixadc_sslservicegroup" "tf_sslservicegroup" {
		servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
		sesstimeout = 50
		sessreuse = "ENABLED"
		ssl3 = "ENABLED"
		snienable = "ENABLED"
		serverauth = "ENABLED"
		sendclosenotify = "YES"
		strictsigdigestcheck = "ENABLED"
	}

	resource "citrixadc_servicegroup" "tf_servicegroup" {
		servicegroupname = "tf_servicegroup"
		servicetype = "SSL"
	}
`

func TestAccSslservicegroup_basic(t *testing.T) {
	if adcTestbed != "STANDALONE_NON_DEFAULT_SSL_PROFILE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_NON_DEFAULT_SSL_PROFILE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslservicegroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslservicegroup_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslservicegroupExist("citrixadc_sslservicegroup.tf_sslservicegroup", nil),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_sslservicegroup", "servicegroupname", "tf_servicegroup"),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_sslservicegroup", "sesstimeout", "50"),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_sslservicegroup", "sessreuse", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_sslservicegroup", "ssl3", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_sslservicegroup", "snienable", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_sslservicegroup", "serverauth", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_sslservicegroup", "sendclosenotify", "YES"),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_sslservicegroup", "strictsigdigestcheck", "ENABLED"),
				),
			},
		},
	})
}

func TestAccSslservicegroup_sdkv2StateUpgrade(t *testing.T) {
	if adcTestbed != "STANDALONE_NON_DEFAULT_SSL_PROFILE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_NON_DEFAULT_SSL_PROFILE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckSslservicegroupDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccSslservicegroup_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSslservicegroupExist("citrixadc_sslservicegroup.tf_sslservicegroup", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccSslservicegroup_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSslservicegroupExist("citrixadc_sslservicegroup.tf_sslservicegroup", nil)),
			},
		},
	})
}

const testAccSslservicegroup_unset_step1 = `
	resource "citrixadc_servicegroup" "tf_servicegroup" {
		servicegroupname = "tf_servicegroup"
		servicetype      = "SSL"
	}

	resource "citrixadc_sslservicegroup" "tf_unset" {
		servicegroupname     = citrixadc_servicegroup.tf_servicegroup.servicegroupname
		# sesstimeout has a NITRO co-requisite (sessReuse==ENABLED); sessreuse and
		# sesstimeout cannot both be non-default at once, so keep sessreuse ENABLED.
		sessreuse            = "ENABLED"
		sesstimeout          = 100
		ssl3                 = "DISABLED"
		tls1                 = "DISABLED"
		tls11                = "DISABLED"
		tls12                = "DISABLED"
		tls13                = "ENABLED"
		snienable            = "ENABLED"
		# ocspstapling omitted: NITRO 3761 "Ocsp stapling option is not supported
		# for SSL Service" -- not settable on a bare SSL servicegroup.
		serverauth           = "ENABLED"
		sendclosenotify      = "NO"
		strictsigdigestcheck = "ENABLED"
		sslclientlogs        = "ENABLED"
	}
`

const testAccSslservicegroup_unset_step2 = `
	resource "citrixadc_servicegroup" "tf_servicegroup" {
		servicegroupname = "tf_servicegroup"
		servicetype      = "SSL"
	}

	resource "citrixadc_sslservicegroup" "tf_unset" {
		servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to documented NITRO defaults).
	}
`

func TestAccSslservicegroup_unset(t *testing.T) {
	if adcTestbed != "STANDALONE_NON_DEFAULT_SSL_PROFILE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_NON_DEFAULT_SSL_PROFILE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslservicegroupDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccSslservicegroup_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslservicegroupExist("citrixadc_sslservicegroup.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_unset", "sessreuse", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_unset", "sesstimeout", "100"),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_unset", "ssl3", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_unset", "tls1", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_unset", "tls11", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_unset", "tls12", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_unset", "tls13", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_unset", "snienable", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_unset", "serverauth", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_unset", "sendclosenotify", "NO"),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_unset", "strictsigdigestcheck", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_unset", "sslclientlogs", "ENABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccSslservicegroup_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslservicegroupExist("citrixadc_sslservicegroup.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_unset", "sessreuse", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_unset", "sesstimeout", "300"),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_unset", "ssl3", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_unset", "tls1", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_unset", "tls11", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_unset", "tls12", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_unset", "tls13", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_unset", "snienable", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_unset", "serverauth", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_unset", "sendclosenotify", "YES"),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_unset", "strictsigdigestcheck", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_unset", "sslclientlogs", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckSslservicegroupADCValue("tf_servicegroup", "ssl3", "ENABLED"),
					testAccCheckSslservicegroupADCValue("tf_servicegroup", "sesstimeout", "300"),
					testAccCheckSslservicegroupADCValue("tf_servicegroup", "sendclosenotify", "YES"),
				),
			},
		},
	})
}

// testAccCheckSslservicegroupADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckSslservicegroupADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Sslservicegroup.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("sslservicegroup %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("sslservicegroup %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func testAccCheckSslservicegroupExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslservicegroup name is set")
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
		data, err := client.FindResource(service.Sslservicegroup.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("sslservicegroup %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslservicegroupDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslservicegroup" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Sslservicegroup.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslservicegroup %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccSslservicegroup_import(t *testing.T) {
	if adcTestbed != "STANDALONE_NON_DEFAULT_SSL_PROFILE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_NON_DEFAULT_SSL_PROFILE.", adcTestbed)
	}
	const resAddr = "citrixadc_sslservicegroup.tf_sslservicegroup"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslservicegroupDestroy,
		Steps: []resource.TestStep{
			{Config: testAccSslservicegroup_basic},
			{
				Config:                  testAccSslservicegroup_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

const testAccSslservicegroupDataSource_basic = `
	resource "citrixadc_sslservicegroup" "tf_sslservicegroup" {
		servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
		sesstimeout = 50
		sessreuse = "ENABLED"
		ssl3 = "ENABLED"
		snienable = "ENABLED"
		serverauth = "ENABLED"
		sendclosenotify = "YES"
		strictsigdigestcheck = "ENABLED"
	}

	resource "citrixadc_servicegroup" "tf_servicegroup" {
		servicegroupname = "tf_servicegroup"
		servicetype = "SSL"
	}

	data "citrixadc_sslservicegroup" "tf_sslservicegroup_datasource" {
		servicegroupname = citrixadc_sslservicegroup.tf_sslservicegroup.servicegroupname
	}
`

func TestAccSslservicegroupDataSource_basic(t *testing.T) {
	if adcTestbed != "STANDALONE_NON_DEFAULT_SSL_PROFILE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE_NON_DEFAULT_SSL_PROFILE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslservicegroupDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslservicegroup.tf_sslservicegroup_datasource", "servicegroupname", "tf_servicegroup"),
					resource.TestCheckResourceAttr("data.citrixadc_sslservicegroup.tf_sslservicegroup_datasource", "sesstimeout", "50"),
					resource.TestCheckResourceAttr("data.citrixadc_sslservicegroup.tf_sslservicegroup_datasource", "sessreuse", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_sslservicegroup.tf_sslservicegroup_datasource", "ssl3", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_sslservicegroup.tf_sslservicegroup_datasource", "snienable", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_sslservicegroup.tf_sslservicegroup_datasource", "serverauth", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_sslservicegroup.tf_sslservicegroup_datasource", "sendclosenotify", "YES"),
					resource.TestCheckResourceAttr("data.citrixadc_sslservicegroup.tf_sslservicegroup_datasource", "strictsigdigestcheck", "ENABLED"),
					resource.TestCheckResourceAttrSet("data.citrixadc_sslservicegroup.tf_sslservicegroup_datasource", "id"),
				),
			},
		},
	})
}
