---
subcategory: "AAA"
---

# Resource: aaapreauthenticationparameter

The aaapreauthenticationparameter resource is used to updateaaapreauthenticationparameter.


## Example usage

```hcl
resource "citrixadc_aaapreauthenticationparameter" "tf_aaapreauthenticationparameter" {
  preauthenticationaction = "DENY"
  deletefiles    = "/var/tmp/*.files"
}
```


## Argument Reference

* `preauthenticationaction` - (Optional) Deny or allow login on the basis of end point analysis results. Possible values: [ ALLOW, DENY ]
* `rule` - (Optional) Name of the Citrix ADC named rule, or an expression, to be evaluated by the EPA tool.
* `killprocess` - (Optional) String specifying the name of a process to be terminated by the EPA tool.
* `deletefiles` - (Optional) String specifying the path(s) to and name(s) of the files to be deleted by the EPA tool, as a string of between 1 and 1023 characters.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaapreauthenticationparameter. It is a unique string prefixed with `tf-aaapreauthenticationparameter-`.
