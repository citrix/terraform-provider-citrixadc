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

type LBVserver struct {
	Name            string
	ServiceType     string
	VIP             string
	Port            int
	PersistenceType string
	LbMethod        string
}

func (lb *LBVserver) Id() string {
	return lb.Name
}

func (c *NetScalerNitroClient) CreateLBVserver(lb *LBVserver) error {
	nitroClient := netscaler.NewNitroClient(c.Endpoint, c.Username, c.Password)
	var lbStruct = new(netscaler.NetscalerLB)
	lbStruct.Name = lb.Name
	lbStruct.ServiceType = lb.ServiceType
	lbStruct.Ipv46 = lb.VIP
	lbStruct.Port = lb.Port
	lbStruct.PersistenceType = lb.PersistenceType
	lbStruct.LbMethod = lb.LbMethod
	_, err := nitroClient.CreateLBVserver(lbStruct)
	if err != nil {
		log.Fatal("Failed to create loadbalancer %s", lb.Name)
		return err
	}
	return nil
}

func (c *NetScalerNitroClient) DeleteLBVserver(lbName string) error {
	nitroClient := netscaler.NewNitroClient(c.Endpoint, c.Username, c.Password)
	err := nitroClient.DeleteLBVserver(lbName)
	if err != nil {
		log.Fatal("Failed to delete loadbalancer %s", lbName)
		return err
	}
	return nil
}

func resourceNetScalerLB() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createLbFunc,
		Read:          readLbFunc,
		Update:        updateLbFunc,
		Delete:        deleteLbFunc,
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
			"lb_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func createLbFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*NetScalerNitroClient)
	var lbName string
	if v, ok := d.GetOk("name"); ok {
		lbName = v.(string)
	} else {
		lbName = resource.PrefixedUniqueId("tf-lb-")
		d.Set("name", lbName)
	}
	lb := LBVserver{
		Name:            lbName,
		VIP:             d.Get("vip").(string),
		Port:            d.Get("port").(int),
		ServiceType:     d.Get("service_type").(string),
		PersistenceType: d.Get("persistence_type").(string),
		LbMethod:        d.Get("lb_method").(string),
	}

	err := client.CreateLBVserver(&lb)
	if err != nil {
		return err
	}

	d.SetId(lb.Id())

	return nil
}

func readLbFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*NetScalerNitroClient)
	lbName := d.Id()
	log.Printf("Reading loadbalancer state %s", lbName)
	found := client.FindResource("lbvserver", lbName)
	if !found {
		log.Printf("Clearing loadbalancer state %s", lbName)
		d.SetId("")
	}
	return nil
}

func updateLbFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func deleteLbFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*NetScalerNitroClient)
	lbName := d.Id()
	err := client.DeleteLBVserver(lbName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
