---
subcategory: "SSL"
---

# Resource: sslservice_sslciphersuite_binding

The sslservice_sslciphersuite_binding resource is used to create binding between sslservice and sslciphersuite.


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

resource "citrixadc_sslcipher" "tfAccsslcipher" {
	ciphergroupname = "tfAccsslcipher"

	# ciphersuitebinding is MANDATORY attribute
	# Any change in the ciphersuitebinding will result in re-creation of the whole sslcipher resource.
	ciphersuitebinding {
		ciphername     = "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256"
		cipherpriority = 1
	}
	ciphersuitebinding {
		ciphername     = "TLS1.2-ECDHE-RSA-AES256-GCM-SHA384"
		cipherpriority = 2
	}
	ciphersuitebinding {
		ciphername     = "TLS1.2-ECDHE-RSA-AES-128-SHA256"
		cipherpriority = 3
	}
}
 resource "citrixadc_sslservice_sslciphersuite_binding" "tf_sslservice_sslcipher_binding" {
	ciphername = citrixadc_sslcipher.tfAccsslcipher.ciphergroupname
	servicename = citrixadc_service.tf_service.name   
}
```


## Argument Reference

* `ciphername` - (Required) The cipher group/alias/individual cipher configuration.
* `description` - (Optional) The cipher suite description.
* `servicename` - (Required) Name of the SSL service for which to set advanced configuration.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslservice_sslciphersuite_binding. t is the concatenation of the `servicename` and `ciphername` attributes separated by a comma.


## Import

A sslservice_sslciphersuite_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_sslservice_sslciphersuite_binding.tf_sslservice_sslciphersuite_binding tf_sslservice_sslciphersuite_binding
```
