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
	return map[string]*schema.Resource{
		"citrixadc_nsversion":                     dataSourceCitrixAdcNsversion(),
		"citrixadc_nitro_info":                    dataSourceCitrixAdcNitroInfo(),
		"citrixadc_sslcipher_sslvserver_bindings": dataSourceCitrixAdcSslcipherSslvserverBindings(),
	}
}

func providerResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"citrixadc_lbmetrictable":                    resourceCitrixAdcLbmetrictable(),
		"citrixadc_sslservice":                       resourceCitrixAdcSslservice(),
		"citrixadc_sslfipskey":                       resourceCitrixAdcSslfipskey(),
		"citrixadc_lbroute6":                         resourceCitrixAdcLbroute6(),
		"citrixadc_sslpolicylabel":                   resourceCitrixAdcSslpolicylabel(),
		"citrixadc_ssllogprofile":                    resourceCitrixAdcSsllogprofile(),
		"citrixadc_sslcacertgroup":                   resourceCitrixAdcSslcacertgroup(),
		"citrixadc_lbvserver":                        resourceCitrixAdcLbvserver(),
		"citrixadc_service":                          resourceCitrixAdcService(),
		"citrixadc_sslpolicy":                        resourceCitrixAdcSslpolicy(),
		"citrixadc_sslparameter":                     resourceCitrixAdcSslparameter(),
		"citrixadc_ssldhparam":                       resourceCitrixAdcSsldhparam(),
		"citrixadc_sslcipher":                        resourceCitrixAdcSslcipher(),
		"citrixadc_servicegroup":                     resourceCitrixAdcServicegroup(),
		"citrixadc_nsacl":                            resourceCitrixAdcNsacl(),
		"citrixadc_nsacls":                           resourceCitrixAdcNsacls(),
		"citrixadc_rnat":                             resourceCitrixAdcRnat(),
		"citrixadc_server":                           resourceCitrixAdcServer(),
		"citrixadc_responderaction":                  resourceCitrixAdcResponderaction(),
		"citrixadc_responderpolicy":                  resourceCitrixAdcResponderpolicy(),
		"citrixadc_responderpolicylabel":             resourceCitrixAdcResponderpolicylabel(),
		"citrixadc_rewriteaction":                    resourceCitrixAdcRewriteaction(),
		"citrixadc_rewritepolicy":                    resourceCitrixAdcRewritepolicy(),
		"citrixadc_rewritepolicylabel":               resourceCitrixAdcRewritepolicylabel(),
		"citrixadc_nsip":                             resourceCitrixAdcNsip(),
		"citrixadc_nsip6":                            resourceCitrixAdcNsip6(),
		"citrixadc_ipset":                            resourceCitrixAdcIpset(),
		"citrixadc_route":                            resourceCitrixAdcRoute(),
		"citrixadc_linkset":                          resourceCitrixAdcLinkset(),
		"citrixadc_nsfeature":                        resourceCitrixAdcNsfeature(),
		"citrixadc_systemgroup":                      resourceCitrixAdcSystemgroup(),
		"citrixadc_systemcmdpolicy":                  resourceCitrixAdcSystemcmdpolicy(),
		"citrixadc_nstcpprofile":                     resourceCitrixAdcNstcpprofile(),
		"citrixadc_nshttpprofile":                    resourceCitrixAdcNshttpprofile(),
		"citrixadc_nslicense":                        resourceCitrixAdcNslicense(),
		"citrixadc_systemfile":                       resourceCitrixAdcSystemfile(),
		"citrixadc_routerdynamicrouting":             resourceCitrixAdcRouterdynamicrouting(),
		"citrixadc_policyexpression":                 resourceCitrixAdcPolicyexpression(),
		"citrixadc_systemextramgmtcpu":               resourceCitrixAdcSystemextramgmtcpu(),
		"citrixadc_netprofile":                       resourceCitrixAdcNetprofile(),
		"citrixadc_nsparam":                          resourceCitrixAdcNsparam(),
		"citrixadc_policydataset":                    resourceCitrixAdcPolicydataset(),
		"citrixadc_policystringmap":                  resourceCitrixAdcPolicystringmap(),
		"citrixadc_transformprofile":                 resourceCitrixAdcTransformprofile(),
		"citrixadc_transformaction":                  resourceCitrixAdcTransformaction(),
		"citrixadc_transformpolicy":                  resourceCitrixAdcTransformpolicy(),
		"citrixadc_quicbridgeprofile":                resourceCitrixAdcQuicbridgeprofile(),
		"citrixadc_policypatset":                     resourceCitrixAdcPolicypatset(),
		"citrixadc_nsvpxparam":                       resourceCitrixAdcNsvpxparam(),
		"citrixadc_nstcpparam":                       resourceCitrixAdcNstcpparam(),
		"citrixadc_iptunnel":                         resourceCitrixAdcIptunnel(),
		"citrixadc_vlan":                             resourceCitrixAdcVlan(),
		"citrixadc_nsmode":                           resourceCitrixAdcNsmode(),
		"citrixadc_lbsipparameters":                  resourceCitrixAdcLbsipparameters(),
		"citrixadc_lbroute":                          resourceCitrixAdcLbroute(),
		"citrixadc_lbgroup":                          resourceCitrixAdcLbgroup(),
		"citrixadc_ssldtlsprofile":                   resourceCitrixAdcSsldtlsprofile(),
		"citrixadc_sslocspresponder":                 resourceCitrixAdcSslocspresponder(),
		"citrixadc_sslvserver":                       resourceCitrixAdcSslvserver(),
		"citrixadc_policyhttpcallout":                resourceCitrixAdcPolicyhttpcallout(),
		"citrixadc_policymap":                        resourceCitrixAdcPolicymap(),
		"citrixadc_policyparam":                      resourceCitrixAdcPolicyparam(),
		"citrixadc_sslservicegroup":                  resourceCitrixAdcSslservicegroup(),
		"citrixadc_rewriteparam":                     resourceCitrixAdcRewriteparam(),
		"citrixadc_responderparam":                   resourceCitrixAdcResponderparam(),
		"citrixadc_vpneula":                          resourceCitrixAdcVpneula(),
		"citrixadc_vpnclientlessaccesspolicy":        resourceCitrixAdcVpnclientlessaccesspolicy(),
		"citrixadc_vpnalwaysonprofile":               resourceCitrixAdcVpnalwaysonprofile(),
		"citrixadc_vpnformssoaction":                 resourceCitrixAdcVpnformssoaction(),
		"citrixadc_vpnclientlessaccessprofile":       resourceCitrixAdcVpnclientlessaccessprofile(),
		"citrixadc_responderhtmlpage":                resourceCitrixAdcResponderhtmlpage(),
		"citrixadc_vpnurl":                           resourceCitrixAdcVpnurl(),
		"citrixadc_vpnsessionaction":                 resourceCitrixAdcVpnsessionaction(),
		"citrixadc_vpnvserver":                       resourceCitrixAdcVpnvserver(),
		"citrixadc_vpnsessionpolicy":                 resourceCitrixAdcVpnsessionpolicy(),
		"citrixadc_vpntrafficaction":                 resourceCitrixAdcVpntrafficaction(),
		"citrixadc_vpnurlaction":                     resourceCitrixAdcVpnurlaction(),
		"citrixadc_vpntrafficpolicy":                 resourceCitrixAdcVpntrafficpolicy(),
		"citrixadc_vpnurlpolicy":                     resourceCitrixAdcVpnurlpolicy(),
		"citrixadc_vpnintranetapplication":           resourceCitrixAdcVpnintranetapplication(),
		"citrixadc_vpnpcoipvserverprofile":           resourceCitrixAdcVpnpcoipvserverprofile(),
		"citrixadc_vpnpcoipprofile":                  resourceCitrixAdcVpnpcoipprofile(),
		"citrixadc_vpnnexthopserver":                 resourceCitrixAdcVpnnexthopserver(),
		"citrixadc_vpnportaltheme":                   resourceCitrixAdcVpnportaltheme(),
		"citrixadc_vpnsamlssoprofile":                resourceCitrixAdcVpnsamlssoprofile(),
		"citrixadc_vpnparameter":                     resourceCitrixAdcVpnparameter(),
		"citrixadc_vxlan":                            resourceCitrixAdcVxlan(),
		"citrixadc_vxlanvlanmap":                     resourceCitrixAdcVxlanvlanmap(),
		"citrixadc_location":                         resourceCitrixAdcLocation(),
		"citrixadc_vrid":                             resourceCitrixAdcVrid(),
		"citrixadc_vrid6":                            resourceCitrixAdcVrid6(),
		"citrixadc_netbridge":                        resourceCitrixAdcNetbridge(),
		"citrixadc_nstimer":                          resourceCitrixAdcNstimer(),
		"citrixadc_nslimitidentifier":                resourceCitrixAdcNslimitidentifier(),
		"citrixadc_nsservicepath":                    resourceCitrixAdcNsservicepath(),
		"citrixadc_nspartition":                      resourceCitrixAdcNspartition(),
		"citrixadc_nsvariable":                       resourceCitrixAdcNsvariable(),
		"citrixadc_nsappflowcollector":               resourceCitrixAdcNsappflowcollector(),
		"citrixadc_nsicapprofile":                    resourceCitrixAdcNsicapprofile(),
		"citrixadc_nsxmlnamespace":                   resourceCitrixAdcNsxmlnamespace(),
		"citrixadc_nstrafficdomain":                  resourceCitrixAdcNstrafficdomain(),
		"citrixadc_nsservicefunction":                resourceCitrixAdcNsservicefunction(),
		"citrixadc_nssimpleacl":                      resourceCitrixAdcNssimpleacl(),
		"citrixadc_mapbmr":                           resourceCitrixAdcMapbmr(),
		"citrixadc_nssimpleacl6":                     resourceCitrixAdcNssimpleacl6(),
		"citrixadc_mapdmr":                           resourceCitrixAdcMapdmr(),
		"citrixadc_nsspparams":                       resourceCitrixAdcNsspparams(),
		"citrixadc_nsconsoleloginprompt":             resourceCitrixAdcNsconsoleloginprompt(),
		"citrixadc_locationparameter":                resourceCitrixAdcLocationparameter(),
		"citrixadc_nsdiameter":                       resourceCitrixAdcNsdiameter(),
		"citrixadc_nsdhcpparams":                     resourceCitrixAdcNsdhcpparams(),
		"citrixadc_nsassignment":                     resourceCitrixAdcNsassignment(),
		"citrixadc_nsratecontrol":                    resourceCitrixAdcNsratecontrol(),
		"citrixadc_l4param":                          resourceCitrixAdcL4param(),
		"citrixadc_rnatparam":                        resourceCitrixAdcRnatparam(),
		"citrixadc_ptp":                              resourceCitrixAdcPtp(),
		"citrixadc_nshttpparam":                      resourceCitrixAdcNshttpparam(),
		"citrixadc_mapdomain":                        resourceCitrixAdcMapdomain(),
		"citrixadc_vridparam":                        resourceCitrixAdcVridparam(),
		"citrixadc_nat64param":                       resourceCitrixAdcNat64param(),
		"citrixadc_nslicenseparameters":              resourceCitrixAdcNslicenseparameters(),
		"citrixadc_iptunnelparam":                    resourceCitrixAdcIptunnelparam(),
		"citrixadc_nsacl6":                           resourceCitrixAdcNsacl6(),
		"citrixadc_nspbr6":                           resourceCitrixAdcNspbr6(),
		"citrixadc_nstcpbufparam":                    resourceCitrixAdcNstcpbufparam(),
		"citrixadc_rsskeytype":                       resourceCitrixAdcRsskeytype(),
		"citrixadc_nat64":                            resourceCitrixAdcNat64(),
		"citrixadc_transformpolicylabel":             resourceCitrixAdcTransformpolicylabel(),
		"citrixadc_snmptrap":                         resourceCitrixAdcSnmptrap(),
		"citrixadc_snmpview":                         resourceCitrixAdcSnmpview(),
		"citrixadc_snmpgroup":                        resourceCitrixAdcSnmpgroup(),
		"citrixadc_snmpengineid":                     resourceCitrixAdcSnmpengineid(),
		"citrixadc_snmpmib":                          resourceCitrixAdcSnmpmib(),
		"citrixadc_snmpmanager":                      resourceCitrixAdcSnmpmanager(),
		"citrixadc_snmpalarm":                        resourceCitrixAdcSnmpalarm(),
		"citrixadc_ntpserver":                        resourceCitrixAdcNtpserver(),
		"citrixadc_systemparameter":                  resourceCitrixAdcSystemparameter(),
		"citrixadc_nstimeout":                        resourceCitrixAdcNstimeout(),
		"citrixadc_nscqaparam":                       resourceCitrixAdcNscqaparam(),
		"citrixadc_nshostname":                       resourceCitrixAdcNshostname(),
		"citrixadc_nslicenseproxyserver":             resourceCitrixAdcNslicenseproxyserver(),
		"citrixadc_snmpcommunity":                    resourceCitrixAdcSnmpcommunity(),
		"citrixadc_lacp":                             resourceCitrixAdcLacp(),
		"citrixadc_route6":                           resourceCitrixAdcRoute6(),
		"citrixadc_nd6":                              resourceCitrixAdcNd6(),
		"citrixadc_nspbr":                            resourceCitrixAdcNspbr(),
		"citrixadc_l3param":                          resourceCitrixAdcL3param(),
		"citrixadc_nd6ravariables":                   resourceCitrixAdcNd6ravariables(),
		"citrixadc_l2param":                          resourceCitrixAdcL2param(),
		"citrixadc_rnat6":                            resourceCitrixAdcRnat6(),
		"citrixadc_ipv6":                             resourceCitrixAdcIpv6(),
		"citrixadc_onlinkipv6prefix":                 resourceCitrixAdcOnlinkipv6prefix(),
		"citrixadc_systembackup":                     resourceCitrixAdcSystembackup(),
		"citrixadc_ntpparam":                         resourceCitrixAdcNtpparam(),
		"citrixadc_tmsessionparameter":               resourceCitrixAdcTmsessionparameter(),
		"citrixadc_tmformssoaction":                  resourceCitrixAdcTmformssoaction(),
		"citrixadc_tmsessionpolicy":                  resourceCitrixAdcTmsessionpolicy(),
		"citrixadc_tmsessionaction":                  resourceCitrixAdcTmsessionaction(),
		"citrixadc_tmtrafficpolicy":                  resourceCitrixAdcTmtrafficpolicy(),
		"citrixadc_ipsecparameter":                   resourceCitrixAdcIpsecparameter(),
		"citrixadc_tmsamlssoprofile":                 resourceCitrixAdcTmsamlssoprofile(),
		"citrixadc_userprotocol":                     resourceCitrixAdcUserprotocol(),
		"citrixadc_uservserver":                      resourceCitrixAdcUservserver(),
		"citrixadc_lldpparam":                        resourceCitrixAdcLldpparam(),
		"citrixadc_tmtrafficaction":                  resourceCitrixAdcTmtrafficaction(),
		"citrixadc_locationfile":                     resourceCitrixAdcLocationfile(),
		"citrixadc_lsnclient":                        resourceCitrixAdcLsnclient(),
		"citrixadc_lsnappsattributes":                resourceCitrixAdcLsnappsattributes(),
		"citrixadc_lsngroup":                         resourceCitrixAdcLsngroup(),
		"citrixadc_lsnappsprofile":                   resourceCitrixAdcLsnappsprofile(),
		"citrixadc_lsnlogprofile":                    resourceCitrixAdcLsnlogprofile(),
		"citrixadc_lsnpool":                          resourceCitrixAdcLsnpool(),
		"citrixadc_lsnrtspalgprofile":                resourceCitrixAdcLsnrtspalgprofile(),
		"citrixadc_lsnsipalgprofile":                 resourceCitrixAdcLsnsipalgprofile(),
		"citrixadc_ntpsync":                          resourceCitrixAdcNtpsync(),
		"citrixadc_lsntransportprofile":              resourceCitrixAdcLsntransportprofile(),
		"citrixadc_lsnip6profile":                    resourceCitrixAdcLsnip6profile(),
		"citrixadc_lsnstatic":                        resourceCitrixAdcLsnstatic(),
		"citrixadc_streamidentifier":                 resourceCitrixAdcStreamidentifier(),
		"citrixadc_pcpserver":                        resourceCitrixAdcPcpserver(),
		"citrixadc_pcpprofile":                       resourceCitrixAdcPcpprofile(),
		"citrixadc_nsweblogparam":                    resourceCitrixAdcNsweblogparam(),
		"citrixadc_spilloveraction":                  resourceCitrixAdcSpilloveraction(),
		"citrixadc_locationfile6":                    resourceCitrixAdcLocationfile6(),
		"citrixadc_subscribergxinterface":            resourceCitrixAdcSubscribergxinterface(),
		"citrixadc_smppparam":                        resourceCitrixAdcSmppparam(),
		"citrixadc_tunneltrafficpolicy":              resourceCitrixAdcTunneltrafficpolicy(),
		"citrixadc_subscriberprofile":                resourceCitrixAdcSubscriberprofile(),
		"citrixadc_subscriberradiusinterface":        resourceCitrixAdcSubscriberradiusinterface(),
		"citrixadc_subscriberparam":                  resourceCitrixAdcSubscriberparam(),
		"citrixadc_lsnhttphdrlogprofile":             resourceCitrixAdcLsnhttphdrlogprofile(),
		"citrixadc_streamselector":                   resourceCitrixAdcStreamselector(),
		"citrixadc_lsnparameter":                     resourceCitrixAdcLsnparameter(),
		"citrixadc_sslcertfile":                      resourceCitrixAdcSslcertfile(),
		"citrixadc_nspbrs":                           resourceCitrixAdcNspbrs(),
		"citrixadc_spilloverpolicy":                  resourceCitrixAdcSpilloverpolicy(),
		"citrixadc_videooptimizationdetectionaction": resourceCitrixAdcVideooptimizationdetectionaction(),
		"citrixadc_videooptimizationdetectionpolicy": resourceCitrixAdcVideooptimizationdetectionpolicy(),
		"citrixadc_videooptimizationpacingaction":    resourceCitrixAdcVideooptimizationpacingaction(),
		"citrixadc_videooptimizationpacingpolicy":    resourceCitrixAdcVideooptimizationpacingpolicy(),
		"citrixadc_lbaction":                         resourceCitrixAdcLbaction(),
		"citrixadc_lbpolicy":                         resourceCitrixAdcLbpolicy(),
		"citricadc_nscapacity":                       resourceCitrixAdcNscapacity(),
		"citrixadc_nscapacity":                       resourceCitrixAdcNscapacity(),
		"citrixadc_nslicenseserver":                  resourceCitrixAdcNslicenseserver(),
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
