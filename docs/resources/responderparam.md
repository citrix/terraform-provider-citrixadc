---
subcategory: "Responder"
---

# Resource: responderparam

The responderparam resource is used to configure the given responder parameters.


## Example usage

```hcl
resource "citrixadc_responderparam" "tf_responderparam" {
	timeout = 5
	undefaction = "RESET"
}
```


## Argument Reference

* `undefaction` - (Optional) Action to perform when policy evaluation creates an UNDEF condition. Available settings function as follows: * NOOP - Send the request to the protected server. * RESET - Reset the request and notify the user's browser, so that the user can resend the request. * DROP - Drop the request without sending a response to the user.
* `timeout` - (Optional) 

