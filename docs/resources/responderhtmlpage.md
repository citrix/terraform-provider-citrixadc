---
subcategory: "Responder"
---

# Resource: responderhtmlpage

The responderhtmlpage resource is used to create html pages to be used with the responder feature.


## Example usage

```hcl
resource "citrixadc_responderhtmlpage" "tf_responder_page" {
    name = "tf_responder_page"
    src = "local://tf_html_page.html"
    depends_on = [citrixadc_systemfile.tf_html_page]
}
```


## Argument Reference

* `cacertfile` - (Optional) CA certificate file name which will be used to verify the peer's certificate. The certificate should be imported using "import ssl certfile" CLI command or equivalent in API or GUI. If certificate name is not configured, then default root CA certificates are used for peer's certificate verification.
* `comment` - (Optional) Any comments to preserve information about the HTML page object.
* `name` - (Required) Name to assign to the HTML page object on the Citrix ADC.
* `overwrite` - (Optional) Overwrites the existing file
* `src` - (Optional) Local path to and name of, or URL (protocol, host, path, and file name) for, the file in which to store the imported HTML page. NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the responderhtmlpage. It has the same value as the `name` attribute.
