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

func TestAccChannel_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_channel.tf_channel"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckChannelDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccChannel_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckChannelExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Channel.Type(), "LA/3"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccChannel_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckChannelExist(resAddr, nil)),
			},
		},
	})
}

func TestAccChannel_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckChannelDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccChannel_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckChannelExist("citrixadc_channel.tf_channel", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccChannel_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckChannelExist("citrixadc_channel.tf_channel", nil)),
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

func TestAccChannel_import(t *testing.T) {
	const resAddr = "citrixadc_channel.tf_channel"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckChannelDestroy,
		Steps: []resource.TestStep{
			{Config: testAccChannel_basic},
			{
				Config:                  testAccChannel_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"speed"},
			},
		},
	})
}

// The channel unset test covers the unset-eligible attributes that are echoed
// by NITRO GET and carry a documented server default: state (ENABLED),
// mtu (1500), tagall (OFF), hamonitor (ON), haheartbeat (ON). Each is given a
// non-default value in step1 and removed from config in step2; the provider
// must unset them so the appliance reverts to the NITRO defaults.
const testAccChannel_unset_step1 = `
	resource "citrixadc_channel" "tf_unset" {
		channel_id  = "LA/3"
		state       = "DISABLED"
		mtu         = 1600
		tagall      = "ON"
		hamonitor   = "OFF"
		haheartbeat = "OFF"
	}
`

const testAccChannel_unset_step2 = `
	resource "citrixadc_channel" "tf_unset" {
		channel_id = "LA/3"
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccChannel_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckChannelDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccChannel_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckChannelExist("citrixadc_channel.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_channel.tf_unset", "state", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_channel.tf_unset", "mtu", "1600"),
					resource.TestCheckResourceAttr("citrixadc_channel.tf_unset", "tagall", "ON"),
					resource.TestCheckResourceAttr("citrixadc_channel.tf_unset", "hamonitor", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_channel.tf_unset", "haheartbeat", "OFF"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccChannel_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckChannelExist("citrixadc_channel.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_channel.tf_unset", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_channel.tf_unset", "mtu", "1500"),
					resource.TestCheckResourceAttr("citrixadc_channel.tf_unset", "tagall", "OFF"),
					resource.TestCheckResourceAttr("citrixadc_channel.tf_unset", "hamonitor", "ON"),
					resource.TestCheckResourceAttr("citrixadc_channel.tf_unset", "haheartbeat", "ON"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckChannelADCValue("LA/3", "state", "ENABLED"),
					testAccCheckChannelADCValue("LA/3", "mtu", "1500"),
					testAccCheckChannelADCValue("LA/3", "tagall", "OFF"),
					testAccCheckChannelADCValue("LA/3", "hamonitor", "ON"),
					testAccCheckChannelADCValue("LA/3", "haheartbeat", "ON"),
				),
			},
		},
	})
}

// testAccCheckChannelADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckChannelADCValue(id, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Channel.Type(), id)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("channel %s not found on appliance", id)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("channel %s: appliance attr %q = %q, want %q (unset did not revert it)", id, attr, got, want)
		}
		return nil
	}
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
