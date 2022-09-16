---
subcategory: "NS"
---

# Resource: nsweblogparam

The nsweblogparam resource is used to create Web log parameters resource.


## Example usage

```hcl
resource "citrixadc_nsweblogparam" "tf_nsweblofparam" {
  buffersizemb  = 32
  customreqhdrs = ["req1", "req2"]
  customrsphdrs = ["res1", "res2"]
}
```


## Argument Reference

* `buffersizemb` - (Optional) Buffer size, in MB, allocated for log transaction data on the system. The maximum value is limited to the memory available on the system. Minimum value =  1 Maximum value =  4294967294LU
* `customreqhdrs` - (Optional) Name(s) of HTTP request headers whose values should be exported by the Web Logging feature. Minimum length =  1
* `customrsphdrs` - (Optional) Name(s) of HTTP response headers whose values should be exported by the Web Logging feature. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsweblogparam. It is a unique string prefixed with "tf-nsweblogparam-"

