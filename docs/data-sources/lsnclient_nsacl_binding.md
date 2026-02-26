---
subcategory: "LSN"
---

# Data Source: lsnclient_nsacl_binding

The lsnclient_nsacl_binding data source allows you to retrieve information about the binding between an LSN client and NSACLs.

## Example Usage

```terraform
data "citrixadc_lsnclient_nsacl_binding" "tf_lsnclient_nsacl_binding" {
  clientname = "my_lsn_client"
  aclname    = "my_acl"
}

output "clientname" {
  value = data.citrixadc_lsnclient_nsacl_binding.tf_lsnclient_nsacl_binding.clientname
}

output "aclname" {
  value = data.citrixadc_lsnclient_nsacl_binding.tf_lsnclient_nsacl_binding.aclname
}
```

## Argument Reference

* `clientname` - (Required) Name for the LSN client entity. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN client is created.
* `aclname` - (Required) Name(s) of any configured extended ACL(s) whose action is ALLOW. The condition specified in the extended ACL rule identifies the traffic from an LSN subscriber for which the Citrix ADC is to perform large scale NAT.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lsnclient_nsacl_binding. It is a system-generated identifier.
* `td` - ID of the traffic domain on which this subscriber or the subscriber network (as specified by the network parameter) belongs. If you do not specify an ID, the subscriber or the subscriber network becomes part of the default traffic domain.
