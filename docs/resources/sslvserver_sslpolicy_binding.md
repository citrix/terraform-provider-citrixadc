---
subcategory: "SSL"
---

# Resource: sslvserver_sslpolicy_binding

The sslvserver_sslpolicy_binding resource is used to bind SSL policies to an SSL virtual server.


## Example usage

```hcl
resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.44"
  name        = "tf_lbvserver"
  port        = 443
  servicetype = "SSL"
  sslprofile  = "ns_default_ssl_profile_frontend"
}

resource "citrixadc_sslpolicy" "tf_sslpolicy" {
  name   = "tf_sslpolicy"
  rule   = "client.ssl.client_cert.exists"
  action = "NOOP"
}

resource "citrixadc_sslvserver_sslpolicy_binding" "tf_binding" {
  vservername = citrixadc_lbvserver.tf_lbvserver.name
  policyname  = citrixadc_sslpolicy.tf_sslpolicy.name
  priority    = 100
  type        = "REQUEST"
}
```


## Argument Reference

* `vservername` - (Required) Name of the SSL virtual server.
* `policyname` - (Required) The name of the SSL policy binding.
* `priority` - (Optional) The priority of the policies bound to this SSL service.
* `type` - (Optional) Bind point to which to bind the policy. REQUEST: Policy evaluation is done at the application above SSL; this bind point is the default and is used for actions based on clientauth and client cert. INTERCEPT_REQ: Policy evaluation is done during the SSL handshake to decide whether to intercept or not (allowed actions: INTERCEPT, BYPASS and RESET). CLIENTHELLO_REQ: Policy evaluation is done while handling the Client Hello Request (allowed actions: RESET, FORWARD and PICKCACERTGRP). Defaults to `"REQUEST"`.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke flag. This attribute is relevant only for ADVANCED policies.
* `labeltype` - (Optional) Type of policy label invocation.
* `labelname` - (Optional) Name of the label to invoke if the current policy rule evaluates to TRUE.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslvserver_sslpolicy_binding. It is the concatenation of the `policyname`, `priority`, `type` and `vservername` attributes separated by a comma.


## Import

A sslvserver_sslpolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_sslvserver_sslpolicy_binding.tf_binding tf_sslpolicy,100,REQUEST,tf_lbvserver
```
