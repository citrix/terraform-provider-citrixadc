---
subcategory: "Compression"
---

# Resource: cmppolicylabel

The cmppolicylabel resource is used to create cmppolicylabel.


## Example usage

```hcl
resource "citrixadc_cmppolicylabel" "tf_cmppolicylabel" {
  labelname = "my_cmppolicy_label"
  type      = "REQ"
}

```


## Argument Reference

* `labelname` - (Required) Name of the HTTP compression policy label. Must begin with a letter, number, or the underscore character (_). Additional characters allowed, after the first character, are the hyphen (-), period (.) pound sign (#), space ( ), at sign (@), equals (=), and colon (:). The name must be unique within the list of policy labels for compression policies. Can be renamed after the policy label is created.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my cmp policylabel" or 'my cmp policylabel'). Minimum length =  1
* `type` - (Required) Type of packets (request packets or response) against which to match the policies bound to this policy label. Possible values: [ REQ, RES, HTTPQUIC_REQ, HTTPQUIC_RES ]

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cmppolicylabel. It has the same value as the `labelname` attribute.


## Import

A cmppolicylabel can be imported using its name, e.g.

```shell
terraform import citrixadc_cmppolicylabel.tf_cmppolicylabel my_cmppolicy_label
```
