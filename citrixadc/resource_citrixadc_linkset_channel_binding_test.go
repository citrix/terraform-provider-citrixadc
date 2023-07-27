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
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"net/url"
	"strings"
	"testing"
)

const testAccLinkset_channel_binding_basic = `

resource "citrixadc_linkset_channel_binding" "tf_linkset_channel_binding" {
	linkset_id = citrixadc_linkset.tf_linkset.linkset_id
	ifnum      = citrixadc_channel.tf_channel.channel_id
  }
  
  
  resource "citrixadc_linkset" "tf_linkset"{
	  linkset_id = "LS/3"
  }
  
  resource "citrixadc_channel" "tf_channel"{
	  channel_id = "0/LA/2"
  }
  
`

const testAccLinkset_channel_binding_basic_step2 = `
resource "citrixadc_linkset" "tf_linkset"{
	linkset_id = "LS/3"
}

resource "citrixadc_channel" "tf_channel"{
	channel_id = "0/LA/2"
}
`

func TestAccLinkset_channel_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLinkset_channel_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLinkset_channel_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLinkset_channel_bindingExist("citrixadc_linkset_channel_binding.tf_linkset_channel_binding", nil),
				),
			},
			{
				Config: testAccLinkset_channel_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLinkset_channel_bindingNotExist("citrixadc_linkset_channel_binding.tf_linkset_channel_binding", "LS/3,0/LA/2"),
				),
			},
		},
	})
}

func testAccCheckLinkset_channel_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No linkset_channel_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 2)

		id := idSlice[0]
		ifnum := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "linkset_channel_binding",
			ResourceName:             url.QueryEscape(url.QueryEscape(id)),
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching ifnum
		found := false
		for _, v := range dataArr {
			if v["ifnum"].(string) == ifnum {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("linkset_channel_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckLinkset_channel_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		id := idSlice[0]
		ifnum := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "linkset_channel_binding",
			ResourceName:             url.QueryEscape(url.QueryEscape(id)),
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching ifnum
		found := false
		for _, v := range dataArr {
			if v["ifnum"].(string) == ifnum {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("linkset_channel_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckLinkset_channel_bindingDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_linkset_channel_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Linkset_channel_binding.Type(), url.QueryEscape(url.QueryEscape(rs.Primary.ID)))
		if err == nil {
			return fmt.Errorf("linkset_channel_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
