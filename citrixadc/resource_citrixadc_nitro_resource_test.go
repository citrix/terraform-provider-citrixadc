/*
Copyright 2022 Citrix Systems, Inc

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
	"io/ioutil"
	"net/url"
	"path"
	"runtime"
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"gopkg.in/yaml.v2"
)

const testAccNitroResource_object_basic_step1 = `
resource "citrixadc_nitro_resource" "tf_lbvserver" {
    workflows_file = "testdata/workflows.yaml"
    workflow = "lbvserver"

    attributes = {
      ipv46       = "10.10.10.33"
    }

    non_updateable_attributes = {
      name        = "tf_lbvserver"
      servicetype = "HTTP"
      port        = 80
    }
}
`
const testAccNitroResource_object_basic_step2 = `
resource "citrixadc_nitro_resource" "tf_lbvserver" {
    workflows_file = "testdata/workflows.yaml"
    workflow = "lbvserver"

    attributes = {
      ipv46       = "10.10.10.44"
    }

    non_updateable_attributes = {
      name        = "tf_lbvserver"
      servicetype = "HTTP"
      port        = 80
    }
}
`
const testAccNitroResource_object_basic_step3 = `
resource "citrixadc_nitro_resource" "tf_lbvserver" {
    workflows_file = "testdata/workflows.yaml"
    workflow = "lbvserver"

    attributes = {
      ipv46       = "10.10.10.44"
    }

    non_updateable_attributes = {
      name        = "tf_lbvserver"
      servicetype = "HTTP"
      port        = 90
    }
}
`

func TestAccNitroResource_object_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckObjectDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNitroResource_object_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists("citrixadc_nitro_resource.tf_lbvserver", nil),
					resource.TestCheckResourceAttr("citrixadc_nitro_resource.tf_lbvserver", "non_updateable_attributes.name", "tf_lbvserver"),
					resource.TestCheckResourceAttr("citrixadc_nitro_resource.tf_lbvserver", "non_updateable_attributes.servicetype", "HTTP"),
					resource.TestCheckResourceAttr("citrixadc_nitro_resource.tf_lbvserver", "non_updateable_attributes.port", "80"),
					resource.TestCheckResourceAttr("citrixadc_nitro_resource.tf_lbvserver", "attributes.ipv46", "10.10.10.33"),
				),
			},
			resource.TestStep{
				Config: testAccNitroResource_object_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists("citrixadc_nitro_resource.tf_lbvserver", nil),
					resource.TestCheckResourceAttr("citrixadc_nitro_resource.tf_lbvserver", "non_updateable_attributes.name", "tf_lbvserver"),
					resource.TestCheckResourceAttr("citrixadc_nitro_resource.tf_lbvserver", "non_updateable_attributes.servicetype", "HTTP"),
					resource.TestCheckResourceAttr("citrixadc_nitro_resource.tf_lbvserver", "non_updateable_attributes.port", "80"),
					resource.TestCheckResourceAttr("citrixadc_nitro_resource.tf_lbvserver", "attributes.ipv46", "10.10.10.44"),
				),
			},
			resource.TestStep{
				Config: testAccNitroResource_object_basic_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectExists("citrixadc_nitro_resource.tf_lbvserver", nil),
					resource.TestCheckResourceAttr("citrixadc_nitro_resource.tf_lbvserver", "non_updateable_attributes.name", "tf_lbvserver"),
					resource.TestCheckResourceAttr("citrixadc_nitro_resource.tf_lbvserver", "non_updateable_attributes.servicetype", "HTTP"),
					resource.TestCheckResourceAttr("citrixadc_nitro_resource.tf_lbvserver", "non_updateable_attributes.port", "90"),
					resource.TestCheckResourceAttr("citrixadc_nitro_resource.tf_lbvserver", "attributes.ipv46", "10.10.10.44"),
				),
			},
		},
	})
}

func testAccCheckObjectExists(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No object primary id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}
		workflowKey := rs.Primary.Attributes["workflow"]
		wf, err := testHelperReadWorkflowDict(workflowKey)
		_ = wf
		if err != nil {
			return err
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

		findParams := service.FindParams{
			ResourceType:             wf["endpoint"].(string),
			ResourceName:             rs.Primary.ID,
			ResourceMissingErrorCode: wf["resource_missing_errorcode"].(int),
		}

		dataArr, err := nsClient.FindResourceArrayWithParams(findParams)

		if err != nil {
			return err
		}

		if len(dataArr) == 0 {
			return fmt.Errorf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		}

		if len(dataArr) > 1 {
			return fmt.Errorf("FindResourceArrayWithParams returned too many results")
		}

		return nil
	}
}

func testAccCheckObjectDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nitro_resource" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		workflowKey := rs.Primary.Attributes["workflow"]
		wf, err := testHelperReadWorkflowDict(workflowKey)
		_ = wf
		if err != nil {
			return err
		}

		findParams := service.FindParams{
			ResourceType:             wf["endpoint"].(string),
			ResourceName:             rs.Primary.ID,
			ResourceMissingErrorCode: wf["resource_missing_errorcode"].(int),
		}

		dataArr, err := nsClient.FindResourceArrayWithParams(findParams)

		if err != nil {
			return err
		}

		if len(dataArr) >= 1 {
			return fmt.Errorf("Object still exists")
		}
	}

	return nil
}

const testAccNitroResource_binding_basic_step1 = `
resource "citrixadc_service" "tf_service" {

    name = "tf_service"
    ip = "192.168.43.33"
    servicetype  = "HTTP"
    port = 80
}

resource "citrixadc_lbvserver" "tf_lbvserver" {

  name        = "tf_lbvserver"
  ipv46       = "10.10.10.33"
  servicetype = "HTTP"
  port        = 80
}

resource "citrixadc_nitro_resource" "tf_lbvserver_service_bind" {
    workflows_file = "testdata/workflows.yaml"
    workflow = "lbvserver_service_binding"

    non_updateable_attributes = {
        name = citrixadc_lbvserver.tf_lbvserver.name
        servicename = citrixadc_service.tf_service.name
        weight = 2
    }
}
`

const testAccNitroResource_binding_basic_step2 = `
resource "citrixadc_service" "tf_service" {

    name = "tf_service"
    ip = "192.168.43.33"
    servicetype  = "HTTP"
    port = 80
}

resource "citrixadc_lbvserver" "tf_lbvserver" {

  name        = "tf_lbvserver"
  ipv46       = "10.10.10.33"
  servicetype = "HTTP"
  port        = 80
}

resource "citrixadc_nitro_resource" "tf_lbvserver_service_bind" {
    workflows_file = "testdata/workflows.yaml"
    workflow = "lbvserver_service_binding"

    non_updateable_attributes = {
        name = citrixadc_lbvserver.tf_lbvserver.name
        servicename = citrixadc_service.tf_service.name
        weight = 3
    }
}
`

func TestAccNitroResource_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBindingDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNitroResource_binding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBindingExists("citrixadc_nitro_resource.tf_lbvserver_service_bind", nil),
					resource.TestCheckResourceAttr("citrixadc_nitro_resource.tf_lbvserver_service_bind", "non_updateable_attributes.weight", "2"),
				),
			},
			resource.TestStep{
				Config: testAccNitroResource_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBindingExists("citrixadc_nitro_resource.tf_lbvserver_service_bind", nil),
					resource.TestCheckResourceAttr("citrixadc_nitro_resource.tf_lbvserver_service_bind", "non_updateable_attributes.weight", "3"),
				),
			},
		},
	})
}

func testAccCheckBindingExists(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No object primary id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}
		workflowKey := rs.Primary.Attributes["workflow"]
		wf, err := testHelperReadWorkflowDict(workflowKey)
		_ = wf
		if err != nil {
			return err
		}
		idSlice := strings.SplitN(rs.Primary.ID, ",", 2)

		primaryId := idSlice[0]
		secondaryId := idSlice[1]

		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		findParams := service.FindParams{
			ResourceType:             wf["endpoint"].(string),
			ResourceName:             primaryId,
			ResourceMissingErrorCode: wf["bound_resource_missing_errorcode"].(int),
		}

		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		if len(dataArr) == 0 {
			return fmt.Errorf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		}
		secondaryIdAttribute := wf["secondary_id_attribute"]
		foundIndex := -1
		for index, binding := range dataArr {
			if fmt.Sprintf("%v", binding[secondaryIdAttribute.(string)]) == secondaryId {
				foundIndex = index
				break
			}
		}

		// Resource is missing
		if foundIndex == -1 {
			return fmt.Errorf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams secondary id attribute not found in array")
		}

		return nil
	}
}

func testAccCheckBindingDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nitro_resource" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No id is set")
		}

		workflowKey := rs.Primary.Attributes["workflow"]
		wf, err := testHelperReadWorkflowDict(workflowKey)
		_ = wf
		if err != nil {
			return err
		}

		idSlice := strings.SplitN(rs.Primary.ID, ",", 2)

		primaryId := idSlice[0]
		secondaryId := idSlice[1]

		client := testAccProvider.Meta().(*NetScalerNitroClient).client

		findParams := service.FindParams{
			ResourceType:             wf["endpoint"].(string),
			ResourceName:             primaryId,
			ResourceMissingErrorCode: wf["bound_resource_missing_errorcode"].(int),
		}

		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		if len(dataArr) == 0 {
			return nil
		}
		secondaryIdAttribute := wf["secondary_id_attribute"]
		foundIndex := -1
		for index, binding := range dataArr {
			if fmt.Sprintf("%v", binding[secondaryIdAttribute.(string)]) == secondaryId {
				foundIndex = index
				break
			}
		}

		// Resource is missing
		if foundIndex != -1 {
			return fmt.Errorf("FindResourceArrayWithParams secondary id attribute found in array")
		}

		return nil

	}

	return nil
}

func testHelperReadWorkflowDict(workflow string) (map[interface{}]interface{}, error) {
	// Get here path
	_, here_filename, _, _ := runtime.Caller(1)

	fileData, err := ioutil.ReadFile(path.Join(path.Dir(here_filename), "testdata", "workflows.yaml"))
	if err != nil {
		return nil, fmt.Errorf("Error reading testdata workflows %s", err)
	}

	var data interface{}
	err = yaml.Unmarshal(fileData, &data)
	if err != nil {
		return nil, err
	}

	workflowsDict, ok := data.(map[interface{}]interface{})["workflow"]
	if !ok {
		return nil, fmt.Errorf("Top level workflow key not found in workflows yaml file")
	}

	specificWorkflow, ok := workflowsDict.(map[interface{}]interface{})[workflow]
	if !ok {
		return nil, fmt.Errorf("Key %v not found in workflows map", workflow)
	}

	return specificWorkflow.(map[interface{}]interface{}), nil
}

const testAccNitroResource_object_by_args_basic_step1 = `

resource "citrixadc_nitro_resource" "tf_snmpmanager" {
  workflows_file = "testdata/workflows.yaml"
  workflow       = "snmpmanager"

  attributes = {
    domainresolveretry   = 10
  }

  non_updateable_attributes = {
    ipaddress = "helo1234.com"
  }
}
`
const testAccNitroResource_object_by_args_basic_step2 = `

resource "citrixadc_nitro_resource" "tf_snmpmanager" {
  workflows_file = "testdata/workflows.yaml"
  workflow       = "snmpmanager"

  attributes = {
    domainresolveretry   = 30
  }

  non_updateable_attributes = {
    ipaddress = "helo1234.com"
  }
}
`

const testAccNitroResource_object_by_args_basic_step3 = `

resource "citrixadc_nitro_resource" "tf_snmpmanager" {
  workflows_file = "testdata/workflows.yaml"
  workflow       = "snmpmanager"

  attributes = {
    domainresolveretry   = 30
  }

  non_updateable_attributes = {
    ipaddress = "helo123456.com"
  }
}
`

func TestAccNitroResource_object_by_args_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckObjectByArgsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNitroResource_object_by_args_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectByArgsExists("citrixadc_nitro_resource.tf_snmpmanager", nil),
					resource.TestCheckResourceAttr("citrixadc_nitro_resource.tf_snmpmanager", "non_updateable_attributes.ipaddress", "helo1234.com"),
					resource.TestCheckResourceAttr("citrixadc_nitro_resource.tf_snmpmanager", "attributes.domainresolveretry", "10"),
				),
			},
			resource.TestStep{
				Config: testAccNitroResource_object_by_args_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectByArgsExists("citrixadc_nitro_resource.tf_snmpmanager", nil),
					resource.TestCheckResourceAttr("citrixadc_nitro_resource.tf_snmpmanager", "non_updateable_attributes.ipaddress", "helo1234.com"),
					resource.TestCheckResourceAttr("citrixadc_nitro_resource.tf_snmpmanager", "attributes.domainresolveretry", "30"),
				),
			},
			resource.TestStep{
				Config: testAccNitroResource_object_by_args_basic_step3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckObjectByArgsExists("citrixadc_nitro_resource.tf_snmpmanager", nil),
					resource.TestCheckResourceAttr("citrixadc_nitro_resource.tf_snmpmanager", "non_updateable_attributes.ipaddress", "helo123456.com"),
					resource.TestCheckResourceAttr("citrixadc_nitro_resource.tf_snmpmanager", "attributes.domainresolveretry", "30"),
				),
			},
		},
	})
}

func testAccCheckObjectByArgsExists(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No object primary id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}
		workflowKey := rs.Primary.Attributes["workflow"]
		wf, err := testHelperReadWorkflowDict(workflowKey)
		_ = wf
		if err != nil {
			return err
		}

		primaryId := rs.Primary.ID
		idItems := strings.Split(primaryId, ",")
		argsMap := make(map[string]string)

		for _, idItem := range idItems {
			idSlice := strings.Split(idItem, ":")
			key := url.QueryEscape(idSlice[0])
			value := url.QueryEscape(idSlice[1])
			argsMap[key] = value
		}

		findParams := service.FindParams{
			ResourceType:             wf["endpoint"].(string),
			ArgsMap:                  argsMap,
			ResourceMissingErrorCode: wf["resource_missing_errorcode"].(int),
		}

		client := testAccProvider.Meta().(*NetScalerNitroClient).client
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return fmt.Errorf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
		}

		if len(dataArr) == 0 {
			return fmt.Errorf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		}

		if len(dataArr) > 1 {
			return fmt.Errorf("FindResourceArrayWithParams returned too many results")
		}

		return nil
	}
}

func testAccCheckObjectByArgsDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nitro_resource" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No id is set")
		}

		workflowKey := rs.Primary.Attributes["workflow"]
		wf, err := testHelperReadWorkflowDict(workflowKey)
		_ = wf
		if err != nil {
			return err
		}

		primaryId := rs.Primary.ID
		idItems := strings.Split(primaryId, ",")
		argsMap := make(map[string]string)

		for _, idItem := range idItems {
			idSlice := strings.Split(idItem, ":")
			key := url.QueryEscape(idSlice[0])
			value := url.QueryEscape(idSlice[1])
			argsMap[key] = value
		}

		findParams := service.FindParams{
			ResourceType:             wf["endpoint"].(string),
			ArgsMap:                  argsMap,
			ResourceMissingErrorCode: wf["resource_missing_errorcode"].(int),
		}

		client := testAccProvider.Meta().(*NetScalerNitroClient).client
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return fmt.Errorf("[DEBUG] citrixadc-provider: Error during FindResourceArrayWithParams %s", err.Error())
		}

		if len(dataArr) != 0 {
			return fmt.Errorf("[DEBUG] citrixadc-provider: FindResourceArrayWithParams returned empty array")
		}

		return nil

	}

	return nil
}
