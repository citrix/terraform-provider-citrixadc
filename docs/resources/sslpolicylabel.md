---
subcategory: "SSL"
---

# Resource: sslpolicylabel

The sslpolicylabel resource is used to configure SSL policy label resource.


## Example usage

```hcl
resource "citrixadc_sslpolicylabel" "demo_sslpolicylabel" {
    labelname = "demo_sslpolicylabel"
    type = "DATA"	
}
```


## Argument Reference

* `labelname` - (Required) Name for the SSL policy label.  Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the policy label is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my label" or 'my label').
* `type` - (Required) Type of policies that the policy label can contain. Possible values: [ CONTROL, DATA ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslpolicylabel. It has the same value as the `labelname` attribute.


## Import

A sslpolicylabel can be imported using its name, e.g.

```shell
terraform import citrixadc_sslpolicylabel.tf_sslpolicylabel tf_sslpolicylabel
```
