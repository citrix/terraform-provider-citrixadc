---
subcategory: "SSL"
---

# Data Source: sslservice_sslcertkey_binding

The sslservice_sslcertkey_binding data source allows you to retrieve information about the binding between an SSL service and an SSL certificate key pair.

## Example Usage

```terraform
data "citrixadc_sslservice_sslcertkey_binding" "example" {
  servicename  = "tf_service"
  certkeyname  = "tf_certkey"
  ca           = true
}

output "ocspcheck" {
  value = data.citrixadc_sslservice_sslcertkey_binding.example.ocspcheck
}

output "skipcaname" {
  value = data.citrixadc_sslservice_sslcertkey_binding.example.skipcaname
}
```

## Argument Reference

* `servicename` - (Required) Name of the SSL service for which to set advanced configuration.
* `certkeyname` - (Required) The certificate key pair binding.
* `ca` - (Required) CA certificate.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslservice_sslcertkey_binding. It is a system-generated identifier.
* `ocspcheck` - Rule to use for the OCSP responder associated with the CA certificate during client authentication. If MANDATORY is specified, deny all SSL clients if the OCSP check fails because of connectivity issues with the remote OCSP server, or any other reason that prevents the OCSP check. With the OPTIONAL setting, allow SSL clients even if the OCSP check fails except when the client certificate is revoked.
* `crlcheck` - The state of the CRL check parameter. (Mandatory/Optional)
* `snicert` - The name of the CertKey. Use this option to bind Certkey(s) which will be used in SNI processing.
* `skipcaname` - The flag is used to indicate whether this particular CA certificate's CA_Name needs to be sent to the SSL client while requesting for client certificate in a SSL handshake.
