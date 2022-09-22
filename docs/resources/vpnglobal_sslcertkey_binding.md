---
subcategory: "VPN"
---

# Resource: vpnglobal_sslcertkey_binding

The vpnglobal_sslcertkey_binding resource is used to bind an ssl CertKey to the global configuration.


## Example usage

```hcl
resource "citrixadc_sslcertkey" "foo" {
  certkey            = "sample_ssl_cert"
  cert               = "/var/tmp/certificate1.crt"
  key                = "/var/tmp/key1.pem"
  notificationperiod = 10
  expirymonitor      = "ENABLED"
}
resource "citrixadc_vpnglobal_sslcertkey_binding" "tf_vpnglobal_slcertkey_binding" {
  certkeyname = citrixadc_sslcertkey.foo.certkey
}
```


## Argument Reference

* `certkeyname` - (Required) SSL certkey to use in signing tokens.
* `cacert` - (Optional) The name of the CA certificate binding.
* `crlcheck` - (Optional) The state of the CRL check parameter (Mandatory/Optional).
* `gotopriorityexpression` - (Optional) Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `ocspcheck` - (Optional) The state of the OCSP check parameter (Mandatory/Optional).
* `userdataencryptionkey` - (Optional) Certificate to be used for encrypting user data like KB Question and Answers, Alternate Email Address, etc.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnglobal_sslcertkey_binding. It has the same value as the `certkeyname` attribute.


## Import

A vpnglobal_sslcertkey_binding can be imported using its certkeyname, e.g.

```shell
terraform import citrixadc_vpnglobal_sslcertkey_binding.tf_vpnglobal_slcertkey_binding sample_ssl_cert
```
