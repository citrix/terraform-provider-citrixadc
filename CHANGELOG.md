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
