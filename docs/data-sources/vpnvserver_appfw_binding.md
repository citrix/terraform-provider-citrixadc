---
subcategory: "VPN"
---

# Data Source `vpnvserver_appfwpolicy_binding`

The vpnvserver_appfwpolicy_binding data source allows you to retrieve information about the bound appfw policies to a vpnvserver

## Example Usage

```terraform
data "citrixadc_vpnvserver_appfwpolicy_binding" "vpnappfwbinding" {
    name = "tf_vpnvserver"
}

output "policy" {
  value = data.citrixadc_vpnvserver_appfwpolicy_binding.vpnappfwbinding.policy
}

output "priority" {
  value = data.citrixadc_vpnvserver_appfwpolicy_binding.vpnappfwbinding.priority
}
```

## Argument Reference

* `name` - (Required) Name of the vpnvserver.


## Attributes Reference

In addition to the argument, the following attributes are available:

* `policy` - The name of the appfw policy bound to the vpnvserver.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `groupextraction` - Binds the authentication policy to a tertiary chain which will be used only for group extraction.
* `priority` - Integer specifying the policy's priority. The lower the number, the higher the priority.
* `secondary` - Bind the authentication policy to the secondary chain.

## Import

A vpnvserver_appfwpolicy_binding can be imported using its name, e.g.

```
terraform import citrixadc_vpnvserver_appfwpolicy_binding.tf_bind tf_vpnvserver,tf_appfwpolicy
```

