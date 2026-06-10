package lsnsession

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LsnsessionResourceModel describes the resource data model.
type LsnsessionResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Clientname types.String `tfsdk:"clientname"`
	Natip      types.String `tfsdk:"natip"`
	Natport2   types.Int64  `tfsdk:"natport2"`
	Nattype    types.String `tfsdk:"nattype"`
	Netmask    types.String `tfsdk:"netmask"`
	Network    types.String `tfsdk:"network"`
	Network6   types.String `tfsdk:"network6"`
	Nodeid     types.Int64  `tfsdk:"nodeid"`
	Td         types.Int64  `tfsdk:"td"`
}

func (r *LsnsessionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lsnsession resource.",
			},
			"clientname": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the LSN Client entity.",
			},
			"natip": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Mapped NAT IP address used in LSN sessions.",
			},
			"natport2": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Mapped NAT port used in the LSN sessions.",
			},
			"nattype": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of sessions to be flushed. If omitted, NITRO applies its server-side default of NAT44.",
			},
			"netmask": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Subnet mask for the IP address specified by the network parameter. Must be supplied together with network.",
			},
			"network": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IP address or network address of subscriber(s).",
			},
			"network6": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IPv6 address of the LSN subscriber or B4 device.",
			},
			"nodeid": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Unique number that identifies the cluster node.",
			},
			"td": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Traffic domain ID of the LSN client entity.",
			},
		},
	}
}

func lsnsessionGetThePayloadFromthePlan(ctx context.Context, data *LsnsessionResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In lsnsessionGetThePayloadFromthePlan Function")

	// Build the flush payload from whichever filter selectors are set.
	lsnsession := make(map[string]interface{})
	if !data.Clientname.IsNull() && !data.Clientname.IsUnknown() {
		lsnsession["clientname"] = data.Clientname.ValueString()
	}
	if !data.Natip.IsNull() && !data.Natip.IsUnknown() {
		lsnsession["natip"] = data.Natip.ValueString()
	}
	if !data.Natport2.IsNull() && !data.Natport2.IsUnknown() {
		lsnsession["natport2"] = int(data.Natport2.ValueInt64())
	}
	if !data.Nattype.IsNull() && !data.Nattype.IsUnknown() {
		lsnsession["nattype"] = data.Nattype.ValueString()
	}
	if !data.Netmask.IsNull() && !data.Netmask.IsUnknown() {
		lsnsession["netmask"] = data.Netmask.ValueString()
	}
	if !data.Network.IsNull() && !data.Network.IsUnknown() {
		lsnsession["network"] = data.Network.ValueString()
	}
	if !data.Network6.IsNull() && !data.Network6.IsUnknown() {
		lsnsession["network6"] = data.Network6.ValueString()
	}
	if !data.Nodeid.IsNull() && !data.Nodeid.IsUnknown() {
		lsnsession["nodeid"] = int(data.Nodeid.ValueInt64())
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		lsnsession["td"] = int(data.Td.ValueInt64())
	}

	return lsnsession
}
