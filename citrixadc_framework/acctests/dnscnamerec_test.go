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
	"net/url"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testAccDnscnamerec_basic = `


resource "citrixadc_dnscnamerec" "dnscnamerec" {
	aliasname = "citrixadc.cloud.com"
    canonicalname = "ctxwsp-citrixadc-fdproxy-global.trafficmanager.net"
    ttl = 3600
}
`

const testAccDnscnamerecDataSource_basic = `

resource "citrixadc_dnscnamerec" "dnscnamerec" {
	aliasname = "tfacc-ds-cname-test.local"
    canonicalname = "tfacc-target.example.com"
    ttl = 3600
}

data "citrixadc_dnscnamerec" "dnscnamerec" {
	aliasname = citrixadc_dnscnamerec.dnscnamerec.aliasname
}
`

func TestAccDnscnamerec_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnscnamerecDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnscnamerec_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnscnamerecExist("citrixadc_dnscnamerec.dnscnamerec", nil),
					resource.TestCheckResourceAttr("citrixadc_dnscnamerec.dnscnamerec", "aliasname", "citrixadc.cloud.com"),
					resource.TestCheckResourceAttr("citrixadc_dnscnamerec.dnscnamerec", "canonicalname", "ctxwsp-citrixadc-fdproxy-global.trafficmanager.net"),
					resource.TestCheckResourceAttr("citrixadc_dnscnamerec.dnscnamerec", "ttl", "3600"),
				),
			},
		},
	})
}

func testAccCheckDnscnamerecExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnscnamerec name is set")
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
		data, err := client.FindResource(service.Dnscnamerec.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("dnscnamerec %s not found", n)
		}

		return nil
	}
}

func testAccCheckDnscnamerecDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnscnamerec" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		argsMap := make(map[string]string)
		argsMap["ecssubnet"] = url.QueryEscape(rs.Primary.Attributes["ecssubnet"])
		findParams := service.FindParams{
			ResourceType: service.Dnscnamerec.Type(),
			ArgsMap:      argsMap,
		}
		_, err := client.FindResourceArrayWithParams(findParams)

		if err == nil {
			return fmt.Errorf("dnscnamerec %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccDnscnamerec_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_dnscnamerec.dnscnamerec"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnscnamerecDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnscnamerec_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnscnamerecExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Dnscnamerec.Type(), "citrixadc.cloud.com"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccDnscnamerec_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckDnscnamerecExist(resAddr, nil)),
			},
		},
	})
}

func TestAccDnscnamerec_import(t *testing.T) {
	const resAddr = "citrixadc_dnscnamerec.dnscnamerec"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnscnamerecDestroy,
		Steps: []resource.TestStep{
			{Config: testAccDnscnamerec_basic},
			{
				Config:                  testAccDnscnamerec_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccDnscnamerec_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDnscnamerecDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccDnscnamerec_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnscnamerecExist("citrixadc_dnscnamerec.dnscnamerec", nil),
				),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccDnscnamerec_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnscnamerecExist("citrixadc_dnscnamerec.dnscnamerec", nil),
				),
			},
		},
	})
}

func TestAccDnscnamerecDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDnscnamerecDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_dnscnamerec.dnscnamerec", "aliasname", "tfacc-ds-cname-test.local"),
					resource.TestCheckResourceAttr("data.citrixadc_dnscnamerec.dnscnamerec", "canonicalname", "tfacc-target.example.com"),
					resource.TestCheckResourceAttr("data.citrixadc_dnscnamerec.dnscnamerec", "ttl", "3600"),
				),
			},
		},
	})
}
