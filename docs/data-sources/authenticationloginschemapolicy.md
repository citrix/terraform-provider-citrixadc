---
subcategory: "Authentication"
---

# Data Source: authenticationloginschemapolicy

The authenticationloginschemapolicy data source allows you to retrieve information about authentication loginschema policies.


## Example usage

```terraform
data "citrixadc_authenticationloginschemapolicy" "tf_loginschemapolicy" {
  name = "my_loginschema_policy"
}

output "rule" {
  value = data.citrixadc_authenticationloginschemapolicy.tf_loginschemapolicy.rule
}
```


## Argument Reference

* `name` - (Required) Name for the LoginSchema policy. This is used for selecting parameters for user logon. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the policy is created.

The following requirement applies only to the Citrix ADC CLI:
If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my policy" or 'my policy').

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `action` - Name of the profile to apply to requests or connections that match this policy.
* NOOP - Do not take any specific action when this policy evaluates to true. This is useful to implicitly go to a different policy set.
* RESET - Reset the client connection by closing it. The client program, such as a browser, will handle this and may inform the user. The client may then resend the request if desired.
* DROP - Drop the request without sending a response to the user.
* `comment` - Any comments to preserve information about this policy.
* `logaction` - Name of messagelog action to use when a request matches this policy.
* `newname` - New name for the LoginSchema policy.
Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.

The following requirement applies only to the Citrix ADC CLI:
If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my loginschemapolicy policy" or 'my loginschemapolicy policy').
* `rule` - Expression which is evaluated to choose a profile for authentication.

The following requirements apply only to the Citrix ADC CLI:
* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.
* If the expression itself includes double quotation marks, escape the quotations by using the \ character.
* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
* `undefaction` - Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Only the above built-in actions can be used.
* `id` - The id of the authenticationloginschemapolicy. It has the same value as the `name` attribute.

### Read-only authenticationloginschemapolicy metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_authenticationloginschemapolicy` resource). They are GET-only/Computed and are `null` when the appliance omits them.

* `hits` - Number of hits.
* `undefhits` - Number of Undef hits.
* `builtin` - Flag to determine if policy is built-in or not (for example `MODIFIABLE`, `DELETABLE`, `IMMUTABLE`, `PARTITION_ALL`). A list of strings.
* `feature` - The feature to be checked while applying this config.
