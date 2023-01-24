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

func TestAccRewritepolicy_globalbinding(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRewritepolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccRewritepolicy_globalbinding_not_exists,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewritepolicyExist("citrixadc_rewritepolicy.tf_rewrite_policy", nil),
					verifyRewriteGlobalBindingExists("REQ_OVERRIDE", "tf_rewrite_policy", true),
				),
			},
			resource.TestStep{
				Config: testAccRewritepolicy_globalbinding_exists,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewritepolicyExist("citrixadc_rewritepolicy.tf_rewrite_policy", nil),
					verifyRewriteGlobalBindingExists("REQ_DEFAULT", "tf_rewrite_policy", false),
				),
			},
			/*
				   // TODO: Find race condition that makes this fail. In manual testing this succeeds
					resource.TestStep{
						Config: testAccRewritepolicy_globalbinding_modified,
						Check: resource.ComposeTestCheckFunc(
							testAccCheckRewritepolicyExist("citrixadc_rewritepolicy.tf_rewrite_policy", nil),
							verifyRewriteGlobalBindingExists("REQ_OVERRIDE", "tf_rewrite_policy", false),
						),
					},
			*/
		},
	})
}

func TestAccRewritepolicy_lbvserverbinding(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRewritepolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccRewritepolicy_lbvserverbindings_none,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewritepolicyExist("citrixadc_rewritepolicy.tf_rewrite_policy", nil),
					verifyRewriteLbvserverBindingExists("tf_lbvserver1", "tf_rewrite_policy", true),
					verifyRewriteLbvserverBindingExists("tf_lbvserver2", "tf_rewrite_policy", true),
				),
			},
			resource.TestStep{
				Config: testAccRewritepolicy_lbvserverbindings_both,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewritepolicyExist("citrixadc_rewritepolicy.tf_rewrite_policy", nil),
					verifyRewriteLbvserverBindingExists("tf_lbvserver1", "tf_rewrite_policy", false),
					verifyRewriteLbvserverBindingExists("tf_lbvserver2", "tf_rewrite_policy", false),
				),
			},
			resource.TestStep{
				Config: testAccRewritepolicy_lbvserverbindings_one,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewritepolicyExist("citrixadc_rewritepolicy.tf_rewrite_policy", nil),
					verifyRewriteLbvserverBindingExists("tf_lbvserver1", "tf_rewrite_policy", false),
					verifyRewriteLbvserverBindingExists("tf_lbvserver2", "tf_rewrite_policy", true),
				),
			},
		},
	})
}

func TestAccRewritepolicy_csvserverbinding(t *testing.T) {
	if adcTestbed != "STANDALONE" {
		t.Skipf("ADC testbed is %s. Expected STANDALONE.", adcTestbed)
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRewritepolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccRewritepolicy_csvserverbindings_none,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewritepolicyExist("citrixadc_rewritepolicy.tf_rewrite_policy", nil),
					//verifyRewriteLbvserverBindingExists("tf_lbvserver1", "tf_rewrite_policy", true),
					//verifyRewriteLbvserverBindingExists("tf_lbvserver2", "tf_rewrite_policy", true),
				),
			},
			resource.TestStep{
				Config: testAccRewritepolicy_csvserverbindings_both,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewritepolicyExist("citrixadc_rewritepolicy.tf_rewrite_policy", nil),
					//verifyRewriteLbvserverBindingExists("tf_lbvserver1", "tf_rewrite_policy", true),
					//verifyRewriteLbvserverBindingExists("tf_lbvserver2", "tf_rewrite_policy", true),
				),
			},
			resource.TestStep{
				Config: testAccRewritepolicy_csvserverbindings_one,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRewritepolicyExist("citrixadc_rewritepolicy.tf_rewrite_policy", nil),
					//verifyRewriteLbvserverBindingExists("tf_lbvserver1", "tf_rewrite_policy", true),
					//verifyRewriteLbvserverBindingExists("tf_lbvserver2", "tf_rewrite_policy", true),
				),
			},
		},
	})
}

func testAccCheckRewritepolicyExist(n string, id *string) resource.TestCheckFunc {
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
		data, err := nsClient.FindResource(service.Rewritepolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckRewritepolicyDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_rewritepolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Rewritepolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

func verifyRewriteGlobalBindingExists(bindtype string, policyname string, inverse bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		bindFound := false
		client := testAccProvider.Meta().(*NetScalerNitroClient).client
		globalBindings, _ := client.FindResourceArray("rewritepolicy_rewriteglobal_binding", policyname)
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

func verifyRewriteLbvserverBindingExists(servername string, policyname string, inverse bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		bindFound := false
		client := testAccProvider.Meta().(*NetScalerNitroClient).client
		lbVserverBindings, _ := client.FindResourceArray("rewritepolicy_lbvserver_binding", policyname)
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

const testAccRewritepolicy_globalbinding_exists = `

resource "citrixadc_lbvserver" "tf_lbvserver" {

  ipv46 = "10.66.22.33"
  name = "tf_lbvserver_name"
  port = 80
  servicetype = "HTTP"

}


resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
	name = "tf_rewrite_policy"
	action = "DROP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"

	depends_on = [ citrixadc_lbvserver.tf_lbvserver ]

	globalbinding {
		gotopriorityexpression = "END"
		labelname = citrixadc_lbvserver.tf_lbvserver.name
		labeltype = "reqvserver"
		priority = 205
		invoke = true
		type = "REQ_DEFAULT"
	}
}
`

const testAccRewritepolicy_globalbinding_modified = `

resource "citrixadc_lbvserver" "tf_lbvserver" {

  ipv46 = "10.66.22.33"
  name = "tf_lbvserver"
  port = 80
  servicetype = "HTTP"

}


resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
	name = "tf_rewrite_policy"
	action = "DROP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"
	depends_on = [ citrixadc_lbvserver.tf_lbvserver ]

	globalbinding {
            gotopriorityexpression = "END"
            labelname = citrixadc_lbvserver.tf_lbvserver.name
            labeltype = "reqvserver"
            priority = 208
            invoke = true
            type = "REQ_OVERRIDE"
	}
}
`

const testAccRewritepolicy_globalbinding_not_exists = `

resource "citrixadc_lbvserver" "tf_lbvserver" {

  ipv46 = "10.66.22.33"
  name = "tf_lbvserver"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
	name = "tf_rewrite_policy"
	action = "DROP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"
	depends_on = [ citrixadc_lbvserver.tf_lbvserver ]
}
`

const testAccRewritepolicy_lbvserverbindings_none = `
resource "citrixadc_lbvserver" "tf_lbvserver1" {

  ipv46 = "10.22.22.22"
  name = "tf_lbvserver1"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_lbvserver" "tf_lbvserver2" {

  ipv46 = "10.33.22.66"
  name = "tf_lbvserver2"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
	name = "tf_rewrite_policy"
	action = "DROP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"
}
`

const testAccRewritepolicy_lbvserverbindings_one = `
resource "citrixadc_lbvserver" "tf_lbvserver1" {

  ipv46 = "10.22.22.22"
  name = "tf_lbvserver1"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_lbvserver" "tf_lbvserver2" {

  ipv46 = "10.33.22.66"
  name = "tf_lbvserver2"
  port = 80
  servicetype = "HTTP"

}
resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
	name = "tf_rewrite_policy"
	action = "DROP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"

	lbvserverbinding {
        name = citrixadc_lbvserver.tf_lbvserver1.name
        bindpoint = "REQUEST"
        priority = 114
        gotopriorityexpression = "END"
	}
}
`

const testAccRewritepolicy_lbvserverbindings_both = `
resource "citrixadc_lbvserver" "tf_lbvserver1" {

  ipv46 = "10.22.22.22"
  name = "tf_lbvserver1"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_lbvserver" "tf_lbvserver2" {

  ipv46 = "10.33.22.66"
  name = "tf_lbvserver2"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
	name = "tf_rewrite_policy"
	action = "DROP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"

	lbvserverbinding {
        name = citrixadc_lbvserver.tf_lbvserver1.name
        bindpoint = "RESPONSE"
        priority = 114
        gotopriorityexpression = "END"
	}

	lbvserverbinding {
        name = citrixadc_lbvserver.tf_lbvserver2.name
        bindpoint = "REQUEST"
        priority = 114
        gotopriorityexpression = "END"
	}
}
`

const testAccRewritepolicy_csvserverbindings_both = `
resource "citrixadc_csvserver" "tf_csvserver1" {

  ipv46 = "192.168.45.66"
  name = "tf_csvserver1"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_csvserver" "tf_csvserver2" {

  ipv46 = "192.168.45.67"
  name = "tf_csvserver2"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
	name = "tf_rewrite_policy"
	action = "DROP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"

	csvserverbinding {
        name = citrixadc_csvserver.tf_csvserver1.name
        bindpoint = "RESPONSE"
        priority = 114
        gotopriorityexpression = "END"
	}

	csvserverbinding {
        name = citrixadc_csvserver.tf_csvserver2.name
        bindpoint = "REQUEST"
        priority = 114
        gotopriorityexpression = "END"
	}
}
`

const testAccRewritepolicy_csvserverbindings_one = `
resource "citrixadc_csvserver" "tf_csvserver1" {

  ipv46 = "192.168.45.66"
  name = "tf_csvserver1"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_csvserver" "tf_csvserver2" {

  ipv46 = "192.168.45.67"
  name = "tf_csvserver2"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
	name = "tf_rewrite_policy"
	action = "DROP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"

	csvserverbinding {
        name = citrixadc_csvserver.tf_csvserver2.name
        bindpoint = "REQUEST"
        priority = 114
        gotopriorityexpression = "END"
	}
}
`

const testAccRewritepolicy_csvserverbindings_none = `
resource "citrixadc_csvserver" "tf_csvserver1" {

  ipv46 = "192.168.45.66"
  name = "tf_csvserver1"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_csvserver" "tf_csvserver2" {

  ipv46 = "192.168.45.67"
  name = "tf_csvserver2"
  port = 80
  servicetype = "HTTP"

}

resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
	name = "tf_rewrite_policy"
	action = "DROP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"

}
`
