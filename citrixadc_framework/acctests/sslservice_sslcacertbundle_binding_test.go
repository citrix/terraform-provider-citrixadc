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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// sslservice_sslcacertbundle_binding joins an SSL service (servicename) with an
// SSL CA cert bundle (cacertbundlename). Participating entities:
//   - citrixadc_service of servicetype SSL (the SSL service is the SSL view of that service)
//   - citrixadc_sslcacertbundle (config lifted from sslcacertbundle_test.go)
//
// NOTE: sslcacertbundle requires a CA-cert-bundle file under /nsconfig/ssl/. The
// PreCheck doSslcacertbundlePreChecks(t) uploads testdata/ca-bundle.pem there.
//
// Composite ID = cacertbundlename:<v>,servicename:<v>. The exist/destroy checks read
// the binding array for the servicename and match on cacertbundlename.

// step1: create the binding
const testAccSslserviceSslcacertbundleBinding_basic_step1 = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_sslsvc_cacert_lb"
  ipv46       = "10.33.55.34"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_service" "tf_service" {
  servicetype = "SSL"
  name        = "tf_sslsvc_cacert"
  ipaddress   = "10.77.33.23"
  ip          = "10.77.33.23"
  port        = "443"
  lbvserver   = citrixadc_lbvserver.tf_lbvserver.name
}

resource "citrixadc_sslcacertbundle" "tf_sslcacertbundle" {
  cacertbundlename = "tf_sslsvc_cacertbundle"
  bundlefile       = "ca-bundle.pem" // CA-cert-bundle file is uploaded to /nsconfig/ssl/ by doSslcacertbundlePreChecks
}

resource "citrixadc_sslservice_sslcacertbundle_binding" "tf_binding" {
  servicename      = citrixadc_service.tf_service.name
  cacertbundlename = citrixadc_sslcacertbundle.tf_sslcacertbundle.cacertbundlename

  depends_on = [citrixadc_service.tf_service, citrixadc_sslcacertbundle.tf_sslcacertbundle]
}
`

// step2: drop the binding (entities remain)
const testAccSslserviceSslcacertbundleBinding_basic_step2 = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_sslsvc_cacert_lb"
  ipv46       = "10.33.55.34"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_service" "tf_service" {
  servicetype = "SSL"
  name        = "tf_sslsvc_cacert"
  ipaddress   = "10.77.33.23"
  ip          = "10.77.33.23"
  port        = "443"
  lbvserver   = citrixadc_lbvserver.tf_lbvserver.name
}

resource "citrixadc_sslcacertbundle" "tf_sslcacertbundle" {
  cacertbundlename = "tf_sslsvc_cacertbundle"
  bundlefile       = "ca-bundle.pem" // CA-cert-bundle file is uploaded to /nsconfig/ssl/ by doSslcacertbundlePreChecks
}
`

func TestAccSslserviceSslcacertbundleBinding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcacertbundlePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslserviceSslcacertbundleBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslserviceSslcacertbundleBinding_basic_step1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslserviceSslcacertbundleBindingExist("citrixadc_sslservice_sslcacertbundle_binding.tf_binding", nil),
					resource.TestCheckResourceAttr("citrixadc_sslservice_sslcacertbundle_binding.tf_binding", "servicename", "tf_sslsvc_cacert"),
					resource.TestCheckResourceAttr("citrixadc_sslservice_sslcacertbundle_binding.tf_binding", "cacertbundlename", "tf_sslsvc_cacertbundle"),
				),
			},
			{
				// Binding dropped; verify it no longer exists on the ADC.
				Config: testAccSslserviceSslcacertbundleBinding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSslserviceSslcacertbundleBindingNotExist("tf_sslsvc_cacert", "tf_sslsvc_cacertbundle"),
				),
			},
		},
	})
}

func testAccCheckSslserviceSslcacertbundleBindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No sslservice_sslcacertbundle_binding ID is set")
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

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}
		servicename := idMap["servicename"]
		cacertbundlename := idMap["cacertbundlename"]

		findParams := service.FindParams{
			ResourceType: service.Sslservice_sslcacertbundle_binding.Type(),
			ResourceName: servicename,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if val, ok := v["cacertbundlename"].(string); ok && val == cacertbundlename {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("sslservice_sslcacertbundle_binding %s not found on ADC", rs.Primary.ID)
		}

		return nil
	}
}

// testAccCheckSslserviceSslcacertbundleBindingNotExist verifies the binding is gone
// while the parent service still exists (step2 keeps the service).
func testAccCheckSslserviceSslcacertbundleBindingNotExist(servicename, cacertbundlename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		findParams := service.FindParams{
			ResourceType:             service.Sslservice_sslcacertbundle_binding.Type(),
			ResourceName:             servicename,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Service may have no bindings at all - acceptable.
			return nil
		}
		for _, v := range dataArr {
			if val, ok := v["cacertbundlename"].(string); ok && val == cacertbundlename {
				return fmt.Errorf("sslservice_sslcacertbundle_binding for %s/%s still exists", servicename, cacertbundlename)
			}
		}
		return nil
	}
}

func testAccCheckSslserviceSslcacertbundleBindingDestroy(s *terraform.State) error {
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_sslservice_sslcacertbundle_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		idMap, _, err := utils.ParseIdString(rs.Primary.ID, nil, nil)
		if err != nil {
			return fmt.Errorf("Error parsing ID %s: %v", rs.Primary.ID, err)
		}
		servicename := idMap["servicename"]
		cacertbundlename := idMap["cacertbundlename"]

		findParams := service.FindParams{
			ResourceType:             service.Sslservice_sslcacertbundle_binding.Type(),
			ResourceName:             servicename,
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)
		if err != nil {
			// Parent gone or no bindings - binding is destroyed.
			continue
		}
		for _, v := range dataArr {
			if val, ok := v["cacertbundlename"].(string); ok && val == cacertbundlename {
				return fmt.Errorf("sslservice_sslcacertbundle_binding %s still exists", rs.Primary.ID)
			}
		}
	}

	return nil
}

const testAccSslserviceSslcacertbundleBindingDataSource_basic = `
resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_sslsvc_cacert_lb"
  ipv46       = "10.33.55.34"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_service" "tf_service" {
  servicetype = "SSL"
  name        = "tf_sslsvc_cacert"
  ipaddress   = "10.77.33.23"
  ip          = "10.77.33.23"
  port        = "443"
  lbvserver   = citrixadc_lbvserver.tf_lbvserver.name
}

resource "citrixadc_sslcacertbundle" "tf_sslcacertbundle" {
  cacertbundlename = "tf_sslsvc_cacertbundle"
  bundlefile       = "ca-bundle.pem" // CA-cert-bundle file is uploaded to /nsconfig/ssl/ by doSslcacertbundlePreChecks
}

resource "citrixadc_sslservice_sslcacertbundle_binding" "tf_binding" {
  servicename      = citrixadc_service.tf_service.name
  cacertbundlename = citrixadc_sslcacertbundle.tf_sslcacertbundle.cacertbundlename

  depends_on = [citrixadc_service.tf_service, citrixadc_sslcacertbundle.tf_sslcacertbundle]
}

data "citrixadc_sslservice_sslcacertbundle_binding" "tf_binding" {
  servicename      = citrixadc_sslservice_sslcacertbundle_binding.tf_binding.servicename
  cacertbundlename = citrixadc_sslservice_sslcacertbundle_binding.tf_binding.cacertbundlename
  depends_on       = [citrixadc_sslservice_sslcacertbundle_binding.tf_binding]
}
`

func TestAccSslserviceSslcacertbundleBindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { doSslcacertbundlePreChecks(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSslserviceSslcacertbundleBindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSslserviceSslcacertbundleBindingDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_sslservice_sslcacertbundle_binding.tf_binding", "servicename", "tf_sslsvc_cacert"),
					resource.TestCheckResourceAttr("data.citrixadc_sslservice_sslcacertbundle_binding.tf_binding", "cacertbundlename", "tf_sslsvc_cacertbundle"),
				),
			},
		},
	})
}
