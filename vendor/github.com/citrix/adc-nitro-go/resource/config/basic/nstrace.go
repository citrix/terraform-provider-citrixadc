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

package basic

/**
* Configuration for nstrace operations resource.
*/
type Nstrace struct {
	/**
	* Number of files to be generated in cycle.
	*/
	Nf *int `json:"nf,omitempty"`
	/**
	* Time per file (sec).
	*/
	Time *int `json:"time,omitempty"`
	/**
	* Size of the captured data. Set 0 for full packet trace.
	*/
	Size *int `json:"size,omitempty"`
	/**
	* Capturing mode for trace. Mode can be any of the following values or combination of these values:
		RX          Received packets before NIC pipelining (Filter does not work when RX capturing mode is ON)
		NEW_RX      Received packets after NIC pipelining
		TX          Transmitted packets
		TXB         Packets buffered for transmission
		IPV6        Translated IPv6 packets
		C2C         Capture C2C message
		NS_FR_TX    TX/TXB packets are not captured in flow receiver.
		MPTCP       MPTCP master flow
		HTTP_QUIC   HTTP-over-QUIC stream data and stream events
		Default mode: NEW_RX TXB 
	*/
	Mode []string `json:"mode,omitempty"`
	/**
	* Use separate trace files for each interface. Works only with cap format.
	*/
	Pernic string `json:"pernic,omitempty"`
	/**
	* Name of the trace file.
	*/
	Filename string `json:"filename,omitempty"`
	/**
	* ID for the trace file name for uniqueness. Should be used only with -name option.
	*/
	Fileid string `json:"fileid,omitempty"`
	/**
	* Filter expression for nstrace. Maximum length of filter is 255 and it can be of following format:
		<expression> [<relop> <expression>]
		<relop> = ( && | || )
		<expression> =:
		CONNECTION.<qualifier>.<qualifier-method>.(<qualifier-value>)
		<qualifier> = SRCIP
		<qualifier-method> = [ EQ | NE ]
		<qualifier-value>  = A valid IPv4 address.
		example = CONNECTION.SRCIP.EQ(127.0.0.1)
		<qualifier> = DSTIP
		<qualifier-method> = [ EQ | NE ]
		<qualifier-value>  = A valid IPv4 address.
		example = CONNECTION.DSTIP.EQ(127.0.0.1)
		<qualifier> = IP
		<qualifier-method> = [ EQ | NE ]
		<qualifier-value>  = A valid IPv4 address.
		example = CONNECTION.IP.EQ(127.0.0.1)
		<qualifier> = SRCIPv6
		<qualifier-method> = [ EQ | NE ]
		<qualifier-value>  = A valid IPv6 address.
		example = CONNECTION.SRCIPv6.EQ(2001:db8:0:0:1::1)
		<qualifier> = DSTIPv6
		<qualifier-method> = [ EQ | NE ]
		<qualifier-value>  = A valid IPv6 address.
		example = CONNECTION.DSTIPv6.EQ(2001:db8:0:0:1::1)
		<qualifier> = IPv6
		<qualifier-method> = [ EQ | NE ]
		<qualifier-value>  = A valid IPv6 address.
		example = CONNECTION.IPv6.EQ(2001:db8:0:0:1::1)
		<qualifier> = SRCPORT
		<qualifier-method> = [ EQ | NE | GT | GE | LT | LE
		| BETWEEN ]
		<qualifier-value>  = A valid port number.
		example = CONNECTION.SRCPORT.EQ(80)
		<qualifier> = DSTPORT
		<qualifier-method> = [ EQ | NE | GT | GE | LT | LE
		| BETWEEN ]
		<qualifier-value>  = A valid port number.
		example = CONNECTION.DSTPORT.EQ(80)
		<qualifier> = PORT
		<qualifier-method> = [ EQ | NE | GT | GE | LT | LE
		| BETWEEN ]
		<qualifier-value>  = A valid port number.
		example = CONNECTION.PORT.EQ(80)
		<qualifier> = VLANID
		<qualifier-method> = [ EQ | NE | GT | GE | LT | LE
		| BETWEEN ]
		<qualifier-value>  = A valid VLAN ID.
		example = CONNECTION.VLANID.EQ(0)
		<qualifier> = CONNID
		<qualifier-method> = [ EQ | NE | GT | GE | LT | LE
		| BETWEEN ]
		<qualifier-value>  = A valid PCB dev number.
		example = CONNECTION.CONNID.EQ(0)
		<qualifier> = PPEID
		<qualifier-method> = [ EQ | NE | GT | GE | LT | LE
		| BETWEEN ]
		<qualifier-value>  = A valid core ID.
		example = CONNECTION.PPEID.EQ(0)
		<qualifier> = SVCNAME
		<qualifier-method> = [ EQ | NE | CONTAINS | STARTSWITH
		| ENDSWITH ]
		<qualifier-value>  = A valid text string.
		example = CONNECTION.SVCNAME.EQ("name")
		<qualifier> = LB_VSERVER.NAME
		<qualifier-method> = [ EQ | NE | CONTAINS | STARTSWITH
		| ENDSWITH ]
		<qualifier-value>  = LB vserver name.
		example = CONNECTION.LB_VSERVER.NAME.EQ("name")
		<qualifier> = CS_VSERVER.NAME
		<qualifier-method> = [ EQ | NE | CONTAINS | STARTSWITH
		| ENDSWITH ]
		<qualifier-value>  = CS vserver name.
		example = CONNECTION.CS_VSERVER.NAME.EQ("name")
		<qualifier> = INTF
		<qualifier-method> = [ EQ | NE ]
		<qualifier-value>  =  A valid interface id in the
		form of x/y.
		example = CONNECTION.INTF.EQ("x/y")
		<qualifier> = SERVICE_TYPE
		<qualifier-method> = [ EQ | NE ]
		<qualifier-value>  = ( SVC_HTTP | FTP | TCP | UDP | SSL |
		SSL_BRIDGE | SSL_TCP | NNTP | RPCSVR | RPCSVRS |
		RPCCLNT | SVC_DNS | ADNS | SNMP | RTSP | DHCPRA | ANY|
		MONITOR | MONITOR_UDP | MONITOR_PING | SIP_UDP |
		SVC_MYSQL | SVC_MSSQL | FIX | SSL_FIX | PKTSTEER |
		SVC_AAA | SERVICE_UNKNOWN )
		example = CONNECTION.SERVICE_TYPE.EQ(ANY)
		<qualifier> = TRAFFIC_DOMAIN_ID
		<qualifier-method> = [ EQ | NE | GT | GE | LT | LE
		| BETWEEN ]
		<qualifier-value>  = A valid traffic domain ID.
		example = CONNECTION.TRAFFIC_DOMAIN_ID.EQ(0)
		eg: start nstrace -filter "CONNECTION.SRCIP.EQ(127.0.0.1) || (CONNECTION.SVCNAME.NE("s1") && CONNECTION.SRCPORT.EQ(80))"
		The filter expression should be given in double quotes.
		common use cases:
		Trace capturing full sized traffic from/to ip 10.102.44.111, excluding loopback traffic
		start nstrace -size 0 -filter "CONNECTION.IP.NE(127.0.0.1) && CONNECTION.IP.EQ(10.102.44.111)"
		Trace capturing all traffic to (terminating at) port 80 or 443
		start nstrace -size 0 -filter "CONNECTION.DSTPORT.EQ(443) || CONNECTION.DSTPORT.EQ(80)"
		Trace capturing all backend traffic specific to service service1 along with corresponding client side traffic
		start nstrace -size 0 -filter "CONNECTION.SVCNAME.EQ("service1")" -link ENABLED
		Trace capturing all traffic through NetScaler interface 1/1
		start nstrace -filter "CONNECTION.INTF.EQ("1/1")"
		Trace capturing all traffic specific through vlan 2
		start nstrace -filter "CONNECTION.VLANID.EQ(2)"
		Trace capturing all frontend (client side) traffic specific to lb vserver vserver1 along with corresponding server side traffic
		start nstrace -size 0 -filter "CONNECTION.LB_VSERVER.NAME.EQ("vserver1")" -link ENABLED 
	*/
	Filter string `json:"filter,omitempty"`
	/**
	* Includes filtered connection's peer traffic.
	*/
	Link string `json:"link,omitempty"`
	/**
	* Nodes on which tracing is started.
	*/
	Nodes []int `json:"nodes,omitempty"`
	/**
	* File size, in MB, treshold for rollover. If free disk space is less than 2GB at the time of rollover, trace will stop
	*/
	Filesize *int `json:"filesize,omitempty"`
	/**
	* Format in which trace will be generated
	*/
	Traceformat string `json:"traceformat,omitempty"`
	/**
	* Specify how traces across PE's are merged
	*/
	Merge string `json:"merge,omitempty"`
	/**
	* Enable or disable runtime temp file cleanup
	*/
	Doruntimecleanup string `json:"doruntimecleanup,omitempty"`
	/**
	* Number of 16KB trace buffers
	*/
	Tracebuffers *int `json:"tracebuffers,omitempty"`
	/**
	* skip RPC packets
	*/
	Skiprpc string `json:"skiprpc,omitempty"`
	/**
	* skip local SSH packets
	*/
	Skiplocalssh string `json:"skiplocalssh,omitempty"`
	/**
	* Capture SSL Master keys. Master keys will not be captured on FIPS machine.
		Warning: The captured keys can be used to decrypt information that may be confidential. The captured key files have to be stored in a secure environment
	*/
	Capsslkeys string `json:"capsslkeys,omitempty"`
	/**
	* Captures Dropped Packets if set to ENABLED.
	*/
	Capdroppkt string `json:"capdroppkt,omitempty"`
	/**
	* Logs packets in appliance's memory and dumps the trace file on stopping the nstrace operation
	*/
	Inmemorytrace string `json:"inmemorytrace,omitempty"`
	/**
	* Unique number that identifies the cluster node.
	*/
	Nodeid *int `json:"nodeid,omitempty"`

	//------- Read only Parameter ---------;

	State string `json:"state,omitempty"`
	Scope string `json:"scope,omitempty"`
	Tracelocation string `json:"tracelocation,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
