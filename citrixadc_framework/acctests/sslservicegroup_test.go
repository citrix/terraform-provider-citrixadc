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
		sslclientlogs = "ENABLED"
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
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup.tf_sslservicegroup", "sslclientlogs", "ENABLED"),
				),
			},
		},
	})
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
		sslclientlogs = "ENABLED"
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
					resource.TestCheckResourceAttr("data.citrixadc_sslservicegroup.tf_sslservicegroup_datasource", "sslclientlogs", "ENABLED"),
					resource.TestCheckResourceAttrSet("data.citrixadc_sslservicegroup.tf_sslservicegroup_datasource", "id"),
				),
			},
		},
	})
}
