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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccResponderpolicy_globalbinding(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckResponderpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResponderpolicy_globalbinding_exists,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResponderpolicyExist("citrixadc_responderpolicy.tf_responder_policy", nil),
					verifyGlobalBindingExists("REQ_OVERRIDE", "tf_responder_policy", false),
				),
			},
			{
				Config: testAccResponderpolicy_globalbinding_modified,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResponderpolicyExist("citrixadc_responderpolicy.tf_responder_policy", nil),
					verifyGlobalBindingExists("REQ_OVERRIDE", "tf_responder_policy", false),
				),
			},
			{
				Config: testAccResponderpolicy_globalbinding_not_exists,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResponderpolicyExist("citrixadc_responderpolicy.tf_responder_policy", nil),
					verifyGlobalBindingExists("REQ_OVERRIDE", "tf_responder_policy", true),
				),
			},
		},
	})
}

func testAccCheckResponderpolicyExist(n string, id *string) resource.TestCheckFunc {
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
		data, err := nsClient.FindResource(service.Responderpolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("Responder policy %s not found", n)
		}

		return nil
	}
}

func verifyGlobalBindingExists(bindtype string, policyname string, inverse bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		bindFound := false
		client := testAccProvider.Meta().(*NetScalerNitroClient).client
		globalBindings, _ := client.FindResourceArray("responderpolicy_responderglobal_binding", policyname)
		for _, val := range globalBindings {
			boundtoSlice := strings.Split(val["boundto"].(string), " ")
			if bindtype == boundtoSlice[1] {
				bindFound = true
				break
			}
		}

		if !inverse {
			if bindFound {
				return nil
			} else {
				return fmt.Errorf("Verify error cannot find bind of type %v for policyname %v\n", bindtype, policyname)
			}
		} else {
			if bindFound {
				return fmt.Errorf("Verify error found exessive bind of type %v for policyname %v\n", bindtype, policyname)
			} else {
				return nil
			}
		}
	}
}

func verifyLbvserverBindingExists(servername string, policyname string, inverse bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		bindFound := false
		client := testAccProvider.Meta().(*NetScalerNitroClient).client
		lbVserverBindings, _ := client.FindResourceArray("responderpolicy_lbvserver_binding", policyname)
		for _, val := range lbVserverBindings {
			boundtoSlice := strings.Split(val["boundto"].(string), " ")
			if servername == boundtoSlice[2] {
				bindFound = true
				break
			}
		}

		if !inverse {
			if bindFound {
				return nil
			} else {
				return fmt.Errorf("Verify error cannot find bind to lbvserver %v for policyname %v\n", servername, policyname)
			}
		} else {
			if bindFound {
				return fmt.Errorf("Verify error found exessive bind to lbvserver %v for policyname %v\n", servername, policyname)
			} else {
				return nil
			}
		}
	}
}

func testAccCheckResponderpolicyDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_responderpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Responderpolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func TestAccResponderpolicy_lbvserverbinding(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckResponderpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResponderpolicy_lbvserverbinding_both,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResponderpolicyExist("citrixadc_responderpolicy.tf_responder_policy", nil),
					verifyLbvserverBindingExists("tf_lbvserver1", "tf_responder_policy", false),
					verifyLbvserverBindingExists("tf_lbvserver2", "tf_responder_policy", false),
				),
			},
			{
				Config: testAccResponderpolicy_lbvserverbinding_one,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResponderpolicyExist("citrixadc_responderpolicy.tf_responder_policy", nil),
					verifyLbvserverBindingExists("tf_lbvserver1", "tf_responder_policy", false),
					verifyLbvserverBindingExists("tf_lbvserver2", "tf_responder_policy", true),
				),
			},
			{
				Config: testAccResponderpolicy_lbvserverbinding_both,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResponderpolicyExist("citrixadc_responderpolicy.tf_responder_policy", nil),
				),
			},
		},
	})
}

func TestAccResponderpolicy_csvserverbinding(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckResponderpolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResponderpolicy_csvserverbinding_both,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResponderpolicyExist("citrixadc_responderpolicy.tf_responder_policy", nil),
					verifyCsvserverBindingExists("tf_csvserver1", "tf_responder_policy", false),
					verifyCsvserverBindingExists("tf_csvserver2", "tf_responder_policy", false),
				),
			},
			{
				Config: testAccResponderpolicy_csvserverbinding_one,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResponderpolicyExist("citrixadc_responderpolicy.tf_responder_policy", nil),
					verifyCsvserverBindingExists("tf_csvserver1", "tf_responder_policy", false),
					verifyCsvserverBindingExists("tf_csvserver2", "tf_responder_policy", true),
				),
			},
			{
				Config: testAccResponderpolicy_csvserverbinding_none,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResponderpolicyExist("citrixadc_responderpolicy.tf_responder_policy", nil),
					verifyCsvserverBindingExists("tf_csvserver1", "tf_responder_policy", true),
					verifyCsvserverBindingExists("tf_csvserver2", "tf_responder_policy", true),
				),
			},
			{
				Config: testAccResponderpolicy_csvserverbinding_one,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResponderpolicyExist("citrixadc_responderpolicy.tf_responder_policy", nil),
					verifyCsvserverBindingExists("tf_csvserver1", "tf_responder_policy", false),
					verifyCsvserverBindingExists("tf_csvserver2", "tf_responder_policy", true),
				),
			},
		},
	})
}

func verifyCsvserverBindingExists(servername string, policyname string, inverse bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		bindFound := false
		client := testAccProvider.Meta().(*NetScalerNitroClient).client
		lbVserverBindings, _ := client.FindResourceArray("responderpolicy_csvserver_binding", policyname)
		for _, val := range lbVserverBindings {
			boundtoSlice := strings.Split(val["boundto"].(string), " ")
			if servername == boundtoSlice[2] {
				bindFound = true
				break
			}
		}

		if !inverse {
			if bindFound {
				return nil
			} else {
				return fmt.Errorf("Verify error cannot find bind to csvserver %v for policyname %v\n", servername, policyname)
			}
		} else {
			if bindFound {
				return fmt.Errorf("Verify error found exessive bind to csvserver %v for policyname %v\n", servername, policyname)
			} else {
				return nil
			}
		}
	}
}

const testAccResponderpolicy_globalbinding_exists = `

resource "citrixadc_lbvserver" "tf_lbvserver" {

  ipv46 = "192.168.43.66"
  name = "tf_lbvserver"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_responderpolicy" "tf_responder_policy" {
  name    = "tf_responder_policy"
  action = "NOOP"
  rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"nosuchthing\")"

  globalbinding {
      invoke = true
      labeltype = "vserver"
      labelname = citrixadc_lbvserver.tf_lbvserver.name
      type = "REQ_OVERRIDE"
      gotopriorityexpression = "END"
      priority = 666
  }
}
`

const testAccResponderpolicy_globalbinding_modified = `

resource "citrixadc_lbvserver" "tf_lbvserver" {

  ipv46 = "192.168.43.66"
  name = "tf_lbvserver"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_responderpolicy" "tf_responder_policy" {
  name    = "tf_responder_policy"
  action = "NOOP"
  rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"nosuchthing\")"

  globalbinding {
      invoke = true
      labeltype = "vserver"
      labelname = citrixadc_lbvserver.tf_lbvserver.name
      type = "REQ_OVERRIDE"
      gotopriorityexpression = "END"
      priority = 777
  }
}
`

const testAccResponderpolicy_globalbinding_not_exists = `

resource "citrixadc_lbvserver" "tf_lbvserver" {

  ipv46 = "192.168.43.66"
  name = "tf_lbvserver"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_responderpolicy" "tf_responder_policy" {
  name    = "tf_responder_policy"
  action = "NOOP"
  rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"nosuchthing\")"
  depends_on = [ citrixadc_lbvserver.tf_lbvserver ]
}
`

const testAccResponderpolicy_lbvserverbinding_both = `

resource "citrixadc_lbvserver" "tf_lbvserver1" {

  ipv46 = "192.168.43.66"
  name = "tf_lbvserver1"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_lbvserver" "tf_lbvserver2" {

  ipv46 = "192.168.43.67"
  name = "tf_lbvserver2"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_responderpolicy" "tf_responder_policy" {
  name    = "tf_responder_policy"
  action = "NOOP"
  rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"nosuchthing\")"

  lbvserverbinding {
      priority = 100
      name = citrixadc_lbvserver.tf_lbvserver1.name
      gotopriorityexpression = "END"
      invoke = false
      bindpoint = "REQUEST"
  }

  lbvserverbinding {
      priority = 100
      name = citrixadc_lbvserver.tf_lbvserver2.name
      gotopriorityexpression = "END"
      invoke = false
      bindpoint = "REQUEST"
  }
}
`

const testAccResponderpolicy_lbvserverbinding_one = `

resource "citrixadc_lbvserver" "tf_lbvserver1" {

  ipv46 = "192.168.43.66"
  name = "tf_lbvserver1"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_lbvserver" "tf_lbvserver2" {

  ipv46 = "192.168.43.67"
  name = "tf_lbvserver2"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_responderpolicy" "tf_responder_policy" {
  name    = "tf_responder_policy"
  action = "NOOP"
  rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"nosuchthing\")"

  lbvserverbinding {
      priority = 100
      name = citrixadc_lbvserver.tf_lbvserver1.name
      gotopriorityexpression = "END"
      invoke = false
      bindpoint = "REQUEST"
  }
}
`

const testAccResponderpolicy_csvserverbinding_both = `

resource "citrixadc_csvserver" "tf_csvserver1" {

  ipv46 = "192.168.43.66"
  name = "tf_csvserver1"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_csvserver" "tf_csvserver2" {

  ipv46 = "192.168.43.67"
  name = "tf_csvserver2"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_responderpolicy" "tf_responder_policy" {
  name    = "tf_responder_policy"
  action = "NOOP"
  rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"nosuchthing\")"

  csvserverbinding {
      priority = 100
      name = citrixadc_csvserver.tf_csvserver1.name
      gotopriorityexpression = "END"
      invoke = false
      bindpoint = "REQUEST"
  }

  csvserverbinding {
      priority = 100
      name = citrixadc_csvserver.tf_csvserver2.name
      gotopriorityexpression = "END"
      invoke = false
      bindpoint = "REQUEST"
  }
}
`

const testAccResponderpolicy_csvserverbinding_one = `

resource "citrixadc_csvserver" "tf_csvserver1" {

  ipv46 = "192.168.43.66"
  name = "tf_csvserver1"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_csvserver" "tf_csvserver2" {

  ipv46 = "192.168.43.67"
  name = "tf_csvserver2"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_responderpolicy" "tf_responder_policy" {
  name    = "tf_responder_policy"
  action = "NOOP"
  rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"nosuchthing\")"

  csvserverbinding {
      priority = 100
      name = citrixadc_csvserver.tf_csvserver1.name
      gotopriorityexpression = "END"
      invoke = false
      bindpoint = "REQUEST"
  }
}
`
const testAccResponderpolicy_csvserverbinding_none = `

resource "citrixadc_csvserver" "tf_csvserver1" {

  ipv46 = "192.168.43.66"
  name = "tf_csvserver1"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_csvserver" "tf_csvserver2" {

  ipv46 = "192.168.43.67"
  name = "tf_csvserver2"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_responderpolicy" "tf_responder_policy" {
  name    = "tf_responder_policy"
  action = "NOOP"
  rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"nosuchthing\")"
}
`
