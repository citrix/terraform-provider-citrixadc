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
	// "os"
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/resource/config/basic"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

func TestAccServicegroup_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckServicegroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServicegroup_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroupExist("citrixadc_servicegroup.foo", nil),

					resource.TestCheckResourceAttr(
						"citrixadc_servicegroup.foo", "servicegroupname", "test_servicegroup"),
					resource.TestCheckResourceAttr(
						"citrixadc_servicegroup.foo", "servicetype", "HTTP"),
				),
			},
		},
	})
}

func testAccCheckServicegroupExist(n string, id *string) resource.TestCheckFunc {
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

		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Servicegroup.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("LB vserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckServicegroupDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_servicegroup" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Servicegroup.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("LB vserver %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

// TODO add testcase when we have servicegroupmembers_by_servername defined
const testAccServicegroup_basic = `

resource "citrixadc_lbvserver" "foo1" {

  name = "foo_lb_1"
  ipv46 = "10.202.11.11"
  port = 80
  servicetype = "HTTP"
}

resource "citrixadc_lbvserver" "foo2" {

  name = "foo_lb_2"
  ipv46 = "10.202.11.12"
  port = 80
  servicetype = "HTTP"
}


resource "citrixadc_servicegroup" "foo" {

  servicegroupname = "test_servicegroup"
  servicetype = "HTTP"
  servicegroupmembers = ["172.20.0.9:80:10", "172.20.0.10:80:10", "172.20.0.11:8080:20"]
  lbvservers = ["foo_lb_1", "foo_lb_2"]
  depends_on = ["citrixadc_lbvserver.foo1", "citrixadc_lbvserver.foo2"]
}
`

func TestAccServicegroup_AssertNonUpdateableAttributes(t *testing.T) {

	// if tfAcc := os.Getenv("TF_ACC"); tfAcc == "" {
	// 	t.Skip("TF_ACC not set. Skipping acceptance test.")
	// }

	c, err := testHelperInstantiateClient("", "", "", false)
	if err != nil {
		t.Fatalf("Failed to instantiate client. %v\n", err)
	}

	// Create resource
	servicegroupName := "tf-acc-servicegroup-test"
	servicegroupType := service.Servicegroup.Type()

	// Defer deletion of actual resource
	defer testHelperEnsureResourceDeletion(c, t, servicegroupType, servicegroupName, nil)

	servicegroupInstance := basic.Servicegroup{
		Servicegroupname: servicegroupName,
		Servicetype:      "HTTP",
	}

	if _, err := c.client.AddResource(servicegroupType, servicegroupName, servicegroupInstance); err != nil {
		t.Logf("Error while creating resource")
		t.Fatal(err)
	}

	//servicetype
	servicegroupInstance.Servicetype = "HTTP"
	testHelperVerifyImmutabilityFunc(c, t, servicegroupType, servicegroupName, servicegroupInstance, "servicetype")
	servicegroupInstance.Servicetype = ""

	//cachetype
	servicegroupInstance.Cachetype = "TRANSPARENT"
	testHelperVerifyImmutabilityFunc(c, t, servicegroupType, servicegroupName, servicegroupInstance, "cachetype")
	servicegroupInstance.Cachetype = ""

	servicegroupInstance = basic.Servicegroup{
		Servicegroupname: servicegroupName,
	}

	//td
	servicegroupInstance.Td = utils.IntPtr(2)
	testHelperVerifyImmutabilityFunc(c, t, servicegroupType, servicegroupName, servicegroupInstance, "td")

	servicegroupInstance = basic.Servicegroup{
		Servicegroupname: servicegroupName,
	}

	//memberport
	servicegroupInstance.Memberport = utils.IntPtr(80)
	testHelperVerifyImmutabilityFunc(c, t, servicegroupType, servicegroupName, servicegroupInstance, "memberport")

	servicegroupInstance = basic.Servicegroup{
		Servicegroupname: servicegroupName,
	}
	//includemembers
	servicegroupInstance.Includemembers = true
	testHelperVerifyImmutabilityFunc(c, t, servicegroupType, servicegroupName, servicegroupInstance, "includemembers")
	servicegroupInstance.Includemembers = false

}

const testAccServicegroupEnableDisable_enabled = `
resource "citrixadc_servicegroup" "tf_enable_disable_test_svcgroup" {
	servicegroupname = "tf_enable_disable_test_svcgroup"
    servicetype = "HTTP"
	servicegroupmembers = []
	comment = "enabled state comment"
	state = "ENABLED"
	graceful = "YES"
	delay = 60
}
`

const testAccServicegroupEnableDisable_disabled = `
resource "citrixadc_servicegroup" "tf_enable_disable_test_svcgroup" {
	servicegroupname = "tf_enable_disable_test_svcgroup"
    servicetype = "HTTP"
	servicegroupmembers = []
	comment = "disabled state comment"
	state = "DISABLED"
	graceful = "YES"
	delay = 60
}
`

func TestAccServicegroup_enable_disable(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckServicegroupDestroy,
		Steps: []resource.TestStep{
			// Create enabled
			{
				Config: testAccServicegroupEnableDisable_enabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroupExist("citrixadc_servicegroup.tf_enable_disable_test_svcgroup", nil),
					resource.TestCheckResourceAttr("citrixadc_servicegroup.tf_enable_disable_test_svcgroup", "state", "ENABLED"),
				),
			},
			// Disable
			{
				Config: testAccServicegroupEnableDisable_disabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroupExist("citrixadc_servicegroup.tf_enable_disable_test_svcgroup", nil),
					resource.TestCheckResourceAttr("citrixadc_servicegroup.tf_enable_disable_test_svcgroup", "state", "DISABLED"),
				),
			},
			// Re enable
			{
				Config: testAccServicegroupEnableDisable_enabled,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroupExist("citrixadc_servicegroup.tf_enable_disable_test_svcgroup", nil),
					resource.TestCheckResourceAttr("citrixadc_servicegroup.tf_enable_disable_test_svcgroup", "state", "ENABLED"),
				),
			},
		},
	})
}

const testAccServicegroupDataSource_basic = `

	resource "citrixadc_servicegroup" "tf_servicegroup" {
		servicegroupname = "test_servicegroup_ds"
		servicetype      = "HTTP"
	}

	data "citrixadc_servicegroup" "tf_servicegroup" {
		servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
		depends_on       = [citrixadc_servicegroup.tf_servicegroup]
	}
`

func TestAccServicegroup_selfHealing(t *testing.T) {
	const resAddr = "citrixadc_servicegroup.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckServicegroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServicegroup_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckServicegroupExist(resAddr, nil)),
			},
			{
				PreConfig: func() {
					client, err := testAccGetFrameworkClient()
					if err != nil {
						t.Fatalf("self-healing: client: %v", err)
					}
					if err := client.DeleteResource(service.Servicegroup.Type(), "test_servicegroup"); err != nil {
						t.Fatalf("self-healing: out-of-band delete failed: %v", err)
					}
				},
				Config: testAccServicegroup_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckServicegroupExist(resAddr, nil)),
			},
		},
	})
}

func TestAccServicegroup_import(t *testing.T) {
	const resAddr = "citrixadc_servicegroup.foo"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckServicegroupDestroy,
		Steps: []resource.TestStep{
			{Config: testAccServicegroup_basic},
			{
				Config:                  testAccServicegroup_basic,
				ResourceName:            resAddr,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"delay", "graceful", "lbvservers", "servicegroupmembers"},
			},
		},
	})
}

func TestAccServicegroup_sdkv2StateUpgrade(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckServicegroupDestroy,
		Steps: []resource.TestStep{
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {Source: "citrix/citrixadc", VersionConstraint: "2.2.0"},
				},
				Config: testAccServicegroup_basic,
				Check:  resource.ComposeTestCheckFunc(testAccCheckServicegroupExist("citrixadc_servicegroup.foo", nil)),
			},
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccServicegroup_basic,
				Check:                    resource.ComposeTestCheckFunc(testAccCheckServicegroupExist("citrixadc_servicegroup.foo", nil)),
			},
		},
	})
}

// The servicegroup unset test covers the mutable attributes wired into
// attributesToUnset. step1 applies non-default values; step2 removes them from
// config and the provider must issue ?action=unset, reverting them to the
// documented NITRO defaults.
const testAccServicegroup_unset_step1 = `
resource "citrixadc_servicegroup" "tf_unset" {
	servicegroupname   = "tf_servicegroup_unset"
	servicetype        = "HTTP"
	appflowlog         = "DISABLED"
	cacheable          = "YES"
	downstateflush     = "DISABLED"
	healthmonitor      = "NO"
	monconnectionclose = "RESET"
	sp                 = "ON"
}
`

const testAccServicegroup_unset_step2 = `
resource "citrixadc_servicegroup" "tf_unset" {
	servicegroupname = "tf_servicegroup_unset"
	servicetype      = "HTTP"
	# All unset-eligible attributes removed -> provider must unset them (revert to
	# NITRO defaults).
}
`

func TestAccServicegroup_unset(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckServicegroupDestroy,
		Steps: []resource.TestStep{
			{
				// Non-default values are applied and persisted.
				Config: testAccServicegroup_unset_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroupExist("citrixadc_servicegroup.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_servicegroup.tf_unset", "appflowlog", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_servicegroup.tf_unset", "cacheable", "YES"),
					resource.TestCheckResourceAttr("citrixadc_servicegroup.tf_unset", "downstateflush", "DISABLED"),
					resource.TestCheckResourceAttr("citrixadc_servicegroup.tf_unset", "healthmonitor", "NO"),
					resource.TestCheckResourceAttr("citrixadc_servicegroup.tf_unset", "monconnectionclose", "RESET"),
					resource.TestCheckResourceAttr("citrixadc_servicegroup.tf_unset", "sp", "ON"),
				),
			},
			{
				// Removing the attributes must unset them: state reverts to the
				// documented NITRO defaults and the implicit post-apply plan is empty.
				Config: testAccServicegroup_unset_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicegroupExist("citrixadc_servicegroup.tf_unset", nil),
					resource.TestCheckResourceAttr("citrixadc_servicegroup.tf_unset", "appflowlog", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_servicegroup.tf_unset", "cacheable", "NO"),
					resource.TestCheckResourceAttr("citrixadc_servicegroup.tf_unset", "downstateflush", "ENABLED"),
					resource.TestCheckResourceAttr("citrixadc_servicegroup.tf_unset", "healthmonitor", "YES"),
					resource.TestCheckResourceAttr("citrixadc_servicegroup.tf_unset", "monconnectionclose", "NONE"),
					resource.TestCheckResourceAttr("citrixadc_servicegroup.tf_unset", "sp", "OFF"),
					// Independent appliance-level confirmation the unset took effect.
					testAccCheckServicegroupADCValue("tf_servicegroup_unset", "appflowlog", "ENABLED"),
					testAccCheckServicegroupADCValue("tf_servicegroup_unset", "healthmonitor", "YES"),
					testAccCheckServicegroupADCValue("tf_servicegroup_unset", "cacheable", "NO"),
				),
			},
		},
	})
}

// testAccCheckServicegroupADCValue asserts an attribute's value directly on the
// appliance (not just in Terraform state), proving the unset actually reverted it.
func testAccCheckServicegroupADCValue(name, attr, want string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}
		data, err := client.FindResource(service.Servicegroup.Type(), name)
		if err != nil {
			return err
		}
		if data == nil {
			return fmt.Errorf("servicegroup %s not found on appliance", name)
		}
		got := strings.TrimSpace(fmt.Sprintf("%v", data[attr]))
		if got != want {
			return fmt.Errorf("servicegroup %s: appliance attr %q = %q, want %q (unset did not revert it)", name, attr, got, want)
		}
		return nil
	}
}

func TestAccServicegroupDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckServicegroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccServicegroupDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_servicegroup.tf_servicegroup", "servicegroupname", "test_servicegroup_ds"),
					resource.TestCheckResourceAttr("data.citrixadc_servicegroup.tf_servicegroup", "servicetype", "HTTP"),
				),
			},
		},
	})
}
