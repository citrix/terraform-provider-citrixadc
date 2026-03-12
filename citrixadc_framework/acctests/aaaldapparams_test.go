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

const testAccAaaldapparams_basic = `
	resource "citrixadc_aaaldapparams" "tf_aaaldapparams" {
		authtimeout = 4
		serverip    = "10.222.74.158"
		passwdchange = "DISABLED"
	}
  
`
const testAccAaaldapparams_update = `
	resource "citrixadc_aaaldapparams" "tf_aaaldapparams" {
		authtimeout = 5
		serverip    = "10.222.74.158"
		passwdchange = "ENABLED"
	}
  
`

func TestAccAaaldapparams_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccAaaldapparams_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaldapparamsExist("citrixadc_aaaldapparams.tf_aaaldapparams", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaldapparams.tf_aaaldapparams", "authtimeout", "4"),
					resource.TestCheckResourceAttr("citrixadc_aaaldapparams.tf_aaaldapparams", "serverip", "10.222.74.158"),
					resource.TestCheckResourceAttr("citrixadc_aaaldapparams.tf_aaaldapparams", "passwdchange", "DISABLED"),
				),
			},
			{
				Config: testAccAaaldapparams_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAaaldapparamsExist("citrixadc_aaaldapparams.tf_aaaldapparams", nil),
					resource.TestCheckResourceAttr("citrixadc_aaaldapparams.tf_aaaldapparams", "authtimeout", "5"),
					resource.TestCheckResourceAttr("citrixadc_aaaldapparams.tf_aaaldapparams", "serverip", "10.222.74.158"),
					resource.TestCheckResourceAttr("citrixadc_aaaldapparams.tf_aaaldapparams", "passwdchange", "ENABLED"),
				),
			},
		},
	})
}

func testAccCheckAaaldapparamsExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No aaaldapparams name is set")
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
		data, err := client.FindResource(service.Aaaldapparams.Type(), "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("aaaldapparams %s not found", n)
		}

		return nil
	}
}

// Add this to the end of aaaldapparams_test.go

const testAccAaaldapparamsDataSource_basic = `

resource "citrixadc_aaaldapparams" "tf_aaaldapparams" {
	serverip = "1.2.3.4"
	serverport = 389
	authtimeout = 5
	ldapbase = "dc=aaa,dc=local"
	ldapbinddn = "cn=Manager,dc=aaa,dc=local"
	ldapbinddnpassword = "secret"
	ldaploginname = "samAccountName"
	searchfilter = "cn"
	groupattrname = "memberOf"
	subattributename = "cn"
	sectype = "PLAINTEXT"
	passwdchange = "ENABLED"
	nestedgroupextraction = "OFF"
	maxnestinglevel = 3
	groupnameidentifier = "samAccountName"
	groupsearchattribute = "memberOf"
	groupsearchsubattribute = "cn"
	groupsearchfilter = "memberOf"
	defaultauthenticationgroup = "default_group"
}

data "citrixadc_aaaldapparams" "tf_aaaldapparams" {
	depends_on = [citrixadc_aaaldapparams.tf_aaaldapparams]
}
`

func TestAccAaaldapparamsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAaaldapparamsDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "serverip", "1.2.3.4"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "serverport", "389"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "authtimeout", "5"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "ldapbase", "dc=aaa,dc=local"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "ldapbinddn", "cn=Manager,dc=aaa,dc=local"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "ldaploginname", "samAccountName"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "searchfilter", "cn"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "groupattrname", "memberOf"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "subattributename", "cn"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "sectype", "PLAINTEXT"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "passwdchange", "ENABLED"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "nestedgroupextraction", "OFF"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "maxnestinglevel", "3"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "groupnameidentifier", "samAccountName"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "groupsearchattribute", "memberOf"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "groupsearchsubattribute", "cn"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "groupsearchfilter", "memberOf"),
					resource.TestCheckResourceAttr("data.citrixadc_aaaldapparams.tf_aaaldapparams", "defaultauthenticationgroup", "default_group"),
				),
			},
		},
	})
}
