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

const testAccStreamidentifier_basic = `

	resource "citrixadc_streamselector" "tf_streamselector" {
		name = "my_streamselector"
		rule = ["HTTP.REQ.URL", "CLIENT.IP.SRC"]
	}
	resource "citrixadc_streamidentifier" "tf_streamidentifier" {
		name         = "my_streamidentifier"
		selectorname = citrixadc_streamselector.tf_streamselector.name
		samplecount  = 10
		sort         = "CONNECTIONS"
		snmptrap     = "ENABLED"
		loglimit	 = 500
		loginterval = 60
		log = "NONE"

	}
  
`

const testAccStreamidentifier_update = `

	resource "citrixadc_streamselector" "tf_streamselector" {
		name = "my_streamselector"
		rule = ["HTTP.REQ.URL", "CLIENT.IP.SRC"]
	}

	resource "citrixadc_streamidentifier" "tf_streamidentifier" {
		name         = "my_streamidentifier"
		selectorname = citrixadc_streamselector.tf_streamselector.name
		samplecount  = 20
		sort         = "REQUESTS"
		snmptrap     = "DISABLED"
		loglimit	 = 600
		loginterval = 120
		log = "SYSLOG"
	}
  
`

func TestAccStreamidentifier_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckStreamidentifierDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccStreamidentifier_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStreamidentifierExist("citrixadc_streamidentifier.tf_streamidentifier", nil),
					resource.TestCheckResourceAttr("citrixadc_streamidentifier.tf_streamidentifier", "name", "my_streamidentifier"),
					resource.TestCheckResourceAttr("citrixadc_streamidentifier.tf_streamidentifier", "selectorname", "my_streamselector"),
					resource.TestCheckResourceAttr("citrixadc_streamidentifier.tf_streamidentifier", "samplecount", "10"),
					resource.TestCheckResourceAttr("citrixadc_streamidentifier.tf_streamidentifier", "sort", "CONNECTIONS"),
					resource.TestCheckResourceAttr("citrixadc_streamidentifier.tf_streamidentifier", "snmptrap", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_streamidentifier.tf_streamidentifier", "loglimit", "500"),
					resource.TestCheckResourceAttr("citrixadc_streamidentifier.tf_streamidentifier", "loginterval", "60"),
					resource.TestCheckResourceAttr("citrixadc_streamidentifier.tf_streamidentifier", "log", "NONE"),
				),
			},
			{
				Config: testAccStreamidentifier_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckStreamidentifierExist("citrixadc_streamidentifier.tf_streamidentifier", nil),
					resource.TestCheckResourceAttr("citrixadc_streamidentifier.tf_streamidentifier", "name", "my_streamidentifier"),
					resource.TestCheckResourceAttr("citrixadc_streamidentifier.tf_streamidentifier", "selectorname", "my_streamselector"),
					resource.TestCheckResourceAttr("citrixadc_streamidentifier.tf_streamidentifier", "samplecount", "20"),
					resource.TestCheckResourceAttr("citrixadc_streamidentifier.tf_streamidentifier", "sort", "REQUESTS"),
					resource.TestCheckResourceAttr("citrixadc_streamidentifier.tf_streamidentifier", "snmptrap", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_streamidentifier.tf_streamidentifier", "loglimit", "600"),
					resource.TestCheckResourceAttr("citrixadc_streamidentifier.tf_streamidentifier", "loginterval", "120"),
					resource.TestCheckResourceAttr("citrixadc_streamidentifier.tf_streamidentifier", "log", "SYSLOG"),
				),
			},
		},
	})
}

func testAccCheckStreamidentifierExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No streamidentifier name is set")
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
		data, err := client.FindResource(service.Streamidentifier.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("streamidentifier %s not found", n)
		}

		return nil
	}
}

func testAccCheckStreamidentifierDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_streamidentifier" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Streamidentifier.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("streamidentifier %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccStreamidentifierDataSource_basic = `
	resource "citrixadc_streamselector" "tf_streamselector" {
		name = "my_streamselector"
		rule = ["HTTP.REQ.URL", "CLIENT.IP.SRC"]
	}
	resource "citrixadc_streamidentifier" "tf_streamidentifier" {
		name         = "my_streamidentifier"
		selectorname = citrixadc_streamselector.tf_streamselector.name
		samplecount  = 10
		sort         = "CONNECTIONS"
		snmptrap     = "ENABLED"
		loglimit	 = 500
		loginterval = 60
		log = "NONE"
	}

	data "citrixadc_streamidentifier" "tf_streamidentifier_datasource" {
		name = citrixadc_streamidentifier.tf_streamidentifier.name
	}
`

func TestAccStreamidentifierDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccStreamidentifierDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_streamidentifier.tf_streamidentifier_datasource", "name", "my_streamidentifier"),
					resource.TestCheckResourceAttr("data.citrixadc_streamidentifier.tf_streamidentifier_datasource", "selectorname", "my_streamselector"),
					resource.TestCheckResourceAttr("data.citrixadc_streamidentifier.tf_streamidentifier_datasource", "samplecount", "10"),
					resource.TestCheckResourceAttr("data.citrixadc_streamidentifier.tf_streamidentifier_datasource", "sort", "CONNECTIONS"),
					resource.TestCheckResourceAttr("data.citrixadc_streamidentifier.tf_streamidentifier_datasource", "snmptrap", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_streamidentifier.tf_streamidentifier_datasource", "loglimit", "500"),
					resource.TestCheckResourceAttr("data.citrixadc_streamidentifier.tf_streamidentifier_datasource", "loginterval", "60"),
					resource.TestCheckResourceAttr("data.citrixadc_streamidentifier.tf_streamidentifier_datasource", "log", "NONE"),
					resource.TestCheckResourceAttrSet("data.citrixadc_streamidentifier.tf_streamidentifier_datasource", "id"),
				),
			},
		},
	})
}
