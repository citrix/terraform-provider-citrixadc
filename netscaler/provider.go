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
package netscaler

import (
	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

type NetScalerNitroClient struct {
	Username  string
	Password  string
	Endpoint  string
	ProxiedNS string
	client    *netscaler.NitroClient
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
			DefaultFunc: schema.EnvDefaultFunc("NS_LOGIN", "nsroot"),
		},
		"password": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "Password to login to the NetScaler",
			DefaultFunc: schema.EnvDefaultFunc("NS_PASSWORD", "nsroot"),
		},
		"endpoint": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The URL to the API",
			DefaultFunc: schema.EnvDefaultFunc("NS_URL", nil),
		},
		"proxiedNS": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The device that is proxied by NMAS",
			DefaultFunc: schema.EnvDefaultFunc("_MPS_API_PROXY_MANAGED_INSTANCE_IP", nil),
		},
	}
}

func providerResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		//"netscaler_lb":  resourceNetScalerLB(),
		"netscaler_lbvserver":    resourceNetScalerLbvserver(),
		"netscaler_service":      resourceNetScalerService(),
		"netscaler_csvserver":    resourceNetScalerCsvserver(),
		"netscaler_cspolicy":     resourceNetScalerCspolicy(),
		"netscaler_sslcertkey":   resourceNetScalerSslcertkey(),
		"netscaler_lbmonitor":    resourceNetScalerLbmonitor(),
		"netscaler_servicegroup": resourceNetScalerServicegroup(),
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	c := NetScalerNitroClient{
		Username:  d.Get("username").(string),
		Password:  d.Get("password").(string),
		Endpoint:  d.Get("endpoint").(string),
		ProxiedNS: d.Get("proxiedNS").(string),
	}
	client := netscaler.NewNitroClient(c.Endpoint, c.Username, c.Password)
	if c.ProxiedNS != "" {
		client = netscaler.NewProxyingNitroClient(c.Endpoint, c.Username, c.Password, c.ProxiedNS)
	}
	c.client = client

	return &c, nil
}
