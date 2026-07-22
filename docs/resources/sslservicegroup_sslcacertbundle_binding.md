---
subcategory: "SSL"
---

# Resource: sslservicegroup_sslcacertbundle_binding

Binds a CA certificate bundle to an SSL service group. The bound CA certificate bundle is used to validate the certificate chain presented by the back-end server during SSL handshakes for the service group.


## Example usage

```hcl
resource "citrixadc_sslservicegroup_sslcacertbundle_binding" "tf_binding" {
  servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
  cacertbundlename = citrixadc_sslcacertbundle.tf_cacertbundle.cacertbundlename
}

resource "citrixadc_servicegroup" "tf_servicegroup" {
  servicegroupname = "tf_servicegroup"
  servicetype      = "SSL"
}

resource "citrixadc_sslcacertbundle" "tf_cacertbundle" {
  cacertbundlename = "tf_cacertbundle"
  src              = "local:ca-bundle.pem"
}
```


## Argument Reference

* `servicegroupname` - (Required) The name of the SSL service group to which the CA certificate bundle is bound. Changing this forces a new resource to be created.
* `cacertbundlename` - (Required) The name of the CA certificate bundle bound to the service group. Changing this forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslservicegroup_sslcacertbundle_binding. It is the concatenation of the `servicegroupname` and `cacertbundlename` attributes separated by a comma.


## Import

A sslservicegroup_sslcacertbundle_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_sslservicegroup_sslcacertbundle_binding.tf_binding tf_servicegroup,tf_cacertbundle
```
