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
