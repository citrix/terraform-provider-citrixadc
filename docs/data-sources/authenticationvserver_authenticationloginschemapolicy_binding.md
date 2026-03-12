---
subcategory: "Authentication"
---

# Data Source: authenticationvserver_authenticationloginschemapolicy_binding

The authenticationvserver_authenticationloginschemapolicy_binding data source allows you to retrieve information about a specific binding between an authentication virtual server and an authentication login schema policy.

## Example Usage

```terraform
data "citrixadc_authenticationvserver_authenticationloginschemapolicy_binding" "example" {
  name   = "tf_authenticationvserver"
  policy = "tf_loginschemapolicy"
}

output "priority" {
  value = data.citrixadc_authenticationvserver_authenticationloginschemapolicy_binding.example.priority
}

output "gotopriorityexpression" {
  value = data.citrixadc_authenticationvserver_authenticationloginschemapolicy_binding.example.gotopriorityexpression
}
```

## Argument Reference

* `name` - (Required) Name of the authentication virtual server to which to bind the policy.
* `policy` - (Required) The name of the policy, if any, bound to the authentication vserver.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `groupextraction` - Applicable only while bindind classic authentication policy as advance authentication policy use nFactor.
* `id` - The id of the authenticationvserver_authenticationloginschemapolicy_binding. It is a system-generated identifier.
* `nextfactor` - Applicable only while binding advance authentication policy as classic authentication policy does not support nFactor.
* `priority` - The priority, if any, of the vpn vserver policy.
* `secondary` - Applicable only while bindind classic authentication policy as advance authentication policy use nFactor.
