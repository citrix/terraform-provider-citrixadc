---
subcategory: "SSL"
---

# Data Source: citrixadc_ssllogprofile

The ssllogprofile data source allows you to retrieve information about SSL log profiles configured on the Citrix ADC.

## Example usage

```terraform
data "citrixadc_ssllogprofile" "demo_ssllogprofile" {
  name = "demo_ssllogprofile"
}

output "ssllogclauth" {
  value = data.citrixadc_ssllogprofile.demo_ssllogprofile.ssllogclauth
}

output "sslloghs" {
  value = data.citrixadc_ssllogprofile.demo_ssllogprofile.sslloghs
}
```

## Argument Reference

* `name` - (Required) The name of the ssllogprofile.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the ssllogprofile. It has the same value as the `name` attribute.
* `ssllogclauth` - log all SSL ClAuth events. Possible values: [ ENABLED, DISABLED ]
* `ssllogclauthfailures` - log all SSL ClAuth error events. Possible values: [ ENABLED, DISABLED ]
* `sslloghs` - log all SSL HS events. Possible values: [ ENABLED, DISABLED ]
* `sslloghsfailures` - log all SSL HS error events. Possible values: [ ENABLED, DISABLED ]

## Import

A ssllogprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_ssllogprofile.demo_ssllogprofile demo_ssllogprofile
```
