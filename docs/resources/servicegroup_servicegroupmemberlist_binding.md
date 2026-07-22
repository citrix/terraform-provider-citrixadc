---
subcategory: "Basic"
---

# Resource: servicegroup\_servicegroupmemberlist\_binding

Manages the complete set of members of a Citrix ADC service group as a single atomic unit. Applying this resource replaces the entire member set of the named service group with the `members` you declare: any existing member that is not present in the list is removed (or disabled, depending on the graceful setting on the service group).

This bulk behavior is what distinguishes it from `servicegroup_servicegroupmember_binding`, which manages a single member at a time and leaves other members untouched. Use this resource when you want Terraform to own the full membership of a service group declaratively; use the single-member binding when you only want to add or remove individual members alongside members managed elsewhere.

~> **Full-membership replacement.** Because every attribute forces replacement, any change to `servicegroupname` or `members` recreates the binding, replacing the full member set of the service group.

~> **Autoscale service group required.** This bulk desired-state member-list operation is only permitted by NITRO on an autoscale-enabled service group. Applying it against a plain (non-autoscale) service group fails with NITRO `errorcode 257, "Operation not permitted"`. The parent `citrixadc_servicegroup` must be created with `autoscale = "API"` (or another autoscale mode supported by your appliance). To manage members of a plain service group one at a time, use `servicegroup_servicegroupmember_binding` instead.


## Example usage

```hcl
resource "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "tf_servicegroup"
  servicetype      = "HTTP"
  autoscale        = "API"
}

resource "citrixadc_servicegroup_servicegroupmemberlist_binding" "tf_binding" {
  servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname

  members = [
    {
      ip     = "10.78.22.33"
      port   = 80
      weight = 10
      state  = "ENABLED"
    },
    {
      ip     = "10.78.22.34"
      port   = 80
      weight = 5
      state  = "ENABLED"
    },
  ]
}
```


## Argument Reference

* `servicegroupname` - (Required) Name of the service group whose member list is managed. Changing this forces a new resource to be created.
* `members` - (Optional) The desired set of service group members. This set replaces the entire member list of the service group; any existing member not present in this set is deleted or disabled based on the graceful setting on the service group. This is a list of nested objects, each describing one member. Changing the member set forces a new resource to be created.

Each `members` object supports the following:

* `ip` - (Optional) IP address of the service group member.
* `port` - (Optional) Port number of the service to be enabled.
* `weight` - (Optional) Weight to assign to this service group member. Specifies the capacity of the member relative to the other members in the load balancing configuration.
* `state` - (Optional) Initial state of the member. Possible values: [ ENABLED, DISABLED ]
* `order` - (Optional) Order number to be assigned to the service group member.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the servicegroup\_servicegroupmemberlist\_binding. It has the same value as the `servicegroupname` attribute.
