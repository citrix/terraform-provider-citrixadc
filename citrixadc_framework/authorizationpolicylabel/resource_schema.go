package authorizationpolicylabel

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authorization"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthorizationpolicylabelResourceModel describes the resource data model.
type AuthorizationpolicylabelResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Labelname types.String `tfsdk:"labelname"`
	Newname   types.String `tfsdk:"newname"`
}

func (r *AuthorizationpolicylabelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authorizationpolicylabel resource.",
			},
			"labelname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the new authorization policy label. \nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the authorization policy is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authorization policy label\" or 'authorization policy label').",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The new name of the auth policy label",
			},
		},
	}
}

func authorizationpolicylabelGetThePayloadFromtheConfig(ctx context.Context, data *AuthorizationpolicylabelResourceModel) authorization.Authorizationpolicylabel {
	tflog.Debug(ctx, "In authorizationpolicylabelGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	authorizationpolicylabel := authorization.Authorizationpolicylabel{}
	if !data.Labelname.IsNull() {
		authorizationpolicylabel.Labelname = data.Labelname.ValueString()
	}
	if !data.Newname.IsNull() {
		authorizationpolicylabel.Newname = data.Newname.ValueString()
	}

	return authorizationpolicylabel
}

func authorizationpolicylabelSetAttrFromGet(ctx context.Context, data *AuthorizationpolicylabelResourceModel, getResponseData map[string]interface{}) *AuthorizationpolicylabelResourceModel {
	tflog.Debug(ctx, "In authorizationpolicylabelSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["labelname"]; ok && val != nil {
		data.Labelname = types.StringValue(val.(string))
	} else {
		data.Labelname = types.StringNull()
	}
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Labelname.ValueString())

	return data
}
