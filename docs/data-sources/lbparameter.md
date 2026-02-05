---
subcategory: "Load Balancing"
---

# Data Source: citrixadc_lbparameter

The `citrixadc_lbparameter` data source is used to retrieve the load balancing parameters configuration.

## Example Usage

```hcl
data "citrixadc_lbparameter" "example" {}

output "lbparameter_details" {
  value = data.citrixadc_lbparameter.example
}
```

## Example Usage with Resource

```hcl
data "citrixadc_lbparameter" "tf_lbparameter" {
  depends_on = [citrixadc_lbparameter.tf_lbparameter]
}

output "configured_lbparameter" {
  value = data.citrixadc_lbparameter.tf_lbparameter
}
```

## Argument Reference

This data source does not require any arguments. It retrieves the current load balancing parameter configuration from the Citrix ADC.

## Attribute Reference

In addition to the above arguments, the following attributes are exported:

* `id` - The id of the lbparameter data source. It is a unique string prefixed with "tf-lbparameter".
* `allowboundsvcremoval` - This is used to enable/disable the option of svc/svcgroup removal, if it is bound to one or more vserver. If it is enabled, the svc/svcgroup can be removed, even if it bound to vservers. If disabled, an error will be thrown, when the user tries to remove a svc/svcgroup without unbinding from its vservers. Possible values: `ENABLED`, `DISABLED`.
* `computedadccookieattribute` - ComputedADCCookieAttribute accepts ns variable as input in form of string starting with $ (to understand how to configure ns variable, please check man add ns variable). Policies can be configured to modify this variable for every transaction and the final value of the variable after policy evaluation will be appended as attribute to Citrix ADC cookie (for example: LB cookie persistence, GSLB sitepersistence, CS cookie persistence, LB group cookie persistence).
* `consolidatedlconn` - To find the service with the fewest connections, the virtual server uses the consolidated connection statistics from all the packet engines. The NO setting allows consideration of only the number of connections on the packet engine that received the new connection. Possible values: `YES`, `NO`.
* `cookiepassphrase` - Use this parameter to specify the passphrase used to generate secured persistence cookie value. It specifies the passphrase with a maximum of 31 characters.
* `dbsttl` - Specify the TTL for DNS record for domain based service. The default value of ttl is 0 which indicates to use the TTL received in DNS response for monitors.
* `dropmqttjumbomessage` - When this option is enabled, MQTT messages of length greater than 64k will be dropped and the client/server connections will be reset. Possible values: `YES`, `NO`.
* `httponlycookieflag` - Include the HttpOnly attribute in persistence cookies. The HttpOnly attribute limits the scope of a cookie to HTTP requests and helps mitigate the risk of cross-site scripting attacks. Possible values: `ENABLED`, `DISABLED`.
* `lbhashalgorithm` - This option dictates the hashing algorithm used for hash based LB methods (URLHASH, DOMAINHASH, SOURCEIPHASH, DESTINATIONIPHASH, SRCIPDESTIPHASH, SRCIPSRCPORTHASH, TOKEN, USER_TOKEN, CALLIDHASH). Possible values: `DEFAULT`, `PRAC`, `JARH`.
* `lbhashfingers` - This option is used to specify the number of fingers to be used in PRAC and JARH algorithms for hash based LB methods. Increasing the number of fingers might give better distribution of traffic at the expense of additional memory.
* `literaladccookieattribute` - String configured as LiteralADCCookieAttribute will be appended as attribute for Citrix ADC cookie (for example: LB cookie persistence, GSLB site persistence, CS cookie persistence, LB group cookie persistence). Sample usage - set lb parameter -LiteralADCCookieAttribute ";SameSite=None".
* `maxpipelinenat` - Maximum number of concurrent requests to allow on a single client connection, which is identified by the <clientip:port>-<vserver ip:port> tuple. This parameter is useful for SSL/TLS, MQTT, and HTTP/2 connections.
* `monitorconnectionclose` - Close monitoring connections by sending the service a connection termination message with the specified bit set. Possible values: `RESET`, `FIN`.
* `monitorskipmaxclient` - When a monitor initiates a connection to a service, do not check to determine whether the number of connections to the service has reached the limit specified by the service's Max Clients setting. Enables monitoring to continue even if the service has reached its connection limit. Possible values: `ENABLED`, `DISABLED`.
* `preferdirectroute` - Perform route lookup for traffic received by the Citrix ADC, and forward the traffic according to configured routes. Do not set this parameter if you want a wildcard virtual server to direct packets received by the appliance to an intermediary device, such as a firewall, even if their destination is directly connected to the appliance. Route lookup is performed after the packets have been processed and returned by the intermediary device. Possible values: `YES`, `NO`.
* `proximityfromself` - Use the ADC location instead of client IP for static proximity LB or GSLB decision.
* `retainservicestate` - This option is used to retain the original state of service or servicegroup member when an enable server command is issued. Possible values: `ON`, `OFF`.
* `sessionsthreshold` - This option is used to set the upper-limit on the number of persistent sessions.
* `startuprrfactor` - Number of requests, per service, for which to apply the round robin load balancing method before switching to the configured load balancing method, thus allowing services to ramp up gradually to full load. Until the specified number of requests is distributed, the Citrix ADC is said to be implementing the slow start mode (or startup round robin).
* `storemqttclientidandusername` - This option allows to store the MQTT clientid and username in transactional logs. Possible values: `YES`, `NO`.
* `undefaction` - Action to perform when policy evaluation creates an UNDEF condition. Available settings function as follows: NOLBACTION - Does not consider LB action in making LB decision, RESET - Reset the request and notify the user, so that the user can resend the request, DROP - Drop the request without sending a response to the user.
* `useencryptedpersistencecookie` - Encode persistence cookie values using SHA2 hash. Possible values: `ENABLED`, `DISABLED`.
* `useportforhashlb` - Include the port number of the service when creating a hash for hash based load balancing methods. With the NO setting, only the IP address of the service is considered when creating a hash. Possible values: `YES`, `NO`.
* `usesecuredpersistencecookie` - Encode persistence cookie values using SHA2 hash. Possible values: `ENABLED`, `DISABLED`.
* `vserverspecificmac` - Allow a MAC-mode virtual server to accept traffic returned by an intermediary device, such as a firewall, to which the traffic was previously forwarded by another MAC-mode virtual server. The second virtual server can then distribute that traffic across the destination server farm. Also useful when load balancing Branch Repeater appliances. Possible values: `ENABLED`, `DISABLED`.
