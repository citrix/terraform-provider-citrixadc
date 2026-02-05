package locationparameter

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/basic"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LocationparameterResourceModel describes the resource data model.
type LocationparameterResourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Context            types.String `tfsdk:"context"`
	Matchwildcardtoany types.String `tfsdk:"matchwildcardtoany"`
	Q1label            types.String `tfsdk:"q1label"`
	Q2label            types.String `tfsdk:"q2label"`
	Q3label            types.String `tfsdk:"q3label"`
	Q4label            types.String `tfsdk:"q4label"`
	Q5label            types.String `tfsdk:"q5label"`
	Q6label            types.String `tfsdk:"q6label"`
}

func (r *LocationparameterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the locationparameter resource.",
			},
			"context": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Context for describing locations. In geographic context, qualifier labels are assigned by default in the following sequence: Continent.Country.Region.City.ISP.Organization. In custom context, the qualifiers labels can have any meaning that you designate.",
			},
			"matchwildcardtoany": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Indicates whether wildcard qualifiers should match any other\nqualifier including non-wildcard while evaluating\nlocation based expressions.\nPossible values: Yes, No, Expression.\n    Yes - Wildcard qualifiers match any other qualifiers.\n    No  - Wildcard qualifiers do not match non-wildcard\n          qualifiers, but match other wildcard qualifiers.\n    Expression - Wildcard qualifiers in an expression\n          match any qualifier in an LDNS location,\n          wildcard qualifiers in the LDNS location do not match\n          non-wildcard qualifiers in an expression",
			},
			"q1label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Label specifying the meaning of the first qualifier. Can be specified for custom context only.",
			},
			"q2label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Label specifying the meaning of the second qualifier. Can be specified for custom context only.",
			},
			"q3label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Label specifying the meaning of the third qualifier. Can be specified for custom context only.",
			},
			"q4label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Label specifying the meaning of the fourth qualifier. Can be specified for custom context only.",
			},
			"q5label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Label specifying the meaning of the fifth qualifier. Can be specified for custom context only.",
			},
			"q6label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Label specifying the meaning of the sixth qualifier. Can be specified for custom context only.",
			},
		},
	}
}

func locationparameterGetThePayloadFromtheConfig(ctx context.Context, data *LocationparameterResourceModel) basic.Locationparameter {
	tflog.Debug(ctx, "In locationparameterGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	locationparameter := basic.Locationparameter{}
	if !data.Context.IsNull() {
		locationparameter.Context = data.Context.ValueString()
	}
	if !data.Matchwildcardtoany.IsNull() {
		locationparameter.Matchwildcardtoany = data.Matchwildcardtoany.ValueString()
	}
	if !data.Q1label.IsNull() {
		locationparameter.Q1label = data.Q1label.ValueString()
	}
	if !data.Q2label.IsNull() {
		locationparameter.Q2label = data.Q2label.ValueString()
	}
	if !data.Q3label.IsNull() {
		locationparameter.Q3label = data.Q3label.ValueString()
	}
	if !data.Q4label.IsNull() {
		locationparameter.Q4label = data.Q4label.ValueString()
	}
	if !data.Q5label.IsNull() {
		locationparameter.Q5label = data.Q5label.ValueString()
	}
	if !data.Q6label.IsNull() {
		locationparameter.Q6label = data.Q6label.ValueString()
	}

	return locationparameter
}

func locationparameterSetAttrFromGet(ctx context.Context, data *LocationparameterResourceModel, getResponseData map[string]interface{}) *LocationparameterResourceModel {
	tflog.Debug(ctx, "In locationparameterSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["context"]; ok && val != nil {
		data.Context = types.StringValue(val.(string))
	} else {
		data.Context = types.StringNull()
	}
	if val, ok := getResponseData["matchwildcardtoany"]; ok && val != nil {
		data.Matchwildcardtoany = types.StringValue(val.(string))
	} else {
		data.Matchwildcardtoany = types.StringNull()
	}
	if val, ok := getResponseData["q1label"]; ok && val != nil {
		data.Q1label = types.StringValue(val.(string))
	} else {
		data.Q1label = types.StringNull()
	}
	if val, ok := getResponseData["q2label"]; ok && val != nil {
		data.Q2label = types.StringValue(val.(string))
	} else {
		data.Q2label = types.StringNull()
	}
	if val, ok := getResponseData["q3label"]; ok && val != nil {
		data.Q3label = types.StringValue(val.(string))
	} else {
		data.Q3label = types.StringNull()
	}
	if val, ok := getResponseData["q4label"]; ok && val != nil {
		data.Q4label = types.StringValue(val.(string))
	} else {
		data.Q4label = types.StringNull()
	}
	if val, ok := getResponseData["q5label"]; ok && val != nil {
		data.Q5label = types.StringValue(val.(string))
	} else {
		data.Q5label = types.StringNull()
	}
	if val, ok := getResponseData["q6label"]; ok && val != nil {
		data.Q6label = types.StringValue(val.(string))
	} else {
		data.Q6label = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("locationparameter-config")

	return data
}
