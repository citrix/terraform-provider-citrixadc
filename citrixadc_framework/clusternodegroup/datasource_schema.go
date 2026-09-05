package clusternodegroup

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ClusternodegroupDataSourceModel is the data-source-specific model, decoupled
// from ClusternodegroupResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only runtime attributes that the resource deliberately
// omits (currentnodemask, activelist, ...). Every non-key attribute is Computed.
type ClusternodegroupDataSourceModel struct {
	Id       types.String `tfsdk:"id"`
	Name     types.String `tfsdk:"name"` // Required lookup key
	Priority types.Int64  `tfsdk:"priority"`
	State    types.String `tfsdk:"state"`
	Sticky   types.String `tfsdk:"sticky"`
	Strict   types.String `tfsdk:"strict"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/clusternodegroup.json). Never settable; populated from GET.
	Currentnodemask          types.Int64 `tfsdk:"currentnodemask"`
	Backupnodemask           types.Int64 `tfsdk:"backupnodemask"`
	Boundedentitiescntfrompe types.Int64 `tfsdk:"boundedentitiescntfrompe"`
	Activelist               types.List  `tfsdk:"activelist"`
	Backuplist               types.List  `tfsdk:"backuplist"`
}

func ClusternodegroupDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority of Nodegroup. This priority is used for all the nodes bound to the nodegroup for Nodegroup selection",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of the nodegroup. All the nodes binding to this nodegroup must have the same state. ACTIVE/SPARE/PASSIVE",
			},
			"sticky": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Only one node can be bound to nodegroup with this option enabled. It specifies whether to prempt the traffic for the entities bound to nodegroup when owner node goes down and rejoins the cluster.\n  * Enabled - When owner node goes down, backup node will become the owner node and takes the traffic for the entities bound to the nodegroup. When bound node rejoins the cluster, traffic for the entities bound to nodegroup will not be steered back to this bound node. Current owner will have the ownership till it goes down.\n  * Disabled - When one of the nodes goes down, a non-nodegroup cluster node is picked up and acts as part of the nodegroup. When the original node of the nodegroup comes up, the backup node will be replaced.",
			},
			"strict": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies whether cluster nodes, that are not part of the nodegroup, will be used as backup for the nodegroup.\n  * Enabled - When one of the nodes goes down, no other cluster node is picked up to replace it. When the node comes up, it will continue being part of the nodegroup.\n  * Disabled - When one of the nodes goes down, a non-nodegroup cluster node is picked up and acts as part of the nodegroup. When the original node of the nodegroup comes up, the backup node will be replaced.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"currentnodemask": schema.Int64Attribute{
				Computed:    true,
				Description: "Bitmap of current nodes in this nodegroup.",
			},
			"backupnodemask": schema.Int64Attribute{
				Computed:    true,
				Description: "Bitmap of backup nodes in this nodegroup.",
			},
			"boundedentitiescntfrompe": schema.Int64Attribute{
				Computed:    true,
				Description: "Count of bounded entities to this nodegroup accoding to PE.",
			},
			"activelist": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Active node list of this nodegroup.",
			},
			"backuplist": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Backup node list of this nodegroup.",
			},
		},
	}
}

// clusternodegroupDataSourceSetAttrFromGet projects a NITRO clusternodegroup GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func clusternodegroupDataSourceSetAttrFromGet(ctx context.Context, data *ClusternodegroupDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In clusternodegroupDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Priority = utils.MapGetInt64(g, "priority")
	data.State = utils.MapGetString(g, "state")
	data.Sticky = utils.MapGetString(g, "sticky")
	data.Strict = utils.MapGetString(g, "strict")

	// Read-only attributes.
	data.Currentnodemask = utils.MapGetInt64(g, "currentnodemask")
	data.Backupnodemask = utils.MapGetInt64(g, "backupnodemask")
	data.Boundedentitiescntfrompe = utils.MapGetInt64(g, "boundedentitiescntfrompe")
	data.Activelist = utils.MapGetStringList(g, "activelist")
	data.Backuplist = utils.MapGetStringList(g, "backuplist")
}
