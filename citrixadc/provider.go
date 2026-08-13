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
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type NetScalerNitroClient struct {
	Username    string
	Password    string
	Endpoint    string
	HttpTimeout int
	NsTimeout   int
	client      *service.NitroClient
	lock        sync.Mutex
}

func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema:         providerSchema(),
		ResourcesMap:   providerResources(),
		DataSourcesMap: providerDataSources(),
	}
	provider.ConfigureContextFunc = func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		terraformVersion := provider.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}
		return providerConfigure(ctx, d, terraformVersion)
	}

	return provider
}

func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"username": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Username to login to the NetScaler",
			DefaultFunc: schema.EnvDefaultFunc("NS_LOGIN", nil),
		},
		"password": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Password to login to the NetScaler",
			DefaultFunc: schema.EnvDefaultFunc("NS_PASSWORD", nil),
		},
		"endpoint": {
			Type:        schema.TypeString,
			Optional:    true,
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
			Description: "Target NS ip. When defined username, password and endpoint must refer to NetScaler Console.",
			DefaultFunc: schema.EnvDefaultFunc("_MPS_API_PROXY_MANAGED_INSTANCE_IP", ""),
		},
		"partition": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Partition to target",
			DefaultFunc: schema.EnvDefaultFunc("NS_PARTITION", nil),
		},
		"do_login": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Perform login to NetScaler",
			Default:     false,
		},
		"is_cloud": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Set to true when using NetScaler Console Cloud",
			Default:     false,
		},
		"http_timeout": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Timeout in seconds for the underlying NITRO HTTP client (Go http.Client.Timeout). It bounds the total duration of each API request so that unreachable endpoints fail fast instead of hanging on the operating system's TCP connection timeout. Can be sourced from the NS_HTTP_TIMEOUT environment variable. When 0 or unset, no client-side timeout is applied.",
			DefaultFunc: func() (interface{}, error) {
				v := os.Getenv("NS_HTTP_TIMEOUT")
				if v == "" {
					return 0, nil
				}
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("the NS_HTTP_TIMEOUT environment variable must be an integer number of seconds")
				}
				return n, nil
			},
		},
		"ns_timeout": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "NITRO session timeout in seconds requested at login. It is sent to the ADC in the login request and controls the idle lifetime of the NITRO session; it only takes effect when 'do_login' is true. Can be sourced from the NS_TIMEOUT environment variable. When 0 or unset, the ADC applies its own default session timeout.",
			DefaultFunc: func() (interface{}, error) {
				v := os.Getenv("NS_TIMEOUT")
				if v == "" {
					return 0, nil
				}
				n, err := strconv.Atoi(v)
				if err != nil {
					return nil, fmt.Errorf("the NS_TIMEOUT environment variable must be an integer number of seconds")
				}
				return n, nil
			},
		},
	}
}

func providerDataSources() map[string]*schema.Resource {
	// All SDK v2 data sources have been migrated to the Plugin Framework provider
	// (citrixadc_framework/custom_resources/{nsversion,nitro_info,sslcipher_sslvserver_bindings})
	// and are served there via the tf6mux. Kept empty here so the muxed servers do
	// not both register the same data-source type names (which the mux rejects).
	return map[string]*schema.Resource{}
}

func providerResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"citricadc_nscapacity": resourceCitrixAdcNscapacity(),
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData, terraformVersion string) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Validate required parameters
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	endpoint := d.Get("endpoint").(string)

	if username == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Missing required parameter",
			Detail:   "The 'username' parameter is required. It can be set via the provider configuration or the NS_LOGIN environment variable.",
		})
	}

	if password == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Missing required parameter",
			Detail:   "The 'password' parameter is required. It can be set via the provider configuration or the NS_PASSWORD environment variable.",
		})
	}

	if endpoint == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Missing required parameter",
			Detail:   "The 'endpoint' parameter is required. It can be set via the provider configuration or the NS_URL environment variable.",
		})
	}

	// Return early if any required parameters are missing
	if len(diags) > 0 {
		return nil, diags
	}

	userHeaders := map[string]string{
		"User-Agent": "terraform-ctxadc",
	}
	c := NetScalerNitroClient{
		Username:    username,
		Password:    password,
		Endpoint:    endpoint,
		HttpTimeout: d.Get("http_timeout").(int),
		NsTimeout:   d.Get("ns_timeout").(int),
	}

	params := service.NitroParams{
		Url:         endpoint,
		Username:    username,
		Password:    password,
		ProxiedNs:   d.Get("proxied_ns").(string),
		SslVerify:   !d.Get("insecure_skip_verify").(bool),
		Headers:     userHeaders,
		IsCloud:     d.Get("is_cloud").(bool),
		HttpTimeout: d.Get("http_timeout").(int),
		Timeout:     d.Get("ns_timeout").(int),
	}
	client, err := service.NewNitroClientFromParams(params)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Nitro client",
			Detail:   err.Error(),
		})
		return nil, diags
	}
	if d.Get("do_login").(bool) {
		client.Login()
	}
	if partition, ok := d.GetOk("partition"); ok {
		nspartition := make(map[string]interface{})
		nspartition["partitionname"] = partition.(string)
		err := client.ActOnResource("nspartition", &nspartition, "Switch")
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to switch partition",
				Detail:   err.Error(),
			})
			return nil, diags
		}
	}

	c.client = client
	log.Printf("[DEBUG] citrixadc-provider: Terraform version imported: %s\n", terraformVersion)

	return &c, diags
}
