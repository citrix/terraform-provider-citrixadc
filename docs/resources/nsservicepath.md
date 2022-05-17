---
subcategory: "NS"
---

# Resource: nsservicepath

The nsservicepath resource is used to create service Path resource.


## Example usage

```hcl
resource "citrixadc_nsservicepath" "tf_servicepath" {
  servicepathname = "tf_servicepath"  
}
```


## Argument Reference

* `servicepathname` - (Required) Name for the Service path. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsservicepath. It has the same value as the `servicepathname` attribute.


## Import

A nsservicepath can be imported using its servicepathname, e.g.

```shell
terraform import citrixadc_nsservicepath.tf_servicepath tf_servicepath
```
