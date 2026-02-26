---
subcategory: "Authentication"
---

# Data Source: authenticationvserver_authenticationpolicy_binding

The authenticationvserver_authenticationpolicy_binding data source allows you to retrieve information about an authenticationvserver to authenticationpolicy binding.


## Example Usage

```terraform
data "citrixadc_authenticationvserver_authenticationpolicy_binding" "tf_bind" {
  name   = "tf_authenticationvserver"
  policy = "tf_authenticationpolicy"
}

output "priority" {
  value = data.citrixadc_authenticationvserver_authenticationpolicy_binding.tf_bind.priority
}

output "nextfactor" {
  value = data.citrixadc_authenticationvserver_authenticationpolicy_binding.tf_bind.nextfactor
}
```


## Argument Reference

* `name` - (Required) Name of the authentication virtual server to which to bind the policy.
* `policy` - (Required) The name of the policy, if any, bound to the authentication vserver.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `groupextraction` - Applicable only while bindind classic authentication policy as advance authentication policy use nFactor
* `id` - The id of the authenticationvserver_authenticationpolicy_binding. It is the concatenation of the `name` and `policy` attributes separated by a comma.
* `nextfactor` - On success invoke label.
* `priority` - The priority, if any, of the vpn vserver policy.
* `secondary` - Applicable only while bindind classic authentication policy as advance authentication policy use nFactor
