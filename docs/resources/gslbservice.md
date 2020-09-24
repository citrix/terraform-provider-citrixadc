---
subcategory: "GSLB"
---

# Resource: gslbservice

This resource is used to manage Global Server Load Balancing service.


## Example usage

```hcl
resource "citrixadc_gslbsite" "site_local" {
  sitename        = "Site-Local"
  siteipaddress   = "172.31.96.234"
  sessionexchange = "DISABLED"
}

resource "citrixadc_gslbservice" "tf_gslbservice" {
  ip          = "172.16.1.121"
  port        = "80"
  servicename = "tf_gslbservice"
  servicetype = "HTTP"
  sitename    = citrixadc_gslbsite.site_local.sitename
}
```


## Argument Reference

* `servicename` - (Optional) Name for the GSLB service.
* `cnameentry` - (Optional) Canonical name of the GSLB service. Used in CNAME-based GSLB.
* `ip` - (Optional) IP address for the GSLB service. Should represent a load balancing, content switching, or VPN virtual server on the Citrix ADC, or the IP address of another load balancing device.
* `servername` - (Optional) Name of the server hosting the GSLB service.
* `servicetype` - (Optional) Type of service to create. Possible values: [ HTTP, FTP, TCP, UDP, SSL, SSL\_BRIDGE, SSL\_TCP, NNTP, ANY, SIP\_UDP, SIP\_TCP, SIP\_SSL, RADIUS, RDP, RTSP, MYSQL, MSSQL, ORACLE ]
* `port` - (Optional) Port on which the load balancing entity represented by this GSLB service listens.
* `publicip` - (Optional) The public IP address that a NAT device translates to the GSLB service's private IP address. Optional.
* `publicport` - (Optional) The public port associated with the GSLB service's public IP address. The port is mapped to the service's private port number. Applicable to the local GSLB service.
* `maxclient` - (Optional) The maximum number of open connections that the service can support at any given time. A GSLB service whose connection count reaches the maximum is not considered when a GSLB decision is made, until the connection count drops below the maximum.
* `healthmonitor` - (Optional) Monitor the health of the GSLB service. Possible values: [ YES, NO ]
* `sitename` - (Optional) Name of the GSLB site to which the service belongs.
* `state` - (Optional) Enable or disable the service. Possible values: [ ENABLED, DISABLED ]
* `cip` - (Optional) In the request that is forwarded to the GSLB service, insert a header that stores the client's IP address. Client IP header insertion is used in connection-proxy based site persistence. Possible values: [ ENABLED, DISABLED ]
* `cipheader` - (Optional) Name for the HTTP header that stores the client's IP address. Used with the Client IP option. If client IP header insertion is enabled on the service and a name is not specified for the header, the Citrix ADC uses the name specified by the cipHeader parameter in the set ns param command or, in the GUI, the Client IP Header parameter in the Configure HTTP Parameters dialog box.
* `sitepersistence` - (Optional) Use cookie-based site persistence. Applicable only to HTTP and SSL GSLB services. Possible values: [ ConnectionProxy, HTTPRedirect, NONE ]
* `cookietimeout` - (Optional) Timeout value, in minutes, for the cookie, when cookie based site persistence is enabled.
* `siteprefix` - (Optional) The site's prefix string. When the service is bound to a GSLB virtual server, a GSLB site domain is generated internally for each bound service-domain pair by concatenating the site prefix of the service and the name of the domain. If the special string NONE is specified, the site-prefix string is unset. When implementing HTTP redirect site persistence, the Citrix ADC redirects GSLB requests to GSLB services by using their site domains.
* `clttimeout` - (Optional) Idle time, in seconds, after which a client connection is terminated. Applicable if connection proxy based site persistence is used.
* `svrtimeout` - (Optional) Idle time, in seconds, after which a server connection is terminated. Applicable if connection proxy based site persistence is used.
* `maxbandwidth` - (Optional) Integer specifying the maximum bandwidth allowed for the service. A GSLB service whose bandwidth reaches the maximum is not considered when a GSLB decision is made, until its bandwidth consumption drops below the maximum.
* `downstateflush` - (Optional) Flush all active transactions associated with the GSLB service when its state transitions from UP to DOWN. Do not enable this option for services that must complete their transactions. Applicable if connection proxy based site persistence is used. Possible values: [ ENABLED, DISABLED ]
* `maxaaausers` - (Optional) Maximum number of SSL VPN users that can be logged on concurrently to the VPN virtual server that is represented by this GSLB service. A GSLB service whose user count reaches the maximum is not considered when a GSLB decision is made, until the count drops below the maximum.
* `monthreshold` - (Optional) Monitoring threshold value for the GSLB service. If the sum of the weights of the monitors that are bound to this GSLB service and are in the UP state is not equal to or greater than this threshold value, the service is marked as DOWN.
* `hashid` - (Optional) Unique hash identifier for the GSLB service, used by hash based load balancing methods.
* `comment` - (Optional) Any comments that you might want to associate with the GSLB service.
* `appflowlog` - (Optional) Enable logging appflow flow information. Possible values: [ ENABLED, DISABLED ]
* `naptrreplacement` - (Optional) The replacement domain name for this NAPTR.
* `naptrorder` - (Optional) An integer specifying the order in which the NAPTR records MUST be processed in order to accurately represent the ordered list of Rules. The ordering is from lowest to highest.
* `naptrservices` - (Optional) Service Parameters applicable to this delegation path.
* `naptrdomainttl` - (Optional) Modify the TTL of the internally created naptr domain.
* `naptrpreference` - (Optional) An integer specifying the preference of this NAPTR among NAPTR records having same order. lower the number, higher the preference.
* `ipaddress` - (Optional) The new IP address of the service.
* `viewname` - (Optional) Name of the DNS view of the service. A DNS view is used in global server load balancing (GSLB) to return a predetermined IP address to a specific group of clients, which are identified by using a DNS policy.
* `viewip` - (Optional) IP address to be used for the given view.
* `weight` - (Optional) Weight to assign to the monitor-service binding. A larger number specifies a greater weight. Contributes to the monitoring threshold, which determines the state of the service.
* `monitornamesvc` - (Optional) Name of the monitor to bind to the service.
* `lbmonitorbinding` - (Optional) A set of lb monitor blocks. Documented below


Lb monitor supports the following:

* `weight` - (Optional) Weight to assign to the monitor-service binding. A larger number specifies a greater weight. Contributes to the monitoring threshold, which determines the state of the service.
* `monitor_name` - (Optional) Monitor name.
* `monstate` - (Optional) State of the monitor bound to gslb service. Possible values: [ ENABLED, DISABLED ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the glsbservice. It has the same value as the `servicename` attribute.


## Import

An instance of the resource can be imported using its name, e.g.

```shell
terraform import citrixadc_gslbservice.tf_gslbservice tf_gslbservice
```
