---
subcategory: "Authorization"
---

# Data Source: authorizationpolicy

The authorizationpolicy data source allows you to retrieve information about authorization policies.


## Example usage

```terraform
data "citrixadc_authorizationpolicy" "tf_authorizationpolicy" {
  name = "my_authorization_policy"
}

output "rule" {
  value = data.citrixadc_authorizationpolicy.tf_authorizationpolicy.rule
}

output "action" {
  value = data.citrixadc_authorizationpolicy.tf_authorizationpolicy.action
}
```


## Argument Reference

* `name` - (Required) Name for the new authorization policy. 
Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the authorization policy is added.

The following requirement applies only to the Citrix ADC CLI:
If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authorization policy" or 'my authorization policy').

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `action` - Action to perform if the policy matches: either allow or deny the request.
* `rule` - Name of the Citrix ADC named rule, or an expression, that the policy uses to perform the authentication.
* `newname` - The new name of the author policy.
* `id` - The id of the authorizationpolicy. It has the same value as the `name` attribute.

### Read-only authorizationpolicy metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_authorizationpolicy` resource). They are GET-only / Computed, and any attribute the appliance does not return is `null`.

* `activepolicy` - Indicates whether policy is bound or not.
* `expressiontype` - Type of policy (Classic/Advanced).
* `hits` - Number of hits.
