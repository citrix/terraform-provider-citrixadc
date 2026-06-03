---
subcategory: "VPN"
---

# Data Source: vpnvserver_secureprivateaccessurl_binding

The vpnvserver_secureprivateaccessurl_binding data source allows you to retrieve information about a Secure Private Access URL bound to a VPN (NetScaler Gateway) virtual server.


## Example usage

```terraform
data "citrixadc_vpnvserver_secureprivateaccessurl_binding" "tf_bind" {
  name                   = "gatewayvserver1"
  secureprivateaccessurl = "https://app.example.com/"
}

output "secureprivateaccessurl" {
  value = data.citrixadc_vpnvserver_secureprivateaccessurl_binding.tf_bind.secureprivateaccessurl
}
```


## Argument Reference

* `name` - (Required) Name of the VPN virtual server.
* `secureprivateaccessurl` - (Required) The configured Secure Private Access URL bound to the VPN virtual server.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_secureprivateaccessurl_binding. It is a comma-separated list of `key:value` pairs in the form `name:<name>,secureprivateaccessurl:<secureprivateaccessurl>`, where each value is URL-encoded.
