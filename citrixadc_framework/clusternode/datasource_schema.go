package clusternode

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ClusternodeDataSourceModel is the data-source-specific model, decoupled from
// ClusternodeResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only runtime/health attributes that the resource
// deliberately omits (clusterhealth, health, masterstate, ...). Every non-key
// attribute is Computed.
type ClusternodeDataSourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Backplane            types.String `tfsdk:"backplane"`
	Clearnodegroupconfig types.String `tfsdk:"clearnodegroupconfig"`
	Delay                types.Int64  `tfsdk:"delay"`
	Force                types.Bool   `tfsdk:"force"`
	Ipaddress            types.String `tfsdk:"ipaddress"`
	Nodegroup            types.String `tfsdk:"nodegroup"`
	Nodeid               types.Int64  `tfsdk:"nodeid"` // Required lookup key
	Priority             types.Int64  `tfsdk:"priority"`
	State                types.String `tfsdk:"state"`
	Tunnelmode           types.String `tfsdk:"tunnelmode"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/clusternode.json). Never settable; populated from GET.
	Clusterhealth              types.String `tfsdk:"clusterhealth"`
	Effectivestate             types.String `tfsdk:"effectivestate"`
	Operationalsyncstate       types.String `tfsdk:"operationalsyncstate"`
	Syncfailurereason          types.String `tfsdk:"syncfailurereason"`
	Masterstate                types.String `tfsdk:"masterstate"`
	Health                     types.String `tfsdk:"health"`
	Syncstate                  types.String `tfsdk:"syncstate"`
	Isconfigurationcoordinator types.Bool   `tfsdk:"isconfigurationcoordinator"`
	Islocalnode                types.Bool   `tfsdk:"islocalnode"`
	Nodersskeymismatch         types.Bool   `tfsdk:"nodersskeymismatch"`
	Nodelicensemismatch        types.Bool   `tfsdk:"nodelicensemismatch"`
	Nodejumbonotsupported      types.Bool   `tfsdk:"nodejumbonotsupported"`
	Nodelist                   types.List   `tfsdk:"nodelist"`
	Ifaceslist                 types.List   `tfsdk:"ifaceslist"`
	Enabledifaces              types.String `tfsdk:"enabledifaces"`
	Disabledifaces             types.String `tfsdk:"disabledifaces"`
	Partialfailifaces          types.String `tfsdk:"partialfailifaces"`
	Hamonifaces                types.String `tfsdk:"hamonifaces"`
	Name                       types.String `tfsdk:"name"`
	Cfgflags                   types.Int64  `tfsdk:"cfgflags"`
	Routemonitor               types.String `tfsdk:"routemonitor"`
	Netmask                    types.String `tfsdk:"netmask"`
}

func ClusternodeDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"backplane": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Interface through which the node communicates with the other nodes in the cluster. Must be specified in the three-tuple form n/c/u, where n represents the node ID and c/u refers to the interface on the appliance.",
			},
			"clearnodegroupconfig": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to remove nodegroup config",
			},
			"delay": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Applicable for Passive node and node becomes passive after this timeout (in minutes)",
			},
			"force": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Node will be removed from cluster without prompting for user confirmation.",
			},
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Citrix ADC IP (NSIP) address of the appliance to add to the cluster. Must be an IPv4 address.",
			},
			"nodegroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The default node group in a Cluster system.",
			},
			"nodeid": schema.Int64Attribute{
				Required:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Preference for selecting a node as the configuration coordinator. The node with the lowest priority value is selected as the configuration coordinator.\nWhen the current configuration coordinator goes down, the node with the next lowest priority is made the new configuration coordinator. When the original node comes back up, it will preempt the new configuration coordinator and take over as the configuration coordinator.\nNote: When priority is not configured for any of the nodes or if multiple nodes have the same priority, the cluster elects one of the nodes as the configuration coordinator.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Admin state of the cluster node. The available settings function as follows:\nACTIVE - The node serves traffic.\nSPARE - The node does not serve traffic unless an ACTIVE node goes down.\nPASSIVE - The node does not serve traffic, unless you change its state. PASSIVE state is useful during temporary maintenance activities in which you want the node to take part in the consensus protocol but not to serve traffic.",
			},
			"tunnelmode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "To set the tunnel mode",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"clusterhealth": schema.StringAttribute{
				Computed:    true,
				Description: "Node clusterd state.",
			},
			"effectivestate": schema.StringAttribute{
				Computed:    true,
				Description: "Node effective health state.",
			},
			"operationalsyncstate": schema.StringAttribute{
				Computed:    true,
				Description: "Node Operational Reconciliation state.",
			},
			"syncfailurereason": schema.StringAttribute{
				Computed:    true,
				Description: "Displays the additional information along with cluster sync status.",
			},
			"masterstate": schema.StringAttribute{
				Computed:    true,
				Description: "Node Master state.",
			},
			"health": schema.StringAttribute{
				Computed:    true,
				Description: "Node Health state.",
			},
			"syncstate": schema.StringAttribute{
				Computed:    true,
				Description: "Enable/Disable the synchronization of cluster configurations on the node.",
			},
			"isconfigurationcoordinator": schema.BoolAttribute{
				Computed:    true,
				Description: "This argument is used to determine whether the node is configuration coordinator (CCO).",
			},
			"islocalnode": schema.BoolAttribute{
				Computed:    true,
				Description: "This argument is used to determine whether it is local node.",
			},
			"nodersskeymismatch": schema.BoolAttribute{
				Computed:    true,
				Description: "This argument is used to determine if there is a RSS key mismatch at cluster node level.",
			},
			"nodelicensemismatch": schema.BoolAttribute{
				Computed:    true,
				Description: "This argument is used to determine if there is a License mismatch at cluster node level.",
			},
			"nodejumbonotsupported": schema.BoolAttribute{
				Computed:    true,
				Description: "This argument is used to determine if Jumbo framework not supported at cluster node level.",
			},
			"nodelist": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Nodelist for displaying Heartbeat not seen interfaces on a cluster node.",
			},
			"ifaceslist": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Interface list corresponding to nodelist for Heartbeat not seen interfaces on a cluster node.",
			},
			"enabledifaces": schema.StringAttribute{
				Computed:    true,
				Description: "Enabled Interfaces on a cluster node.",
			},
			"disabledifaces": schema.StringAttribute{
				Computed:    true,
				Description: "Disabled Interfaces on a cluster node.",
			},
			"partialfailifaces": schema.StringAttribute{
				Computed:    true,
				Description: "Partial Failure Interfaces on a cluster node.",
			},
			"hamonifaces": schema.StringAttribute{
				Computed:    true,
				Description: "Hamon Interfaces on a cluster node.",
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the state specific nodegroup.",
			},
			"cfgflags": schema.Int64Attribute{
				Computed:    true,
				Description: "Flag indicates whether the node is bound to cluster nodegroup.",
			},
			"routemonitor": schema.StringAttribute{
				Computed:    true,
				Description: "The IP address (IPv4 or IPv6).",
			},
			"netmask": schema.StringAttribute{
				Computed:    true,
				Description: "The netmask.",
			},
		},
	}
}

// clusternodeDataSourceSetAttrFromGet projects a NITRO clusternode GET response
// onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func clusternodeDataSourceSetAttrFromGet(ctx context.Context, data *ClusternodeDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In clusternodeDataSourceSetAttrFromGet Function")

	data.Nodeid = utils.MapGetInt64(g, "nodeid")
	if v, ok := g["nodeid"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Backplane = utils.MapGetString(g, "backplane")
	data.Clearnodegroupconfig = utils.MapGetString(g, "clearnodegroupconfig")
	data.Delay = utils.MapGetInt64(g, "delay")
	data.Force = utils.MapGetBool(g, "force")
	data.Ipaddress = utils.MapGetString(g, "ipaddress")
	data.Nodegroup = utils.MapGetString(g, "nodegroup")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.State = utils.MapGetString(g, "state")
	data.Tunnelmode = utils.MapGetString(g, "tunnelmode")

	// Read-only attributes.
	data.Clusterhealth = utils.MapGetString(g, "clusterhealth")
	data.Effectivestate = utils.MapGetString(g, "effectivestate")
	data.Operationalsyncstate = utils.MapGetString(g, "operationalsyncstate")
	data.Syncfailurereason = utils.MapGetString(g, "syncfailurereason")
	data.Masterstate = utils.MapGetString(g, "masterstate")
	data.Health = utils.MapGetString(g, "health")
	data.Syncstate = utils.MapGetString(g, "syncstate")
	data.Isconfigurationcoordinator = utils.MapGetBool(g, "isconfigurationcoordinator")
	data.Islocalnode = utils.MapGetBool(g, "islocalnode")
	data.Nodersskeymismatch = utils.MapGetBool(g, "nodersskeymismatch")
	data.Nodelicensemismatch = utils.MapGetBool(g, "nodelicensemismatch")
	data.Nodejumbonotsupported = utils.MapGetBool(g, "nodejumbonotsupported")
	data.Nodelist = utils.MapGetStringList(g, "nodelist")
	data.Ifaceslist = utils.MapGetStringList(g, "ifaceslist")
	data.Enabledifaces = utils.MapGetString(g, "enabledifaces")
	data.Disabledifaces = utils.MapGetString(g, "disabledifaces")
	data.Partialfailifaces = utils.MapGetString(g, "partialfailifaces")
	data.Hamonifaces = utils.MapGetString(g, "hamonifaces")
	data.Name = utils.MapGetString(g, "name")
	data.Cfgflags = utils.MapGetInt64(g, "cfgflags")
	data.Routemonitor = utils.MapGetString(g, "routemonitor")
	data.Netmask = utils.MapGetString(g, "netmask")
}
