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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

const testAccNscqaparam_basic = `


	resource "citrixadc_nscqaparam" "tf_nscqaparam" {
		harqretxdelay = 5
		net1label = "2g"
		minrttnet1 = 25
		lr1coeflist = "intercept=4.95,thruputavg=5.92,iaiavg=-189.48,rttmin=-15.75,loaddelayavg=0.01,noisedelayavg=-2.59"
		lr1probthresh = 0.2
		net1cclscale = "25,50,75"
		net1csqscale = "25,50,75"
		net1logcoef = " 1.49,3.62,-0.14,1.84,4.83"
		net2label = "3g"
		net3label = "4g"
	}
`
const testAccNscqaparam_update = `


	resource "citrixadc_nscqaparam" "tf_nscqaparam" {
		harqretxdelay = 10
		net1label = "3g"
		minrttnet1 = 30
		lr1coeflist = "intercept=4.95,thruputavg=5.92,iaiavg=-189.48,rttmin=-15.75,loaddelayavg=0.01,noisedelayavg=-2.59"
		lr1probthresh = 0.2
		net1cclscale = "25,50,75"
		net1csqscale = "25,50,75"
		net1logcoef = " 1.49,3.62,-0.14,1.84,4.83"
		net2label = "4g"
		net3label = "5g"
	}
`

func TestAccNscqaparam_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNscqaparam_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNscqaparamExist("citrixadc_nscqaparam.tf_nscqaparam", nil),
					resource.TestCheckResourceAttr("citrixadc_nscqaparam.tf_nscqaparam", "harqretxdelay", "5"),
					resource.TestCheckResourceAttr("citrixadc_nscqaparam.tf_nscqaparam", "net1label", "2g"),
					resource.TestCheckResourceAttr("citrixadc_nscqaparam.tf_nscqaparam", "net2label", "3g"),
					resource.TestCheckResourceAttr("citrixadc_nscqaparam.tf_nscqaparam", "net3label", "4g"),
					resource.TestCheckResourceAttr("citrixadc_nscqaparam.tf_nscqaparam", "minrttnet1", "25"),
				),
			},
			{
				Config: testAccNscqaparam_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNscqaparamExist("citrixadc_nscqaparam.tf_nscqaparam", nil),
					resource.TestCheckResourceAttr("citrixadc_nscqaparam.tf_nscqaparam", "harqretxdelay", "10"),
					resource.TestCheckResourceAttr("citrixadc_nscqaparam.tf_nscqaparam", "net1label", "3g"),
					resource.TestCheckResourceAttr("citrixadc_nscqaparam.tf_nscqaparam", "net2label", "4g"),
					resource.TestCheckResourceAttr("citrixadc_nscqaparam.tf_nscqaparam", "net3label", "5g"),
					resource.TestCheckResourceAttr("citrixadc_nscqaparam.tf_nscqaparam", "minrttnet1", "30"),
				),
			},
		},
	})
}

func testAccCheckNscqaparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nscqaparam name is set")
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
		data, err := client.FindResource("nscqaparam", "")

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("nscqaparam %s not found", n)
		}

		return nil
	}
}
