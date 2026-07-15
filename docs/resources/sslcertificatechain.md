---
subcategory: "SSL"
---

# Resource: sslcertificatechain

Forming a certificate chain links an existing server certificate-key pair to its intermediate CA certificate(s) on the Citrix ADC so that the appliance presents the complete trust chain to clients during the SSL handshake. Building the chain lets clients validate the server certificate up to a trusted root without having to obtain the intermediates separately. This resource triggers the chain-building operation for a previously configured `sslcertkey`.


## Example usage

```hcl
resource "citrixadc_sslcertkey" "servercert1" {
  certkey = "servercert1"
  cert    = "/nsconfig/ssl/server.crt"
  key     = "/nsconfig/ssl/server.key"
}

resource "citrixadc_sslcertificatechain" "tf_certchain" {
  certkeyname = citrixadc_sslcertkey.servercert1.certkey
}
```


## Argument Reference

* `certkeyname` - (Required) Name of the certificate-key pair for which the certificate chain is formed. This must reference an existing `sslcertkey` on the appliance. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslcertificatechain. It has the same value as the `certkeyname` attribute.


## Import

A sslcertificatechain can be imported using its certkeyname, e.g.

```shell
terraform import citrixadc_sslcertificatechain.tf_sslcertificatechain tf_sslcertificatechain
```


## Note on deletion

The Citrix ADC NITRO API exposes no delete endpoint for `sslcertificatechain`. The certificate chain that this resource forms cannot be removed through NITRO. When you run `terraform destroy` (or otherwise remove this resource), Terraform only drops the resource from its state; the chain that was formed remains in effect on the appliance.
