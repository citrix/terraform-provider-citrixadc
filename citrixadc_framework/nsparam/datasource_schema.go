package nsparam

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func NsparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"advancedanalyticsstats": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Disable/Enable advanace analytics stats",
			},
			"aftpallowrandomsourceport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow the FTP server to come from a random source port for active FTP data connections",
			},
			"cip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable the insertion of the actual client IP address into the HTTP header request passed from the client to one, some, or all servers attached to the system. The passed address can then be accessed through a minor modification to the server.\n* If the CIP header is specified, it will be used as the client IP header.\n* If the CIP header is not specified, the value that has been set will be used as the client IP header.",
			},
			"cipheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Text that will be used as the client IP address header.",
			},
			"cookieversion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Version of the cookie inserted by the system.",
			},
			"crportrange": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Port range for cache redirection services.",
			},
			"exclusivequotamaxclient": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Percentage of maxClient threshold to be divided equally among PEs.",
			},
			"exclusivequotaspillover": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Percentage of spillover threshold to be divided equally among PEs.",
			},
			"ftpportrange": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum and maximum port (port range) that FTP services are allowed to use.",
			},
			"grantquotamaxclient": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Percentage of shared pool value granted to PE once PE exhausts the local exclusive quota. Where shared pool is the remaining maxclient quota after distribution of exclusive quota to PEs.\n\nExample: In a 2 PE NetScaler system if configured maxclient is 100 and exclusive quota is 80 percent then each PE will get 40 as local exclusive quota and 20 will be in shared pool. If configured grantQuota is 20 percent, then after exhausting its local exclusive quota PE borrows from shared pool in chunks of 4 i.e. 20 percent of 20.",
			},
			"grantquotaspillover": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Percentage of shared pool value granted to PE once PE exhausts the local exclusive quota. Where shared pool is the remaining spillover quota after distribution of exclusive quota to PEs.\n\nExample: In a 2 PE NetScaler system if configured spillover is 100 and exclusive quota is 80 percent then each PE will get 40 as local exclusive quota and 20 will be in shared pool. If configured grantQuota is 20 percent, then after exhausting its local exclusive quota PE borrows from shared pool in chunks of 4 i.e. 20 percent of 20.",
			},
			"httpport": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "HTTP ports on the web server. This allows the system to perform connection off-load for any client request that has a destination port matching one of these configured ports.",
			},
			"icaports": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "The ICA ports on the Web server. This allows the system to perform connection off-load for any client request that has a destination port matching one of these configured ports.",
			},
			"internaluserlogin": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enables/disables the internal user from logging in to the appliance. Before disabling internal user login, you must have key-based authentication set up on the appliance. The file name for the key pair must be \"ns_comm_key\".",
			},
			"ipttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Set the IP Time to Live (TTL) and Hop Limit value for all outgoing packets from Citrix ADC.",
			},
			"maxconn": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of connections that will be made from the appliance to the web server(s) attached to it. The value entered here is applied globally to all attached servers.",
			},
			"maxreq": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of requests that the system can pass on a particular connection between the appliance and a server attached to it. Setting this value to 0 allows an unlimited number of requests to be passed. This value is overridden by the maximum number of requests configured on the individual service.",
			},
			"mgmthttpport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This allow the configuration of management HTTP port.",
			},
			"mgmthttpsport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This allows the configuration of management HTTPS port.",
			},
			"pmtumin": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum path MTU value that Citrix ADC will process in the ICMP fragmentation needed message. If the ICMP message contains a value less than this value, then this value is used instead.",
			},
			"pmtutimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Interval, in minutes, for flushing the PMTU entries.",
			},
			"proxyprotocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Disable/Enable v1 or v2 proxy protocol header for client info insertion",
			},
			"securecookie": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable secure flag for persistence cookie.",
			},
			"secureicaports": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "The Secure ICA ports on the Web server. This allows the system to perform connection off-load for any\n            client request that has a destination port matching one of these configured ports.",
			},
			"servicepathingressvlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "VLAN on which the subscriber traffic arrives on the appliance.",
			},
			"tcpcip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable the insertion of the client TCP/IP header in TCP payload passed from the client to one, some, or all servers attached to the system. The passed address can then be accessed through a minor modification to the server.",
			},
			"timezone": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Time zone for the Citrix ADC. Name of the time zone should be specified as argument.",
			},
			"useproxyport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable use_proxy_port setting",
			},
		},
	}
}
