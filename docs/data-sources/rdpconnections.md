---
subcategory: "RDP"
---

# Data Source: rdpconnections

The rdpconnections data source retrieves information about active RDP proxy connections on the Citrix ADC, equivalent to `show rdp connections`. Use it to inspect the endpoint and target addressing of live RDP proxy sessions established through the NetScaler Gateway / VPN RDP proxy feature. An optional `username` filter narrows the lookup to a single user.

~> **Note** The RDP Proxy feature (NetScaler Gateway / VPN RDP proxy) must be in use for any RDP connections to exist. If there are no active RDP proxy sessions, the data source returns an empty result (no attributes populated) rather than an error.


## Example usage

### List active RDP proxy connections

```hcl
data "citrixadc_rdpconnections" "all" {}

output "rdp_endpoint_ip" {
  value = data.citrixadc_rdpconnections.all.endpointip
}
```

### Filter by user name

```hcl
data "citrixadc_rdpconnections" "user" {
  username = "jdoe"
}

output "rdp_target" {
  value = "${data.citrixadc_rdpconnections.user.targetip}:${data.citrixadc_rdpconnections.user.targetport}"
}
```


## Argument Reference

* `username` - (Optional) User name for which to display connections. When omitted, the first available active connection is returned.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier for the data source (`rdpconnections`).
* `endpointip` - IP address of the RDP connection endpoint (the client side).
* `endpointport` - Port of the RDP connection endpoint (1-65535).
* `targetip` - IP address of the RDP connection target (the backend RDP server).
* `targetport` - Port of the RDP connection target (1-65535).
* `peid` - Packet engine (core) ID handling the RDP connection.
