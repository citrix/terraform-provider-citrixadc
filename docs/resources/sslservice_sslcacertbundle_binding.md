---
subcategory: "SSL"
---

# Resource: sslservice_sslcacertbundle_binding

Binds a CA certificate bundle to an SSL service so the service can present the bundled CA certificate chain during client-certificate authentication in the TLS handshake. Use this when an SSL service must request and validate client certificates issued under a set of CAs distributed as a single bundle.


## Example usage

```hcl
resource "citrixadc_lbvserver" "tf_lbvserver" {
  name        = "tf_lbvserver"
  ipv46       = "10.10.10.44"
  port        = 443
  servicetype = "SSL"
  sslprofile  = "ns_default_ssl_profile_frontend"
}

resource "citrixadc_service" "tf_service" {
  name        = "tf_service"
  servicetype = "SSL"
  port        = 443
  ip          = "10.77.33.22"
  lbvserver   = citrixadc_lbvserver.tf_lbvserver.name
}

resource "citrixadc_sslcacertbundle" "tf_cacertbundle" {
  cacertbundlename = "tf_cacertbundle"
}

resource "citrixadc_sslservice_sslcacertbundle_binding" "tf_binding" {
  servicename      = citrixadc_service.tf_service.name
  cacertbundlename = citrixadc_sslcacertbundle.tf_cacertbundle.cacertbundlename
  skipcacertbundle = false
}
```


## Argument Reference

* `servicename` - (Required) Name of the SSL service for which to set advanced configuration. Changing this forces a new resource to be created.
* `cacertbundlename` - (Required) CA certbundle name bound to the service. Changing this forces a new resource to be created.
* `skipcacertbundle` - (Optional) The flag is used to indicate whether all CA_names in this particular CA certificate bundle needs to be sent to the SSL client while requesting for client certificate in a SSL handshake. Changing this forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslservice_sslcacertbundle_binding. It is the concatenation of the `cacertbundlename` and `servicename` unique attributes, formatted as comma-separated `key:value` pairs (for example, `cacertbundlename:tf_cacertbundle,servicename:tf_service`).


## Import

A sslservice_sslcacertbundle_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_sslservice_sslcacertbundle_binding.tf_binding cacertbundlename:tf_cacertbundle,servicename:tf_service
```
