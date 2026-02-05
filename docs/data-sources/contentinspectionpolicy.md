---
subcategory: "Content Inspection"
---

# Data Source: citrixadc_contentinspectionpolicy

The `citrixadc_contentinspectionpolicy` data source is used to retrieve information about an existing Content Inspection Policy configured on a Citrix ADC appliance.

## Example usage

```hcl
# Retrieve a content inspection policy by name
data "citrixadc_contentinspectionpolicy" "example" {
  name = "my_ci_policy"
}

# Use the retrieved data in other resources
output "policy_rule" {
  value = data.citrixadc_contentinspectionpolicy.example.rule
}

output "policy_action" {
  value = data.citrixadc_contentinspectionpolicy.example.action
}

```

## Argument Reference

The following arguments are required:

* `name` - (Required) Name for the contentInspection policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the contentInspection policy is added.

## Attribute Reference

In addition to the arguments, the following attributes are exported:

* `id` - The ID of the content inspection policy. It has the same value as the `name` attribute.
* `action` - Name of the contentInspection action to perform if the request matches this contentInspection policy. There are also some built-in actions which can be used. These are:
    * NOINSPECTION - Send the request from the client to the server or response from the server to the client without sending it to Inspection device for Content Inspection.
    * RESET - Resets the client connection by closing it. The client program, such as a browser, will handle this and may inform the user. The client may then resend the request if desired.
    * DROP - Drop the request without sending a response to the user.
* `comment` - Any type of information about this contentInspection policy.
* `logaction` - Name of the messagelog action to use for requests that match this policy.
* `newname` - New name for the contentInspection policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
* `rule` - Expression that the policy uses to determine whether to execute the specified action.
* `undefaction` - Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Only the above built-in actions can be used.
