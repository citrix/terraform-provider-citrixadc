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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

const testAccDataSourceHanode = `
data "citrixadc_hanode" "hanode" {
	hanode_id = 0
}
`

func TestAccDataSourceHanode_basic(t *testing.T) {
	nsip := os.Getenv("NSIP")
	if nsip == "" {
		nsurl := os.Getenv("NS_URL")
		// Try to extract IP from NS_URL (assuming format like http(s)://<ip>[:port])
		if nsurl != "" {
			// Remove protocol if present
			start := 0
			if idx := len("https://"); len(nsurl) > idx && nsurl[:idx] == "https://" {
				start = idx
			} else if idx := len("http://"); len(nsurl) > idx && nsurl[:idx] == "http://" {
				start = idx
			}
			hostport := nsurl[start:]
			// Split by '/' to remove any path
			if slashIdx := len(hostport); slashIdx > 0 {
				if idx := indexOf(hostport, "/"); idx != -1 {
					hostport = hostport[:idx]
				}
			}
			// Split by ':' to remove port if present
			if colonIdx := indexOf(hostport, ":"); colonIdx != -1 {
				nsip = hostport[:colonIdx]
			} else {
				nsip = hostport
			}
		}
	}
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceHanode,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_hanode.hanode", "ipaddress", nsip),
				),
			},
		},
	})
}

// Helper function to find index of a character in string
func indexOf(s string, sep string) int {
	for i := 0; i < len(s); i++ {
		if string(s[i]) == sep {
			return i
		}
	}
	return -1
}
