---
subcategory: "Load Balancing"
---

# Resource: lbvserver_dnspolicy64_binding

The lbvserver_dnspolicy64_binding resource is used to add a dns policy64 to lbvserver.


## Example usage

```hcl
resource "citrixadc_lbvserver_dnspolicy64_binding" "tf_lbvserver_dnspolicy64_binding" {
        name = "tf_lbvserver"
        policyname = "tf_dnspolicy64"
        priority = 1
}
```


## Argument Reference

* `policyname` - (Required) Name of the policy bound to the LB vserver.
* `priority` - (Optional) Priority.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `name` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.  CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver'). .
* `bindpoint` - (Optional) Bind point to which to bind the policy. Applicable only to compression, rewrite, videooptimization and cache policies. Possible values: [ REQUEST, RESPONSE ]
* `invoke` - (Optional) Invoke policies bound to a virtual server or policy label.
* `labeltype` - (Optional) Type of policy label to invoke. Applicable only to rewrite, videooptimization and cache policies. Available settings function as follows: * reqvserver - Evaluate the request against the request-based policies bound to the specified virtual server. * resvserver - Evaluate the response against the response-based policies bound to the specified virtual server. * policylabel - invoke the request or response against the specified user-defined policy label. Possible values: [ reqvserver, resvserver, policylabel ]
* `labelname` - (Optional) Name of the virtual server or user-defined policy label to invoke if the policy evaluates to TRUE.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbvserver_dnspolicy64_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A lbvserver_dnspolicy64_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_lbvserver_dnspolicy64_binding.tf_lbvserver_dnspolicy64_binding tf_lbvserver,tf_dnspolicy64
```