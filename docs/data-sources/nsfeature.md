---
subcategory: "NS"
---

# Data Source `nsfeature`

The nsfeature data source allows you to retrieve information about the features enabled on the Citrix ADC appliance.


## Example usage

```terraform
data "citrixadc_nsfeature" "features" {
}

output "load_balancing_enabled" {
  value = data.citrixadc_nsfeature.features.lb
}

output "ssl_enabled" {
  value = data.citrixadc_nsfeature.features.ssl
}

output "content_switching_enabled" {
  value = data.citrixadc_nsfeature.features.cs
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `id` - The id of the nsfeature datasource. It has a static value of `nsfeature-config`.
* `wl` - Web Logging feature status.
* `sp` - Surge Protection feature status.
* `lb` - Load Balancing feature status.
* `cs` - Content Switching feature status.
* `cr` - Cache Redirection feature status.
* `cmp` - Compression feature status.
* `pq` - Priority Queuing feature status.
* `ssl` - SSL Offloading feature status.
* `gslb` - Global Server Load Balancing feature status.
* `hdosp` - DoS Protection feature status.
* `cf` - Content Filtering feature status.
* `ic` - Integrated Caching feature status.
* `sslvpn` - SSL VPN feature status.
* `aaa` - AAA feature status.
* `ospf` - OSPF Routing feature status.
* `rip` - RIP Routing feature status.
* `bgp` - BGP Routing feature status.
* `rewrite` - Rewrite feature status.
* `ipv6pt` - IPv6 Protocol Translation feature status.
* `appfw` - Application Firewall feature status.
* `responder` - Responder feature status.
* `htmlinjection` - HTML Injection feature status.
* `push` - Push feature status.
* `appflow` - AppFlow feature status.
* `cloudbridge` - CloudBridge feature status.
* `isis` - ISIS Routing feature status.
* `ch` - Call Home feature status.
* `appqoe` - AppQoE feature status.
* `contentaccelerator` - Content Accelerator feature status.
* `rise` - RISE feature status.
* `feo` - Front End Optimization feature status.
* `lsn` - Large Scale NAT feature status.
* `rdpproxy` - RDP Proxy feature status.
* `rep` - Reputation feature status.
* `urlfiltering` - URL Filtering feature status.
* `videooptimization` - Video Optimization feature status.
* `forwardproxy` - Forward Proxy feature status.
* `sslinterception` - SSL Interception feature status.
* `adaptivetcp` - Adaptive TCP feature status.
* `cqa` - Connection Quality Analytics feature status.
* `ci` - Content Inspection feature status.
* `bot` - Bot Management feature status.
* `apigateway` - API Gateway feature status.
