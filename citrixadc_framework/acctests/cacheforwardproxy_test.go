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
	"strconv"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCacheforwardproxyDestroy,
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

func TestAccCacheforwardproxy_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_cacheforwardproxy.tf_cacheforwardproxy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCacheforwardproxyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCacheforwardproxy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCacheforwardproxyExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgs(service.Cacheforwardproxy.Type(), "10.222.74.185", []string{"port:5000"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccCacheforwardproxy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCacheforwardproxyExist(resAddr, nil)),
			},
		},
	})
}

func TestAccCacheforwardproxy_import(t *testing.T) {
	const resAddr = "citrixadc_cacheforwardproxy.tf_cacheforwardproxy"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCacheforwardproxyDestroy,
		Steps: []resource.TestStep{
			{Config: testAccCacheforwardproxy_basic},
			{
				Config:                  testAccCacheforwardproxy_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
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
		client, err := testAccGetFrameworkClient()
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
	client, err := testAccGetFrameworkClient()
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

func TestAccCacheforwardproxy_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckCacheforwardproxyDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccCacheforwardproxy_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCacheforwardproxyExist("citrixadc_cacheforwardproxy.tf_cacheforwardproxy", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccCacheforwardproxy_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckCacheforwardproxyExist("citrixadc_cacheforwardproxy.tf_cacheforwardproxy", nil)),
			},
		},
	})
}

func TestAccCacheforwardproxyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccCacheforwardproxyDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_cacheforwardproxy.tf_cacheforwardproxy_ds", "ipaddress", "10.222.74.187"),
					resource.TestCheckResourceAttr("data.citrixadc_cacheforwardproxy.tf_cacheforwardproxy_ds", "port", "6000"),
				),
			},
		},
	})
}

const testAccCacheforwardproxyDataSource_basic = `

resource "citrixadc_cacheforwardproxy" "tf_cacheforwardproxy_ds" {
    ipaddress  = "10.222.74.187"
    port       = 6000
}

data "citrixadc_cacheforwardproxy" "tf_cacheforwardproxy_ds" {
    ipaddress = citrixadc_cacheforwardproxy.tf_cacheforwardproxy_ds.ipaddress
    depends_on = [citrixadc_cacheforwardproxy.tf_cacheforwardproxy_ds]
}

`
