---
subcategory: "Content Inspection"
---

# Data Source: contentinspectionpolicylabel

The contentinspectionpolicylabel data source allows you to retrieve information about a Content Inspection policy label.


## Example usage

```terraform
data "citrixadc_contentinspectionpolicylabel" "tf_contentinspectionpolicylabel" {
  labelname = "my_ci_policylabel"
}

output "type" {
  value = data.citrixadc_contentinspectionpolicylabel.tf_contentinspectionpolicylabel.type
}

output "comment" {
  value = data.citrixadc_contentinspectionpolicylabel.tf_contentinspectionpolicylabel.comment
}
```


## Argument Reference

* `labelname` - (Required) Name for the contentInspection policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `comment` - Any comments to preserve information about this contentInspection policy label.
* `type` - Type of packets (request or response packets) against which to match the policies bound to this policy label.
* `newname` - New name for the contentInspection policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
* `id` - The id of the contentinspectionpolicylabel. It has the same value as the `labelname` attribute.

### Read-only contentinspectionpolicylabel metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_contentinspectionpolicylabel` resource) and are Computed / GET-only. Any attribute the appliance does not return is `null`.

* `numpol` - Number of polices bound to label.
* `hits` - Number of times policy label was invoked.
* `priority` - Specifies the priority of the policy.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `labeltype` - Type of invocation (for example `reqvserver`, `resvserver`, `policylabel`).
* `invoke_labelname` - If labelType is policylabel, name of the policy label to invoke; if labelType is reqvserver or resvserver, name of the virtual server to which to forward the request or response.
* `flowtype` - Flowtype of the bound contentInspection policy.
* `isdefault` - A value of true is returned if it is a default cipolicylabel.
