---
subcategory: "SSL"
---

# Data Source: sslcacertgroup_sslcertkey_binding

The sslcacertgroup_sslcertkey_binding data source allows you to retrieve information about the binding between an SSL CA certificate group and an SSL certificate key.

## Example Usage

```terraform
data "citrixadc_sslcacertgroup_sslcertkey_binding" "sslcacertgroup_sslcertkey_binding_demo" {
  cacertgroupname = "ns_callout_certs1"
  certkeyname     = "tf_cacertkey"
}

output "ocspcheck" {
  value = data.citrixadc_sslcacertgroup_sslcertkey_binding.sslcacertgroup_sslcertkey_binding_demo.ocspcheck
}

output "crlcheck" {
  value = data.citrixadc_sslcacertgroup_sslcertkey_binding.sslcacertgroup_sslcertkey_binding_demo.crlcheck
}
```

## Argument Reference

* `cacertgroupname` - (Required) Name given to the CA certificate group. The name will be used to add the CA certificates to the group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my file" or 'my file').
* `certkeyname` - (Required) Name for the certkey added to the Citrix ADC. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the certificate-key pair is created.The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my cert" or 'my cert').

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `crlcheck` - The state of the CRL check parameter. (Mandatory/Optional)
* `ocspcheck` - The state of the OCSP check parameter. (Mandatory/Optional)
* `id` - The id of the sslcacertgroup_sslcertkey_binding. It is a system-generated identifier.
