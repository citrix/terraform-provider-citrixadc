---
subcategory: "Integrated Caching"
---

# Resource: cacheobject_flush

This resource is used to flush objects from the Citrix ADC integrated cache.


## Example usage

### Flush a single cached object by locator

```hcl
resource "citrixadc_cacheobject_flush" "flush_by_locator" {
  locator = 12345
}
```

### Flush a cached object by URL and host

```hcl
resource "citrixadc_cacheobject_flush" "flush_by_url" {
  url        = "/images/logo.png"
  host       = "www.example.com"
  port       = 80
  httpmethod = "GET"
}
```


## Argument Reference

You must specify either `locator` OR both `url` and `host` (not both forms together). Changing any argument forces the action to be re-fired.

* `locator` - (Optional) ID of the cached object. Supply either `locator` OR (`url` + `host`), not both. Changing this attribute forces the action to be re-fired.
* `url` - (Optional) URL of the particular object to flush. `host` must be specified along with the URL. Changing this attribute forces the action to be re-fired.
* `host` - (Optional) Host name of the object. `url` must also be specified. Changing this attribute forces the action to be re-fired.
* `port` - (Optional) Host port of the object. You must also set `host`. Changing this attribute forces the action to be re-fired.
* `groupname` - (Optional) Name of the content group to which the object belongs. Restricts the action to objects belonging to the specified content group. You must also set `host`. Changing this attribute forces the action to be re-fired.
* `httpmethod` - (Optional) HTTP request method that caused the object to be stored. Changing this attribute forces the action to be re-fired. Possible values: [ `GET`, `POST` ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cacheobject_flush resource. It is set to `cacheobject_flush`.
