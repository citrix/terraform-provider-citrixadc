---
subcategory: "Rewrite"
---

# Resource: rewritepolicylabel_rewritepolicy_binding

The rewritepolicylabel_rewritepolicy_binding resource is used to bind a rewrite policy to rewrite policy label.


## Example usage

```hcl
resource "citrixadc_rewritepolicylabel_rewritepolicy_binding" "tf_rewritepolicylabel_rewritepolicy_binding" {
	labelname = citrixadc_rewritepolicylabel.tf_rewritepolicylabel.labelname
	policyname = citrixadc_rewritepolicy.tf_rewrite_policy.name
	gotopriorityexpression = "END"
	priority = 5   
}

resource "citrixadc_rewritepolicylabel" "tf_rewritepolicylabel" {
	labelname = "tf_rewritepolicylabel"
	transform = "http_req"
}

resource "citrixadc_rewritepolicy" "tf_rewrite_policy" {
	name = "tf_rewrite_policy"
	action = "DROP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"
}
```


## Argument Reference

* `policyname` - (Required) Name of the rewrite policy to bind to the policy label.
* `priority` - (Required) Specifies the priority of the policy.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Suspend evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.
* `labeltype` - (Optional) Type of invocation. Available settings function as follows: * reqvserver - Forward the request to the specified request virtual server. * resvserver - Forward the response to the specified response virtual server. * policylabel - Invoke the specified policy label. Possible values: [ reqvserver, resvserver, policylabel ]
* `invoke_labelname` - (Optional) * If labelType is policylabel, name of the policy label to invoke.  * If labelType is reqvserver or resvserver, name of the virtual server to which to forward the request or response.
* `labelname` - (Required) Name of the rewrite policy label to which to bind the policy.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the rewritepolicylabel_rewritepolicy_binding. It is the concatenation of the `labelname` and `policyname` attributes separated by a comma.


## Import

A rewritepolicylabel_rewritepolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_rewritepolicylabel_rewritepolicy_binding.tf_rewritepolicylabel_rewritepolicy_binding tf_rewritepolicylabel,policyname
```
