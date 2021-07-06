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
	"bytes"
	"errors"
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccNsacls_basic(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNsaclsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNsacls_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsaclsExist("citrixadc_nsacls.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_nsacls.foo", "acl.1629125634.aclname", "restricttcp2"),
					resource.TestCheckResourceAttr("citrixadc_nsacls.foo", "acl.1629125634.protocol", "TCP"),
					resource.TestCheckResourceAttr("citrixadc_nsacls.foo", "acl.1629125634.destipval", "192.168.199.52"),
					resource.TestCheckResourceAttr("citrixadc_nsacls.foo", "acl.1056203765.aclname", "allowudp"),
					resource.TestCheckResourceAttr("citrixadc_nsacls.foo", "acl.1056203765.protocol", "UDP"),
					resource.TestCheckResourceAttr("citrixadc_nsacls.foo", "acl.1056203765.destipval", "192.168.45.55"),
				),
			},
		},
	})
}

func testAccCheckNsaclsExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsacls name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		deviceAcls, err := nsClient.FindAllResources(service.Nsacl.Type())

		if err != nil {
			return err
		}
		acl1 := map[string]interface{}{
			"aclname":    "restricttcp2",
			"protocol":   "TCP",
			"aclaction":  "DENY",
			"destipval":  "192.168.199.52",
			"srcportval": "149-1524",
			"priority":   "25",
		}
		acl2 := map[string]interface{}{
			"aclname":    "allowudp",
			"protocol":   "UDP",
			"aclaction":  "ALLOW",
			"destipval":  "192.168.45.55",
			"srcportval": "490-1024",
			"priority":   "100",
		}

		acl3 := map[string]interface{}{
			"aclname":   "restrictvlan",
			"aclaction": "DENY",
			"vlan":      "2000",
			"priority":  "250",
		}
		found1, found2, found3 := false, false, false
		for _, acl := range deviceAcls {
			if testMapEquals(acl1, acl) {
				found1 = true
			}
			if testMapEquals(acl2, acl) {
				found2 = true
			}
			if testMapEquals(acl3, acl) {
				found3 = true
			}
		}
		if found1 && found2 && found3 {
			//fmt.Printf("netscaler-provider testNsAcls Found acls\n")
		} else {
			//fmt.Printf("netscaler-provider testNsAcls Did not find all acls\n")
			return fmt.Errorf("netscaler-provider testNsAcls Did not find all acls")
		}

		return nil
	}
}

func TestAccNsacls_update(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNsaclsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNsacls_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsaclsExist("citrixadc_nsacls.foo", nil),
					resource.TestCheckResourceAttr("citrixadc_nsacls.foo", "acl.1629125634.aclname", "restricttcp2"),
				),
			},
			resource.TestStep{
				Config: testAccNsacls_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNsaclsUpdateExist("citrixadc_nsacls.foo", nil),
				),
			},
		},
	})
}

func testAccCheckNsaclsDestroy(s *terraform.State) error {

	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
	deviceAcls, err := nsClient.FindAllResources(service.Nsacl.Type())
	if err != nil {
		return err
	}

	messageBuffer := bytes.NewBuffer(make([]byte, 0, 0))

	foundDanglingAcl := false
	for _, acl := range deviceAcls {
		log.Printf("acl found %v\n", acl)
		for _, aclname := range []string{"restricttcp2", "restrictvlan", "allowudp"} {
			if acl["aclname"] == aclname {
				foundDanglingAcl = true
				if _, err := messageBuffer.WriteString(fmt.Sprintf("Dangling acl %s\n", aclname)); err != nil {
					return errors.New("Error appending acl name to message buffer")
				}
			}
		}
	}
	if foundDanglingAcl {
		return fmt.Errorf("citrixadc-provider testAccCheckNsaclsDestroy: %s", messageBuffer.String())
	}

	return nil
}

func testAccCheckNsaclsUpdateExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No nsacls name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

		deviceAcls, err := nsClient.FindAllResources(service.Nsacl.Type())
		if err != nil {
			return err
		}
		acl1 := map[string]interface{}{
			"aclname":    "restricttcp2",
			"protocol":   "TCP",
			"aclaction":  "DENY",
			"destipval":  "192.168.22.22",
			"srcportval": "222-2222",
			"priority":   "25",
		}
		acl2 := map[string]interface{}{
			"aclname":   "restrictvlan",
			"aclaction": "DENY",
			"vlan":      "2000",
			"priority":  "250",
		}

		found1, found2 := false, false
		for _, acl := range deviceAcls {
			if testMapEquals(acl1, acl) {
				found1 = true
			}
			if testMapEquals(acl2, acl) {
				found2 = true
			}

		}
		if found1 && found2 {
			//fmt.Printf("netscaler-provider testNsAclsUpdate Found acls\n")
		} else {
			//fmt.Printf("netscaler-provider testNsAclsUpdate Did not find all acls\n")
			return fmt.Errorf("netscaler-provider testNsAcls Did not find all acls")
		}

		return nil
	}
}

func testMapEquals(m1 map[string]interface{}, m2 map[string]interface{}) bool {
	//test that m2 has the field values present in m1
	eq := true
	for k, v := range m1 {
		eq = eq && reflect.DeepEqual(m2[k], v)
	}
	return eq
}

const testAccNsacls_basic = `


resource "citrixadc_nsacls" "foo" {
 acl  {
    aclname = "restricttcp2"
    protocol = "TCP"
    aclaction = "DENY"
    destipval = "192.168.199.52"
    srcportval = "149-1524"
    priority = "25"
  }

  acl  {
    aclname = "allowudp"
    protocol = "UDP"
    aclaction = "ALLOW"
    destipval = "192.168.45.55"
    srcportval = "490-1024"
       priority = "100"
  }

  acl  {
    aclname = "restrictvlan"
    aclaction = "DENY"
    vlan = "2000"
    priority = "250"
  }


}
`
const testAccNsacls_update = `


resource "citrixadc_nsacls" "foo" {
 acl  {
    aclname = "restricttcp2"
    protocol = "TCP"
    aclaction = "DENY"
    destipval = "192.168.22.22"
    srcportval = "222-2222"
    priority = "25"
  }

  acl  {
    aclname = "restrictvlan"
    aclaction = "DENY"
    vlan = "2000"
    priority = "250"
  }


}
`
