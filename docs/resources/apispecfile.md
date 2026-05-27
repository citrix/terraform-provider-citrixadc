---
subcategory: "API Definition"
---

# Resource: apispecfile

The apispecfile resource is used to import an API specification file from a remote URL onto the Citrix ADC. The imported spec file can subsequently be referenced by other API-Definition resources (e.g. `citrixadc_apispec`) for schema validation, routing, and protection.

The underlying NITRO operation is an `Import` action: the appliance downloads the spec from the supplied URL and stores it locally under the given name. NITRO exposes no update endpoint for this object, so every attribute forces resource replacement when changed.


## Example usage

```hcl
resource "citrixadc_apispecfile" "tf_apispecfile" {
  name      = "my_apispecfile"
  src       = "http://www.example.com/petstore.yaml"
  overwrite = true
}
```


## Argument Reference

* `name` - (Required) Name to assign to the imported spec file. Must begin with an ASCII alphanumeric or underscore (`_`) character, and must contain only ASCII alphanumeric, underscore, hash (`#`), period (`.`), space, colon (`:`), at (`@`), equals (`=`), and hyphen (`-`) characters. If the name includes one or more spaces, enclose the name in double or single quotation marks when using the Citrix ADC CLI (for example, `"my file"` or `'my file'`). Changing this attribute forces resource replacement.
* `src` - (Required) URL specifying the protocol, host, and path, including file name, to the spec file to be imported. For example, `http://www.example.com/spec_file`. The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access. Changing this attribute forces resource replacement.
* `overwrite` - (Optional) Overwrite any existing schema file of the same name. This value is consumed by the NITRO `Import` action at create time and is not echoed back by the appliance on subsequent reads. Changing this attribute forces resource replacement.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the apispecfile. It has the same value as the `name` attribute.


## Import

An apispecfile can be imported using its name, e.g.

```shell
terraform import citrixadc_apispecfile.tf_apispecfile my_apispecfile
```
