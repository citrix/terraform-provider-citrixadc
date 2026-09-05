package clusterinstance

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ClusterinstanceDataSourceModel is the data-source-specific model, decoupled
// from ClusterinstanceResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only runtime/status attributes that the resource
// deliberately omits (adminstate, status, mismatch flags, ...). Every non-key
// attribute is Computed.
type ClusterinstanceDataSourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Backplanebasedview         types.String `tfsdk:"backplanebasedview"`
	Clid                       types.Int64  `tfsdk:"clid"` // Required lookup key
	Clusterproxyarp            types.String `tfsdk:"clusterproxyarp"`
	Deadinterval               types.Int64  `tfsdk:"deadinterval"`
	Dfdretainl2params          types.String `tfsdk:"dfdretainl2params"`
	Hellointerval              types.Int64  `tfsdk:"hellointerval"`
	Inc                        types.String `tfsdk:"inc"`
	Nodegroup                  types.String `tfsdk:"nodegroup"`
	Preemption                 types.String `tfsdk:"preemption"`
	Processlocal               types.String `tfsdk:"processlocal"`
	Quorumtype                 types.String `tfsdk:"quorumtype"`
	Retainconnectionsoncluster types.String `tfsdk:"retainconnectionsoncluster"`
	Secureheartbeats           types.String `tfsdk:"secureheartbeats"`
	Syncstatusstrictmode       types.String `tfsdk:"syncstatusstrictmode"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/clusterinstance.json). Never settable; populated from GET.
	Adminstate                 types.String `tfsdk:"adminstate"`
	Propstate                  types.String `tfsdk:"propstate"`
	Validmtu                   types.Int64  `tfsdk:"validmtu"`
	Heterogeneousflag          types.String `tfsdk:"heterogeneousflag"`
	Operationalstate           types.String `tfsdk:"operationalstate"`
	Status                     types.String `tfsdk:"status"`
	Rsskeymismatch             types.Bool   `tfsdk:"rsskeymismatch"`
	Penummismatch              types.Bool   `tfsdk:"penummismatch"`
	Nodegroupstatewarning      types.Bool   `tfsdk:"nodegroupstatewarning"`
	Licensemismatch            types.Bool   `tfsdk:"licensemismatch"`
	Jumbonotsupported          types.Bool   `tfsdk:"jumbonotsupported"`
	Clustertunnelmodemismatch  types.Bool   `tfsdk:"clustertunnelmodemismatch"`
	Clusternoheartbeatonnode   types.Bool   `tfsdk:"clusternoheartbeatonnode"`
	Clusternolinksetmbf        types.Bool   `tfsdk:"clusternolinksetmbf"`
	Clusternospottedip         types.Bool   `tfsdk:"clusternospottedip"`
	Clusterclipfailure         types.Bool   `tfsdk:"clusterclipfailure"`
	Clusterhbhmacerrordetected types.Bool   `tfsdk:"clusterhbhmacerrordetected"`
	Nodepenummismatch          types.Bool   `tfsdk:"nodepenummismatch"`
	Operationalpropstate       types.String `tfsdk:"operationalpropstate"`
}

func ClusterinstanceDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"backplanebasedview": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "View based on heartbeat only on bkplane interface",
			},
			"clid": schema.Int64Attribute{
				Required:    true,
				Description: "Unique number that identifies the cluster.",
			},
			"clusterproxyarp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This field controls the proxy arp feature in cluster. By default the flag is enabled.",
			},
			"deadinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Amount of time, in seconds, after which nodes that do not respond to the heartbeats are assumed to be down.If the value is less than 3 sec, set the helloInterval parameter to 200 msec",
			},
			"dfdretainl2params": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "flag to add ext l2 header during steering. By default the flag is disabled.",
			},
			"hellointerval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Interval, in milliseconds, at which heartbeats are sent to each cluster node to check the health status.Set the value to 200 msec, if the deadInterval parameter is less than 3 sec",
			},
			"inc": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option is required if the cluster nodes reside on different networks.",
			},
			"nodegroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The node group in a Cluster system used for transition from L2 to L3.",
			},
			"preemption": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Preempt a cluster node that is configured as a SPARE if an ACTIVE node becomes available.",
			},
			"processlocal": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "By turning on this option packets destined to a service in a cluster will not under go any steering.",
			},
			"quorumtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Quorum Configuration Choices  - \"Majority\" (recommended) requires majority of nodes to be online for the cluster to be UP. \"None\" relaxes this requirement.",
			},
			"retainconnectionsoncluster": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option enables you to retain existing connections on a node joining a Cluster system or when a node is being configured for passive timeout. By default, this option is disabled.",
			},
			"secureheartbeats": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "By turning on this option cluster heartbeats will have security enabled.",
			},
			"syncstatusstrictmode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "strict mode for sync status of cluster. Depending on the the mode if there are any errors while applying config, sync status is displayed accordingly. By default the flag is disabled.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"adminstate": schema.StringAttribute{
				Computed:    true,
				Description: "Cluster Admin State.",
			},
			"propstate": schema.StringAttribute{
				Computed:    true,
				Description: "Enable/Disable the execution of commands on the cluster. This will not impact the execution of commands on individual cluster nodes by using the NSIP.",
			},
			"validmtu": schema.Int64Attribute{
				Computed:    true,
				Description: "Correct MTU value that has to be set on backplane.",
			},
			"heterogeneousflag": schema.StringAttribute{
				Computed:    true,
				Description: "This field diplay if heterogeneous is detected in the cluster system.",
			},
			"operationalstate": schema.StringAttribute{
				Computed:    true,
				Description: "Cluster Operational State.",
			},
			"status": schema.StringAttribute{
				Computed:    true,
				Description: "Cluster Operational State.",
			},
			"rsskeymismatch": schema.BoolAttribute{
				Computed:    true,
				Description: "This argument is used to determine if there is a RSS key mismatch at cluster instance level.",
			},
			"penummismatch": schema.BoolAttribute{
				Computed:    true,
				Description: "This argument is used to determine if there is a PE number mismatch at cluster instance level.",
			},
			"nodegroupstatewarning": schema.BoolAttribute{
				Computed:    true,
				Description: "This argument is used to determine whether all the cluster nodes are bound to nodegroup with state set.",
			},
			"licensemismatch": schema.BoolAttribute{
				Computed:    true,
				Description: "This argument is used to determine if there is a License mismatch at cluster instance level.",
			},
			"jumbonotsupported": schema.BoolAttribute{
				Computed:    true,
				Description: "This argument is used to determine if Jumbo framework is not supported at cluster instance level.",
			},
			"clustertunnelmodemismatch": schema.BoolAttribute{
				Computed:    true,
				Description: "This argument is used to determine if different tunnel mode configured on cluster nodes.",
			},
			"clusternoheartbeatonnode": schema.BoolAttribute{
				Computed:    true,
				Description: "HB is not seen on the backplane interface of member node.",
			},
			"clusternolinksetmbf": schema.BoolAttribute{
				Computed:    true,
				Description: "MBF is enabled but linkset is not configured.",
			},
			"clusternospottedip": schema.BoolAttribute{
				Computed:    true,
				Description: "There is no spotted SNIP or MIP.",
			},
			"clusterclipfailure": schema.BoolAttribute{
				Computed:    true,
				Description: "CLIP movement failure. CLIP is not attached to CCO.",
			},
			"clusterhbhmacerrordetected": schema.BoolAttribute{
				Computed:    true,
				Description: "There is cluster hb hmac error detected, it could be due to version mismatch.",
			},
			"nodepenummismatch": schema.BoolAttribute{
				Computed:    true,
				Description: "This argument is used to determine if there is a PE mismatch at cluster node level.",
			},
			"operationalpropstate": schema.StringAttribute{
				Computed:    true,
				Description: "Cluster Operational Propagation State.",
			},
		},
	}
}

// clusterinstanceDataSourceSetAttrFromGet projects a NITRO clusterinstance GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func clusterinstanceDataSourceSetAttrFromGet(ctx context.Context, data *ClusterinstanceDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In clusterinstanceDataSourceSetAttrFromGet Function")

	data.Clid = utils.MapGetInt64(g, "clid")
	if v, ok := g["clid"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Backplanebasedview = utils.MapGetString(g, "backplanebasedview")
	data.Clusterproxyarp = utils.MapGetString(g, "clusterproxyarp")
	data.Deadinterval = utils.MapGetInt64(g, "deadinterval")
	data.Dfdretainl2params = utils.MapGetString(g, "dfdretainl2params")
	data.Hellointerval = utils.MapGetInt64(g, "hellointerval")
	data.Inc = utils.MapGetString(g, "inc")
	data.Nodegroup = utils.MapGetString(g, "nodegroup")
	data.Preemption = utils.MapGetString(g, "preemption")
	data.Processlocal = utils.MapGetString(g, "processlocal")
	data.Quorumtype = utils.MapGetString(g, "quorumtype")
	data.Retainconnectionsoncluster = utils.MapGetString(g, "retainconnectionsoncluster")
	data.Secureheartbeats = utils.MapGetString(g, "secureheartbeats")
	data.Syncstatusstrictmode = utils.MapGetString(g, "syncstatusstrictmode")

	// Read-only attributes.
	data.Adminstate = utils.MapGetString(g, "adminstate")
	data.Propstate = utils.MapGetString(g, "propstate")
	data.Validmtu = utils.MapGetInt64(g, "validmtu")
	data.Heterogeneousflag = utils.MapGetString(g, "heterogeneousflag")
	data.Operationalstate = utils.MapGetString(g, "operationalstate")
	data.Status = utils.MapGetString(g, "status")
	data.Rsskeymismatch = utils.MapGetBool(g, "rsskeymismatch")
	data.Penummismatch = utils.MapGetBool(g, "penummismatch")
	data.Nodegroupstatewarning = utils.MapGetBool(g, "nodegroupstatewarning")
	data.Licensemismatch = utils.MapGetBool(g, "licensemismatch")
	data.Jumbonotsupported = utils.MapGetBool(g, "jumbonotsupported")
	data.Clustertunnelmodemismatch = utils.MapGetBool(g, "clustertunnelmodemismatch")
	data.Clusternoheartbeatonnode = utils.MapGetBool(g, "clusternoheartbeatonnode")
	data.Clusternolinksetmbf = utils.MapGetBool(g, "clusternolinksetmbf")
	data.Clusternospottedip = utils.MapGetBool(g, "clusternospottedip")
	data.Clusterclipfailure = utils.MapGetBool(g, "clusterclipfailure")
	data.Clusterhbhmacerrordetected = utils.MapGetBool(g, "clusterhbhmacerrordetected")
	data.Nodepenummismatch = utils.MapGetBool(g, "nodepenummismatch")
	data.Operationalpropstate = utils.MapGetString(g, "operationalpropstate")
}
