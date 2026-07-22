---
subcategory: "SSL"
---

# Resource: sslcertkeybundle_change

The sslcertkeybundle_change resource re-reads and updates an existing SSL certificate-key bundle on the Citrix ADC from its on-disk bundle file. Use it when the underlying X509 certificate bundle file (for example, after a certificate renewal) has been replaced on the appliance and you want the ADC to reload the certificate, private-key, and any intermediate certificates from that file without deleting and re-adding the `sslcertkeybundle` object.

~> **One-shot action.** This resource performs the `change` action for `sslcertkeybundle` (CLI: `update ssl certKeyBundle <certkeybundlename>`); it does not create a new persistent object on the appliance. Each `terraform apply` that creates or replaces this resource performs the change once, and changing any argument forces a new change (replacement).


## Example usage

```hcl
resource "citrixadc_sslcertkeybundle_change" "tf_sslcertkeybundle_change" {
  certkeybundlename = "servercertbundle1"
  bundlefile        = "/nsconfig/ssl/servercertbundle1.pem"
  passplain         = var.sslcertkeybundle_change_passplain
}

variable "sslcertkeybundle_change_passplain" {
  type      = string
  sensitive = true
}
```


## Argument Reference

* `certkeybundlename` - (Required) Name given to the certKeyBundle. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Maximum length: 127. Changing this value forces the resource to be recreated (re-running the change action against the new bundle).
* `bundlefile` - (Optional) Name of and, optionally, path to the X509 certificate bundle file that is used to form the certificate-key bundle. `/nsconfig/ssl/` is the default path. Maximum length: 255. Changing this value forces the resource to be recreated.
* `passplain` - (Optional, Sensitive) Pass phrase used to encrypt the private-key. Required when the certificate bundle file contains an encrypted private-key in PEM format. Maximum length: 31. Changing this value forces the resource to be recreated.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the sslcertkeybundle_change resource. It has the format `sslcertkeybundle_change-<certkeybundlename>` (for example, `sslcertkeybundle_change-servercertbundle1`).
