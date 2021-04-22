---
subcategory: "Content Switching"
---

# Resource: csvserver\_cspolicy\_binding

The csvserver\_cspolicy\_binding resource is used to create bindings between a content switching virtual server and a content switcing policy.


## Example usage

```hcl
resource "citrixadc_csvserver" "tf_csvserver" {
  ipv46       = "10.10.10.22"
  name        = "tf_csvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.33"
  name        = "tf_lbvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_cspolicy" "tf_cspolicy" {
  policyname      = "tf_cspolicy"
  rule            = "CLIENT.IP.SRC.SUBNET(24).EQ(10.217.85.0)"
}

resource "citrixadc_csvserver_cspolicy_binding" "tf_csvscspolbind" {
    name = citrixadc_csvserver.tf_csvserver.name
    policyname = citrixadc_cspolicy.tf_cspolicy.policyname
    priority = 100
    targetlbvserver = citrixadc_lbvserver.tf_lbvserver.name
}
```


## Argument Reference

* `policyname` - (Optional) Policies bound to this vserver.
* `targetlbvserver` - (Optional) target vserver name.
* `priority` - (Optional) Priority for the policy.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `bindpoint` - (Optional) The bindpoint to which the policy is bound. Possible values: [ REQUEST, RESPONSE, ICA_REQUEST, OTHERTCP_REQUEST ]
* `invoke` - (Optional) Invoke flag.
* `labeltype` - (Optional) The invocation type. Possible values: [ reqvserver, resvserver, policylabel ]
* `labelname` - (Optional) Name of the label invoked.
* `name` - (Optional) Name of the content switching virtual server to which the content switching policy applies.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the csvserver\_cspolicy\_binding . It is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A csvserver\_cspolicy\_binding  can be imported using its id, e.g.

```shell
terraform import citrixadc_csvserver_cspolicy_binding.tf_csvscspolbind tf_csvserver,tf_cspolicy
```
