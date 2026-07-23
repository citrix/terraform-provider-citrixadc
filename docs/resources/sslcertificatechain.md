---
subcategory: "SSL"
---

# Resource: sslcertificatechain

This resource is used to build the SSL certificate chain for a certificate-key pair on the Citrix ADC.


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
