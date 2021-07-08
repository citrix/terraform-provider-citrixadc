---
subcategory: "NS"
---

# Resource: nsmode

The nsmode resource is used to enable or disable ADC modes.


## Example usage

```hcl
resource "citrixadc_nsmode" "tf_nsmode" {
    usip = true
	cka = false
}
```


## Argument Reference

The following arguments can set to `true` or `false` to enable or disable the corresponding mode.

* `fr` - (Optional)
* `l2` - (Optional)
* `usip` - (Optional)
* `cka` - (Optional)
* `tcpb` - (Optional)
* `mbf` - (Optional)
* `edge` - (Optional)
* `usnip` - (Optional)
* `l3` - (Optional)
* `pmtud` - (Optional)
* `mediaclassification` - (Optional)
* `sradv` - (Optional)
* `dradv` - (Optional)
* `iradv` - (Optional)
* `sradv6` - (Optional)
* `dradv6` - (Optional)
* `bridgebpdus` - (Optional)
* `ulfd` - (Optional)


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsmode resource. It is a random string prefixed with "tf-nsmode-"
