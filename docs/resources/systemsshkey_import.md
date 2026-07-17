---
subcategory: "System"
---

# Resource: systemsshkey_import

Imports an SSH key (public or private) onto the Citrix ADC from a remote or local source URL. SSH keys are used to establish trusted, password-less SSH connectivity between the ADC and external hosts (for example, for secure file transfers or scripted administration). Creating this resource performs the NITRO `Import` action against the provided source location; deleting it removes the imported key from the appliance.

Because the appliance only exposes an `Import` action (not a conventional create) and offers no update endpoint, all attributes force replacement when changed. The `src` source URL is consumed at import time and is **not** returned by the NITRO GET API, so it is never read back into state.


## Example usage

```hcl
resource "citrixadc_systemsshkey_import" "private_key" {
  name       = "id_rsa_admin"
  src        = "http://192.0.2.10/keys/id_rsa_admin"
  sshkeytype = "PRIVATE"
}
```

Importing a public key from a local appliance path:

```hcl
resource "citrixadc_systemsshkey_import" "public_key" {
  name       = "id_rsa_admin.pub"
  src        = "local://id_rsa_admin.pub"
  sshkeytype = "PUBLIC"
}
```


## Argument Reference

The following arguments are supported:

* `name` - (Required) Name to assign to the imported SSH key on the Citrix ADC. Changing this value forces a new resource to be created.
* `src` - (Required) URL (protocol, host, path, and file name) from where the SSH key file is imported. This value is the import source only: it is consumed by the `Import` action and is **not** returned by the NITRO GET API, so the provider preserves the user-configured value in state and never reads it back. Note: the import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access. Changing this value forces a new resource to be created.
* `sshkeytype` - (Required) The type of the SSH key, whether public or private. Changing this value forces a new resource to be created. Possible values: [ PRIVATE, PUBLIC ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The composite identifier of the imported SSH key. It is the concatenation of the URL-encoded `name` and `sshkeytype` values in the form `name:<name>,sshkeytype:<sshkeytype>` (for example, `name:id_rsa_admin,sshkeytype:PRIVATE`).


## Import

A systemsshkey_import can be imported using its composite ID, in the form `name:<name>,sshkeytype:<sshkeytype>`:

```shell
terraform import citrixadc_systemsshkey_import.private_key name:id_rsa_admin,sshkeytype:PRIVATE
```

Note: `src` is not returned by the NITRO API and therefore is not populated by import. Set it in configuration to match the original import source.
