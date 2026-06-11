---
subcategory: "Policy"
---

# Data Source: policypatsetfile

The policypatsetfile data source allows you to retrieve information about a pattern-set file that has been imported into the Citrix ADC.


## Example usage

```terraform
data "citrixadc_policypatsetfile" "example" {
  name = "tf_patsetfile"
}

output "patsetfile_src" {
  value = data.citrixadc_policypatsetfile.example.src
}
```


## Argument Reference

* `name` - (Required) Name of the imported patset file to look up. Unique name of the pattern set. Not case sensitive. Must begin with an ASCII letter or underscore (_) character and must contain only alphanumeric and underscore characters.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the policypatsetfile. It has the same value as the `name` attribute.
* `charset` - Character set associated with the characters in the string.
* `comment` - Any comments to preserve information about this patsetfile.
* `delimiter` - Patset file patterns delimiter.
* `overwrite` - Overwrites the existing file.
* `src` - URL in protocol, host, path, and file name format from where the patset file was imported.
