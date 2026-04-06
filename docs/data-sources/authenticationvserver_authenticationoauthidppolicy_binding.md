---
subcategory: "Authentication"
---

# Data Source: authenticationvserver_authenticationoauthidppolicy_binding

The authenticationvserver_authenticationoauthidppolicy_binding data source allows you to retrieve information about a specific binding between an authentication virtual server and an authentication OAuth IDP policy.

## Example Usage

```terraform
data "citrixadc_authenticationvserver_authenticationoauthidppolicy_binding" "example" {
  name   = "tf_authenticationvserver"
  policy = "tf_idppolicy"
}

output "priority" {
  value = data.citrixadc_authenticationvserver_authenticationoauthidppolicy_binding.example.priority
}

output "nextfactor" {
  value = data.citrixadc_authenticationvserver_authenticationoauthidppolicy_binding.example.nextfactor
}
```

## Argument Reference

* `name` - (Required) Name of the authentication virtual server to which to bind the policy.
* `policy` - (Required) The name of the policy, if any, bound to the authentication vserver.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `groupextraction` - Applicable only while bindind classic authentication policy as advance authentication policy use nFactor.
* `id` - The id of the authenticationvserver_authenticationoauthidppolicy_binding. It is a system-generated identifier.
* `nextfactor` - On success invoke label.
* `priority` - The priority, if any, of the vpn vserver policy.
* `secondary` - Applicable only while bindind classic authentication policy as advance authentication policy use nFactor.
