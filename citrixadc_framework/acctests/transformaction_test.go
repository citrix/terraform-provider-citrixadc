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

const testAccTransformaction_basic_step1 = `
resource "citrixadc_transformprofile" "tf_trans_profile1" {
  name = "tf_trans_profile1"
}

resource "citrixadc_transformprofile" "tf_trans_profile2" {
  name = "tf_trans_profile2"
}

resource "citrixadc_transformaction" "tf_trans_action" {
  name = "tf_trans_action"
  profilename = citrixadc_transformprofile.tf_trans_profile1.name
  priority = 100
  requrlfrom = "http://m3.mydomain.com/(.*)"
  requrlinto = "https://exp-proxy-v1.api.mydomain.com/$1"
  resurlfrom = "https://exp-proxy-v1.api.mydomain.com/(.*)"
  resurlinto = "https://m3.mydomain.com/$1"
}
`

const testAccTransformaction_basic_step2 = `
resource "citrixadc_transformprofile" "tf_trans_profile1" {
  name = "tf_trans_profile1"
}

resource "citrixadc_transformprofile" "tf_trans_profile2" {
  name = "tf_trans_profile2"
}

resource "citrixadc_transformaction" "tf_trans_action" {
  name = "tf_trans_action"
  profilename = citrixadc_transformprofile.tf_trans_profile2.name
  priority = 100
  requrlfrom = "http://m4.mydomain.com/(.*)"
  requrlinto = "https://exp-proxy-v1.api.mydomain.com/$1"
  resurlfrom = "https://exp-proxy-v1.api.mydomain.com/(.*)"
  resurlinto = "https://m4.mydomain.com/$1"
}
`

const testAccTransformaction_basic_step3 = `
resource "citrixadc_transformprofile" "tf_trans_profile1" {
  name = "tf_trans_profile1"
}

resource "citrixadc_transformprofile" "tf_trans_profile2" {
  name = "tf_trans_profile2"
}

resource "citrixadc_transformaction" "tf_trans_action" {
  name = "tf_trans_action"
  profilename = citrixadc_transformprofile.tf_trans_profile2.name
  priority = 110
  requrlfrom = "http://m5.mydomain.com/(.*)"
  requrlinto = "https://exp-proxy-v1.api.mydomain.com/$1"
  resurlfrom = "https://exp-proxy-v1.api.mydomain.com/(.*)"
  resurlinto = "https://m5.mydomain.com/$1"
}
`

func TestAccTransformaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTransformactionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTransformaction_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTransformactionExist("citrixadc_transformaction.tf_trans_action", nil),
				),
			},
			{
				Config: testAccTransformaction_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTransformactionExist("citrixadc_transformaction.tf_trans_action", nil),
				),
			},
			{
				Config: testAccTransformaction_basic_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTransformactionExist("citrixadc_transformaction.tf_trans_action", nil),
				),
			},
		},
	})
}

func testAccCheckTransformactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No transformaction name is set")
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
		data, err := client.FindResource(service.Transformaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("transformaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckTransformactionDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_transformaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Transformaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("transformaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccTransformactionDataSource_basic = `
resource "citrixadc_transformprofile" "tf_trans_profile1" {
  name = "tf_trans_profile1"
}

resource "citrixadc_transformaction" "tf_trans_action" {
  name = "tf_trans_action"
  profilename = citrixadc_transformprofile.tf_trans_profile1.name
  priority = 100
  requrlfrom = "http://m3.mydomain.com/(.*)"
  requrlinto = "https://exp-proxy-v1.api.mydomain.com/$1"
  resurlfrom = "https://exp-proxy-v1.api.mydomain.com/(.*)"
  resurlinto = "https://m3.mydomain.com/$1"
}

data "citrixadc_transformaction" "tf_trans_action" {
    name = citrixadc_transformaction.tf_trans_action.name
}
`

func TestAccTransformactionDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTransformactionDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_transformaction.tf_trans_action", "name", "tf_trans_action"),
					resource.TestCheckResourceAttr("data.citrixadc_transformaction.tf_trans_action", "profilename", "tf_trans_profile1"),
					resource.TestCheckResourceAttr("data.citrixadc_transformaction.tf_trans_action", "priority", "100"),
					resource.TestCheckResourceAttr("data.citrixadc_transformaction.tf_trans_action", "requrlfrom", "http://m3.mydomain.com/(.*)"),
					resource.TestCheckResourceAttr("data.citrixadc_transformaction.tf_trans_action", "requrlinto", "https://exp-proxy-v1.api.mydomain.com/$1"),
					resource.TestCheckResourceAttr("data.citrixadc_transformaction.tf_trans_action", "resurlfrom", "https://exp-proxy-v1.api.mydomain.com/(.*)"),
					resource.TestCheckResourceAttr("data.citrixadc_transformaction.tf_trans_action", "resurlinto", "https://m3.mydomain.com/$1"),
				),
			},
		},
	})
}
