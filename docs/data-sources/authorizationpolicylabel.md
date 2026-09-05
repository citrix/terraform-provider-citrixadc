---
subcategory: "Authorization"
---

# Data Source: authorizationpolicylabel

The authorizationpolicylabel data source allows you to retrieve information about an existing authorization policy label.


## Example usage

```terraform
data "citrixadc_authorizationpolicylabel" "tf_authorizationpolicylabel" {
  labelname = "my_authorizationpolicylabel"
}

output "id" {
  value = data.citrixadc_authorizationpolicylabel.tf_authorizationpolicylabel.id
}

output "labelname" {
  value = data.citrixadc_authorizationpolicylabel.tf_authorizationpolicylabel.labelname
}
```


## Argument Reference

* `labelname` - (Required) Name for the new authorization policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the authorization policy is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authorizationpolicylabel. It has the same value as the `labelname` attribute.
* `newname` - The new name of the auth policy label.

### Read-only authorizationpolicylabel metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_authorizationpolicylabel` resource). They are GET-only / Computed, and any attribute the appliance does not return is `null`.

* `numpol` - Number of polices bound to label.
* `hits` - Number of times policy label was invoked.
* `policyname` - Name of the authorization policy to bind to the policy label.
* `priority` - Specifies the priority of the policy.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `labeltype` - Type of invocation (reqvserver, resvserver, policylabel).
* `invoke_labelname` - Name of the policy label to invoke if the current policy evaluates to TRUE, the invoke parameter is set, and Label Type is set to Policy Label.
* `flowtype` - Flowtype of the bound authorization policy.
* `description` - Description of the policylabel.
