---
subcategory: "Database"
---

# Data Source: dbuser

The dbuser data source allows you to retrieve information about database users on Citrix ADC.


## Example usage

```terraform
data "citrixadc_dbuser" "tf_dbuser_ds" {
  username = "dbuser1"
}

output "username" {
  value = data.citrixadc_dbuser.tf_dbuser_ds.username
}

output "loggedin" {
  value = data.citrixadc_dbuser.tf_dbuser_ds.loggedin
}
```


## Argument Reference

* `username` - (Required) Name of the database user. Must be the same as the user name specified in the database.
* `loggedin` - (Optional) Display the names of all database users currently logged on to the Citrix ADC.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dbuser. It has the same value as the `username` attribute.
* `password` - Password for logging on to the database. Must be the same as the password specified in the database.
* `password_wo_version` - Version tracker for the `password_wo` write-only attribute.
