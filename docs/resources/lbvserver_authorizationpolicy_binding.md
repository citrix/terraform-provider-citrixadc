---
subcategory: "Load Balancing"
---

# Resource: lbvserver_authorizationpolicy_binding

The lbvserver_authorizationpolicy_bindingresource is used to add an authorization policy to lbmonitor.


## Example usage

```hcl
resource "citrixadc_lbvserver_authorizationpolicy_binding" "tf_lbvserver_authorizationpolicy_binding" {
  name = "test_lbvserver"
  policyname = "tf_authorizationpolicy"
}
```


## Argument Reference

* `policyname` - (Required) Name of the policy bound to the LB vserver.
* `priority` - (Optional) Priority.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `bindpoint` - (Optional) The bindpoint to which the policy is bound. Possible values: [ REQUEST, RESPONSE ]
* `invoke` - (Optional) Invoke policies bound to a virtual server or policy label.
* `labeltype` - (Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]
* `labelname` - (Optional) Name of the label invoked.
* `name` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.  CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver'). .


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver_authorizationpolicy_binding. It has the same value as the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A lbvserver\_authorizationpolicy\_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_lbvserver_authorizationpolicy_binding.tf_lbvserver_authorizationpolicy_binding test_lbvserver,tf_authorizationpolicy
```