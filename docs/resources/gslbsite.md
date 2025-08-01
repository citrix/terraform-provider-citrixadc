---
subcategory: "GSLB"
---

# Resource: gslbsite

This resource is used to manage Global Service Load Balancing site.


## Example usage

```hcl
resource "citrixadc_gslbsite" "tf_site_local" {
  sitename        = "tf_site_local"
  siteipaddress   = "172.31.96.234"
  sessionexchange = "DISABLED"
}
```


## Argument Reference

* `sitename` - (Optional) Name for the GSLB site.
* `sitetype` - (Optional) Type of site to create. If the type is not specified, the appliance automatically detects and sets the type on the basis of the IP address being assigned to the site. If the specified site IP address is owned by the appliance (for example, a MIP address or SNIP address), the site is a local site. Otherwise, it is a remote site. Possible values: [ REMOTE, LOCAL ]
* `siteipaddress` - (Optional) IP address for the GSLB site. The GSLB site uses this IP address to communicate with other GSLB sites. For a local site, use any IP address that is owned by the appliance (for example, a SNIP or MIP address, or the IP address of the ADNS service).
* `publicip` - (Optional) Public IP address for the local site. Required only if the appliance is deployed in a private address space and the site has a public IP address hosted on an external firewall or a NAT device.
* `metricexchange` - (Optional) Exchange metrics with other sites. Metrics are exchanged by using Metric Exchange Protocol (MEP). The appliances in the GSLB setup exchange health information once every second.

    If you disable metrics exchange, you can use only static load balancing methods (such as round robin, static proximity, or the hash-based methods), and if you disable metrics exchange when a dynamic load balancing method (such as least connection) is in operation, the appliance falls back to round robin. Also, if you disable metrics exchange, you must use a monitor to determine the state of GSLB services. Otherwise, the service is marked as DOWN.

    Possible values: [ ENABLED, DISABLED ]

* `nwmetricexchange` - (Optional) Exchange, with other GSLB sites, network metrics such as round-trip time (RTT), learned from communications with various local DNS (LDNS) servers used by clients. RTT information is used in the dynamic RTT load balancing method, and is exchanged every 5 seconds. Possible values: [ ENABLED, DISABLED ]
* `sessionexchange` - (Optional) Exchange persistent session entries with other GSLB sites every five seconds. Possible values: [ ENABLED, DISABLED ]
* `triggermonitor` - (Optional) Specify the conditions under which the GSLB service must be monitored by a monitor, if one is bound. Available settings function as follows:

    * ALWAYS - Monitor the GSLB service at all times.
    * MEPDOWN - Monitor the GSLB service only when the exchange of metrics through the Metrics Exchange Protocol (MEP) is disabled.
    * MEPDOWN\_SVCDOWN - Monitor the service in either of the following situations:
    The exchange of metrics through MEP is disabled.
    The exchange of metrics through MEP is enabled but the status of the service, learned through metrics exchange, is DOWN.

* `parentsite` - (Optional) Parent site of the GSLB site, in a parent-child topology.
* `clip` - (Optional) Cluster IP address. Specify this parameter to connect to the remote cluster site for GSLB auto-sync. Note: The cluster IP address is defined when creating the cluster.
* `publicclip` - (Optional) IP address to be used to globally access the remote cluster when it is deployed behind a NAT. It can be same as the normal cluster IP address.
* `naptrreplacementsuffix` - (Optional) The naptr replacement suffix configured here will be used to construct the naptr replacement field in NAPTR record.
* `backupparentlist` - (Optional) The list of backup gslb sites configured in preferred order. Need to be parent gsb sites.
* `sitepassword` - (Optional) Password to be used for mep communication between gslb site nodes.
* `newname` - (Optional) New name for the GSLB site.
  
## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the gslbsite. It has the same value as the `sitename` attribute.


## Import

An instance of the resource can be imported using its sitename, e.g.

```shell
terraform import citrixadc_gslbsite.tf_site_local tf_site_local
```
