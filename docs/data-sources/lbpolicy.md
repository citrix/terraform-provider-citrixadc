---
subcategory: "Load Balancing"
---

# Data Source `lbpolicy`

The lbpolicy data source allows you to retrieve information about a load balancing policy.


## Example usage

```terraform
data "citrixadc_lbpolicy" "tf_lbpolicy" {
  name = "my_lbpolicy"
}

output "rule" {
  value = data.citrixadc_lbpolicy.tf_lbpolicy.rule
}

output "action" {
  value = data.citrixadc_lbpolicy.tf_lbpolicy.action
}
```


## Argument Reference

* `name` - (Required) Name of the LB policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the LB policy is added.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbpolicy. It has the same value as the `name` attribute.
* `action` - Name of action to use if the request matches this LB policy.
* `comment` - Any type of information about this LB policy.
* `logaction` - Name of the messagelog action to use for requests that match this policy.
* `newname` - New name for the LB policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
* `rule` - Expression against which traffic is evaluated.
* `undefaction` - Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Available settings function as follows:
  * NOLBACTION - Does not consider LB actions in making LB decision.
  * RESET - Reset the request and notify the user, so that the user can resend the request.
  * DROP - Drop the request without sending a response to the user.


## Import

A lbpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_lbpolicy.tf_lbpolicy my_lbpolicy
```
