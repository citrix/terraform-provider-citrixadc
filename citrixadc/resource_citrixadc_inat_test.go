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
	"os"
	"regexp"
	"testing"

	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccInat_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckInatDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccInat_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckInatExist("citrixadc_inat.foo", nil),

					resource.TestCheckResourceAttr(
						"citrixadc_inat.foo", "name", "ip4ip"),
					resource.TestCheckResourceAttr(
						"citrixadc_inat.foo", "privateip", "192.168.1.1"),
					resource.TestCheckResourceAttr(
						"citrixadc_inat.foo", "publicip", "172.16.1.2"),
					resource.TestCheckResourceAttr(
						"citrixadc_inat.foo", "tcpproxy", "ENABLED"),
					resource.TestCheckResourceAttr(
						"citrixadc_inat.foo", "usnip", "ON"),
				),
			},
		},
	})
}

func testAccCheckInatExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No lb vserver name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Inat.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Inat rule %s not found", n)
		}

		return nil
	}
}

func testAccCheckInatDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_inat" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Inat.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Inat rule %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccInat_basic = `


resource "citrixadc_inat" "foo" {

  name = "ip4ip"
  privateip = "192.168.1.1"
  publicip = "172.16.1.2"
  tcpproxy = "ENABLED"
  usnip = "ON"

}
`

func TestAccInat_AssertNonUpdateableAttributes(t *testing.T) {

	if tfAcc := os.Getenv("TF_ACC"); tfAcc == "" {
		t.Skip("TF_ACC not set. Skipping acceptance test.")
	}

	c, err := testHelperInstantiateClient("", "", "", false)
	if err != nil {
		t.Fatalf("Failed to instantiate client. %v\n", err)
	}

	// Create resource
	inatName := "tf-acc-inat-test"
	inatType := service.Inat.Type()

	// Defer deletion of actual resource
	defer testHelperEnsureResourceDeletion(c, t, inatType, inatName, nil)

	inatInstance := network.Inat{
		Name:      inatName,
		Privateip: "192.168.1.1",
		Publicip:  "172.16.1.2",
	}

	if _, err := c.client.AddResource(inatType, inatName, inatInstance); err != nil {
		t.Logf("Error while creating resource")
		t.Fatal(err)
	}

	// publicip
	inatInstance.Publicip = "172.16.1.3"
	testHelperVerifyImmutabilityFunc(c, t, inatType, inatName, inatInstance, "publicip")
	inatInstance.Publicip = ""

	// td
	inatInstance.Td = 1
	testHelperVerifyImmutabilityFunc(c, t, inatType, inatName, inatInstance, "td")
	inatInstance.Td = 0

	// name
	newName := "inat-new-name"
	inatInstance.Name = newName
	if _, err := c.client.UpdateResource(inatType, inatName, inatInstance); err != nil {
		r := regexp.MustCompile(fmt.Sprintf("errorcode.*258.*No such resource \\[name, %s\\]", newName))
		if r.Match([]byte(err.Error())) {
			t.Logf("Succesfully verified immutability of attribute name")
		} else {
			t.Errorf("Error while assesing immutability of attribute name")
			t.Fatal(err)
		}
	}
}
