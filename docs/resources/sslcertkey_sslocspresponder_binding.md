---
subcategory: "SSL"
---

# Resource: sslcertkey_sslocspresponder_binding

The sslcertkey_sslocspresponder_binding resource is used to bind sslocspresponder to sslcertkey.


## Example usage

```hcl
resource "citrixadc_sslcertkey" "tf_sslcertkey" {
  certkey            = "tf_sslcertkey"
  cert               = "/nsconfig/ssl/certificate1.crt"
  key                = "/nsconfig/ssl/key1.pem"
  notificationperiod = 40
  expirymonitor      = "ENABLED"
}
resource "citrixadc_sslocspresponder" "tf_sslocspresponder" {
  name = "tf_sslocspresponder"
  url  = "http://www.google.com"
}
resource "citrixadc_sslcertkey_sslocspresponder_binding" "tf_binding" {
  certkey       = citrixadc_sslcertkey.tf_sslcertkey.certkey
  ocspresponder = citrixadc_sslocspresponder.tf_sslocspresponder.name
  priority      = 90
}
```


## Argument Reference


* `certkey` - (Required) Name of the certificate-key pair.
* `ocspresponder` - (Required) OCSP responders bound to this certkey
* `priority` - (Required) ocsp priority


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslcertkey_sslocspresponder_binding. It is the concatenation of both `certkey` and `ocspresponder` attributes seperated by comma. Ex: `tf_sslcertkey,tf_sslocspresponder` is `id` for above example.


## Import

A sslcertkey_sslocspresponder_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_sslcertkey_sslocspresponder_binding.tf_binding tf_sslcertkey,tf_sslocspresponder
```
