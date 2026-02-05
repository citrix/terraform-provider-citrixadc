---
subcategory: "Rewrite"
---

# Data Source: citrixadc_rewritepolicylabel

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
