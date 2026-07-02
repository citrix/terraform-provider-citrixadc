package dnssubnetcache

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DnssubnetcacheResourceModel describes the resource data model.
// dnssubnetcache is an action-only (flush) resource with no persistent object.
// nodeid is a NITRO GET-only cluster filter (Pattern 15) and is intentionally not
// modelled: it is never part of the flush write payload.
type DnssubnetcacheResourceModel struct {
	Id        types.String `tfsdk:"id"`
	All       types.Bool   `tfsdk:"all"`
	Ecssubnet types.String `tfsdk:"ecssubnet"`
}

func (r *DnssubnetcacheResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dnssubnetcache resource.",
			},
			"all": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Flush all the ECS subnets from the DNS cache.",
			},
			"ecssubnet": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "ECS Subnet.",
			},
		},
	}
}

func dnssubnetcacheGetThePayloadFromthePlan(ctx context.Context, data *DnssubnetcacheResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In dnssubnetcacheGetThePayloadFromthePlan Function")

	// Build the flush action payload. Only ecssubnet/all are valid flush inputs;
	// nodeid is a GET-only filter and must never be sent.
	dnssubnetcache := make(map[string]interface{})
	if !data.All.IsNull() && !data.All.IsUnknown() {
		dnssubnetcache["all"] = data.All.ValueBool()
	}
	if !data.Ecssubnet.IsNull() && !data.Ecssubnet.IsUnknown() {
		dnssubnetcache["ecssubnet"] = data.Ecssubnet.ValueString()
	}

	return dnssubnetcache
}
