---
subcategory: "CI"
---

# Resource: contentinspectionparameter

The contentinspectionparameter resource is used to create contentinspectionparameter.


## Example usage

```hcl
resource "citrixadc_contentinspectionparameter" "tf_contentinspectionparameter" {
undefaction = "RESET"
}
```


## Argument Reference

* `undefaction` - (Optional) Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an error condition in evaluating the expression. Available settings function as follows: * NOINSPECTION - Do not Inspect the traffic. * RESET - Reset the connection and notify the user's browser, so that the user can resend the request. * DROP - Drop the message without sending a response to the user.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the contentinspectionparameter. It is a unique string prefixed with  `tf-contentinspectionparameter-`.