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
	"log"
	"sync"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type NetScalerNitroClient struct {
	Username string
	Password string
	Endpoint string
	client   *service.NitroClient
	lock     sync.Mutex
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
		"citrixadc_lbmetrictable":                                   resourceCitrixAdcLbmetrictable(),
		"citrixadc_sslservice":                                      resourceCitrixAdcSslservice(),
		"citrixadc_sslfipskey":                                      resourceCitrixAdcSslfipskey(),
		"citrixadc_lbroute6":                                        resourceCitrixAdcLbroute6(),
		"citrixadc_sslpolicylabel":                                  resourceCitrixAdcSslpolicylabel(),
		"citrixadc_ssllogprofile":                                   resourceCitrixAdcSsllogprofile(),
		"citrixadc_sslcacertgroup":                                  resourceCitrixAdcSslcacertgroup(),
		"citrixadc_botpolicy":                                       resourceCitrixAdcBotpolicy(),
		"citrixadc_lbvserver":                                       resourceCitrixAdcLbvserver(),
		"citrixadc_service":                                         resourceCitrixAdcService(),
		"citrixadc_csvserver":                                       resourceCitrixAdcCsvserver(),
		"citrixadc_cspolicy":                                        resourceCitrixAdcCspolicy(),
		"citrixadc_csaction":                                        resourceCitrixAdcCsaction(),
		"citrixadc_sslaction":                                       resourceCitrixAdcSslaction(),
		"citrixadc_sslpolicy":                                       resourceCitrixAdcSslpolicy(),
		"citrixadc_sslparameter":                                    resourceCitrixAdcSslparameter(),
		"citrixadc_ssldhparam":                                      resourceCitrixAdcSsldhparam(),
		"citrixadc_sslcipher":                                       resourceCitrixAdcSslcipher(),
		"citrixadc_servicegroup":                                    resourceCitrixAdcServicegroup(),
		"citrixadc_nsacl":                                           resourceCitrixAdcNsacl(),
		"citrixadc_nsacls":                                          resourceCitrixAdcNsacls(),
		"citrixadc_inat":                                            resourceCitrixAdcInat(),
		"citrixadc_rnat":                                            resourceCitrixAdcRnat(),
		"citrixadc_gslbvserver":                                     resourceCitrixAdcGslbvserver(),
		"citrixadc_gslbservice":                                     resourceCitrixAdcGslbservice(),
		"citrixadc_server":                                          resourceCitrixAdcServer(),
		"citrixadc_responderaction":                                 resourceCitrixAdcResponderaction(),
		"citrixadc_responderpolicy":                                 resourceCitrixAdcResponderpolicy(),
		"citrixadc_responderpolicylabel":                            resourceCitrixAdcResponderpolicylabel(),
		"citrixadc_rewriteaction":                                   resourceCitrixAdcRewriteaction(),
		"citrixadc_rewritepolicy":                                   resourceCitrixAdcRewritepolicy(),
		"citrixadc_rewritepolicylabel":                              resourceCitrixAdcRewritepolicylabel(),
		"citrixadc_nsip":                                            resourceCitrixAdcNsip(),
		"citrixadc_nsip6":                                           resourceCitrixAdcNsip6(),
		"citrixadc_nsconfig_save":                                   resourceCitrixAdcNsconfigSave(),
		"citrixadc_nsconfig_clear":                                  resourceCitrixAdcNsconfigClear(),
		"citrixadc_nsconfig_update":                                 resourceCitrixAdcNsconfigUpdate(),
		"citrixadc_ipset":                                           resourceCitrixAdcIpset(),
		"citrixadc_route":                                           resourceCitrixAdcRoute(),
		"citrixadc_linkset":                                         resourceCitrixAdcLinkset(),
		"citrixadc_nsfeature":                                       resourceCitrixAdcNsfeature(),
		"citrixadc_systemgroup":                                     resourceCitrixAdcSystemgroup(),
		"citrixadc_systemcmdpolicy":                                 resourceCitrixAdcSystemcmdpolicy(),
		"citrixadc_interface":                                       resourceCitrixAdcInterface(),
		"citrixadc_nstcpprofile":                                    resourceCitrixAdcNstcpprofile(),
		"citrixadc_nshttpprofile":                                   resourceCitrixAdcNshttpprofile(),
		"citrixadc_nslicense":                                       resourceCitrixAdcNslicense(),
		"citrixadc_clusterfiles_syncer":                             resourceCitrixAdcClusterfilesSyncer(),
		"citrixadc_systemfile":                                      resourceCitrixAdcSystemfile(),
		"citrixadc_auditmessageaction":                              resourceCitrixAdcAuditmessageaction(),
		"citrixadc_auditsyslogpolicy":                               resourceCitrixAdcAuditsyslogpolicy(),
		"citrixadc_rebooter":                                        resourceCitrixAdcRebooter(),
		"citrixadc_installer":                                       resourceCitrixAdcInstaller(),
		"citrixadc_pinger":                                          resourceCitrixAdcPinger(),
		"citrixadc_routerdynamicrouting":                            resourceCitrixAdcRouterdynamicrouting(),
		"citrixadc_policyexpression":                                resourceCitrixAdcPolicyexpression(),
		"citrixadc_systemextramgmtcpu":                              resourceCitrixAdcSystemextramgmtcpu(),
		"citrixadc_netprofile":                                      resourceCitrixAdcNetprofile(),
		"citrixadc_nsparam":                                         resourceCitrixAdcNsparam(),
		"citrixadc_policydataset":                                   resourceCitrixAdcPolicydataset(),
		"citrixadc_password_resetter":                               resourceCitrixAdcPasswordResetter(),
		"citrixadc_appfwprofile":                                    resourceCitrixAdcAppfwprofile(),
		"citrixadc_appfwpolicy":                                     resourceCitrixAdcAppfwpolicy(),
		"citrixadc_appfwfieldtype":                                  resourceCitrixAdcAppfwfieldtype(),
		"citrixadc_appfwpolicylabel":                                resourceCitrixAdcAppfwpolicylabel(),
		"citrixadc_appfwjsoncontenttype":                            resourceCitrixAdcAppfwjsoncontenttype(),
		"citrixadc_appfwxmlcontenttype":                             resourceCitrixAdcAppfwxmlcontenttype(),
		"citrixadc_policystringmap":                                 resourceCitrixAdcPolicystringmap(),
		"citrixadc_transformprofile":                                resourceCitrixAdcTransformprofile(),
		"citrixadc_transformaction":                                 resourceCitrixAdcTransformaction(),
		"citrixadc_transformpolicy":                                 resourceCitrixAdcTransformpolicy(),
		"citrixadc_quicbridgeprofile":                               resourceCitrixAdcQuicbridgeprofile(),
		"citrixadc_policypatset":                                    resourceCitrixAdcPolicypatset(),
		"citrixadc_filterpolicy":                                    resourceCitrixAdcFilterpolicy(),
		"citrixadc_lbvserver_filterpolicy_binding":                  resourceCitrixAdcLbvserver_filterpolicy_binding(),
		"citrixadc_csvserver_filterpolicy_binding":                  resourceCitrixAdcCsvserver_filterpolicy_binding(),
		"citrixadc_cmppolicy":                                       resourceCitrixAdcCmppolicy(),
		"citrixadc_nsvpxparam":                                      resourceCitrixAdcNsvpxparam(),
		"citrixadc_nstcpparam":                                      resourceCitrixAdcNstcpparam(),
		"citrixadc_dnsnsrec":                                        resourceCitrixAdcDnsnsrec(),
		"citrixadc_dnssoarec":                                       resourceCitrixAdcDnssoarec(),
		"citrixadc_iptunnel":                                        resourceCitrixAdcIptunnel(),
		"citrixadc_vlan":                                            resourceCitrixAdcVlan(),
		"citrixadc_nsmode":                                          resourceCitrixAdcNsmode(),
		"citrixadc_botprofile":                                      resourceCitrixAdcBotprofile(),
		"citrixadc_botpolicylabel":                                  resourceCitrixAdcBotpolicylabel(),
		"citrixadc_lbsipparameters":                                 resourceCitrixAdcLbsipparameters(),
		"citrixadc_lbroute":                                         resourceCitrixAdcLbroute(),
		"citrixadc_lbgroup":                                         resourceCitrixAdcLbgroup(),
		"citrixadc_ssldtlsprofile":                                  resourceCitrixAdcSsldtlsprofile(),
		"citrixadc_sslocspresponder":                                resourceCitrixAdcSslocspresponder(),
		"citrixadc_csparameter":                                     resourceCitrixAdcCsparameter(),
		"citrixadc_cspolicylabel":                                   resourceCitrixAdcCspolicylabel(),
		"citrixadc_sslvserver":                                      resourceCitrixAdcSslvserver(),
		"citrixadc_policyhttpcallout":                               resourceCitrixAdcPolicyhttpcallout(),
		"citrixadc_policymap":                                       resourceCitrixAdcPolicymap(),
		"citrixadc_policyparam":                                     resourceCitrixAdcPolicyparam(),
		"citrixadc_sslservicegroup":                                 resourceCitrixAdcSslservicegroup(),
		"citrixadc_rewriteparam":                                    resourceCitrixAdcRewriteparam(),
		"citrixadc_responderparam":                                  resourceCitrixAdcResponderparam(),
		"citrixadc_vpneula":                                         resourceCitrixAdcVpneula(),
		"citrixadc_vpnclientlessaccesspolicy":                       resourceCitrixAdcVpnclientlessaccesspolicy(),
		"citrixadc_vpnalwaysonprofile":                              resourceCitrixAdcVpnalwaysonprofile(),
		"citrixadc_vpnformssoaction":                                resourceCitrixAdcVpnformssoaction(),
		"citrixadc_vpnclientlessaccessprofile":                      resourceCitrixAdcVpnclientlessaccessprofile(),
		"citrixadc_filterglobal_filterpolicy_binding":               resourceCitrixAdcFilterglobal_filterpolicy_binding(),
		"citrixadc_dnsparameter":                                    resourceCitrixAdcDnsparameter(),
		"citrixadc_responderhtmlpage":                               resourceCitrixAdcResponderhtmlpage(),
		"citrixadc_authorizationpolicy":                             resourceCitrixAdcAuthorizationpolicy(),
		"citrixadc_vpnurl":                                          resourceCitrixAdcVpnurl(),
		"citrixadc_vpnsessionaction":                                resourceCitrixAdcVpnsessionaction(),
		"citrixadc_vpnvserver":                                      resourceCitrixAdcVpnvserver(),
		"citrixadc_vpnsessionpolicy":                                resourceCitrixAdcVpnsessionpolicy(),
		"citrixadc_vpntrafficaction":                                resourceCitrixAdcVpntrafficaction(),
		"citrixadc_vpnurlaction":                                    resourceCitrixAdcVpnurlaction(),
		"citrixadc_vpntrafficpolicy":                                resourceCitrixAdcVpntrafficpolicy(),
		"citrixadc_vpnurlpolicy":                                    resourceCitrixAdcVpnurlpolicy(),
		"citrixadc_authenticationvserver":                           resourceCitrixAdcAuthenticationvserver(),
		"citrixadc_authenticationauthnprofile":                      resourceCitrixAdcAuthenticationauthnprofile(),
		"citrixadc_authenticationpolicylabel":                       resourceCitrixAdcAuthenticationpolicylabel(),
		"citrixadc_authenticationpolicy":                            resourceCitrixAdcAuthenticationpolicy(),
		"citrixadc_vpnintranetapplication":                          resourceCitrixAdcVpnintranetapplication(),
		"citrixadc_vpnpcoipvserverprofile":                          resourceCitrixAdcVpnpcoipvserverprofile(),
		"citrixadc_vpnpcoipprofile":                                 resourceCitrixAdcVpnpcoipprofile(),
		"citrixadc_vpnnexthopserver":                                resourceCitrixAdcVpnnexthopserver(),
		"citrixadc_vpnportaltheme":                                  resourceCitrixAdcVpnportaltheme(),
		"citrixadc_vpnsamlssoprofile":                               resourceCitrixAdcVpnsamlssoprofile(),
		"citrixadc_hafailover":                                      resourceCitrixAdcHafailover(),
		"citrixadc_authenticationldappolicy":                        resourceCitrixAdcAuthenticationldappolicy(),
		"citrixadc_authenticationlocalpolicy":                       resourceCitrixAdcAuthenticationlocalpolicy(),
		"citrixadc_authenticationnoauthaction":                      resourceCitrixAdcAuthenticationnoauthaction(),
		"citrixadc_authenticationsamlaction":                        resourceCitrixAdcAuthenticationsamlaction(),
		"citrixadc_authenticationepaaction":                         resourceCitrixAdcAuthenticationepaaction(),
		"citrixadc_authenticationloginschema":                       resourceCitrixAdcAuthenticationloginschema(),
		"citrixadc_authenticationsamlpolicy":                        resourceCitrixAdcAuthenticationsamlpolicy(),
		"citrixadc_authenticationdfapolicy":                         resourceCitrixAdcAuthenticationdfapolicy(),
		"citrixadc_authenticationwebauthaction":                     resourceCitrixAdcAuthenticationwebauthaction(),
		"citrixadc_authenticationradiuspolicy":                      resourceCitrixAdcAuthenticationradiuspolicy(),
		"citrixadc_authenticationstorefrontauthaction":              resourceCitrixAdcAuthenticationstorefrontauthaction(),
		"citrixadc_authenticationwebauthpolicy":                     resourceCitrixAdcAuthenticationwebauthpolicy(),
		"citrixadc_authenticationcitrixauthaction":                  resourceCitrixAdcAuthenticationcitrixauthaction(),
		"citrixadc_authenticationtacacspolicy":                      resourceCitrixAdcAuthenticationtacacspolicy(),
		"citrixadc_authenticationcertaction":                        resourceCitrixAdcAuthenticationcertaction(),
		"citrixadc_authenticationoauthidppolicy":                    resourceCitrixAdcAuthenticationoauthidppolicy(),
		"citrixadc_authenticationcertpolicy":                        resourceCitrixAdcAuthenticationcertpolicy(),
		"citrixadc_authenticationloginschemapolicy":                 resourceCitrixAdcAuthenticationloginschemapolicy(),
		"citrixadc_authenticationnegotiatepolicy":                   resourceCitrixAdcAuthenticationnegotiatepolicy(),
		"citrixadc_authenticationsamlidpprofile":                    resourceCitrixAdcAuthenticationsamlidpprofile(),
		"citrixadc_authenticationsamlidppolicy":                     resourceCitrixAdcAuthenticationsamlidppolicy(),
		"citrixadc_vpnparameter":                                    resourceCitrixAdcVpnparameter(),
		"citrixadc_vxlan":                                           resourceCitrixAdcVxlan(),
		"citrixadc_vxlanvlanmap":                                    resourceCitrixAdcVxlanvlanmap(),
		"citrixadc_appfwconfidfield":                                resourceCitrixAdcAppfwconfidfield(),
		"citrixadc_filteraction":                                    resourceCitrixAdcFilteraction(),
		"citrixadc_botsignature":                                    resourceCitrixAdcBotsignature(),
		"citrixadc_appfwurlencodedformcontenttype":                  resourceCitrixAdcAppfwurlencodedformcontenttype(),
		"citrixadc_nitro_resource":                                  resourceCitrixAdcNintroResource(),
		"citrixadc_appfwsignatures":                                 resourceCitrixAdcAppfwsignatures(),
		"citrixadc_appfwlearningsettings":                           resourceCitrixAdcAppfwlearningsettings(),
		"citrixadc_appfwhtmlerrorpage":                              appfwhtmlerrorpage(),
		"citrixadc_appfwxmlerrorpage":                               resourceCitrixAdcAppfwxmlerrorpage(),
		"citrixadc_location":                                        resourceCitrixAdcLocation(),
		"citrixadc_service_dospolicy_binding":                       resourceCitrixAdcService_dospolicy_binding(),
		"citrixadc_fis":                                             resourceCitrixAdcFis(),
		"citrixadc_vrid":                                            resourceCitrixAdcVrid(),
		"citrixadc_vrid6":                                           resourceCitrixAdcVrid6(),
		"citrixadc_forwardingsession":                               resourceCitrixAdcForwardingsession(),
		"citrixadc_netbridge":                                       resourceCitrixAdcNetbridge(),
		"citrixadc_appfwjsonerrorpage":                              resourceCitrixAdcAppfwjsonerrorpage(),
		"citrixadc_nstimer":                                         resourceCitrixAdcNstimer(),
		"citrixadc_nslimitidentifier":                               resourceCitrixAdcNslimitidentifier(),
		"citrixadc_nsservicepath":                                   resourceCitrixAdcNsservicepath(),
		"citrixadc_nspartition":                                     resourceCitrixAdcNspartition(),
		"citrixadc_nsvariable":                                      resourceCitrixAdcNsvariable(),
		"citrixadc_nsappflowcollector":                              resourceCitrixAdcNsappflowcollector(),
		"citrixadc_nsicapprofile":                                   resourceCitrixAdcNsicapprofile(),
		"citrixadc_nsxmlnamespace":                                  resourceCitrixAdcNsxmlnamespace(),
		"citrixadc_nstrafficdomain":                                 resourceCitrixAdcNstrafficdomain(),
		"citrixadc_nsservicefunction":                               resourceCitrixAdcNsservicefunction(),
		"citrixadc_nssimpleacl":                                     resourceCitrixAdcNssimpleacl(),
		"citrixadc_mapbmr":                                          resourceCitrixAdcMapbmr(),
		"citrixadc_nssimpleacl6":                                    resourceCitrixAdcNssimpleacl6(),
		"citrixadc_mapdmr":                                          resourceCitrixAdcMapdmr(),
		"citrixadc_appfwwsdl":                                       resourceCitrixAdcAppfwwsdl(),
		"citrixadc_nsspparams":                                      resourceCitrixAdcNsspparams(),
		"citrixadc_nsconsoleloginprompt":                            resourceCitrixAdcNsconsoleloginprompt(),
		"citrixadc_extendedmemoryparam":                             resourceCitrixAdcExtendedmemoryparam(),
		"citrixadc_appfwmultipartformcontenttype":                   resourceCitrixAdcAppfwmultipartformcontenttype(),
		"citrixadc_locationparameter":                               resourceCitrixAdcLocationparameter(),
		"citrixadc_nsdiameter":                                      resourceCitrixAdcNsdiameter(),
		"citrixadc_nsdhcpparams":                                    resourceCitrixAdcNsdhcpparams(),
		"citrixadc_nsassignment":                                    resourceCitrixAdcNsassignment(),
		"citrixadc_appfwxmlschema":                                  resourceCitrixAdcAppfwxmlschema(),
		"citrixadc_nsratecontrol":                                   resourceCitrixAdcNsratecontrol(),
		"citrixadc_l4param":                                         resourceCitrixAdcL4param(),
		"citrixadc_arpparam":                                        resourceCitrixAdcArpparam(),
		"citrixadc_rnatparam":                                       resourceCitrixAdcRnatparam(),
		"citrixadc_ptp":                                             resourceCitrixAdcPtp(),
		"citrixadc_nshttpparam":                                     resourceCitrixAdcNshttpparam(),
		"citrixadc_mapdomain":                                       resourceCitrixAdcMapdomain(),
		"citrixadc_vridparam":                                       resourceCitrixAdcVridparam(),
		"citrixadc_bridgegroup":                                     resourceCitrixAdcBridgegroup(),
		"citrixadc_nat64param":                                      resourceCitrixAdcNat64param(),
		"citrixadc_nslicenseparameters":                             resourceCitrixAdcNslicenseparameters(),
		"citrixadc_appalgparam":                                     resourceCitrixAdcAppalgparam(),
		"citrixadc_iptunnelparam":                                   resourceCitrixAdcIptunnelparam(),
		"citrixadc_nsacl6":                                          resourceCitrixAdcNsacl6(),
		"citrixadc_nspbr6":                                          resourceCitrixAdcNspbr6(),
		"citrixadc_nstcpbufparam":                                   resourceCitrixAdcNstcpbufparam(),
		"citrixadc_ip6tunnel":                                       resourceCitrixAdcIp6tunnel(),
		"citrixadc_rsskeytype":                                      resourceCitrixAdcRsskeytype(),
		"citrixadc_nat64":                                           resourceCitrixAdcNat64(),
		"citrixadc_dnsaddrec":                                       resourceCitrixAdcDnsaddrec(),
		"citrixadc_dnspolicy":                                       resourceCitrixAdcDnspolicy(),
		"citrixadc_dnsaction":                                       resourceCitrixAdcDnsaction(),
		"citrixadc_dnsprofile":                                      resourceCitrixAdcDnsprofile(),
		"citrixadc_bridgetable":                                     resourceCitrixAdcBridgetable(),
		"citrixadc_dnsview":                                         resourceCitrixAdcDnsview(),
		"citrixadc_dnsmxrec":                                        resourceCitrixAdcDnsmxrec(),
		"citrixadc_dnspolicy64":                                     resourceCitrixAdcDnspolicy64(),
		"citrixadc_dnsaction64":                                     resourceCitrixAdcDnsaction64(),
		"citrixadc_dnssrvrec":                                       resourceCitrixAdcDnssrvrec(),
		"citrixadc_dnsnaptrrec":                                     resourceCitrixAdcDnsnaptrrec(),
		"citrixadc_dnscnamerec":                                     resourceCitrixAdcDnscnamerec(),
		"citrixadc_dnsaaaarec":                                      resourceCitrixAdcDnsaaaarec(),
		"citrixadc_dnsnameserver":                                   resourceCitrixAdcDnsnameserver(),
		"citrixadc_dnssuffix":                                       resourceCitrixAdcDnssuffix(),
		"citrixadc_dnspolicylabel":                                  resourceCitrixAdcDnspolicylabel(),
		"citrixadc_dnsptrrec":                                       resourceCitrixAdcDnsptrrec(),
		"citrixadc_ip6tunnelparam":                                  resourceCitrixAdcIp6tunnelparam(),
		"citrixadc_dnszone":                                         resourceCitrixAdcDnszone(),
		"citrixadc_crpolicy":                                        resourceCitrixAdcCrpolicy(),
		"citrixadc_gslbparameter":                                   resourceCitrixAdcGslbparameter(),
		"citrixadc_transformpolicylabel":                            resourceCitrixAdcTransformpolicylabel(),
		"citrixadc_gslbservicegroup":                                resourceCitrixAdcGslbservicegroup(),
		"citrixadc_authorizationpolicylabel":                        resourceCitrixAdcAuthorizationpolicylabel(),
		"citrixadc_gslbservicegroup_lbmonitor_binding":              resourceCitrixAdcGslbservicegroup_lbmonitor_binding(),
		"citrixadc_crvserver_filterpolicy_binding":                  resourceCitrixAdcCrvserver_filterpolicy_binding(),
		"citrixadc_dnstxtrec":                                       resourceCitrixAdcDnstxtrec(),
		"citrixadc_crvserver":                                       resourceCitrixAdcCrvserver(),
		"citrixadc_appflowcollector":                                resourceCitrixAdcAppflowcollector(),
		"citrixadc_appflowpolicylabel":                              resourceCitrixAdcAppflowpolicylabel(),
		"citrixadc_appflowaction":                                   resourceCitrixAdcAppflowaction(),
		"citrixadc_appflowpolicy":                                   resourceCitrixAdcAppflowpolicy(),
		"citrixadc_snmptrap":                                        resourceCitrixAdcSnmptrap(),
		"citrixadc_snmpview":                                        resourceCitrixAdcSnmpview(),
		"citrixadc_snmpgroup":                                       resourceCitrixAdcSnmpgroup(),
		"citrixadc_snmpengineid":                                    resourceCitrixAdcSnmpengineid(),
		"citrixadc_opoption":                                        resourceCitrixAdcSnmpoption(),
		"citrixadc_snmpmib":                                         resourceCitrixAdcSnmpmib(),
		"citrixadc_snmpmanager":                                     resourceCitrixAdcSnmpmanager(),
		"citrixadc_snmpalarm":                                       resourceCitrixAdcSnmpalarm(),
		"citrixadc_ntpserver":                                       resourceCitrixAdcNtpserver(),
		"citrixadc_systemparameter":                                 resourceCitrixAdcSystemparameter(),
		"citrixadc_nstimeout":                                       resourceCitrixAdcNstimeout(),
		"citrixadc_clusternodegroup":                                resourceCitrixAdcClusternodegroup(),
		"citrixadc_clusternodegroup_gslbsite_binding":               resourceCitrixAdcClusternodegroup_gslbsite_binding(),
		"citrixadc_clusternodegroup_lbvserver_binding":              resourceCitrixAdcClusternodegroup_lbvserver_binding(),
		"citrixadc_clusternodegroup_clusternode_binding":            resourceCitrixAdcClusternodegroup_clusternode_binding(),
		"citrixadc_clusternodegroup_csvserver_binding":              resourceCitrixAdcClusternodegroup_csvserver_binding(),
		"citrixadc_clusternodegroup_crvserver_binding":              resourceCitrixAdcClusternodegroup_crvserver_binding(),
		"citrixadc_clusternodegroup_gslbvserver_binding":            resourceCitrixAdcClusternodegroup_gslbvserver_binding(),
		"citrixadc_clusterinstance":                                 resourceCitrixAdcClusterinstance(),
		"citrixadc_clusternodegroup_service_binding":                resourceCitrixAdcClusternodegroup_service_binding(),
		"citrixadc_clusternode":                                     resourceCitrixAdcClusternode(),
		"citrixadc_clusternodegroup_nslimitidentifier_binding":      resourceCitrixAdcClusternodegroup_nslimitidentifier_binding(),
		"citrixadc_clusternodegroup_vpnvserver_binding":             resourceCitrixAdcClusternodegroup_vpnvserver_binding(),
		"citrixadc_clusternodegroup_streamidentifier_binding":       resourceCitrixAdcClusternodegroup_streamidentifier_binding(),
		"citrixadc_clusternodegroup_authenticationvserver_binding":  resourceCitrixAdcClusternodegroup_authenticationvserver_binding(),
		"citrixadc_nscqaparam":                                      resourceCitrixAdcNscqaparam(),
		"citrixadc_nshostname":                                      resourceCitrixAdcNshostname(),
		"citrixadc_nslicenseproxyserver":                            resourceCitrixAdcNslicenseproxyserver(),
		"citrixadc_snmpcommunity":                                   resourceCitrixAdcSnmpcommunity(),
		"citrixadc_lacp":                                            resourceCitrixAdcLacp(),
		"citrixadc_route6":                                          resourceCitrixAdcRoute6(),
		"citrixadc_nd6":                                             resourceCitrixAdcNd6(),
		"citrixadc_nspbr":                                           resourceCitrixAdcNspbr(),
		"citrixadc_l3param":                                         resourceCitrixAdcL3param(),
		"citrixadc_arp":                                             resourceCitrixAdcArp(),
		"citrixadc_nd6ravariables":                                  resourceCitrixAdcNd6ravariables(),
		"citrixadc_l2param":                                         resourceCitrixAdcL2param(),
		"citrixadc_rnat6":                                           resourceCitrixAdcRnat6(),
		"citrixadc_inatparam":                                       resourceCitrixAdcInatparam(),
		"citrixadc_ipv6":                                            resourceCitrixAdcIpv6(),
		"citrixadc_onlinkipv6prefix":                                resourceCitrixAdcOnlinkipv6prefix(),
		"citrixadc_aaacertparams":                                   resourceCitrixAdcAaacertparams(),
		"citrixadc_aaaotpparameter":                                 resourceCitrixAdcAaaotpparameter(),
		"citrixadc_aaaparameter":                                    resourceCitrixAdcAaaparameter(),
		"citrixadc_aaapreauthenticationpolicy":                      resourceCitrixAdcAaapreauthenticationpolicy(),
		"citrixadc_aaapreauthenticationparameter":                   resourceCitrixAdcAaapreauthenticationparameter(),
		"citrixadc_aaagroup":                                        resourceCitrixAdcAaagroup(),
		"citrixadc_systemcollectionparam":                           resourceCitrixAdcSystemcollectionparam(),
		"citrixadc_systembackup":                                    resourceCitrixAdcSystembackup(),
		"citrixadc_systembackup_restore":                            resourceCitrixAdcSystembackupRestore(),
		"citrixadc_cmppolicylabel":                                  resourceCitrixAdcCmppolicylabel(),
		"citrixadc_cmpaction":                                       resourceCitrixAdcCmpaction(),
		"citrixadc_cmpparameter":                                    resourceCitrixAdcCmpparameter(),
		"citrixadc_icaaction":                                       resourceCitrixAdcIcaaction(),
		"citrixadc_icaaccessprofile":                                resourceCitrixAdcIcaaccessprofile(),
		"citrixadc_icalatencyprofile":                               resourceCitrixAdcIcalatencyprofile(),
		"citrixadc_ntpparam":                                        resourceCitrixAdcNtpparam(),
		"citrixadc_icapolicy":                                       resourceCitrixAdcIcapolicy(),
		"citrixadc_tmsessionparameter":                              resourceCitrixAdcTmsessionparameter(),
		"citrixadc_tmformssoaction":                                 resourceCitrixAdcTmformssoaction(),
		"citrixadc_tmsessionpolicy":                                 resourceCitrixAdcTmsessionpolicy(),
		"citrixadc_tmsessionaction":                                 resourceCitrixAdcTmsessionaction(),
		"citrixadc_tmtrafficpolicy":                                 resourceCitrixAdcTmtrafficpolicy(),
		"citrixadc_ipsecparameter":                                  resourceCitrixAdcIpsecparameter(),
		"citrixadc_tmsamlssoprofile":                                resourceCitrixAdcTmsamlssoprofile(),
		"citrixadc_interfacepair":                                   resourceCitrixAdcInterfacepair(),
		"citrixadc_channel":                                         resourceCitrixAdcChannel(),
		"citrixadc_userprotocol":                                    resourceCitrixAdcUserprotocol(),
		"citrixadc_uservserver":                                     resourceCitrixAdcUservserver(),
		"citrixadc_auditsyslogparams":                               resourceCitrixAdcAuditsyslogparams(),
		"citrixadc_auditnslogparams":                                resourceCitrixAdcAuditnslogparams(),
		"citrixadc_lldpparam":                                       resourceCitrixAdcLldpparam(),
		"citrixadc_feopolicy":                                       resourceCitrixAdcFeopolicy(),
		"citrixadc_ipsecalgprofile":                                 resourceCitrixAdcIpsecalgprofile(),
		"citrixadc_dbdbprofile":                                     resourceCitrixAdcDbdbprofile(),
		"citrixadc_feoaction":                                       resourceCitrixAdcFeoaction(),
		"citrixadc_hanode":                                          resourceCitrixAdcHanode(),
		"citrixadc_admparameter":                                    resourceCitrixAdcAdmparameter(),
		"citrixadc_tmtrafficaction":                                 resourceCitrixAdcTmtrafficaction(),
		"citrixadc_systembackup_create":                             resourceCitrixAdcSystemCreatebackup(),
		"citrixadc_locationfile":                                    resourceCitrixAdcLocationfile(),
		"citrixadc_locationfile_import":                             resourceCitrixAdcLocationImportfile(),
		"citrixadc_feoparameter":                                    resourceCitrixAdcFeoparameter(),
		"citrixadc_appqoepolicy":                                    resourceCitrixAdcAppqoepolicy(),
		"citrixadc_appqoecustomresp":                                resourceCitrixAdcAppqoecustomresp(),
		"citrixadc_appqoeparameter":                                 resourceCitrixAdcAppqoeparameter(),
		"citrixadc_appqoeaction":                                    resourceCitrixAdcAppqoeaction(),
		"citrixadc_contentinspectionpolicy":                         resourceCitrixAdcContentinspectionpolicy(),
		"citrixadc_contentinspectionparameter":                      resourceCitrixAdcContentinspectionparameter(),
		"citrixadc_contentinspectioncallout":                        resourceCitrixAdcContentinspectioncallout(),
		"citrixadc_contentinspectionpolicylabel":                    resourceCitrixAdcContentinspectionpolicylabel(),
		"citrixadc_contentinspectionprofile":                        resourceCitrixAdcContentinspectionprofile(),
		"citrixadc_lsnclient":                                       resourceCitrixAdcLsnclient(),
		"citrixadc_lsnappsattributes":                               resourceCitrixAdcLsnappsattributes(),
		"citrixadc_lsngroup":                                        resourceCitrixAdcLsngroup(),
		"citrixadc_lsnappsprofile":                                  resourceCitrixAdcLsnappsprofile(),
		"citrixadc_lsnlogprofile":                                   resourceCitrixAdcLsnlogprofile(),
		"citrixadc_lsnpool":                                         resourceCitrixAdcLsnpool(),
		"citrixadc_lsnrtspalgprofile":                               resourceCitrixAdcLsnrtspalgprofile(),
		"citrixadc_lsnsipalgprofile":                                resourceCitrixAdcLsnsipalgprofile(),
		"citrixadc_ntpsync":                                         resourceCitrixAdcNtpsync(),
		"citrixadc_cacheparameter":                                  resourceCitrixAdcCacheparameter(),
		"citrixadc_lsntransportprofile":                             resourceCitrixAdcLsntransportprofile(),
		"citrixadc_lsnip6profile":                                   resourceCitrixAdcLsnip6profile(),
		"citrixadc_lsnstatic":                                       resourceCitrixAdcLsnstatic(),
		"citrixadc_streamidentifier":                                resourceCitrixAdcStreamidentifier(),
		"citrixadc_contentinspectionaction":                         resourceCitrixAdcContentinspectionaction(),
		"citrixadc_pcpserver":                                       resourceCitrixAdcPcpserver(),
		"citrixadc_pcpprofile":                                      resourceCitrixAdcPcpprofile(),
		"citrixadc_nsweblogparam":                                   resourceCitrixAdcNsweblogparam(),
		"citrixadc_autoscaleaction":                                 resourceCitrixAdcAutoscaleaction(),
		"citrixadc_autoscalepolicy":                                 resourceCitrixAdcAutoscalepolicy(),
		"citrixadc_spilloveraction":                                 resourceCitrixAdcSpilloveraction(),
		"citrixadc_locationfile6":                                   resourceCitrixAdcLocationfile6(),
		"citrixadc_locationfile6_import":                            resourceCitrixAdcLocationfile6Import(),
		"citrixadc_subscribergxinterface":                           resourceCitrixAdcSubscribergxinterface(),
		"citrixadc_cachecontentgroup":                               resourceCitrixAdcCachecontentgroup(),
		"citrixadc_cacheforwardproxy":                               resourceCitrixAdcCacheforwardproxy(),
		"citrixadc_auditnslogpolicy":                                resourceCitrixAdcAuditnslogpolicy(),
		"citrixadc_auditnslogaction":                                resourceCitrixAdcAuditnslogaction(),
		"citrixadc_icaparameter":                                    resourceCitrixAdcIcaparameter(),
		"citrixadc_smppparam":                                       resourceCitrixAdcSmppparam(),
		"citrixadc_cachepolicy":                                     resourceCitrixAdcCachepolicy(),
		"citrixadc_cachepolicylabel":                                resourceCitrixAdcCachepolicylabel(),
		"citrixadc_cacheselector":                                   resourceCitrixAdcCacheselector(),
		"citrixadc_tunneltrafficpolicy":                             resourceCitrixAdcTunneltrafficpolicy(),
		"citrixadc_subscriberprofile":                               resourceCitrixAdcSubscriberprofile(),
		"citrixadc_subscriberradiusinterface":                       resourceCitrixAdcSubscriberradiusinterface(),
		"citrixadc_subscriberparam":                                 resourceCitrixAdcSubscriberparam(),
		"citrixadc_lsnhttphdrlogprofile":                            resourceCitrixAdcLsnhttphdrlogprofile(),
		"citrixadc_streamselector":                                  resourceCitrixAdcStreamselector(),
		"citrixadc_lsnparameter":                                    resourceCitrixAdcLsnparameter(),
		"citrixadc_sslcertfile":                                     resourceCitrixAdcSslcertfile(),
		"citrixadc_nspbrs":                                          resourceCitrixAdcNspbrs(),
		"citrixadc_rnat_clear":                                      resourceCitrixAdcRnatClear(),
		"citrixadc_change_password":                                 resourceCitrixAdcChangePassword(),
		"citrixadc_spilloverpolicy":                                 resourceCitrixAdcSpilloverpolicy(),
		"citrixadc_videooptimizationdetectionaction":                resourceCitrixAdcVideooptimizationdetectionaction(),
		"citrixadc_videooptimizationdetectionpolicy":                resourceCitrixAdcVideooptimizationdetectionpolicy(),
		"citrixadc_aaapreauthenticationaction":                      resourceCitrixAdcAaapreauthenticationaction(),
		"citrixadc_videooptimizationpacingaction":                   resourceCitrixAdcVideooptimizationpacingaction(),
		"citrixadc_videooptimizationpacingpolicy":                   resourceCitrixAdcVideooptimizationpacingpolicy(),
		"citrixadc_lbaction":                                        resourceCitrixAdcLbaction(),
		"citrixadc_lbpolicy":                                        resourceCitrixAdcLbpolicy(),
		"citrixadc_gslbservicegroup_gslbservicegroupmember_binding": resourceCitrixAdcGslbservicegroup_gslbservicegroupmember_binding(),
		"citricadc_nscapacity":                                      resourceCitrixAdcNscapacity(),
		"citrixadc_nslicenseserver":                                 resourceCitrixAdcNslicenseserver(),
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
		Username: username,
		Password: password,
		Endpoint: endpoint,
	}

	params := service.NitroParams{
		Url:       endpoint,
		Username:  username,
		Password:  password,
		ProxiedNs: d.Get("proxied_ns").(string),
		SslVerify: !d.Get("insecure_skip_verify").(bool),
		Headers:   userHeaders,
		IsCloud:   d.Get("is_cloud").(bool),
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
