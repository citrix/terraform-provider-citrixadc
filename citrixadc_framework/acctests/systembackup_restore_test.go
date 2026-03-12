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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

const testAccSystembackupRestore_basic = `

	resource "citrixadc_systembackup_create" "tf_systembackup_create" {
		filename         = "my_backup_file"
		level            = "basic"
		uselocaltimezone = "true"
	}
	resource "citrixadc_systembackup_restore" "tf_systembackup_restore" {
		filename   = "${citrixadc_systembackup_create.tf_systembackup_create.filename}.tgz"
		skipbackup = "false"
	}
  
`

func TestAccSystembackupRestore_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             nil,
		Steps: []resource.TestStep{
			{
				Config: testAccSystembackupRestore_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSystembackupRestoreExist("citrixadc_systembackup_restore.tf_systembackup_restore", nil),
				),
			},
		},
	})
}

func testAccCheckSystembackupRestoreExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No systembackup name is set")
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
