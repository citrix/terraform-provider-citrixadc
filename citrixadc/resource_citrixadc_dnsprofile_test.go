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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

const testAccDnsprofile_add = `


	resource "citrixadc_dnsprofile" "tf_add" {
		
		dnsprofilename      = "tf_profile1"
  		dnsquerylogging     = "DISABLED"
  		dnsanswerseclogging = "DISABLED"
  		dnsextendedlogging  = "DISABLED"
  		dnserrorlogging     = "DISABLED"
  		cacherecords        = "ENABLED"
  		cachenegativeresponses="ENABLED"
  		dropmultiqueryrequest="DISABLED"
  		cacheecsresponses ="DISABLED"
		recursiveresolution = "ENABLED"
		insertecs = "ENABLED"
		replaceecs = "ENABLED"
		maxcacheableecsprefixlength = 16
		maxcacheableecsprefixlength6 = 16
		
	}
`
const testAccDnsprofile_update = `


	resource "citrixadc_dnsprofile" "tf_add" {
		
		dnsprofilename      = "tf_profile1"
  		dnsquerylogging     = "DISABLED"
  		dnsanswerseclogging = "DISABLED"
  		dnsextendedlogging  = "DISABLED"
  		dnserrorlogging     = "DISABLED"
  		cacherecords        = "ENABLED"
  		cachenegativeresponses="ENABLED"
  		dropmultiqueryrequest="ENABLED"
  		cacheecsresponses ="DISABLED"
		recursiveresolution = "DISABLED"
		insertecs = "DISABLED"
		replaceecs = "DISABLED"
		maxcacheableecsprefixlength = 18
		maxcacheableecsprefixlength6 = 18
		
	}
`

func TestAccDnsprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDnsprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsprofile_add,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsprofileExist("citrixadc_dnsprofile.tf_add", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "dnsprofilename", "tf_profile1"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "dnsquerylogging", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "dnsanswerseclogging", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "dnsextendedlogging", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "dnserrorlogging", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "cacherecords", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "cachenegativeresponses", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "dropmultiqueryrequest", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "cacheecsresponses", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "recursiveresolution", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "insertecs", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "replaceecs", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "maxcacheableecsprefixlength", "16"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "maxcacheableecsprefixlength6", "16"),
				),
			},
			{
				Config: testAccDnsprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsprofileExist("citrixadc_dnsprofile.tf_add", nil),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "dnsprofilename", "tf_profile1"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "dnsquerylogging", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "dnsanswerseclogging", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "dnsextendedlogging", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "dnserrorlogging", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "cacherecords", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "cachenegativeresponses", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "dropmultiqueryrequest", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "cacheecsresponses", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "recursiveresolution", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "insertecs", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "replaceecs", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "maxcacheableecsprefixlength", "18"),
					resource.TestCheckResourceAttr("citrixadc_dnsprofile.tf_add", "maxcacheableecsprefixlength6", "18"),
				),
			},
		},
	})
}

func testAccCheckDnsprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dnsprofile name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Dnsprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("dnsprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckDnsprofileDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_dnsprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Dnsprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("dnsprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
