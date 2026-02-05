---
subcategory: "DB"
---

# Data Source: citrixadc_dbuser

The dbuser data source allows you to retrieve information about database users.


## Example usage

```terraform
resource "citrixadc_dbuser" "tf_dbuser" {
  username = "user1"
  password = "1234"
}

data "citrixadc_dbuser" "tf_dbuser_ds" {
  username = citrixadc_dbuser.tf_dbuser.username
  loggedin = false
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
* `loggedin` - (Required) Display the names of all database users currently logged on to the Citrix ADC.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dbuser. It has the same value as the `username` attribute.
* `password` - Password for logging on to the database. Must be the same as the password specified in the database.
