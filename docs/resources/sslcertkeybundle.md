---
subcategory: "SSL"
---

# Resource: sslcertkeybundle

This resource is used to manage SSL certificate-key bundles.


## Example usage

### Using passplain (sensitive attribute - persisted in state)

```hcl
variable "sslcertkeybundle_passplain" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslcertkeybundle" "tf_certkeybundle" {
  certkeybundlename = "web-certkey-bundle"
  bundlefile        = "/nsconfig/ssl/server_bundle.pem"
  passplain         = var.sslcertkeybundle_passplain
}
```

### Using passplain_wo (write-only/ephemeral - NOT persisted in state)

The `passplain_wo` attribute provides an ephemeral path for the private-key pass phrase. The value is sent to the Citrix ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the value changes, increment `passplain_wo_version`.

```hcl
variable "sslcertkeybundle_passplain" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslcertkeybundle" "tf_certkeybundle" {
  certkeybundlename    = "web-certkey-bundle"
  bundlefile           = "/nsconfig/ssl/server_bundle.pem"
  passplain_wo         = var.sslcertkeybundle_passplain
  passplain_wo_version = 1
}
```

To rotate the secret, update the variable value and bump the version:

```hcl
resource "citrixadc_sslcertkeybundle" "tf_certkeybundle" {
  certkeybundlename    = "web-certkey-bundle"
  bundlefile           = "/nsconfig/ssl/server_bundle.pem"
  passplain_wo         = var.sslcertkeybundle_passplain
  passplain_wo_version = 2  # Bumped to trigger update
}
```


## Argument Reference

* `certkeybundlename` - (Required) Name given to the certKeyBundle. The name will be used to bind/unbind the certkey bundle to a VIP. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my file" or 'my file'). Changing this attribute forces a new resource to be created.
* `bundlefile` - (Required) Name of and, optionally, path to the X509 certificate bundle file that is used to form the certificate-key bundle. The certificate bundle file must already be present on the appliance's hard-disk drive or solid-state drive. `/nsconfig/ssl/` is the default path. The certificate bundle file consists of a list of certificates and one key in PEM format.
* `passplain` - (Optional, Sensitive) Pass phrase used to encrypt the private-key. Required when the certificate bundle file contains an encrypted private-key in PEM format. The value is persisted in Terraform state (encrypted). See also `passplain_wo` for an ephemeral alternative.
* `passplain_wo` - (Optional, Sensitive, WriteOnly) Same as `passplain`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `passplain_wo_version`. If both `passplain` and `passplain_wo` are set, `passplain_wo` takes precedence.
* `passplain_wo_version` - (Optional) An integer version tracker for `passplain_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslcertkeybundle. It has the same value as the `certkeybundlename` attribute.


## Import

A sslcertkeybundle can be imported using its name, e.g.

```shell
terraform import citrixadc_sslcertkeybundle.tf_certkeybundle web-certkey-bundle
```

Note: the `passplain` / `passplain_wo` pass phrase is never returned by the NITRO API and therefore cannot be recovered on import. After importing, set it again in your configuration if the bundle's private key is encrypted.
