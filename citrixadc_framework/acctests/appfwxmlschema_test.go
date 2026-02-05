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

const testAccAppfwxmlschema_basic = `
	resource "citrixadc_systemfile" "tf_xmlschema" {
		filename     = "appfwxmlschema.xml"
		filelocation = "/var/tmp"
		filecontent  = file("testdata/appfwxmlschema.xml")
	}
	resource "citrixadc_appfwxmlschema" "tf_appfwxmlschema" {
		name       = "tf_appfwxmlschema"
		src        = "local://appfwxmlschema.xml"
		depends_on = [citrixadc_systemfile.tf_xmlschema]
		comment    = "TestingExample"
	}
`

func TestAccAppfwxmlschema_basic(t *testing.T) {
	t.Skip("TODO: Need to find a way to test this resource!")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwxmlschemaDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwxmlschema_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwxmlschemaExist("citrixadc_appfwxmlschema.tf_appfwxmlschema", nil),
				),
			},
		},
	})
}

func testAccCheckAppfwxmlschemaExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwxmlschema name is set")
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
		data, err := client.FindResource(service.Appfwxmlschema.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appfwxmlschema %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwxmlschemaDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwxmlschema" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Appfwxmlschema.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwxmlschema %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccAppfwxmlschemaDataSource_basic = `
	resource "citrixadc_systemfile" "tf_xmlschema_ds" {
		filename     = "appfwxmlschema_ds.xml"
		filelocation = "/var/tmp"
		filecontent  = file("testdata/appfwxmlschema.xml")
	}
	resource "citrixadc_appfwxmlschema" "tf_appfwxmlschema_ds" {
		name       = "tf_appfwxmlschema_ds"
		src        = "local://appfwxmlschema_ds.xml"
		depends_on = [citrixadc_systemfile.tf_xmlschema_ds]
		comment    = "TestingDataSource"
	}
	
	data "citrixadc_appfwxmlschema" "tf_appfwxmlschema_ds" {
		name = citrixadc_appfwxmlschema.tf_appfwxmlschema_ds.name
	}
`

func TestAccAppfwxmlschemaDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwxmlschemaDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwxmlschemaDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appfwxmlschema.tf_appfwxmlschema_ds", "name", "tf_appfwxmlschema_ds"),
				),
			},
		},
	})
}
