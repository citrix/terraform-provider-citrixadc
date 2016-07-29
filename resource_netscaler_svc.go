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
	"errors"
	netscaler "github.com/chiradeep/terraform-provider-netscaler/netscaler"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"

	"fmt"
	"log"
)

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
			"ip": &schema.Schema{
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
			"lb": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func createSvcFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*NetScalerNitroClient).client
	lbName := d.Get("lb").(string)
	lbFound := client.ResourceExists("lbvserver", lbName)
	if !lbFound {
		log.Printf("No lb with name %s found", lbName)
		return errors.New(fmt.Sprintf("No lb with name %s found", lbName))
	}
	var svcName string
	if v, ok := d.GetOk("name"); ok {
		svcName = v.(string)
	} else {
		svcName = resource.PrefixedUniqueId("tf-svc-" + lbName + "-")
		d.Set("name", svcName)
	}
	log.Printf("****Creating service %s", svcName)
	svc := netscaler.NetscalerService{
		Name:        svcName,
		Ip:          d.Get("ip").(string),
		Port:        d.Get("port").(int),
		ServiceType: d.Get("service_type").(string),
	}

	err := client.AddAndBindService(lbName, &svc)
	if err != nil {
		return err
	}

	d.SetId(svcName)

	return nil
}

func readSvcFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*NetScalerNitroClient).client
	svcName := d.Id()
	log.Printf("****Reading service state %s", svcName)
	found := client.ResourceExists("service", svcName)
	if !found {
		log.Printf("Clearing service state %s", svcName)
		d.SetId("")
	}
	return nil
}

func updateSvcFunc(d *schema.ResourceData, meta interface{}) error {
	svcName := d.Id()
	log.Printf("****Updating service state %s", svcName)
	d.Set("name", svcName) //FIXME: why?
	return nil
}

func deleteSvcFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*NetScalerNitroClient).client
	svcName := d.Id()
	err := client.DeleteService(svcName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
