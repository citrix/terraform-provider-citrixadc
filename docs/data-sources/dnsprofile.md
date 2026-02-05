---
subcategory: "DNS"
---

# Data Source `dnsprofile`

The dnsprofile data source allows you to retrieve information about a DNS profile configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_dnsprofile" "tf_dnsprofile" {
  dnsprofilename = "tf_profile1"
}

output "dns_query_logging" {
  value = data.citrixadc_dnsprofile.tf_dnsprofile.dnsquerylogging
}

output "dns_cache_records" {
  value = data.citrixadc_dnsprofile.tf_dnsprofile.cacherecords
}

output "dns_recursive_resolution" {
  value = data.citrixadc_dnsprofile.tf_dnsprofile.recursiveresolution
}
```


## Argument Reference

* `dnsprofilename` - (Required) Name of the DNS profile.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnsprofile. It is the same as the `dnsprofilename`.
* `cacheecsresponses` - Cache DNS responses with EDNS Client Subnet(ECS) option in the DNS cache. When disabled, the appliance stops caching responses with ECS option. This is relevant to proxy configuration. Enabling/disabling support of ECS option when Citrix ADC is authoritative for a GSLB domain is supported using a knob in GSLB vserver. In all other modes, ECS option is ignored.
* `cachenegativeresponses` - Cache negative responses in the DNS cache. When disabled, the appliance stops caching negative responses except referral records. This applies to all configurations - proxy, end resolver, and forwarder. However, cached responses are not flushed. The appliance does not serve negative responses from the cache until this parameter is enabled again.
* `cacherecords` - Cache resource records in the DNS cache. Applies to resource records obtained through proxy configurations only. End resolver and forwarder configurations always cache records in the DNS cache, and you cannot disable this behavior. When you disable record caching, the appliance stops caching server responses. However, cached records are not flushed. The appliance does not serve requests from the cache until record caching is enabled again.
* `dnsanswerseclogging` - DNS answer section; if enabled, answer section in the response will be logged.
* `dnserrorlogging` - DNS error logging; if enabled, whenever error is encountered in DNS module reason for the error will be logged.
* `dnsextendedlogging` - DNS extended logging; if enabled, authority and additional section in the response will be logged.
* `dnsquerylogging` - DNS query logging; if enabled, DNS query information such as DNS query id, DNS query flags , DNS domain name and DNS query type will be logged.
* `dropmultiqueryrequest` - Drop the DNS requests containing multiple queries. When enabled, DNS requests containing multiple queries will be dropped. In case of proxy configuration by default the DNS request containing multiple queries is forwarded to the backend and in case of ADNS and Resolver configuration NOCODE error response will be sent to the client.
* `insertecs` - Insert ECS Option on DNS query.
* `maxcacheableecsprefixlength` - The maximum ecs prefix length that will be cached.
* `maxcacheableecsprefixlength6` - The maximum ecs prefix length that will be cached for IPv6 subnets.
* `recursiveresolution` - DNS recursive resolution; if enabled, will do recursive resolution for DNS query when the profile is associated with ADNS service, CS Vserver and DNS action.
* `replaceecs` - Replace ECS Option on DNS query.
