---
subcategory: "NS"
---

# Data Source `nscqaparam`

The nscqaparam data source allows you to retrieve information about NS Connection Quality Assessment (CQA) parameters configuration.


## Example usage

```terraform
data "citrixadc_nscqaparam" "tf_nscqaparam" {
}

output "harqretxdelay" {
  value = data.citrixadc_nscqaparam.tf_nscqaparam.harqretxdelay
}

output "net1label" {
  value = data.citrixadc_nscqaparam.tf_nscqaparam.net1label
}

output "net2label" {
  value = data.citrixadc_nscqaparam.tf_nscqaparam.net2label
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `harqretxdelay` - HARQ retransmission delay (in ms).
* `net1label` - Name of the network label.
* `minrttnet1` - MIN RTT (in ms) for the first network.
* `lr1coeflist` - Coefficients values for Label1.
* `lr1probthresh` - Probability threshold values for LR model to differentiate between NET1 and reset(NET2 and NET3).
* `net1cclscale` - Three congestion level scores limits corresponding to None, Low, Medium.
* `net1csqscale` - Three signal quality level scores limits corresponding to Excellent, Good, Fair.
* `net1logcoef` - Connection quality ranking Log coefficients of network 1.
* `net2label` - Name of the network label 2.
* `minrttnet2` - MIN RTT (in ms) for the second network.
* `lr2coeflist` - Coefficients values for Label 2.
* `lr2probthresh` - Probability threshold values for LR model to differentiate between NET2 and NET3.
* `net2cclscale` - Three congestion level scores limits corresponding to None, Low, Medium.
* `net2csqscale` - Three signal quality level scores limits corresponding to Excellent, Good, Fair.
* `net2logcoef` - Connection quality ranking Log coefficients of network 2.
* `net3label` - Name of the network label 3.
* `minrttnet3` - MIN RTT (in ms) for the third network.
* `net3cclscale` - Three congestion level scores limits corresponding to None, Low, Medium.
* `net3csqscale` - Three signal quality level scores limits corresponding to Excellent, Good, Fair.
* `net3logcoef` - Connection quality ranking Log coefficients of network 3.

## Attribute Reference

* `id` - The id of the nscqaparam. It is a system-generated identifier.
