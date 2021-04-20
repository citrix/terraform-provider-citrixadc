---
subcategory: "Rewrite"
---

# Resource: rewritepolicylabel

The rewritepolicylabel resource is used to create rewrite policy labels.


## Example usage

```hcl
resource "citrixadc_rewritepolicylabel" "tf_rewritepolicylabel" {
	labelname = "tf_rewritepolicylabel"
    transform = "http_req"
	comment = "Some comment"
}
```


## Argument Reference

* `labelname` - (Optional) Name for the rewrite policy label.
* `transform` - (Optional) Types of transformations allowed by the policies bound to the label. For Rewrite, the following types are supported: * http_req - HTTP requests * http_res - HTTP responses * othertcp_req - Non-HTTP TCP requests * othertcp_res - Non-HTTP TCP responses * url - URLs * text - Text strings * clientless_vpn_req - Citrix ADC clientless VPN requests * clientless_vpn_res - Citrix ADC clientless VPN responses * sipudp_req - SIP requests * sipudp_res - SIP responses * diameter_req - DIAMETER requests * diameter_res - DIAMETER responses * radius_req - RADIUS requests * radius_res - RADIUS responses * dns_req - DNS requests * dns_res - DNS responses. Possible values: [ http_req, http_res, othertcp_req, othertcp_res, url, text, clientless_vpn_req, clientless_vpn_res, sipudp_req, sipudp_res, siptcp_req, siptcp_res, diameter_req, diameter_res, radius_req, radius_res, dns_req, dns_res ]
* `comment` - (Optional) Any comments to preserve information about this rewrite policy label.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the rewritepolicylabel. It has the same value as the `labelname` attribute.


## Import

A rewritepolicylabel can be imported using its labelname, e.g.

```shell
terraform import citrixadc_rewritepolicylabel.tf_rewritepolicylabel tf_rewritepolicylabel
```
