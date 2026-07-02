package rnatsession

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// RnatsessionResourceModel describes the resource data model.
type RnatsessionResourceModel struct {
	Id      types.String `tfsdk:"id"`
	Aclname types.String `tfsdk:"aclname"`
	Natip   types.String `tfsdk:"natip"`
	Netmask types.String `tfsdk:"netmask"`
	Network types.String `tfsdk:"network"`
}

func (r *RnatsessionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the rnatsession resource.",
			},
			"aclname": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of any configured extended ACL whose action is ALLOW.",
			},
			"natip": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The NAT IP address defined for the RNAT entry.",
			},
			"netmask": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Subnet mask associated with the network address.",
			},
			"network": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IPv4 network address on whose traffic you want the Citrix ADC to do RNAT processing.",
			},
		},
	}
}

func rnatsessionGetThePayloadFromthePlan(ctx context.Context, data *RnatsessionResourceModel) network.Rnatsession {
	tflog.Debug(ctx, "In rnatsessionGetThePayloadFromthePlan Function")

	// Create API request body from the model
	rnatsession := network.Rnatsession{}
	if !data.Aclname.IsNull() && !data.Aclname.IsUnknown() {
		rnatsession.Aclname = data.Aclname.ValueString()
	}
	if !data.Natip.IsNull() && !data.Natip.IsUnknown() {
		rnatsession.Natip = data.Natip.ValueString()
	}
	if !data.Netmask.IsNull() && !data.Netmask.IsUnknown() {
		rnatsession.Netmask = data.Netmask.ValueString()
	}
	if !data.Network.IsNull() && !data.Network.IsUnknown() {
		rnatsession.Network = data.Network.ValueString()
	}

	return rnatsession
}
