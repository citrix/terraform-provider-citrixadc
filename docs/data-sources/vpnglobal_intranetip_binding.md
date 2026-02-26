---
subcategory: "VPN"
---

# Data Source: vpnglobal_intranetip_binding

The vpnglobal_intranetip_binding data source allows you to retrieve information about a binding between vpnglobal configuration and an intranet IP address or range.

## Example Usage

```terraform
data "citrixadc_vpnglobal_intranetip_binding" "tf_bind" {
  intranetip = "2.3.4.5"
}

output "intranetip" {
  value = data.citrixadc_vpnglobal_intranetip_binding.tf_bind.intranetip
}

output "netmask" {
  value = data.citrixadc_vpnglobal_intranetip_binding.tf_bind.netmask
}
```

## Argument Reference

The following arguments are required:

* `intranetip` - (Required) The intranet ip address or range.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnglobal_intranetip_binding. It is a system-generated identifier.
* `netmask` - The intranet ip address or range's netmask.
* `gotopriorityexpression` - Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
