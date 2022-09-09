---
subcategory: "Appqoe"
---

# Resource: appqoecustomresp

The appqoecustomresp resource is used to create appqoecustomresp.


## Example usage

```hcl
resource "citrixadc_appqoecustomresp" "tf_appqoecustomresp" {
  name   = "my_appqoecustomresp"
  src   = "local://index.html"
}
```


## Argument Reference

* `name` - (Required) Indicates name of the custom response HTML page to import/update.
* `src` - (Optional) URL \(protocol, host, path, and file name\) from where the location file will be imported.
NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appqoecustomresp. It has the same value as the `name` attribute.


## Import

A appqoecustomresp can be imported using its name, e.g.

```shell
terraform import citrixadc_appqoecustomresp.tf_appqoecustomresp my_appqoecustomresp
```
