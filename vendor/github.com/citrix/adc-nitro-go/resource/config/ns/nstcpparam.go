/*
* Copyright (c) 2021 Citrix Systems, Inc.
*
*   Licensed under the Apache License, Version 2.0 (the "License");
*   you may not use this file except in compliance with the License.
*   You may obtain a copy of the License at
*
*       http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*   distributed under the License is distributed on an "AS IS" BASIS,
*   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*   See the License for the specific language governing permissions and
*   limitations under the License.
*/

package ns

/**
* Configuration for tcp parameters resource.
*/
type Nstcpparam struct {
	/**
	* Enable or disable window scaling.
	*/
	Ws string `json:"ws,omitempty"`
	/**
	* Factor used to calculate the new window size.
		This argument is needed only when the window scaling is enabled.
	*/
	Wsval uint32 `json:"wsval,omitempty"`
	/**
	* Enable or disable Selective ACKnowledgement (SACK).
	*/
	Sack string `json:"sack,omitempty"`
	/**
	* Enable or disable maximum segment size (MSS) learning for virtual servers.
	*/
	Learnvsvrmss string `json:"learnvsvrmss,omitempty"`
	/**
	* Maximum number of TCP segments allowed in a burst.
	*/
	Maxburst uint32 `json:"maxburst,omitempty"`
	/**
	* Initial maximum upper limit on the number of TCP packets that can be outstanding on the TCP link to the server.
	*/
	Initialcwnd uint32 `json:"initialcwnd,omitempty"`
	/**
	* TCP Receive buffer size
	*/
	Recvbuffsize uint32 `json:"recvbuffsize,omitempty"`
	/**
	* Timeout for TCP delayed ACK, in milliseconds.
	*/
	Delayedack uint32 `json:"delayedack,omitempty"`
	/**
	* Flag to switch on RST on down services.
	*/
	Downstaterst string `json:"downstaterst,omitempty"`
	/**
	* Enable or disable the Nagle algorithm on TCP connections.
	*/
	Nagle string `json:"nagle,omitempty"`
	/**
	* Limit the number of persist (zero window) probes.
	*/
	Limitedpersist string `json:"limitedpersist,omitempty"`
	/**
	* Maximum size of out-of-order packets queue. A value of 0 means no limit.
	*/
	Oooqsize uint32 `json:"oooqsize,omitempty"`
	/**
	* Send immediate positive acknowledgement (ACK) on receipt of TCP packets with PUSH flag.
	*/
	Ackonpush string `json:"ackonpush,omitempty"`
	/**
	* Maximum number of TCP packets allowed per maximum segment size (MSS).
	*/
	Maxpktpermss uint32 `json:"maxpktpermss,omitempty"`
	/**
	* Maximum limit on the number of packets that should be retransmitted on receiving a partial ACK.
	*/
	Pktperretx int32 `json:"pktperretx,omitempty"`
	/**
	* Minimum retransmission timeout, in milliseconds, specified in 10-millisecond increments (value must yield a whole number if divided by 10).
	*/
	Minrto int32 `json:"minrto,omitempty"`
	/**
	* Multiplier that determines the rate at which slow start increases the size of the TCP transmission window after each acknowledgement of successful transmission.
	*/
	Slowstartincr int32 `json:"slowstartincr,omitempty"`
	/**
	* Maximum number of probes that Citrix ADC can send out in 10 milliseconds, to dynamically learn a service. Citrix ADC probes for the existence of the origin in case of wildcard virtual server or services.
	*/
	Maxdynserverprobes uint32 `json:"maxdynserverprobes,omitempty"`
	/**
	* Maximum threshold. After crossing this threshold number of outstanding probes for origin, the Citrix ADC reduces the number of connection retries for probe connections.
	*/
	Synholdfastgiveup uint32 `json:"synholdfastgiveup,omitempty"`
	/**
	* Limit the number of client connections (SYN) waiting for status of single probe. Any new SYN packets will be dropped.
	*/
	Maxsynholdperprobe uint32 `json:"maxsynholdperprobe,omitempty"`
	/**
	* Limit the number of client connections (SYN) waiting for status of probe system wide. Any new SYN packets will be dropped.
	*/
	Maxsynhold uint32 `json:"maxsynhold,omitempty"`
	/**
	* Duration, in seconds, to sample the Maximum Segment Size (MSS) of the services. The Citrix ADC determines the best MSS to set for the virtual server based on this sampling. The argument to enable maximum segment size (MSS) for virtual servers must be enabled.
	*/
	Msslearninterval uint32 `json:"msslearninterval,omitempty"`
	/**
	* Frequency, in seconds, at which the virtual servers learn the Maximum segment size (MSS) from the services. The argument to enable maximum segment size (MSS) for virtual servers must be enabled.
	*/
	Msslearndelay uint32 `json:"msslearndelay,omitempty"`
	/**
	* Maximum number of connections to hold in the TCP TIME_WAIT state on a packet engine. New connections entering TIME_WAIT state are proactively cleaned up.
	*/
	Maxtimewaitconn uint32 `json:"maxtimewaitconn,omitempty"`
	/**
	* Update last activity for KA probes
	*/
	Kaprobeupdatelastactivity string `json:"kaprobeupdatelastactivity,omitempty"`
	/**
	* When 'syncookie' is disabled in the TCP profile that is bound to the virtual server or service, and the number of TCP SYN+ACK retransmission by Citrix ADC for that virtual server or service crosses this threshold, the Citrix ADC responds by using the TCP SYN-Cookie mechanism.
	*/
	Maxsynackretx uint32 `json:"maxsynackretx,omitempty"`
	/**
	* Detect TCP SYN packet flood and send an SNMP trap.
	*/
	Synattackdetection string `json:"synattackdetection,omitempty"`
	/**
	* Flush an existing connection if no memory can be obtained for new connection.
		HALF_CLOSED_AND_IDLE: Flush a connection that is closed by us but not by peer, or failing that, a connection that is past configured idle time.  New connection fails if no such connection can be found.
		FIFO: If no half-closed or idle connection can be found, flush the oldest non-management connection, even if it is active.  New connection fails if the oldest few connections are management connections.
		Note: If you enable this setting, you should also consider lowering the zombie timeout and half-close timeout, while setting the Citrix ADC timeout.
		See Also: connFlushThres argument below.
	*/
	Connflushifnomem string `json:"connflushifnomem,omitempty"`
	/**
	* Flush an existing connection (as configured through -connFlushIfNoMem FIFO) if the system has more than specified number of connections, and a new connection is to be established.  Note: This value may be rounded down to be a whole multiple of the number of packet engines running.
	*/
	Connflushthres uint32 `json:"connflushthres,omitempty"`
	/**
	* Accept DATA_FIN/FAST_CLOSE on passive subflow
	*/
	Mptcpconcloseonpassivesf string `json:"mptcpconcloseonpassivesf,omitempty"`
	/**
	* Use MPTCP DSS checksum
	*/
	Mptcpchecksum string `json:"mptcpchecksum,omitempty"`
	/**
	* The timeout value in seconds for idle mptcp subflows. If this timeout is not set, idle subflows are cleared after cltTimeout of vserver
	*/
	Mptcpsftimeout uint64 `json:"mptcpsftimeout,omitempty"`
	/**
	* The minimum idle time value in seconds for idle mptcp subflows after which the sublow is replaced by new incoming subflow if maximum subflow limit is reached. The priority for replacement is given to those subflow without any transaction
	*/
	Mptcpsfreplacetimeout uint64 `json:"mptcpsfreplacetimeout,omitempty"`
	/**
	* Maximum number of subflow connections supported in established state per mptcp connection.
	*/
	Mptcpmaxsf uint32 `json:"mptcpmaxsf,omitempty"`
	/**
	* Maximum number of subflow connections supported in pending join state per mptcp connection.
	*/
	Mptcpmaxpendingsf uint32 `json:"mptcpmaxpendingsf,omitempty"`
	/**
	* Maximum system level pending join connections allowed.
	*/
	Mptcppendingjointhreshold uint32 `json:"mptcppendingjointhreshold,omitempty"`
	/**
	* Number of RTO's at subflow level, after which MPCTP should start using other subflow.
	*/
	Mptcprtostoswitchsf uint32 `json:"mptcprtostoswitchsf,omitempty"`
	/**
	* When enabled, if NS receives a DSS on a backup subflow, NS will start using that subflow to send data. And if disabled, NS will continue to transmit on current chosen subflow. In case there is some error on a subflow (like RTO's/RST etc.) then NS can choose a backup subflow irrespective of this tunable.
	*/
	Mptcpusebackupondss string `json:"mptcpusebackupondss,omitempty"`
	/**
	* Number of RTO's after which a connection should be freed.
	*/
	Tcpmaxretries uint32 `json:"tcpmaxretries,omitempty"`
	/**
	* Allow subflows to close immediately on FIN before the DATA_FIN exchange is completed at mptcp level.
	*/
	Mptcpimmediatesfcloseonfin string `json:"mptcpimmediatesfcloseonfin,omitempty"`
	/**
	* Allow to send DATA FIN or FAST CLOSE on mptcp connection while sending FIN or RST on the last subflow.
	*/
	Mptcpclosemptcpsessiononlastsfclose string `json:"mptcpclosemptcpsessiononlastsfclose,omitempty"`
	/**
	* Allow MPTCP subflows to send TCP RST Reason (MP_TCPRST) Option while sending TCP RST.
	*/
	Mptcpsendsfresetoption string `json:"mptcpsendsfresetoption,omitempty"`
	/**
	* Allow to select option ACK or RESET to force the closure of an MPTCP connection abruptly.
	*/
	Mptcpfastcloseoption string `json:"mptcpfastcloseoption,omitempty"`
	/**
	* Timeout in seconds after which a new TFO Key is computed for generating TFO Cookie. If zero, the same key is used always. If timeout is less than 120seconds, NS defaults to 120seconds timeout.
	*/
	Tcpfastopencookietimeout uint64 `json:"tcpfastopencookietimeout,omitempty"`
	/**
	* Timeout for the server to function in syncookie mode after the synattack. This is valid if TCP syncookie is disabled on the profile and server acts in non syncookie mode by default.
	*/
	Autosyncookietimeout uint32 `json:"autosyncookietimeout,omitempty"`
	/**
	* The amount of time in seconds, after which a TCP connnection in the TCP TIME-WAIT state is flushed.
	*/
	Tcpfintimeout uint64 `json:"tcpfintimeout,omitempty"`
	/**
	* If enabled, non-negotiated TCP options are removed from the received packet while proxying it. By default, non-negotiated TCP options would be replaced by NOPs in the proxied packets. This option is not applicable for Citrix ADC generated packets.
	*/
	Compacttcpoptionnoop string `json:"compacttcpoptionnoop,omitempty"`

	//------- Read only Parameter ---------;

	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`

}
