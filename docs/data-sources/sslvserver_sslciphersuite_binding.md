---
subcategory: "SSL"
---

# Data Source: sslvserver_sslciphersuite_binding

The sslvserver_sslciphersuite_binding data source allows you to retrieve information about the binding between an SSL virtual server and an SSL cipher suite.

## Example Usage

```terraform
data "citrixadc_sslvserver_sslciphersuite_binding" "tf_sslvserver_sslciphersuite_binding" {
  vservername = "tf_sslvserver_ds"
  ciphername  = "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256"
}

output "description" {
  value = data.citrixadc_sslvserver_sslciphersuite_binding.tf_sslvserver_sslciphersuite_binding.description
}
```

## Argument Reference

* `vservername` - (Required) Name of the SSL virtual server.
* `ciphername` - (Required) The cipher group/alias/individual cipher configuration.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `description` - The cipher suite description.
* `id` - The id of the sslvserver_sslciphersuite_binding. It is a system-generated identifier.
