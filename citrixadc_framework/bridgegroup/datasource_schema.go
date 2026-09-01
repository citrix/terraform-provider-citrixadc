package bridgegroup

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// BridgegroupDataSourceModel is the data-source-specific model, decoupled from
// BridgegroupResourceModel. A data source is a pure read surface, so it can
// expose the full GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits.
type BridgegroupDataSourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Dynamicrouting     types.String `tfsdk:"dynamicrouting"`
	Bridgegroupid      types.Int64  `tfsdk:"bridgegroup_id"`
	Ipv6dynamicrouting types.String `tfsdk:"ipv6dynamicrouting"`

	// Read-only (GET-only) metadata from the NITRO read-only set
	// (zion73x_readonly/bridgegroup.json). Never settable; populated from GET.
	Flags         types.Bool   `tfsdk:"flags"`
	Portbitmap    types.Int64  `tfsdk:"portbitmap"`
	Tagbitmap     types.Int64  `tfsdk:"tagbitmap"`
	Ifaces        types.String `tfsdk:"ifaces"`
	Tagifaces     types.String `tfsdk:"tagifaces"`
	Rnat          types.Bool   `tfsdk:"rnat"`
	Partitionname types.String `tfsdk:"partitionname"`
}

func BridgegroupDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"dynamicrouting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable dynamic routing for this bridgegroup.",
			},
			"bridgegroup_id": schema.Int64Attribute{
				Required:    true,
				Description: "An integer that uniquely identifies the bridge group.",
			},
			"ipv6dynamicrouting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable all IPv6 dynamic routing protocols on all VLANs bound to this bridgegroup. Note: For the ENABLED setting to work, you must configure IPv6 dynamic routing protocols from the VTYSH command line.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"flags": schema.BoolAttribute{
				Computed:    true,
				Description: "Temporary flag used for internal purpose.",
			},
			"portbitmap": schema.Int64Attribute{
				Computed:    true,
				Description: "Member interfaces of this bridge group.",
			},
			"tagbitmap": schema.Int64Attribute{
				Computed:    true,
				Description: "Tagged members of this bridge group.",
			},
			"ifaces": schema.StringAttribute{
				Computed:    true,
				Description: "Names of all member interfaces of this bridge group.",
			},
			"tagifaces": schema.StringAttribute{
				Computed:    true,
				Description: "Names of all tagged member interfaces of this bridge group.",
			},
			"rnat": schema.BoolAttribute{
				Computed:    true,
				Description: "Temporary flag used for internal purpose.",
			},
			"partitionname": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the Partition to which this vlan bound to.",
			},
		},
	}
}

// bridgegroupDataSourceSetAttrFromGet projects a NITRO bridgegroup GET response
// onto the data-source model. Attributes are filled from the GET (or left Null
// when the GET omits them) via the shared utils.MapGet* helpers.
func bridgegroupDataSourceSetAttrFromGet(ctx context.Context, data *BridgegroupDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In bridgegroupDataSourceSetAttrFromGet Function")

	if v, ok := g["id"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Bridgegroupid = utils.MapGetInt64(g, "id")
	data.Dynamicrouting = utils.MapGetString(g, "dynamicrouting")
	data.Ipv6dynamicrouting = utils.MapGetString(g, "ipv6dynamicrouting")

	// Read-only metadata.
	data.Flags = utils.MapGetBool(g, "flags")
	data.Portbitmap = utils.MapGetInt64(g, "portbitmap")
	data.Tagbitmap = utils.MapGetInt64(g, "tagbitmap")
	data.Ifaces = utils.MapGetString(g, "ifaces")
	data.Tagifaces = utils.MapGetString(g, "tagifaces")
	data.Rnat = utils.MapGetBool(g, "rnat")
	data.Partitionname = utils.MapGetString(g, "partitionname")
}
