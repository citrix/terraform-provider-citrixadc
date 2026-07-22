---
subcategory: "Integrated Caching"
---

# Resource: cacheobject_expire

The cacheobject_expire resource forcibly expires objects held in the Citrix ADC integrated cache so that the next matching request is revalidated against the origin server instead of being served stale. It is an action-only resource: applying it invokes the NITRO `expire` action against the integrated cache. This is useful after publishing new content when you want cached copies of a specific URL, host, or object locator to be marked expired immediately.

Each apply performs the expire; changing any argument re-fires the action with the new inputs.

To re-run the action (for example, to expire the same object again), taint the resource or bump a distinguishing input value so Terraform re-creates it.

> The integrated caching (IC) feature must be enabled on the Citrix ADC before this action can be performed (`enable ns feature IC`).

You must supply either `locator` OR the combination of `url` and `host` (optionally refined with `port`, `groupname`, and `httpmethod`), but not both. This choice is enforced at plan time.


## Example usage

### Expire a single cached object by locator

```hcl
resource "citrixadc_cacheobject_expire" "expire_by_locator" {
  locator = 12345
}
```

### Expire a cached object by URL and host

```hcl
resource "citrixadc_cacheobject_expire" "expire_by_url" {
  url        = "/images/logo.png"
  host       = "www.example.com"
  port       = 80
  httpmethod = "GET"
}
```


## Argument Reference

You must specify either `locator` OR both `url` and `host` (not both forms together). Changing any argument forces the action to be re-fired.

* `locator` - (Optional) ID of the cached object. Supply either `locator` OR (`url` + `host`), not both. Changing this attribute forces the action to be re-fired.
* `url` - (Optional) URL of the particular object to expire. `host` must be specified along with the URL. Changing this attribute forces the action to be re-fired.
* `host` - (Optional) Host name of the object. `url` must also be specified. Changing this attribute forces the action to be re-fired.
* `port` - (Optional) Host port of the object. You must also set `host`. Changing this attribute forces the action to be re-fired.
* `groupname` - (Optional) Name of the content group to which the object belongs. Restricts the action to objects belonging to the specified content group. You must also set `host`. Changing this attribute forces the action to be re-fired.
* `httpmethod` - (Optional) HTTP request method that caused the object to be stored. Changing this attribute forces the action to be re-fired. Possible values: [ `GET`, `POST` ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cacheobject_expire resource. It is set to `cacheobject_expire`.
