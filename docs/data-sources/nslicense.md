---
subcategory: "NS"
---

# citrixadc_nslicense (Data Source)

Data source for querying Citrix ADC license information. This data source retrieves information about the currently installed license and the features that are enabled on the ADC.

## Example Usage

```hcl
data "citrixadc_nslicense" "example" {
}

# Output license information
output "licensing_mode" {
  value = data.citrixadc_nslicense.example.licensingmode
}

output "is_platinum" {
  value = data.citrixadc_nslicense.example.isplatinumlic
}

output "lb_enabled" {
  value = data.citrixadc_nslicense.example.lb
}
```

## Argument Reference

This data source does not require any arguments.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - The ID of the nslicense datasource (always "nslicense").
* `licensingmode` - Licensing mode (e.g., EXPRESS, POOLED).
* `modelid` - Model ID of the appliance.

### License Type Attributes

* `isstandardlic` - Standard license is applied.
* `isenterpriselic` - Enterprise license is applied.
* `isplatinumlic` - Platinum license is applied.
* `issgwylic` - Secure Gateway license is applied.
* `isswglic` - SWG license is applied.

### Feature License Attributes

* `wl` - Web Logging feature is licensed.
* `sp` - Surge Protection feature is licensed.
* `lb` - Load Balancing feature is licensed.
* `cs` - Content Switching feature is licensed.
* `cr` - Cache Redirection feature is licensed.
* `cmp` - Compression feature is licensed.
* `delta` - Delta Compression feature is licensed.
* `ssl` - SSL Offloading feature is licensed.
* `gslb` - Global Server Load Balancing feature is licensed.
* `gslbp` - GSLB Proximity feature is licensed.
* `routing` - Routing feature is licensed.
* `cf` - Content Filtering feature is licensed.
* `contentaccelerator` - Content Accelerator feature is licensed.
* `ic` - Integrated Caching feature is licensed.
* `sslvpn` - SSL VPN feature is licensed.
* `f_sslvpn_users` - Number of SSL VPN users licensed.
* `f_ica_users` - Number of ICA users licensed.
* `aaa` - AAA (Authentication, Authorization, Accounting) feature is licensed.
* `ospf` - OSPF routing feature is licensed.
* `rip` - RIP routing feature is licensed.
* `bgp` - BGP routing feature is licensed.
* `rewrite` - Rewrite feature is licensed.
* `ipv6pt` - IPv6 Protocol Translation feature is licensed.
* `appfw` - Application Firewall feature is licensed.
* `responder` - Responder feature is licensed.
* `agee` - AGEE feature is licensed.
* `nsxn` - NetScaler XN feature is licensed.
* `push` - Push feature is licensed.
* `appflow` - AppFlow feature is licensed.
* `cloudbridge` - CloudBridge feature is licensed.
* `cloudbridgeappliance` - CloudBridge Appliance feature is licensed.
* `cloudextenderappliance` - CloudExtender Appliance feature is licensed.
* `isis` - ISIS routing feature is licensed.
* `cluster` - Cluster feature is licensed.
* `ch` - Call Home feature is licensed.
* `appqoe` - AppQoE feature is licensed.
* `appflowica` - AppFlow for ICA feature is licensed.
* `feo` - Front End Optimization feature is licensed.
* `lsn` - Large Scale NAT feature is licensed.
* `rdpproxy` - RDP Proxy feature is licensed.
* `rep` - Reputation feature is licensed.
* `urlfiltering` - URL Filtering feature is licensed.
* `videooptimization` - Video Optimization feature is licensed.
* `forwardproxy` - Forward Proxy feature is licensed.
* `sslinterception` - SSL Interception feature is licensed.
* `remotecontentinspection` - Remote Content Inspection feature is licensed.
* `adaptivetcp` - Adaptive TCP feature is licensed.
* `cqa` - CQA feature is licensed.
* `bot` - Bot Management feature is licensed.
* `apigateway` - API Gateway feature is licensed.
