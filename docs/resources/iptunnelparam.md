---
subcategory: "Network"
---

# Resource: iptunnelparam

The iptunnelparam resource is used to create ip tunnel parameter resource.


## Example usage

```hcl
resource "citrixadc_iptunnelparam" "tf_iptunnelparam" {
  dropfrag             = "NO"
  dropfragcputhreshold = 1
  srciproundrobin      = "NO"
  enablestrictrx       = "NO"
  enablestricttx       = "NO"
  useclientsourceip    = "NO"
}
```


## Argument Reference

* `srcip` - (Optional) Common source-IP address for all tunnels. For a specific tunnel, this global setting is overridden if you have specified another source IP address. Must be a MIP or SNIP address. Minimum length =  1
* `dropfrag` - (Optional) Drop any IP packet that requires fragmentation before it is sent through the tunnel. Possible values: [ YES, NO ]
* `dropfragcputhreshold` - (Optional) Threshold value, as a percentage of CPU usage, at which to drop packets that require fragmentation to use the IP tunnel. Applies only if dropFragparameter is set to NO. The default value, 0, specifies that this parameter is not set. Minimum value =  1 Maximum value =  100
* `srciproundrobin` - (Optional) Use a different source IP address for each new session through a particular IP tunnel, as determined by round robin selection of one of the SNIP addresses. This setting is ignored if a common global source IP address has been specified for all the IP tunnels. This setting does not apply to a tunnel for which a source IP address has been specified. Possible values: [ YES, NO ]
* `enablestrictrx` - (Optional) Strict PBR check for IPSec packets received through tunnel. Possible values: [ YES, NO ]
* `enablestricttx` - (Optional) Strict PBR check for packets to be sent IPSec protected. Possible values: [ YES, NO ]
* `mac` - (Optional) The shared MAC used for shared IP between cluster nodes/HA peers.
* `useclientsourceip` - (Optional) Use client source IP as source IP for outer tunnel IP header. Possible values: [ YES, NO ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the iptunnelparam. It is a unique string prefixed with "tf-iptunnelparam-"

