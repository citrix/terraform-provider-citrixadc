---
subcategory: "Basic"
---

# Resource: servicegroup\_servicegroupmemberlist\_binding

This resource is used to manage the complete member set of a service group as a single unit on the Citrix ADC.

~> **Full-membership replacement:** Applying this resource replaces the entire member set and every attribute forces replacement. Requires an autoscale-enabled service group (NITRO `errorcode 257` otherwise).


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
