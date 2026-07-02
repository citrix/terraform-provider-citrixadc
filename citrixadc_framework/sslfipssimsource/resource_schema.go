package sslfipssimsource

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslfipssimsourceResourceModel describes the resource data model.
type SslfipssimsourceResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Certfile     types.String `tfsdk:"certfile"`
	Sourcesecret types.String `tfsdk:"sourcesecret"`
	Targetsecret types.String `tfsdk:"targetsecret"`
}

func (r *SslfipssimsourceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		// WARNING: DISRUPTIVE and FIPS-only. sslfipssimsource exposes only the
		// `enable` and `init` NITRO actions (no get/add/delete). It requires
		// dedicated FIPS hardware and is unsupported on non-FIPS VPX appliances.
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslfipssimsource resource.",
			},
			"certfile": schema.StringAttribute{
				// Required for the init action (Pattern 8: tfdata wrongly marks optional).
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for and, optionally, path to the source FIPS appliance's certificate file. /nsconfig/ssl/ is the default path.",
			},
			"sourcesecret": schema.StringAttribute{
				// Required for the enable action (Pattern 8).
				Required:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for and, optionally, path to the source FIPS appliance's secret data. /nsconfig/ssl/ is the default path.",
			},
			"targetsecret": schema.StringAttribute{
				// Required for the enable action (Pattern 8).
				Required:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of and, optionally, path to the target FIPS appliance's secret data. /nsconfig/ssl/ is the default path.",
			},
		},
	}
}

func sslfipssimsourceGetThePayloadFromthePlan(ctx context.Context, data *SslfipssimsourceResourceModel) ssl.Sslfipssimsource {
	tflog.Debug(ctx, "In sslfipssimsourceGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslfipssimsource := ssl.Sslfipssimsource{}
	if !data.Certfile.IsNull() && !data.Certfile.IsUnknown() {
		sslfipssimsource.Certfile = data.Certfile.ValueString()
	}
	if !data.Sourcesecret.IsNull() && !data.Sourcesecret.IsUnknown() {
		sslfipssimsource.Sourcesecret = data.Sourcesecret.ValueString()
	}
	if !data.Targetsecret.IsNull() && !data.Targetsecret.IsUnknown() {
		sslfipssimsource.Targetsecret = data.Targetsecret.ValueString()
	}

	return sslfipssimsource
}
