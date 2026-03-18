---
subcategory: "LSN"
---

# Data Source: lsnclient_nsacl6_binding

The lsnclient_nsacl6_binding data source allows you to retrieve information about lsnclient_nsacl6_binding.

## Example Usage

```terraform
data "citrixadc_lsnclient_nsacl6_binding" "tf_lsnclient_nsacl6_binding" {
  clientname = "my_lsn_client"
  acl6name   = "my_acl6"
}

output "acl6name" {
  value = data.citrixadc_lsnclient_nsacl6_binding.tf_lsnclient_nsacl6_binding.acl6name
}

output "clientname" {
  value = data.citrixadc_lsnclient_nsacl6_binding.tf_lsnclient_nsacl6_binding.clientname
}
```

## Argument Reference

* `acl6name` - (Required) Name of any configured extended ACL6 whose action is ALLOW. The condition specified in the extended ACL6 rule is used as the condition for the LSN client.
* `clientname` - (Required) Name for the LSN client entity. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN client is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "lsn client1" or 'lsn client1').

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lsnclient_nsacl6_binding. It is a system-generated identifier.
* `td` - ID of the traffic domain on which this subscriber or the subscriber network (as specified by the network parameter) belongs. If you do not specify an ID, the subscriber or the subscriber network becomes part of the default traffic domain.
