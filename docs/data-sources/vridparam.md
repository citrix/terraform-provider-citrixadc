---
subcategory: "Network"
---

# Data Source: vridparam

The vridparam data source allows you to retrieve information about VRID (Virtual Router ID) global parameters.

## Example usage

```terraform
data "citrixadc_vridparam" "example" {
}

output "hellointerval" {
  value = data.citrixadc_vridparam.example.hellointerval
}

output "deadinterval" {
  value = data.citrixadc_vridparam.example.deadinterval
}
```

## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `deadinterval` - Number of seconds after which a peer node in active-active mode is marked down if vrrp advertisements are not received from the peer node.
* `hellointerval` - Interval, in milliseconds, between vrrp advertisement messages sent to the peer node in active-active mode.
* `sendtomaster` - Forward packets to the master node, in an active-active mode configuration, if the virtual server is in the backup state and sharing is disabled.
