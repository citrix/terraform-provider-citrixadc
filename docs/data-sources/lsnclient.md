---
subcategory: "LSN"
---

# Data Source: citrixadc_lsnclient

The lsnclient data source allows you to retrieve information about LSN (Large Scale NAT) client entities.

## Example usage

```terraform
data "citrixadc_lsnclient" "tf_lsnclient" {
  clientname = "my_lsnclient"
}

output "clientname" {
  value = data.citrixadc_lsnclient.tf_lsnclient.clientname
}

output "id" {
  value = data.citrixadc_lsnclient.tf_lsnclient.id
}
```

## Argument Reference

* `clientname` - (Required) Name for the LSN client entity. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN client is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "lsn client1" or 'lsn client1').

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lsnclient. It has the same value as the `clientname` attribute.

## Import

A lsnclient can be imported using its clientname, e.g.

```shell
terraform import citrixadc_lsnclient.tf_lsnclient my_lsnclient
```
