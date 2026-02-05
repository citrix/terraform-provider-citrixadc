---
subcategory: "AppQoE"
---

# Data Source `appqoeparameter`

The appqoeparameter data source allows you to retrieve information about AppQoE parameters configuration.


## Example usage

```terraform
data "citrixadc_appqoeparameter" "tf_appqoeparameter" {
}

output "sessionlife" {
  value = data.citrixadc_appqoeparameter.tf_appqoeparameter.sessionlife
}

output "avgwaitingclient" {
  value = data.citrixadc_appqoeparameter.tf_appqoeparameter.avgwaitingclient
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `sessionlife` - Time, in seconds, between the first time and the next time the AppQoE alternative content window is displayed. The alternative content window is displayed only once during a session for the same browser accessing a configured URL, so this parameter determines the length of a session.
* `avgwaitingclient` - Average number of client connections, that can sit in service waiting queue.
* `maxaltrespbandwidth` - Maximum bandwidth which will determine whether to send alternate content response.
* `dosattackthresh` - Average number of client connection that can queue up on vserver level without triggering DoS mitigation module.

## Attribute Reference

* `id` - The id of the appqoeparameter. It is a system-generated identifier.
