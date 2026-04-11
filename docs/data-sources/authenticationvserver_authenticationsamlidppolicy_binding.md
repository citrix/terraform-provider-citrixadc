---
subcategory: "Authentication"
---

# Data Source: authenticationvserver_authenticationsamlidppolicy_binding

The authenticationvserver_authenticationsamlidppolicy_binding data source allows you to retrieve information about the binding between an authentication virtual server and an authentication SAML IDP policy.

## Example Usage

```terraform
data "citrixadc_authenticationvserver_authenticationsamlidppolicy_binding" "tf_bind" {
  name   = "tf_authenticationvserver"
  policy = "tf_samlidppolicy"
}

output "priority" {
  value = data.citrixadc_authenticationvserver_authenticationsamlidppolicy_binding.tf_bind.priority
}

output "secondary" {
  value = data.citrixadc_authenticationvserver_authenticationsamlidppolicy_binding.tf_bind.secondary
}
```

## Argument Reference

* `name` - (Required) Name of the authentication virtual server to which to bind the policy.
* `policy` - (Required) The name of the policy, if any, bound to the authentication vserver.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `groupextraction` - Applicable only while bindind classic authentication policy as advance authentication policy use nFactor
* `nextfactor` - On success invoke label.
* `priority` - The priority, if any, of the vpn vserver policy.
* `secondary` - Applicable only while bindind classic authentication policy as advance authentication policy use nFactor
* `id` - The id of the authenticationvserver_authenticationsamlidppolicy_binding. It is a system-generated identifier.
