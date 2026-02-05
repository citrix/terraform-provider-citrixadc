---
subcategory: "Front End Optimization"
---

# Data Source: citrixadc_feoaction

This data source is used to retrieve information about an existing Front End Optimization (FEO) action.

## Example Usage

```hcl
data "citrixadc_feoaction" "example" {
  name = "my_feoaction"
}
```

## Argument Reference

* `name` - (Required) The name of the front end optimization action.

## Attribute Reference

In addition to the argument, the following attributes are exported:

* `id` - The ID of the FEO action (same as `name`).
* `cachemaxage` - Maxage for cache extension.
* `clientsidemeasurements` - Send AppFlow records about the web pages optimized by this action. The records provide FEO statistics, such as the number of HTTP requests that have been reduced for this page.
* `convertimporttolink` - Convert CSS import statements to HTML link tags.
* `csscombine` - Combine one or more CSS files into one file.
* `cssimginline` - Inline small images (less than 2KB) referred within CSS files as background-URLs.
* `cssinline` - Inline CSS files, whose size is less than 2KB, within the main page.
* `cssminify` - Remove comments and whitespaces from CSSs.
* `cssmovetohead` - Move any CSS file present within the body tag of an HTML page to the head tag.
* `dnsshards` - Set of domain names that replaces the parent domain.
* `domainsharding` - Domain name of the server.
* `htmlminify` - Remove comments and whitespaces from an HTML page.
* `imggiftopng` - Convert GIF image formats to PNG formats.
* `imginline` - Inline images whose size is less than 2KB.
* `imglazyload` - Download images, only when the user scrolls the page to view them.
* `imgshrinktoattrib` - Shrink image dimensions as per the height and width attributes specified in the &lt;img&gt; tag.
* `imgtojpegxr` - Convert JPEG, GIF, PNG image formats to JXR format.
* `imgtowebp` - Convert JPEG, GIF, PNG image formats to WEBP format.
* `jpgoptimize` - Remove non-image data such as comments from JPEG images.
* `jsinline` - Convert linked JavaScript files (less than 2KB) to inline JavaScript files.
* `jsminify` - Remove comments and whitespaces from JavaScript.
* `jsmovetoend` - Move any JavaScript present in the body tag to the end of the body tag.
* `pageextendcache` - Extend the time period during which the browser can use the cached resource.
