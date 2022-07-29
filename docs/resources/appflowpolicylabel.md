---
subcategory: "AppFlow"
---

# Resource: appflowpolicylabel

The appflowpolicylabel resource is used to create appflowpolicylabel.


## Example usage

```hcl
resource "citrixadc_appflowpolicylabel" "tf_appflowpolicylabel" {
  labelname       = "tf_policylabel"
  policylabeltype = "OTHERTCP"
}
```


## Argument Reference

* `labelname` - (Optional) Name of the AppFlow policy label. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my appflow policylabel" or 'my appflow policylabel').
* `policylabeltype` - (Optional) Type of traffic evaluated by the policies bound to the policy label.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appflowpolicylabel. It has the same value as the `labelname` attribute.


## Import

An appflowpolicylabel can be imported using its name, e.g.

```shell
terraform import citrixadc_appflowpolicylabel.tf_appflowpolicylabel tf_policylabel
```
