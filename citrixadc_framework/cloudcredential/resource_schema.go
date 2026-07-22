package cloudcredential

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cloud"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CloudcredentialResourceModel describes the resource data model.
type CloudcredentialResourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Applicationid              types.String `tfsdk:"applicationid"`
	Applicationsecret          types.String `tfsdk:"applicationsecret"`
	ApplicationsecretWo        types.String `tfsdk:"applicationsecret_wo"`
	ApplicationsecretWoVersion types.Int64  `tfsdk:"applicationsecret_wo_version"`
	Tenantidentifier           types.String `tfsdk:"tenantidentifier"`
}

func (r *CloudcredentialResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cloudcredential resource.",
			},
			"applicationid": schema.StringAttribute{
				Required:    true,
				Description: "Application ID of the Credentials",
			},
			"applicationsecret": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Application Secret of the Credentials.",
			},
			"applicationsecret_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "Application Secret of the Credentials.",
			},
			"applicationsecret_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a applicationsecret_wo update.",
			},
			"tenantidentifier": schema.StringAttribute{
				Required:    true,
				Description: "Tenant ID of the Credentials",
			},
		},
	}
}

func cloudcredentialGetThePayloadFromthePlan(ctx context.Context, data *CloudcredentialResourceModel) cloud.Cloudcredential {
	tflog.Debug(ctx, "In cloudcredentialGetThePayloadFromthePlan Function")

	// Create API request body from the model
	cloudcredential := cloud.Cloudcredential{}
	if !data.Applicationid.IsNull() && !data.Applicationid.IsUnknown() {
		cloudcredential.Applicationid = data.Applicationid.ValueString()
	}
	if !data.Applicationsecret.IsNull() && !data.Applicationsecret.IsUnknown() {
		cloudcredential.Applicationsecret = data.Applicationsecret.ValueString()
	}
	// Skip write-only attribute: applicationsecret_wo
	// Skip version tracker attribute: applicationsecret_wo_version
	if !data.Tenantidentifier.IsNull() && !data.Tenantidentifier.IsUnknown() {
		cloudcredential.Tenantidentifier = data.Tenantidentifier.ValueString()
	}

	return cloudcredential
}

func cloudcredentialGetThePayloadFromtheConfig(ctx context.Context, data *CloudcredentialResourceModel, payload *cloud.Cloudcredential) {
	tflog.Debug(ctx, "In cloudcredentialGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: applicationsecret_wo -> applicationsecret
	if !data.ApplicationsecretWo.IsNull() {
		applicationsecretWo := data.ApplicationsecretWo.ValueString()
		if applicationsecretWo != "" {
			payload.Applicationsecret = applicationsecretWo
		}
	}
}

func cloudcredentialSetAttrFromGet(ctx context.Context, data *CloudcredentialResourceModel, getResponseData map[string]interface{}) *CloudcredentialResourceModel {
	tflog.Debug(ctx, "In cloudcredentialSetAttrFromGet Function")

	// NITRO GET for cloudcredential returns ONLY the read-only "isset" flag
	// (and "_nextgenapiresource"); it does NOT echo back applicationid,
	// tenantidentifier, or applicationsecret. Therefore we must NOT overwrite
	// these three from the GET response - preserve the values already present
	// in the plan/state, otherwise they would be clobbered to null on every
	// Read, causing "inconsistent result after apply" and perpetual diffs.
	// applicationid       - retain from plan/state (not returned by GET)
	// tenantidentifier    - retain from plan/state (not returned by GET)
	// applicationsecret   - retain from plan/state (secret, never returned by GET)
	// applicationsecret_wo / applicationsecret_wo_version - retain from plan/state

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("cloudcredential-config")

	return data
}
