---
subcategory: "NS"
---

# Resource: nstcpparam

The nstcpparam resource is used to update the ADC tcp parameters.


## Example usage

```hcl
resource "citrixadc_nstcpparam" "tf_tcpparam" {
  ws                                  = "ENABLED"
  wsval                               = 8
  sack                                = "ENABLED"
  learnvsvrmss                        = "DISABLED"
  maxburst                            = 6
  initialcwnd                         = 10
  delayedack                          = 100
  downstaterst                        = "DISABLED"
  nagle                               = "DISABLED"
  limitedpersist                      = "ENABLED"
  oooqsize                            = 300
  ackonpush                           = "ENABLED"
  maxpktpermss                        = 0
  pktperretx                          = 1
  minrto                              = 1000
  slowstartincr                       = 2
  maxdynserverprobes                  = 7
  synholdfastgiveup                   = 1024
  maxsynholdperprobe                  = 128
  maxsynhold                          = 16384
  msslearninterval                    = 180
  msslearndelay                       = 3600
  maxtimewaitconn                     = 7000
  maxsynackretx                       = 100
  synattackdetection                  = "ENABLED"
  connflushifnomem                    = "NONE "
  connflushthres                      = 4294967295
  mptcpconcloseonpassivesf            = "ENABLED"
  mptcpchecksum                       = "ENABLED"
  mptcpsftimeout                      = 0
  mptcpsfreplacetimeout               = 10
  mptcpmaxsf                          = 4
  mptcpmaxpendingsf                   = 4
  mptcppendingjointhreshold           = 0
  mptcprtostoswitchsf                 = 2
  mptcpusebackupondss                 = "ENABLED"
  tcpmaxretries                       = 7
  mptcpimmediatesfcloseonfin          = "DISABLED"
  mptcpclosemptcpsessiononlastsfclose = "DISABLED"
  tcpfastopencookietimeout            = 0
  autosyncookietimeout                = 30
  tcpfintimeout                       = 40
}
```


## Argument Reference

* `ws` - (Optional) Enable or disable window scaling. Possible values: [ ENABLED, DISABLED ]
* `wsval` - (Optional) Factor used to calculate the new window size. This argument is needed only when the window scaling is enabled.
* `sack` - (Optional) Enable or disable Selective ACKnowledgement (SACK). Possible values: [ ENABLED, DISABLED ]
* `learnvsvrmss` - (Optional) Enable or disable maximum segment size (MSS) learning for virtual servers. Possible values: [ ENABLED, DISABLED ]
* `maxburst` - (Optional) 
* `initialcwnd` - (Optional) Initial maximum upper limit on the number of TCP packets that can be outstanding on the TCP link to the server.
* `recvbuffsize` - (Optional) TCP Receive buffer size.
* `delayedack` - (Optional) Timeout for TCP delayed ACK, in milliseconds.
* `downstaterst` - (Optional) Flag to switch on RST on down services. Possible values: [ ENABLED, DISABLED ]
* `nagle` - (Optional) Enable or disable the Nagle algorithm on TCP connections. Possible values: [ ENABLED, DISABLED ]
* `limitedpersist` - (Optional) Limit the number of persist (zero window) probes. Possible values: [ ENABLED, DISABLED ]
* `oooqsize` - (Optional) 
* `ackonpush` - (Optional) Send immediate positive acknowledgement (ACK) on receipt of TCP packets with PUSH flag. Possible values: [ ENABLED, DISABLED ]
* `maxpktpermss` - (Optional) 
* `pktperretx` - (Optional) 
* `minrto` - (Optional) 
* `slowstartincr` - (Optional) Multiplier that determines the rate at which slow start increases the size of the TCP transmission window after each acknowledgement of successful transmission.
* `maxdynserverprobes` - (Optional) 
* `synholdfastgiveup` - (Optional) 
* `maxsynholdperprobe` - (Optional) Limit the number of client connections (SYN) waiting for status of single probe. Any new SYN packets will be dropped.
* `maxsynhold` - (Optional) Limit the number of client connections (SYN) waiting for status of probe system wide. Any new SYN packets will be dropped.
* `msslearninterval` - (Optional) Duration, in seconds, to sample the Maximum Segment Size (MSS) of the services. The Citrix ADC determines the best MSS to set for the virtual server based on this sampling. The argument to enable maximum segment size (MSS) for virtual servers must be enabled.
* `msslearndelay` - (Optional) Frequency, in seconds, at which the virtual servers learn the Maximum segment size (MSS) from the services. The argument to enable maximum segment size (MSS) for virtual servers must be enabled.
* `maxtimewaitconn` - (Optional) 
* `kaprobeupdatelastactivity` - (Optional) Update last activity for KA probes. Possible values: [ ENABLED, DISABLED ]
* `maxsynackretx` - (Optional) When 'syncookie' is disabled in the TCP profile that is bound to the virtual server or service, and the number of TCP SYN+ACK retransmission by Citrix ADC for that virtual server or service crosses this threshold, the Citrix ADC responds by using the TCP SYN-Cookie mechanism.
* `synattackdetection` - (Optional) Detect TCP SYN packet flood and send an SNMP trap. Possible values: [ ENABLED, DISABLED ]
* `connflushifnomem` - (Optional) Flush an existing connection if no memory can be obtained for new connection. HALF_CLOSED_AND_IDLE: Flush a connection that is closed by us but not by peer, or failing that, a connection that is past configured idle time.  New connection fails if no such connection can be found. FIFO: If no half-closed or idle connection can be found, flush the oldest non-management connection, even if it is active.  New connection fails if the oldest few connections are management connections. Note: If you enable this setting, you should also consider lowering the zombie timeout and half-close timeout, while setting the Citrix ADC timeout. See Also: connFlushThres argument below. Possible values: [ NONE , HALFCLOSED_AND_IDLE, FIFO ]
* `connflushthres` - (Optional) Flush an existing connection (as configured through -connFlushIfNoMem FIFO) if the system has more than specified number of connections, and a new connection is to be established.  Note: This value may be rounded down to be a whole multiple of the number of packet engines running.
* `mptcpconcloseonpassivesf` - (Optional) Accept DATA_FIN/FAST_CLOSE on passive subflow. Possible values: [ ENABLED, DISABLED ]
* `mptcpchecksum` - (Optional) Use MPTCP DSS checksum. Possible values: [ ENABLED, DISABLED ]
* `mptcpsftimeout` - (Optional) The timeout value in seconds for idle mptcp subflows. If this timeout is not set, idle subflows are cleared after cltTimeout of vserver.
* `mptcpsfreplacetimeout` - (Optional) The minimum idle time value in seconds for idle mptcp subflows after which the sublow is replaced by new incoming subflow if maximum subflow limit is reached. The priority for replacement is given to those subflow without any transaction.
* `mptcpmaxsf` - (Optional) 
* `mptcpmaxpendingsf` - (Optional) 
* `mptcppendingjointhreshold` - (Optional) 
* `mptcprtostoswitchsf` - (Optional) Number of RTO's at subflow level, after which MPCTP should start using other subflow.
* `mptcpusebackupondss` - (Optional) When enabled, if NS receives a DSS on a backup subflow, NS will start using that subflow to send data. And if disabled, NS will continue to transmit on current chosen subflow. In case there is some error on a subflow (like RTO's/RST etc.) then NS can choose a backup subflow irrespective of this tunable. Possible values: [ ENABLED, DISABLED ]
* `tcpmaxretries` - (Optional) Number of RTO's after which a connection should be freed.
* `mptcpimmediatesfcloseonfin` - (Optional) Allow subflows to close immediately on FIN before the DATA_FIN exchange is completed at mptcp level. Possible values: [ ENABLED, DISABLED ]
* `mptcpclosemptcpsessiononlastsfclose` - (Optional) Allow to send DATA FIN or FAST CLOSE on mptcp connection while sending FIN or RST on the last subflow. Possible values: [ ENABLED, DISABLED ]
* `tcpfastopencookietimeout` - (Optional) Timeout in seconds after which a new TFO Key is computed for generating TFO Cookie. If zero, the same key is used always. If timeout is less than 120seconds, NS defaults to 120seconds timeout.
* `autosyncookietimeout` - (Optional) Timeout for the server to function in syncookie mode after the synattack. This is valid if TCP syncookie is disabled on the profile and server acts in non syncookie mode by default.
* `tcpfintimeout` - (Optional) The amount of time in seconds, after which a TCP connnection in the TCP TIME-WAIT state is flushed.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nstcpparam. It is a unique string prefixed with "tf-nstcpparam-"
