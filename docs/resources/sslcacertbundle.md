---
subcategory: "SSL"
---

# Resource: sslcacertbundle

A CA certificate bundle groups a list of trusted Certificate Authority (CA) certificates into a single named entity on the Citrix ADC. Bundling the issuing-CA chain together lets you bind one object during SSL client/server certificate verification instead of binding each CA certificate individually. Use this resource to register a CA bundle file that already exists on the appliance's disk.

The bundle is added from a file on the appliance and cannot be modified in place; changing either argument forces the resource to be re-created.


## Example usage

```hcl
resource "citrixadc_sslcacertbundle" "tf_cacertbundle" {
  cacertbundlename = "trusted-ca-bundle"
  bundlefile       = "/nsconfig/ssl/ca_bundle.pem"
}
```


## Argument Reference

* `cacertbundlename` - (Required) Name given to the CA certbundle. The name will be used for bind/unbind/update operations. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my file" or 'my file'). Changing this attribute forces a new resource to be created.
* `bundlefile` - (Required) Name of and, optionally, path to the X509 CA certificate bundle file that is used to form the cacertbundle entity. The CA certificate bundle file must already be present on the appliance's hard-disk drive or solid-state drive. `/nsconfig/ssl/` is the default path. The CA certificate bundle file consists of a list of certificates. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslcacertbundle. It has the same value as the `cacertbundlename` attribute.


## Import

A sslcacertbundle can be imported using its name, e.g.

```shell
terraform import citrixadc_sslcacertbundle.tf_cacertbundle trusted-ca-bundle
```
