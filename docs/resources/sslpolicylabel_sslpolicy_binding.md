---
subcategory: "SSL"
---

# Resource: sslpolicylabel_sslpolicy_binding

The sslpolicylabel_sslpolicy_binding resource is used to bind an SSL policy to an SSL policy label.


## Example usage

```hcl
resource "citrixadc_sslaction" "certinsertact" {
  name       = "certinsertact"
  clientcert = "ENABLED"
  certheader = "CERT"
}

resource "citrixadc_sslpolicy" "certinsert_pol" {
  name   = "certinsert_pol"
  rule   = "false"
  action = citrixadc_sslaction.certinsertact.name
}

resource "citrixadc_sslpolicylabel" "ssl_pol_label" {
  labelname = "ssl_pol_label"
  type      = "DATA"
}

resource "citrixadc_sslpolicylabel_sslpolicy_binding" "demo_binding" {
  labelname              = citrixadc_sslpolicylabel.ssl_pol_label.labelname
  policyname             = citrixadc_sslpolicy.certinsert_pol.name
  priority               = 56
  gotopriorityexpression = "END"
  labeltype              = "policylabel"
  invokelabelname        = "ssl_pol_label"
  invoke                 = true
}
```


## Argument Reference

* `labelname` - (Required) Name of the SSL policy label to which to bind policies.
* `policyname` - (Required) Name of the SSL policy to bind to the policy label.
* `priority` - (Required) Specifies the priority of the policy.
* `labeltype` - (Required) Type of policy label invocation. Possible values: [ vserver, service, policylabel ]
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invokelabelname` - (Optional) Name of the label to invoke if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke policies bound to a policy label. After the invoked policies are evaluated, the flow returns to the policy with the next priority.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslpolicylabel_sslpolicy_binding. It is the concatenation of the `labelname` and `policyname` attributes separated by a comma.


## Import

A sslpolicylabel_sslpolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_sslpolicylabel_sslpolicy_binding.demo_binding ssl_pol_label,certinsert_pol
```
