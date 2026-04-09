/*
Copyright 2025 Citrix Systems, Inc

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

package provider

import (
	"context"
	"os"

	adcnitrogoservice "github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaacertparams"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaaglobal_aaapreauthenticationpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaagroup"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaagroup_aaauser_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaagroup_auditnslogpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaagroup_auditsyslogpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaagroup_authorizationpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaagroup_intranetip_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaagroup_tmsessionpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaagroup_vpnintranetapplication_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaagroup_vpnsessionpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaagroup_vpntrafficpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaagroup_vpnurl_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaagroup_vpnurlpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaakcdaccount"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaaldapparams"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaaotpparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaaparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaapreauthenticationaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaapreauthenticationparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaapreauthenticationpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaaradiusparams"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaassoprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaatacacsparams"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaauser"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaauser_auditnslogpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaauser_auditsyslogpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaauser_authorizationpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaauser_intranetip6_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaauser_intranetip_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaauser_tmsessionpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaauser_vpnintranetapplication_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaauser_vpnsessionpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaauser_vpntrafficpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaauser_vpnurl_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/aaauser_vpnurlpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/analyticsglobal_analyticsprofile_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/analyticsprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appalgparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appflowaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appflowaction_analyticsprofile_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appflowcollector"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appflowglobal_appflowpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appflowparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appflowpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appflowpolicylabel"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appflowpolicylabel_appflowpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwconfidfield"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwfieldtype"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwglobal_appfwpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwglobal_auditnslogpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwglobal_auditsyslogpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwhtmlerrorpage"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwjsoncontenttype"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwjsonerrorpage"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwlearningsettings"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwmultipartformcontenttype"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwpolicylabel"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwpolicylabel_appfwpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwprofile_cmdinjection_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwprofile_contenttype_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwprofile_cookieconsistency_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwprofile_creditcardnumber_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwprofile_crosssitescripting_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwprofile_csrftag_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwprofile_denyurl_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwprofile_excluderescontenttype_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwprofile_fieldconsistency_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwprofile_fieldformat_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwprofile_fileuploadtype_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwprofile_jsoncmdurl_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwprofile_jsondosurl_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwprofile_jsonsqlurl_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwprofile_jsonxssurl_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwprofile_logexpression_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwprofile_safeobject_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwprofile_sqlinjection_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwprofile_starturl_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwprofile_trustedlearningclients_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwprofile_xmlattachmenturl_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwprofile_xmldosurl_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwprofile_xmlsqlinjection_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwprofile_xmlvalidationurl_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwprofile_xmlwsiurl_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwprofile_xmlxss_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwsettings"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwsignatures"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwurlencodedformcontenttype"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwwsdl"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwxmlcontenttype"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwxmlerrorpage"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwxmlschema"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appqoeaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appqoecustomresp"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appqoeparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appqoepolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/arp"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/arpparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/auditmessageaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/auditnslogaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/auditnslogglobal_auditnslogpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/auditnslogparams"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/auditnslogpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/auditsyslogaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/auditsyslogglobal_auditsyslogpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/auditsyslogparams"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/auditsyslogpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationauthnprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationcaptchaaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationcertaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationcertpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationcitrixauthaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationdfaaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationdfapolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationemailaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationepaaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationldapaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationldappolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationlocalpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationloginschema"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationloginschemapolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationnegotiateaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationnegotiatepolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationnoauthaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationoauthaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationoauthidppolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationoauthidpprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationpolicylabel"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationpolicylabel_authenticationpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationpushservice"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationradiusaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationradiuspolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationsamlaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationsamlidppolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationsamlidpprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationsamlpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationstorefrontauthaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationtacacsaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationtacacspolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationvserver"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationvserver_auditnslogpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationvserver_auditsyslogpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationvserver_authenticationcertpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationvserver_authenticationldappolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationvserver_authenticationlocalpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationvserver_authenticationloginschemapolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationvserver_authenticationnegotiatepolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationvserver_authenticationoauthidppolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationvserver_authenticationpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationvserver_authenticationradiuspolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationvserver_authenticationsamlidppolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationvserver_authenticationsamlpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationvserver_authenticationtacacspolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationvserver_authenticationwebauthpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationvserver_cachepolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationvserver_cspolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationvserver_responderpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationvserver_rewritepolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationvserver_tmsessionpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationvserver_vpnportaltheme_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationwebauthaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationwebauthpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authorizationpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authorizationpolicylabel"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authorizationpolicylabel_authorizationpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/autoscaleaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/autoscalepolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/autoscaleprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/botglobal_botpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/botpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/botpolicylabel"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/botpolicylabel_botpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/botprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/botprofile_blacklist_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/botprofile_captcha_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/botprofile_ipreputation_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/botprofile_logexpression_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/botprofile_ratelimit_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/botprofile_tps_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/botprofile_trapinsertionurl_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/botprofile_whitelist_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/botsettings"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/botsignature"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/bridgegroup"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/bridgegroup_nsip6_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/bridgegroup_nsip_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/bridgegroup_vlan_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/bridgetable"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/cachecontentgroup"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/cacheforwardproxy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/cacheglobal_cachepolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/cacheparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/cachepolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/cachepolicylabel"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/cachepolicylabel_cachepolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/cacheselector"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/channel"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/clusterinstance"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/clusternode"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/clusternode_routemonitor_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/clusternodegroup"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/clusternodegroup_authenticationvserver_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/clusternodegroup_clusternode_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/clusternodegroup_crvserver_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/clusternodegroup_csvserver_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/clusternodegroup_gslbsite_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/clusternodegroup_gslbvserver_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/clusternodegroup_lbvserver_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/clusternodegroup_nslimitidentifier_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/clusternodegroup_service_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/clusternodegroup_streamidentifier_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/clusternodegroup_vpnvserver_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/cmpaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/cmpglobal_cmppolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/cmpparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/cmppolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/cmppolicylabel"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/cmppolicylabel_cmppolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/contentinspectionaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/contentinspectioncallout"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/contentinspectionglobal_contentinspectionpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/contentinspectionparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/contentinspectionpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/contentinspectionpolicylabel"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/contentinspectionpolicylabel_contentinspectionpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/contentinspectionprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/crpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/crvserver"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/crvserver_analyticsprofile_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/crvserver_appflowpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/crvserver_appfwpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/crvserver_appqoepolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/crvserver_cachepolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/crvserver_cmppolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/crvserver_crpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/crvserver_cspolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/crvserver_feopolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/crvserver_icapolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/crvserver_lbvserver_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/crvserver_policymap_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/crvserver_responderpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/crvserver_rewritepolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/crvserver_spilloverpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/csaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/csparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/cspolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/cspolicylabel"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/csvserver"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/csvserver_analyticsprofile_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/csvserver_appflowpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/csvserver_appfwpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/csvserver_appqoepolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/csvserver_auditnslogpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/csvserver_auditsyslogpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/csvserver_authorizationpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/csvserver_botpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/csvserver_cachepolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/csvserver_cmppolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/csvserver_contentinspectionpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/csvserver_cspolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/csvserver_feopolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/csvserver_gslbvserver_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/csvserver_lbvserver_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/csvserver_responderpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/csvserver_rewritepolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/csvserver_spilloverpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/csvserver_tmtrafficpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/csvserver_transformpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/csvserver_vpnvserver_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dbdbprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dbuser"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnsaaaarec"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnsaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnsaction64"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnsaddrec"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnscnamerec"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnsglobal_dnspolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnskey"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnsmxrec"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnsnameserver"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnsnaptrrec"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnsnsrec"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnsparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnspolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnspolicy64"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnspolicylabel"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnspolicylabel_dnspolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnsprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnsptrrec"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnssoarec"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnssrvrec"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnssuffix"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnstxtrec"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnsview"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnszone"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/extendedmemoryparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/feoaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/feoglobal_feopolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/feoparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/feopolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/fis"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/forwardingsession"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/gslbparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/gslbservice"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/gslbservice_dnsview_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/gslbservice_lbmonitor_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/gslbservicegroup"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/gslbservicegroup_gslbservicegroupmember_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/gslbservicegroup_lbmonitor_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/gslbsite"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/gslbvserver"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/gslbvserver_domain_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/gslbvserver_gslbservice_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/gslbvserver_gslbservicegroup_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/gslbvserver_lbpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/gslbvserver_spilloverpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/hanode"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/hanode_routemonitor6_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/hanode_routemonitor_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/icaaccessprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/icaaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/icaglobal_icapolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/icalatencyprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/icaparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/icapolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/inat"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/inatparam"
	Interface "github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/interface"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/interfacepair"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/ip6tunnel"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/ip6tunnelparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/ipsecalgprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/ipsecparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/ipsecprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/ipset"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/ipset_nsip6_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/ipset_nsip_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/iptunnel"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/iptunnelparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/ipv6"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/l2param"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/l3param"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/l4param"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lacp"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbgroup"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbgroup_lbvserver_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbmetrictable"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbmetrictable_metric_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbmonitor"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbmonitor_metric_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbmonitor_sslcertkey_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbroute"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbroute6"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbsipparameters"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbvserver"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbvserver_analyticsprofile_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbvserver_appflowpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbvserver_appfwpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbvserver_appqoepolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbvserver_auditsyslogpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbvserver_authorizationpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbvserver_botpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbvserver_cachepolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbvserver_cmppolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbvserver_contentinspectionpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbvserver_dnspolicy64_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbvserver_feopolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbvserver_lbpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbvserver_responderpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbvserver_rewritepolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbvserver_service_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbvserver_servicegroup_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbvserver_spilloverpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbvserver_tmtrafficpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbvserver_transformpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbvserver_videooptimizationdetectionpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbvserver_videooptimizationpacingpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/linkset"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/linkset_channel_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lldpparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/location"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/locationfile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/locationfile6"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/locationparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lsnappsattributes"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lsnappsprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lsnappsprofile_lsnappsattributes_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lsnappsprofile_port_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lsnclient"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lsnclient_network6_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lsnclient_network_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lsnclient_nsacl6_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lsnclient_nsacl_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lsngroup"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lsngroup_lsnappsprofile_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lsngroup_lsnhttphdrlogprofile_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lsngroup_lsnlogprofile_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lsngroup_lsnpool_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lsngroup_lsntransportprofile_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lsngroup_pcpserver_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lsnhttphdrlogprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lsnip6profile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lsnlogprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lsnparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lsnpool"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lsnrtspalgprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lsnsipalgprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lsnstatic"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lsntransportprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/mapbmr"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/mapbmr_bmrv4network_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/mapdmr"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/mapdomain"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/mapdomain_mapbmr_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nat64"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nat64param"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nd6"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nd6ravariables"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nd6ravariables_onlinkipv6prefix_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/netbridge"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/netbridge_iptunnel_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/netbridge_nsip6_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/netbridge_nsip_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/netbridge_vlan_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/netprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/netprofile_natrule_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/netprofile_srcportset_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsacl"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsacl6"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsacls"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsappflowcollector"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsassignment"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nscapacity"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsconsoleloginprompt"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nscqaparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsdhcpparams"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsdiameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsencryptionkey"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsencryptionparams"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsfeature"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nshmackey"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nshostname"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nshttpparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nshttpprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsicapprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsip"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsip6"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nslaslicense_offline"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nslicense"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nslicenseparameters"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nslicenseproxyserver"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nslicenseserver"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nslimitidentifier"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsmode"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nspartition"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nspartition_bridgegroup_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nspartition_vlan_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nspartition_vxlan_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nspbr"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nspbr6"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsratecontrol"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsrpcnode"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsservicefunction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsservicepath"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsservicepath_nsservicefunction_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nssimpleacl"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nssimpleacl6"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsspparams"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nstcpbufparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nstcpparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nstcpprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nstimeout"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nstimer"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nstrafficdomain"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nstrafficdomain_bridgegroup_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nstrafficdomain_vlan_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nstrafficdomain_vxlan_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsvariable"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsvpxparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsweblogparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsxmlnamespace"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/ntpparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/ntpserver"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/onlinkipv6prefix"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/pcpprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/pcpserver"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/policydataset"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/policydataset_value_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/policyexpression"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/policyhttpcallout"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/policymap"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/policyparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/policypatset"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/policypatset_pattern_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/policystringmap"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/policystringmap_pattern_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/ptp"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/quicbridgeprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/radiusnode"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/rdpclientprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/rdpserverprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/reputationsettings"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/responderaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/responderglobal_responderpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/responderhtmlpage"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/responderparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/responderpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/responderpolicylabel"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/responderpolicylabel_responderpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/rewriteaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/rewriteglobal_rewritepolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/rewriteparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/rewritepolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/rewritepolicylabel"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/rewritepolicylabel_rewritepolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/rnat"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/rnat6"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/rnat6_nsip6_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/rnat_nsip_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/rnatglobal_auditsyslogpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/rnatparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/route"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/route6"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/routerdynamicrouting"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/rsskeytype"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/server"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/service_lbmonitor_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/servicegroup"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/servicegroup_lbmonitor_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/servicegroup_servicegroupmember_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/smppparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/smppuser"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/snmpalarm"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/snmpcommunity"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/snmpengineid"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/snmpgroup"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/snmpmanager"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/snmpmib"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/snmpoption"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/snmptrap"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/snmptrap_snmpuser_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/snmpuser"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/snmpview"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/spilloveraction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/spilloverpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslcacertgroup"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslcacertgroup_sslcertkey_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslcertfile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslcertkey"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslcertkey_sslocspresponder_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslcipher"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslcipher_sslciphersuite_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslcrl"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/ssldtlsprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslfipskey"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslhsmkey"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/ssllogprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslocspresponder"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslpolicylabel"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslpolicylabel_sslpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslprofile_ecccurve_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslprofile_sslcertkey_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslprofile_sslcipher_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslservice"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslservice_ecccurve_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslservice_sslcertkey_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslservice_sslciphersuite_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslservicegroup"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslservicegroup_ecccurve_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslservicegroup_sslcertkey_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslservicegroup_sslciphersuite_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslvserver"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslvserver_ecccurve_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslvserver_sslcertkey_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslvserver_sslciphersuite_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslvserver_sslpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/streamidentifier"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/streamselector"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/subscribergxinterface"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/subscriberparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/subscriberprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/subscriberradiusinterface"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/systembackup"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/systemcmdpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/systemextramgmtcpu"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/systemfile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/systemglobal_auditnslogpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/systemglobal_authenticationldappolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/systemglobal_authenticationlocalpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/systemglobal_authenticationpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/systemglobal_authenticationradiuspolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/systemglobal_authenticationtacacspolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/systemgroup"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/systemgroup_nspartition_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/systemgroup_systemcmdpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/systemgroup_systemuser_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/systemparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/systemuser"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/systemuser_nspartition_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/systemuser_systemcmdpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/tmformssoaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/tmglobal_auditnslogpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/tmglobal_auditsyslogpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/tmglobal_tmtrafficpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/tmsamlssoprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/tmsessionaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/tmsessionparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/tmsessionpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/tmtrafficaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/tmtrafficpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/transformaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/transformglobal_transformpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/transformpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/transformpolicylabel"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/transformpolicylabel_transformpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/transformprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/tunnelglobal_tunneltrafficpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/tunneltrafficpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/userprotocol"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/uservserver"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/videooptimizationdetectionaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/videooptimizationdetectionpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/videooptimizationpacingaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/videooptimizationpacingpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vlan"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vlan_channel_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vlan_interface_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vlan_nsip6_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vlan_nsip_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnalwaysonprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnclientlessaccesspolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnclientlessaccessprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpneula"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnformssoaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnglobal_appcontroller_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnglobal_auditsyslogpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnglobal_authenticationcertpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnglobal_authenticationldappolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnglobal_authenticationlocalpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnglobal_authenticationnegotiatepolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnglobal_authenticationradiuspolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnglobal_authenticationsamlpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnglobal_authenticationtacacspolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnglobal_domain_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnglobal_intranetip6_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnglobal_intranetip_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnglobal_sharefileserver_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnglobal_sslcertkey_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnglobal_staserver_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnglobal_vpnclientlessaccesspolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnglobal_vpneula_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnglobal_vpnintranetapplication_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnglobal_vpnnexthopserver_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnglobal_vpnportaltheme_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnglobal_vpnsessionpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnglobal_vpntrafficpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnglobal_vpnurl_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnglobal_vpnurlpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnintranetapplication"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnnexthopserver"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnpcoipprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnpcoipvserverprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnportaltheme"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnsamlssoprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnsessionaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnsessionpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpntrafficaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpntrafficpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnurl"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnurlaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnurlpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_aaapreauthenticationpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_analyticsprofile_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_appcontroller_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_appflowpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_appfwpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_auditnslogpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_auditsyslogpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_authenticationcertpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_authenticationdfapolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_authenticationldappolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_authenticationlocalpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_authenticationloginschemapolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_authenticationnegotiatepolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_authenticationoauthidppolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_authenticationradiuspolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_authenticationsamlidppolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_authenticationsamlpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_authenticationtacacspolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_authenticationwebauthpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_cachepolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_cspolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_feopolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_icapolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_intranetip6_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_intranetip_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_responderpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_rewritepolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_sharefileserver_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_staserver_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_vpnclientlessaccesspolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_vpneula_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_vpnintranetapplication_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_vpnnexthopserver_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_vpnportaltheme_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_vpnsessionpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_vpntrafficpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_vpnurl_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnvserver_vpnurlpolicy_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vrid"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vrid6"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vridparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vxlan"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vxlan_nsip6_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vxlan_nsip_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vxlan_srcip_binding"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vxlanvlanmap"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vxlanvlanmap_vxlan_binding"
)

// Ensure CitrixAdcFrameworkProvider satisfies various provider interfaces.
var _ provider.Provider = &CitrixAdcFrameworkProvider{}

// CitrixAdcFrameworkProvider defines the provider implementation.
type CitrixAdcFrameworkProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// CitrixAdcFrameworkProviderModel describes the provider data model.
type CitrixAdcFrameworkProviderModel struct {
	Username           types.String `tfsdk:"username"`
	Password           types.String `tfsdk:"password"`
	Endpoint           types.String `tfsdk:"endpoint"`
	InsecureSkipVerify types.Bool   `tfsdk:"insecure_skip_verify"`
	ProxiedNs          types.String `tfsdk:"proxied_ns"`
	Partition          types.String `tfsdk:"partition"`
	DoLogin            types.Bool   `tfsdk:"do_login"`
	IsCloud            types.Bool   `tfsdk:"is_cloud"`
}

// ProviderData contains the configured client for data sources and resources.
type ProviderData struct {
	Client   *adcnitrogoservice.NitroClient
	Username string
	Password string
	Endpoint string
}

func (p *CitrixAdcFrameworkProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "citrixadc"
	resp.Version = p.version
}

func (p *CitrixAdcFrameworkProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"username": schema.StringAttribute{
				Description: "Username to login to the NetScaler",
				Optional:    true,
			},
			"password": schema.StringAttribute{
				Description: "Password to login to the NetScaler",
				Optional:    true,
			},
			"endpoint": schema.StringAttribute{
				Description: "The URL to the API",
				Optional:    true,
			},
			"insecure_skip_verify": schema.BoolAttribute{
				Description: "Ignore validity of endpoint TLS certificate if true",
				Optional:    true,
			},
			"proxied_ns": schema.StringAttribute{
				Description: "Target NS ip. When defined username, password and endpoint must refer to NetScaler Console.",
				Optional:    true,
			},
			"partition": schema.StringAttribute{
				Description: "Partition to target",
				Optional:    true,
			},
			"do_login": schema.BoolAttribute{
				Description: "Perform login to NetScaler",
				Optional:    true,
			},
			"is_cloud": schema.BoolAttribute{
				Description: "Set to true when using NetScaler Console Cloud",
				Optional:    true,
			},
		},
	}
}

func (p *CitrixAdcFrameworkProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data CitrixAdcFrameworkProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.
	// Validate required parameters
	username := data.Username.ValueString()
	if username == "" {
		username = os.Getenv("NS_LOGIN")
	}

	password := data.Password.ValueString()
	if password == "" {
		password = os.Getenv("NS_PASSWORD")
	}

	endpoint := data.Endpoint.ValueString()
	if endpoint == "" {
		endpoint = os.Getenv("NS_URL")
	}

	// Check if required parameters are empty and add errors
	if username == "" {
		resp.Diagnostics.AddError(
			"Missing required parameter",
			"The 'username' parameter is required. It can be set via the provider configuration or the NS_LOGIN environment variable.",
		)
	}

	if password == "" {
		resp.Diagnostics.AddError(
			"Missing required parameter",
			"The 'password' parameter is required. It can be set via the provider configuration or the NS_PASSWORD environment variable.",
		)
	}

	if endpoint == "" {
		resp.Diagnostics.AddError(
			"Missing required parameter",
			"The 'endpoint' parameter is required. It can be set via the provider configuration or the NS_URL environment variable.",
		)
	}

	// Return early if any required parameters are missing
	if resp.Diagnostics.HasError() {
		return
	}

	proxiedNs := os.Getenv("_MPS_API_PROXY_MANAGED_INSTANCE_IP")
	if !data.ProxiedNs.IsNull() {
		proxiedNs = data.ProxiedNs.ValueString()
	}

	partition := os.Getenv("NS_PARTITION")
	if !data.Partition.IsNull() {
		partition = data.Partition.ValueString()
	}

	insecureSkipVerify := false
	if !data.InsecureSkipVerify.IsNull() {
		insecureSkipVerify = data.InsecureSkipVerify.ValueBool()
	}

	doLogin := false
	if !data.DoLogin.IsNull() {
		doLogin = data.DoLogin.ValueBool()
	}

	isCloud := false
	if !data.IsCloud.IsNull() {
		isCloud = data.IsCloud.ValueBool()
	}

	userHeaders := map[string]string{
		"User-Agent": "terraform-ctxadc",
	}

	params := adcnitrogoservice.NitroParams{
		Url:       endpoint,
		Username:  username,
		Password:  password,
		ProxiedNs: proxiedNs,
		SslVerify: !insecureSkipVerify,
		Headers:   userHeaders,
		IsCloud:   isCloud,
	}

	client, err := adcnitrogoservice.NewNitroClientFromParams(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create Nitro client",
			"Unable to create client:\n\n"+err.Error(),
		)
		return
	}

	if doLogin {
		client.Login()
	}

	if partition != "" {
		nspartition := make(map[string]interface{})
		nspartition["partitionname"] = partition
		err := client.ActOnResource("nspartition", &nspartition, "Switch")
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to switch partition",
				"Unable to switch partition:\n\n"+err.Error(),
			)
			return
		}
	}

	resp.DataSourceData = &client
	resp.ResourceData = &client

	tflog.Info(ctx, "Configured CitrixADC Framework Provider", map[string]any{"success": true})
}

func (p *CitrixAdcFrameworkProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		sslcertkey.NewSslCertKeyResource,
		sslcertkey.NewSslCertKeyUpdateResource,
		vpnvserver_appfwpolicy_binding.NewVpnvserverAppfwpolicyBindingResource,
		nslaslicense_offline.NewNSLASLicenseOfflineResource,
		csvserver_lbvserver_binding.NewCsvserverLbvserverBindingResource,
		authenticationldapaction.NewAuthenticationldapactionResource,
		authenticationradiusaction.NewAuthenticationradiusactionResource,
		authenticationtacacsaction.NewAuthenticationtacacsactionResource,
		lbparameter.NewLbparameterResource,
		nsrpcnode.NewNsrpcnodeResource,
		snmpuser.NewSnmpuserResource,
		systemuser.NewSystemuserResource,
	}
}

func (p *CitrixAdcFrameworkProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		sslcertkey.SslCertKeyDataSource,
		aaacertparams.AAacertparamsDataSource,
		aaaglobal_aaapreauthenticationpolicy_binding.AAaglobalAaapreauthenticationpolicyBindingDataSource,
		aaagroup.AAagroupDataSource,
		aaagroup_aaauser_binding.AAagroupAaauserBindingDataSource,
		aaakcdaccount.AAakcdaccountDataSource,
		aaaldapparams.AAaldapparamsDataSource,
		aaaotpparameter.AAaotpparameterDataSource,
		aaaparameter.AAaparameterDataSource,
		aaapreauthenticationaction.AAapreauthenticationactionDataSource,
		aaapreauthenticationparameter.AAapreauthenticationparameterDataSource,
		aaapreauthenticationpolicy.AAapreauthenticationpolicyDataSource,
		aaaradiusparams.AAaradiusparamsDataSource,
		aaassoprofile.AAassoprofileDataSource,
		aaatacacsparams.AAatacacsparamsDataSource,
		aaauser.AAauserDataSource,
		analyticsprofile.ANalyticsprofileDataSource,
		appalgparam.APpalgparamDataSource,
		appflowaction.APpflowactionDataSource,
		appflowcollector.APpflowcollectorDataSource,
		appflowparam.APpflowparamDataSource,
		appflowpolicy.APpflowpolicyDataSource,
		appflowpolicylabel.APpflowpolicylabelDataSource,
		appfwconfidfield.APpfwconfidfieldDataSource,
		appfwfieldtype.APpfwfieldtypeDataSource,
		appfwhtmlerrorpage.APpfwhtmlerrorpageDataSource,
		appfwjsoncontenttype.APpfwjsoncontenttypeDataSource,
		appfwjsonerrorpage.APpfwjsonerrorpageDataSource,
		appfwlearningsettings.APpfwlearningsettingsDataSource,
		appfwmultipartformcontenttype.APpfwmultipartformcontenttypeDataSource,
		appfwpolicy.APpfwpolicyDataSource,
		appfwpolicylabel.APpfwpolicylabelDataSource,
		appfwprofile.APpfwprofileDataSource,
		appfwsettings.APpfwsettingsDataSource,
		appfwsignatures.APpfwsignaturesDataSource,
		appfwurlencodedformcontenttype.APpfwurlencodedformcontenttypeDataSource,
		appfwwsdl.APpfwwsdlDataSource,
		appfwxmlcontenttype.APpfwxmlcontenttypeDataSource,
		appfwxmlerrorpage.APpfwxmlerrorpageDataSource,
		appfwxmlschema.APpfwxmlschemaDataSource,
		appqoeaction.APpqoeactionDataSource,
		appqoecustomresp.APpqoecustomrespDataSource,
		appqoeparameter.APpqoeparameterDataSource,
		appqoepolicy.APpqoepolicyDataSource,
		arp.ARpDataSource,
		arpparam.ARpparamDataSource,
		auditmessageaction.AUditmessageactionDataSource,
		auditnslogaction.AUditnslogactionDataSource,
		auditnslogparams.AUditnslogparamsDataSource,
		auditnslogpolicy.AUditnslogpolicyDataSource,
		auditsyslogaction.AUditsyslogactionDataSource,
		auditsyslogparams.AUditsyslogparamsDataSource,
		auditsyslogpolicy.AUditsyslogpolicyDataSource,
		authenticationauthnprofile.AUthenticationauthnprofileDataSource,
		authenticationcaptchaaction.AUthenticationcaptchaactionDataSource,
		authenticationcertaction.AUthenticationcertactionDataSource,
		authenticationcertpolicy.AUthenticationcertpolicyDataSource,
		authenticationcitrixauthaction.AUthenticationcitrixauthactionDataSource,
		authenticationdfaaction.AUthenticationdfaactionDataSource,
		authenticationdfapolicy.AUthenticationdfapolicyDataSource,
		authenticationemailaction.AUthenticationemailactionDataSource,
		authenticationepaaction.AUthenticationepaactionDataSource,
		authenticationldapaction.AUthenticationldapactionDataSource,
		authenticationldappolicy.AUthenticationldappolicyDataSource,
		authenticationlocalpolicy.AUthenticationlocalpolicyDataSource,
		authenticationloginschema.AUthenticationloginschemaDataSource,
		authenticationloginschemapolicy.AUthenticationloginschemapolicyDataSource,
		authenticationnegotiateaction.AUthenticationnegotiateactionDataSource,
		authenticationnegotiatepolicy.AUthenticationnegotiatepolicyDataSource,
		authenticationnoauthaction.AUthenticationnoauthactionDataSource,
		authenticationoauthaction.AUthenticationoauthactionDataSource,
		authenticationoauthidppolicy.AUthenticationoauthidppolicyDataSource,
		authenticationoauthidpprofile.AUthenticationoauthidpprofileDataSource,
		authenticationpolicy.AUthenticationpolicyDataSource,
		authenticationpolicylabel.AUthenticationpolicylabelDataSource,
		authenticationpushservice.AUthenticationpushserviceDataSource,
		authenticationradiusaction.AUthenticationradiusactionDataSource,
		authenticationradiuspolicy.AUthenticationradiuspolicyDataSource,
		authenticationsamlaction.AUthenticationsamlactionDataSource,
		authenticationsamlidppolicy.AUthenticationsamlidppolicyDataSource,
		authenticationsamlidpprofile.AUthenticationsamlidpprofileDataSource,
		authenticationsamlpolicy.AUthenticationsamlpolicyDataSource,
		authenticationstorefrontauthaction.AUthenticationstorefrontauthactionDataSource,
		authenticationtacacsaction.AUthenticationtacacsactionDataSource,
		authenticationtacacspolicy.AUthenticationtacacspolicyDataSource,
		authenticationvserver.AUthenticationvserverDataSource,
		authenticationwebauthaction.AUthenticationwebauthactionDataSource,
		authenticationwebauthpolicy.AUthenticationwebauthpolicyDataSource,
		authorizationpolicy.AUthorizationpolicyDataSource,
		authorizationpolicylabel.AUthorizationpolicylabelDataSource,
		autoscaleaction.AUtoscaleactionDataSource,
		autoscalepolicy.AUtoscalepolicyDataSource,
		autoscaleprofile.AUtoscaleprofileDataSource,
		botpolicy.BOtpolicyDataSource,
		botpolicylabel.BOtpolicylabelDataSource,
		botprofile.BOtprofileDataSource,
		botsettings.BOtsettingsDataSource,
		botsignature.BOtsignatureDataSource,
		bridgegroup.BRidgegroupDataSource,
		bridgetable.BRidgetableDataSource,
		cachecontentgroup.CAchecontentgroupDataSource,
		cacheforwardproxy.CAcheforwardproxyDataSource,
		cacheparameter.CAcheparameterDataSource,
		cachepolicy.CAchepolicyDataSource,
		cachepolicylabel.CAchepolicylabelDataSource,
		cacheselector.CAcheselectorDataSource,
		channel.CHannelDataSource,
		clusterinstance.CLusterinstanceDataSource,
		clusternode.CLusternodeDataSource,
		clusternodegroup.CLusternodegroupDataSource,
		cmpaction.CMpactionDataSource,
		cmpparameter.CMpparameterDataSource,
		cmppolicy.CMppolicyDataSource,
		cmppolicylabel.CMppolicylabelDataSource,
		contentinspectionaction.COntentinspectionactionDataSource,
		contentinspectioncallout.COntentinspectioncalloutDataSource,
		contentinspectionparameter.COntentinspectionparameterDataSource,
		contentinspectionpolicy.COntentinspectionpolicyDataSource,
		contentinspectionpolicylabel.COntentinspectionpolicylabelDataSource,
		contentinspectionprofile.COntentinspectionprofileDataSource,
		crpolicy.CRpolicyDataSource,
		crvserver.CRvserverDataSource,
		csaction.CSactionDataSource,
		csparameter.CSparameterDataSource,
		cspolicy.CSpolicyDataSource,
		cspolicylabel.CSpolicylabelDataSource,
		csvserver.CSvserverDataSource,
		dbdbprofile.DBdbprofileDataSource,
		dbuser.DBuserDataSource,
		dnsaaaarec.DNsaaaarecDataSource,
		dnsaction.DNsactionDataSource,
		dnsaction64.DNsaction64DataSource,
		dnsaddrec.DNsaddrecDataSource,
		dnscnamerec.DNscnamerecDataSource,
		dnskey.DNskeyDataSource,
		dnsmxrec.DNsmxrecDataSource,
		dnsnameserver.DNsnameserverDataSource,
		dnsnaptrrec.DNsnaptrrecDataSource,
		dnsnsrec.DNsnsrecDataSource,
		dnsparameter.DNsparameterDataSource,
		dnspolicy.DNspolicyDataSource,
		dnspolicy64.DNspolicy64DataSource,
		dnspolicylabel.DNspolicylabelDataSource,
		dnsprofile.DNsprofileDataSource,
		dnsptrrec.DNsptrrecDataSource,
		dnssoarec.DNssoarecDataSource,
		dnssrvrec.DNssrvrecDataSource,
		dnssuffix.DNssuffixDataSource,
		dnstxtrec.DNstxtrecDataSource,
		dnsview.DNsviewDataSource,
		dnszone.DNszoneDataSource,
		extendedmemoryparam.EXtendedmemoryparamDataSource,
		feoaction.FEoactionDataSource,
		feoparameter.FEoparameterDataSource,
		feopolicy.FEopolicyDataSource,
		fis.FIsDataSource,
		forwardingsession.FOrwardingsessionDataSource,
		gslbparameter.GSlbparameterDataSource,
		gslbservice.GSlbserviceDataSource,
		gslbservicegroup.GSlbservicegroupDataSource,
		gslbsite.GSlbsiteDataSource,
		gslbvserver.GSlbvserverDataSource,
		hanode.HAnodeDataSource,
		icaaccessprofile.ICaaccessprofileDataSource,
		icaaction.ICaactionDataSource,
		icalatencyprofile.ICalatencyprofileDataSource,
		icaparameter.ICaparameterDataSource,
		icapolicy.ICapolicyDataSource,
		inat.INatDataSource,
		inatparam.INatparamDataSource,
		interfacepair.INterfacepairDataSource,
		ip6tunnel.IP6tunnelDataSource,
		ip6tunnelparam.IP6tunnelparamDataSource,
		ipsecalgprofile.IPsecalgprofileDataSource,
		ipsecparameter.IPsecparameterDataSource,
		ipsecprofile.IPsecprofileDataSource,
		ipset.IPsetDataSource,
		iptunnel.IPtunnelDataSource,
		iptunnelparam.IPtunnelparamDataSource,
		ipv6.IPv6DataSource,
		l2param.L2ParamDataSource,
		l3param.L3ParamDataSource,
		l4param.L4ParamDataSource,
		lacp.LAcpDataSource,
		lbaction.LBactionDataSource,
		lbgroup.LBgroupDataSource,
		lbmetrictable.LBmetrictableDataSource,
		lbmonitor.LBmonitorDataSource,
		lbpolicy.LBpolicyDataSource,
		lbprofile.LBprofileDataSource,
		lbroute.LBrouteDataSource,
		lbroute6.LBroute6DataSource,
		lbsipparameters.LBsipparametersDataSource,
		lbvserver.LBvserverDataSource,
		linkset.LInksetDataSource,
		lldpparam.LLdpparamDataSource,
		location.LOcationDataSource,
		locationfile.LOcationfileDataSource,
		locationfile6.LOcationfile6DataSource,
		locationparameter.LOcationparameterDataSource,
		lsnappsattributes.LSnappsattributesDataSource,
		lsnappsprofile.LSnappsprofileDataSource,
		lsnclient.LSnclientDataSource,
		lsngroup.LSngroupDataSource,
		lsnhttphdrlogprofile.LSnhttphdrlogprofileDataSource,
		lsnip6profile.LSnip6profileDataSource,
		lsnlogprofile.LSnlogprofileDataSource,
		lsnparameter.LSnparameterDataSource,
		lsnpool.LSnpoolDataSource,
		lsnrtspalgprofile.LSnrtspalgprofileDataSource,
		lsnsipalgprofile.LSnsipalgprofileDataSource,
		lsnstatic.LSnstaticDataSource,
		lsntransportprofile.LSntransportprofileDataSource,
		mapbmr.MApbmrDataSource,
		mapdmr.MApdmrDataSource,
		mapdomain.MApdomainDataSource,
		nat64.NAt64DataSource,
		nat64param.NAt64paramDataSource,
		nd6.ND6DataSource,
		nd6ravariables.ND6ravariablesDataSource,
		netbridge.NEtbridgeDataSource,
		netprofile.NEtprofileDataSource,
		nsacl.NSaclDataSource,
		nsacl6.NSacl6DataSource,
		nsacls.NSaclsDataSource,
		nsappflowcollector.NSappflowcollectorDataSource,
		nsassignment.NSassignmentDataSource,
		nscapacity.NScapacityDataSource,
		nsconsoleloginprompt.NSconsoleloginpromptDataSource,
		nscqaparam.NScqaparamDataSource,
		nsdhcpparams.NSdhcpparamsDataSource,
		nsdiameter.NSdiameterDataSource,
		nsencryptionkey.NSencryptionkeyDataSource,
		nsencryptionparams.NSencryptionparamsDataSource,
		nsfeature.NSfeatureDataSource,
		nshmackey.NShmackeyDataSource,
		nshostname.NShostnameDataSource,
		nshttpparam.NShttpparamDataSource,
		nshttpprofile.NShttpprofileDataSource,
		nsicapprofile.NSicapprofileDataSource,
		nsip.NSipDataSource,
		nsip6.NSip6DataSource,
		nslicense.NSlicenseDataSource,
		nslicenseparameters.NSlicenseparametersDataSource,
		nslicenseproxyserver.NSlicenseproxyserverDataSource,
		nslicenseserver.NSlicenseserverDataSource,
		nslimitidentifier.NSlimitidentifierDataSource,
		nsmode.NSmodeDataSource,
		nsparam.NSparamDataSource,
		nspartition.NSpartitionDataSource,
		nspbr.NSpbrDataSource,
		nspbr6.NSpbr6DataSource,
		nsratecontrol.NSratecontrolDataSource,
		nsrpcnode.NSrpcnodeDataSource,
		nsservicefunction.NSservicefunctionDataSource,
		nsservicepath.NSservicepathDataSource,
		nssimpleacl.NSsimpleaclDataSource,
		nssimpleacl6.NSsimpleacl6DataSource,
		nsspparams.NSspparamsDataSource,
		nstcpbufparam.NStcpbufparamDataSource,
		nstcpparam.NStcpparamDataSource,
		nstcpprofile.NStcpprofileDataSource,
		nstimeout.NStimeoutDataSource,
		nstimer.NStimerDataSource,
		nstrafficdomain.NStrafficdomainDataSource,
		nsvariable.NSvariableDataSource,
		nsvpxparam.NSvpxparamDataSource,
		nsweblogparam.NSweblogparamDataSource,
		nsxmlnamespace.NSxmlnamespaceDataSource,
		ntpparam.NTpparamDataSource,
		ntpserver.NTpserverDataSource,
		onlinkipv6prefix.ONlinkipv6prefixDataSource,
		pcpprofile.PCpprofileDataSource,
		pcpserver.PCpserverDataSource,
		policydataset.POlicydatasetDataSource,
		policyexpression.POlicyexpressionDataSource,
		policyhttpcallout.POlicyhttpcalloutDataSource,
		policymap.POlicymapDataSource,
		policyparam.POlicyparamDataSource,
		policypatset.POlicypatsetDataSource,
		policystringmap.POlicystringmapDataSource,
		ptp.PTpDataSource,
		quicbridgeprofile.QUicbridgeprofileDataSource,
		radiusnode.RAdiusnodeDataSource,
		rdpclientprofile.RDpclientprofileDataSource,
		rdpserverprofile.RDpserverprofileDataSource,
		reputationsettings.REputationsettingsDataSource,
		responderaction.REsponderactionDataSource,
		responderhtmlpage.REsponderhtmlpageDataSource,
		responderparam.REsponderparamDataSource,
		responderpolicy.REsponderpolicyDataSource,
		responderpolicylabel.REsponderpolicylabelDataSource,
		rewriteaction.REwriteactionDataSource,
		rewriteparam.REwriteparamDataSource,
		rewritepolicy.REwritepolicyDataSource,
		rewritepolicylabel.REwritepolicylabelDataSource,
		rnat.RNatDataSource,
		rnat6.RNat6DataSource,
		rnatparam.RNatparamDataSource,
		route.ROuteDataSource,
		route6.ROute6DataSource,
		routerdynamicrouting.ROuterdynamicroutingDataSource,
		rsskeytype.RSskeytypeDataSource,
		server.SErverDataSource,
		service.SErviceDataSource,
		servicegroup.SErvicegroupDataSource,
		smppparam.SMppparamDataSource,
		smppuser.SMppuserDataSource,
		snmpalarm.SNmpalarmDataSource,
		snmpcommunity.SNmpcommunityDataSource,
		snmpengineid.SNmpengineidDataSource,
		snmpgroup.SNmpgroupDataSource,
		snmpmanager.SNmpmanagerDataSource,
		snmpmib.SNmpmibDataSource,
		snmpoption.SNmpoptionDataSource,
		snmptrap.SNmptrapDataSource,
		snmpuser.SNmpuserDataSource,
		snmpview.SNmpviewDataSource,
		spilloveraction.SPilloveractionDataSource,
		spilloverpolicy.SPilloverpolicyDataSource,
		sslaction.SSlactionDataSource,
		sslcacertgroup.SSlcacertgroupDataSource,
		sslcertfile.SSlcertfileDataSource,
		sslcipher.SSlcipherDataSource,
		sslcrl.SSlcrlDataSource,
		ssldtlsprofile.SSldtlsprofileDataSource,
		sslfipskey.SSlfipskeyDataSource,
		sslhsmkey.SSlhsmkeyDataSource,
		ssllogprofile.SSllogprofileDataSource,
		sslocspresponder.SSlocspresponderDataSource,
		sslparameter.SSlparameterDataSource,
		sslpolicy.SSlpolicyDataSource,
		sslpolicylabel.SSlpolicylabelDataSource,
		sslprofile.SSlprofileDataSource,
		sslservice.SSlserviceDataSource,
		sslservicegroup.SSlservicegroupDataSource,
		sslvserver.SSlvserverDataSource,
		streamidentifier.STreamidentifierDataSource,
		streamselector.STreamselectorDataSource,
		subscribergxinterface.SUbscribergxinterfaceDataSource,
		subscriberparam.SUbscriberparamDataSource,
		subscriberprofile.SUbscriberprofileDataSource,
		subscriberradiusinterface.SUbscriberradiusinterfaceDataSource,
		systembackup.SYstembackupDataSource,
		systemcmdpolicy.SYstemcmdpolicyDataSource,
		systemextramgmtcpu.SYstemextramgmtcpuDataSource,
		systemfile.SYstemfileDataSource,
		systemgroup.SYstemgroupDataSource,
		systemparameter.SYstemparameterDataSource,
		systemuser.SYstemuserDataSource,
		tmformssoaction.TMformssoactionDataSource,
		tmsamlssoprofile.TMsamlssoprofileDataSource,
		tmsessionaction.TMsessionactionDataSource,
		tmsessionparameter.TMsessionparameterDataSource,
		tmsessionpolicy.TMsessionpolicyDataSource,
		tmtrafficaction.TMtrafficactionDataSource,
		tmtrafficpolicy.TMtrafficpolicyDataSource,
		transformaction.TRansformactionDataSource,
		transformpolicy.TRansformpolicyDataSource,
		transformpolicylabel.TRansformpolicylabelDataSource,
		transformprofile.TRansformprofileDataSource,
		tunneltrafficpolicy.TUnneltrafficpolicyDataSource,
		userprotocol.USerprotocolDataSource,
		uservserver.UServserverDataSource,
		videooptimizationdetectionaction.VIdeooptimizationdetectionactionDataSource,
		videooptimizationdetectionpolicy.VIdeooptimizationdetectionpolicyDataSource,
		videooptimizationpacingaction.VIdeooptimizationpacingactionDataSource,
		videooptimizationpacingpolicy.VIdeooptimizationpacingpolicyDataSource,
		vlan.VLanDataSource,
		vpnalwaysonprofile.VPnalwaysonprofileDataSource,
		vpnclientlessaccesspolicy.VPnclientlessaccesspolicyDataSource,
		vpnclientlessaccessprofile.VPnclientlessaccessprofileDataSource,
		vpneula.VPneulaDataSource,
		vpnformssoaction.VPnformssoactionDataSource,
		vpnintranetapplication.VPnintranetapplicationDataSource,
		vpnnexthopserver.VPnnexthopserverDataSource,
		vpnparameter.VPnparameterDataSource,
		vpnpcoipprofile.VPnpcoipprofileDataSource,
		vpnpcoipvserverprofile.VPnpcoipvserverprofileDataSource,
		vpnportaltheme.VPnportalthemeDataSource,
		vpnsamlssoprofile.VPnsamlssoprofileDataSource,
		vpnsessionaction.VPnsessionactionDataSource,
		vpnsessionpolicy.VPnsessionpolicyDataSource,
		vpntrafficaction.VPntrafficactionDataSource,
		vpntrafficpolicy.VPntrafficpolicyDataSource,
		vpnurl.VPnurlDataSource,
		vpnurlaction.VPnurlactionDataSource,
		vpnurlpolicy.VPnurlpolicyDataSource,
		vpnvserver.VPnvserverDataSource,
		vrid.VRidDataSource,
		vrid6.VRid6DataSource,
		vridparam.VRidparamDataSource,
		vxlan.VXlanDataSource,
		vxlanvlanmap.VXlanvlanmapDataSource,
		Interface.INterfaceDataSource,
		aaagroup_auditnslogpolicy_binding.AAagroupAuditnslogpolicyBindingDataSource,
		aaagroup_auditsyslogpolicy_binding.AAagroupAuditsyslogpolicyBindingDataSource,
		aaagroup_authorizationpolicy_binding.AAagroupAuthorizationpolicyBindingDataSource,
		aaagroup_intranetip_binding.AAagroupIntranetipBindingDataSource,
		aaagroup_tmsessionpolicy_binding.AAagroupTmsessionpolicyBindingDataSource,
		aaagroup_vpnintranetapplication_binding.AAagroupVpnintranetapplicationBindingDataSource,
		aaagroup_vpnsessionpolicy_binding.AAagroupVpnsessionpolicyBindingDataSource,
		aaagroup_vpntrafficpolicy_binding.AAagroupVpntrafficpolicyBindingDataSource,
		aaagroup_vpnurl_binding.AAagroupVpnurlBindingDataSource,
		aaagroup_vpnurlpolicy_binding.AAagroupVpnurlpolicyBindingDataSource,
		aaauser_auditnslogpolicy_binding.AAauserAuditnslogpolicyBindingDataSource,
		aaauser_auditsyslogpolicy_binding.AAauserAuditsyslogpolicyBindingDataSource,
		aaauser_authorizationpolicy_binding.AAauserAuthorizationpolicyBindingDataSource,
		aaauser_intranetip6_binding.AAauserIntranetip6BindingDataSource,
		aaauser_intranetip_binding.AAauserIntranetipBindingDataSource,
		aaauser_tmsessionpolicy_binding.AAauserTmsessionpolicyBindingDataSource,
		aaauser_vpnintranetapplication_binding.AAauserVpnintranetapplicationBindingDataSource,
		aaauser_vpnsessionpolicy_binding.AAauserVpnsessionpolicyBindingDataSource,
		aaauser_vpntrafficpolicy_binding.AAauserVpntrafficpolicyBindingDataSource,
		aaauser_vpnurl_binding.AAauserVpnurlBindingDataSource,
		aaauser_vpnurlpolicy_binding.AAauserVpnurlpolicyBindingDataSource,
		analyticsglobal_analyticsprofile_binding.ANalyticsglobalAnalyticsprofileBindingDataSource,
		appflowaction_analyticsprofile_binding.APpflowactionAnalyticsprofileBindingDataSource,
		appflowglobal_appflowpolicy_binding.APpflowglobalAppflowpolicyBindingDataSource,
		appflowpolicylabel_appflowpolicy_binding.APpflowpolicylabelAppflowpolicyBindingDataSource,
		appfwglobal_appfwpolicy_binding.APpfwglobalAppfwpolicyBindingDataSource,
		appfwglobal_auditnslogpolicy_binding.APpfwglobalAuditnslogpolicyBindingDataSource,
		appfwglobal_auditsyslogpolicy_binding.APpfwglobalAuditsyslogpolicyBindingDataSource,
		appfwpolicylabel_appfwpolicy_binding.APpfwpolicylabelAppfwpolicyBindingDataSource,
		appfwprofile_cmdinjection_binding.APpfwprofileCmdinjectionBindingDataSource,
		appfwprofile_contenttype_binding.APpfwprofileContenttypeBindingDataSource,
		appfwprofile_cookieconsistency_binding.APpfwprofileCookieconsistencyBindingDataSource,
		appfwprofile_creditcardnumber_binding.APpfwprofileCreditcardnumberBindingDataSource,
		appfwprofile_crosssitescripting_binding.APpfwprofileCrosssitescriptingBindingDataSource,
		appfwprofile_csrftag_binding.APpfwprofileCsrftagBindingDataSource,
		appfwprofile_denyurl_binding.APpfwprofileDenyurlBindingDataSource,
		appfwprofile_excluderescontenttype_binding.APpfwprofileExcluderescontenttypeBindingDataSource,
		appfwprofile_fieldconsistency_binding.APpfwprofileFieldconsistencyBindingDataSource,
		appfwprofile_fieldformat_binding.APpfwprofileFieldformatBindingDataSource,
		appfwprofile_fileuploadtype_binding.APpfwprofileFileuploadtypeBindingDataSource,
		appfwprofile_jsoncmdurl_binding.APpfwprofileJsoncmdurlBindingDataSource,
		appfwprofile_jsondosurl_binding.APpfwprofileJsondosurlBindingDataSource,
		appfwprofile_jsonsqlurl_binding.APpfwprofileJsonsqlurlBindingDataSource,
		appfwprofile_jsonxssurl_binding.APpfwprofileJsonxssurlBindingDataSource,
		appfwprofile_logexpression_binding.APpfwprofileLogexpressionBindingDataSource,
		appfwprofile_safeobject_binding.APpfwprofileSafeobjectBindingDataSource,
		appfwprofile_sqlinjection_binding.APpfwprofileSqlinjectionBindingDataSource,
		appfwprofile_starturl_binding.APpfwprofileStarturlBindingDataSource,
		appfwprofile_trustedlearningclients_binding.APpfwprofileTrustedlearningclientsBindingDataSource,
		appfwprofile_xmlattachmenturl_binding.APpfwprofileXmlattachmenturlBindingDataSource,
		appfwprofile_xmldosurl_binding.APpfwprofileXmldosurlBindingDataSource,
		appfwprofile_xmlsqlinjection_binding.APpfwprofileXmlsqlinjectionBindingDataSource,
		appfwprofile_xmlvalidationurl_binding.APpfwprofileXmlvalidationurlBindingDataSource,
		appfwprofile_xmlwsiurl_binding.APpfwprofileXmlwsiurlBindingDataSource,
		appfwprofile_xmlxss_binding.APpfwprofileXmlxssBindingDataSource,
		auditnslogglobal_auditnslogpolicy_binding.AUditnslogglobalAuditnslogpolicyBindingDataSource,
		auditsyslogglobal_auditsyslogpolicy_binding.AUditsyslogglobalAuditsyslogpolicyBindingDataSource,
		authenticationpolicylabel_authenticationpolicy_binding.AUthenticationpolicylabelAuthenticationpolicyBindingDataSource,
		authenticationvserver_auditnslogpolicy_binding.AUthenticationvserverAuditnslogpolicyBindingDataSource,
		authenticationvserver_auditsyslogpolicy_binding.AUthenticationvserverAuditsyslogpolicyBindingDataSource,
		authenticationvserver_authenticationcertpolicy_binding.AUthenticationvserverAuthenticationcertpolicyBindingDataSource,
		authenticationvserver_authenticationldappolicy_binding.AUthenticationvserverAuthenticationldappolicyBindingDataSource,
		authenticationvserver_authenticationlocalpolicy_binding.AUthenticationvserverAuthenticationlocalpolicyBindingDataSource,
		authenticationvserver_authenticationloginschemapolicy_binding.AUthenticationvserverAuthenticationloginschemapolicyBindingDataSource,
		authenticationvserver_authenticationnegotiatepolicy_binding.AUthenticationvserverAuthenticationnegotiatepolicyBindingDataSource,
		authenticationvserver_authenticationoauthidppolicy_binding.AUthenticationvserverAuthenticationoauthidppolicyBindingDataSource,
		authenticationvserver_authenticationpolicy_binding.AUthenticationvserverAuthenticationpolicyBindingDataSource,
		authenticationvserver_authenticationradiuspolicy_binding.AUthenticationvserverAuthenticationradiuspolicyBindingDataSource,
		authenticationvserver_authenticationsamlidppolicy_binding.AUthenticationvserverAuthenticationsamlidppolicyBindingDataSource,
		authenticationvserver_authenticationsamlpolicy_binding.AUthenticationvserverAuthenticationsamlpolicyBindingDataSource,
		authenticationvserver_authenticationtacacspolicy_binding.AUthenticationvserverAuthenticationtacacspolicyBindingDataSource,
		authenticationvserver_authenticationwebauthpolicy_binding.AUthenticationvserverAuthenticationwebauthpolicyBindingDataSource,
		authenticationvserver_cachepolicy_binding.AUthenticationvserverCachepolicyBindingDataSource,
		authenticationvserver_cspolicy_binding.AUthenticationvserverCspolicyBindingDataSource,
		authenticationvserver_responderpolicy_binding.AUthenticationvserverResponderpolicyBindingDataSource,
		authenticationvserver_rewritepolicy_binding.AUthenticationvserverRewritepolicyBindingDataSource,
		authenticationvserver_tmsessionpolicy_binding.AUthenticationvserverTmsessionpolicyBindingDataSource,
		authenticationvserver_vpnportaltheme_binding.AUthenticationvserverVpnportalthemeBindingDataSource,
		authorizationpolicylabel_authorizationpolicy_binding.AUthorizationpolicylabelAuthorizationpolicyBindingDataSource,
		botglobal_botpolicy_binding.BOtglobalBotpolicyBindingDataSource,
		botpolicylabel_botpolicy_binding.BOtpolicylabelBotpolicyBindingDataSource,
		botprofile_blacklist_binding.BOtprofileBlacklistBindingDataSource,
		botprofile_captcha_binding.BOtprofileCaptchaBindingDataSource,
		botprofile_ipreputation_binding.BOtprofileIpreputationBindingDataSource,
		botprofile_logexpression_binding.BOtprofileLogexpressionBindingDataSource,
		botprofile_ratelimit_binding.BOtprofileRatelimitBindingDataSource,
		botprofile_tps_binding.BOtprofileTpsBindingDataSource,
		botprofile_trapinsertionurl_binding.BOtprofileTrapinsertionurlBindingDataSource,
		botprofile_whitelist_binding.BOtprofileWhitelistBindingDataSource,
		bridgegroup_nsip6_binding.BRidgegroupNsip6BindingDataSource,
		bridgegroup_nsip_binding.BRidgegroupNsipBindingDataSource,
		bridgegroup_vlan_binding.BRidgegroupVlanBindingDataSource,
		cacheglobal_cachepolicy_binding.CAcheglobalCachepolicyBindingDataSource,
		cachepolicylabel_cachepolicy_binding.CAchepolicylabelCachepolicyBindingDataSource,
		clusternode_routemonitor_binding.CLusternodeRoutemonitorBindingDataSource,
		clusternodegroup_authenticationvserver_binding.CLusternodegroupAuthenticationvserverBindingDataSource,
		clusternodegroup_clusternode_binding.CLusternodegroupClusternodeBindingDataSource,
		clusternodegroup_crvserver_binding.CLusternodegroupCrvserverBindingDataSource,
		clusternodegroup_csvserver_binding.CLusternodegroupCsvserverBindingDataSource,
		clusternodegroup_gslbsite_binding.CLusternodegroupGslbsiteBindingDataSource,
		clusternodegroup_gslbvserver_binding.CLusternodegroupGslbvserverBindingDataSource,
		clusternodegroup_lbvserver_binding.CLusternodegroupLbvserverBindingDataSource,
		clusternodegroup_nslimitidentifier_binding.CLusternodegroupNslimitidentifierBindingDataSource,
		clusternodegroup_service_binding.CLusternodegroupServiceBindingDataSource,
		clusternodegroup_streamidentifier_binding.CLusternodegroupStreamidentifierBindingDataSource,
		clusternodegroup_vpnvserver_binding.CLusternodegroupVpnvserverBindingDataSource,
		cmpglobal_cmppolicy_binding.CMpglobalCmppolicyBindingDataSource,
		cmppolicylabel_cmppolicy_binding.CMppolicylabelCmppolicyBindingDataSource,
		contentinspectionglobal_contentinspectionpolicy_binding.COntentinspectionglobalContentinspectionpolicyBindingDataSource,
		contentinspectionpolicylabel_contentinspectionpolicy_binding.COntentinspectionpolicylabelContentinspectionpolicyBindingDataSource,
		crvserver_analyticsprofile_binding.CRvserverAnalyticsprofileBindingDataSource,
		crvserver_appflowpolicy_binding.CRvserverAppflowpolicyBindingDataSource,
		crvserver_appfwpolicy_binding.CRvserverAppfwpolicyBindingDataSource,
		crvserver_appqoepolicy_binding.CRvserverAppqoepolicyBindingDataSource,
		crvserver_cachepolicy_binding.CRvserverCachepolicyBindingDataSource,
		crvserver_cmppolicy_binding.CRvserverCmppolicyBindingDataSource,
		crvserver_crpolicy_binding.CRvserverCrpolicyBindingDataSource,
		crvserver_cspolicy_binding.CRvserverCspolicyBindingDataSource,
		crvserver_feopolicy_binding.CRvserverFeopolicyBindingDataSource,
		crvserver_icapolicy_binding.CRvserverIcapolicyBindingDataSource,
		crvserver_lbvserver_binding.CRvserverLbvserverBindingDataSource,
		crvserver_policymap_binding.CRvserverPolicymapBindingDataSource,
		crvserver_responderpolicy_binding.CRvserverResponderpolicyBindingDataSource,
		crvserver_rewritepolicy_binding.CRvserverRewritepolicyBindingDataSource,
		crvserver_spilloverpolicy_binding.CRvserverSpilloverpolicyBindingDataSource,
		csvserver_analyticsprofile_binding.CSvserverAnalyticsprofileBindingDataSource,
		csvserver_appflowpolicy_binding.CSvserverAppflowpolicyBindingDataSource,
		csvserver_appfwpolicy_binding.CSvserverAppfwpolicyBindingDataSource,
		csvserver_appqoepolicy_binding.CSvserverAppqoepolicyBindingDataSource,
		csvserver_auditnslogpolicy_binding.CSvserverAuditnslogpolicyBindingDataSource,
		csvserver_auditsyslogpolicy_binding.CSvserverAuditsyslogpolicyBindingDataSource,
		csvserver_authorizationpolicy_binding.CSvserverAuthorizationpolicyBindingDataSource,
		csvserver_botpolicy_binding.CSvserverBotpolicyBindingDataSource,
		csvserver_cachepolicy_binding.CSvserverCachepolicyBindingDataSource,
		csvserver_cmppolicy_binding.CSvserverCmppolicyBindingDataSource,
		csvserver_contentinspectionpolicy_binding.CSvserverContentinspectionpolicyBindingDataSource,
		csvserver_cspolicy_binding.CSvserverCspolicyBindingDataSource,
		csvserver_feopolicy_binding.CSvserverFeopolicyBindingDataSource,
		csvserver_gslbvserver_binding.CSvserverGslbvserverBindingDataSource,
		csvserver_responderpolicy_binding.CSvserverResponderpolicyBindingDataSource,
		csvserver_rewritepolicy_binding.CSvserverRewritepolicyBindingDataSource,
		csvserver_spilloverpolicy_binding.CSvserverSpilloverpolicyBindingDataSource,
		csvserver_tmtrafficpolicy_binding.CSvserverTmtrafficpolicyBindingDataSource,
		csvserver_transformpolicy_binding.CSvserverTransformpolicyBindingDataSource,
		csvserver_vpnvserver_binding.CSvserverVpnvserverBindingDataSource,
		dnsglobal_dnspolicy_binding.DNsglobalDnspolicyBindingDataSource,
		dnspolicylabel_dnspolicy_binding.DNspolicylabelDnspolicyBindingDataSource,
		feoglobal_feopolicy_binding.FEoglobalFeopolicyBindingDataSource,
		gslbservice_dnsview_binding.GSlbserviceDnsviewBindingDataSource,
		gslbservice_lbmonitor_binding.GSlbserviceLbmonitorBindingDataSource,
		gslbservicegroup_gslbservicegroupmember_binding.GSlbservicegroupGslbservicegroupmemberBindingDataSource,
		gslbservicegroup_lbmonitor_binding.GSlbservicegroupLbmonitorBindingDataSource,
		gslbvserver_domain_binding.GSlbvserverDomainBindingDataSource,
		gslbvserver_gslbservice_binding.GSlbvserverGslbserviceBindingDataSource,
		gslbvserver_gslbservicegroup_binding.GSlbvserverGslbservicegroupBindingDataSource,
		gslbvserver_lbpolicy_binding.GSlbvserverLbpolicyBindingDataSource,
		gslbvserver_spilloverpolicy_binding.GSlbvserverSpilloverpolicyBindingDataSource,
		hanode_routemonitor6_binding.HAnodeRoutemonitor6BindingDataSource,
		hanode_routemonitor_binding.HAnodeRoutemonitorBindingDataSource,
		icaglobal_icapolicy_binding.ICaglobalIcapolicyBindingDataSource,
		ipset_nsip6_binding.IPsetNsip6BindingDataSource,
		ipset_nsip_binding.IPsetNsipBindingDataSource,
		lbgroup_lbvserver_binding.LBgroupLbvserverBindingDataSource,
		lbmetrictable_metric_binding.LBmetrictableMetricBindingDataSource,
		lbmonitor_metric_binding.LBmonitorMetricBindingDataSource,
		lbmonitor_sslcertkey_binding.LBmonitorSslcertkeyBindingDataSource,
		lbvserver_analyticsprofile_binding.LBvserverAnalyticsprofileBindingDataSource,
		lbvserver_appflowpolicy_binding.LBvserverAppflowpolicyBindingDataSource,
		lbvserver_appfwpolicy_binding.LBvserverAppfwpolicyBindingDataSource,
		lbvserver_appqoepolicy_binding.LBvserverAppqoepolicyBindingDataSource,
		lbvserver_auditsyslogpolicy_binding.LBvserverAuditsyslogpolicyBindingDataSource,
		lbvserver_authorizationpolicy_binding.LBvserverAuthorizationpolicyBindingDataSource,
		lbvserver_botpolicy_binding.LBvserverBotpolicyBindingDataSource,
		lbvserver_cachepolicy_binding.LBvserverCachepolicyBindingDataSource,
		lbvserver_cmppolicy_binding.LBvserverCmppolicyBindingDataSource,
		lbvserver_contentinspectionpolicy_binding.LBvserverContentinspectionpolicyBindingDataSource,
		lbvserver_dnspolicy64_binding.LBvserverDnspolicy64BindingDataSource,
		lbvserver_feopolicy_binding.LBvserverFeopolicyBindingDataSource,
		lbvserver_lbpolicy_binding.LBvserverLbpolicyBindingDataSource,
		lbvserver_responderpolicy_binding.LBvserverResponderpolicyBindingDataSource,
		lbvserver_rewritepolicy_binding.LBvserverRewritepolicyBindingDataSource,
		lbvserver_service_binding.LBvserverServiceBindingDataSource,
		lbvserver_servicegroup_binding.LBvserverServicegroupBindingDataSource,
		lbvserver_spilloverpolicy_binding.LBvserverSpilloverpolicyBindingDataSource,
		lbvserver_tmtrafficpolicy_binding.LBvserverTmtrafficpolicyBindingDataSource,
		lbvserver_transformpolicy_binding.LBvserverTransformpolicyBindingDataSource,
		lbvserver_videooptimizationdetectionpolicy_binding.LBvserverVideooptimizationdetectionpolicyBindingDataSource,
		lbvserver_videooptimizationpacingpolicy_binding.LBvserverVideooptimizationpacingpolicyBindingDataSource,
		vpnglobal_appcontroller_binding.VPnglobalAppcontrollerBindingDataSource,
		vpnglobal_auditsyslogpolicy_binding.VPnglobalAuditsyslogpolicyBindingDataSource,
		vpnglobal_authenticationcertpolicy_binding.VPnglobalAuthenticationcertpolicyBindingDataSource,
		vpnglobal_authenticationldappolicy_binding.VPnglobalAuthenticationldappolicyBindingDataSource,
		vpnglobal_authenticationlocalpolicy_binding.VPnglobalAuthenticationlocalpolicyBindingDataSource,
		vpnglobal_authenticationnegotiatepolicy_binding.VPnglobalAuthenticationnegotiatepolicyBindingDataSource,
		vpnglobal_authenticationradiuspolicy_binding.VPnglobalAuthenticationradiuspolicyBindingDataSource,
		vpnglobal_authenticationsamlpolicy_binding.VPnglobalAuthenticationsamlpolicyBindingDataSource,
		vpnglobal_authenticationtacacspolicy_binding.VPnglobalAuthenticationtacacspolicyBindingDataSource,
		vpnglobal_domain_binding.VPnglobalDomainBindingDataSource,
		vpnglobal_intranetip6_binding.VPnglobalIntranetip6BindingDataSource,
		vpnglobal_intranetip_binding.VPnglobalIntranetipBindingDataSource,
		vpnglobal_sharefileserver_binding.VPnglobalSharefileserverBindingDataSource,
		vpnglobal_sslcertkey_binding.VPnglobalSslcertkeyBindingDataSource,
		vpnglobal_staserver_binding.VPnglobalStaserverBindingDataSource,
		vpnglobal_vpnclientlessaccesspolicy_binding.VPnglobalVpnclientlessaccesspolicyBindingDataSource,
		vpnglobal_vpneula_binding.VPnglobalVpneulaBindingDataSource,
		vpnglobal_vpnintranetapplication_binding.VPnglobalVpnintranetapplicationBindingDataSource,
		vpnglobal_vpnnexthopserver_binding.VPnglobalVpnnexthopserverBindingDataSource,
		vpnglobal_vpnportaltheme_binding.VPnglobalVpnportalthemeBindingDataSource,
		vpnglobal_vpnsessionpolicy_binding.VPnglobalVpnsessionpolicyBindingDataSource,
		vpnglobal_vpntrafficpolicy_binding.VPnglobalVpntrafficpolicyBindingDataSource,
		vpnglobal_vpnurl_binding.VPnglobalVpnurlBindingDataSource,
		vpnglobal_vpnurlpolicy_binding.VPnglobalVpnurlpolicyBindingDataSource,
		vpnvserver_aaapreauthenticationpolicy_binding.VPnvserverAaapreauthenticationpolicyBindingDataSource,
		vpnvserver_analyticsprofile_binding.VPnvserverAnalyticsprofileBindingDataSource,
		vpnvserver_appcontroller_binding.VPnvserverAppcontrollerBindingDataSource,
		vpnvserver_appflowpolicy_binding.VPnvserverAppflowpolicyBindingDataSource,
		vpnvserver_auditnslogpolicy_binding.VPnvserverAuditnslogpolicyBindingDataSource,
		vpnvserver_auditsyslogpolicy_binding.VPnvserverAuditsyslogpolicyBindingDataSource,
		vpnvserver_authenticationcertpolicy_binding.VPnvserverAuthenticationcertpolicyBindingDataSource,
		vpnvserver_authenticationdfapolicy_binding.VPnvserverAuthenticationdfapolicyBindingDataSource,
		vpnvserver_authenticationldappolicy_binding.VPnvserverAuthenticationldappolicyBindingDataSource,
		vpnvserver_authenticationlocalpolicy_binding.VPnvserverAuthenticationlocalpolicyBindingDataSource,
		vpnvserver_authenticationloginschemapolicy_binding.VPnvserverAuthenticationloginschemapolicyBindingDataSource,
		vpnvserver_authenticationnegotiatepolicy_binding.VPnvserverAuthenticationnegotiatepolicyBindingDataSource,
		vpnvserver_authenticationoauthidppolicy_binding.VPnvserverAuthenticationoauthidppolicyBindingDataSource,
		vpnvserver_authenticationradiuspolicy_binding.VPnvserverAuthenticationradiuspolicyBindingDataSource,
		vpnvserver_authenticationsamlidppolicy_binding.VPnvserverAuthenticationsamlidppolicyBindingDataSource,
		vpnvserver_authenticationsamlpolicy_binding.VPnvserverAuthenticationsamlpolicyBindingDataSource,
		vpnvserver_authenticationtacacspolicy_binding.VPnvserverAuthenticationtacacspolicyBindingDataSource,
		vpnvserver_authenticationwebauthpolicy_binding.VPnvserverAuthenticationwebauthpolicyBindingDataSource,
		vpnvserver_cachepolicy_binding.VPnvserverCachepolicyBindingDataSource,
		vpnvserver_cspolicy_binding.VPnvserverCspolicyBindingDataSource,
		vpnvserver_feopolicy_binding.VPnvserverFeopolicyBindingDataSource,
		vpnvserver_icapolicy_binding.VPnvserverIcapolicyBindingDataSource,
		vpnvserver_intranetip6_binding.VPnvserverIntranetip6BindingDataSource,
		vpnvserver_intranetip_binding.VPnvserverIntranetipBindingDataSource,
		vpnvserver_responderpolicy_binding.VPnvserverResponderpolicyBindingDataSource,
		vpnvserver_rewritepolicy_binding.VPnvserverRewritepolicyBindingDataSource,
		vpnvserver_sharefileserver_binding.VPnvserverSharefileserverBindingDataSource,
		vpnvserver_staserver_binding.VPnvserverStaserverBindingDataSource,
		vpnvserver_vpnclientlessaccesspolicy_binding.VPnvserverVpnclientlessaccesspolicyBindingDataSource,
		vpnvserver_vpneula_binding.VPnvserverVpneulaBindingDataSource,
		vpnvserver_vpnintranetapplication_binding.VPnvserverVpnintranetapplicationBindingDataSource,
		vpnvserver_vpnnexthopserver_binding.VPnvserverVpnnexthopserverBindingDataSource,
		vpnvserver_vpnportaltheme_binding.VPnvserverVpnportalthemeBindingDataSource,
		vpnvserver_vpnsessionpolicy_binding.VPnvserverVpnsessionpolicyBindingDataSource,
		vpnvserver_vpntrafficpolicy_binding.VPnvserverVpntrafficpolicyBindingDataSource,
		vpnvserver_vpnurl_binding.VPnvserverVpnurlBindingDataSource,
		vpnvserver_vpnurlpolicy_binding.VPnvserverVpnurlpolicyBindingDataSource,
		tmglobal_auditnslogpolicy_binding.TMglobalAuditnslogpolicyBindingDataSource,
		tmglobal_auditsyslogpolicy_binding.TMglobalAuditsyslogpolicyBindingDataSource,
		tmglobal_tmtrafficpolicy_binding.TMglobalTmtrafficpolicyBindingDataSource,
		transformglobal_transformpolicy_binding.TRansformglobalTransformpolicyBindingDataSource,
		transformpolicylabel_transformpolicy_binding.TRansformpolicylabelTransformpolicyBindingDataSource,
		tunnelglobal_tunneltrafficpolicy_binding.TUnnelglobalTunneltrafficpolicyBindingDataSource,
		vlan_channel_binding.VLanChannelBindingDataSource,
		vlan_interface_binding.VLanInterfaceBindingDataSource,
		vlan_nsip6_binding.VLanNsip6BindingDataSource,
		vlan_nsip_binding.VLanNsipBindingDataSource,
		vxlan_nsip6_binding.VXlanNsip6BindingDataSource,
		vxlan_nsip_binding.VXlanNsipBindingDataSource,
		vxlan_srcip_binding.VXlanSrcipBindingDataSource,
		vxlanvlanmap_vxlan_binding.VXlanvlanmapVxlanBindingDataSource,
		sslcacertgroup_sslcertkey_binding.SSlcacertgroupSslcertkeyBindingDataSource,
		sslcertkey_sslocspresponder_binding.SSlcertkeySslocspresponderBindingDataSource,
		sslcipher_sslciphersuite_binding.SSlcipherSslciphersuiteBindingDataSource,
		sslpolicylabel_sslpolicy_binding.SSlpolicylabelSslpolicyBindingDataSource,
		sslprofile_ecccurve_binding.SSlprofileEcccurveBindingDataSource,
		sslprofile_sslcertkey_binding.SSlprofileSslcertkeyBindingDataSource,
		sslprofile_sslcipher_binding.SSlprofileSslcipherBindingDataSource,
		sslservice_ecccurve_binding.SSlserviceEcccurveBindingDataSource,
		sslservice_sslcertkey_binding.SSlserviceSslcertkeyBindingDataSource,
		sslservice_sslciphersuite_binding.SSlserviceSslciphersuiteBindingDataSource,
		sslservicegroup_ecccurve_binding.SSlservicegroupEcccurveBindingDataSource,
		sslservicegroup_sslcertkey_binding.SSlservicegroupSslcertkeyBindingDataSource,
		sslservicegroup_sslciphersuite_binding.SSlservicegroupSslciphersuiteBindingDataSource,
		sslvserver_ecccurve_binding.SSlvserverEcccurveBindingDataSource,
		sslvserver_sslcertkey_binding.SSlvserverSslcertkeyBindingDataSource,
		sslvserver_sslciphersuite_binding.SSlvserverSslciphersuiteBindingDataSource,
		sslvserver_sslpolicy_binding.SSlvserverSslpolicyBindingDataSource,
		systemglobal_auditnslogpolicy_binding.SYstemglobalAuditnslogpolicyBindingDataSource,
		systemglobal_authenticationldappolicy_binding.SYstemglobalAuthenticationldappolicyBindingDataSource,
		systemglobal_authenticationlocalpolicy_binding.SYstemglobalAuthenticationlocalpolicyBindingDataSource,
		systemglobal_authenticationpolicy_binding.SYstemglobalAuthenticationpolicyBindingDataSource,
		systemglobal_authenticationradiuspolicy_binding.SYstemglobalAuthenticationradiuspolicyBindingDataSource,
		systemglobal_authenticationtacacspolicy_binding.SYstemglobalAuthenticationtacacspolicyBindingDataSource,
		systemgroup_nspartition_binding.SYstemgroupNspartitionBindingDataSource,
		systemgroup_systemcmdpolicy_binding.SYstemgroupSystemcmdpolicyBindingDataSource,
		systemgroup_systemuser_binding.SYstemgroupSystemuserBindingDataSource,
		systemuser_nspartition_binding.SYstemuserNspartitionBindingDataSource,
		systemuser_systemcmdpolicy_binding.SYstemuserSystemcmdpolicyBindingDataSource,
		linkset_channel_binding.LInksetChannelBindingDataSource,
		lsnappsprofile_lsnappsattributes_binding.LSnappsprofileLsnappsattributesBindingDataSource,
		lsnappsprofile_port_binding.LSnappsprofilePortBindingDataSource,
		lsnclient_network6_binding.LSnclientNetwork6BindingDataSource,
		lsnclient_network_binding.LSnclientNetworkBindingDataSource,
		lsnclient_nsacl6_binding.LSnclientNsacl6BindingDataSource,
		lsnclient_nsacl_binding.LSnclientNsaclBindingDataSource,
		lsngroup_lsnappsprofile_binding.LSngroupLsnappsprofileBindingDataSource,
		lsngroup_lsnhttphdrlogprofile_binding.LSngroupLsnhttphdrlogprofileBindingDataSource,
		lsngroup_lsnlogprofile_binding.LSngroupLsnlogprofileBindingDataSource,
		lsngroup_lsnpool_binding.LSngroupLsnpoolBindingDataSource,
		lsngroup_lsntransportprofile_binding.LSngroupLsntransportprofileBindingDataSource,
		lsngroup_pcpserver_binding.LSngroupPcpserverBindingDataSource,
		mapbmr_bmrv4network_binding.MApbmrBmrv4networkBindingDataSource,
		mapdomain_mapbmr_binding.MApdomainMapbmrBindingDataSource,
		nd6ravariables_onlinkipv6prefix_binding.ND6ravariablesOnlinkipv6prefixBindingDataSource,
		netbridge_iptunnel_binding.NEtbridgeIptunnelBindingDataSource,
		netbridge_nsip6_binding.NEtbridgeNsip6BindingDataSource,
		netbridge_nsip_binding.NEtbridgeNsipBindingDataSource,
		netbridge_vlan_binding.NEtbridgeVlanBindingDataSource,
		netprofile_natrule_binding.NEtprofileNatruleBindingDataSource,
		netprofile_srcportset_binding.NEtprofileSrcportsetBindingDataSource,
		nspartition_bridgegroup_binding.NSpartitionBridgegroupBindingDataSource,
		nspartition_vlan_binding.NSpartitionVlanBindingDataSource,
		nspartition_vxlan_binding.NSpartitionVxlanBindingDataSource,
		nsservicepath_nsservicefunction_binding.NSservicepathNsservicefunctionBindingDataSource,
		nstrafficdomain_bridgegroup_binding.NStrafficdomainBridgegroupBindingDataSource,
		nstrafficdomain_vlan_binding.NStrafficdomainVlanBindingDataSource,
		nstrafficdomain_vxlan_binding.NStrafficdomainVxlanBindingDataSource,
		policydataset_value_binding.POlicydatasetValueBindingDataSource,
		policypatset_pattern_binding.POlicypatsetPatternBindingDataSource,
		policystringmap_pattern_binding.POlicystringmapPatternBindingDataSource,
		responderglobal_responderpolicy_binding.REsponderglobalResponderpolicyBindingDataSource,
		responderpolicylabel_responderpolicy_binding.REsponderpolicylabelResponderpolicyBindingDataSource,
		rewriteglobal_rewritepolicy_binding.REwriteglobalRewritepolicyBindingDataSource,
		rewritepolicylabel_rewritepolicy_binding.REwritepolicylabelRewritepolicyBindingDataSource,
		rnat6_nsip6_binding.RNat6Nsip6BindingDataSource,
		rnat_nsip_binding.RNatNsipBindingDataSource,
		rnatglobal_auditsyslogpolicy_binding.RNatglobalAuditsyslogpolicyBindingDataSource,
		service_lbmonitor_binding.SErviceLbmonitorBindingDataSource,
		servicegroup_lbmonitor_binding.SErvicegroupLbmonitorBindingDataSource,
		servicegroup_servicegroupmember_binding.SErvicegroupServicegroupmemberBindingDataSource,
		snmptrap_snmpuser_binding.SNmptrapSnmpuserBindingDataSource,
		vpnvserver_appfwpolicy_binding.VpnvserverAppfwpolicyBindingDataSource,
		csvserver_lbvserver_binding.CSvserverLbvserverBindingDataSource,
		lbparameter.LBparameterDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CitrixAdcFrameworkProvider{
			version: version,
		}
	}
}
