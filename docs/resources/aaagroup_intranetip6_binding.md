---
subcategory: "AAA"
---

# Resource: aaagroup_intranetip6_binding

Assigns a range of intranet IPv6 addresses to an AAA group on the Citrix ADC. Members of the group are allocated addresses from this range when they establish a VPN session, allowing the ADC to hand out predictable, group-scoped intranet IPv6 connectivity instead of per-user assignments.


## Example usage

```hcl
resource "citrixadc_aaagroup" "tf_aaagroup" {
  groupname = "my_group"
}

resource "citrixadc_aaagroup_intranetip6_binding" "tf_aaagroup_intranetip6_binding" {
  groupname   = citrixadc_aaagroup.tf_aaagroup.groupname
  intranetip6 = "2001:db8::1"
  numaddr     = 1
}
```


## Argument Reference

* `groupname` - (Required) Name of the group that you are binding. Changing this forces a new resource to be created.
* `intranetip6` - (Required) The intranet IPv6 address (start of the range) bound to the group. Changing this forces a new resource to be created.
* `numaddr` - (Required) Number of IPv6 addresses bound, starting with `intranetip6`. Changing this forces a new resource to be created.

~> **Note** This resource has no NITRO update endpoint and every attribute forces replacement. Any change to `groupname`, `intranetip6`, or `numaddr` recreates the binding.

~> **Note** `gotopriorityexpression` is not applicable to the intranet IPv6 branch of `bind aaa group` and is not accepted by the ADC on this binding; it is therefore not exposed as a configurable argument here.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaagroup_intranetip6_binding. It is a comma-separated set of `key:value` pairs in the form `groupname:<groupname>,intranetip6:<intranetip6>,numaddr:<numaddr>`. The `intranetip6` value is URL-encoded inside the id so that the IPv6 colons (`:`) become `%3A` and do not collide with the `key:value` and comma delimiters.


## Import

A aaagroup_intranetip6_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_aaagroup_intranetip6_binding.tf_aaagroup_intranetip6_binding "groupname:my_group,intranetip6:2001%3Adb8%3A%3A1,numaddr:1"
```

Note that the IPv6 colons are URL-encoded (`:` becomes `%3A`) in the id used for import, exactly as the provider stores it in state.
