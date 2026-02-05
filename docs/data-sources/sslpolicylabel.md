---
subcategory: "SSL"
---

# Data Source `sslpolicylabel`

The sslpolicylabel data source allows you to retrieve information about an existing sslpolicylabel.


## Example usage

```terraform
data "citrixadc_sslpolicylabel" "tf_sslpolicylabel" {
  labelname = "tf_sslpolicylabel"
}

output "labelname" {
  value = data.citrixadc_sslpolicylabel.tf_sslpolicylabel.labelname
}

output "type" {
  value = data.citrixadc_sslpolicylabel.tf_sslpolicylabel.type
}
```


## Argument Reference

* `labelname` - (Required) Name for the SSL policy label. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the policy label is created.

The following requirement applies only to the Citrix ADC CLI:
If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my label" or 'my label').

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslpolicylabel. It has the same value as the `labelname` attribute.
* `type` - Type of policies that the policy label can contain.
