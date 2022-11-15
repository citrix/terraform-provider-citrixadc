## 1.26.0 (Nov 15, 2022)

FEATURES

* **New Resource** change_password
* **New Resource** rnat_clear

BUG FIXES

* **citrixadc_lbmonitor** Added `ipaddress` attribute in `lbmonitor` resource
* **citrixadc_nd6ravariables** Type conversion for `vlan` attribute in `nd6ravariables` resource.
* **citrixadc_service**: Removed the default values for some attributes in `service` resource.
* **citrixadc_systemuser**: Removed unwanted function call to `client.SetPassword()` in read operation.

ENHANCEMENTS

* **citrixadc_rnat**: Supported Add, update and Delete operation in `rnat` resource(for ADC version 13.0 and above) and created new `rnat_clear` resource which supports old rnat resouce or operation
`clear` for ADC version 12.0. Refer `rnat` and `rnat_clear` resource documentation for more details.
* **citrixadc_cspolicy**: Updated `cspolicy` resource to support to the latest ADC versions i.e., Added attribute `boundto` which is present in `csvserver_cspolicy_binding` for ADC version 13.1.
## 1.25.0 (Nov 4, 2022)

FEATURES

* **New Resource** sslcertkey_update
* **New Resource** nspbrs
* **New Resource** sslcertkey_sslocspresponder_binding
* **New Resource** sslcertfile
* **New Data-Source** hanode

EXAMPLE USECASES

* **New Usecase** SSL Offloading in Citrix ADC using terraform modules

ENHANCEMENTS

* **Enhancements** policydataset_value_binding supported CIDR for IP addresses
* **Enhancements** Updated documentation

BUG FIXES

* **Bug Fix** a small fix in naparam
* **Bug Fix** Updated systemuser resource


## 1.24.0 (Sept 28, 2022)

FEATURES

* **New Resource** aaaglobal_aaapreauthenticationpolicy_binding
* **New Resource** lsnparameter
* **New Resource** streamselector
* **New Resource** lsngroup_lsntransportprofile_binding
* **New Resource** lsngroup_pcpserver_binding
* **New Resource** lsnhttphdrlogprofile
* **New Resource** subscriberparam
* **New Resource** subscriberradiusinterface
* **New Resource** subscriberprofile
* **New Resource** tunneltrafficpolicy
* **New Resource** cacheselector
* **New Resource** cacheglobal_cachepolicy_binding
* **New Resource** sslservicegroup_sslciphersuite_binding
* **New Resource** cachepolicylabel
* **New Resource** cachepolicylabel_cachepolicy_binding
* **New Resource** cachepolicy
* **New Resource** smppparam
* **New Resource** icaparameter
* **New Resource** tunnelglobal_tunneltrafficpolicy_binding
* **New Resource** auditnslogaction
* **New Resource** auditnslogpolicy
* **New Resource** cacheforwardproxy
* **New Resource** linkset_channel_binding
* **New Resource** smppuser
* **New Resource** cachecontentgroup
* **New Resource** subscribergxinterface
* **New Resource** locationfile6
* **New Resource** spilloveraction
* **New Resource** autoscalepolicy
* **New Resource** autoscaleprofile
* **New Resource** autoscaleaction
* **New Resource** nsweblogparam
* **New Usecase** Added ADC usecases for Responded and Rewrite Resources

ENHANCEMENTS

* **Enhancements** Updated documentation
* **Enhancements** Removed examples from example folder which are available in terraform registry

BUG FIXES
* **Bugfix** fixed `inc` attribute not adding to the `hanode` POST payload

## 1.23.0 (Sept 15, 2022)

FEATURES

* **New Resource** rdpserverprofile
* **New Resource** rdpclientprofile
* **New Resource** pcpprofile
* **New Resource** pcpserver
* **New Resource** contentinspectionaction
* **New Resource** hanode_routemonitor6_binding
* **New Resource** hanode_routemonitor_binding
* **New Resource** streamidentifier
* **New Resource** lsngroup_lsnpool_binding
* **New Resource** lsngroup_lsnlogprofile_binding
* **New Resource** lsngroup_lsnhttphdrlogprofile_binding
* **New Resource** lsngroup_lsnappsprofile_binding
* **New Resource** lsnclient_nsacl6_binding
* **New Resource** lsnclient_nsacl_binding
* **New Resource** lsnclient_network6_binding
* **New Resource** lsnclient_network_binding
* **New Resource** lsnappsprofile_port_binding
* **New Resource** lsnappsprofile_lsnappsattributes_binding
* **New Resource** lsnstatic
* **New Resource** lsnip6profile
* **New Resource** lsntransportprofile
* **New Resource** cacheparameter
* **New Resource** ntpsync
* **New Resource** lsnsipalgprofile
* **New Resource** lsnrtspalgprofile
* **New Resource** lsnpool
* **New Resource** lsnlogprofile
* **New Resource** lsnappsprofile
* **New Resource** lsngroup
* **New Resource** lsnappsattributes
* **New Resource** lsnclient
* **New Resource** contentinspectionpolicylabel_contentinspectionpolicy_binding
* **New Resource** contentinspectionglobal_contentinspectionpolicy_binding
* **New Resource** contentinspectionprofile
* **New Resource** contentinspectionpolicylabel
* **New Resource** contentinspectioncallout
* **New Resource** contentinspectionparameter
* **New Resource** contentinspectionpolicy
* **New Resource** appqoeaction
* **New Resource** appqoeparameter
* **New Resource** customresp
* **New Resource** appqoepolicy
* **New Resource** feoparameter
* **New Resource** locationfile
* **New Resource** feoaction
* **New Resource** systembackup
* **New Resource** auditsyslogglobal_auditsyslogpolicy_binding
* **New Resource** tmtrafficaction
* **New Resource** admparameter
* **New Resource** hanode
* **New Resource** feoaction
* **New Resource** auditnslogglobal_auditnslogpolicy_binding
* **New Resource** dbdbprofile
* **New Resource** ipsecalgprofile
* **New Resource** feopolicy
* **New Resource** feologbal_feopolicy_binding
* **New Resource** lldpparam
* **New Resource** analyticsprofile
* **New Resource** auditnslogparams
* **New Resource** auditsyslogparams
* **New Resource** dbuser
* **New Resource** uservserver
* **New Resource** userprotocol
* **New Resource** channel
* **New Resource** interfacepair
* **New Resource** ipsecprofile
* **New Resource** tmsamlssoprofile
* **New Resource** tmglobal_tmtrafficpolicy_binding
* **New Resource** ipsecparameter
* **New Resource** tmtrafficpolicy
* **New Resource** tmsessionaction
* **New Resource** tmsessionpolicy
* **New Resource** tmformssoaction

ENHANCEMENTS

* **Enhancements** Updated policydataset_value_binding Resource
* **Enhancements** updated appfwprofile attributes
* **Enhancements** Fixed nsip documentation
* **Enhancements** Added enable/disable functionality


## 1.22.0 (Aug 25, 2022)

BUG FIXES

* Fixed vendor module dependencies

## 1.21.0 (Aug 25, 2022)

FEATURES

* **New Resource** tmsessionparameter
* **New Resource** tmglobal_auditnslogpolicy_binding
* **New Resource** tmglobal_auditsyslogpolicy_binding
* **New Resource** icaglobal_icapolicy_binding
* **New Resource** icapolicy
* **New Resource** ntpparam
* **New Resource** icalatencyprofile
* **New Resource** icaaccessprofile
* **New Resource** icaaction
* **New Resource** cmppolicylabel_cmppolicy_binding
* **New Resource** cmpparameter
* **New Resource** cmpaction
* **New Resource** cmppolicylabel
* **New Resource** systembackup
* **New Resource** systemcollectionparam
* **New Resource** aaauser_intranetip6_binding
* **New Resource** systemglobal_authenticationtacacspolicy_binding
* **New Resource** systemglobal_auditnslogpolicy_binding
* **New Resource** systemglobal_authenticationlocalpolicy_binding
* **New Resource** aaaradiusparams
* **New Resource** systemgroup_nspartition_binding
* **New Resource** systemuser_nspartition_binding
* **New Resource** systemglobal_authenticationpolicy_binding
* **New Resource** systemglobal_authenticationradiuspolicy_binding
* **New Resource** aaauser_vpnurl_binding
* **New Resource** aaagroup_auditnslogpolicy_binding
* **New Resource** aaagroup_aaauser_binding
* **New Resource** aaauser_intranetip_binding
* **New Resource** aaagroup_intranetip_binding
* **New Resource** aaagroup_vpnintranetapplication_binding
* **New Resource** aaauser_vpnurlpolicy_binding
* **New Resource** aaagroup_auditsyslogpolicy_binding
* **New Resource** aaagroup_vpnsessionpolicy_binding
* **New Resource** aaauser_tmsessionpolicy_binding
* **New Resource** aaauser_auditsyslogpolicy_binding
* **New Resource** aaauser_vpntrafficpolicy_binding
* **New Resource** aaagroup_tmsessionpolicy_binding
* **New Resource** aaauser_vpnsessionpolicy_binding
* **New Resource** aaagroup_vpnurlpolicy_binding
* **New Resource** aaagroup_vpntrafficpolicy_binding
* **New Resource** aaauser_authorizationpolicy_binding
* **New Resource** aaagroup_vpnurl_binding
* **New Resource** aaauser_vpnintranetapplication_binding
* **New Resource** aaauser_auditnslogpolicy_binding
* **New Resource** aaagroup_authorizationpolicy_binding
* **New Resource** aaakcdaccount
* **New Resource** aaatacacsparams
* **New Resource** aaagroup
* **New Resource** aaapreauthenticationparameter
* **New Resource** aaapreauthenticationpolicy
* **New Resource** aaaparameter
* **New Resource** aaauser
* **New Resource** aaaotpparameter
* **New Resource** aaassoprofile
* **New Resource** vlan_nsip6_binding
* **New Resource** aaacertparams
* **New Resource** aaaldapparams
* **New Resource** vlan_channel_binding
* **New Resource** onlinkipv6prefix
* **New Resource** rnat6_nsip6_binding
* **New Resource** rnat_nsip_binding
* **New Resource** netbridge_nsip6_binding
* **New Resource** netbridge_nsip_binding
* **New Resource** nd6ravariables_onlinkipv6prefix_binding
* **New Resource** ipset_nsip6_binding
* **New Resource** ipset_nsip_binding
* **New Resource** ipv6
* **New Resource** inatparam
* **New Resource** rnat6
* **New Resource** l2param
* **New Resource** nd6ravariables
* **New Resource** arp
* **New Resource** l3param
* **New Resource** nspbr
* **New Resource** nd6
* **New Resource** route6
* **New Resource** lacp
* **New Resource** snmpcommunity
* **New Resource** nslicenseproxyserver
* **New Resource** nsencryptionparams
* **New Resource** nshostname
* **New Resource** nscqaparam
* **New Resource** responderglobal_responderpolicy_binding
* **New Resource** clusternodegroup_authenticationvserver_binding
* **New Resource** clusternode_routemonitor_binding
* **New Resource** clusternodegroup_streamidentifier_binding
* **New Resource** clusternodegroup_vpnvserver_binding
* **New Resource** clusternodegroup_nslimitidentifier_binding
* **New Resource** clusternode
* **New Resource** clusternodegroup_service_binding
* **New Resource** clusterinstance
* **New Resource** clusternodegroup_gslbvserver_binding
* **New Resource** clusternodegroup_crvserver_binding
* **New Resource** clusternodegroup_csvserver_binding
* **New Resource** clusternodegroup_clusternode_binding
* **New Resource** clusternodegroup_lbvserver_binding

ENHANCEMENTS

* **Documentation Update**
* Update issue templates for issues and PRs

BUG FIXES

* **Bug Fix** Removed duplicate entry
* Fixed nsip resource with Required attributes

## 1.20.0 (August 05, 2022)

FEATURES

* **New Resource** `clusternodegroup_gslbsite_binding`
* **New Resource** `cmpglobal_cmppolicy_binding`
* **New Resource** `clusternodegroup`
* **New Resource** `nstimeout`
* **New Resource** `systemparameter`
* **New Resource** `ntpserver`
* **New Resource** `snmpalarm`
* **New Resource** `snmpmanager`
* **New Resource** `snmpmib`
* **New Resource** `snmpoption`
* **New Resource** `snmpengineid`
* **New Resource** `snmpuser`
* **New Resource** `snmpgroup`
* **New Resource** `snmpview`
* **New Resource** `snmptrap`
* **New Resource** `appflowglobal_appflowpolicy_binding`
* **New Resource** `appflowpolicylabel_appflowpolicy_binding`
* **New Resource** `appflowaction_analyticsprofile_binding`
* **New Resource** `appflowpolicy`
* **New Resource** `appflowaction`
* **New Resource** `crvserver_policymap_binding`
* **New Resource** `appflowpolicylabel`
* **New Resource** `appflowparam`
* **New Resource** `appflowcollector`

## 1.19.0 (July 29, 2022)

FEATURES

* **New Resource** `dnsnameserver`
* **New Resource** `crvserver_spilloverpolicy_binding`
* **New Resource** `crvserver_cspolicy_binding`
* **New Resource** `crvserver_filterpolicy_binding`
* **New Resource** `crvserver_appqoepolicy_bind`
* **New Resource** `crvserver_analyticsprofile_binding`
* **New Resource** `crvserver_icapolicy_binding`
* **New Resource** `crvserver_cachepolicy_binding`
* **New Resource** `crvserver_cmppolicy_binding`
* **New Resource** `crvserver_feopolicy_binding`

## 1.18.0 (July 26, 2022)

FEATURES

* **New Resource** `dnssuffix `
* **New Resource** `dnspolicylabel`
* **New Resource** `dnsptrrec`
* **New Resource** `ip6tunnelparam`
* **New Resource** `dnstxtrec`
* **New Resource** `dnscnamerec`
* **New Resource** `dnsnaptrrec`
* **New Resource** `dnssrvrec`
* **New Resource** `dnsaction`
* **New Resource** `dnsaction64`
* **New Resource** `dnszone`
* **New Resource** `dnspolicy`
* **New Resource** `dnsaddrec`
* **New Resource** `dnspolicy64`
* **New Resource** `dnsmxrec`
* **New Resource** `authorizationpolicylabel`
* **New Resource** `transformpolicylabel`
* **New Resource** `gslbservicegroup`
* **New Resource** `authorizationpolicylabel_authorizationpolicy_binding`
* **New Resource** `transformpolicylabel_policy_binding`
* **New Resource** `transformglobal_transformpolicy_binding`
* **New Resource** `transformpolicylabel_policy_binding`
* **New Resource** `gslbparameter`
* **New Resource** `gslbservice_lbmonitor_binding`
* **New Resource** `gslbvserver_domain_binding`
* **New Resource** `gslbvserver_gslbservicegroup_binding`
* **New Resource** `gslbservicegroup_lbmonitor_binding`
* **New Resource** `gslbvserver_spilloverpolicy_binding`
* **New Resource** `gslbservice_dnsview_binding`
* **New Resource** `gslbvserver_gslbservice_binding`
* **New Resource** `crpolicy`
* **New Resource** `crvserver`
* **New Resource** `crvserver_appfwpolicy_binding`
* **New Resource** `dnskey`
* **New Resource** `dnspolicylabel_dnspolicy_binding`
* **New Resource** `crvserver_crpolicy_binding`
* **New Resource** `crvserver_responderpolicy_binding`
* **New Resource** `crvserver_rewritepolicy_binding`
* **New Resource** `crvserver_appflowpolicy_binding`

ENHANCEMENTS

* Rearranged examples folders in github repository
* Added example with secure CS vserver
* Update HA pair upgrade script

BUG FIXES

* Fix glsbservicegroup to retrieve servicegroupname correctly.

## 1.17.0 (June 30, 2022)

FEATURES

* **New Resource** `appalgparam`
* **New Resource** `iptunnelparam`
* **New Resource** `nsacl6`
* **New Resource** `nspbr6`
* **New Resource** `nstcpbufparam`
* **New Resource** `ip6tunnel`
* **New Resource** `rsskeytype`
* **New Resource** `nat64`
* **New Resource** `netbridge_iptunnel_binding`
* **New Resource** `dnsprofile`
* **New Resource** `mapdomain_mapbmr_binding`
* **New Resource** `netbridge_vlan_binding`
* **New Resource** `bridgetable`
* **New Resource** `dnsview`

* **New Data Source** `nitro_info`

BUG FIXES

* Fixed `lbvserver` resource to accept zero value for `timeout` attribute.


## 1.16.0 (June 1, 2022)

FEATURES

* **New Resource** `nsservicepath`
* **New Resource** `nspartition`
* **New Resource** `nsvariable`
* **New Resource** `nsappflowcollector`
* **New Resource** `nsicapprofile`
* **New Resource** `nsxmlnamespace`
* **New Resource** `nstrafficdomain`
* **New Resource** `nsencryptionkey`
* **New Resource** `nsservicefunction`
* **New Resource** `nssimpleacl`
* **New Resource** `mapbmr`
* **New Resource** `nssimpleacl6`
* **New Resource** `mapdmr`
* **New Resource** `appfwwsdl`
* **New Resource** `nsspparams`
* **New Resource** `nsconsoleloginprompt`
* **New Resource** `extendedmemoryparam`
* **New Resource** `appfwmultipartformcontenttype`
* **New Resource** `locationparameter`
* **New Resource** `nstrafficdomain_vlan_binding`
* **New Resource** `nsservicepath_nsservicefunction_binding`
* **New Resource** `nsdiameter`
* **New Resource** `nspartition_vxlan_binding`
* **New Resource** `nspartition_vlan_binding`
* **New Resource** `nsdhcpparams`
* **New Resource** `nstrafficdomain_vxlan_binding`
* **New Resource** `nsassignment`
* **New Resource** `appfwxmlschema`
* **New Resource** `nsratecontrol`
* **New Resource** `l4param`
* **New Resource** `arpparam`
* **New Resource** `rnatparam`
* **New Resource** `ptp`
* **New Resource** `nshttpparam`
* **New Resource** `mapdomain`
* **New Resource** `mapbmr_bmrv4network_binding`
* **New Resource** `vridparam`
* **New Resource** `bridgegroup`
* **New Resource** `bridgegroup_vlan_binding`
* **New Resource** `nspartition_bridgegroup_binding`
* **New Resource** `nstrafficdomain_bridgegroup_binding`
* **New Resource** `nat64param`
* **New Resource** `nslicenseparameters`
* **New Resource** `bridgegroup_nsip_binding`
* **New Resource** `bridgegroup_nsip6_binding`

* **New Data Source** `sslcipher_sslvserver_bindings`

BUG FIXES

* Fixed resource missing errorcode in `sslvserver_sslciphersuite_binding`


## 1.15.0 (May 16, 2022)

FEATURES

* **New Resource** `vxlan_srcip_binding`
* **New Resource** `botprofile_ipreputation_binding`
* **New Resource** `botprofile_blacklist_binding`
* **New Resource** `botprofile_ratelimit_binding`
* **New Resource** `botglobal_botpolicy_binding`
* **New Resource** `appfwprofile_contenttype_binding`
* **New Resource** `appfwprofile_excluderescontenttype_binding`
* **New Resource** `appfwglobal_appfwpolicy_binding`
* **New Resource** `appfwprofile_csrftag_binding`
* **New Resource** `appfwglobal_auditsyslogpolicy_binding`
* **New Resource** `appfwglobal_auditnslogpolicy_binding`
* **New Resource** `appfwprofile_creditcardnumber_binding`
* **New Resource** `appfwpolicylabel_appfwpolicy_binding`
* **New Resource** `appfwprofile_cmdinjection_binding`
* **New Resource** `appfwprofile_xmlxss_binding`
* **New Resource** `appfwprofile_fieldformat_binding`
* **New Resource** `appfwprofile_safeobject_binding`
* **New Resource** `appfwprofile_xmlwsiurl_binding`
* **New Resource** `appfwprofile_jsoncmdurl_binding`
* **New Resource** `botsignature`
* **New Resource** `appfwurlencodedformcontenttype`
* **New Resource** `appfwprofile_jsonsqlurl_binding`
* **New Resource** `appfwprofile_xmlsqlinjection_binding`
* **New Resource** `appfwprofile_fieldconsistency_binding`
* **New Resource** `appfwprofile_jsonxssurl_binding`
* **New Resource** `appfwprofile_logexpression_binding`
* **New Resource** `appfwprofile_fileuploadtype_binding`
* **New Resource** `appfwprofile_jsondosurl_binding`
* **New Resource** `appfwprofile_xmlattachmenturl_binding`
* **New Resource** `appfwprofile_trustedlearningclients_binding`
* **New Resource** `appfwprofile_xmldosurl_binding`
* **New Resource** `nitro_resource`
* **New Resource** `appfwprofile_xmlvalidationurl_binding`
* **New Resource** `appfwsignatures`
* **New Resource** `appfwlearningsettings`
* **New Resource** `appfwhtmlerrorpage`
* **New Resource** `appfwxmlerrorpage`
* **New Resource** `location`
* **New Resource** `service_dospolicy_binding`
* **New Resource** `fis`
* **New Resource** `vrid`
* **New Resource** `netprofile_natrule_binding`
* **New Resource** `radiusnode`
* **New Resource** `service_lbmonitor_binding`
* **New Resource** `vrid6`
* **New Resource** `forwardingsession`
* **New Resource** `netbridge`
* **New Resource** `netprofile_srcportset_binding`
* **New Resource** `rnatglobal_auditsyslogpolicy_binding`
* **New Resource** `appfwjsonerrorpage`
* **New Resource** `nstimer`
* **New Resource** `nslimitidentifier`
* **New Resource** `nshmackey`


## 1.14.0 (April 20, 2022)

FEATURES

* **New Resource** `authenticationvserver_auditnslogpolicy_binding`
* **New Resource** `vpnglobal_domain_binding`
* **New Resource** `vpnglobal_intranetip_binding`
* **New Resource** `vpnglobal_sharefileserver_binding`
* **New Resource** `vpnglobal_vpnclientlessaccesspolicy_binding`
* **New Resource** `vpnvserver_intranetip_binding`
* **New Resource** `vpnvserver_sharefileserver_binding`
* **New Resource** `authenticationnegotiatepolicy`
* **New Resource** `vpnvserver_aaapreauthenticationpolicy_binding`
* **New Resource** `vpnvserver_icapolicy_binding`
* **New Resource** `authenticationvserver_authenticationloginschemapolicy_binding`
* **New Resource** `vpnvserver_authenticationnegotiatepolicy_binding`
* **New Resource** `authenticationvserver_authenticationnegotiatepolicy_binding`
* **New Resource** `vpnvserver_authenticationoauthidppolicy_binding`
* **New Resource** `authenticationsamlidpprofile`
* **New Resource** `authenticationvserver_tmsessionpolicy_binding`
* **New Resource** `authenticationvserver_authenticationoauthidppolicy_binding`
* **New Resource** `vpnvserver_vpnurlpolicy_binding`
* **New Resource** `systemglobal_authenticationldappolicy_binding`
* **New Resource** `vpnvserver_feopolicy_binding`
* **New Resource** `authenticationsamlidppolicy`
* **New Resource** `vpnvserver_authenticationloginschemapolicy_binding`
* **New Resource** `vpnvserver_vpnclientlessaccesspolicy_binding`
* **New Resource** `vpnparameter`
* **New Resource** `vpnvserver_staserver_binding`
* **New Resource** `vpnglobal_staserver_binding`
* **New Resource** `vpnglobal_authenticationnegotiatepolicy_binding`
* **New Resource** `vpnglobal_intranetip6_binding`
* **New Resource** `vpnvserver_intranetip6_binding`
* **New Resource** `vpnvserver_authenticationsamlidppolicy_binding`
* **New Resource** `authenticationvserver_authenticationsamlidppolicy_binding`
* **New Resource** `authenticationvserver_cachepolicy_binding`
* **New Resource** `vxlan`
* **New Resource** `vxlanvlanmap`
* **New Resource** `appfwconfidfield`
* **New Resource** `vxlan_nsip_binding`
* **New Resource** `vxlan_nsip6_binding`
* **New Resource** `vxlanvlanmap_vxlan_binding`
* **New Resource** `botpolicylabel_botpolicy_binding`
* **New Resource** `botprofile_captcha_binding`
* **New Resource** `botprofile_tps_binding`
* **New Resource** `botprofile_trapinsertionurl_binding`
* **New Resource** `botprofile_logexpression_binding`
* **New Resource** `botprofile_whitelist_binding`
* **New Resource** `filteraction`

BUG FIXES

* Fixed erroneous category `vpnnexthopserever` documentation
* Fixed erroneous category `vpnvserver_authenticationcertpolicy_binding` documentation

## 1.13.0 (March 31, 2022)

FEATURES

* **New Resource** `authenticationldappolicy`
* **New Resource** `authenticationlocalpolicy`
* **New Resource** `authenticationnoauthaction`
* **New Resource** `authenticationsamlaction`
* **New Resource** `vpnvserver_cachepolicy_binding`
* **New Resource** `vpnvserver_appcontroller_binding`
* **New Resource** `authenticationepaaction`
* **New Resource** `authenticationloginschema`
* **New Resource** `authenticationdfaaction`
* **New Resource** `authenticationsamlpolicy`
* **New Resource** `authenticationradiusaction`
* **New Resource** `authenticationdfapolicy`
* **New Resource** `authenticationemailaction`
* **New Resource** `authenticationwebauthaction`
* **New Resource** `authenticationcaptchaaction`
* **New Resource** `authenticationradiuspolicy`
* **New Resource** `authenticationtacacsaction`
* **New Resource** `authenticationstorefrontauthaction`
* **New Resource** `authenticationwebauthpolicy`
* **New Resource** `authenticationcitrixauthaction`
* **New Resource** `authenticationtacacspolicy`
* **New Resource** `authenticationoauthidpprofile`
* **New Resource** `authenticationcertaction`
* **New Resource** `authenticationpushservice`
* **New Resource** `authenticationoauthidppolicy`
* **New Resource** `authenticationcertpolicy`
* **New Resource** `authenticationnegotiateaction`
* **New Resource** `authenticationloginschemapolicy`
* **New Resource** `vpnvserver_authenticationdfapolicy_binding`
* **New Resource** `vpnglobal_vpnnexthopserver_binding`
* **New Resource** `vpnglobal_vpnportaltheme_binding`
* **New Resource** `vpnglobal_authenticationcertpolicy_binding`
* **New Resource** `vpnglobal_authenticationradiuspolicy_binding`
* **New Resource** `vpnglobal_authenticationsamlpolicy_binding`
* **New Resource** `vpnglobal_authenticationtacacspolicy_binding`
* **New Resource** `vpnvserver_authenticationcertpolicy_binding`
* **New Resource** `vpnvserver_authenticationlocalpolicy_binding`
* **New Resource** `vpnvserver_authenticationsamlpolicy_binding`
* **New Resource** `vpnvserver_authenticationtacacspolicy_binding`
* **New Resource** `vpnvserver_authenticationwebauthpolicy_binding`
* **New Resource** `authenticationvserver_vpnportaltheme_binding`
* **New Resource** `authenticationvserver_authenticationwebauthpolicy_binding`
* **New Resource** `authenticationvserver_authenticationcertpolicy_binding`
* **New Resource** `authenticationvserver_rewritepolicy_binding`
* **New Resource** `authenticationvserver_authenticationldappolicy_binding`
* **New Resource** `authenticationvserver_auditsyslogpolicy_binding`
* **New Resource** `authenticationvserver_authenticationsamlpolicy_binding`
* **New Resource** `authenticationvserver_authenticationtacacspolicy_binding`
* **New Resource** `authenticationvserver_authenticationlocalpolicy_binding`
* **New Resource** `authenticationvserver_cspolicy_binding`
* **New Resource** `authenticationvserver_responderpolicy_binding`
* **New Resource** `authenticationvserver_authenticationradiuspolicy_binding`

## 1.12.0 (March 15, 2022)

FEATURES

* **New Data Source** `nsversion`
* **New Resource** `vpnvserver_vpnurl_binding`
* **New Resource** `vpnvserver_rewritepolicy_binding`
* **New Resource** `vpnvserver_responderpolicy_binding`
* **New Resource** `vpnvserver_cspolicy_binding`
* **New Resource** `authenticationoauthaction`

## 1.11.0 (March 2, 2022)

FEATURES

* **New Resource** `vpnglobal_authenticationldappolicy_binding`
* **New Resource** `vpnglobal_authenticationlocalpolicy_binding`
* **New Resource** `vpnintranetapplication`
* **New Resource** `vpnpcoipvserverprofile`
* **New Resource** `vpnpcoipprofile`
* **New Resource** `vpnglobal_vpnurl_binding`
* **New Resource** `vpnglobal_vpnurlpolicy_binding`
* **New Resource** `vpnglobal_vpneula_binding`
* **New Resource** `vpnvserver_authenticationldappolicy_binding`
* **New Resource** `vpnvserver_authenticationradiuspolicy_binding`
* **New Resource** `vpnnexthopserver`
* **New Resource** `vpnportaltheme`
* **New Resource** `vpnsamlssoprofile`
* **New Resource** `vpnglobal_auditsyslogpolicy_binding`
* **New Resource** `vpnglobal_vpnintranetapplication_binding`
* **New Resource** `vpnvserver_auditsyslogpolicy_binding`
* **New Resource** `vpnvserver_auditnslogpolicy_binding`
* **New Resource** `vpnvserver_appflowpolicy_binding`
* **New Resource** `vpnvserver_analyticsprofile_binding`
* **New Resource** `vpnvserver_vpneula_binding`
* **New Resource** `vpnvserver_vpnintranetapplication_binding`
* **New Resource** `vpnvserver_vpnnexthopserver_binding`
* **New Resource** `vpnvserver_vpnportaltheme_binding`
* **New Resource** `vpnvserver_vpntrafficpolicy_binding`
* **New Resource** `hafailover`


BUG FIXES

* Fix resource `service` to calculate service state correctly.

## 1.10.0 (February 22, 2022)

FEATURES

* **New Resource** `vpnurl`
* **New Resource** `vpnsessionaction`
* **New Resource** `vpnvserver`
* **New Resource** `vpnsessionpolicy`
* **New Resource** `vpntrafficaction`
* **New Resource** `vpnurlaction`
* **New Resource** `vpnvserver_vpnsessionpolicy_binding`
* **New Resource** `vpntrafficpolicy`
* **New Resource** `vpnurlpolicy`
* **New Resource** `authenticationvserver`
* **New Resource** `authenticationldapaction`
* **New Resource** `vpnglobal_sslcertkey_binding`
* **New Resource** `vpnglobal_vpntrafficpolicy_binding`
* **New Resource** `vpnglobal_vpnsessionpolicy_binding`
* **New Resource** `authenticationauthnprofile`
* **New Resource** `authenticationpolicylabel`
* **New Resource** `authenticationpolicy`
* **New Resource** `authenticationvserver_authenticationpolicy_binding`
* **New Resource** `authenticationpolicylabel_authenticationpolicy_binding`

ENHANCEMENTS

* Add apigateway feature in `nsfeature` resource.

BUG FIXES

* Fix resource `lbvserver` to calculate server state correctly.

## 1.9.0 (January 11, 2022)

FEATURES

* **New Resource** `authorizationpolicy`

## 1.8.0 (November 11, 2021)

FEATURES

* **New Resource** `filterglobal_filterpolicy_binding`
* **New Resource** `dnsparameter`
* **New Resource** `appfwsettings`
* **New Resource** `responderhtmlpage`

ENHANCEMENTS

* Add import operation for `servicegroup_servicegroupmemeber_binding` resource
* Implement login operation and partition targetting in provider
* Introduce no read flag in servicegroup member binding
* Enable use of proxy through HTTP\_PROXY and HTTPS\_PROXY environment variables

BUG FIXES

* Fix port logic error in `servicegroup_servicegroupmemeber_binding` resource
* Fix default lb monitor handling in `service` resource

## 1.7.0 (September 29, 2021)

FEATURES

* **New Resource** `csparameter`
* **New Resource** `cspolicylabel`
* **New Resource** `csvserver_analyticsprofile_binding`
* **New Resource** `csvserver_appqoepolicy_binding`
* **New Resource** `csvserver_auditnslogpolicy_binding`
* **New Resource** `csvserver_auditsyslogpolicy_binding`
* **New Resource** `csvserver_authorizationpolicy_binding`
* **New Resource** `csvserver_botpolicy_binding`
* **New Resource** `csvserver_cachepolicy_binding`
* **New Resource** `csvserver_contentinspectionpolicy_binding`
* **New Resource** `csvserver_feopolicy_binding`
* **New Resource** `csvserver_gslbvserver_binding`
* **New Resource** `csvserver_spilloverpolicy_binding`
* **New Resource** `csvserver_tmtrafficpolicy_binding`
* **New Resource** `csvserver_vpnvserver_binding`
* **New Resource** `policyhttpcallout`
* **New Resource** `policymap`
* **New Resource** `policyparam`
* **New Resource** `responderparam`
* **New Resource** `responderpolicylabel_responderpolicy_binding`
* **New Resource** `rewriteglobal_rewritepolicy_binding`
* **New Resource** `rewriteparam`
* **New Resource** `rewritepolicylabel_rewritepolicy_binding`
* **New Resource** `sslcacertgroup_sslcertkey_binding`
* **New Resource** `ssldtlsprofile`
* **New Resource** `sslfipskey`
* **New Resource** `ssllogprofile`
* **New Resource** `sslocspresponder`
* **New Resource** `sslpolicylabel`
* **New Resource** `sslpolicylabel_sslpolicy_binding`
* **New Resource** `sslprofile_sslcertkey_binding`
* **New Resource** `sslservice`
* **New Resource** `sslservice_ecccurve_binding`
* **New Resource** `sslservice_sslcertkey_binding`
* **New Resource** `sslservice_sslciphersuite_binding`
* **New Resource** `sslservicegroup`
* **New Resource** `sslservicegroup_ecccurve_binding`
* **New Resource** `sslservicegroup_sslcertkey_binding`
* **New Resource** `sslvserver`
* **New Resource** `sslvserver_ecccurve_binding`
* **New Resource** `sslvserver_sslciphersuite_binding`
* **New Resource** `vpnalwaysonprofile`
* **New Resource** `vpnclientlessaccesspolicy`
* **New Resource** `vpnclientlessaccessprofile`
* **New Resource** `vpneula`
* **New Resource** `vpnformssoaction`
* **New Resource** `vpnglobal_appcontroller_binding`

ENHANCEMENTS

* Update nshttpprofile resource to include latest options

## 1.6.0 (September 14, 2021)

FEATURES

* **New Resource** `botpolicy`
* **New Resource** `botpolicylabel`
* **New Resource** `botprofile`
* **New Resource** `botsettings`
* **New Resource** `lbgroup`
* **New Resource** `lbgroup_lbvserver_binding`
* **New Resource** `lbmetrictable`
* **New Resource** `lbmonitor_metric_binding`
* **New Resource** `lbmonitor_sslcertkey_binding`
* **New Resource** `lbprofile`
* **New Resource** `lbroute`
* **New Resource** `lbroute6`
* **New Resource** `lbsipparameters`
* **New Resource** `lbvserver_analyticsprofile_binding`
* **New Resource** `lbvserver_appflowpolicy_binding`
* **New Resource** `lbvserver_appqoepolicy_binding`
* **New Resource** `lbvserver_auditsyslogpolicy_binding`
* **New Resource** `lbvserver_authorizationpolicy_binding`
* **New Resource** `lbvserver_botpolicy_binding`
* **New Resource** `lbvserver_cachepolicy_binding`
* **New Resource** `lbvserver_contentinspectionpolicy_binding`
* **New Resource** `lbvserver_dnspolicy64_binding`
* **New Resource** `lbvserver_feopolicy_binding`
* **New Resource** `lbvserver_spilloverpolicy_binding`
* **New Resource** `lbvserver_tmtrafficpolicy_binding`
* **New Resource** `lbvserver_videooptimizationdetectionpolicy_binding`
* **New Resource** `lbvserver_videooptimizationpacingpolicy_binding`
* **New Resource** `sslcacertgroup`

## 1.5.0 (July 16, 2021)

ENHANCEMENTS

* add `save_on_destroy` option in `nsconfig_save` resrouce

## 1.4.0 (July 13, 2021)

FEATURES

* **New Resource** `vlan_nsip_binding`
* **New Resource** `vlan_interface_binding`
* **New Resource** `nsmode`

ENHANCEMENTS

* Extra cluster configuration options. Now a SNIP and vtysh commands can be added per node added to the cluster.
* Update to latest adc-nitro-go library.

BUG FIXES

* Correct the way the `route` resource is read from the ADC. Now takes into account the `ownergroup` attribute.
* Correct `ownernode` argument of `nsip` resource to TypeString. Now ownernode will be applied correctly.

## 1.3.0 (July 6, 2021)

ENHANCEMENTS

* Migrate NITRO library to citrix/adc-nitro-go

## 1.2.0 (June 30, 2021)

FEATURES

* **New Resource** `citrixadc_iptunnel`
* **New Resource** `citrixadc_lbparameter`
* **New Resource** `citrixadc_vlan`

BUG FIXES

* resource/citrixadc\_cluster: Check `masterstate` instead of `health` for determining when a node has succesfully joined the cluster.

## 1.1.0 (June 7, 2021)

FEATURES

* **New Resource** `citrixadc_lbvserver_servicegroup_binding`

ENHANCEMENTS

* resource/citrixadc\_servicegroup: Do not read `lbvservers` attribute when not set. Needed for maintaining idempotency in conjuction with  `citrixadc_lbvserver_servicegroup_binding` resource.

NOTES

* Correct documentation of citrixadc\_nslicense to list `ssh_host_pubkey` attribute as `Required`.

## 1.0.1 (May 13, 2021)

NOTES

* Correct category in documentation for csvserver\_appfwpolicy\_binding

## 1.0.0 (May 10, 2021)

FEATURES

* **New Resource** `appfwfieldtype`
* **New Resource** `appfwjsoncontenttype`
* **New Resource** `appfwpolicy`
* **New Resource** `appfwpolicylabel`
* **New Resource** `appfwprofile_denyurl_binding`
* **New Resource** `appfwprofile`
* **New Resource** `appfwprofile_starturl_binding`
* **New Resource** `appfwprofile_sqlinjection_binding`
* **New Resource** `appfwprofile_crosssitescripting_binding`
* **New Resource** `appfwprofile_cookieconsistency_binding`
* **New Resource** `appfwxmlcontenttype`
* **New Resource** `auditmessageaction`
* **New Resource** `auditsyslogaction`
* **New Resource** `auditsyslogpolicy`
* **New Resource** `clusterfiles_syncer`
* **New Resource** `cluster`
* **New Resource** `cmppolicy`
* **New Resource** `csaction`
* **New Resource** `cspolicy`
* **New Resource** `csvserver_appfwpolicy_binding`
* **New Resource** `csvserver_cmppolicy_binding`
* **New Resource** `csvserver_cspolicy_binding`
* **New Resource** `csvserver_filterpolicy_binding`
* **New Resource** `csvserver`
* **New Resource** `csvserver_responderpolicy_binding`
* **New Resource** `csvserver_rewritepolicy_binding`
* **New Resource** `csvserver_transformpolicy_binding`
* **New Resource** `dnsnsrec`
* **New Resource** `dnssoarec`
* **New Resource** `filterpolicy`
* **New Resource** `gslbservice`
* **New Resource** `gslbsite`
* **New Resource** `gslbvserver`
* **New Resource** `inat`
* **New Resource** `installer`
* **New Resource** `interface`
* **New Resource** `ipset`
* **New Resource** `lbmonitor`
* **New Resource** `lbvserver_appfwpolicy_binding`
* **New Resource** `lbvserver_cmppolicy_binding`
* **New Resource** `lbvserver_filterpolicy_binding`
* **New Resource** `lbvserver`
* **New Resource** `lbvserver_responderpolicy_binding`
* **New Resource** `lbvserver_rewritepolicy_binding`
* **New Resource** `lbvserver_service_binding`
* **New Resource** `lbvserver_transformpolicy_binding`
* **New Resource** `linkset`
* **New Resource** `netprofile`
* **New Resource** `nsacl`
* **New Resource** `nsacls`
* **New Resource** `nscapacity`
* **New Resource** `nsconfig_clear`
* **New Resource** `nsconfig_save`
* **New Resource** `nsconfig_update`
* **New Resource** `nsfeature`
* **New Resource** `nshttpprofile`
* **New Resource** `nsip6`
* **New Resource** `nsip`
* **New Resource** `nslicense`
* **New Resource** `nslicenseserver`
* **New Resource** `nsparam`
* **New Resource** `nsrpcnode`
* **New Resource** `nstcpparam`
* **New Resource** `nstcpprofile`
* **New Resource** `nsvpxparam`
* **New Resource** `password_resetter`
* **New Resource** `pinger`
* **New Resource** `policydataset`
* **New Resource** `policydataset_value_binding`
* **New Resource** `policyexpression`
* **New Resource** `policypatset`
* **New Resource** `policypatset_pattern_binding`
* **New Resource** `policystringmap`
* **New Resource** `policystringmap_pattern_binding`
* **New Resource** `quicbridgeprofile`
* **New Resource** `rebooter`
* **New Resource** `responderaction`
* **New Resource** `responderpolicy`
* **New Resource** `responderpolicylabel`
* **New Resource** `rewriteaction`
* **New Resource** `rewritepolicy`
* **New Resource** `rewritepolicylabel`
* **New Resource** `rnat`
* **New Resource** `route`
* **New Resource** `routerdynamicrouting`
* **New Resource** `server`
* **New Resource** `service`
* **New Resource** `servicegroup`
* **New Resource** `servicegroup_lbmonitor_binding`
* **New Resource** `servicegroup_servicegroupmember_binding`
* **New Resource** `sslaction`
* **New Resource** `sslcertkey`
* **New Resource** `sslcipher`
* **New Resource** `ssldhparam`
* **New Resource** `sslparameter`
* **New Resource** `sslpolicy`
* **New Resource** `sslprofile`
* **New Resource** `sslprofile_sslcipher_binding`
* **New Resource** `sslvserver_sslcertkey_binding`
* **New Resource** `sslvserver_sslpolicy_binding`
* **New Resource** `systemcmdpolicy`
* **New Resource** `systemextramgmtcpu`
* **New Resource** `systemfile`
* **New Resource** `systemgroup`
* **New Resource** `systemuser`
* **New Resource** `transformaction`
* **New Resource** `transformpolicy`
* **New Resource** `transformprofile`
