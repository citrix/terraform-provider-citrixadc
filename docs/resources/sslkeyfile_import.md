---
subcategory: "SSL"
---

# Resource: sslkeyfile_import

The sslkeyfile_import resource imports a private key file onto the Citrix ADC so it can be paired with a certificate to terminate SSL traffic. The key file is fetched from a remote source at creation time and stored under the given name on the appliance. If the key is encrypted, the passphrase used to decrypt it can be supplied via the `password` attributes.

This resource maps to the NITRO `Import` action for the `sslkeyfile` object. The ADC exposes this object only through that import action (there is no in-place update endpoint), so every attribute forces a new resource when changed.


## Example usage

### Using password (sensitive attribute - persisted in state)

```hcl
variable "sslkeyfile_import_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslkeyfile_import" "tf_sslkeyfile_import" {
  name     = "servercert1key"
  src      = "http://www.example.com/key_file.pem"
  password = var.sslkeyfile_import_password
}
```

### Using password_wo (write-only/ephemeral - NOT persisted in state)

The `password_wo` attribute provides an ephemeral path for the key file passphrase. The value is sent to the Citrix ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the value changes, increment `password_wo_version`.

```hcl
variable "sslkeyfile_import_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslkeyfile_import" "tf_sslkeyfile_import" {
  name                = "servercert1key"
  src                 = "http://www.example.com/key_file.pem"
  password_wo         = var.sslkeyfile_import_password
  password_wo_version = 1
}
```

To rotate the secret, update the variable value and bump the version:

```hcl
resource "citrixadc_sslkeyfile_import" "tf_sslkeyfile_import" {
  name                = "servercert1key"
  src                 = "http://www.example.com/key_file.pem"
  password_wo         = var.sslkeyfile_import_password
  password_wo_version = 2  # Bumped to trigger replacement
}
```

A key file that is not passphrase-protected can be imported without any of the `password` attributes:

```hcl
resource "citrixadc_sslkeyfile_import" "tf_sslkeyfile_import" {
  name = "servercert1key"
  src  = "http://www.example.com/key_file.pem"
}
```


## Argument Reference

* `name` - (Required) Name to assign to the imported key file. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my file" or 'my file'). Changing this value forces a new resource to be created.
* `src` - (Required) URL specifying the protocol, host, and path, including file name, to the key file to be imported. For example, `http://www.example.com/key_file`. This is the import source consumed at creation time; the NITRO GET response does not faithfully echo it back, so the provider preserves the user-configured value in state. Changing this value forces a new resource to be created. Note: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access, and the issuer certificate of the HTTPS server is not present in the specific path on the Citrix ADC to authenticate the HTTPS server.
* `password` - (Optional, Sensitive) Passphrase used to decrypt the imported private key, if the key file is encrypted. The value is persisted in Terraform state (encrypted at rest only if your backend supports it). See also `password_wo` for an ephemeral alternative. Changing this value forces a new resource to be created.
* `password_wo` - (Optional, Sensitive, WriteOnly) Same as `password`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `password_wo_version`. If both `password` and `password_wo` are set, `password_wo` takes precedence. Changing this value forces a new resource to be created.
* `password_wo_version` - (Optional) An integer version tracker for `password_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger replacement. Defaults to `1`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslkeyfile_import resource. It has the same value as the `name` attribute.


## Import

An sslkeyfile_import resource can be imported using its name, e.g.

```shell
terraform import citrixadc_sslkeyfile_import.tf_sslkeyfile_import servercert1key
```

Note: The `password` / `password_wo` secret is never returned by the NITRO API, so it cannot be recovered through import. After importing, ensure your configuration supplies the passphrase if the key is encrypted.
