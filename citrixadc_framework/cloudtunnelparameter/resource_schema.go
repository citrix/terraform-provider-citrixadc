package cloudtunnelparameter

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cloudtunnel"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CloudtunnelparameterResourceModel describes the resource data model.
type CloudtunnelparameterResourceModel struct {
	Id                             types.String `tfsdk:"id"`
	Controllerfqdn                 types.String `tfsdk:"controllerfqdn"`
	Fqdn                           types.String `tfsdk:"fqdn"`
	Resourcelocation               types.String `tfsdk:"resourcelocation"`
	Subnetresourcelocationmappings types.String `tfsdk:"subnetresourcelocationmappings"`
}

func (r *CloudtunnelparameterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cloudtunnelparameter resource.",
			},
			"controllerfqdn": schema.StringAttribute{
				Optional:    true,
				Description: "0",
			},
			"fqdn": schema.StringAttribute{
				Optional:    true,
				Description: "0",
			},
			"resourcelocation": schema.StringAttribute{
				Optional:    true,
				Description: "0",
			},
			"subnetresourcelocationmappings": schema.StringAttribute{
				Optional:    true,
				Description: "0",
			},
		},
	}
}

func cloudtunnelparameterGetThePayloadFromthePlan(ctx context.Context, data *CloudtunnelparameterResourceModel) cloudtunnel.Cloudtunnelparameter {
	tflog.Debug(ctx, "In cloudtunnelparameterGetThePayloadFromthePlan Function")

	// Create API request body from the model
	cloudtunnelparameter := cloudtunnel.Cloudtunnelparameter{}
	if !data.Controllerfqdn.IsNull() && !data.Controllerfqdn.IsUnknown() {
		cloudtunnelparameter.Controllerfqdn = data.Controllerfqdn.ValueString()
	}
	if !data.Fqdn.IsNull() && !data.Fqdn.IsUnknown() {
		cloudtunnelparameter.Fqdn = data.Fqdn.ValueString()
	}
	if !data.Resourcelocation.IsNull() && !data.Resourcelocation.IsUnknown() {
		cloudtunnelparameter.Resourcelocation = data.Resourcelocation.ValueString()
	}
	if !data.Subnetresourcelocationmappings.IsNull() && !data.Subnetresourcelocationmappings.IsUnknown() {
		cloudtunnelparameter.Subnetresourcelocationmappings = data.Subnetresourcelocationmappings.ValueString()
	}

	return cloudtunnelparameter
}

// cloudtunnelparameterSetAttrFromGet is the RESOURCE state setter.
//
// cloudtunnelparameter is a singleton SET-GET resource whose 4 attributes are all
// Optional-only (not Computed), so Terraform requires the post-apply state to match
// config exactly. The feature is platform-gated: GET/show may return
// "Feature not supported in this release" (handled gracefully in Read) and, even when
// available, may not echo every value back verbatim. Adopting the server view could
// therefore clobber the user's configured value to null and produce a perpetual diff
// / "inconsistent result after apply" (Pattern 7).
//
// Therefore the resource setter PRESERVES the plan/state values and only (re)affirms
// the static ID. The datasource, which has no prior state to preserve, uses the
// separate faithful setter below.
func cloudtunnelparameterSetAttrFromGet(ctx context.Context, data *CloudtunnelparameterResourceModel, getResponseData map[string]interface{}) *CloudtunnelparameterResourceModel {
	tflog.Debug(ctx, "In cloudtunnelparameterSetAttrFromGet Function (preserving plan/state values)")

	// Do not overwrite the settable attributes from the GET response: preserve the
	// existing model values verbatim (write-back safe on feature-gated platforms).

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("cloudtunnelparameter-config")

	return data
}

// cloudtunnelparameterSetAttrFromGetForDatasource is the DATASOURCE state setter.
//
// A datasource has no prior plan/state to preserve, so it faithfully copies the
// attributes the GET response returns.
func cloudtunnelparameterSetAttrFromGetForDatasource(ctx context.Context, data *CloudtunnelparameterResourceModel, getResponseData map[string]interface{}) *CloudtunnelparameterResourceModel {
	tflog.Debug(ctx, "In cloudtunnelparameterSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["controllerfqdn"]; ok && val != nil {
		data.Controllerfqdn = types.StringValue(val.(string))
	} else {
		data.Controllerfqdn = types.StringNull()
	}
	if val, ok := getResponseData["fqdn"]; ok && val != nil {
		data.Fqdn = types.StringValue(val.(string))
	} else {
		data.Fqdn = types.StringNull()
	}
	if val, ok := getResponseData["resourcelocation"]; ok && val != nil {
		data.Resourcelocation = types.StringValue(val.(string))
	} else {
		data.Resourcelocation = types.StringNull()
	}
	if val, ok := getResponseData["subnetresourcelocationmappings"]; ok && val != nil {
		data.Subnetresourcelocationmappings = types.StringValue(val.(string))
	} else {
		data.Subnetresourcelocationmappings = types.StringNull()
	}

	// Set ID for the datasource (no Create step to set it)
	data.Id = types.StringValue("cloudtunnelparameter-config")

	return data
}
