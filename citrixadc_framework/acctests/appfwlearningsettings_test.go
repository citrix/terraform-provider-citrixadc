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
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

const testAccAppfwlearningsettings_add = `
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
	resource "citrixadc_appfwlearningsettings" "tf_learningsetting" {
		profilename                        = citrixadc_appfwprofile.tf_appfwprofile.name
		starturlminthreshold               = 9
		starturlpercentthreshold           = 10
		cookieconsistencyminthreshold      = 2
		cookieconsistencypercentthreshold  = 1
		csrftagminthreshold                = 2
		csrftagpercentthreshold            = 10
		fieldconsistencyminthreshold       = 20
		fieldconsistencypercentthreshold   = 8
		crosssitescriptingminthreshold     = 10
		crosssitescriptingpercentthreshold = 1
		sqlinjectionminthreshold           = 10
		sqlinjectionpercentthreshold       = 1
		fieldformatminthreshold            = 10
		fieldformatpercentthreshold        = 1
		creditcardnumberminthreshold       = 1
		creditcardnumberpercentthreshold   = 0
		contenttypeminthreshold            = 1
		contenttypepercentthreshold        = 0
	}
`
const testAccAppfwlearningsettings_update = `
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name                     = "tf_appfwprofile"
		type                     = ["HTML"]
	}
	resource "citrixadc_appfwlearningsettings" "tf_learningsetting" {
		profilename                        = citrixadc_appfwprofile.tf_appfwprofile.name
		starturlminthreshold               = 1
		starturlpercentthreshold           = 0
		cookieconsistencyminthreshold      = 1
		cookieconsistencypercentthreshold  = 0
		csrftagminthreshold                = 1
		csrftagpercentthreshold            = 0
		fieldconsistencyminthreshold       = 20
		fieldconsistencypercentthreshold   = 8
		crosssitescriptingminthreshold     = 10
		crosssitescriptingpercentthreshold = 1
		sqlinjectionminthreshold           = 10
		sqlinjectionpercentthreshold       = 1
		fieldformatminthreshold            = 10
		fieldformatpercentthreshold        = 1
		creditcardnumberminthreshold       = 1
		creditcardnumberpercentthreshold   = 0
		contenttypeminthreshold            = 1
		contenttypepercentthreshold        = 0
	}
`

func TestAccAppfwlearningsettings_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil, //testAccCheckAppfwlearningsettingsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwlearningsettings_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwlearningsettingsExist("citrixadc_appfwlearningsettings.tf_learningsetting", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwlearningsettings.tf_learningsetting", "profilename", "tf_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwlearningsettings.tf_learningsetting", "starturlminthreshold", "9"),
					resource.TestCheckResourceAttr("citrixadc_appfwlearningsettings.tf_learningsetting", "starturlpercentthreshold", "10"),
					resource.TestCheckResourceAttr("citrixadc_appfwlearningsettings.tf_learningsetting", "cookieconsistencyminthreshold", "2"),
					resource.TestCheckResourceAttr("citrixadc_appfwlearningsettings.tf_learningsetting", "cookieconsistencypercentthreshold", "1"),
					resource.TestCheckResourceAttr("citrixadc_appfwlearningsettings.tf_learningsetting", "csrftagminthreshold", "2"),
				),
			},
			{
				Config: testAccAppfwlearningsettings_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwlearningsettingsExist("citrixadc_appfwlearningsettings.tf_learningsetting", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwlearningsettings.tf_learningsetting", "profilename", "tf_appfwprofile"),
					resource.TestCheckResourceAttr("citrixadc_appfwlearningsettings.tf_learningsetting", "starturlminthreshold", "1"),
					resource.TestCheckResourceAttr("citrixadc_appfwlearningsettings.tf_learningsetting", "starturlpercentthreshold", "0"),
					resource.TestCheckResourceAttr("citrixadc_appfwlearningsettings.tf_learningsetting", "cookieconsistencyminthreshold", "1"),
					resource.TestCheckResourceAttr("citrixadc_appfwlearningsettings.tf_learningsetting", "cookieconsistencypercentthreshold", "0"),
					resource.TestCheckResourceAttr("citrixadc_appfwlearningsettings.tf_learningsetting", "csrftagminthreshold", "1"),
				),
			},
		},
	})
}

func testAccCheckAppfwlearningsettingsExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No appfwlearningsettings name is set")
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
		data, err := client.FindResource(service.Appfwlearningsettings.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("appfwlearningsettings %s not found", n)
		}

		return nil
	}
}
