---
subcategory: "SSL"
---

# Data Source: sslservicegroup_sslcertkey_binding

The sslservicegroup_sslcertkey_binding data source allows you to retrieve information about an SSL certificate key binding to an SSL service group.

## Example Usage

```terraform
data "citrixadc_sslservicegroup_sslcertkey_binding" "tf_sslservicegroup_sslcertkey_binding" {
  servicegroupname = "tf_servicegroup"
  certkeyname      = "tf_sslcertkey"
  ca               = false
}

output "crlcheck" {
  value = data.citrixadc_sslservicegroup_sslcertkey_binding.tf_sslservicegroup_sslcertkey_binding.crlcheck
}

output "ocspcheck" {
  value = data.citrixadc_sslservicegroup_sslcertkey_binding.tf_sslservicegroup_sslcertkey_binding.ocspcheck
}
```

## Argument Reference

* `servicegroupname` - (Required) The name of the SSL service group to which the certificate is bound.
* `certkeyname` - (Required) The name of the certificate bound to the SSL service group.
* `ca` - (Required) CA certificate.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslservicegroup_sslcertkey_binding. It is a system-generated identifier.
* `ocspcheck` - The state of the OCSP check parameter. (Mandatory/Optional)
* `snicert` - The name of the CertKey. Use this option to bind Certkey(s) which will be used in SNI processing.
* `crlcheck` - The state of the CRL check parameter. (Mandatory/Optional)
