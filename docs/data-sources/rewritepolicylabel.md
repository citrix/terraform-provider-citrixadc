---
subcategory: "Rewrite"
---

# Data Source: rewritepolicylabel

This data source retrieves information about a specific rewrite policy label.

## Example Usage

```hcl
data "citrixadc_rewritepolicylabel" "example" {
  labelname = "my_rewrite_label"
}

output "label_transform" {
  value = data.citrixadc_rewritepolicylabel.example.transform
}
```

## Argument Reference

* `labelname` - (Required) Name of the rewrite policy label.

## Attribute Reference

In addition to the argument, the following attributes are exported:

* `id` - The ID of the rewrite policy label.
* `transform` - Types of transformations allowed by the policies bound to the label. Possible values include: http_req, http_res, othertcp_req, othertcp_res, url, text, clientless_vpn_req, clientless_vpn_res, sipudp_req, sipudp_res, diameter_req, diameter_res, radius_req, radius_res, dns_req, dns_res, mqtt_req, mqtt_res.
* `comment` - Any comments to preserve information about this rewrite policy label.
* `newname` - New name for the rewrite policy label.

### Read-only rewritepolicylabel metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_rewritepolicylabel` resource). They are GET-only / Computed. Any attribute the appliance does not return is `null`.

* `numpol` - Number of polices bound to label.
* `hits` - Number of times policy label was invoked.
* `priority` - Specifies the priority of the policy.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `labeltype` - Type of invocation. Possible values = reqvserver, resvserver, resHttpEventvserver, policylabel.
* `invoke_labelname` - If labelType is policylabel, name of the policy label to invoke. If labelType is reqvserver or resvserver, name of the virtual server to which to forward the request or response.
* `flowtype` - Flowtype of the bound rewrite policy.
* `description` - Description of the policylabel.
* `isdefault` - A value of true is returned if it is a default rewritepolicylabel.
* `builtin` - Flag to determine if rewrite policy label is built-in or not. A list of strings. Possible values = MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL.
* `feature` - The feature to be checked while applying this config.
