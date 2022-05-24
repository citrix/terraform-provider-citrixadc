---
subcategory: "Network"
---

# Resource: arpparam

The arpparam resource is used to create Global arp parameters resource.


## Example usage

```hcl
resource "citrixadc_arpparam" "tf_arpparam" {
  timeout         = 1200
  spoofvalidation = "DISABLED"
}
```


## Argument Reference

* `timeout` - (Optional) Time-out value (aging time) for the dynamically learned ARP entries, in seconds. The new value applies only to ARP entries that are dynamically learned after the new value is set. Previously existing ARP entries expire after the previously configured aging time. Minimum value =  5 Maximum value =  1200
* `spoofvalidation` - (Optional) enable/disable arp spoofing validation. Possible values: [ ENABLED, DISABLED ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the arpparam. It is a unique string prefixed with "tf-arpparam-"

