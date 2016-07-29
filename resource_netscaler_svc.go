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
package main

import (
	netscaler "github.com/chiradeep/terraform-provider-netscaler/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"log"
)

type Service struct {
	Name        string
	ServiceType string
	Ip          string
	Port        int
}

func (svc *Service) Id() string {
	return svc.Name
}

func (c *NetScalerNitroClient) CreateService(svc *Service) error {
	nitroClient := netscaler.NewNitroClient(c.Endpoint, c.Username, c.Password)
	var svcStruct = new(netscaler.NetscalerService)
	svcStruct.Name = svc.Name
	svcStruct.ServiceType = svc.ServiceType
	svcStruct.Ip = svc.Ip
	svcStruct.Port = svc.Port
	_, err := nitroClient.CreateService(svcStruct)
	if err != nil {
		log.Fatal("Failed to create service %s", svc.Name)
		return err
	}
	return nil
}

func (c *NetScalerNitroClient) DeleteService(svcName string) error {
	nitroClient := netscaler.NewNitroClient(c.Endpoint, c.Username, c.Password)
	err := nitroClient.DeleteService(svcName)
	if err != nil {
		log.Fatal("Failed to delete service %s", svcName)
		return err
	}
	return nil
}

func resourceNetScalerSvc() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createSvcFunc,
		Read:          readSvcFunc,
		Update:        updateSvcFunc,
		Delete:        deleteSvcFunc,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"service_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"persistence_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"svc_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func createSvcFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*NetScalerNitroClient)
	var svcName string
	if v, ok := d.GetOk("name"); ok {
		svcName = v.(string)
	} else {
		svcName = resource.PrefixedUniqueId("tf-svc-")
		d.Set("name", svcName)
	}
	svc := Service{
		Name:        svcName,
		Ip:          d.Get("ip").(string),
		Port:        d.Get("port").(int),
		ServiceType: d.Get("service_type").(string),
	}

	err := client.CreateService(&svc)
	if err != nil {
		return err
	}

	d.SetId(svc.Id())

	return nil
}

func readSvcFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*NetScalerNitroClient)
	svcName := d.Id()
	log.Printf("Reading service state %s", svcName)
	found := client.FindResource("svcvserver", svcName)
	if !found {
		log.Printf("Clearing service state %s", svcName)
		d.SetId("")
	}
	return nil
}

func updateSvcFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func deleteSvcFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*NetScalerNitroClient)
	svcName := d.Id()
	err := client.DeleteService(svcName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
