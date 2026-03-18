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

const testAccPolicyhttpcallout_basic = `
resource "citrixadc_policyhttpcallout" "tf_policyhttpcallout" {
	name = "tf_policyhttpcallout"
	bodyexpr = "client.ip.src"
	cacheforsecs = 5
	comment = "Demo comment"
	headers = ["cip(client.ip.src)", "hdr(http.req.header(\"HDR\"))"]
	hostexpr = "http.req.header(\"Host\")"
	httpmethod = "GET"
	parameters = ["param1(\"name1\")", "param2(http.req.header(\"hdr\"))"]
	resultexpr = "http.res.body(10000).length"
	returntype = "TEXT"
	scheme = "http"
	vserver = citrixadc_lbvserver.tf_lbvserver.name
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
	name = "tf_lbvserver"
	ipv46 = "10.202.11.11"
	port = 80
	servicetype = "HTTP"
}
`

const testAccPolicyhttpcallout_basic_update_list_attributes = `
	resource "citrixadc_policyhttpcallout" "tf_policyhttpcallout" {
		name = "tf_policyhttpcallout"
		headers = ["Name(\"MyHeader\")"]
		parameters = ["param_update(\"name_update\")"]
	}

	resource "citrixadc_lbvserver" "tf_lbvserver" {
		name = "tf_lbvserver"
		ipv46 = "10.202.11.11"
		port = 80
		servicetype = "HTTP"
	}
`

const testAccPolicyhttpcallout_basic_update_other_attributes = `
	resource "citrixadc_policyhttpcallout" "tf_policyhttpcallout" {
		name = "tf_policyhttpcallout"
		fullreqexpr = "\"GET\" + http.req.url + \"HTTP/\" + http.req.version.major + \".\" + http.req.version.minor.sub(1) + \"r\\nHost:10.101.10.10\\r\\nAccept: */*\\r\\n\\r\\n\""
		cacheforsecs = 6
		comment = "Demo comment updated"
		resultexpr = "http.res.body(10001).length"
		returntype = "TEXT"
		scheme = "https"
		vserver = citrixadc_lbvserver.tf_lbvserver_update.name
	}

	resource "citrixadc_lbvserver" "tf_lbvserver_update" {
		name = "tf_lbvserver_update"
		ipv46 = "10.202.11.12"
		port = 80
		servicetype = "HTTP"
	}
`

const testAccPolicyhttpcallout_basic_update_return_type = `
	resource "citrixadc_policyhttpcallout" "tf_policyhttpcallout" {
		name = "tf_policyhttpcallout"
		returntype = "BOOL"
	}

	resource "citrixadc_lbvserver" "tf_lbvserver_update" {
		name = "tf_lbvserver_update"
		ipv46 = "10.202.11.12"
		port = 80
		servicetype = "HTTP"
	}
`

func TestAccPolicyhttpcallout_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckPolicyhttpcalloutDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicyhttpcallout_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyhttpcalloutExist("citrixadc_policyhttpcallout.tf_policyhttpcallout", nil),
					resource.TestCheckResourceAttr("citrixadc_policyhttpcallout.tf_policyhttpcallout", "name", "tf_policyhttpcallout"),
					resource.TestCheckResourceAttr("citrixadc_policyhttpcallout.tf_policyhttpcallout", "bodyexpr", "client.ip.src"),
					resource.TestCheckResourceAttr("citrixadc_policyhttpcallout.tf_policyhttpcallout", "cacheforsecs", "5"),
					resource.TestCheckResourceAttr("citrixadc_policyhttpcallout.tf_policyhttpcallout", "comment", "Demo comment"),
					resource.TestCheckResourceAttr("citrixadc_policyhttpcallout.tf_policyhttpcallout", "headers.0", "cip(client.ip.src)"),
					resource.TestCheckResourceAttr("citrixadc_policyhttpcallout.tf_policyhttpcallout", "headers.1", "hdr(http.req.header(\"HDR\"))"),
					resource.TestCheckResourceAttr("citrixadc_policyhttpcallout.tf_policyhttpcallout", "hostexpr", "http.req.header(\"Host\")"),
					resource.TestCheckResourceAttr("citrixadc_policyhttpcallout.tf_policyhttpcallout", "httpmethod", "GET"),
					resource.TestCheckResourceAttr("citrixadc_policyhttpcallout.tf_policyhttpcallout", "parameters.0", "param1(\"name1\")"),
					resource.TestCheckResourceAttr("citrixadc_policyhttpcallout.tf_policyhttpcallout", "parameters.1", "param2(http.req.header(\"hdr\"))"),
					resource.TestCheckResourceAttr("citrixadc_policyhttpcallout.tf_policyhttpcallout", "resultexpr", "http.res.body(10000).length"),
					resource.TestCheckResourceAttr("citrixadc_policyhttpcallout.tf_policyhttpcallout", "returntype", "TEXT"),
					resource.TestCheckResourceAttr("citrixadc_policyhttpcallout.tf_policyhttpcallout", "scheme", "http"),
					resource.TestCheckResourceAttr("citrixadc_policyhttpcallout.tf_policyhttpcallout", "vserver", "tf_lbvserver"),
				),
			},
			{
				Config: testAccPolicyhttpcallout_basic_update_list_attributes,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyhttpcalloutExist("citrixadc_policyhttpcallout.tf_policyhttpcallout", nil),
					resource.TestCheckResourceAttr("citrixadc_policyhttpcallout.tf_policyhttpcallout", "headers.0", "Name(\"MyHeader\")"),
					resource.TestCheckResourceAttr("citrixadc_policyhttpcallout.tf_policyhttpcallout", "parameters.0", "param_update(\"name_update\")"),
				),
			},
			{
				Config: testAccPolicyhttpcallout_basic_update_other_attributes,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyhttpcalloutExist("citrixadc_policyhttpcallout.tf_policyhttpcallout", nil),
					resource.TestCheckResourceAttr("citrixadc_policyhttpcallout.tf_policyhttpcallout", "fullreqexpr", "\"GET\" + http.req.url + \"HTTP/\" + http.req.version.major + \".\" + http.req.version.minor.sub(1) + \"r\\nHost:10.101.10.10\\r\\nAccept: */*\\r\\n\\r\\n\""),
					resource.TestCheckResourceAttr("citrixadc_policyhttpcallout.tf_policyhttpcallout", "bodyexpr", ""),
					resource.TestCheckResourceAttr("citrixadc_policyhttpcallout.tf_policyhttpcallout", "cacheforsecs", "6"),
					resource.TestCheckResourceAttr("citrixadc_policyhttpcallout.tf_policyhttpcallout", "comment", "Demo comment updated"),
					resource.TestCheckResourceAttr("citrixadc_policyhttpcallout.tf_policyhttpcallout", "hostexpr", ""),
					resource.TestCheckResourceAttr("citrixadc_policyhttpcallout.tf_policyhttpcallout", "resultexpr", "http.res.body(10001).length"),
					resource.TestCheckResourceAttr("citrixadc_policyhttpcallout.tf_policyhttpcallout", "returntype", "TEXT"),
					resource.TestCheckResourceAttr("citrixadc_policyhttpcallout.tf_policyhttpcallout", "scheme", "https"),
					resource.TestCheckResourceAttr("citrixadc_policyhttpcallout.tf_policyhttpcallout", "vserver", "tf_lbvserver_update"),
				),
			},
			{
				Config: testAccPolicyhttpcallout_basic_update_return_type,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPolicyhttpcalloutExist("citrixadc_policyhttpcallout.tf_policyhttpcallout", nil),
					resource.TestCheckResourceAttr("citrixadc_policyhttpcallout.tf_policyhttpcallout", "returntype", "BOOL"),
				),
			},
		},
	})
}

func testAccCheckPolicyhttpcalloutExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No policyhttpcallout name is set")
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
		data, err := client.FindResource(service.Policyhttpcallout.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("policyhttpcallout %s not found", n)
		}

		return nil
	}
}

func testAccCheckPolicyhttpcalloutDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_policyhttpcallout" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Policyhttpcallout.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("policyhttpcallout %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccPolicyhttpcalloutDataSource_basic = `
	resource "citrixadc_lbvserver" "tf_lbvserver_ds" {
		name        = "tf_lbvserver_ds"
		ipv46       = "10.202.11.13"
		port        = 80
		servicetype = "HTTP"
	}

	resource "citrixadc_policyhttpcallout" "tf_policyhttpcallout_ds" {
		name         = "tf_policyhttpcallout_ds"
		httpmethod   = "GET"
		scheme       = "http"
		vserver      = citrixadc_lbvserver.tf_lbvserver_ds.name
		returntype   = "TEXT"
		comment      = "Test datasource callout"
	}

	data "citrixadc_policyhttpcallout" "tf_policyhttpcallout_ds" {
		name = citrixadc_policyhttpcallout.tf_policyhttpcallout_ds.name
	}
`

func TestAccPolicyhttpcalloutDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicyhttpcalloutDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_policyhttpcallout.tf_policyhttpcallout_ds", "name", "tf_policyhttpcallout_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_policyhttpcallout.tf_policyhttpcallout_ds", "httpmethod", "GET"),
					resource.TestCheckResourceAttr("data.citrixadc_policyhttpcallout.tf_policyhttpcallout_ds", "scheme", "http"),
					resource.TestCheckResourceAttr("data.citrixadc_policyhttpcallout.tf_policyhttpcallout_ds", "returntype", "TEXT"),
					resource.TestCheckResourceAttr("data.citrixadc_policyhttpcallout.tf_policyhttpcallout_ds", "comment", "Test datasource callout"),
				),
			},
		},
	})
}
