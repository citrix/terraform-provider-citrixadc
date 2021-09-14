---
subcategory: "SSL"
---

# Resource: sslcacertgroup_sslcertkey_binding

The sslprofile_sslcipher_binding resource is used to create bindings between sslcacertgroup and sslcertkey.


## Example usage

```hcl
resource "citrixadc_sslcacertgroup_sslcertkey_binding" "sslcacertgroup_sslcertkey_binding_demo" {	
    cacertgroupname = citrixadc_sslcacertgroup.ns_callout_certs1.cacertgroupname
    certkeyname = citrixadc_sslcertkey.tf_cacertkey.certkey
    ocspcheck = "Mandatory"
}

resource "citrixadc_sslcertkey" "tf_sslcertkey" {
    certkey = "tf_sslcertkey"
    cert = "/var/tmp/certificate1.crt"
    key = "/var/tmp/key1.pem"
    notificationperiod = 40
    expirymonitor = "ENABLED"
}

resource "citrixadc_sslcertkey" "tf_cacertkey" {
    certkey = "tf_cacertkey"
    cert = "/var/tmp/ca.crt"
}
    
resource "citrixadc_sslcacertgroup" "ns_callout_certs1" {
    cacertgroupname = "ns_callout_certs1"
}
```


## Argument Reference

* `certkeyname` - (Required) Name for the certkey added to the Citrix ADC. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the certificate-key pair is created.The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my cert" or 'my cert').
* `crlcheck` - (Optional) The state of the CRL check parameter. (Mandatory/Optional). Possible values: [ Mandatory, Optional ]
* `ocspcheck` - (Optional) The state of the OCSP check parameter. (Mandatory/Optional). Possible values: [ Mandatory, Optional ]
* `cacertgroupname` - (Required) Name given to the CA certificate group. The name will be used to add the CA certificates to the group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my file" or 'my file').


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslcacertgroup_sslcertkey_binding. It has is the conatenation of the `cacertgroupname` and `certkeyname` attributes.


## Import

A sslcacertgroup_sslcertkey_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_sslcacertgroup_sslcertkey_binding.tf_sslcacertgroup_sslcertkey_binding tf_sslcacertgroup_sslcertkey_binding
```
