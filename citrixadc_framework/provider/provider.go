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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/analyticsprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appalgparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appflowaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appflowcollector"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appflowparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appflowpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appflowpolicylabel"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwconfidfield"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwfieldtype"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwhtmlerrorpage"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwjsoncontenttype"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwjsonerrorpage"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwlearningsettings"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwmultipartformcontenttype"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwpolicylabel"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/appfwprofile"
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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/auditnslogparams"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/auditnslogpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/auditsyslogaction"
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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationwebauthaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authenticationwebauthpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authorizationpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/authorizationpolicylabel"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/autoscaleaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/autoscalepolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/autoscaleprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/botpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/botpolicylabel"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/botprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/botsettings"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/botsignature"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/bridgegroup"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/bridgetable"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/cachecontentgroup"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/cacheforwardproxy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/cacheparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/cachepolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/cachepolicylabel"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/cacheselector"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/channel"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/clusterinstance"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/clusternode"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/clusternodegroup"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/cmpaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/cmpparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/cmppolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/cmppolicylabel"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/contentinspectionaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/contentinspectioncallout"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/contentinspectionparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/contentinspectionpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/contentinspectionpolicylabel"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/contentinspectionprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/crpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/crvserver"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/csaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/csparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/cspolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/cspolicylabel"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/csvserver"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dbdbprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dbuser"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnsaaaarec"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnsaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnsaction64"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnsaddrec"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnscnamerec"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnskey"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnsmxrec"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnsnameserver"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnsnaptrrec"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnsnsrec"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnsparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnspolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnspolicy64"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/dnspolicylabel"
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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/feoparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/feopolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/fis"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/forwardingsession"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/gslbparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/gslbservice"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/gslbservicegroup"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/gslbsite"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/gslbvserver"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/hanode"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/icaaccessprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/icaaction"
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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/iptunnel"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/iptunnelparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/ipv6"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/l2param"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/l3param"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/l4param"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lacp"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbgroup"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbmetrictable"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbmonitor"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbroute"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbroute6"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbsipparameters"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lbvserver"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/linkset"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lldpparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/location"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/locationfile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/locationfile6"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/locationparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lsnappsattributes"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lsnappsprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lsnclient"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/lsngroup"
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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/mapdmr"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/mapdomain"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nat64"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nat64param"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nd6"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nd6ravariables"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/netbridge"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/netprofile"
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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nslicense"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nslicenseparameters"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nslicenseproxyserver"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nslicenseserver"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nslimitidentifier"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsmode"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nspartition"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nspbr"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nspbr6"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsratecontrol"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsrpcnode"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsservicefunction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsservicepath"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nssimpleacl"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nssimpleacl6"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nsspparams"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nstcpbufparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nstcpparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nstcpprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nstimeout"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nstimer"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/nstrafficdomain"
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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/policyexpression"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/policyhttpcallout"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/policymap"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/policyparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/policypatset"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/policystringmap"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/ptp"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/quicbridgeprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/radiusnode"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/rdpclientprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/rdpserverprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/reputationsettings"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/responderaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/responderhtmlpage"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/responderparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/responderpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/responderpolicylabel"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/rewriteaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/rewriteparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/rewritepolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/rewritepolicylabel"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/rnat"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/rnat6"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/rnatparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/route"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/route6"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/routerdynamicrouting"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/rsskeytype"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/server"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/servicegroup"
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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/snmpuser"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/snmpview"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/spilloveraction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/spilloverpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslcacertgroup"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslcertfile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslcertkey"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslcipher"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslcrl"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/ssldtlsprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslfipskey"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslhsmkey"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/ssllogprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslocspresponder"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslpolicylabel"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslservice"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslservicegroup"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/sslvserver"
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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/systemgroup"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/systemparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/systemuser"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/tmformssoaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/tmsamlssoprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/tmsessionaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/tmsessionparameter"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/tmsessionpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/tmtrafficaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/tmtrafficpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/transformaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/transformpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/transformpolicylabel"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/transformprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/tunneltrafficpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/userprotocol"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/uservserver"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/videooptimizationdetectionaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/videooptimizationdetectionpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/videooptimizationpacingaction"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/videooptimizationpacingpolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vlan"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnalwaysonprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnclientlessaccesspolicy"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnclientlessaccessprofile"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpneula"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vpnformssoaction"
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
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vrid"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vrid6"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vridparam"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vxlan"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/vxlanvlanmap"
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
				Description: "Target NS ip. When defined username, password and endpoint must refer to MAS.",
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

	userHeaders := map[string]string{
		"User-Agent": "terraform-ctxadc-framework",
	}

	params := adcnitrogoservice.NitroParams{
		Url:       endpoint,
		Username:  username,
		Password:  password,
		ProxiedNs: proxiedNs,
		SslVerify: !insecureSkipVerify,
		Headers:   userHeaders,
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

	// providerData := &ProviderData{
	// 	Client:   client,
	// 	Username: username,
	// 	Password: password,
	// 	Endpoint: endpoint,
	// }

	resp.DataSourceData = &client
	resp.ResourceData = &client

	tflog.Info(ctx, "Configured CitrixADC Framework Provider", map[string]any{"success": true})
}

func (p *CitrixAdcFrameworkProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		lbparameter.NewLbParameterResource,
		sslcertkey.NewSslCertKeyResource,
	}
}

func (p *CitrixAdcFrameworkProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		lbparameter.LBParameterDataSource,
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
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CitrixAdcFrameworkProvider{
			version: version,
		}
	}
}
