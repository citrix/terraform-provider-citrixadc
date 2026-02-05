---
subcategory: "Network"
---

# Data Source `nstcpprofile`

The nstcpprofile data source allows you to retrieve information about a TCP profile configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_nstcpprofile" "tf_nstcpprofile" {
  name = "test_profile"
}

output "ws" {
  value = data.citrixadc_nstcpprofile.tf_nstcpprofile.ws
}

output "ackaggregation" {
  value = data.citrixadc_nstcpprofile.tf_nstcpprofile.ackaggregation
}
```


## Argument Reference

* `name` - (Required) Name for a TCP profile.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `ackaggregation` - Enable or disable ACK Aggregation.
* `ackonpush` - Send immediate positive acknowledgement (ACK) on receipt of TCP packets with PUSH flag.
* `applyadaptivetcp` - Apply Adaptive TCP optimizations.
* `buffersize` - TCP buffering size, in bytes.
* `burstratecontrol` - TCP Burst Rate Control DISABLED/FIXED/DYNAMIC.
* `clientiptcpoption` - Client IP in TCP options.
* `clientiptcpoptionnumber` - ClientIP TCP Option number.
* `delayedack` - Timeout for TCP delayed ACK, in milliseconds.
* `downstaterst` - Flag to switch on RST on down services.
* `dropestconnonfin` - Silently drop tcp half closed connections.
* `drophalfclosedconnontimeout` - Silently drop tcp half closed connections on idle timeout.
* `dsack` - Enable or disable DSACK.
* `dupackthresh` - TCP dupack threshold.
* `dynamicreceivebuffering` - Enable or disable dynamic receive buffering.
* `ecn` - Enable or disable TCP Explicit Congestion Notification.
* `establishclientconn` - Establishing Client Client connection on First data/ Final-ACK / Automatic.
* `fack` - Enable or disable FACK (Forward ACK).
* `flavor` - Set TCP congestion control algorithm.
* `frto` - Enable or disable FRTO (Forward RTO-Recovery).
* `hystart` - Enable or disable HYSTART.
* `initialcwnd` - Initial maximum upper limit on the number of TCP packets outstanding.
* `ka` - Send periodic TCP keep-alive (KA) probes.
* `kaconnidletime` - Duration for which a connection can be idle, after which a KA probe is sent.
* `kamaxprobes` - Maximum number of TCP keep-alive probes to be sent.
* `kaprobeinterval` - Time interval between TCP keep-alive probes.
* `kaprobeupdatelastactivity` - Update last activity for KA probes.
* `maxburst` - Maximum number of TCP segments allowed in a burst.
* `maxcwnd` - TCP Maximum Congestion Window.
* `maxpktpermss` - Maximum number of TCP packets allowed per maximum segment size (MSS).
* `minrto` - Minimum retransmission timeout, in milliseconds.
* `mpcapablecbit` - MPTCP Capable C bit.
* `mptcp` - Enable or disable Multipath TCP.
* `mptcpdropdataonpreestsf` - Drop data on Pre-Established subflow.
* `mptcpfastopen` - Enable or disable Multipath TCP fastopen.
* `mptcpsessiontimeout` - MPTCP session timeout.
* `mss` - Maximum segment size.
* `nagle` - Enable or disable the Nagle algorithm on TCP connections.
* `oooqsize` - Maximum size of out-of-order packets queue.
* `pktperretx` - Maximum limit on the number of packets retransmitted.
* `rateqmax` - Maximum connection queue size in bytes.
* `rstmaxack` - Enable or disable acceptance of RST that is out of window.
* `rstwindowattenuate` - Enable or disable RST window attenuation.
* `sack` - Enable or disable Selective ACKnowledgement (SACK).
* `sendbuffsize` - TCP Send Buffer Size.
* `sendclientportintcpoption` - Send Client Port number in TCP options.
* `slowstartincr` - Multiplier for the slow-start window size.
* `spoofsyndrop` - Enable or disable drop of invalid SYN packets.
* `syncookie` - Enable or disable SYN Cookie for TCP handshake.
* `taillossprobe` - TCP Tail Loss Probe optimizations.
* `tcpfastopen` - Enable or disable TCP Fast Open.
* `tcpfastopencookiesize` - TCP Fast Open Cookie size.
* `tcpfastopencookietimeout` - Timeout for TCP Fast Open cookies.
* `tcpmode` - TCP Optimization modes TRANSPARENT / ENDPOINT.
* `tcprate` - TCP connection payload send rate in Kb/s.
* `tcpsegoffload` - Offload TCP segmentation to the NIC.
* `timestamp` - Enable or disable TCP timestamps.
* `ws` - Enable or disable window scaling.
* `wsval` - Window scaling factor.
