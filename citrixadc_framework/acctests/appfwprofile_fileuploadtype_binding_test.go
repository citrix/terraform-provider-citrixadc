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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccAppfwprofile_fileuploadtype_binding_basic = `
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
	resource "citrixadc_appfwprofile_fileuploadtype_binding" "tf_binding1" {
		name                   = citrixadc_appfwprofile.tf_appfwprofile.name
		fileuploadtype         = "tf_uploadtype"
		as_fileuploadtypes_url = "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$"
		filetype               = ["pdf", "text"]
	}
	resource "citrixadc_appfwprofile_fileuploadtype_binding" "tf_binding2" {
		name                   = citrixadc_appfwprofile.tf_appfwprofile.name
		fileuploadtype         = "tf_uploadtype"
		as_fileuploadtypes_url = "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/v1/resource/temp$"
		filetype               = ["pdf", "text"]
	}
	resource "citrixadc_appfwprofile_fileuploadtype_binding" "tf_binding3" {
		name                   = citrixadc_appfwprofile.tf_appfwprofile.name
		fileuploadtype         = "tf_uploadtype"
		as_fileuploadtypes_url = "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/v1/resource/temp$"
		filetype               = ["text"]
	}
`

const testAccAppfwprofile_fileuploadtype_binding_basic_step2 = `
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
`

func TestAccAppfwprofile_fileuploadtype_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAppfwprofile_fileuploadtype_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwprofile_fileuploadtype_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_fileuploadtype_bindingExist("citrixadc_appfwprofile_fileuploadtype_binding.tf_binding1", nil),
					// resource.TestCheckResourceAttr("citrixadc_appfwprofile_fileuploadtype_binding.tf_binding1", "name", "tf_appfwprofile"),
					// resource.TestCheckResourceAttr("citrixadc_appfwprofile_fileuploadtype_binding.tf_binding1", "fileuploadtype", "tf_uploadtype"),
					// resource.TestCheckResourceAttr("citrixadc_appfwprofile_fileuploadtype_binding.tf_binding1", "as_fileuploadtypes_url", "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$"),
					// testAccCheckAppfwprofile_fileuploadtype_bindingExist("citrixadc_appfwprofile_fileuploadtype_binding.tf_binding2", nil),
					// resource.TestCheckResourceAttr("citrixadc_appfwprofile_fileuploadtype_binding.tf_binding2", "name", "tf_appfwprofile"),
					// resource.TestCheckResourceAttr("citrixadc_appfwprofile_fileuploadtype_binding.tf_binding2", "fileuploadtype", "tf_uploadtype"),
					// resource.TestCheckResourceAttr("citrixadc_appfwprofile_fileuploadtype_binding.tf_binding2", "as_fileuploadtypes_url", "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/v1/resource/temp$"),
					// testAccCheckAppfwprofile_fileuploadtype_bindingExist("citrixadc_appfwprofile_fileuploadtype_binding.tf_binding3", nil),
					// resource.TestCheckResourceAttr("citrixadc_appfwprofile_fileuploadtype_binding.tf_binding3", "name", "tf_appfwprofile"),
					// resource.TestCheckResourceAttr("citrixadc_appfwprofile_fileuploadtype_binding.tf_binding3", "fileuploadtype", "tf_uploadtype"),
					// resource.TestCheckResourceAttr("citrixadc_appfwprofile_fileuploadtype_binding.tf_binding3", "as_fileuploadtypes_url", "^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/v1/resource/temp$"),
				),
			},
			{
				Config: testAccAppfwprofile_fileuploadtype_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwprofile_fileuploadtype_bindingNotExist("citrixadc_appfwprofile_fileuploadtype_binding.tf_binding1", "tf_appfwprofile,tf_uploadtype,^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/document/content$,text;pdf"),
					// testAccCheckAppfwprofile_fileuploadtype_bindingNotExist("citrixadc_appfwprofile_fileuploadtype_binding.tf_binding2", "tf_appfwprofile,tf_uploadtype,^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/v1/resource/temp$,pdf;text"),
					// testAccCheckAppfwprofile_fileuploadtype_bindingNotExist("citrixadc_appfwprofile_fileuploadtype_binding.tf_binding3", "tf_appfwprofile,tf_uploadtype,^https://sd2\\-zgw\\.test\\.ctxns\\.com/api/v1/resource/temp$,text"),
				),
			},
		},
	})
}

func testAccCheckAppfwprofile_fileuploadtype_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwprofile_fileuploadtype_binding id is set")
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

		idSlice := strings.SplitN(bindingId, ",", 4)

		name := idSlice[0]
		fileuploadtype := idSlice[1]
		as_fileuploadtypes_url := idSlice[2]
		filetype := idSlice[3]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_fileuploadtype_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["fileuploadtype"].(string) == fileuploadtype {
				if v["as_fileuploadtypes_url"].(string) == as_fileuploadtypes_url {
					// Check if filetype matches (convert slice to space-separated string for comparison)
					dataFiletype := ""
					if v["filetype"] != nil {
						if filetypeSlice, ok := v["filetype"].([]interface{}); ok {
							dataFiletype = strings.Join(utils.ToStringList(filetypeSlice), ";")
						}
					}
					if dataFiletype == filetype {
						found = true
						break
					}
				}
			}
		}

		if !found {
			return fmt.Errorf("appfwprofile_fileuploadtype_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_fileuploadtype_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 4)

		name := idSlice[0]
		fileuploadtype := idSlice[1]
		as_fileuploadtypes_url := idSlice[2]
		filetype := idSlice[3]

		findParams := service.FindParams{
			ResourceType:             "appfwprofile_fileuploadtype_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching secondIdComponent
		found := false
		for _, v := range dataArr {
			if v["fileuploadtype"].(string) == fileuploadtype {
				if v["as_fileuploadtypes_url"].(string) == as_fileuploadtypes_url {
					// Check if filetype matches (convert slice to space-separated string for comparison)
					dataFiletype := ""
					if v["filetype"] != nil {
						if filetypeSlice, ok := v["filetype"].([]interface{}); ok {
							dataFiletype = strings.Join(utils.ToStringList(filetypeSlice), ";")
						}
					}
					if dataFiletype == filetype {
						found = true
						break
					}
				}
			}
		}

		if found {
			return fmt.Errorf("appfwprofile_fileuploadtype_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckAppfwprofile_fileuploadtype_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_appfwprofile_fileuploadtype_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource("appfwprofile_fileuploadtype_binding", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("appfwprofile_fileuploadtype_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
