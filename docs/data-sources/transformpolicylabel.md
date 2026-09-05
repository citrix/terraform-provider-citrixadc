---
subcategory: "Transform"
---

# Data Source: transformpolicylabel

The transformpolicylabel data source allows you to retrieve information about a URL Transformation policy label.


## Example usage

```terraform
data "citrixadc_transformpolicylabel" "transformpolicylabel" {
  labelname = "label_1"
}

output "policylabeltype" {
  value = data.citrixadc_transformpolicylabel.transformpolicylabel.policylabeltype
}
```


## Argument Reference

* `labelname` - (Required) Name for the policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `newname` - New name for the policy label.
* `policylabeltype` - Types of transformations allowed by the policies bound to the label. For URL transformation, always http_req (HTTP Request).
* `id` - The id of the transformpolicylabel. It has the same value as the `labelname` attribute.

### Read-only transformpolicylabel metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_transformpolicylabel` resource). They are GET-only/Computed and are `null` when the appliance does not return them.

* `numpol` - Number of polices bound to label.
* `hits` - Number of times policy label was invoked.
* `priority` - Specifies the priority of the policy.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `labeltype` - Type of invocation (reqvserver, policylabel).
* `invoke_labelname` - Name of the policy label.
* `description` - Description of the policylabel.
