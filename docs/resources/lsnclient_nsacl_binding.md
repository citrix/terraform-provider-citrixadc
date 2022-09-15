---
subcategory: "Lsn"
---

# Resource: lsnclient_nsacl_binding

The lsnclient_nsacl_binding resource is used to create lsnclient_nsacl_binding.


## Example usage

```hcl
resource "citrixadc_lsnclient_nsacl_binding" "tf_lsnclient_nsacl_binding" {
  clientname = "my_lsnclient"
  aclname    = "my_acl"
}

```


## Argument Reference

* `aclname` - (Required) Name(s) of any configured extended ACL(s) whose action is ALLOW. The condition specified in the extended ACL rule identifies the traffic from an LSN subscriber for which the Citrix ADC is to perform large scale NAT.
* `clientname` - (Required) Name for the LSN client entity. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN client is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "lsn client1" or 'lsn client1').
* `td` - (Optional) ID of the traffic domain on which this subscriber or the subscriber network (as specified by the network parameter) belongs.  If you do not specify an ID, the subscriber or the subscriber network becomes part of the default traffic domain.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lsnclient_nsacl_binding. It is the concatenation of `clientname` and `aclname` attributes separated by a comma.


## Import

A lsnclient_nsacl_binding> can be imported using its name, e.g.

```shell
terraform import citrixadc_lsnclient_nsacl_binding.tf_lsnclient_nsacl_binding my_lsnclient,my_acl
```
