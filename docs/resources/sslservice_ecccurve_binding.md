---
subcategory: "SSL"
---

# Resource: sslservice_ecccurve_binding

The sslservice_ecccurve_binding resource is used to bind sslservice and ecccurve.


## Example usage

```hcl
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
}

resource "citrixadc_sslservice_ecccurve_binding" "tf_sslservice_ecccurve_binding" {
	ecccurvename = "P_256"
	servicename = citrixadc_service.tf_service.name
	
}
```


## Argument Reference

* `ecccurvename` - (Required) Named ECC curve bound to service/vserver. Possible values: [ ALL, P_224, P_256, P_384, P_521 ]
* `servicename` - (Required) Name of the SSL service for which to set advanced configuration.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslservice_ecccurve_binding. It is the concatenation of the `servicename` and `ecccurvename` attributes separated by a comma.


## Import

A sslservice_ecccurve_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_sslservice_ecccurve_binding.tf_sslservice_ecccurve_binding tf_sslservice_ecccurve_binding
```
