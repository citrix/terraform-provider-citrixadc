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
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// TestAccNSLASLicenseOffline_missingHostPubKey verifies that ssh_host_pubkey is a
// Required argument on citrixadc_nslaslicense_offline (CTXMYT-2537). Omitting it
// must fail at plan time, which guarantees the SCP license transfer can never
// fall back to an unverified SSH host key (ssh.InsecureIgnoreHostKey). The error
// is raised by Terraform core during config validation, before any ADC/SCP I/O,
// so this step needs no reachable appliance and is safe to run everywhere.
func TestAccNSLASLicenseOffline_missingHostPubKey(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccNSLASLicenseOffline_missingHostPubKey,
				ExpectError: regexp.MustCompile(`The argument "ssh_host_pubkey" is required|Missing required argument`),
			},
		},
	})
}

const testAccNSLASLicenseOffline_missingHostPubKey = `
resource "citrixadc_nslaslicense_offline" "tf_las_offline" {
  entitlement_name = "VPX 10000 Premium"
  las_secrets_json = "/tmp/las_secrets.json"
}
`

// TestAccNSLASLicenseOffline_wrongHostPubKey verifies that supplying a
// well-formed but INCORRECT ssh_host_pubkey causes the SCP license transfer to
// abort with a host-key mismatch instead of silently trusting the peer — proving
// the ssh.FixedHostKey verification restored in CTXMYT-2537 is actually enforced
// (a MITM presenting a different host key is rejected before the nsroot password
// is sent).
//
// It drives the real offline-licensing workflow up to the first SCP download
// (NITRO version check + license-activation-data GET, then SCPDownload), so it
// requires a reachable, non-CPX appliance and a valid las_secrets_json. It is
// Skip-gated for the same reason the other SSH/license tests in this package are;
// un-skip and run it manually against a disposable appliance with real values
// substituted. The key below is a throwaway ed25519 key that does not match any
// real ADC host key.
func TestAccNSLASLicenseOffline_wrongHostPubKey(t *testing.T) {
	t.Skip("Requires a reachable, LAS-capable ADC and a valid las_secrets_json; run manually")
	if isCpxRun {
		t.Skip("ssh does not work correctly with CPX")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccNSLASLicenseOffline_wrongHostPubKey,
				ExpectError: regexp.MustCompile("host key mismatch"),
			},
		},
	})
}

const testAccNSLASLicenseOffline_wrongHostPubKey = `
resource "citrixadc_nslaslicense_offline" "tf_las_offline" {
  entitlement_name = "VPX 10000 Premium"
  las_secrets_json = "/tmp/las_secrets.json"
  ssh_host_pubkey  = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIHPEytt2ZGu260pcnAF8uTKF4jv7FxQ5/EPR6TL/3S3d ctxmyt-2537-wrong-key-fixture"
}
`
