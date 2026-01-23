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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

const testAccNsacl6_add = `

	resource "citrixadc_nstrafficdomain" "tf_trafficdomain" {
		td        = 2
		aliasname = "tf_trafficdomain"
		vmac      = "ENABLED"
	}
	resource "citrixadc_nsacl6" "tf_nsacl6" {
		acl6name   = "tf_nsacl6"
		acl6action = "ALLOW"
		td         = citrixadc_nstrafficdomain.tf_trafficdomain.td
		logstate   = "ENABLED"
		stateful   = "YES"
		ratelimit  = 120
		state      = "ENABLED"
		priority   = 20
		protocol   = "ICMPV6"
	}
`
const testAccNsacl6_update = `

	resource "citrixadc_nstrafficdomain" "tf_trafficdomain" {
		td        = 2
		aliasname = "tf_trafficdomain"
		vmac      = "ENABLED"
	}
	resource "citrixadc_nsacl6" "tf_nsacl6" {
		acl6name   = "tf_nsacl6"
		acl6action = "ALLOW"
		td         = citrixadc_nstrafficdomain.tf_trafficdomain.td
		logstate   = "ENABLED"
		stateful   = "NO"
		ratelimit  = 100
		state      = "DISABLED"
		priority   = 30
		protocol   = "TCP"
	}
`

func TestAccNsacl6_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsacl6Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsacl6_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsacl6Exist("citrixadc_nsacl6.tf_nsacl6", nil),
					resource.TestCheckResourceAttr("citrixadc_nsacl6.tf_nsacl6", "acl6name", "tf_nsacl6"),
					resource.TestCheckResourceAttr("citrixadc_nsacl6.tf_nsacl6", "stateful", "YES"),
					resource.TestCheckResourceAttr("citrixadc_nsacl6.tf_nsacl6", "ratelimit", "120"),
					resource.TestCheckResourceAttr("citrixadc_nsacl6.tf_nsacl6", "state", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_nsacl6.tf_nsacl6", "protocol", "ICMPV6"),
				),
			},
			{
				Config: testAccNsacl6_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsacl6Exist("citrixadc_nsacl6.tf_nsacl6", nil),
					resource.TestCheckResourceAttr("citrixadc_nsacl6.tf_nsacl6", "acl6name", "tf_nsacl6"),
					resource.TestCheckResourceAttr("citrixadc_nsacl6.tf_nsacl6", "stateful", "NO"),
					resource.TestCheckResourceAttr("citrixadc_nsacl6.tf_nsacl6", "ratelimit", "100"),
					resource.TestCheckResourceAttr("citrixadc_nsacl6.tf_nsacl6", "state", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_nsacl6.tf_nsacl6", "protocol", "TCP"),
				),
			},
		},
	})
}

func testAccCheckNsacl6Exist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsacl6 name is set")
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
		data, err := client.FindResource(service.Nsacl6.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nsacl6 %s not found", n)
		}

		return nil
	}
}

func testAccCheckNsacl6Destroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nsacl6" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Nsacl6.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nsacl6 %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
