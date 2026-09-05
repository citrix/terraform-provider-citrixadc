---
subcategory: "Cache Redirection"
---

# Resource: crvserver_policymap_binding

The crvserver_policymap_binding resource is used to create CR vserver Policymap Binding.


## Example usage

```hcl
resource "citrixadc_crvserver" "crvserver" {
  name        = "my_vserver"
  servicetype = "HTTP"
  ipv46       = "10.102.80.55"
  port        = 8090
  cachetype   = "REVERSE"
}
resource "citrixadc_policymap" "tf_policymap" {
  mappolicyname = "ia_mappol123"
  sd            = "amazon.com"
  td            = "apple.com"
}
resource "citrixadc_lbvserver" "foo_lbvserver" {
  name        = "test_lbvserver"
  servicetype = "HTTP"
  ipv46       = "192.122.3.31"
  port        = 8000
  comment     = "hello"
}
resource "citrixadc_service" "tf_service" {
  lbvserver   = citrixadc_lbvserver.foo_lbvserver.name
  name        = "tf_service"
  port        = 8081
  ip          = "10.33.4.5"
  servicetype = "HTTP"
  cachetype   = "TRANSPARENT"
}
resource "citrixadc_crvserver_policymap_binding" "crvserver_policymap_binding" {
  name          = citrixadc_crvserver.crvserver.name
  policyname    = citrixadc_policymap.tf_policymap.mappolicyname
  targetvserver = citrixadc_lbvserver.foo_lbvserver.name
  depends_on = [
    citrixadc_service.tf_service
  ]
}

```


## Argument Reference

* `name` - (Required) Name of the cache redirection virtual server to which to bind the cache redirection policy.
* `policyname` - (Required) Policies bound to this vserver.
* `bindpoint` - (Optional) For a rewrite policy, the bind point to which to bind the policy. Note: This parameter applies only to rewrite policies, because content switching policies are evaluated only at request time.
* `gotopriorityexpression` - (Optional) Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.
* `invoke` - (Optional) Invoke a policy label if this policy's rule evaluates to TRUE.
* `labelname` - (Optional) Name of the label to be invoked.
* `labeltype` - (Optional) Type of label to be invoked.
* `priority` - (Optional) An unsigned integer that determines the priority of the policy relative to other policies bound to this cache redirection virtual server. The lower the value, higher the priority. Note: This option is available only when binding content switching, filtering, and compression policies to a cache redirection virtual server.
* `targetvserver` - (Optional) The CSW target server names.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the crvserver_policymap_binding. It is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A crvserver_policymap_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_crvserver_policymap_binding.crvserver_policymap_binding my_vserver,ia_mappol123
```
