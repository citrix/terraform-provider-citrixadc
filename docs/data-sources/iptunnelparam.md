---
subcategory: "Network"
---

# Data Source `iptunnelparam`

The iptunnelparam data source allows you to retrieve information about IP tunnel parameters configuration.


## Example usage

```terraform
data "citrixadc_iptunnelparam" "tf_iptunnelparam" {
}

output "dropfrag" {
  value = data.citrixadc_iptunnelparam.tf_iptunnelparam.dropfrag
}

output "enablestrictrx" {
  value = data.citrixadc_iptunnelparam.tf_iptunnelparam.enablestrictrx
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `dropfrag` - Drop any IP packet that requires fragmentation before it is sent through the tunnel. Possible values: `YES`, `NO`.
* `dropfragcputhreshold` - Threshold value, as a percentage of CPU usage, at which to drop packets that require fragmentation to use the IP tunnel. Applies only if dropFragparameter is set to NO. The default value, 0, specifies that this parameter is not set.
* `enablestrictrx` - Strict PBR check for IPSec packets received through tunnel. Possible values: `YES`, `NO`.
* `enablestricttx` - Strict PBR check for packets to be sent IPSec protected. Possible values: `YES`, `NO`.
* `mac` - The shared MAC used for shared IP between cluster nodes/HA peers.
* `srcip` - Common source-IP address for all tunnels. For a specific tunnel, this global setting is overridden if you have specified another source IP address. Must be a MIP or SNIP address.
* `srciproundrobin` - Use a different source IP address for each new session through a particular IP tunnel, as determined by round robin selection of one of the SNIP addresses. This setting is ignored if a common global source IP address has been specified for all the IP tunnels. This setting does not apply to a tunnel for which a source IP address has been specified. Possible values: `YES`, `NO`.
* `useclientsourceip` - Use client source IP as source IP for outer tunnel IP header. Possible values: `YES`, `NO`.
* `id` - The id of the iptunnelparam. It is a system-generated identifier.
