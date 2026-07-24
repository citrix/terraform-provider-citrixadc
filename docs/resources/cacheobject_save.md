---
subcategory: "Integrated Caching"
---

# Resource: cacheobject_save

This resource is used to save objects from the Citrix ADC integrated cache to disk.


## Example usage

### Save all cached objects

```hcl
resource "citrixadc_cacheobject_save" "save_all" {}
```

### Save a cached object to the secondary node

```hcl
resource "citrixadc_cacheobject_save" "save_to_secondary" {
  locator     = 12345
  tosecondary = "YES"
}
```


## Argument Reference

All arguments are optional; with none supplied, a save-all is performed. Changing any argument forces the action to be re-fired.

* `locator` - (Optional) ID of the cached object. Changing this attribute forces the action to be re-fired.
* `tosecondary` - (Optional) Save the object onto the secondary node. Applies only to the `save` action. Changing this attribute forces the action to be re-fired. Possible values: [ `YES`, `NO` ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cacheobject_save resource. It is set to `cacheobject_save`.
