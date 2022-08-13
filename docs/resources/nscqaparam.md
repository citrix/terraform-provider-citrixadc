---
subcategory: "NS"
---

# Resource: nscqaparam

The nscqaparam resource is used to update nscqaparam.


## Example usage

```hcl
resource "citrixadc_nscqaparam" "tf_nscqaparam" {
    harqretxdelay = 5
    net1label = "2g"
    minrttnet1 = 25
    lr1coeflist = "intercept=4.95,thruputavg=5.92,iaiavg=-189.48,rttmin=-15.75,loaddelayavg=0.01,noisedelayavg=-2.59"
    lr1probthresh = 0.2
    net1cclscale = "25,50,75"
    net1csqscale = "25,50,75"
    net1logcoef = " 1.49,3.62,-0.14,1.84,4.83"
    net2label = "3g"
    net3label = "4g"
}
```


## Argument Reference

* `harqretxdelay` - (Optional) HARQ retransmission delay (in ms). Minimum value =  1 Maximum value =  64000
* `net1label` - (Optional) Name of the network label. Maximum length =  15
* `minrttnet1` - (Optional) MIN RTT (in ms) for the first network. Minimum value =  0 Maximum value =  6400
* `lr1coeflist` - (Optional) coefficients values for Label1.
* `lr1probthresh` - (Optional) Probability threshold values for LR model to differentiate between NET1 and reset(NET2 and NET3). Minimum value =  0 Maximum value =  1
* `net1cclscale` - (Optional) Three congestion level scores limits corresponding to None, Low, Medium.
* `net1csqscale` - (Optional) Three signal quality level scores limits corresponding to Excellent, Good, Fair.
* `net1logcoef` - (Optional) Connection quality ranking Log coefficients of network 1.
* `net2label` - (Optional) Name of the network label 2. Maximum length =  15
* `minrttnet2` - (Optional) MIN RTT (in ms) for the second network. Minimum value =  0 Maximum value =  6400
* `lr2coeflist` - (Optional) coefficients values for Label 2.
* `lr2probthresh` - (Optional) Probability threshold values for LR model to differentiate between NET2 and NET3. Minimum value =  0 Maximum value =  1
* `net2cclscale` - (Optional) Three congestion level scores limits corresponding to None, Low, Medium.
* `net2csqscale` - (Optional) Three signal quality level scores limits corresponding to Excellent, Good, Fair.
* `net2logcoef` - (Optional) Connnection quality ranking Log coefficients of network 2.
* `net3label` - (Optional) Name of the network label 3. Maximum length =  15
* `minrttnet3` - (Optional) MIN RTT (in ms) for the third network. Minimum value =  0 Maximum value =  6400
* `net3cclscale` - (Optional) Three congestion level scores limits corresponding to None, Low, Medium.
* `net3csqscale` - (Optional) Three signal quality level scores limits corresponding to Excellent, Good, Fair.
* `net3logcoef` - (Optional) Connection quality ranking Log coefficients of network 3.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nscqaparam. It is a unique string prefixed with `tf-nscqaparam-`.

