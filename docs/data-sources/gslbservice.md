---
subcategory: "GSLB"
---

# Data Source `gslbservice`

The gslbservice data source allows you to retrieve information about a GSLB service.


## Example usage

```terraform
data "citrixadc_gslbservice" "tf_gslbservice" {
  servicename = "my_gslbservice"
}

output "ipaddress" {
  value = data.citrixadc_gslbservice.tf_gslbservice.ipaddress
}

output "port" {
  value = data.citrixadc_gslbservice.tf_gslbservice.port
}

output "servicetype" {
  value = data.citrixadc_gslbservice.tf_gslbservice.servicetype
}

output "sitename" {
  value = data.citrixadc_gslbservice.tf_gslbservice.sitename
}
```


## Argument Reference

* `servicename` - (Required) Name for the GSLB service. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the GSLB service is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `appflowlog` - Enable logging appflow flow information.
* `cip` - In the request that is forwarded to the GSLB service, insert a header that stores the client's IP address. Client IP header insertion is used in connection-proxy based site persistence.
* `cipheader` - Name for the HTTP header that stores the client's IP address. Used with the Client IP option. If client IP header insertion is enabled on the service and a name is not specified for the header, the Citrix ADC uses the name specified by the cipHeader parameter in the set ns param command or, in the GUI, the Client IP Header parameter in the Configure HTTP Parameters dialog box.
* `clttimeout` - Idle time, in seconds, after which a client connection is terminated. Applicable if connection proxy based site persistence is used.
* `cnameentry` - Canonical name of the GSLB service. Used in CNAME-based GSLB.
* `comment` - Any comments that you might want to associate with the GSLB service.
* `cookietimeout` - Timeout value, in minutes, for the cookie, when cookie based site persistence is enabled.
* `downstateflush` - Flush all active transactions associated with the GSLB service when its state transitions from UP to DOWN. Do not enable this option for services that must complete their transactions. Applicable if connection proxy based site persistence is used.
* `hashid` - Unique hash identifier for the GSLB service, used by hash based load balancing methods.
* `healthmonitor` - Monitor the health of the GSLB service.
* `ip` - IP address for the GSLB service. Should represent a load balancing, content switching, or VPN virtual server on the Citrix ADC, or the IP address of another load balancing device.
* `ipaddress` - The new IP address of the service.
* `maxaaausers` - Maximum number of SSL VPN users that can be logged on concurrently to the VPN virtual server that is represented by this GSLB service. A GSLB service whose user count reaches the maximum is not considered when a GSLB decision is made, until the count drops below the maximum.
* `maxbandwidth` - Integer specifying the maximum bandwidth allowed for the service. A GSLB service whose bandwidth reaches the maximum is not considered when a GSLB decision is made, until its bandwidth consumption drops below the maximum.
* `maxclient` - The maximum number of open connections that the service can support at any given time. A GSLB service whose connection count reaches the maximum is not considered when a GSLB decision is made, until the connection count drops below the maximum.
* `monitor_name_svc` - Name of the monitor to bind to the service.
* `monthreshold` - Monitoring threshold value for the GSLB service. If the sum of the weights of the monitors that are bound to this GSLB service and are in the UP state is not equal to or greater than this threshold value, the service is marked as DOWN.
* `naptrdomainttl` - Modify the TTL of the internally created naptr domain.
* `naptrorder` - An integer specifying the order in which the NAPTR records MUST be processed in order to accurately represent the ordered list of Rules. The ordering is from lowest to highest.
* `naptrpreference` - An integer specifying the preference of this NAPTR among NAPTR records having same order. lower the number, higher the preference.
* `naptrreplacement` - The replacement domain name for this NAPTR.
* `naptrservices` - Service Parameters applicable to this delegation path.
* `newname` - New name for the GSLB service.
* `port` - Port on which the load balancing entity represented by this GSLB service listens.
* `publicip` - The public IP address that a NAT device translates to the GSLB service's private IP address. Optional.
* `publicport` - The public port associated with the GSLB service's public IP address. The port is mapped to the service's private port number. Applicable to the local GSLB service. Optional.
* `servername` - Name of the server hosting the GSLB service.
* `servicetype` - Type of service to create.
* `sitename` - Name of the GSLB site to which the service belongs.
* `sitepersistence` - Use cookie-based site persistence. Applicable only to HTTP and SSL GSLB services.
* `siteprefix` - The site's prefix string. When the service is bound to a GSLB virtual server, a GSLB site domain is generated internally for each bound service-domain pair by concatenating the site prefix of the service and the name of the domain. If the special string NONE is specified, the site-prefix string is unset. When implementing HTTP redirect site persistence, the Citrix ADC redirects GSLB requests to GSLB services by using their site domains.
* `state` - Enable or disable the service.
* `svrtimeout` - Idle time, in seconds, after which a server connection is terminated. Applicable if connection proxy based site persistence is used.
* `viewip` - IP address to be used for the given view.
* `viewname` - Name of the DNS view of the service. A DNS view is used in global server load balancing (GSLB) to return a predetermined IP address to a specific group of clients, which are identified by using a DNS policy.
* `weight` - Weight to assign to the monitor-service binding. A larger number specifies a greater weight. Contributes to the monitoring threshold, which determines the state of the service.

## Attribute Reference

* `id` - The id of the gslbservice. It has the same value as the `servicename` attribute.


## Import

A gslbservice can be imported using its name, e.g.

```shell
terraform import citrixadc_gslbservice.tf_gslbservice my_gslbservice
```
