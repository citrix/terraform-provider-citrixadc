package vpnglobal_staserver_binding

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnglobalStaserverBindingDataSourceModel is the data-source-specific model,
// decoupled from the resource model. A data source is a pure read surface (Read
// only; no plan/apply lifecycle), so it can expose the full GET projection: the
// read/write attributes (as Computed outputs) AND the read-only attributes that
// the resource deliberately omits (staauthid, stastate).
type VpnglobalStaserverBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Staaddresstype         types.String `tfsdk:"staaddresstype"`
	Staserver              types.String `tfsdk:"staserver"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/vpnglobal_staserver_binding.json).
	Staauthid types.String `tfsdk:"staauthid"`
	Stastate  types.String `tfsdk:"stastate"`
}

func VpnglobalStaserverBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"staaddresstype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of the STA server address(ipv4/v6).",
			},
			"staserver": schema.StringAttribute{
				Required:    true,
				Description: "Configured Secure Ticketing Authority (STA) server.",
			},

			// Read-only (GET-only) metadata surfaced by the data source. All Computed.
			"staauthid": schema.StringAttribute{
				Computed:    true,
				Description: "Authority ID of the STA Server. Authority ID is used to match incoming STA Tickets in the SOCKS/CGP protocol with the right STA Server.",
			},
			"stastate": schema.StringAttribute{
				Computed:    true,
				Description: "State of the STA Server. If Authority ID is set then STA Server is UP else DOWN. Possible values: [ UP, DOWN ]",
			},
		},
	}
}

// vpnglobal_staserver_bindingDataSourceSetAttrFromGet projects a NITRO GET
// response onto the data-source model using the shared utils.MapGet* helpers.
func vpnglobal_staserver_bindingDataSourceSetAttrFromGet(ctx context.Context, data *VpnglobalStaserverBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In vpnglobal_staserver_bindingDataSourceSetAttrFromGet Function")

	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Staaddresstype = utils.MapGetString(g, "staaddresstype")
	data.Staserver = utils.MapGetString(g, "staserver")

	// Read-only (GET-only) metadata.
	data.Staauthid = utils.MapGetString(g, "staauthid")
	data.Stastate = utils.MapGetString(g, "stastate")

	// Set ID for the datasource (single unique attribute - plain staserver).
	data.Id = types.StringValue(data.Staserver.ValueString())
}
