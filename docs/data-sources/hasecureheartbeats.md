---
subcategory: "HA"
---

# Data Source: hasecureheartbeats

The hasecureheartbeats data source allows you to retrieve information about the
HA secure heartbeats parameters configuration.


## Example usage

```terraform
data "citrixadc_hasecureheartbeats" "tf_hasecureheartbeats" {
}

output "state" {
  value = data.citrixadc_hasecureheartbeats.tf_hasecureheartbeats.state
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `state` - By enabling this option, HA heartbeats are securely exchanged between nodes. Possible values: [ ENABLED, DISABLED ]
* `hapsk` - Pre shared key to be used for securing HA heartbeats. (Sensitive; returned by the appliance only in encrypted form or omitted.)
* `hapsk_wo` - Write-only (ephemeral) equivalent of `hapsk`. Never persisted to state.
* `hapsk_wo_version` - Version tracker for `hapsk_wo`.
* `id` - The id of the hasecureheartbeats. It is a system-generated identifier.
