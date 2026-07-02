---
subcategory: "Utility"
---

# Data Source: filesystemencryption

The filesystemencryption data source retrieves the file system encryption status of a Citrix ADC appliance. Use it to check whether the platform supports encryption (`supportedstate`) and whether the file system is currently encrypted (`effectivestate`) — for example, to gate the `citrixadc_filesystemencryption` resource on a supported platform.

This is a nameless singleton data source: no lookup key is required. On a cluster you may optionally target a specific node with `nodeid`.


## Example usage

```terraform
data "citrixadc_filesystemencryption" "example" {}

output "filesystemencryption_supportedstate" {
  value = data.citrixadc_filesystemencryption.example.supportedstate
}

output "filesystemencryption_effectivestate" {
  value = data.citrixadc_filesystemencryption.example.effectivestate
}
```

To query a specific cluster node:

```terraform
data "citrixadc_filesystemencryption" "node1" {
  nodeid = 1
}
```


## Argument Reference

* `nodeid` - (Optional) Unique number that identifies the cluster node to query.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the filesystemencryption data source (static string, `"filesystemencryption-config"`).
* `supportedstate` - Whether the platform supports file system encryption. Possible values: [ DISABLED, ENABLED, UNKNOWN ].
* `effectivestate` - The current encrypted state of the file system. Possible values: [ ENABLED, DISABLED ].
