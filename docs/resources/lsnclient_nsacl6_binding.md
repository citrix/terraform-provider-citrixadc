---
subcategory: "Lsn"
---

# Resource: lsnclient_nsacl6_binding

The lsnclient_nsacl6_binding resource is used to create lsnclient_nsacl6_binding.


## Example usage

```hcl
resource "citrixadc_lsnclient_nsacl6_binding" "tf_lsnclient_nsacl6_binding" {
  clientname = "my_lsn_client"
  acl6name   = "my_acl6"
}
```


## Argument Reference

* `acl6name` - (Required) Name of any configured extended ACL6 whose action is ALLOW. The condition specified in the extended ACL6 rule is used as the condition for the LSN client.
* `clientname` - (Required) Name for the LSN client entity. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN client is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "lsn client1" or 'lsn client1').
* `td` - (Optional) ID of the traffic domain on which this subscriber or the subscriber network (as specified by the network parameter) belongs.  If you do not specify an ID, the subscriber or the subscriber network becomes part of the default traffic domain.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lsnclient_nsacl6_binding. It is the concatenation of `clientname` and `acl6name` attributes separated by a comma.


## Import

A lsnclient_nsacl6_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_lsnclient_nsacl6_binding.tf_lsnclient_nsacl6_binding my_lsn_client,my_acl6
```
