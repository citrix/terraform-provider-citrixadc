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
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// cloudtrafficroutes is a cloud-deployment resource: it references a cloud VPC
// network (targetvpcnetwork) and cloud next-hop routing. It can only be created
// on an ADC running in a supported cloud environment (AWS/GCP/Azure) with a
// pre-existing VPC network. A standalone lab VPX rejects the configuration, so
// these tests skip unless that fixture is available.
const cloudtrafficroutesSkipReason = "needs cloud deployment fixture: an ADC in a supported cloud (AWS/GCP/Azure) with an existing targetvpcnetwork VPC network"

const testAccCloudtrafficroutes_basic = `
resource "citrixadc_cloudtrafficroutes" "tf_cloudtrafficroutes" {
	name             = "tf_cloudtrafficroutes"
	targetvpcnetwork = "tf_vpc_network"
	destrange        = "10.0.0.0/24"
	nexthopip        = "192.168.1.1"
	ownernode        = 0
}
`

const testAccCloudtrafficroutes_update = `
resource "citrixadc_cloudtrafficroutes" "tf_cloudtrafficroutes" {
	name             = "tf_cloudtrafficroutes"
	targetvpcnetwork = "tf_vpc_network"
	destrange        = "10.0.1.0/24"
	nexthopip        = "192.168.1.2"
	ownernode        = 1
}
`

func TestAccCloudtrafficroutes_basic(t *testing.T) {
	t.Skip(cloudtrafficroutesSkipReason)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudtrafficroutesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudtrafficroutes_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudtrafficroutesExist("citrixadc_cloudtrafficroutes.tf_cloudtrafficroutes", nil),
					resource.TestCheckResourceAttr("citrixadc_cloudtrafficroutes.tf_cloudtrafficroutes", "name", "tf_cloudtrafficroutes"),
					resource.TestCheckResourceAttr("citrixadc_cloudtrafficroutes.tf_cloudtrafficroutes", "targetvpcnetwork", "tf_vpc_network"),
					resource.TestCheckResourceAttr("citrixadc_cloudtrafficroutes.tf_cloudtrafficroutes", "destrange", "10.0.0.0/24"),
					resource.TestCheckResourceAttr("citrixadc_cloudtrafficroutes.tf_cloudtrafficroutes", "nexthopip", "192.168.1.1"),
				),
			},
			{
				Config: testAccCloudtrafficroutes_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudtrafficroutesExist("citrixadc_cloudtrafficroutes.tf_cloudtrafficroutes", nil),
					resource.TestCheckResourceAttr("citrixadc_cloudtrafficroutes.tf_cloudtrafficroutes", "name", "tf_cloudtrafficroutes"),
					resource.TestCheckResourceAttr("citrixadc_cloudtrafficroutes.tf_cloudtrafficroutes", "destrange", "10.0.1.0/24"),
					resource.TestCheckResourceAttr("citrixadc_cloudtrafficroutes.tf_cloudtrafficroutes", "nexthopip", "192.168.1.2"),
				),
			},
		},
	})
}

func TestAccCloudtrafficroutes_import(t *testing.T) {
	t.Skip(cloudtrafficroutesSkipReason)
	const resAddr = "citrixadc_cloudtrafficroutes.tf_cloudtrafficroutes"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudtrafficroutesDestroy,
		Steps: []resource.TestStep{
			{Config: testAccCloudtrafficroutes_basic},
			{
				Config:            testAccCloudtrafficroutes_basic,
				ResourceName:      resAddr,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// Step 1: every unset-eligible attribute set to a valid non-default value.
const testAccCloudtrafficroutes_unset_step1 = `
resource "citrixadc_cloudtrafficroutes" "tf_unset" {
	name             = "tf_cloudtrafficroutes_unset"
	targetvpcnetwork = "tf_vpc_network"
	destrange        = "10.0.0.0/24"
	nexthopip        = "192.168.1.1"
	ownernode        = 0
}
`

// Step 2: eligible attributes removed from config -> provider must unset them.
const testAccCloudtrafficroutes_unset_step2 = `
resource "citrixadc_cloudtrafficroutes" "tf_unset" {
	name = "tf_cloudtrafficroutes_unset"
}
`

func TestAccCloudtrafficroutes_unset(t *testing.T) {
	t.Skip(cloudtrafficroutesSkipReason)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudtrafficroutesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudtrafficroutes_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudtrafficroutesExist("citrixadc_cloudtrafficroutes.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_cloudtrafficroutes.tf_unset", "destrange", "10.0.0.0/24"),
					resource.TestCheckResourceAttr("citrixadc_cloudtrafficroutes.tf_unset", "nexthopip", "192.168.1.1"),
				),
			},
			{
				Config: testAccCloudtrafficroutes_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudtrafficroutesExist("citrixadc_cloudtrafficroutes.tf_unset", nil),
				),
			},
		},
	})
}

// TestAccCloudtrafficroutes_selfHealing verifies the provider re-creates the route
// when it is deleted out-of-band between apply steps (drift recovery).
func TestAccCloudtrafficroutes_selfHealing(t *testing.T) {
	t.Skip(cloudtrafficroutesSkipReason)
	const resAddr = "citrixadc_cloudtrafficroutes.tf_cloudtrafficroutes"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudtrafficroutesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudtrafficroutes_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCloudtrafficroutesExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Cloudtrafficroutes.Type(), "tf_cloudtrafficroutes"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccCloudtrafficroutes_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckCloudtrafficroutesExist(resAddr, nil)),
			},
		},
	})
}

const testAccCloudtrafficroutesDataSource_basic = `
resource "citrixadc_cloudtrafficroutes" "tf_cloudtrafficroutes" {
	name             = "tf_cloudtrafficroutes"
	targetvpcnetwork = "tf_vpc_network"
	destrange        = "10.0.0.0/24"
	nexthopip        = "192.168.1.1"
}

data "citrixadc_cloudtrafficroutes" "test" {
	name = citrixadc_cloudtrafficroutes.tf_cloudtrafficroutes.name
}
`

func TestAccCloudtrafficroutesDataSource_basic(t *testing.T) {
	t.Skip(cloudtrafficroutesSkipReason)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudtrafficroutesDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_cloudtrafficroutes.test", "name", "tf_cloudtrafficroutes"),
					resource.TestCheckResourceAttrSet("data.citrixadc_cloudtrafficroutes.test", "id"),
					resource.TestCheckResourceAttr("data.citrixadc_cloudtrafficroutes.test", "destrange", "10.0.0.0/24"),
				),
			},
		},
	})
}

func testAccCheckCloudtrafficroutesExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cloudtrafficroutes name is set")
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
		data, err := client.FindResource(service.Cloudtrafficroutes.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("cloudtrafficroutes %s not found", n)
		}

		return nil
	}
}

func testAccCheckCloudtrafficroutesDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_cloudtrafficroutes" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Cloudtrafficroutes.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("cloudtrafficroutes %s still exists", rs.Primary.ID)
		}
	}

	return nil
}
