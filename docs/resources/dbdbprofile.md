---
subcategory: "Db"
---

# Resource: dbdbprofile

The dbdbprofile resource is used to create dbdbprofile.


## Example usage

```hcl
resource "citrixadc_dbdbprofile" "tf_dbdbprofile" {
  name           = "my_dbprofile"
  stickiness     = "YES"
  conmultiplex   = "ENABLED"
  interpretquery = "YES"
}
```


## Argument Reference

* `name` - (Required) Name for the database profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Cannot be changed after the profile is created. CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my profile" or 'my profile'). . Minimum length =  1 Maximum length =  127
* `interpretquery` - (Optional) If ENABLED, inspect the query and update the connection information, if required. If DISABLED, forward the query to the server. Possible values: [ YES, NO ]
* `stickiness` - (Optional) If the queries are related to each other, forward to the same backend server. Possible values: [ YES, NO ]
* `kcdaccount` - (Optional) Name of the KCD account that is used for Windows authentication. Minimum length =  1 Maximum length =  127
* `conmultiplex` - (Optional) Use the same server-side connection for multiple client-side requests. Default is enabled. Possible values: [ ENABLED, DISABLED ]
* `enablecachingconmuxoff` - (Optional) Enable caching when connection multiplexing is OFF. Possible values: [ ENABLED, DISABLED ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dbdbprofile. It has the same value as the `name` attribute.


## Import

A dbdbprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_dbdbprofile.tf_dbdbprofile my_dbprofile
```
