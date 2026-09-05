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

### Read-only lbpolicylabel metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_lbpolicylabel` resource). They are GET-only/Computed and are `null` when the appliance omits them.

* `numpol` - Number of policies bound to the label.
* `hits` - Number of times the policy label was invoked.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `labeltype` - Type of policy label to invoke. Possible values: `reqvserver`, `policylabel`.
* `invoke_labelname` - If labelType is policylabel, name of the policy label to invoke; if labelType is reqvserver, name of the virtual server.
