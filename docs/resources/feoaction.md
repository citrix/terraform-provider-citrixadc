---
subcategory: "Front-end-optimization"
---

# Resource: feoaction

The feoaction resource is used to create feoaction.


## Example usage

```hcl
resource "citrixadc_feoaction" "tf_feoaction" {
  name              = "my_feoaction"
  cachemaxage       = 50
  imgshrinktoattrib = "true"
  imggiftopng       = "true"
}
```


## Argument Reference

* `name` - (Required) The name of the front end optimization action. Minimum length =  1
* `pageextendcache` - (Optional) Extend the time period during which the browser can use the cached resource.
* `cachemaxage` - (Optional) Maxage for cache extension. Minimum value =  0 Maximum value =  360
* `imgshrinktoattrib` - (Optional) Shrink image dimensions as per the height and width attributes specified in the <img> tag.
* `imggiftopng` - (Optional) Convert GIF image formats to PNG formats.
* `imgtowebp` - (Optional) Convert JPEG, GIF, PNG image formats to WEBP format.
* `imgtojpegxr` - (Optional) Convert JPEG, GIF, PNG image formats to JXR format.
* `imginline` - (Optional) Inline images whose size is less than 2KB.
* `cssimginline` - (Optional) Inline small images (less than 2KB) referred within CSS files as background-URLs.
* `jpgoptimize` - (Optional) Remove non-image data such as comments from JPEG images.
* `imglazyload` - (Optional) Download images, only when the user scrolls the page to view them.
* `cssminify` - (Optional) Remove comments and whitespaces from CSSs.
* `cssinline` - (Optional) Inline CSS files, whose size is less than 2KB, within the main page.
* `csscombine` - (Optional) Combine one or more CSS files into one file.
* `convertimporttolink` - (Optional) Convert CSS import statements to HTML link tags.
* `jsminify` - (Optional) Remove comments and whitespaces from JavaScript.
* `jsinline` - (Optional) Convert linked JavaScript files (less than 2KB) to inline JavaScript files.
* `htmlminify` - (Optional) Remove comments and whitespaces from an HTML page.
* `cssmovetohead` - (Optional) Move any CSS file present within the body tag of an HTML page to the head tag.
* `jsmovetoend` - (Optional) Move any JavaScript present in the body tag to the end of the body tag.
* `domainsharding` - (Optional) Domain name of the server.
* `dnsshards` - (Optional) Set of domain names that replaces the parent domain.
* `clientsidemeasurements` - (Optional) Send AppFlow records about the web pages optimized by this action. The records provide FEO statistics, such as the number of HTTP requests that have been reduced for this page. You must enable the Appflow feature before enabling this parameter.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the feoaction. It has the same value as the `name` attribute.


## Import

A feoaction can be imported using its name, e.g.

```shell
terraform import citrixadc_feoaction.tf_feoaction my_feoaction
```
