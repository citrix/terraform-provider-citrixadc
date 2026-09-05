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
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

const testAccCloudroutes_basic = `


resource "citrixadc_cloudroutes" "tf_cloudroutes" {
	name             = "tf_cloudroutes"
	routesvpcnetwork = "client_vpc"
	vipsubnet        = "192.168.10.0/24"
	vipvpcnetwork    = "vip_vpc"
	clientipaddress  = "192.168.10.5"
}

`
const testAccCloudroutes_update = `


resource "citrixadc_cloudroutes" "tf_cloudroutes" {
	name             = "tf_cloudroutes"
	routesvpcnetwork = "client_vpc"
	vipsubnet        = "192.168.20.0/24"
	vipvpcnetwork    = "vip_vpc2"
	clientipaddress  = "192.168.20.5"
}

`

func TestAccCloudroutes_basic(t *testing.T) {
	// cloudroutes is a cloud-feature resource: the referenced VPC networks must
	// exist on a cloud-provisioned ADC. Guard live runs behind this note.
	t.Skip("needs fixture: cloudroutes requires a cloud-provisioned ADC with the referenced VPC networks (routesvpcnetwork/vipvpcnetwork). Adjust values or remove skip to run against such an environment.")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudroutesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudroutes_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudroutesExist("citrixadc_cloudroutes.tf_cloudroutes", nil),
					resource.TestCheckResourceAttr("citrixadc_cloudroutes.tf_cloudroutes", "name", "tf_cloudroutes"),
					resource.TestCheckResourceAttr("citrixadc_cloudroutes.tf_cloudroutes", "vipsubnet", "192.168.10.0/24"),
					resource.TestCheckResourceAttr("citrixadc_cloudroutes.tf_cloudroutes", "clientipaddress", "192.168.10.5"),
				),
			},
			{
				Config: testAccCloudroutes_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudroutesExist("citrixadc_cloudroutes.tf_cloudroutes", nil),
					resource.TestCheckResourceAttr("citrixadc_cloudroutes.tf_cloudroutes", "name", "tf_cloudroutes"),
					resource.TestCheckResourceAttr("citrixadc_cloudroutes.tf_cloudroutes", "vipsubnet", "192.168.20.0/24"),
					resource.TestCheckResourceAttr("citrixadc_cloudroutes.tf_cloudroutes", "clientipaddress", "192.168.20.5"),
				),
			},
		},
	})
}

func TestAccCloudroutes_import(t *testing.T) {
	const resAddr = "citrixadc_cloudroutes.tf_cloudroutes"
	t.Skip("needs fixture: cloudroutes requires a cloud-provisioned ADC with the referenced VPC networks. Remove skip to run against such an environment.")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudroutesDestroy,
		Steps: []resource.TestStep{
			{Config: testAccCloudroutes_basic},
			{
				Config:            testAccCloudroutes_basic,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckCloudroutesExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cloudroutes name is set")
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
		data, err := client.FindResource("cloudroutes", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("cloudroutes %s not found", n)
		}

		return nil
	}
}

func testAccCheckCloudroutesDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_cloudroutes" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("cloudroutes", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("cloudroutes %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccCloudroutesDataSource_basic(t *testing.T) {
	t.Skip("needs fixture: cloudroutes requires a cloud-provisioned ADC with the referenced VPC networks. Remove skip to run against such an environment.")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudroutesDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_cloudroutes.test", "name", "tf_cloudroutes"),
					resource.TestCheckResourceAttrSet("data.citrixadc_cloudroutes.test", "id"),
					resource.TestCheckResourceAttrSet("data.citrixadc_cloudroutes.test", "vipsubnet"),
				),
			},
		},
	})
}

const testAccCloudroutesDataSource_basic = `
resource "citrixadc_cloudroutes" "tf_cloudroutes" {
	name             = "tf_cloudroutes"
	routesvpcnetwork = "client_vpc"
	vipsubnet        = "192.168.10.0/24"
	vipvpcnetwork    = "vip_vpc"
	clientipaddress  = "192.168.10.5"
}

data "citrixadc_cloudroutes" "test" {
	name = citrixadc_cloudroutes.tf_cloudroutes.name
}
`

// Step 1: every unset-eligible attribute set to a valid non-default value.
const testAccCloudroutes_unset_step1 = `
resource "citrixadc_cloudroutes" "tf_unset" {
	name             = "tf_test_cloudroutes_unset"
	routesvpcnetwork = "client_vpc"
	vipsubnet        = "192.168.10.0/24"
	vipvpcnetwork    = "vip_vpc"
	clientipaddress  = "192.168.10.5"
}
`

// Step 2: unset-eligible attributes removed from config -> provider must unset them.
const testAccCloudroutes_unset_step2 = `
resource "citrixadc_cloudroutes" "tf_unset" {
	name             = "tf_test_cloudroutes_unset"
	routesvpcnetwork = "client_vpc"
}
`

func TestAccCloudroutes_unset(t *testing.T) {
	t.Skip("needs fixture: cloudroutes requires a cloud-provisioned ADC with the referenced VPC networks. Remove skip to run against such an environment.")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudroutesDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values apply and persist.
				Config: testAccCloudroutes_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudroutesExist("citrixadc_cloudroutes.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_cloudroutes.tf_unset", "vipsubnet", "192.168.10.0/24"),
					resource.TestCheckResourceAttr("citrixadc_cloudroutes.tf_unset", "vipvpcnetwork", "vip_vpc"),
					resource.TestCheckResourceAttr("citrixadc_cloudroutes.tf_unset", "clientipaddress", "192.168.10.5"),
				),
			},
			{
				// Removing them must unset -> appliance reverts them.
				Config: testAccCloudroutes_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudroutesExist("citrixadc_cloudroutes.tf_unset", nil),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckCloudroutesADCValue("tf_test_cloudroutes_unset", "vipsubnet", ""),
					testAccCheckCloudroutesADCValue("tf_test_cloudroutes_unset", "vipvpcnetwork", ""),
					testAccCheckCloudroutesADCValue("tf_test_cloudroutes_unset", "clientipaddress", ""),
				),
			},
		},
	})
}

// testAccCheckCloudroutesADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckCloudroutesADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Cloudroutes.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("cloudroutes %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("cloudroutes %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

// TestAccCloudroutes_selfHealing verifies the provider re-creates the route when
// it is deleted out-of-band between apply steps (drift recovery).
func TestAccCloudroutes_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_cloudroutes.tf_cloudroutes"
	t.Skip("needs fixture: cloudroutes requires a cloud-provisioned ADC with the referenced VPC networks. Remove skip to run against such an environment.")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudroutesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudroutes_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCloudroutesExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Cloudroutes.Type(), "tf_cloudroutes"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccCloudroutes_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCloudroutesExist(resAddr, nil)),
			},
		},
	})
}
