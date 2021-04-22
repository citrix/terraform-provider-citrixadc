---
subcategory: "SSL"
---

# Resource: sslvserver\_sslcertkey\_binding

The sslvserver\_sslcertkey\_binding resource is used to bind ssl certificates to SSL vservers.


## Example usage

```hcl
resource "citrixadc_sslcertkey" "tf_sslcertkey" {
  certkey            = "tf_sslcertkey"
  cert               = "/var/tmp/certificate2.crt"
  key                = "/var/tmp/key2.pem"
  notificationperiod = 40
  expirymonitor      = "ENABLED"
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
  ipv46       = "10.10.10.44"
  name        = "tf_lbvserver"
  port        = 443
  servicetype = "SSL"
  sslprofile  = "ns_default_ssl_profile_frontend"
}


resource "citrixadc_sslvserver_sslcertkey_binding" "tf_binding" {
  vservername = citrixadc_lbvserver.tf_lbvserver.name
  certkeyname = citrixadc_sslcertkey.tf_sslcertkey.certkey
  snicert     = true
}
```


## Argument Reference

* `certkeyname` - (Required) The name of the certificate key pair binding.
* `crlcheck` - (Optional) The state of the CRL check parameter. (Mandatory/Optional). Possible values: [ Mandatory, Optional ]
* `ocspcheck` - (Optional) The state of the OCSP check parameter. (Mandatory/Optional). Possible values: [ Mandatory, Optional ]
* `ca` - (Optional) CA certificate.
* `snicert` - (Optional) The name of the CertKey. Use this option to bind Certkey(s) which will be used in SNI processing.
* `skipcaname` - (Optional) The flag is used to indicate whether this particular CA certificate's CA_Name needs to be sent to the SSL client while requesting for client certificate in a SSL handshake.
* `vservername` - (Required) Name of the SSL virtual server.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslvserver\_sslcertkey\_binding. It is the concatenation of the `vservername` and `certkeyname` attributes separated by a comma.


## Import

A sslvserver\_sslcertkey\_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_sslvserver_sslcertkey_binding.tf_binding tf_lbvserver,tf_sslcertkey
```
