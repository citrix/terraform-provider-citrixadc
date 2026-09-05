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
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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

func TestAccAppfwlearningsettings_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.0.0"},
				},
				Config: testAccAppfwlearningsettings_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppfwlearningsettingsExist("citrixadc_appfwlearningsettings.tf_learningsetting", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{expectNoReplace()},
				},
				Config: testAccAppfwlearningsettings_add,
				Check:  resource.ComposeTestCheckFunc(testAccCheckAppfwlearningsettingsExist("citrixadc_appfwlearningsettings.tf_learningsetting", nil)),
			},
		},
	})
}

func TestAccAppfwlearningsettings_import(t *testing.T) {
	const resAddr = "citrixadc_appfwlearningsettings.tf_learningsetting"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccAppfwlearningsettings_add},
			{
				Config:                  testAccAppfwlearningsettings_add,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
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

const testAccAppfwlearningsettingsDataSource_basic = `

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

	data "citrixadc_appfwlearningsettings" "tf_learningsetting" {
		profilename = citrixadc_appfwlearningsettings.tf_learningsetting.profilename
		depends_on = [citrixadc_appfwlearningsettings.tf_learningsetting]
	}
`

// testAccAppfwlearningsettings_unset covers every spec-unsettable attribute.
// Step 1 sets each to a valid non-default value; step 2 removes them all from
// config, so the provider must issue a NITRO unset that reverts each to its
// documented default (minthreshold=1, percentthreshold=0, graceperiod=10080).
const testAccAppfwlearningsettings_unset_step1 = `
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name = "tf_appfwprofile_unset"
		type = ["HTML"]
	}
	resource "citrixadc_appfwlearningsettings" "tf_unset" {
		profilename                             = citrixadc_appfwprofile.tf_appfwprofile.name
		contenttypeautodeploygraceperiod        = 20
		contenttypeminthreshold                 = 5
		contenttypepercentthreshold             = 10
		cookieconsistencyautodeploygraceperiod  = 20
		cookieconsistencyminthreshold           = 5
		cookieconsistencypercentthreshold       = 10
		creditcardnumberminthreshold            = 5
		creditcardnumberpercentthreshold        = 10
		crosssitescriptingautodeploygraceperiod = 20
		crosssitescriptingminthreshold          = 5
		crosssitescriptingpercentthreshold      = 10
		csrftagautodeploygraceperiod            = 20
		csrftagminthreshold                     = 5
		csrftagpercentthreshold                 = 10
		fieldconsistencyautodeploygraceperiod   = 20
		fieldconsistencyminthreshold            = 5
		fieldconsistencypercentthreshold        = 10
		fieldformatautodeploygraceperiod        = 20
		fieldformatminthreshold                 = 5
		fieldformatpercentthreshold             = 10
		sqlinjectionautodeploygraceperiod       = 20
		sqlinjectionminthreshold                = 5
		sqlinjectionpercentthreshold            = 10
		starturlautodeploygraceperiod           = 20
		starturlminthreshold                    = 5
		starturlpercentthreshold                = 10
		xmlattachmentminthreshold               = 5
		xmlattachmentpercentthreshold           = 10
		xmlwsiminthreshold                      = 5
		xmlwsipercentthreshold                  = 10
	}
`

const testAccAppfwlearningsettings_unset_step2 = `
	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name = "tf_appfwprofile_unset"
		type = ["HTML"]
	}
	resource "citrixadc_appfwlearningsettings" "tf_unset" {
		profilename = citrixadc_appfwprofile.tf_appfwprofile.name
		# All unset-eligible attributes removed from config -> the provider must
		# unset them (revert to NITRO defaults).
	}
`

func TestAccAppfwlearningsettings_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccAppfwlearningsettings_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwlearningsettingsExist("citrixadc_appfwlearningsettings.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwlearningsettings.tf_unset", "contenttypeautodeploygraceperiod", "20"),
					resource.TestCheckResourceAttr("citrixadc_appfwlearningsettings.tf_unset", "contenttypeminthreshold", "5"),
					resource.TestCheckResourceAttr("citrixadc_appfwlearningsettings.tf_unset", "contenttypepercentthreshold", "10"),
					resource.TestCheckResourceAttr("citrixadc_appfwlearningsettings.tf_unset", "starturlminthreshold", "5"),
					resource.TestCheckResourceAttr("citrixadc_appfwlearningsettings.tf_unset", "starturlpercentthreshold", "10"),
					resource.TestCheckResourceAttr("citrixadc_appfwlearningsettings.tf_unset", "xmlwsiminthreshold", "5"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccAppfwlearningsettings_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAppfwlearningsettingsExist("citrixadc_appfwlearningsettings.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_appfwlearningsettings.tf_unset", "contenttypeautodeploygraceperiod", "10080"),
					resource.TestCheckResourceAttr("citrixadc_appfwlearningsettings.tf_unset", "contenttypeminthreshold", "1"),
					resource.TestCheckResourceAttr("citrixadc_appfwlearningsettings.tf_unset", "contenttypepercentthreshold", "0"),
					resource.TestCheckResourceAttr("citrixadc_appfwlearningsettings.tf_unset", "starturlminthreshold", "1"),
					resource.TestCheckResourceAttr("citrixadc_appfwlearningsettings.tf_unset", "starturlpercentthreshold", "0"),
					resource.TestCheckResourceAttr("citrixadc_appfwlearningsettings.tf_unset", "xmlwsiminthreshold", "1"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckAppfwlearningsettingsADCValue("tf_appfwprofile_unset", "contenttypeautodeploygraceperiod", "10080"),
					testAccCheckAppfwlearningsettingsADCValue("tf_appfwprofile_unset", "starturlminthreshold", "1"),
					testAccCheckAppfwlearningsettingsADCValue("tf_appfwprofile_unset", "starturlpercentthreshold", "0"),
				),
			},
		},
	})
}

// testAccCheckAppfwlearningsettingsADCValue asserts an attribute's value directly
// on the appliance (not just in Terraform state), proving the unset reverted it.
func testAccCheckAppfwlearningsettingsADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Appfwlearningsettings.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("appfwlearningsettings %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("appfwlearningsettings %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccAppfwlearningsettingsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAppfwlearningsettingsDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_appfwlearningsettings.tf_learningsetting", "profilename", "tf_appfwprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwlearningsettings.tf_learningsetting", "starturlminthreshold", "9"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwlearningsettings.tf_learningsetting", "starturlpercentthreshold", "10"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwlearningsettings.tf_learningsetting", "cookieconsistencyminthreshold", "2"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwlearningsettings.tf_learningsetting", "cookieconsistencypercentthreshold", "1"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwlearningsettings.tf_learningsetting", "csrftagminthreshold", "2"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwlearningsettings.tf_learningsetting", "csrftagpercentthreshold", "10"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwlearningsettings.tf_learningsetting", "fieldconsistencyminthreshold", "20"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwlearningsettings.tf_learningsetting", "fieldconsistencypercentthreshold", "8"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwlearningsettings.tf_learningsetting", "crosssitescriptingminthreshold", "10"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwlearningsettings.tf_learningsetting", "crosssitescriptingpercentthreshold", "1"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwlearningsettings.tf_learningsetting", "sqlinjectionminthreshold", "10"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwlearningsettings.tf_learningsetting", "sqlinjectionpercentthreshold", "1"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwlearningsettings.tf_learningsetting", "fieldformatminthreshold", "10"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwlearningsettings.tf_learningsetting", "fieldformatpercentthreshold", "1"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwlearningsettings.tf_learningsetting", "creditcardnumberminthreshold", "1"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwlearningsettings.tf_learningsetting", "creditcardnumberpercentthreshold", "0"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwlearningsettings.tf_learningsetting", "contenttypeminthreshold", "1"),
					resource.TestCheckResourceAttr("data.citrixadc_appfwlearningsettings.tf_learningsetting", "contenttypepercentthreshold", "0"),
				),
			},
		},
	})
}
