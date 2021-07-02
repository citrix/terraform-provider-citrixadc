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
* Configuration for TCP profile resource.
*/
type Nstcpprofile struct {
	/**
	* Name for a TCP profile. Must begin with a letter, number, or the underscore \(_\) character. Other characters allowed, after the first character, are the hyphen \(-\), period \(.\), hash \(\#\), space \( \), at \(@\), colon \(:\), and equal \(=\) characters. The name of a TCP profile cannot be changed after it is created.
		CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks \(for example, "my tcp profile" or 'my tcp profile'\).
	*/
	Name string `json:"name,omitempty"`
	/**
	* Enable or disable window scaling.
	*/
	Ws string `json:"ws,omitempty"`
	/**
	* Enable or disable Selective ACKnowledgement (SACK).
	*/
	Sack string `json:"sack,omitempty"`
	/**
	* Factor used to calculate the new window size.
		This argument is needed only when window scaling is enabled.
	*/
	Wsval uint32 `json:"wsval,omitempty"`
	/**
	* Enable or disable the Nagle algorithm on TCP connections.
	*/
	Nagle string `json:"nagle,omitempty"`
	/**
	* Send immediate positive acknowledgement (ACK) on receipt of TCP packets with PUSH flag.
	*/
	Ackonpush string `json:"ackonpush,omitempty"`
	/**
	* Maximum number of octets to allow in a TCP data segment.
	*/
	Mss uint32 `json:"mss,omitempty"`
	/**
	* Maximum number of TCP segments allowed in a burst.
	*/
	Maxburst uint32 `json:"maxburst,omitempty"`
	/**
	* Initial maximum upper limit on the number of TCP packets that can be outstanding on the TCP link to the server.
	*/
	Initialcwnd uint32 `json:"initialcwnd,omitempty"`
	/**
	* Timeout for TCP delayed ACK, in milliseconds.
	*/
	Delayedack uint32 `json:"delayedack,omitempty"`
	/**
	* Maximum size of out-of-order packets queue. A value of 0 means no limit.
	*/
	Oooqsize uint32 `json:"oooqsize,omitempty"`
	/**
	* Maximum number of TCP packets allowed per maximum segment size (MSS).
	*/
	Maxpktpermss uint32 `json:"maxpktpermss,omitempty"`
	/**
	* Maximum limit on the number of packets that should be retransmitted on receiving a partial ACK.
	*/
	Pktperretx uint32 `json:"pktperretx,omitempty"`
	/**
	* Minimum retransmission timeout, in milliseconds, specified in 10-millisecond increments (value must yield a whole number if divided by  10).
	*/
	Minrto uint32 `json:"minrto,omitempty"`
	/**
	* Multiplier that determines the rate at which slow start increases the size of the TCP transmission window after each acknowledgement of successful transmission.
	*/
	Slowstartincr uint32 `json:"slowstartincr,omitempty"`
	/**
	* TCP buffering size, in bytes.
	*/
	Buffersize uint32 `json:"buffersize,omitempty"`
	/**
	* Enable or disable the SYNCOOKIE mechanism for TCP handshake with clients. Disabling SYNCOOKIE prevents SYN attack protection on the Citrix ADC.
	*/
	Syncookie string `json:"syncookie,omitempty"`
	/**
	* Update last activity for the connection after receiving keep-alive (KA) probes.
	*/
	Kaprobeupdatelastactivity string `json:"kaprobeupdatelastactivity,omitempty"`
	/**
	* Set TCP congestion control algorithm.
	*/
	Flavor string `json:"flavor,omitempty"`
	/**
	* Enable or disable dynamic receive buffering. When enabled, allows the receive buffer to be adjusted dynamically based on memory and network conditions.
		Note: The buffer size argument must be set for dynamic adjustments to take place.
	*/
	Dynamicreceivebuffering string `json:"dynamicreceivebuffering,omitempty"`
	/**
	* Send periodic TCP keep-alive (KA) probes to check if peer is still up.
	*/
	Ka string `json:"ka,omitempty"`
	/**
	* Duration, in seconds, for the connection to be idle, before sending a keep-alive (KA) probe.
	*/
	Kaconnidletime uint32 `json:"kaconnidletime,omitempty"`
	/**
	* Number of keep-alive (KA) probes to be sent when not acknowledged, before assuming the peer to be down.
	*/
	Kamaxprobes uint32 `json:"kamaxprobes,omitempty"`
	/**
	* Time interval, in seconds, before the next keep-alive (KA) probe, if the peer does not respond.
	*/
	Kaprobeinterval uint32 `json:"kaprobeinterval,omitempty"`
	/**
	* TCP Send Buffer Size
	*/
	Sendbuffsize uint32 `json:"sendbuffsize,omitempty"`
	/**
	* Enable or disable Multipath TCP.
	*/
	Mptcp string `json:"mptcp,omitempty"`
	/**
	* Establishing Client Client connection on First data/ Final-ACK / Automatic
	*/
	Establishclientconn string `json:"establishclientconn,omitempty"`
	/**
	* Offload TCP segmentation to the NIC. If set to AUTOMATIC, TCP segmentation will be offloaded to the NIC, if the NIC supports it.
	*/
	Tcpsegoffload string `json:"tcpsegoffload,omitempty"`
	/**
	* Enable or disable RST window attenuation to protect against spoofing. When enabled, will reply with corrective ACK when a sequence number is invalid.
	*/
	Rstwindowattenuate string `json:"rstwindowattenuate,omitempty"`
	/**
	* Enable or disable acceptance of RST that is out of window yet echoes highest ACK sequence number. Useful only in proxy mode.
	*/
	Rstmaxack string `json:"rstmaxack,omitempty"`
	/**
	* Enable or disable drop of invalid SYN packets to protect against spoofing. When disabled, established connections will be reset when a SYN packet is received.
	*/
	Spoofsyndrop string `json:"spoofsyndrop,omitempty"`
	/**
	* Enable or disable TCP Explicit Congestion Notification.
	*/
	Ecn string `json:"ecn,omitempty"`
	/**
	* Enable or disable silently dropping the data on Pre-Established subflow. When enabled, DSS data packets are dropped silently instead of dropping the connection when data is received on pre established subflow.
	*/
	Mptcpdropdataonpreestsf string `json:"mptcpdropdataonpreestsf,omitempty"`
	/**
	* Enable or disable Multipath TCP fastopen. When enabled, DSS data packets are accepted before receiving the third ack of SYN handshake.
	*/
	Mptcpfastopen string `json:"mptcpfastopen,omitempty"`
	/**
	* MPTCP session timeout in seconds. If this value is not set, idle MPTCP sessions are flushed after vserver's client idle timeout.
	*/
	Mptcpsessiontimeout uint32 `json:"mptcpsessiontimeout,omitempty"`
	/**
	* Enable or Disable TCP Timestamp option (RFC 1323)
	*/
	Timestamp string `json:"timestamp,omitempty"`
	/**
	* Enable or disable DSACK.
	*/
	Dsack string `json:"dsack,omitempty"`
	/**
	* Enable or disable ACK Aggregation.
	*/
	Ackaggregation string `json:"ackaggregation,omitempty"`
	/**
	* Enable or disable FRTO (Forward RTO-Recovery).
	*/
	Frto string `json:"frto,omitempty"`
	/**
	* TCP Maximum Congestion Window.
	*/
	Maxcwnd uint32 `json:"maxcwnd,omitempty"`
	/**
	* Enable or disable FACK (Forward ACK).
	*/
	Fack string `json:"fack,omitempty"`
	/**
	* TCP Optimization modes TRANSPARENT / ENDPOINT.
	*/
	Tcpmode string `json:"tcpmode,omitempty"`
	/**
	* Enable or disable TCP Fastopen. When enabled, NS can receive or send Data in SYN or SYN-ACK packets.
	*/
	Tcpfastopen string `json:"tcpfastopen,omitempty"`
	/**
	* Enable or disable CUBIC Hystart
	*/
	Hystart string `json:"hystart,omitempty"`
	/**
	* TCP dupack threshold.
	*/
	Dupackthresh uint32 `json:"dupackthresh,omitempty"`
	/**
	* TCP Burst Rate Control DISABLED/FIXED/DYNAMIC. FIXED requires a TCP rate to be set.
	*/
	Burstratecontrol string `json:"burstratecontrol,omitempty"`
	/**
	* TCP connection payload send rate in Kb/s
	*/
	Tcprate uint32 `json:"tcprate,omitempty"`
	/**
	* Maximum connection queue size in bytes, when BurstRateControl is used
	*/
	Rateqmax uint32 `json:"rateqmax,omitempty"`
	/**
	* Silently drop tcp half closed connections on idle timeout
	*/
	Drophalfclosedconnontimeout string `json:"drophalfclosedconnontimeout,omitempty"`
	/**
	* Silently drop tcp established connections on idle timeout
	*/
	Dropestconnontimeout string `json:"dropestconnontimeout,omitempty"`
	/**
	* Apply Adaptive TCP optimizations
	*/
	Applyadaptivetcp string `json:"applyadaptivetcp,omitempty"`
	/**
	* TCP FastOpen Cookie size. This accepts only even numbers. Odd number is trimmed down to nearest even number.
	*/
	Tcpfastopencookiesize uint32 `json:"tcpfastopencookiesize,omitempty"`
	/**
	* TCP tail loss probe optimizations
	*/
	Taillossprobe string `json:"taillossprobe,omitempty"`
	/**
	* Client IP in TCP options
	*/
	Clientiptcpoption string `json:"clientiptcpoption,omitempty"`
	/**
	* ClientIP TCP Option number
	*/
	Clientiptcpoptionnumber uint32 `json:"clientiptcpoptionnumber,omitempty"`
	/**
	* Set C bit in MP-CAPABLE Syn-Ack sent by Citrix ADC
	*/
	Mpcapablecbit string `json:"mpcapablecbit,omitempty"`

	//------- Read only Parameter ---------;

	Refcnt string `json:"refcnt,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`

}
