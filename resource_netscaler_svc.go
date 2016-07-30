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
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"service_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
	_ = readSvcFunc(d, meta)

	return nil
}

func readSvcFunc(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*NetScalerNitroClient).client
	svcName := d.Id()
	log.Printf("****Reading service state %s", svcName)
	data, err := client.FindResource("service", svcName)
	if err != nil {
		log.Printf("Clearing service %s", svcName)
		d.SetId("")
		return nil
	}
	/*{"name": "tf-svc-sample_lb2-qmdh42onpnf6fdqaipo5fwrsay", "numofconnections": 0, "servername": "10.33.44.54", "policyname": "10.33.44.54", "servicetype": "HTTP",
	  "serviceconftype2": "Configured", "port": 80, "gslb": "NONE", "cachetype": "SERVER", "maxclient": "0", "maxreq": "0", "cacheable": "NO", "cip": "DISABLED",
	   "usip": "NO", "pathmonitor": "NO", "pathmonitorindv": "NO", "useproxyport": "YES", "sc": "OFF", "dup_weight": "0", "dup_state": "ENABLED", "sp": "ON",
	   "rtspsessionidremap": "OFF", "failedprobes": "1", "clttimeout": 180, "totalprobes": "180", "svrtimeout": 360, "totalfailedprobes": "360", "publicip": "0.0.0.0",
	   "publicport": 80, "customserverid": "None", "cka": "NO", "tcpb": "NO", "processlocal": "DISABLED", "cmp": "NO", "maxbandwidth": "0", "accessdown": "NO",
	   "svrstate": "DOWN", "delay": 0, "ipaddress": "10.33.44.54", "monthreshold": "0", "monstate": "ENABLED", "monitor_state": "Unknown", "monstatcode": 0,
	   "lastresponse": "Probing ownership is with some other node in cluster.", "responsetime": "0", "riseapbrstatsmsgcode": 0, "riseapbrstatsmsgcode2": 1, "monstatparam1": 0,
	   "monstatparam2": 0, "monstatparam3": 0, "downstateflush": "ENABLED", "statechangetimesec": "Fri Jul 29 22:32:58 2016",
	   "statechangetimemsec": "67", "tickssincelaststatechange": "13580", "stateupdatereason": "0",
	   "clmonview": "0", "graceful": "NO", "monitortotalprobes": "0", "monitortotalfailedprobes": "0", "monitorcurrentfailedprobes": "0", "healthmonitor": "YES",
	   "appflowlog": "ENABLED", "serviceipstr": "10.33.44.54", "td": "0", "passive": false }
	*/
	d.Set("name", data["name"])
	d.Set("service_type", data["servicetype"])
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
