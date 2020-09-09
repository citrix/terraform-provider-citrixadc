---
subcategory: "Responder"
---

# Resource: responderpolicy

The responderpolicy resource is used to create responder policies.


## Example usage

```hcl
resource "citrixadc_responderpolicy" "tf_responder_policy" {
  name    = "tf_responder_policy"
  action = "NOOP"
  rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"nosuchthing\")"

  globalbinding {
      invoke = true
      labeltype = "vserver"
      labelname = citrixadc_lbvserver.tf_lbvserver.name
      type = "REQ_OVERRIDE"
      gotopriorityexpression = "END"
      priority = 100
  }
}
```


## Argument Reference

* `name` - (Optional) Name for the responder policy.
* `rule` - (Optional) Expression that the policy uses to determine whether to respond to the specified request.
* `action` - (Optional) Name of the responder action to perform if the request matches this responder policy. There are also some built-in actions which can be used. These are: * NOOP - Send the request to the protected server instead of responding to it. * RESET - Reset the client connection by closing it. The client program, such as a browser, will handle this and may inform the user. The client may then resend the request if desired. * DROP - Drop the request without sending a response to the user.
* `undefaction` - (Optional) Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Only the above built-in actions can be used.
* `comment` - (Optional) Any type of information about this responder policy.
* `logaction` - (Optional) Name of the messagelog action to use for requests that match this policy.
* `appflowaction` - (Optional) AppFlow action to invoke for requests that match this policy.
* `globalbinding` - (Optional) A global binding block. Documented below.
* `lbvserverbinding` - (Optional) A lbvserver binding block. Documented below.
* `csvserverbinding` - (Optional) A csvserver binding block. Documented below.


A global binding block supports the following:

* `type` - (Optional) Specifies the bind point whose policies you want to display. Available settings function as follows: * REQ_OVERRIDE - Request override. Binds the policy to the priority request queue. * REQ_DEFAULT - Binds the policy to the default request queue. * OTHERTCP_REQ_OVERRIDE - Binds the policy to the non-HTTP TCP priority request queue. * OTHERTCP_REQ_DEFAULT - Binds the policy to the non-HTTP TCP default request queue.. * SIPUDP_REQ_OVERRIDE - Binds the policy to the SIP UDP priority response queue.. * SIPUDP_REQ_DEFAULT - Binds the policy to the SIP UDP default response queue. * RADIUS_REQ_OVERRIDE - Binds the policy to the RADIUS priority response queue.. * RADIUS_REQ_DEFAULT - Binds the policy to the RADIUS default response queue. * MSSQL_REQ_OVERRIDE - Binds the policy to the Microsoft SQL priority response queue.. * MSSQL_REQ_DEFAULT - Binds the policy to the Microsoft SQL default response queue. * MYSQL_REQ_OVERRIDE - Binds the policy to the MySQL priority response queue. * MYSQL_REQ_DEFAULT - Binds the policy to the MySQL default response queue. Possible values: [ REQ_OVERRIDE, REQ_DEFAULT, OVERRIDE, DEFAULT, OTHERTCP_REQ_OVERRIDE, OTHERTCP_REQ_DEFAULT, SIPUDP_REQ_OVERRIDE, SIPUDP_REQ_DEFAULT, SIPTCP_REQ_OVERRIDE, SIPTCP_REQ_DEFAULT, MSSQL_REQ_OVERRIDE, MSSQL_REQ_DEFAULT, MYSQL_REQ_OVERRIDE, MYSQL_REQ_DEFAULT, NAT_REQ_OVERRIDE, NAT_REQ_DEFAULT, DIAMETER_REQ_OVERRIDE, DIAMETER_REQ_DEFAULT, RADIUS_REQ_OVERRIDE, RADIUS_REQ_DEFAULT, DNS_REQ_OVERRIDE, DNS_REQ_DEFAULT ]
* `priority` - (Optional) Specifies the priority of the policy.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.
* `labeltype` - (Optional) Type of invocation, Available settings function as follows: * vserver - Forward the request to the specified virtual server. * policylabel - Invoke the specified policy label. Possible values: [ vserver, policylabel ]
* `labelname` - (Optional) Name of the policy label to invoke. If the current policy evaluates to TRUE, the invoke parameter is set, and Label Type is policylabel.
* `globalbindtype` - (Optional) . Possible values: [ SYSTEM_GLOBAL, VPN_GLOBAL, RNAT_GLOBAL ]

A lbvserver binding block supports the following:

* `priority` - (Optional) Priority.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke policies bound to a virtual server or policy label.
* `labeltype` - (Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]
* `labelname` - (Optional) Name of the label invoked.
* `name` - (Optional) Name for the virtual server.
* `bindpoint` - (Optional) Bind point to which to bind the policy. Applicable only to compression, rewrite, videooptimization and cache policies. Possible values: [ REQUEST, RESPONSE ]

A csvserver binding block supports the following:

* `priority` - (Optional) Priority for the policy.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke flag.
* `labeltype` - (Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]
* `labelname` - (Optional) Name of the label invoked.
* `name` - (Optional) Name of the content switching virtual server to which the content switching policy applies.
* `targetlbvserver` - (Optional) Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE. Example: bind cs vs cs1 -policyname pol1 -priority 101 -targetLBVserver lb1 Note: Use this parameter only in case of Content Switching policy bind operations to a CS vserver.
* `bindpoint` - (Optional) For a rewrite policy, the bind point to which to bind the policy. Note: This parameter applies only to rewrite policies, because content switching policies are evaluated only at request time. Possible values: [ REQUEST, RESPONSE, ICA_REQUEST, OTHERTCP_REQUEST ]



## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the responderpolicy. It has the same value as the `name` attribute.


## Import

A responderpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_responderpolicy.tf_responder_policy tf_responder_policy
```
