package streamselector

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/stream"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// StreamselectorResourceModel describes the resource data model.
type StreamselectorResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Rule types.List   `tfsdk:"rule"`
}

func (r *StreamselectorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the streamselector resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the selector. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. If the name includes one or more spaces, and you are using the Citrix ADC CLI, enclose the name in double or single quotation marks (for example, \"my selector\" or 'my selector').",
			},
			"rule": schema.ListAttribute{
				ElementType: types.StringType,
				Required:    true,
				Description: "Set of up to five expressions. Maximum length: 7499 characters. Each expression must identify a specific request characteristic, such as the client's IP address (with CLIENT.IP.SRC) or requested server resource (with HTTP.REQ.URL).\nNote: If two or more selectors contain the same expressions in different order, a separate set of records is created for each selector.",
			},
		},
	}
}

func streamselectorGetThePayloadFromthePlan(ctx context.Context, data *StreamselectorResourceModel) stream.Streamselector {
	tflog.Debug(ctx, "In streamselectorGetThePayloadFromthePlan Function")

	// Create API request body from the model
	streamselector := stream.Streamselector{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		streamselector.Name = data.Name.ValueString()
	}
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		var ruleList []string
		data.Rule.ElementsAs(ctx, &ruleList, false)
		streamselector.Rule = ruleList
	}

	return streamselector
}

func streamselectorSetAttrFromGet(ctx context.Context, data *StreamselectorResourceModel, getResponseData map[string]interface{}) *StreamselectorResourceModel {
	tflog.Debug(ctx, "In streamselectorSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["rule"]; ok && val != nil {
		switch v := val.(type) {
		case []interface{}:
			stringList := utils.ToStringList(v)
			listValue, _ := types.ListValueFrom(ctx, types.StringType, stringList)
			data.Rule = listValue
		case string:
			listValue, _ := types.ListValueFrom(ctx, types.StringType, []string{v})
			data.Rule = listValue
		default:
			data.Rule = types.ListNull(types.StringType)
		}
	} else {
		data.Rule = types.ListNull(types.StringType)
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
