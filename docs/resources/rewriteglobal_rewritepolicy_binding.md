---
subcategory: "Rewrite"
---

# Resource: rewriteglobal_rewritepolicy_binding

The rewriteglobal_rewritepolicy_binding resource is used to bind a rewrite policy to global, in order to apply it to the entire traffic handled by the Citrix ADC.


## Example usage

```hcl
resource "citrixadc_rewriteglobal_rewritepolicy_binding" "tf_rewriteglobal_rewritepolicy_binding" {
	policyname = citrixadc_rewritepolicy.tf_rewrite_policy.name
	priority = 5
	type = "REQ_DEFAULT"
	globalbindtype = "SYSTEM_GLOBAL"
	gotopriorityexpression = "END"
	invoke = "true"
	labelname = citrixadc_rewritepolicylabel.tf_rewritepolicylabel.labelname
	labeltype = "policylabel"
}

resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
	name = "tf_rewrite_policy"
	action = "DROP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"
}

resource "citrixadc_rewritepolicylabel" "tf_rewritepolicylabel" {
	labelname = "tf_rewritepolicylabel"
	transform = "http_req"
}
```


## Argument Reference

* `policyname` - (Required) Name of the rewrite policy.
* `type` - (Required) The bindpoint to which to policy is bound. Possible values: [ REQ_OVERRIDE, REQ_DEFAULT, RES_OVERRIDE, RES_DEFAULT, OTHERTCP_REQ_OVERRIDE, OTHERTCP_REQ_DEFAULT, OTHERTCP_RES_OVERRIDE, OTHERTCP_RES_DEFAULT, SIPUDP_REQ_OVERRIDE, SIPUDP_REQ_DEFAULT, SIPUDP_RES_OVERRIDE, SIPUDP_RES_DEFAULT, SIPTCP_REQ_OVERRIDE, SIPTCP_REQ_DEFAULT, SIPTCP_RES_OVERRIDE, SIPTCP_RES_DEFAULT, DIAMETER_REQ_OVERRIDE, DIAMETER_REQ_DEFAULT, DIAMETER_RES_OVERRIDE, DIAMETER_RES_DEFAULT, RADIUS_REQ_OVERRIDE, RADIUS_REQ_DEFAULT, RADIUS_RES_OVERRIDE, RADIUS_RES_DEFAULT, DNS_REQ_OVERRIDE, DNS_REQ_DEFAULT, DNS_RES_OVERRIDE, DNS_RES_DEFAULT ]
* `priority` - (Required) Specifies the priority of the policy.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.
* `labeltype` - (Optional) Type of invocation. Available settings function as follows: * reqvserver - Forward the request to the specified request virtual server. * resvserver - Forward the response to the specified response virtual server. * policylabel - Invoke the specified policy label. Possible values: [ reqvserver, resvserver, policylabel ]
* `labelname` - (Optional) * If labelType is policylabel, name of the policy label to invoke.  * If labelType is reqvserver or resvserver, name of the virtual server to which to forward the request of response.
* `globalbindtype` - (Optional) . Possible values: [ SYSTEM_GLOBAL, VPN_GLOBAL, RNAT_GLOBAL ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the rewriteglobal_rewritepolicy_binding. It is the concatenation of the `policyname`, `priority` and `type` attributes separated by a comma.


## Import

A rewriteglobal_rewritepolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_rewriteglobal_rewritepolicy_binding.tf_rewriteglobal_rewritepolicy_binding tf_rewrite_policy,5,REQ_DEFAULT
```
