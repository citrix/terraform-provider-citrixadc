---
subcategory: "SSL"
---

# Data Source `sslcipher_sslvserver_bindings`

The sslcipher_sslvserver_bindings data source allows you to retrieve information about the bound sslvserver to a sslcipher

## Example Usage

```terraform
data "citrixadc_sslcipher_sslvserver_bindings" "sslbindings" {
    ciphername = "tfsslcipher"
}

```

## Argument Reference

* `ciphername` - (Required) Name of the cipher.


## Attributes Reference

* `bound_sslvservers` - The list of the bound sslvservers names.
