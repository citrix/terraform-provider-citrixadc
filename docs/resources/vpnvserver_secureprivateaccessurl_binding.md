---
subcategory: "VPN"
---

# Resource: vpnvserver_secureprivateaccessurl_binding

Binds a Secure Private Access URL to a VPN (NetScaler Gateway) virtual server so that the application reachable at that URL is published through the gateway. Use this resource to associate a specific Secure Private Access application URL with an individual VPN vserver.


## Example usage

```hcl
resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "gatewayvserver1"
  servicetype = "SSL"
  ipv46       = "10.222.74.150"
  port        = 443
}

resource "citrixadc_vpnvserver_secureprivateaccessurl_binding" "tf_bind" {
  name                   = citrixadc_vpnvserver.tf_vpnvserver.name
  secureprivateaccessurl = "https://app.example.com/"
}
```


## Argument Reference

* `name` - (Required) Name of the VPN virtual server to which the Secure Private Access URL is bound. Changing this forces a new resource to be created.
* `secureprivateaccessurl` - (Required) The configured Secure Private Access URL to bind to the VPN virtual server. This is the literal URL string (the bound value, not a reference to another resource). Maximum length = 255. Changing this forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_secureprivateaccessurl_binding. It is a comma-separated list of `key:value` pairs in the form `name:<name>,secureprivateaccessurl:<secureprivateaccessurl>`, where each value is URL-encoded.


## Import

A vpnvserver_secureprivateaccessurl_binding can be imported using its id, which is a comma-separated list of `key:value` pairs (with URL-encoded values), e.g.

```shell
terraform import citrixadc_vpnvserver_secureprivateaccessurl_binding.tf_bind name:gatewayvserver1,secureprivateaccessurl:https%3A%2F%2Fapp.example.com%2F
```
