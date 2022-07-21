---
subcategory: "Cache Redirection"
---

# Resource: <fillme>

The <resource> resource is used to create <fillme>.


## Example usage

```hcl
<fillme>
```


## Argument Reference

* `bindpoint` - (Optional) For a rewrite policy, the bind point to which to bind the policy. Note: This parameter applies only to rewrite policies, because content switching policies are evaluated only at request time.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke a policy label if this policy's rule evaluates to TRUE (valid only for default-syntax policies such as application firewall, transform, integrated cache, rewrite, responder, and content switching).
* `labelname` - (Optional) Name of the label to be invoked.
* `labeltype` - (Optional) Type of label to be invoked.
* `name` - (Optional) Name of the cache redirection virtual server to which to bind the cache redirection policy.
* `policyname` - (Optional) Policies bound to this vserver.
* `priority` - (Optional) The priority for the policy.
* `targetvserver` - (Optional) Name of the virtual server to which content is forwarded. Applicable only if the policy is a map policy and the cache redirection virtual server is of type REVERSE.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the <fillme>. It has the same value as the `name` attribute.


## Import

A <resource> can be imported using its name, e.g.

```shell
terraform import citrixadc_csaction.tf_csaction tf_csaction
```
