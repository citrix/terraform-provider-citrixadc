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

const testAccChannel_basic = `


	resource "citrixadc_channel" "tf_channel" {
		channel_id = "LA/3"
		tagall     = "ON"
		speed      = "1000"
	}
`
const testAccChannel_update = `


	resource "citrixadc_channel" "tf_channel" {
		channel_id = "LA/3"
		tagall     = "OFF"
		speed      = "100"
	}
`

func TestAccChannel_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckChannelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccChannel_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckChannelExist("citrixadc_channel.tf_channel", nil),
					resource.TestCheckResourceAttr("citrixadc_channel.tf_channel", "channel_id", "LA/3"),
					resource.TestCheckResourceAttr("citrixadc_channel.tf_channel", "tagall", "ON"),
					resource.TestCheckResourceAttr("citrixadc_channel.tf_channel", "speed", "1000"),
				),
			},
			{
				Config: testAccChannel_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckChannelExist("citrixadc_channel.tf_channel", nil),
					resource.TestCheckResourceAttr("citrixadc_channel.tf_channel", "channel_id", "LA/3"),
					resource.TestCheckResourceAttr("citrixadc_channel.tf_channel", "tagall", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_channel.tf_channel", "speed", "100"),
				),
			},
		},
	})
}

func testAccCheckChannelExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No channel name is set")
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
		data, err := client.FindResource(service.Channel.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("channel %s not found", n)
		}

		return nil
	}
}

func testAccCheckChannelDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_channel" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Channel.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("channel %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccChannelDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccChannelDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_channel.tf_channel_ds", "channelid", "LA/3"),
					resource.TestCheckResourceAttr("data.citrixadc_channel.tf_channel_ds", "tagall", "ON"),
				),
			},
		},
	})
}

const testAccChannelDataSource_basic = `

resource "citrixadc_channel" "tf_channel_ds" {
	channel_id = "LA/3"
	tagall     = "ON"
}

data "citrixadc_channel" "tf_channel_ds" {
	channelid = citrixadc_channel.tf_channel_ds.channel_id
	depends_on = [citrixadc_channel.tf_channel_ds]
}

`
