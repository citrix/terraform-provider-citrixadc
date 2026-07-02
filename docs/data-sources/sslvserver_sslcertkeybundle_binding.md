---
subcategory: "SSL"
---

# Data Source: sslvserver_sslcertkeybundle_binding

The sslvserver_sslcertkeybundle_binding data source allows you to retrieve information about a certificate-key bundle binding to an SSL virtual server.

## Example usage

```terraform
data "citrixadc_sslvserver_sslcertkeybundle_binding" "tf_binding" {
  vservername       = "tf_lbvserver"
  certkeybundlename = "tf_certkeybundle"
  snicertkeybundle  = false
}

output "vservername" {
  value = data.citrixadc_sslvserver_sslcertkeybundle_binding.tf_binding.vservername
}

output "certkeybundlename" {
  value = data.citrixadc_sslvserver_sslcertkeybundle_binding.tf_binding.certkeybundlename
}
```

## Argument Reference

The following arguments are required:

* `vservername` - (Required) Name of the SSL virtual server.
* `certkeybundlename` - (Required) Certkeybundle name bound to the vserver.
* `snicertkeybundle` - (Required) Use this option to bind certkeybundle which will be used in SNI processing.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslvserver_sslcertkeybundle_binding. It is the concatenation of the `vservername` and `certkeybundlename` attributes separated by a comma.
