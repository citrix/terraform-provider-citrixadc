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
* Configuration for TCP/IP connection table resource.
*/
type Nsconnectiontable struct {
	/**
	* The maximum length of filter expression is 255 and it can be of following format:
		<expression> [<relop> <expression>]
		<relop> = ( && | || )
		<expression> =:
		CONNECTION.<qualifier>.<qualifier-method>.(<qualifier-value>)
		<qualifier> = SRCIP
		<qualifier-method> = [ EQ | NE ]
		<qualifier-value>  = A valid IPv4 address
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
		<qualifier> = SVCNAME
		<qualifier-method> = [ EQ | NE | CONTAINS | STARTSWITH
		| ENDSWITH ]
		<qualifier-value>  = service name.
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
		<qualifier-value>  = A valid interface id in the form of
		x/y (n/x/y in case of cluster interface).
		examle = CONNECTION.INTF.EQ("0/1/1")
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
		<qualifier> = IDLETIME
		<qualifier-method> = [ EQ | NE | GT | GE | LT | LE
		| BETWEEN ]
		<qualifier-value>  = A positive integer indicating the
		idletime.
		example = CONNECTION.IDLETIME.LT(100)
		<qualifier> = TCPSTATE
		<qualifier-method> = [ EQ | NE ]
		<qualifier-value>  = ( CLOSE_WAIT | CLOSED | CLOSING |
		ESTABLISHED | FIN_WAIT_1 | FIN_WAIT_2 | LAST_ACK |
		LISTEN | SYN_RECEIVED | SYN_SENT | TIME_WAIT |
		NOT_APPLICABLE)
		example = CONNECTION.TCPSTATE.EQ(LISTEN)
		<qualifier> = SERVICE_TYPE
		<qualifier-method> = [ EQ | NE ]
		<qualifier-value>  = ( SVC_HTTP | FTP | TCP | UDP | SSL |
		SSL_BRIDGE | SSL_TCP | NNTP | RPCSVR | RPCSVRS |
		RPCCLNT | SVC_DNS | ADNS | SNMP | RTSP | DHCPRA | NAT | ANY |
		MONITOR | MONITOR_UDP | MONITOR_PING | SIP_UDP |
		SVC_MYSQL | SVC_MSSQL | SERVICE_UNKNOWN )
		example = CONNECTION.SERVICE_TYPE.EQ(ANY)
		<qualifier> = TRAFFIC_DOMAIN_ID
		<qualifier-method> = [ EQ | NE | GT | GE | LT | LE
		| BETWEEN ]
		<qualifier-value>  = A valid traffic domain ID.
		example = CONNECTION.TRAFFIC_DOMAIN_ID.EQ(0)
		common usecases:
		Filtering out loopback connections and view present
		connections through netscaler
		show connectiontable "CONNECTION.IP.NE(127.0.0.1) &&
		CONNECTION.TCPSTATE.EQ(ESTABLISHED)" -detail full
		show connections from a particular sourceip and targeted
		to port 80
		show connectiontable "CONNECTION.SRCIP.EQ(10.102.1.91) &&
		CONNECTION.DSTPORT.EQ(80)"
		show connection particular to a service and its linked
		client connections
		show connectiontable CONNECTION.SVCNAME.EQ("S1")
		-detail link
		show connections for a particular servicetype(e.g.http)
		show connectiontable CONNECTION.SERVICE_TYPE.EQ(TCP)
		viewing connections that have been idle for a long time
		show connectiontable CONNECTION.IDLETIME.GT(100)
		show connections particular to a service and idle
		for a long time
		show connectiontable "CONNECTION.SVCNAME.EQ(\\"S1\\") &&
		CONNECTION.IDLETIME.GT(100)"
		show connections for a particular interface
		show connectiontable CONNECTION.INTF.EQ("1/1")
		show connections for a particular interface and vlan
		show connectiontable "CONNECTION.INTF.EQ(\\"1/1\\") &&
		CONNECTION.VLANID.EQ(1)"
	*/
	Filterexpression string `json:"filterexpression,omitempty"`
	/**
	* Display link information if available
	*/
	Link bool `json:"link,omitempty"`
	/**
	* Display name instead of IP for local entities
	*/
	Filtername bool `json:"filtername,omitempty"`
	/**
	* Specify display options for the connection table.
		* LINK - Displays the linked PCB (Protocol Control Block).
		* NAME - Displays along with the service name.
		* CONNFAILOVER - Displays PCB with connection failover.
		* FULL - Displays all available details.
	*/
	Detail []string `json:"detail,omitempty"`
	/**
	* Display listening services only
	*/
	Listen bool `json:"listen,omitempty"`
	/**
	* Unique number that identifies the cluster node.
	*/
	Nodeid int `json:"nodeid,omitempty"`

	//------- Read only Parameter ---------;

	Sourceip string `json:"sourceip,omitempty"`
	Sourceport string `json:"sourceport,omitempty"`
	Destip string `json:"destip,omitempty"`
	Destport string `json:"destport,omitempty"`
	Svctype string `json:"svctype,omitempty"`
	Idletime string `json:"idletime,omitempty"`
	State string `json:"state,omitempty"`
	Linksourceip string `json:"linksourceip,omitempty"`
	Linksourceport string `json:"linksourceport,omitempty"`
	Linkdestip string `json:"linkdestip,omitempty"`
	Linkdestport string `json:"linkdestport,omitempty"`
	Linkservicetype string `json:"linkservicetype,omitempty"`
	Linkidletime string `json:"linkidletime,omitempty"`
	Linkstate string `json:"linkstate,omitempty"`
	Entityname string `json:"entityname,omitempty"`
	Linkentityname string `json:"linkentityname,omitempty"`
	Connid string `json:"connid,omitempty"`
	Linkconnid string `json:"linkconnid,omitempty"`
	Connproperties string `json:"connproperties,omitempty"`
	Optionflags string `json:"optionflags,omitempty"`
	Nswsvalue string `json:"nswsvalue,omitempty"`
	Peerwsvalue string `json:"peerwsvalue,omitempty"`
	Mss string `json:"mss,omitempty"`
	Retxretrycnt string `json:"retxretrycnt,omitempty"`
	Rcvwnd string `json:"rcvwnd,omitempty"`
	Advwnd string `json:"advwnd,omitempty"`
	Sndcwnd string `json:"sndcwnd,omitempty"`
	Iss string `json:"iss,omitempty"`
	Irs string `json:"irs,omitempty"`
	Rcvnxt string `json:"rcvnxt,omitempty"`
	Maxack string `json:"maxack,omitempty"`
	Sndnxt string `json:"sndnxt,omitempty"`
	Sndunack string `json:"sndunack,omitempty"`
	Httpendseq string `json:"httpendseq,omitempty"`
	Httpstate string `json:"httpstate,omitempty"`
	Trcount string `json:"trcount,omitempty"`
	Priority string `json:"priority,omitempty"`
	Httpreqver string `json:"httpreqver,omitempty"`
	Httprequest string `json:"httprequest,omitempty"`
	Httprspcode string `json:"httprspcode,omitempty"`
	Rttsmoothed string `json:"rttsmoothed,omitempty"`
	Rttvariance string `json:"rttvariance,omitempty"`
	Outoforderpkts string `json:"outoforderpkts,omitempty"`
	Linkoptionflag string `json:"linkoptionflag,omitempty"`
	Linknswsvalue string `json:"linknswsvalue,omitempty"`
	Linkpeerwsvalue string `json:"linkpeerwsvalue,omitempty"`
	Targetnodeidnnm string `json:"targetnodeidnnm,omitempty"`
	Sourcenodeidnnm string `json:"sourcenodeidnnm,omitempty"`
	Channelidnnm string `json:"channelidnnm,omitempty"`
	Msgversionnnm string `json:"msgversionnnm,omitempty"`
	Td string `json:"td,omitempty"`
	Maxrcvbuf string `json:"maxrcvbuf,omitempty"`
	Linkmaxrcvbuf string `json:"linkmaxrcvbuf,omitempty"`
	Rxqsize string `json:"rxqsize,omitempty"`
	Linkrxqsize string `json:"linkrxqsize,omitempty"`
	Maxsndbuf string `json:"maxsndbuf,omitempty"`
	Linkmaxsndbuf string `json:"linkmaxsndbuf,omitempty"`
	Txqsize string `json:"txqsize,omitempty"`
	Linktxqsize string `json:"linktxqsize,omitempty"`
	Flavor string `json:"flavor,omitempty"`
	Linkflavor string `json:"linkflavor,omitempty"`
	Bwestimate string `json:"bwestimate,omitempty"`
	Linkbwestimate string `json:"linkbwestimate,omitempty"`
	Rttmin string `json:"rttmin,omitempty"`
	Linkrttmin string `json:"linkrttmin,omitempty"`
	Name string `json:"name,omitempty"`
	Linkname string `json:"linkname,omitempty"`
	Tcpmode string `json:"tcpmode,omitempty"`
	Linktcpmode string `json:"linktcpmode,omitempty"`
	Realtimertt string `json:"realtimertt,omitempty"`
	Linkrealtimertt string `json:"linkrealtimertt,omitempty"`
	Sndbuf string `json:"sndbuf,omitempty"`
	Linksndbuf string `json:"linksndbuf,omitempty"`
	Nsbtcpwaitq string `json:"nsbtcpwaitq,omitempty"`
	Linknsbtcpwaitq string `json:"linknsbtcpwaitq,omitempty"`
	Nsbretxq string `json:"nsbretxq,omitempty"`
	Linknsbretxq string `json:"linknsbretxq,omitempty"`
	Sackblocks string `json:"sackblocks,omitempty"`
	Linksackblocks string `json:"linksackblocks,omitempty"`
	Congstate string `json:"congstate,omitempty"`
	Linkcongstate string `json:"linkcongstate,omitempty"`
	Sndrecoverle string `json:"sndrecoverle,omitempty"`
	Linksndrecoverle string `json:"linksndrecoverle,omitempty"`
	Creditsinbytes string `json:"creditsinbytes,omitempty"`
	Linkcredits string `json:"linkcredits,omitempty"`
	Rateinbytes string `json:"rateinbytes,omitempty"`
	Linkrateinbytes string `json:"linkrateinbytes,omitempty"`
	Rateschedulerqueue string `json:"rateschedulerqueue,omitempty"`
	Linkrateschedulerqueue string `json:"linkrateschedulerqueue,omitempty"`
	Burstratecontrol string `json:"burstratecontrol,omitempty"`
	Linkburstratecontrol string `json:"linkburstratecontrol,omitempty"`
	Cqabifavg string `json:"cqabifavg,omitempty"`
	Cqathruputavg string `json:"cqathruputavg,omitempty"`
	Cqarcvwndavg string `json:"cqarcvwndavg,omitempty"`
	Cqaiai1mspct string `json:"cqaiai1mspct,omitempty"`
	Cqaiai2mspct string `json:"cqaiai2mspct,omitempty"`
	Cqasamples string `json:"cqasamples,omitempty"`
	Cqaiaisamples string `json:"cqaiaisamples,omitempty"`
	Cqanetclass string `json:"cqanetclass,omitempty"`
	Cqaccl string `json:"cqaccl,omitempty"`
	Cqacsq string `json:"cqacsq,omitempty"`
	Cqaiaiavg string `json:"cqaiaiavg,omitempty"`
	Cqaisiavg string `json:"cqaisiavg,omitempty"`
	Cqarcvwndmin string `json:"cqarcvwndmin,omitempty"`
	Cqaretxcorr string `json:"cqaretxcorr,omitempty"`
	Cqaretxcong string `json:"cqaretxcong,omitempty"`
	Cqaretxpackets string `json:"cqaretxpackets,omitempty"`
	Cqaloaddelayavg string `json:"cqaloaddelayavg,omitempty"`
	Cqanoisedelayavg string `json:"cqanoisedelayavg,omitempty"`
	Cqarttmax string `json:"cqarttmax,omitempty"`
	Cqarttmin string `json:"cqarttmin,omitempty"`
	Cqarttavg string `json:"cqarttavg,omitempty"`
	Adaptivetcpprofname string `json:"adaptivetcpprofname,omitempty"`
	Outoforderblocks string `json:"outoforderblocks,omitempty"`
	Outoforderflushedcount string `json:"outoforderflushedcount,omitempty"`
	Outoforderbytes string `json:"outoforderbytes,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
