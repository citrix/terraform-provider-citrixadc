---
subcategory: "SSL"
---

# Resource: sslvserver\_sslcertkeybundle\_binding

Binds a certificate-key bundle to an SSL virtual server on the Citrix ADC. Use this resource to associate a server certificate (and its private key) bundle with an SSL vserver so it can be presented during the SSL handshake, optionally as an SNI certificate.


## Example usage

```hcl
resource "citrixadc_sslcertkeybundle" "tf_certkeybundle" {
  name = "tf_certkeybundle"
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_lbvserver"
  ipv46       = "10.10.10.44"
  port        = 443
  servicetype = "SSL"
  sslprofile  = "ns_default_ssl_profile_frontend"
}

resource "citrixadc_sslvserver_sslcertkeybundle_binding" "tf_binding" {
  vservername       = citrixadc_lbvserver.tf_lbvserver.name
  certkeybundlename = citrixadc_sslcertkeybundle.tf_certkeybundle.name
  snicertkeybundle  = false
}
```


## Argument Reference

* `vservername` - (Required) Name of the SSL virtual server. Changing this attribute forces a new resource to be created.
* `certkeybundlename` - (Required) Certkeybundle name bound to the vserver. Changing this attribute forces a new resource to be created.
* `snicertkeybundle` - (Optional) Use this option to bind certkeybundle which will be used in SNI processing. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslvserver\_sslcertkeybundle\_binding. It is the concatenation of the `vservername` and `certkeybundlename` attributes separated by a comma.


## Import

A sslvserver\_sslcertkeybundle\_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_sslvserver_sslcertkeybundle_binding.tf_binding tf_lbvserver,tf_certkeybundle
```
