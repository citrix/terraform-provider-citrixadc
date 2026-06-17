package nsconfig

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func NsconfigDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"async": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Using this option will run the operation in async mode and return the job id. The job ID can be used later to track the conversion progress via show ns job <id> Command. This option is mostly useful for API to avoid timeouts for large input configuration",
			},
			"all": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this option to do saveconfig for all partitions",
			},
			"changedpassword": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to list all passwords changed which would not work when downgraded to older releases. Takes config file as input, if no input specified, running configuration is considered. Command => query ns config -changedpassword / query ns config -changedpassword /nsconfig/ns.conf",
			},
			"cip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The option to control (enable or disable) the insertion of the actual client IP address into the HTTP header request passed from the client to one, some, or all servers attached to the system.\nThe passed address can then be accessed through a minor modification to the server.\nl    If cipHeader is specified, it will be used as the client IP header.\nl    If it is not specified, then the value that has been set by the set ns config CLI command will be used as the client IP header.",
			},
			"cipheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The text that will be used as the client IP header.",
			},
			"config": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "configuration File to be used to find weak passwords, if not specified, running config is taken as input.",
			},
			"config1": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Location of the configurations.",
			},
			"config2": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Location of the configurations.",
			},
			"configfile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Full path of config file to be converted to nitro",
			},
			"cookieversion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The version of the cookie inserted by system.",
			},
			"crportrange": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Port range for cache redirection services.",
			},
			"exclusivequotamaxclient": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The percentage of maxClient to be given to PEs",
			},
			"exclusivequotaspillover": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The percentage of spillover threshold to be given to PEs",
			},
			"force": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Configurations will be cleared without prompting for confirmation.",
			},
			"ftpportrange": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Port range configured for FTP services.",
			},
			"grantquotamaxclient": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The percentage of shared quota to be granted at a time for maxClient",
			},
			"grantquotaspillover": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The percentage of shared quota to be granted at a time for spillover",
			},
			"httpport": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "The HTTP ports on the Web server. This allows the system to perform connection off-load for any client request that has a destination port matching one of these configured ports.",
			},
			"ifnum": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Interfaces of the appliances that must be bound to the NSVLAN.",
			},
			"ignoredevicespecific": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Suppress device specific differences.",
			},
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the Citrix ADC. Commonly referred to as NSIP address. This parameter is mandatory to bring up the appliance.",
			},
			"level": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Types of configurations to be cleared.\n* basic: Clears all configurations except the following:\n  - NSIP, default route (gateway), static routes, MIPs, and SNIPs\n  - Network settings (DG, VLAN, RHI and DNS settings)\n  - Cluster settings\n  - HA node definitions\n  - Feature and mode settings\n  - nsroot password\n* extended: Clears the same configurations as the 'basic' option. In addition, it clears the feature and mode settings.\n* full: Clears all configurations except NSIP, default route, and interface settings.\nNote: When you clear the configurations through the cluster IP address, by specifying the level as 'full', the cluster is deleted and all cluster nodes become standalone appliances. The 'basic' and 'extended' levels are propagated to the cluster nodes.",
			},
			"maxconn": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The maximum number of connections that will be made from the system to the web server(s) attached to it. The value entered here is applied globally to all attached servers.",
			},
			"maxreq": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The maximum number of requests that the system can pass on a particular connection between the system and a server attached to it. Setting this value to 0 allows an unlimited number of requests to be passed.",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Netmask corresponding to the IP address. This parameter is mandatory to bring up the appliance.",
			},
			"nsvlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "VLAN (NSVLAN) for the subnet on which the IP address resides.",
			},
			"outtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Format to display the difference in configurations.",
			},
			"pmtumin": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The minimum Path MTU.",
			},
			"pmtutimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The timeout value in minutes.",
			},
			"rbaconfig": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RBA configurations and TACACS policies bound to system global will not be cleared if RBA is set to NO.This option is applicable only for BASIC level of clear configuration.Default is YES, which will clear rba configurations.",
			},
			"responsefile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Full path of file to store the nitro graph. If not specified, nitro graph is returned as part of the API response.",
			},
			"securecookie": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "enable/disable secure flag for persistence cookie",
			},
			"securemanagementtd": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This positive integer identifies Management traffic domain. If not specified, defaults to 4094",
			},
			"securemanagementtraffic": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This enabled secure management traffic handling.",
			},
			"tagged": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies that the interfaces will be added as 802.1q tagged interfaces. Packets sent on these interface on this VLAN will have an additional 4-byte 802.1q tag which identifies the VLAN.\nTo use 802.1q tagging, the switch connected to the appliance's interfaces must also be configured for tagging.",
			},
			"template": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "File that contains the commands to be compared.",
			},
			"timezone": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the timezone",
			},
			"weakpassword": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to list all weak passwords (not adhering to strong password requirements). Takes config file as input, if no input specified, running configuration is considered. Command => query ns config -weakpassword  / query ns config -weakpassword /nsconfig/ns.conf",
			},
		},
	}
}
