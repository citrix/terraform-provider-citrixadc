---
subcategory: "Authentication"
---

# Data Source: authenticationvserver_tmsessionpolicy_binding

The authenticationvserver_tmsessionpolicy_binding data source allows you to retrieve information about an authenticationvserver to tmsessionpolicy binding.


## Example Usage

```terraform
data "citrixadc_authenticationvserver_tmsessionpolicy_binding" "tf_bind" {
  name   = "tf_authenticationvserver"
  policy = "my_tmsession_policy"
}

output "priority" {
  value = data.citrixadc_authenticationvserver_tmsessionpolicy_binding.tf_bind.priority
}

output "nextfactor" {
  value = data.citrixadc_authenticationvserver_tmsessionpolicy_binding.tf_bind.nextfactor
}
```


## Argument Reference

* `name` - (Required) Name of the authentication virtual server to which to bind the policy.
* `policy` - (Required) The name of the policy, if any, bound to the authentication vserver.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `groupextraction` - Applicable only while bindind classic authentication policy as advance authentication policy use nFactor
* `id` - The id of the authenticationvserver_tmsessionpolicy_binding. It is the concatenation of the `name` and `policy` attributes separated by a comma.
* `nextfactor` - Applicable only while binding advance authentication policy as classic authentication policy does not support nFactor
* `priority` - The priority, if any, of the vpn vserver policy.
* `secondary` - Applicable only while bindind classic authentication policy as advance authentication policy use nFactor
