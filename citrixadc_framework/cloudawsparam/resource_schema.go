package cloudawsparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cloud"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

)

// CloudawsparamResourceModel describes the resource data model.
type CloudawsparamResourceModel struct {
	Id types.String `tfsdk:"id"`
	Rolearn types.String `tfsdk:"rolearn"`
}

func (r *CloudawsparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cloudawsparam resource.",
			},
			"rolearn": schema.StringAttribute{
				Required:    true,
				Description: "IAM Role ARN",
			},
		},
	}
}

func cloudawsparamGetThePayloadFromthePlan(ctx context.Context, data *CloudawsparamResourceModel) cloud.Cloudawsparam {
	tflog.Debug(ctx, "In cloudawsparamGetThePayloadFromthePlan Function")

	// Create API request body from the model
	cloudawsparam := cloud.Cloudawsparam{}
	if !data.Rolearn.IsNull() && !data.Rolearn.IsUnknown() {
		cloudawsparam.Rolearn = data.Rolearn.ValueString()
	}

	return cloudawsparam
}

func cloudawsparamSetAttrFromGet(ctx context.Context, data *CloudawsparamResourceModel, getResponseData map[string]interface{}) *CloudawsparamResourceModel {
	tflog.Debug(ctx, "In cloudawsparamSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["rolearn"]; ok && val != nil {
		data.Rolearn = types.StringValue(val.(string))
	} else {
		data.Rolearn = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("cloudawsparam-config")

	return data
}