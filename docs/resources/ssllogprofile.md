---
subcategory: "SSL"
---

# Resource: ssllogprofile

The ssllogprofile resource is used to create the ADC SSL logprofile.


## Example usage

```hcl
resource "citrixadc_ssllogprofile" "foo" {
    name = "foo"
    ssllogclauth = "DISABLED"
    ssllogclauthfailures = "ENABLED"
    sslloghs = "ENABLED"
    sslloghsfailures = "ENABLED"	
}
```


## Argument Reference

* `name` - (Required) The name of the ssllogprofile.
* `ssllogclauth` - (Optional) log all SSL ClAuth events. Possible values: [ ENABLED, DISABLED ]
* `ssllogclauthfailures` - (Optional) log all SSL ClAuth error events. Possible values: [ ENABLED, DISABLED ]
* `sslloghs` - (Optional) log all SSL HS events. Possible values: [ ENABLED, DISABLED ]
* `sslloghsfailures` - (Optional) log all SSL HS error events. Possible values: [ ENABLED, DISABLED ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the ssllogprofile. It has the same value as the `name` attribute.


## Import

A ssllogprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_ssllogprofile.tf_ssllogprofile tf_ssllogprofile
```
