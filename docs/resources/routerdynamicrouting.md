---
subcategory: "Network"
---

# Resource: routerdynamcirouting

The routerdynamcirouting resource is used to create dynamic routing entries.


## Example usage

```hcl
resource "citrixadc_routerdynamicrouting" "tf_dynamicrouting" {
    commandlines = [
        "router bgp 101",
        "neighbor 192.168.5.1 remote-as 100",
        "redistribute kernel",
    ]
}
```

~> The `commandlines` attribute can accept up to 711 characters. If your configuration exceeds this limit, please split it into multiple resources to ensure proper functionality.

## Argument Reference

* `commandlines` - (Optional) A list of commands to be executed.
* `nodeid` - (Optional) Unique number that identifies the cluster node.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the routerdynamcirouting. It is a random string prefixed with "tf-routerdynamicrouting-"
