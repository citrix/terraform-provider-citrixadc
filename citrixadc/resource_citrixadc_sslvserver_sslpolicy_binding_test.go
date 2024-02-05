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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccSslvserver_sslpolicy_binding_lbvserver(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSslvserver_sslpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslvserver_sslpolicy_binding_lbvserver_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_sslpolicy_bindingExist("citrixadc_sslvserver_sslpolicy_binding.tf_binding_lb", nil),
				),
			},
			{
				Config: testAccSslvserver_sslpolicy_binding_lbvserver_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_sslpolicy_bindingExist("citrixadc_sslvserver_sslpolicy_binding.tf_binding_lb", nil),
				),
			},
		},
	})
}

func TestAccSslvserver_sslpolicy_binding_csvserver(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSslvserver_sslpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslvserver_sslpolicy_binding_csvserver_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_sslpolicy_bindingExist("citrixadc_sslvserver_sslpolicy_binding.tf_binding_cs", nil),
				),
			},
			{
				Config: testAccSslvserver_sslpolicy_binding_csvserver_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslvserver_sslpolicy_bindingExist("citrixadc_sslvserver_sslpolicy_binding.tf_binding_cs", nil),
				),
			},
		},
	})
}

func testAccCheckSslvserver_sslpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lb vserver name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		idSlice := strings.Split(rs.Primary.ID, ",")

		vservername := idSlice[0]
		policyname := idSlice[1]

		client := testAccProvider.Meta().(*NetScalerNitroClient).client
		findParams := service.FindParams{
			ResourceType:             "sslvserver_sslpolicy_binding",
			ResourceName:             vservername,
			ResourceMissingErrorCode: 461,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		if len(dataArr) == 0 {
			return fmt.Errorf("sslvserver sslpolicy binding zero array result")
		}

		foundIndex := -1
		for i, v := range dataArr {
			if v["policyname"].(string) == policyname {
				foundIndex = i
				break
			}
		}

		if foundIndex == -1 {
			return fmt.Errorf("sslvserver sslpolicy binding does not exist")
		}

		return nil
	}
}

func testAccCheckSslvserver_sslpolicy_bindingDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslvserver_sslpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		idSlice := strings.Split(rs.Primary.ID, ",")

		vservername := idSlice[0]

		findParams := service.FindParams{
			ResourceType:             "sslvserver_sslpolicy_binding",
			ResourceName:             vservername,
			ResourceMissingErrorCode: 461,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		if err != nil {
			return err
		}

		if len(dataArr) > 0 {
			return fmt.Errorf("sslvserver sslpolicy binding still exists")
		}

	}

	return nil
}

const testAccSslvserver_sslpolicy_binding_lbvserver_step1 = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_lbvserver"
  ipv46       = "192.168.34.21"
  port        = "443"
  servicetype = "SSL"
  sslprofile = "ns_default_ssl_profile_frontend"
}


resource "citrixadc_sslpolicy" "tf_sslpolicy" {
  name   = "tf_sslpolicy"
  rule   = "true"
  action = "NOOP"
}

resource "citrixadc_sslvserver_sslpolicy_binding" "tf_binding_lb" {
    vservername = citrixadc_lbvserver.tf_lbvserver.name
    policyname = citrixadc_sslpolicy.tf_sslpolicy.name
    priority = 333
    type = "REQUEST"
}
`

const testAccSslvserver_sslpolicy_binding_lbvserver_step2 = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_lbvserver"
  ipv46       = "192.168.34.21"
  port        = "443"
  servicetype = "SSL"
  sslprofile = "ns_default_ssl_profile_frontend"
}


resource "citrixadc_sslpolicy" "tf_sslpolicy" {
  name   = "tf_sslpolicy"
  rule   = "true"
  action = "NOOP"
}

resource "citrixadc_sslvserver_sslpolicy_binding" "tf_binding_lb" {
    vservername = citrixadc_lbvserver.tf_lbvserver.name
    policyname = citrixadc_sslpolicy.tf_sslpolicy.name
    priority = 444
    type = "REQUEST"
}
`

const testAccSslvserver_sslpolicy_binding_csvserver_step1 = `
resource "citrixadc_csvserver" "tf_csvserver" {

  name        = "tf_csvserver"
  ipv46       = "10.202.11.11"
  port        = 9090
  servicetype = "SSL"
  sslprofile = "ns_default_ssl_profile_frontend"
}

resource "citrixadc_sslpolicy" "tf_sslpolicy" {
  name   = "tf_sslpolicy"
  rule   = "true"
  action = "NOOP"
}

resource "citrixadc_sslvserver_sslpolicy_binding" "tf_binding_cs" {
    vservername = citrixadc_csvserver.tf_csvserver.name
    policyname = citrixadc_sslpolicy.tf_sslpolicy.name
    priority = 333
    type = "REQUEST"
}
`

const testAccSslvserver_sslpolicy_binding_csvserver_step2 = `
resource "citrixadc_csvserver" "tf_csvserver" {

  name        = "tf_csvserver"
  ipv46       = "10.202.11.11"
  port        = 9090
  servicetype = "SSL"
  sslprofile = "ns_default_ssl_profile_frontend"
}

resource "citrixadc_sslpolicy" "tf_sslpolicy" {
  name   = "tf_sslpolicy"
  rule   = "true"
  action = "NOOP"
}

resource "citrixadc_sslvserver_sslpolicy_binding" "tf_binding_cs" {
    vservername = citrixadc_csvserver.tf_csvserver.name
    policyname = citrixadc_sslpolicy.tf_sslpolicy.name
    priority = 444
    type = "REQUEST"
}
`
