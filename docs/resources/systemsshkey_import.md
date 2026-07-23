---
subcategory: "System"
---

# Resource: systemsshkey_import

This resource is used to import an SSH key onto the Citrix ADC from a source URL.


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
* `src` - (Required) URL (protocol, host, path, and file name) from where the SSH key file is imported. Note: the import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access. Changing this value forces a new resource to be created.
* `sshkeytype` - (Required) The type of the SSH key, whether public or private. Changing this value forces a new resource to be created. Possible values: [ PRIVATE, PUBLIC ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The composite identifier of the imported SSH key. It is the concatenation of the URL-encoded `name` and `sshkeytype` values in the form `name:<name>,sshkeytype:<sshkeytype>` (for example, `name:id_rsa_admin,sshkeytype:PRIVATE`).


## Import

A systemsshkey_import can be imported using its composite ID, in the form `name:<name>,sshkeytype:<sshkeytype>`:

```shell
terraform import citrixadc_systemsshkey_import.private_key name:id_rsa_admin,sshkeytype:PRIVATE
```

Note: `src` is not populated by import. Set it in configuration to match the original import source.
