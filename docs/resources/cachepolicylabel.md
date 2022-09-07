---
subcategory: "Cache"
---

# Resource: cachepolicylabel

The cachepolicylabel resource is used to create cachepolicylabel.


## Example usage

```hcl
resource "citrixadc_cachepolicylabel" "policylabel1" {
    labelname = "policylabel1"
    evaluates = "REQ"
}```


## Argument Reference

* `evaluates` - (Optional) When to evaluate policies bound to this label: request-time or response-time.
* `labelname` - (Optional) Name for the label. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the label is created.
* `newname` - (Optional) New name for the cache-policy label. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cachepolicylabel. It has the same value as the `name` attribute.


## Import

A cachepolicylabel can be imported using its name, e.g.

```shell
terraform import citrixadc_csaction.tf_csaction tf_csaction
```
