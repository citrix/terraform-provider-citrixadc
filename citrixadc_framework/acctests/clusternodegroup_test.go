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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccClusternodegroup_basic = `

resource "citrixadc_clusternodegroup" "tf_clusternodegroup" {
	name   = "my_clusternode"
	strict = "YES"
	}
`
const testAccClusternodegroup_update = `

resource "citrixadc_clusternodegroup" "tf_clusternodegroup" {
	name   = "my_clusternode"
	strict = "NO"
	}
`

const testAccClusternodegroupDataSource_basic = `

resource "citrixadc_clusternodegroup" "tf_clusternodegroup" {
	name   = "my_clusternode_ds"
	strict = "YES"
}

data "citrixadc_clusternodegroup" "tf_clusternodegroup_ds" {
	name = citrixadc_clusternodegroup.tf_clusternodegroup.name
}
`

func TestAccClusternodegroup_basic(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckClusternodegroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClusternodegroup_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroupExist("citrixadc_clusternodegroup.tf_clusternodegroup", nil),
					resource.TestCheckResourceAttr("citrixadc_clusternodegroup.tf_clusternodegroup", "name", "my_clusternode"),
					resource.TestCheckResourceAttr("citrixadc_clusternodegroup.tf_clusternodegroup", "strict", "YES"),
				),
			},
			{
				Config: testAccClusternodegroup_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroupExist("citrixadc_clusternodegroup.tf_clusternodegroup", nil),
					resource.TestCheckResourceAttr("citrixadc_clusternodegroup.tf_clusternodegroup", "name", "my_clusternode"),
					resource.TestCheckResourceAttr("citrixadc_clusternodegroup.tf_clusternodegroup", "strict", "NO"),
				),
			},
		},
	})
}

func testAccCheckClusternodegroupExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No clusternodegroup name is set")
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
		data, err := client.FindResource(service.Clusternodegroup.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("clusternodegroup %s not found", n)
		}

		return nil
	}
}

func testAccCheckClusternodegroupDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_clusternodegroup" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Clusternodegroup.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("clusternodegroup %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccClusternodegroup_unset_step1 = `

resource "citrixadc_clusternodegroup" "tf_unset" {
	name   = "tf_unset_clusternode"
	strict = "YES"
}
`

const testAccClusternodegroup_unset_step2 = `

resource "citrixadc_clusternodegroup" "tf_unset" {
	name = "tf_unset_clusternode"
	# strict removed from config -> provider must unset it (revert to NITRO default "NO").
}
`

func TestAccClusternodegroup_unset(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckClusternodegroupDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default value is applied and persisted.
				Config: testAccClusternodegroup_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroupExist("citrixadc_clusternodegroup.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_clusternodegroup.tf_unset", "strict", "YES"),
				),
			},
			{
				// Removing strict must unset it: state reverts to the documented
				// NITRO default ("NO") and the implicit post-apply plan is empty.
				Config: testAccClusternodegroup_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroupExist("citrixadc_clusternodegroup.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_clusternodegroup.tf_unset", "strict", "NO"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckClusternodegroupADCValue("tf_unset_clusternode", "strict", "NO"),
				),
			},
		},
	})
}

// testAccCheckClusternodegroupADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckClusternodegroupADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Clusternodegroup.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("clusternodegroup %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("clusternodegroup %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccClusternodegroup_selfHealing(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	const resAddr = "citrixadc_clusternodegroup.tf_clusternodegroup"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckClusternodegroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClusternodegroup_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckClusternodegroupExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Clusternodegroup.Type(), "my_clusternode"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccClusternodegroup_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckClusternodegroupExist(resAddr, nil)),
			},
		},
	})
}

func TestAccClusternodegroup_import(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	const resAddr = "citrixadc_clusternodegroup.tf_clusternodegroup"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckClusternodegroupDestroy,
		Steps: []resource.TestStep{
			{Config: testAccClusternodegroup_basic},
			{
				Config:                  testAccClusternodegroup_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccClusternodegroup_sdkv2StateUpgrade(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckClusternodegroupDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccClusternodegroup_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckClusternodegroupExist("citrixadc_clusternodegroup.tf_clusternodegroup", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccClusternodegroup_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckClusternodegroupExist("citrixadc_clusternodegroup.tf_clusternodegroup", nil)),
			},
		},
	})
}

func TestAccClusternodegroupDataSource_basic(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccClusternodegroupDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_clusternodegroup.tf_clusternodegroup_ds", "name", "my_clusternode_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_clusternodegroup.tf_clusternodegroup_ds", "strict", "YES"),
				),
			},
		},
	})
}
