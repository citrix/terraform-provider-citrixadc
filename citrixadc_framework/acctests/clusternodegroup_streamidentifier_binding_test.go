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

const testAccClusternodegroup_streamidentifier_binding_basic = `

	resource "citrixadc_streamselector" "tf_streamselector" {
		name = "my_streamselector"
		rule = ["HTTP.REQ.URL", "CLIENT.IP.SRC"]
	}
	resource "citrixadc_streamidentifier" "tf_streamidentifier" {
		name         = "my_streamidentifier"
		selectorname = citrixadc_streamselector.tf_streamselector.name
		samplecount  = 10
		sort         = "CONNECTIONS"
		snmptrap     = "ENABLED"
		loglimit	 = 500
		loginterval = 60
		log = "NONE"
	}

	resource "citrixadc_clusternodegroup_streamidentifier_binding" "tf_clusternodegroup_streamidentifier_binding" {
		name           = "my_tf_group"
		identifiername = citrixadc_streamidentifier.tf_streamidentifier.name
	}
`

const testAccClusternodegroup_streamidentifier_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion
`

func TestAccClusternodegroup_streamidentifier_binding_basic(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckClusternodegroup_streamidentifier_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClusternodegroup_streamidentifier_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroup_streamidentifier_bindingExist("citrixadc_clusternodegroup_streamidentifier_binding.tf_clusternodegroup_streamidentifier_binding", nil),
				),
			},
			{
				Config: testAccClusternodegroup_streamidentifier_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClusternodegroup_streamidentifier_bindingNotExist("citrixadc_clusternodegroup_streamidentifier_binding.tf_clusternodegroup_streamidentifier_binding", "my_group,my_identifier"),
				),
			},
		},
	})
}

func testAccCheckClusternodegroup_streamidentifier_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No clusternodegroup_streamidentifier_binding id is set")
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

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 2)

		name := idSlice[0]
		identifiername := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "clusternodegroup_streamidentifier_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching identifiername
		found := false
		for _, v := range dataArr {
			if v["identifiername"].(string) == identifiername {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("clusternodegroup_streamidentifier_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckClusternodegroup_streamidentifier_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		name := idSlice[0]
		identifiername := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "clusternodegroup_streamidentifier_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching identifiername
		found := false
		for _, v := range dataArr {
			if v["identifiername"].(string) == identifiername {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("clusternodegroup_streamidentifier_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckClusternodegroup_streamidentifier_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_clusternodegroup_streamidentifier_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Clusternodegroup_streamidentifier_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("clusternodegroup_streamidentifier_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccClusternodegroup_streamidentifier_bindingDataSource_basic = `

	resource "citrixadc_streamselector" "tf_streamselector" {
		name = "my_streamselector"
		rule = ["HTTP.REQ.URL", "CLIENT.IP.SRC"]
	}
	resource "citrixadc_streamidentifier" "tf_streamidentifier" {
		name         = "my_streamidentifier"
		selectorname = citrixadc_streamselector.tf_streamselector.name
		samplecount  = 10
		sort         = "CONNECTIONS"
		snmptrap     = "ENABLED"
		loglimit	 = 500
		loginterval = 60
		log = "NONE"
	}

resource "citrixadc_clusternodegroup_streamidentifier_binding" "tf_clusternodegroup_streamidentifier_binding" {
name           = "my_tf_group"
identifiername = citrixadc_streamidentifier.tf_streamidentifier.name
}

data "citrixadc_clusternodegroup_streamidentifier_binding" "tf_clusternodegroup_streamidentifier_binding" {
name           = citrixadc_clusternodegroup_streamidentifier_binding.tf_clusternodegroup_streamidentifier_binding.name
identifiername = citrixadc_clusternodegroup_streamidentifier_binding.tf_clusternodegroup_streamidentifier_binding.identifiername
depends_on = [citrixadc_clusternodegroup_streamidentifier_binding.tf_clusternodegroup_streamidentifier_binding]
}
`

func TestAccclusternodegroup_streamidentifier_bindingDataSource_basic(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccClusternodegroup_streamidentifier_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_clusternodegroup_streamidentifier_binding.tf_clusternodegroup_streamidentifier_binding", "name", "my_tf_group"),
					resource.TestCheckResourceAttr("data.citrixadc_clusternodegroup_streamidentifier_binding.tf_clusternodegroup_streamidentifier_binding", "identifiername", "my_streamidentifier"),
				),
			},
		},
	})
}
