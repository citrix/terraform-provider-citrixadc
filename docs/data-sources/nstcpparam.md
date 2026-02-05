---
subcategory: "Network"
---

# Data Source `nstcpparam`

The nstcpparam data source allows you to retrieve information about global TCP parameters configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_nstcpparam" "tf_nstcpparam" {
}

output "delayedack" {
  value = data.citrixadc_nstcpparam.tf_nstcpparam.delayedack
}

output "maxburst" {
  value = data.citrixadc_nstcpparam.tf_nstcpparam.maxburst
}
```


## Argument Reference

This data source has no required arguments.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `ackonpush` - Send immediate positive acknowledgement (ACK) on receipt of TCP packets with PUSH flag.
* `autosyncookietimeout` - Timeout for the server to function in syncookie mode after the synattack.
* `compacttcpoptionnoop` - If enabled, non-negotiated TCP options are removed from the received packet while proxying it.
* `connflushifnomem` - Flush an existing connection if no memory can be obtained for new connection.
* `connflushthres` - Flush an existing connection if the system has more than specified number of connections.
* `delayedack` - Timeout for TCP delayed ACK, in milliseconds.
* `delinkclientserveronrst` - If enabled, Delink client and server connection, when there is outstanding data.
* `downstaterst` - Flag to switch on RST on down services.
* `enhancedisngeneration` - If enabled, increase the ISN variation in SYN-ACKs sent by the NetScaler.
* `initialcwnd` - Initial maximum upper limit on the number of TCP packets outstanding on the TCP link to the server.
* `kaprobeupdatelastactivity` - Update last activity for KA probes.
* `learnvsvrmss` - Enable or disable maximum segment size (MSS) learning for virtual servers.
* `limitedpersist` - Limit the number of persist (zero window) probes.
* `maxburst` - Maximum number of TCP segments allowed in a burst.
* `maxdynserverprobes` - Maximum number of probes that Citrix ADC can send out in 10 milliseconds.
* `maxpktpermss` - Maximum number of TCP packets allowed per maximum segment size (MSS).
* `maxsynackretx` - Maximum number of TCP SYN+ACK retransmission threshold.
* `maxsynhold` - Maximum number of half-open TCP connections.
* `maxtimewaitconn` - Maximum number of connections in TIME_WAIT state.
* `minrto` - Minimum retransmission timeout, in milliseconds.
* `mptcpchecksum` - Use MPTCP DSS checksum.
* `mptcpclosemptcpsessiononlastsfclose` - MPTCP session is closed when last subflow is closed.
* `mptcpconcloseonpassivesf` - Connection close is propagated to all subflows.
* `mptcpfastcloseoption` - Fast close option for MPTCP.
* `mptcpimmediatesfcloseonfin` - Allow subflow FIN even when data pending.
* `mptcpmaxpendingsf` - Maximum pending subflows per MPTCP connection.
* `mptcpmaxsf` - Maximum subflows per MPTCP connection.
* `mptcppendingjointhreshold` - Pending join threshold.
* `mptcpreliableaddaddr` - Reliable ADDADDR signal.
* `mptcprtostoswitchsf` - RTO threshold for subflow switching.
* `mptcpsendsfresetoption` - Send RST on subflow close.
* `mptcpsfreplacetimeout` - Subflow replace timeout.
* `mptcpsftimeout` - Subflow timeout.
* `mptcpusebackupondss` - Use backup subflow when DSS lost.
* `msslearndelay` - Frequency at which the virtual servers learn the MSS from the services.
* `msslearninterval` - Duration for which msslearndelay is in effect.
* `nagle` - Enable or disable the Nagle algorithm on TCP connections.
* `oooqsize` - Maximum size of out-of-order packets queue.
* `pktperretx` - Maximum limit on the number of packets that should be retransmitted on receiving a partial ACK.
* `rfc5961chlgacklimit` - RFC 5961 challenge ACK limit.
* `sack` - Enable or disable Selective ACKnowledgement (SACK).
* `slowstartincr` - Multiplier for the slow-start window size.
* `synattackdetection` - Detect TCP SYN packet flood and send an SNMP trap.
* `synholdfastgiveup` - Limit the number of client connections allowed per IP address.
* `tcpfastopencookietimeout` - Timeout for TCP Fast Open cookies.
* `tcpfintimeout` - TCP FIN timeout.
* `tcpmaxretries` - Maximum number of retries for TCP.
* `ws` - Enable or disable window scaling.
* `wsval` - Window scaling factor.
