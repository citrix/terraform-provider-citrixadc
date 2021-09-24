---
subcategory: "SSL"
---

# Resource: sslservice_sslcertkey_binding

The sslservice_sslcertkey_binding resource is used to create binding between sslservice and sslcertkey.


## Example usage

```hcl
resource "citrixadc_sslcertkey" "tf_certkey" {
	certkey = "tf_certkey"
	cert = "/nsconfig/ssl/ns-root.cert"
	key = "/nsconfig/ssl/ns-root.key"
	notificationperiod = 40
	expirymonitor = "ENABLED"
}

resource "citrixadc_sslservice" "demo_sslservice" {
	cipherredirect = "DISABLED"
	clientauth = "DISABLED"
	dh = "DISABLED"
	dhcount = 0
	dhkeyexpsizelimit = "DISABLED"
	dtls12 = "DISABLED"
	ersa = "DISABLED"
	redirectportrewrite = "DISABLED"
	serverauth = "ENABLED"
	servicename = citrixadc_service.tf_service.name
	sessreuse = "ENABLED"
	sesstimeout = 300
	snienable = "DISABLED"
	ssl2 = "DISABLED"
	ssl3 = "ENABLED"
	sslredirect = "DISABLED"
	sslv2redirect = "DISABLED"
	tls1 = "ENABLED"
	tls11 = "ENABLED"
	tls12 = "ENABLED"
	tls13 = "DISABLED"
	
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
	ipv46       = "10.10.10.44"
	name        = "tf_lbvserver"
	port        = 443
	servicetype = "SSL"
	sslprofile  = "ns_default_ssl_profile_frontend"
}

resource "citrixadc_service" "tf_service" {
	name = "tf_service"
	servicetype = "SSL"
	port = 443 
	lbvserver = citrixadc_lbvserver.tf_lbvserver.name
	ip = "10.77.33.22"

	depends_on = [citrixadc_sslcertkey.tf_certkey]

}
resource "citrixadc_sslservice_sslcertkey_binding" "tf_sslservice_sslcertkey_binding" {
	certkeyname = citrixadc_sslcertkey.tf_certkey.certkey
	servicename = citrixadc_service.tf_service.name
	ca = true
	ocspcheck = "Optional"
}
```


## Argument Reference

* `certkeyname` - (Required) The certificate key pair binding.
* `crlcheck` - (Optional) The state of the CRL check parameter. (Mandatory/Optional). Possible values: [ Mandatory, Optional ]
* `ocspcheck` - (Optional) Rule to use for the OCSP responder associated with the CA certificate during client authentication. If MANDATORY is specified, deny all SSL clients if the OCSP check fails because of connectivity issues with the remote OCSP server, or any other reason that prevents the OCSP check. With the OPTIONAL setting, allow SSL clients even if the OCSP check fails except when the client certificate is revoked. Possible values: [ Mandatory, Optional ]
* `ca` - (Optional) CA certificate.
* `snicert` - (Optional) The name of the CertKey. Use this option to bind Certkey(s) which will be used in SNI processing.
* `skipcaname` - (Optional) The flag is used to indicate whether this particular CA certificate's CA_Name needs to be sent to the SSL client while requesting      for client certificate in a SSL handshake.
* `servicename` - (Required) Name of the SSL service for which to set advanced configuration.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslservice_sslcertkey_binding. It is the concatenation of the `servicename` and `certkeyname` attributes separated by a comma.



## Import

A sslservice_sslcertkey_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_sslservice_sslcertkey_binding.tf_sslservice_sslcertkey_binding tf_sslservice_sslcertkey_binding
```
