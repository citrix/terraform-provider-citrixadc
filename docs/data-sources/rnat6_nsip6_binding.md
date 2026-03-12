---
subcategory: "Network"
---

# Data Source: rnat6_nsip6_binding

The rnat6_nsip6_binding data source allows you to retrieve information about an IPv6 NAT IP address binding to an RNAT6 rule.

## Example Usage

```terraform
data "citrixadc_rnat6_nsip6_binding" "tf_rnat6_nsip6_binding" {
  name   = "my_rnat6"
  natip6 = "2001:db8:85a3::8a2e:370:7334"
}

output "name" {
  value = data.citrixadc_rnat6_nsip6_binding.tf_rnat6_nsip6_binding.name
}

output "natip6" {
  value = data.citrixadc_rnat6_nsip6_binding.tf_rnat6_nsip6_binding.natip6
}

output "ownergroup" {
  value = data.citrixadc_rnat6_nsip6_binding.tf_rnat6_nsip6_binding.ownergroup
}
```

## Argument Reference

* `name` - (Required) Name of the RNAT6 rule to which to bind NAT IPs.
* `natip6` - (Required) Nat IP Address.
* `ownergroup` - (Optional) The owner node group in a Cluster for this rnat rule.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the rnat6_nsip6_binding. It is a system-generated identifier.
