---
subcategory: "DNS"
---

# Data Source `dnszone`

The dnszone data source allows you to retrieve information about DNS zones configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_dnszone" "tf_dnszone" {
  zonename = "example.com"
  type     = "ALL"
}

output "proxymode" {
  value = data.citrixadc_dnszone.tf_dnszone.proxymode
}

output "dnssecoffload" {
  value = data.citrixadc_dnszone.tf_dnszone.dnssecoffload
}

output "nsec" {
  value = data.citrixadc_dnszone.tf_dnszone.nsec
}
```


## Argument Reference

* `zonename` - (Required) Name of the zone to retrieve.
* `type` - (Required) Type of zone to display. Mutually exclusive with the DNS Zone (zoneName) parameter. Available settings function as follows:
  * `ADNS` - Display all the zones for which the Citrix ADC is authoritative.
  * `PROXY` - Display all the zones for which the Citrix ADC is functioning as a proxy server.
  * `ALL` - Display all the zones configured on the appliance.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnszone. It has the same value as the `zonename` attribute.
* `dnssecoffload` - Enable dnssec offload for this zone.
* `keyname` - Name of the public/private DNS key pair with which to sign the zone. You can sign a zone with up to four keys.
* `nsec` - Enable nsec generation for dnssec offload.
* `proxymode` - Deploy the zone in proxy mode. Enable in the following scenarios:
  * The load balanced DNS servers are authoritative for the zone and all resource records that are part of the zone.
  * The load balanced DNS servers are authoritative for the zone, but the Citrix ADC owns a subset of the resource records that belong to the zone (partial zone ownership configuration). Typically seen in global server load balancing (GSLB) configurations, in which the appliance responds authoritatively to queries for GSLB domain names but forwards queries for other domain names in the zone to the load balanced servers.
  
  In either scenario, do not create the zone's Start of Authority (SOA) and name server (NS) resource records on the appliance.
  
  Disable if the appliance is authoritative for the zone, but make sure that you have created the SOA and NS records on the appliance before you create the zone.
