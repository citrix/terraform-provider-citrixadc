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

const testAccSystemparameter_basic = `

resource "citrixadc_systemparameter" "tf_systemparameter" {
    rbaonresponse = "ENABLED"
    natpcbforceflushlimit = 3000
    natpcbrstontimeout = "DISABLED"
    timeout = 500
    doppler = "ENABLED"
	pwdhistorycount = 5
	warnpriorndays = 10
	passwordhistorycontrol = "ENABLED"
	maxsessionperuser = 10
	daystoexpire = 45
}
`

const testAccSystemparameter_update = `

resource "citrixadc_systemparameter" "tf_systemparameter" {
    rbaonresponse = "DISABLED"
    natpcbforceflushlimit = 2000
    natpcbrstontimeout = "ENABLED"
    timeout = 600
    doppler = "DISABLED"
	pwdhistorycount = 10
	warnpriorndays = 15
	passwordhistorycontrol = "DISABLED"
	maxsessionperuser = 15
	daystoexpire = 50
}
`

const testAccSystemparameterDataSource_basic = `

resource "citrixadc_systemparameter" "tf_systemparameter" {
    rbaonresponse = "ENABLED"
    natpcbforceflushlimit = 3000
    natpcbrstontimeout = "DISABLED"
    timeout = 500
    doppler = "ENABLED"
	pwdhistorycount = 5
	warnpriorndays = 10
	passwordhistorycontrol = "ENABLED"
	maxsessionperuser = 10
	daystoexpire = 45
}

data "citrixadc_systemparameter" "tf_systemparameter" {
    depends_on = [citrixadc_systemparameter.tf_systemparameter]
}
`

func TestAccSystemparameter_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemparameter_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemparameterExist("citrixadc_systemparameter.tf_systemparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_systemparameter", "rbaonresponse", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_systemparameter", "natpcbforceflushlimit", "3000"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_systemparameter", "natpcbrstontimeout", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_systemparameter", "timeout", "500"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_systemparameter", "doppler", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_systemparameter", "pwdhistorycount", "5"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_systemparameter", "warnpriorndays", "10"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_systemparameter", "passwordhistorycontrol", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_systemparameter", "maxsessionperuser", "10"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_systemparameter", "daystoexpire", "45"),
				),
			},
			{
				Config: testAccSystemparameter_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemparameterExist("citrixadc_systemparameter.tf_systemparameter", nil),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_systemparameter", "rbaonresponse", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_systemparameter", "natpcbforceflushlimit", "2000"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_systemparameter", "natpcbrstontimeout", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_systemparameter", "timeout", "600"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_systemparameter", "doppler", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_systemparameter", "pwdhistorycount", "10"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_systemparameter", "warnpriorndays", "15"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_systemparameter", "passwordhistorycontrol", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_systemparameter", "maxsessionperuser", "15"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_systemparameter", "daystoexpire", "50"),
				),
			},
		},
	})
}

func TestAccSystemparameter_import(t *testing.T) {
	const resAddr = "citrixadc_systemparameter.tf_systemparameter"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{Config: testAccSystemparameter_basic},
			{
				Config:                  testAccSystemparameter_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCheckSystemparameterExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No systemparameter name is set")
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
		data, err := client.FindResource(service.Systemparameter.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("systemparameter %s not found", n)
		}

		return nil
	}
}

func TestAccSystemparameter_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccSystemparameter_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckSystemparameterExist("citrixadc_systemparameter.tf_systemparameter", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccSystemparameter_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckSystemparameterExist("citrixadc_systemparameter.tf_systemparameter", nil)),
			},
		},
	})
}

// testAccSystemparameter_unset_step1 sets every unset-eligible attribute to a
// valid NON-default value.
const testAccSystemparameter_unset_step1 = `
resource "citrixadc_systemparameter" "tf_unset" {
    rbaonresponse           = "DISABLED"
    natpcbforceflushlimit   = 3000
    natpcbrstontimeout      = "ENABLED"
    timeout                 = 500
    doppler                 = "DISABLED"
    googleanalytics         = "ENABLED"
    totalauthtimeout        = 30
    cliloglevel             = "DEBUG"
    reauthonauthparamchange = "ENABLED"
    removesensitivefiles    = "ENABLED"
    restrictedtimeout       = "ENABLED"
}
`

// testAccSystemparameter_unset_step2 removes all unset-eligible attributes so
// the provider must unset them (revert to the documented NITRO defaults).
const testAccSystemparameter_unset_step2 = `
resource "citrixadc_systemparameter" "tf_unset" {
}
`

func TestAccSystemparameter_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccSystemparameter_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemparameterExist("citrixadc_systemparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_unset", "rbaonresponse", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_unset", "natpcbforceflushlimit", "3000"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_unset", "natpcbrstontimeout", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_unset", "timeout", "500"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_unset", "doppler", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_unset", "googleanalytics", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_unset", "totalauthtimeout", "30"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_unset", "cliloglevel", "DEBUG"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_unset", "reauthonauthparamchange", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_unset", "removesensitivefiles", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_unset", "restrictedtimeout", "ENABLED"),
				),
			},
			{
				// Removing the attributes must unset them: state (read back from
				// the appliance) reverts to the documented NITRO defaults, and the
				// implicit post-apply plan must be empty.
				Config: testAccSystemparameter_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemparameterExist("citrixadc_systemparameter.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_unset", "rbaonresponse", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_unset", "natpcbforceflushlimit", "2147483647"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_unset", "natpcbrstontimeout", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_unset", "timeout", "900"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_unset", "doppler", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_unset", "googleanalytics", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_unset", "totalauthtimeout", "20"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_unset", "cliloglevel", "INFORMATIONAL"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_unset", "reauthonauthparamchange", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_unset", "removesensitivefiles", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_systemparameter.tf_unset", "restrictedtimeout", "DISABLED"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckSystemparameterADCValue("rbaonresponse", "ENABLED"),
					testAccCheckSystemparameterADCValue("cliloglevel", "INFORMATIONAL"),
					testAccCheckSystemparameterADCValue("timeout", "900"),
				),
			},
		},
	})
}

// testAccCheckSystemparameterADCValue asserts an attribute's value directly on
// the appliance (not just in Terraform state), proving the unset actually
// reverted it.
func testAccCheckSystemparameterADCValue(attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Systemparameter.Type(), "")
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("systemparameter not found on appliance")
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("systemparameter: appliance attr %q = %q, want %q (unset did not revert it)", attr, got, want)
		}
		return nil
	}
}

func TestAccSystemparameterDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemparameterDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_systemparameter.tf_systemparameter", "rbaonresponse", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_systemparameter.tf_systemparameter", "natpcbforceflushlimit", "3000"),
					resource.TestCheckResourceAttr("data.citrixadc_systemparameter.tf_systemparameter", "natpcbrstontimeout", "DISABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_systemparameter.tf_systemparameter", "timeout", "500"),
					resource.TestCheckResourceAttr("data.citrixadc_systemparameter.tf_systemparameter", "doppler", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_systemparameter.tf_systemparameter", "pwdhistorycount", "5"),
					resource.TestCheckResourceAttr("data.citrixadc_systemparameter.tf_systemparameter", "warnpriorndays", "10"),
					resource.TestCheckResourceAttr("data.citrixadc_systemparameter.tf_systemparameter", "passwordhistorycontrol", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_systemparameter.tf_systemparameter", "maxsessionperuser", "10"),
					resource.TestCheckResourceAttr("data.citrixadc_systemparameter.tf_systemparameter", "daystoexpire", "45"),
				),
			},
		},
	})
}
