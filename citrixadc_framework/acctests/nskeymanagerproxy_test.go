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

// nskeymanagerproxy is a create+delete (no-update) resource. All schema attributes
// are RequiresReplace, so the basic test only creates and verifies. ValidateConfig
// enforces that exactly one of serverip / servername is set; this test uses serverip,
// which is also the resource ID and the GET/DELETE path key.
const testAccNskeymanagerproxy_basic_step1 = `
resource "citrixadc_nskeymanagerproxy" "tf_nskeymanagerproxy" {
  serverip = "192.0.2.50"
  port     = 1443
}

`

func TestAccNskeymanagerproxy_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNskeymanagerproxyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNskeymanagerproxy_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNskeymanagerproxyExist("citrixadc_nskeymanagerproxy.tf_nskeymanagerproxy", nil),
					resource.TestCheckResourceAttr("citrixadc_nskeymanagerproxy.tf_nskeymanagerproxy", "serverip", "192.0.2.50"),
					resource.TestCheckResourceAttr("citrixadc_nskeymanagerproxy.tf_nskeymanagerproxy", "port", "1443"),
				),
			},
		},
	})
}

func TestAccNskeymanagerproxy_import(t *testing.T) {
	t.Skip("TODO: Requires review")
	const resAddr = "citrixadc_nskeymanagerproxy.tf_nskeymanagerproxy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNskeymanagerproxyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNskeymanagerproxy_basic_step1,
			},
			{
				Config:                  testAccNskeymanagerproxy_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckNskeymanagerproxyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nskeymanagerproxy name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Nskeymanagerproxy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nskeymanagerproxy %s not found", n)
		}

		return nil
	}
}

func testAccCheckNskeymanagerproxyDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nskeymanagerproxy" {
			continue
		}
		_, err := client.FindResource(service.Nskeymanagerproxy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nskeymanagerproxy %s still exists", rs.Primary.ID)
		}
	}
	return nil
}

const testAccNskeymanagerproxyDataSource_basic = `

resource "citrixadc_nskeymanagerproxy" "tf_nskeymanagerproxy" {
  serverip = "192.0.2.50"
  port     = 1443
}

data "citrixadc_nskeymanagerproxy" "tf_nskeymanagerproxy" {
  serverip   = citrixadc_nskeymanagerproxy.tf_nskeymanagerproxy.serverip
  depends_on = [citrixadc_nskeymanagerproxy.tf_nskeymanagerproxy]
}
`

func TestAccNskeymanagerproxyDataSource_basic(t *testing.T) {
	t.Skip("TODO: Requires review")
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNskeymanagerproxyDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nskeymanagerproxy.tf_nskeymanagerproxy", "serverip", "192.0.2.50"),
					resource.TestCheckResourceAttr("data.citrixadc_nskeymanagerproxy.tf_nskeymanagerproxy", "port", "1443"),
				),
			},
		},
	})
}
