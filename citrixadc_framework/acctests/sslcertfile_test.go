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
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testAccSslcertfile_basic = `
	resource "citrixadc_sslcertfile" "tf_sslcertfile" {
		name = "tf_sslcertfile"
		src = "local://certificate1.crt"
	}
`

func TestAccSslcertfile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslcertfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcertfile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslcertfileExist("citrixadc_sslcertfile.tf_sslcertfile", nil),
				),
			},
		},
	})
}

func testAccCheckSslcertfileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslcertfile name is set")
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
		dataArr, err := client.FindAllResources(service.Sslcertfile.Type())

		if err != nil {
			return err
		}
		found := false
		for _, v := range dataArr {
			if v["name"].(string) == rs.Primary.Attributes["name"] {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("sslcertfile %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslcertfileDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslcertfile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		dataArr, err := client.FindAllResources(service.Sslcertfile.Type())

		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if v["name"].(string) == rs.Primary.Attributes["name"] {
				found = true
				break
			}
		}
		if found {
			return fmt.Errorf("sslcertfile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccSslcertfile_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_sslcertfile.tf_sslcertfile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslcertfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcertfile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSslcertfileExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResourceWithArgs(service.Sslcertfile.Type(), "", []string{"name:tf_sslcertfile"}); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccSslcertfile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSslcertfileExist(resAddr, nil)),
			},
		},
	})
}

func TestAccSslcertfile_import(t *testing.T) {
	const resAddr = "citrixadc_sslcertfile.tf_sslcertfile"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslcertfileDestroy,
		Steps: []resource.TestStep{
			{Config: testAccSslcertfile_basic},
			{
				Config:                  testAccSslcertfile_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"src"},
			},
		},
	})
}

func TestAccSslcertfile_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckSslcertfileDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccSslcertfile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSslcertfileExist("citrixadc_sslcertfile.tf_sslcertfile", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccSslcertfile_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSslcertfileExist("citrixadc_sslcertfile.tf_sslcertfile", nil)),
			},
		},
	})
}

func TestAccSslcertfileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslcertfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslcertfileDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslcertfile.tf_sslcertfile_ds", "name", "tf_sslcertfile_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_sslcertfile.tf_sslcertfile_ds", "src", "certificate1.crt"),
				),
			},
		},
	})
}

const testAccSslcertfileDataSource_basic = `

resource "citrixadc_sslcertfile" "tf_sslcertfile_ds" {
	name = "tf_sslcertfile_ds"
	src = "local://certificate1.crt"
}

data "citrixadc_sslcertfile" "tf_sslcertfile_ds" {
	name = citrixadc_sslcertfile.tf_sslcertfile_ds.name
	depends_on = [citrixadc_sslcertfile.tf_sslcertfile_ds]
}
`
