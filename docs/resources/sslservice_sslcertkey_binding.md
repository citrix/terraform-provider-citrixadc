---
subcategory: "SSL"
---

# Resource: sslservice_sslcertkey_binding

The sslservice_sslcertkey_binding resource is used to bind an SSL certificate key pair to an SSL service.


## Example usage

```hcl
resource "citrixadc_sslcertkey" "tf_certkey" {
  certkey            = "tf_certkey"
  cert               = "/nsconfig/ssl/ns-root.cert"
  key                = "/nsconfig/ssl/ns-root.key"
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

resource "citrixadc_service" "tf_service" {
  name        = "tf_service"
  servicetype = "SSL"
  port        = 443
  lbvserver   = citrixadc_lbvserver.tf_lbvserver.name
  ip          = "10.77.33.22"

  depends_on = [citrixadc_sslcertkey.tf_certkey]
}

resource "citrixadc_sslservice_sslcertkey_binding" "tf_binding" {
  servicename = citrixadc_service.tf_service.name
  certkeyname = citrixadc_sslcertkey.tf_certkey.certkey
  ca          = true
  ocspcheck   = "Optional"
}
```


## Argument Reference

* `servicename` - (Required) Name of the SSL service for which to set advanced configuration.
* `certkeyname` - (Required) The certificate key pair binding.
* `ca` - (Optional) CA certificate.
* `crlcheck` - (Optional) The state of the CRL check parameter. (Mandatory/Optional). Possible values: [ Mandatory, Optional ]
* `ocspcheck` - (Optional) Rule to use for the OCSP responder associated with the CA certificate during client authentication. If MANDATORY is specified, deny all SSL clients if the OCSP check fails because of connectivity issues with the remote OCSP server, or any other reason that prevents the OCSP check. With the OPTIONAL setting, allow SSL clients even if the OCSP check fails except when the client certificate is revoked. Possible values: [ Mandatory, Optional ]
* `snicert` - (Optional) The name of the CertKey. Use this option to bind Certkey(s) which will be used in SNI processing.
* `skipcaname` - (Optional) The flag is used to indicate whether this particular CA certificate's CA_Name needs to be sent to the SSL client while requesting for client certificate in a SSL handshake.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslservice_sslcertkey_binding. It is the concatenation of the `servicename`, `certkeyname`, `snicert` and `ca` attributes separated by a comma.


## Import

A sslservice_sslcertkey_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_sslservice_sslcertkey_binding.tf_binding tf_service,tf_certkey,true,true
```
