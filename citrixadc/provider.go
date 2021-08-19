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

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/hashicorp/terraform/version"
)

type NetScalerNitroClient struct {
	Username string
	Password string
	Endpoint string
	client   *service.NitroClient
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
		"citrixadc_botsettings":                             resourceCitrixAdcBotsettings(),
		"citrixadc_botpolicy":                               resourceCitrixAdcBotpolicy(),
		"citrixadc_lbmonitor_metric_binding":                resourceCitrixAdcLbmonitor_metric_binding(),
		"citrixadc_lbvserver":                               resourceCitrixAdcLbvserver(),
		"citrixadc_service":                                 resourceCitrixAdcService(),
		"citrixadc_csvserver":                               resourceCitrixAdcCsvserver(),
		"citrixadc_cspolicy":                                resourceCitrixAdcCspolicy(),
		"citrixadc_csaction":                                resourceCitrixAdcCsaction(),
		"citrixadc_sslaction":                               resourceCitrixAdcSslaction(),
		"citrixadc_sslpolicy":                               resourceCitrixAdcSslpolicy(),
		"citrixadc_sslcertkey":                              resourceCitrixAdcSslcertkey(),
		"citrixadc_sslprofile":                              resourceCitrixAdcSslprofile(),
		"citrixadc_sslparameter":                            resourceCitrixAdcSslparameter(),
		"citrixadc_ssldhparam":                              resourceCitrixAdcSsldhparam(),
		"citrixadc_sslcipher":                               resourceCitrixAdcSslcipher(),
		"citrixadc_lbmonitor":                               resourceCitrixAdcLbmonitor(),
		"citrixadc_servicegroup":                            resourceCitrixAdcServicegroup(),
		"citrixadc_nsacl":                                   resourceCitrixAdcNsacl(),
		"citrixadc_nsacls":                                  resourceCitrixAdcNsacls(),
		"citrixadc_inat":                                    resourceCitrixAdcInat(),
		"citrixadc_rnat":                                    resourceCitrixAdcRnats(),
		"citrixadc_gslbsite":                                resourceCitrixAdcGslbsite(),
		"citrixadc_gslbvserver":                             resourceCitrixAdcGslbvserver(),
		"citrixadc_gslbservice":                             resourceCitrixAdcGslbservice(),
		"citrixadc_server":                                  resourceCitrixAdcServer(),
		"citrixadc_responderaction":                         resourceCitrixAdcResponderaction(),
		"citrixadc_responderpolicy":                         resourceCitrixAdcResponderpolicy(),
		"citrixadc_responderpolicylabel":                    resourceCitrixAdcResponderpolicylabel(),
		"citrixadc_rewriteaction":                           resourceCitrixAdcRewriteaction(),
		"citrixadc_rewritepolicy":                           resourceCitrixAdcRewritepolicy(),
		"citrixadc_rewritepolicylabel":                      resourceCitrixAdcRewritepolicylabel(),
		"citrixadc_nsip":                                    resourceCitrixAdcNsip(),
		"citrixadc_nsip6":                                   resourceCitrixAdcNsip6(),
		"citrixadc_nsconfig_save":                           resourceCitrixAdcNsconfigSave(),
		"citrixadc_nsconfig_clear":                          resourceCitrixAdcNsconfigClear(),
		"citrixadc_nsconfig_update":                         resourceCitrixAdcNsconfigUpdate(),
		"citrixadc_ipset":                                   resourceCitrixAdcIpset(),
		"citrixadc_route":                                   resourceCitrixAdcRoute(),
		"citrixadc_linkset":                                 resourceCitrixAdcLinkset(),
		"citrixadc_nsfeature":                               resourceCitrixAdcNsfeature(),
		"citrixadc_systemuser":                              resourceCitrixAdcSystemuser(),
		"citrixadc_systemgroup":                             resourceCitrixAdcSystemgroup(),
		"citrixadc_systemcmdpolicy":                         resourceCitrixAdcSystemcmdpolicy(),
		"citrixadc_interface":                               resourceCitrixAdcInterface(),
		"citrixadc_nstcpprofile":                            resourceCitrixAdcNstcpprofile(),
		"citrixadc_nshttpprofile":                           resourceCitrixAdcNshttpprofile(),
		"citrixadc_nslicense":                               resourceCitrixAdcNslicense(),
		"citrixadc_cluster":                                 resourceCitrixAdcCluster(),
		"citrixadc_clusterfiles_syncer":                     resourceCitrixAdcClusterfilesSyncer(),
		"citrixadc_systemfile":                              resourceCitrixAdcSystemfile(),
		"citrixadc_auditmessageaction":                      resourceCitrixAdcAuditmessageaction(),
		"citrixadc_auditsyslogaction":                       resourceCitrixAdcAuditsyslogaction(),
		"citrixadc_auditsyslogpolicy":                       resourceCitrixAdcAuditsyslogpolicy(),
		"citrixadc_rebooter":                                resourceCitrixAdcRebooter(),
		"citrixadc_installer":                               resourceCitrixAdcInstaller(),
		"citrixadc_pinger":                                  resourceCitrixAdcPinger(),
		"citrixadc_nsrpcnode":                               resourceCitrixAdcNsrpcnode(),
		"citrixadc_routerdynamicrouting":                    resourceCitrixAdcRouterdynamicrouting(),
		"citrixadc_policyexpression":                        resourceCitrixAdcPolicyexpression(),
		"citrixadc_systemextramgmtcpu":                      resourceCitrixAdcSystemextramgmtcpu(),
		"citrixadc_netprofile":                              resourceCitrixAdcNetprofile(),
		"citrixadc_servicegroup_lbmonitor_binding":          resourceCitrixAdcServicegroup_lbmonitor_binding(),
		"citrixadc_nsparam":                                 resourceCitrixAdcNsparam(),
		"citrixadc_sslvserver_sslpolicy_binding":            resourceCitrixAdcSslvserver_sslpolicy_binding(),
		"citrixadc_sslprofile_sslcipher_binding":            resourceCitrixAdcSslprofile_sslcipher_binding(),
		"citrixadc_policydataset":                           resourceCitrixAdcPolicydataset(),
		"citrixadc_policydataset_value_binding":             resourceCitrixAdcPolicydataset_value_binding(),
		"citrixadc_password_resetter":                       resourceCitrixAdcPasswordResetter(),
		"citrixadc_csvserver_cspolicy_binding":              resourceCitrixAdcCsvserver_cspolicy_binding(),
		"citrixadc_appfwprofile":                            resourceCitrixAdcAppfwprofile(),
		"citrixadc_appfwpolicy":                             resourceCitrixAdcAppfwpolicy(),
		"citrixadc_appfwfieldtype":                          resourceCitrixAdcAppfwfieldtype(),
		"citrixadc_appfwpolicylabel":                        resourceCitrixAdcAppfwpolicylabel(),
		"citrixadc_appfwjsoncontenttype":                    resourceCitrixAdcAppfwjsoncontenttype(),
		"citrixadc_appfwxmlcontenttype":                     resourceCitrixAdcAppfwxmlcontenttype(),
		"citrixadc_appfwprofile_starturl_binding":           resourceCitrixAdcAppfwprofileStarturlBinding(),
		"citrixadc_appfwprofile_denyurl_binding":            resourceCitrixAdcAppfwprofileDenyurlBinding(),
		"citrixadc_nslicenseserver":                         resourceCitrixAdcNslicenseserver(),
		"citrixadc_nscapacity":                              resourceCitrixAdcNscapacity(),
		"citrixadc_lbvserver_service_binding":               resourceCitrixAdcLbvserver_service_binding(),
		"citrixadc_policystringmap":                         resourceCitrixAdcPolicystringmap(),
		"citrixadc_policystringmap_pattern_binding":         resourceCitrixAdcPolicystringmap_pattern_binding(),
		"citrixadc_transformprofile":                        resourceCitrixAdcTransformprofile(),
		"citrixadc_transformaction":                         resourceCitrixAdcTransformaction(),
		"citrixadc_transformpolicy":                         resourceCitrixAdcTransformpolicy(),
		"citrixadc_lbvserver_transformpolicy_binding":       resourceCitrixAdcLbvserver_transformpolicy_binding(),
		"citrixadc_csvserver_transformpolicy_binding":       resourceCitrixAdcCsvserver_transformpolicy_binding(),
		"citrixadc_sslvserver_sslcertkey_binding":           resourceCitrixAdcSslvserver_sslcertkey_binding(),
		"citrixadc_servicegroup_servicegroupmember_binding": resourceCitrixAdcServicegroup_servicegroupmember_binding(),
		"citrixadc_quicbridgeprofile":                       resourceCitrixAdcQuicbridgeprofile(),
		"citrixadc_policypatset":                            resourceCitrixAdcPolicypatset(),
		"citrixadc_policypatset_pattern_binding":            resourceCitrixAdcPolicypatset_pattern_binding(),
		"citrixadc_filterpolicy":                            resourceCitrixAdcFilterpolicy(),
		"citrixadc_lbvserver_filterpolicy_binding":          resourceCitrixAdcLbvserver_filterpolicy_binding(),
		"citrixadc_csvserver_filterpolicy_binding":          resourceCitrixAdcCsvserver_filterpolicy_binding(),
		"citrixadc_cmppolicy":                               resourceCitrixAdcCmppolicy(),
		"citrixadc_lbvserver_cmppolicy_binding":             resourceCitrixAdcLbvserver_cmppolicy_binding(),
		"citrixadc_csvserver_cmppolicy_binding":             resourceCitrixAdcCsvserver_cmppolicy_binding(),
		"citrixadc_lbvserver_responderpolicy_binding":       resourceCitrixAdcLbvserver_responderpolicy_binding(),
		"citrixadc_csvserver_responderpolicy_binding":       resourceCitrixAdcCsvserver_responderpolicy_binding(),
		"citrixadc_csvserver_rewritepolicy_binding":         resourceCitrixAdcCsvserver_rewritepolicy_binding(),
		"citrixadc_lbvserver_rewritepolicy_binding":         resourceCitrixAdcLbvserver_rewritepolicy_binding(),
		"citrixadc_nsvpxparam":                              resourceCitrixAdcNsvpxparam(),
		"citrixadc_nstcpparam":                              resourceCitrixAdcNstcpparam(),
		"citrixadc_dnsnsrec":                                resourceCitrixAdcDnsnsrec(),
		"citrixadc_dnssoarec":                               resourceCitrixAdcDnssoarec(),
		"citrixadc_lbvserver_appfwpolicy_binding":           resourceCitrixAdcLbvserver_appfwpolicy_binding(),
		"citrixadc_csvserver_appfwpolicy_binding":           resourceCitrixAdcCsvserver_appfwpolicy_binding(),
		"citrixadc_appfwprofile_cookieconsistency_binding":  resourceCitrixAdcAppfwprofile_cookieconsistency_binding(),
		"citrixadc_appfwprofile_crosssitescripting_binding": resourceCitrixAdcAppfwprofile_crosssitescripting_binding(),
		"citrixadc_appfwprofile_sqlinjection_binding":       resourceCitrixAdcAppfwprofile_sqlinjection_binding(),
		"citrixadc_lbvserver_servicegroup_binding":          resourceCitrixAdcLbvserver_servicegroup_binding(),
		"citrixadc_lbparameter":                             resourceCitrixAdcLbparameter(),
		"citrixadc_iptunnel":                                resourceCitrixAdcIptunnel(),
		"citrixadc_vlan":                                    resourceCitrixAdcVlan(),
		"citrixadc_vlan_interface_binding":                  resourceCitrixAdcVlan_interface_binding(),
		"citrixadc_vlan_nsip_binding":                       resourceCitrixAdcVlan_nsip_binding(),
		"citrixadc_nsmode":                                  resourceCitrixAdcNsmode(),
		"citrixadc_botprofile":                              resourceCitrixAdcBotprofile(),
		"citrixadc_botpolicylabel":                          resourceCitrixAdcBotpolicylabel(),
		"citrixadc_lbsipparameters":                         resourceCitrixAdcLbsipparameters(),
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	userHeaders := map[string]string{
		"User-Agent": "terraform-ctxadc",
	}
	c := NetScalerNitroClient{
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
		Endpoint: d.Get("endpoint").(string),
	}

	params := service.NitroParams{
		Url:       d.Get("endpoint").(string),
		Username:  d.Get("username").(string),
		Password:  d.Get("password").(string),
		ProxiedNs: d.Get("proxied_ns").(string),
		SslVerify: !d.Get("insecure_skip_verify").(bool),
		Headers:   userHeaders,
	}
	client, err := service.NewNitroClientFromParams(params)
	if err != nil {
		return nil, err
	}

	c.client = client
	log.Printf("[DEBUG] citrixadc-provider: Terraform version imported: %s\n", version.Version)

	return &c, nil
}
