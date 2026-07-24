---
subcategory: "SSL"
---

# Resource: sslvserver\_sslcacertbundle\_binding

This resource is used to manage the binding between a CA certificate bundle and an SSL virtual server.


## Example usage

```hcl
resource "citrixadc_sslcacertbundle" "tf_cacertbundle" {
  cacertbundlename = "tf_cacertbundle"
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_lbvserver"
  ipv46       = "10.10.10.44"
  port        = 443
  servicetype = "SSL"
  sslprofile  = "ns_default_ssl_profile_frontend"
}

resource "citrixadc_sslvserver_sslcacertbundle_binding" "tf_binding" {
  vservername      = citrixadc_lbvserver.tf_lbvserver.name
  cacertbundlename = citrixadc_sslcacertbundle.tf_cacertbundle.cacertbundlename
  skipcacertbundle = false
}
```


## Argument Reference

* `vservername` - (Required) Name of the SSL virtual server. Changing this attribute forces a new resource to be created.
* `cacertbundlename` - (Required) CA certbundle name bound to the vserver. Changing this attribute forces a new resource to be created.
* `skipcacertbundle` - (Optional) The flag is used to indicate whether this particular CA certificate's CA_Name needs to be sent to the SSL client while requesting for client certificate in a SSL handshake. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslvserver\_sslcacertbundle\_binding. It is the concatenation of the `vservername` and `cacertbundlename` attributes separated by a comma.


## Import

A sslvserver\_sslcacertbundle\_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_sslvserver_sslcacertbundle_binding.tf_binding tf_lbvserver,tf_cacertbundle
```
