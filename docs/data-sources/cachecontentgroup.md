---
subcategory: "Integrated Caching"
---

# Data Source `cachecontentgroup`

The cachecontentgroup data source allows you to retrieve information about an existing cachecontentgroup.


## Example usage

```terraform
data "citrixadc_cachecontentgroup" "tf_cachecontentgroup" {
  name = "my_cachecontentgroup"
}

output "name" {
  value = data.citrixadc_cachecontentgroup.tf_cachecontentgroup.name
}

output "heurexpiryparam" {
  value = data.citrixadc_cachecontentgroup.tf_cachecontentgroup.heurexpiryparam
}

output "prefetch" {
  value = data.citrixadc_cachecontentgroup.tf_cachecontentgroup.prefetch
}

output "quickabortsize" {
  value = data.citrixadc_cachecontentgroup.tf_cachecontentgroup.quickabortsize
}
```


## Argument Reference

* `name` - (Required) Name for the content group. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the content group is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cachecontentgroup. It has the same value as the `name` attribute.
* `absexpiry` - Local time, up to 4 times a day, at which all objects in the content group must expire.
* `absexpirygmt` - Coordinated Universal Time (GMT), up to 4 times a day, when all objects in the content group must expire.
* `alwaysevalpolicies` - Force policy evaluation for each response arriving from the origin server. Cannot be set to YES if the Prefetch parameter is also set to YES.
* `cachecontrol` - Insert a Cache-Control header into the response.
* `expireatlastbyte` - Force expiration of the content immediately after the response is downloaded (upon receipt of the last byte of the response body). Applicable only to positive responses.
* `flashcache` - Perform flash cache. Mutually exclusive with Poll Every Time (PET) on the same content group.
* `heurexpiryparam` - Heuristic expiry time, in percent of the duration, since the object was last modified.
* `hitparams` - Parameters to use for parameterized hit evaluation of an object. Up to 128 parameters can be specified. Mutually exclusive with the Hit Selector parameter.
* `hitselector` - Selector for evaluating whether an object gets stored in a particular content group. A selector is an abstraction for a collection of PIXL expressions.
* `host` - Flush only objects that belong to the specified host. Do not use except with parameterized invalidation. Also, the Invalidation Restricted to Host parameter for the group must be set to YES.
* `ignoreparamvaluecase` - Ignore case when comparing parameter values during parameterized hit evaluation. (Parameter value case is ignored by default during parameterized invalidation.)
* `ignorereloadreq` - Ignore any request to reload a cached object from the origin server. To guard against Denial of Service attacks, set this parameter to YES. For RFC-compliant behavior, set it to NO.
* `ignorereqcachinghdrs` - Ignore Cache-Control and Pragma headers in the incoming request.
* `insertage` - Insert an Age header into the response. An Age header contains information about the age of the object, in seconds, as calculated by the integrated cache.
* `insertetag` - Insert an ETag header in the response. With ETag header insertion, the integrated cache does not serve full responses on repeat requests.
* `insertvia` - Insert a Via header into the response.
* `invalparams` - Parameters for parameterized invalidation of an object. You can specify up to 8 parameters. Mutually exclusive with invalSelector.
* `invalrestrictedtohost` - Take the host header into account during parameterized invalidation.
* `invalselector` - Selector for invalidating objects in the content group. A selector is an abstraction for a collection of PIXL expressions.
* `lazydnsresolve` - Perform DNS resolution for responses only if the destination IP address in the request does not match the destination IP address of the cached response.
* `matchcookies` - Evaluate for parameters in the cookie header also.
* `maxressize` - Maximum size of a response that can be cached in this content group.
* `memlimit` - Maximum amount of memory that the cache can use. The effective limit is based on the available memory of the Citrix ADC.
* `minhits` - Number of hits that qualifies a response for storage in this content group.
* `minressize` - Minimum size of a response that can be cached in this content group. Default minimum response size is 0.
* `persistha` - Setting persistHA to YES causes IC to save objects in contentgroup to Secondary node in HA deployment.
* `pinned` - Do not flush objects from this content group under memory pressure.
* `polleverytime` - Always poll for the objects in this content group. That is, retrieve the objects from the origin server whenever they are requested.
* `prefetch` - Attempt to refresh objects that are about to go stale.
* `prefetchmaxpending` - Maximum number of outstanding prefetches that can be queued for the content group.
* `prefetchperiod` - Time period, in seconds before an object's calculated expiry time, during which to attempt prefetch.
* `prefetchperiodmillisec` - Time period, in milliseconds before an object's calculated expiry time, during which to attempt prefetch.
* `query` - Query string specifying individual objects to flush from this group by using parameterized invalidation. If this parameter is not set, all objects are flushed from the group.
* `quickabortsize` - If the size of an object that is being downloaded is less than or equal to the quick abort value, and a client aborts during the download, the cache stops downloading the response. If the object is larger than the quick abort size, the cache continues to download the response.
* `relexpiry` - Relative expiry time, in seconds, after which to expire an object cached in this content group.
* `relexpirymillisec` - Relative expiry time, in milliseconds, after which to expire an object cached in this content group.
* `removecookies` - Remove cookies from responses.
* `selectorvalue` - Value of the selector to be used for flushing objects from the content group. Requires that an invalidation selector be configured for the content group.
* `tosecondary` - Content group whose objects are to be sent to secondary.
* `type` - The type of the content group.
* `weaknegrelexpiry` - Relative expiry time, in seconds, for expiring negative responses. This value is used only if the expiry time cannot be determined from any other source. It is applicable only to the following status codes: 307, 403, 404, and 410.
* `weakposrelexpiry` - Relative expiry time, in seconds, for expiring positive responses with response codes between 200 and 399. Cannot be used in combination with other Expiry attributes. Similar to -relExpiry but has lower precedence.
