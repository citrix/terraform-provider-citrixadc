---
subcategory: "Responder"
---

# Resource: responderpolicylabel

The responderpolicylabel resource is used to create responder policy labels.


## Example usage

```hcl
resource "citrixadc_responderpolicylabel" "tf_responder_policylabel" {
	labelname = "tf_responder_policylabel"
	policylabeltype = "HTTP"
	comment = "Some comment"
}
```


## Argument Reference

* `labelname` - (Optional) Name for the responder policy label.
* `policylabeltype` - (Optional) Type of responses sent by the policies bound to this policy label. Types are: * HTTP - HTTP responses. * OTHERTCP - NON-HTTP TCP responses. * SIP_UDP - SIP responses. * RADIUS - RADIUS responses. * MYSQL - SQL responses in MySQL format. * MSSQL - SQL responses in Microsoft SQL format. * NAT - NAT response. Possible values: [ HTTP, OTHERTCP, SIP_UDP, SIP_TCP, MYSQL, MSSQL, NAT, DIAMETER, RADIUS, DNS ]
* `comment` - (Optional) Any comments to preserve information about this responder policy label.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the responderpolicylabel. It has the same value as the `labelname` attribute.


## Import

A responderpolicylabel can be imported using its labelname, e.g.

```shell
terraform import citrixadc_responderpolicylabel.tf_responder_policylabel tf_responder_policylabel
```
