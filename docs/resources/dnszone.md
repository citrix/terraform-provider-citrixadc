---
subcategory: "DNS"
---

# Resource: dnszone

The dnszone resource is used to create DNS zone.


## Example usage

```hcl
resource "citrixadc_dnszone" "dnszone" {
	zonename      = "tf_zone1"
	proxymode     = "YES"
	dnssecoffload = "DISABLED"
	nsec          = "DISABLED"
  }
```


## Argument Reference

* `zonename` - (Required) Name of the zone to create.
* `proxymode` - (Reqiured) Deploy the zone in proxy mode. Enable in the following scenarios: * The load balanced DNS servers are authoritative for the zone and all resource records that are part of the zone.  * The load balanced DNS servers are authoritative for the zone, but the Citrix ADC owns a subset of the resource records that belong to the zone (partial zone ownership configuration). Typically seen in global server load balancing (GSLB) configurations, in which the appliance responds authoritatively to queries for GSLB domain names but forwards queries for other domain names in the zone to the load balanced servers. In either scenario, do not create the zone's Start of Authority (SOA) and name server (NS) resource records on the appliance.  Disable if the appliance is authoritative for the zone, but make sure that you have created the SOA and NS records on the appliance before you create the zone.
* `dnssecoffload` - (Optional) Enable dnssec offload for this zone.
* `keyname` - (Optional) Name of the public/private DNS key pair with which to sign the zone. You can sign a zone with up to four keys.
* `nsec` - (Optional) Enable nsec generation for dnssec offload.
* `type` - (Optional) Type of zone to display. Mutually exclusive with the DNS Zone (zoneName) parameter. Available settings function as follows: * ADNS - Display all the zones for which the Citrix ADC is authoritative. * PROXY - Display all the zones for which the Citrix ADC is functioning as a proxy server. * ALL - Display all the zones configured on the appliance.



## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnszone. It has the same value as the `zonename` attribute.


## Import

A dnszone can be imported using its name, e.g.

```shell
terraform import citrixadc_dnszone.dnszone tf_zone1
```
