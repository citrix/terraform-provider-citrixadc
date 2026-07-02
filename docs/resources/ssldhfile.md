---
subcategory: "SSL"
---

# Resource: ssldhfile

The ssldhfile resource imports a Diffie-Hellman (DH) parameters file onto the Citrix ADC. DH parameter files supply the key-exchange parameters used to enable forward secrecy with DHE cipher suites on SSL virtual servers and services. The file is fetched from a remote source at creation time and stored under the given name on the appliance.

Because the ADC NITRO API exposes this object only through an import action (no in-place update), every attribute forces a new resource when changed.


## Example usage

```hcl
resource "citrixadc_ssldhfile" "tf_ssldhfile" {
  name = "dh2048"
  src  = "http://www.example.com/dh_file.pem"
}
```


## Argument Reference

* `name` - (Required) Name to assign to the imported DH file. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my file" or 'my file'). Changing this value forces a new resource to be created.
* `src` - (Required) URL specifying the protocol, host, and path, including file name, to the DH file to be imported. For example, `http://www.example.com/dh_file`. This is the import source consumed at creation time; the NITRO GET response does not faithfully echo it back, so the provider preserves the user-configured value in state. Changing this value forces a new resource to be created. Note: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access, and the issuer certificate of the HTTPS server is not present in the specific path on the Citrix ADC to authenticate the HTTPS server.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the ssldhfile. It has the same value as the `name` attribute.


## Import

An ssldhfile can be imported using its name, e.g.

```shell
terraform import citrixadc_ssldhfile.tf_ssldhfile dh2048
```
