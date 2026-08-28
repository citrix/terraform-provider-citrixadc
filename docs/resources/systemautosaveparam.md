---
subcategory: "System"
---

# Resource: systemautosaveparam

The systemautosaveparam resource is used to configure the system autosave
parameters. This is a singleton (global) configuration resource.


## Example usage

```hcl
resource "citrixadc_systemautosaveparam" "tf_systemautosaveparam" {
  status                = "ENABLED"
  periodicsave          = "ENABLED"
  periodicsavefrequency = 1440
}
```


## Argument Reference

* `status` - (Optional) Configure autosave feature. Available options are: `DEFAULT` - NetScaler decides the default option for the autosave feature. `DISABLED` - Autosave feature is disabled. `ENABLED` - Autosave feature is enabled. Possible values: [ DEFAULT, DISABLED, ENABLED ]
* `periodicsave` - (Optional) Enable or disable periodic save of autosave configuration. If enabled, `saveconfig` will be done periodically for all partitions including default. Possible values: [ ENABLED, DISABLED ]
* `periodicsavefrequency` - (Optional) Frequency in multiple of 60 minutes for periodic save of autosave configuration. Default value is 720 minutes. Minimum value = 60, Maximum value = 7200.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemautosaveparam. Because this is a singleton resource, it has a fixed identifier.
