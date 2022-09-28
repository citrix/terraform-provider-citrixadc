---
subcategory: "Integrated Caching"
---

# Resource: cachepolicylabel

The cachepolicylabel resource is used to create cachepolicylabel.


## Example usage

```hcl
resource "citrixadc_cachepolicylabel" "tf_policylabel" {
    labelname = "my_cachepolicylabel"
    evaluates = "REQ"
}
```

## Argument Reference

* `evaluates` - (Required) When to evaluate policies bound to this label: request-time or response-time.
* `labelname` - (Required) Name for the label. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the label is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cachepolicylabel. It has the same value as the `labelname` attribute.

## Import

A cachepolicylabel can be imported using its name, e.g.

```shell
terraform import citrix_policylabel.tf_policylabel my_cachepolicylabel
```