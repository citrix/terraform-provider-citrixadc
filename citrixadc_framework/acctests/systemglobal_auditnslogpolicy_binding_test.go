/*
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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testAccSystemglobal_auditnslogpolicy_binding_basic = `

	resource "citrixadc_auditnslogaction" "tf_auditnslogaction" {
		name     = "tf_auditnslogaction"
		serverip = "1.1.1.1"
		loglevel = ["ALERT", "CRITICAL"]
	}
	resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
		name   = "tf_auditnslogpolicy"
		rule   = "ns_true"
		action = citrixadc_auditnslogaction.tf_auditnslogaction.name
	}

	resource "citrixadc_systemglobal_auditnslogpolicy_binding" "tf_systemglobal_auditnslogpolicy_binding" {
		policyname = citrixadc_auditnslogpolicy.tf_auditnslogpolicy.name
		priority   = 50
	}
`

const testAccSystemglobal_auditnslogpolicy_binding_basic_step2 = `
	# Keep the above bound resources without the actual binding to check proper deletion

	resource "citrixadc_auditnslogaction" "tf_auditnslogaction" {
		name     = "tf_auditnslogaction"
		serverip = "1.1.1.1"
		loglevel = ["ALERT", "CRITICAL"]
	}
	resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
		name   = "tf_auditnslogpolicy"
		rule   = "ns_true"
		action = citrixadc_auditnslogaction.tf_auditnslogaction.name
	}
`

func TestAccSystemglobal_auditnslogpolicy_binding_import(t *testing.T) {
	const resAddr = "citrixadc_systemglobal_auditnslogpolicy_binding.tf_systemglobal_auditnslogpolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemglobal_auditnslogpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{Config: testAccSystemglobal_auditnslogpolicy_binding_basic},
			{Config: testAccSystemglobal_auditnslogpolicy_binding_basic, ResourceName: resAddr, ImportState: true, ImportStateVerify: true, ImportStateVerifyIgnore: []string{}},
		},
	})
}

func TestAccSystemglobal_auditnslogpolicy_binding_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSystemglobal_auditnslogpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemglobal_auditnslogpolicy_binding_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemglobal_auditnslogpolicy_bindingExist("citrixadc_systemglobal_auditnslogpolicy_binding.tf_systemglobal_auditnslogpolicy_binding", nil),
				),
			},
			{
				Config: testAccSystemglobal_auditnslogpolicy_binding_basic_step2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemglobal_auditnslogpolicy_bindingNotExist("citrixadc_systemglobal_auditnslogpolicy_binding.tf_systemglobal_auditnslogpolicy_binding", "tf_auditnslogpolicy"),
				),
			},
		},
	})
}

func testAccCheckSystemglobal_auditnslogpolicy_bindingExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No systemglobal_auditnslogpolicy_binding id is set")
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

		policyname := rs.Primary.ID

		findParams := service.FindParams{
			ResourceType:             "systemglobal_auditnslogpolicy_binding",
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to find the one with the matching policyname
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("systemglobal_auditnslogpolicy_binding %s not found", n)
		}

		return nil
	}
}

func testAccCheckSystemglobal_auditnslogpolicy_bindingNotExist(n string, id string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		// Use the shared utility function to get a configured client
		client, err := testAccGetFrameworkClient()
		if err != nil {
			return fmt.Errorf("Failed to get test client: %v", err)
		}

		policyname := id
		findParams := service.FindParams{
			ResourceType:             "systemglobal_auditnslogpolicy_binding",
			ResourceMissingErrorCode: 258,
		}
		dataArr, err := client.FindResourceArrayWithParams(findParams)

		// Unexpected error
		if err != nil {
			return err
		}

		// Iterate through results to hopefully not find the one with the matching policyname
		found := false
		for _, v := range dataArr {
			if v["policyname"].(string) == policyname {
				found = true
				break
			}
		}

		if found {
			return fmt.Errorf("systemglobal_auditnslogpolicy_binding %s was found, but it should have been destroyed", n)
		}

		return nil
	}
}

func testAccCheckSystemglobal_auditnslogpolicy_bindingDestroy(s *terraform.State) error {
	// Use the shared utility function to get a configured client
	client, err := testAccGetFrameworkClient()
	if err != nil {
		return fmt.Errorf("Failed to get test client: %v", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_systemglobal_auditnslogpolicy_binding" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := client.FindResource(service.Systemglobal_auditnslogpolicy_binding.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("systemglobal_auditnslogpolicy_binding %s still exists", rs.Primary.ID)
		}

	}

	return nil
}

const testAccSystemglobal_auditnslogpolicy_bindingDataSource_basic = `

	resource "citrixadc_auditnslogaction" "tf_auditnslogaction" {
		name     = "tf_auditnslogaction"
		serverip = "1.1.1.1"
		loglevel = ["ALERT", "CRITICAL"]
	}
	resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
		name   = "tf_auditnslogpolicy"
		rule   = "ns_true"
		action = citrixadc_auditnslogaction.tf_auditnslogaction.name
	}

	resource "citrixadc_systemglobal_auditnslogpolicy_binding" "tf_systemglobal_auditnslogpolicy_binding" {
		policyname = citrixadc_auditnslogpolicy.tf_auditnslogpolicy.name
		priority   = 50
	}

	data "citrixadc_systemglobal_auditnslogpolicy_binding" "tf_systemglobal_auditnslogpolicy_binding" {
		policyname = citrixadc_systemglobal_auditnslogpolicy_binding.tf_systemglobal_auditnslogpolicy_binding.policyname
	}
`

func TestAccSystemglobal_auditnslogpolicy_bindingDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSystemglobal_auditnslogpolicy_bindingDataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.citrixadc_systemglobal_auditnslogpolicy_binding.tf_systemglobal_auditnslogpolicy_binding", "policyname", "tf_auditnslogpolicy"),
					resource.TestCheckResourceAttr("data.citrixadc_systemglobal_auditnslogpolicy_binding.tf_systemglobal_auditnslogpolicy_binding", "priority", "50"),
				),
			},
		},
	})
}

// testAccSystemglobal_auditnslogpolicy_binding_upgrade_basic reuses the _basic config
// (binding + all prerequisite resources). It is valid under BOTH the SDK v2 2.2.0
// schema and the current Framework schema because the migration restored the SDK v2
// attribute names.
const testAccSystemglobal_auditnslogpolicy_binding_upgrade_basic = `

	resource "citrixadc_auditnslogaction" "tf_auditnslogaction" {
		name     = "tf_auditnslogaction"
		serverip = "1.1.1.1"
		loglevel = ["ALERT", "CRITICAL"]
	}
	resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
		name   = "tf_auditnslogpolicy"
		rule   = "ns_true"
		action = citrixadc_auditnslogaction.tf_auditnslogaction.name
	}

	resource "citrixadc_systemglobal_auditnslogpolicy_binding" "tf_systemglobal_auditnslogpolicy_binding" {
		policyname = citrixadc_auditnslogpolicy.tf_auditnslogpolicy.name
		priority   = 50
	}
`

// TestAccSystemglobal_auditnslogpolicy_binding_sdkv2StateUpgrade verifies that state
// written by the last SDK v2 release is correctly upgraded when the same config is
// subsequently managed by the current Framework provider. Step 1 creates the binding
// with citrix/citrixadc 2.2.0 (writes the legacy id "tf_auditnslogpolicy" — the SDK v2
// d.SetId(policyname)). Step 2 refreshes/plans/applies the same config through the
// Framework provider, exercising ParseIdString on the legacy id; the Framework
// recomputes the id on Read (SetAttrFromGet). For this single-key resource the
// canonical new-format id is the plain policyname, so it stays "tf_auditnslogpolicy".
func TestAccSystemglobal_auditnslogpolicy_binding_sdkv2StateUpgrade(t *testing.T) {
	resourceAddr := "citrixadc_systemglobal_auditnslogpolicy_binding.tf_systemglobal_auditnslogpolicy_binding"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckSystemglobal_auditnslogpolicy_bindingDestroy,
		Steps: []resource.TestStep{
			// Step 1: create with the last SDK v2 release -> state carries the legacy id.
			{
				ExternalProviders: map[string]resource.ExternalProvider{
					"citrixadc": {
						Source:            "citrix/citrixadc",
						VersionConstraint: "2.2.0",
					},
				},
				Config: testAccSystemglobal_auditnslogpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemglobal_auditnslogpolicy_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "tf_auditnslogpolicy"),
				),
			},
			// Step 2: refresh/plan/apply the SAME config through the current Framework
			// provider. The legacy-id state is read via ParseIdString and the id is
			// recomputed on Read; for this single-key resource it remains "tf_auditnslogpolicy".
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config:                   testAccSystemglobal_auditnslogpolicy_binding_upgrade_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystemglobal_auditnslogpolicy_bindingExist(resourceAddr, nil),
					resource.TestCheckResourceAttr(resourceAddr, "id", "tf_auditnslogpolicy"),
				),
			},
		},
	})
}
