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

// kafkacluster is a CREATE-ONLY, immutable named resource. Its only attribute,
// name, is RequiresReplace, so there is no in-place update step. Step 2 uses a
// different name to force a replacement (destroy + recreate).

const testAccKafkacluster_basic_step1 = `

	resource "citrixadc_kafkacluster" "tf_kafkacluster" {
		name = "tf_kafkacluster"
	}
`

const testAccKafkacluster_basic_step2 = `

	resource "citrixadc_kafkacluster" "tf_kafkacluster" {
		name = "tf_kafkacluster_updated"
	}
`

func TestAccKafkacluster_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckKafkaclusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccKafkacluster_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckKafkaclusterExist("citrixadc_kafkacluster.tf_kafkacluster", nil),
					resource.TestCheckResourceAttr("citrixadc_kafkacluster.tf_kafkacluster", "name", "tf_kafkacluster"),
				),
			},
			{
				// name is RequiresReplace: this forces destroy + recreate.
				Config: testAccKafkacluster_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckKafkaclusterExist("citrixadc_kafkacluster.tf_kafkacluster", nil),
					resource.TestCheckResourceAttr("citrixadc_kafkacluster.tf_kafkacluster", "name", "tf_kafkacluster_updated"),
				),
			},
		},
	})
}

func TestAccKafkacluster_import(t *testing.T) {
	const resAddr = "citrixadc_kafkacluster.tf_kafkacluster"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckKafkaclusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccKafkacluster_basic_step1,
			},
			{
				Config:                  testAccKafkacluster_basic_step1,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckKafkaclusterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No kafkacluster name is set")
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
		data, err := client.FindResource(service.Kafkacluster.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("kafkacluster %s not found", n)
		}

		return nil
	}
}

func testAccCheckKafkaclusterDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_kafkacluster" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Kafkacluster.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("kafkacluster %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccKafkacluster_DataSource_basic = `

	resource "citrixadc_kafkacluster" "tf_kafkacluster" {
		name = "tf_kafkacluster"
	}

	data "citrixadc_kafkacluster" "tf_kafkacluster_data" {
		name       = citrixadc_kafkacluster.tf_kafkacluster.name
		depends_on = [citrixadc_kafkacluster.tf_kafkacluster]
	}
`

func TestAccKafkacluster_DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckKafkaclusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccKafkacluster_DataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_kafkacluster.tf_kafkacluster_data", "name", "tf_kafkacluster"),
				),
			},
		},
	})
}
