## 1.43.3 (July 24, 2025)

BUG FIXES

* **citrixadc_sslvserver_sslcertkey_binding**: Fix for binding a certificate with snicert option.
* **citrixadc_appfwprofile_crosssitescripting_binding**: Fixing issues while creating HTML Cross Site Scripting Relaxations having optional attributes.

UPDATES

* **nsacls**: Updated documentation to include description for `acls_apply_trigger` attribute

## 1.43.2 (July 21, 2025)

FEATURES

* **citrixadc_nsacls**: Add `acls_apply_trigger` attribute that allows users to run apply nsacls every run if value is Yes. Includes validation and automatic reset mechanism for repeatable triggers

BUG FIXES

* **test.go**: Replace non-constant format string in `t.Errorf` with a constant format string for error reporting 
* **adc-nitro-go**: Fix URL encoding for resource names from PathEscape to QueryEscape
* **citrixadc_sslcertkey**: Fix for updating sslcertkey's passplain attribute value.
* **citrixadc_sslvserver_sslcertkey_binding**: Fix for binding a certificate with snicert option.
* **golang.org/x/oauth2**: Version upgrade to 0.27.0.

## 1.43.1 (May 12, 2025)

BUG FIXES

* **Go version**: Handled bug due to go version update and updated libraries. Updated go.mod file with  `godebug tlsrsakex=1` setting [#1251]

[#1251]: https://github.com/citrix/terraform-provider-citrixadc/issues/1251

## 1.43.0 (Feb 27, 2025)

FEATURES

* **New Resource**: citrixadc_sslrsakey [#1176]
* **New Resource**: citrixadc_sslecdsakey [#1176]

BUG FIXES

* **citrixadc_nsrpcnode**: handled clearing of `rpcnode` resource if it is deleted manually. [#1187]
* **citrixadc_sslprofile**: Updated sslprofile to handle the unbind of ecccurve and ciphers by introducing the attributes. [#1231]

ENHANCEMENTS

* **GoLang version**: Updated GoLang version from 1.19.0 to 1.23.0
* **citrixadc_sslcertkey**: Updated the behavior of cert and key attribute in the schema of sslcertkey resource to handle the updation of cert and key. [#1236]
* **citrixadc_nitro_info**: Updated the datasource nitro_info with query_args attribute. [#1178]
* **citrixadc_sslcertkey_update**: Added linkcertkeyname attribute to sslcertkey_update resource. [#1235]
* **External Libraries**: Updated the external libraries to their stable versions.

[#1187]: https://github.com/citrix/terraform-provider-citrixadc/issues/1187
[#1235]: https://github.com/citrix/terraform-provider-citrixadc/issues/1235
[#1178]: https://github.com/citrix/terraform-provider-citrixadc/issues/1178
[#1176]: https://github.com/citrix/terraform-provider-citrixadc/issues/1176
[#1236]: https://github.com/citrix/terraform-provider-citrixadc/issues/1236
[#1231]: https://github.com/citrix/terraform-provider-citrixadc/issues/1231

## 1.42.0 (Nov 21, 2024)

BUG FIXES

* **citrixadc_lbvserver**: Handled `timeout` attribute to accept zero value. [#1169]
* **citrixadc_appfwprofile_denyurl_binding**: Updated attributes schema behavior for attributes with `computed` as true.
* **documentation**: Updated docs with additional `order` attribute for supported resources.

ENHANCEMENTS

* **citrixadc_aaaparameter**: Updated aaaparameter resource with additional supported attributes.
* **citrixadc_gslbvserver**: Updated gslbvserver resource to handle backupvserver attribute and updated required attribute.
* **citrixadc_systemuser**: Updated systemuser resource to handle updating of attributes of nsroot user.

[#1169]: https://github.com/citrix/terraform-provider-citrixadc/issues/1169

## 1.41.0 (Oct 15, 2024)

FEATURES

* **New Resource**: citrixadc_sslprofile_ecccurve_binding
* **New Resource**: citrixadc_systemuser_systemcmdpolicy_binding
* **New Resource**: citrixadc_systemgroup_systemcmdpolicy_binding 
* **New Resource**: citrixadc_systemgroup_systemuser_binding
* **New Resource**: citrixadc_sslcipher_sslciphersuite_binding

BUG FIXES

* **citrixadc_lbmonitor**: Updated read func to handle the unexpected values for parameter units3 and interval. [#1165]
* **citrixadc_snmpalarmr**: Updated snmpalarm resource to accept zero value to time attribute and removed the snmpalarm struct and updated with map.
* **citrixadc_sslcipher**: In the read function, using GetAll insted of Get API call.
* **citrixadc_ip6tunnelparam**: Udpated read function, to handle the attribute that was not received from NetScaler.
* **citrixadc_dnsnsrec**: Updated the read function, to handle cases where the resource is not present on NetScaler.

ENHANCEMENTS

* **citrixadc_systemuser**: Updated systemuser resource with additional supported `allowedmanagementinterface` attribute.
* **External Libraries**: Updated the external libraries terratest and go-git to their stable versions.

UPDATES
* **DEPRECATED SOON**
    * **citrixadc_sslprofile**: The attributes `ecccurvebindings` and `cipherbindings` in `citrixadc_sslprofile` resource will be deprecated soon. Please use `citrixadc_sslprofile_ecccurve_binding` to bind `ecccurve` and `citrixadc_sslprofile_sslcipher_binding` to bind `sslcipher` to `sslprofile`.
    * **citrixadc_systemuser**: The attribute `cmdpolicybinding` in `citrixadc_systemuser` resource will be deprecated soon. Please use `citrixadc_systemuser_systemcmdpolicy_binding` to bind `systemcmdpolicy` to `systemuser`.
    * **citrixadc_systemgroup**: The attributes `cmdpolicybinding` and `systemusers` in `citrixadc_systemgroup` resource will be deprecated soon. Please use `citrixadc_systemgroup_systemcmdpolicy_binding` to bind `systemcmdpolicy` and `citrixadc_systemgroup_systemuser_binding` to bind `systemuser` to `systemgroup`.
    * **citrixadc_sslcipher**: The attribute `ciphersuitebinding` in `citrixadc_sslcipher` resource will be deprecated soon. Please use `citrixadc_sslcipher_sslciphersuite_binding` to bind `sslciphersuite` to `sslcipher`.

[#1165]: https://github.com/citrix/terraform-provider-citrixadc/issues/1165

## 1.40.1 (Sept 17, 2024)

BUG FIXES

* **citrixadc_appfwsignatures**: Updated appfwsignatures attribute Schema Behaviors to avoid unnecessary recreating of resource. [#1162]

ENHANCEMENTS

* **External Libraries**: Updated the versions of external libraries used in the project to stable version. Update the go version from 1.18 to 1.19.


[#1162]: https://github.com/citrix/terraform-provider-citrixadc/issues/1162

## 1.40.0 (Sept 2, 2024)

BUG FIXES

* **citrixadc_cachecontentgroup**: Updated read func to accept `cachecontrol` attribute as string insted of int [#1171]
* **citrixadc_appfwprofile_crosssitescripting_binding**: Updated delete method with supported query-parameters [#1177]
* **citrixadc_appfwprofile_csrftag_binding**: Updated delete method to escape the special character from query-parameter [#1175]
* **citrixadc_appfwprofile_sqlinjection_binding**: Updated delete method with supported query-parameters

ENHANCEMENTS

* **citrixadc_csvserver**: Updated the resource with additional supported `httpsredirecturl` attributes. [#1167] 
* **routerdynamicrouting**: updated `routerdynamcirouting` docs with limitation note for `commandlines` attribute  [#1179]
* **adc-nitro-go**: Updated the adc-nitro-go client library to encode the resourcename value before appending it to url. [#1180]

[#1171]: https://github.com/citrix/terraform-provider-citrixadc/issues/1171
[#1167]: https://github.com/citrix/terraform-provider-citrixadc/issues/1167
[#1177]: https://github.com/citrix/terraform-provider-citrixadc/issues/1177
[#1175]: https://github.com/citrix/terraform-provider-citrixadc/issues/1175
[#1179]: https://github.com/citrix/terraform-provider-citrixadc/issues/1179
[#1180]: https://github.com/citrix/terraform-provider-citrixadc/issues/1180

## 1.39.0 (May 10, 2024)

BUG FIXES

* **citrixadc_channel**: Updated read func to convert certain attributes from string to int [#1154]
* **citrixadc_interface**: Updated read func to convert certain attributes from string to int [#1154]
* **citrixadc_vlan**: Updated read func to convert certain attributes from string to int [#1154]
* **citrixadc_vlan_channel_binding**: Updated read func to convert certain attributes from string to int [#1154]
* **citrixadc_vlan_interface_binding**: Updated read func to convert certain attributes from string to int [#1154]
* **citrixadc_vlan_nsip_binding**: Updated read func to convert certain attributes from string to int [#1154]
* **citrixadc_dnssrvrec**: Updated Read func to handle importing operation and handled deletion [#1157]
* Handled `priority` argument which was not being imported for multiple 'binding' resources [#1153]

ENHANCEMENTS

* **citrixadc_gslbservicegroup_gslbservicegroupmember_binding**: Updated resource with additional supported `order` parameter [#1151] 
* **citrixadc_gslbvserver_gslbservice_binding**: Updated resource with additional supported `order` parameter [#1151] 
* **citrixadc_gslbvserver_gslbservicegroup_binding**: Updated resource with additional supported `order` parameter [#1151] 

[#1154]: https://github.com/citrix/terraform-provider-citrixadc/issues/1154
[#1157]: https://github.com/citrix/terraform-provider-citrixadc/issues/1157
[#1151]: https://github.com/citrix/terraform-provider-citrixadc/issues/1151
[#1153]: https://github.com/citrix/terraform-provider-citrixadc/issues/1153


## 1.38.0 (April 19, 2024)

FEATURES

* **New Resource**: citrixadc_sslcrl [#1109]
* **New Resource**: citrixadc_lbaction [#1129]
* **New Resource**: citrixadc_lbpolicy [#1129]
* **New Resource**: citrixadc_gslbvserver_lbpolicy_binding [#1129]
* **New Resource**: citrixadc_lbvserver_lbpolicy_binding [#1129]
* **New Resource**: citrixadc_analyticsglobal_analyticsprofile_binding [#1117] 
* **New Resource**: citrixadc_gslbservicegroup_gslbservicegroupmember_binding [#1127] 

BUG FIXES

* **citrixadc_vpnsessionaction**: Updated the read function to not to set value for some attributes from vpnsessionaction resource. [#1110]
* **citrixadc_lbmonitor**: Updated the resource struct to unset omitempty for `deviation` attribute [#1123]
* **citrixadc_lbvserver**: Updated read func to convert certain attributes from string to int [#1126]
* **citrixadc_service**: Updated read func to convert certain attributes from string to int [#1126]
* **citrixadc_servicegroup**: Updated read func to convert certain attributes from string to int [#1126]
* **citrixadc_aaagroup**: Updated read func to convert certain attributes from string to int [#1134]
* **citrixadc_dnssrvrec**: Updated read func to convert certain attributes from string to int [#1134]
* **citrixadc_gslbservice**: Updated read func to convert certain attributes from string to int [#1134]
* **citrixadc_gslbservice_lbmonitor_binding**: Updated read func to convert certain attributes from string to int [#1134]
* **citrixadc_gslbservicegroup**: Updated read func to convert certain attributes from string to int [#1134]
* **citrixadc_gslbservicegroup_lbmonitor_binding**: Updated read func to convert certain attributes from string to int [#1134]
* **citrixadc_gslbvserver**: Updated read func to convert certain attributes from string to int [#1134]
* **citrixadc_gslbvserver_gslbservice_binding**: Updated read func to convert certain attributes from string to int [#1134]
* **citrixadc_lbvserver**: Updated read func to convert certain attributes from string to int [#1134]
* **citrixadc_route**: Updated read func to convert certain attributes from string to int [#1134]
* **citrixadc_route6**: Updated read func to convert certain attributes from string to int [#1134]
* **citrixadc_service**: Updated read func to convert certain attributes from string to int [#1134]
* **citrixadc_servicegroup**: Updated read func to convert certain attributes from string to int [#1134]
* **citrixadc_service_lbmonitor_binding**: Updated read func to convert certain attributes from string to int [#1134]
* **citrixadc_servicegroup_lbmonitor_binding**: Updated read func to convert certain attributes from string to int [#1134]
* **citrixadc_servicegroup_servicegroupmember_binding**: Updated read func to convert certain attributes from string to int [#1134]

ENHANCEMENTS

* **citrixadc_appfwprofile**: Updated citrixadc_appfwprofile resource with additional supported attributes and addressed [#1085] issue. [#1135] 
* **citrixadc_appfwsignatures**: Updated citrixadc_appfwsignatures resource with additional supported attributes. [#1113] [#1112]


[#1110]: https://github.com/citrix/terraform-provider-citrixadc/issues/1110
[#1109]: https://github.com/citrix/terraform-provider-citrixadc/issues/1109
[#1126]: https://github.com/citrix/terraform-provider-citrixadc/issues/1126
[#1129]: https://github.com/citrix/terraform-provider-citrixadc/issues/1129
[#1135]: https://github.com/citrix/terraform-provider-citrixadc/issues/1135
[#1134]: https://github.com/citrix/terraform-provider-citrixadc/issues/1134
[#1085]: https://github.com/citrix/terraform-provider-citrixadc/issues/1085
[#1117]: https://github.com/citrix/terraform-provider-citrixadc/issues/1117
[#1127]: https://github.com/citrix/terraform-provider-citrixadc/issues/1127
[#1113]: https://github.com/citrix/terraform-provider-citrixadc/issues/1113
[#1112]: https://github.com/citrix/terraform-provider-citrixadc/issues/1112
[#1123]: https://github.com/citrix/terraform-provider-citrixadc/issues/1123


## 1.37.0 (Oct 19, 2023)

BUG FIXES

* **citrixadc_snmpalarm**: Updated the read function to not to set value for `normalvalue` attribute, that is not recieved from the NetScaler.
* **citrixadc_lbvserver**: Updated attribute `timeout` to accept 0 value from the user and set defalut value to 2.
* **citrixadc_sslprofile**: Updated updated function to handle the updation of `sesstimeout` attribute.
* **citrixadc_nsparam**: Updated the read function to not to set value for `timezone` attribute, as we recieve different value from the NetScaler.
* **citrixadc_lbmonitor**: Updated the read function to not to set value for `respcode` attribute, as we recieve different value from the NetScaler.
* **citrixadc_sslprofile**: Added ForceNew property to `sslprofiletype` argument.

ENHANCEMENTS

* **lbvserver_service_binding**: Updated the resource with additional supported `order` attribute.
* **citrixadc_analyticsprofile**: Updated the resource with additional supported `servemode` and `schemafile` attributes.
* **citrixadc_nsrpcnode**: Updated the resource with additional supported `validatecert` attribute and updated docs for `secure` attribute.
* **citrixadc_dnsaddrec**: Updated read operation to make it backward compatible, appended the attribute value to old Id. 
  
## 1.36.0 (July 27, 2023)

BUG FIXES

* **citrixadc_systemuser**: Updated Read function to set the username attribute value, that we get from the NetScaler.
* **citrixadc_dnsaddrec**: Updated Read function to split the id and parse the data that we get from the NetScaler to match the Id.
* **citrixadc_servicegroup**: Updated update operation to formulate appropriate playload whenever there is change in `cipheader` attribute.
* **citrixadc_snmpcommunity**: Updated `permissions` attribute property to required from optional, updated documentation with possible values.

ENHANCEMENTS

* **citrixadc_nstcpprofile**: Updated the resource with additional supported `sendclientportintcpoption` and `slowstartthreshold` attributes and also handled converting of string to integer for some attributes in the read function.
* **citrixadc_botsettings**: Updated the resource with additional supported `defaultnonintrusiveprofile` attribute and also supported import functionality.
* **citrixadc_lbvserver**: Updated the resource with additional supported `probeport` and `probeprotocol` attribute.
* **citrixadc_csvserver**: Updated the resource with additional supported `redirectfromport` attributes.


## 1.35.0 (May 26, 2023)

FEATURES

* **New Resource**: videooptimizationpacingaction 
* **New Resource**: videooptimizationpacingpolicy 

BUG FIXES

* **citrixadc_ntpsync**: Updated create operation to call appropriate API call.
* **citrixadc_appfwprofile**: Updated the read function to not to set values for some attributes that are not recieved from the NetScaler.
* **citrixadc_vpnvserver_appflowpolicy_binding**: Updated Id of the resource and Updated read operation to make it backward compatible. 

ENHANCEMENTS

* **citrixadc_appfwprofile_jsonsqlurl_binding**: Updated ppfwprofile_jsonsqlurl_binding resource with additional attributes.
* **citrixadc_lbparameter**: Added `sessionsthreshold` attribute to the schema of the resource.


## 1.34.0 (May 05, 2023)

FEATURES

* **New Resource**: lbmetrictable_metric_binding 
* **New Resource**: citrixadc_videooptimizationdetectionaction 
* **New Resource**: citrixadc_videooptimizationdetectionpolicy 
* **New Resource**: citrixadc_aaapreauthenticationaction 

BUG FIXES

* **citrixadc_snmptrap**: Updated read operation to make it backward compatible, appended the attribute value to old Id. 
* **citrixadc_policystringmap_pattern_binding**: Updated delete operation with QueryEscape function while making API call.
* **citrixadc_nspbr**: Updated the read function to not to set values for some attributes that are not recieved from the NetScaler. In the Update function, added the dependent attributes into the payload when there is update called to some attributes.
* **citrixadc_dnsnameserver**: Updated read operation to make it backward compatible, appended an attribute to old Id. 
* **citrixadc_systemfile**: Updated the read function to handle recreating of resource in case file doesnot exist in NetScaler.
* **citrixadc_hanode**: Updated read operation to handle the correct state of `hastatus` attribute. 


## 1.33.0 (April 03, 2023)

FEATURES

* **New Resource** snmptrap_snmpuser_binding

BUG FIXES

* **citrixadc_authenticationldapaction**: Updated the read function to not to set `ldapbinddnpassword` attribute from the NetScaler.
* **citrixadc_dnsnameserver**: Updated the delete function and updated computed as true for some of the attributes.
* **citrixadc_lbmonitor**: Updated the read function to not to set some attributes that are received as Hash value from NetScaler.
* **citrixadc_lbvserver**: Updated the resource struct to set omitempty for `timeout` attribute and made computed as true for `sslprofile` attribute.
* **citrixadc_sslprofile_sslcipher_binding**: Updated the schema to make the `cipherpriority` as the optional attribute.
* **citrixadc_sslprofile**: Updated the schema to make the `ecccurvebindings` as the optional attribute.
* **citrixadc_vpnvserver_vpnsessionpolicy_binding**: Updated the read function to not to set some attributes that are not received from NetScaler.
* **citrixadc_snmptrap**: Updated id of resource, it is now the concatenation of `trapclass`, `trapdestination` and `version` attributes seperated by comma. User can now add more snmptrap with different different combination of these value.  

ENHANCEMENTS

* **citrixadc_appfwpolicy**: Suppored Import to the resource.
* **citrixadc_appfwprofile**: Suppored Import to the resource.
* **citrixadc_lbvserver_appfwpolicy_binding**: Suppored Import to the resource.

## 1.32.0 (Feb 16, 2023)

FEATURES

* **New Resource** sslcertreq
* **New Resource** sslcert

BUG FIXES

* **citrixadc_route**: Updated the read function for fetching the resource instance from the NetScaler.
* **citrixadc_nspbr**: Handled setting of srcip value of nspbr resource.
* **citrixadc_appfwjsoncontenttype**: Escaped jsoncontenttype value for proper api calls and supported the Import functionality for the endpoint.
* **citrixadc_appfwxmlcontenttype**: Escaped xmlcontenttype value for proper api calls and supported the Import functionality for the endpoint.
* **provider.go**: Updated the nspartation to include only the `partitionname` in the request payload in provider.go file.

ENHANCEMENTS

* **citrixadc_systemgroup**: Updated the systemgroup resource with additional supported `allowedmanagementinterface` attribute.

## 1.31.0 (Feb 06, 2023)

BUG FIXES

* **citrixadc_csvserver**: Updated the csvserver resource with additional supported `persistencetype` attribute 
* **citrixadc_systemparameter**: Updated the systemparameter resource with additional supported `maxclient` attribute
* **citrixadc_auditsyslogpolicy**: Updated with the url QueryEscape for the attributes for the proper API calls.
* **citrixadc_auditsyslogglobal_auditsyslogpolicy_binding**: Updated with the url QueryEscape for the attributes for the proper API calls.

## 1.30.0 (Jan 25, 2023)

BUG FIXES

* **citrixadc_ntpserver**: Handled the ntpserver resource instance missing in the read function.
* **citrixadc_route**: Updated Read function to make the Import work as required.

ENHANCEMENTS

* Migrated to terraform sdkv1 and resolved the transitive dependencies.

## 1.29.0 (Jan 05, 2023)

ENHANCEMENTS

* **citrixadc_sslvserver**: Suppored Import to the resource.
* **citrixadc_policypatset**: Suppored Import to the resource.
* **citrixadc_policypatset_pattern_binding**: Suppored Import to the resource.

## 1.28.0 (Dec 20, 2022)

BUG FIXES

* **citrixadc_nscapacity**: Included warm rebooting of NetScaler after changing the license bandwidth, to read the latest snapshot of nscapacity. Removed some unnecessary attributes from schema.

## 1.27.0 (Dec 14, 2022)

FEATURES

* **New Resource** spilloverpolicy

BUG FIXES

* **citrixadc_dnsnameserver**: Updated id of resource, it is now the concatenation of `ip` (or `dnsvservername`) and `type` attributes seperated by comma. User can now add more dnsnameserver with same `ip` and different `type` value.  
* **citrixadc_dnsaddrec**: Updated id of resource, it is now the concatenation of `hostname` and `ipaddress` attributes separated by a comma. User can now create dnsaddrec with same `hostname` and multiple `ipaddress`.
* **citrixadc_nspartition**: Supporting the zero value for certatin attributes of `nspartition` and type conversion for attribute.

ENHANCEMENTS

* **citrixadc_botprofile**: Added missed attributes to resource `signaturemultipleuseragentheaderaction`, `signaturenouseragentheaderaction`, ` devicefingerprintmobile`, `kmjavascriptname`, `kmeventspostbodylimit`, `kmdetection`, and `clientipexpression`, 
* **citrixadc_systemuser**: Added condition to check the username and to throw error if the user wants to change the Admin password.

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
