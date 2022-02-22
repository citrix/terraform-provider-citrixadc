---
subcategory: "NS"
---

# Resource: nsfeature

The nsfeature resource is used to enable or disable ADC features.


## Example usage

```hcl
resource "citrixadc_nsfeature" "tf_nsfeature" {
    cs = true
    lb = true
    ssl = false
    appfw = false
}
```


## Argument Reference

The following arguments can set to `true` or `false` to enable or disable the corresponding feature.

* `wl` - (Optional) Web Logging.
* `sp` - (Optional) Surge Protection.
* `lb` - (Optional) Load Balancing.
* `cs` - (Optional) Content Switching.
* `cr` - (Optional) Cache Redirect.
* `cmp` - (Optional) Compression.
* `pq` - (Optional) Priority Queuing.
* `ssl` - (Optional) Secure Sockets Layer.
* `gslb` - (Optional) Global Server Load Balancing.
* `hdosp` - (Optional) HTTP DOS protection.
* `cf` - (Optional) Content Filter.
* `ic` - (Optional) Integrated Caching.
* `sslvpn` - (Optional) SSL VPN.
* `aaa` - (Optional) AAA.
* `ospf` - (Optional) OSPF Routing.
* `rip` - (Optional) RIP Routing.
* `bgp` - (Optional) BGP Routing.
* `rewrite` - (Optional) Rewrite.
* `ipv6pt` - (Optional) IPv6 protocol translation.
* `appfw` - (Optional) Application Firewall.
* `responder` - (Optional) Responder.
* `htmlinjection` - (Optional) HTML Injection.
* `push` - (Optional) Citrix ADC Push.
* `appflow` - (Optional) AppFlow.
* `cloudbridge` - (Optional) CloudBridge.
* `isis` - (Optional) ISIS Routing.
* `ch` - (Optional) Call Home.
* `appqoe` - (Optional) AppQoS.
* `contentaccelerator` - (Optional) Transparent Integrated Caching.
* `rise` - (Optional)
* `feo` - (Optional) Optimize Web content (HTML, CSS, JavaScript, images).
* `lsn` - (Optional) Large Scale NAT.
* `rdpproxy` - (Optional) RDPPROXY.
* `rep` - (Optional) Reputation Services.
* `urlfiltering` - (Optional) URL Filtering.
* `videooptimization` - (Optional) Video Optimization.
* `forwardproxy` - (Optional) Forward Proxy.
* `sslinterception` - (Optional) SSL Interception.
* `adaptivetcp` - (Optional) Adaptive TCP.
* `cqa` - (Optional) Connection Quality Analytics.
* `ci` - (Optional) Content Inspection.
* `bot` - (Optional) Bot Management.
* `apigateway` - (Optional) API Gateway.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsfeature resource. It is a random string prefixed with "tf-nsfeature-"
