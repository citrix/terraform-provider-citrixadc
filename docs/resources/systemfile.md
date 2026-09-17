---
subcategory: "System"
---

# Resource: systemfile

The systemfile resource is used to upload files to the target ADC.


## Example usage

```hcl
resource "citrixadc_systemfile" "tf_file" {
    filename = "hello.txt"
    filelocation = "/var/tmp"
    filecontent = "hello"
}
```

### Write-only content (ephemeral)

Use `filecontent_wo` instead of `filecontent` to keep sensitive file content (for
example, a private key sourced from Vault) **out of Terraform state**. The content is
supplied at apply time and is never written to state, plan, or logs. Because the
appliance has no in-place update for `systemfile`, a content change is signalled by
incrementing `filecontent_wo_version`, which re-uploads the file (destroy + recreate).

```hcl
variable "key_pem" {
  type      = string
  sensitive = true
}

resource "citrixadc_systemfile" "tf_key" {
    filename               = "server.key"
    filelocation           = "/nsconfig/ssl"
    filecontent_wo         = var.key_pem
    filecontent_wo_version = 1 # bump this to re-upload when the content changes
}
```

> **Note:** `filecontent_wo` keeps the content out of Terraform state, but the ADC still
> returns the file content on a NITRO GET (so the `citrixadc_systemfile` **data source**
> and any NITRO caller can still read it). This is state hygiene, not end-to-end secrecy.


## Argument Reference

Exactly one of `filecontent` or `filecontent_wo` must be set.

* `filename` - (Optional) Name of the file. It should not include filepath.
* `filecontent` - (Optional) File content (plaintext, or base64 when `is_base64_encoded` is true). Stored in Terraform state. Mutually exclusive with `filecontent_wo`.
* `filecontent_wo` - (Optional, [Write-only](https://developer.hashicorp.com/terraform/language/resources/ephemeral/write-only)) File content that is **not** stored in Terraform state. Requires Terraform 1.11+. Change `filecontent_wo_version` to signal an update. Mutually exclusive with `filecontent`.
* `filecontent_wo_version` - (Optional) Version tracker for `filecontent_wo`. Increment it to trigger a re-upload of the file. Defaults to `1`.
* `filelocation` - (Optional) Location of the file on Citrix ADC.
* `fileencoding` - (Optional) Encoding type of the file content. Defaults to `BASE64`.
* `is_base64_encoded` - (Optional) Set to true when the content is already base64 encoded; otherwise the provider base64-encodes it before upload. Applies to both `filecontent` and `filecontent_wo`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemfile. It is the fullpath of the system file.


## Import

A systemfile can be imported using its full path, e.g.

```shell
terraform import citrixadc_systemfile.tf_file /var/tmp/hello.txt
```

Import is supported for the plaintext `filecontent` path. A file managed with
`filecontent_wo` should not be imported: the write-only content cannot be recovered
into state, and importing it would transiently place the content in state before the
next apply.
