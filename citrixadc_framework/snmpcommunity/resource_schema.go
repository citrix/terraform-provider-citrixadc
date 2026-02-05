package snmpcommunity

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/snmp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SnmpcommunityResourceModel describes the resource data model.
type SnmpcommunityResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Communityname types.String `tfsdk:"communityname"`
	Permissions   types.String `tfsdk:"permissions"`
}

func (r *SnmpcommunityResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the snmpcommunity resource.",
			},
			"communityname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The SNMP community string. Can consist of 1 to 31 characters that include uppercase and lowercase letters,numbers and special characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the string includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my string\" or 'my string').",
			},
			"permissions": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The SNMP V1 or V2 query-type privilege that you want to associate with this SNMP community.",
			},
		},
	}
}

func snmpcommunityGetThePayloadFromtheConfig(ctx context.Context, data *SnmpcommunityResourceModel) snmp.Snmpcommunity {
	tflog.Debug(ctx, "In snmpcommunityGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	snmpcommunity := snmp.Snmpcommunity{}
	if !data.Communityname.IsNull() {
		snmpcommunity.Communityname = data.Communityname.ValueString()
	}
	if !data.Permissions.IsNull() {
		snmpcommunity.Permissions = data.Permissions.ValueString()
	}

	return snmpcommunity
}

func snmpcommunitySetAttrFromGet(ctx context.Context, data *SnmpcommunityResourceModel, getResponseData map[string]interface{}) *SnmpcommunityResourceModel {
	tflog.Debug(ctx, "In snmpcommunitySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["communityname"]; ok && val != nil {
		data.Communityname = types.StringValue(val.(string))
	} else {
		data.Communityname = types.StringNull()
	}
	if val, ok := getResponseData["permissions"]; ok && val != nil {
		data.Permissions = types.StringValue(val.(string))
	} else {
		data.Permissions = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Communityname.ValueString())

	return data
}
