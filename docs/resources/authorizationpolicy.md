---
subcategory: "Authorization"
---

# Resource: authorizationpolicy

The `authorizationpolicy` resource is used to create Authorization Policy.

## Example usage

```hcl
resource "citrixadc_authorizationpolicy" "authorize1" {
  name   = "tp-authorize-1"
  rule   = "true"
  action = "ALLOW"
}

resource "citrixadc_authorizationpolicy" "authorize2" {
  name   = "tp-authorize-2"
  rule   = "true"
  action = "DENY"
}

```

## Argument Reference

* `name` - (Required) Name for the new authorization policy.  Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the authorization policy is added.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authorization policy" or 'my authorization policy').
* `rule` - (Optional) Name of the Citrix ADC named rule, or an expression, that the policy uses to perform the authentication.
* `action` - (Optional) Action to perform if the policy matches: either allow or deny the request.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the `authorizationpolicy`. It has the same value as the `name` attribute.
