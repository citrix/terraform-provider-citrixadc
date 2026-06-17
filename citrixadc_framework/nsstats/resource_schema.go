package nsstats

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NsstatsResourceModel describes the resource data model.
type NsstatsResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Cleanuplevel types.String `tfsdk:"cleanuplevel"`
}

func (r *NsstatsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsstats resource.",
			},
			"cleanuplevel": schema.StringAttribute{
				Required:    true,
				Description: "The level of stats to be cleared. 'global' option will clear global counters only, 'all' option will clear all device counters also along with global counters. For both the cases only 'ever incrementing counters' i.e. total counters will be cleared.\nPossible values = global, all",
			},
		},
	}
}

// nsstatsGetThePayloadFromthePlan builds the action payload, including only the set args.
func nsstatsGetThePayloadFromthePlan(ctx context.Context, data *NsstatsResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In nsstatsGetThePayloadFromthePlan Function")

	payload := map[string]interface{}{}
	if !data.Cleanuplevel.IsNull() && !data.Cleanuplevel.IsUnknown() {
		payload["cleanuplevel"] = data.Cleanuplevel.ValueString()
	}

	return payload
}
