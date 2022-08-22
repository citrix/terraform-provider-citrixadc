---
subcategory: "Compression"
---

# Resource: cmppolicylabel_cmppolicy_binding

The cmppolicylabel_cmppolicy_binding resource is used to create cmppolicylabel_cmppolicy_binding.


## Example usage

```hcl
resource "citrixadc_cmppolicylabel_cmppolicy_binding" "tf_cmppolicylabel_cmppolicy_binding" {
  policyname = citrixadc_cmppolicy.tf_cmppolicy.name
  labelname  = "my_cmppolicy_label"
  priority   = 100
}

resource "citrixadc_cmppolicy" "tf_cmppolicy" {
  name      = "tf_cmppolicy"
  rule      = "HTTP.RES.HEADER(\"Content-Type\").CONTAINS(\"text\")"
  resaction = "COMPRESS"
}
```


## Argument Reference

* `policyname` - (Optional) The compression policy name.
* `priority` - (Optional) Specifies the priority of the policy.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke policies bound to a virtual server or a user-defined policy label. After the invoked policies are evaluated, the flow returns to the policy with the next higher priority number in the original label.
* `labeltype` - (Optional) Type of policy label invocation. Possible values: [ reqvserver, resvserver, policylabel ]
* `invoke_labelname` - (Optional) Name of the label to invoke if the current policy evaluates to TRUE.
* `labelname` - (Optional) Name of the HTTP compression policy label to which to bind the policy. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cmppolicylabel_cmppolicy_binding. It is the concatenation of  `labelname` and `policyname` attributes separated by a comma.


## Import

A cmppolicylabel_cmppolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_cmppolicylabel_cmppolicy_binding.tf_cmppolicylabel_cmppolicy_binding my_cmppolicy_label,tf_cmppolicy
```
