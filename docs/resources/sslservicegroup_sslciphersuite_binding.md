---
subcategory: "SSL"
---

# Resource: sslservicegroup_sslciphersuite_binding

The sslservicegroup_sslciphersuite_binding resource is used to create sslservicegroup_sslciphersuite_binding.


## Example usage

```hcl
resource "citrixadc_sslservicegroup_sslciphersuite_binding" "tf_sslservicegroup_sslciphersuite_binding" {
  ciphername       = citrixadc_sslcipher.tf_sslcipher.ciphergroupname
  servicegroupname = citrixadc_gslbservicegroup.tf_gslbservicegroup.servicegroupname
}
resource "citrixadc_sslcipher" "tf_sslcipher" {
    ciphergroupname = "my_ciphersuite"
   
    ciphersuitebinding {
        ciphername     = "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256"
        cipherpriority = 1
    }    
}

resource "citrixadc_gslbservicegroup" "tf_gslbservicegroup" {
  servicegroupname = "my_gslbvservicegroup"
  servicetype      = "SSL_TCP"
  cip              = "DISABLED"
  healthmonitor    = "NO"
  sitename         = citrixadc_gslbsite.site_local.sitename
}
resource "citrixadc_gslbsite" "site_local" {
  sitename        = "Site-Local"
  siteipaddress   = "172.31.96.234"
  sessionexchange = "DISABLED"
}
```


## Argument Reference

* `ciphername` - (Required) The name of the cipher group/alias/name configured for the SSL service group.
* `description` - (Optional) The description of the cipher.
* `servicegroupname` - (Required) The name of the SSL service to which the SSL policy needs to be bound.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslservicegroup_sslciphersuite_binding. It is the concatenation of `servicegroupname` and `ciphername` attributes separated by a comma.


## Import

A sslservicegroup_sslciphersuite_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_sslservicegroup_sslciphersuite_binding.tf_sslservicegroup_sslciphersuite_binding  my_gslbvservicegroup,my_ciphersuite
```
