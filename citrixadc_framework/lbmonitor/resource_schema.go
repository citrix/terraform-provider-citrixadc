package lbmonitor

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/lb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LbmonitorResourceModel describes the resource data model.
type LbmonitorResourceModel struct {
	Id                               types.String `tfsdk:"id"`
	Snmpoid                          types.String `tfsdk:"snmpoid"`
	Acctapplicationid                types.List   `tfsdk:"acctapplicationid"`
	Action                           types.String `tfsdk:"action"`
	Alertretries                     types.Int64  `tfsdk:"alertretries"`
	Application                      types.String `tfsdk:"application"`
	Attribute                        types.String `tfsdk:"attribute"`
	Authapplicationid                types.List   `tfsdk:"authapplicationid"`
	Basedn                           types.String `tfsdk:"basedn"`
	Binddn                           types.String `tfsdk:"binddn"`
	Customheaders                    types.String `tfsdk:"customheaders"`
	Database                         types.String `tfsdk:"database"`
	Destip                           types.String `tfsdk:"destip"`
	Destport                         types.Int64  `tfsdk:"destport"`
	Deviation                        types.Int64  `tfsdk:"deviation"`
	Dispatcherip                     types.String `tfsdk:"dispatcherip"`
	Dispatcherport                   types.Int64  `tfsdk:"dispatcherport"`
	Domain                           types.String `tfsdk:"domain"`
	Downtime                         types.Int64  `tfsdk:"downtime"`
	Evalrule                         types.String `tfsdk:"evalrule"`
	Failureretries                   types.Int64  `tfsdk:"failureretries"`
	Filename                         types.String `tfsdk:"filename"`
	Filter                           types.String `tfsdk:"filter"`
	Firmwarerevision                 types.Int64  `tfsdk:"firmwarerevision"`
	Group                            types.String `tfsdk:"group"`
	Grpchealthcheck                  types.String `tfsdk:"grpchealthcheck"`
	Grpcservicename                  types.String `tfsdk:"grpcservicename"`
	Grpcstatuscode                   types.List   `tfsdk:"grpcstatuscode"`
	Hostipaddress                    types.String `tfsdk:"hostipaddress"`
	Hostname                         types.String `tfsdk:"hostname"`
	Httprequest                      types.String `tfsdk:"httprequest"`
	Inbandsecurityid                 types.String `tfsdk:"inbandsecurityid"`
	Interval                         types.Int64  `tfsdk:"interval"`
	Ipaddress                        types.List   `tfsdk:"ipaddress"`
	Iptunnel                         types.String `tfsdk:"iptunnel"`
	Kcdaccount                       types.String `tfsdk:"kcdaccount"`
	Lasversion                       types.String `tfsdk:"lasversion"`
	Logonpointname                   types.String `tfsdk:"logonpointname"`
	Lrtm                             types.String `tfsdk:"lrtm"`
	Maxforwards                      types.Int64  `tfsdk:"maxforwards"`
	Metric                           types.String `tfsdk:"metric"`
	Metrictable                      types.String `tfsdk:"metrictable"`
	Metricthreshold                  types.Int64  `tfsdk:"metricthreshold"`
	Metricweight                     types.Int64  `tfsdk:"metricweight"`
	Monitorname                      types.String `tfsdk:"monitorname"`
	Mqttclientidentifier             types.String `tfsdk:"mqttclientidentifier"`
	Mqttversion                      types.Int64  `tfsdk:"mqttversion"`
	Mssqlprotocolversion             types.String `tfsdk:"mssqlprotocolversion"`
	Netprofile                       types.String `tfsdk:"netprofile"`
	Oraclesid                        types.String `tfsdk:"oraclesid"`
	Originhost                       types.String `tfsdk:"originhost"`
	Originrealm                      types.String `tfsdk:"originrealm"`
	Password                         types.String `tfsdk:"password"`
	Productname                      types.String `tfsdk:"productname"`
	Query                            types.String `tfsdk:"query"`
	Querytype                        types.String `tfsdk:"querytype"`
	Radaccountsession                types.String `tfsdk:"radaccountsession"`
	Radaccounttype                   types.Int64  `tfsdk:"radaccounttype"`
	Radapn                           types.String `tfsdk:"radapn"`
	Radframedip                      types.String `tfsdk:"radframedip"`
	Radkey                           types.String `tfsdk:"radkey"`
	Radmsisdn                        types.String `tfsdk:"radmsisdn"`
	Radnasid                         types.String `tfsdk:"radnasid"`
	Radnasip                         types.String `tfsdk:"radnasip"`
	Recv                             types.String `tfsdk:"recv"`
	Respcode                         types.List   `tfsdk:"respcode"`
	Resptimeout                      types.Int64  `tfsdk:"resptimeout"`
	Resptimeoutthresh                types.Int64  `tfsdk:"resptimeoutthresh"`
	Retries                          types.Int64  `tfsdk:"retries"`
	Reverse                          types.String `tfsdk:"reverse"`
	Rtsprequest                      types.String `tfsdk:"rtsprequest"`
	Scriptargs                       types.String `tfsdk:"scriptargs"`
	Scriptname                       types.String `tfsdk:"scriptname"`
	Secondarypassword                types.String `tfsdk:"secondarypassword"`
	Secure                           types.String `tfsdk:"secure"`
	Secureargs                       types.String `tfsdk:"secureargs"`
	Send                             types.String `tfsdk:"send"`
	Servicegroupname                 types.String `tfsdk:"servicegroupname"`
	Servicename                      types.String `tfsdk:"servicename"`
	Sipmethod                        types.String `tfsdk:"sipmethod"`
	Sipreguri                        types.String `tfsdk:"sipreguri"`
	Sipuri                           types.String `tfsdk:"sipuri"`
	Sitepath                         types.String `tfsdk:"sitepath"`
	Snmpcommunity                    types.String `tfsdk:"snmpcommunity"`
	Snmpthreshold                    types.String `tfsdk:"snmpthreshold"`
	Snmpversion                      types.String `tfsdk:"snmpversion"`
	Sqlquery                         types.String `tfsdk:"sqlquery"`
	Sslprofile                       types.String `tfsdk:"sslprofile"`
	State                            types.String `tfsdk:"state"`
	Storedb                          types.String `tfsdk:"storedb"`
	Storefrontacctservice            types.String `tfsdk:"storefrontacctservice"`
	Storefrontcheckbackendservices   types.String `tfsdk:"storefrontcheckbackendservices"`
	Storename                        types.String `tfsdk:"storename"`
	Successretries                   types.Int64  `tfsdk:"successretries"`
	Supportedvendorids               types.List   `tfsdk:"supportedvendorids"`
	Tos                              types.String `tfsdk:"tos"`
	Tosid                            types.Int64  `tfsdk:"tosid"`
	Transparent                      types.String `tfsdk:"transparent"`
	Trofscode                        types.Int64  `tfsdk:"trofscode"`
	Trofsstring                      types.String `tfsdk:"trofsstring"`
	Type                             types.String `tfsdk:"type"`
	Units1                           types.String `tfsdk:"units1"`
	Units2                           types.String `tfsdk:"units2"`
	Units3                           types.String `tfsdk:"units3"`
	Units4                           types.String `tfsdk:"units4"`
	Username                         types.String `tfsdk:"username"`
	Validatecred                     types.String `tfsdk:"validatecred"`
	Vendorid                         types.Int64  `tfsdk:"vendorid"`
	Vendorspecificacctapplicationids types.List   `tfsdk:"vendorspecificacctapplicationids"`
	Vendorspecificauthapplicationids types.List   `tfsdk:"vendorspecificauthapplicationids"`
	Vendorspecificvendorid           types.Int64  `tfsdk:"vendorspecificvendorid"`
}

func (r *LbmonitorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lbmonitor resource.",
			},
			"snmpoid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SNMP OID for SNMP monitors.",
			},
			"acctapplicationid": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "List of Acct-Application-Id attribute value pairs (AVPs) for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. A maximum of eight of these AVPs are supported in a monitoring message.",
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DOWN"),
				Description: "Action to perform when the response to an inline monitor (a monitor of type HTTP-INLINE) indicates that the service is down. A service monitored by an inline monitor is considered DOWN if the response code is not one of the codes that have been specified for the Response Code parameter.\nAvailable settings function as follows:\n* NONE - Do not take any action. However, the show service command and the show lb monitor command indicate the total number of responses that were checked and the number of consecutive error responses received after the last successful probe.\n* LOG - Log the event in NSLOG or SYSLOG.\n* DOWN - Mark the service as being down, and then do not direct any traffic to the service until the configured down time has expired. Persistent connections to the service are terminated as soon as the service is marked as DOWN. Also, log the event in NSLOG or SYSLOG.",
			},
			"alertretries": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of consecutive probe failures after which the appliance generates an SNMP trap called monProbeFailed.",
			},
			"application": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the application used to determine the state of the service. Applicable to monitors of type CITRIX-XML-SERVICE.",
			},
			"attribute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Attribute to evaluate when the LDAP server responds to the query. Success or failure of the monitoring probe depends on whether the attribute exists in the response. Optional.",
			},
			"authapplicationid": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "List of Auth-Application-Id attribute value pairs (AVPs) for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. A maximum of eight of these AVPs are supported in a monitoring CER message.",
			},
			"basedn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The base distinguished name of the LDAP service, from where the LDAP server can begin the search for the attributes in the monitoring query. Required for LDAP service monitoring.",
			},
			"binddn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The distinguished name with which an LDAP monitor can perform the Bind operation on the LDAP server. Optional. Applicable to LDAP monitors.",
			},
			"customheaders": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom header string to include in the monitoring probes.",
			},
			"database": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the database to connect to during authentication.",
			},
			"destip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the service to which to send probes. If the parameter is set to 0, the IP address of the server to which the monitor is bound is considered the destination IP address.",
			},
			"destport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "TCP or UDP port to which to send the probe. If the parameter is set to 0, the port number of the service to which the monitor is bound is considered the destination port. For a monitor of type USER, however, the destination port is the port number that is included in the HTTP request sent to the dispatcher. Does not apply to monitors of type PING.",
			},
			"deviation": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time value added to the learned average response time in dynamic response time monitoring (DRTM). When a deviation is specified, the appliance learns the average response time of bound services and adds the deviation to the average. The final value is then continually adjusted to accommodate response time variations over time. Specified in milliseconds, seconds, or minutes.",
			},
			"dispatcherip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the dispatcher to which to send the probe.",
			},
			"dispatcherport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number on which the dispatcher listens for the monitoring probe.",
			},
			"domain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Domain in which the XenDesktop Desktop Delivery Controller (DDC) servers or Web Interface servers are present. Required by CITRIX-XD-DDC and CITRIX-WI-EXTENDED monitors for logging on to the DDC servers and Web Interface servers, respectively.",
			},
			"downtime": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(30),
				Description: "Time duration for which to wait before probing a service that has been marked as DOWN. Expressed in milliseconds, seconds, or minutes.",
			},
			"evalrule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that evaluates the database server's response to a MYSQL-ECV or MSSQL-ECV monitoring query. Must produce a Boolean result. The result determines the state of the server. If the expression returns TRUE, the probe succeeds.\nFor example, if you want the appliance to evaluate the error message to determine the state of the server, use the rule MYSQL.RES.ROW(10) .TEXT_ELEM(2).EQ(\"MySQL\").",
			},
			"failureretries": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of retries that must fail, out of the number specified for the Retries parameter, for a service to be marked as DOWN. For example, if the Retries parameter is set to 10 and the Failure Retries parameter is set to 6, out of the ten probes sent, at least six probes must fail if the service is to be marked as DOWN. The default value of 0 means that all the retries must fail if the service is to be marked as DOWN.",
			},
			"filename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of a file on the FTP server. The appliance monitors the FTP service by periodically checking the existence of the file on the server. Applicable to FTP-EXTENDED monitors.",
			},
			"filter": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Filter criteria for the LDAP query. Optional.",
			},
			"firmwarerevision": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Firmware-Revision value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.",
			},
			"group": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of a newsgroup available on the NNTP service that is to be monitored. The appliance periodically generates an NNTP query for the name of the newsgroup and evaluates the response. If the newsgroup is found on the server, the service is marked as UP. If the newsgroup does not exist or if the search fails, the service is marked as DOWN. Applicable to NNTP monitors.",
			},
			"grpchealthcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to enable or disable gRPC health check service.",
			},
			"grpcservicename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to specify gRPC service name on which gRPC health check need to be performed",
			},
			"grpcstatuscode": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "gRPC status codes for which to mark the service as UP. The default value is 12(health check unimplemented). If the gRPC status code 0 is received from the backend this configuration is ignored.",
			},
			"hostipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Host-IP-Address value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. If Host-IP-Address is not specified, the appliance inserts the mapped IP (MIP) address or subnet IP (SNIP) address from which the CER request (the monitoring probe) is sent.",
			},
			"hostname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Hostname in the FQDN format (Example: porche.cars.org). Applicable to STOREFRONT monitors.",
			},
			"httprequest": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "HTTP request to send to the server (for example, \"HEAD /file.html\").",
			},
			"inbandsecurityid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Inband-Security-Id for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.",
			},
			"interval": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(5),
				Description: "Time interval between two successive probes. Must be greater than the value of Response Time-out.",
			},
			"ipaddress": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Set of IP addresses expected in the monitoring response from the DNS server, if the record type is A or AAAA. Applicable to DNS monitors.",
			},
			"iptunnel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send the monitoring probe to the service through an IP tunnel. A destination IP address must be specified.",
			},
			"kcdaccount": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "KCD Account used by MSSQL monitor",
			},
			"lasversion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Version number of the Citrix Advanced Access Control Logon Agent. Required by the CITRIX-AAC-LAS monitor.",
			},
			"logonpointname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the logon point that is configured for the Citrix Access Gateway Advanced Access Control software. Required if you want to monitor the associated login page or Logon Agent. Applicable to CITRIX-AAC-LAS and CITRIX-AAC-LOGINPAGE monitors.",
			},
			"lrtm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Calculate the least response times for bound services. If this parameter is not enabled, the appliance does not learn the response times of the bound services. Also used for LRTM load balancing.",
			},
			"maxforwards": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Maximum number of hops that the SIP request used for monitoring can traverse to reach the server. Applicable only to monitors of type SIP-UDP.",
			},
			"metric": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Metric name in the metric table, whose setting is changed. A value zero disables the metric and it will not be used for load calculation",
			},
			"metrictable": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Metric table to which to bind metrics.",
			},
			"metricthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Threshold to be used for that metric.",
			},
			"metricweight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The weight for the specified service metric with respect to others.",
			},
			"monitorname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the monitor. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\n\nCLI Users:  If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my monitor\" or 'my monitor').",
			},
			"mqttclientidentifier": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Client id to be used in Connect command",
			},
			"mqttversion": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(4),
				Description: "Version of MQTT protocol used in connect message, default is version 3.1.1 [4]",
			},
			"mssqlprotocolversion": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("70"),
				Description: "Version of MSSQL server that is to be monitored.",
			},
			"netprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the network profile.",
			},
			"oraclesid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the service identifier that is used to connect to the Oracle database during authentication.",
			},
			"originhost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Origin-Host value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.",
			},
			"originrealm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Origin-Realm value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password that is required for logging on to the RADIUS, NNTP, FTP, FTP-EXTENDED, MYSQL, MSSQL, POP3, CITRIX-AG, CITRIX-XD-DDC, CITRIX-WI-EXTENDED, CITRIX-XNC-ECV or CITRIX-XDM server. Used in conjunction with the user name specified for the User Name parameter.",
			},
			"productname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Product-Name value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.",
			},
			"query": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Domain name to resolve as part of monitoring the DNS service (for example, example.com).",
			},
			"querytype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of DNS record for which to send monitoring queries. Set to Address for querying A records, AAAA for querying AAAA records, and Zone for querying the SOA record.",
			},
			"radaccountsession": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Account Session ID to be used in Account Request Packet. Applicable to monitors of type RADIUS_ACCOUNTING.",
			},
			"radaccounttype": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Account Type to be used in Account Request Packet. Applicable to monitors of type RADIUS_ACCOUNTING.",
			},
			"radapn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Called Station Id to be used in Account Request Packet. Applicable to monitors of type RADIUS_ACCOUNTING.",
			},
			"radframedip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Source ip with which the packet will go out . Applicable to monitors of type RADIUS_ACCOUNTING.",
			},
			"radkey": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Authentication key (shared secret text string) for RADIUS clients and servers to exchange. Applicable to monitors of type RADIUS and RADIUS_ACCOUNTING.",
			},
			"radmsisdn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Calling Stations Id to be used in Account Request Packet. Applicable to monitors of type RADIUS_ACCOUNTING.",
			},
			"radnasid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "NAS-Identifier to send in the Access-Request packet. Applicable to monitors of type RADIUS.",
			},
			"radnasip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Network Access Server (NAS) IP address to use as the source IP address when monitoring a RADIUS server. Applicable to monitors of type RADIUS and RADIUS_ACCOUNTING.",
			},
			"recv": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String expected from the server for the service to be marked as UP. Applicable to TCP-ECV, HTTP-ECV, and UDP-ECV monitors.",
			},
			"respcode": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Response codes for which to mark the service as UP. For any other response code, the action performed depends on the monitor type. HTTP monitors and RADIUS monitors mark the service as DOWN, while HTTP-INLINE monitors perform the action indicated by the Action parameter.",
			},
			"resptimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(2),
				Description: "Amount of time for which the appliance must wait before it marks a probe as FAILED.  Must be less than the value specified for the Interval parameter.\n\nNote: For UDP-ECV monitors for which a receive string is not configured, response timeout does not apply. For UDP-ECV monitors with no receive string, probe failure is indicated by an ICMP port unreachable error received from the service.",
			},
			"resptimeoutthresh": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Response time threshold, specified as a percentage of the Response Time-out parameter. If the response to a monitor probe has not arrived when the threshold is reached, the appliance generates an SNMP trap called monRespTimeoutAboveThresh. After the response time returns to a value below the threshold, the appliance generates a monRespTimeoutBelowThresh SNMP trap. For the traps to be generated, the \"MONITOR-RTO-THRESHOLD\" alarm must also be enabled.",
			},
			"retries": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(3),
				Description: "Maximum number of probes to send to establish the state of a service for which a monitoring probe failed.",
			},
			"reverse": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Mark a service as DOWN, instead of UP, when probe criteria are satisfied, and as UP instead of DOWN when probe criteria are not satisfied.",
			},
			"rtsprequest": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RTSP request to send to the server (for example, \"OPTIONS *\").",
			},
			"scriptargs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String of arguments for the script. The string is copied verbatim into the request.",
			},
			"scriptname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Path and name of the script to execute. The script must be available on the Citrix ADC, in the /nsconfig/monitors/ directory.",
			},
			"secondarypassword": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Secondary password that users might have to provide to log on to the Access Gateway server. Applicable to CITRIX-AG monitors.",
			},
			"secure": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use a secure SSL connection when monitoring a service. Applicable only to TCP based monitors. The secure option cannot be used with a CITRIX-AG monitor, because a CITRIX-AG monitor uses a secure connection by default.",
			},
			"secureargs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "List of arguments for the script which should be secure",
			},
			"send": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String to send to the service. Applicable to TCP-ECV, HTTP-ECV, and UDP-ECV monitors.",
			},
			"servicegroupname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the service group to which the monitor is to be bound.",
			},
			"servicename": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the service to which the monitor is bound.",
			},
			"sipmethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SIP method to use for the query. Applicable only to monitors of type SIP-UDP.",
			},
			"sipreguri": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SIP user to be registered. Applicable only if the monitor is of type SIP-UDP and the SIP Method parameter is set to REGISTER.",
			},
			"sipuri": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SIP URI string to send to the service (for example, sip:sip.test). Applicable only to monitors of type SIP-UDP.",
			},
			"sitepath": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL of the logon page. For monitors of type CITRIX-WEB-INTERFACE, to monitor a dynamic page under the site path, terminate the site path with a slash (/). Applicable to CITRIX-WEB-INTERFACE, CITRIX-WI-EXTENDED and CITRIX-XDM monitors.",
			},
			"snmpcommunity": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Community name for SNMP monitors.",
			},
			"snmpthreshold": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Threshold for SNMP monitors.",
			},
			"snmpversion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SNMP version to be used for SNMP monitors.",
			},
			"sqlquery": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SQL query for a MYSQL-ECV or MSSQL-ECV monitor. Sent to the database server after the server authenticates the connection.",
			},
			"sslprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SSL Profile associated with the monitor",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "State of the monitor. The DISABLED setting disables not only the monitor being configured, but all monitors of the same type, until the parameter is set to ENABLED. If the monitor is bound to a service, the state of the monitor is not taken into account when the state of the service is determined.",
			},
			"storedb": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Store the database list populated with the responses to monitor probes. Used in database specific load balancing if MSSQL-ECV/MYSQL-ECV  monitor is configured.",
			},
			"storefrontacctservice": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Enable/Disable probing for Account Service. Applicable only to Store Front monitors. For multi-tenancy configuration users my skip account service",
			},
			"storefrontcheckbackendservices": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option will enable monitoring of services running on storefront server. Storefront services are monitored by probing to a Windows service that runs on the Storefront server and exposes details of which storefront services are running.",
			},
			"storename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Store Name. For monitors of type STOREFRONT, STORENAME is an optional argument defining storefront service store name. Applicable to STOREFRONT monitors.",
			},
			"successretries": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Number of consecutive successful probes required to transition a service's state from DOWN to UP.",
			},
			"supportedvendorids": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "List of Supported-Vendor-Id attribute value pairs (AVPs) for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. A maximum eight of these AVPs are supported in a monitoring message.",
			},
			"tos": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Probe the service by encoding the destination IP address in the IP TOS (6) bits.",
			},
			"tosid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The TOS ID of the specified destination IP. Applicable only when the TOS parameter is set.",
			},
			"transparent": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The monitor is bound to a transparent device such as a firewall or router. The state of a transparent device depends on the responsiveness of the services behind it. If a transparent device is being monitored, a destination IP address must be specified. The probe is sent to the specified IP address by using the MAC address of the transparent device.",
			},
			"trofscode": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Code expected when the server is under maintenance",
			},
			"trofsstring": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String expected from the server for the service to be marked as trofs. Applicable to HTTP-ECV/TCP-ECV monitors.",
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "Type of monitor that you want to create.",
			},
			"units1": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("SEC"),
				Description: "Unit of measurement for the Deviation parameter. Cannot be changed after the monitor is created.",
			},
			"units2": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("SEC"),
				Description: "Unit of measurement for the Down Time parameter. Cannot be changed after the monitor is created.",
			},
			"units3": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("SEC"),
				Description: "monitor interval units",
			},
			"units4": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("SEC"),
				Description: "monitor response timeout units",
			},
			"username": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "User name with which to probe the RADIUS, NNTP, FTP, FTP-EXTENDED, MYSQL, MSSQL, POP3, CITRIX-AG, CITRIX-XD-DDC, CITRIX-WI-EXTENDED, CITRIX-XNC or CITRIX-XDM server.",
			},
			"validatecred": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Validate the credentials of the Xen Desktop DDC server user. Applicable to monitors of type CITRIX-XD-DDC.",
			},
			"vendorid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Vendor-Id value for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers.",
			},
			"vendorspecificacctapplicationids": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "List of Vendor-Specific-Acct-Application-Id attribute value pairs (AVPs) to use for monitoring Diameter servers. A maximum of eight of these AVPs are supported in a monitoring message. The specified value is combined with the value of vendorSpecificVendorId to obtain the Vendor-Specific-Application-Id AVP in the CER monitoring message.",
			},
			"vendorspecificauthapplicationids": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "List of Vendor-Specific-Auth-Application-Id attribute value pairs (AVPs) for the Capabilities-Exchange-Request (CER) message to use for monitoring Diameter servers. A maximum of eight of these AVPs are supported in a monitoring message. The specified value is combined with the value of vendorSpecificVendorId to obtain the Vendor-Specific-Application-Id AVP in the CER monitoring message.",
			},
			"vendorspecificvendorid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Vendor-Id to use in the Vendor-Specific-Application-Id grouped attribute-value pair (AVP) in the monitoring CER message. To specify Auth-Application-Id or Acct-Application-Id in Vendor-Specific-Application-Id, use vendorSpecificAuthApplicationIds or vendorSpecificAcctApplicationIds, respectively. Only one Vendor-Id is supported for all the Vendor-Specific-Application-Id AVPs in a CER monitoring message.",
			},
		},
	}
}

func lbmonitorGetThePayloadFromtheConfig(ctx context.Context, data *LbmonitorResourceModel) lb.Lbmonitor {
	tflog.Debug(ctx, "In lbmonitorGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lbmonitor := lb.Lbmonitor{}
	if !data.Snmpoid.IsNull() {
		lbmonitor.Snmpoid = data.Snmpoid.ValueString()
	}
	if !data.Action.IsNull() {
		lbmonitor.Action = data.Action.ValueString()
	}
	if !data.Alertretries.IsNull() {
		lbmonitor.Alertretries = utils.IntPtr(int(data.Alertretries.ValueInt64()))
	}
	if !data.Application.IsNull() {
		lbmonitor.Application = data.Application.ValueString()
	}
	if !data.Attribute.IsNull() {
		lbmonitor.Attribute = data.Attribute.ValueString()
	}
	if !data.Basedn.IsNull() {
		lbmonitor.Basedn = data.Basedn.ValueString()
	}
	if !data.Binddn.IsNull() {
		lbmonitor.Binddn = data.Binddn.ValueString()
	}
	if !data.Customheaders.IsNull() {
		lbmonitor.Customheaders = data.Customheaders.ValueString()
	}
	if !data.Database.IsNull() {
		lbmonitor.Database = data.Database.ValueString()
	}
	if !data.Destip.IsNull() {
		lbmonitor.Destip = data.Destip.ValueString()
	}
	if !data.Destport.IsNull() {
		lbmonitor.Destport = utils.IntPtr(int(data.Destport.ValueInt64()))
	}
	if !data.Deviation.IsNull() {
		lbmonitor.Deviation = utils.IntPtr(int(data.Deviation.ValueInt64()))
	}
	if !data.Dispatcherip.IsNull() {
		lbmonitor.Dispatcherip = data.Dispatcherip.ValueString()
	}
	if !data.Dispatcherport.IsNull() {
		lbmonitor.Dispatcherport = utils.IntPtr(int(data.Dispatcherport.ValueInt64()))
	}
	if !data.Domain.IsNull() {
		lbmonitor.Domain = data.Domain.ValueString()
	}
	if !data.Downtime.IsNull() {
		lbmonitor.Downtime = utils.IntPtr(int(data.Downtime.ValueInt64()))
	}
	if !data.Evalrule.IsNull() {
		lbmonitor.Evalrule = data.Evalrule.ValueString()
	}
	if !data.Failureretries.IsNull() {
		lbmonitor.Failureretries = utils.IntPtr(int(data.Failureretries.ValueInt64()))
	}
	if !data.Filename.IsNull() {
		lbmonitor.Filename = data.Filename.ValueString()
	}
	if !data.Filter.IsNull() {
		lbmonitor.Filter = data.Filter.ValueString()
	}
	if !data.Firmwarerevision.IsNull() {
		lbmonitor.Firmwarerevision = utils.IntPtr(int(data.Firmwarerevision.ValueInt64()))
	}
	if !data.Group.IsNull() {
		lbmonitor.Group = data.Group.ValueString()
	}
	if !data.Grpchealthcheck.IsNull() {
		lbmonitor.Grpchealthcheck = data.Grpchealthcheck.ValueString()
	}
	if !data.Grpcservicename.IsNull() {
		lbmonitor.Grpcservicename = data.Grpcservicename.ValueString()
	}
	if !data.Hostipaddress.IsNull() {
		lbmonitor.Hostipaddress = data.Hostipaddress.ValueString()
	}
	if !data.Hostname.IsNull() {
		lbmonitor.Hostname = data.Hostname.ValueString()
	}
	if !data.Httprequest.IsNull() {
		lbmonitor.Httprequest = data.Httprequest.ValueString()
	}
	if !data.Inbandsecurityid.IsNull() {
		lbmonitor.Inbandsecurityid = data.Inbandsecurityid.ValueString()
	}
	if !data.Interval.IsNull() {
		lbmonitor.Interval = utils.IntPtr(int(data.Interval.ValueInt64()))
	}
	if !data.Iptunnel.IsNull() {
		lbmonitor.Iptunnel = data.Iptunnel.ValueString()
	}
	if !data.Kcdaccount.IsNull() {
		lbmonitor.Kcdaccount = data.Kcdaccount.ValueString()
	}
	if !data.Lasversion.IsNull() {
		lbmonitor.Lasversion = data.Lasversion.ValueString()
	}
	if !data.Logonpointname.IsNull() {
		lbmonitor.Logonpointname = data.Logonpointname.ValueString()
	}
	if !data.Lrtm.IsNull() {
		lbmonitor.Lrtm = data.Lrtm.ValueString()
	}
	if !data.Maxforwards.IsNull() {
		lbmonitor.Maxforwards = utils.IntPtr(int(data.Maxforwards.ValueInt64()))
	}
	if !data.Metric.IsNull() {
		lbmonitor.Metric = data.Metric.ValueString()
	}
	if !data.Metrictable.IsNull() {
		lbmonitor.Metrictable = data.Metrictable.ValueString()
	}
	if !data.Metricthreshold.IsNull() {
		lbmonitor.Metricthreshold = utils.IntPtr(int(data.Metricthreshold.ValueInt64()))
	}
	if !data.Metricweight.IsNull() {
		lbmonitor.Metricweight = utils.IntPtr(int(data.Metricweight.ValueInt64()))
	}
	if !data.Monitorname.IsNull() {
		lbmonitor.Monitorname = data.Monitorname.ValueString()
	}
	if !data.Mqttclientidentifier.IsNull() {
		lbmonitor.Mqttclientidentifier = data.Mqttclientidentifier.ValueString()
	}
	if !data.Mqttversion.IsNull() {
		lbmonitor.Mqttversion = utils.IntPtr(int(data.Mqttversion.ValueInt64()))
	}
	if !data.Mssqlprotocolversion.IsNull() {
		lbmonitor.Mssqlprotocolversion = data.Mssqlprotocolversion.ValueString()
	}
	if !data.Netprofile.IsNull() {
		lbmonitor.Netprofile = data.Netprofile.ValueString()
	}
	if !data.Oraclesid.IsNull() {
		lbmonitor.Oraclesid = data.Oraclesid.ValueString()
	}
	if !data.Originhost.IsNull() {
		lbmonitor.Originhost = data.Originhost.ValueString()
	}
	if !data.Originrealm.IsNull() {
		lbmonitor.Originrealm = data.Originrealm.ValueString()
	}
	if !data.Password.IsNull() {
		lbmonitor.Password = data.Password.ValueString()
	}
	if !data.Productname.IsNull() {
		lbmonitor.Productname = data.Productname.ValueString()
	}
	if !data.Query.IsNull() {
		lbmonitor.Query = data.Query.ValueString()
	}
	if !data.Querytype.IsNull() {
		lbmonitor.Querytype = data.Querytype.ValueString()
	}
	if !data.Radaccountsession.IsNull() {
		lbmonitor.Radaccountsession = data.Radaccountsession.ValueString()
	}
	if !data.Radaccounttype.IsNull() {
		lbmonitor.Radaccounttype = utils.IntPtr(int(data.Radaccounttype.ValueInt64()))
	}
	if !data.Radapn.IsNull() {
		lbmonitor.Radapn = data.Radapn.ValueString()
	}
	if !data.Radframedip.IsNull() {
		lbmonitor.Radframedip = data.Radframedip.ValueString()
	}
	if !data.Radkey.IsNull() {
		lbmonitor.Radkey = data.Radkey.ValueString()
	}
	if !data.Radmsisdn.IsNull() {
		lbmonitor.Radmsisdn = data.Radmsisdn.ValueString()
	}
	if !data.Radnasid.IsNull() {
		lbmonitor.Radnasid = data.Radnasid.ValueString()
	}
	if !data.Radnasip.IsNull() {
		lbmonitor.Radnasip = data.Radnasip.ValueString()
	}
	if !data.Recv.IsNull() {
		lbmonitor.Recv = data.Recv.ValueString()
	}
	if !data.Resptimeout.IsNull() {
		lbmonitor.Resptimeout = utils.IntPtr(int(data.Resptimeout.ValueInt64()))
	}
	if !data.Resptimeoutthresh.IsNull() {
		lbmonitor.Resptimeoutthresh = utils.IntPtr(int(data.Resptimeoutthresh.ValueInt64()))
	}
	if !data.Retries.IsNull() {
		lbmonitor.Retries = utils.IntPtr(int(data.Retries.ValueInt64()))
	}
	if !data.Reverse.IsNull() {
		lbmonitor.Reverse = data.Reverse.ValueString()
	}
	if !data.Rtsprequest.IsNull() {
		lbmonitor.Rtsprequest = data.Rtsprequest.ValueString()
	}
	if !data.Scriptargs.IsNull() {
		lbmonitor.Scriptargs = data.Scriptargs.ValueString()
	}
	if !data.Scriptname.IsNull() {
		lbmonitor.Scriptname = data.Scriptname.ValueString()
	}
	if !data.Secondarypassword.IsNull() {
		lbmonitor.Secondarypassword = data.Secondarypassword.ValueString()
	}
	if !data.Secure.IsNull() {
		lbmonitor.Secure = data.Secure.ValueString()
	}
	if !data.Secureargs.IsNull() {
		lbmonitor.Secureargs = data.Secureargs.ValueString()
	}
	if !data.Send.IsNull() {
		lbmonitor.Send = data.Send.ValueString()
	}
	if !data.Servicegroupname.IsNull() {
		lbmonitor.Servicegroupname = data.Servicegroupname.ValueString()
	}
	if !data.Servicename.IsNull() {
		lbmonitor.Servicename = data.Servicename.ValueString()
	}
	if !data.Sipmethod.IsNull() {
		lbmonitor.Sipmethod = data.Sipmethod.ValueString()
	}
	if !data.Sipreguri.IsNull() {
		lbmonitor.Sipreguri = data.Sipreguri.ValueString()
	}
	if !data.Sipuri.IsNull() {
		lbmonitor.Sipuri = data.Sipuri.ValueString()
	}
	if !data.Sitepath.IsNull() {
		lbmonitor.Sitepath = data.Sitepath.ValueString()
	}
	if !data.Snmpcommunity.IsNull() {
		lbmonitor.Snmpcommunity = data.Snmpcommunity.ValueString()
	}
	if !data.Snmpthreshold.IsNull() {
		lbmonitor.Snmpthreshold = data.Snmpthreshold.ValueString()
	}
	if !data.Snmpversion.IsNull() {
		lbmonitor.Snmpversion = data.Snmpversion.ValueString()
	}
	if !data.Sqlquery.IsNull() {
		lbmonitor.Sqlquery = data.Sqlquery.ValueString()
	}
	if !data.Sslprofile.IsNull() {
		lbmonitor.Sslprofile = data.Sslprofile.ValueString()
	}
	if !data.State.IsNull() {
		lbmonitor.State = data.State.ValueString()
	}
	if !data.Storedb.IsNull() {
		lbmonitor.Storedb = data.Storedb.ValueString()
	}
	if !data.Storefrontacctservice.IsNull() {
		lbmonitor.Storefrontacctservice = data.Storefrontacctservice.ValueString()
	}
	if !data.Storefrontcheckbackendservices.IsNull() {
		lbmonitor.Storefrontcheckbackendservices = data.Storefrontcheckbackendservices.ValueString()
	}
	if !data.Storename.IsNull() {
		lbmonitor.Storename = data.Storename.ValueString()
	}
	if !data.Successretries.IsNull() {
		lbmonitor.Successretries = utils.IntPtr(int(data.Successretries.ValueInt64()))
	}
	if !data.Tos.IsNull() {
		lbmonitor.Tos = data.Tos.ValueString()
	}
	if !data.Tosid.IsNull() {
		lbmonitor.Tosid = utils.IntPtr(int(data.Tosid.ValueInt64()))
	}
	if !data.Transparent.IsNull() {
		lbmonitor.Transparent = data.Transparent.ValueString()
	}
	if !data.Trofscode.IsNull() {
		lbmonitor.Trofscode = utils.IntPtr(int(data.Trofscode.ValueInt64()))
	}
	if !data.Trofsstring.IsNull() {
		lbmonitor.Trofsstring = data.Trofsstring.ValueString()
	}
	if !data.Type.IsNull() {
		lbmonitor.Type = data.Type.ValueString()
	}
	if !data.Units1.IsNull() {
		lbmonitor.Units1 = data.Units1.ValueString()
	}
	if !data.Units2.IsNull() {
		lbmonitor.Units2 = data.Units2.ValueString()
	}
	if !data.Units3.IsNull() {
		lbmonitor.Units3 = data.Units3.ValueString()
	}
	if !data.Units4.IsNull() {
		lbmonitor.Units4 = data.Units4.ValueString()
	}
	if !data.Username.IsNull() {
		lbmonitor.Username = data.Username.ValueString()
	}
	if !data.Validatecred.IsNull() {
		lbmonitor.Validatecred = data.Validatecred.ValueString()
	}
	if !data.Vendorid.IsNull() {
		lbmonitor.Vendorid = utils.IntPtr(int(data.Vendorid.ValueInt64()))
	}
	if !data.Vendorspecificvendorid.IsNull() {
		lbmonitor.Vendorspecificvendorid = utils.IntPtr(int(data.Vendorspecificvendorid.ValueInt64()))
	}

	return lbmonitor
}

func lbmonitorSetAttrFromGet(ctx context.Context, data *LbmonitorResourceModel, getResponseData map[string]interface{}) *LbmonitorResourceModel {
	tflog.Debug(ctx, "In lbmonitorSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["Snmpoid"]; ok && val != nil {
		data.Snmpoid = types.StringValue(val.(string))
	} else {
		data.Snmpoid = types.StringNull()
	}
	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else {
		data.Action = types.StringNull()
	}
	if val, ok := getResponseData["alertretries"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Alertretries = types.Int64Value(intVal)
		}
	} else {
		data.Alertretries = types.Int64Null()
	}
	if val, ok := getResponseData["application"]; ok && val != nil {
		data.Application = types.StringValue(val.(string))
	} else {
		data.Application = types.StringNull()
	}
	if val, ok := getResponseData["attribute"]; ok && val != nil {
		data.Attribute = types.StringValue(val.(string))
	} else {
		data.Attribute = types.StringNull()
	}
	if val, ok := getResponseData["basedn"]; ok && val != nil {
		data.Basedn = types.StringValue(val.(string))
	} else {
		data.Basedn = types.StringNull()
	}
	if val, ok := getResponseData["binddn"]; ok && val != nil {
		data.Binddn = types.StringValue(val.(string))
	} else {
		data.Binddn = types.StringNull()
	}
	if val, ok := getResponseData["customheaders"]; ok && val != nil {
		data.Customheaders = types.StringValue(val.(string))
	} else {
		data.Customheaders = types.StringNull()
	}
	if val, ok := getResponseData["database"]; ok && val != nil {
		data.Database = types.StringValue(val.(string))
	} else {
		data.Database = types.StringNull()
	}
	if val, ok := getResponseData["destip"]; ok && val != nil {
		data.Destip = types.StringValue(val.(string))
	} else {
		data.Destip = types.StringNull()
	}
	if val, ok := getResponseData["destport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Destport = types.Int64Value(intVal)
		}
	} else {
		data.Destport = types.Int64Null()
	}
	if val, ok := getResponseData["deviation"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Deviation = types.Int64Value(intVal)
		}
	} else {
		data.Deviation = types.Int64Null()
	}
	if val, ok := getResponseData["dispatcherip"]; ok && val != nil {
		data.Dispatcherip = types.StringValue(val.(string))
	} else {
		data.Dispatcherip = types.StringNull()
	}
	if val, ok := getResponseData["dispatcherport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Dispatcherport = types.Int64Value(intVal)
		}
	} else {
		data.Dispatcherport = types.Int64Null()
	}
	if val, ok := getResponseData["domain"]; ok && val != nil {
		data.Domain = types.StringValue(val.(string))
	} else {
		data.Domain = types.StringNull()
	}
	if val, ok := getResponseData["downtime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Downtime = types.Int64Value(intVal)
		}
	} else {
		data.Downtime = types.Int64Null()
	}
	if val, ok := getResponseData["evalrule"]; ok && val != nil {
		data.Evalrule = types.StringValue(val.(string))
	} else {
		data.Evalrule = types.StringNull()
	}
	if val, ok := getResponseData["failureretries"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Failureretries = types.Int64Value(intVal)
		}
	} else {
		data.Failureretries = types.Int64Null()
	}
	if val, ok := getResponseData["filename"]; ok && val != nil {
		data.Filename = types.StringValue(val.(string))
	} else {
		data.Filename = types.StringNull()
	}
	if val, ok := getResponseData["filter"]; ok && val != nil {
		data.Filter = types.StringValue(val.(string))
	} else {
		data.Filter = types.StringNull()
	}
	if val, ok := getResponseData["firmwarerevision"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Firmwarerevision = types.Int64Value(intVal)
		}
	} else {
		data.Firmwarerevision = types.Int64Null()
	}
	if val, ok := getResponseData["group"]; ok && val != nil {
		data.Group = types.StringValue(val.(string))
	} else {
		data.Group = types.StringNull()
	}
	if val, ok := getResponseData["grpchealthcheck"]; ok && val != nil {
		data.Grpchealthcheck = types.StringValue(val.(string))
	} else {
		data.Grpchealthcheck = types.StringNull()
	}
	if val, ok := getResponseData["grpcservicename"]; ok && val != nil {
		data.Grpcservicename = types.StringValue(val.(string))
	} else {
		data.Grpcservicename = types.StringNull()
	}
	if val, ok := getResponseData["hostipaddress"]; ok && val != nil {
		data.Hostipaddress = types.StringValue(val.(string))
	} else {
		data.Hostipaddress = types.StringNull()
	}
	if val, ok := getResponseData["hostname"]; ok && val != nil {
		data.Hostname = types.StringValue(val.(string))
	} else {
		data.Hostname = types.StringNull()
	}
	if val, ok := getResponseData["httprequest"]; ok && val != nil {
		data.Httprequest = types.StringValue(val.(string))
	} else {
		data.Httprequest = types.StringNull()
	}
	if val, ok := getResponseData["inbandsecurityid"]; ok && val != nil {
		data.Inbandsecurityid = types.StringValue(val.(string))
	} else {
		data.Inbandsecurityid = types.StringNull()
	}
	if val, ok := getResponseData["interval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Interval = types.Int64Value(intVal)
		}
	} else {
		data.Interval = types.Int64Null()
	}
	if val, ok := getResponseData["iptunnel"]; ok && val != nil {
		data.Iptunnel = types.StringValue(val.(string))
	} else {
		data.Iptunnel = types.StringNull()
	}
	if val, ok := getResponseData["kcdaccount"]; ok && val != nil {
		data.Kcdaccount = types.StringValue(val.(string))
	} else {
		data.Kcdaccount = types.StringNull()
	}
	if val, ok := getResponseData["lasversion"]; ok && val != nil {
		data.Lasversion = types.StringValue(val.(string))
	} else {
		data.Lasversion = types.StringNull()
	}
	if val, ok := getResponseData["logonpointname"]; ok && val != nil {
		data.Logonpointname = types.StringValue(val.(string))
	} else {
		data.Logonpointname = types.StringNull()
	}
	if val, ok := getResponseData["lrtm"]; ok && val != nil {
		data.Lrtm = types.StringValue(val.(string))
	} else {
		data.Lrtm = types.StringNull()
	}
	if val, ok := getResponseData["maxforwards"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxforwards = types.Int64Value(intVal)
		}
	} else {
		data.Maxforwards = types.Int64Null()
	}
	if val, ok := getResponseData["metric"]; ok && val != nil {
		data.Metric = types.StringValue(val.(string))
	} else {
		data.Metric = types.StringNull()
	}
	if val, ok := getResponseData["metrictable"]; ok && val != nil {
		data.Metrictable = types.StringValue(val.(string))
	} else {
		data.Metrictable = types.StringNull()
	}
	if val, ok := getResponseData["metricthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Metricthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Metricthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["metricweight"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Metricweight = types.Int64Value(intVal)
		}
	} else {
		data.Metricweight = types.Int64Null()
	}
	if val, ok := getResponseData["monitorname"]; ok && val != nil {
		data.Monitorname = types.StringValue(val.(string))
	} else {
		data.Monitorname = types.StringNull()
	}
	if val, ok := getResponseData["mqttclientidentifier"]; ok && val != nil {
		data.Mqttclientidentifier = types.StringValue(val.(string))
	} else {
		data.Mqttclientidentifier = types.StringNull()
	}
	if val, ok := getResponseData["mqttversion"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mqttversion = types.Int64Value(intVal)
		}
	} else {
		data.Mqttversion = types.Int64Null()
	}
	if val, ok := getResponseData["mssqlprotocolversion"]; ok && val != nil {
		data.Mssqlprotocolversion = types.StringValue(val.(string))
	} else {
		data.Mssqlprotocolversion = types.StringNull()
	}
	if val, ok := getResponseData["netprofile"]; ok && val != nil {
		data.Netprofile = types.StringValue(val.(string))
	} else {
		data.Netprofile = types.StringNull()
	}
	if val, ok := getResponseData["oraclesid"]; ok && val != nil {
		data.Oraclesid = types.StringValue(val.(string))
	} else {
		data.Oraclesid = types.StringNull()
	}
	if val, ok := getResponseData["originhost"]; ok && val != nil {
		data.Originhost = types.StringValue(val.(string))
	} else {
		data.Originhost = types.StringNull()
	}
	if val, ok := getResponseData["originrealm"]; ok && val != nil {
		data.Originrealm = types.StringValue(val.(string))
	} else {
		data.Originrealm = types.StringNull()
	}
	if val, ok := getResponseData["password"]; ok && val != nil {
		data.Password = types.StringValue(val.(string))
	} else {
		data.Password = types.StringNull()
	}
	if val, ok := getResponseData["productname"]; ok && val != nil {
		data.Productname = types.StringValue(val.(string))
	} else {
		data.Productname = types.StringNull()
	}
	if val, ok := getResponseData["query"]; ok && val != nil {
		data.Query = types.StringValue(val.(string))
	} else {
		data.Query = types.StringNull()
	}
	if val, ok := getResponseData["querytype"]; ok && val != nil {
		data.Querytype = types.StringValue(val.(string))
	} else {
		data.Querytype = types.StringNull()
	}
	if val, ok := getResponseData["radaccountsession"]; ok && val != nil {
		data.Radaccountsession = types.StringValue(val.(string))
	} else {
		data.Radaccountsession = types.StringNull()
	}
	if val, ok := getResponseData["radaccounttype"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Radaccounttype = types.Int64Value(intVal)
		}
	} else {
		data.Radaccounttype = types.Int64Null()
	}
	if val, ok := getResponseData["radapn"]; ok && val != nil {
		data.Radapn = types.StringValue(val.(string))
	} else {
		data.Radapn = types.StringNull()
	}
	if val, ok := getResponseData["radframedip"]; ok && val != nil {
		data.Radframedip = types.StringValue(val.(string))
	} else {
		data.Radframedip = types.StringNull()
	}
	if val, ok := getResponseData["radkey"]; ok && val != nil {
		data.Radkey = types.StringValue(val.(string))
	} else {
		data.Radkey = types.StringNull()
	}
	if val, ok := getResponseData["radmsisdn"]; ok && val != nil {
		data.Radmsisdn = types.StringValue(val.(string))
	} else {
		data.Radmsisdn = types.StringNull()
	}
	if val, ok := getResponseData["radnasid"]; ok && val != nil {
		data.Radnasid = types.StringValue(val.(string))
	} else {
		data.Radnasid = types.StringNull()
	}
	if val, ok := getResponseData["radnasip"]; ok && val != nil {
		data.Radnasip = types.StringValue(val.(string))
	} else {
		data.Radnasip = types.StringNull()
	}
	if val, ok := getResponseData["recv"]; ok && val != nil {
		data.Recv = types.StringValue(val.(string))
	} else {
		data.Recv = types.StringNull()
	}
	if val, ok := getResponseData["resptimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Resptimeout = types.Int64Value(intVal)
		}
	} else {
		data.Resptimeout = types.Int64Null()
	}
	if val, ok := getResponseData["resptimeoutthresh"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Resptimeoutthresh = types.Int64Value(intVal)
		}
	} else {
		data.Resptimeoutthresh = types.Int64Null()
	}
	if val, ok := getResponseData["retries"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Retries = types.Int64Value(intVal)
		}
	} else {
		data.Retries = types.Int64Null()
	}
	if val, ok := getResponseData["reverse"]; ok && val != nil {
		data.Reverse = types.StringValue(val.(string))
	} else {
		data.Reverse = types.StringNull()
	}
	if val, ok := getResponseData["rtsprequest"]; ok && val != nil {
		data.Rtsprequest = types.StringValue(val.(string))
	} else {
		data.Rtsprequest = types.StringNull()
	}
	if val, ok := getResponseData["scriptargs"]; ok && val != nil {
		data.Scriptargs = types.StringValue(val.(string))
	} else {
		data.Scriptargs = types.StringNull()
	}
	if val, ok := getResponseData["scriptname"]; ok && val != nil {
		data.Scriptname = types.StringValue(val.(string))
	} else {
		data.Scriptname = types.StringNull()
	}
	if val, ok := getResponseData["secondarypassword"]; ok && val != nil {
		data.Secondarypassword = types.StringValue(val.(string))
	} else {
		data.Secondarypassword = types.StringNull()
	}
	if val, ok := getResponseData["secure"]; ok && val != nil {
		data.Secure = types.StringValue(val.(string))
	} else {
		data.Secure = types.StringNull()
	}
	if val, ok := getResponseData["secureargs"]; ok && val != nil {
		data.Secureargs = types.StringValue(val.(string))
	} else {
		data.Secureargs = types.StringNull()
	}
	if val, ok := getResponseData["send"]; ok && val != nil {
		data.Send = types.StringValue(val.(string))
	} else {
		data.Send = types.StringNull()
	}
	if val, ok := getResponseData["servicegroupname"]; ok && val != nil {
		data.Servicegroupname = types.StringValue(val.(string))
	} else {
		data.Servicegroupname = types.StringNull()
	}
	if val, ok := getResponseData["servicename"]; ok && val != nil {
		data.Servicename = types.StringValue(val.(string))
	} else {
		data.Servicename = types.StringNull()
	}
	if val, ok := getResponseData["sipmethod"]; ok && val != nil {
		data.Sipmethod = types.StringValue(val.(string))
	} else {
		data.Sipmethod = types.StringNull()
	}
	if val, ok := getResponseData["sipreguri"]; ok && val != nil {
		data.Sipreguri = types.StringValue(val.(string))
	} else {
		data.Sipreguri = types.StringNull()
	}
	if val, ok := getResponseData["sipuri"]; ok && val != nil {
		data.Sipuri = types.StringValue(val.(string))
	} else {
		data.Sipuri = types.StringNull()
	}
	if val, ok := getResponseData["sitepath"]; ok && val != nil {
		data.Sitepath = types.StringValue(val.(string))
	} else {
		data.Sitepath = types.StringNull()
	}
	if val, ok := getResponseData["snmpcommunity"]; ok && val != nil {
		data.Snmpcommunity = types.StringValue(val.(string))
	} else {
		data.Snmpcommunity = types.StringNull()
	}
	if val, ok := getResponseData["snmpthreshold"]; ok && val != nil {
		data.Snmpthreshold = types.StringValue(val.(string))
	} else {
		data.Snmpthreshold = types.StringNull()
	}
	if val, ok := getResponseData["snmpversion"]; ok && val != nil {
		data.Snmpversion = types.StringValue(val.(string))
	} else {
		data.Snmpversion = types.StringNull()
	}
	if val, ok := getResponseData["sqlquery"]; ok && val != nil {
		data.Sqlquery = types.StringValue(val.(string))
	} else {
		data.Sqlquery = types.StringNull()
	}
	if val, ok := getResponseData["sslprofile"]; ok && val != nil {
		data.Sslprofile = types.StringValue(val.(string))
	} else {
		data.Sslprofile = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}
	if val, ok := getResponseData["storedb"]; ok && val != nil {
		data.Storedb = types.StringValue(val.(string))
	} else {
		data.Storedb = types.StringNull()
	}
	if val, ok := getResponseData["storefrontacctservice"]; ok && val != nil {
		data.Storefrontacctservice = types.StringValue(val.(string))
	} else {
		data.Storefrontacctservice = types.StringNull()
	}
	if val, ok := getResponseData["storefrontcheckbackendservices"]; ok && val != nil {
		data.Storefrontcheckbackendservices = types.StringValue(val.(string))
	} else {
		data.Storefrontcheckbackendservices = types.StringNull()
	}
	if val, ok := getResponseData["storename"]; ok && val != nil {
		data.Storename = types.StringValue(val.(string))
	} else {
		data.Storename = types.StringNull()
	}
	if val, ok := getResponseData["successretries"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Successretries = types.Int64Value(intVal)
		}
	} else {
		data.Successretries = types.Int64Null()
	}
	if val, ok := getResponseData["tos"]; ok && val != nil {
		data.Tos = types.StringValue(val.(string))
	} else {
		data.Tos = types.StringNull()
	}
	if val, ok := getResponseData["tosid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Tosid = types.Int64Value(intVal)
		}
	} else {
		data.Tosid = types.Int64Null()
	}
	if val, ok := getResponseData["transparent"]; ok && val != nil {
		data.Transparent = types.StringValue(val.(string))
	} else {
		data.Transparent = types.StringNull()
	}
	if val, ok := getResponseData["trofscode"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Trofscode = types.Int64Value(intVal)
		}
	} else {
		data.Trofscode = types.Int64Null()
	}
	if val, ok := getResponseData["trofsstring"]; ok && val != nil {
		data.Trofsstring = types.StringValue(val.(string))
	} else {
		data.Trofsstring = types.StringNull()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}
	if val, ok := getResponseData["units1"]; ok && val != nil {
		data.Units1 = types.StringValue(val.(string))
	} else {
		data.Units1 = types.StringNull()
	}
	if val, ok := getResponseData["units2"]; ok && val != nil {
		data.Units2 = types.StringValue(val.(string))
	} else {
		data.Units2 = types.StringNull()
	}
	if val, ok := getResponseData["units3"]; ok && val != nil {
		data.Units3 = types.StringValue(val.(string))
	} else {
		data.Units3 = types.StringNull()
	}
	if val, ok := getResponseData["units4"]; ok && val != nil {
		data.Units4 = types.StringValue(val.(string))
	} else {
		data.Units4 = types.StringNull()
	}
	if val, ok := getResponseData["username"]; ok && val != nil {
		data.Username = types.StringValue(val.(string))
	} else {
		data.Username = types.StringNull()
	}
	if val, ok := getResponseData["validatecred"]; ok && val != nil {
		data.Validatecred = types.StringValue(val.(string))
	} else {
		data.Validatecred = types.StringNull()
	}
	if val, ok := getResponseData["vendorid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vendorid = types.Int64Value(intVal)
		}
	} else {
		data.Vendorid = types.Int64Null()
	}
	if val, ok := getResponseData["vendorspecificvendorid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vendorspecificvendorid = types.Int64Value(intVal)
		}
	} else {
		data.Vendorspecificvendorid = types.Int64Null()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated
	data.Id = types.StringValue(fmt.Sprintf("%s", data.Monitorname.ValueString()))

	return data
}
