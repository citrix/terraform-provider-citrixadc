---
subcategory: "Content Switching"
---

# Resource: cspolicy

The resource is used to configure content switching policies.


## Example usage

```hcl
resource "citrixadc_csvserver" "tf_csvserver" {
  ipv46       = "10.10.10.22"
  name        = "tf_csvserver"
  port        = 80
  servicetype = "HTTP"
}

resource "citrixadc_cspolicy" "tf_cspolicy" {
  policyname = "tf_cspolicy"
  rule       = "CLIENT.IP.SRC.SUBNET(24).EQ(10.217.85.0)"
}

resource "citrixadc_csvserver_cspolicy_binding" "tf_csvscspolbind" {
  name                   = citrixadc_csvserver.tf_csvserver.name
  policyname             = citrixadc_cspolicy.tf_cspolicy.policyname
  priority               = 100
  gotopriorityexpression = "NEXT"
}
```

!>
[**DEPRECATED**] Please use `csvserver_cspolicy_binding` to bind `csvserver` to `cspolicy` insted of this resource. the support for binding `csvserver` to `cspolicy` in `cspolicy` resource will get deprecated soon.

## Argument Reference

* `policyname` - (Optional) Name of the content switching policy
* `url` - (Optional) URL string that is matched with the URL of a request. Can contain a wildcard character. Specify the string value in the following format: [[prefix] [*]] [.suffix].
* `rule` - (Optional) Expression, or name of a named expression, against which traffic is evaluated.
* `domain` - (Optional) The domain name. The string value can range to 63 characters.
* `boundto` - (Optional) The boundto name. The string value can range to 63 characters.
* `action` - (Optional) Content switching action that names the target load balancing virtual server to which the traffic is switched.
* `logaction` - (Optional) The log action associated with the content switching policy.
* `csvserver` - [DEPRECATED] Content switching vserver that this policy will bind to.
* `priority`  - (Optional) Priority for the binding with the Content switching vserver.
* `targetlbvserver` - [DEPRECATED] Targe load balancing vserver that will be used for the binding with the Content switching vserver.
* `forcenew_id_set` - (Optional) A list of terraform resource ids. Any change in the list of values will trigger the recreation of the cspolicy resource.

    Its main intent is to force the rebinding with the Content switching vserver defined in `csvserver` should it be deleted and recreated.

    The same applies for the LB vserver defined in `targetlbvserver`.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cspolicy. It has the same value as the `policyname` attribute.


## Import

An instance of the resource can be imported using its `policyname`, e.g.

```shell
terraform import citrixadc_cspolicy.tf_cspolicy tf_cspolicy
```
