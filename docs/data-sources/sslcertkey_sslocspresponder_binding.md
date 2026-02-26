---
subcategory: "SSL"
---

# Data Source: sslcertkey_sslocspresponder_binding

The sslcertkey_sslocspresponder_binding data source allows you to retrieve information about an OCSP responder binding to an SSL certificate-key pair.

## Example Usage

```terraform
data "citrixadc_sslcertkey_sslocspresponder_binding" "tf_binding" {
  certkey       = "tf_sslcertkey"
  ocspresponder = "tf_sslocspresponder"
}

output "priority" {
  value = data.citrixadc_sslcertkey_sslocspresponder_binding.tf_binding.priority
}
```

## Argument Reference

* `certkey` - (Required) Name of the certificate-key pair.
* `ocspresponder` - (Required) OCSP responders bound to this certkey.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslcertkey_sslocspresponder_binding. It is a system-generated identifier.
* `priority` - OCSP priority.
* `ca` - The certificate-key pair being unbound is a Certificate Authority (CA) certificate. If you choose this option, the certificate-key pair is unbound from the list of CA certificates that were bound to the specified SSL virtual server or SSL service.
