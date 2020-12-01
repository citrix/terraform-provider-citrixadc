---
subcategory: "SSL"
---

# Resource: sslvserver_sslpolicy_binding

The sslvserver_sslpolicy_binding resource is used to create bindings between sslvservers and sslpolicies.


## Example usage

```hcl
resource "citrixadc_sslvserver_sslpolicy_binding" "tf_binding" {
    vservername = "tf_lbvserver"
    policyname = "tf_sslpolicy"
    priority = 100
    type = "REQUEST"
}
```


## Argument Reference

* `policyname` - (Required) The name of the SSL policy binding.
* `priority` - (Optional) The priority of the policies bound to this SSL service.
* `type` - (Optional) Bind point to which to bind the policy. Possible Values: REQUEST, INTERCEPT_REQ and CLIENTHELLO_REQ. These bindpoints mean: 1. REQUEST: Policy evaluation will be done at appplication above SSL. This bindpoint is default and is used for actions based on clientauth and client cert. 2. INTERCEPT_REQ: Policy evaluation will be done during SSL handshake to decide whether to intercept or not. Actions allowed with this type are: INTERCEPT, BYPASS and RESET. 3. CLIENTHELLO_REQ: Policy evaluation will be done during handling of Client Hello Request. Action allowed with this type is: RESET, FORWARD and PICKCACERTGRP. Possible values: [ INTERCEPT_REQ, REQUEST, CLIENTHELLO_REQ ]
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke flag. This attribute is relevant only for ADVANCED policies.
* `labeltype` - (Optional) Type of policy label invocation. Possible values: [ vserver, service, policylabel ]
* `labelname` - (Optional) Name of the label to invoke if the current policy rule evaluates to TRUE.
* `vservername` - (Required) Name of the SSL virtual server.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslvserver_sslpolicy_binding. It is the concatenation of the `vservername` and `policyname` attributes.


## Import

A sslvserver_sslpolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_sslvserver_sslpolicy_binding.tf_binding tf_lbvserver,tf_sslpolicy
```
