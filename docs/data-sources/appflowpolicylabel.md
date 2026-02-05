---
subcategory: "AppFlow"
---

# Data Source `appflowpolicylabel`

The appflowpolicylabel data source allows you to retrieve information about an existing appflowpolicylabel.


## Example usage

```terraform
data "citrixadc_appflowpolicylabel" "tf_appflowpolicylabel" {
  labelname = "tf_policylabel"
}

output "labelname" {
  value = data.citrixadc_appflowpolicylabel.tf_appflowpolicylabel.labelname
}

output "policylabeltype" {
  value = data.citrixadc_appflowpolicylabel.tf_appflowpolicylabel.policylabeltype
}
```


## Argument Reference

* `labelname` - (Required) Name of the AppFlow policy label. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.

The following requirement applies only to the Citrix ADC CLI:
If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my appflow policylabel" or 'my appflow policylabel').

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appflowpolicylabel. It has the same value as the `labelname` attribute.
* `newname` - New name for the policy label. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
* `policylabeltype` - Type of traffic evaluated by the policies bound to the policy label.
