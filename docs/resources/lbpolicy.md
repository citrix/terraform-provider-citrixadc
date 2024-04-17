---
subcategory: "Load Balancing"
---

# Resource: lbpolicy

The lbpolicy resource is used to configure lb policy.


## Example usage

```hcl
resource "citrixadc_lbaction" "tf_lbact" {
  name  = "tf_lbact"
  type  = "SELECTIONORDER"
  value = [1]
}

resource "citrixadc_lbpolicy" "tf_pol" {
  name   = "tf_pol"
  rule   = "true"
  action = citrixadc_lbaction.tf_lbact.name
}

```


## Argument Reference

* `name` - (Required) Name of the LB policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the LB policy is added. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my lb policy" or 'my lb policy').
* `rule` - (Required) Expression against which traffic is evaluated.
* `action` - (Required) Name of action to use if the request matches this LB policy.
* `comment` - (Optional) Any type of information about this LB policy.
* `logaction` - (Optional) Name of the messagelog action to use for requests that match this policy.
* `newname` - (Optional) New name for the LB policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my lb policy" or 'my lb policy').
* `undefaction` - (Optional) Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Available settings function as follows: * NOLBACTION - Does not consider LB actions in making LB decision. * RESET - Reset the request and notify the user, so that the user can resend the request. * DROP - Drop the request without sending a response to the user.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbpolicy resource. It has the same value as the `name` attribute.


## Import

A lbpolicy resource can be imported using its name, e.g.

```shell
terraform import citrixadc_lbpolicy.tf_pol tf_pol
```
