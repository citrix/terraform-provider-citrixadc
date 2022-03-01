---
subcategory: "High Availability"
---

# Resource: hafailover

The hafailover resource is used to trigger the high availability fail over operation.


## Example usage

```hcl
resource "citrixadc_hafailover" "tf_failover" {
  ipaddress = "10.222.74.152"
  state     = "Primary"
  force     = true
}
```


## Argument Reference

* `ipaddress` - (Required) IP address of the target High Availability node.
* `state` - (Required) Desired High Availability node state. Can be either `Primary` or `Secondary`.
* `force` - (Optional) Force a failover without prompting for confirmation.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the hafailover resource. It is a unique string prefixed with "tf-hafailover-"
