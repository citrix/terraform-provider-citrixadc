---
subcategory: "Load Balancing"
---

# Data Source: lbvserver_contentinspectionpolicy_binding

The lbvserver_contentinspectionpolicy_binding data source allows you to retrieve information about a content inspection policy binding to an lbvserver.


## Example Usage

```terraform
data "citrixadc_lbvserver_contentinspectionpolicy_binding" "tf_lbvserver_contentinspectionpolicy_binding" {
  name       = "tf_lbvserver"
  bindpoint  = "REQUEST"
  policyname = "tf_contentinspectionpolicy"
}

output "bindpoint" {
  value = data.citrixadc_lbvserver_contentinspectionpolicy_binding.tf_lbvserver_contentinspectionpolicy_binding.bindpoint
}

output "gotopriorityexpression" {
  value = data.citrixadc_lbvserver_contentinspectionpolicy_binding.tf_lbvserver_contentinspectionpolicy_binding.gotopriorityexpression
}
```


## Argument Reference

* `name` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created. CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver').
* `bindpoint` - (Required) The bindpoint to which the policy is bound.
* `policyname` - (Required) Name of the policy bound to the LB vserver.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - Invoke policies bound to a virtual server or policy label.
* `labelname` - Name of the label invoked.
* `priority` - Priority.
* `labeltype` - The invocation type.
* `order` - Integer specifying the order of the service. A larger number specifies a lower order. Defines the order of the service relative to the other services in the load balancing vserver's bindings. Determines the priority given to the service among all the services bound.
* `id` - The id of the lbvserver_contentinspectionpolicy_binding. It is a system-generated identifier.
