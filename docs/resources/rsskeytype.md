---
subcategory: "Network"
---

# Resource: rsskeytype

The rsskeytype resource is used to create RSS key type resource.


## Example usage

```hcl
resource "citrixadc_rsskeytype" "tf_rsskeytype" {
  rsstype = "ASYMMETRIC"
}
```


## Argument Reference

* `rsstype` - (Required) Type of RSS key, possible values are SYMMETRIC and ASYMMETRIC. Possible values: [ ASYMMETRIC, SYMMETRIC ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the rsskeytype. It is a unique string prefixed with "tf-rsskeytype-"

