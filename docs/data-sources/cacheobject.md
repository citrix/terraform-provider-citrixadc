---
subcategory: "Integrated Caching"
---

# Data Source: cacheobject

The cacheobject data source lets you retrieve details about objects currently stored in the Citrix ADC integrated cache (equivalent to `show cache object`). Because the NITRO API exposes only a get-all endpoint, the data source fetches all cached objects and returns the first one matching the filters you supply. All returned attributes are read-only.

> The integrated caching (IC) feature must be enabled on the Citrix ADC (`enable ns feature IC`).


## Example usage

### Look up a cached object by URL and host

```terraform
data "citrixadc_cacheobject" "example" {
  url  = "/images/logo.png"
  host = "www.example.com"
}

output "cacheobject_locator" {
  value = data.citrixadc_cacheobject.example.locator
}

output "cacheobject_httpstatus" {
  value = data.citrixadc_cacheobject.example.httpstatus
}
```

### Look up a cached object by locator

```terraform
data "citrixadc_cacheobject" "by_locator" {
  locator = 12345
}
```


## Argument Reference

The following filter attributes are optional. When set, they narrow the get-all result to the first matching cached object. If none are specified, the first cached object returned by the ADC is used.

* `group` - (Optional) Name of the content group whose objects should be listed.
* `groupname` - (Optional) Name of the content group to which the object belongs. Restricts results to objects belonging to the specified content group. `host` must also be set.
* `host` - (Optional) Host name of the object. `url` must also be specified.
* `url` - (Optional) URL of the particular object whose details are required. `host` must be specified along with the URL.
* `port` - (Optional) Host port of the object. `host` must also be set.
* `httpmethod` - (Optional) HTTP request method that caused the object to be stored.
* `httpstatus` - (Optional) HTTP status of the object.
* `ignoremarkerobjects` - (Optional) Ignore marker objects when filtering.
* `includenotreadyobjects` - (Optional) Include responses that have not yet reached a minimum number of hits before being cached.
* `locator` - (Optional) ID of the cached object.
* `nodeid` - (Optional) Unique number that identifies the cluster node.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier for the matched object, derived from `locator` (for example, `locator:12345`) or `url` (for example, `url:/images/logo.png`) when available.
* `action` - Not applicable for reads (action-only on the resource); present for model compatibility.
* `group` - Name of the content group whose objects were listed.
* `groupname` - Name of the content group to which the object belongs.
* `host` - Host name of the object.
* `url` - URL of the object.
* `port` - Host port of the object.
* `httpmethod` - HTTP request method that caused the object to be stored.
* `httpstatus` - HTTP status of the object.
* `ignoremarkerobjects` - Whether marker objects are ignored.
* `includenotreadyobjects` - Whether responses not yet reaching the minimum hit count are included.
* `locator` - ID of the cached object.
* `nodeid` - Unique number that identifies the cluster node.
* `tosecondary` - Whether the object is saved onto the secondary node.
