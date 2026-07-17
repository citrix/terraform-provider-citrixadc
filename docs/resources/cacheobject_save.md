---
subcategory: "Integrated Caching"
---

# Resource: cacheobject_save

The cacheobject_save resource persists objects currently held in the Citrix ADC integrated cache to disk, so that cached content survives a reboot or can be replicated to the secondary node in a high-availability pair. It is an action-only resource: applying it invokes the NITRO `save` action against the integrated cache. This is useful for preserving a warm cache across appliance restarts or for pushing the cache contents to the secondary node.

This resource does not create, read, or manage a persistent object on the appliance. There is no NITRO GET endpoint for the save action, so there is no corresponding data source. Each apply performs the save; because every input attribute is marked `RequiresReplace`, changing any argument destroys the resource from state and re-fires the action with the new inputs. Read and Delete are no-ops.

Unlike the expire and flush actions, `save` has no mandatory arguments: with no arguments a save-all is performed. Optionally, `locator` and `tosecondary` may be supplied.

To re-run the action (for example, to save the cache again), taint the resource or bump a distinguishing input value so Terraform re-creates it.

> The integrated caching (IC) feature must be enabled on the Citrix ADC before this action can be performed (`enable ns feature IC`).


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

* `id` - A synthetic identifier for this action-only resource. It is a fixed string with the value `cacheobject_save`. It does not correspond to any object on the Citrix ADC.
