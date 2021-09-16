---
subcategory: "Rewrite"
---

# Resource: rewriteparam

The rewriteparam resource is used to configure the given rewrite parameters.


## Example usage

```hcl
resource "citrixadc_rewriteparam" "tf_rewriteparam" {
	timeout = 5
	undefaction = "RESET"
}
```


## Argument Reference

* `undefaction` - (Optional) Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an error condition in evaluating the expression. Available settings function as follows: * NOREWRITE - Do not modify the message. * RESET - Reset the connection and notify the user's browser, so that the user can resend the request. * DROP - Drop the message without sending a response to the user.
* `timeout` - (Optional) 

