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
	"log"
	"sync"

	"github.com/chiradeep/go-nitro/netscaler"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/hashicorp/terraform/version"
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
		"proxied_ns": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Target NS ip. When defined username, password and endpoint must refer to MAS.",
			DefaultFunc: schema.EnvDefaultFunc("_MPS_API_PROXY_MANAGED_INSTANCE_IP", ""),
		},
	}
}

func providerResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"citrixadc_lbvserver":            resourceCitrixAdcLbvserver(),
		"citrixadc_service":              resourceCitrixAdcService(),
		"citrixadc_csvserver":            resourceCitrixAdcCsvserver(),
		"citrixadc_cspolicy":             resourceCitrixAdcCspolicy(),
		"citrixadc_csaction":             resourceCitrixAdcCsaction(),
		"citrixadc_sslcertkey":           resourceCitrixAdcSslcertkey(),
		"citrixadc_sslprofile":           resourceCitrixAdcSslprofile(),
		"citrixadc_sslparameter":         resourceCitrixAdcSslparameter(),
		"citrixadc_lbmonitor":            resourceCitrixAdcLbmonitor(),
		"citrixadc_servicegroup":         resourceCitrixAdcServicegroup(),
		"citrixadc_nsacl":                resourceCitrixAdcNsacl(),
		"citrixadc_nsacls":               resourceCitrixAdcNsacls(),
		"citrixadc_inat":                 resourceCitrixAdcInat(),
		"citrixadc_rnat":                 resourceCitrixAdcRnats(),
		"citrixadc_gslbsite":             resourceCitrixAdcGslbsite(),
		"citrixadc_gslbvserver":          resourceCitrixAdcGslbvserver(),
		"citrixadc_gslbservice":          resourceCitrixAdcGslbservice(),
		"citrixadc_server":               resourceCitrixAdcServer(),
		"citrixadc_responderaction":      resourceCitrixAdcResponderaction(),
		"citrixadc_responderpolicy":      resourceCitrixAdcResponderpolicy(),
		"citrixadc_responderpolicylabel": resourceCitrixAdcResponderpolicylabel(),
		"citrixadc_rewriteaction":        resourceCitrixAdcRewriteaction(),
		"citrixadc_rewritepolicy":        resourceCitrixAdcRewritepolicy(),
		"citrixadc_rewritepolicylabel":   resourceCitrixAdcRewritepolicylabel(),
		"citrixadc_nsip":                 resourceCitrixAdcNsip(),
		"citrixadc_nsip6":                resourceCitrixAdcNsip6(),
		"citrixadc_ipset":                resourceCitrixAdcIpset(),
		"citrixadc_nsfeature":            resourceCitrixAdcNsfeature(),
		"citrixadc_systemuser":           resourceCitrixAdcSystemuser(),
		"citrixadc_systemgroup":          resourceCitrixAdcSystemgroup(),
		"citrixadc_systemcmdpolicy":      resourceCitrixAdcSystemcmdpolicy(),
		"citrixadc_interface":            resourceCitrixAdcInterface(),
		"citrixadc_nstcpprofile":         resourceCitrixAdcNstcpprofile(),
		"citrixadc_nslicense":            resourceCitrixAdcNslicense(),
		"citrixadc_cluster":              resourceCitrixAdcCluster(),
		"citrixadc_clusterfiles_syncer":  resourceCitrixAdcClusterfilesSyncer(),
		"citrixadc_systemfile":           resourceCitrixAdcSystemfile(),
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
		ProxiedNs: d.Get("proxied_ns").(string),
		SslVerify: !d.Get("insecure_skip_verify").(bool),
	}
	client, err := netscaler.NewNitroClientFromParams(params)
	if err != nil {
		return nil, err
	}

	c.client = client
	log.Printf("[DEBUG] citrixadc-provider: Terraform version imported: %s\n", version.Version)

	return &c, nil
}
