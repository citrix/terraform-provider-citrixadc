---
subcategory: "NS"
---

# Data Source: nstestlicense

The nstestlicense data source retrieves the read-only license feature flags of the NITRO `nstestlicense` object via its keyless `get (all)` endpoint. It is read-only and non-destructive (it does not apply a license).


## Example usage

```terraform
data "citrixadc_nstestlicense" "example" {
}

output "licensingmode" {
  value = data.citrixadc_nstestlicense.example.licensingmode
}
```


## Argument Reference

This data source takes no arguments.


## Attribute Reference

The following attributes are available:

* `id` - The id of the nstestlicense data source. It is a synthetic value (`nstestlicense-config`).
* `wl` - Web Logging.
* `sp` - Surge Protection.
* `lb` - Load Balancing.
* `cs` - Content Switching.
* `cr` - Cache Redirect.
* `cmp` - Compression.
* `delta` - Delta Compression.
* `ssl` - Secure Sockets Layer.
* `gslb` - Global Server Load Balancing.
* `gslbp` - GSLB Proximity.
* `routing` - Routing.
* `cf` - Content Filter.
* `contentaccelerator` - Transparent Integrated Caching.
* `ic` - Integrated Caching.
* `sslvpn` - SSL VPN.
* `f_sslvpn_users` - Number of licensed users allowed by this license.
* `f_ica_users` - Number of licensed users allowed by ICAONLY license.
* `aaa` - AAA.
* `ospf` - OSPF Routing.
* `rip` - RIP Routing.
* `bgp` - BGP Routing.
* `rewrite` - Rewrite.
* `ipv6pt` - IPv6 protocol translation.
* `appfw` - Application Firewall.
* `responder` - Responder.
* `agee` - Access Gateway Enterprise Edition.
* `nsxn` - NSXN.
* `modelid` - Model Number ID.
* `push` - Citrix ADC Push.
* `appflow` - AppFlow.
* `cloudbridge` - CloudBridge.
* `cloudbridgeappliance` - CloudBridge Appliance.
* `cloudextenderappliance` - Cloud Extender Appliance.
* `isis` - ISIS Routing.
* `cluster` - Clustering.
* `ch` - Call Home.
* `appqoe` - AppQoS.
* `appflowica` - Appflow for ICA.
* `isstandardlic` - Standard License.
* `isenterpriselic` - Enterprise License.
* `isplatinumlic` - Platinum License.
* `issgwylic` - Simple Gateway License.
* `isswglic` - Secure Web Gateway License.
* `feo` - Front End Optimization.
* `lsn` - Large Scale NAT.
* `licensingmode` - Pooled Licensed. Default value: `Local`.
* `daystoexpiration` - Days to expire.
* `rdpproxy` - RDPPROXY.
* `rep` - Reputation Services.
* `urlfiltering` - URL Filtering.
* `videooptimization` - Video Optimization.
* `forwardproxy` - Forward Proxy.
* `sslinterception` - SSL Interception.
* `remotecontentinspection` - Remote Content Inspection.
* `adaptivetcp` - Adaptive TCP.
* `cqa` - Connection Quality Analytics.
* `bot` - Bot Management.
* `apigateway` - API Gateway.
* `nextgenapiresource` - Read-only attribute (NITRO key `_nextgenapiresource`).
