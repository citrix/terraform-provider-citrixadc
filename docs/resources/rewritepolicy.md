---
subcategory: "Rewrite"
---

# Resource: rewritepolicy

The rewritepolicy resource is used to create rewrite policies.


## Example usage

```hcl
resource "citrixadc_rewritepolicy" "tf_rewritepolicy" {
	name = "tf_rewritepolicy"
	action = "DROP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"helloandby\")"

	globalbinding {
		gotopriorityexpression = "END"
		labelname = citrixadc_lbvserver.tf_lbvserver.name
		labeltype = "reqvserver"
		priority = 205
		invoke = true
		type = "REQ_DEFAULT"
	}
}
```


## Argument Reference

* `name` - (Optional) Name for the rewrite policy.
* `rule` - (Optional) Expression against which traffic is evaluated. The following requirements apply only to the Citrix ADC CLI: * If the expression includes one or more spaces, enclose the entire expression in double quotation marks. * If the expression itself includes double quotation marks, escape the quotations by using the \ character. * Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
* `action` - (Optional) Name of the rewrite action to perform if the request or response matches this rewrite policy. There are also some built-in actions which can be used. These are: * NOREWRITE - Send the request from the client to the server or response from the server to the client without making any changes in the message. * RESET - Resets the client connection by closing it. The client program, such as a browser, will handle this and may inform the user. The client may then resend the request if desired. * DROP - Drop the request without sending a response to the user.
* `undefaction` - (Optional) Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Only the above built-in actions can be used.
* `comment` - (Optional) Any comments to preserve information about this rewrite policy.
* `logaction` - (Optional) Name of messagelog action to use when a request matches this policy.
* `globalbinding` - (Optional) A global binding block, documented below.
* `lbvserverbinding` - (Optional) A lbvserver binding block, documented below.
* `csvserverbinding` - (Optional) A csvserver binding block, documented below.

A global binding supports the following:

* `type` - (Optional) The bindpoint to which to policy is bound. Possible values: [ REQ_OVERRIDE, REQ_DEFAULT, RES_OVERRIDE, RES_DEFAULT, OTHERTCP_REQ_OVERRIDE, OTHERTCP_REQ_DEFAULT, OTHERTCP_RES_OVERRIDE, OTHERTCP_RES_DEFAULT, SIPUDP_REQ_OVERRIDE, SIPUDP_REQ_DEFAULT, SIPUDP_RES_OVERRIDE, SIPUDP_RES_DEFAULT, SIPTCP_REQ_OVERRIDE, SIPTCP_REQ_DEFAULT, SIPTCP_RES_OVERRIDE, SIPTCP_RES_DEFAULT, DIAMETER_REQ_OVERRIDE, DIAMETER_REQ_DEFAULT, DIAMETER_RES_OVERRIDE, DIAMETER_RES_DEFAULT, RADIUS_REQ_OVERRIDE, RADIUS_REQ_DEFAULT, RADIUS_RES_OVERRIDE, RADIUS_RES_DEFAULT, DNS_REQ_OVERRIDE, DNS_REQ_DEFAULT, DNS_RES_OVERRIDE, DNS_RES_DEFAULT ]
* `priority` - (Optional) Specifies the priority of the policy.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.
* `labeltype` - (Optional) Type of invocation. Available settings function as follows: * reqvserver - Forward the request to the specified request virtual server. * resvserver - Forward the response to the specified response virtual server. * policylabel - Invoke the specified policy label. Possible values: [ reqvserver, resvserver, policylabel ]
* `labelname` - (Optional) * If labelType is policylabel, name of the policy label to invoke.  * If labelType is reqvserver or resvserver, name of the virtual server to which to forward the request of response.
* `globalbindtype` - (Optional) . Possible values: [ SYSTEM_GLOBAL, VPN_GLOBAL, RNAT_GLOBAL ]

A lbvserver binding supports the following:

* `priority` - (Optional) Priority.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `bindpoint` - (Optional) The bindpoint to which the policy is bound. Possible values: [ REQUEST, RESPONSE ]
* `invoke` - (Optional) Invoke policies bound to a virtual server or policy label.
* `labeltype` - (Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]
* `labelname` - (Optional) Name of the label invoked.
* `name` - (Optional) Name for the virtual server.

A csvserver binding supports the following:

* `priority` - (Optional) Priority for the policy.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `bindpoint` - (Optional) The bindpoint to which the policy is bound. Possible values: [ REQUEST, RESPONSE, ICA_REQUEST, OTHERTCP_REQUEST ]
* `invoke` - (Optional) Invoke flag.
* `labeltype` - (Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]
* `labelname` - (Optional) Name of the label invoked.
* `name` - (Optional) Name of the content switching virtual server to which the content switching policy applies.
* `targetlbvserver` - (Optional) Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE. Example: bind cs vs cs1 -policyname pol1 -priority 101 -targetLBVserver lb1 Note: Use this parameter only in case of Content Switching policy bind operations to a CS vserver.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the rewritepolicy. It has the same value as the `name` attribute.


## Import

A rewritepolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_rewritepolicy.tf_rewritepolicy tf_rewritepolicy
```
