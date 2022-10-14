---
subcategory: "SSL"
---

# Resource: sslcertfile

The sslcertfile resource is used to manage ssl certfile.


## Example usage

```hcl
resource "citrixadc_sslcertfile" "tf_sslcertfile" {
  name = "tf_sslcertfile"
  src  = "local://certificate1.crt"
}
```


## Argument Reference

* `name` - (Required) Name to assign to the imported certificate file. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my file" or 'my file').
* `src` - (Required) URL specifying the protocol, host, and path, including file name, to the certificate file to be imported. For example, http://www.example.com/cert_file. NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslcertfile. It has the same value as the `name` attribute.


## Import

A sslcertfile can be imported using its name, e.g.

```shell
terraform import citrixadc_sslcertfile.tf_sslcertfile tf_sslcertfile
```
