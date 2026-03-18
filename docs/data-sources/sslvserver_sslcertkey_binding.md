---
subcategory: "SSL"
---

# Data Source: sslvserver_sslcertkey_binding

The sslvserver_sslcertkey_binding data source allows you to retrieve information about an SSL certificate key binding to an SSL virtual server.

## Example usage

```terraform
data "citrixadc_sslvserver_sslcertkey_binding" "tf_binding" {
  vservername = "tf_lbvserver"
  certkeyname = "tf_sslcertkey"
  ca          = false
  snicert     = false
}

output "vservername" {
  value = data.citrixadc_sslvserver_sslcertkey_binding.tf_binding.vservername
}

output "certkeyname" {
  value = data.citrixadc_sslvserver_sslcertkey_binding.tf_binding.certkeyname
}
```

## Argument Reference

The following arguments are required:

* `vservername` - (Required) Name of the SSL virtual server.
* `certkeyname` - (Required) The name of the certificate key pair binding.
* `ca` - (Required) CA certificate.
* `snicert` - (Required) The name of the CertKey. Use this option to bind Certkey(s) which will be used in SNI processing.

The following arguments are optional:

* `crlcheck` - (Optional) The state of the CRL check parameter. (Mandatory/Optional)
* `ocspcheck` - (Optional) The state of the OCSP check parameter. (Mandatory/Optional)

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslvserver_sslcertkey_binding. It is a system-generated identifier.
* `skipcaname` - The flag is used to indicate whether this particular CA certificate's CA_Name needs to be sent to the SSL client while requesting for client certificate in a SSL handshake.
