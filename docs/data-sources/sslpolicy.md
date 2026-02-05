# citrixadc_sslpolicy

This datasource is used to retrieve information about an existing SSL policy on the Citrix ADC.

## Example Usage

```hcl
resource "citrixadc_sslaction" "foo" {
  name                   = "tf_sslaction"
  clientauth             = "DOCLIENTAUTH"
  clientcertverification = "Mandatory"
}

resource "citrixadc_sslpolicy" "foo" {
  name   = "tf_sslpolicy"
  rule   = "false"
  action = citrixadc_sslaction.foo.name
}

data "citrixadc_sslpolicy" "foo" {
  name = citrixadc_sslpolicy.foo.name
}
```

## Argument Reference

* `name` - (Required) Name of the SSL policy to retrieve. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.

## Attribute Reference

In addition to the argument above, the following attributes are exported:

* `id` - The ID of the SSL policy.
* `action` - Name of the built-in or user-defined action to perform on the request. Available built-in actions are NOOP, RESET, DROP, CLIENTAUTH, NOCLIENTAUTH, INTERCEPT AND BYPASS.
* `comment` - Any comments associated with this policy.
* `reqaction` - The name of the action to be performed on the request. Refer to 'add ssl action' command to add a new action. Builtin actions like NOOP, RESET, DROP, CLIENTAUTH and NOCLIENTAUTH are also allowed.
* `rule` - Expression, against which traffic is evaluated.
* `undefaction` - Name of the action to be performed when the result of rule evaluation is undefined. Possible values for control policies: CLIENTAUTH, NOCLIENTAUTH, NOOP, RESET, DROP. Possible values for data policies: NOOP, RESET, DROP and BYPASS.
