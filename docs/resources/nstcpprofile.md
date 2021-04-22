---
subcategory: "NS"
---

# Resource: nstcpprofile

The nstcpprofile resource is used to manage the TCP profile in the target ADC.


## Example usage

```hcl
resource "citrixadc_nstcpprofile" "tf_nsprofile" {
    name = "tf_nsprofile"
    ws = "ENABLED"
    ackaggregation = "DISABLED"
}
```


## Argument Reference

* `name` - (Required) Name for a TCP profile. The name of a TCP profile cannot be changed after it is created.
* `ws` - (Optional) Enable or disable window scaling. Possible values: [ ENABLED, DISABLED ]
* `sack` - (Optional) Enable or disable Selective ACKnowledgement (SACK). Possible values: [ ENABLED, DISABLED ]
* `wsval` - (Optional) Factor used to calculate the new window size. This argument is needed only when window scaling is enabled.
* `nagle` - (Optional) Enable or disable the Nagle algorithm on TCP connections. Possible values: [ ENABLED, DISABLED ]
* `ackonpush` - (Optional) Send immediate positive acknowledgement (ACK) on receipt of TCP packets with PUSH flag. Possible values: [ ENABLED, DISABLED ]
* `mss` - (Optional) 
* `maxburst` - (Optional) 
* `initialcwnd` - (Optional) Initial maximum upper limit on the number of TCP packets that can be outstanding on the TCP link to the server.
* `delayedack` - (Optional) Timeout for TCP delayed ACK, in milliseconds.
* `oooqsize` - (Optional) 
* `maxpktpermss` - (Optional) 
* `pktperretx` - (Optional) 
* `minrto` - (Optional) 
* `slowstartincr` - (Optional) Multiplier that determines the rate at which slow start increases the size of the TCP transmission window after each acknowledgement of successful transmission.
* `buffersize` - (Optional) TCP buffering size, in bytes.
* `syncookie` - (Optional) Enable or disable the SYNCOOKIE mechanism for TCP handshake with clients. Disabling SYNCOOKIE prevents SYN attack protection on the Citrix ADC. Possible values: [ ENABLED, DISABLED ]
* `kaprobeupdatelastactivity` - (Optional) Update last activity for the connection after receiving keep-alive (KA) probes. Possible values: [ ENABLED, DISABLED ]
* `flavor` - (Optional) Set TCP congestion control algorithm. Possible values: [ Default, Westwood, BIC, CUBIC, Nile ]
* `dynamicreceivebuffering` - (Optional) Enable or disable dynamic receive buffering. When enabled, allows the receive buffer to be adjusted dynamically based on memory and network conditions. Note: The buffer size argument must be set for dynamic adjustments to take place. Possible values: [ ENABLED, DISABLED ]
* `ka` - (Optional) Send periodic TCP keep-alive (KA) probes to check if peer is still up. Possible values: [ ENABLED, DISABLED ]
* `kaconnidletime` - (Optional) Duration, in seconds, for the connection to be idle, before sending a keep-alive (KA) probe.
* `kamaxprobes` - (Optional) Number of keep-alive (KA) probes to be sent when not acknowledged, before assuming the peer to be down.
* `kaprobeinterval` - (Optional) Time interval, in seconds, before the next keep-alive (KA) probe, if the peer does not respond.
* `sendbuffsize` - (Optional) TCP Send Buffer Size.
* `mptcp` - (Optional) Enable or disable Multipath TCP. Possible values: [ ENABLED, DISABLED ]
* `establishclientconn` - (Optional) Establishing Client Client connection on First data/ Final-ACK / Automatic. Possible values: [ AUTOMATIC, CONN_ESTABLISHED, ON_FIRST_DATA ]
* `tcpsegoffload` - (Optional) Offload TCP segmentation to the NIC. If set to AUTOMATIC, TCP segmentation will be offloaded to the NIC, if the NIC supports it. Possible values: [ AUTOMATIC, DISABLED ]
* `rstwindowattenuate` - (Optional) Enable or disable RST window attenuation to protect against spoofing. When enabled, will reply with corrective ACK when a sequence number is invalid. Possible values: [ ENABLED, DISABLED ]
* `rstmaxack` - (Optional) Enable or disable acceptance of RST that is out of window yet echoes highest ACK sequence number. Useful only in proxy mode. Possible values: [ ENABLED, DISABLED ]
* `spoofsyndrop` - (Optional) Enable or disable drop of invalid SYN packets to protect against spoofing. When disabled, established connections will be reset when a SYN packet is received. Possible values: [ ENABLED, DISABLED ]
* `ecn` - (Optional) Enable or disable TCP Explicit Congestion Notification. Possible values: [ ENABLED, DISABLED ]
* `mptcpdropdataonpreestsf` - (Optional) Enable or disable silently dropping the data on Pre-Established subflow. When enabled, DSS data packets are dropped silently instead of dropping the connection when data is received on pre established subflow. Possible values: [ ENABLED, DISABLED ]
* `mptcpfastopen` - (Optional) Enable or disable Multipath TCP fastopen. When enabled, DSS data packets are accepted before receiving the third ack of SYN handshake. Possible values: [ ENABLED, DISABLED ]
* `mptcpsessiontimeout` - (Optional) MPTCP session timeout in seconds. If this value is not set, idle MPTCP sessions are flushed after vserver's client idle timeout.
* `timestamp` - (Optional) Enable or Disable TCP Timestamp option (RFC 1323). Possible values: [ ENABLED, DISABLED ]
* `dsack` - (Optional) Enable or disable DSACK. Possible values: [ ENABLED, DISABLED ]
* `ackaggregation` - (Optional) Enable or disable ACK Aggregation. Possible values: [ ENABLED, DISABLED ]
* `frto` - (Optional) Enable or disable FRTO (Forward RTO-Recovery). Possible values: [ ENABLED, DISABLED ]
* `maxcwnd` - (Optional) TCP Maximum Congestion Window.
* `fack` - (Optional) Enable or disable FACK (Forward ACK). Possible values: [ ENABLED, DISABLED ]
* `tcpmode` - (Optional) TCP Optimization modes TRANSPARENT / ENDPOINT. Possible values: [ TRANSPARENT, ENDPOINT ]
* `tcpfastopen` - (Optional) Enable or disable TCP Fastopen. When enabled, NS can receive or send Data in SYN or SYN-ACK packets. Possible values: [ ENABLED, DISABLED ]
* `hystart` - (Optional) Enable or disable CUBIC Hystart. Possible values: [ ENABLED, DISABLED ]
* `dupackthresh` - (Optional) TCP dupack threshold.
* `burstratecontrol` - (Optional) TCP Burst Rate Control DISABLED/FIXED/DYNAMIC. FIXED requires a TCP rate to be set. Possible values: [ DISABLED, FIXED, DYNAMIC ]
* `tcprate` - (Optional) TCP connection payload send rate in Kb/s.
* `rateqmax` - (Optional) 
* `drophalfclosedconnontimeout` - (Optional) Silently drop tcp half closed connections on idle timeout. Possible values: [ ENABLED, DISABLED ]
* `dropestconnontimeout` - (Optional) Silently drop tcp established connections on idle timeout. Possible values: [ ENABLED, DISABLED ]
* `applyadaptivetcp` - (Optional) Apply Adaptive TCP optimizations. Possible values: [ ENABLED, DISABLED ]
* `tcpfastopencookiesize` - (Optional) TCP FastOpen Cookie size. This accepts only even numbers. Odd number is trimmed down to nearest even number.
* `taillossprobe` - (Optional) TCP tail loss probe optimizations. Possible values: [ ENABLED, DISABLED ]
* `clientiptcpoption` - (Optional) Client IP in TCP options. Possible values: [ ENABLED, DISABLED ]
* `clientiptcpoptionnumber` - (Optional) ClientIP TCP Option number.
* `mpcapablecbit` - (Optional)


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nstcpprofile. It has the same value as the `name` attribute.


## Import

A nstcpprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_nstcpprofile.tf_nsprofile tf_nsprofile
```
