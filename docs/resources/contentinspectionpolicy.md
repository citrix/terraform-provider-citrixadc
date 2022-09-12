---
subcategory: "CI"
---

# Resource: contentinspectionpolicy

The contentinspectionpolicy resource is used to create contentinspectionpolicy.


## Example usage

```hcl
resource "citrixadc_contentinspectionpolicy" "tf_contentinspectionpolicy" {
  name   = "my_ci_policy"
  rule   = "true"
  action = "RESET"
}

```


## Argument Reference

* `name` - (Required) Name for the contentInspection policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the contentInspection policy is added.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my contentInspection policy" or 'my contentInspection policy').
* `rule` - (Required) Expression that the policy uses to determine whether to execute the specified action.
* `action` - (Required) Name of the contentInspection action to perform if the request matches this contentInspection policy.     There are also some built-in actions which can be used. These are:     * NOINSPECTION - Send the request from the client to the server or response from the server to the client without sending it to Inspection device for Content Inspection.     * RESET - Resets the client connection by closing it. The client program, such as a browser, will handle this and may inform the user. The client may then resend the request if desired.     * DROP - Drop the request without sending a response to the user.
* `comment` - (Optional) Any type of information about this contentInspection policy.
* `logaction` - (Optional) Name of the messagelog action to use for requests that match this policy.
* `undefaction` - (Optional) Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Only the above built-in actions can be used.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the contentinspectionpolicy. It has the same value as the `name` attribute.


## Import

A contentinspectionpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_contentinspectionpolicy.tf_contentinspectionpolicy my_ci_policy
```
