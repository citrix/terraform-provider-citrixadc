---
subcategory: "Load Balancing"
---

# Data Source: lbpolicylabel

The lbpolicylabel data source allows you to retrieve information about an LB policy label.


## Example usage

```terraform
data "citrixadc_lbpolicylabel" "example" {
  labelname = "http_redirect_label"
}

output "lbpolicylabel_policylabeltype" {
  value = data.citrixadc_lbpolicylabel.example.policylabeltype
}
```


## Argument Reference

* `labelname` - (Required) Name of the LB policy label to look up.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbpolicylabel. It has the same value as the `labelname` attribute.
* `policylabeltype` - Protocols supported by the policy label. Possible values: [ HTTP, DNS, OTHERTCP, SIP_UDP, SIP_TCP, MYSQL, MSSQL, ORACLE, NAT, DIAMETER, RADIUS, MQTT, QUIC_BRIDGE, HTTP_QUIC ]
* `comment` - Any comments to preserve information about this LB policy label.
* `newname` - New name for the LB policy label (rename-only attribute; typically null on read).
