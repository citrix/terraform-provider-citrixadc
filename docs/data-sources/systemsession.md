---
subcategory: "System"
---

# Data Source: systemsession

Retrieves information about an administrative session on the Citrix ADC, looked up by its session ID (`sid`). This is the read-only counterpart to the `citrixadc_systemsession` resource (which kills sessions): use it to inspect details such as the logged-in user, login time, and client IP address of an active session.


## Example usage

```terraform
data "citrixadc_systemsession" "session" {
  sid = 12
}

output "session_user" {
  value = data.citrixadc_systemsession.session.username
}

output "session_client_ip" {
  value = data.citrixadc_systemsession.session.clientipaddress
}
```


## Argument Reference

* `sid` - (Required) ID of the system session about which to display information.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The identifier of the data source lookup.
* `username` - Name of the user who is logged in.
* `logintime` - Time when the user logged in.
* `logintimelocal` - Time (local) when the user logged in.
* `lastactivitytime` - Time of last activity in the session.
* `lastactivitytimelocal` - Time (local) of last activity in the session.
* `expirytime` - Time when the session expires.
* `numofconnections` - Number of connections in the session.
* `currentconn` - Indicates the current connection.
* `clienttype` - Type of client used for the session.
* `partitionname` - Name of the partition for the session.
* `clientipaddress` - Client IP address for the session.
