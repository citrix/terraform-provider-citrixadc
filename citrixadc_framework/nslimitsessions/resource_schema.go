package nslimitsessions

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NslimitsessionsResourceModel describes the resource data model.
// nslimitsessions is an action-only resource (clear); the model is kept
// minimal: a synthetic id plus the clear payload key, limitidentifier.
// detail is a GET-only filter argument and lives only in the datasource.
type NslimitsessionsResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Limitidentifier types.String `tfsdk:"limitidentifier"`
}

func (r *NslimitsessionsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nslimitsessions resource.",
			},
			"limitidentifier": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the rate limit identifier for which to display the sessions.",
			},
		},
	}
}

func nslimitsessionsGetThePayloadFromthePlan(ctx context.Context, data *NslimitsessionsResourceModel) ns.Nslimitsessions {
	tflog.Debug(ctx, "In nslimitsessionsGetThePayloadFromthePlan Function")

	// Create API request body from the model. The clear action accepts only
	// limitidentifier (detail is a GET-only filter, excluded here).
	nslimitsessions := ns.Nslimitsessions{}
	if !data.Limitidentifier.IsNull() && !data.Limitidentifier.IsUnknown() {
		nslimitsessions.Limitidentifier = data.Limitidentifier.ValueString()
	}

	return nslimitsessions
}

// nslimitsessionsSetAttrFromGetForDatasource faithfully copies the GET response
// into the datasource model (the datasource has its own model so the resource
// model can stay minimal).
func nslimitsessionsSetAttrFromGetForDatasource(ctx context.Context, data *NslimitsessionsDataSourceModel, getResponseData map[string]interface{}) *NslimitsessionsDataSourceModel {
	tflog.Debug(ctx, "In nslimitsessionsSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["limitidentifier"]; ok && val != nil {
		data.Limitidentifier = types.StringValue(val.(string))
	} else {
		data.Limitidentifier = types.StringNull()
	}
	if val, ok := getResponseData["detail"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.Detail = types.BoolValue(b)
		}
	}
	if val, ok := getResponseData["timeout"]; ok && val != nil {
		data.Timeout = types.StringValue(utils.ToString(val))
	} else {
		data.Timeout = types.StringNull()
	}
	if val, ok := getResponseData["hits"]; ok && val != nil {
		data.Hits = types.StringValue(utils.ToString(val))
	} else {
		data.Hits = types.StringNull()
	}
	if val, ok := getResponseData["drop"]; ok && val != nil {
		data.Drop = types.StringValue(utils.ToString(val))
	} else {
		data.Drop = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(utils.ToString(val))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["unit"]; ok && val != nil {
		data.Unit = types.StringValue(utils.ToString(val))
	} else {
		data.Unit = types.StringNull()
	}

	// Datasource has no Create; set its ID from the lookup key.
	data.Id = types.StringValue(data.Limitidentifier.ValueString())

	return data
}
