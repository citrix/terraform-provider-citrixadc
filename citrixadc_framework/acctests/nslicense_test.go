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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccNslicense_basic(t *testing.T) {
	t.Skip("ssh does not work correctly with CPX")
	if isCpxRun {
		t.Skip("ssh does not work correctly with CPX")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNslicense_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNslicenseExist("citrixadc_nslicense.tf_license", nil),
				),
			},
		},
	})
}

func testAccCheckNslicenseExist(n string, id *string) resource.TestCheckFunc {
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

		return nil
	}
}

const testAccNslicense_basic = `
resource "citrixadc_nslicense" "tf_license" {

    license_file = "CNS_V10000_SERVER_PLT_Retail.lic"
    ssh_host_pubkey = "AAAAB3NzaC1kc3MAAACBAJ2yRBTkiCIR94oYfCmabSKPi7EjyYS7FxLPV0j8zAsWtarS16UTiPyW+tTpd5I9HNJPqIkGKTLWw9DakXd+lyRnXAusOGXIfiV+wdyLn8hg/T/dZMA7r6QssIfIUza0Bqcjn7eCGStwOHzSgUH5qS9YmIlZuZ/hKOU3Zs0N7wF5AAAAFQC6iGTFqcADv83/ItKiHh+6pEWe2wAAAIB3pcTDxth1IASlwNgzm1HQYaOm5ttcGve468w3c97BpzEXCbiwObd6T8Ynt2GFMH2NHOFFqid4nRXvT2Ba5JlLYgyrDTU53J6eDxkXtBSuxTcMss0P6EEtcqOJzi1e+OZWFPKtxaIsKBtScBw+S/dNFkY4H+Eo5vl5/ChahdOchAAAAIBd4sHyDMVWzI6vG9Z/HYNM6los0fXqCL8ait+LpFN5+hOScdDKNgzIfM5md35ToV6cM28nPQL3bum3sLLO4R4o5Rqp3QFW82+mipswjzycNIgKy3gcSSuFA7ALivIsZUxqpyQYU7GyBKnkJsf5om0tcr7PawHL08CqJf0/mLXZcw=="
}
`

func TestAccNslicenseDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNslicenseDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.citrixadc_nslicense.test", "id"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nslicense.test", "licensingmode"),
					resource.TestCheckResourceAttrSet("data.citrixadc_nslicense.test", "modelid"),
					resource.TestCheckResourceAttr("data.citrixadc_nslicense.test", "lb", "true"),
					resource.TestCheckResourceAttr("data.citrixadc_nslicense.test", "ssl", "true"),
				),
			},
		},
	})
}

const testAccNslicenseDataSource_basic = `
data "citrixadc_nslicense" "test" {
}
`
