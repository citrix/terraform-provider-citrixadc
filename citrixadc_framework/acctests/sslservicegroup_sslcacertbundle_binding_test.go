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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// sslservicegroup_sslcacertbundle_binding joins an SSL servicegroup with a CA cert
// bundle. The bundle's bundlefile must ALREADY EXIST on the appliance under
// /nsconfig/ssl/ (see sslcacertbundle_test.go). The bundlefile value is a
// TODO_PLACEHOLDER - replace "ns-root.pem" with a real CA-cert-bundle file present
// on your testbed.

const testAccSslservicegroup_sslcacertbundle_binding_basic = `
	resource "citrixadc_servicegroup" "tf_servicegroup" {
		servicegroupname = "tf_servicegroup"
		servicetype      = "SSL"
	}

	resource "citrixadc_sslcacertbundle" "tf_sslcacertbundle" {
		cacertbundlename = "tf_sslcacertbundle"
		bundlefile       = "ca-bundle.pem" // TODO_PLACEHOLDER: CA-cert-bundle file must pre-exist on the appliance under /nsconfig/ssl/
	}

	resource "citrixadc_sslservicegroup_sslcacertbundle_binding" "tf_sslservicegroup_sslcacertbundle_binding" {
		cacertbundlename = citrixadc_sslcacertbundle.tf_sslcacertbundle.cacertbundlename
		servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
		depends_on       = [citrixadc_servicegroup.tf_servicegroup, citrixadc_sslcacertbundle.tf_sslcacertbundle]
	}
`

const testAccSslservicegroup_sslcacertbundle_binding_basic_step2 = `
	resource "citrixadc_servicegroup" "tf_servicegroup" {
		servicegroupname = "tf_servicegroup"
		servicetype      = "SSL"
	}

	resource "citrixadc_sslcacertbundle" "tf_sslcacertbundle" {
		cacertbundlename = "tf_sslcacertbundle"
		bundlefile       = "ca-bundle.pem" // TODO_PLACEHOLDER: CA-cert-bundle file must pre-exist on the appliance under /nsconfig/ssl/
	}
`

func TestAccSslservicegroup_sslcacertbundle_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcacertbundlePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslservicegroup_sslcacertbundle_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslservicegroup_sslcacertbundle_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslservicegroup_sslcacertbundle_bindingExist("citrixadc_sslservicegroup_sslcacertbundle_binding.tf_sslservicegroup_sslcacertbundle_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup_sslcacertbundle_binding.tf_sslservicegroup_sslcacertbundle_binding", "cacertbundlename", "tf_sslcacertbundle"),
					resource.TestCheckResourceAttr("citrixadc_sslservicegroup_sslcacertbundle_binding.tf_sslservicegroup_sslcacertbundle_binding", "servicegroupname", "tf_servicegroup"),
				),
			},
			{
				Config: testAccSslservicegroup_sslcacertbundle_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslservicegroup_sslcacertbundle_bindingNotExist("citrixadc_sslservicegroup_sslcacertbundle_binding.tf_sslservicegroup_sslcacertbundle_binding", "tf_servicegroup,tf_sslcacertbundle"),
				),
			},
		},
	})
}

func testAccCheckSslservicegroup_sslcacertbundle_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslservicegroup_sslcacertbundle_binding id is set")
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

		servicegroupname := rs.Primary.Attributes["servicegroupname"]
		cacertbundlename := rs.Primary.Attributes["cacertbundlename"]

		findParams := service.FindParams{
			ResourceType:             "sslservicegroup_sslcacertbundle_binding",
			ResourceName:             servicegroupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching cacertbundlename
		found := false
		for _, v := range dataArr {
			if v["cacertbundlename"].(string) == cacertbundlename {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("sslservicegroup_sslcacertbundle_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckSslservicegroup_sslcacertbundle_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		if !strings.Contains(id, ",") {
			return fmt.Errorf("Invalid id string %v. The id string must contain a comma.", id)
		}
		idSlice := strings.SplitN(id, ",", 2)

		servicegroupname := idSlice[0]
		cacertbundlename := idSlice[1]

		findParams := service.FindParams{
			ResourceType:             "sslservicegroup_sslcacertbundle_binding",
			ResourceName:             servicegroupname,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the matching cacertbundlename
		found := false
		for _, v := range dataArr {
			if v["cacertbundlename"].(string) == cacertbundlename {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("sslservicegroup_sslcacertbundle_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckSslservicegroup_sslcacertbundle_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslservicegroup_sslcacertbundle_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Sslservicegroup_sslcacertbundle_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("sslservicegroup_sslcacertbundle_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccSslservicegroup_sslcacertbundle_bindingDataSource_basic = `
	resource "citrixadc_servicegroup" "tf_servicegroup" {
		servicegroupname = "tf_servicegroup"
		servicetype      = "SSL"
	}

	resource "citrixadc_sslcacertbundle" "tf_sslcacertbundle" {
		cacertbundlename = "tf_sslcacertbundle"
		bundlefile       = "ca-bundle.pem" // CA-cert-bundle file is uploaded to /nsconfig/ssl/ by doSslcacertbundlePreChecks
	}

	resource "citrixadc_sslservicegroup_sslcacertbundle_binding" "tf_sslservicegroup_sslcacertbundle_binding" {
		cacertbundlename = citrixadc_sslcacertbundle.tf_sslcacertbundle.cacertbundlename
		servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
		depends_on       = [citrixadc_servicegroup.tf_servicegroup, citrixadc_sslcacertbundle.tf_sslcacertbundle]
	}

	data "citrixadc_sslservicegroup_sslcacertbundle_binding" "tf_sslservicegroup_sslcacertbundle_binding_ds" {
		servicegroupname = citrixadc_sslservicegroup_sslcacertbundle_binding.tf_sslservicegroup_sslcacertbundle_binding.servicegroupname
		cacertbundlename = citrixadc_sslservicegroup_sslcacertbundle_binding.tf_sslservicegroup_sslcacertbundle_binding.cacertbundlename
		depends_on       = [citrixadc_sslservicegroup_sslcacertbundle_binding.tf_sslservicegroup_sslcacertbundle_binding]
	}
`

func TestAccSslservicegroup_sslcacertbundle_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcacertbundlePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSslservicegroup_sslcacertbundle_bindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslservicegroup_sslcacertbundle_binding.tf_sslservicegroup_sslcacertbundle_binding_ds", "servicegroupname", "tf_servicegroup"),
					resource.TestCheckResourceAttr("data.citrixadc_sslservicegroup_sslcacertbundle_binding.tf_sslservicegroup_sslcacertbundle_binding_ds", "cacertbundlename", "tf_sslcacertbundle"),
				),
			},
		},
	})
}
