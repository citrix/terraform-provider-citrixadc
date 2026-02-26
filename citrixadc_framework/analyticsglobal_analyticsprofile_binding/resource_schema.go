package analyticsglobal_analyticsprofile_binding

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/analytics"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AnalyticsglobalAnalyticsprofileBindingResourceModel describes the resource data model.
type AnalyticsglobalAnalyticsprofileBindingResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Analyticsprofile types.String `tfsdk:"analyticsprofile"`
}

func (r *AnalyticsglobalAnalyticsprofileBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the analyticsglobal_analyticsprofile_binding resource.",
			},
			"analyticsprofile": schema.StringAttribute{
				Required:    true,
				Description: "Name of the analytics profile bound.",
			},
		},
	}
}

func analyticsglobal_analyticsprofile_bindingGetThePayloadFromtheConfig(ctx context.Context, data *AnalyticsglobalAnalyticsprofileBindingResourceModel) analytics.Analyticsglobalanalyticsprofilebinding {
	tflog.Debug(ctx, "In analyticsglobal_analyticsprofile_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	analyticsglobal_analyticsprofile_binding := analytics.Analyticsglobalanalyticsprofilebinding{}
	if !data.Analyticsprofile.IsNull() {
		analyticsglobal_analyticsprofile_binding.Analyticsprofile = data.Analyticsprofile.ValueString()
	}

	return analyticsglobal_analyticsprofile_binding
}

func analyticsglobal_analyticsprofile_bindingSetAttrFromGet(ctx context.Context, data *AnalyticsglobalAnalyticsprofileBindingResourceModel, getResponseData map[string]interface{}) *AnalyticsglobalAnalyticsprofileBindingResourceModel {
	tflog.Debug(ctx, "In analyticsglobal_analyticsprofile_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["analyticsprofile"]; ok && val != nil {
		data.Analyticsprofile = types.StringValue(val.(string))
	} else {
		data.Analyticsprofile = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Analyticsprofile.ValueString())

	return data
}
