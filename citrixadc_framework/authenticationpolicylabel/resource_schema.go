package authenticationpolicylabel

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationpolicylabelResourceModel describes the resource data model.
type AuthenticationpolicylabelResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Comment     types.String `tfsdk:"comment"`
	Labelname   types.String `tfsdk:"labelname"`
	Loginschema types.String `tfsdk:"loginschema"`
	Newname     types.String `tfsdk:"newname"`
	Type        types.String `tfsdk:"type"`
}

func (r *AuthenticationpolicylabelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationpolicylabel resource.",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments to preserve information about this authentication policy label.",
			},
			"labelname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new authentication policy label.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication policy label\" or 'authentication policy label').",
			},
			"loginschema": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Login schema associated with authentication policy label. Login schema defines the UI rendering by providing customization option of the fields. If user intervention is not needed for a given factor such as group extraction, a loginSchema whose authentication schema is \"noschema\" should be used.",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The new name of the auth policy label",
			},
			"type": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("AAATM_REQ"),
				Description: "Type of feature (aaatm or rba) against which to match the policies bound to this policy label.",
			},
		},
	}
}

func authenticationpolicylabelGetThePayloadFromtheConfig(ctx context.Context, data *AuthenticationpolicylabelResourceModel) authentication.Authenticationpolicylabel {
	tflog.Debug(ctx, "In authenticationpolicylabelGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	authenticationpolicylabel := authentication.Authenticationpolicylabel{}
	if !data.Comment.IsNull() {
		authenticationpolicylabel.Comment = data.Comment.ValueString()
	}
	if !data.Labelname.IsNull() {
		authenticationpolicylabel.Labelname = data.Labelname.ValueString()
	}
	if !data.Loginschema.IsNull() {
		authenticationpolicylabel.Loginschema = data.Loginschema.ValueString()
	}
	if !data.Newname.IsNull() {
		authenticationpolicylabel.Newname = data.Newname.ValueString()
	}
	if !data.Type.IsNull() {
		authenticationpolicylabel.Type = data.Type.ValueString()
	}

	return authenticationpolicylabel
}

func authenticationpolicylabelSetAttrFromGet(ctx context.Context, data *AuthenticationpolicylabelResourceModel, getResponseData map[string]interface{}) *AuthenticationpolicylabelResourceModel {
	tflog.Debug(ctx, "In authenticationpolicylabelSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["labelname"]; ok && val != nil {
		data.Labelname = types.StringValue(val.(string))
	} else {
		data.Labelname = types.StringNull()
	}
	if val, ok := getResponseData["loginschema"]; ok && val != nil {
		data.Loginschema = types.StringValue(val.(string))
	} else {
		data.Loginschema = types.StringNull()
	}
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Labelname.ValueString())

	return data
}
