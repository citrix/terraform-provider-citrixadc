package cluster

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ClusterResourceModel describes the resource data model. It mirrors the
// backward-compatible SDK v2 "citrixadc_cluster" schema: a high-level cluster
// bootstrap orchestrator whose identity is the cluster instance id (clid) and
// which manages the cluster instance, its member nodes (clusternode blocks) and
// (in L3/INC mode) node groups (clusternodegroup blocks).
type ClusterResourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Backplanebasedview         types.String `tfsdk:"backplanebasedview"`
	Clid                       types.Int64  `tfsdk:"clid"`
	Clip                       types.String `tfsdk:"clip"`
	Deadinterval               types.Int64  `tfsdk:"deadinterval"`
	Hellointerval              types.Int64  `tfsdk:"hellointerval"`
	Inc                        types.String `tfsdk:"inc"`
	Nodegroup                  types.String `tfsdk:"nodegroup"`
	Preemption                 types.String `tfsdk:"preemption"`
	Processlocal               types.String `tfsdk:"processlocal"`
	Quorumtype                 types.String `tfsdk:"quorumtype"`
	Retainconnectionsoncluster types.String `tfsdk:"retainconnectionsoncluster"`

	// Client-side control knobs (not NITRO attributes) with SDK v2 defaults.
	BootstrapPollDelay        types.String `tfsdk:"bootstrap_poll_delay"`
	BootstrapPollInterval     types.String `tfsdk:"bootstrap_poll_interval"`
	BootstrapPollTimeout      types.String `tfsdk:"bootstrap_poll_timeout"`
	BootstrapTotalTimeout     types.String `tfsdk:"bootstrap_total_timeout"`
	ClipMigrationPollDelay    types.String `tfsdk:"clip_migration_poll_delay"`
	ClipMigrationPollInterval types.String `tfsdk:"clip_migration_poll_interval"`
	ClipMigrationPollTimeout  types.String `tfsdk:"clip_migration_poll_timeout"`
	ClipMigrationTotalTimeout types.String `tfsdk:"clip_migration_total_timeout"`
	NodeAddPollDelay          types.String `tfsdk:"node_add_poll_delay"`
	NodeAddPollInterval       types.String `tfsdk:"node_add_poll_interval"`
	NodeAddTotalTimeout       types.String `tfsdk:"node_add_total_timeout"`

	Clusternodegroup types.Set `tfsdk:"clusternodegroup"`
	Clusternode      types.Set `tfsdk:"clusternode"`
}

// ClusternodegroupModel is a single clusternodegroup block element.
type ClusternodegroupModel struct {
	Name     types.String `tfsdk:"name"`
	Priority types.Int64  `tfsdk:"priority"`
	State    types.String `tfsdk:"state"`
	Sticky   types.String `tfsdk:"sticky"`
	Strict   types.String `tfsdk:"strict"`
}

// ClusternodeModel is a single clusternode block element.
type ClusternodeModel struct {
	Backplane            types.String `tfsdk:"backplane"`
	Clearnodegroupconfig types.String `tfsdk:"clearnodegroupconfig"`
	Delay                types.Int64  `tfsdk:"delay"`
	Ipaddress            types.String `tfsdk:"ipaddress"`
	Nodegroup            types.String `tfsdk:"nodegroup"`
	Nodeid               types.Int64  `tfsdk:"nodeid"`
	Priority             types.Int64  `tfsdk:"priority"`
	State                types.String `tfsdk:"state"`
	Tunnelmode           types.String `tfsdk:"tunnelmode"`
	Endpoint             types.String `tfsdk:"endpoint"`
	Username             types.String `tfsdk:"username"`
	Password             types.String `tfsdk:"password"`
	InsecureSkipVerify   types.Bool   `tfsdk:"insecure_skip_verify"`
	SnipNetmask          types.String `tfsdk:"snip_netmask"`
	SnipIpaddress        types.String `tfsdk:"snip_ipaddress"`
	Addsnip              types.Bool   `tfsdk:"addsnip"`
	VtyshEnable          types.Bool   `tfsdk:"vtysh_enable"`
	Vtysh                types.List   `tfsdk:"vtysh"`
}

func (r *ClusterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version:     1,
		Description: "Configuration for cluster resource. Bootstraps and manages a Citrix ADC cluster (cluster instance, member nodes and node groups). Applying this resource is disruptive.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cluster resource (the cluster instance id, clid).",
			},
			"backplanebasedview": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "View based on heartbeat only on bkplane interface.",
			},
			"clid": schema.Int64Attribute{
				Required:    true,
				Description: "Unique number that identifies the cluster.",
			},
			"clip": schema.StringAttribute{
				Required:    true,
				Description: "Cluster IP address.",
			},
			"deadinterval": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Description: "Amount of time, in seconds, after which nodes that do not respond to the heartbeats are assumed to be down.",
			},
			"hellointerval": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Description: "Interval, in milliseconds, at which heartbeats are sent to each cluster node to check the health status.",
			},
			"inc": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "This option is required if the cluster nodes reside on different networks.",
			},
			"nodegroup": schema.StringAttribute{
				Optional:    true,
				Description: "The default node group in a Cluster system.",
			},
			"preemption": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "Preempt a cluster node that is configured as a SPARE if an ACTIVE node becomes available.",
			},
			"processlocal": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "By turning on this option packets destined to a service in a cluster will not undergo any steering.",
			},
			"quorumtype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "Quorum Configuration Choices - Majority (recommended) or None.",
			},
			"retainconnectionsoncluster": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "Retain existing connections on a node joining a Cluster system or when a node is being configured for passive timeout.",
			},

			// Client-side control knobs with SDK v2 defaults.
			"bootstrap_poll_delay": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("60s"),
				Description: "How long to wait before the first bootstrap poll.",
			},
			"bootstrap_poll_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("60s"),
				Description: "Interval between bootstrap polls.",
			},
			"bootstrap_poll_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("10s"),
				Description: "Timeout for each individual bootstrap HTTP poll.",
			},
			"bootstrap_total_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("10m"),
				Description: "Timeout for the whole bootstrap operation.",
			},
			"clip_migration_poll_delay": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("10s"),
				Description: "Delay before the first CLIP-migration poll.",
			},
			"clip_migration_poll_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("30s"),
				Description: "Interval between CLIP-migration polls.",
			},
			"clip_migration_poll_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("10s"),
				Description: "Timeout for each individual CLIP-migration poll HTTP request.",
			},
			"clip_migration_total_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("10m"),
				Description: "Timeout for the whole CLIP-migration operation.",
			},
			"node_add_poll_delay": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("10s"),
				Description: "Delay before the first node-add poll.",
			},
			"node_add_poll_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("30s"),
				Description: "Interval between node-add polls.",
			},
			"node_add_total_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("10m"),
				Description: "Timeout for the whole node-add operation.",
			},
		},
		Blocks: map[string]schema.Block{
			"clusternodegroup": schema.SetNestedBlock{
				Description: "Cluster node groups (used in L3/INC cluster mode).",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Optional:    true,
							Description: "Name of the nodegroup.",
						},
						"priority": schema.Int64Attribute{
							Optional:    true,
							Description: "Priority of the nodegroup.",
						},
						"state": schema.StringAttribute{
							Optional:    true,
							Description: "State of the nodegroup.",
						},
						"sticky": schema.StringAttribute{
							Optional:    true,
							Description: "Only one node can be bound to nodegroup with this option enabled.",
						},
						"strict": schema.StringAttribute{
							Optional:    true,
							Description: "Whether cluster nodes not part of the nodegroup are used as backup.",
						},
					},
				},
			},
			"clusternode": schema.SetNestedBlock{
				Description: "Cluster member nodes. At least one node is required.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"backplane": schema.StringAttribute{
							Optional:    true,
							Description: "Interface through which the node communicates with the other nodes (n/c/u).",
						},
						"clearnodegroupconfig": schema.StringAttribute{
							Optional:    true,
							Description: "Option to remove nodegroup config.",
						},
						"delay": schema.Int64Attribute{
							Optional:    true,
							Description: "Passive node becomes passive after this timeout (in minutes).",
						},
						"ipaddress": schema.StringAttribute{
							Required:    true,
							Description: "Citrix ADC IP (NSIP) address of the appliance to add to the cluster.",
						},
						"nodegroup": schema.StringAttribute{
							Optional:    true,
							Description: "The default node group in a Cluster system.",
						},
						"nodeid": schema.Int64Attribute{
							Required:    true,
							Description: "Unique number that identifies the cluster node.",
						},
						"priority": schema.Int64Attribute{
							Optional:    true,
							Description: "Preference for selecting a node as the configuration coordinator.",
						},
						"state": schema.StringAttribute{
							Optional:    true,
							Description: "Admin state of the cluster node (ACTIVE/SPARE/PASSIVE).",
						},
						"tunnelmode": schema.StringAttribute{
							Optional:    true,
							Description: "Tunnel mode.",
						},
						"endpoint": schema.StringAttribute{
							Required:    true,
							Description: "NITRO endpoint (URL) of the node, used while bootstrapping/joining it to the cluster.",
						},
						"username": schema.StringAttribute{
							Optional:    true,
							Sensitive:   true,
							Description: "Username for the node NITRO endpoint. Defaults to the provider username.",
						},
						"password": schema.StringAttribute{
							Optional:    true,
							Sensitive:   true,
							Description: "Password for the node NITRO endpoint. Defaults to the provider password.",
						},
						"insecure_skip_verify": schema.BoolAttribute{
							Optional:    true,
							Description: "Ignore validity of the endpoint TLS certificate if true.",
						},
						"snip_netmask": schema.StringAttribute{
							Optional:    true,
							Description: "Netmask of the SNIP to add to the CLIP before joining.",
						},
						"snip_ipaddress": schema.StringAttribute{
							Optional:    true,
							Description: "SNIP address to add to the CLIP before joining.",
						},
						"addsnip": schema.BoolAttribute{
							Optional:    true,
							Description: "Add the node SNIP to the CLIP before joining.",
						},
						"vtysh_enable": schema.BoolAttribute{
							Optional:    true,
							Description: "Run the vtysh commands while adding the node.",
						},
						"vtysh": schema.ListAttribute{
							Optional:    true,
							ElementType: types.StringType,
							Description: "Optional VTYSH commands to apply while adding the node.",
						},
					},
				},
			},
		},
	}
}
