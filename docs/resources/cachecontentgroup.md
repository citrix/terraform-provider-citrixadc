---
subcategory: "Integrated Caching"
---

# Resource: cachecontentgroup

The cachecontentgroup resource is used to create cachecontentgroup.


## Example usage

```hcl
resource "citrixadc_cachecontentgroup" "tf_cachecontentgroup" {
  name                 = "my_cachecontentgroup"
  heurexpiryparam      = 30
  prefetch             = "YES"
  quickabortsize       = 40
  ignorereqcachinghdrs = "YES"
}
```


## Argument Reference

* `name` - (Required) Name for the content group.  Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the content group is created.
* `absexpiry` - (Optional) Local time, up to 4 times a day, at which all objects in the content group must expire.   CLI Users: For example, to specify that the objects in the content group should expire by 11:00 PM, type the following command: add cache contentgroup <contentgroup name> -absexpiry 23:00  To specify that the objects in the content group should expire at 10:00 AM, 3 PM, 6 PM, and 11:00 PM, type: add cache contentgroup <contentgroup name> -absexpiry 10:00 15:00 18:00 23:00
* `absexpirygmt` - (Optional) Coordinated Universal Time (GMT), up to 4 times a day, when all objects in the content group must expire.
* `alwaysevalpolicies` - (Optional) Force policy evaluation for each response arriving from the origin server. Cannot be set to YES if the Prefetch parameter is also set to YES.
* `cachecontrol` - (Optional) Insert a Cache-Control header into the response.
* `expireatlastbyte` - (Optional) Force expiration of the content immediately after the response is downloaded (upon receipt of the last byte of the response body). Applicable only to positive responses.
* `flashcache` - (Optional) Perform flash cache. Mutually exclusive with Poll Every Time (PET) on the same content group.
* `heurexpiryparam` - (Optional) Heuristic expiry time, in percent of the duration, since the object was last modified.
* `hitparams` - (Optional) Parameters to use for parameterized hit evaluation of an object. Up to 128 parameters can be specified. Mutually exclusive with the Hit Selector parameter.
* `hitselector` - (Optional) Selector for evaluating whether an object gets stored in a particular content group. A selector is an abstraction for a collection of PIXL expressions.
* `host` - (Optional) Flush only objects that belong to the specified host. Do not use except with parameterized invalidation. Also, the Invalidation Restricted to Host parameter for the group must be set to YES.
* `ignoreparamvaluecase` - (Optional) Ignore case when comparing parameter values during parameterized hit evaluation. (Parameter value case is ignored by default during parameterized invalidation.)
* `ignorereloadreq` - (Optional) Ignore any request to reload a cached object from the origin server. To guard against Denial of Service attacks, set this parameter to YES. For RFC-compliant behavior, set it to NO.
* `ignorereqcachinghdrs` - (Optional) Ignore Cache-Control and Pragma headers in the incoming request.
* `insertage` - (Optional) Insert an Age header into the response. An Age header contains information about the age of the object, in seconds, as calculated by the integrated cache.
* `insertetag` - (Optional) Insert an ETag header in the response. With ETag header insertion, the integrated cache does not serve full responses on repeat requests.
* `insertvia` - (Optional) Insert a Via header into the response.
* `invalparams` - (Optional) Parameters for parameterized invalidation of an object. You can specify up to 8 parameters. Mutually exclusive with invalSelector.
* `invalrestrictedtohost` - (Optional) Take the host header into account during parameterized invalidation.
* `invalselector` - (Optional) Selector for invalidating objects in the content group. A selector is an abstraction for a collection of PIXL expressions.
* `lazydnsresolve` - (Optional) Perform DNS resolution for responses only if the destination IP address in the request does not match the destination IP address of the cached response.
* `matchcookies` - (Optional) Evaluate for parameters in the cookie header also.
* `maxressize` - (Optional) Maximum size of a response that can be cached in this content group.
* `memlimit` - (Optional) Maximum amount of memory that the cache can use. The effective limit is based on the available memory of the Citrix ADC.
* `minhits` - (Optional) Number of hits that qualifies a response for storage in this content group.
* `minressize` - (Optional) Minimum size of a response that can be cached in this content group.  Default minimum response size is 0.
* `persistha` - (Optional) Setting persistHA to YES causes IC to save objects in contentgroup to Secondary node in HA deployment.
* `pinned` - (Optional) Do not flush objects from this content group under memory pressure.
* `polleverytime` - (Optional) Always poll for the objects in this content group. That is, retrieve the objects from the origin server whenever they are requested.
* `prefetch` - (Optional) Attempt to refresh objects that are about to go stale.
* `prefetchmaxpending` - (Optional) Maximum number of outstanding prefetches that can be queued for the content group.
* `prefetchperiod` - (Optional) Time period, in seconds before an object's calculated expiry time, during which to attempt prefetch.
* `prefetchperiodmillisec` - (Optional) Time period, in milliseconds before an object's calculated expiry time, during which to attempt prefetch.
* `query` - (Optional) Query string specifying individual objects to flush from this group by using parameterized invalidation. If this parameter is not set, all objects are flushed from the group.
* `quickabortsize` - (Optional) If the size of an object that is being downloaded is less than or equal to the quick abort value, and a client aborts during the download, the cache stops downloading the response. If the object is larger than the quick abort size, the cache continues to download the response.
* `relexpiry` - (Optional) Relative expiry time, in seconds, after which to expire an object cached in this content group.
* `relexpirymillisec` - (Optional) Relative expiry time, in milliseconds, after which to expire an object cached in this content group.
* `removecookies` - (Optional) Remove cookies from responses.
* `selectorvalue` - (Optional) Value of the selector to be used for flushing objects from the content group. Requires that an invalidation selector be configured for the content group.
* `tosecondary` - (Optional) content group whose objects are to be sent to secondary.
* `type` - (Optional) The type of the content group.
* `weaknegrelexpiry` - (Optional) Relative expiry time, in seconds, for expiring negative responses. This value is used only if the expiry time cannot be determined from any other source. It is applicable only to the following status codes: 307, 403, 404, and 410.
* `weakposrelexpiry` - (Optional) Relative expiry time, in seconds, for expiring positive responses with response codes between 200 and 399. Cannot be used in combination with other Expiry attributes. Similar to -relExpiry but has lower precedence.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cachecontentgroup. It has the same value as the `name` attribute.


## Import

A cachecontentgroup can be imported using its name, e.g.

```shell
terraform import citrixadc_cachecontentgroup.tf_cachecontentgroup my_cachecontentgroup
```
