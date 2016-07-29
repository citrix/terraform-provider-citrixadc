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
	client := meta.(*NetScalerNitroClient).client
	var lbName string
	if v, ok := d.GetOk("name"); ok {
		lbName = v.(string)
	} else {
		lbName = resource.PrefixedUniqueId("tf-lb-")
		d.Set("name", lbName)
	}
	lb := netscaler.NetscalerLB{
		Name:            lbName,
		Ipv46:           d.Get("vip").(string),
		Port:            d.Get("port").(int),
		ServiceType:     d.Get("service_type").(string),
		PersistenceType: d.Get("persistence_type").(string),
		LbMethod:        d.Get("lb_method").(string),
	}

	_, err := client.CreateLBVserver(&lb)
	if err != nil {
		return err
	}

	d.SetId(lbName)
	err = readLbFunc(d, meta)
	if err != nil {
		log.Printf("?? we just created this loadbalancer but we can't read it ?? %s", lbName)
		return nil
	}
	return nil
}

func readLbFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*NetScalerNitroClient).client
	lbName := d.Id()
	log.Printf("Reading loadbalancer state %s", lbName)
	data, err := client.FindResource("lbvserver", lbName)
	if err != nil {
		log.Printf("Clearing loadbalancer state %s", lbName)
		d.SetId("")
		return nil
	}
	/* { "name": "sample_lb2", "insertvserveripport": "OFF", "ipv46": "10.71.136.151", "ippattern": "0.0.0.0", "ipmask": "*", "listenpolicy": "NONE",
	   "ipmapping": "0.0.0.0", "port": 443, "range": "1", "servicetype": "HTTP", "type": "ADDRESS", "curstate": "DOWN", "effectivestate": "DOWN", "status": 1,
	   "lbrrreason": 0, "cachetype": "SERVER", "authentication": "OFF", "authn401": "OFF", "dynamicweight": "0", "priority": "0", "clttimeout": "180",
	   "somethod": "NONE", "sopersistence": "DISABLED", "sopersistencetimeout": "2", "healththreshold": "0", "lbmethod": "LEASTCONNECTION", "backuplbmethod": "ROUNDROBIN",
	   "dataoffset": "0", "health": "0", "datalength": "0", "ruletype": "0", "m": "IP", "persistencetype": "NONE", "timeout": 2, "persistmask": "255.255.255.255",
	   "v6persistmasklen": "128", "persistencebackup": "NONE", "cacheable": "NO", "pq": "OFF", "sc": "OFF", "rtspnat": "OFF", "sessionless": "DISABLED", "map": "OFF",
	   "connfailover": "DISABLED", "redirectportrewrite": "DISABLED", "downstateflush": "ENABLED", "disableprimaryondown": "DISABLED", "gt2gb": "DISABLED", "consolidatedlconn":
	   "GLOBAL", "consolidatedlconngbl": "YES", "thresholdvalue": 0, "invoke": false, "version": 0, "totalservices": "2", "activeservices": "0",
	   "statechangetimesec": "Fri Jul 29 19:14:02 2016", "statechangetimeseconds": "1469819642", "statechangetimemsec": "382", "tickssincelaststatechange": "728421",
	   "hits": "0", "pipolicyhits": "0", "push": "DISABLED", "pushlabel": "none", "pushmulticlients": "NO", "policysubtype": "0", "l2conn": "OFF", "appflowlog": "ENABLED",
	   "isgslb": false, "icmpvsrresponse": "PASSIVE", "rhistate": "PASSIVE", "newservicerequestunit": "PER_SECOND", "vsvrbindsvcip": "10.71.136.151", "vsvrbindsvcport": 0,
	   "skippersistency": "None", "td": "0", "minautoscalemembers": "0", "maxautoscalemembers": "0", "macmoderetainvlan": "DISABLED", "dns64": "DISABLED", "bypassaaaa": "NO",
	   "processlocal": "DISABLED", "vsvrdynconnsothreshold": "0" }
	*/
	d.Set("name", data["name"])
	d.Set("persistence_type", data["persistencetype"])
	d.Set("lb_method", data["lbmethod"])
	d.Set("service_type", data["servicetype"])

	return nil

}

func updateLbFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func deleteLbFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*NetScalerNitroClient).client
	lbName := d.Id()
	err := client.DeleteLBVserver(lbName)
	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}
