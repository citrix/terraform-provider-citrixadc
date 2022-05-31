---
subcategory: "Network"
---

# Resource: appalgparam

The appalgparam resource is used to create AppAlg Param resource.


## Example usage

```hcl
resource "citrixadc_appalgparam" "tf_appalgparam" {
  pptpgreidletimeout = 9000
}
```


## Argument Reference

* `pptpgreidletimeout` - (Required) Interval in sec, after which data sessions of PPTP GRE is cleared. Minimum value =  1 Maximum value =  9000


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appalgparam. It is a unique string prefixed with "tf-appalgparam-"
