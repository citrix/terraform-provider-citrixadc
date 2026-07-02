---
subcategory: "Integrated Caching"
---

# Resource: cacheobject

The cacheobject resource operates on the runtime objects stored in the Citrix ADC integrated cache. Unlike most resources, `cacheobject` does not manage a persistent configuration object. Cached objects are created and evicted by the traffic engine, not by the configuration API. The NITRO API exposes only read (get-all/count) and the action verbs `expire`, `flush`, and `save`. This resource therefore fires a single cache action when it is applied.

## Action-only behavior

Because there is no add/set/delete of a persistent object, this resource behaves differently from typical Terraform resources:

* **Apply fires one action.** When the resource is created, it performs the configured `action` (`expire`, `flush`, or `save`) against the integrated cache.
* **Read is a no-op.** A fired action leaves no re-findable object keyed to the resource's synthetic ID, so Terraform preserves prior state unchanged and never detects drift.
* **Update forces re-creation.** Every input attribute is marked `RequiresReplace`. Changing any input destroys the resource from state and re-fires the action with the new inputs.
* **Delete is a no-op.** The prior action is not reversible through the configuration API; removing the resource simply drops it from Terraform state.

To re-run an action (for example, to flush the cache again), either taint the resource or bump a distinguishing input value so Terraform re-creates it.

> The integrated caching (IC) feature must be enabled on the Citrix ADC before these actions can be performed (`enable ns feature IC`).


## Example usage

### Flush a single cached object by locator

```hcl
resource "citrixadc_cacheobject" "tf_flush_by_locator" {
  action  = "flush"
  locator = 12345
}
```

### Expire cached objects by URL and host

For `expire` and `flush` you must supply either `locator` OR the combination of `url` and `host` (optionally refined with `port`, `groupname`, and `httpmethod`).

```hcl
resource "citrixadc_cacheobject" "tf_expire_by_url" {
  action     = "expire"
  url        = "/images/logo.png"
  host       = "www.example.com"
  port       = 80
  httpmethod = "GET"
}
```

### Save cached objects to disk

The `save` action does not require `locator` or `url`/`host`; a save-all is allowed with no additional arguments. Optionally, `locator` and `tosecondary` can be supplied.

```hcl
resource "citrixadc_cacheobject" "tf_save" {
  action      = "save"
  tosecondary = "NO"
}
```


## Argument Reference

* `action` - (Optional) The cache action to perform on the cached object(s). Because this is an action-only runtime object (not a persistent configuration), applying the resource fires this action. Changing this attribute forces the action to be re-fired. Defaults to `"flush"`. Possible values: [ `expire`, `flush`, `save` ]
* `locator` - (Optional) ID of the cached object. For `expire`/`flush`, supply either `locator` OR (`url` + `host`), not both. For `save`, `locator` is optional. Changing this attribute forces the action to be re-fired.
* `url` - (Optional) URL of the particular object to act on. `host` must be specified along with the URL. Used with `expire`/`flush`. Changing this attribute forces the action to be re-fired.
* `host` - (Optional) Host name of the object. `url` must also be specified. Used with `expire`/`flush`. Changing this attribute forces the action to be re-fired.
* `port` - (Optional) Host port of the object. `host` must also be set. Used with `expire`/`flush`. Changing this attribute forces the action to be re-fired.
* `groupname` - (Optional) Name of the content group to which the object belongs. Restricts the action to objects belonging to the specified content group. `host` must also be set. Used with `expire`/`flush`. Changing this attribute forces the action to be re-fired.
* `httpmethod` - (Optional) HTTP request method that caused the object to be stored. Used with `expire`/`flush`. Changing this attribute forces the action to be re-fired. Possible values: [ `GET`, `POST` ]
* `httpstatus` - (Optional) HTTP status of the object. Changing this attribute forces the action to be re-fired.
* `ignoremarkerobjects` - (Optional) Ignore marker objects. Marker objects are created when a response exceeds the maximum or minimum response size for the content group, or has not yet received the minimum number of hits for the content group. Changing this attribute forces the action to be re-fired. Possible values: [ `ON`, `OFF` ]
* `includenotreadyobjects` - (Optional) Include responses that have not yet reached a minimum number of hits before being cached. Changing this attribute forces the action to be re-fired. Possible values: [ `ON`, `OFF` ]
* `group` - (Optional) Name of the content group whose objects should be acted on. Changing this attribute forces the action to be re-fired.
* `tosecondary` - (Optional) Save the object onto the secondary node. Applies only to the `save` action. Changing this attribute forces the action to be re-fired. Possible values: [ `YES`, `NO` ]
* `nodeid` - (Optional) Unique number that identifies the cluster node. Changing this attribute forces the action to be re-fired.

### Constraints

* For `action = "expire"` or `action = "flush"`, you must specify either `locator` OR both `url` and `host` — but not both `locator` and (`url`/`host`) together. This is enforced at plan time.
* For `action = "save"`, no identifying arguments are required (save-all is allowed). `locator` and `tosecondary` may optionally be supplied.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier for the fired action. Because `cacheobject` is not a persistent object, the ID is derived from the action and the primary identity supplied (for example, `flush:locator:12345` or `expire:url:/images/logo.png`), so repeated applies with different inputs produce distinct IDs.
