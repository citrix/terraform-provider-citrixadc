---
subcategory: "System"
---

# Data Source: systemsshkey

Retrieves information about an SSH key that has been imported onto the Citrix ADC, looked up by its `name` and `sshkeytype`.

Note: the original source URL (`src`) is not returned by the NITRO GET API and is therefore never populated by this data source.


## Example usage

```terraform
data "citrixadc_systemsshkey" "private_key" {
  name       = "id_rsa_admin"
  sshkeytype = "PRIVATE"
}

output "ssh_key_id" {
  value = data.citrixadc_systemsshkey.private_key.id
}
```


## Argument Reference

* `name` - (Required) Name of the imported SSH key to look up on the Citrix ADC.
* `sshkeytype` - (Required) The type of the SSH key to look up. Possible values: [ PRIVATE, PUBLIC ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The composite identifier of the SSH key, in the form `name:<name>,sshkeytype:<sshkeytype>`.
* `src` - Always null. The source URL used to import the key is not returned by the NITRO GET API.
