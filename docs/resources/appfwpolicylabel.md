---
subcategory: "Application Firewall"
---

# Resource: appfwpolicylabel

The `appfwpolicylabel` resource is used to create Application Firewall Policy Label resource.

## Example usage

``` hcl
resource "citrixadc_appfwpolicylabel" "demo_appfwpolicylabel" {
  labelname = "demo_appfwpolicylabel"
  policylabeltype = "http_req"
```

## Argument Reference

* `labelname` - Name for the policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the policy label is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my policy label" or 'my policy label').
* `policylabeltype` - (Optional) Type of transformations allowed by the policies bound to the label. Always http_req for application firewall policy labels. Possible values: [ http_req ]

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the `appfwpolicylabel`. It has the same value as the `labelname` attribute.
