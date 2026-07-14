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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// kafkacluster_servicegroup_binding is an immutable binding (bind/unbind only,
// no NITRO set/update endpoint). Both name (parent kafkacluster) and
// servicegroupname are RequiresReplace, so there is NO in-place update step.
// The basic test only creates the binding (step1) and then removes it while
// keeping the participating entities (step2) to verify clean unbind.
//
// Participating entity config is lifted from:
//   - kafkacluster_test.go   (resource "citrixadc_kafkacluster")
//   - servicegroup_test.go   (resource "citrixadc_servicegroup", KAFKA_BROKER servicetype)

const testAccKafkacluster_servicegroup_binding_basic_step1 = `

	resource "citrixadc_kafkacluster" "tf_kafkacluster" {
		name = "tf_kafkacluster"
	}

	resource "citrixadc_servicegroup" "tf_servicegroup" {
		servicegroupname = "tf_kafka_servicegroup"
		servicetype      = "KAFKA_BROKER"
		bootstrap        = "YES"
	}

	resource "citrixadc_kafkacluster_servicegroup_binding" "tf_binding" {
		name             = citrixadc_kafkacluster.tf_kafkacluster.name
		servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
		depends_on = [
			citrixadc_kafkacluster.tf_kafkacluster,
			citrixadc_servicegroup.tf_servicegroup,
		]
	}
`

const testAccKafkacluster_servicegroup_binding_basic_step2 = `
	# Keep the participating entities without the binding to verify clean unbind.
	resource "citrixadc_kafkacluster" "tf_kafkacluster" {
		name = "tf_kafkacluster"
	}

	resource "citrixadc_servicegroup" "tf_servicegroup" {
		servicegroupname = "tf_kafka_servicegroup"
		servicetype      = "KAFKA_BROKER"
		bootstrap        = "YES"
	}
`

func TestAccKafkacluster_servicegroup_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckKafkacluster_servicegroup_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccKafkacluster_servicegroup_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckKafkacluster_servicegroup_bindingExist("citrixadc_kafkacluster_servicegroup_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_kafkacluster_servicegroup_binding.tf_binding", "name", "tf_kafkacluster"),
					resource.TestCheckResourceAttr("citrixadc_kafkacluster_servicegroup_binding.tf_binding", "servicegroupname", "tf_kafka_servicegroup"),
				),
			},
			{
				Config: testAccKafkacluster_servicegroup_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckKafkacluster_servicegroup_bindingNotExist("tf_kafkacluster", "tf_kafka_servicegroup"),
				),
			},
		},
	})
}

func testAccCheckKafkacluster_servicegroup_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No kafkacluster_servicegroup_binding id is set")
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

		// ID format is "name:<enc>,servicegroupname:<enc>"
		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %q: %v", rs.Primary.ID, err)
		}

		name := idMap["name"]
		servicegroupname := idMap["servicegroupname"]

		findParams := service.FindParams{
			ResourceType:             service.Kafkacluster_servicegroup_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if val, ok := v["servicegroupname"].(string); ok && val == servicegroupname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("kafkacluster_servicegroup_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckKafkacluster_servicegroup_bindingNotExist(name string, servicegroupname string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		findParams := service.FindParams{
			ResourceType:             service.Kafkacluster_servicegroup_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if val, ok := v["servicegroupname"].(string); ok && val == servicegroupname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("kafkacluster_servicegroup_binding (%s,%s) was found, but it should have been destroyed", name, servicegroupname)
		}

		return nil
	}
}

func testAccCheckKafkacluster_servicegroup_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_kafkacluster_servicegroup_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No id is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %q: %v", rs.Primary.ID, err)
		}

		name := idMap["name"]
		servicegroupname := idMap["servicegroupname"]

		findParams := service.FindParams{
			ResourceType:             service.Kafkacluster_servicegroup_binding.Type(),
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Parent gone / no bindings returned means the binding is destroyed.
			continue
		}

		for _, v := range dataArr {
			if val, ok := v["servicegroupname"].(string); ok && val == servicegroupname {
				return fmt.Errorf("kafkacluster_servicegroup_binding %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

const testAccKafkacluster_servicegroup_binding_DataSource_basic = `

	resource "citrixadc_kafkacluster" "tf_kafkacluster" {
		name = "tf_kafkacluster"
	}

	resource "citrixadc_servicegroup" "tf_servicegroup" {
		servicegroupname = "tf_kafka_servicegroup"
		servicetype      = "KAFKA_BROKER"
		bootstrap        = "YES"
	}

	resource "citrixadc_kafkacluster_servicegroup_binding" "tf_binding" {
		name             = citrixadc_kafkacluster.tf_kafkacluster.name
		servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
		depends_on = [
			citrixadc_kafkacluster.tf_kafkacluster,
			citrixadc_servicegroup.tf_servicegroup,
		]
	}

	data "citrixadc_kafkacluster_servicegroup_binding" "tf_binding" {
		name             = citrixadc_kafkacluster.tf_kafkacluster.name
		servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
		depends_on       = [citrixadc_kafkacluster_servicegroup_binding.tf_binding]
	}
`

func TestAccKafkacluster_servicegroup_binding_DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckKafkacluster_servicegroup_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccKafkacluster_servicegroup_binding_DataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_kafkacluster_servicegroup_binding.tf_binding", "name", "tf_kafkacluster"),
					resource.TestCheckResourceAttr("data.citrixadc_kafkacluster_servicegroup_binding.tf_binding", "servicegroupname", "tf_kafka_servicegroup"),
				),
			},
		},
	})
}
