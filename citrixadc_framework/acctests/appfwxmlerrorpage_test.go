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

const testAccAppfwxmlerrorpage_basic = `
	resource "citrixadc_systemfile" "tf_xmlerrorpage" {
		filename     = "appfwxmlerrorpage.xml"
		filelocation = "/var/tmp"
		filecontent  = file("testdata/appfwxmlerrorpage.xml")
	}
	resource "citrixadc_appfwxmlerrorpage" "tf_appfwxmlerrorpage" {
		name       = "tf_appfwxmlerrorpage"
		src        = "local://appfwxmlerrorpage.xml"
		depends_on = [citrixadc_systemfile.tf_xmlerrorpage]
		comment    = "TestingExample"
	}
`

func TestAccAppfwxmlerrorpage_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwxmlerrorpageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwxmlerrorpage_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwxmlerrorpageExist("citrixadc_appfwxmlerrorpage.tf_appfwxmlerrorpage", nil),
				),
			},
		},
	})
}

func testAccCheckAppfwxmlerrorpageExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwxmlerrorpage name is set")
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
		data, err := client.FindResource(service.Appfwxmlerrorpage.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appfwxmlerrorpage %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwxmlerrorpageDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwxmlerrorpage" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Appfwxmlerrorpage.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwxmlerrorpage %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAppfwxmlerrorpageDataSource_basic = `
	resource "citrixadc_systemfile" "tf_xmlerrorpage_ds" {
		filename     = "appfwxmlerrorpage_ds.xml"
		filelocation = "/var/tmp"
		filecontent  = file("testdata/appfwxmlerrorpage.xml")
	}
	resource "citrixadc_appfwxmlerrorpage" "tf_appfwxmlerrorpage_ds" {
		name       = "tf_appfwxmlerrorpage_ds"
		src        = "local://appfwxmlerrorpage_ds.xml"
		depends_on = [citrixadc_systemfile.tf_xmlerrorpage_ds]
		comment    = "TestingExampleDataSource"
	}
	
	data "citrixadc_appfwxmlerrorpage" "tf_appfwxmlerrorpage_ds" {
		name = citrixadc_appfwxmlerrorpage.tf_appfwxmlerrorpage_ds.name
	}
`

func TestAccAppfwxmlerrorpageDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwxmlerrorpageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwxmlerrorpageDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appfwxmlerrorpage.tf_appfwxmlerrorpage_ds", "name", "tf_appfwxmlerrorpage_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwxmlerrorpage.tf_appfwxmlerrorpage_ds", "src", "appfwxmlerrorpage_ds.xml"),
				),
			},
		},
	})
}
