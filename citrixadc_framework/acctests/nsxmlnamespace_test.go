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

const testAccNsxmlnamespace_add = `
	resource "citrixadc_nsxmlnamespace" "tf_nsxmlnamespace" {
		prefix      = "tf_nsxmlnamespace"
		namespace   = "http://www.w3.org/2001/04/xmlenc#"
		description = "Description"
	}
`
const testAccNsxmlnamespace_update = `
	resource "citrixadc_nsxmlnamespace" "tf_nsxmlnamespace" {
		prefix      = "tf_nsxmlnamespace"
		namespace   = "http://www.w3.org/2001/04/xmlenc#"
		description = "Description_sample"
	}
`

func TestAccNsxmlnamespace_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsxmlnamespaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsxmlnamespace_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsxmlnamespaceExist("citrixadc_nsxmlnamespace.tf_nsxmlnamespace", nil),
					resource.TestCheckResourceAttr("citrixadc_nsxmlnamespace.tf_nsxmlnamespace", "prefix", "tf_nsxmlnamespace"),
					resource.TestCheckResourceAttr("citrixadc_nsxmlnamespace.tf_nsxmlnamespace", "description", "Description"),
				),
			},
			{
				Config: testAccNsxmlnamespace_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsxmlnamespaceExist("citrixadc_nsxmlnamespace.tf_nsxmlnamespace", nil),
					resource.TestCheckResourceAttr("citrixadc_nsxmlnamespace.tf_nsxmlnamespace", "prefix", "tf_nsxmlnamespace"),
					resource.TestCheckResourceAttr("citrixadc_nsxmlnamespace.tf_nsxmlnamespace", "description", "Description_sample"),
				),
			},
		},
	})
}

func TestAccNsxmlnamespace_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_nsxmlnamespace.tf_nsxmlnamespace"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsxmlnamespaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsxmlnamespace_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsxmlnamespaceExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Nsxmlnamespace.Type(), "tf_nsxmlnamespace"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccNsxmlnamespace_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsxmlnamespaceExist(resAddr, nil)),
			},
		},
	})
}

func TestAccNsxmlnamespace_import(t *testing.T) {
	const resAddr = "citrixadc_nsxmlnamespace.tf_nsxmlnamespace"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsxmlnamespaceDestroy,
		Steps: []resource.TestStep{
			{Config: testAccNsxmlnamespace_add},
			{
				Config:                  testAccNsxmlnamespace_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckNsxmlnamespaceExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsxmlnamespace name is set")
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
		data, err := client.FindResource(service.Nsxmlnamespace.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nsxmlnamespace %s not found", n)
		}

		return nil
	}
}

func testAccCheckNsxmlnamespaceDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nsxmlnamespace" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Nsxmlnamespace.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nsxmlnamespace %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
func TestAccNsxmlnamespace_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckNsxmlnamespaceDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccNsxmlnamespace_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckNsxmlnamespaceExist("citrixadc_nsxmlnamespace.tf_nsxmlnamespace", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccNsxmlnamespace_add,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckNsxmlnamespaceExist("citrixadc_nsxmlnamespace.tf_nsxmlnamespace", nil)),
			},
		},
	})
}

// testAccCheckNsxmlnamespaceADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it. An empty want asserts the attribute is absent/empty on the box.
func testAccCheckNsxmlnamespaceADCValue(prefix, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Nsxmlnamespace.Type(), prefix)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("nsxmlnamespace %s not found on appliance", prefix)
		}
		got := ""
		if v, ok := data[attr]; ok && v != nil {
			got = fmt.Sprintf("%v", v)
		}
		if got != want {
			return fmt.Errorf("nsxmlnamespace %s: appliance attr %q = %q, want %q (unset did not revert it)", prefix, attr, got, want)
		}
		return nil
	}
}

const testAccNsxmlnamespace_unset_step1 = `
resource "citrixadc_nsxmlnamespace" "tf_unset" {
	prefix      = "tf_nsxmlnamespace_unset"
	namespace   = "http://www.w3.org/2001/04/xmlenc#"
	description = "unset_desc"
}
`

const testAccNsxmlnamespace_unset_step2 = `
resource "citrixadc_nsxmlnamespace" "tf_unset" {
	prefix    = "tf_nsxmlnamespace_unset"
	namespace = "http://www.w3.org/2001/04/xmlenc#"
	# description removed from config -> the provider must unset it (revert to
	# the NITRO default: no configured description).
}
`

func TestAccNsxmlnamespace_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNsxmlnamespaceDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccNsxmlnamespace_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsxmlnamespaceExist("citrixadc_nsxmlnamespace.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nsxmlnamespace.tf_unset", "description", "unset_desc"),
					testAccCheckNsxmlnamespaceADCValue("tf_nsxmlnamespace_unset", "description", "unset_desc"),
				),
			},
			{
				// Removing description must unset it: state (read back from the
				// appliance) reverts to the NITRO default, and the implicit
				// post-apply plan must be empty.
				Config: testAccNsxmlnamespace_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsxmlnamespaceExist("citrixadc_nsxmlnamespace.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_nsxmlnamespace.tf_unset", "description", ""),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckNsxmlnamespaceADCValue("tf_nsxmlnamespace_unset", "description", ""),
				),
			},
		},
	})
}

func TestAccNsxmlnamespaceDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNsxmlnamespaceDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_nsxmlnamespace.test", "prefix", "tf_nsxmlnamespace_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_nsxmlnamespace.test", "namespace", "http://www.w3.org/2001/04/xmlenc#"),
					resource.TestCheckResourceAttr("data.citrixadc_nsxmlnamespace.test", "description", "Datasource test"),
				),
			},
		},
	})
}

const testAccNsxmlnamespaceDataSource_basic = `
resource "citrixadc_nsxmlnamespace" "tf_nsxmlnamespace_ds" {
	prefix      = "tf_nsxmlnamespace_ds"
	namespace   = "http://www.w3.org/2001/04/xmlenc#"
	description = "Datasource test"
}

data "citrixadc_nsxmlnamespace" "test" {
	prefix = citrixadc_nsxmlnamespace.tf_nsxmlnamespace_ds.prefix
}
`
