---
subcategory: "Load Balancing"
---

# Resource: lbpolicylabel

This resource is used to manage LB policy labels.


## Example usage

```hcl
resource "citrixadc_lbpolicylabel" "tf_lbpolicylabel" {
  labelname       = "http_redirect_label"
  policylabeltype = "HTTP"
  comment         = "Reusable HTTP policy bank for redirect rules"
}
```


## Argument Reference

* `labelname` - (Required) Name for the LB policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, `"my lb policy label"` or `'my lb policy label'`). Changing this value forces a new resource to be created.
* `policylabeltype` - (Optional) Protocols supported by the policy label. Changing this value forces a new resource to be created. Defaults to `"HTTP"`. Possible values: [ HTTP, DNS, OTHERTCP, SIP_UDP, SIP_TCP, MYSQL, MSSQL, ORACLE, NAT, DIAMETER, RADIUS, MQTT, QUIC_BRIDGE, HTTP_QUIC ]
* `comment` - (Optional) Any comments to preserve information about this LB policy label. Changing this value forces a new resource to be created.
* `newname` - (Optional) New name for the LB policy label. This is a rename-only, advanced attribute. Most users should leave it unset and instead control the name through `labelname`. Must follow the same naming rules as `labelname`.

~> **Note:** This resource is immutable. Changing `labelname`, `policylabeltype`, or `comment` forces Terraform to destroy and recreate the resource.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbpolicylabel. It has the same value as the `labelname` attribute.


## Import

An lbpolicylabel can be imported using its name (the `labelname`), e.g.

```shell
terraform import citrixadc_lbpolicylabel.tf_lbpolicylabel http_redirect_label
```
