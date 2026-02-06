
package acctests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccVpnvserverAppfwpolicyBinding_basic = `

	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name = "tf_appfwprofile"
	}
	resource "citrixadc_appfwpolicy" "tf_appfwpolicy" {
		name        = "tf_appfwpolicy"
		profilename = citrixadc_appfwprofile.tf_appfwprofile.name
		rule        = "true"
	}
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name        = "tf_vpnvserver"
		servicetype = "SSL"
		ipv46       = "0.0.0.0"
		port        = 0
	}
	resource "citrixadc_vpnvserver_appfwpolicy_binding" "tf_bind" {
		name      = citrixadc_vpnvserver.tf_vpnvserver.name
		policy    = citrixadc_appfwpolicy.tf_appfwpolicy.name
		priority  = 100
		gotopriorityexpression = "END"
	}
`

const testAccVpnvserverAppfwpolicyBinding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion

	resource "citrixadc_appfwprofile" "tf_appfwprofile" {
		name = "tf_appfwprofile"
	}
	resource "citrixadc_appfwpolicy" "tf_appfwpolicy" {
		name        = "tf_appfwpolicy"
		profilename = citrixadc_appfwprofile.tf_appfwprofile.name
		rule        = "true"
	}
	resource "citrixadc_vpnvserver" "tf_vpnvserver" {
		name        = "tf_vpnvserver"
		servicetype = "SSL"
		ipv46       = "0.0.0.0"
		port        = 0
	}
`

func TestAccVpnvserverAppfwpolicyBinding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVpnvserverAppfwpolicyBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpnvserverAppfwpolicyBinding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserverAppfwpolicyBindingExist("citrixadc_vpnvserver_appfwpolicy_binding.tf_bind", nil),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver_appfwpolicy_binding.tf_bind", "name", "tf_vpnvserver"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver_appfwpolicy_binding.tf_bind", "policy", "tf_appfwpolicy"),
					resource.TestCheckResourceAttr("citrixadc_vpnvserver_appfwpolicy_binding.tf_bind", "priority", "100"),
				),
			},
			{
				Config: testAccVpnvserverAppfwpolicyBinding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpnvserverAppfwpolicyBindingNotExist("citrixadc_vpnvserver_appfwpolicy_binding.tf_bind", "tf_vpnvserver,tf_appfwpolicy"),
				),
			},
		},
	})
}

func testAccCheckVpnvserverAppfwpolicyBindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No vpnvserver_appfwpolicy_binding id is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		// Get a configured client from the test helper
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		bindingId := rs.Primary.ID

		idSlice := strings.SplitN(bindingId, ",", 2)

		name := idSlice[0]
		policy := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "vpnvserver_appfwpolicy_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching policy
		found := false
		for _, v := range dataArr {
			if v["policy"].(string) == policy {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("vpnvserver_appfwpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckVpnvserverAppfwpolicyBindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Get a configured client from the test helper
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		name := idSlice[0]
		policy := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "vpnvserver_appfwpolicy_binding",
			ResourceName:             name,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching policy
		found := false
		for _, v := range dataArr {
			if v["policy"].(string) == policy {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("vpnvserver_appfwpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckVpnvserverAppfwpolicyBindingDestroy(s *terraform.State) error {
	// Get a configured client from the test helper
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_vpnvserver_appfwpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Vpnvserver_appfwpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("vpnvserver_appfwpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
