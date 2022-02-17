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

	"github.com/citrix/adc-nitro-go/resource/config/ns"
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
		"do_login": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "Perform the login operation and acquire a session token to be used for subsequent requests.",
			DefaultFunc: schema.EnvDefaultFunc("NS_DO_LOGIN", false),
		},
		"partition": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Partition to target",
			DefaultFunc: schema.EnvDefaultFunc("NS_PARTITION", nil),
		},
	}
}

func providerResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"citrixadc_lbmetrictable":                                      resourceCitrixAdcLbmetrictable(),
		"citrixadc_sslservice_sslcertkey_binding":                      resourceCitrixAdcSslservice_sslcertkey_binding(),
		"citrixadc_sslservice_sslciphersuite_binding":                  resourceCitrixAdcSslservice_sslciphersuite_binding(),
		"citrixadc_sslservice_ecccurve_binding":                        resourceCitrixAdcSslservice_ecccurve_binding(),
		"citrixadc_sslservice":                                         resourceCitrixAdcSslservice(),
		"citrixadc_sslfipskey":                                         resourceCitrixAdcSslfipskey(),
		"citrixadc_sslcacertgroup_sslcertkey_binding":                  resourceCitrixAdcSslcacertgroup_sslcertkey_binding(),
		"citrixadc_lbroute6":                                           resourceCitrixAdcLbroute6(),
		"citrixadc_sslpolicylabel":                                     resourceCitrixAdcSslpolicylabel(),
		"citrixadc_ssllogprofile":                                      resourceCitrixAdcSsllogprofile(),
		"citrixadc_sslprofile_sslcertkey_binding":                      resourceCitrixAdcSslprofile_sslcertkey_binding(),
		"citrixadc_sslpolicylabel_sslpolicy_binding":                   resourceCitrixAdcSslpolicylabel_sslpolicy_binding(),
		"citrixadc_lbvserver_botpolicy_binding":                        resourceCitrixAdcLbvserver_botpolicy_binding(),
		"citrixadc_lbvserver_auditsyslogpolicy_binding":                resourceCitrixAdcLbvserver_auditsyslogpolicy_binding(),
		"citrixadc_sslcacertgroup":                                     resourceCitrixAdcSslcacertgroup(),
		"citrixadc_botsettings":                                        resourceCitrixAdcBotsettings(),
		"citrixadc_botpolicy":                                          resourceCitrixAdcBotpolicy(),
		"citrixadc_lbvserver_analyticsprofile_binding":                 resourceCitrixAdcLbvserver_analyticsprofile_binding(),
		"citrixadc_lbvserver_appqoepolicy_binding":                     resourceCitrixAdcLbvserver_appqoepolicy_binding(),
		"citrixadc_lbmonitor_metric_binding":                           resourceCitrixAdcLbmonitor_metric_binding(),
		"citrixadc_lbvserver":                                          resourceCitrixAdcLbvserver(),
		"citrixadc_service":                                            resourceCitrixAdcService(),
		"citrixadc_csvserver":                                          resourceCitrixAdcCsvserver(),
		"citrixadc_cspolicy":                                           resourceCitrixAdcCspolicy(),
		"citrixadc_csaction":                                           resourceCitrixAdcCsaction(),
		"citrixadc_sslaction":                                          resourceCitrixAdcSslaction(),
		"citrixadc_sslpolicy":                                          resourceCitrixAdcSslpolicy(),
		"citrixadc_sslcertkey":                                         resourceCitrixAdcSslcertkey(),
		"citrixadc_sslprofile":                                         resourceCitrixAdcSslprofile(),
		"citrixadc_sslparameter":                                       resourceCitrixAdcSslparameter(),
		"citrixadc_ssldhparam":                                         resourceCitrixAdcSsldhparam(),
		"citrixadc_sslcipher":                                          resourceCitrixAdcSslcipher(),
		"citrixadc_lbmonitor":                                          resourceCitrixAdcLbmonitor(),
		"citrixadc_servicegroup":                                       resourceCitrixAdcServicegroup(),
		"citrixadc_nsacl":                                              resourceCitrixAdcNsacl(),
		"citrixadc_nsacls":                                             resourceCitrixAdcNsacls(),
		"citrixadc_inat":                                               resourceCitrixAdcInat(),
		"citrixadc_rnat":                                               resourceCitrixAdcRnats(),
		"citrixadc_gslbsite":                                           resourceCitrixAdcGslbsite(),
		"citrixadc_gslbvserver":                                        resourceCitrixAdcGslbvserver(),
		"citrixadc_gslbservice":                                        resourceCitrixAdcGslbservice(),
		"citrixadc_server":                                             resourceCitrixAdcServer(),
		"citrixadc_responderaction":                                    resourceCitrixAdcResponderaction(),
		"citrixadc_responderpolicy":                                    resourceCitrixAdcResponderpolicy(),
		"citrixadc_responderpolicylabel":                               resourceCitrixAdcResponderpolicylabel(),
		"citrixadc_rewriteaction":                                      resourceCitrixAdcRewriteaction(),
		"citrixadc_rewritepolicy":                                      resourceCitrixAdcRewritepolicy(),
		"citrixadc_rewritepolicylabel":                                 resourceCitrixAdcRewritepolicylabel(),
		"citrixadc_nsip":                                               resourceCitrixAdcNsip(),
		"citrixadc_nsip6":                                              resourceCitrixAdcNsip6(),
		"citrixadc_nsconfig_save":                                      resourceCitrixAdcNsconfigSave(),
		"citrixadc_nsconfig_clear":                                     resourceCitrixAdcNsconfigClear(),
		"citrixadc_nsconfig_update":                                    resourceCitrixAdcNsconfigUpdate(),
		"citrixadc_ipset":                                              resourceCitrixAdcIpset(),
		"citrixadc_route":                                              resourceCitrixAdcRoute(),
		"citrixadc_linkset":                                            resourceCitrixAdcLinkset(),
		"citrixadc_nsfeature":                                          resourceCitrixAdcNsfeature(),
		"citrixadc_systemuser":                                         resourceCitrixAdcSystemuser(),
		"citrixadc_systemgroup":                                        resourceCitrixAdcSystemgroup(),
		"citrixadc_systemcmdpolicy":                                    resourceCitrixAdcSystemcmdpolicy(),
		"citrixadc_interface":                                          resourceCitrixAdcInterface(),
		"citrixadc_nstcpprofile":                                       resourceCitrixAdcNstcpprofile(),
		"citrixadc_nshttpprofile":                                      resourceCitrixAdcNshttpprofile(),
		"citrixadc_nslicense":                                          resourceCitrixAdcNslicense(),
		"citrixadc_cluster":                                            resourceCitrixAdcCluster(),
		"citrixadc_clusterfiles_syncer":                                resourceCitrixAdcClusterfilesSyncer(),
		"citrixadc_systemfile":                                         resourceCitrixAdcSystemfile(),
		"citrixadc_auditmessageaction":                                 resourceCitrixAdcAuditmessageaction(),
		"citrixadc_auditsyslogaction":                                  resourceCitrixAdcAuditsyslogaction(),
		"citrixadc_auditsyslogpolicy":                                  resourceCitrixAdcAuditsyslogpolicy(),
		"citrixadc_rebooter":                                           resourceCitrixAdcRebooter(),
		"citrixadc_installer":                                          resourceCitrixAdcInstaller(),
		"citrixadc_pinger":                                             resourceCitrixAdcPinger(),
		"citrixadc_nsrpcnode":                                          resourceCitrixAdcNsrpcnode(),
		"citrixadc_routerdynamicrouting":                               resourceCitrixAdcRouterdynamicrouting(),
		"citrixadc_policyexpression":                                   resourceCitrixAdcPolicyexpression(),
		"citrixadc_systemextramgmtcpu":                                 resourceCitrixAdcSystemextramgmtcpu(),
		"citrixadc_netprofile":                                         resourceCitrixAdcNetprofile(),
		"citrixadc_servicegroup_lbmonitor_binding":                     resourceCitrixAdcServicegroup_lbmonitor_binding(),
		"citrixadc_nsparam":                                            resourceCitrixAdcNsparam(),
		"citrixadc_sslvserver_sslpolicy_binding":                       resourceCitrixAdcSslvserver_sslpolicy_binding(),
		"citrixadc_sslprofile_sslcipher_binding":                       resourceCitrixAdcSslprofile_sslcipher_binding(),
		"citrixadc_policydataset":                                      resourceCitrixAdcPolicydataset(),
		"citrixadc_policydataset_value_binding":                        resourceCitrixAdcPolicydataset_value_binding(),
		"citrixadc_password_resetter":                                  resourceCitrixAdcPasswordResetter(),
		"citrixadc_csvserver_cspolicy_binding":                         resourceCitrixAdcCsvserver_cspolicy_binding(),
		"citrixadc_appfwprofile":                                       resourceCitrixAdcAppfwprofile(),
		"citrixadc_appfwpolicy":                                        resourceCitrixAdcAppfwpolicy(),
		"citrixadc_appfwfieldtype":                                     resourceCitrixAdcAppfwfieldtype(),
		"citrixadc_appfwpolicylabel":                                   resourceCitrixAdcAppfwpolicylabel(),
		"citrixadc_appfwjsoncontenttype":                               resourceCitrixAdcAppfwjsoncontenttype(),
		"citrixadc_appfwxmlcontenttype":                                resourceCitrixAdcAppfwxmlcontenttype(),
		"citrixadc_appfwprofile_starturl_binding":                      resourceCitrixAdcAppfwprofileStarturlBinding(),
		"citrixadc_appfwprofile_denyurl_binding":                       resourceCitrixAdcAppfwprofileDenyurlBinding(),
		"citrixadc_nslicenseserver":                                    resourceCitrixAdcNslicenseserver(),
		"citrixadc_nscapacity":                                         resourceCitrixAdcNscapacity(),
		"citrixadc_lbvserver_service_binding":                          resourceCitrixAdcLbvserver_service_binding(),
		"citrixadc_policystringmap":                                    resourceCitrixAdcPolicystringmap(),
		"citrixadc_policystringmap_pattern_binding":                    resourceCitrixAdcPolicystringmap_pattern_binding(),
		"citrixadc_transformprofile":                                   resourceCitrixAdcTransformprofile(),
		"citrixadc_transformaction":                                    resourceCitrixAdcTransformaction(),
		"citrixadc_transformpolicy":                                    resourceCitrixAdcTransformpolicy(),
		"citrixadc_lbvserver_transformpolicy_binding":                  resourceCitrixAdcLbvserver_transformpolicy_binding(),
		"citrixadc_csvserver_transformpolicy_binding":                  resourceCitrixAdcCsvserver_transformpolicy_binding(),
		"citrixadc_sslvserver_sslcertkey_binding":                      resourceCitrixAdcSslvserver_sslcertkey_binding(),
		"citrixadc_servicegroup_servicegroupmember_binding":            resourceCitrixAdcServicegroup_servicegroupmember_binding(),
		"citrixadc_quicbridgeprofile":                                  resourceCitrixAdcQuicbridgeprofile(),
		"citrixadc_policypatset":                                       resourceCitrixAdcPolicypatset(),
		"citrixadc_policypatset_pattern_binding":                       resourceCitrixAdcPolicypatset_pattern_binding(),
		"citrixadc_filterpolicy":                                       resourceCitrixAdcFilterpolicy(),
		"citrixadc_lbvserver_filterpolicy_binding":                     resourceCitrixAdcLbvserver_filterpolicy_binding(),
		"citrixadc_csvserver_filterpolicy_binding":                     resourceCitrixAdcCsvserver_filterpolicy_binding(),
		"citrixadc_cmppolicy":                                          resourceCitrixAdcCmppolicy(),
		"citrixadc_lbvserver_cmppolicy_binding":                        resourceCitrixAdcLbvserver_cmppolicy_binding(),
		"citrixadc_csvserver_cmppolicy_binding":                        resourceCitrixAdcCsvserver_cmppolicy_binding(),
		"citrixadc_lbvserver_responderpolicy_binding":                  resourceCitrixAdcLbvserver_responderpolicy_binding(),
		"citrixadc_csvserver_responderpolicy_binding":                  resourceCitrixAdcCsvserver_responderpolicy_binding(),
		"citrixadc_csvserver_rewritepolicy_binding":                    resourceCitrixAdcCsvserver_rewritepolicy_binding(),
		"citrixadc_lbvserver_rewritepolicy_binding":                    resourceCitrixAdcLbvserver_rewritepolicy_binding(),
		"citrixadc_nsvpxparam":                                         resourceCitrixAdcNsvpxparam(),
		"citrixadc_nstcpparam":                                         resourceCitrixAdcNstcpparam(),
		"citrixadc_dnsnsrec":                                           resourceCitrixAdcDnsnsrec(),
		"citrixadc_dnssoarec":                                          resourceCitrixAdcDnssoarec(),
		"citrixadc_lbvserver_appfwpolicy_binding":                      resourceCitrixAdcLbvserver_appfwpolicy_binding(),
		"citrixadc_csvserver_appfwpolicy_binding":                      resourceCitrixAdcCsvserver_appfwpolicy_binding(),
		"citrixadc_appfwprofile_cookieconsistency_binding":             resourceCitrixAdcAppfwprofile_cookieconsistency_binding(),
		"citrixadc_appfwprofile_crosssitescripting_binding":            resourceCitrixAdcAppfwprofile_crosssitescripting_binding(),
		"citrixadc_appfwprofile_sqlinjection_binding":                  resourceCitrixAdcAppfwprofile_sqlinjection_binding(),
		"citrixadc_lbvserver_servicegroup_binding":                     resourceCitrixAdcLbvserver_servicegroup_binding(),
		"citrixadc_lbparameter":                                        resourceCitrixAdcLbparameter(),
		"citrixadc_iptunnel":                                           resourceCitrixAdcIptunnel(),
		"citrixadc_vlan":                                               resourceCitrixAdcVlan(),
		"citrixadc_vlan_interface_binding":                             resourceCitrixAdcVlan_interface_binding(),
		"citrixadc_vlan_nsip_binding":                                  resourceCitrixAdcVlan_nsip_binding(),
		"citrixadc_nsmode":                                             resourceCitrixAdcNsmode(),
		"citrixadc_botprofile":                                         resourceCitrixAdcBotprofile(),
		"citrixadc_botpolicylabel":                                     resourceCitrixAdcBotpolicylabel(),
		"citrixadc_lbsipparameters":                                    resourceCitrixAdcLbsipparameters(),
		"citrixadc_lbprofile":                                          resourceCitrixAdcLbprofile(),
		"citrixadc_lbvserver_contentinspectionpolicy_binding":          resourceCitrixAdcLbvserver_contentinspectionpolicy_binding(),
		"citrixadc_lbvserver_tmtrafficpolicy_binding":                  resourceCitrixAdcLbvserver_tmtrafficpolicy_binding(),
		"citrixadc_lbvserver_feopolicy_binding":                        resourceCitrixAdcLbvserver_feopolicy_binding(),
		"citrixadc_lbvserver_videooptimizationdetectionpolicy_binding": resourceCitrixAdcLbvserver_videooptimizationdetectionpolicy_binding(),
		"citrixadc_lbvserver_videooptimizationpacingpolicy_binding":    resourceCitrixAdcLbvserver_videooptimizationpacingpolicy_binding(),
		"citrixadc_lbvserver_spilloverpolicy_binding":                  resourceCitrixAdcLbvserver_spilloverpolicy_binding(),
		"citrixadc_lbvserver_dnspolicy64_binding":                      resourceCitrixAdcLbvserver_dnspolicy64_binding(),
		"citrixadc_lbmonitor_sslcertkey_binding":                       resourceCitrixAdcLbmonitor_sslcertkey_binding(),
		"citrixadc_lbvserver_authorizationpolicy_binding":              resourceCitrixAdcLbvserver_authorizationpolicy_binding(),
		"citrixadc_lbvserver_appflowpolicy_binding":                    resourceCitrixAdcLbvserver_appflowpolicy_binding(),
		"citrixadc_lbvserver_cachepolicy_binding":                      resourceCitrixAdcLbvserver_cachepolicy_binding(),
		"citrixadc_lbroute":                                            resourceCitrixAdcLbroute(),
		"citrixadc_lbgroup":                                            resourceCitrixAdcLbgroup(),
		"citrixadc_lbgroup_lbvserver_binding":                          resourceCitrixAdcLbgroup_lbvserver_binding(),
		"citrixadc_ssldtlsprofile":                                     resourceCitrixAdcSsldtlsprofile(),
		"citrixadc_sslocspresponder":                                   resourceCitrixAdcSslocspresponder(),
		"citrixadc_csparameter":                                        resourceCitrixAdcCsparameter(),
		"citrixadc_cspolicylabel":                                      resourceCitrixAdcCspolicylabel(),
		"citrixadc_csvserver_analyticsprofile_binding":                 resourceCitrixAdcCsvserver_analyticsprofile_binding(),
		"citrixadc_csvserver_appqoepolicy_binding":                     resourceCitrixAdcCsvserver_appqoepolicy_binding(),
		"citrixadc_sslvserver":                                         resourceCitrixAdcSslvserver(),
		"citrixadc_sslservicegroup_sslcertkey_binding":                 resourceCitrixAdcSslservicegroup_sslcertkey_binding(),
		"citrixadc_sslvserver_sslciphersuite_binding":                  resourceCitrixAdcSslvserver_sslciphersuite_binding(),
		"citrixadc_csvserver_botpolicy_binding":                        resourceCitrixAdcCsvserver_botpolicy_binding(),
		"citrixadc_csvserver_auditsyslogpolicy_binding":                resourceCitrixAdcCsvserver_auditsyslogpolicy_binding(),
		"citrixadc_csvserver_auditnslogpolicy_binding":                 resourceCitrixAdcCsvserver_auditnslogpolicy_binding(),
		"citrixadc_csvserver_authorizationpolicy_binding":              resourceCitrixAdcCsvserver_authorizationpolicy_binding(),
		"citrixadc_csvserver_cachepolicy_binding":                      resourceCitrixAdcCsvserver_cachepolicy_binding(),
		"citrixadc_csvserver_contentinspectionpolicy_binding":          resourceCitrixAdcCsvserver_contentinspectionpolicy_binding(),
		"citrixadc_csvserver_feopolicy_binding":                        resourceCitrixAdcCsvserver_feopolicy_binding(),
		"citrixadc_csvserver_gslbvserver_binding":                      resourceCitrixAdcCsvserver_gslbvserver_binding(),
		"citrixadc_csvserver_spilloverpolicy_binding":                  resourceCitrixAdcCsvserver_spilloverpolicy_binding(),
		"citrixadc_csvserver_tmtrafficpolicy_binding":                  resourceCitrixAdcCsvserver_tmtrafficpolicy_binding(),
		"citrixadc_csvserver_vpnvserver_binding":                       resourceCitrixAdcCsvserver_vpnvserver_binding(),
		"citrixadc_policyhttpcallout":                                  resourceCitrixAdcPolicyhttpcallout(),
		"citrixadc_policymap":                                          resourceCitrixAdcPolicymap(),
		"citrixadc_policyparam":                                        resourceCitrixAdcPolicyparam(),
		"citrixadc_sslservicegroup":                                    resourceCitrixAdcSslservicegroup(),
		"citrixadc_rewriteparam":                                       resourceCitrixAdcRewriteparam(),
		"citrixadc_rewritepolicylabel_rewritepolicy_binding":           resourceCitrixAdcRewritepolicylabel_rewritepolicy_binding(),
		"citrixadc_sslservicegroup_ecccurve_binding":                   resourceCitrixAdcSslservicegroup_ecccurve_binding(),
		"citrixadc_sslvserver_ecccurve_binding":                        resourceCitrixAdcSslvserver_ecccurve_binding(),
		"citrixadc_responderpolicylabel_responderpolicy_binding":       resourceCitrixAdcResponderpolicylabel_responderpolicy_binding(),
		"citrixadc_responderparam":                                     resourceCitrixAdcResponderparam(),
		"citrixadc_vpneula":                                            resourceCitrixAdcVpneula(),
		"citrixadc_vpnclientlessaccesspolicy":                          resourceCitrixAdcVpnclientlessaccesspolicy(),
		"citrixadc_vpnalwaysonprofile":                                 resourceCitrixAdcVpnalwaysonprofile(),
		"citrixadc_rewriteglobal_rewritepolicy_binding":                resourceCitrixAdcRewriteglobal_rewritepolicy_binding(),
		"citrixadc_vpnformssoaction":                                   resourceCitrixAdcVpnformssoaction(),
		"citrixadc_vpnglobal_appcontroller_binding":                    resourceCitrixAdcVpnglobal_appcontroller_binding(),
		"citrixadc_vpnclientlessaccessprofile":                         resourceCitrixAdcVpnclientlessaccessprofile(),
		"citrixadc_filterglobal_filterpolicy_binding":                  resourceCitrixAdcFilterglobal_filterpolicy_binding(),
		"citrixadc_dnsparameter":                                       resourceCitrixAdcDnsparameter(),
		"citrixadc_appfwsettings":                                      resourceCitrixAdcAppfwsettings(),
		"citrixadc_responderhtmlpage":                                  resourceCitrixAdcResponderhtmlpage(),
		"citrixadc_authorizationpolicy":                                resourceCitrixAdcAuthorizationpolicy(),
		"citrixadc_vpnurl":                                             resourceCitrixAdcVpnurl(),
		"citrixadc_vpnsessionaction":                                   resourceCitrixAdcVpnsessionaction(),
		"citrixadc_vpnvserver":                                         resourceCitrixAdcVpnvserver(),
		"citrixadc_vpnsessionpolicy":                                   resourceCitrixAdcVpnsessionpolicy(),
		"citrixadc_vpntrafficaction":                                   resourceCitrixAdcVpntrafficaction(),
		"citrixadc_vpnurlaction":                                       resourceCitrixAdcVpnurlaction(),
		"citrixadc_vpnvserver_vpnsessionpolicy_binding":                resourceCitrixAdcVpnvserver_vpnsessionpolicy_binding(),
		"citrixadc_vpntrafficpolicy":                                   resourceCitrixAdcVpntrafficpolicy(),
		"citrixadc_vpnurlpolicy":                                       resourceCitrixAdcVpnurlpolicy(),
		"citrixadc_authenticationvserver":                              resourceCitrixAdcAuthenticationvserver(),
		"citrixadc_authenticationldapaction":                           resourceCitrixAdcAuthenticationldapaction(),
		"citrixadc_vpnglobal_sslcertkey_binding":                       resourceCitrixAdcVpnglobal_sslcertkey_binding(),
		"citrixadc_vpngobal_vpntrafficpolicy_binding":                  resourceCitrixAdcVpnglobal_vpntrafficpolicy_binding(),
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

	if d.Get("do_login").(bool) {
		client.Login()
	}

	if partition, ok := d.GetOk("partition"); ok {
		nspartition := ns.Nspartition{
			Partitionname: partition.(string),
		}
		err := client.ActOnResource("nspartition", &nspartition, "Switch")
		if err != nil {
			return nil, err
		}
	}

	c.client = client
	log.Printf("[DEBUG] citrixadc-provider: Terraform version imported: %s\n", version.Version)

	return &c, nil
}
