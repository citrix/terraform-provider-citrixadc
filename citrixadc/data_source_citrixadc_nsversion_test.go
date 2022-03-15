/*
Copyright 2020 Citrix Systems, Inc

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
	"testing"
	"regexp"

	"github.com/hashicorp/terraform/helper/resource"
)

const testAccNsversion = `
data "citrixadc_nsversion" "nsversion" {
	installedversion = true
}
`

func TestAccNsversion_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNsversion,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("data.citrixadc_nsversion.nsversion", "version", regexp.MustCompile("^Netscaler")),
				),
			},
		},
	})
}
