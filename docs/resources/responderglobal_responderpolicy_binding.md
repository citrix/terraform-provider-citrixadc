---
subcategory: "Responder"
---

# Resource: responderglobal_responderpolicy_binding

The responderglobal_responderpolicy_binding resource is used to create responderglobal_responderpolicy_binding.


## Example usage

```hcl
resource "citrixadc_responderglobal_responderpolicy_binding" "tf_responderglobal_responderpolicy_binding" {
  globalbindtype = "SYSTEM_GLOBAL"
  priority   = 50
  policyname =citrixadc_responderpolicy.tf_responderpolicy.name
}


resource "citrixadc_responderpolicy" "tf_responderpolicy" {
	name    = "tf_responderpolicy"
	action = "NOOP"
	rule = "HTTP.REQ.URL.PATH_AND_QUERY.CONTAINS(\"nosuchthing\")"
}
```


## Argument Reference

* `policyname` - (Required) Name of the responder policy.
* `priority` - (Required) Specifies the priority of the policy.
* `type` - (Optional) Specifies the bind point whose policies you want to display. Available settings function as follows: * REQ_OVERRIDE - Request override. Binds the policy to the priority request queue. * REQ_DEFAULT - Binds the policy to the default request queue. * OTHERTCP_REQ_OVERRIDE - Binds the policy to the non-HTTP TCP priority request queue. * OTHERTCP_REQ_DEFAULT - Binds the policy to the non-HTTP TCP default request queue.. * SIPUDP_REQ_OVERRIDE - Binds the policy to the SIP UDP priority response queue.. * SIPUDP_REQ_DEFAULT - Binds the policy to the SIP UDP default response queue. * RADIUS_REQ_OVERRIDE - Binds the policy to the RADIUS priority response queue.. * RADIUS_REQ_DEFAULT - Binds the policy to the RADIUS default response queue. * MSSQL_REQ_OVERRIDE - Binds the policy to the Microsoft SQL priority response queue.. * MSSQL_REQ_DEFAULT - Binds the policy to the Microsoft SQL default response queue. * MYSQL_REQ_OVERRIDE - Binds the policy to the MySQL priority response queue. * MYSQL_REQ_DEFAULT - Binds the policy to the MySQL default response queue. * HTTPQUIC_REQ_OVERRIDE - Binds the policy to the HTTP_QUIC override response queue. * HTTPQUIC_REQ_DEFAULT - Binds the policy to the HTTP_QUIC default response queue. Possible values: [ REQ_OVERRIDE, REQ_DEFAULT, OVERRIDE, DEFAULT, OTHERTCP_REQ_OVERRIDE, OTHERTCP_REQ_DEFAULT, SIPUDP_REQ_OVERRIDE, SIPUDP_REQ_DEFAULT, SIPTCP_REQ_OVERRIDE, SIPTCP_REQ_DEFAULT, MSSQL_REQ_OVERRIDE, MSSQL_REQ_DEFAULT, MYSQL_REQ_OVERRIDE, MYSQL_REQ_DEFAULT, NAT_REQ_OVERRIDE, NAT_REQ_DEFAULT, DIAMETER_REQ_OVERRIDE, DIAMETER_REQ_DEFAULT, RADIUS_REQ_OVERRIDE, RADIUS_REQ_DEFAULT, DNS_REQ_OVERRIDE, DNS_REQ_DEFAULT, MQTT_REQ_OVERRIDE, MQTT_REQ_DEFAULT, MQTT_JUMBO_REQ_OVERRIDE, MQTT_JUMBO_REQ_DEFAULT, QUIC_OVERRIDE, QUIC_DEFAULT, HTTPQUIC_REQ_OVERRIDE, HTTPQUIC_REQ_DEFAULT ]
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.
* `labeltype` - (Optional) Type of invocation, Available settings function as follows: * vserver - Forward the request to the specified virtual server. * policylabel - Invoke the specified policy label. Possible values: [ vserver, policylabel ]
* `labelname` - (Optional) Name of the policy label to invoke. If the current policy evaluates to TRUE, the invoke parameter is set, and Label Type is policylabel.
* `globalbindtype` - (Optional) . Possible values: [ SYSTEM_GLOBAL, VPN_GLOBAL, RNAT_GLOBAL ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the responderglobal_responderpolicy_binding. It has the same value as the `policyname` attribute.


## Import

A responderglobal_responderpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_responderglobal_responderpolicy_binding.tf_responderglobal_responderpolicy_binding tf_responderpolicy
```
