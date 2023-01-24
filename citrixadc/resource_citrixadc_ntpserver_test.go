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
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

const testAccNtpserver_basic_ip = `

resource "citrixadc_ntpserver" "tf_ntpserver" {
	serverip          = "10.222.74.200"
	minpoll            = 5
	maxpoll            = 9
	preferredntpserver = "NO"
  
  }
  
`
const testAccNtpserver_update_ip = `

resource "citrixadc_ntpserver" "tf_ntpserver" {
	serverip         = "10.222.74.200"
	minpoll            = 6
	maxpoll            = 10
	preferredntpserver = "YES"
  
  } 
`

func TestAccNtpserver_ip(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNtpserverDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNtpserver_basic_ip,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNtpserverExist("citrixadc_ntpserver.tf_ntpserver", nil),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "serverip", "10.222.74.200"),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "minpoll", "5"),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "maxpoll", "9"),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "preferredntpserver", "NO"),

				),
			},
			resource.TestStep{
				Config: testAccNtpserver_update_ip,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNtpserverExist("citrixadc_ntpserver.tf_ntpserver", nil),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "serverip", "10.222.74.200"),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "minpoll", "6"),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "maxpoll", "10"),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "preferredntpserver", "YES"),
				),
			},		
		},
	})
}


const testAccNtpserver_basic_servername = `

resource "citrixadc_ntpserver" "tf_ntpserver" {
	servername         = "www.example.com"
	minpoll            = 5
	maxpoll            = 9
	preferredntpserver = "NO"
  
  }
  
`
const testAccNtpserver_update_servername = `

resource "citrixadc_ntpserver" "tf_ntpserver" {
	servername        = "www.example.com"
	minpoll            = 6
	maxpoll            = 10
	preferredntpserver = "YES"
  
  }
  
`

func TestAccNtpserver_servername(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNtpserverDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccNtpserver_basic_servername,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNtpserverExist("citrixadc_ntpserver.tf_ntpserver", nil),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "servername", "www.example.com"),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "minpoll", "5"),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "maxpoll", "9"),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "preferredntpserver", "NO"),

				),
			},
			resource.TestStep{
				Config: testAccNtpserver_update_servername,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNtpserverExist("citrixadc_ntpserver.tf_ntpserver", nil),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "servername", "www.example.com"),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "minpoll", "6"),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "maxpoll", "10"),
					resource.TestCheckResourceAttr("citrixadc_ntpserver.tf_ntpserver", "preferredntpserver", "YES"),
				),
			},		
		},
	})
}
func testAccCheckNtpserverExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ntpserver name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}
		ntpserverName := rs.Primary.ID
		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		dataArr, err := nsClient.FindAllResources(service.Ntpserver.Type())

		if err != nil {
			return err
		}

		found := false
		for _, v := range dataArr {
			if v["serverip"] == ntpserverName || v["servername"] == ntpserverName {
				found = true
				break
			}
	
		}
		if !found {
			return fmt.Errorf("ntpserver %s not found", n)
		}

		return nil
	}
}

func testAccCheckNtpserverDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_ntpserver" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		ntpserverName := rs.Primary.ID
		dataArr, err := nsClient.FindAllResources(service.Ntpserver.Type())
		
		if err != nil {
			return err
		}
		
		found := false
		for _, v := range dataArr {
			if v["serverip"] == ntpserverName || v["servername"] == ntpserverName {
				found = true
				break
			}
		}
		
		if found {
			return fmt.Errorf("ntpserver %s still exists", ntpserverName)
		}
	}
	return nil
}