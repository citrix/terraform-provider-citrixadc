---
subcategory: "CMP"
---

# Resource: cmppolicy

The cmppolicy resource is used to create compression policies.


## Example usage

```hcl
resource "citrixadc_cmppolicy" "tf_cmppolicy" {
    name = "tf_cmppolicy"
    rule = "HTTP.RES.HEADER(\"Content-Type\").CONTAINS(\"text\")"
    resaction = "COMPRESS"
}
```


## Argument Reference

* `name` - (Optional) Name of the HTTP compression policy. Must begin with an ASCII alphabetic or underscore (\_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the policy is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my cmp policy" or 'my cmp policy').
* `rule` - (Optional) Expression that determines which HTTP requests or responses match the compression policy. The following requirements apply only to the Citrix ADC CLI: * If the expression includes one or more spaces, enclose the entire expression in double quotation marks. * If the expression itself includes double quotation marks, escape the quotations by using the \ character. * Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
* `resaction` - (Optional) The built-in or user-defined compression action to apply to the response when the policy matches a request or response.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the compression policy. It has the same value as the `name` attribute.


## Import

A cmppolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_cmppolicy.tf_cmppolicy tf_cmppolicy
```
