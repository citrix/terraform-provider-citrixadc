---
subcategory: "SSL"
---

# Resource: sslservicegroup_sslcertkey_binding

The sslservicegroup_sslcertkey_binding resource is used to add an sslcertkey to sslservicegroup.


## Example usage

```hcl
resource "citrixadc_sslservicegroup_sslcertkey_binding" "tf_sslservicegroup_sslcertkey_binding" {
	ca = false
	certkeyname = citrixadc_sslcertkey.tf_sslcertkey.certkey
	servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
}

resource "citrixadc_sslcertkey" "tf_sslcertkey" {
	certkey = "tf_sslcertkey"
	cert = "/var/tmp/certificate1.crt"
	key = "/var/tmp/key1.pem"
}

resource "citrixadc_servicegroup" "tf_servicegroup" {
	servicegroupname = "tf_servicegroup"
	servicetype = "SSL"
}
```


## Argument Reference

* `certkeyname` - (Required) The name of the certificate bound to the SSL service group.
* `crlcheck` - (Optional) The state of the CRL check parameter. (Mandatory/Optional). Possible values: [ Mandatory, Optional ]
* `ocspcheck` - (Optional) The state of the OCSP check parameter. (Mandatory/Optional). Possible values: [ Mandatory, Optional ]
* `ca` - (Optional) CA certificate.
* `snicert` - (Optional) The name of the CertKey. Use this option to bind Certkey(s) which will be used in SNI processing.
* `servicegroupname` - (Required) The name of the SSL service to which the SSL policy needs to be bound.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslservicegroup_sslcertkey_binding. It is the concatenation of the `servicegroupname` and `certkeyname` attributes separated by a comma.


## Import

A sslservicegroup_sslcertkey_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_sslservicegroup_sslcertkey_binding.tf_sslservicegroup_sslcertkey_binding tf_servicegroup,tf_sslcertkey
```
