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
	"sync"
)

type NetScalerNitroClient struct {
	Username string
	Password string
	Endpoint string
	client   *netscaler.NitroClient
	lock     sync.Mutex
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
		"username": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Username to login to the NetScaler",
			DefaultFunc: schema.EnvDefaultFunc("NS_LOGIN", "nsroot"),
		},
		"password": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Password to login to the NetScaler",
			DefaultFunc: schema.EnvDefaultFunc("NS_PASSWORD", "nsroot"),
		},
		"endpoint": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The URL to the API",
			DefaultFunc: schema.EnvDefaultFunc("NS_URL", nil),
		},
		"insecure_skip_verify": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Ignore validity of endpoint TLS certificate if true",
			Default:     false,
		},
	}
}

func providerResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"netscaler_lbvserver":    resourceNetScalerLbvserver(),
		"netscaler_service":      resourceNetScalerService(),
		"netscaler_csvserver":    resourceNetScalerCsvserver(),
		"netscaler_cspolicy":     resourceNetScalerCspolicy(),
		"netscaler_sslcertkey":   resourceNetScalerSslcertkey(),
		"netscaler_lbmonitor":    resourceNetScalerLbmonitor(),
		"netscaler_servicegroup": resourceNetScalerServicegroup(),
		"netscaler_nsacl":        resourceNetScalerNsacl(),
		"netscaler_nsacls":       resourceNetScalerNsacls(),
		"netscaler_inat":         resourceNetScalerInat(),
		"netscaler_rnat":         resourceNetScalerRnats(),
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	c := NetScalerNitroClient{
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
		Endpoint: d.Get("endpoint").(string),
	}

	params := netscaler.NitroParams{
		Url:       d.Get("endpoint").(string),
		Username:  d.Get("username").(string),
		Password:  d.Get("password").(string),
		SslVerify: !d.Get("insecure_skip_verify").(bool),
	}
	client, err := netscaler.NewNitroClientFromParams(params)
	if err != nil {
		return nil, err
	}

	c.client = client

	return &c, nil
}
