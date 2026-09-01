package lacp

import (
	"context"
	"fmt"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LacpDataSourceModel is the data-source-specific model, decoupled from
// LacpResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits
// (devicename, mac, flags, lacpkey, clustersyspriority, clustermac). Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type LacpDataSourceModel struct {
	Id          types.String `tfsdk:"id"`
	Ownernode   types.Int64  `tfsdk:"ownernode"`
	Syspriority types.Int64  `tfsdk:"syspriority"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/lacp.json). Never settable; populated from GET.
	Devicename         types.String `tfsdk:"devicename"`
	Mac                types.String `tfsdk:"mac"`
	Flags              types.Int64  `tfsdk:"flags"`
	Lacpkey            types.Int64  `tfsdk:"lacpkey"`
	Clustersyspriority types.Int64  `tfsdk:"clustersyspriority"`
	Clustermac         types.String `tfsdk:"clustermac"`
}

func LacpDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ownernode": schema.Int64Attribute{
				Required:    true,
				Description: "The owner node in a cluster for which we want to set the lacp priority. Owner node can vary from 0 to 31. Ownernode value of 254 is used for Cluster.",
			},
			"syspriority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority number that determines which peer device of an LACP LA channel can have control over the LA channel. This parameter is globally applied to all LACP channels on the Citrix ADC. The lower the number, the higher the priority.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"devicename": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the channel.",
			},
			"mac": schema.StringAttribute{
				Computed:    true,
				Description: "LACP system MAC.",
			},
			"flags": schema.Int64Attribute{
				Computed:    true,
				Description: "Flags of this channel.",
			},
			"lacpkey": schema.Int64Attribute{
				Computed:    true,
				Description: "LACP key of this channel.",
			},
			"clustersyspriority": schema.Int64Attribute{
				Computed:    true,
				Description: "LACP system (Cluster) priority.",
			},
			"clustermac": schema.StringAttribute{
				Computed:    true,
				Description: "LACP system (Cluster) mac.",
			},
		},
	}
}

// lacpDataSourceSetAttrFromGet projects a NITRO lacp GET response onto the
// data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func lacpDataSourceSetAttrFromGet(ctx context.Context, data *LacpDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In lacpDataSourceSetAttrFromGet Function")

	// ownernode is the config-supplied lookup key; keep it (and use it as the id)
	// even when the GET does not echo it back.
	if v := utils.MapGetInt64(g, "ownernode"); !v.IsNull() {
		data.Ownernode = v
	}
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Ownernode.ValueInt64()))

	// Read/write attributes as read-back outputs.
	data.Syspriority = utils.MapGetInt64(g, "syspriority")

	// Read-only metadata.
	data.Devicename = utils.MapGetString(g, "devicename")
	data.Mac = utils.MapGetString(g, "mac")
	data.Flags = utils.MapGetInt64(g, "flags")
	data.Lacpkey = utils.MapGetInt64(g, "lacpkey")
	data.Clustersyspriority = utils.MapGetInt64(g, "clustersyspriority")
	data.Clustermac = utils.MapGetString(g, "clustermac")
}
