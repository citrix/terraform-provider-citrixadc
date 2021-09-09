---
subcategory: "SSL"
---

# Resource: sslpolicylabel_sslpolicy_binding

The sslpolicylabel_sslpolicy_bindingresource is used to create bindings between sslpolicylabel and sslpolicy.


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
	type = "DATA"	
}

resource "citrixadc_sslpolicylabel_sslpolicy_binding" "demo_sslpolicylabel_sslpolicy_binding" {
	gotopriorityexpression = "END"
	invoke = true
	labelname = citrixadc_sslpolicylabel.ssl_pol_label.labelname
	labeltype = "policylabel"
	policyname = citrixadc_sslpolicy.certinsert_pol.name
	priority = 56       
	invokelabelname = "ssl_pol_label"
}
```


## Argument Reference

* `policyname` - (Required) Name of the SSL policy to bind to the policy label.
* `priority` - (Required) Specifies the priority of the policy.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `labeltype` - (Optional) Type of policy label invocation. Possible values: [ vserver, service, policylabel ]
* `invoke_labelname` - (Optional) Name of the label to invoke if the current policy rule evaluates to TRUE.
* `labelname` - (Required) Name of the SSL policy label to which to bind policies.
* `invoke` - (Optional) Invoke policies bound to a policy label. After the invoked policies are evaluated, the flow returns to the policy with the next priority.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslpolicylabel_sslpolicy_binding. It has is the conatenation of the `labelname` and `policyname` attributes.

## Import

A sslpolicylabel_sslpolicy_bindingcan be imported using its id, e.g.

```shell
terraform import citrixadc_sslpolicylabel_sslpolicy_binding.tf_sslpolicylabel_sslpolicy_binding tf_sslpolicylabel_sslpolicy_binding
```
