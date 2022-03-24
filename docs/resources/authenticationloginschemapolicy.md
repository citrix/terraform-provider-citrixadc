---
subcategory: "Authentication"
---

# Resource: authenticationloginschemapolicy

The authenticationloginschemapolicy resource is used to create authentication loginschema policy resource.


## Example usage

```hcl
resource "citrixadc_authenticationloginschema" "tf_loginschema" {
  name                    = "tf_loginschema"
  authenticationschema    = "LoginSchema/SingleAuth.xml"
  ssocredentials          = "YES"
  authenticationstrength  = "30"
  passwordcredentialindex = "10"
}
resource "citrixadc_authenticationloginschemapolicy" "tf_loginschemapolicy" {
  name      = "tf_loginschemapolicy"
  rule      = "true"
  action    = citrixadc_authenticationloginschema.tf_loginschema.name
  comment   = "samplenew_testing"
}
```


## Argument Reference

* `name` - (Required) Name for the LoginSchema policy. This is used for selecting parameters for user logon. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the policy is created.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my policy" or 'my policy').
* `rule` - (Required) Expression which is evaluated to choose a profile for authentication.  The following requirements apply only to the Citrix ADC CLI: * If the expression includes one or more spaces, enclose the entire expression in double quotation marks. * If the expression itself includes double quotation marks, escape the quotations by using the \ character.  * Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
* `action` - (Required) Name of the profile to apply to requests or connections that match this policy. * NOOP - Do not take any specific action when this policy evaluates to true. This is useful to implicitly go to a different policy set. * RESET - Reset the client connection by closing it. The client program, such as a browser, will handle this and may inform the user. The client may then resend the request if desired. * DROP - Drop the request without sending a response to the user.
* `comment` - (Optional) Any comments to preserve information about this policy.
* `logaction` - (Optional) Name of messagelog action to use when a request matches this policy.
* `newname` - (Optional) New name for the LoginSchema policy.  Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my loginschemapolicy policy" or 'my loginschemapolicy policy').
* `undefaction` - (Optional) Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Only the above built-in actions can be used.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationloginschemapolicy. It has the same value as the `name` attribute.


## Import

A authenticationloginschemapolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy tf_loginschemapolicy
```
