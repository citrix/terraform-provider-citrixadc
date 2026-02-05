---
subcategory: "SSL"
---

# Data Source `sslcertfile`

The sslcertfile data source allows you to retrieve information about SSL certificate files.


## Example usage

```terraform
data "citrixadc_sslcertfile" "tf_sslcertfile" {
  name = "certificate1"
}

output "src" {
  value = data.citrixadc_sslcertfile.tf_sslcertfile.src
}
```


## Argument Reference

* `name` - (Optional) Name to assign to the imported certificate file. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my file" or 'my file').

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `src` - URL specifying the protocol, host, and path, including file name, to the certificate file to be imported. For example, http://www.example.com/cert_file.
NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access, and the issuer certificate of the HTTPS server is not present in the specific path on NetScaler to authenticate the HTTPS server.

## Attribute Reference

* `id` - The id of the sslcertfile. It has the same value as the `name` attribute.


## Import

A sslcertfile can be imported using its name, e.g.

```shell
terraform import citrixadc_sslcertfile.tf_sslcertfile certificate1
```
