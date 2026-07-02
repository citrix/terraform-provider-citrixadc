package cloudparaminternal

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cloud"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CloudparaminternalResourceModel describes the resource data model.
type CloudparaminternalResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Nonftumode types.String `tfsdk:"nonftumode"`
}

func (r *CloudparaminternalResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cloudparaminternal resource.",
			},
			"nonftumode": schema.StringAttribute{
				Optional:    true,
				Description: "Indicates if GUI in in FTU mode or not",
			},
		},
	}
}

func cloudparaminternalGetThePayloadFromthePlan(ctx context.Context, data *CloudparaminternalResourceModel) cloud.Cloudparaminternal {
	tflog.Debug(ctx, "In cloudparaminternalGetThePayloadFromthePlan Function")

	// Create API request body from the model
	cloudparaminternal := cloud.Cloudparaminternal{}
	if !data.Nonftumode.IsNull() && !data.Nonftumode.IsUnknown() {
		cloudparaminternal.Nonftumode = data.Nonftumode.ValueString()
	}

	return cloudparaminternal
}

// cloudparaminternalSetAttrFromGet is the RESOURCE state setter.
//
// cloudparaminternal is a singleton SET-GET resource. nonftumode is Optional-only
// (not Computed) in the schema, so Terraform requires the post-apply state to match
// config exactly. GET/show on some platforms returns "Operation not supported on
// this platform" (handled gracefully in Read) and on others may not echo the value
// back verbatim. Adopting the server view could therefore clobber the user's value
// to null and produce a perpetual diff / "inconsistent result after apply" (Pattern 7).
//
// Therefore the resource setter PRESERVES the plan/state value for nonftumode and
// only (re)affirms the static ID. The datasource, which has no prior state to
// preserve, uses the separate faithful setter below.
func cloudparaminternalSetAttrFromGet(ctx context.Context, data *CloudparaminternalResourceModel, getResponseData map[string]interface{}) *CloudparaminternalResourceModel {
	tflog.Debug(ctx, "In cloudparaminternalSetAttrFromGet Function (preserving plan/state values)")

	// Do not overwrite nonftumode from the GET response: preserve the existing
	// model value verbatim (write-back safe on platforms where GET is gated).

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("cloudparaminternal-config")

	return data
}

// cloudparaminternalSetAttrFromGetForDatasource is the DATASOURCE state setter.
//
// A datasource has no prior plan/state to preserve, so it faithfully copies the
// attribute the GET response returns.
func cloudparaminternalSetAttrFromGetForDatasource(ctx context.Context, data *CloudparaminternalResourceModel, getResponseData map[string]interface{}) *CloudparaminternalResourceModel {
	tflog.Debug(ctx, "In cloudparaminternalSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["nonftumode"]; ok && val != nil {
		data.Nonftumode = types.StringValue(val.(string))
	} else {
		data.Nonftumode = types.StringNull()
	}

	// Set ID for the datasource (no Create step to set it)
	data.Id = types.StringValue("cloudparaminternal-config")

	return data
}
