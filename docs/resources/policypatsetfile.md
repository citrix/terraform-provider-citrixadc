---
subcategory: "Policy"
---

# Resource: policypatsetfile

The policypatsetfile resource imports a pattern-set file into the Citrix ADC. Importing the file loads a collection of patterns from an external source (an HTTP/HTTPS/FTP URL or a file already present on the appliance) and registers it under a named pattern set, which can then be referenced in policy expressions for pattern-set matching.

Importing the file is a one-shot operation. Every configurable attribute forces replacement, so any change to the configuration destroys the existing imported file and re-imports it. Destroying the resource removes the imported pattern-set file from the appliance.


## Example usage

```hcl
resource "citrixadc_policypatsetfile" "tf_policypatsetfile" {
  name      = "tf_patsetfile"
  src       = "http://10.0.0.1/patterns/blocklist.txt"
  overwrite = true
  delimiter = "10"
  charset   = "UTF_8"
  comment   = "Imported blocklist patterns"
}
```

Importing a file that already resides on the appliance using the `local` keyword:

```hcl
resource "citrixadc_policypatsetfile" "tf_local_patsetfile" {
  name = "tf_local_patsetfile"
  src  = "local:blocklist.txt"
}
```


## Argument Reference

* `name` - (Required) Name to assign to the imported patset file. Unique name of the pattern set. Not case sensitive. Must begin with an ASCII letter or underscore (_) character and must contain only alphanumeric and underscore characters. Changing this attribute forces a new resource to be created.
* `src` - (Required) URL in protocol, host, path, and file name format from where the patset file will be imported. If the file is already present on the appliance, it can be imported using the `local` keyword (for example, `local:filename`). Note: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access. Changing this attribute forces a new resource to be created.
* `charset` - (Optional) Character set associated with the characters in the string. Possible values: [ ASCII, UTF_8 ]. Changing this attribute forces a new resource to be created.
* `comment` - (Optional) Any comments to preserve information about this patsetfile. Changing this attribute forces a new resource to be created.
* `delimiter` - (Optional) Patset file patterns delimiter. Defaults to `10`. Changing this attribute forces a new resource to be created.
* `overwrite` - (Optional) Overwrites the existing file. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the policypatsetfile. It has the same value as the `name` attribute.


## Import

A policypatsetfile can be imported using its name, e.g.

```shell
terraform import citrixadc_policypatsetfile.tf_policypatsetfile tf_patsetfile
```
