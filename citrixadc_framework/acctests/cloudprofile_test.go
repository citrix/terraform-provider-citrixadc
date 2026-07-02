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

// cloudprofile is an IMMUTABLE named resource (create + delete only; no NITRO update).
// All attributes are RequiresReplace, so any config change forces resource recreation.
// The basic test therefore uses a single create+read step.
//
// NOTE: ipaddress and vservername may be testbed-specific. If the appliance rejects
// the sample values below, replace the TODO_PLACEHOLDER-marked values with valid ones
// for your ADC (e.g. a free IP in the subnet).

const testAccCloudprofile_basic_step1 = `
resource "citrixadc_cloudprofile" "tf_cloudprofile" {
  name                     = "tf_cloudprofile"
  type                     = "autoscale"
  vservername              = "tf_cloudprofile_vserver"
  servicetype              = "HTTP"
  ipaddress                = "192.0.2.100"
  port                     = 80
  servicegroupname         = "tf_cloudprofile_svcgrp"
  boundservicegroupsvctype = "HTTP"
  vsvrbindsvcport          = 80
  graceful                 = "NO"
}

`

func TestAccCloudprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudprofile_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudprofileExist("citrixadc_cloudprofile.tf_cloudprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_cloudprofile.tf_cloudprofile", "name", "tf_cloudprofile"),
					resource.TestCheckResourceAttr("citrixadc_cloudprofile.tf_cloudprofile", "type", "autoscale"),
					resource.TestCheckResourceAttr("citrixadc_cloudprofile.tf_cloudprofile", "vservername", "tf_cloudprofile_vserver"),
					resource.TestCheckResourceAttr("citrixadc_cloudprofile.tf_cloudprofile", "servicetype", "HTTP"),
					resource.TestCheckResourceAttr("citrixadc_cloudprofile.tf_cloudprofile", "ipaddress", "192.0.2.100"),
					resource.TestCheckResourceAttr("citrixadc_cloudprofile.tf_cloudprofile", "port", "80"),
					resource.TestCheckResourceAttr("citrixadc_cloudprofile.tf_cloudprofile", "servicegroupname", "tf_cloudprofile_svcgrp"),
					resource.TestCheckResourceAttr("citrixadc_cloudprofile.tf_cloudprofile", "boundservicegroupsvctype", "HTTP"),
					resource.TestCheckResourceAttr("citrixadc_cloudprofile.tf_cloudprofile", "vsvrbindsvcport", "80"),
					resource.TestCheckResourceAttr("citrixadc_cloudprofile.tf_cloudprofile", "graceful", "NO"),
				),
			},
		},
	})
}

func testAccCheckCloudprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No cloudprofile name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Cloudprofile.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("cloudprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckCloudprofileDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_cloudprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Cloudprofile.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("cloudprofile %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

const testAccCloudprofileDataSource_basic = `

resource "citrixadc_cloudprofile" "tf_cloudprofile" {
  name                     = "tf_cloudprofile"
  type                     = "autoscale"
  vservername              = "tf_cloudprofile_vserver"
  servicetype              = "HTTP"
  ipaddress                = "192.0.2.100"
  port                     = 80
  servicegroupname         = "tf_cloudprofile_svcgrp"
  boundservicegroupsvctype = "HTTP"
  vsvrbindsvcport          = 80
  graceful                 = "NO"
}

data "citrixadc_cloudprofile" "tf_cloudprofile" {
  name       = citrixadc_cloudprofile.tf_cloudprofile.name
  depends_on = [citrixadc_cloudprofile.tf_cloudprofile]
}
`

func TestAccCloudprofileDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudprofileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudprofileDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_cloudprofile.tf_cloudprofile", "name", "tf_cloudprofile"),
					resource.TestCheckResourceAttr("data.citrixadc_cloudprofile.tf_cloudprofile", "type", "autoscale"),
					resource.TestCheckResourceAttr("data.citrixadc_cloudprofile.tf_cloudprofile", "vservername", "tf_cloudprofile_vserver"),
					resource.TestCheckResourceAttr("data.citrixadc_cloudprofile.tf_cloudprofile", "servicetype", "HTTP"),
					resource.TestCheckResourceAttr("data.citrixadc_cloudprofile.tf_cloudprofile", "servicegroupname", "tf_cloudprofile_svcgrp"),
					resource.TestCheckResourceAttr("data.citrixadc_cloudprofile.tf_cloudprofile", "boundservicegroupsvctype", "HTTP"),
				),
			},
		},
	})
}
