---
subcategory: "SSL"
---

# Resource: sslcertbundle

A certificate bundle imports a set of certificates packaged together (for example, an end-entity certificate together with its intermediate chain) onto the Citrix ADC from a remote or local source. Use this resource to pull a bundle file into the appliance's SSL store in a single operation so it can be referenced by SSL configuration.

The bundle is brought in through the NITRO `Import` action. There is no in-place update; changing either argument forces the resource to be re-created.


## Example usage

```hcl
resource "citrixadc_sslcertbundle" "tf_certbundle" {
  name = "web-cert-bundle"
  src  = "http://www.example.com/cert_bundle_file"
}
```

Importing from a file already on the appliance:

```hcl
resource "citrixadc_sslcertbundle" "tf_certbundle" {
  name = "web-cert-bundle"
  src  = "local:cert_bundle_file"
}
```


## Argument Reference

* `name` - (Required) Name to assign to the imported certificate bundle. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my file" or 'my file'). Changing this attribute forces a new resource to be created.
* `src` - (Required) URL specifying the protocol, host, and path, including file name, to the certificate bundle to be imported. For example, `http://www.example.com/cert_bundle_file`. Note: This value is the import source only. It is not round-tripped by the NITRO GET, so the provider preserves the user-configured value in Terraform state. The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access and the issuer certificate of that HTTPS server is not present in the expected path on the Citrix ADC. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslcertbundle. It has the same value as the `name` attribute.


## Import

A sslcertbundle can be imported using its name, e.g.

```shell
terraform import citrixadc_sslcertbundle.tf_certbundle web-cert-bundle
```
