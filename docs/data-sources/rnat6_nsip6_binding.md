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
* `ownergroup` - (Optional) The owner node group in a Cluster for this rnat rule. Defaults to `"DEFAULT_NG"`.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the rnat6_nsip6_binding. It is a system-generated identifier.

### Read-only rnat6_nsip6_binding metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_rnat6_nsip6_binding` resource). They are GET-only/Computed. Any attribute the appliance does not return is `null`.

* `td` - Integer value that uniquely identifies the traffic domain in which the entity is configured. Omitted (null) for the default traffic domain (ID 0).
