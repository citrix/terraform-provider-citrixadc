package sslfipssimtarget

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

// SslfipssimtargetResourceModel describes the resource data model.
type SslfipssimtargetResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Certfile     types.String `tfsdk:"certfile"`
	Keyvector    types.String `tfsdk:"keyvector"`
	Sourcesecret types.String `tfsdk:"sourcesecret"`
	Targetsecret types.String `tfsdk:"targetsecret"`
}

func (r *SslfipssimtargetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		// WARNING: DISRUPTIVE and FIPS-only. sslfipssimtarget exposes only the
		// `enable` and `init` NITRO actions (no get/add/delete). It requires
		// dedicated FIPS hardware and is unsupported on non-FIPS VPX appliances.
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslfipssimtarget resource.",
			},
			"certfile": schema.StringAttribute{
				// Required for the init action (Pattern 8: tfdata wrongly marks optional).
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of and, optionally, path to the source FIPS appliance's certificate file. /nsconfig/ssl/ is the default path.",
			},
			"keyvector": schema.StringAttribute{
				// Required for the enable action (Pattern 8).
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of and, optionally, path to the target FIPS appliance's key vector. /nsconfig/ssl/ is the default path.",
			},
			"sourcesecret": schema.StringAttribute{
				// Required for the enable action (Pattern 8).
				Required:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of and, optionally, path to the source FIPS appliance's secret data. /nsconfig/ssl/ is the default path.",
			},
			"targetsecret": schema.StringAttribute{
				// Required for the init action (Pattern 8).
				Required:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for and, optionally, path to the target FIPS appliance's secret data. The default input path for the secret data is /nsconfig/ssl/.",
			},
		},
	}
}

func sslfipssimtargetGetThePayloadFromthePlan(ctx context.Context, data *SslfipssimtargetResourceModel) ssl.Sslfipssimtarget {
	tflog.Debug(ctx, "In sslfipssimtargetGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslfipssimtarget := ssl.Sslfipssimtarget{}
	if !data.Certfile.IsNull() && !data.Certfile.IsUnknown() {
		sslfipssimtarget.Certfile = data.Certfile.ValueString()
	}
	if !data.Keyvector.IsNull() && !data.Keyvector.IsUnknown() {
		sslfipssimtarget.Keyvector = data.Keyvector.ValueString()
	}
	if !data.Sourcesecret.IsNull() && !data.Sourcesecret.IsUnknown() {
		sslfipssimtarget.Sourcesecret = data.Sourcesecret.ValueString()
	}
	if !data.Targetsecret.IsNull() && !data.Targetsecret.IsUnknown() {
		sslfipssimtarget.Targetsecret = data.Targetsecret.ValueString()
	}

	return sslfipssimtarget
}
