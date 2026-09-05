package lsngroup

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LsngroupDataSourceModel is the data-source-specific model, decoupled from
// LsngroupResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the configurable attributes (as
// Computed outputs) AND the read-only attributes the resource deliberately
// omits. Every non-key attribute is Computed.
type LsngroupDataSourceModel struct {
	Id             types.String `tfsdk:"id"`
	Allocpolicy    types.String `tfsdk:"allocpolicy"`
	Clientname     types.String `tfsdk:"clientname"`
	Ftp            types.String `tfsdk:"ftp"`
	Ftpcm          types.String `tfsdk:"ftpcm"`
	Groupname      types.String `tfsdk:"groupname"` // Required lookup key
	Ip6profile     types.String `tfsdk:"ip6profile"`
	Logging        types.String `tfsdk:"logging"`
	Nattype        types.String `tfsdk:"nattype"`
	Portblocksize  types.Int64  `tfsdk:"portblocksize"`
	Pptp           types.String `tfsdk:"pptp"`
	Rtspalg        types.String `tfsdk:"rtspalg"`
	Sessionlogging types.String `tfsdk:"sessionlogging"`
	Sessionsync    types.String `tfsdk:"sessionsync"`
	Sipalg         types.String `tfsdk:"sipalg"`
	Snmptraplimit  types.Int64  `tfsdk:"snmptraplimit"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/lsngroup.json). Never settable; populated from GET.
	Groupid types.Int64 `tfsdk:"groupid"`
}

func LsngroupDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source to read LSN (Large Scale NAT) group configuration.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"allocpolicy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "NAT IP and PORT block allocation policy for Deterministic NAT. Supported Policies are,\n1: PORTS: Port blocks from single NATIP will be allocated to LSN subscribers sequentially. After all blocks are exhausted, port blocks from next NATIP will be allocated and so on.\n2: IPADDRS(Default): One port block from each NATIP will be allocated and once all the NATIPs are over second port block from each NATIP will be allocated and so on.\nTo understand better if we assume port blocks of all NAT IPs as two dimensional array, PORTS policy follows \"row major order\" and IPADDRS policy follows \"column major order\" while allocating port blocks.\nExample:\nClient IPs: 2.2.2.1, 2.2.2.2 and 2.2.2.3\nNAT IPs and PORT Blocks: \n4.4.4.1:PB1, PB2, PB3,., PBn\n4.4.4.2: PB1, PB2, PB3,., PBn\nPORTS Policy: \n2.2.2.1 => 4.4.4.1:PB1\n2.2.2.2 => 4.4.4.1:PB2\n2.2.2.3 => 4.4.4.1:PB3\nIPADDRS Policy:\n2.2.2.1 => 4.4.4.1:PB1\n2.2.2.2 => 4.4.4.2:PB1\n2.2.2.3 => 4.4.4.1:PB2",
			},
			"clientname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the LSN client entity to be associated with the LSN group. You can associate only one LSN client entity with an LSN group.You cannot remove this association or replace with another LSN client entity once the LSN group is created.",
			},
			"ftp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable Application Layer Gateway (ALG) for the FTP protocol. For some application-layer protocols, the IP addresses and protocol port numbers are usually communicated in the packet's payload. When acting as an ALG, the Citrix ADC changes the packet's payload to ensure that the protocol continues to work over LSN. \n\nNote:  The Citrix ADC also includes ALG for ICMP and TFTP protocols. ALG for the ICMP protocol is enabled by default, and there is no provision to disable it. ALG for the TFTP protocol is disabled by default. ALG is enabled automatically for an LSN group when you bind a UDP LSN application profile, with endpoint-independent-mapping, endpoint-independent filtering, and destination port as 69 (well-known port for TFTP), to the LSN group.",
			},
			"ftpcm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable the FTP connection mirroring for specified LSN group. Connection mirroring (CM or connection failover) refers to keeping active an established TCP or UDP connection when a failover occurs.",
			},
			"groupname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the LSN group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN group is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"lsn group1\" or 'lsn group1').",
			},
			"ip6profile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the LSN ip6 profile to associate with the specified LSN group. An ip6 profile can be associated with a group only during group creation.\n\nBy default, no LSN ip6 profile is associated with an LSN group during its creation. Only one ip6profile can be associated with a group.",
			},
			"logging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log mapping entries and sessions created or deleted for this LSN group. The Citrix ADC logs LSN sessions for this LSN group only when both logging and session logging parameters are enabled.\n\nThe ADC uses its existing syslog and audit log framework to log LSN information. You must enable global level LSN logging by enabling the LSN parameter in the related NSLOG action and SYLOG action entities. When the Logging parameter is enabled, the Citrix ADC generates log messages related to LSN mappings and LSN sessions of this LSN group. The ADC then sends these log messages to servers associated with the NSLOG action and SYSLOG actions entities. \n\nA log message for an LSN mapping entry consists of the following information:\n* NSIP address of the Citrix ADC\n* Time stamp\n* Entry type (MAPPING or SESSION)\n* Whether the LSN mapping entry is created or deleted\n* Subscriber's IP address, port, and traffic domain ID\n* NAT IP address and port\n* Protocol name\n* Destination IP address, port, and traffic domain ID might be  present, depending on the following conditions:\n** Destination IP address and port are not logged for Endpoint-Independent mapping\n** Only Destination IP address (and not port) is logged for Address-Dependent mapping\n** Destination IP address and port are logged for Address-Port-Dependent mapping",
			},
			"nattype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of NAT IP address and port allocation (from the bound LSN pools) for subscribers:\n\nAvailable options function as follows:\n\n* Deterministic - Allocate a NAT IP address and a block of ports to each subscriber (of the LSN client bound to the LSN group). The Citrix ADC sequentially allocates NAT resources to these subscribers. The Citrix ADC ADC assigns the first block of ports (block size determined by the port block size parameter of the LSN group) on the beginning NAT IP address to the beginning subscriber IP address. The next range of ports is assigned to the next subscriber, and so on, until the NAT address does not have enough ports for the next subscriber. In this case, the first port block on the next NAT address is used for the subscriber, and so on.  Because each subscriber now receives a deterministic NAT IP address and a block of ports, a subscriber can be identified without any need for logging. For a connection, a subscriber can be identified based only on the NAT IP address and port, and the destination IP address and port. The maximum number of LSN subscribers allowed, globally, is 1 million.  \n\n* Dynamic - Allocate a random NAT IP address and a port from the LSN NAT pool for a subscriber's connection. If port block allocation is enabled (in LSN pool) and a port block size is specified (in the LSN group), the Citrix ADC allocates a random NAT IP address and a block of ports for a subscriber when it initiates a connection for the first time. The ADC allocates this NAT IP address and a port (from the allocated block of ports) for different connections from this subscriber. If all the ports are allocated (for different subscriber's connections) from the subscriber's allocated port block, the ADC allocates a new random port block for the subscriber.",
			},
			"portblocksize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Size of the NAT port block to be allocated for each subscriber.\n\nTo set this parameter for Dynamic NAT, you must enable the port block allocation parameter in the bound LSN pool. For Deterministic NAT, the port block allocation parameter is always  enabled, and you cannot disable it.\n\nIn Dynamic NAT, the Citrix ADC allocates a random NAT port block, from the available NAT port pool of an NAT IP address, for each subscriber. For a subscriber, if all the ports are allocated from the subscriber's allocated port block, the ADC allocates a new random port block for the subscriber.\n\nThe default port block size is 256 for Deterministic NAT, and 0 for Dynamic NAT.",
			},
			"pptp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable the PPTP Application Layer Gateway.",
			},
			"rtspalg": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable the RTSP ALG.",
			},
			"sessionlogging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log sessions created or deleted for the LSN group. The Citrix ADC logs LSN sessions for this LSN group only when both logging and session logging parameters are enabled.\n\nA log message for an LSN session consists of the following information:\n* NSIP address of the Citrix ADC\n* Time stamp\n* Entry type (MAPPING or SESSION)\n* Whether the LSN session is created or removed\n* Subscriber's IP address, port, and traffic domain ID\n* NAT IP address and port\n* Protocol name\n* Destination IP address, port, and traffic domain ID",
			},
			"sessionsync": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "In a high availability (HA) deployment, synchronize information of all LSN sessions related to this LSN group with the secondary node. After a failover, established TCP connections and UDP packet flows are kept active and resumed on the secondary node (new primary).\n\nFor this setting to work, you must enable the global session synchronization parameter.",
			},
			"sipalg": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable the SIP ALG.",
			},
			"snmptraplimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of SNMP Trap messages that can be generated for the LSN group in one minute.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"groupid": schema.Int64Attribute{
				Computed:    true,
				Description: "Unique identifier assigned to the LSN group by the appliance.",
			},
		},
	}
}

// lsngroupDataSourceSetAttrFromGet projects a NITRO lsngroup GET response onto
// the data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func lsngroupDataSourceSetAttrFromGet(ctx context.Context, data *LsngroupDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In lsngroupDataSourceSetAttrFromGet Function")

	if v, ok := g["groupname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Groupname = types.StringValue(utils.AnyToString(v))
	}

	data.Allocpolicy = utils.MapGetString(g, "allocpolicy")
	data.Clientname = utils.MapGetString(g, "clientname")
	data.Ftp = utils.MapGetString(g, "ftp")
	data.Ftpcm = utils.MapGetString(g, "ftpcm")
	data.Ip6profile = utils.MapGetString(g, "ip6profile")
	data.Logging = utils.MapGetString(g, "logging")
	data.Nattype = utils.MapGetString(g, "nattype")
	data.Portblocksize = utils.MapGetInt64(g, "portblocksize")
	data.Pptp = utils.MapGetString(g, "pptp")
	data.Rtspalg = utils.MapGetString(g, "rtspalg")
	data.Sessionlogging = utils.MapGetString(g, "sessionlogging")
	data.Sessionsync = utils.MapGetString(g, "sessionsync")
	data.Sipalg = utils.MapGetString(g, "sipalg")
	data.Snmptraplimit = utils.MapGetInt64(g, "snmptraplimit")

	// Read-only attributes.
	data.Groupid = utils.MapGetInt64(g, "groupid")
}
