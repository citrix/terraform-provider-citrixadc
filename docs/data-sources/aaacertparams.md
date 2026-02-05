---
subcategory: "AAA"
---

# Data Source `aaacertparams`

The aaacertparams data source allows you to retrieve information about AAA certificate parameters configuration.


## Example usage

```terraform
data "citrixadc_aaacertparams" "tf_aaacertparams" {
}

output "usernamefield" {
  value = data.citrixadc_aaacertparams.tf_aaacertparams.usernamefield
}

output "groupnamefield" {
  value = data.citrixadc_aaacertparams.tf_aaacertparams.groupnamefield
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `defaultauthenticationgroup` - This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
* `groupnamefield` - Client certificate field that specifies the group, in the format <field>:<subfield>.
* `usernamefield` - Client certificate field that contains the username, in the format <field>:<subfield>.

## Attribute Reference

* `id` - The id of the aaacertparams. It is a system-generated identifier.
