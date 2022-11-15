---
subcategory: "Content Switching"
---

# Resource: cspolicy

The resource is used to configure content switching policies.


## Example usage

```hcl
resource "citrixadc_csvserver" "foo_csvserver" {

  ipv46       = "10.202.11.11"
  name        = "tst_policy_cs"
  port        = 9090
  servicetype = "SSL"
  comment     = "hello"
  sslprofile  = "ns_default_ssl_profile_frontend"
}

resource "citrixadc_lbvserver" "foo_lbvserver" {

  name        = "tst_policy_lb"
  servicetype = "HTTP"
  ipv46       = "192.122.3.3"
  port        = 8000
  comment     = "hello"
}

resource "citrixadc_cspolicy" "tf_cspolicy" {
  csvserver       = citrixadc_csvserver.foo_csvserver.name
  targetlbvserver = citrixadc_lbvserver.foo_lbvserver.name
  policyname      = "tf_cspolicy"
  rule            = "CLIENT.IP.SRC.SUBNET(24).EQ(10.217.85.0)"
  priority        = 10

  # Any change in the following id set will force recreation of the cs policy
  forcenew_id_set = [
    citrixadc_lbvserver.foo_lbvserver.id,
    citrixadc_csvserver.foo_csvserver.id,
  ]
}
```

!>
To bind `csvserver` to `cspolicy` please use `csvserver_cspolicy_binding` insted of this resource. As this support for binding in `cspolicy` resource will get deprecated soon.

## Argument Reference

* `policyname` - (Optional) Name of the content switching policy
* `url` - (Optional) URL string that is matched with the URL of a request. Can contain a wildcard character. Specify the string value in the following format: [[prefix] [*]] [.suffix].
* `rule` - (Optional) Expression, or name of a named expression, against which traffic is evaluated.
* `domain` - (Optional) The domain name. The string value can range to 63 characters.
* `boundto` - (Optional) The boundto name. The string value can range to 63 characters.
* `action` - (Optional) Content switching action that names the target load balancing virtual server to which the traffic is switched.
* `logaction` - (Optional) The log action associated with the content switching policy.
* `csvserver` - (Required) Content switching vserver that this policy will bind to.
* `priority`  - (Optional) Priority for the binding with the Content switching vserver.
* `targetlbvserver` - (Optional) Targe load balancing vserver that will be used for the binding with the Content switching vserver.
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
