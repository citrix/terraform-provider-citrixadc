---
subcategory: "SSL"
---

# Resource: sslservice_sslpolicy_binding

Binds an SSL policy to an SSL service so that the policy's rule is evaluated for traffic handled by the service, letting you apply SSL actions (such as enforcing handshake controls or rejecting connections) on a per-service basis. Use the priority and goto-priority expression to control the order in which multiple bound policies are evaluated.


## Example usage

```hcl
resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_lbvserver"
  ipv46       = "10.10.10.44"
  port        = 443
  servicetype = "SSL"
}

resource "citrixadc_service" "tf_service" {
  name        = "tf_service"
  servicetype = "SSL"
  port        = 443
  ip          = "10.77.33.22"
  lbvserver   = citrixadc_lbvserver.tf_lbvserver.name
}

resource "citrixadc_sslpolicy" "tf_sslpolicy" {
  name   = "tf_sslpolicy"
  rule   = "client.ssl.version.ge(3)"
  action = "NOOP"
}

resource "citrixadc_sslservice_sslpolicy_binding" "tf_binding" {
  servicename = citrixadc_service.tf_service.name
  policyname  = citrixadc_sslpolicy.tf_sslpolicy.name
  priority    = 100
}
```


## Argument Reference

* `servicename` - (Required) Name of the SSL service for which to set advanced configuration. Changing this forces a new resource to be created.
* `policyname` - (Required) The SSL policy binding. Changing this forces a new resource to be created.
* `priority` - (Optional) The priority of the policies bound to this SSL service. Changing this forces a new resource to be created.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE. Changing this forces a new resource to be created.
* `invoke` - (Optional) Invoke flag. This attribute is relevant only for ADVANCED policies. Changing this forces a new resource to be created.
* `labelname` - (Optional) Name of the label to invoke if the current policy rule evaluates to TRUE. Changing this forces a new resource to be created.
* `labeltype` - (Optional) Type of policy label invocation. Changing this forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslservice_sslpolicy_binding. It is the concatenation of the `policyname`, `priority`, and `servicename` unique attributes, formatted as comma-separated `key:value` pairs (for example, `policyname:tf_sslpolicy,priority:100,servicename:tf_service`).


## Import

A sslservice_sslpolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_sslservice_sslpolicy_binding.tf_binding policyname:tf_sslpolicy,priority:100,servicename:tf_service
```
