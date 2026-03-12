---
subcategory: "VPN"
---

# Data Source: vpnglobal_sslcertkey_binding

The vpnglobal_sslcertkey_binding data source allows you to retrieve information about a vpnglobal_sslcertkey_binding.


## Example Usage

```terraform
data "citrixadc_vpnglobal_sslcertkey_binding" "tf_vpnglobal_sslcertkey_binding" {
  certkeyname = "sample_ssl_cert"
}

output "certkeyname" {
  value = data.citrixadc_vpnglobal_sslcertkey_binding.tf_vpnglobal_sslcertkey_binding.certkeyname
}

output "crlcheck" {
  value = data.citrixadc_vpnglobal_sslcertkey_binding.tf_vpnglobal_sslcertkey_binding.crlcheck
}
```


## Argument Reference

* `certkeyname` - (Optional) SSL certkey to use in signing tokens. Only RSA cert key is allowed
* `cacert` - (Optional) The name of the CA certificate binding.
* `userdataencryptionkey` - (Optional) Certificate to be used for encrypting user data like KB Question and Answers, Alternate Email Address, etc.

**Note:** At least one of the above arguments should be provided to identify the specific binding.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `crlcheck` - The state of the CRL check parameter (Mandatory/Optional).
* `gotopriorityexpression` - Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `id` - The id of the vpnglobal_sslcertkey_binding. It is a system-generated identifier.
* `ocspcheck` - The state of the OCSP check parameter (Mandatory/Optional).
