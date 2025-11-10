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
	"strconv"
	"testing"
)

const testAccCacheforwardproxy_basic = `
	resource "citrixadc_cacheforwardproxy" "tf_cacheforwardproxy" {
		ipaddress  = "10.222.74.185"
		port        = 5000
	}
`
const testAccCacheforwardproxy_update = `
	resource "citrixadc_cacheforwardproxy" "tf_cacheforwardproxy" {
		ipaddress  = "10.222.74.186"
		port        = 5500
	}
`

func TestAccCacheforwardproxy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckCacheforwardproxyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCacheforwardproxy_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCacheforwardproxyExist("citrixadc_cacheforwardproxy.tf_cacheforwardproxy", nil),
					resource.TestCheckResourceAttr("citrixadc_cacheforwardproxy.tf_cacheforwardproxy", "ipaddress", "10.222.74.185"),
					resource.TestCheckResourceAttr("citrixadc_cacheforwardproxy.tf_cacheforwardproxy", "port", "5000"),
				),
			},
			{
				Config: testAccCacheforwardproxy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCacheforwardproxyExist("citrixadc_cacheforwardproxy.tf_cacheforwardproxy", nil),
					resource.TestCheckResourceAttr("citrixadc_cacheforwardproxy.tf_cacheforwardproxy", "ipaddress", "10.222.74.186"),
					resource.TestCheckResourceAttr("citrixadc_cacheforwardproxy.tf_cacheforwardproxy", "port", "5500"),
				),
			},
		},
	})
}

func testAccCheckCacheforwardproxyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cacheforwardproxy name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Use the shared utility function to get a configured client
		client, err := testAccGetClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		dataArr, err := client.FindAllResources(service.Cacheforwardproxy.Type())

		if err != nil {
			return err
		}
		found := false
		for _, v := range dataArr {
			if v["ipaddress"].(string) == rs.Primary.Attributes["ipaddress"] &&
				strconv.Itoa(int(v["port"].(float64))) == rs.Primary.Attributes["port"] {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("cacheforwardproxy %s not found", n)
		}

		return nil
	}
}

func testAccCheckCacheforwardproxyDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_cacheforwardproxy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		dataArr, err := client.FindAllResources(service.Cacheforwardproxy.Type())

		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if v["ipaddress"].(string) == rs.Primary.Attributes["ipaddress"] &&
				strconv.Itoa(int(v["port"].(float64))) == rs.Primary.Attributes["port"] {
				found = true
				break
			}
		}
		if found {
			return fmt.Errorf("cacheforwardproxy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
