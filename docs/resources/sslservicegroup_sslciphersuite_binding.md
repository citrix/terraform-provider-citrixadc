---
subcategory: "SSL"
---

# Resource: sslservicegroup_sslciphersuite_binding

The sslservicegroup_sslciphersuite_binding resource is used to add a sslciphersuite to sslservicegroup.


## Example usage

```hcl
resource "citrixadc_sslservicegroup_sslciphersuite_binding" "tf_sslservicegroup_sslciphersuite_binding" {
	ciphername = citrixadc_sslcipher.tf_sslcipher.ciphergroupname
	servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
}

resource "citrixadc_servicegroup" "tf_servicegroup" {
	servicegroupname = "tf_servicegroup"
	servicetype = "SSL"
}

resource "citrixadc_sslcipher" "tf_sslcipher" {
	ciphergroupname = "tf_sslcipher"
	ciphersuitebinding {
		ciphername     = "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256"
		cipherpriority = 1
	}
}
```


## Argument Reference

* `ciphername` - (Required) The name of the cipher group/alias/name configured for the SSL service group.
* `description` - (Optional) The description of the cipher.
* `servicegroupname` - (Required) The name of the SSL service to which the SSL policy needs to be bound.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslservicegroup_sslciphersuite_binding. It is the concatenation of the `servicegroupname` and `ciphername` attributes separated by a comma.


## Import

A sslservicegroup_sslciphersuite_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_sslservicegroup_sslciphersuite_binding.tf_sslservicegroup_sslciphersuite_binding tf_servicegroup,tf_sslcipher
```
