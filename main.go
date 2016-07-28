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
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"

	"log"
)

type NetScalerNitroClient struct {
	Username string
	Password string
	Endpoint string
}

type LBVserver struct {
	Name            string
	ServiceType     string
	VIP             string
	Port            int
	PersistenceType string
	LbMethod        string
}

func (lb *LBVserver) Id() string {
	return "id-" + lb.Name + "!"
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

func main() {
	opts := plugin.ServeOpts{
		ProviderFunc: Provider,
	}
	plugin.Serve(&opts)
}

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema:        providerSchema(),
		ResourcesMap:  providerResources(),
		ConfigureFunc: providerConfigure,
	}
}

func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"username": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "Username to login to the NetScaler",
		},
		"password": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "Password to login to the NetScaler",
		},
		"endpoint": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The URL to the API",
		},
	}
}

func providerResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"netscaler_lb": &schema.Resource{
			SchemaVersion: 1,
			Create:        createFunc,
			Read:          readFunc,
			Update:        updateFunc,
			Delete:        deleteFunc,
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
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	client := NetScalerNitroClient{
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
		Endpoint: d.Get("endpoint").(string),
	}

	return &client, nil
}

func createFunc(d *schema.ResourceData, meta interface{}) error {
	/*
		Name            string
		ServiceType     string
		VIP             string
		Port            int
		PersistenceType string
		LbMethod        string
	*/
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

func readFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func updateFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func deleteFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}
