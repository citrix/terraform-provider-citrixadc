package nslimitselector

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NslimitselectorResourceModel describes the resource data model.
type NslimitselectorResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Rule         types.List   `tfsdk:"rule"`
	Selectorname types.String `tfsdk:"selectorname"`
}

func (r *NslimitselectorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nslimitselector resource.",
			},
			"rule": schema.ListAttribute{
				ElementType: types.StringType,
				Required:    true,
				Description: "0",
			},
			"selectorname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "0",
			},
		},
	}
}

func nslimitselectorGetThePayloadFromthePlan(ctx context.Context, data *NslimitselectorResourceModel) ns.Nslimitselector {
	tflog.Debug(ctx, "In nslimitselectorGetThePayloadFromthePlan Function")

	// Create API request body from the model
	nslimitselector := ns.Nslimitselector{}
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		var ruleList []string
		data.Rule.ElementsAs(ctx, &ruleList, false)
		nslimitselector.Rule = ruleList
	}
	if !data.Selectorname.IsNull() && !data.Selectorname.IsUnknown() {
		nslimitselector.Selectorname = data.Selectorname.ValueString()
	}

	return nslimitselector
}

func nslimitselectorSetAttrFromGet(ctx context.Context, data *NslimitselectorResourceModel, getResponseData map[string]interface{}) *NslimitselectorResourceModel {
	tflog.Debug(ctx, "In nslimitselectorSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["rule"]; ok && val != nil {
		if sliceVal, ok := val.([]interface{}); ok {
			stringList := utils.ToStringList(sliceVal)
			listValue, _ := types.ListValueFrom(ctx, types.StringType, stringList)
			data.Rule = listValue
		} else {
			data.Rule = types.ListNull(types.StringType)
		}
	} else {
		data.Rule = types.ListNull(types.StringType)
	}
	// NITRO returns this object under the "streamselector" key with the field
	// named "name" rather than "selectorname". Accept both for robustness.
	if val, ok := getResponseData["selectorname"]; ok && val != nil {
		data.Selectorname = types.StringValue(val.(string))
	} else if val, ok := getResponseData["name"]; ok && val != nil {
		data.Selectorname = types.StringValue(val.(string))
	} else {
		data.Selectorname = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Selectorname.ValueString()))

	return data
}
