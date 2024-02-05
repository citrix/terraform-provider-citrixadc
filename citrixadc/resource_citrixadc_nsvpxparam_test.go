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

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const testAccNsvpxparam_basic_step1 = `
resource "citrixadc_nsvpxparam" "tf_vpxparam" {
    cpuyield = "YES"
    masterclockcpu1 = "YES"
}
`

const testAccNsvpxparam_basic_step2 = `
resource "citrixadc_nsvpxparam" "tf_vpxparam" {
    cpuyield = "DEFAULT"
    masterclockcpu1 = "NO"
}
`

const testAccNsvpxparam_cluster_step1 = `
resource "citrixadc_nsvpxparam" "tf_vpxparam0" {
    cpuyield = "YES"
    masterclockcpu1 = "YES"
    ownernode = 0
}

resource "citrixadc_nsvpxparam" "tf_vpxparam1" {
    cpuyield = "YES"
    masterclockcpu1 = "YES"
    ownernode = 1
}
`

const testAccNsvpxparam_cluster_step2 = `
resource "citrixadc_nsvpxparam" "tf_vpxparam0" {
    cpuyield = "DEFAULT"
    masterclockcpu1 = "NO"
    ownernode = 0
}

resource "citrixadc_nsvpxparam" "tf_vpxparam1" {
    cpuyield = "DEFAULT"
    masterclockcpu1 = "NO"
    ownernode = 1
}
`

func TestAccNsvpxparam_basic(t *testing.T) {
	// if isCluster {
	// 	t.Skip("Use case is applicable for standalone VPX")
	// }
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNsvpxparamDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsvpxparam_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsvpxparamExist("citrixadc_nsvpxparam.tf_vpxparam", nil),
				),
			},
			{
				Config: testAccNsvpxparam_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsvpxparamExist("citrixadc_nsvpxparam.tf_vpxparam", nil),
				),
			},
		},
	})
}

func TestAccNsvpxparam_cluster(t *testing.T) {
	if adcTestbed != "CLUSTER" {
		t.Skipf("ADC testbed is %s. Expected CLUSTER.", adcTestbed)
	}
	// if !isCluster {
	// 	t.Skip("Use case is not applicable to non clustered VPX")
	// }
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNsvpxparamDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNsvpxparam_cluster_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsvpxparamExist("citrixadc_nsvpxparam.tf_vpxparam0", nil),
					testAccCheckNsvpxparamExist("citrixadc_nsvpxparam.tf_vpxparam1", nil),
				),
			},
			{
				Config: testAccNsvpxparam_cluster_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsvpxparamExist("citrixadc_nsvpxparam.tf_vpxparam0", nil),
					testAccCheckNsvpxparamExist("citrixadc_nsvpxparam.tf_vpxparam1", nil),
				),
			},
		},
	})
}

func testAccCheckNsvpxparamExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsvpxparam name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		return nil
	}
}

func testAccCheckNsvpxparamDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_nsvpxparam" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("nsvpxparam", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("nsvpxparam %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
