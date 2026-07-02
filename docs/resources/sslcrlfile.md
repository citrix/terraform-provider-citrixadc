---
subcategory: "SSL"
---

# Resource: sslcrlfile

The sslcrlfile resource imports a Certificate Revocation List (CRL) file onto the Citrix ADC so it can be used by SSL CRL configurations to determine which certificates have been revoked. The CRL is fetched from a remote source at creation time and stored under the given name on the appliance.

Because the ADC NITRO API exposes this object only through an import action (no in-place update), every attribute forces a new resource when changed.


## Example usage

```hcl
resource "citrixadc_sslcrlfile" "tf_sslcrlfile" {
  name = "crl1"
  src  = "http://www.example.com/crl_file.crl"
}
```


## Argument Reference

* `name` - (Required) Name to assign to the imported CRL file. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my file" or 'my file'). Changing this value forces a new resource to be created.
* `src` - (Required) URL specifying the protocol, host, and path, including file name to the CRL file to be imported. For example, `http://www.example.com/crl_file`. This is the import source consumed at creation time; the NITRO GET response does not faithfully echo it back, so the provider preserves the user-configured value in state. Changing this value forces a new resource to be created. Note: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access, and the issuer certificate of the HTTPS server is not present in the specific path on the Citrix ADC to authenticate the HTTPS server.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslcrlfile. It has the same value as the `name` attribute.


## Import

An sslcrlfile can be imported using its name, e.g.

```shell
terraform import citrixadc_sslcrlfile.tf_sslcrlfile crl1
```
